package gocumber

type DocumentLocation struct {
	Filename string
	Line     int
}

type TableData struct {
	Columns []string
	Data    []map[string]string
}

type Step struct {
	Scenario     *Scenario
	Text         string
	StartsAt     DocumentLocation
	Verb         string
	OriginalVerb string
	TableData    *TableData
	DocString    string
}

type Scenario struct {
	Feature    *Feature
	Name       string
	StartsAt   DocumentLocation
	Tags       []string
	Background bool
	Steps      []Step
	TableData  *TableData
	DocString  string
}

type Feature struct {
	Name                     string
	StartsAt                 DocumentLocation
	ConditionsOfSatisfaction []string
	Tags                     []string
	Language                 string
	Background               *Scenario
	Scenario                 []Scenario
}
