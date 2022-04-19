package commands

import (
	"fmt"
	"stew/pkg/templates/repositories"

	"github.com/go-git/go-git/v5"
)

func Clone(gitUrl, clonePath string) error {

	// Clone project template.
	_, err := git.PlainClone(
		clonePath,
		false,
		&git.CloneOptions{
			URL: gitUrl,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func DownloadTemplate(templateName, clonePath string) error {
	gitUrl := repositories.MicroservicesTemplates[templateName]
	// clone template to path
	ShowMessage("info", fmt.Sprintf("Cloning Template for %s at location : %s", templateName, clonePath), true, false)
	err := Clone(gitUrl, clonePath)
	return err
}
