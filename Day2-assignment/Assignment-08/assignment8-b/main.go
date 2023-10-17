/*b. Recursion
Write a function that calculates the factorial of a given integer. Use recursion for this task.*/

package main

import "fmt"

func main() {
	var num int // declare int variable
	fmt.Println("enter a number to get its factorial")
	fmt.Scan(&num) //taking user input

	result := Factorial(num)                          //calling factorial func
	fmt.Printf("factorial of %d id %d ", num, result) //printing factorial
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1) //recursion call of Factorial func
}
