package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error! Enter a file to read!")
	}

	inputfile := os.Args[1]

		//opening the text file
	input, err := os.OpenFile(inputfile, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	}
	//Automatic file closing function
	defer input.Close()

	content, err := io.ReadAll(input)
	if err != nil {
		fmt.Println(err)
	}

	splited := strings.Split(string(content), "\n")
	fmt.Println(splited)

}
