package digit_root

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(7, DigitalRoot(16))
	assert.Equal(6, DigitalRoot(942))
}

func TestParseDigits(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{4, 3, 2, 1}, parseDigits(1234, make([]int, 0)))
	assert.Equal([]int{0, 1}, parseDigits(10, make([]int, 0)))
	assert.Equal([]int{1}, parseDigits(1, make([]int, 0)))
	assert.Equal([]int{0, 0, 0, 1}, parseDigits(1000, make([]int, 0)))
}
