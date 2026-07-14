package main

import (
	"errors"
	"os"
	"strings"
)

// LoadBanner function takes as input file name as string, read the file and load it into memory
// Split the content of the file on new line
// store the splited content in chunks of 8 lines and a seperator in a map 
// returns the map as rune slice of string
func LoadBanner(inputfile string) (map[rune][]string, error) {
	cMap := make(map[rune][]string)
	
	// read file content
	data, err := os.ReadFile(inputfile)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty banner file")
	}

	// split the read content by new line
	lines := strings.Split(string(data), "\n")

	if len(lines) < 856 {
		return nil, errors.New("incomplete banner file")
	}

	// make chunks of 8 line and save them with a new line in a map 
	for char := ' '; char <= '~'; char++{
		start := (int(char) - 32) * 9
		cMap[char] = lines[start+1 : start+9]
	}
	return cMap, nil

}
