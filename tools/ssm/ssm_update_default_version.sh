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

set -e

##########################################
# This script is used in CD workflow
# to make the second largest version
# in SSM package as default
# Error: https://github.com/aws-observability/aws-otel-collector/runs/5014507743?check_suite_focus=true
##########################################

function error_exit() {
    echo "$1" 1>&2
    exit 1
}

function check_deps() {
    test -f $(which aws) || error_exit "aws command not detected in path, please install it"
    test -f $(which docker) || error_exit "docker command not detected in path, please install it"
}

function parse_environment_input() {
    if [[ -z "${version}" ]]; then
        error_exit "Missing input for flag version"
    fi

    if [[ -z "${ssm_package_name}" ]]; then
        error_exit "Missing input for flag ssm_package_name"
    fi
}

function update_second_largest_target_document_as_default() {
    list_document_versions=$(aws ssm list-document-versions --name "${ssm_package_name}" --output json)
    target_document_version=$(echo $list_document_versions | docker run --rm -i stedolan/jq -c -r ".DocumentVersions|.[]|select(.VersionName==\"${version}\")|.DocumentVersion")

    if [[ -z $target_document_version ]]; then
        error_exit "Document version not found for ${version}"
    fi

    second_largest_document_version="$((target_document_version - 1))"
    aws ssm describe-document --name ${ssm_package_name} --document-version "${second_largest_document_version}" >/dev/null &&
        aws ssm update-document-default-version --name ${ssm_package_name} --document-version "${second_largest_document_version}"

}

check_deps
parse_environment_input
update_second_largest_target_document_as_default
