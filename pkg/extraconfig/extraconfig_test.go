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
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtraConfigBasic(t *testing.T) {
	tests := []struct {
		testname      string
		testFile      string
		expectedError assert.ErrorAssertionFunc
		checkFunc     func(test *testing.T, config *ExtraConfig) bool
	}{
		{
			"get extra config",
			"extraconfig.txt",
			assert.NoError,
			func(test *testing.T, config *ExtraConfig) bool {
				return assert.Equal(test, config.LoggingLevel, "DEBUG")
			},
		},
		{
			"no logging level",
			"extraconfigwithoutcfg.txt",
			assert.NoError,
			func(test *testing.T, config *ExtraConfig) bool {
				return assert.Equal(test, config.LoggingLevel, "")
			},
		},
		{
			"no file",
			"notexistedextraconfig.txt",
			assert.Error,
			func(test *testing.T, config *ExtraConfig) bool {
				return true
			},
		},
		{
			"custom aws profile",
			"extraconfig.txt",
			assert.NoError,
			func(test *testing.T, config *ExtraConfig) bool {
				return assert.Equal(test, config.AwsProfile, "test")
			},
		},
		{
			"no custom profile",
			"extraconfigwithoutcfg.txt",
			assert.NoError,
			func(test *testing.T, config *ExtraConfig) bool {
				return assert.Equal(test, config.AwsProfile, "")
			},
		},
		{
			"custom aws creds file",
			"extraconfig.txt",
			assert.NoError,
			func(test *testing.T, config *ExtraConfig) bool {
				return assert.Equal(test, config.AwsCredentialFile, "~/.aws/credentials")
			},
		},
		{
			"extra config with any var",
			"extraconfig.txt",
			assert.NoError,
			func(test *testing.T, config *ExtraConfig) bool {
				return assert.Equal(test, "anyVal", os.Getenv("Any_Var"))
			},
		},
		{
			"no custom was creds file",
			"extraconfigwithoutcfg.txt",
			assert.NoError,
			func(test *testing.T, config *ExtraConfig) bool {
				return assert.Equal(test, config.AwsCredentialFile, "")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			t.Cleanup(func() {
				os.Clearenv()
			})
			setFilePathHelper(test.testFile)
			extraConfig, err := GetExtraConfig()
			test.expectedError(t, err)
			assert.True(t, test.checkFunc(t, extraConfig))
		})
	}
}

func setFilePathHelper(filename string) {
	fp := filepath.Join(".", "testdata", filename)
	UnixExtraConfigPath = fp
	WindowsExtraConfigPath = fp
}
