package asciiartsproject

import "strings"

func Splitter(data string) []string {
	lines := strings.Split(data, "\n")

	return lines
}