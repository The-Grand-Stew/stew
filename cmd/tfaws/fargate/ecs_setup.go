package fargate

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"strings"
)

func ECRSetup(infraPath string, tfvars map[string]string) (string, error) {
	path := filepath.Join(infraPath, "aws-ecs-fargate", "ecs-fargate")
	// check for creds
	logging.ShowMessage("info", "Checking for Credentials..", true, true)
	commands.CheckCredentials()
	currentDir, _ := os.Getwd()
	// generate vars file
	varsPath := filepath.Join(path, "vars.tfvars")
	logging.ShowMessage("info", "Generating tfvars file..", true, true)
	err := commands.GenerateVarsFile(tfvars, varsPath)
	if err != nil {
		return "", err
	}
	logging.ShowMessage("info", "Generating terragrunt file..", true, true)
	// generate terragrunt file
	commands.GenerateTerragruntFile(tfvars, path)
	// run terragrunt apply for ecr only
	os.Chdir(path)
	logging.ShowMessage("info", "Creating ECR..", true, true)
	target := "aws_ecr_repository.app"
	err = commands.TerragruntInit()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	err = commands.TerragruntApplySpecific(target)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	repository_url, err := commands.TerragruntOutput("repository_url")
	// get the ecr image registry id
	os.Chdir(currentDir)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return repository_url, nil

}

//TODO: create common funciton for all functions in this package
func FargateSetup(infraPath string, tfvars map[string]string) error {
	path := filepath.Join(infraPath, "aws-ecs-fargate", "ecs-fargate")
	// check for creds
	commands.CheckCredentials()
	currentDir, _ := os.Getwd()
	// generate vars file
	varsPath := filepath.Join(path, "vars.tfvars")
	logging.ShowMessage("info", "Generating tfvars file..", true, true)
	err := commands.GenerateVarsFile(tfvars, varsPath)
	if err != nil {
		return err
	}
	logging.ShowMessage("info", "Generating terragrunt file..", true, true)
	// generate terragrunt file
	commands.GenerateTerragruntFile(tfvars, path)
	// run terragrunt apply for ecr only
	os.Chdir(path)
	err = commands.TerragruntInit()
	if err != nil {
		fmt.Println(err)
		return err
	}
	logging.ShowMessage("info", "Deploying to Fargate..", true, true)
	err = commands.TerragruntApply(true)
	if err != nil {
		return err
	}
	invoke_url, err := commands.TerragruntOutput("invoke_url")
	// get the ecr image registry id
	os.Chdir(currentDir)
	if err != nil {
		return err
	}
	logging.ShowMessage("success", "Your Service is Deployed! Here's your API Gateway invoke variable is: ", true, true)
	logging.ShowMessage("success", invoke_url, true, true)
	return nil
}

func Deploy(infraPath, imageName string, tfvars map[string]string) error {
	// create ecr repo
	repositoryUrl, err := ECRSetup(infraPath, tfvars)
	if err != nil {
		return err
	}
	logging.ShowMessage("info", "Repository url: "+repositoryUrl, true, false)
	// push to ecr
	registry := strings.Split(repositoryUrl, "/")[0]
	logging.ShowMessage("info", "Pushing Image to registry", true, false)
	err = commands.DockerLogin(tfvars["region"], registry, "aws")
	if err != nil {
		return err
	}
	image, err := commands.DockerTagAndPush(imageName, repositoryUrl)
	if err != nil {
		return err
	}
	// deploy to fg
	tfvars["app_image"] = image
	logging.ShowMessage("info", "Starting Fargate setup...", true, false)
	err = FargateSetup(infraPath, tfvars)
	// showError(err)
	return err
}
