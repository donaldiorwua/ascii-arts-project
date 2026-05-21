package asciiartsproject

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArts() {
	inputfile := "standard.txt"
	inputtext := os.Args[1]

	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(data) == 0 {
		fmt.Println("empty banner file")
		return
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 856 {
		fmt.Println("incomplete banner file")
		return
	}

	inputtext = strings.ReplaceAll(inputtext, "\n", "\n")
	inputtext = strings.ReplaceAll(inputtext, "\r", "\n")

	if inputtext == "" {
		return
	}
	if inputtext == "\\n" {
		fmt.Print("\n")
		return
	}
	text := strings.Split(inputtext, "\n")
	for _, word := range text {
		if word == "" {
			fmt.Print("\n")
			continue
		}
		for char := 0; char < 8; char++ {
			for _, ch := range word {
				start := (int(ch)-32)*9 + 1
				block := lines[start+char]
				fmt.Print(block)
			}
			fmt.Println()
		}
	}
}
