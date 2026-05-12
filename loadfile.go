package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(inputfile string) (map[rune][]string, error) {
	cMap := make(map[rune][]string)
	index := 0
	asciicode := 32

	data, err := os.ReadFile(inputfile)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return cMap, errors.New("Empty banner file")
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) != 856 {
		return cMap, errors.New("Incomplete banner file")
	}

	for index+8 < len(lines) {
		if index+8 > len(lines) {
			break
		}
		block := lines[index : index+8]
		cMap[rune(asciicode)] = block
		index += 9
		asciicode++
	}
	return cMap, nil

}
