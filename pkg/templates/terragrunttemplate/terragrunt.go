package terragrunttemplate

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
        bucket          = "{{ .region }}-{{ .project}}-{{ .environment }}-terraform-state"
        region          = "eu-west-1"
        key             = "${path_relative_to_include()}/{{ .name }}.tfstate"
        dynamodb_table  = "{{ .region }}-{{ .project}}-{{ .environment }}-terraform-lock"
    }
}
`
