package which_x_for_that_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	a := assert.New(t)
	result := Solve(8.0)
	a.Equal(0.7034648345913732, result)
}
