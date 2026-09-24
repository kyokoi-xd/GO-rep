func reversePrefix(s string, k int) string {
	r := []rune(s)
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}