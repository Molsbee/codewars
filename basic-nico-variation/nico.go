package basic_nico_variation

import (
	"math"
	"sort"
	"strings"
)

func Nico(key, message string) string {
	sortedKey := strings.Split(key, "")
	sort.Strings(sortedKey)

	var keyIndexShift []int
	for _, char := range key {
		for index, c := range sortedKey {
			if string(char) == c {
				keyIndexShift = append(keyIndexShift, index+1)
			}
		}
	}

	messageLength := len(message)
	messageCharacters := strings.Split(message, "")
	rows := int(math.Ceil(float64(messageLength) / float64(len(key))))
	for i := 0; i < rows*len(keyIndexShift)-messageLength; i++ {
		messageCharacters = append(messageCharacters, " ")
	}

	matrix := make([][]string, rows)
	row := 0
	for i, c := range messageCharacters {
		if i != 0 && i%len(key) == 0 {
			row++
		}
		matrix[row] = append(matrix[row], c)
	}

	for i := 0; i < len(matrix); i++ {
		line := make([]string, len(keyIndexShift))
		for idx, position := range keyIndexShift {
			line[position-1] = matrix[i][idx]
		}
		matrix[i] = line
	}

	encryptedMessage := ""
	for _, row := range matrix {
		for _, col := range row {
			encryptedMessage += col
		}
	}
	return encryptedMessage
}
