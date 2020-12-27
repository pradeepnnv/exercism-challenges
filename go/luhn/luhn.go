package luhn

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Valid returns true if the given number is a valid Luhn number
func Valid(num string) bool {
	num = strings.ReplaceAll(num, " ", "")
	bn := []byte(num)
	//Find if any invalid characters are available
	match, err := regexp.Match("^[0-9]*$", bn)
	if err != nil {
		log.Fatal("Error processing regex pattern", err)
	}

	//Return false if length is less than 2 or has invalid characters
	if len(num) <= 1 || !match {
		return false
	}
	//Reverse the string
	var r strings.Builder
	for i := 0; i < len(bn); i++ {
		r.WriteByte(bn[len(bn)-i-1])
	}
	num = r.String()

	var b strings.Builder
	sum := 0
	for i, digit := range []byte(num) {
		n, _ := strconv.Atoi(string(digit))
		if i%2 != 0 {
			if n*2 > 9 {
				n = n*2 - 9
			} else {
				n = n * 2
			}
		}
		sum += n
		b.WriteString(strconv.Itoa(n))
	}

	if sum%10 != 0 {
		return false
	}
	return true
}
