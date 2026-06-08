package main

import "strings"


// splitinput functino splits the user inputs on new line character 
// to be used by the ASCII art gernarator function
// it returns a slice of strings 
func SplitInput(inputText string) []string {
	lines := strings.Split(inputText, "\\n")

	return lines
}
