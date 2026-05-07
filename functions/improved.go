package asciiartsproject

import (
	"fmt"
	"os"
	"strings"
)

func Improved()  {
	if len(os.Args) < 2 {
		fmt.Println("make sure your arguments contain a file name and text to transform")
	}

	inputfile := os.Args[1]
	inputText := os.Args[2]

	if !strings.HasSuffix(inputfile, ".txt"){
		inputfile = inputfile + ".txt"
	}

	input, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	content := strings.Split(string(input), "\n")

	cMap := make(map[rune][]string)

	index := 0
	asciicode := 32

	for index+8 < len(content){
		if index+8 > len(content){
			break
		}
		charHeight := content[index : index+8]
		cMap[rune(asciicode)] = charHeight
		index += 9
		asciicode++
	}

	if inputText == ""{
		fmt.Println("Empty words are not allowed!\nEnter a word!")
		return
	}	
	words := strings.Split(inputText, "\\n")

	for _, line := range words{
		if line == ""{
			continue
		}
		for i := range 8{
			for _, char := range line{
				if char < 32 || char > 126 {
					fmt.Println("Error: Invalid character! Enter ASCII pritable characters")
					return
				}
				if block, ok := cMap[char]; ok{
					fmt.Print(block[i])
				}else{
					fmt.Print(cMap[' '])
				}
			}
			fmt.Println()
		}
	}

}