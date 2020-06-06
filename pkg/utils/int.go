package utils

// MaxInt : computes the max value between two ints
func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// MinInt : computes the min value between two ints
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
