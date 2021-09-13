package payment

import (
	"bank/pkg/bank/types"
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
