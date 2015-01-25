package gocumber

type DocumentLocation struct {
	Filename string
	Line     int
}

type Scenario struct {
	Name       string
	StartsAt   DocumentLocation
	Tags       []string
	Background bool
	Steps      []string
	Data       map[string]string
}

type Feature struct {
	Name                     string
	StartsAt                 DocumentLocation
	ConditionsOfSatisfaction []string
	Tags                     []string
	Language                 string
	Background               Scenario
	Scenario                 []Scenario
}
