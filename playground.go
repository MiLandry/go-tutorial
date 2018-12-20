package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func LongestWord(sen string) string {

	// Make a Regex to say we only want
	reg, err := regexp.Compile("[^a-zA-Z ]+")
	if err != nil {
		log.Fatal(err)
	}

	//filter out punctuation
	filtered := reg.ReplaceAllString(sen, "")
	// fmt.Println("filtered", filtered, "\n")

	//get words into array
	words := strings.Split(filtered, " ")
	// fmt.Println("words", words, "\n")

	var result string
	// return biggest word
	for _, value := range words {
		// fmt.Println(value, i)
		if len(value) > len(result) {
			result = value
		}

	}

	// code goes here
	// Note: feel free to modify the return type of this function
	return result

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	// fmt.Println(LongestWord(readline()))
	fmt.Println(LongestWord("Argument goes here"))

}
