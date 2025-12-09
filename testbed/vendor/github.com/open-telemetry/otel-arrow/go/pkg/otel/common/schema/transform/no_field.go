/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package transform

import "github.com/apache/arrow-go/v18/arrow"

// NoField is a FieldTransform that returns nil, so in practice it removes the
// field.
type NoField struct{}

func (t *NoField) Transform(_ *arrow.Field) *arrow.Field {
	return nil
}

func (t *NoField) RevertCounters() {}

func (t *NoField) Path() string {
	return "undefined"
}
