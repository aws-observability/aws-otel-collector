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

COLLECTOR_VERSION_FULL=$(grep "go.opentelemetry.io/collector " ../../../go.mod | cut -d " " -f 2)
COLLECTOR_VERSION_MAJOR_MINOR=$(echo "$COLLECTOR_VERSION_FULL" | cut -d "." -f 1,2)

OPERATOR_TAG_FULL=$(curl https://api.github.com/repos/open-telemetry/opentelemetry-operator/tags | jq -r "sort_by(.name) | reverse | map(select( .name | startswith(\"${COLLECTOR_VERSION_MAJOR_MINOR}\"))) | first | .name")
OPERATOR_TAG_MAJOR_MINOR=$(echo "$OPERATOR_TAG_FULL" | cut -d "." -f 1,2)

if [ "$OPERATOR_TAG_MAJOR_MINOR" = "$COLLECTOR_VERSION_MAJOR_MINOR" ]; then
    echo "operator-tag=$OPERATOR_TAG_FULL"
else
    echo "NO OPERATOR IMAGE EXISTS FOR THIS COLLECTOR IMAGE"
    exit 1
fi
