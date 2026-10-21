import "strings"

func findOcurrences(text string, first string, second string) []string {
	var third []string
	words := strings.Split(text, " ")

	for i := 0; i < len(words)-2; i++ {
		if words[i] == first && words[i+1] == second {
			third = append(third, words[i+2])
		}
	}
	return third
}