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

	isValid := isValidPassword(password)

	if isValid {
		fmt.Println("Password is valid.")
	} else {
		fmt.Println("Password is invalid.")
	}
}

func isValidPassword(password string) bool {
	if len([]rune(password)) < 5 {
		return false
	}

	for _, char := range password {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	return true
}
