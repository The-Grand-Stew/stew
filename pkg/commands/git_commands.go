package commands

import (
"github.com/go-git/go-git/v5"
"os"
"strings"
"fmt"
)

func Clone(gitUrl, clonePath string) error {
	// // Clone project template.
	// _, err := git.PlainClone(
	// 	clonePath,
	// 	false,
	// 	&git.CloneOptions{
	// 		URL: gitUrl,

	// 	},
	// )
	// if err != nil {
	// 	return err
	// }
	// return nil
	// AASHRIT CHANGES, REMOVE AND UNCOMMENT ABOVE

	username:= os.Getenv("GIT_USER");
	password:= os.Getenv("GIT_PASS");
	giturltrimmed:=strings.Replace(gitUrl,"https://","",1)
	url := fmt.Sprintf("https://%s:%s@%s", username, password, giturltrimmed)
	fmt.Println(url)
	_, err := git.PlainClone(
		clonePath,
		false,
		&git.CloneOptions{
			URL: url,

		},
	)
	if err != nil {
		return err
	}
	return nil
}
