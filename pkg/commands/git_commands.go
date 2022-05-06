package commands

import (
	"errors"
	"fmt"
	"os"
	"stew/pkg/templates/repositories"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func Clone(gitUrl, clonePath string) error {
	token := os.Getenv("HOMEBREW_GITHUB_API_TOKEN")
	if token == "" {
		return errors.New("Github token not found! Cannot download any of our repositories! please make sure to run `export HOMEBREW_GITHUB_API_TOKEN=<TOKEN>`")
	}
	// Clone project template.
	_, err := git.PlainClone(
		clonePath,
		false,
		&git.CloneOptions{
			URL: gitUrl,
			Auth: &http.BasicAuth{
				Username: "StewCook",
				Password: token,
			},
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
