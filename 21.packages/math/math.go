package math

func Sum[T int | float64](x, y T) T {
	return x + y
}

func Subtract[T int | float64](x, y T) T {
	return x - y
}

func Divide[T int | float64](x, y T) T {
	return x / y
}

func Multipy[T int | float64](x, y T) T {
	return x * y
}
