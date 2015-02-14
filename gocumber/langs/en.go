package langs

import "github.com/sheriff/gocumber/gocumber"

func Given(regexStr string, body func(gocumber.StepContext)) {
	gocumber.Add("Given", regexStr, body)
}
func When(regexStr string, body func(gocumber.StepContext)) {
	gocumber.Add("When", regexStr, body)
}
func Then(regexStr string, body func(gocumber.StepContext)) {
	gocumber.Add("Then", regexStr, body)
}
func Any(regexStr string, body func(gocumber.StepContext)) {
	gocumber.Add("Any", regexStr, body)
}
