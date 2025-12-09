/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package common

import (
	"sort"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type (
	AttributeEntry struct {
		Key   string
		Value pcommon.Value
	}

	AttributeEntries []AttributeEntry
)

var _ sort.Interface = AttributeEntries{}

func (ae AttributeEntries) Len() int {
	return len(ae)
}

func (ae AttributeEntries) Swap(i, j int) {
	ae[i], ae[j] = ae[j], ae[i]
}

func (ae AttributeEntries) Less(i, j int) bool {
	return ae[i].Key < ae[j].Key
}
