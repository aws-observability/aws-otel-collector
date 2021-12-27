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
# to delete image tags from docker hub registry
# when the validation test in CD are failing
##########################################

#Define variables to use in the environment
#version=`cat ${local_packages_home}/VERSION`
dockerhub_url="https://hub.docker.com"

function error_exit() {
  echo "$1" 1>&2
  exit 1
}

function check_deps() {
  test -f $(which jq) || error_exit "jq command not detected in path, please install it"
  test -f $(which curl) || error_exit "curl command not detected in path, please install it"
}

function log_in_dockerhub(){
  #Get access token from docker hub to authorize when deleting
  dockerhub_log_in_respond=$(curl -s -X "POST" \
    --retry 2 --retry-delay 5 --retry-max-time 30 --connect-timeout 5 --max-time 10 \
    -H "Content-Type: application/json" \
    -d "{\"username\": \"${dockerhub_username}\", \"password\": \"${dockerhub_password}\"}" \
    "${dockerhub_url}/v2/users/login/")
  dockerhub_log_in_error=$(jq -r '.detail' <<< ${dockerhub_log_in_respond})

  #Checking if there is any error when delete (faster failing)
  if [[ "${dockerhub_log_in_error}" != 'null'  ]]; then
    error_exit "Cannot log in because of error: ${dockerhub_log_in_error}"
  fi

  dockerhub_token=$(jq -r '.token' <<< ${dockerhub_log_in_respond})

}

function delete_image_from_dockerhub(){
  #Delete the image from docker hub's registry
  dockerhub_delete_respond=$(curl -L -X "DELETE" \
    --retry 2 --retry-delay 5 --retry-max-time 30 --connect-timeout 5 --max-time 10 \
    -H "Accept: application/json" \
    -H "Authorization: JWT $dockerhub_token" \
    "${dockerhub_url}/v2/repositories/${image_namespace}/${image_repo}/tags/${version}/" | jq -r '.detail')

  if [[ "${dockerhub_delete_respond}" != ''  ]]; then
    error_exit "Cannot delete because of error: ${dockerhub_delete_respond}"
    exit 1
  fi

  echo "Delete successfully the image ${image_namespace}/${image_repo}:${version}"
}

check_deps
log_in_dockerhub
delete_image_from_dockerhub




