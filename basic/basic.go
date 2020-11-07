package basic

// Abs calculates absolute value
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// AbsWrong is a wrong implemented Abs
func AbsWrong(a int) int {
	return a
}
