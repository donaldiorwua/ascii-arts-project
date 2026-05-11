package main

import (
	asciiartsproject "ascii-asrts-project/functions"
	"fmt"
	"os"
)

func main() {
	//asciiartsproject.Improved()
	if len(os.Args) < 2 {
		fmt.Println("Error! Enter a file to read!")
	}
	inputfile := "standard.txt"
	inputText := os.Args[1]

	text := asciiartsproject.FileLoader(inputfile)
	lines := asciiartsproject.Splitter(text)
	cMap := asciiartsproject.CharMap(lines)
	result := asciiartsproject.Renderer(inputText, cMap)

	fmt.Print(result)

}
