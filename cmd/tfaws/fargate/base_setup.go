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
	commands.ShowMessage("info", fmt.Sprintf("Cloning initial terraform scripts at %s", path), true, false)
	err := commands.Clone(url, path)
	if err != nil {
		return err
	}
	// check for creds
	commands.ShowMessage("info", "Checking for AWS credentials", true, false)
	commands.CheckCredentials()

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
	err = commands.TerragruntInit()
	if err != nil {
		fmt.Println(err)
		return err
	}
	commands.ShowMessage("info", "Applying terraform..", true, true)
	err = commands.TerragruntApply(true)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)

	return nil
}
