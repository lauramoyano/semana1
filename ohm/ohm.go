package ohm

// Estructuras de control (*) se utiliza para declarar punteros
func Ohm(v *int, r *int, i *int) *int {
	if v != nil && r != nil && i != nil {
		return nil
	}
	if v != nil && r != nil {
		result := *v / *r
		return &result
	}
	if v != nil && i != nil {
		result := *v / *i
		return &result
	}
	if r != nil && i != nil {
		result := *r * *i
		return &result
	}
	return nil
}
