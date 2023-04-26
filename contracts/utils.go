package contracts

import (
	"fmt"

	"github.com/Alec1017/fable-sdk/types"
)

func NewFactoryToken(creatorAddr string, subdenom string) types.FactoryToken {
	return types.FactoryToken{
		CreatorAddr: creatorAddr,
		Denom:       fmt.Sprintf(`factory/%s/%s`, creatorAddr, subdenom),
		Subdenom:    subdenom,
	}
}
