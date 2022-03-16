package git

import (
	"github.com/go-git/go-git/v5"
)

func Clone(template, clonePath string) error {
	gitUrl := microservicesTemplates[template]
	// Get current directory.

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

func PullGist() error {
	return nil
}
