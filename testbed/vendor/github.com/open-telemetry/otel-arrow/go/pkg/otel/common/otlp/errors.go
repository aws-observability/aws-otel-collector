/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package otlp

import "errors"

var (
	ErrMissingRelatedData  = errors.New("missing related data")
	ErrInvalidTypeCode     = errors.New("invalid type code")
	ErrInvalidFieldId      = errors.New("invalid field id")
	ErrParentIDMissing     = errors.New("parent id missing")
	ErrInvalidAttrName     = errors.New("invalid attribute name")
	ErrMissingTypeMetadata = errors.New("missing type metadata")
)
