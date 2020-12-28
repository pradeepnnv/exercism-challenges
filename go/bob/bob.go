package bob

import (
	"fmt"
	"regexp"
	"strings"
)

// Hey returns Bob's response based on the input remark
func Hey(remark string) string {
	// var r string
	fmt.Println(remark)
	if m, _ := regexp.Match("[\t\n]*", []byte(remark)); m { //Address Bob without actually saying anything
		return "Fine. Be that way!"
	} else if strings.Trim(remark, " ") == "" {
		return "Fine. Be that way!"
	} else if m, _ := regexp.Match("^[A-Z' ]+[?]$", []byte(remark)); m { //Yell a Question
		return "Calm down, I know what I'm doing!"
	}
	// else if m, _ := regexp.Match("^[A-Z| |0-9|,]*[!]*$", []byte(remark)); m { //Yell
	// 	return "Whoa, chill out!"
	// }

	words := strings.Fields(remark)
	// for i, w := range words {
	// 	if m, _ := regexp.Match("^[A-Z]*[?]$", []byte(w)); m { //Yell a Question
	// 		return "Calm down, I know what I'm doing!"
	// 	} else if m, _ := regexp.Match("[?]$", []byte(w)); m { //Question
	// 		return "Sure."
	// 	} else if m, _ := regexp.Match("^[A-Z]{2,}.*[!]*[^.]$", []byte(w)); m { //Yell
	// 		if i == len(words)-1 {
	// 			return "Whoa, chill out!"
	// 		}
	// 	}
	// }
	var w string
	if len(words) > 0 {
		w = words[len(words)-1] //w is the final word
	} else {
		return "Whatever."
	}
	// m, _ := regexp.Match("^[A-Z]{2,}[?]$", []byte(w)); m { //Yell a Question
	// 	return "Calm down, I know what I'm doing!"
	// } else
	if m, _ := regexp.Match("[?]$", []byte(w)); m { //Question
		return "Sure."
	} else if m, _ := regexp.Match("^[A-Z]{2,}.*[!]*[^.]$", []byte(w)); m { //Yell
		return "Whoa, chill out!"
	}

	return "Whatever."
}
