package vo

import (
	"errors"
	"regexp"
)

type Name struct {
	value string 
}

var nameRegex = regexp.MustCompile(`^([a-zA-Z '-]+)$`)

func NewName(name string) (Name, error) {
	if name == "" {
		return Name{}, errors.New("name must not be empty")
	}

	if !nameRegex.MatchString(name) {
		return Name{}, errors.New("name can only contain letters, spaces, apostrophes, and hyphens")
	}

	return Name{value: name}, nil
}

func (n Name) String() string {
	return n.value
}

func (n Name) Equals(other Name) bool {
	return n.value == other.value
}