/*c.Manipulating Slices
Description: In this assignment, you'll work with slices and perform basic operations on them.
Instructions:
1. Create an empty integer slice called numbers.
2. Add/Append the following numbers to the slice, one at a time, in this order: 5, 10, 15, 20,
25.
3. Print the numbers slice after each addition/append.
4. Calculate and print the length and capacity of the number slice at the end.
5. Remove the number 15 from the numbers slice.(No need to create a new slice)
6. Print the number slice after the removal.*/

package main

import "fmt"

func main() {
	numbers := []int{} //Create an empty integer slice called numbers.
	for i, j := 0, 0; i < 5; i++ {
		j = j + 5
		numbers = append(numbers, j)                     //Add/Append the following numbers to the slice, one at a time, in this order: 5, 10, 15, 20,25.
		fmt.Printf("appended %d to numbers sclice\n", j) //printing the appended value
		fmt.Println(numbers)                             //Printing the numbers slice after each addition/append.
	}

	fmt.Printf("lenght : %d and capacity : %d\n", len(numbers), cap(numbers)) //Calculating and printing the length and capacity of the number slice at the end.
	numbers = append(numbers[:2], numbers[3:]...)                             //Removeing the number 15 from the numbers slice by appending
	fmt.Println("element 15 removed : ", numbers)                             //Print the number slice after the removal
}
