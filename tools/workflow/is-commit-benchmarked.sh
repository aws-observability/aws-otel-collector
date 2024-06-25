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

if [ -z "${COMMIT_SHA}" ]; then
    echo "Must have COMMIT_SHA set"
    exit 1
fi

if [ -z "${GH_PAGES_BRANCH}" ]; then
    GH_PAGES_BRANCH="gh-pages"
fi

HAS_COMMIT=false

git switch "${GH_PAGES_BRANCH}"
if [ -f benchmark/trend/data.js ]; then
    HAS_COMMIT=$(sed "s/window.BENCHMARK_DATA = //" benchmark/trend/data.js | docker run --rm -i ghcr.io/jqlang/jq -c ".entries.Benchmark | any(.commit.id == \"${COMMIT_SHA}\")")
fi
echo "has-commit=${HAS_COMMIT}" >> "$GITHUB_OUTPUT"
git checkout -
