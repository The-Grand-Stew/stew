package commands

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GenerateTag() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func DockerBuild(imageName string, imageTag string) (string, error) {
	image := fmt.Sprintf("%s:%s", imageName, imageTag)
	options := []string{"build", "--platform", "linux/amd64", "-t", image, "."}
	err := ExecCommand("docker", options, false)

	return image, err
}

func DockerTagAndPush(imageName string, repositoryUrl string) (string, error) {
	tag := strings.Split(imageName, ":")[1]
	imageRepo := repositoryUrl + ":" + tag
	err := ExecCommand("docker", []string{"tag", imageName, imageRepo}, false)
	err = ExecCommand("docker", []string{"push", imageRepo}, false)
	return imageRepo, err
}

func DockerLogin(region, registry, cloud string) error {
	switch cloud {
	case "aws":
		command := fmt.Sprintf("aws ecr get-login-password --region %s | docker login --username AWS --password-stdin %s", region, registry)

		out, err := exec.Command("bash", "-c", command).Output()
		if err != nil {
			return err
		}
		fmt.Println(string(out))

		return err
	}
	return nil

}
