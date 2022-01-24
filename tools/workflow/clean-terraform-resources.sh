#! /bin/bash

# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# A copy of the License is located at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

set -ex

#Set environment variables
terraform_state_s3_bucket="aws-otel-test-terraform-state"
yesterday=$(docker run --rm ubuntu date -d "yesterday" "+%Y-%m-%d") # ensure we are using gnudate

function error_exit() {
  echo "$1" 1>&2
  exit 1
}

function check_deps() {
  test -f $(which aws) || error_exit "aws command is not detected in path, please install it"
  test -f $(which git) || error_exit "git command is not detected in path, please install it"
  test -f $(which terraform) || error_exit "terraform command is not detected in path, please install it"
}

function check_if_test_framework_exist(){
  if [ ! -d "testing-framework/" ]; then
    git clone "https://github.com/aws-observability/aws-otel-test-framework.git" "testing-framework/"
  fi
}
function clean_test_framework(){
  rm -rf "testing-framework"
}

function terraform_destroy_state() {
	key_name=$1
	echo "destroy the terraform resource"
	echo "${key_name}"
	# check if the object still exists
	still_exists=$(aws s3api head-object --bucket "${terraform_state_s3_bucket}" --key "${key_name}" || echo '')
	if [ -n "${still_exists}" ]; then

    # Split keys based on this format: https://github.com/khanhntd/aws-otel-test-framework/blob/terraform/terraform/add_on/remote_state/main.tf#L34
    key_name_split=(${key_name//// })
    platform="${key_name_split[${#key_name_split[@]}-3]}"
    testcase_name="${key_name_split[${#key_name_split[@]}-2]}"

	  if [[ $platform != "ec2" && $platform != "ecs" && $platform != "eks" ]]; then
	    error_exit "${platform} is not a supported platform."
	  fi

    if [[ ! -f "testing-framework/terraform/testcases/${testcase_name}/parameters.tfvars" ]] ; then
      error_exit "${testcase_name} is not a supported test case."
    fi

    # Download the terraform state from s3 bucket
    echo "download terraform state from s3 bucket: ${terraform_state_s3_bucket}/${key_name}"
    aws s3 cp "s3://${terraform_state_s3_bucket}/${key_name}" "testing-framework/terraform/${platform}/terraform.tfstate"

    #Destroy resources created by test case
    cd "testing-framework/terraform/${platform}"
    terraform init
    terraform destroy --auto-approve -var-file="../testcases/${testcase_name}/parameters.tfvars" #Some EKS needs to have parameters.trvars in order to destroy in the right eks cluster
    cd -

    #Remove terraform state after destroying
    rm -rf "testing-framework/terraform/${platform}/terraform.tfstate"
    aws s3 rm "s3://${terraform_state_s3_bucket}/${key_name}"

    echo "finish destroying state from s3 bucket: ${terraform_state_s3_bucket}/${key_name}"
	fi
}

s3_keys=$(aws s3api list-objects-v2 --bucket "${terraform_state_s3_bucket}" --query "Contents[?LastModified < '${yesterday}'].[Key]" --output text )

check_deps
check_if_test_framework_exist
for s3_key in $s3_keys; do
    terraform_destroy_state ${s3_key}
done
clean_test_framework

