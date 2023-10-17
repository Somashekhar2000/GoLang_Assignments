/*d. Struct, functions, package
1. Create a Circle struct In a model package, with a Radius field.
2. Create a shape package.Inside shape package create a function to calculate Area of a
Circle.
3. Call this Area func from your main function.*/

package main

import (
	"assignment8-d/model"
	"assignment8-d/shape"
	"fmt"
)

func main() {
	c := model.Circle{ //creating circle struct
		Radius: 5,
	}

	fmt.Printf("Area of circle is %f.\n", shape.AreaofACircle(c)) //calling AreaofACircle func to get area of it
}
