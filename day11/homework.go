package main

import "fmt"

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	for i, value := range numbers {
		numbers[i] = value * 3
	}
	fmt.Println(numbers)

	numbers = []int{12, 7, 25, 4, 18, 10, 7}
	min := numbers[0]
	max := numbers[0]
	counteven := 0
	sum := 0
	for _, value := range numbers {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
		if value%2 == 0 {
			counteven++
		}
		sum += value
	}
	average := float64(sum) / float64(len(numbers))
	fmt.Printf("Numbers: %v\nMin: %d\nMax: %d\nEven count: %d\nSum: %d\nAverage: %.2f\n", numbers, min, max, counteven, sum, average)
}
