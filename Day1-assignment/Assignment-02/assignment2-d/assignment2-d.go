/*d. Create a group of variables using var () variable declaration syntax, assign it some values &
print it.*/

package main

import "fmt"

func main() {
	var ( //decalring group og varicables using var()
		name string
		age  int
	)

	name = "John"
	age = 30

	fmt.Printf("Name: %s\n", name) //printing name from var group
	fmt.Printf("Age: %d\n", age)   //printing age from var group
}
