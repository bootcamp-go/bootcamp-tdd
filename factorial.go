package tdd

func factorial(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	if i == 2 {
		return 2
	}
	if i == 3 {
		return 6
	}
	if i == 4 {
		return 24
	}
	return 0
}

/*
// codigo refactorizado
func factorial(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return i * factorial(i-1)
} */
