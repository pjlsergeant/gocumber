package gocumber

  %%{
    machine lexer;
    alphtype byte;

    action begin_content {
      // contentStart = p;
      // currentLine = lineNumber;
      // if(keyword != null) {
      //   startCol = p - lastNewline - (keyword.length() + 1);
      // }
    }

    action start_docstring {
      // currentLine = lineNumber;
      // startCol = p - lastNewline;
    }

    action begin_docstring_content {
      // contentStart = p;
    }

    action start_docstring_content_type {
      // docstringContentTypeStart = p;
    }

    action end_docstring_content_type {
      // docstringContentTypeEnd = p;
    }

    action store_docstring_content {
      // String con = unindent(startCol, substring(data, contentStart, nextKeywordStart-1).replaceFirst("(\\r?\\n)?([\\t ])*\\Z", "").replace("\\\"\\\"\\\"", "\"\"\""));
      // String conType = substring(data, docstringContentTypeStart, docstringContentTypeEnd).trim();
      // listener.docString(conType, con, currentLine);
    }

    action store_feature_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.feature(keyword, nameDescription[0], nameDescription[1], currentLine);
      // if(nextKeywordStart != -1) p = nextKeywordStart - 1;
      // nextKeywordStart = -1;
    }

    action store_background_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.background(keyword, nameDescription[0], nameDescription[1], currentLine);
      // if(nextKeywordStart != -1) p = nextKeywordStart - 1;
      // nextKeywordStart = -1;
    }

    action store_scenario_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.scenario(keyword, nameDescription[0], nameDescription[1], currentLine);
      // if(nextKeywordStart != -1) p = nextKeywordStart - 1;
      // nextKeywordStart = -1;
    }

    action store_scenario_outline_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.scenarioOutline(keyword, nameDescription[0], nameDescription[1], currentLine);
      // if(nextKeywordStart != -1) p = nextKeywordStart - 1;
      // nextKeywordStart = -1;
    }

    action store_examples_content {
      // String[] nameDescription = nameAndUnindentedDescription(startCol, keywordContent(data, p, eof, nextKeywordStart, contentStart));
      // listener.examples(keyword, nameDescription[0], nameDescription[1], currentLine);
      // if(nextKeywordStart != -1) p = nextKeywordStart - 1;
      // nextKeywordStart = -1;
    }

    action store_step_content {
      // listener.step(keyword, substring(data, contentStart, p).trim(), currentLine);
    }

    action store_comment_content {
      // listener.comment(substring(data, contentStart, p).trim(), lineNumber);
      // keywordStart = -1;
    }

    action store_tag_content {
      // listener.tag(substring(data, contentStart, p).trim(), currentLine);
      // keywordStart = -1;
    }

    action inc_line_number {
      // lineNumber++;
    }

    action last_newline {
      // lastNewline = p + 1;
    }

    action start_keyword {
      // if(keywordStart == -1) keywordStart = p;
    }

    action end_keyword {
      // keyword = substring(data, keywordStart, p).replaceFirst(":$","");
      // keywordStart = -1;
    }

    action next_keyword_start {
      // nextKeywordStart = p;
    }

    action start_row {
    //   p = p - 1;
    //   currentRow = new ArrayList<String>();
    //   currentLine = lineNumber;
    }

    action begin_cell_content {
      // contentStart = p;
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
      // if(cs < lexer_first_final) {
      //   String content = currentLineContent(data, lastNewline);
      //   throw new LexingError("Lexing error on line " + lineNumber + ": '" + content + "'. See http://wiki.github.com/cucumber/gherkin/lexingerror for more information.");
      // } else {
      //   listener.eof();
      // }
    }

    include lexer_common "lexer_common_english.rl";
  }%%

  // START: write data noerror;
  %% write data noerror;
  // END: write data noerror;

func ParseFeature(data []byte) (feature Feature, err error) {
  cs, p, pe, eof := 0, 0, len(data), len(data)

  // public void scan(String source)  {
  //   String input = source + "\n%_FEATURE_END_%";
  //   byte[] data = null;
  //   try {
  //     data = input.getBytes("UTF-8");
  //   } catch(UnsupportedEncodingException e) {
  //     throw new RuntimeException(e);
  //   }
  //   int cs, p = 0, pe = data.length;
  //   int eof = pe;

  //   int lineNumber = 1;
  //   int lastNewline = 0;

  //   int contentStart = -1;
  //   int currentLine = -1;
  //   int docstringContentTypeStart = -1;
  //   int docstringContentTypeEnd = -1;
  //   int startCol = -1;
  //   int nextKeywordStart = -1;
  //   int keywordStart = -1;
  //   String keyword = null;
  //   List<String> currentRow = null;

  // START: write init
    %% write init;
  // END: write init

  // START: write exec
    %% write exec;
  // END: write init


    return Feature{}, nil
}
