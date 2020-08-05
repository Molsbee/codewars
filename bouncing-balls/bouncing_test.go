package bouncing_balls

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBouncingBall(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(3, BouncingBall(3, 0.66, 1.5))
	assert.Equal(3, BouncingBall(40, 0.4, 10))
	assert.Equal(-1, BouncingBall(10, 0.6, 10))
	assert.Equal(-1, BouncingBall(40, 1, 10))
	assert.Equal(-1, BouncingBall(5, -1, 1.5))
	assert.Equal(1, BouncingBall(2, .5, 1))
}
