package main

import "fmt"

func main() {
	name := "Alice"
	age := 30
	salary := 50000.016
	student := true

	fmt.Printf("\nName: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Salary: %.2f\n", salary)
	fmt.Printf("Is Student: %t\n", student)
	fmt.Printf("After 5 years, %s will be %d years old.\n", name, age+5)
}
