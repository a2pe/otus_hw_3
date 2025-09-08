package main

import (
	"fmt"
	"strings"
)

// version 1
//func main() {
//	var oneNumber int
//	_, err := fmt.Scan(&oneNumber)
//	if err != nil {
//		oneNumber = 8 // default value
//	}
//
//	secondNumber := oneNumber
//	for i := 0; i < secondNumber; i++ {
//		if i%2 == 0 {
//			fmt.Printf(strings.Repeat("# ", secondNumber/2+1))
//			fmt.Printf("\n")
//		} else {
//			fmt.Printf(strings.Repeat(" #", secondNumber/2))
//			fmt.Printf("\n")
//		}
//	}
//}

func main() {
	var oneNumber int
	_, err := fmt.Scan(&oneNumber)
	if err != nil {
		oneNumber = 8 // default value
	}

	secondNumber := oneNumber
	for i := 0; i < secondNumber; i++ {
		if oneNumber%2 != 0 {
			if i%2 == 0 {
				fmt.Printf(strings.Repeat("# ", secondNumber/2+1))
				fmt.Printf("\n")
			} else {
				fmt.Printf(strings.Repeat(" #", secondNumber/2))
				//fmt.Printf(strings.Repeat("# ", secondNumber/2)) // для обычного прямоугольника
				fmt.Printf("\n")
			}
		} else {
			if i%2 == 0 {
				fmt.Printf(strings.Repeat("# ", secondNumber/2))
				fmt.Printf("\n")
			} else {
				fmt.Printf(strings.Repeat(" #", secondNumber/2))
				fmt.Printf("\n")
			}
		}
	}
}
