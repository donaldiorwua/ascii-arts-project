package asciiartsproject

import "strings"


func CharMap(words []string) map[rune][]string {
	cMap := make(map[rune][]string)

	index := 0
	asciicode := 32
	joined := strings.Join(words, " ")
	lines := Splitter(joined)

	for index < len(lines) {
		if index+8 > len(lines) {
			break
		}
		block := lines[index : index+8]
		cMap[rune(asciicode)] = block
		index += 8
		asciicode++
		index += 9
	}
	return cMap
}