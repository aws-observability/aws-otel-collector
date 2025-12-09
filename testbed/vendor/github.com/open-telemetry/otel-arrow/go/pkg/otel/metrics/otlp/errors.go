/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package otlp

import (
	"errors"
)

var (
	ErrNotArrayInt32   = errors.New("not an arrow array.Int32")
	ErrNotArrayUint64  = errors.New("not an arrow array.Uint64")
	ErrNotArrayFloat64 = errors.New("not an arrow array.Float64")
	ErrNotArrayList    = errors.New("not an arrow array.List")
)
