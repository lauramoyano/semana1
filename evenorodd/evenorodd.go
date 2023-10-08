package evenorodd

// Estructuras de control
func EvenOrOdd(a int) string {
	if a%2 != 0 {
		return "odd"
	}
	if a%2 == 0 {
		return "even"
	}
	return "idk"
}
