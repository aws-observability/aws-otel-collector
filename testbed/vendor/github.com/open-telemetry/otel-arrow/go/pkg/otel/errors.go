/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package otel

import (
	"errors"
)

var (
	ErrMultipleTracesRecords     = errors.New("multiple traces records found")
	ErrMultipleSpanEventsRecords = errors.New("multiple span events records found")
	ErrDuplicatePayloadType      = errors.New("duplicate payload type")
	UnknownPayloadType           = errors.New("unknown payload type")
)
