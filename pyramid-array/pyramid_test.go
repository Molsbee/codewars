package pyramid_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPyramid(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([][]int{}, Pyramid(0))
	assert.Equal([][]int{{1}}, Pyramid(1))
	assert.Equal([][]int{{1}, {1, 1}}, Pyramid(2))
	assert.Equal([][]int{{1}, {1, 1}, {1, 1, 1}}, Pyramid(3))
}
