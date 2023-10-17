/*j.Slice Concatenation
Description: Create a program that concatenates two slices into a single slice.
Instructions:
1. Create two integer slices, slice1 and slice2, each with at least 3 different numbers.
2. Write code to concatenate slice2 to the end of slice1.
3. Print the concatenated slice.*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers1 := []int{} //creating slice number1
	numbers2 := []int{} //creating slice number2

	for i := 0; i < 3; i++ {
		numbers1 = append(numbers1, rand.Intn(11)) //numbers1 with randon values integer
		numbers2 = append(numbers2, rand.Intn(11)) //numbers2 with randon values integer
	}

	fmt.Println("slice number1 : ", numbers1) //printing the numbers1 slice
	fmt.Println("slice number2 : ", numbers2) //printing the numbers2 slice

	fmt.Println("appened slice2 to slice1 : ", Appending2Slices(numbers1, numbers2)) //printing the appended sclice by calling Appending2Slices func
}

func Appending2Slices(slice1, sclice2 []int) []int { //defining the Appending2Slices func
	return append(slice1, sclice2...) //returning concatenated slice2 to the end of slice1.
}
