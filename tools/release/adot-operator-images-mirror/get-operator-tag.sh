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

COLLECTOR_TAG=$(cat go.mod | grep "go.opentelemetry.io/collector " | cut -d " " -f2)
COLLECTOR_PATCH_VER=${COLLECTOR_TAG: -1}
COLLECTOR_TAG_WITHOUT_PATCH_VER=$(echo $COLLECTOR_TAG | sed 's/.$//')

OPERATOR_TAGS=$(echo $(curl https://api.github.com/repos/open-telemetry/opentelemetry-operator/tags) | jq  '.[] | .name')

if [[ $COLLECTOR_PATCH_VER == "0" ]]; then
    # Check if any Operator image exists for the Collector image
    if echo $OPERATOR_TAGS | grep -q $COLLECTOR_TAG_WITHOUT_PATCH_VER; then
        # Checks if there is more than one operator tag that has the same major/minor version; if only one, then there are no operator patches
        if [[ $(echo "$OPERATOR_TAGS" | grep "$COLLECTOR_TAG_WITHOUT_PATCH_VER" -c) == "1" ]]; then
            # NO COLLECTOR PATCH, NO PATCH FOR OPERATOR IMAGE - SAME IMAGE TAG USED
            echo "::set-output name=operator-tag::$COLLECTOR_TAG"
        else
            # NO COLLECTOR PATCH, PATCH FOR OPERATOR IMAGE - IMAGE TAG USED == vx.x.n
            #  retrives 1st item from list of Operator tags (most recent patch) and cleans it up with cut/sed to get the tag
            NEW_OPERATOR_TAG=$(echo "$OPERATOR_TAGS" | grep "$COLLECTOR_TAG_WITHOUT_PATCH_VER" -m 1 | sed 's|[",]||g')
            echo "::set-output name=operator-tag::$NEW_OPERATOR_TAG"
        fi
    else
        # NO COLLECTOR PATCH, NO OPERATOR IMAGE - FAIL WORKFLOW
        echo "NO OPERATOR IMAGE EXISTS FOR THIS COLLECTOR IMAGE"
        exit 1
    fi
else
    # Check if any Operator image exists for the Collector image
    if echo $OPERATOR_TAGS | grep -q $COLLECTOR_TAG_WITHOUT_PATCH_VER; then
        # Checks if there is more than one operator tag that has the same major/minor version; if only one, then there are no operator patches
        if [[ $(echo "$OPERATOR_TAGS" | grep "$COLLECTOR_TAG_WITHOUT_PATCH_VER" -c) == "1" ]]; then
            # COLLECTOR PATCH, NO PATCH FOR OPERATOR IMAGE - IMAGE TAG USED == vx.x.0
            # NEW_OPERATOR_TAG = "vx.x." + "0" = "vx.x.0"
            NEW_OPERATOR_TAG="${COLLECTOR_TAG_WITHOUT_PATCH_VER}0"
            echo "::set-output name=operator-tag::$NEW_OPERATOR_TAG"
        else
            # COLLECTOR PATCH, PATCH FOR OPERATOR IMAGE - SAME IMAGE TAG USED
            echo "::set-output name=operator-tag::$COLLECTOR_TAG"
        fi
    else
        # COLLECTOR PATCH, NO OPERATOR IMAGE - FAIL WORKFLOW
        echo "NO OPERATOR IMAGE EXISTS FOR THIS COLLECTOR IMAGE"
        exit 1
    fi
fi
