package not_very_secure

import "regexp"

func alphanumeric(str string) bool {
	if len(str) > 0 {
		if match, err := regexp.MatchString("^[a-zA-Z0-9]*$", str); err == nil {
			return match
		}
	}

	return false
}
