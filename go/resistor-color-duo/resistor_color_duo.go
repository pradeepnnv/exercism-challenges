package resistorcolorduo

import (
	"fmt"
	"strconv"
	"strings"
)

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	val := colorCode(colors[0]) + colorCode(colors[1])

	v, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
	}
	return v
}

// ColorCode returns the resistance value of the given color.
func colorCode(color string) string {
	color = strings.ToLower(color)
	switch color {
	case "black":
		return "0"
	case "brown":
		return "1"
	case "red":
		return "2"
	case "orange":
		return "3"
	case "yellow":
		return "4"
	case "green":
		return "5"
	case "blue":
		return "6"
	case "violet":
		return "7"
	case "grey":
		return "8"
	case "white":
		return "9"
	}
	return "-1"
}
