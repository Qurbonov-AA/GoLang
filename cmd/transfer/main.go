package main

import (
	"bank/pkg/bank/card"
	"fmt"
)

func main() {
	result := card.IssueCard("TJS", "White", "Infinity")
	fmt.Println(result)
}
