package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1

	fmt.Println("\nGuess the number between 1 and 100:")
	var guess int
	attempts := 0
	for fmt.Scanln(&guess); guess != secret; fmt.Scanln(&guess) {
		attempts++
		if guess < secret {
			fmt.Println("Too low!")
		} else if guess > secret {
			fmt.Println("Too high!")
		}
	}
	fmt.Println("Correct!", "You guessed it in", attempts, "attempts.")
}
