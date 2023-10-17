/*a. Create a package named calculator.
Inside the calculator package, define a below functions that
1. Takes an integer number & calculates the square of a number & returns it.
2. Takes two integers as input and calculates their sum & return it,
3.Takes two integers as input and calculates their difference & return it,
4 .Takes two integers as input from the user and calculates their product & return it,
5. Takes two integers as input and calculates their quotient & remainder & return both
values.
Import calculator package in your main file & Call the functions from your main function using
some dummy values & print the results.
*/

package main

import (
	"assignment3-a/calculator"
	"fmt"
)

func main() {
	num1 := 4 // declaring int variable num1 using short declation

	fmt.Printf("printing square of a integer %d : %d\n", num1, calculator.SquaringInteger(num1)) //printing the sqaure of num1 by calling SquaringInteger func of calculator package

	num1 = 10  // reinitializing int variable num1
	num2 := 20 // declaring int variable num2 using short declation

	fmt.Printf("printing sum of 2 integer  %d,%d is : %d\n", num1, num2, calculator.Summing2Integers(num1, num2)) // printing the sum of num1 & num2 by calling Summing2Integers func of calculator package

	num1, num2 = 20, 10                                                                                                             // reinitializing int variable num1 & num2
	fmt.Printf("printing difference of 2 integer  %d,%d is : %d\n", num1, num2, calculator.DiffereneceBetween2Integers(num1, num2)) // printing the differnce between num1 & num2 by calling DiffereneceBetween2Integers func of calculator package

	fmt.Printf("printing product of 2 integer's %d,%d is : %d\n", num1, num2, calculator.ProductOf2Integers(num1, num2)) //printing the product of num1 & num2 by calling ProductOf2Integers func of calculator package

	num1, num2 = 20, 4                                                                                                // reinitializing int variable num1 & num2
	quotient, remainder := calculator.DivisionFunction(num1, num2)                                                    //calling DivisionFunction func of calculator package to get quotient and remainder
	fmt.Printf("printing quotient and remainder of 2 integer's %d,%d is : %d, %d\n", num1, num2, quotient, remainder) //printing Quotient and remainder of num1 & num2

}
