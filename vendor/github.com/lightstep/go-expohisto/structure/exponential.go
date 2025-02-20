// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package structure // import "github.com/lightstep/go-expohisto/structure"

import (
	"fmt"

	"github.com/lightstep/go-expohisto/mapping"
	"github.com/lightstep/go-expohisto/mapping/exponent"
	"github.com/lightstep/go-expohisto/mapping/logarithm"
)

type (
	// Histogram observes counts observations in exponentially-spaced
	// buckets.  It is configured with a maximum scale factor
	// which determines resolution.  Scale is automatically
	// adjusted to accommodate the range of input data.
	//
	// Note that the generic type `N` determines the type of the
	// Sum, Min, and Max fields.  Bucket boundaries are handled in
	// floating point regardless of the type of N.
	Histogram[N ValueType] struct {
		// maxSize is the maximum capacity of the positive and
		// negative ranges.  it is set by Init(), preserved by
		// Copy and Move.
		maxSize int32

		// sum is the sum of all Updates reflected in the
		// aggregator.  It has the same type number as the
		// corresponding sdkinstrument.Descriptor.
		sum N
		// count is incremented by 1 per Update.
		count uint64
		// zeroCount is incremented by 1 when the measured
		// value is exactly 0.
		zeroCount uint64
		// min is set when count > 0
		min N
		// max is set when count > 0
		max N
		// positive holds the positive values
		positive Buckets
		// negative holds the negative values in these buckets
		// by their absolute value.
		negative Buckets
		// mapping corresponds to the current scale, is shared
		// by both positive and negative ranges.
		mapping mapping.Mapping
	}

	// Buckets stores counts for measurement values in the range
	// (0, +Inf).
	Buckets struct {
		// backing is a slice of nil, []uint8, []uint16, []uint32, or []uint64
		backing bucketsBacking

		// The term "index" refers to the number of the
		// histogram bucket used to determine its boundaries.
		// The lower-boundary of a bucket is determined by
		// formula base**index and the upper-boundary of a
		// bucket is base**(index+1).  Index values are signed
		// to account for values less than or equal to 1.
		//
		// Note that the width of this field is determined by
		// the field being stated as int32 in the OTLP
		// protocol.  The meaning of this field can be
		// extended to wider types, however this it would
		// would be an extremely high-resolution histogram.

		// indexBase is index of the 0th position in the
		// backing array, i.e., backing[0] is the count
		// in the bucket with index `indexBase`.
		indexBase int32

		// indexStart is the smallest index value represented
		// in the backing array.
		indexStart int32

		// indexEnd is the largest index value represented in
		// the backing array.
		indexEnd int32
	}

	// ValueType is an interface constraint for the numeric type
	// aggregated by this histogram.
	ValueType interface {
		int64 | float64
	}

	// bucketsCount are the possible backing array widths.
	bucketsCount interface {
		uint8 | uint16 | uint32 | uint64
	}

	// bucketsVarwidth is a variable-width slice of unsigned int counters.
	bucketsVarwidth[N bucketsCount] struct {
		counts []N
	}

	// bucketsBacking is implemented by bucektsVarwidth[N].
	bucketsBacking interface {
		// size returns the physical size of the backing
		// array, which is >= buckets.Len() the number allocated.
		//
		// Note this is logically an unsigned quantity,
		// however it creates fewer type conversions in the
		// code with this as int32, because: (a) this is not
		// allowed to grow to outside the range of a signed
		// int32, and (b) this is frequently involved in
		// arithmetic with signed index values.
		size() int32
		// growTo grows a backing array and copies old entries
		// into their correct new positions.
		growTo(newSize, oldPositiveLimit, newPositiveLimit int32)
		// reverse reverse the items in a backing array in the
		// range [from, limit).
		reverse(from, limit int32)
		// emptyBucket empties the count from a bucket, for
		// moving into another.
		emptyBucket(src int32) uint64
		// tryIncrement increments a bucket by `incr`, returns
		// false if the result would overflow the current
		// backing width.
		tryIncrement(bucketIndex int32, incr uint64) bool
		// countAt returns the count in a specific bucket.
		countAt(pos uint32) uint64
		// reset resets all buckets to zero count.
		reset()
	}

	// highLow is used to establish the maximum range of bucket
	// indices needed, in order to establish the best value of the
	// scale parameter.
	highLow struct {
		low  int32
		high int32
	}

	// Int64 is an integer-valued histogram.
	Int64 = Histogram[int64]

	// Float64 is a float64-valued histogram.
	Float64 = Histogram[float64]
)

