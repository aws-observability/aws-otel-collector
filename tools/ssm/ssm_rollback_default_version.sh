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
    test -f "$(which aws)" || error_exit "aws command not detected in path, please install it"
    test -f "$(which docker)" || error_exit "docker command not detected in path, please install it"
    test -f "$(which sed)" || error_exit "sed command not detected in path, please install it"
    test -f "$(which sort)" || error_exit "sort command not detected in path, please install it"
}

function parse_environment_input() {
    if [[ -z "${version}" ]]; then
        error_exit "Missing input for flag version"
    fi

    if [[ -z "${ssm_package_name}" ]]; then
        error_exit "Missing input for flag ssm_package_name"
    fi
}

function rollback_second_largest_target_document_as_default() {
    document_version_names=$(aws ssm list-document-versions --name "${ssm_package_name}" --output json | docker run --rm -i ghcr.io/jqlang/jq -c -r ".DocumentVersions|.[]|.VersionName" | sed 's/"//g' | sort -V -r)

    # Convert document versions name to array
    # shellcheck disable=SC2206
    document_versions_array=($document_version_names)

    latest_document_version_name=${document_versions_array[0]}
    second_latest_document_version_name=${document_versions_array[1]}

    if [[ -z ${latest_document_version_name} || -z ${second_latest_document_version_name} ]]; then
        error_exit "Cannot find rollback document version."
    fi

    if [[ ${latest_document_version_name} == "${version}" ]]; then
      rollback_version_name="${second_latest_document_version_name}"
    else
      rollback_version_name="${latest_document_version_name}"
    fi

    echo "Roll-back the default SSM version to ${rollback_version_name}"

    rollback_document_version=$(aws ssm describe-document --name "${ssm_package_name}" --version-name "${rollback_version_name}" --output json | docker run --rm -i ghcr.io/jqlang/jq -c -r ".Document|.DocumentVersion")
    aws ssm update-document-default-version --name "${ssm_package_name}" --document-version "${rollback_document_version}"
}

check_deps
parse_environment_input
rollback_second_largest_target_document_as_default
