/*e. Create a group of constants using var () group constant declaration syntax, assign it some
values & print it. Try to change its values & check if there is some error.
*/

package main

import "fmt"

func main() {
	const ( //declaring group of constants using var () group constant declaration syntax
		name = "John"
		age  = 30
	)

	fmt.Printf("printing name in constants var group : %s\n", name) // printing name from above cost group
	fmt.Printf("printing age in constants var group : %d\n", age)   // printing age from above cost group

	// name = "Steven" //cannot assign to name to const type

}
