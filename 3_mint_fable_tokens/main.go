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

	// Treasury contract init message (mints 10 billion FABLE tokens)
	treasuryMintMsg := fmt.Sprintf(`{
		"token_execute":{
			"mint_and_send":{
				"amount":{
					"denom":"%s", 
					"amount":"10000000000000000"
				}, 
				"recipient":"sei1jgttaksncs09m8zwmpew97achehqaxlv43hlzw"
			}
		}
	}`, contracts.FABLE_TOKEN.Denom)

	// Create a wrapper message that the admin will execute
	admin_msg := fmt.Sprintf(`{
		"execute_admin_msgs":{
			"msgs":[
				{
					"wasm":{
						"execute":{
							"contract_addr":"%s",
							"msg":"%s",
							"funds":[]
						}
					}
				}
			]
		}
	}`,
		contracts.FABLE_DAO_TREASURY.Addr,
		utils.Base64Encode(treasuryMintMsg),
	)

	// Execute the contract call
	response, err := seiClient.ExecuteContract(
		contracts.FABLE_DAO_CORE.Addr,
		admin_msg,
		[]sdk.Coin{},
		client.GasFee(sdk.NewCoin("usei", sdk.NewInt(19000))),
		client.GasLimit(uint64(1900000)),
	)

	// Check for execution errors
	if err != nil {
		panic(err)
	}

	// Print the raw log response
	fmt.Println(response.RawLog)
}
