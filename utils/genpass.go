package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PasswordCriteria struct {
	Length    int
	Symbols   bool
	Numbers   bool
	Uppercase bool
}

func GeneratePassword(length int, symbols, numbers, uppercase bool) (string, error) {
	criteria := PasswordCriteria{
		Length:    length,
		Symbols:   symbols,
		Numbers:   numbers,
		Uppercase: uppercase,
	}
	
	return generatePasswordWithCriteria(criteria)
}

func generatePasswordWithCriteria(criteria PasswordCriteria) (string, error) {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num := "0123456789"
	sym := "!@#$%^&*()-_=+[]{}<>?/"

	chars := lower
	if criteria.Uppercase {
		chars += upper
	}
	if criteria.Numbers {
		chars += num
	}
	if criteria.Symbols {
		chars += sym
	}

	if len(chars) == 0 {
		return "", fmt.Errorf("you must include at least one character set")
	}

	password := make([]byte, criteria.Length)
	for i := range password {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password[i] = chars[nBig.Int64()]
	}

	if !validatePassword(string(password), criteria) {
		return generatePasswordWithCriteria(criteria)
	}

	return string(password), nil
}

func validatePassword(password string, criteria PasswordCriteria) bool {
	hasUpper := !criteria.Uppercase
	hasNumber := !criteria.Numbers
	hasSymbol := !criteria.Symbols

	for _, char := range password {
		if criteria.Uppercase && 'A' <= char && char <= 'Z' {
			hasUpper = true
		}
		if criteria.Numbers && '0' <= char && char <= '9' {
			hasNumber = true
		}
		if criteria.Symbols {
			// Check if the character is a symbol
			c := string(char)
			if c == "!" || c == "@" || c == "#" || c == "$" || c == "%" || c == "^" ||
			   c == "&" || c == "*" || c == "(" || c == ")" || c == "-" || c == "_" ||
			   c == "=" || c == "+" || c == "[" || c == "]" || c == "{" || c == "}" ||
			   c == "<" || c == ">" || c == "?" || c == "/" {
				hasSymbol = true
			}
		}
	}

	return hasUpper && hasNumber && hasSymbol
}