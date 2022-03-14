package git

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func Clone(template string) error {
	gitUrl := gitRegistries[template]
	// Get current directory.
	currentDir, _ := os.Getwd()
	// Set project folder.
	clonePath := filepath.Join(currentDir, template)
	// Clone project template.
	_, err := git.PlainClone(
		clonePath,
		false,
		&git.CloneOptions{
			URL: "https://" + gitUrl,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
