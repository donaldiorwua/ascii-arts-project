package main

import "strings"

// Function to generate ASCII arts takes as inut text string from the standard input and 
// Map of rune slice of strings made from the ASCII banner file
// Prints ASCII arts at standard output mirrowing the inputed text
func GenerateArt(text string, cMap map[rune][]string) string {
	var result strings.Builder

	// check for empty user input
	if text == "" {
		return ""
	}

	if text == "\\n" {
		return "\n"
	}

	words := SplitInput(text)

	// handle new line in user input
	for index, lines := range words {
		if lines == "" {
			if index != len(lines)-1 {
				result.WriteString("\n")
			}
			continue
		}

		// generate ASCII arts
		for j := 0; j < 8; j++ {
			for _, char := range lines {
				result.WriteString(cMap[char][j])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
