/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package common

import (
	"errors"
)

var (
	ErrInvalidKeyMap         = errors.New("invalid key map")
	ErrUnsupportedCborType   = errors.New("unsupported cbor type")
	ErrInvalidTypeConversion = errors.New("invalid type conversion")

	ErrInvalidSpanIDLength  = errors.New("invalid span id length")
	ErrInvalidTraceIDLength = errors.New("invalid trace id length")

	ErrNotArraySparseUnion = errors.New("not an arrow array.SparseUnion")
	ErrNotArrayMap         = errors.New("not an arrow array.Map")
)
