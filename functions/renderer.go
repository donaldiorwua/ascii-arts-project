package asciiartsproject

import (
	"fmt"
	"strings"
)

func Renderer(text string, cMap map[rune][]string) string {
	rendered := strings.Split(text, "\\n")

	for _, line := range rendered {
		if line == "" {
			fmt.Println("Error! Empty line detected!")
			continue
		}
		for i := range 8 {
			for _, char := range line {
				if block, ok := cMap[char]; ok {
					fmt.Print(block[i])
				}else if char == ' ' {
					fmt.Print(cMap[' '][i])
				}
			}
			fmt.Println()
		}
	}

	return strings.Join(rendered, "\n")
}
