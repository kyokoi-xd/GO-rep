package main

import "fmt"

func main() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age < 13 {
		fmt.Println(" Under 13")
	} else if age < 18 {
		fmt.Println(" Teenager")
	} else if age < 65 {
		fmt.Println(" Adult")
	} else {
		fmt.Println(" Senior")
	}
}
