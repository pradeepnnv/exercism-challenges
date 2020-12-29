package diffsquares

//SumOfSquares returns the sum of squares of of first N natural numbers
func SumOfSquares(num int) int {
	if num < 1 {
		return 0
	}
	sum := 0
	for i := 1; i <= num; i++ {
		sum += (i * i)
	}
	return sum
}

//SquareOfSum returns the square of sum of first N natural numbers
func SquareOfSum(num int) int {
	if num < 1 {
		return 0
	}
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum * sum
}

//Difference returns difference of sum of squares and square of sum of first N Natural numbers
func Difference(num int) int {
	if num < 1 {
		return 0
	}
	return SquareOfSum(num) - SumOfSquares(num)
}
