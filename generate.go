package main

import "strings"

func GenerateArts(text string, cMap map[rune][]string) string {
	var result strings.Builder
	if text == "" {
		return ""
	}

	if text == "\\n" {
		return "\n"
	}

	words := SplitInput(text)

	for index, lines := range words {
		if lines == "" {
			if index != len(words)-1 {
				result.WriteString("\n")
			}
			continue
		}
		for j := 0; j < 8; j++ {
			for _, char := range lines {
				result.WriteString(cMap[char][j])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
