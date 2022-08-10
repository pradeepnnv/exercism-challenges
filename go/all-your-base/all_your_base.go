package allyourbase

import (
	"fmt"
	"strconv"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}
	for _, v := range inputDigits {
		if v >= inputBase || v < 0 {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
	}
	var num []int
	if inputBase == outputBase {
		copy(inputDigits, num)
	}
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}
	b10 := convertToBase10(inputBase, inputDigits)
	if b10 == 0 {
		return []int{0}, nil
	}
	if outputBase == 10 {
		s := strconv.Itoa(b10)
		num = stringToIntSlice(s)

		return num, nil
	}

	for b10 > 0 {
		rem := b10 % outputBase
		num = append(num, rem)
		b10 /= outputBase
	}

	reverseSlice(num)
	return num, nil
}

func reverseSlice(input []int) {
	i := 0
	j := len(input) - 1
	for i < len(input)/2 {
		input[i], input[j] = input[j], input[i]
		i++
		j--
	}
}

func stringToIntSlice(input string) []int {
	// num := make([]int, len(input))
	num := []int{}
	for i, r := range input {
		if i == 0 && int(r-'0') == 0 {
			continue
		}
		num = append(num, int(r-'0'))
	}
	return num
}

func convertToBase10(inputBase int, inputDigits []int) int {
	n := 0

	for i := len(inputDigits); i > 0; i-- {

		n += pow(inputBase, i-1) * inputDigits[len(inputDigits)-i]
	}

	return n
}

func pow(num, exp int) int {
	v := 1
	for i := 0; i < exp; i++ {
		v *= num
	}
	return v
}
