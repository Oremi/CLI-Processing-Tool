package main

import (
	f "github.com/Oremi/CLI-Processing-Tool/processor"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("USAGE: go run . input_file.txt output_file.txt")
		return
	}

	input_file := os.Args[1]
	output_file := os.Args[2]

	data, err := os.ReadFile(input_file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	result := strings.Fields(string(data))
	result = f.DecimalConversion(result)
	result = f.CaseConversion(result)
	result = f.VowelHandler(result)
	new_result := strings.Join(result, " ")
	new_result = f.PunctuationHandler(new_result)
	new_result = f.QuoteHandler(new_result)
	err = os.WriteFile(output_file, []byte(new_result), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
