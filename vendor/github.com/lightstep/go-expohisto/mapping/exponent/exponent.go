// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exponent // import "github.com/lightstep/go-expohisto/mapping/exponent"

import (
	"fmt"
	"math"

	"github.com/lightstep/go-expohisto/mapping"
	"github.com/lightstep/go-expohisto/mapping/internal"
)

const (
	// MinScale defines the point at which the exponential mapping
	// function becomes useless for float64.  With scale -10, ignoring
	// subnormal values, bucket indices range from -1 to 1.
	MinScale int32 = -10

	// MaxScale is the largest scale supported in this code.  Use
	// ../logarithm for larger scales.
	MaxScale int32 = 0
)

type exponentMapping struct {
	shift uint8 // equals negative scale
}

// exponentMapping is used for negative scales, effectively a
// mapping of the base-2 logarithm of the exponent.
var prebuiltMappings = [-MinScale + 1]exponentMapping{
	{10},
	{9},
	{8},
	{7},
	{6},
	{5},
	{4},
	{3},
	{2},
	{1},
	{0},
}

// NewMapping constructs an exponential mapping function, used for scales <= 0.
func NewMapping(scale int32) (mapping.Mapping, error) {
	if scale > MaxScale {
		return nil, fmt.Errorf("exponent mapping requires scale <= 0")
	}
	if scale < MinScale {
		return nil, fmt.Errorf("scale too low")
	}
	return &prebuiltMappings[scale-MinScale], nil
}

// minNormalLowerBoundaryIndex is the largest index such that
// base**index is <= MinValue.  A histogram bucket with this index
// covers the range (base**index, base**(index+1)], including
// MinValue.
func (e *exponentMapping) minNormalLowerBoundaryIndex() int32 {
	idx := int32(internal.MinNormalExponent) >> e.shift
	if e.shift < 2 {
		// For scales -1 and 0 the minimum value 2**-1022
		// is a power-of-two multiple, meaning it belongs
		// to the index one less.
		idx--
	}
	return idx
}

// maxNormalLowerBoundaryIndex is the index such that base**index
// equals the largest representable boundary.  A histogram bucket with this
// index covers the range (0x1p+1024/base, 0x1p+1024], which includes
// MaxValue; note that this bucket is incomplete, since the upper
// boundary cannot be represented.  One greater than this index
// corresponds with the bucket containing values > 0x1p1024.
func (e *exponentMapping) maxNormalLowerBoundaryIndex() int32 {
	return int32(internal.MaxNormalExponent) >> e.shift
}

// MapToIndex implements mapping.Mapping.
func (e *exponentMapping) MapToIndex(value float64) int32 {
	// Note: we can assume not a 0, Inf, or NaN; positive sign bit.
	if value < internal.MinValue {
		return e.minNormalLowerBoundaryIndex()
	}

	// Extract the raw exponent.
	rawExp := internal.GetNormalBase2(value)

	// In case the value is an exact power of two, compute a
	// correction of -1:
	correction := int32((internal.GetSignificand(value) - 1) >> internal.SignificandWidth)

	// Note: bit-shifting does the right thing for negative
	// exponents, e.g., -1 >> 1 == -1.
	return (rawExp + correction) >> e.shift
}

// LowerBoundary implements mapping.Mapping.
func (e *exponentMapping) LowerBoundary(index int32) (float64, error) {
	if min := e.minNormalLowerBoundaryIndex(); index < min {
		return 0, mapping.ErrUnderflow
	}

	if max := e.maxNormalLowerBoundaryIndex(); index > max {
		return 0, mapping.ErrOverflow
	}

	return math.Ldexp(1, int(index<<e.shift)), nil
}

// Scale implements mapping.Mapping.
func (e *exponentMapping) Scale() int32 {
	return -int32(e.shift)
}
