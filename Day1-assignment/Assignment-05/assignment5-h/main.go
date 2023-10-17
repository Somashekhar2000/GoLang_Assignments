/*h.Count Even and Odd Numbers
Description: Create a program that counts and prints the number of even and odd numbers in a
slice.
Instructions:
1. Create an integer slice called numbers with at least 8 numbers (some even and some
odd).
2. Write code to count the number of even and odd numbers in the numbers slice*/

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

	fmt.Println(numbers)                  //printing the numbers slice
	even, odd := countEvenAndOdd(numbers) //calling countEvenAndOdd func that returns the count of even and odd numbers in the sclice numbers

	fmt.Printf("even numbers : %d\n odd numbers : %d", even, odd) //printing the count of even and odd numbers

}

func countEvenAndOdd(s []int) (int, int) { //defing countEvenAndOdd func
	var even int //declaring even count variable
	var odd int  //declarinf odd count variable
	for _value := range s {
		if _value%2 == 0 { //check if number is even
			even++ //increase in even number count
		}
		odd++ //incerase the odd number count
	}

	return even, odd //returning the even and odd counts
}
