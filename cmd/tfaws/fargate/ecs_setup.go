package fargate

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
)

func ECRSetup(infraPath string, tfvars map[string]string) (string, error) {
	path := filepath.Join(infraPath, "aws-ecs-fargate", "ecs-fargate")
	// check for creds
	err := commands.CheckCredentials()
	if err != nil {
		return "", err
	}
	currentDir, _ := os.Getwd()
	// generate vars file
	varsPath := filepath.Join(path, "vars.tfvars")
	commands.ShowMessage("info", "Generating tfvars file..", true, true)
	err = commands.GenerateVarsFile(tfvars, varsPath)
	if err != nil {
		return "", err
	}
	commands.ShowMessage("info", "Generating terragrunt file..", true, true)
	// generate terragrunt file
	commands.GenerateTerragruntFile(tfvars, path)
	// run terragrunt apply for ecr only
	os.Chdir(path)
	commands.ShowMessage("info", "Creating ECR..", true, true)
	target := "aws_ecr_repository.app"
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
	err := commands.CheckCredentials()
	if err != nil {
		return err
	}
	currentDir, _ := os.Getwd()
	// generate vars file
	varsPath := filepath.Join(path, "vars.tfvars")
	commands.ShowMessage("info", "Generating tfvars file..", true, true)
	err = commands.GenerateVarsFile(tfvars, varsPath)
	if err != nil {
		return err
	}
	commands.ShowMessage("info", "Generating terragrunt file..", true, true)
	// generate terragrunt file
	commands.GenerateTerragruntFile(tfvars, path)
	// run terragrunt apply for ecr only
	os.Chdir(path)
	commands.ShowMessage("info", "Deploying to Fargate..", true, true)
	err = commands.TerragruntApply(true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// get the ecr image registry id
	os.Chdir(currentDir)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
