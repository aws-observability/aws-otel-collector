// Copyright 2021 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"math"

	"github.com/prometheus/prometheus/model/histogram"
	"github.com/prometheus/prometheus/tsdb/chunkenc"
)

// MemoizedSeriesIterator wraps an iterator with a buffer to look back the previous element.
<<<<<<< HEAD
=======
//
// This iterator regards integer histograms as float histograms; calls to Seek() will never return chunkenc.Histogram.
// This iterator deliberately does not implement chunkenc.Iterator.
>>>>>>> main
type MemoizedSeriesIterator struct {
	it    chunkenc.Iterator
	delta int64

	lastTime  int64
	valueType chunkenc.ValueType

	// Keep track of the previously returned value.
	prevTime           int64
	prevValue          float64
<<<<<<< HEAD
	prevHistogram      *histogram.Histogram
	prevFloatHistogram *histogram.FloatHistogram
	// TODO(beorn7): MemoizedSeriesIterator is currently only used by the
	// PromQL engine, which only works with FloatHistograms. For better
	// performance, we could change MemoizedSeriesIterator to also only
	// handle FloatHistograms.
=======
	prevFloatHistogram *histogram.FloatHistogram
>>>>>>> main
}

// NewMemoizedEmptyIterator is like NewMemoizedIterator but it's initialised with an empty iterator.
func NewMemoizedEmptyIterator(delta int64) *MemoizedSeriesIterator {
	return NewMemoizedIterator(chunkenc.NewNopIterator(), delta)
}

// NewMemoizedIterator returns a new iterator that buffers the values within the
// time range of the current element and the duration of delta before.
func NewMemoizedIterator(it chunkenc.Iterator, delta int64) *MemoizedSeriesIterator {
	bit := &MemoizedSeriesIterator{
		delta:    delta,
		prevTime: math.MinInt64,
	}
	bit.Reset(it)

	return bit
}

// Reset the internal state to reuse the wrapper with the provided iterator.
func (b *MemoizedSeriesIterator) Reset(it chunkenc.Iterator) {
	b.it = it
	b.lastTime = math.MinInt64
	b.prevTime = math.MinInt64
	b.valueType = it.Next()
}

// PeekPrev returns the previous element of the iterator. If there is none buffered,
// ok is false.
<<<<<<< HEAD
func (b *MemoizedSeriesIterator) PeekPrev() (t int64, v float64, h *histogram.Histogram, fh *histogram.FloatHistogram, ok bool) {
	if b.prevTime == math.MinInt64 {
		return 0, 0, nil, nil, false
	}
	return b.prevTime, b.prevValue, b.prevHistogram, b.prevFloatHistogram, true
=======
func (b *MemoizedSeriesIterator) PeekPrev() (t int64, v float64, fh *histogram.FloatHistogram, ok bool) {
	if b.prevTime == math.MinInt64 {
		return 0, 0, nil, false
	}
	return b.prevTime, b.prevValue, b.prevFloatHistogram, true
>>>>>>> main
}

// Seek advances the iterator to the element at time t or greater.
func (b *MemoizedSeriesIterator) Seek(t int64) chunkenc.ValueType {
	t0 := t - b.delta

	if b.valueType != chunkenc.ValNone && t0 > b.lastTime {
		// Reset the previously stored element because the seek advanced
		// more than the delta.
		b.prevTime = math.MinInt64

		b.valueType = b.it.Seek(t0)
<<<<<<< HEAD
		if b.valueType == chunkenc.ValNone {
			return chunkenc.ValNone
=======
		switch b.valueType {
		case chunkenc.ValNone:
			return chunkenc.ValNone
		case chunkenc.ValHistogram:
			b.valueType = chunkenc.ValFloatHistogram
>>>>>>> main
		}
		b.lastTime = b.it.AtT()
	}
	if b.lastTime >= t {
		return b.valueType
	}
	for b.Next() != chunkenc.ValNone {
		if b.lastTime >= t {
			return b.valueType
		}
	}

	return chunkenc.ValNone
}

<<<<<<< HEAD
// Next advances the iterator to the next element.
=======
// Next advances the iterator to the next element. Note that this does not check whether the element being buffered is
// within the time range of the current element and the duration of delta before.
>>>>>>> main
func (b *MemoizedSeriesIterator) Next() chunkenc.ValueType {
	// Keep track of the previous element.
	switch b.valueType {
	case chunkenc.ValNone:
		return chunkenc.ValNone
	case chunkenc.ValFloat:
		b.prevTime, b.prevValue = b.it.At()
<<<<<<< HEAD
		b.prevHistogram = nil
		b.prevFloatHistogram = nil
	case chunkenc.ValHistogram:
		b.prevValue = 0
		b.prevTime, b.prevHistogram = b.it.AtHistogram()
		_, b.prevFloatHistogram = b.it.AtFloatHistogram()
	case chunkenc.ValFloatHistogram:
		b.prevValue = 0
		b.prevHistogram = nil
=======
		b.prevFloatHistogram = nil
	case chunkenc.ValHistogram, chunkenc.ValFloatHistogram:
		b.prevValue = 0
>>>>>>> main
		b.prevTime, b.prevFloatHistogram = b.it.AtFloatHistogram()
	}

	b.valueType = b.it.Next()
	if b.valueType != chunkenc.ValNone {
		b.lastTime = b.it.AtT()
	}
<<<<<<< HEAD
=======
	if b.valueType == chunkenc.ValHistogram {
		b.valueType = chunkenc.ValFloatHistogram
	}
>>>>>>> main
	return b.valueType
}

// At returns the current float element of the iterator.
func (b *MemoizedSeriesIterator) At() (int64, float64) {
	return b.it.At()
}

<<<<<<< HEAD
// AtHistogram returns the current histogram element of the iterator.
func (b *MemoizedSeriesIterator) AtHistogram() (int64, *histogram.Histogram) {
	return b.it.AtHistogram()
}

=======
>>>>>>> main
// AtFloatHistogram returns the current float-histogram element of the iterator.
func (b *MemoizedSeriesIterator) AtFloatHistogram() (int64, *histogram.FloatHistogram) {
	return b.it.AtFloatHistogram()
}

<<<<<<< HEAD
// AtT returns the current timestamp of the iterator.
func (b *MemoizedSeriesIterator) AtT() int64 {
	return b.it.AtT()
}

=======
>>>>>>> main
// Err returns the last encountered error.
func (b *MemoizedSeriesIterator) Err() error {
	return b.it.Err()
}
