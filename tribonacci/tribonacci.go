package tribonacci

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	}

	if n <= 3 {
		return signature[0:n]
	}

	sequence := make([]float64, n)
	sequence[0] = signature[0]
	sequence[1] = signature[1]
	sequence[2] = signature[2]
	for i := 3; i < n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2] + sequence[i-3]
	}

	return sequence
}
