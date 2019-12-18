package common

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// GetPassword displays a message and prompts user to enter a password
func GetPassword(msg string) []byte {
	for {
		fmt.Println(msg)
		fmt.Print("> ")
		password, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Printf("invalid input: %s\n", err)
		} else {
			fmt.Printf("\n")
			return password
		}
	}
}
