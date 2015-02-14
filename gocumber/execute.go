package gocumber

import "testing"
import "fmt"

const (
	passed = iota
	pending
	failed
)

type Executor struct {
	StepMatcher *StepMatcher
	Feature     Feature
	Stash       map[string]interface{}
	T           *testing.T
}

func (e *Executor) RunScenario(sc Scenario) int {

	// Check we already have a stash
	if e.Stash == nil {
		e.Stash = make(map[string]interface{})
	}

	// Use the global stepmatcher if we don't have one injected
	if e.StepMatcher == nil {
		e.StepMatcher = Steps
	}

	scenarioStatus := passed
	shortCircuit := false

	for _, step := range sc.Steps {
		stepStatus := e.RunStep(step, shortCircuit)

		fmt.Printf("Step [%s] returning %i\n", step.Text, stepStatus)

		if scenarioStatus == passed {
			if stepStatus == pending {
				scenarioStatus = pending
				shortCircuit = true
			}

			if stepStatus == failed {
				scenarioStatus = failed
				shortCircuit = true
			}
		}
	}

	return scenarioStatus
}

func (e Executor) RunStep(s Step, skip bool) int {

	// New testing context
	t := testing.T{}

	// Build the step context
	sc := StepContext{
		Stash: e.Stash,
		Step:  s,
		T:     &t,
	}

	// Can we find a matching step definition
	sd := e.StepMatcher.Match(&sc)

	status := passed

	if sd != nil {
		t.Logf("Executing for %s", s.Text)
		sd(sc)
		if t.Failed() {
			status = failed
		} else {
			status = passed
		}
	} else {
		status = pending
	}

	return status
}
