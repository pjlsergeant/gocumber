package main

import "fmt"
import "crypto/md5"
import "github.com/sheriff/gocumber/gocumber"

func main() {

	f := gocumber.Feature_digestfeature()
	s := f.Scenario[0].Steps[0]
	//fmt.Printf("%+v\n", f)

	z := gocumber.StepMatcher{}
	z.Add("Given", "a Digest (.+) object", func(s gocumber.StepContext) {
		switch s.Matches[1] {
		case "MD5":
			s.Stash["hash_object"] = md5.New()
		default:
			panic("Don't know about hash type: " + s.Matches[1])
		}
	})
	z.Add("When", "I've added \"(.+)\" to the object", func(s gocumber.StepContext) {
		s.Stash["hash_input"] = s.Stash["hash_input"].(string) + s.Matches[1]
	})

	zz := make(map[string]interface{})
	sc := gocumber.StepContext{Step: &s, Stash: zz}

	rf := z.Match(&sc)
	rf(sc)

	fmt.Printf("Stash: %+v\n", sc.Stash["object"])
}
