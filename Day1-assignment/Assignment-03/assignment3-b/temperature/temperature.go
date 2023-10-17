package temperature

func FahrenheitToCelsiusConvertion(fahrenheit int) float64 {
	return (float64)(fahrenheit-32) * 5 / 9
}
