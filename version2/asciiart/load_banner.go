package asciiart

import (
	"errors"
	"os"
	"strings"
)

//function to read the inputs from banner file 
// Takes inputfile as string and returns the data as slice of strings
func LoadBanner(inputfile string)([]string, error) {
	if inputfile == ""{
		return nil, errors.New("nil file")
	}

	data, err := os.ReadFile(inputfile)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("empty banner file")
	}
	lines := strings.Split(string(data), "\n")

	if len(lines) < 856 {
		return nil, errors.New("incomplete banner file")
	}


	return lines, nil

}
