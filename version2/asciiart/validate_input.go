package asciiart

import (
	"errors"
	"fmt"
	"strings"
)

// Function to validate user inputs to ensure only allowed inputs are supplied
// Takes user inputs as strings and returns strings
func ValidateInput(input string) (string, error) {
	input = strings.ReplaceAll(input, "\n", "\n")
	if input == "" {
		return "", nil
	}
	if input == "\\n" {
		fmt.Print("\n")
		return "", nil
	}
	for _, word := range input {
		if word < 32 || word > 126 {
			return "", errors.New("invalid character")
		}
	}
	return input, nil
}
