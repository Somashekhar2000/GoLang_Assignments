/*a. Create an array of integers with 5 elements and print each element. (use simple for loop)*/

package main

import "fmt"

func main() {
	array1 := [5]int{10, 20, 30, 40, 50} //array of 5 elements
	for i := 0; i < len(array1); i++ {   //simple for loop
		fmt.Printf("Element %d is : %d\n", i+1, array1[i]) //printing each element of array
	}
}
