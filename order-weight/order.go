package order_weight

import (
	"sort"
	"strconv"
	"strings"
)

func OrderWeight(s string) string {
	parts := strings.Split(s, " ")
	sort.Slice(parts, func(i, j int) bool {
		s1 := sumDigits(parts[i])
		s2 := sumDigits(parts[j])
		if s1 == s2 {
			return parts[i] < parts[j]
		}

		return s1 < s2
	})

	return strings.Join(parts, " ")
}

func sumDigits(s string) (sum int) {
	digits := strings.Split(s, "")
	for _, digit := range digits {
		d, _ := strconv.Atoi(digit)
		sum += d
	}
	return sum
}
