package asciiartsproject

import (
	"fmt"
	"strings"
)

func Renderer(text string, cMap map[rune][]string) string {
	if text == "" {
		return ""
	}
	rendered := strings.Split(text, "\\n")

	for index, line := range rendered {
		if line != "" {

			for i := range 8 {
				for _, char := range line {
					if char < 32 || char > 126 {
						fmt.Println("Error! Character is not Ascii Printable character")
						return ""
					}
					if block, ok := cMap[char]; ok {
						fmt.Print(block[i])
					}
				}
				fmt.Println()
			}
		}
		if index != len(rendered)-1 {
			fmt.Println()
		}
	}

	return ""
}
