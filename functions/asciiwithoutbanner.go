package asciiartsproject

import (
	"fmt"
	"strings"
)

func StringToArt(input string) string {

	digits := map[rune][]string{
		'0': {
			" ___",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'1': {
			" /|  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			" ___",
			"|   |",
			"   / ",
			" /   ",
			"|___ ",
		},
		'3': {
			" ___ ",
			"   / ",
			"  /  ",
			"  |  ",
			" ___|",
		},
		'4': {
			"     ",
			"|    ",
			"|    ",
			"|_|_ ",
			"  |  ",
		},
		'5': {
			" ___ ",
			"|    ",
			"| __ ",
			"    |",
			" ___|",
		},
		'6': {
			" ___ ",
			"|    ",
			"|___ ",
			"|   |",
			"|___|",
		},
		'7': {
			" ___",
			"|   |",
			"    |",
			"    |",
			"    |",
		},
		'8': {
			" ___",
			"|   |",
			"|___|",
			"|   |",
			"|___|",
		},
		'9': {
			" ___",
			"|   |",
			"|___|",
			"    |",
			" ___|",
		},
	}
	result := []string{}

	for _, line := range digits {
		result = append(result, line...)
	}

	word := strings.Split(input, "\n")
	var output []string
	block := make([]string, 5)
	for _, char := range word {
		for _, ch := range char {
			pattern := digits[ch]
			if ch < '0' || ch > '9' {
				fmt.Print("invalid charactrer: ")
				return char
			}
			for i := 0; i < 5; i++ {
				block[i] += pattern[i]
			}
		}
		output = append(output, block...)
	}
	return strings.Join(output, "\n")
}
