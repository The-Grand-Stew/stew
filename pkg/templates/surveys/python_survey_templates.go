package surveys

import "github.com/AlecAivazis/survey/v2"

var (
	PythonQuestions = []*survey.Question{
		{
			Name: "pyFramework",
			Prompt: &survey.Select{
				Message: "Choose a framework:",
				Options: []string{
					"fastapi",
				},
				Default:  "fastapi",
				PageSize: 1,
			},
			Validate: survey.Required,
		},
	}
)
