/*a. Declare a variable of type int and assign it a value of 42. Print the variable*/
package main

import "fmt"

var intNum1 int = 41 // global variable declaration using var key word

func main() {
	intNum2 := 42                                                                           // local  variable  using short  declaration function
	var intNum3 int = 43                                                                    //local  variable  using var keyword  declaration
	fmt.Printf("printing global int variable intNum1 : %d\n", intNum1)                      //printing global int variable intNum1
	fmt.Printf("printing int variable  using short  declaration function :  %d\n", intNum2) // printing int variable using short  declaration function
	fmt.Printf("printing int variable using var keyword intNum3 :  %d\n", intNum3)          //printing int variable using var keyword intNum3
}
