/*i.Double the Values
Description: Write a program that doubles the values of all elements in a slice.
Instructions:
1. Create an integer slice called numbers with at least 6 different numbers.
2. Write code to double the values of all elements in the numbers slice.
3. Print the modified slice.
Note: Update the values in an original slice, don't create a new slice with doubled values.*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := []int{} //Creating an integer slice called numbers

	for i := 0; i < 8; i++ {
		numbers = append(numbers, rand.Intn(11)) //numbers with randon values integer
	}

	fmt.Println("sclice before doubling values :", numbers) //printing the numbers slice

	DoublingValues(&numbers)                               //calling DoublingValues func to double values in numbers slice
	fmt.Println("sclice after doubling values :", numbers) //printing the numbers slice with values doubled

}

func DoublingValues(s *[]int) { //defining func DoublingValues accepting pointer slice
	a := *s //assigning sclice to another slice
	for i := range a {
		a[i] += a[i] //doubling the values and assigning back to the same index
	}
}
