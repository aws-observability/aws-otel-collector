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
# This script is used in CI/CD workflow
# for both upload testing packages to s3
# and release packages to s3
# with different s3 bucket as the parameter
#
# Env vars
# 1. s3_bucket_name
# 2. upload_to_latest
##########################################

# check environment vars
if [ -z "${s3_bucket_name}" ]; then
    s3_bucket_name="aws-otel-collector-test"
    echo "use default s3_bucket_name: ${s3_bucket_name}"
else
    echo "load s3_bucket_name from env var: ${s3_bucket_name}"
fi

if [ -z "${upload_to_latest}" ]; then
    upload_to_latest=0
    echo "upload_to_latest is set to 0 by default"
else
    echo "upload_to_latest is set by env var to ${upload_to_latest}"
fi

# define vars
local_packages_home="build/packages"
package_name="aws-otel-collector"

# check if all the required files are there
declare -a required_files=(
    "${local_packages_home}/VERSION"
    "${local_packages_home}/linux/amd64/${package_name}.rpm"
    "${local_packages_home}/linux/arm64/${package_name}.rpm"
    "${local_packages_home}/debian/amd64/${package_name}.deb"
    "${local_packages_home}/debian/arm64/${package_name}.deb"
    "${local_packages_home}/windows/amd64/${package_name}.msi"
)
for i in "${required_files[@]}"; do
    if [ ! -f "${i}" ]; then
        echo "${i} does not exist"
        exit 1
    fi
done

# upload packages to s3
version=$(cat ${local_packages_home}/VERSION)
declare -a local_to_s3_path=(
    "${local_packages_home}/linux/amd64/${package_name}.rpm amazon_linux/amd64/latest/${package_name}.rpm"
    "${local_packages_home}/linux/amd64/${package_name}.rpm redhat/amd64/latest/${package_name}.rpm"
    "${local_packages_home}/linux/amd64/${package_name}.rpm centos/amd64/latest/${package_name}.rpm"
    "${local_packages_home}/linux/amd64/${package_name}.rpm suse/amd64/latest/${package_name}.rpm"
    "${local_packages_home}/debian/amd64/${package_name}.deb ubuntu/amd64/latest/${package_name}.deb"
    "${local_packages_home}/debian/amd64/${package_name}.deb debian/amd64/latest/${package_name}.deb"
    "${local_packages_home}/windows/amd64/${package_name}.msi windows/amd64/latest/${package_name}.msi"
    "${local_packages_home}/linux/arm64/${package_name}.rpm amazon_linux/arm64/latest/${package_name}.rpm"
    "${local_packages_home}/linux/arm64/${package_name}.rpm redhat/arm64/latest/${package_name}.rpm"
    "${local_packages_home}/linux/arm64/${package_name}.rpm centos/arm64/latest/${package_name}.rpm"
    "${local_packages_home}/linux/arm64/${package_name}.rpm suse/arm64/latest/${package_name}.rpm"
    "${local_packages_home}/debian/arm64/${package_name}.deb ubuntu/arm64/latest/${package_name}.deb"
    "${local_packages_home}/debian/arm64/${package_name}.deb debian/arm64/latest/${package_name}.deb"
)
for i in "${local_to_s3_path[@]}"; do
    local_path=$(echo "${i}" | awk -F ' ' '{print $1}')
    s3_latest_key=$(echo "${i}" | awk -F ' ' '{print $2}')
    s3_version_key="${s3_latest_key/latest/${version}}"
    s3_latest_url="s3://${s3_bucket_name}/${s3_latest_key}"
    s3_version_url="s3://${s3_bucket_name}/${s3_version_key}"

    echo "check if package is already there: ${s3_version_url}"
    aws s3api head-object --bucket "${s3_bucket_name}" --key "${s3_version_key}" >/dev/null || version_not_exist=true

    if [ "${version_not_exist}" ]; then
        echo "upload package ${local_path} to ${s3_version_url}"
        aws s3 cp "${local_path}" "${s3_version_url}" --acl public-read
    else
        echo "package ${s3_version_url} is already there, skip it"
    fi

    if [ "$upload_to_latest" -eq 1 ]; then
        echo "upload package ${local_path} to ${s3_latest_url}"
        aws s3 cp "${local_path}" "${s3_latest_url}" --acl public-read
    else
        echo "skip publishing ${s3_latest_url}"
    fi

done
