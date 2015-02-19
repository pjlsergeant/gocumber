package main

//import "fmt"

import _ "github.com/sheriff/gocumber/gocumber/examples"
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
	f, err := gocumber.ParseFeature(dat, "basic.feature")
	check(err)

	// fmt.Printf("Feature: %+v\n", f.Scenarios[0])

	// Load the step definitions in to a StepMatcher
	sm := gocumber.Steps
	// gocumber.AddDigestDefinitions(&sm)

	e := gocumber.Executor{
		StepMatcher: sm,
		Feature:     f,
	}

	for _, s := range f.Scenarios {
		e.RunScenario(s)
	}

	//fmt.Printf("Stash: %+v\n", e.Stash)
}
