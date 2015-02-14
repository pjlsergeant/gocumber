package gocumber

import "fmt"
import "strings"
import "regexp"
import "strconv"

// USEFUL URLs:
//    https://github.com/bnoordhuis/ragel/blob/master/examples/go/url.rl
//    http://thingsaaronmade.com/blog/a-simple-intro-to-writing-a-lexer-with-ragel.html
//    https://raw.githubusercontent.com/cucumber/gherkin/master/ragel/lexer.java.rl.erb

  %%{
    machine lexer;
    alphtype byte;

    action begin_content {
      contentStart = p
      currentLine = lineNumber

      if(len(keyword) != 0) {
        startCol = p - lastNewline - (len(keyword) + 1);
      }
    }

    action start_docstring {
      currentLine = lineNumber
      startCol = p - lastNewline
    }

    action begin_docstring_content {
      contentStart = p
    }

    action start_docstring_content_type {
      docstringContentTypeStart = p
    }

    action end_docstring_content_type {
      docstringContentTypeEnd = p
    }

    action store_docstring_content {
      rawcon := data[contentStart:nextKeywordStart-1]

      con := unindent( startCol, rawcon )
      con = strings.Replace( con, "\\\"\\\"\\\"", "\"\"\"", -1 )
      con = strings.TrimLeft( con, " \t\r\n" )

      conType := string(data[docstringContentTypeStart:docstringContentTypeEnd])
      conType = strings.TrimSpace( conType )

      fmt.Printf("DocString Type:[%s]\nDocString Cont:[%s]\n", conType, con)
      // listener.docString(conType, con, currentLine);
    }

    action store_feature_content {
      kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
      name, description := nameAndUnindentedDescription( startCol, kcon );

      fmt.Printf("Item Typ: [Feature]\nItem Key: [%s]\nItem Nam: [%s]\nItem Des: [%s]\n", keyword, name, description)
      // listener.feature(keyword, nameDescription[0], nameDescription[1], currentLine);

      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_background_content {
      kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
      name, description := nameAndUnindentedDescription( startCol, kcon );

      fmt.Printf("Item Typ: [Background]\nItem Key: [%s]\nItem Nam: [%s]\nItem Des: [%s]\n", keyword, name, description)
      // listener.background(keyword, nameDescription[0], nameDescription[1], currentLine);
      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_scenario_content {
            kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
      name, description := nameAndUnindentedDescription( startCol, kcon );

      fmt.Printf("Item Typ: [Scenario]\nItem Key: [%s]\nItem Nam: [%s]\nItem Des: [%s]\n", keyword, name, description)
      // listener.scenario(keyword, nameDescription[0], nameDescription[1], currentLine);
      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_scenario_outline_content {
            kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
      name, description := nameAndUnindentedDescription( startCol, kcon );

      fmt.Printf("Item Typ: [Scenario Outline]\nItem Key: [%s]\nItem Nam: [%s]\nItem Des: [%s]\n", keyword, name, description)
      // listener.scenarioOutline(keyword, nameDescription[0], nameDescription[1], currentLine);
      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_examples_content {
            kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
      name, description := nameAndUnindentedDescription( startCol, kcon );

      fmt.Printf("Item Typ: [Examples]\nItem Key: [%s]\nItem Nam: [%s]\nItem Des: [%s]\n", keyword, name, description)
      // listener.examples(keyword, nameDescription[0], nameDescription[1], currentLine);
      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_step_content {
        con := string( data[contentStart:p] )
        con = strings.TrimSpace( con )
        fmt.Printf("Store Thing: [Step]\nStore Keyword: [%s]\nStore Content: [%s]\n", keyword, con )
      // listener.step(keyword, substring(data, contentStart, p).trim(), currentLine);
    }

    action store_comment_content {
        con := string( data[contentStart:p] )
        con = strings.TrimSpace( con )
        fmt.Printf("Store Thing: [Comment]\nStore Keyword: [%s]\nStore Content: [%s]\n", keyword, con )
      // listener.comment(substring(data, contentStart, p).trim(), lineNumber);
      keywordStart = -1
    }

    action store_tag_content {
              con := string( data[contentStart:p] )
        con = strings.TrimSpace( con )
        fmt.Printf("Store Thing: [Tag]\nStore Keyword: [%s]\nStore Content: [%s]\n", keyword, con )
      // listener.tag(substring(data, contentStart, p).trim(), currentLine);
      keywordStart = -1
    }

    action inc_line_number {
      lineNumber++;
    }

    action last_newline {
      lastNewline = p + 1
    }

    action start_keyword {
      if (keywordStart == -1) { keywordStart = p }
    }

    action end_keyword {
      rawKeyword := string(data[keywordStart:p])
      rawKeyword = strings.Replace(rawKeyword, ":", "", 1)
      keyword = rawKeyword[0:]
      keywordStart = -1
    }

    action next_keyword_start {
      nextKeywordStart = p
    }

    action start_row {
      p = p - 1
      currentRow = currentRow[:0]
      currentLine = lineNumber
    }

    action begin_cell_content {
      contentStart = p;
    }

    action store_cell_content {
      con := string(data[contentStart:p])
      con = strings.TrimSpace(con)
      con = strings.Replace(con, "\\|", "|", -1)
      con = strings.Replace(con, "\\n", "\n", -1)
      con = strings.Replace(con, "\\\\", "\\", -1)
      // fmt.Printf("store_cell_conent: [%s]\n", con)
      currentRow = append( currentRow, con )
    }

    action store_row {
      fmt.Printf("store_row: [%s] [%s]\n", currentRow, currentLine )
      // listener.row(currentRow, currentLine);
    }

    action end_feature {
      if(cs < lexer_first_final) {
          content := currentLineContent( data, lastNewline )
          panic(fmt.Sprintf("Lexing error on line %d: '%s'. See http://wiki.github.com/cucumber/gherkin/lexingerror for more information.", lineNumber, content))
      } else {
        fmt.Printf("EOF\n")
         // listener.eof();
      }
    }

    include lexer_common "lexer_common_english.rl";
  }%%

  // START: write data noerror;
  %% write data noerror;
  // END: write data noerror;

func currentLineContent(data []byte, lastNewline int) (string) {
    current := string(data[lastNewline:])
    return strings.TrimSpace( current )
}

func unindent(startCol int, text []byte) (string) {

    regex, err := regexp.Compile("(?m)^[\t ]{0," + strconv.Itoa(startCol) + "}")

    if ( err != nil ) {
      panic(err)
    }
    result := regex.ReplaceAll( text, text[:0] )
    return string(result)
}

func keywordContent( data []byte, p int, eof int, nextKeywordStart int, contentStart int ) ([]byte) {
    endPoint := nextKeywordStart
    if ( (nextKeywordStart == -1 || (p == eof)) ) {
        endPoint = p
    }
    con := data[contentStart:endPoint]
    return con
}


func nameAndUnindentedDescription(startCol int, textBytes []byte) (string, string) {
    text := unindent( startCol, textBytes )
    text = strings.TrimSpace( text )
    lines := strings.Split(text, "\n")

      for index,element := range lines {
        lines[index] = strings.TrimSpace(element)
      }

    if ( len(lines) == 0 ) {
      return "", ""
    } else if ( len(lines) == 1 ) {
      return lines[0],""
      } else {
        return lines[0], strings.Join(lines[1:], "\n")
      }
}

func ParseFeature(data []byte) (feature Feature, err error) {

  // Original ragel parser assumes this will be there, who am I to argue?
  data = append(data, []byte("%_FEATURE_END_%")...)

  cs := 0 // No idea what this is
  p := 0 // Position?
  pe := len(data) // No idea
  eof := len(data) // Location of EOF

  lineNumber := 1
  lastNewline := 0

  contentStart := -1
  currentLine := -1

  docstringContentTypeStart := -1
  docstringContentTypeEnd := -1
  startCol := -1;
  nextKeywordStart := -1
  keywordStart := -1

  var keyword string
  var currentRow []string

  // START: write init
    %% write init;
  // END: write init

  // START: write exec
    %% write exec;
  // END: write init


    return Feature{}, nil
}
