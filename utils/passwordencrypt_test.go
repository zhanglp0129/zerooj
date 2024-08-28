package utils

import (
	"fmt"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	password := "12345678"
	ret, err := PasswordEncrypt(password)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
