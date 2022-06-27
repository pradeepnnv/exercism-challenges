package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

func (i FancyNumber) Number() int {
	v, err := strconv.Atoi(i.n)
	if err != nil {
		fmt.Println("Error: ", err)
		v = 0
	}
	return v
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	v, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}
	val, err := strconv.Atoi(v.Value())
	if err != nil {
		return 0
	}
	return val
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	_, ok := fnb.(FancyNumber)
	val := 0.0
	if !ok {
		return fmt.Sprintf("This is a fancy box containing the number %0.1f", 0.0)
	}

	val, err := strconv.ParseFloat(fnb.Value(), 64)
	if err != nil {
		fmt.Println("Formatting error")
		return ""
	}
	return fmt.Sprintf("This is a fancy box containing the number %0.1f", val)

}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {

	var str string
	switch V := i.(type) {
	case int:
		str = DescribeNumber(float64(V))
	case float64:
		str = DescribeNumber(V)
	case FancyNumber:
		str = DescribeFancyNumberBox(V)
	case NumberBox:
		str = DescribeNumberBox(V)
	case FancyNumberBox:
		str = DescribeFancyNumberBox(V)
	default:
		str = "Return to sender"
	}
	return str
}
