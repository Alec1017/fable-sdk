package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Definitions for environment variables
type Env struct {
	PrivateKey string
	NodeUri    string
	ChainId    string
}

// Reads in variable from environment, and checks that it is set
func readInVariable(env string) {

	// Read variable in from environment
	_ = viper.BindEnv(env)

	// Make sure the variable is set
	if !viper.IsSet(env) || viper.GetString(env) == "" {
		errorMessage := fmt.Sprintf("Environment variable error: %s is missing", env)
		panic(errorMessage)
	}
}

// Load environment variables either from .env file
// or from the environment
func GetEnv() *Env {

	// Setup config for .env file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	// Read in the environment variable file
	if err := viper.ReadInConfig(); err != nil {

		// Check if the error response is because the config file wasnt found
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, we can ignore
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	// Read in from environment variables, and check that it is set
	readInVariable("PRIVATE_KEY")
	readInVariable("NODE_URI")
	readInVariable("CHAIN_ID")

	// Retrieve config variables
	return &Env{
		PrivateKey: viper.GetString("PRIVATE_KEY"),
		NodeUri:    viper.GetString("NODE_URI"),
		ChainId:    viper.GetString("CHAIN_ID"),
	}
}
