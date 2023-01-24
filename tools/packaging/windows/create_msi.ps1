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

# Install the required software
mkdir -p .\build\packages\windows\amd64\
choco install wixtoolset --force -y
$env:Path += ";C:\Program Files (x86)\WiX Toolset v3.11\bin"
refreshenv

# export release version number
export RELEASE_VERSION=$( cat VERSION | cut -c 2- )

# create msi
candle -arch x64  .\tools\packaging\windows\aws-otel-collector.wxs
light .\aws-otel-collector.wixobj

mv aws-otel-collector.msi .\build\packages\windows\amd64\
