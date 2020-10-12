package main

const EPSILON = 1e-7

func gte(a, b float64) bool {
	return a-b < EPSILON
}

func lte(a, b float64) bool {
	return a-b < -EPSILON
}

func eq(a, b float64) bool {
	return gte(a, b) && gte(b, a)
}

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func median(a, b, c uint32) uint32 {
	// return min(min(max(a, min(b, c)), max(b, min(a, c))), max(c, min(b, a)))
	return max(min(a, b), min(max(a, b), c)) // for the lol
}
