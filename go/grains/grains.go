package grains

import (
	"fmt"
	"math"
)

//Total returns the total number of grains on Chessboard
func Total() uint64 {
	var total uint64
	for i := 0; i <= 63; i++ {
		total += uint64(math.Pow(2.0, float64(i)))
	}
	return total
}

//Square returns the number of grains on a given square in Chessboard
func Square(squareNumber int) (uint64, error) {
	if squareNumber <= 0 || squareNumber > 64 {
		return 0, fmt.Errorf("invalid square number")
	}
	return uint64(math.Pow(2.0, float64(squareNumber-1))), nil
}
