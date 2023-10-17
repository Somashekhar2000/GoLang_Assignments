/*l.Calculate the Sum of Even Numbers
Description: Write a program that calculates and prints the sum of all even numbers in a slice.
Instructions:
1. Create an integer slice called numbers with at least 6 different numbers (some even and
some odd).
2. Write code to calculate the sum of all even numbers in the numbers slice.
3. Print the sum of even numbers.*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := []int{} //creating numbers slice

	for i := 0; i < 10; i++ {
		numbers = append(numbers, rand.Intn(10)) //adding randon 6 different numbers with even and odd
	}

	fmt.Println(numbers)                                                                                  //printing numbers slice
	fmt.Printf("sum of all even numbers in the numbers slice : %d", SumOfAllEvenNumbersInSclice(numbers)) //printing the sum aof all even numbers in numbers slice by callinf SumOfAllEvenNumbersInSclice func
}

func SumOfAllEvenNumbersInSclice(slice1 []int) int { //defining SumOfAllEvenNumbersInSclice func
	var sum int
	for _, value := range slice1 {
		if value%2 == 0 { //check if value is even
			sum += value //sum the even value
		}
	}
	return sum //return sum of even numbers
}
