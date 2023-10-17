/*c.Package, Exported & Unexported members
1. Create a model package
2. In the model package, Create a Person struct with fields for Name and Age.
3. Create a person package.
4. In person package creates a function called PrintPerson that prints a person's
information.(Name & Age)
Call the PrintPerson function from main.*/

package main

import (
	"assignment8-c/model"
	"assignment8-c/person"
)

func main() {
	p := model.Person{ //creating person struct
		Name: "Soma",
		Age:  23,
	}

	person.PrintPerson(p) //calling PrintPerson func
}
