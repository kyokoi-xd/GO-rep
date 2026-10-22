func mapWordWeights(words []string, weights []int) string {
	var res string
	for _, word := range words {
		sum := 0
		for _, char := range word {
			indexChar := int(char - 'a')
			sum += weights[indexChar]
		}
		res += string(rune('a' + 25 - (sum % 26)))
	}
	return res
}