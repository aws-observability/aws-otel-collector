/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package arrow

import (
	"github.com/apache/arrow-go/v18/arrow"
)

func StructFromSparseUnion(dt *arrow.SparseUnionType, code int8) *arrow.StructType {
	codes := dt.TypeCodes()
	for i, c := range codes {
		if c == code {
			return dt.Fields()[i].Type.(*arrow.StructType)
		}
	}
	return nil
}
