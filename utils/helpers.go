package utils

import (
	"github.com/Alec1017/fable-sdk/config"

	"github.com/sei-protocol/golang-sdk/client"
)

func NewDefaultSeiClient() *client.Client {
	// load environemnt variables
	env := config.GetEnv()

	// Load account from private key
	privKey := LoadPrivKeyFromHex(env.PrivateKey)

	// create sei SDK client
	seiClient := client.NewClient(
		env.NodeUri,
		client.ChainID(env.ChainId),
		client.PrivateKey(privKey),
		client.BroadcastMode("block"),
	)

	return seiClient
}

func ExecuteAdminMsgs() {

}
