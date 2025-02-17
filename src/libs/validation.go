package libs

import (
	"errors"
	"regexp"
)

func Validation(email, pass string) error {
	if len(pass) < 8 {
		return errors.New("password must contain at least 8 characters")
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	r := emailRegex.MatchString(email)
	if !r {
		return errors.New("invalid email type")
	}

	return nil
}
