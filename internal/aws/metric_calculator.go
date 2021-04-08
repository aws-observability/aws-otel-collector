// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
)

const (
	cleanInterval = 5 * time.Minute
)

// CalculateFunc defines how to process metric values by the calculator. It
// passes previously received MetricValue, and the current raw value and timestamp
// as parameters. Returns true if the calculation is executed successfully.
type CalculateFunc func(prev *MetricValue, val interface{}, timestamp time.Time) (interface{}, bool)

func NewFloat64DeltaCalculator() MetricCalculator {
	return NewMetricCalculator(calculateDelta)
}

func calculateDelta(prev *MetricValue, val interface{}, timestamp time.Time) (interface{}, bool) {
	deltaValue := val.(float64)
	if prev != nil {
		deltaValue = deltaValue - prev.RawValue.(float64)
		if deltaValue < 0 {
			return float64(0), true
		}
	}
	return deltaValue, true
}

// MetricCalculator is a calculator used to adjust metric values based on its previous record.
type MetricCalculator struct {
	// lock on write
	lock sync.Mutex
	// cache stores data with expiry time. The expiry is not supported at the moment.
	cache *MapWithExpiry
	// calculateFunc is the delegation for data processing
	calculateFunc CalculateFunc
}

func NewMetricCalculator(calculateFunc CalculateFunc) MetricCalculator {
	return MetricCalculator{
		cache:         NewMapWithExpiry(cleanInterval),
		calculateFunc: calculateFunc,
	}
}

// Calculate accepts a new metric value identified by matricName and labels, and delegates
// the calculation with value and timestamp back to CalculateFunc for the result. Returns
// true if the calculation is executed successfully.
func (rm *MetricCalculator) Calculate(metricName string, labels map[string]string, value interface{}, timestamp time.Time) (interface{}, bool) {
	k := NewKey(metricName, labels)
	cacheStore := rm.cache

	var result interface{}
	done := false

	rm.lock.Lock()
	defer rm.lock.Unlock()

	prev, exists := cacheStore.Get(k)
	result, done = rm.calculateFunc(prev, value, timestamp)
	if !exists || done {
		cacheStore.Set(k, MetricValue{
			RawValue:  value,
			Timestamp: timestamp,
		})
	}
	return result, done
}

type Key struct {
	MetricName   string
	MetricLabels attribute.Distinct
}

func NewKey(metricName string, labels map[string]string) Key {
	var kvs []attribute.KeyValue
	var sortable attribute.Sortable
	for k, v := range labels {
		kvs = append(kvs, attribute.String(k, v))
	}
	set := attribute.NewSetWithSortable(kvs, &sortable)

	dedupSortedLabels := set.Equivalent()
	return Key{
		MetricName:   metricName,
		MetricLabels: dedupSortedLabels,
	}
}

type MetricValue struct {
	RawValue  interface{}
	Timestamp time.Time
}

// MapWithExpiry act like a map which provide a method to clean up expired entries
type MapWithExpiry struct {
	lock    *sync.Mutex
	ttl     time.Duration
	entries map[interface{}]*MetricValue
}

func NewMapWithExpiry(ttl time.Duration) *MapWithExpiry {
	return &MapWithExpiry{lock: &sync.Mutex{}, ttl: ttl, entries: make(map[interface{}]*MetricValue)}
}

func (m *MapWithExpiry) Get(key Key) (*MetricValue, bool) {
	v, ok := m.entries[key]
	return v, ok
}

func (m *MapWithExpiry) Set(key Key, value MetricValue) {
	m.entries[key] = &value
}

func (m *MapWithExpiry) CleanUp(now time.Time) {
	for k, v := range m.entries {
		if now.Sub(v.Timestamp) >= m.ttl {
			delete(m.entries, k)
		}
	}
}

func (m *MapWithExpiry) Size() int {
	return len(m.entries)
}

func (m *MapWithExpiry) Lock() {
	m.lock.Lock()
}

func (m *MapWithExpiry) Unlock() {
	m.lock.Unlock()
}
