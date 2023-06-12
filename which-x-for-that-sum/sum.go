package which_x_for_that_sum

import "math"

func Solve(m float64) float64 {
	return (2*m + 1 - math.Sqrt(4*m+1)) / (2 * m)
}
