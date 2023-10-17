/*e. Find the Maximum Element
Description: Write a program that finds and prints the maximum element in a slice of integers.
Instructions:
1. Create an integer slice called numbers with at least 5 different numbers in random
order..
2. Find and print the maximum element in the numbers slice.*/

package main

import (
	"fmt"
	"math/rand"
)

func MaxNum(s []int) int {
	var maxNum int
	for i := range s {
		if s[i] > maxNum { //check if the greater than the next number
			maxNum = s[i] //assigning the index element to maxNum number
		}
	}
	return maxNum //returning the assigned maxNum
}

func main() {
	numbers := []int{} ///Creating an integer slice called numbers
	for i := 0; i < 5; i++ {
		numbers = append(numbers, rand.Intn(50)) //slice numbers with at least 5 different numbers in random order
	}
	fmt.Println(numbers)                                                       //printing the numbers sclice with random order
	fmt.Printf("maximum element in the numbers slice is %d	", MaxNum(numbers)) //printing the maximum number by calling MaxNum func
}
