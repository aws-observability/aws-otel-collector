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

#Ensure we are using gnudate
yesterday=$(docker run --rm ubuntu date -d "yesterday" "+%Y-%m-23")


terraform_destroy() {
	key_name=$1
	echo "destroy the terraform resource"
	echo "${key_name}"
	# check if the object still exists
	still_exists=$(aws s3api head-object --bucket "${terraform_state_s3_bucket}" --key "${key_name}" || echo '')
	if [ -n "${still_exists}" ]; then
	  if [[ $key_name == *"/ec2/"* ]]; then platform_folder="ec2"; fi
	  if [[ $key_name == *"/ecs/"* ]]; then platform_folder="ecs"; fi
	  if [[ $key_name == *"/eks/"* ]]; then platform_folder="eks"; fi

	  # Set output in case the destroy fails
    echo "::set-output name=key-name::${key_name}"
    echo "remove s3 key: ${key_name}"

    # Download the terraform state from s3 bucket
    echo "download terraform state from s3 bucket: ${terraform_state_s3_bucket}/${key_name}"
    aws s3 cp "s3://${terraform_state_s3_bucket}/${key_name}" "testing-framework/terraform/${platform_folder}/terraform.tfstate"

    #Destroy resources created by test case
    if [ -d "testing-framework/terraform/${platform_folder}" ]; then
        cd "testing-framework/terraform/${platform_folder}"
        terraform init
        terraform destroy -auto-approve
        cd -
    fi

    #Remove terraform state after destroying
    rm -rf "testing-framework/terraform/${platform_folder}/terraform.tfstate"



	fi
}

s3_keys=$(aws s3api list-objects-v2 --bucket "${terraform_state_s3_bucket}" --query "Contents[?LastModified < '${yesterday}'].[Key]" --output text )

for s3_key in $s3_keys; do
    terraform_destroy ${s3_key}
done

