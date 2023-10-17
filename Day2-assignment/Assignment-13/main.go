package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func CountWords(input1 string) int {
	sentences := strings.FieldsFunc(input1, func(r rune) bool { // Split the paragraph into sentences based on ".", "?", and "!"
		return r == '.' || r == '?' || r == '!'
	})

	totalWords := 0 // Initialize variables to keep track of the total number of words and sentences

	for _, sentence := range sentences { // Loop through each sentence
		words := strings.Fields(sentence) // Split each sentence into words based on spaces
		totalWords += len(words)          // Update the total number of words
	}

	return totalWords //return words count
}

func main() {
	fileContent, err := ioutil.ReadFile("sample.txt") //calling ReadFile func that conerts string to []byte
	if err != nil {                                   //check if anu error
		fmt.Println(err) //print error
		return
	}
	// fmt.Println(fileContent)
	paragraph := string(fileContent)      //calling string func to convert []byte to string
	fmt.Println(paragraph)                //printing string
	averageWords := CountWords(paragraph) //calling CountWords func to count words

	fmt.Printf("Total words : %d\n", averageWords) //printing total words
}

func isPunctuation(r rune) bool {
	return r == '.' || r == '?' || r == '!'
}
