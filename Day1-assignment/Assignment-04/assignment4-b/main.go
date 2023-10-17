/*b. Create a simple guessing game.
. Generate a random number between 1 and 10, and let the user guess the number.
(use rand.Intn() method to generate random numbers)
Example on how to use rand package to generate random number:
https://gobyexample.com/random-numbers
Ask the user for the inputs and Provide hints to the user like "too high" or "too low" until they
guess the correct number. If the user guesses the number correctly say “you guessed it right” &
quit.*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	randomNum := rand.Intn(10) //Generateing a random number between 1 and 10

	for { //for recursive input from the user
		var number string // var string to hold the console input
		fmt.Println("Enter number from 1 to 10 :")
		_, err := fmt.Scanln(&number) //holding error from scanner
		if err != nil {               // checking for error from input
			fmt.Println("Error :", err) //printing error from input
			continue                    // restart for loop
		}

		num, err := strconv.Atoi(number) //type casting value from string to int datatype
		if err != nil {                  //checking for error while type casting
			fmt.Println("Enter proper integer number from 1 to 10 : ")
			continue // restart for loop
		}

		if num >= 1 && num <= 10 { //condition check of user input in range 1 to 10
			if num < randomNum { // checking if input is less than the random number
				fmt.Println("too low!.....Try again")
				continue // restart for loop
			}
			if num > randomNum { // checking if input is greater than the random number
				fmt.Println("too high!.....Try again")
				continue // restart for loop
			}
			fmt.Println("you guessed it right") //printing if the guess if correct ie. input is equal to randomnumber
			break
		}
		fmt.Println("Enter number in range 1 to 10 only")
	}
}
