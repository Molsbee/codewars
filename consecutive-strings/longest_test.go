package consecutive_strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsec(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("abigailtheta", LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
	assert.Equal("oocccffuucccjjjkkkjyyyeehh", LongestConsec([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1))
	assert.Equal("", LongestConsec([]string{}, 3))
	assert.Equal("oocccffuucccjjjkkkjyyyeehh", LongestConsec([]string{"oocccffuucccjjjkkkjyyyeehh"}, 1))
	assert.Equal("wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck", LongestConsec([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2))
}
