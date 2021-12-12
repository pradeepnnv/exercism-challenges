package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return fmt.Sprintf("%s\n%s\n%s", strings.Repeat("*", numStarsPerLine),
		welcomeMsg,
		strings.Repeat("*", numStarsPerLine))
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	lines := strings.Split(oldMsg, "\n")
	cm := ""
	for _, l := range lines {

		if !strings.Contains(l, "****") {
			fmt.Println("found a line")
			cm = strings.Trim(l, "*")
			cm = strings.Trim(cm, " ")
		}
	}
	return cm
}
