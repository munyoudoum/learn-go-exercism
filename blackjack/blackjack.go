package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king", "queen", "jack", "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerSum := ParseCard(card1) + ParseCard(card2)
	dealerSum := ParseCard(dealerCard)
	switch {

	case playerSum == 22:
		return "P"
	case playerSum <= 20 && playerSum >= 17:
		return "S"
	case playerSum < 17 && playerSum >= 12:
		if dealerSum >= 7 {
			return "H"
		}
		return "S"
	case playerSum < 12:
		return "H"
	case playerSum <= 21:
		if dealerSum < 10 {
			return "W"
		}
		return "S"
	default:
		return "idk"
	}
}
