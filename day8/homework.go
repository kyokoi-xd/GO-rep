package main

import "fmt"

func main() {
	var n int

	fmt.Print("Menu:\n1. Add\n2. Delete\n3. Find\n4. Exit\n")
	fmt.Print("Enter number of operation: ")
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
	fmt.Print("Enter day of the week (1-7): ")
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
	default:
		fmt.Println("Unknown day")
	}

	var firstNum, secondNum float64
	var operation int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNum)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNum)

	fmt.Println("Choose operation:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	fmt.Print("Enter number of operation: ")
	fmt.Scan(&operation)

	switch operation {
	case 1:
		fmt.Println("Result:", add(firstNum, secondNum))
	case 2:
		fmt.Println("Result:", subtract(firstNum, secondNum))
	case 3:
		fmt.Println("Result:", multiply(firstNum, secondNum))
	case 4:
		result, err := divide(firstNum, secondNum)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Result:", result)
		}
	default:
		fmt.Println("Unknown operation")
	}
}

func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a float64, b float64) float64 {
	return a - b
}

func multiply(a float64, b float64) float64 {
	return a * b
}

func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
