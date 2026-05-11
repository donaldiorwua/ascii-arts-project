package asciiartsproject

func CharMap(lines []string) map[rune][]string {
	cMap := make(map[rune][]string)

	index := 0
	asciicode := 32

	for index < len(lines) {
		if index+8 > len(lines) {
			break
		}
		block := lines[index : index+8]
		cMap[rune(asciicode)] = block
		index += 9
		asciicode++
	}
	return cMap
}
