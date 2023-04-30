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

	// leaderboard module init message
	leaderboardInitMsg := fmt.Sprintf(`{
		"should_error": false
	}`)

	// Update the voting/staking module with a new contract
	fableDaoCoreGeneralPurposeModuleUpdateMsg := fmt.Sprintf(`{
		"update_general_purpose_modules": {
			"to_add": [
				{
					"admin": {
						"core_module": {}
					},
					"code_id": %d,
					"label": "Fable DAO Leaderboard Module",
					"msg": "%s"
				}
			],
			"to_remove": [
				"sei1t6k44ltqmr9alpenr8fu2g6rl8st0z8y9pl4vgg2t5p3dcqyhdyq3ny0qy"
			],
			"to_update": []
		}
	}`,
		contracts.FABLE_DAO_LEADERBOARD.CodeId,
		utils.Base64Encode(leaderboardInitMsg),
	)

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
		contracts.FABLE_DAO_CORE.Addr,
		utils.Base64Encode(fableDaoCoreGeneralPurposeModuleUpdateMsg),
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
