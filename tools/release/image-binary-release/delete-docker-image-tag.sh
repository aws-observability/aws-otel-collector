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

#Get access token from docker hub to authorize when deleting
dockerhub_token=$(curl -s -X "POST" \
  -H "Content-Type: application/json" \
  -d "{\"username\": \"${dockerhub_username}\", \"password\": \"${dockerhub_password}\"}" \
  "${dockerhub_url}/v2/users/login/" | jq -r .token)

#Checking if the input username or input password is invalid (faster failing)
if [ ${dockerhub_token} == "null" ];  then
  echo "Username or password is not valid."
  exit 1
fi

#Delete the image from docker hub's registry
curl -i -L -X "DELETE" \
  -H "Accept: application/json" \
  -H "Authorization: JWT $dockerhub_token" \
  "${dockerhub_url}/v2/repositories/${image_namespace}/${image_repo}/tags/${version}/"