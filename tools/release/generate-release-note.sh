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

VERSION=$(cat VERSION)

sed "s/__VERSION__/$VERSION/g" tools/release/header.md.template >header
sed "s/__VERSION__/$VERSION/g" tools/release/downloading-links.md.template >downloading-links
# Formats changelog sections: **<label>:** becomes ### <label>
# Skips the first 4 lines (title and version number)
sed -e 1,4d -e "s/^\*\*\(.*\)\:\*\*/### \1/g" "docs/releases/${VERSION}.md" >changelog

cat header changelog downloading-links >release-note
