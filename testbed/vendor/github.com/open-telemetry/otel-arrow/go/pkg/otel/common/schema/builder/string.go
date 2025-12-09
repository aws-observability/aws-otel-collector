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

type StringBuilder struct {
	builder       array.Builder
	transformNode *schema.TransformNode
	updateRequest *update.SchemaUpdateRequest
}

// NewStringBuilder creates a new StringBuilder.
func NewStringBuilder(builder array.Builder, transformNode *schema.TransformNode, updateRequest *update.SchemaUpdateRequest) *StringBuilder {
	return &StringBuilder{
		builder:       builder,
		transformNode: transformNode,
		updateRequest: updateRequest,
	}
}

func (b *StringBuilder) AppendNull() {
	if b.builder != nil {
		b.builder.AppendNull()
		return
	}
}

func (b *StringBuilder) Append(value string) {
	if b.builder != nil {
		switch builder := b.builder.(type) {
		case *array.StringBuilder:
			builder.Append(value)
		case *array.BinaryDictionaryBuilder:
			if err := builder.AppendString(value); err != nil {
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
		b.transformNode.RemoveOptional()
		b.updateRequest.Inc(&update.NewFieldEvent{FieldName: b.transformNode.Path()})
		b.updateRequest = nil // No need to report this again.
	}
}

func (b *StringBuilder) AppendNonEmpty(value string) {
	if b.builder != nil {
		if value == "" {
			b.builder.AppendNull()
			return
		}

		switch builder := b.builder.(type) {
		case *array.StringBuilder:
			builder.Append(value)
		case *array.BinaryDictionaryBuilder:
			if err := builder.AppendString(value); err != nil {
				// Should never happen.
				panic(err)
			}
		default:
			// Should never happen.
			panic("unknown builder type")
		}

		return
	}

	if value != "" && b.updateRequest != nil {
		b.transformNode.RemoveOptional()
		b.updateRequest.Inc(&update.NewFieldEvent{FieldName: b.transformNode.Path()})
		b.updateRequest = nil // No need to report this again.
	}
}
