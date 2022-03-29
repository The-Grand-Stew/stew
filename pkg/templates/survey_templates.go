package templates

import (
	"github.com/AlecAivazis/survey/v2"
)

// type CreateAnswers struct {
// 	Language      string
// 	Backend       string
// 	Proxy         string
// 	AgreeCreation bool `survey:"agree"`
// }

var (
	SurveyIconsConfig = func(icons *survey.IconSet) {
		icons.Question.Format = "cyan"
		icons.Question.Text = "Q:"
		icons.Help.Format = "blue"
		icons.Help.Text = "Help ->"
		icons.Error.Format = "red"
		icons.Error.Text = "Note ->"
	}
	AppQuestion = []*survey.Question{
		{
			Name:     "appNameå",
			Prompt:   &survey.Input{Message: "What do you want to call your app?"},
			Validate: survey.Required,
		},
	}
	LanguageQuestion = []*survey.Question{
		{
			Name: "language",
			Prompt: &survey.Select{
				Message: "Choose a language:",
				Options: []string{
					"go",
					"python",
				},
				Default:  "go",
				PageSize: 2,
			},
			Validate: survey.Required,
		},
	}
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

	DatabaseQuestions = []*survey.Question{
		{
			Name: "database",
			Prompt: &survey.Select{
				Message: "Choose a database:",
				Options: []string{
					"postgres",
					"mysql",
					"redis",
					"n/a",
				},
				Default:  "postgres",
				PageSize: 4,
			},
			Validate: survey.Required,
		},
	}

	AgreeCreation = []*survey.Question{
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "Are you sure that we can proceed ? ;)",
				Default: true,
			},
		},
	}
)
