package gocumber

import "crypto/md5"
import "encoding/hex"

import (
	"github.com/stretchr/testify/assert"
)

func AddDigestDefinitions(sm *StepMatcher) {

	sm.Add("Given", "a Digest (.+) object", func(s StepContext) {
		switch s.Matches[1] {
		case "MD5":
			s.Stash["hash_object"] = md5.New()
		default:
			panic("Don't know about hash type: " + s.Matches[1])
		}
	})

	sm.Add("When", "I've added \"(.+)\" to the object", func(s StepContext) {
		if s.Stash["hash_input"] == nil {
			s.Stash["hash_input"] = ""
		}
		s.Stash["hash_input"] = s.Stash["hash_input"].(string) + s.Matches[1]
	})

	sm.Add("Then", "the hex output is \"(.+)\"", func(s StepContext) {
		hash_array := md5.Sum([]byte(s.Stash["hash_input"].(string)))
		hash_slice := hash_array[0:]

		// hashed
		hash_hex := hex.EncodeToString(hash_slice)
		assert.Equal(s.T, hash_hex, s.Matches[1], "Matches!")
	})
}
