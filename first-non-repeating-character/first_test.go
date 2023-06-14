package first_non_repeating_character

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var samples = [][]string{
	{"stress", "t"},
	{"sTreSS", "T"},
	{"a", "a"},
	{"moonmen", "e"},
	{"", ""},
	{"~><#~><", "#"},
}

func Test(t *testing.T) {
	a := assert.New(t)
	for _, sample := range samples {
		answer := FirstNonRepeating(sample[0])
		a.Equal(sample[1], answer)
	}
}
