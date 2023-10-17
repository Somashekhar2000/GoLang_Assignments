/*c.Create a constant called pi and assign it the value of Pi (math.Pi). Print the constant.*/
package main

import (
	"fmt"
	"math"
)

const pi1 = math.Pi // declaring const in global

func main() {
	const pi2 = math.Pi // declaring const in local

	fmt.Printf("printing const pi1 assigned math.pi : %f\n", pi1) //printing declared const in global
	fmt.Printf("printing const pi2 assigned math.pi : %f\n", pi2) //printing declared const in local
}
