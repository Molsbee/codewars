package snail

func Snail(snaipMap [][]int) []int {
	snail := []int{}
	for len(snaipMap) > 0 {
		// Remove top row
		for _, col := range snaipMap[0] {
			snail = append(snail, col)
		}
		snaipMap = snaipMap[1:]
		if isMatrixEmpty(snaipMap) {
			break
		}

		// Remove last column
		for row := 0; row < len(snaipMap); row++ {
			snail = append(snail, snaipMap[row][len(snaipMap)])
			snaipMap[row] = snaipMap[row][:len(snaipMap)]
		}
		if isMatrixEmpty(snaipMap) {
			break
		}

		// Remove bottom row
		for col := len(snaipMap) - 1; col >= 0; col-- {
			snail = append(snail, snaipMap[len(snaipMap)-1][col])
		}
		snaipMap = snaipMap[:len(snaipMap)-1]
		if isMatrixEmpty(snaipMap) {
			break
		}

		// Remove first column
		for row := len(snaipMap) - 1; row >= 0; row-- {
			snail = append(snail, snaipMap[row][0])
			snaipMap[row] = snaipMap[row][1:]
		}
		if isMatrixEmpty(snaipMap) {
			break
		}
	}
	return snail
}

func isMatrixEmpty(matrix [][]int) bool {
	return len(matrix) == 0 || len(matrix[0]) == 0
}
