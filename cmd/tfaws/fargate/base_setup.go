package fargate

import (
	"fmt"
	"os"
	"path/filepath"
	"stew/pkg/commands"
	"stew/pkg/templates/repositories"
)

func BaseSetup(infraPath string, tfvars map[string]string) error {
	// clone the repo for setup
	url := repositories.CloudInfraTemplates["aws-ecs-fargate"]
	path := filepath.Join(infraPath, "aws-ecs-fargate")
	err := commands.Clone(url, path)
	if err != nil {
		return err
	}
	// check for creds
	err = commands.CheckCredentials()
	if err != nil {
		return err
	}
	currentDir, _ := os.Getwd()
	baseSetupPath := filepath.Join(path, "base-setup")
	// generate vars file
	varsPath := filepath.Join(baseSetupPath, "vars.tfvars")
	commands.ShowMessage("info", "Generating tfvars file..", true, true)
	err = commands.GenerateVarsFile(tfvars, varsPath)
	if err != nil {
		return err
	}
	commands.ShowMessage("info", "Generating terragrunt file..", true, true)
	// generate terragrunt file
	commands.GenerateTerragruntFile(tfvars, baseSetupPath)
	// run terragrunt init/plan/apply
	os.Chdir(baseSetupPath)
	commands.ShowMessage("info", "Applying terraform..", true, true)
	err = commands.TerragruntApply(true)
	fmt.Println(err)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil
}
