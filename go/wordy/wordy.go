package wordy

import (
	"fmt"
	"strconv"
	"strings"
)

// Answer returns the answer to the question in the input variable.
// Return variables: answer - integer & ok - status of answer
func Answer(question string) (int, bool) {
	// convert to lower case
	question = strings.ToLower(question)
	// remove question mark at the end if it's there
	question = strings.TrimSuffix(question, "?")
	if !strings.HasPrefix(question, "what is") {
		return 0, false
	}
	//remove prefix "what is"?
	question = strings.TrimPrefix(question, "what is")
	//trim all spaces
	question = strings.TrimSpace(question)

	words := strings.Split(strings.TrimSpace(question), " ")
	if len(words) == 1 {
		num, err := strconv.Atoi(question)
		if err != nil {
			return 0, false
		}
		return num, true
	}
	//supported operations - plus, minus, mutiplied & divided
	if len(words) <= 4 {
		return findValue(words[0], words[len(words)-1], strings.Join(words[1:len(words)-1], " "))
	} else {
		//multiple operations
		var oper string
		o1 := words[0]
		o2 := words[2]
		oper = words[1]
		if words[2] == "by" {
			oper = words[1] + " by"
			o2 = words[3]
		}

		val, ok := findValue(o1, o2, oper)
		if !ok {
			return 0, false
		}
		o1 = fmt.Sprint(val)
		o2 = words[4]
		oper = words[3]
		if len(words) >= 6 && words[5] == "by" {
			oper = words[4] + " by"
			o2 = words[6]
		} else if words[4] == "by" {
			o2 = words[5]
			oper = words[3] + " by"
		}
		return findValue(o1, o2, oper)

	}
}

func findValue(op1, op2, operand string) (int, bool) {
	o1, err := strconv.Atoi(op1)
	if err != nil {
		return 0, false
	}
	o2, err := strconv.Atoi(op2)
	if err != nil {
		return 0, false
	}
	switch operand {
	case "plus":
		return o1 + o2, true
	case "minus":
		return o1 - o2, true
	case "multiplied by":
		return o1 * o2, true
	case "divided by":
		return o1 / o2, true
	default:
		return 0, false
	}
}
