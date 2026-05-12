package main

import "fmt"

func Validate(input string) (rune, error) {
	for _, char := range input {
		if char < ' ' || char > '~' {
			return char, fmt.Errorf("Invalid character: %v", char)
		}
	}
	return 0, nil
}
