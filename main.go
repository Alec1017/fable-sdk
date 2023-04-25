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
	nodeURI := "https://rpc.atlantic-2.seinetwork.io/"
	chainID := "atlantic-2"
	broadcastMode := "block"
	instantiateGasFee := sdk.NewCoin("usei", sdk.NewInt(30000))
	instantiateGasLimit := uint64(3000000)

	// create sei SDK client
	seiClient := seiSdk.NewClient(
		nodeURI,
		seiSdk.ChainID(chainID),
		seiSdk.PrivateKey(privKey),
		seiSdk.BroadcastMode(broadcastMode),
	)

	// Label the code ID for each Fable contract
	fableDaoCoreCodeId := uint64(175)
	fableDaoLeaderboardCodeId := uint64(176)
	fableDaoTreasuryCodeId := uint64(177)
	fableDaoVotingNativeStakedCodeId := uint64(289)
	daoProposalSignatoryCodeId := uint64(179)

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
	}`, account.String())

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
	}`, account.String())

	// Leaderboard contract & treasury contract init message
	leaderboardAndTreasuryInitMsg := fmt.Sprintf(`{
		"should_error": false
	}`)

	// Encode the init messages into base64
	proposalModuleInitMsgEncoded := base64.StdEncoding.EncodeToString([]byte(proposalModuleInitMsg))
	votingStakingModuleInitMsgEncoded := base64.StdEncoding.EncodeToString([]byte(votingStakingModuleInitMsg))
	leaderboardAndTreasuryInitMsgEncoded := base64.StdEncoding.EncodeToString([]byte(leaderboardAndTreasuryInitMsg))

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
	}`, account.String(),
		daoProposalSignatoryCodeId,
		proposalModuleInitMsgEncoded,
		fableDaoLeaderboardCodeId,
		leaderboardAndTreasuryInitMsgEncoded,
		fableDaoTreasuryCodeId,
		leaderboardAndTreasuryInitMsgEncoded,
		fableDaoVotingNativeStakedCodeId,
		votingStakingModuleInitMsgEncoded,
	)

	response, err := seiClient.InstantiateContract(
		fableDaoCoreCodeId,
		"Fable DAO Core",
		fableDaoCoreInstantiateMsg,
		[]sdk.Coin{},
		false,
		seiSdk.GasFee(instantiateGasFee),
		seiSdk.GasLimit(instantiateGasLimit),
	)

	if err != nil {
		panic(err)
	}

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
