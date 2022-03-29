package stew

import (
	"fmt"
	"stew/cmd/gofiber"
	"stew/cmd/pyfastapi"
	"stew/pkg/configs"
	"stew/pkg/templates"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use:     "domain",
	Aliases: []string{"add-domain"},
	Short:   "Add a domain",
	Long:    "\nCreate domains",
	RunE:    runDomainCommand,
}

func addDomains(template string, domains []string) error {
	var err error
	switch template {
	case "go-fiber":
		err = gofiber.AddModel(domains)
	case "python-fastapi":
		err = pyfastapi.AddModel(domains)
	}
	return err
}

func runDomainCommand(cmd *cobra.Command, args []string) error {
	//load the config file
	cfg, err := configs.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Detected Language and framework", cfg.Language, cfg.Framework)
	// Ask for a database
	var domains string
	err = survey.Ask(templates.DomainQuestion, &domains, survey.WithIcons(templates.SurveyIconsConfig))
	if err != nil {
		fmt.Println(err)
	}
	cfg.Domains = strings.Split(domains, ",")
	template := cfg.Language + "-" + cfg.Framework
	addDomains(template, cfg.Domains)
	cfg.UpdateConfig()
	fmt.Println("Success!!")
	return nil

}

func init() {
	rootCmd.AddCommand(domainCmd)
}
