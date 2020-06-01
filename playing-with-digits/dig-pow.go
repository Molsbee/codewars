package playing_with_digits

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	digitStrings := strconv.Itoa(n)

	total := 0
	for i := 0; i < len(digitStrings); i++ {
		digit, _ := strconv.Atoi(string(digitStrings[i]))
		total += int(math.Pow(float64(digit), float64(p+i)))
	}

	if total%n != 0 {
		return -1
	}

	return total / n
}
