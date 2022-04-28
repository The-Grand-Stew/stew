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

	fmt.Println(url)
	path := filepath.Join(infraPath, "tf-aws-setup")
	err := commands.Clone(url, path)
	fmt.Println(path)
	if err != nil {
		return err
	}
	// check for creds ??
	err = commands.CheckCredentials()
	if err != nil {
		return err
	}
	// generate vars file
	varsPath := filepath.Join(path, "vars.tfvars")
	err = commands.GenerateVarsFile(tfvars, varsPath)
	if err != nil {
		return err
	}
	// run terragrunt init/plan/apply
	currentDir, _ := os.Getwd()
	os.Chdir(path)
	err = commands.TerragruntApply(true)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil
}
