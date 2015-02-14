package examples

import g "github.com/sheriff/gocumber/gocumber"
import d "github.com/sheriff/gocumber/gocumber/langs"
import "crypto/md5"
import "encoding/hex"

import (
	"github.com/stretchr/testify/assert"
)

func init() {
	d.Given("a Digest (.+) object", func(s g.StepContext) {
		switch s.Matches[1] {
		case "MD5":
			s.Stash["hash_object"] = md5.New()
		default:
			panic("Don't know about hash type: " + s.Matches[1])
		}
	})

	d.When("I've added \"(.+)\" to the object", func(s g.StepContext) {
		if s.Stash["hash_input"] == nil {
			s.Stash["hash_input"] = ""
		}
		s.Stash["hash_input"] = s.Stash["hash_input"].(string) + s.Matches[1]
	})

	d.Then("the hex output is \"(.+)\"", func(s g.StepContext) {
		hash_array := md5.Sum([]byte(s.Stash["hash_input"].(string)))
		hash_slice := hash_array[0:]

		// hashed
		hash_hex := hex.EncodeToString(hash_slice)
		assert.Equal(s.T, hash_hex, s.Matches[1], "Matches!")
	})
}
