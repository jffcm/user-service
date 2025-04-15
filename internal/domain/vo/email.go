package vo

import (
	"errors"
	"regexp"
)

type Email struct {
	value string 
}

var emailRegex = regexp.MustCompile(`^([A-Za-z]+)([0-9]+)?([A-Za-z0-9\.\_]+)?\@(([A-Za-z]+)([0-9]+)?([A-Za-z0-9\.\_]+)?)((\.)([a-zA-Z]+))$`)

func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email{}, errors.New("email must not be empty")
	}

	if !emailRegex.MatchString(email) {
		return Email{}, errors.New("invalid email format. Example: example@domain.com")
	}

	return Email{value: email}, nil
}

func (e Email) String() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	return e.value == other.value
}