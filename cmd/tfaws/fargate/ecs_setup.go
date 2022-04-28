package fargate

import (
	"os"
	"path/filepath"
	"stew/pkg/commands"
)

//TODO: create common funciton for all functions in this package
func FargateSetup(infraPath string, tfvars map[string]string) error {
	path := filepath.Join(infraPath, "aws-ecs-fargate", "ecs-cluster")
	// check for creds
	err := commands.CheckCredentials()
	if err != nil {
		return err
	}
	currentDir, _ := os.Getwd()
	// generate vars file
	varsPath := filepath.Join(path, "vars.tfvars")
	err = commands.GenerateVarsFile(tfvars, varsPath)
	if err != nil {
		return err
	}
	// run terragrunt init/plan/apply
	os.Chdir(path)
	err = commands.TerragruntApply(false)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil
}
