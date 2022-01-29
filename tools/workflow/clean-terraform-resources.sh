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

# ensure we are using gnudate
yesterday=$(docker run --rm ubuntu date -d "yesterday" "+%Y-%m-%d")

terraform_destroy() {
     key_name=$1
     echo "destroy the terraform resource"
     # check if the object still exists
     still_exists=$(aws s3api head-object --bucket soaking-terraform-state --key "${key_name}" || echo '')
     if [ -n "${still_exists}" ]; then
          echo "download s3 object: soaking-terraform-state/${key_name}"
          aws s3 cp "s3://soaking-terraform-state/${key_name}" "downloaded_terraform/${key_name}"
          tar xvf "downloaded_terraform/${key_name}"

          # Set output in case the destroy fails
          echo "::set-output name=key-name::${key_name}"
          echo "remove s3 key: ${key_name}"
          aws s3 rm "s3://soaking-terraform-state/${key_name}"

          if [ -d testing-framework/terraform/soaking ]; then
               cd testing-framework/terraform/soaking
               terraform init
               terraform destroy -auto-approve
               cd -
          fi
          rm -rf testing-framework
          rm -rf "downloaded_terraform/${key_name}"
     fi
}

s3_keys=$(aws s3api list-objects-v2 --bucket "soaking-terraform-state" --query "Contents[?LastModified < '${yesterday}'].Key")

echo "${s3_keys}" | docker run --rm -i stedolan/jq -c -r '.[]' | while read -r s3_key; do
     terraform_destroy "${s3_key}"
done
