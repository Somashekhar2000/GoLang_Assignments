/*m.Slice Comparison
Description: Write a program that compares two slices and prints whether they are equal or not.
Instructions:
1. Create two integer slices, slice1 and slice2, with at least 5 numbers in each.
2. Write code to compare whether slice1 and slice2 are equal (have the same elements in
the same order).
3. Print whether the slices are equal or not.*/

package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30, 40, 50} //creating slice1 with 5 elements
	slice2 := []int{10, 10, 30, 40, 50} //creating slice2 with 5 elements

	if !SlicesEqualityCheck(slice1, slice2) { //check for bool return by calling SlicesEqualityCheck func
		fmt.Println("Both the slices are not equal") //printing if both sclies not equal
		return
	}
	fmt.Println("Both the slices are qeual") //printing if both sclies equal

}

func SlicesEqualityCheck(slice1, slice2 []int) bool { //defining SlicesEqualityCheck func to check equality
	if len(slice1) == len(slice2) { //check if both slices length are equal
		for i := 0; i < len(slice1); i++ {
			if slice1[i] == slice2[i] { //check if each elements atre equal
				continue // continue if elemenst are equal
			}
			return false //return false if elemenst are not equal
		}
		return true //return true if elemnts are equal
	}
	return false //return if lengths are not equal
}
