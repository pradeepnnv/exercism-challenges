package acronym

import "strings"

// Abbreviate will create an abbreviation for a given string
// Example: Input `Portable Network Graphics ` will be converted to `PNG`
// func Abbreviate(s string) string {
// 	var abr string
// 	for _, word := range strings.Split(s, " ") {
// 		word = strings.Trim(word, "_")
// 		word = strings.Trim(word, "-")
// 		// If the word contains '-', split and use them again
// 		if strings.Contains(word, "-") {
// 			for _, splitWord := range strings.Split(word, "-") {
// 				abr += strings.ToUpper(string([]byte(splitWord)[0]))
// 			}
// 			continue
// 		}
// 		if len(word) > 0 {
// 			abr += strings.ToUpper(string([]byte(word)[0]))
// 		}
// 	}
// 	return abr
// }

// Abbreviate e.g. Ruby on Rails as ROR
func Abbreviate(s string) (out string) {

	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)

	words := strings.Fields(s)

	for i := range words {
		out += string(words[i][0])
	}
	return strings.ToUpper(out)
}
