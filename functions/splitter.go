package asciiartsproject

import (
	"fmt"
	"strings"
)

func Splitter(data string) []string {
	if len(data) == 0 {
		fmt.Println("Error! Banner file is empty!")
		return nil
	}
	lines := strings.Split(data, "\n")
	if lines[0] == "" {
		lines = lines[1:]
	}

	return lines
}
