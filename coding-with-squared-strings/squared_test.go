package coding_with_squared_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsec(t *testing.T) {
	a := assert.New(t)
	originalString := "I.was.going.fishing.that.morning.at.ten.o'clock"
	code := Code(originalString)

	encodedString := "c.nhsoI\nltiahi.\noentinw\ncng.nga\nk..mg.s\n\voao.f.\n\v'trtig"
	a.Equal(encodedString, code)

	decodedString := Decode(encodedString)
	a.Equal(originalString, decodedString)
}
