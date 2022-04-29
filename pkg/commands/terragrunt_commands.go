package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func TerragruntInit() error {
	err := ExecCommand("terragrunt", []string{"init"}, false)
	return err
}

func TerragruntPlan() error {
	err := ExecCommand("terragrunt", []string{"plan"}, false)
	return err

}

func TerragruntApply(skipApprove bool) error {
	options := []string{"apply"}
	if skipApprove {
		options = append(options, "-auto-approve")
	}
	err := ExecCommand("terragrunt", options, false)
	return err

}

func TerragruntApplySpecific(target string) error {
	options := []string{"apply", "-auto-approve", "-target=" + target}
	err := ExecCommand("terragrunt", options, false)
	fmt.Println(err)
	return err

}

func TerragruntOutput(outputValue string) (string, error) {
	options := []string{"output", "-raw", outputValue}
	output, err := ExecCommandWithOutput("terragrunt", options)
	fmt.Println("output", output)
	return output, err
}

func GenerateVarsFile(vars map[string]string, varsPath string) error {
	var varString []string
	for key, value := range vars {
		// generate a list of vars
		varString = append(varString, fmt.Sprintf(`%s="%s"`, key, value))
	}
	// write to file
	f, err := os.Create(varsPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	data := strings.Join(varString, "\n")
	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

const TerragruntTemplate string = `terraform {
    extra_arguments "conditional_vars" {
        commands = [
            "init",
            "apply",
            "plan",
            "destroy",
            "refresh",
            "taint",
            "import"
        ]

        required_var_files = [   
            "${get_parent_terragrunt_dir()}/vars.tfvars"
        ]
    }
}   

remote_state {
    backend = "s3"
    config = {
        bucket          = "{{ .region }}-{{ .project }}-{{ .environment }}-terraform-state"
        region          = "eu-west-1"
        key             = "${path_relative_to_include()}/{{ .name }}.tfstate"
        dynamodb_table  = "{{ .region }}-{{ .project}}-{{ .environment }}-terraform-lock"
    }
}
`

func GenerateTerragruntFile(vars map[string]string, infraPath string) error {
	path := filepath.Join(infraPath, "terragrunt.hcl")
	t, err := template.New("modelTemplate").Parse(TerragruntTemplate)
	if err != nil {
		return err
	}
	// create output path to write the template
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	// vomit output to model file
	t.Execute(f, vars)
	f.Close()
	return nil
}
