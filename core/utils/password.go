package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type Password interface {
	HashPassword(pass *string)
}

type password struct{}

func NewHashPassword() Password {
	return &password{}
}

func (p *password) HashPassword(pass *string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(*pass), bcrypt.DefaultCost)
	*pass = string(hash)
}
