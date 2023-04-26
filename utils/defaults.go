package utils

import (
	"github.com/Alec1017/fable-sdk/config"

	seiSdk "github.com/Alec1017/golang-sdk/core"
)

func NewDefaultSeiClient() *seiSdk.Client {
	// load environemnt variables
	env := config.GetEnv()

	// Load account from private key
	privKey := LoadPrivKeyFromHex(env.PrivateKey)

	// create sei SDK client
	seiClient := seiSdk.NewClient(
		env.NodeUri,
		seiSdk.ChainID(env.ChainId),
		seiSdk.PrivateKey(privKey),
		seiSdk.BroadcastMode("block"),
	)

	return seiClient
}
