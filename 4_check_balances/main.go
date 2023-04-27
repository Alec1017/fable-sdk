package main

import (
	"fmt"

	"github.com/Alec1017/fable-sdk/contracts"
	"github.com/Alec1017/fable-sdk/utils"
)

func main() {
	// create sei SDK client
	seiClient := utils.NewDefaultSeiClient()

	balance := seiClient.GetBankBalance(seiClient.Account.String(), contracts.FABLE_TOKEN.Denom)

	fmt.Println(balance.Amount.String())
	fmt.Println(balance.Denom)
}
