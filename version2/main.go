package main

import (
	"ascii-art/asciiart"
	"fmt"
	"os"
)

func main() {
	var inputtext string
	var inputfile string

	switch {
	case len(os.Args) < 2:
		fmt.Println(`ensure you enter: go run . "Hello"`)
		return
	case len(os.Args) == 2:

		inputtext = os.Args[1]
		inputfile = "standard"

	case len(os.Args) == 3:

		inputtext = os.Args[1]
		inputfile = os.Args[2]

	case len(os.Args) > 3:
		fmt.Println("Argument should not be more than 3")
		return
	}

	data, err := asciiart.LoadBanner(inputfile + ".txt")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		result, err := asciiart.BuildArt(inputtext, data)
		if err != nil {
			fmt.Println("Error!", err)
			return
		}

		fmt.Print(result)
}
