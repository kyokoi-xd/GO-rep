package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	fmt.Printf("\nHello, %s!\nNext year, you will be %d years old.\n", name, age+1)
}
