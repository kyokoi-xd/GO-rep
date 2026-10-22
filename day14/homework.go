package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var password string

	fmt.Print("Enter a password:")
	reader := bufio.NewReader(os.Stdin)
	password, _ = reader.ReadString('\n')
	password = strings.TrimSpace(password)

	isValid := isStrongPassword(password)

	if isValid {
		fmt.Println("Password is valid.")
	} else {
		fmt.Println("Password is invalid.")
	}
}

func isStrongPassword(password string) bool {
	if len([]rune(password)) < 8 {
		return false
	}

	flagUpper := false
	flagLower := false
	flagDigit := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			flagUpper = true
		} else if char >= 'a' && char <= 'z' {
			flagLower = true
		} else if char >= '0' && char <= '9' {
			flagDigit = true
		} else {
			return false
		}
	}

	return flagUpper && flagLower && flagDigit
}
