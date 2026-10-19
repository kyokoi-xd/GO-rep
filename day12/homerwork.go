package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)

	res := countRunes(s)
	fmt.Println(res)
}

func countRunes(s string) int {
	count := 0

	for range s {
		count++
	}
	return count
}
