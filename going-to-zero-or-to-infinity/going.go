package going_to_zero_or_to_infinity

import (
	"math"
	"math/big"
)

func Going(n int) float64 {
	nfac := factorial(n)
	sum := 0.0
	for i := 1; i < n; i++ {
		fac := factorial(i)
		s, _ := fac.Quo(fac, nfac).Float64()
		sum += s
	}

	precision := math.Pow10(6)
	sum = float64(int(sum*precision)) / precision
	return 1.0 + sum
}

func factorial(n int) *big.Float {
	fact := big.NewFloat(1.0)
	for i := 1.0; i <= float64(n); i++ {
		fact = fact.Mul(fact, big.NewFloat(i))
	}
	return fact
}
