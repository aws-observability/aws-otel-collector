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

bucket_name="soaking-terraform-state" 
# ensure we are using gnudate
yesterday=$(docker run ubuntu date -d 'yesterday' '+%Y-%m-%d')

terraform_destroy() {
	key_name=${1}
	echo "destroy the terraform resource"
	echo "download s3 object: ${bucket_name}/${key_name}"
	aws s3 cp "s3://${bucket_name}/${key_name}" "downloaded_terraform/${key_name}"
	tar xvf "downloaded_terraform/${key_name}"
	if [ -d testing-framework/terraform/soaking ]; then
		cd testing-framework/terraform/soaking
		terraform init
		terraform destroy -auto-approve
		cd -
	fi	
	echo "remove s3 key: ${key_name}"
	# use hard code bucket name here in case we mistakenly delete false bucket
	aws s3 rm s3://soaking-terraform-state/"${key_name}"
	rm -rf testing-framework
}

s3_keys=$(aws s3api list-objects-v2 --bucket ${bucket_name} --query "Contents[?LastModified < '${yesterday}'].Key")

echo "${s3_keys}" | docker run --rm -i stedolan/jq -c '.[]' | while read -r i; do
	key_name=${i:1:$((${#i}-2))}
	terraform_destroy "${key_name}"
done 

