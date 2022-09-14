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

Write-Output 'Installing ADOTCollector.'
Start-Process msiexec.exe -Wait -ArgumentList '/i aws-otel-collector.msi'

$Cmd = "${Env:ProgramFiles}\Amazon\AWSOTelCollector\aws-otel-collector-ctl.ps1"
$Conf = "${Env:ProgramFiles}\Amazon\AWSOTelCollector\config.yaml"

if (Test-Path Env:SSM_CONFIG) {
    $config = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($Env:SSM_CONFIG))
    $Utf8NoBomEncoding = New-Object System.Text.UTF8Encoding $False
    [System.IO.File]::WriteAllLines($Conf, $config, $Utf8NoBomEncoding)
}

& "${Cmd}" -a start
