package main

import (
	asciiartsproject "ascii-asrts-project/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error! Enter a file to read!")
	}

	inputfile := os.Args[1]
	inputText := os.Args[2]

	text := asciiartsproject.FileLoader(inputfile)
	lines := asciiartsproject.Splitter(text)
	cMap := asciiartsproject.CharMap(lines)
			fmt.Println(cMap)

	result := asciiartsproject.Renderer(inputText, cMap)

	fmt.Println(result)

}
