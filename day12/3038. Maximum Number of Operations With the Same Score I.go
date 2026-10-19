func maxOperations(nums []int) int {
	var count int
	score := nums[0] + nums[1]
	for i := 0; i+1 < len(nums); i += 2 {
		a, b := nums[i], nums[i+1]
		sum := a + b
		if sum != score {
			return count
		}
		count++
	}
	return count
}