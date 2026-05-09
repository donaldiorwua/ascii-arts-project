package asciiartsproject

import (
	"fmt"
	"os"
	"strings"
)

func Improved()  {
	if len(os.Args) < 2 {
		fmt.Println("Error: make sure your arguments contain a file name and text to transform")
		return
	}

	inputfile := os.Args[1]
	inputText := os.Args[2]

	if !strings.HasSuffix(inputfile, ".txt"){
		inputfile = inputfile + ".txt"
	}
	//file loader
	//Opens and read file content into memory
	input, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// splits the content of the read file into lines
	if len(input) == 0 {
		fmt.Println("Error: Banner file is Empty!")
		return
	}
	content := strings.Split(string(input), "\n")
	
	// Mapping cahracters t their lines
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

	// Rendering the string chacters into their ascii arts equivelence
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