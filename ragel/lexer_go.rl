package gocumber

import "fmt"
import "strings"

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
      // String con = unindent(startCol, substring(data, contentStart, nextKeywordStart-1).replaceFirst("(\\r?\\n)?([\\t ])*\\Z", "").replace("\\\"\\\"\\\"", "\"\"\""));
      // String conType = substring(data, docstringContentTypeStart, docstringContentTypeEnd).trim();
      // listener.docString(conType, con, currentLine);
    }

    action store_feature_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.feature(keyword, nameDescription[0], nameDescription[1], currentLine);
      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_background_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.background(keyword, nameDescription[0], nameDescription[1], currentLine);
      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_scenario_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.scenario(keyword, nameDescription[0], nameDescription[1], currentLine);
      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_scenario_outline_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.scenarioOutline(keyword, nameDescription[0], nameDescription[1], currentLine);
      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_examples_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.examples(keyword, nameDescription[0], nameDescription[1], currentLine);
      if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
      nextKeywordStart = -1
    }

    action store_step_content {
      // listener.step(keyword, substring(data, contentStart, p).trim(), currentLine);
    }

    action store_comment_content {
      // listener.comment(substring(data, contentStart, p).trim(), lineNumber);
      keywordStart = -1
    }

    action store_tag_content {
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
      // String con = substring(data, contentStart, p).trim();
      // currentRow.add(con
      //   .replace("\\|", "|")
      //   .replace("\\n", "\n")
      //   .replace("\\\\", "\\")
      // );
    }

    action store_row {
      // listener.row(currentRow, currentLine);
    }

    action end_feature {
      if(cs < lexer_first_final) {
          content := currentLineContent( data, lastNewline )
          panic(fmt.Sprintf("Lexing error on line %d: '%s'. See http://wiki.github.com/cucumber/gherkin/lexingerror for more information.", lineNumber, content))
          // Silence warnings about these being unused...
          panic( currentLine )
          panic(contentStart)
          panic(docstringContentTypeStart)
          panic(docstringContentTypeEnd)
          panic(startCol)
      }
      // } else {
      //   listener.eof();
      // }
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