// Init initializes a new histogram.
func (h *Histogram[N]) Init(cfg Config) {
	cfg, _ = cfg.Validate()

	h.maxSize = cfg.maxSize

	m, _ := newMapping(logarithm.MaxScale)
	h.mapping = m
}

// Sum implements aggregation.Histogram.
func (h *Histogram[N]) Sum() N {
	return h.sum
}

// Min implements aggregation.Histogram.
func (h *Histogram[N]) Min() N {
	return h.min
}

// Max implements aggregation.Histogram.
func (h *Histogram[N]) Max() N {
	return h.max
}

// Count implements aggregation.Histogram.
func (h *Histogram[N]) Count() uint64 {
	return h.count
}

// Scale implements aggregation.Histogram.
func (h *Histogram[N]) Scale() int32 {
	if h.count == h.zeroCount {
		// all zeros! scale doesn't matter, use zero.
		return 0
	}
	return h.mapping.Scale()
}

// ZeroCount implements aggregation.Histogram.
func (h *Histogram[N]) ZeroCount() uint64 {
	return h.zeroCount
}

// Positive implements aggregation.Histogram.
func (h *Histogram[N]) Positive() *Buckets {
	return &h.positive
}

// Negative implements aggregation.Histogram.
func (h *Histogram[N]) Negative() *Buckets {
	return &h.negative
}

// Offset implements aggregation.Bucket.
func (b *Buckets) Offset() int32 {
	return b.indexStart
}

// Len implements aggregation.Bucket.
func (b *Buckets) Len() uint32 {
	if b.backing == nil {
		return 0
	}
	if b.indexEnd == b.indexStart && b.At(0) == 0 {
		return 0
	}
	return uint32(b.indexEnd - b.indexStart + 1)
}

// At returns the count of the bucket at a position in the logical
// array of counts.
func (b *Buckets) At(pos0 uint32) uint64 {
	pos := pos0
	bias := uint32(b.indexBase - b.indexStart)

	if pos < bias {
		pos += uint32(b.backing.size())
	}
	pos -= bias

	return b.backing.countAt(pos)
}

// Clear resets a histogram to the empty state without changing
// backing array.
func (h *Histogram[N]) Clear() {
	h.positive.clear()
	h.negative.clear()
	h.sum = 0
	h.count = 0
	h.zeroCount = 0
	h.min = 0
	h.max = 0
	h.mapping, _ = newMapping(logarithm.MaxScale)
}

// clear zeros the backing array.
func (b *Buckets) clear() {
	b.indexStart = 0
	b.indexEnd = 0
	b.indexBase = 0
	if b.backing != nil {
		b.backing.reset()
	}
}

func newMapping(scale int32) (mapping.Mapping, error) {
	if scale <= 0 {
		return exponent.NewMapping(scale)
	}
	return logarithm.NewMapping(scale)
}

// Swap exchanges the contents of `h` and `dest`.
func (h *Histogram[N]) Swap(dest *Histogram[N]) {
	*dest, *h = *h, *dest
}

// CopyInto copies `h` into `dest`.
func (h *Histogram[N]) CopyInto(dest *Histogram[N]) {
	dest.Clear()
	dest.MergeFrom(h)
}

