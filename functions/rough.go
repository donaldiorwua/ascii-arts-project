func fullwork{
if len(os.Args) < 3 {
		fmt.Println("Make sure there is a banner file and a text to be transformed")
	}

	inputfile := os.Args[1]
	inputText := os.Args[2]

	input := FileLoader(inputfile)
	splitted := Spliter(input)
	mapped := Mapper(splitted)
	rendered := Renderer(inputText, mapped)

	err := os.WriteFile("result.txt", []byte(rendered), 0644)
	if err == nil{
	fmt.Println("Result written to result.txt")
	}
}

func FileLoader(inputfile string) string {
	lines, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return string(lines)
}

func Spliter(lines string) []string {
	words := strings.Split(lines, "\n")

	return words
}

func Mapper(lines []string) map[rune][]string {
	cMap := make(map[rune][]string)

	index := 0
	asciicode := 32 

	for index < len(lines){
		if index+8 > len(lines){
			fmt.Println("Error, Number of lines are not upto 8")
			break
		}
		block := lines[index : index+8]
		cMap[rune(asciicode)] = block
		index += 9
		asciicode++
	}
	return cMap
}

func Renderer(inputText string, cMap map[rune][]string) string {
	words := strings.Split(inputText, "\\n")
	var result []string

	for _, line := range words{
		if line == ""{
			fmt.Println("Empty inputs now allowed, Enter a word to transform to ASCII Arts", " ")
		}
		for i := range 8{
			for _, char := range line{
				if block, ok := cMap[char]; ok{
					fmt.Print(block[i])
					result = append(result, block[i])
				} else {
					fmt.Print(cMap[' '][i])
				}
			}
			fmt.Println()
		}
	}
	return ""

}