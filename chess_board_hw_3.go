package main

import (
	"fmt"
	"strings"
)

func main() {
	var oneNumber int
	_, err := fmt.Scan(&oneNumber)
	if err != nil {
		oneNumber = 8 // default value if the number is not provided
	}

	secondNumber := oneNumber
	for i := 0; i < secondNumber; i++ {
		if i%2 == 0 {
			fmt.Printf(strings.Repeat(" #", secondNumber-1))
			fmt.Printf("\n")
		} else {
			fmt.Printf(strings.Repeat("# ", secondNumber))
			fmt.Printf("\n")
		}
	}
}
