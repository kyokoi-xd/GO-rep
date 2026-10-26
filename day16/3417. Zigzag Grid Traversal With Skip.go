func zigzagTraversal(grid [][]int) []int {
	var result []int

	for indexRow := 0; indexRow < len(grid); indexRow++ {
		if indexRow%2 == 0 {
			for indexNum := 0; indexNum < len(grid[indexRow]); indexNum++ {
				temp := grid[indexRow]
				if indexNum%2 == 0 {
					result = append(result, temp[indexNum])
				}
			}
		} else {
			for indexNum := len(grid[indexRow]) - 1; indexNum > 0; indexNum-- {
				temp := grid[indexRow]
				if indexNum%2 != 0 {
					result = append(result, temp[indexNum])
				}
			}
		}
	}

	return result
}