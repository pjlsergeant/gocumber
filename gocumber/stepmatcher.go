package gocumber

import "regexp"
import "testing"

// So basically a step matcher needs to take a StepContext,
// and it needs to return it with the matches populated...
type StepMatcherInterface interface {
	Match(StepContext) func(StepContext)
}

type StepMatcher struct {
	Given []StepDefinition
	When  []StepDefinition
	Then  []StepDefinition
	Any   []StepDefinition
}

type StepDefinition struct {
	Regexp *regexp.Regexp
	Body   func(StepContext)
}

type StepContext struct {
	Stash         map[string]interface{}
	Matches       []string
	MatchIndicies []int
	Step          Step
	T             *testing.T
}

func (m *StepMatcher) Add(verb string, regexStr string, body func(StepContext)) {
	regex := regexp.MustCompile(regexStr)
	definition := StepDefinition{regex, body}

	switch verb {
	case "Given":
		m.Given = append(m.Given, definition)
	case "When":
		m.When = append(m.When, definition)
	case "Then":
		m.Then = append(m.Then, definition)
	case "Any":
		m.Any = append(m.Any, definition)
	default:
		panic("You can only define handlers for Given/When/Then/Any, but you have: " + verb)
	}

}

func (m StepMatcher) Match(c *StepContext) (f func(StepContext)) {
	var stepJar *[]StepDefinition

	// We should have checked this when parsing the feature file,
	// so panic if this has gone weird along the way
	switch c.Step.Verb {
	case "Given":
		stepJar = &m.Given
	case "When":
		stepJar = &m.When
	case "Then":
		stepJar = &m.Then
	default:
		panic("Somehow your verb isn't Given/When/Then, it's " + c.Step.Verb)
	}

	for _, jar := range []*[]StepDefinition{stepJar, &m.Any} {
		for _, definition := range *jar {

			// Do we match at all?
			matches := definition.Regexp.FindStringSubmatchIndex(c.Step.Text)

			// If not, on to the next!
			if matches == nil {
				continue
			}

			// Ahhh, but if we do, we need to do some work
			(*c).MatchIndicies = matches

			// If you were expecting this code to be performant, and not geared
			// to programmer laziness, you missed the fact that this is a Ruby
			// tool :-P
			matchStrings := definition.Regexp.FindStringSubmatch(c.Step.Text)
			(*c).Matches = matchStrings

			f = definition.Body
			return f
		}
	}

	return f
}
