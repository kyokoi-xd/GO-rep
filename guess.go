package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := rand.Intn(100) + 1
	fmt.Println("Guess a number between 1 and 100:")
	reader := bufio.NewReader(os.Stdin)
	attempts := 10
	for attempts > 0 {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}
		if guess < n {
			fmt.Println("Too low! Try again:")
		} else {
			fmt.Println("Too high! Try again:")
		}
		attempts--
		if guess == n {
			fmt.Println("Congratulations! You guessed the number.")
			break
		}
		if attempts == 0 {
			fmt.Println("Sorry, you've run out of attempts. The number was:", n)
		}
	}
}
