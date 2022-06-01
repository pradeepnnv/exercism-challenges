package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[ERR\])|^(\[INF\])|^(\[TC\])|^(\[DBG\])|^(\[WRN\])|^(\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	var lines []string
	re := regexp.MustCompile(`<[\*\~\=\-]*>`)
	lines = re.Split(text, -1)
	return lines
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`\"[\w\s]*[passwordPASSWORD]{8}[\w\s]*\"`)
	count := 0
	for _, s := range lines {
		if re.MatchString(s) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	if !re.MatchString(text) {
		return text
	}

	var b strings.Builder

	for _, s := range re.Split(text, -1) {
		if !strings.Contains(s, "end-of-line") {
			b.WriteString(s)
		}
	}

	return b.String()
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w*)`)
	var output []string
	for _, line := range lines {

		if !re.MatchString(line) {
			output = append(output, line)
			continue
		}
		// fmt.Println(re.FindStringSubmatch(line))
		// fmt.Printf("[USR] %s %s\n",
		// 	re.FindStringSubmatch(line)[1],
		// 	line)
		output = append(output,
			fmt.Sprintf("[USR] %s %s",
				re.FindStringSubmatch(line)[1],
				line))
	}
	// fmt.Println(output)
	return output
}
