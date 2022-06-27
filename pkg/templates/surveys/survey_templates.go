package surveys

import (
	"github.com/AlecAivazis/survey/v2"
)

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
					"node",
					"java",
				},
				Default:  "go",
				PageSize: 3,
			},
			Validate: survey.Required,
		},
	}
	// HttpIntegrationQuestion = []*survey.Question{
	// 	{
	// 		Name: "httpintegration",
	// 		Prompt: &survey.Select{
	// 			Message: "Does your lambda integrate with API gateway",
	// 			Options: []string{
	// 				"yes",
	// 				"no",
	// 			},
	// 			Default:  "yes",
	// 			PageSize: 2,
	// 		},
	// 		Validate: survey.Required,
	// 	},
	// }
	LambdaNameQuestion = []*survey.Question{
		{
			Name:     "lambdaname",
			Prompt:   &survey.Input{Message: "What is the name of your lambda function"},
			Validate: survey.Required,
		},
	}
	HttpMethodQuestion = []*survey.Question{
		{
			Name: "httpmethod",
			Prompt: &survey.Select{
				Message: "Which HTTP method does your lambda use",
				Options: []string{
					"get",
					"post",
					"put",
					"delete",
				},
				Default:  "get",
				PageSize: 4,
			},
			Validate: survey.Required,
		},
	}
	HttpPathQuestion = []*survey.Question{
		{
			Name:     "apipath",
			Prompt:   &survey.Input{Message: "What is the api path relative to /?"},
			Validate: survey.Required,
		},
	}
	RuntimeQuestion = []*survey.Question{
		{
			Name: "runtime",
			Prompt: &survey.Select{
				Message: "Choose a runtime:",
				Options: []string{
					"go1.x",
					"nodejs12.x",
					"nodejs14.x",
				},
				Default:  "go1.x",
				PageSize: 3,
			},
			Validate: survey.Required,
		},
	}
	NodeQuestions = []*survey.Question{
		{
			Name: "nodeFramework",
			Prompt: &survey.Select{
				Message: "Choose a framework:",
				Options: []string{
					"express",
				},
				Default:  "express",
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
					"mongodb",
				},
				Default:  "postgres",
				PageSize: 2,
			},
			Validate: survey.Required,
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
