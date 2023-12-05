package utils

func IntPow(base, power int) int {
	if power == 0 {
		return 1
	}
	result := base
	for i := 2; i <= power; i++ {
		result *= base
	}
	return result
}
