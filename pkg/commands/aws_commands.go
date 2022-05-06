package commands

import (
	"os"
	"stew/pkg/templates/surveys"

	"github.com/AlecAivazis/survey/v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func CheckCredentials() {
	//TODO: look inot chain provider extra env cred is not needed
	creds := credentials.NewEnvCredentials()
	_, errEnv := creds.Get()

	if errEnv != nil {

		ShowMessage("info", "AWS Credentials not found, setting credentials", true, false)
		var AWSAccessKeyId, AWSSecretAccessKey, AWSSessionToken string
		survey.Ask(surveys.AWSAccessKeyId, &AWSAccessKeyId)
		survey.Ask(surveys.AWSSecretAccessKey, &AWSSecretAccessKey)
		survey.Ask(surveys.AWSSessionToken, &AWSSessionToken)
		os.Setenv("AWS_SECRET_ACCESS_KEY", AWSAccessKeyId)
		os.Setenv("AWS_ACCESS_KEY_ID", AWSSecretAccessKey)
		os.Setenv("AWS_SESSION_TOKEN", AWSSessionToken)
	}

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
