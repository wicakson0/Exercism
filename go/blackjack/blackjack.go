package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	var action string

	switch {
	case (card1Value == 11) && (card2Value == 11):
		action = "P"
	case (card1Value+card2Value == 21) && (dealerCardValue < 10):
		action = "W"
	case (card1Value+card2Value == 21) && (dealerCardValue >= 10):
		action = "S"
	case (card1Value+card2Value >= 17) && (card1Value+card2Value <= 20):
		action = "S"
	case (card1Value+card2Value >= 12) && (card1Value+card2Value <= 16) && (dealerCardValue < 7):
		action = "S"
	case (card1Value+card2Value >= 12) && (card1Value+card2Value <= 16) && (dealerCardValue >= 7):
		action = "H"
	case (card1Value+card2Value <= 11):
		action = "H"
	}
	return action
}
