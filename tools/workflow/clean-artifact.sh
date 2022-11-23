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

# this is the command to send the dispatch event to the nightly-clean-artifact workflow
# please specify the TOKEN with the github token
# example of how to use it: TOKEN=e7*******************************3d0d ./clean-artifact.sh

OWNER=mxiamxia
REPO=aws-opentelemetry-collector

curl -u="${OWNER}:${TOKEN}" -X POST --data '{"event_type": "clean-artifacts"}' "https://api.github.com/repos/${OWNER}/${REPO}/dispatches"
