package main

import (
	"encoding/base64"
	"fmt"

	"github.com/Alec1017/fable-sdk/config"
	"github.com/Alec1017/fable-sdk/utils"

	seiSdk "github.com/Alec1017/golang-sdk/core"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	// load environemnt variables
	env := config.GetEnv()

	// Load account from private key
	privKey := utils.LoadPrivKeyFromHex(env.PrivateKey)
	account := utils.LoadAccountFromHex(env.PrivateKey)

	// Define configuration values
	nodeURI := "https://sei-testnet-rpc.polkachu.com/"
	chainID := "atlantic-2"
	broadcastMode := "block"
	instantiateGasFee := sdk.NewCoin("usei", sdk.NewInt(30000))
	instantiateGasLimit := uint64(3000000)

	fable_dao_core_contract_addr := "sei10ptghqesxycs2x0enhu4p0vyjxhs76ze4ey29lv0507uzl6tg45sdmd0tk"
	fableDaoVotingNativeStakedCodeId := uint64(289)

	// create sei SDK client
	seiClient := seiSdk.NewClient(
		nodeURI,
		seiSdk.ChainID(chainID),
		seiSdk.PrivateKey(privKey),
		seiSdk.BroadcastMode(broadcastMode),
	)

	// Voting & staking module init message
	votingStakingModuleInitMsg := fmt.Sprintf(`{
		"owner": {
			"core_module":{}
		},
		"manager": "%s",
		"denom": "factory/sei1x9rszpesgkk486l4lpztxhaz7vjhgcdjuqhsxx9m5zycuvwce64s5h9gj3/RACE",
		"unstaking_duration": {
			"time":300
		}
	}`, account.String())

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
		fableDaoVotingNativeStakedCodeId,
		base64.StdEncoding.EncodeToString([]byte(votingStakingModuleInitMsg)),
	)

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
		fable_dao_core_contract_addr,
		base64.StdEncoding.EncodeToString([]byte(fableDaoCoreVotingStakingMsg)),
	)

	response, err := seiClient.ExecuteContract(
		fable_dao_core_contract_addr,
		1,
		admin_msg,
		[]sdk.Coin{},
		seiSdk.GasFee(instantiateGasFee),
		seiSdk.GasLimit(instantiateGasLimit),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.RawLog)
}
