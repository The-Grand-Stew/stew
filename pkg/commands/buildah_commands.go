package commands

import (
	"fmt"
	"strconv"
	"time"
)

func GenerateTag() string {
	return strconv.FormatInt(time.Now().Unix(), 10)

}

func BuildahBuild(imageName string) error {
	imageTag := GenerateTag()
	options := []string{"bud", fmt.Sprintf("%s:%s", imageName, imageTag), "."}
	err := ExecCommand("buildah", options, false)
	return err
}

func BuildahTagAndPush(imageName string, imageRepo string) error {
	err := ExecCommand("terragrunt", []string{"plan"}, false)
	return err
}

func BuildahPush(skipApprove bool) error {
	options := []string{"apply"}
	if skipApprove {
		options = append(options, "-auto-approve")
	}
	err := ExecCommand("terragrunt", options, false)
	return err

}
