/*a. Creating a Struct
Description: Define a struct for representing an Employee and create instances of it.
Instructions:
1. Define a struct called Employee with fields for Id, Name, Age, and City.
2. Create two instances of the Employee struct with different values.
3. Print the details of each employee.*/

package main

import "fmt"

type Employee struct { //defining Employee struct
	Id   int
	Name string
	Age  int
	City string
}

func main() {
	emp1 := Employee{ //creating employee struct
		Id:   101,
		Name: "Somashekhatr",
		Age:  23,
		City: "Bengaluru",
	}

	emp2 := Employee{ //creating employee struct
		Id:   102,
		Name: "Siddarth",
		Age:  24,
		City: "Mysore",
	}

	fmt.Print("employee1 details : \n", "Id : ", emp1.Id, "\nName : ", emp1.Name, "\nAge : ", emp1.Age, "\nCity : ", emp1.City, "\n") //printing employee1 details seperately
	fmt.Print("employee2 details : \n", "Id : ", emp2.Id, "\nName : ", emp2.Name, "\nAge : ", emp2.Age, "\nCity : ", emp2.City, "\n") //printing employee2 details seperately
}
