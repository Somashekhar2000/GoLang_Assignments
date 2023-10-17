package main

import "fmt"

func divide(numerator, denominator int) { //divide function to create panic
	if denominator == 0 { //check if denominator is zero
		panic("denominator is zero") //intentionally creating panic
	}
	fmt.Println("no problem in denominator")
}

func safeDivide(numerator, denominator int) { // defining safeDivide func
	defer recoveryfunc()           // calling recover function to handle panic
	divide(numerator, denominator) // calling divide func
}

func main() {
	safeDivide(10, 0) //calling safeDivide func
}

func recoveryfunc() { //defining recoveryfunc
	m := recover() //calling revcover func to handle panic ad give err msg

	if m != nil { //check if msg is nil
		fmt.Println(m) //print msg
		fmt.Println("recovered from the panic")
	}
}
