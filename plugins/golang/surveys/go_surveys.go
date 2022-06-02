package surveys

import (
	"github.com/AlecAivazis/survey/v2"
)

var (
	GoFramework = []*survey.Question{
		{
			Name: "goFramework",
			Prompt: &survey.Select{
				Message: "Choose a framework:",
				Options: []string{
					"fiber",
					"echo",
				},
				Default:  "fiber",
				PageSize: 2,
			},
			Validate: survey.Required,
		},
	}

	GoDatabase = []*survey.Question{
		{
			Name: "goDatabase",
			Prompt: &survey.Select{
				Message: "Choose a database:",
				Options: []string{
					"postgres",
					"mysql",
					"mongo",
				},
				Default:  "postgres",
				PageSize: 3,
			},
			Validate: survey.Required,
		},
	}
)
