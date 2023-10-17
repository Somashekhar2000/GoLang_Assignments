package calculator

func SquaringInteger(num int) int { //defining SquaringInteger func
	return num * num //returning the square of num
}

func Summing2Integers(num1, num2 int) int { //defining Summing2Integers func
	return num1 + num2 //returning the sum of num1 & num2
}

func DiffereneceBetween2Integers(num1, num2 int) int { //defining DiffereneceBetween2Integers func
	return num1 - num2 //returning the difference of num1 & num2
}

func ProductOf2Integers(num1, num2 int) int { //defining ProductOf2Integers func
	return num1 * num2 //returning the product of num1 & num2
}

func DivisionFunction(num1, num2 int) (int, int) { //defining DivisionFunction func
	return num1 / num2, num1 % num2 //returning the quotient & remainder of num1 & num2
}
