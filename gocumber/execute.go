package gocumber

import "fmt"
import "testing"

type Executor struct {
	StepMatcher StepMatcher
	Feature     Feature
	Stash       map[string]interface{}
	T           *testing.T
}

func (e *Executor) RunStep(s Step) {

	// Check we already have a stash
	if e.Stash == nil {
		e.Stash = make(map[string]interface{})
	}

	// Build the step context
	sc := StepContext{
		Stash: e.Stash,
		Step:  s,
		T:     e.T,
	}

	// Can we find a matching step definition
	sd := e.StepMatcher.Match(&sc)

	if sd != nil {
		fmt.Println("Executing for %s", s.Text)
		sd(sc)
		if e.T.Failed() {
			fmt.Println("\tThat failed")
		} else {
			fmt.Println("\tThat passed")
		}
	} else {
		fmt.Println("No match found for %s", s.Text)
	}

}
