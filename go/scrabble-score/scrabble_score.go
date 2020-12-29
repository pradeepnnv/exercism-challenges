package scrabble

import (
	"strings"
)

//Score returns the scrabble score of a given word
func Score(word string) int {
	score := 0
	if strings.Trim(word, " ") == "" {
		return 0
	}
	scoreMap := make(map[string]int)

	scoreMap["A"] = 1
	scoreMap["E"] = 1
	scoreMap["I"] = 1
	scoreMap["O"] = 1
	scoreMap["U"] = 1
	scoreMap["L"] = 1
	scoreMap["N"] = 1
	scoreMap["R"] = 1
	scoreMap["S"] = 1
	scoreMap["T"] = 1

	scoreMap["D"] = 2
	scoreMap["G"] = 2

	scoreMap["B"] = 3
	scoreMap["C"] = 3
	scoreMap["M"] = 3
	scoreMap["P"] = 3

	scoreMap["F"] = 4
	scoreMap["H"] = 4
	scoreMap["V"] = 4
	scoreMap["W"] = 4
	scoreMap["Y"] = 4

	scoreMap["K"] = 5

	scoreMap["J"] = 8
	scoreMap["X"] = 8

	scoreMap["Q"] = 10
	scoreMap["Z"] = 10

	for _, c := range strings.ToUpper(word) {
		score += scoreMap[string(c)]
	}
	return score
}
