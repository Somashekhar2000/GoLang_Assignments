/*a. Implement a method called checkEvenOdd , it takes int as an input.
check if it is even or odd.
If it is even then print "Number is even",
else print "Number is odd"
*/

package main

import (
	"fmt"
	"strconv"
)

func checkEvenOdd(num int) { //defining checkEvenOdd func
	if num%2 == 0 { //condition to check if num is even of odd
		fmt.Printf("%d is Even number", num) //printing if num even
		return
	}
	fmt.Printf("%d is Odd number", num) //printing if num odd
}

func main() {

	for { //for recursive input from the user
		var number string // var string to hold the console input
		fmt.Println("Enter an integer :")
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

		checkEvenOdd(num) //calling checkEvenOdd func
		break             // out of for loop
	}

}
