#!/bin/sh

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

BINARY=aoc_linux_x86_64
if [[ $TARGETPLATFORM == "linux/arm64" ]]; then
  BINARY=aoc_linux_aarch64
fi
echo binary name is ${BINARY}
mkdir -p /workspace/build/linux
if [[ ! -f /workspace/${BINARY} ]]; then
  echo "Binary ${BINARY} does not exist."
  exit 1
fi
cp /workspace/${BINARY} /workspace/build/linux/aoc_linux
