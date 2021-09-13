package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func Max(payments []types.Payment) types.Payment {
	var max types.Money
	var id int
	for index, payment := range payments {
		if max < payment.Amount {
			max = payment.Amount
			id = index
		}
	}
	return payments[id]
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	payments := []types.PaymentSource{}
	for index, card := range cards {
		if card.Balance > 0 && card.Active == true {
			payments = append(payments, types.PaymentSource{Type: card.Currency, Number: card.PAN, Balance: card.Balance})
			fmt.Println(index)

		}
	}
	return payments

}
