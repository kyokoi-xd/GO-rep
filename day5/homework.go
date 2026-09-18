package main

import "fmt"

func main() {
	var name string
	var age int
	var salary float64
	var developer bool

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Print("Enter your salary: ")
	fmt.Scanln(&salary)
	fmt.Print("Are you a developer (true/false): ")
	fmt.Scanln(&developer)

	fmt.Println("\nUser Information:")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Developer:", developer)
}
