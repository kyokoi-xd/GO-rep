import "fmt"

func findComplement(num int) int {
	var binary string
	for num > 0 {
		remainder := num % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		num /= 2
	}
	var digits []int
	for _, value := range binary {
		if value == '1' {
			digits = append(digits, 0)
		} else {
			digits = append(digits, 1)
		}
	}

	var decimal int
	weight := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 1 {
			decimal += weight
		}
		weight *= 2
	}
	return decimal
}