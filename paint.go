package main

import "fmt"

func main() {
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	total += amount
	fmt.Printf("%0.2f liters needed\n", amount)
	amount = paintNeeded(5.2, 3.5)
	total += amount
	fmt.Printf("%0.2f liters needed\n", amount)
	fmt.Printf("Total: %0.2f liters needed\n", total)
}

func paintNeeded(width, height float64) float64 {
	area := width * height
	paint := area / 10.0
	return paint
}
