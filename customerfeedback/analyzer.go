package main

import (
	"fmt"

	"os"
)

func HasA(word string) bool {
	for _, r := range word {
		if r == 'a' || r == 'A' {
			return true
		}
	}
	return false
}

func main() {
	matches := []string{}

	for i := 1; i < len(os.Args); i++ {
		word := os.Args[i]

		if HasA(word) {
			matches = append(matches, word)
		}
	}
	for _, word := range matches {
		fmt.Println(word)
	}
	fmt.Println(matches)

}
