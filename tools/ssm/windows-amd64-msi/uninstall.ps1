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

Set-StrictMode -Version 2.0
$ErrorActionPreference = "Stop"

Write-Output 'Uninstalling ADOTCollector.'

$Cmd = "${Env:ProgramFiles}\Amazon\AWSOTelCollector\aws-otel-collector-ctl.ps1"

if (Test-Path -LiteralPath "${Cmd}" -PathType Leaf) {
    & "${Cmd}" -a stop
}

[System.Management.ManagementObject[]] $agents = Get-WmiObject -Class Win32_Product -ComputerName . -Filter "Name='AWS OTel Collector'"

if($agents) {
    for ($i = 0; $i -lt @($agents).Count; $i ++) {
        @($agents)[$i].Uninstall()
    }
}
