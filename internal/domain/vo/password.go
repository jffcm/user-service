package vo

import (
	"errors"
)

type Password struct {
	value string
}

// var passwordRegex = regexp.MustCompile(`^(.{0,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{4,})|(.{1,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{3,})|(.{2,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{2,})|(.{3,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{1,})|(.{4,}(([a-zA-Z][^a-zA-Z])|([^a-zA-Z][a-zA-Z])).{0,})$`)

func NewPassword(password string) (Password, error) {
	if len(password) < 6 {
		return Password{}, errors.New("password must be at least 6 characters long")
	}

	// if !passwordRegex.MatchString(password) {
	// 	return Password{}, errors.New("password must be at least 6 characters long and contain a mix of letters and non-letter characters (e.g., numbers or symbols)")
	// }

	return Password{value: password}, nil
}

func (p Password) String() string {
	return p.value
}

func (p Password) Equals(other Password) bool {
	return p.value == other.value
}
