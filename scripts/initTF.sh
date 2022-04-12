#!/bin/bash
TF_VARS=$1
DEFAULT_VARS="variables.tf"

if [ ! -f "$DEFAULT_VARS" ]; then
    echo "variables.tf not found, make sure you are running this in your root terraform module"
    exit 1
fi

if [ ! -f "$TF_VARS" ]; then
    echo "Terraform vars file $TF_VARS not found"
    exit 1
fi
#extract required variables from tf vars file - note appName, project and environment are mandatory
VAR_REGION=$(sed -n 's/region[[:space:]]*=[[:space:]]*"\(.*\)"/\1/ p' "$TF_VARS")
APP_NAME=$(sed -n 's/appName[[:space:]]*=[[:space:]]*"\(.*\)"/\1/ p' "$TF_VARS")
PROJECT=$(sed -n 's/project[[:space:]]*=[[:space:]]*"\(.*\)"/\1/ p' "$TF_VARS" | awk '{print tolower($0)}')
DEFAULT_REGION=$(sed -n '/region/ {;n; s/.*default.*=.*"\(.*\)"/\1/ p;}' "$DEFAULT_VARS")
REGION="${VAR_REGION:-$DEFAULT_REGION}"
SHORT_REGION=$(echo "$REGION" | sed -n 's/\(.*\)-[[:digit:]]/\1/p')

VAR_ENV=$(sed -n 's/environment[[:space:]]*=[[:space:]]*"\(.*\)"/\1/ p' "$TF_VARS")
DEFAULT_ENV=$(sed -n '/environment/ {;n; s/.*default.*=.*"\(.*\)"/\1/ p;}' "$DEFAULT_VARS")
ENV="${VAR_ENV:-$DEFAULT_ENV}"

if [ -z "$APP_NAME" ] || [ -z "$PROJECT" ] || [ -z "$VAR_REGION" ] || [ -z "$VAR_ENV" ]; then
    echo "You are missing some mandatory variables in your tfvars file: ensure - appName,project,region & environment are present"
    exit 1
fi

echo 'terraform {\nbackend "s3" {}\n}' >backend.tf
BUCKET_NAME="${SHORT_REGION}-${PROJECT}-${ENV}-terraform-state"
echo "Initializing backend with..."
echo "  Region: $REGION"
echo "  Environment: $ENV"
echo "  Full s3 bucket path: s3://${BUCKET_NAME}/${APP_NAME}.tfstate"

#check if tf backend state bucket exists, if not create it
bucketstatus=$(aws s3api head-bucket --bucket "${BUCKET_NAME}" 2>&1)
if echo "${bucketstatus}" | grep 'Not Found'; then
    echo "bucket doesn't exist; attempting to create it"
    aws s3api create-bucket --bucket $BUCKET_NAME --region $REGION

elif echo "${bucketstatus}" | grep 'Forbidden'; then
    echo "Bucket exists but not owned; use credentials of the bucket owner and retry"
    exit 1
elif echo "${bucketstatus}" | grep 'Bad Request'; then
    echo "Bucket name specified is less than 3 or greater than 63 characters; Hint: Try to reduce the length of the project variable in the tf vars file"
else
    echo "Bucket owned and exists; Wohoo, sit back and let your terraform initialize"
fi

# remove old .terraform folder and initialize based on new tfvars passed in
rm -rf .terraform
terraform init \
    -backend-config="bucket=${SHORT_REGION}-${PROJECT}-${ENV}-terraform-state" \
    -backend-config="key=${APP_NAME}.tfstate" \
    -backend-config="region=${REGION}"
