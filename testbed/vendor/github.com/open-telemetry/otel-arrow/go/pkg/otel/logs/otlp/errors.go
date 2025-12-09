/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package otlp

import "errors"

var (
	ErrBodyNotSparseUnion = errors.New("body is not a sparse union")
)
