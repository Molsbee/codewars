package encrypt_this

import (
	"fmt"
	"strings"
)

func EncryptThis(text string) string {
	trimmedText := strings.Trim(text, " ")
	if trimmedText == "" {
		return ""
	}

	parts := strings.Split(text, " ")
	encryptedText := ""
	for i := 0; i < len(parts); i++ {
		encryptedText += encrypt(parts[i])
		if i != len(parts)-1 {
			encryptedText += " "
		}
	}

	return encryptedText
}

func encrypt(word string) string {
	runes := []rune(word)
	if len(runes) == 1 {
		return fmt.Sprintf("%d", runes[0])
	} else if len(runes) > 2 {
		temp := runes[1]
		runes[1] = runes[len(runes)-1]
		runes[len(runes)-1] = temp
	}

	return fmt.Sprintf("%d%s", runes[0], string(runes[1:]))
}
