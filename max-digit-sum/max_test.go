package max_digit_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, Max(1))
	assert.Equal(2, Max(2))
	assert.Equal(18, Max(18))
	assert.Equal(99, Max(100))
	assert.Equal(9, Max(10))
	assert.Equal(99, Max(110))
	assert.Equal(1999, Max(2090))
	assert.Equal(29999999999, Max(30316357896))
	assert.Equal(29999999999, Max(30000000000))
	assert.Equal(5899999999, Max(5981222416))
}
