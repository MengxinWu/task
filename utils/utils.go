package utils

// IntInSlice int in slice.
func IntInSlice(a int, s []int) bool {
	for _, t := range s {
		if t == a {
			return true
		}
	}
	return false
}
