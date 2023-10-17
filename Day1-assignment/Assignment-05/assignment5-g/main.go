/*g.Reverse a Slice
Description: Write a program that reverses the order of elements in a slice.
Instructions:
1. Create an integer slice called numbers with at least 5 different numbers.
2. Write code to reverse the order of elements in the numbers slice.
3. Print the reversed slice.*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := []int{} //Creating an integer slice called numbers

	for i := 0; i < 6; i++ {
		numbers = append(numbers, rand.Intn(11)) //numbers with randon values integer
	}

	fmt.Println(numbers)  //printing the numbers slice
	reverseSlice(numbers) //calling reverseSlice method

}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 { //swaping the elemenst to reverse order using for loop
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s) //printing the reversed sclice
}
