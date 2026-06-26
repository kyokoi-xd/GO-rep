package main

import "fmt"

func main() {
	var sum int
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	var numberFactorial int
	result := 1
	fmt.Println("Enter a number to calculate its factorial:")
	fmt.Scanln(&numberFactorial)

	for i := 1; i <= numberFactorial; i++ {
		result *= i
	}
	fmt.Println(result)
}
