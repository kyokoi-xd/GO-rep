func numOfUnplacedFruits(fruits []int, baskets []int) int {
	var unplaced int
	flag := false
	for i := 0; i < len(fruits); i++ {
		for j := 0; j < len(baskets); j++ {
			if fruits[i] <= baskets[j] {
				flag = false
				baskets[j] = 0
				break
			}
			flag = true
		}
		if flag {
			unplaced++
		}
	}
	return unplaced
}