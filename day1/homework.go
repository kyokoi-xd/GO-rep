package main

import "fmt"

func main() {
	var age int
	var number int
	var firstNumber int
	var SecondNumber int
	var operation string

	fmt.Println("\nEnter your age: ")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("You are a minor.")
	} else {
		fmt.Println("You are an adult.")
	}

	fmt.Println("\nEnter a number: ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	fmt.Println("\nEnter first number: ")
	fmt.Scanln(&firstNumber)
	fmt.Println("\nEnter second number: ")
	fmt.Scanln(&SecondNumber)

	fmt.Println("\nEnter operation (+, -, *, /): ")
	fmt.Scanln(&operation)

	if operation == "+" {
		fmt.Println("Result:", firstNumber+SecondNumber)
	} else if operation == "-" {
		fmt.Println("Result:", firstNumber-SecondNumber)
	} else if operation == "*" {
		fmt.Println("Result:", firstNumber*SecondNumber)
	} else if operation == "/" {
		if SecondNumber != 0 {
			fmt.Println("Result:", firstNumber/SecondNumber)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	} else {
		fmt.Println("Invalid operation.")
	}
}
