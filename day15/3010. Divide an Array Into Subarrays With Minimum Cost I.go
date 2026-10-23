func minimumCost(nums []int) int {
	minFirst, minSecond := 51, 51

	for i := 1; i < len(nums); i++ {
		if nums[i] < minFirst {
			minSecond = minFirst
			minFirst = nums[i]
		} else if nums[i] < minSecond {
			minSecond = nums[i]
		}
	}
	sum := nums[0] + minFirst + minSecond
	return sum
}