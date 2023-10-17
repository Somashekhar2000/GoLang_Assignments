/*b. Declare a variable of type float64 and assign it a value of 3.14. Print the variable.
 */
package main

import "fmt"

var floatNum1 float64 = 3.14 // global float variable declaration using var key wordon

func main() {
	floatNum2 := 3.14                                                                                     // local  float variable  using short  declaration function
	var floatNum3 float64 = 3.14                                                                          //local float  variable  using var keyword  declaration
	fmt.Printf("printing global float64 variable floatNum1 : %f\n", floatNum1)                            //printing global float variable floatNum1
	fmt.Printf("printing float64 variable using short  declaration function floatNum2 : %f\n", floatNum2) // printing float variable using short  declaration function floatNum2
	fmt.Printf("printing float64 variable using var keyword floatNum2 : %f\n", floatNum3)                 //printing float variable using var keyword floatNum3
}
