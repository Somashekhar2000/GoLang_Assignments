/*Assignment 12: Interface Implementation
Description: Create a program that demonstrates interface implementation with multiple structs.
Instructions:
1. Define an interface Animal with a method MakeSound() that returns a string.
2. Create two structs Dog and Cat, both implementing the Animal interface with their
respective MakeSound() methods.
3. Create instances of Dog and Cat and call their MakeSound() methods to print the
sounds they make.*/

package main

import "fmt"

type Animal interface { //defining Animal interface with MakeSound method
	MakeSound() string
}

type Dog struct { //defining dog struct
}

type Cat struct { //defining cat struct
}

func (d Dog) MakeSound() string { //overriding MakeSound by dog struct
	return "bow....bow...."
}

func (c Cat) MakeSound() string { //overriding MakeSound by cat struct
	return "meow....meow...."
}

func main() {
	c := Cat{} //crating cat struct
	d := Dog{} //creating dog struct

	fmt.Printf("Dog's sound is %s.\n", d.MakeSound()) //calling MakeSound func by dog and printing its sound
	fmt.Printf("Cat's sound is %s.", c.MakeSound())   //calling MakeSound func by cat and printing its sound
}
