package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	result := Withdraw(&card, 10_000_00)
	fmt.Println(result.Balance)

	//Output:
	//1000000

}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Active: true}
	result := Withdraw(&card, 30_000_00)
	fmt.Println(result.Balance)
	//Output:
	//0

}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 0, Active: false}
	result := Withdraw(&card, 5_000_00)
	fmt.Println(result.Balance)
	//Output:
	//0
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	result := Withdraw(&card, 40_000_00)
	fmt.Println(result.Balance)
	//Output:
	//2000000
}

func ExampleDeposit_limitEqual() {
	card := types.Card{Balance: 20_000_00, Active: true}
	result := Deposit(&card, 10_000_00)
	fmt.Println(result.Balance)

	//Output:
	//3000000

}
func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	result := Deposit(&card, 10_000_00)
	fmt.Println(result.Balance)

	//Output:
	//2000000

}

func ExampleDeposit_limitAbove() {
	card := types.Card{Balance: 20_000_00, Active: true}
	result := Deposit(&card, 60_000_00)
	fmt.Println(result.Balance)

	//Output:
	//2000000

}
