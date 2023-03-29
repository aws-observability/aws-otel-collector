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

ECR_TAGS=$(aws ecr-public describe-image-tags --repository-name adot-operator --region us-east-1)
OPERATOR_TAGS=$(curl https://api.github.com/repos/open-telemetry/opentelemetry-operator/tags)

if grep -q "$VERSION" <<< "$OPERATOR_TAGS" && ! ( grep -q "$VERSION" <<< "$ECR_TAGS" ); then
        echo "update-operator=true" >> "$GITHUB_OUTPUT"
else
        echo "update-operator=false" >> "$GITHUB_OUTPUT"
fi
