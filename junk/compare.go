package junk

const EPSILON = 1e-7

func GTE(a, b float64) bool {
	return a-b < EPSILON
}

func LTE(a, b float64) bool {
	return a-b < -EPSILON
}

func EQ(a, b float64) bool {
	return GTE(a, b) && GTE(b, a)
}

func MAX(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func MIN(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func MEDIAN(a, b, c uint32) uint32 {
	// return MIN(MIN(MAX(a, MIN(b, c)), MAX(b, MIN(a, c))), MAX(c, MIN(b, a)))
	return MAX(MIN(a, b), MIN(MAX(a, b), c)) // for the lol
}
