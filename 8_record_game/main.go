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

	// add a new leaderboard game
	updateLeaderboardMsg := fmt.Sprintf(`{
		"record_game": {
			 "id": "1",
			 "team_addrs": [
				 "sei1jgttaksncs09m8zwmpew97achehqaxlv43hlzw",
				 "sei1lhnkvh7s2xfy6hdwehd4e2m3n26hl3j2snx8wp",
				 "sei1qdw5wemyrtrgzh3sh40xtuhekm4cmxqzv66ef0"
			 ],
			 "team_scores": [
				 "23",
				 "46",
				 "11"
			 ],
			 "winner": "sei1lhnkvh7s2xfy6hdwehd4e2m3n26hl3j2snx8wp"
		}
	}`)

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
		contracts.FABLE_DAO_LEADERBOARD.Addr,
		utils.Base64Encode(updateLeaderboardMsg),
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
