package asciiart

import (
	"strings"
)

// function to build ascii arts, taking inputs from the result of reading the banner file
// and user inputs from standard input
func BuildArt(str string, text []string) (string, error) {
	// validate user inputs to cleean them
	validated, err := ValidateInput(str)
	if err != nil {
		return "", err
	}

	// split the user input by new line
	splistr := Splitting(validated)
	var words strings.Builder

	// Iterate through the splited inputs to handle new line characters
	for _, char := range splistr {
		if char == "" {
			words.WriteString("\n")
			continue
		}
		// build the ascii using the runed characters from user inputs and ascii lines
		// from the banner file
		for i := range 8 {
			for _, cha := range char {
				stat := (int(cha)-32)*9 + 1
				word := text[stat : stat+8]
				words.WriteString(word[i])

			}
			words.WriteString("\n")
		}
	}
	return words.String(), nil
}
