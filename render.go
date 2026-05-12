package main

func RenderLine(text string, cMap map[rune][]string) []string {
	result := make([]string, 8)

	for _, char := range text {
		for j := 0; j < 8; j++ {
			result[j] += cMap[char][j]
		}
	}
	return result
}
