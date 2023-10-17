package person

import (
	"assignment8-c/model"
	"fmt"
)

func PrintPerson(p model.Person) { //defining PrintPerson func take person struct input
	fmt.Printf("person name is : %s.\n", p.Name) //printing person name
	fmt.Printf("person age is : %d.\n", p.Age)   //printing person age
}
