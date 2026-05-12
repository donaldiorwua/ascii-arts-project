package main

import "fmt"

func Validate(input string) (rune, error) {
	for _, char := range input {
		if char < ' ' || char > '~' {
			return char, fmt.Println()("Invalid character")
		}
	}
	return 0, nil
}
