/*b. Create a package called temperature.
Write a program/function which accepts temperature in Fahrenheit & converts Fahrenheit to
Celsius.
The Fahrenheit to Celsius formula is expressed as °C = (°F - 32) × 5/9; in which the value of F
(Fahrenheit) is substituted and we get the value in Celsius. This formula converts the given
temperature value from Fahrenheit to Celsius.*/

package main

import (
	"assassignment3-b/temperature"
	"fmt"
	"strconv"
)

func main() {

	var fahrenheit int //declaring var int fahrenheit to hold input from user

	for { //for recursive input from the user
		var input string // var string to hold the console input
		fmt.Println("Enter an integer the temperature value from Fahrenheit to Celsius.: ")
		_, err := fmt.Scanln(&input) //holding error from scanner
		if err != nil {              // checking for error from input
			fmt.Println("Error:", err) //printing error from input
			continue                   // restart for loop
		}

		num, parseErr := strconv.Atoi(input) //type casting value from string to int datatype
		if parseErr != nil {                 //checking for error while type casting
			fmt.Println("Error: Please enter a valid integer.")
			continue // restart for loop
		}

		fahrenheit = num //assigning input to fahrenheit
		break
	}

	fmt.Printf("%d°F is equal to %f°C", fahrenheit, temperature.FahrenheitToCelsiusConvertion(fahrenheit)) //printing the coverted temeprature from fahrenheit to Celsius by calling FahrenheitToCelsiusConvertion func of temperature package

}
