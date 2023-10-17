/*
f.Search for an Element
Description: Create a program that searches for a specific number in a slice and prints whether
it's found or not.
Instructions:
1. Create an integer slice called numbers with at least 5 numbers.
2. Choose a number to search for.(take a number input from user using command line)
3. Write code to search for that number in the numbers slice.
4. Print whether the number was found or not.
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	numbers := []int{} //Creating an integer slice called numbers
	var searchNum int

	for i := 0; i < 5; i++ {
		numbers = append(numbers, rand.Intn(11)) //numbers with randon values integer
	}

	fmt.Println(numbers) //printing the numbers slice

	for { //for recursive input from the user
		var number string // var string to hold the console input
		fmt.Println("Enter an integer with in 1 to 11:")
		_, err := fmt.Scanln(&number) //holding error from scanner
		if err != nil {               // checking for error from input
			fmt.Println("Error :", err) //printing error from input
			continue                    // restart for loop
		}

		num, err := strconv.Atoi(number) //type casting value from string to int datatype
		if err != nil {                  //checking for error while type casting
			fmt.Println("Invalid input enter proper Integer ")
			continue // restart for loop
		}

		if num < 11 { //check if num input is less than 11
			searchNum = num
			break
		}
	}

	if !Searching(numbers, searchNum) { //calling Searching func and checking the conditon if number found or not
		fmt.Printf("%d not found", searchNum)
		return
	}

	fmt.Printf("%d  found", searchNum)

}

func Searching(s []int, num int) bool { //defining Searching func
	for _, value := range s {
		if value == num { //conditon if value in sclice is equal to given value
			return true //return true if found
		}

	}
	return false //return false if not found
}
