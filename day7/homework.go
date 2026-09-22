package main

import "fmt"

func main() {
	var n int
	var sum int
	var even int

	fmt.Println("Enter n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(i)
		sum += i
		if i%2 == 0 {
			even++
		}
	}
	fmt.Println("Sum:", sum)
	fmt.Println("Even numbers:", even)
}
