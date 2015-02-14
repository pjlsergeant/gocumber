package main

import "fmt"

import ex "github.com/sheriff/gocumber/gocumber/examples"
import "github.com/sheriff/gocumber/gocumber"
import "io/ioutil"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Load the feature
	dat, err := ioutil.ReadFile("./examples/digest/basic.feature")
	check(err)
	f, err := gocumber.ParseFeature(dat)
	check(err)

	// fmt.Printf("Feature: %+v\n", f)

	f = ex.Feature_digestfeature()

	// Load the step definitions in to a StepMatcher
	sm := gocumber.Steps
	// gocumber.AddDigestDefinitions(&sm)

	e := gocumber.Executor{
		StepMatcher: sm,
		Feature:     f,
	}

	e.RunScenario(f.Scenario[0])

	fmt.Printf("Stash: %+v\n", e.Stash)
}
