package validator

import (
	"regexp"
	"unicode"
)

func ValidateEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(regex, email)
	if err != nil {
		return false
	}
	return matched
}

func ValidatePassword(password string) bool {
	if len(password) > 16 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true

		case unicode.IsLower(char):
			hasLower = true

		case unicode.IsDigit(char):
			hasDigit = true

		default:
			return false
		}
	}
	return hasUpper && hasLower && hasDigit
}
