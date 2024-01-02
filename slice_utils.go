package utils

func UniqueSlice[T string | ~int](repeatedSlice []T) []T {
	keys := make(map[T]bool)
	var list []T
	for _, entry := range repeatedSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
