package commands

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func CheckCredentials() error {
	//TODO: look inot chain provider extra env cred is not needed
	// Retrieve the credentials value from env vars
	creds := credentials.NewEnvCredentials()
	_, errEnv := creds.Get()
	creds = credentials.NewChainCredentials([]credentials.Provider{})
	if errEnv == nil {
		return nil
	}
	// Retrieve the credentials value from chain
	_, errChain := creds.Get()
	if errChain != nil {
		// handle error
		return errors.New("Credentials not found, please configure your AWS credentials manually for now.(Configuration from stew coming soon)")
	}
	return nil
}

func CreateECRRepository(repoName string) (*string, error) {
	svc := ecr.New(session.New())
	input := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String("application/" + repoName),
	}

	result, err := svc.CreateRepository(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == ecr.ErrCodeRepositoryAlreadyExistsException {
				ShowMessage("info", "Repository Exists", true, true)
				return result.Repository.RepositoryUri, nil
			}

		}
		return nil, err
	}
	return result.Repository.RepositoryUri, nil

}
