package utils

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setSeiBech32Prefix() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("sei", sdk.Bech32PrefixAccPub)
}

func LoadPrivKeyFromHex(hexPrivKey string) *secp256k1.PrivKey {

	// Ensure Sei bech32 prefix is properly set when creating the account
	setSeiBech32Prefix()

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

	// Ensure Sei bech32 prefix is properly set when creating the account
	setSeiBech32Prefix()

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
