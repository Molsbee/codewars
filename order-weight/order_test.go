package order_weight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderWeight(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("2000 103 123 4444 99", OrderWeight("103 123 4444 99 2000"))
	assert.Equal("11 11 2000 10003 22 123 1234000 44444444 9999", OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"))
	assert.Equal("", OrderWeight(""))
}
