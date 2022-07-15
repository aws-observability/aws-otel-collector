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

COLLECTOR_TAG=$(cat go.mod | grep "go.opentelemetry.io/collector " | cut -d " " -f 2 | cut -d "." -f 1,2)

OPERATOR_TAG=$(curl https://api.github.com/repos/open-telemetry/opentelemetry-operator/tags | jq -r "map(select( .name | startswith(\"${COLLECTOR_TAG}\"))) | first | .name")

if [[ $OPERATOR_TAG == "" ]]; then
    echo "NO OPERATOR IMAGE EXISTS FOR THIS COLLECTOR IMAGE"
    exit 1
else
    echo "::set-output name=operator-tag::$OPERATOR_TAG"
fi
