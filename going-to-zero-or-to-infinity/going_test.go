package going_to_zero_or_to_infinity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoing(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1.275, Going(5))
	assert.Equal(1.2125, Going(6))
	assert.Equal(1.173214, Going(7))
	assert.Equal(1.146651, Going(8))
	assert.Equal(1.052786, Going(20))
	assert.Equal(1.034525, Going(30))
	assert.Equal(1.020416, Going(50))
	assert.Equal(1.008929, Going(113))
	assert.Equal(1.005025, Going(200))
	assert.Equal(1.001742, Going(575))
}
