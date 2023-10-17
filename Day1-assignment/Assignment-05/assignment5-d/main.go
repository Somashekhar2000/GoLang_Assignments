/*
d.Find the Sum and Average
Description: Create a program that calculates the sum and average of a list of numbers stored
in a slice.
Instructions:
1. Create an integer slice called numbers and initialize it with at least 5 numbers.
2. Calculate and print the sum of all the numbers in the numbers slice.
3. Calculate and print the average of the numbers.
*/
package main

import "assignment5-d/sumaverage"

func main() {
	slice1 := []int{10, 20, 5, 30, 41}      //Create an integer slice called numbers and initialize it with at least 5 numbers.
	sumaverage.SumAndAverageOfSlice(slice1) //calling SumAndAverageOfSlice func of package sumaverage

}
