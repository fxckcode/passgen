package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type EncryptCriteria struct {
	Password string
	HashCost int
}

func EncryptPassword(password string, hashCost int) (string, error) {
	criteria := EncryptCriteria{
		Password: password,
		HashCost: hashCost,
	}

	return encryptPasswordWithCriteria(criteria)
}

func encryptPasswordWithCriteria(criteria EncryptCriteria) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(criteria.Password), criteria.HashCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}