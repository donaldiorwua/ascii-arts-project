package main

import (
	"fmt"
	"os"
)

func main() {
	//asciiartsproject.Improved()
	if len(os.Args) != 2 {
		fmt.Println("Error! Usage: go run . \"Hello\"")
		return
	}
	inputfile := "standard.txt"
	inputText := os.Args[1]

	text, err := LoadBanner(inputfile)
	if err != nil {
		return
	}

	result := GenerateArt(inputText, text)

	fmt.Print(result)

}
