package pkg

func IfAllEmpty[T comparable](v ...T) bool {
	var zero T
	var i int
	for _, val := range v {
		if val == zero {
			i++
		}
	}
	return i == len(v)
}

func IfAnyEmpty[T comparable](v ...T) bool {
	var zero T
	for _, val := range v {
		if val == zero {
			return true
		}
	}
	return false
}
