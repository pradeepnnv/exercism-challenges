package phonenumber

import (
	"fmt"
	"regexp"
	"strings"
)

const regPattern string = `^\+?1?\s*\(?([2-9][0-9]{2})\)?\-?\s*\.?([2-9][0-9]{2})\-?\s*\.?([0-9]{4})$`

func Number(phoneNumber string) (string, error) {
	r := regexp.MustCompile(regPattern)
	phoneNumber = strings.TrimSpace(phoneNumber)
	m := r.Match([]byte(phoneNumber))
	if !m {
		return "", fmt.Errorf("invalid phone number")
	}
	matches := r.FindAllStringSubmatch(phoneNumber, -1)

	areaCode := matches[0][1]
	exchangeCode := matches[0][2]
	subscriberNumber := matches[0][3]

	return fmt.Sprintf("%s%s%s", areaCode, exchangeCode, subscriberNumber), nil
}

func AreaCode(phoneNumber string) (string, error) {
	r := regexp.MustCompile(regPattern)
	phoneNumber = strings.TrimSpace(phoneNumber)
	m := r.Match([]byte(phoneNumber))
	if !m {
		return "", fmt.Errorf("invalid phone number")
	}
	matches := r.FindAllStringSubmatch(phoneNumber, -1)

	return matches[0][1], nil
}

func Format(phoneNumber string) (string, error) {
	r := regexp.MustCompile(regPattern)
	phoneNumber = strings.TrimSpace(phoneNumber)
	m := r.Match([]byte(phoneNumber))
	if !m {
		return "", fmt.Errorf("invalid phone number")
	}
	matches := r.FindAllStringSubmatch(phoneNumber, -1)

	areaCode := matches[0][1]
	exchangeCode := matches[0][2]
	subscriberNumber := matches[0][3]

	return fmt.Sprintf("(%s) %s-%s", areaCode, exchangeCode, subscriberNumber), nil
}
