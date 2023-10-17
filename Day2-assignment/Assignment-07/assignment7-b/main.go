/*b.Struct Methods
Description: Create a struct with methods for common operations.
Instructions:
1. Define a struct called Rectangle with fields for Width and Height.
2. Create methods for calculating the area and perimeter of a rectangle on this struct.
3. Create an instance of the Rectangle struct and use the methods to calculate its area and
perimeter.
4. Print the results.*/

package main

import "fmt"

type Rectangle struct { //defining Rectangle struct
	Width  int
	Height int
}

func (r Rectangle) Perimeter() int { //calling Perimeter method by rectangle
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() int { //calling area method by rectangle
	return r.Width * r.Height
}

func main() {
	rectangle := Rectangle{ //creating rectangle struct
		Width:  10,
		Height: 20,
	}

	fmt.Println("Perimeter of rectangle is : ", rectangle.Perimeter()) //calling perimeter method ny rectangle
	fmt.Println("Area of rectangle is : ", rectangle.Area())           //calling area method ny rectangle

}
