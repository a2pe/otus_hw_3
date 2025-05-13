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
	for oneNumber > 0 {
		for i := 0; i <= secondNumber; i++ {
			if i == 0 && oneNumber == secondNumber {
				fmt.Printf(strings.Repeat(" ___", oneNumber))
				fmt.Printf("\n")
			}
			if i != secondNumber {
				fmt.Printf("|___")
			} else {
				fmt.Printf("|\n")
			}
		}
		oneNumber -= 1
	}
}
