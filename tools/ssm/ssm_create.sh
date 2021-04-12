#!/usr/bin/env bash
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

set -eu

PKG_NAME=$1
REL_VERSION=$2
S3_BUCKET=$3

output=$(aws ssm list-documents --filter Key=Owner,Values=Self Key=DocumentType,Values=Package Key=Name,Values=${PKG_NAME})
echo ${output}
status=$(echo ${output} | python3 -c "import sys, json; print('empty') if len(json.load(sys.stdin)['DocumentIdentifiers']) == 0 else print('exist')")
if [[ ${status} == "empty" ]]
then
    aws ssm create-document \
        --name "${PKG_NAME}" \
        --content file://build/packages/ssm/manifest.json \
        --attachments Key="SourceUrl",Values="https://s3.amazonaws.com/${S3_BUCKET}" \
        --version-name ${REL_VERSION} \
        --document-type Package
else
    output=$(aws ssm update-document \
        --name "${PKG_NAME}" \
        --content file://build/packages/ssm/manifest.json \
        --attachments Key="SourceUrl",Values="https://s3.amazonaws.com/${S3_BUCKET}" \
        --version-name ${REL_VERSION} \
        --document-version "\$LATEST")

    echo ${output}
    last_version=$(echo ${output} | python3 -c "import sys, json; print(json.load(sys.stdin)['DocumentDescription']['LatestVersion'])")
    echo ${last_version}

    aws ssm update-document-default-version \
        --name "${PKG_NAME}" \
        --document-version "${last_version}"
fi
