package snail

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne(t *testing.T) {
	a := assert.New(t)
	snail := Snail([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	a.Equal([]int{1, 2, 3, 6, 9, 8, 7, 4, 5}, snail)
}

func TestTwo(t *testing.T) {
	a := assert.New(t)
	snail := Snail([][]int{})
	a.Equal([]int{}, snail)
}

func TestThree(t *testing.T) {
	a := assert.New(t)
	snail := Snail([][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	})
	a.Equal([]int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13}, snail)
}
