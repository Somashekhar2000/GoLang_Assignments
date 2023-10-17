/*k.Slice Deduplication
Description: Create a program that removes duplicate elements from a slice.
Instructions:
1. Create an integer slice called numbers with some duplicate values.
2. Write code to remove duplicates from the numbers slice.
3. Print the deduplicated slice.
*/

package main

import "fmt"

func main() {
	numbers := []int{12, 5, 8, 23, 12, 9, 1, 1}                                   //Creating an integer slice called numbers
	fmt.Println("sclice with duplicate values : ", numbers)                       //printing numbers slice
	fmt.Println("sclice without duplicate removed : ", RemoveDuplicates(numbers)) //calling RemoveDuplicates func and printing sclice removed duplicates

}

func RemoveDuplicates(slice1 []int) []int { //definfing RemoveDuplicates func
	map1 := make(map[int]bool) //creating map1
	slice2 := []int{}          //creating slice2

	for _, value := range slice1 {
		if !map1[value] { //check if value already exists
			map1[value] = true             //adding calue to map
			slice2 = append(slice2, value) //appending value to new slice
		}
	}
	return slice2 //returning new slice
}
