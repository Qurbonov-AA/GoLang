package card

import (
	"bank/pkg/bank/types"
	"math"
)

const withdrawLimit = 20_000_00

func IssueCard(currency types.Currency, color string, name string) types.Card {

	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: types.TJS,
		Active:   true,
		Color:    color,
		Name:     name,
	}
	return card
}

func Withdraw(card *types.Card, amount types.Money) *types.Card {

	if amount < 0 {
		return card
	}
	if amount > withdrawLimit {
		return card
	}
	if card.Active == false {
		return card
	}
	if card.Balance < amount {
		return card
	}

	card.Balance -= amount

	return card
}

func Deposit(card *types.Card, amount types.Money) *types.Card {

	if card.Active == true && amount <= 50_000_00 && amount > 0 {

		card.Balance += amount
		return card

	} else {
		return card
	}

}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) *types.Card {

	bonus := math.Trunc((float64(card.MinBalance) * float64(percent) * float64(daysInMonth)) / float64(100*daysInYear))
	if card.Active == true && card.Balance > 0 && bonus <= 500_000.0 {
		card.Balance += types.Money(bonus)
		return card
	} else {
		return card
	}

}

func Total(card []types.Card) types.Money {
	sum := types.Money(0)

	for index := range card {
		if card[index].Balance > 0 && card[index].Active == true {
			sum += card[index].Balance
		}
	}

	return sum
}
