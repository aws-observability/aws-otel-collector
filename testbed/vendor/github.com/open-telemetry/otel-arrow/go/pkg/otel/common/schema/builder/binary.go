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

// BinaryBuilder is a wrapper around the arrow BinaryBuilder.
type BinaryBuilder struct {
	builder       array.Builder
	transformNode *schema.TransformNode
	updateRequest *update.SchemaUpdateRequest
}

// Append appends a value to the underlying builder and updates the
// transform node if the builder is nil.
func (b *BinaryBuilder) Append(value []byte) {
	if b.builder != nil {
		if value == nil {
			b.builder.AppendNull()
			return
		}

		switch builder := b.builder.(type) {
		case *array.BinaryBuilder:
			builder.Append(value)
		case *array.BinaryDictionaryBuilder:
			if err := builder.Append(value); err != nil {
				// Should never happen.
				panic(err)
			}
		default:
			// Should never happen.
			panic("unknown builder type")
		}
		return
	}

	if b.updateRequest != nil {
		// If the builder is nil, then the transform node is not optional.
		b.transformNode.RemoveOptional()
		b.updateRequest.Inc(&update.NewFieldEvent{FieldName: b.transformNode.Path()})
		b.updateRequest = nil // No need to report this again.
	}
}

// AppendNonNil appends a value to the underlying builder and updates the
// transform node if the builder is nil.
// Note: nil values are not appended to the builder.
func (b *BinaryBuilder) AppendNonNil(value []byte) {
	if b.builder != nil {
		if value == nil {
			b.builder.AppendNull()
			return
		}

		switch builder := b.builder.(type) {
		case *array.BinaryBuilder:
			builder.Append(value)
		case *array.BinaryDictionaryBuilder:
			if err := builder.Append(value); err != nil {
				// Should never happen.
				panic(err)
			}
		default:
			// Should never happen.
			panic("unknown builder type")
		}
		return
	}

	if value != nil && b.updateRequest != nil {
		// If the builder is nil, then the transform node is not optional.
		b.transformNode.RemoveOptional()
		b.updateRequest.Inc(&update.NewFieldEvent{FieldName: b.transformNode.Path()})
		b.updateRequest = nil // No need to report this again.
	}
}

// AppendNull appends a null value to the underlying builder. If the builder is
// nil we do nothing as we have no information about the presence of this field
// in the data.
func (b *BinaryBuilder) AppendNull() {
	if b.builder != nil {
		b.builder.AppendNull()
		return
	}
}

// FixedSizeBinaryBuilder is a wrapper around the arrow FixedSizeBinaryBuilder.
type FixedSizeBinaryBuilder struct {
	builder       array.Builder
	transformNode *schema.TransformNode
	updateRequest *update.SchemaUpdateRequest
}

// Append appends a value to the underlying builder and updates the
// transform node if the builder is nil.
func (b *FixedSizeBinaryBuilder) Append(value []byte) {
	if b.builder != nil {
		if value == nil {
			b.builder.AppendNull()
			return
		}

		switch builder := b.builder.(type) {
		case *array.FixedSizeBinaryBuilder:
			builder.Append(value)
		case *array.FixedSizeBinaryDictionaryBuilder:
			if err := builder.Append(value); err != nil {
				// Should never happen.
				panic(err)
			}
		default:
			// Should never happen.
			panic("unknown builder type")
		}

		return
	}

	if value != nil && b.updateRequest != nil {
		// If the builder is nil, then the transform node is not optional.
		b.transformNode.RemoveOptional()
		b.updateRequest.Inc(&update.NewFieldEvent{FieldName: b.transformNode.Path()})
		b.updateRequest = nil // No need to report this again.
	}
}

// AppendNull appends a null value to the underlying builder. If the builder is
// nil we do nothing as we have no information about the presence of this field
// in the data.
func (b *FixedSizeBinaryBuilder) AppendNull() {
	if b.builder != nil {
		b.builder.AppendNull()
		return
	}
}