// Update supports updating a histogram with a single count.
func (h *Histogram[N]) Update(number N) {
	h.UpdateByIncr(number, 1)
}

// UpdateByIncr supports updating a histogram with a non-negative
// increment.
func (h *Histogram[N]) UpdateByIncr(number N, incr uint64) {
	value := float64(number)

	// Maintain min and max
	if h.count == 0 {
		h.min = number
		h.max = number
	} else {
		if number < h.min {
			h.min = number
		}
		if number > h.max {
			h.max = number
		}
	}

	// Note: Not checking for overflow here. TODO.
	h.count += incr

	if value == 0 {
		h.zeroCount += incr
		return
	}

	// Sum maintains the original type, otherwise we use the floating point value.
	h.sum += number * N(incr)

	var b *Buckets
	if value > 0 {
		b = &h.positive
	} else {
		value = -value
		b = &h.negative
	}

	h.update(b, value, incr)
}

// downscale subtracts `change` from the current mapping scale.
func (h *Histogram[N]) downscale(change int32) {
	if change == 0 {
		return
	}
	if change < 0 {
		panic(fmt.Sprint("impossible change of scale", change))
	}
	newScale := h.mapping.Scale() - change

	h.positive.downscale(change)
	h.negative.downscale(change)
	var err error
	h.mapping, err = newMapping(newScale)
	if err != nil {
		panic(fmt.Sprint("impossible scale", newScale))
	}
}

// changeScale computes how much downscaling is needed by shifting the
// high and low values until they are separated by no more than size.
func changeScale(hl highLow, size int32) int32 {
	var change int32
	for hl.high-hl.low >= size {
		hl.high >>= 1
		hl.low >>= 1
		change++
	}
	return change
}

// update increments the appropriate buckets for a given absolute
// value by the provided increment.
func (h *Histogram[N]) update(b *Buckets, value float64, incr uint64) {
	index := h.mapping.MapToIndex(value)

	hl, success := h.incrementIndexBy(b, index, incr)
	if success {
		return
	}

	h.downscale(changeScale(hl, h.maxSize))

	index = h.mapping.MapToIndex(value)
	if _, success := h.incrementIndexBy(b, index, incr); !success {
		panic("downscale logic error")
	}
}

// incrementIndexBy determines if the index lies inside the current range
// [indexStart, indexEnd] and, if not, returns the minimum size (up to
// maxSize) will satisfy the new value.
func (h *Histogram[N]) incrementIndexBy(b *Buckets, index int32, incr uint64) (highLow, bool) {
	if incr == 0 {
		// Skipping a bunch of work for 0 increment.  This
		// happens when merging sparse data, for example.
		// This also happens UpdateByIncr is used with a 0
		// increment, means it can be safely skipped.
		return highLow{}, true
	}
	if b.Len() == 0 {
		if b.backing == nil {
			b.backing = &bucketsVarwidth[uint8]{
				counts: []uint8{0},
			}
		}
		b.indexStart = index
		b.indexEnd = b.indexStart
		b.indexBase = b.indexStart
	} else if index < b.indexStart {
		if span := b.indexEnd - index; span >= h.maxSize {
			// rescale needed: mapped value to the right
			return highLow{
				low:  index,
				high: b.indexEnd,
			}, false
		} else if span >= b.backing.size() {
			h.grow(b, span+1)
		}
		b.indexStart = index
	} else if index > b.indexEnd {
		if span := index - b.indexStart; span >= h.maxSize {
			// rescale needed: mapped value to the left
			return highLow{
				low:  b.indexStart,
				high: index,
			}, false
		} else if span >= b.backing.size() {
			h.grow(b, span+1)
		}
		b.indexEnd = index
	}

	bucketIndex := index - b.indexBase
	if bucketIndex < 0 {
		bucketIndex += b.backing.size()
	}
	b.incrementBucket(bucketIndex, incr)
	return highLow{}, true
}

