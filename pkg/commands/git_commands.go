package commands

import (
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
