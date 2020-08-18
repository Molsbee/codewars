package pick_peaks

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	pos := make([]int, 0)
	peaks := make([]int, 0)

	foundPeakIndex := 0
	foundPeakValue := 0
	increasing := false
	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			increasing = true
			foundPeakIndex = i
			foundPeakValue = array[i]
		} else if array[i] < array[i-1] && increasing == true {
			pos = append(pos, foundPeakIndex)
			peaks = append(peaks, foundPeakValue)
			increasing = false
		}
	}

	return PosPeaks{
		Pos:   pos,
		Peaks: peaks,
	}
}
