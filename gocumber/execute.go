package gocumber

import "testing"
import "fmt"
import "io"
import "bytes"
import "os"

const (
	passed = iota
	pending
	failed
)

type Executor struct {
	StepMatcher *StepMatcher
	Feature     Feature
	T           *testing.T
}

func (e *Executor) RunScenario(sc Scenario) int {

	// Use the global stepmatcher if we don't have one injected
	if e.StepMatcher == nil {
		e.StepMatcher = Steps
	}

	scenarioStatus := passed
	shortCircuit := false

	for _, step := range sc.Steps {
		stepStatus, output := e.RunStep(step, shortCircuit)

		fmt.Printf("Step [%s] ", step.Text)

		switch stepStatus {
		case passed:
			fmt.Println("Passed")
		case failed:
			fmt.Println("Failed")
		case pending:
			fmt.Println("Pending")
		}

		if stepStatus == failed {
			fmt.Printf("Error: [%s]\n", output)
		}

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

func makeInternalTest(sc *StepContext, f func(StepContext)) func(*testing.T) {
	return func(t *testing.T) {
		sc.T = t
		f(*sc)
	}
}

var fakeMatchStep = func(string, string) (bool, error) {
	return true, nil
}

// You throw the switch and the hideous creatures awakes, rises, and lurches
// forward. You're simultaneously elated and terrified that something so
// unnatural could work at all. When you realized what you've unleashed - the
// pure immorality of it - your creation reaches out with its bloodied, mangled
// arms and strangles you.

func testSandbox(name string, sc *StepContext, testFunction func(StepContext)) (bool, string) {
	// Wrap the returned test function to make it accept only testing.T
	wrapped := makeInternalTest(sc, testFunction)

	// Stick it in to testing.InternalTest format
	internalTest := testing.InternalTest{
		Name: name,
		F:    wrapped,
	}

	fakeM := testing.MainStart(fakeMatchStep, []testing.InternalTest{internalTest}, []testing.InternalBenchmark{}, []testing.InternalExample{})

	// Steal STDOUT
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Throw the switch!
	osReturnCode := fakeM.Run()

	// Get the output and return stdout
	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	return (osReturnCode == 0), out
}

func (e Executor) RunStep(s Step, skip bool) (int, string) {

	// Build the step context
	sc := StepContext{
		Step: s,
	}

	// Can we find a matching step definition
	sd := e.StepMatcher.Match(&sc)

	if skip {
		return pending, ""
	} else if sd != nil {
		success, output := testSandbox(s.Text, &sc, sd)

		if success {
			return passed, output
		} else {
			return failed, output
		}
	} else {
		return pending, ""
	}
}
