package main

import (
	"errors"
	"fmt"
)

func Validate(input map[rune][]string) (error) {
	if input == nil{
		return errors.New("nil map not allowed")
	}
	if len(input) == 0 {
		return errors.New("empty banner file")
	}
	if len(input) != 95 {
		return errors.New("incomplet banner file")
	}

	for char, block := range input {
		if char < 32 || char > 126 {
			return fmt.Errorf("Invalid character: %v", char)
		}
		if len(block) !=8 {
			return errors.New("incomplete block")
		}
		for _, line := range block{
			if line == ""{
				return errors.New("empty lines not allowed")
			}
		}
	}
		
	return nil
}
