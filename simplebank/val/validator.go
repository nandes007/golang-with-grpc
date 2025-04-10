package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername    = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidateFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}

	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 20); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowecase letters, digits, or underscore")
	}

	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 20); err != nil {
		return err
	}

	if !isValidateFullName(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}

	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}

	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be positive integer")
	}

	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
