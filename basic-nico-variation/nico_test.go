package basic_nico_variation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var samples = [][]string{
	{"crazy", "secretinformation", "cseerntiofarmit on  "},
	{"abc", "abcd", "abcd  "},
	{"ba", "1234567890", "2143658709"},
	{"key", "key", "eky"},
}

func Test(t *testing.T) {
	a := assert.New(t)
	for _, sample := range samples {
		actual := Nico(sample[0], sample[1])
		a.Equal(sample[2], actual)
	}
}
