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
yesterday=$(docker run --rm ubuntu date -d "yesterday" "+%Y-%m-22")


terraform_destroy() {
	key_name=$1
	echo "destroy the terraform resource"
	# check if the object still exists
	still_exists=$(aws s3api head-object --bucket soaking-terraform-state --key "${key_name}" || echo '')
	if [ -n "${still_exists}" ]; then
	  platform_folder=""
	  if "${key_name}" =~ "ec2"; then platform_folder="ec2"; else { if "${key_name}" =~ "eks"; then platform_folder="eks"; else platform_folder="ecs"; fi; }; fi
	  echo "${platform_folder}"


	fi
}

s3_keys=$(aws s3api list-objects-v2 --bucket "${terraform_state_s3_bucket}" --query "Contents[?LastModified < '${yesterday}'].Key")

echo "${s3_keys}" | docker run --rm -i stedolan/jq -c -r '.[]' | while read -r s3_key; do
  terraform_destroy "${s3_key}"
done

