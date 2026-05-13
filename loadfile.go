package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(inputfile string) (map[rune][]string, error) {
	cMap := make(map[rune][]string)
	
	data, err := os.ReadFile(inputfile)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return cMap, errors.New("empty banner file")
	}
	lines := strings.Split(string(data), "\n")

	if len(lines) < 856 {
		return cMap, errors.New("incomplete banner file")
	}

	for char := ' '; char <= '~'; char++{
		start := (int(char) - 32) * 9
		cMap[char] = lines[start+1 : start+9]
	}
	return cMap, nil

}
