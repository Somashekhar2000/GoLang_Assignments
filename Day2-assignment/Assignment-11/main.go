/*Assignment 11: Struct Composition
Task Description: Create a program that showcases struct composition.
Instructions:
1. Define a struct named Address that has three fields - Street, City, and ZipCode.
2. Define another struct named Person that has two fields - Name and embed Address struct in
Person.
3. Create an instance of the Person struct and initialize both the Name and Address fields in
your main function..
4. Print the details of the person, including their name and address.*/

package main

import "fmt"

type Address struct { //defining Address struct with fields
	Street  string
	City    string
	ZipCode int
}

type Person struct { //defining Address Person with fields
	Name         string
	PersonAdress Address //field PersonAdress as adress struct
}

func main() {
	p1 := Person{ //creating person struct
		Name: "Somashekhar",
		PersonAdress: Address{
			Street:  "Rukmini Nagr",
			City:    "Bengaluru",
			ZipCode: 560073,
		},
	}

	fmt.Println("Person Name : ", p1.Name)          //printing person name
	fmt.Println("Person Adress :", p1.PersonAdress) //printing person adress ie struct

}
