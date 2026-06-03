package asciiartsproject

import "strings"

func Converter(input string) string {
	if input == "" {
		return ""
	}
	if input == "\n" {
		return "\n"
	}
	digits := map[rune][]string{
		'0': {
			" ___ ",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			"____ ",
			"    |",
			" ___|",
			"|    ",
			"|___ ",
		},
		'3': {
			" ___ ",
			"    |",
			" ___|",
			"    |",
			" ___|",
		},
		'4': {
			"|    ",
			"|    ",
			"|_|__",
			"  |  ",
			"     ",
		},
		'5': {
			" ___ ",
			"|    ",
			"|___ ",
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
			" ___ ",
			"|   |",
			"    |",
			"    |",
			"    |",
		},
		'8': {
			" ___ ",
			"|   |",
			"|___|",
			"|   |",
			"|___|",
		},
		'9': {
			" ___ ",
			"|   |",
			"|___|",
			"    |",
			" ___|",
		},
	}

	words := strings.Split(input, "\n")
	result := []string{}

	for _, word := range words {

		for _, ch := range word {
			block := make([]string, 5)

			pattern, ok := digits[ch]
			if !ok {
				return "invalid digit"
			}

			for i := range 5 {
				block[i] += pattern[i]
			}
			result = append(result, block...)
		}

	}
	return strings.Join(result, "\n") + "\n"
}
