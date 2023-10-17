/*b. Create a slice of strings with the names of your favorite fruits. Use a for range loop to print
each fruit.*/

package main

import "fmt"

func main() {
	sclice1 := []string{"Mango", "Jackfruit", "Banana", "Orange", "Pineapple", "Chikku"} //slice with names of fruits

	for _, name := range sclice1 { //for range loop
		fmt.Println(name) //printing each fruit name
	}

}
