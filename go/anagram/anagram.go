package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {

	list := []string{}

	//create map of subject
	sMap := make(map[rune]int)
	subject = strings.ToLower(subject)
	for _, r := range subject {
		sMap[r]++
	}

	for _, s := range candidates {
		c := strings.ToLower(s)
		if len(subject) != len(c) || subject == c {
			continue
		}
		//create map of each candidate
		cMap := make(map[rune]int)
		for _, r := range c {
			cMap[r]++
		}

		isAnagram := true
		for k, v := range sMap {
			if cMap[k] != v {
				isAnagram = false
			}
		}
		if isAnagram {
			list = append(list, s)
		}

	}
	return list
}
