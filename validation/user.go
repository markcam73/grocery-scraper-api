package validation

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateUserParams(name, email string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}

	if strings.TrimSpace(email) == "" {
		return errors.New("email cannot be empty")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}
