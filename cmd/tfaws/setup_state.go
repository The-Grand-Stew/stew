package tfaws

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates/repositories"
)

func SetupTfStateResources(infraPath string, tfvars map[string]string) error {
	// clone the repo for setup
	url := repositories.CloudInfraTemplates["aws-setup"]
	path := filepath.Join(infraPath, "tf-aws-setup")
	err := commands.Clone(url, path)
	if err != nil {
		return err
	}
	logging.ShowMessage("info", "Checking for AWS credentials", true, false)
	// check for creds
	commands.CheckCredentials()
	// generate vars file
	varsPath := filepath.Join(path, "vars.tfvars")
	logging.ShowMessage("info", "Generating tfvars file..", true, true)
	err = commands.GenerateVarsFile(tfvars, varsPath)
	if err != nil {
		return err
	}
	// run terragrunt init/plan/apply
	currentDir, _ := os.Getwd()
	os.Chdir(path)
	logging.ShowMessage("info", "Applying terraform..", true, true)
	err = commands.TerragruntApply(true)
	if err != nil {
		fmt.Println(err)
		return err
	}
	os.Chdir(currentDir)
	return nil
}
