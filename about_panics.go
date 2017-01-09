package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 4

func aboutPanics() {
	assert(divideFourBy(__divisor__) == 1) // panics are exceptional errors at runtime

	n := divideFourBy(2)
	assert(n == 2) // panics are exceptional errors at runtime
}
