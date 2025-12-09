/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package events

// Events is a collection of events that occurred while using RecordBuilderExt
type Events struct {
	// Dictionary fields that have overflowed and that have been downgraded to
	// their value type.
	DictionariesWithOverflow map[string]bool

	// Dictionary fields that have their dictionary index type changed.
	DictionariesIndexTypeChanged map[string]string
}
