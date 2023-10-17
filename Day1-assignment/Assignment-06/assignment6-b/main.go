/*b.Map Iteration
Description: Write a program that iterates over a map and prints its key-value pairs.
Instructions:
1. Create a map called fruits with fruit names as keys and their corresponding quantities as
values.
2. Add at least four different fruits and quantities to the map.
3. Write code to iterate over the fruits map and print each fruit and its quantity.
*/

package main

import "fmt"

func main() {
	fruits := make(map[string]int) //creating fruit as map

	fruits["mango"] = 10 //adding elemenst to map
	fruits["apple"] = 20
	fruits["grape"] = 40
	fruits["chikku"] = 70

	for k, v := range fruits {
		fmt.Println(k, " - ", v) //printing map
	}
}
