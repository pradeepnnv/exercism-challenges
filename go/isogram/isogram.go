package isogram

import (
	"strings"
)

//IsIsogram determines if the given string is an isogram or not
func IsIsogram(word string) bool {
	isIsogram := true

	//Map of characters and their counts in the word
	charMap := make(map[rune]int)
	for _, char := range []rune(strings.ToLower(word)) {
		charMap[char]++
	}
	for k, v := range charMap {
		if v > 1 && (k != 45 && k != 32) {
			isIsogram = false
			break
		}
	}
	return isIsogram
}
