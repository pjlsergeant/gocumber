package main

import "fmt"
import "github.com/sheriff/gocumber/gocumber"

func main() {
	// Load the feature
	f := gocumber.Feature_digestfeature()

	// Load the step definitions in to a StepMatcher
	sm := gocumber.StepMatcher{}
	gocumber.AddDigestDefinitions(&sm)

	e := gocumber.Executor{
		StepMatcher: sm,
		Feature:     f,
	}

	e.RunScenario(f.Scenario[0])

	fmt.Printf("Stash: %+v\n", e.Stash)
}
