package asciiartsproject

import (
	"fmt"
	"os"
)


func FileLoader(inputfile string) string {
	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}