// powTwoRoundedUp computes the next largest power-of-two, which
// ensures power-of-two slices are allocated.
func powTwoRoundedUp(v int32) int32 {
	// The following expression computes the least power-of-two
	// that is >= v.  There are a number of tricky ways to
	// do this, see https://stackoverflow.com/questions/466204/rounding-up-to-next-power-of-2
	//
	// One equivalent expression:
	//
	// v = int32(1) << (32 - bits.LeadingZeros32(uint32(v-1)))
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

// grow resizes the backing array by doubling in size up to maxSize.
// this extends the array with a bunch of zeros and copies the
// existing counts to the same position.
func (h *Histogram[N]) grow(b *Buckets, needed int32) {
	size := b.backing.size()
	bias := b.indexBase - b.indexStart
	oldPositiveLimit := size - bias
	newSize := powTwoRoundedUp(needed)
	if newSize > h.maxSize {
		newSize = h.maxSize
	}
	newPositiveLimit := newSize - bias
	b.backing.growTo(newSize, oldPositiveLimit, newPositiveLimit)
}

// downscale first rotates, then collapses 2**`by`-to-1 buckets.
func (b *Buckets) downscale(by int32) {
	b.rotate()

	size := 1 + b.indexEnd - b.indexStart
	each := int64(1) << by
	inpos := int32(0)
	outpos := int32(0)

	for pos := b.indexStart; pos <= b.indexEnd; {
		mod := int64(pos) % each
		if mod < 0 {
			mod += each
		}
		for i := mod; i < each && inpos < size; i++ {
			b.relocateBucket(outpos, inpos)
			inpos++
			pos++
		}
		outpos++
	}

	b.indexStart >>= by
	b.indexEnd >>= by
	b.indexBase = b.indexStart
}

// rotate shifts the backing array contents so that indexStart ==
// indexBase to simplify the downscale logic.
func (b *Buckets) rotate() {
	bias := b.indexBase - b.indexStart

	if bias == 0 {
		return
	}

	// Rotate the array so that indexBase == indexStart
	b.indexBase = b.indexStart

	b.backing.reverse(0, b.backing.size())
	b.backing.reverse(0, bias)
	b.backing.reverse(bias, b.backing.size())
}

// relocateBucket adds the count in counts[src] to counts[dest] and
// resets count[src] to zero.
func (b *Buckets) relocateBucket(dest, src int32) {
	if dest == src {
		return
	}

	b.incrementBucket(dest, b.backing.emptyBucket(src))
}

// incrementBucket increments the backing array index by `incr`.
func (b *Buckets) incrementBucket(bucketIndex int32, incr uint64) {
	for {
		if b.backing.tryIncrement(bucketIndex, incr) {
			return
		}

		switch bt := b.backing.(type) {
		case *bucketsVarwidth[uint8]:
			b.backing = widenBuckets[uint8, uint16](bt)
		case *bucketsVarwidth[uint16]:
			b.backing = widenBuckets[uint16, uint32](bt)
		case *bucketsVarwidth[uint32]:
			b.backing = widenBuckets[uint32, uint64](bt)
		case *bucketsVarwidth[uint64]:
			// Problem. The exponential histogram has overflowed a uint64.
			// However, this shouldn't happen because the total count would
			// overflow first.
			panic("bucket overflow must be avoided")
		}
	}
}

// Merge combines data from `o` into `h`.
func (h *Histogram[N]) MergeFrom(o *Histogram[N]) {
	if h.count == 0 {
		h.min = o.min
		h.max = o.max
	} else if o.count != 0 {
		if o.min < h.min {
			h.min = o.min
		}
		if o.max > h.max {
			h.max = o.max
		}
	}

	// Note: Not checking for overflow here. TODO.
	h.sum += o.sum
	h.count += o.count
	h.zeroCount += o.zeroCount

	minScale := int32min(h.Scale(), o.Scale())

	hlp := h.highLowAtScale(&h.positive, minScale)
	hlp = hlp.with(o.highLowAtScale(&o.positive, minScale))

	hln := h.highLowAtScale(&h.negative, minScale)
	hln = hln.with(o.highLowAtScale(&o.negative, minScale))

	minScale = int32min(
		minScale-changeScale(hlp, h.maxSize),
		minScale-changeScale(hln, h.maxSize),
	)

	h.downscale(h.Scale() - minScale)

	h.mergeBuckets(&h.positive, o, &o.positive, minScale)
	h.mergeBuckets(&h.negative, o, &o.negative, minScale)
}

// mergeBuckets translates index values from another histogram into
// the corresponding buckets of this histogram.
func (h *Histogram[N]) mergeBuckets(mine *Buckets, other *Histogram[N], theirs *Buckets, scale int32) {
	theirOffset := theirs.Offset()
	theirChange := other.Scale() - scale

	for i := uint32(0); i < theirs.Len(); i++ {
		_, success := h.incrementIndexBy(
			mine,
			(theirOffset+int32(i))>>theirChange,
			theirs.At(i),
		)
		if !success {
			panic("incorrect merge scale")
		}
	}
}

// highLowAtScale is an accessory for Merge() to calculate ideal combined scale.
func (h *Histogram[N]) highLowAtScale(b *Buckets, scale int32) highLow {
	if b.Len() == 0 {
		return highLow{
			low:  0,
			high: -1,
		}
	}
	shift := h.Scale() - scale
	return highLow{
		low:  b.indexStart >> shift,
		high: b.indexEnd >> shift,
	}
}

// with is an accessory for Merge() to calculate ideal combined scale.
func (h *highLow) with(o highLow) highLow {
	if o.empty() {
		return *h
	}
	if h.empty() {
		return o
	}
	return highLow{
		low:  int32min(h.low, o.low),
		high: int32max(h.high, o.high),
	}
}

// empty indicates whether there are any values in a highLow.
func (h *highLow) empty() bool {
	return h.low > h.high
}

func int32min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func int32max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// bucketsVarwidth[]
//
// Each of the methods below is generic with respect to the underlying
// backing array.  See the interface-level comments.

func (b *bucketsVarwidth[N]) countAt(pos uint32) uint64 {
	return uint64(b.counts[pos])
}

func (b *bucketsVarwidth[N]) reset() {
	for i := range b.counts {
		b.counts[i] = 0
	}
}

func (b *bucketsVarwidth[N]) size() int32 {
	return int32(len(b.counts))
}

func (b *bucketsVarwidth[N]) growTo(newSize, oldPositiveLimit, newPositiveLimit int32) {
	tmp := make([]N, newSize)
	copy(tmp[newPositiveLimit:], b.counts[oldPositiveLimit:])
	copy(tmp[0:oldPositiveLimit], b.counts[0:oldPositiveLimit])
	b.counts = tmp
}

func (b *bucketsVarwidth[N]) reverse(from, limit int32) {
	num := ((from + limit) / 2) - from
	for i := int32(0); i < num; i++ {
		b.counts[from+i], b.counts[limit-i-1] = b.counts[limit-i-1], b.counts[from+i]
	}
}

func (b *bucketsVarwidth[N]) emptyBucket(src int32) uint64 {
	tmp := b.counts[src]
	b.counts[src] = 0
	return uint64(tmp)
}

func (b *bucketsVarwidth[N]) tryIncrement(bucketIndex int32, incr uint64) bool {
	var limit = uint64(N(0) - 1)
	if uint64(b.counts[bucketIndex])+incr <= limit {
		b.counts[bucketIndex] += N(incr)
		return true
	}
	return false
}

func widenBuckets[From, To bucketsCount](in *bucketsVarwidth[From]) *bucketsVarwidth[To] {
	tmp := make([]To, len(in.counts))
	for i := range in.counts {
		tmp[i] = To(in.counts[i])
	}
	return &bucketsVarwidth[To]{counts: tmp}
}
