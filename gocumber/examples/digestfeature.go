package examples

import g "github.com/sheriff/gocumber/gocumber"

func Feature_digestfeature() (f g.Feature) {
	f = g.Feature{
		Name:                     "Simple tests of Digest.pm",
		StartsAt:                 g.DocumentLocation{"digestfeature.feature", 2},
		ConditionsOfSatisfaction: []string{"As a developer planning to use Digest.pm", "I want to test the basic functionality of Digest.pm", "In order to have confidence in it"},
		Tags:     []string{},
		Language: "en",
		Background: &g.Scenario{
			Name:       "",
			StartsAt:   g.DocumentLocation{"digestfeature.feature", 7},
			Tags:       []string{},
			Background: true,
			Steps: []g.Step{g.Step{
				Text:         "a usable \"Digest\" class",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 8},
				Verb:         "Given",
				OriginalVerb: "Given",
				TableData:    nil,
				DocString:    nil,
			}},
			TableData: nil,
			DocString: nil,
		},
		Scenarios: []g.Scenario{g.Scenario{
			Name:       "Check MD5",
			StartsAt:   g.DocumentLocation{"digestfeature.feature", 10},
			Tags:       []string{},
			Background: false,
			Steps: []g.Step{g.Step{
				Text:         "a Digest MD5 object",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 11},
				Verb:         "Given",
				OriginalVerb: "Given",
				TableData:    nil,
				DocString:    nil,
			}, g.Step{
				Text:         "I've added \"foo bar baz\" to the object",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 12},
				Verb:         "When",
				OriginalVerb: "When",
				TableData:    nil,
				DocString:    nil,
			}, g.Step{
				Text:         "I've added \"bat ban shan\" to the object",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 13},
				Verb:         "When",
				OriginalVerb: "And",
				TableData:    nil,
				DocString:    nil,
			}, g.Step{
				Text:         "the hex output is \"bcb56b3dd4674d5d7459c95e4c8a41d5\"",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 14},
				Verb:         "Then",
				OriginalVerb: "Then",
				TableData:    nil,
				DocString:    nil,
			}, g.Step{
				Text:         "the base64 output is \"1B2M2Y8AsgTpgAmY7PhCfg\"",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 15},
				Verb:         "Then",
				OriginalVerb: "Then",
				TableData:    nil,
				DocString:    nil,
			}},
			TableData: nil,
			DocString: nil,
		}, g.Scenario{
			Name:       "Check SHA-1",
			StartsAt:   g.DocumentLocation{"digestfeature.feature", 17},
			Tags:       []string{},
			Background: false,
			Steps: []g.Step{g.Step{
				Text:         "a Digest SHA-1 object",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 18},
				Verb:         "Given",
				OriginalVerb: "Given",
				TableData:    nil,
				DocString:    nil,
			}, g.Step{
				Text:         "I've added \"<data>\" to the object",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 19},
				Verb:         "When",
				OriginalVerb: "When",
				TableData:    nil,
				DocString:    nil,
			}, g.Step{
				Text:         "the hex output is \"<output>\"",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 20},
				Verb:         "Then",
				OriginalVerb: "Then",
				TableData:    nil,
				DocString:    nil,
			}},
			TableData: &g.TableData{
				Columns: []string{},
				Data:    []map[string]string{{"data": "foo", "output": "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33"}, {"data": "bar", "output": "62cdb7020ff920e5aa642c3d4066950dd1f01f4d"}, {"output": "bbe960a25ea311d21d40669e93df2003ba9b90a2", "data": "baz"}},
			},
			DocString: nil,
		}, g.Scenario{
			Name:       "MD5 longer data",
			StartsAt:   g.DocumentLocation{"digestfeature.feature", 27},
			Tags:       []string{},
			Background: false,
			Steps: []g.Step{g.Step{
				Text:         "a Digest MD5 object",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 28},
				Verb:         "Given",
				OriginalVerb: "Given",
				TableData:    nil,
				DocString:    nil,
			}, g.Step{
				Text:         "I've added the following to the object",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 29},
				Verb:         "When",
				OriginalVerb: "When",
				TableData:    nil,
				DocString: &g.DocString{
					ContentType: "text/plain",
					Content:     "Here is a chunk of text that works a bit like a HereDoc. We'll split\noff indenting space from the lines in it up to the indentation of the\nfirst \"\"\"\n",
					StartsAt:    g.DocumentLocation{"digestfeature.feature", 31},
				},
			}, g.Step{
				Text:         "the hex output is \"75ad9f578e43b863590fae52d5d19ce6\"",
				StartsAt:     g.DocumentLocation{"digestfeature.feature", 35},
				Verb:         "Then",
				OriginalVerb: "Then",
				TableData:    nil,
				DocString:    nil,
			}},
			TableData: nil,
			DocString: nil,
		}},
	}

	return f
}
