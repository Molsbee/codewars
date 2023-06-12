package coding_with_squared_strings

import (
	"math"
	"strings"
)

func Code(s string) string {
	characters := strings.Split(s, "")
	sq := math.Ceil(math.Sqrt(float64(len(characters))))
	squaredSize := int(sq)

	matrix := createMatrix(characters, squaredSize)

	code := ""
	for col := 0; col < squaredSize; col++ {
		for row := squaredSize - 1; row >= 0; row-- {
			code += matrix[row][col]
			if row == 0 && col != squaredSize-1 {
				code += "\n"
			}
		}
	}

	return code
}

func Decode(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	characters := strings.Split(s, "")
	sq := math.Ceil(math.Sqrt(float64(len(characters))))
	squaredSize := int(sq)

	matrix := createMatrix(characters, squaredSize)

	decode := ""
	for col := squaredSize - 1; col >= 0; col-- {
		for row := 0; row < squaredSize; row++ {
			char := matrix[row][col]
			if char != string(rune(11)) {
				decode += char
			}
		}
	}

	return decode
}

func createMatrix(characters []string, size int) (matrix [][]string) {
	for r := 0; r < size; r++ {
		var row []string
		for c := 0; c < size; c++ {
			if r*size+c < len(characters) {
				row = append(row, characters[r*size+c])
			} else {
				row = append(row, string(rune(11)))
			}
		}
		matrix = append(matrix, row)
	}
	return
}
