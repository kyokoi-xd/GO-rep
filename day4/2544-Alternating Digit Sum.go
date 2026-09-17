func alternateDigitSum(n int) int {
	if n < 10 {
		return n
	}
	var digits []int
	for n > 0 {
		digit := n % 10
		digits = append([]int{digit}, digits...)
		n = n / 10
	}
	var sum int
	for i, digit := range digits {
		if i%2 == 0 {
			sum += digit
		} else {
			sum -= digit
		}
	}
	return sum
}