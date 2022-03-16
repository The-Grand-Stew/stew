package main

import (
	"stew/cmd/gofiber"
)

func main() {
	// var language, framework, database string
	// var err error
	// var template []*survey.Question
	// // Ask Language
	// err = survey.Ask(templates.LanguageQuestion, &language, survey.WithIcons(templates.SurveyIconsConfig))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // Get Frameworks
	// switch language {
	// case "go":
	// 	template = templates.GoQuestions
	// case "python":
	// 	template = templates.PythonQuestions
	// }
	// survey.Ask(template, &framework, survey.WithIcons(templates.SurveyIconsConfig))
	// fmt.Println(framework)
	// // Get Database
	// survey.Ask(templates.DatabaseQuestions, &database, survey.WithIcons(templates.SurveyIconsConfig))
	// fmt.Println(database)
	// // create a template
	// registry := language + "_" + framework + "_" + database
	// //clone template
	// git.Clone(registry)
	gofiber.DownloadTemplate("myapp")
	gofiber.AddPostgres("myapp")

}
