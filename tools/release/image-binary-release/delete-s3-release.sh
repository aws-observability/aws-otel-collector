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
# to delete binaries which were uploaded to S3
# when the validation test in CD are failing
##########################################

#Define variables to use in the environment
package_name="aws-otel-collector"

function error_exit() {
    echo "$1" 1>&2
    exit 1
}

function check_deps() {
    test -f "$(which aws)" || error_exit "aws command not detected in path, please install it"
}

function parse_environment_input() {
    if [[ -z "${version}" ]]; then
        error_exit "Missing input for flag version"
    fi

    if [[ -z "${s3_bucket_name}" ]]; then
        error_exit "Missing input for flag s3_bucket_name"
    fi

    if [ -z "${delete_to_latest}" ]; then
        delete_to_latest=0
        echo "Flag delete_to_latest is set to 0 by default"
    else
        echo "Flag delete_to_latest is set by env var to ${delete_to_latest}"
    fi
}

function check_exist_and_delete_object_s3() {
    s3_key=$1
    s3_url=$2

    echo "Check if package is there: ${s3_url}"
    aws s3api head-object --bucket "${s3_bucket_name}" --key "${s3_key}" >/dev/null || not_exist=true

    if [ "${not_exist}" ]; then
        echo "Skip delete since package ${s3_url} is not there to delete"
    else
        echo "Begin to delete ${s3_url}"
        aws s3 rm "${s3_url}"
    fi
}

function delete_s3_objects_from_s3_bucket() {
    # Get the path key for each binary and delete based on these keys
    declare -a s3_path=(
        "amazon_linux/amd64/latest/${package_name}.rpm"
        "redhat/amd64/latest/${package_name}.rpm"
        "centos/amd64/latest/${package_name}.rpm"
        "suse/amd64/latest/${package_name}.rpm"
        "ubuntu/amd64/latest/${package_name}.deb"
        "debian/amd64/latest/${package_name}.deb"
        "windows/amd64/latest/${package_name}.msi"
        "amazon_linux/arm64/latest/${package_name}.rpm"
        "redhat/arm64/latest/${package_name}.rpm"
        "centos/arm64/latest/${package_name}.rpm"
        "suse/arm64/latest/${package_name}.rpm"
        "ubuntu/arm64/latest/${package_name}.deb"
        "debian/arm64/latest/${package_name}.deb"
    )

    for i in "${s3_path[@]}"; do
        s3_latest_key="${i}"
        s3_version_key="${s3_latest_key//latest/${version}}"
        s3_version_url="s3://${s3_bucket_name}/${s3_version_key}"
        s3_latest_url="s3://${s3_bucket_name}/${s3_latest_key}"

        check_exist_and_delete_object_s3 "${s3_version_key}" "${s3_version_url}"

        if [ "$delete_to_latest" -eq 1 ]; then
            check_exist_and_delete_object_s3 "${s3_latest_key}" "${s3_latest_url}"
        fi
    done

    echo "Finish deleting script"

}

check_deps
parse_environment_input
delete_s3_objects_from_s3_bucket
