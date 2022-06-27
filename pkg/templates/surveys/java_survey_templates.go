package surveys

import "github.com/AlecAivazis/survey/v2"

var (
	JavaQuestions = []*survey.Question{
		{
			Name: "javaFramework",
			Prompt: &survey.Select{
				Message: "Choose a framework:",
				Options: []string{
					"springboot",
				},
				Default:  "springboot",
				PageSize: 1,
			},
			Validate: survey.Required,
		},
	}
)
