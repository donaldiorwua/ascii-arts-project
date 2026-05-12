package main

import "strings"

func SplitInput(inputText string) []string {
	lines := strings.Split(inputText, "\\n")

	return lines
}
