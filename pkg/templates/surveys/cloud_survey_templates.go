package surveys

import "github.com/AlecAivazis/survey/v2"

var (
	CloudProviderQuestion = []*survey.Question{
		{
			Name: "cloud",
			Prompt: &survey.Select{
				Message: "Choose a cloud provider:",
				Options: []string{
					"aws",
					"gcp",
				},
				Default:  "aws",
				PageSize: 2,
			},
			Validate: survey.Required,
		},
	}

	CloudInfraTypeQuestion = []*survey.Question{
		{
			Name: "infraType",
			Prompt: &survey.Select{
				Message: "How are your microservices going to run?",
				Options: []string{
					"container-based",
					"serverless",
				},
				Default:  "container-based",
				PageSize: 2,
			},
			Validate: survey.Required,
		},
	}
	EnvNameQuestion = []*survey.Question{
		{
			Name:     "infraType",
			Prompt:   &survey.Input{Message: "What environment are you deploying for (eg. dev or nonprod)?"},
			Validate: survey.Required,
		},
	}
)
