package examples

import g "github.com/sheriff/gocumber/gocumber"
import d "github.com/sheriff/gocumber/gocumber/langs"

//import "crypto"
import "hash"
import "crypto/md5"
import "encoding/base64"
import "encoding/hex"

import (
	"github.com/stretchr/testify/assert"
)

func init() {

	var hash_object hash.Hash

	d.Given("a Digest (.+) object", func(s g.StepContext) {
		switch s.Matches[1] {
		case "MD5":
			hash_object = md5.New()
		default:
			panic("Don't know about hash type: " + s.Matches[1])
		}
	})

	d.When("I've added \"(.+)\" to the object", func(s g.StepContext) {
		hash_object.Write([]byte(s.Matches[1]))
	})

	d.Then("the hex output is \"(.+)\"", func(s g.StepContext) {
		assert.Equal(s.T, s.Matches[1], hex.EncodeToString(hash_object.Sum(nil)), "Matches!")
	})

	d.Then("the base64 output is \"(.+)\"", func(s g.StepContext) {
		assert.Equal(s.T, s.Matches[1], base64.StdEncoding.EncodeToString(hash_object.Sum(nil)), "Matches!")
	})
}