package main

import "fmt"

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(containsDuplicate(nums))
	fmt.Print(firstDuplicate(nums))
}

func firstDuplicate(nums []int) (int, bool) {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return num, true
		}
		seen[num] = true
	}
	return 0, false
}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}
