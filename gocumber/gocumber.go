package gocumber

type DocumentLocation struct {
	Filename string
	Line     int
}

type DocString struct {
	StartsAt    DocumentLocation
	ContentType string
	Content     string
}

type TableData struct {
	StartsAt DocumentLocation
	Columns  []string
	Data     []map[string]string
}

type Step struct {
	Scenario     *Scenario
	Text         string
	StartsAt     DocumentLocation
	Verb         string
	OriginalVerb string
	TableData    *TableData
	DocString    *DocString
}

type Scenario struct {
	Feature    *Feature
	Name       string
	StartsAt   DocumentLocation
	Tags       []string
	Background bool
	Steps      []Step
	TableData  *TableData
	DocString  *DocString
}

type Feature struct {
	Name                     string
	StartsAt                 DocumentLocation
	ConditionsOfSatisfaction []string
	Tags                     []string
	Language                 string
	Background               *Scenario
	Scenarios                []Scenario
}
