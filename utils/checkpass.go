package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type CheckCriteria struct {
	Password string
	Hash     string
}

func CheckPassword(password string, hash string) (bool, error) {
	criteria := CheckCriteria{
		Password: password,
		Hash:     hash,
	}

	return checkPasswordWithCriteria(criteria)
}

func checkPasswordWithCriteria(criteria CheckCriteria) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(criteria.Hash), []byte(criteria.Password))
	if err != nil {
		return false, err
	}
	return true, nil
}

