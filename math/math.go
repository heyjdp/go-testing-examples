package math

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Divide(x, y int) float64 {
	if y == 0 {
		return 0.0 // Mathematically this is shite!
	}
	var x_float = float64(x)
	var y_float = float64(y)
	return x_float / y_float
}

func Multiply(x, y int) int {
	return x * y
}
