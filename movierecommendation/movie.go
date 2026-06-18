package main

import (
	"fmt"
	"os"
)

func HasE(word string) bool {
	for _, r := range word {
		if r == 'e' || r == 'E' { //finds movies containing the letters
			return true
		}
	}
	return false
}

func main() {
	movies := []string{}                // a slice that stores the movies
	for a := 1; a < len(os.Args); a++ { // for loop that checks the words
		word := os.Args[a]

		if HasE(word) {
			movies = append(movies, word) //add it to the slice
		}

	}
	for _, movie := range movies { // goes through the movies slice and gives the items
		fmt.Println(movie)
	}
}
