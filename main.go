package main

import "fmt"
import "github.com/sheriff/gocumber/gocumber"

func main() {

	f := gocumber.Feature_digestfeature()
	fmt.Printf("%+v\n", f)

}
