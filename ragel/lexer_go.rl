package gocumber

import "fmt"
import "strings"
import "regexp"
import "strconv"

const (
  typeBackground = iota
  typeScenario
  typeOutline
)

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

      fb.addDocString( currentLine, con, conType )
    }

    action store_feature_content {
        kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        name, description := nameAndUnindentedDescription( startCol, kcon );

        fb.addFeature( currentLine, name, description )

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    }

    action store_background_content {
        kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        name, _ := nameAndUnindentedDescription( startCol, kcon );

        fb.addScenario( currentLine, typeBackground, name )

        // background := Scenario{
        //     Feature:    &feature,
        //     Name:       name,
        //     StartsAt:   DocumentLocation{ Filename: filename, Line: currentLine },
        //     Tags:       tags,
        //     Background: true,
        //     Steps:      []Step{},
        //     TableData:  nil,
        // }

        // lastScenario = &background
        // tags = tags[:0]

        // if ( examplesTable ) {
        //     background.TableData = makeTable(currentTable)
        //     currentTable = currentTable[:0]
        //     examplesTable = false
        // }

        // feature.Background = &background

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    }

    action store_scenario_content {
        kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        name, _ := nameAndUnindentedDescription( startCol, kcon );

        fb.addScenario( currentLine, typeScenario, name )

        // scenario := Scenario{
        //     Feature:    &feature,
        //     Name:       name,
        //     StartsAt:   DocumentLocation{ Filename: filename, Line: currentLine },
        //     Tags:       tags,
        //     Background: false,
        //     Steps:      []Step{},
        //     TableData:  nil,
        // }

        // lastScenario = &scenario
        // tags = tags[:0]

        // if ( examplesTable ) {
        //     scenario.TableData = makeTable(currentTable)
        //     currentTable = currentTable[:0]
        //     examplesTable = false
        // }

        // feature.Scenario = append( feature.Scenario, scenario )

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    }

    action store_scenario_outline_content {
        kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        name, _ := nameAndUnindentedDescription( startCol, kcon );

        fb.addScenario( currentLine, typeOutline, name )

        // scenario := Scenario{
        //     Feature:    &feature,
        //     Name:       name,
        //     StartsAt:   DocumentLocation{ Filename: filename, Line: currentLine },
        //     Tags:       tags,
        //     Background: false,
        //     Steps:      []Step{},
        //     TableData:  nil,
        // }

        // lastScenario = &scenario
        // tags = tags[:0]

        // if ( examplesTable ) {
        //     scenario.TableData = makeTable(currentTable)
        //     currentTable = currentTable[:0]
        //     examplesTable = false
        // }

        // feature.Scenario = append( feature.Scenario, scenario )

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    }

    action store_examples_content {
        //kcon := keywordContent(data, p, eof, nextKeywordStart, contentStart)
        //name, description := nameAndUnindentedDescription( startCol, kcon );

        fb.addExamples( currentLine )

        // examplesTable = true;

        if(nextKeywordStart != -1) { p = nextKeywordStart - 1 }
        nextKeywordStart = -1
    }

    action store_step_content {
        con := string( data[contentStart:p] )
        con = strings.TrimSpace( con )

        // step := Step{
        //     Scenario:     lastScenario,
        //     Text:         con,
        //     StartsAt:DocumentLocation{ Filename: filename, Line: currentLine },
        //     Verb: keyword,
        //     OriginalVerb: keyword,
        //     TableData: nil,
        //     DocString: "",
        // }

        fb.addStep( currentLine, keyword, con )

        // //fmt.Printf("Before: [%+v]\n", lastScenario.Steps)
        // lastScenario.Steps = append( lastScenario.Steps, step )
        //fmt.Printf("After : [%+v]\n", lastScenario.Steps)
    }

    action store_comment_content {
        con := string( data[contentStart:p] )
        con = strings.TrimSpace( con )
        // Discard comments
        // fmt.Printf("Store Thing: [Comment]\nStore Keyword: [%s]\nStore Content: [%s]\n", keyword, con )
        // listener.comment(substring(data, contentStart, p).trim(), lineNumber);
        keywordStart = -1
    }

    action store_tag_content {
        con := string( data[contentStart:p] )
        con = strings.TrimSpace( con )
        con = strings.TrimLeft( con, "@" ) // Don't need the @

        fb.addTag( currentLine, con )

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
      //currentRow = currentRow[:0]
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
        fb.addCell( con )
        // currentRow = append( currentRow, con )
    }

    action store_row {
        fb.rowEnd()
        // currentTable = append( currentTable, currentRow )
    }

    action end_feature {
      if(cs < lexer_first_final) {
          content := currentLineContent( data, lastNewline )
          panic(fmt.Sprintf("Lexing error on line %d: '%s'. See http://wiki.github.com/cucumber/gherkin/lexingerror for more information.", lineNumber, content))
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


func nameAndUnindentedDescription(startCol int, textBytes []byte) (string, []string) {
    text := unindent( startCol, textBytes )
    text = strings.TrimSpace( text )
    lines := strings.Split(text, "\n")

      for index,element := range lines {
        lines[index] = strings.TrimSpace(element)
      }

    if ( len(lines) == 0 ) {
      return "", []string{}
    } else if ( len(lines) == 1 ) {
      return lines[0], []string{}
      } else {
        return lines[0], lines[1:]
      }
}

func makeTable(dl *DocumentLocation, rows [][]string) (*TableData) {
    columns := rows[0]
    bodyRows := rows[1:0]

    panic("Finish makeTable()")

    table := &TableData{
      StartsAt: *dl,
    }
    return table
}

type FeatureBuilder struct {
    feature *Feature
    scenario *Scenario
    step *Step
    table [][]string
    row []string
    tags []string
    examplesTable *DocumentLocation
    filename string
    lastVerb string
}

func (fb *FeatureBuilder) addFeature ( currentLine int, name string, cos []string ) {
    //fmt.Printf("[%d] addFeature: [%s] {%s}\n", currentLine, name, strings.Join( cos, "\n" ) )
    //fmt.Printf("[%d] .. tags: [%s]\n", currentLine, strings.Join( fb.getTags(), ", "))

    fb.feature = &Feature{
        Name: name,
        StartsAt: DocumentLocation{ Filename: fb.filename, Line: currentLine },
        ConditionsOfSatisfaction: cos,
        Tags: fb.getTags(),
        Language: "en",
        Background: nil,
        Scenarios: []Scenario{},
    }
}

func (fb *FeatureBuilder) getFeature () (*Feature) {
    return fb.feature
}

// Tags
func (fb *FeatureBuilder) addTag (currentLine int, tag string) {
    // fmt.Printf("[%d] addTag: [%s]\n", currentLine, tag)
    fb.tags = append( fb.tags, tag )
}

func (fb *FeatureBuilder) getTags () ([]string) {
    tags := fb.tags
    fb.tags = fb.tags[:0]
    return tags
}

func (fb *FeatureBuilder) addCell (content string) {
    fb.row = append( fb.row, content )
}

func (fb *FeatureBuilder) rowEnd() {
    fb.table = append( fb.table, fb.row )
    fb.row = fb.row[:0]
}

func (fb *FeatureBuilder) getRows () ([][]string) {
    rows := fb.table
    fb.table = fb.table[:0]
    return rows
}

func (fb *FeatureBuilder) addScenario ( currentLine int, stype int, name string ) {
    fmt.Printf("[%d] addScenario: [%d][%s]\n", currentLine, stype, name)
    fmt.Printf("[%d] .. tags: [%s]\n", currentLine, strings.Join(fb.getTags(), ", "))

    if ( fb.examplesTable != nil ) {
        dl := fb.examplesTable
        fb.examplesTable = nil
        fb.step.TableData = makeTable( dl, fb.getRows() )
    }

    scenario := Scenario{
        Feature: fb.feature,
        Name:       name,
        StartsAt:   DocumentLocation{ Filename: fb.filename, Line: currentLine },
        Tags:       fb.getTags(),
        Background: false,
        Steps:      []Step{},
        TableData:  nil,
    }

    if ( stype == typeBackground ) {
        scenario.Background = true
        fb.feature.Background = &scenario
        fb.scenario = &scenario
    } else {
        fb.feature.Scenarios = append( fb.feature.Scenarios, scenario )
        sref := &fb.feature.Scenarios[ len(fb.feature.Scenarios) - 1 ]
        fb.scenario = sref
    }

    fb.lastVerb = ""
}

func (fb *FeatureBuilder) addStep ( currentLine int, verb string, content string ) {
    fmt.Printf("[%d] addStep [%s] [%s]\n", currentLine, verb, content )

    // Deal with "and" as a verb
    originalVerb := verb
    if ( verb == "And" ) {
        if ( fb.lastVerb == "" ) {
            panic("Can't use 'And' as first step of a scenario")
        } else {
            verb = fb.lastVerb
        }
    }
    fb.lastVerb = verb

    step := Step{
        Scenario: fb.scenario,
        Text: content,
        StartsAt:   DocumentLocation{ Filename: fb.filename, Line: currentLine },
        Verb: verb,
        OriginalVerb: originalVerb,
        TableData: nil,
        DocString: nil,
    }

    fb.scenario.Steps = append( fb.scenario.Steps, step )
    sref := &fb.scenario.Steps[ len(fb.scenario.Steps) - 1 ]
    fb.step = sref
}

func (fb *FeatureBuilder) addDocString (currentLine int, docstring string, contentType string) {
    if ( len(contentType) < 1 ) {
        contentType = "text/plain"
    }

    d := &DocString{
        ContentType: docstring,
        StartsAt:   DocumentLocation{ Filename: fb.filename, Line: currentLine },
        Content: docstring,
    }

    fb.step.DocString = d

}

func (fb *FeatureBuilder) addExamples ( currentLine int ) {
    fb.examplesTable = &DocumentLocation{ Filename: fb.filename, Line: currentLine }
    //fmt.Printf("[%d] addExamples\n", currentLine);
}

func ParseFeature(data []byte, filename string) (feature Feature, err error) {

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

    fb := FeatureBuilder{
        filename: filename,
    }

  // START: write init
    %% write init;
  // END: write init

  // START: write exec
    %% write exec;
  // END: write init

    feature = *fb.getFeature()
    fmt.Printf("Feature: %+v\n", feature.Scenarios)

    return
}
