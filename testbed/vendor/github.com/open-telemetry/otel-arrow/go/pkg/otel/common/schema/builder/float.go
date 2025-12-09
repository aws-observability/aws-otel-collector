/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package builder

import (
	"github.com/apache/arrow-go/v18/arrow/array"

	"github.com/open-telemetry/otel-arrow/go/pkg/otel/common/schema"
	"github.com/open-telemetry/otel-arrow/go/pkg/otel/common/schema/update"
)

// Float64Builder is a wrapper around the arrow array builder for float64.
type Float64Builder struct {
	builder       *array.Float64Builder
	transformNode *schema.TransformNode
	updateRequest *update.SchemaUpdateRequest
}

// Append adds a value to the underlying builder or updates the transform node
// if the builder is nil.
func (b *Float64Builder) Append(value float64) {
	if b.builder != nil {
		b.builder.Append(value)
		return
	}

	if b.updateRequest != nil {
		// If the builder is nil, then the transform node is not optional.
		b.transformNode.RemoveOptional()
		b.updateRequest.Inc(&update.NewFieldEvent{FieldName: b.transformNode.Path()})
		b.updateRequest = nil // No need to report this again.
	}
}

// AppendNonZero adds a value to the underlying builder or updates the transform node
// if the builder is nil.
// Note: 0 values are not appended to the builder.
func (b *Float64Builder) AppendNonZero(value float64) {
	if b.builder != nil {
		if value == 0 {
			b.builder.AppendNull()
			return
		}

		b.builder.Append(value)
		return
	}

	if value != 0 && b.updateRequest != nil {
		// If the builder is nil, then the transform node is not optional.
		b.transformNode.RemoveOptional()
		b.updateRequest.Inc(&update.NewFieldEvent{FieldName: b.transformNode.Path()})
		b.updateRequest = nil // No need to report this again.
	}
}

// AppendNull adds a null value to the underlying builder. If the builder is
// nil we do nothing as we have no information about the presence of this field
// in the data.
func (b *Float64Builder) AppendNull() {
	if b.builder != nil {
		b.builder.AppendNull()
		return
	}
}
