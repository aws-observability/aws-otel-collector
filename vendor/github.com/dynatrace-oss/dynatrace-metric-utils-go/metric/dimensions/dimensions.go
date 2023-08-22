// Copyright 2021 Dynatrace LLC
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dimensions

import (
	"log"

	"github.com/dynatrace-oss/dynatrace-metric-utils-go/normalize"
)

// Dimension is a key-value pair mapping string to string.
type Dimension struct {
	Key   string
	Value string
}

// NormalizedDimensionList is a holder structure, containing normalized but not necessarily unique Dimension key-value pairs.
type NormalizedDimensionList struct {
	dimensions []Dimension
}

// NewNormalizedDimensionList creates a new Dimension set. All dimensions in the set are normalized, but it mights still contain duplicate keys.
// Dimensions with invalid keys (after normalization) are dropped.
func NewNormalizedDimensionList(dims ...Dimension) NormalizedDimensionList {
	return NormalizedDimensionList{dimensions: normalizeDimensions(dims...)}
}

// pass a function that transforms a slice of dimensions to a string.
// That way, the code for the actual serialization can be stored in the
// serialization package without exporting the dimensions in normalized dimensions
// which in turn restricts manipulation of already normalized values.
func (ds NormalizedDimensionList) Format(formatter func([]Dimension) string) string {
	return formatter(ds.dimensions)
}

func NewDimension(key, val string) Dimension {
	return Dimension{Key: key, Value: val}
}

// normalizeDimensions normalizes all dimensions passed to it.
// No duplicate elimination will be done at this stage, but invalid dimensions (e.g. empty key after normalization)
// are removed. The order of the passed dimensions is retained.
func normalizeDimensions(dims ...Dimension) []Dimension {
	// this is basically a set, but golang does not offer a set type
	normalizedDims := []Dimension{}

	for _, dim := range dims {
		k, err := normalize.DimensionKey(dim.Key)
		if err != nil {
			log.Printf("normalization for '%s' returned invalid key. Skipping...", dim.Key)
			continue
		}

		normalizedDims = append(normalizedDims, NewDimension(k, normalize.DimensionValue(dim.Value)))
	}

	return normalizedDims

}

// MergeLists combines one or more NormalizedDimensionList instances into one. Dimensions in sets passed further to the right but containing the
// same keys as sets further to the left will overwrite the values. The resulting set contains no duplicate keys. If duplicate
// keys appear in different sets, the value of the resulting set will be the one from the last set passed to this function and
// containing the key. The order of keys in the resulting set is not guaranteed.
func MergeLists(normalizedDimensionLists ...NormalizedDimensionList) NormalizedDimensionList {
	if len(normalizedDimensionLists) == 0 {
		return NewNormalizedDimensionList()
	}

	uniqueDimensions := make(map[string]Dimension)
	for _, set := range normalizedDimensionLists {
		for _, dim := range set.dimensions {
			// since dimension sets are already normalized there is no need to do it again here.
			uniqueDimensions[dim.Key] = dim
		}
	}

	uniqueDimSlice := make([]Dimension, 0, len(uniqueDimensions))

	for _, v := range uniqueDimensions {
		uniqueDimSlice = append(uniqueDimSlice, v)
	}

	return NormalizedDimensionList{dimensions: uniqueDimSlice}
}
