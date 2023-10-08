package bmi

func CalculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}
