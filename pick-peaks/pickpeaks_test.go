package pick_peaks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPickPeaks(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		PickPeaks([]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}),
	)

	assert.Equal(
		PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}),
	)

	assert.Equal(
		PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}),
	)

	assert.Equal(
		PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
		PickPeaks([]int{2, 1, 3, 1, 2, 2, 2, 2, 1}),
	)

	assert.Equal(
		PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		PickPeaks([]int{2, 1, 3, 1, 2, 2, 2, 2}),
	)
}
