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

	// Proposal module initialization message
	proposalModuleInitMsg := fmt.Sprintf(`{
		"allow_revoting": false,
		"close_proposal_on_execution_failure": true,
		"signatories": [
			"%s"
		],
		"max_voting_period": {
			"time":432000
		},
		"min_voting_period": {
			"time":0
		},
		"pre_propose_info": {
			"anyone_may_propose": {}
		},
		"only_members_execute": true,
		"threshold": {
			"absolute_count": {
					"threshold": "1"
			}
		}
	}`, seiClient.Account.String())

	// Voting & staking module init message
	votingStakingModuleInitMsg := fmt.Sprintf(`{
		"owner": {
			"core_module":{}
		},
		"manager": "%s",
		"denom": "usei",
		"unstaking_duration": {
			"time":300
		}
	}`, seiClient.Account.String())

	// Leaderboard contract & treasury contract init message
	leaderboardAndTreasuryInitMsg := fmt.Sprintf(`{
		"should_error": false
	}`)

	// Core module init message
	fableDaoCoreInstantiateMsg := fmt.Sprintf(`{
		"admin": "%s",
		"automatically_add_cw20s": false,
		"automatically_add_cw721s": false,
		"description": "Fable DAO POC",
		"image_url": "TBA",
		"name": "Fable DAO Core",
		"proposal_modules_instantiate_info": [
			{
				"admin": {
					"core_module": {}
				},
				"code_id": %d,
				"label": "Fable DAO Signatory Proposal Module",
				"msg": "%s"
			}
		],
		"general_purpose_module_instantiate_info": [
			{
				"admin": {
					"core_module": {}
				},
				"code_id": %d,
				"label": "Fable DAO Leaderboard Module",
				"msg": "%s"
			},
			{
				"admin": {
					"core_module": {}
				},
				"code_id": %d,
				"label": "Fable Treasury Module",
				"msg": "%s"
			}
		],
		"voting_module_instantiate_info": {
			"admin": {
				"core_module": {}
			},
			"code_id": %d,
			"label": "Fable DAO Voting n Staking Module",
			"msg": "%s"
		}
	}`, seiClient.Account.String(),
		contracts.DAO_PROPOSAL_SIGNATORY.CodeId,
		utils.Base64Encode(proposalModuleInitMsg),
		contracts.FABLE_DAO_LEADERBOARD.CodeId,
		utils.Base64Encode(leaderboardAndTreasuryInitMsg),
		contracts.FABLE_DAO_TREASURY.CodeId,
		utils.Base64Encode(leaderboardAndTreasuryInitMsg),
		contracts.FABLE_DAO_VOTING_STAKING.CodeId,
		utils.Base64Encode(votingStakingModuleInitMsg),
	)

	// Instantiate the contract
	response, err := seiClient.InstantiateContract(
		contracts.FABLE_DAO_CORE.CodeId,
		"Fable DAO Core",
		fableDaoCoreInstantiateMsg,
		[]sdk.Coin{},
		false,
		client.GasFee(sdk.NewCoin("usei", sdk.NewInt(30000))),
		client.GasLimit(uint64(3000000)),
	)

	// check for deployment errors
	if err != nil {
		panic(err)
	}

	// log all code IDs and their deployment addresses
	for _, log := range response.Logs {
		for _, event := range log.Events {
			if event.Type != "instantiate" {
				continue
			}
			for _, attribute := range event.Attributes {
				if attribute.Key == "_contract_address" {
					fmt.Println(fmt.Sprintf("Contract Address: %s", attribute.Value))
				}

				if attribute.Key == "code_id" {
					fmt.Println(fmt.Sprintf("Code ID: %s", attribute.Value))
				}

			}
		}
	}
}
