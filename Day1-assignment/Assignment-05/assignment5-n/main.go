/*n.Find the Index of an Element
Description: Write a program that finds and prints the index of a specific number in a slice.
Instructions:
1. Create an integer slice called numbers with at least 6 different numbers.
2. Choose a number to search for in the numbers slice.
3. Write code to find the index of the chosen number.
4. Print the index (or a message if the number is not found).
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := []int{} //creating a slice numbers

	for i := 0; i < 6; i++ {
		numbers = append(numbers, rand.Intn(50)) //appending 5 random numbers to slice
	}

	fmt.Println("slice :", numbers) //printing slice

	fmt.Println("enter the number from the above slice")
	var num int
	fmt.Scan(&num)                                     //taking input value from the user
	index := FindingIndexOfNumberInSlice(numbers, num) //calling FindingIndexOfNumberInSlice func to get index

	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", num, index) //print index if found
	} else {
		fmt.Printf("Element %d not found in the array\n", num) //print msg if not found
	}
}

func FindingIndexOfNumberInSlice(arr []int, target int) int { //defining FindingIndexOfNumberInSlice func
	for i, value := range arr {
		if value == target { //check if value of index is equal to the given target
			return i
		}
	}

	return -1
}
