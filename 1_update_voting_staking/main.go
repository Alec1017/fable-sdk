package main

import (
	"fmt"

	"github.com/Alec1017/fable-sdk/contracts"
	"github.com/Alec1017/fable-sdk/utils"

	seiSdk "github.com/Alec1017/golang-sdk/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	// create sei SDK client
	seiClient := utils.NewDefaultSeiClient()

	// Voting & staking module init message
	votingStakingModuleInitMsg := fmt.Sprintf(`{
		"owner": {
			"core_module":{}
		},
		"manager": "%s",
		"denom": "%s",
		"unstaking_duration": {
			"time":300
		}
	}`, seiClient.Account.String(),
		contracts.FABLE_TOKEN.Denom,
	)

	// Update the voting/staking module with a new contract
	fableDaoCoreVotingStakingMsg := fmt.Sprintf(`{
		"update_voting_module": {
			"module": {
				"admin": {
					"core_module": {}
				},
				"code_id": %d,
				"label": "Fable DAO Voting n Staking Module",
				"msg": "%s"
			}
		}
	}`,
		contracts.FABLE_DAO_VOTING_NATIVE_STAKED.CodeId,
		utils.Base64Encode(votingStakingModuleInitMsg),
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
		utils.Base64Encode(fableDaoCoreVotingStakingMsg),
	)

	// Execute the contract call
	response, err := seiClient.ExecuteContract(
		contracts.FABLE_DAO_CORE.Addr,
		admin_msg,
		[]sdk.Coin{},
		seiSdk.GasFee(sdk.NewCoin("usei", sdk.NewInt(19000))),
		seiSdk.GasLimit(uint64(1900000)),
	)

	// Check for execution errors
	if err != nil {
		panic(err)
	}

	// Print the raw log response
	fmt.Println(response.RawLog)
}
