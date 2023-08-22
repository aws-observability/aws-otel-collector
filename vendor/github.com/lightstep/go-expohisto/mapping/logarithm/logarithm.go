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

package logarithm // import "github.com/lightstep/go-expohisto/mapping/logarithm"

import (
	"fmt"
	"math"
	"sync"

	"github.com/lightstep/go-expohisto/mapping"
	"github.com/lightstep/go-expohisto/mapping/internal"
)

const (
	// MinScale ensures that the ../exponent mapper is used for
	// zero and negative scale values.  Do not use the logarithm
	// mapper for scales <= 0.
	MinScale int32 = 1

	// MaxScale is selected as the largest scale that is possible
	// in current code, considering there are 10 bits of base-2
	// exponent combined with scale-bits of range.  At this scale,
	// the growth factor is 0.0000661%.
	//
	// Scales larger than 20 complicate the logic in cmd/prebuild,
	// because math/big overflows when exponent is math.MaxInt32
	// (== the index of math.MaxFloat64 at scale=21),
	//
	// At scale=20, index values are in the interval [-0x3fe00000,
	// 0x3fffffff], having 31 bits of information.  This is
	// sensible given that the OTLP exponential histogram data
	// point uses a signed 32 bit integer for indices.
	MaxScale int32 = 20

	// MinValue is the smallest normal number.
	MinValue = internal.MinValue

	// MaxValue is the largest normal number.
	MaxValue = internal.MaxValue
)

// logarithmMapping contains the constants used to implement the
// exponential mapping function for a particular scale > 0.
type logarithmMapping struct {
	// scale is between MinScale and MaxScale. The exponential
	// base is defined as 2**(2**(-scale)).
	scale int32

	// scaleFactor is used and computed as follows:
	// index = log(value) / log(base)
	// = log(value) / log(2^(2^-scale))
	// = log(value) / (2^-scale * log(2))
	// = log(value) * (1/log(2) * 2^scale)
	// = log(value) * scaleFactor
	// where:
	// scaleFactor = (1/log(2) * 2^scale)
	// = math.Log2E * math.Exp2(scale)
	// = math.Ldexp(math.Log2E, scale)
	// Because multiplication is faster than division, we define scaleFactor as a multiplier.
	// This implementation was copied from a Java prototype. See:
	// https://github.com/newrelic-experimental/newrelic-sketch-java/blob/1ce245713603d61ba3a4510f6df930a5479cd3f6/src/main/java/com/newrelic/nrsketch/indexer/LogIndexer.java
	// for the equations used here.
	scaleFactor float64

	// log(boundary) = index * log(base)
	// log(boundary) = index * log(2^(2^-scale))
	// log(boundary) = index * 2^-scale * log(2)
	// boundary = exp(index * inverseFactor)
	// where:
	// inverseFactor = 2^-scale * log(2)
	// = math.Ldexp(math.Ln2, -scale)
	inverseFactor float64
}

var (
	_ mapping.Mapping = &logarithmMapping{}

	prebuiltMappingsLock sync.Mutex
	prebuiltMappings     = map[int32]*logarithmMapping{}
)

// NewMapping constructs a logarithm mapping function, used for scales > 0.
func NewMapping(scale int32) (mapping.Mapping, error) {
	// An assumption used in this code is that scale is > 0.  If
	// scale is <= 0 it's better to use the exponent mapping.
	if scale < MinScale || scale > MaxScale {
		// scale 20 can represent the entire float64 range
		// with a 30 bit index, and we don't handle larger
		// scales to simplify range tests in this package.
		return nil, fmt.Errorf("scale out of bounds")
	}
	prebuiltMappingsLock.Lock()
	defer prebuiltMappingsLock.Unlock()

	if p := prebuiltMappings[scale]; p != nil {
		return p, nil
	}
	l := &logarithmMapping{
		scale:         scale,
		scaleFactor:   math.Ldexp(math.Log2E, int(scale)),
		inverseFactor: math.Ldexp(math.Ln2, int(-scale)),
	}
	prebuiltMappings[scale] = l
	return l, nil
}

// minNormalLowerBoundaryIndex is the index such that base**index equals
// MinValue.  A histogram bucket with this index covers the range
// (MinValue, MinValue*base].  One less than this index corresponds
// with the bucket containing values <= MinValue.
func (l *logarithmMapping) minNormalLowerBoundaryIndex() int32 {
	return int32(internal.MinNormalExponent << l.scale)
}

// maxNormalLowerBoundaryIndex is the index such that base**index equals the
// greatest representable lower boundary.  A histogram bucket with this
// index covers the range (0x1p+1024/base, 0x1p+1024], which includes
// MaxValue; note that this bucket is incomplete, since the upper
// boundary cannot be represented.  One greater than this index
// corresponds with the bucket containing values > 0x1p1024.
func (l *logarithmMapping) maxNormalLowerBoundaryIndex() int32 {
	return (int32(internal.MaxNormalExponent+1) << l.scale) - 1
}

// MapToIndex implements mapping.Mapping.
func (l *logarithmMapping) MapToIndex(value float64) int32 {
	// Note: we can assume not a 0, Inf, or NaN; positive sign bit.
	if value <= MinValue {
		return l.minNormalLowerBoundaryIndex() - 1
	}

	// Exact power-of-two correctness: an optional special case.
	if internal.GetSignificand(value) == 0 {
		exp := internal.GetNormalBase2(value)
		return (exp << l.scale) - 1
	}

	// Non-power of two cases.  Use Floor(x) to round the scaled
	// logarithm.  We could use Ceil(x)-1 to achieve the same
	// result, though Ceil() is typically defined as -Floor(-x)
	// and typically not performed in hardware, so this is likely
	// less code.
	index := int32(math.Floor(math.Log(value) * l.scaleFactor))

	if max := l.maxNormalLowerBoundaryIndex(); index >= max {
		return max
	}
	return index
}

// LowerBoundary implements mapping.Mapping.
func (l *logarithmMapping) LowerBoundary(index int32) (float64, error) {
	if max := l.maxNormalLowerBoundaryIndex(); index >= max {
		if index == max {
			// Note that the equation on the last line of this
			// function returns +Inf.  Use the alternate equation.
			return 2 * math.Exp(float64(index-(int32(1)<<l.scale))*l.inverseFactor), nil
		}
		return 0, mapping.ErrOverflow
	}
	if min := l.minNormalLowerBoundaryIndex(); index <= min {
		if index == min {
			return MinValue, nil
		} else if index == min-1 {
			// Similar to the logic above, the math.Exp()
			// formulation is not accurate for subnormal
			// values.
			return math.Exp(float64(index+(int32(1)<<l.scale))*l.inverseFactor) / 2, nil
		}
		return 0, mapping.ErrUnderflow
	}
	return math.Exp(float64(index) * l.inverseFactor), nil
}

// Scale implements mapping.Mapping.
func (l *logarithmMapping) Scale() int32 {
	return l.scale
}
