/*Assignment 10: Function with Variadic Parameters
Description: Create a program that defines a function with variadic parameters and uses it.
Instructions:
1. Create a function called sum that takes an arbitrary number of integers and returns their
sum.
2. Use this function to calculate and print the sum of different sets of numbers.
Note: Example:https://gobyexample.com/variadic-functions*/

package main

import "fmt"

func sum(nums ...int) int { //defining sum func that accepts variadic int parameter
	sum := 0 //initializing sum to 0
	for _, value := range nums {
		sum += value //adding values
	}

	return sum //returning sum
}

func main() {
	nums := []int{10, 20, 30, 40, 50} //creating nums slice with values
	fmt.Println(sum(nums...))         //calling sum func to print sum of elements in slice
}
