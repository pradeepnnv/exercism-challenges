package darts

func Score(x, y float64) int {
	sum := x*x + y*y

	if sum <= 1 {
		return 10
	} else if sum > 1 && sum <= 25 {
		return 5
	} else if sum > 25 && sum <= 100 {
		return 1
	}
	return 0
}
