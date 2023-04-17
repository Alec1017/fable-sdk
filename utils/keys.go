package utils

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func LoadPrivKeyFromHex(hexPrivKey string) *secp256k1.PrivKey {

	// Decode the hex private key
	data, err := hex.DecodeString(hexPrivKey)
	if err != nil {
		panic(err)
	}

	// Create private key
	privKey := &secp256k1.PrivKey{
		Key: data,
	}

	return privKey
}

func LoadAccountFromHex(hexPrivKey string) sdk.AccAddress {

	// Decode the hex private key
	data, err := hex.DecodeString(hexPrivKey)
	if err != nil {
		panic(err)
	}

	// Create private key
	privKey := &secp256k1.PrivKey{
		Key: data,
	}

	// convert keypair into an account
	account := sdk.AccAddress(privKey.PubKey().Address())

	return account
}
