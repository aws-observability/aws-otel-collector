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

package lambdacomponents

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponents(t *testing.T) {
	factories, err := Components()
	require.NoError(t, err)
	exporters := factories.Exporters
	// aws exporters
	assert.True(t, exporters["awsxray"] != nil)
	assert.True(t, exporters["awsprometheusremotewrite"] != nil)

	// core exporters
	assert.True(t, exporters["logging"] != nil)
	assert.True(t, exporters["otlphttp"] != nil)

	receivers := factories.Receivers
	assert.True(t, receivers["otlp"] != nil)
}
