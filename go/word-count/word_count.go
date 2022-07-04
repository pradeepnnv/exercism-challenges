package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	wc := make(map[string]int)
	var sep string
	switch {
	case strings.Contains(phrase, " "):
		sep = " "
	case strings.Contains(phrase, "\n"):
		sep = "\n"
	case strings.Contains(phrase, ","):
		sep = ","
	default:
		sep = " "
	}

	if sep != "," {
		phrase = strings.ReplaceAll(phrase, ",", "")
		//remove any punctuation
		var sb strings.Builder
		for _, c := range phrase {
			if c == '\'' || !unicode.IsPunct(c) && !unicode.IsSymbol(c) && c != '`' {
				sb.WriteRune(c)
			}
		}
		phrase = sb.String()
	}
	for _, w := range strings.Split(phrase, sep) {
		w = strings.ReplaceAll(w, "\n", "")
		if w == "" {
			continue
		}
		if w[0] == '\'' && w[len(w)-1] == '\'' {
			w = w[1 : len(w)-1]
		}

		wc[w]++
	}
	return wc
}
