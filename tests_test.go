package main

import "fmt"
import "github.com/sheriff/gocumber/gocumber"
import "testing"

func TestTests(t *testing.T) {

	// Load the feature
	f := gocumber.Feature_digestfeature()

	// Load the step definitions in to a StepMatcher
	sm := gocumber.StepMatcher{}
	gocumber.AddDigestDefinitions(&sm)

	e := gocumber.Executor{
		StepMatcher: sm,
		Feature:     f,
		T:           t,
	}

	e.RunStep(f.Scenario[0].Steps[0])
	e.RunStep(f.Scenario[0].Steps[1])
	e.RunStep(f.Scenario[0].Steps[2])
	e.RunStep(f.Scenario[0].Steps[3])

	fmt.Printf("Stash: %+v\n", e.Stash)
}
