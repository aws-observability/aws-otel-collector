/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
	extraConfig, err := GetExtraConfig()
	assert.NoError(t, err)
	assert.Equal(t, extraConfig.AwsProfile, "test")
}

func TestGetExtraConfigWithoutCustomAwsProfile(t *testing.T) {
	unixExtraConfigPath = "./test/extraconfigwithoutcfg.txt"
	extraConfig, err := GetExtraConfig()
	assert.NoError(t, err)
	assert.Equal(t, extraConfig.AwsProfile, "")
}

func TestGetExtraConfigWithCustomAwsCredsFile(t *testing.T) {
	unixExtraConfigPath = "./test/extraconfig.txt"
	extraConfig, err := GetExtraConfig()
	assert.NoError(t, err)
	assert.Equal(t, extraConfig.AwsCredentialFile, "~/.aws/credentials")
}

func TestGetExtraConfigWithoutCustomAwsCredsFile(t *testing.T) {
	unixExtraConfigPath = "./test/extraconfigwithoutcfg.txt"
	extraConfig, err := GetExtraConfig()
	assert.NoError(t, err)
	assert.Equal(t, extraConfig.AwsCredentialFile, "")
}
