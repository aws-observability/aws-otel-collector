// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics"

import (
<<<<<<< HEAD
=======
	"errors"
>>>>>>> main
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

func calculateDelta(prev *MetricValue, val interface{}, _ time.Time) (interface{}, bool) {
	var deltaValue float64
	if prev != nil {
		deltaValue = val.(float64) - prev.RawValue.(float64)
	} else {
		return deltaValue, false
	}
	return deltaValue, true
}

// MetricCalculator is a calculator used to adjust metric values based on its previous record.
<<<<<<< HEAD
=======
// Shutdown() must be called to clean up goroutines before program exit.
>>>>>>> main
type MetricCalculator struct {
	// lock on write
	lock sync.Mutex
	// cache stores data with expiry time. The expiry is not supported at the moment.
	cache *MapWithExpiry
	// calculateFunc is the delegation for data processing
	calculateFunc CalculateFunc
}

<<<<<<< HEAD
=======
// NewMetricCalculator Creates a metric calculator that enforces a five-minute time to live on cache entries.
>>>>>>> main
func NewMetricCalculator(calculateFunc CalculateFunc) MetricCalculator {
	return MetricCalculator{
		cache:         NewMapWithExpiry(cleanInterval),
		calculateFunc: calculateFunc,
	}
}

// Calculate accepts a new metric value identified by metric key (consists of metric metadata and labels),
// https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/eacfde3fcbd46ba60a6db0e9a41977390c4883bd/internal/aws/metrics/metric_calculator.go#L88-L91
// and delegates the calculation with value and timestamp back to CalculateFunc for the result. Returns
// true if the calculation is executed successfully.
func (rm *MetricCalculator) Calculate(mKey Key, value interface{}, timestamp time.Time) (interface{}, bool) {
	cacheStore := rm.cache

	var result interface{}
<<<<<<< HEAD
	done := false
=======
	var done bool
>>>>>>> main

	rm.lock.Lock()
	defer rm.lock.Unlock()

<<<<<<< HEAD
=======
	// need to also lock cache to avoid the cleanup from removing entries while they are being processed.
	// This is only likely to happen when data points come in close to expiration date.
	rm.cache.Lock()
	defer rm.cache.Unlock()

>>>>>>> main
	prev, exists := cacheStore.Get(mKey)
	result, done = rm.calculateFunc(prev, value, timestamp)
	if !exists || done {
		cacheStore.Set(mKey, MetricValue{
			RawValue:  value,
			Timestamp: timestamp,
		})
	}
	return result, done
}

<<<<<<< HEAD
=======
func (rm *MetricCalculator) Shutdown() error {
	return rm.cache.Shutdown()
}

>>>>>>> main
type Key struct {
	MetricMetadata interface{}
	MetricLabels   attribute.Distinct
}

func NewKey(metricMetadata interface{}, labels map[string]string) Key {
<<<<<<< HEAD
	var kvs []attribute.KeyValue
=======
	kvs := make([]attribute.KeyValue, 0, len(labels))
>>>>>>> main
	var sortable attribute.Sortable
	for k, v := range labels {
		kvs = append(kvs, attribute.String(k, v))
	}
	set := attribute.NewSetWithSortable(kvs, &sortable)

	dedupSortedLabels := set.Equivalent()
	return Key{
		MetricMetadata: metricMetadata,
		MetricLabels:   dedupSortedLabels,
	}
}

type MetricValue struct {
	RawValue  interface{}
	Timestamp time.Time
}

<<<<<<< HEAD
// MapWithExpiry act like a map which provide a method to clean up expired entries
type MapWithExpiry struct {
	lock    *sync.Mutex
	ttl     time.Duration
	entries map[interface{}]*MetricValue
}

func NewMapWithExpiry(ttl time.Duration) *MapWithExpiry {
	return &MapWithExpiry{lock: &sync.Mutex{}, ttl: ttl, entries: make(map[interface{}]*MetricValue)}
=======
// MapWithExpiry act like a map which provides a method to clean up expired entries.
// MapWithExpiry is not thread safe and locks must be managed by the owner of the Map by the use of Lock() and Unlock()
type MapWithExpiry struct {
	lock     *sync.Mutex
	ttl      time.Duration
	entries  map[interface{}]*MetricValue
	doneChan chan struct{}
}

// NewMapWithExpiry automatically starts a sweeper to enforce the maps TTL. ShutDown() must be called to ensure that these
// go routines are properly cleaned up ShutDown() must be called.
func NewMapWithExpiry(ttl time.Duration) *MapWithExpiry {
	m := &MapWithExpiry{lock: &sync.Mutex{}, ttl: ttl, entries: make(map[interface{}]*MetricValue), doneChan: make(chan struct{})}
	go m.sweep(m.CleanUp)
	return m
>>>>>>> main
}

func (m *MapWithExpiry) Get(key Key) (*MetricValue, bool) {
	v, ok := m.entries[key]
	return v, ok
}

func (m *MapWithExpiry) Set(key Key, value MetricValue) {
	m.entries[key] = &value
}

<<<<<<< HEAD
=======
func (m *MapWithExpiry) sweep(removeFunc func(time2 time.Time)) {
	ticker := time.NewTicker(m.ttl)
	for {
		select {
		case currentTime := <-ticker.C:
			m.lock.Lock()
			removeFunc(currentTime)
			m.lock.Unlock()
		case <-m.doneChan:
			ticker.Stop()
			return
		}
	}
}

func (m *MapWithExpiry) Shutdown() error {
	select {
	case <-m.doneChan:
		return errors.New("shutdown called on an already closed channel")
	default:
		close(m.doneChan)

	}
	return nil
}

>>>>>>> main
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
