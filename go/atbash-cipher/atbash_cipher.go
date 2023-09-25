package atbash

import (
	"strings"
)

func Atbash(s string) string {
	var es strings.Builder
	i := 0
	for _, c := range strings.ToLower(s) {
		ec := rune(0)
		if c >= 97 && c <= 122 {
			ec = 97 + (122 - c)
		} else if c >= 48 && c <= 57 {
			ec = c
		} else {
			continue
		}
		es.WriteRune(ec)
		i++
		if i == 5 {
			es.WriteString(" ")
			i = 0
		}
	}
	return strings.TrimSpace(es.String())
}
