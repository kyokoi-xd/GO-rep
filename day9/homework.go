package main

import "fmt"

func main() {
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
		result, err := save_divide(firstNum, secondNum)
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

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
