package main

import (
	"encoding/json"
	"fmt"

	"github.com/Alec1017/fable-sdk/contracts"
	"github.com/Alec1017/fable-sdk/types"
	"github.com/Alec1017/fable-sdk/utils"
)

func main() {
	// create sei SDK client
	seiClient := utils.NewDefaultSeiClient()

	stakersByTeamQuery := `{
		StakersByTeam: {}
	}`

	response, err := seiClient.QueryContract(
		stakersByTeamQuery,
		contracts.FABLE_DAO_VOTING_STAKING.Addr,
	)

	if err != nil {
		panic(err)
	}

	var resp types.StakersByTeamResponse
	json.Unmarshal(response.Data.Bytes(), &resp)

	fmt.Println(resp)
}
