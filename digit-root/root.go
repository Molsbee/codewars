package digit_root

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	digits := parseDigits(n, make([]int, 0))
	sum := 0
	for _, d := range digits {
		sum += d
	}
	return DigitalRoot(sum)
}

func parseDigits(n int, digits []int) []int {
	if n < 10 {
		return append(digits, n)
	}

	digits = append(digits, n%10)
	return parseDigits(n/10, digits)
}
