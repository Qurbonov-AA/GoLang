package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{{ID: 1, Amount: 200_000_00}, {ID: 2, Amount: 300_000_00}, {ID: 99, Amount: 350_000_00}}
	max := Max(payments)
	fmt.Println(max.ID)

	//99
}
