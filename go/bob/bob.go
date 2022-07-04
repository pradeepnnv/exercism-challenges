// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"fmt"
	"regexp"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	r := "Whatever."
	if regexp.MustCompile(`[A-Z0-9_][?]$`).MatchString(remark) { //yell a question at him
		fmt.Println("Yelled Question")
		r = "Calm down, I know what I'm doing!"
	} else if regexp.MustCompile(`[A-Z0-9][!]?$`).MatchString(remark) { // if you YELL AT HIM (in all capitals).
		fmt.Println("Yelled statement")
		r = "Whoa, chill out!"

	} else if regexp.MustCompile(`\w[?]$`).MatchString(remark) { //check if it's a question (not yelled)
		fmt.Println("Normal Question")
		r = "Sure."
	} else if !regexp.MustCompile(`[^a-zA-Z0-9_]+$`).MatchString(remark) { // if you address him without actually saying anything.
		fmt.Println("Say nothing")
		r = "Fine. Be that way!"
	}
	return r
}
