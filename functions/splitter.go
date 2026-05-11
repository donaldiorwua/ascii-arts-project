package asciiartsproject

import (
	"fmt"
	"strings"
)

func Splitter(data string) []string {
	if len(data) == 0 {
		fmt.Println("Error! Banner file is empty!")
	}
	lines := strings.Split(strings.TrimRight(data, "\n"), "\n")

	return lines
}
