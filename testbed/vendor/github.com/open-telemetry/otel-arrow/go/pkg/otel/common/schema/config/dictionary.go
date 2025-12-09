/*
 * Copyright The OpenTelemetry Authors
 * SPDX-License-Identifier: Apache-2.0
 */

package builder

import "math"

// Dictionary is a configuration for a dictionary field.
// The MaxCard is the maximum cardinality of the dictionary field. If the
// cardinality of the dictionary field is higher than MaxCard, then the
// dictionary field will be automatically converted to its base type.
//
// if MaxCard is equal to 0, then the dictionary field will be converted to its
// base type no matter what.
type Dictionary struct {
	MinCard        uint64
	MaxCard        uint64
	ResetThreshold float64
}

// NewDictionary creates a new dictionary configuration with the given maximum
// cardinality.
func NewDictionary(maxCard uint64, resetThreshold float64) *Dictionary {
	// If `maxCard` is 0 (no dictionary configuration), then the dictionary
	// field will be converted to its base type no matter what. So, the minimum
	// cardinality will be set to 0.
	minCard := uint64(math.MaxUint8)
	if maxCard < minCard {
		minCard = maxCard
	}
	return &Dictionary{
		MinCard:        minCard,
		MaxCard:        maxCard,
		ResetThreshold: resetThreshold,
	}
}

// NewDictionaryFrom creates a new dictionary configuration from a prototype
// dictionary configuration with the given minimum cardinality.
func NewDictionaryFrom(minCard uint64, dicProto *Dictionary) *Dictionary {
	// If `maxCard` is 0 (no dictionary configuration), then the dictionary
	// field will be converted to its base type no matter what. So, the minimum
	// cardinality will be set to 0.
	if dicProto.MaxCard < minCard {
		minCard = dicProto.MaxCard
	}
	return &Dictionary{
		MinCard:        minCard,
		MaxCard:        dicProto.MaxCard,
		ResetThreshold: dicProto.ResetThreshold,
	}
}
