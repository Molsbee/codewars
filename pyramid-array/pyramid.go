package pyramid_array

func Pyramid(n int) [][]int {
	list := make([][]int, n)
	for i := 1; i <= n; i++ {
		subList := make([]int, i)
		for j := 1; j <= i; j++ {
			subList[j-1] = 1
		}
		list[i-1] = subList
	}

	return list
}
