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

#! /bin/bash

## this is the script to bump the version and create the release note
## please run this script whenever you want to bump the version instead of directly modifying the VERSION file
## below is an example to run this script: 
## RELEASE_VERSION=v0.1.8 GITHUB_TOKEN=e75***********fa3d0d ./tools/release/bump-version-and-create-release-note.sh

# get the current version
VERSION=`cat VERSION`
OUTPUT="docs/releases/${RELEASE_VERSION}.md"

# generate release note
docker run --rm -it -v "`pwd`":/usr/local/src/your-app ferrarimarco/github-changelog-generator \
  --user aws-observability \
	--project aws-otel-collector \
	-t ${GITHUB_TOKEN} \
	--since-tag ${VERSION} \
	--future-release ${RELEASE_VERSION} \
	--output ${OUTPUT} \
	--exclude-labels bumpversion

# bump the version
echo ${RELEASE_VERSION} > VERSION

# git commit
git add VERSION docs/releases/${RELEASE_VERSION}.md
git commit -m "bump version to ${RELEASE_VERSION}"
