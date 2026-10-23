package main

import "fmt"

func main() {
	numbers := []int{4, 2, 4, 1, 2, 4, 3, 2}
	frequency := make(map[int]int)

	for _, num := range numbers {
		frequency[num]++
	}

	for num, count := range frequency {
		fmt.Printf("%d occurs %d times\n", num, count)
	}
}
