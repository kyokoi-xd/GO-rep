package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	numberAnalyzer(n)
}

func numberAnalyzer(n int) {
	evenCount := 0
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			evenCount++
		}
		sum += i
		fmt.Println(i)
	}
	fmt.Printf("Even numbers: %d\n", evenCount)
	fmt.Printf("Sum: %d\n", sum)
}
