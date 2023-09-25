package bob

import (
	"fmt"
	"strings"
)

// Hey returns response of Bob.
func Hey(remark string) string {
	if strings.HasSuffix(remark, "?") {
		return "Sure."
	} else if IsAllCaps(remark) {
		return "Whoa, chill out!"
	}
	return ""
}

func IsAllCaps(remark string) bool {
	fmt.Println(remark)
	for _, c := range remark {
		fmt.Println(c)
		if c <= 65 || c >= 90 {
			return false
		}
	}
	return true
}
