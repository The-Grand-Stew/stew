package surveys

import "github.com/AlecAivazis/survey/v2"

var (
	AWSRegion = []*survey.Question{
		{
			Name: "region",
			Prompt: &survey.Select{
				Message: "Choose your region:",
				Options: []string{
					"eu-west-1",
					"us-east-1",
				},
				Default:  "eu-west-1",
				PageSize: 2,
			},
			Validate: survey.Required,
		},
	}

	AWSEnvironment = []*survey.Question{
		{
			Name: "env",
			Prompt: &survey.Input{
				Message: "What do you want to call your environment (default: dev) ?",
				Default: "dev",
			},
			Validate: survey.Required,
		},
	}

	AWSSecretAccessKey = []*survey.Question{
		{
			Name: "cred",
			Prompt: &survey.Password{
				Message: "Enter value for AWS_SECRET_ACCESS_KEY: ",
			},
			Validate: survey.Required,
		},
	}

	AWSAccessKeyId = []*survey.Question{
		{
			Name: "cred",
			Prompt: &survey.Password{
				Message: "Enter value for AWS_ACCESS_KEY_ID: ",
			},
			Validate: survey.Required,
		},
	}

	AWSSessionToken = []*survey.Question{
		{
			Name: "cred",
			Prompt: &survey.Password{
				Message: "Enter value for AWS_SESSION_TOKEN (optional): ",
			},
		},
	}

	AWSComponentQuestion = []*survey.Question{
		{
			Name: "region",
			Prompt: &survey.Select{
				Message: "How do you want to deploy your services?",
				Options: []string{
					"ecs-fargate",
				},
				Default:  "ecs-fargate",
				PageSize: 1,
			},
			Validate: survey.Required,
		},
	}
)
