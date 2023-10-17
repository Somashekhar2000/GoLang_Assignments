/*a.Simple map operation
Description: Create a simple program that uses a map to store and retrieve information.
Instructions:
1. Create a map called studentGrades with student names as keys and their corresponding
grades as values.
2. Add at least three students to the map with their grades.
3. Print the grades of each student.
4. Delete one student's entry from the map.
5. Now print the students map again.
Note: student grades are float values between 0 to 10. Example 5.6, 9.2 etc
*/

package main

import "fmt"

func main() {
	studentGrades := make(map[string]float64) //creating studentGrades as map

	studentGrades["Soma"] = 8.2 //adding elements to map
	studentGrades["Siddarth"] = 9.5
	studentGrades["vishnu"] = 7.3
	fmt.Println("all students in map-----")
	for k, v := range studentGrades {
		fmt.Println(k, " has grade ", v) //printing map elements
	}

	delete(studentGrades, "vishnu") //deleting record from map

	fmt.Println("all students in map after deleting a student-----")
	for k, v := range studentGrades {
		fmt.Println(k, " has grade ", v) //printing map elements
	}
}
