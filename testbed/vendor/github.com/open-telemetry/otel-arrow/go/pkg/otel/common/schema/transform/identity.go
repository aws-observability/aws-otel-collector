/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package transform

import "github.com/apache/arrow-go/v18/arrow"

// IdentityField is a FieldTransform that returns a copy of the field.
type IdentityField struct {
	path string
}

func NewIdentityField(path string) *IdentityField {
	return &IdentityField{path: path}
}

func (t *IdentityField) Transform(field *arrow.Field) *arrow.Field {
	return &arrow.Field{Name: field.Name, Type: field.Type, Nullable: field.Nullable, Metadata: field.Metadata}
}

func (t *IdentityField) RevertCounters() {}

func (t *IdentityField) Path() string {
	return t.path
}
