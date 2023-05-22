package utils

// Check if an int array contains a value.
func SliceContains[T comparable](array []T, value T) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}
