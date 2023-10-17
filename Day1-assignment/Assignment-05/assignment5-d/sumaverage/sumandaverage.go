package sumaverage

import "fmt"

func SumAndAverageOfSlice(s []int) {
	sum := 0
	for i := range s {
		sum += s[i] //summing the int elements
	}

	fmt.Printf("sum of a list of numbers stored in given slice is %d\n", sum)                       // print the sum of all the numbers in the numbers slice.
	fmt.Printf("average of a list of numbers stored in given slice is %f\n", (float64)(sum/len(s))) // print the average of all the numbers in the numbers slice.
}
