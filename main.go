package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/Alec1017/fable-sdk/config"
	seiSdk "github.com/Alec1017/golang-sdk/core"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {

	// load environemnt variables
	env := config.GetEnv()

	// Instantiate sei client
	data, err := hex.DecodeString(env.PrivateKey)
	if err != nil {
		panic(err)
	}
	privKey := &secp256k1.PrivKey{
		Key: data,
	}
	account := sdk.AccAddress(privKey.PubKey().Address())
	seiClient := seiSdk.NewClientWithDefaultConfig(privKey)

	// Label the code ID for each Fable contract
	fableDaoCoreCodeId := uint64(1)
	fableDaoLeaderboardCodeId := uint64(2)
	fableDaoTreasuryCodeId := uint64(3)
	fableDaoVotingNativeStakedCodeId := uint64(4)
	daoProposalSignatoryCodeId := uint64(5)

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
	}`, account.String(), daoProposalSignatoryCodeId, proposalModuleInitMsgEncoded, fableDaoLeaderboardCodeId, leaderboardAndTreasuryInitMsgEncoded,
		fableDaoTreasuryCodeId, leaderboardAndTreasuryInitMsgEncoded, fableDaoVotingNativeStakedCodeId, votingStakingModuleInitMsgEncoded)

	response, err := seiClient.InstantiateContract(fableDaoCoreCodeId, fableDaoCoreInstantiateMsg)

	if err != nil {
		panic(err)
	}

	contractAddr := seiSdk.GetEventAttributeValue(*response, "instantiate", "_contract_address")

	fmt.Println(fmt.Sprintf("Fable DAO Core Address: %s", contractAddr))
}
