package surveys

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
		icons.Question.Text = ""
		icons.Help.Format = "blue"
		icons.Help.Text = "Help ->"
		icons.Error.Format = "red"
		icons.Error.Text = "Note ->"
	}

	ProjectQuestion = []*survey.Question{
		{
			Name:     "projectName",
			Prompt:   &survey.Input{Message: "What do you want to call your project?"},
			Validate: survey.Required,
		},
	}

	FirstAppQuestion = []*survey.Question{
		{
			Name:     "appName",
			Prompt:   &survey.Input{Message: "What do you want to call your first microservice?"},
			Validate: survey.Required,
		},
	}

	AppQuestion = []*survey.Question{
		{
			Name:     "appName",
			Prompt:   &survey.Input{Message: "What do you want to call your microservice?"},
			Validate: survey.Required,
		},
	}

	CreateMicroserviceAssurance = []*survey.Question{
		{
			Name: "agreeMicroservice",
			Prompt: &survey.Confirm{
				Message: "create a new microservice ? ",
				Default: false,
			},
		},
	}

	DomainQuestion = []*survey.Question{
		{
			Name:     "domain",
			Prompt:   &survey.Input{Message: "What is your additional domain called?"},
			Validate: survey.Required,
		},
	}
	PortQuestion = []*survey.Question{
		{
			Name:     "port",
			Prompt:   &survey.Input{Message: "What port will your service run on?"},
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

	DatabaseQuestions = []*survey.Question{
		{
			Name: "database",
			Prompt: &survey.Select{
				Message: "Choose a database:",
				Options: []string{
					"postgres",
					"mysql",
				},
				Default:  "postgres",
				PageSize: 2,
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

func GenerateAppListTemplate(apps []string) []*survey.Question {
	appList := []*survey.Question{
		{
			Name: "appList",
			Prompt: &survey.Select{
				Message: "Choose your service:",
				Options: apps,
			},
			Validate: survey.Required,
		},
	}
	return appList
}
