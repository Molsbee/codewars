package josephus_survivor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var samples = [][]int{
	{7, 3, 4},
	{11, 19, 10},
	{40, 3, 28},
	{100, 1, 100},
	{1, 300, 1},
}

func Test(t *testing.T) {
	a := assert.New(t)
	for _, sample := range samples {
		answer := JosephusSurvivor(sample[0], sample[1])
		a.Equal(sample[2], answer)
	}
}
