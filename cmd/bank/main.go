package main

import (
	"bank/pkg/bank/deposit"
	"fmt"
)

func main() {
	fmt.Println(deposit.Calculate(0, "TJS"))
	fmt.Println(deposit.Calculate(0, "USD"))
	fmt.Println(deposit.Calculate(500_000, "TJS"))
	fmt.Println(deposit.Calculate(500_000, "USD"))
	fmt.Println(deposit.Calculate(1_000_000, "TJS"))
	fmt.Println(deposit.Calculate(1_000_000, "USD"))
}
