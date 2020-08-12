package max_digit_sum

func Max(v int) int {
	b := 1
	ans := v
	for {
		cur := (v-1)*b + (b - 1)
		sumOfCur := sumOfDigits(cur)
		sumOfAns := sumOfDigits(ans)
		if sumOfCur > sumOfAns || sumOfCur == sumOfAns && cur > ans {
			ans = cur
		}

		v /= 10
		b *= 10

		if v == 0 {
			break
		}
	}
	return ans
}

func sumOfDigits(i int) (sum int) {
	for {
		sum += i % 10
		i /= 10
		if i == 0 {
			break
		}
	}
	return
}
