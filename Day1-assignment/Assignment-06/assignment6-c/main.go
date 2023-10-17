/*c. Map of Maps
Description: Create a program that uses a map of maps to store and retrieve data.
Instructions:
1. Create a map called studentData where each student's name is the key, and the value is
another map containing information such as Age, Grade, and City.
2. Add at least three students to the studentData map with their respective information.
3. Retrieve and print the details of each student.*/

package main

import "fmt"

func main() {
	studentData := make(map[string]map[string]interface{})                                            // this is a map which contains key as string and value as a another amp
	studentData["Somashekhar"] = map[string]interface{}{"Age": 23, "Grade": "A", "City": "Bengaluru"} //adding elements to map
	studentData["SIddarth"] = map[string]interface{}{"Age": 21, "Grade": "B", "City": "Mysore"}
	studentData["Vishnu"] = map[string]interface{}{"Age": 23, "Grade": "C", "City": "Tumkur"}
	for k, v := range studentData {
		fmt.Print(k, ": ")
		for k1, v1 := range v {
			fmt.Print(k1, ": ", v1, " ") //printing map
		}
		fmt.Println()
	}

}
