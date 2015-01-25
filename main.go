package main

import "fmt"
import "github.com/sheriff/gocumber/gocumber"

func main() {

	// Sample document location
	d := gocumber.DocumentLocation{"myfile.feature", 15}

	// A scenario...
	s := gocumber.Scenario{
		Name:       "My Scenario",
		StartsAt:   d,
		Tags:       []string{"foo", "bar", "baz"},
		Background: false,
		Steps:      []string{},
		Data:       map[string]string{"zoomba": "roomba"},
	}

	f := gocumber.Feature{"My Feature", d, []string{"cos", "cos2"}, []string{}, "en", s, []gocumber.Scenario{s}}
	fmt.Println(f)
	fmt.Println("Hi")
}
