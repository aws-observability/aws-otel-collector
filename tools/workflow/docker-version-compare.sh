#!/bin/bash

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
# to check to see what type of version
# release is being done
#
# Env vars
# 1. TARGET_VERSION
# 2. REPO_NAME
#
# Step outputs
# (All booleans: true/false)
# 1. latest-or-newer
# 2. major-update
# 3. minor-update
# 4. patch-update
##########################################

# splits the semver v{major}.{minor}.{patch}
split_version() {
    if [[ $1 =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
        echo "${BASH_REMATCH[1]} ${BASH_REMATCH[2]} ${BASH_REMATCH[3]}"
    fi
}

# check environment vars
if [ -z "${TARGET_VERSION}" ]; then
    echo "Must have TARGET_VERSION set"
    exit 1
fi

if [ -z "${REPO_NAME}" ]; then
    REPO_NAME="amazon/aws-otel-collector"
fi

TOKEN=$(curl -s "https://auth.docker.io/token?service=registry.docker.io&scope=repository:${REPO_NAME}:pull" | docker run --rm -i ghcr.io/jqlang/jq -c ".token" | sed 's/"//g')
# retrieves all the available tags and sorts them in reverse version order
TAGS=$(curl -s "https://registry.hub.docker.com/v2/${REPO_NAME}/tags/list" -H "Authorization: Bearer ${TOKEN}" | docker run --rm -i ghcr.io/jqlang/jq -c ".tags[]" | sed 's/"//g' | sort -V -r)
# get the first tag in the reversed list
LATEST_VERSION=$(echo "${TAGS}" | sed -n '1p')

echo "Comparing ${LATEST_VERSION}(latest) / ${TARGET_VERSION}(target)"

# shellcheck disable=SC2207
LATEST_PARTS=($(split_version "${LATEST_VERSION}"))
# shellcheck disable=SC2207
TARGET_PARTS=($(split_version "${TARGET_VERSION}"))
# shellcheck disable=SC2128
if [ -z "${LATEST_PARTS}" ] || [ -z "${TARGET_PARTS}" ]; then
    echo "Unable to split versions: ${LATEST_VERSION}(latest) / ${TARGET_VERSION}(target)"
    exit 1
fi

LATEST_MAJOR=${LATEST_PARTS[0]}
LATEST_MINOR=${LATEST_PARTS[1]}
LATEST_PATCH=${LATEST_PARTS[2]}

TARGET_MAJOR=${TARGET_PARTS[0]}
TARGET_MINOR=${TARGET_PARTS[1]}
TARGET_PATCH=${TARGET_PARTS[2]}

MAJOR_UPDATE=false
MINOR_UPDATE=false
PATCH_UPDATE=false

if [ "${TARGET_MAJOR}" -gt "${LATEST_MAJOR}" ]; then
    MAJOR_UPDATE=true
elif [ "${TARGET_MAJOR}" -eq "${LATEST_MAJOR}" ]; then
    if [ "${TARGET_MINOR}" -gt "${LATEST_MINOR}" ]; then
        MINOR_UPDATE=true
    elif [ "${TARGET_MINOR}" -eq "${LATEST_MINOR}" ]; then
        if [ "${TARGET_PATCH}" -gt "${LATEST_PATCH}" ]; then
            PATCH_UPDATE=true
        elif [ "${TARGET_PATCH}" -eq "${LATEST_PATCH}" ]; then
            SAME_VERSION=true
        fi
    fi
fi

# If any of the updates are true, then LATEST_OR_NEWER = true
# LATEST_OR_NEWER would be false if the target version is less than the current latest version
if [ "${MAJOR_UPDATE}" == "true" ] || [ "${MINOR_UPDATE}" == "true" ] || [ "${PATCH_UPDATE}" == "true" ] || [ "${SAME_VERSION}" == "true" ]; then 
    LATEST_OR_NEWER=true
else
    LATEST_OR_NEWER=false
fi

{
  echo "major-update=${MAJOR_UPDATE}"
  echo "minor-update=${MINOR_UPDATE}"
  echo "patch-update=${PATCH_UPDATE}"
  echo "latest-or-newer=${LATEST_OR_NEWER}"
  echo "same-version=${SAME_VERSION}"
} >> "$GITHUB_OUTPUT"

