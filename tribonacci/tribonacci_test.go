package tribonacci

import "testing"

func Test(t *testing.T) {
	Expect(t, []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105}, Tribonacci([3]float64{1, 1, 1}, 10))
	Expect(t, []float64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44}, Tribonacci([3]float64{0, 0, 1}, 10))
	Expect(t, []float64{0, 1, 1, 2, 4, 7, 13, 24, 44, 81}, Tribonacci([3]float64{0, 1, 1}, 10))
	Expect(t, []float64{1, 0, 0, 1, 1, 2, 4, 7, 13, 24}, Tribonacci([3]float64{1, 0, 0}, 10))
	Expect(t, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Tribonacci([3]float64{0, 0, 0}, 10))
	Expect(t, []float64{1, 2, 3, 6, 11, 20, 37, 68, 125, 230}, Tribonacci([3]float64{1, 2, 3}, 10))
	Expect(t, []float64{3, 2, 1, 6, 9, 16, 31, 56, 103, 190}, Tribonacci([3]float64{3, 2, 1}, 10))
	Expect(t, []float64{1}, Tribonacci([3]float64{1, 1, 1}, 1))
	Expect(t, []float64{}, Tribonacci([3]float64{300, 200, 100}, 0))
	Expect(t, []float64{0.5, 0.5, 0.5, 1.5, 2.5, 4.5, 8.5, 15.5, 28.5, 52.5, 96.5, 177.5, 326.5, 600.5, 1104.5, 2031.5, 3736.5, 6872.5, 12640.5, 23249.5, 42762.5, 78652.5, 144664.5, 266079.5, 489396.5, 900140.5, 1655616.5, 3045153.5, 5600910.5, 10301680.5}, Tribonacci([3]float64{0.5, 0.5, 0.5}, 30))
}

func Expect(t *testing.T, expected []float64, actual []float64) {
	if len(actual) != len(expected) {
		t.Errorf("actual and expected contain different number of arguments - expected: %v actual: %v", expected, actual)
		t.Fail()
	}
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("actual: %f does not equal expected: %f", actual[i], expected[i])
		}
	}
}
