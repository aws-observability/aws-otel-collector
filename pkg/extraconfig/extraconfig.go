/*
 * Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package extraconfig

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

var (
	unixExtraConfigPath    = "/opt/aws/aws-otel-collector/etc/.env"
	windowsExtraConfigPath = "C:\\ProgramData\\Amazon\\AwsOTelCollector\\\\.env"
)

type ExtraConfig struct {
	LoggingLevel string
}

// GetExtraConfig returns the extra configs.
func GetExtraConfig() (*ExtraConfig, error) {
	// read .env from os
	file, err := os.Open(getConfigFilePath())
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read its content line by line and handle it as keyvalue pairs
	extraConfigMap := map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				extraConfigMap[key] = value

			}
		}
	}

	// check error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// build extraconfig object
	extraConfig := ExtraConfig{}
	if val, ok := extraConfigMap["loggingLevel"]; ok {
		extraConfig.LoggingLevel = val
	}

	return &extraConfig, nil
}

// getConfigFilePath return the path base on os
func getConfigFilePath() string {
	if runtime.GOOS == "windows" {
		return windowsExtraConfigPath
	}
	return unixExtraConfigPath
}
