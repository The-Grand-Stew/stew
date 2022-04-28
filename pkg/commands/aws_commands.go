package commands

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

func CheckCredentials() error {
	//TODO: look inot chain provider extra env cred is not needed
	// Retrieve the credentials value from env vars
	creds := credentials.NewEnvCredentials()
	credValue, errEnv := creds.Get()
	creds = credentials.NewChainCredentials([]credentials.Provider{})
	if errEnv == nil {
		return nil
	}
	// Retrieve the credentials value from chain
	credValue, errChain := creds.Get()
	fmt.Println(credValue)
	if errChain != nil {
		// handle error
		return errors.New("Credentials not found, please configure your AWS credentials manually for now.(Configuration from stew coming soon)")
	}
	return nil
}
