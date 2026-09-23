package main

import "fmt"

func main() {
	var n int

	fmt.Print("Menu:\n1. Add\n2. Delete\n3. Find\n4. Exit\n")
	fmt.Println("Enter number of operation: ")
	fmt.Scan(&n)

	switch n {
	case 1:
		fmt.Println("Add item")
	case 2:
		fmt.Println("Delete item")
	case 3:
		fmt.Println("Find item")
	case 4:
		fmt.Println("Goodbye")
	default:
		fmt.Println("Unknown operation")
	}

	var day int
	fmt.Println("Enter day of the week (1-7): ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}
}
