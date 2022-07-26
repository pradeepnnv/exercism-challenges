package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for k, v := range in {
		for _, c := range v {
			out[strings.ToLower(string(c))] = k
		}
	}
	return out
}
