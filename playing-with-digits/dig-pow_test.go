package playing_with_digits

import "testing"

func Test(t *testing.T) {
	a := DigPow(89, 1)
	AssertEqual(t, 1, a)

	a = DigPow(92, 1)
	AssertEqual(t, -1, a)

	a = DigPow(695, 2)
	AssertEqual(t, 2, a)

	a = DigPow(46288, 3)
	AssertEqual(t, 51, a)

	a = DigPow(46288, 5)
	AssertEqual(t, 3263, a)
}

func AssertEqual(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected: %d actual: %d", expected, actual)
	}
}
