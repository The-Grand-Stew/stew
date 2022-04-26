package surveys

import "github.com/AlecAivazis/survey/v2"

var (
	GoQuestions = []*survey.Question{
		{
			Name: "goFramework",
			Prompt: &survey.Select{
				Message: "Choose a framework:",
				Options: []string{
					"fiber",
				},
				Default:  "fiber",
				PageSize: 1,
			},
			Validate: survey.Required,
		},
	}
)
