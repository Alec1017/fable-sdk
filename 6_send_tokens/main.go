package main

import (
	"fmt"

	"github.com/Alec1017/fable-sdk/contracts"
	"github.com/Alec1017/fable-sdk/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/golang-sdk/client"
)

func main() {
	// create sei SDK client
	seiClient := utils.NewDefaultSeiClient()

	// Define address to send tokens to
	toAddr, err := sdk.AccAddressFromBech32("sei19alhhjzqsg25fae3m6fpxm4pl8fplm3d23uvzl")
	if err != nil {
		panic(err)
	}

	// Define FABLE funds to send
	funds := []sdk.Coin{
		sdk.NewCoin(contracts.FABLE_TOKEN.Denom, sdk.NewInt(100000000)),
	}

	// Send funds
	response, err := seiClient.BankSend(
		toAddr,
		funds,
		client.GasFee(sdk.NewCoin("usei", sdk.NewInt(2000))),
		client.GasLimit(uint64(200000)),
	)

	// Check for execution errors
	if err != nil {
		panic(err)
	}

	// Print the raw log response
	fmt.Println(response.RawLog)

}
