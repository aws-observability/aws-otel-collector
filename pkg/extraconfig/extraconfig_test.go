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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExtraConfig(t *testing.T) {
	unixExtraConfigPath = "./test/extraconfig.txt"
	extraConfig, err := GetExtraConfig()

	assert.NoError(t, err)
	assert.Equal(t, extraConfig.LoggingLevel, "DEBUG")
}

func TestGetExtraConfigWithNoLoggingLevel(t *testing.T) {
	unixExtraConfigPath = "./test/extraconfigwithoutcfg.txt"
	extraConfig, err := GetExtraConfig()

	assert.NoError(t, err)
	assert.Equal(t, extraConfig.LoggingLevel, "")
}

func TestGetExtraConfigWithNoFile(t *testing.T) {
	unixExtraConfigPath = "./notexistedextraconfig.txt"
	_, err := GetExtraConfig()
	assert.Error(t, err)
}

func TestGetExtraConfigWithCustomAwsProfile(t *testing.T) {
	unixExtraConfigPath = "./test/extraconfig.txt"
	defer os.Unsetenv("AWS_PROFILE")
	_, _ = GetExtraConfig()
	profileName := os.Getenv("AWS_PROFILE")
	assert.Equal(t, "test", profileName)
}

func TestGetExtraConfigWithoutCustomAwsProfile(t *testing.T) {
	unixExtraConfigPath = "./test/extraconfigwithoutcfg.txt"
	defer os.Unsetenv("AWS_PROFILE")
	_, _ = GetExtraConfig()
	profileName := os.Getenv("AWS_PROFILE")
	assert.Equal(t, "", profileName)
}

func TestGetExtraConfigWithCustomAwsCredsFile(t *testing.T) {
	unixExtraConfigPath = "./test/extraconfig.txt"
	defer os.Unsetenv("AWS_CREDENTIAL_PROFILES_FILE")
	_, _ = GetExtraConfig()
	awsCredsFile := os.Getenv("AWS_CREDENTIAL_PROFILES_FILE")
	assert.Equal(t, "~/.aws/credentials", awsCredsFile)
}

func TestGetExtraConfigWithoutCustomAwsCredsFile(t *testing.T) {
	unixExtraConfigPath = "./test/extraconfigwithoutcfg.txt"
	defer os.Unsetenv("AWS_CREDENTIAL_PROFILES_FILE")
	_, _ = GetExtraConfig()
	awsCredsFile := os.Getenv("AWS_CREDENTIAL_PROFILES_FILE")
	assert.Equal(t, "", awsCredsFile)
}
