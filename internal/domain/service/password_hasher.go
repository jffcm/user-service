package service

import (
    "golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
    Hash(password string) (string, error)
    Compare(hashedPassword, password string) bool
}

type bcryptPasswordHasher struct{}

func NewBcryptPasswordHasher() PasswordHasher {
    return &bcryptPasswordHasher{}
}

func (b *bcryptPasswordHasher) Hash(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashedPassword), err
}

func (b *bcryptPasswordHasher) Compare(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}