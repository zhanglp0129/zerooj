package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
)

const (
	salt = "zerooj"
)

func PasswordEncrypt(password string) (string, error) {
	h := sha256.New()
	_, err := io.WriteString(h, password)
	if err != nil {
		return "", err
	}
	_, err = io.WriteString(h, salt)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
