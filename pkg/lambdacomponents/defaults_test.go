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

	"github.com/stretchr/testify/require"
)

const (
	exportersCount = 7
	receiversCount = 1
)

func TestComponents(t *testing.T) {
	factories, err := Components()
	require.NoError(t, err)
	exporters := factories.Exporters
	require.Len(t, exporters, exportersCount)
	// aws exporters
	require.NotNil(t, exporters["awsxray"])
	require.NotNil(t, exporters["awsemf"])
	require.NotNil(t, exporters["awsprometheusremotewrite"])
	// core exporters
	require.NotNil(t, exporters["logging"])
	require.NotNil(t, exporters["otlp"])
	require.NotNil(t, exporters["otlphttp"])
	// other exporters
	require.NotNil(t, exporters["prometheus"])

	receivers := factories.Receivers
	require.Len(t, receivers, receiversCount)
	// core receivers
	require.NotNil(t, receivers["otlp"])
}
