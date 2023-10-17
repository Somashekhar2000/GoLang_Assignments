/*Assignment 9: Using Interfaces
Description: Create a program that demonstrates the use of interfaces.
Instructions:
1. Define an interface called Shape with a method Area() that returns a float64.
2. Create two struct types, Circle and Rectangle, both implementing the Shape interface
with their respective Area() methods.
3. Create instances of both Circle and Rectangle and calculate and print their areas using
the Area() method.
*/

package main

import (
	"fmt"
	"math"
)

type Shape interface { //defining shape interface with area method
	Area() float64
}

type Circle struct { //defining circle struct with field
	Radius float64
}

type Rectangle struct { //defining Rectangle struct with fields
	Lenght int
	Width  int
}

func (c Circle) Area() float64 { //overrinding area method by circle struct
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 { //overrinding area method by Rectangle struct
	return float64(r.Lenght) * float64(r.Width)
}

func main() {
	c := Circle{ //creating circle struct
		Radius: 5,
	}

	r := Rectangle{ //creating Rectangle strcut
		Lenght: 5,
		Width:  7,
	}

	fmt.Printf("Area of Circle : %f.\n", c.Area()) //calling area method by circle struct
	fmt.Printf("Area of Rectangle : %f", r.Area()) //calling area method by Rectangle struct
}
