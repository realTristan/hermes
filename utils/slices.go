package utils

// SliceContains is a generic function that checks if a given value is present in a given slice of comparable values.
// Parameters:
//   - array ([]T): The slice to search for the value.
//   - value (T): The value to search for in the slice.
//
// Returns:
//   - bool: true if the value is present in the slice, false otherwise.
func SliceContains[T comparable](array []T, value T) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			return true
		}
	}
	return false
}
