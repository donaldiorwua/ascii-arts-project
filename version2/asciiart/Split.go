package asciiart

import "strings"

//Function to split user inputs
//Takes user input as strings and returns a slice of strings
func Splitting (input string) []string {
	if input == "" {
		return []string{}
	}

	lines := strings.Split(input, `\n`)
	return lines
}