package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StakerDataResponse struct {
	Address sdk.AccAddress
	Team    sdk.AccAddress
	Balance uint64
}

type StakerResponse struct {
	Staker StakerDataResponse
}

type TeamStakersResponse struct {
	Team    sdk.AccAddress
	Stakers []StakerResponse
}

type StakersByTeamResponse struct {
	Teams []TeamStakersResponse
}
