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

	// Update the subdaos in the core module
	updateSubdaosMsg := fmt.Sprintf(`{
		"update_sub_daos": {
			"to_add": [
				{
					"addr": "sei1jgttaksncs09m8zwmpew97achehqaxlv43hlzw",
					"charter": "Red Pandas"
				},
				{
					"addr": "sei1qdw5wemyrtrgzh3sh40xtuhekm4cmxqzv66ef0",
					"charter": "Green Mambas"
				},
				{
					"addr": "sei1lhnkvh7s2xfy6hdwehd4e2m3n26hl3j2snx8wp",
					"charter": "Blue Whales"
				}
			], 
			"to_remove": [
				"sei1e6hu590ndqsuss568wh4du73g3qxn07udcyw28",
				"sei1swuakr3uvqq72ag34h99pa38edm839vd0h2q8g",
				"sei1265wwrle7k5yrcapzk7grh5487hhhxshpeflrz"
			]
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
		contracts.FABLE_DAO_CORE.Addr,
		utils.Base64Encode(updateSubdaosMsg),
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
