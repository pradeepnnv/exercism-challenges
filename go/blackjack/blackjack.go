package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	val := 0
	switch card {
	case "ace":
		val = 11
	case "two":
		val = 2
	case "three":
		val = 3
	case "four":
		val = 4
	case "five":
		val = 5
	case "six":
		val = 6
	case "seven":
		val = 7
	case "eight":
		val = 8
	case "nine":
		val = 9
	case "ten":
		val = 10
	case "jack":
		val = 10
	case "queen":
		val = 10
	case "king":
		val = 10
	default:
		val = 0

	}
	return val
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	d := ""
	if !isBlackjack {
		d = "P"
	} else {
		switch dealerScore {
		case 11:
			d = "S"
		case 4:
			d = "S"
		case 10:
			d = "S"
		default:
			d = "W"
		}
	}
	return d
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	d := ""
	if handScore >= 17 {
		d = "S"
	} else if handScore <= 11 {
		d = "H"
	} else if handScore >= 12 && handScore <= 16 {
		if dealerScore >= 7 {
			d = "H"
		} else {
			d = "S"
		}
	}
	return d
}
