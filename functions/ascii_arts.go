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
	
	if len(lines) < 855 {
		fmt.Println("incomplete banner file")
		return
	}

	
	if inputtext == "" {
		return
	}
	
	text := strings.Split(inputtext, `\n`)
	for i, word := range text {
		if word == "" {
			if i == 0 || text[i-1] != ""{
				fmt.Println()
			}
		
			continue
		}
		for char := range 8{
			for _, ch := range word {
				start := (int(ch)-32)* 9 +1
				block := lines[start + char]
				fmt.Print(block)
			}
			fmt.Println()
		}
	}
}