func shuffle(nums []int, n int) []int {
	res := make([]int, 0, len(nums))
	left, right := 0, n
	for left < n {
		res = append(res, nums[left])
		res = append(res, nums[right])
		left++
		right++
	}
	return res
}