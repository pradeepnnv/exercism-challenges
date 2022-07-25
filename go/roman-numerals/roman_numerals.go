package romannumerals

import (
	"fmt"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3000 {
		return "", fmt.Errorf("invalid input")
	}

	thousands := input / 1000
	hundreds := (input - thousands*1000) / 100
	tens := (input - thousands*1000 - hundreds*100) / 10
	ones := (input - thousands*1000 - hundreds*100 - tens*10)

	var builder strings.Builder

	builder.WriteString(createPartRomanNumeral(thousands, "M", "", ""))
	builder.WriteString(createPartRomanNumeral(hundreds, "C", "D", "M"))
	builder.WriteString(createPartRomanNumeral(tens, "X", "L", "C"))
	builder.WriteString(createPartRomanNumeral(ones, "I", "V", "X"))
	return builder.String(), nil
}

func createPartRomanNumeral(c int, n string, n5 string, n10 string) string {
	var builder strings.Builder
	if c == 5 {
		builder.WriteString(n5)
	} else if c <= 3 {
		for i := 1; i <= c; i++ {
			builder.WriteString(n)
		}
	} else if c == 4 {
		// builder.WriteString(n + n5)
		builder.WriteString(n)
		builder.WriteString(n5)
	} else if c == 7 {
		// builder.WriteString(n5 + n + n)
		builder.WriteString(n5)
		builder.WriteString(n)
		builder.WriteString(n)
	} else if c == 6 {
		// builder.WriteString(n5 + n)
		builder.WriteString(n5)
		builder.WriteString(n)
	} else if c == 8 {
		// builder.WriteString(n5 + n + n + n)
		builder.WriteString(n5)
		builder.WriteString(n)
		builder.WriteString(n)
		builder.WriteString(n)
	} else if c == 9 {
		// builder.WriteString(n + n10)
		builder.WriteString(n)
		builder.WriteString(n10)
	}

	return builder.String()

}
