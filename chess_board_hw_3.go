package main

import (
	"fmt"
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
			if i != secondNumber {
				fmt.Printf("#")
			} else {
				fmt.Printf("\n")
			}
		}
		oneNumber -= 1
	}
}
