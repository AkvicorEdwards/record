package util

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

func GetPassword() string {
	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	return string(bytePassword)
}
