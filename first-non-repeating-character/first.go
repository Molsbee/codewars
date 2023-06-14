package first_non_repeating_character

import (
	"strings"
)

func FirstNonRepeating(str string) string {
	characters := strings.Split(str, "")
	for _, c := range characters {
		occurrenceCount := 0
		lowerCaseCharacter := strings.ToLower(c)
		upperCaseCharacter := strings.ToUpper(c)
		if lowerCaseCharacter != upperCaseCharacter {
			occurrenceCount = strings.Count(str, lowerCaseCharacter) + strings.Count(str, upperCaseCharacter)
		} else {
			occurrenceCount = strings.Count(str, c)
		}

		if occurrenceCount == 1 {
			return c
		}
	}
	return ""
}
