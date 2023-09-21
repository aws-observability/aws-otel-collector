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

package metric // import "go.opentelemetry.io/otel/sdk/metric"

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
<<<<<<< HEAD
=======
	"sync/atomic"
>>>>>>> main

	"go.opentelemetry.io/otel/internal/global"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/sdk/instrumentation"
<<<<<<< HEAD
	"go.opentelemetry.io/otel/sdk/metric/aggregation"
	"go.opentelemetry.io/otel/sdk/metric/internal"
=======
	"go.opentelemetry.io/otel/sdk/metric/internal"
	"go.opentelemetry.io/otel/sdk/metric/internal/aggregate"
>>>>>>> main
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
)

var (
	errCreatingAggregators     = errors.New("could not create all aggregators")
	errIncompatibleAggregation = errors.New("incompatible aggregation")
	errUnknownAggregation      = errors.New("unrecognized aggregation")
<<<<<<< HEAD
	errUnknownTemporality      = errors.New("unrecognized temporality")
)

type aggregator interface {
	Aggregation() metricdata.Aggregation
}

// instrumentSync is a synchronization point between a pipeline and an
// instrument's Aggregators.
=======
)

// instrumentSync is a synchronization point between a pipeline and an
// instrument's aggregate function.
>>>>>>> main
type instrumentSync struct {
	name        string
	description string
	unit        string
<<<<<<< HEAD
	aggregator  aggregator
=======
	compAgg     aggregate.ComputeAggregation
>>>>>>> main
}

func newPipeline(res *resource.Resource, reader Reader, views []View) *pipeline {
	if res == nil {
		res = resource.Empty()
	}
	return &pipeline{
<<<<<<< HEAD
		resource:     res,
		reader:       reader,
		views:        views,
		aggregations: make(map[instrumentation.Scope][]instrumentSync),
=======
		resource: res,
		reader:   reader,
		views:    views,
		// aggregations is lazy allocated when needed.
>>>>>>> main
	}
}

// pipeline connects all of the instruments created by a meter provider to a Reader.
// This is the object that will be `Reader.register()` when a meter provider is created.
//
<<<<<<< HEAD
// As instruments are created the instrument should be checked if it exists in the
// views of a the Reader, and if so each aggregator should be added to the pipeline.
=======
// As instruments are created the instrument should be checked if it exists in
// the views of a the Reader, and if so each aggregate function should be added
// to the pipeline.
>>>>>>> main
type pipeline struct {
	resource *resource.Resource

	reader Reader
	views  []View

	sync.Mutex
	aggregations   map[instrumentation.Scope][]instrumentSync
	callbacks      []func(context.Context) error
	multiCallbacks list.List
}

// addSync adds the instrumentSync to pipeline p with scope. This method is not
// idempotent. Duplicate calls will result in duplicate additions, it is the
// callers responsibility to ensure this is called with unique values.
func (p *pipeline) addSync(scope instrumentation.Scope, iSync instrumentSync) {
	p.Lock()
	defer p.Unlock()
	if p.aggregations == nil {
		p.aggregations = map[instrumentation.Scope][]instrumentSync{
			scope: {iSync},
		}
		return
	}
	p.aggregations[scope] = append(p.aggregations[scope], iSync)
}

// addCallback registers a single instrument callback to be run when
// `produce()` is called.
func (p *pipeline) addCallback(cback func(context.Context) error) {
	p.Lock()
	defer p.Unlock()
	p.callbacks = append(p.callbacks, cback)
}

type multiCallback func(context.Context) error

// addMultiCallback registers a multi-instrument callback to be run when
// `produce()` is called.
func (p *pipeline) addMultiCallback(c multiCallback) (unregister func()) {
	p.Lock()
	defer p.Unlock()
	e := p.multiCallbacks.PushBack(c)
	return func() {
		p.Lock()
		p.multiCallbacks.Remove(e)
		p.Unlock()
	}
}

// produce returns aggregated metrics from a single collection.
//
// This method is safe to call concurrently.
func (p *pipeline) produce(ctx context.Context, rm *metricdata.ResourceMetrics) error {
	p.Lock()
	defer p.Unlock()

	var errs multierror
	for _, c := range p.callbacks {
		// TODO make the callbacks parallel. ( #3034 )
		if err := c(ctx); err != nil {
			errs.append(err)
		}
		if err := ctx.Err(); err != nil {
			rm.Resource = nil
			rm.ScopeMetrics = rm.ScopeMetrics[:0]
			return err
		}
	}
	for e := p.multiCallbacks.Front(); e != nil; e = e.Next() {
		// TODO make the callbacks parallel. ( #3034 )
		f := e.Value.(multiCallback)
		if err := f(ctx); err != nil {
			errs.append(err)
		}
		if err := ctx.Err(); err != nil {
			// This means the context expired before we finished running callbacks.
			rm.Resource = nil
			rm.ScopeMetrics = rm.ScopeMetrics[:0]
			return err
		}
	}

	rm.Resource = p.resource
	rm.ScopeMetrics = internal.ReuseSlice(rm.ScopeMetrics, len(p.aggregations))

	i := 0
	for scope, instruments := range p.aggregations {
		rm.ScopeMetrics[i].Metrics = internal.ReuseSlice(rm.ScopeMetrics[i].Metrics, len(instruments))
		j := 0
		for _, inst := range instruments {
<<<<<<< HEAD
			data := inst.aggregator.Aggregation()
			if data != nil {
=======
			data := rm.ScopeMetrics[i].Metrics[j].Data
			if n := inst.compAgg(&data); n > 0 {
>>>>>>> main
				rm.ScopeMetrics[i].Metrics[j].Name = inst.name
				rm.ScopeMetrics[i].Metrics[j].Description = inst.description
				rm.ScopeMetrics[i].Metrics[j].Unit = inst.unit
				rm.ScopeMetrics[i].Metrics[j].Data = data
				j++
			}
		}
		rm.ScopeMetrics[i].Metrics = rm.ScopeMetrics[i].Metrics[:j]
		if len(rm.ScopeMetrics[i].Metrics) > 0 {
			rm.ScopeMetrics[i].Scope = scope
			i++
		}
	}

	rm.ScopeMetrics = rm.ScopeMetrics[:i]

	return errs.errorOrNil()
}

// inserter facilitates inserting of new instruments from a single scope into a
// pipeline.
type inserter[N int64 | float64] struct {
<<<<<<< HEAD
	// aggregators is a cache that holds Aggregators inserted into the
	// underlying reader pipeline. This cache ensures no duplicate Aggregators
	// are inserted into the reader pipeline and if a new request during an
	// instrument creation asks for the same Aggregator the same instance is
	// returned.
	aggregators *cache[streamID, aggVal[N]]
=======
	// aggregators is a cache that holds aggregate function inputs whose
	// outputs have been inserted into the underlying reader pipeline. This
	// cache ensures no duplicate aggregate functions are inserted into the
	// reader pipeline and if a new request during an instrument creation asks
	// for the same aggregate function input the same instance is returned.
	aggregators *cache[instID, aggVal[N]]
>>>>>>> main

	// views is a cache that holds instrument identifiers for all the
	// instruments a Meter has created, it is provided from the Meter that owns
	// this inserter. This cache ensures during the creation of instruments
	// with the same name but different options (e.g. description, unit) a
	// warning message is logged.
<<<<<<< HEAD
	views *cache[string, streamID]
=======
	views *cache[string, instID]
>>>>>>> main

	pipeline *pipeline
}

<<<<<<< HEAD
func newInserter[N int64 | float64](p *pipeline, vc *cache[string, streamID]) *inserter[N] {
	if vc == nil {
		vc = &cache[string, streamID]{}
	}
	return &inserter[N]{
		aggregators: &cache[streamID, aggVal[N]]{},
=======
func newInserter[N int64 | float64](p *pipeline, vc *cache[string, instID]) *inserter[N] {
	if vc == nil {
		vc = &cache[string, instID]{}
	}
	return &inserter[N]{
		aggregators: &cache[instID, aggVal[N]]{},
>>>>>>> main
		views:       vc,
		pipeline:    p,
	}
}

// Instrument inserts the instrument inst with instUnit into a pipeline. All
// views the pipeline contains are matched against, and any matching view that
<<<<<<< HEAD
// creates a unique Aggregator will be inserted into the pipeline and included
// in the returned slice.
//
// The returned Aggregators are ensured to be deduplicated and unique. If
// another view in another pipeline that is cached by this inserter's cache has
// already inserted the same Aggregator for the same instrument, that
// Aggregator instance is returned.
//
// If another instrument has already been inserted by this inserter, or any
// other using the same cache, and it conflicts with the instrument being
// inserted in this call, an Aggregator matching the arguments will still be
// returned but an Info level log message will also be logged to the OTel
// global logger.
//
// If the passed instrument would result in an incompatible Aggregator, an
// error is returned and that Aggregator is not inserted or returned.
//
// If an instrument is determined to use a Drop aggregation, that instrument is
// not inserted nor returned.
func (i *inserter[N]) Instrument(inst Instrument) ([]internal.Aggregator[N], error) {
	var (
		matched bool
		aggs    []internal.Aggregator[N]
	)

	errs := &multierror{wrapped: errCreatingAggregators}
	// The cache will return the same Aggregator instance. Use this fact to
	// compare pointer addresses to deduplicate Aggregators.
	seen := make(map[internal.Aggregator[N]]struct{})
=======
// creates a unique aggregate function will have its output inserted into the
// pipeline and its input included in the returned slice.
//
// The returned aggregate function inputs are ensured to be deduplicated and
// unique. If another view in another pipeline that is cached by this
// inserter's cache has already inserted the same aggregate function for the
// same instrument, that functions input instance is returned.
//
// If another instrument has already been inserted by this inserter, or any
// other using the same cache, and it conflicts with the instrument being
// inserted in this call, an aggregate function input matching the arguments
// will still be returned but an Info level log message will also be logged to
// the OTel global logger.
//
// If the passed instrument would result in an incompatible aggregate function,
// an error is returned and that aggregate function output is not inserted nor
// is its input returned.
//
// If an instrument is determined to use a Drop aggregation, that instrument is
// not inserted nor returned.
func (i *inserter[N]) Instrument(inst Instrument) ([]aggregate.Measure[N], error) {
	var (
		matched  bool
		measures []aggregate.Measure[N]
	)

	errs := &multierror{wrapped: errCreatingAggregators}
	seen := make(map[uint64]struct{})
>>>>>>> main
	for _, v := range i.pipeline.views {
		stream, match := v(inst)
		if !match {
			continue
		}
		matched = true

<<<<<<< HEAD
		agg, err := i.cachedAggregator(inst.Scope, inst.Kind, stream)
		if err != nil {
			errs.append(err)
		}
		if agg == nil { // Drop aggregator.
			continue
		}
		if _, ok := seen[agg]; ok {
			// This aggregator has already been added.
			continue
		}
		seen[agg] = struct{}{}
		aggs = append(aggs, agg)
	}

	if matched {
		return aggs, errs.errorOrNil()
=======
		in, id, err := i.cachedAggregator(inst.Scope, inst.Kind, stream)
		if err != nil {
			errs.append(err)
		}
		if in == nil { // Drop aggregation.
			continue
		}
		if _, ok := seen[id]; ok {
			// This aggregate function has already been added.
			continue
		}
		seen[id] = struct{}{}
		measures = append(measures, in)
	}

	if matched {
		return measures, errs.errorOrNil()
>>>>>>> main
	}

	// Apply implicit default view if no explicit matched.
	stream := Stream{
		Name:        inst.Name,
		Description: inst.Description,
		Unit:        inst.Unit,
	}
<<<<<<< HEAD
	agg, err := i.cachedAggregator(inst.Scope, inst.Kind, stream)
	if err != nil {
		errs.append(err)
	}
	if agg != nil {
		// Ensured to have not seen given matched was false.
		aggs = append(aggs, agg)
	}
	return aggs, errs.errorOrNil()
}

// aggVal is the cached value in an aggregators cache.
type aggVal[N int64 | float64] struct {
	Aggregator internal.Aggregator[N]
	Err        error
}

// cachedAggregator returns the appropriate Aggregator for an instrument
// configuration. If the exact instrument has been created within the
// inst.Scope, that Aggregator instance will be returned. Otherwise, a new
// computed Aggregator will be cached and returned.
//
// If the instrument configuration conflicts with an instrument that has
// already been created (e.g. description, unit, data type) a warning will be
// logged at the "Info" level with the global OTel logger. A valid new
// Aggregator for the instrument configuration will still be returned without
// an error.
//
// If the instrument defines an unknown or incompatible aggregation, an error
// is returned.
func (i *inserter[N]) cachedAggregator(scope instrumentation.Scope, kind InstrumentKind, stream Stream) (internal.Aggregator[N], error) {
	switch stream.Aggregation.(type) {
	case nil, aggregation.Default:
		// Undefined, nil, means to use the default from the reader.
		stream.Aggregation = i.pipeline.reader.aggregation(kind)
	}

	if err := isAggregatorCompatible(kind, stream.Aggregation); err != nil {
		return nil, fmt.Errorf(
=======
	in, _, err := i.cachedAggregator(inst.Scope, inst.Kind, stream)
	if err != nil {
		errs.append(err)
	}
	if in != nil {
		// Ensured to have not seen given matched was false.
		measures = append(measures, in)
	}
	return measures, errs.errorOrNil()
}

var aggIDCount uint64

// aggVal is the cached value in an aggregators cache.
type aggVal[N int64 | float64] struct {
	ID      uint64
	Measure aggregate.Measure[N]
	Err     error
}

// cachedAggregator returns the appropriate aggregate input and output
// functions for an instrument configuration. If the exact instrument has been
// created within the inst.Scope, those aggregate function instances will be
// returned. Otherwise, new computed aggregate functions will be cached and
// returned.
//
// If the instrument configuration conflicts with an instrument that has
// already been created (e.g. description, unit, data type) a warning will be
// logged at the "Info" level with the global OTel logger. Valid new aggregate
// functions for the instrument configuration will still be returned without an
// error.
//
// If the instrument defines an unknown or incompatible aggregation, an error
// is returned.
func (i *inserter[N]) cachedAggregator(scope instrumentation.Scope, kind InstrumentKind, stream Stream) (meas aggregate.Measure[N], aggID uint64, err error) {
	switch stream.Aggregation.(type) {
	case nil:
		// Undefined, nil, means to use the default from the reader.
		stream.Aggregation = i.pipeline.reader.aggregation(kind)
		switch stream.Aggregation.(type) {
		case nil, AggregationDefault:
			// If the reader returns default or nil use the default selector.
			stream.Aggregation = DefaultAggregationSelector(kind)
		default:
			// Deep copy and validate before using.
			stream.Aggregation = stream.Aggregation.copy()
			if err := stream.Aggregation.err(); err != nil {
				orig := stream.Aggregation
				stream.Aggregation = DefaultAggregationSelector(kind)
				global.Error(
					err, "using default aggregation instead",
					"aggregation", orig,
					"replacement", stream.Aggregation,
				)
			}
		}
	case AggregationDefault:
		stream.Aggregation = DefaultAggregationSelector(kind)
	}

	if err := isAggregatorCompatible(kind, stream.Aggregation); err != nil {
		return nil, 0, fmt.Errorf(
>>>>>>> main
			"creating aggregator with instrumentKind: %d, aggregation %v: %w",
			kind, stream.Aggregation, err,
		)
	}

<<<<<<< HEAD
	id := i.streamID(kind, stream)
	// If there is a conflict, the specification says the view should
	// still be applied and a warning should be logged.
	i.logConflict(id)
	cv := i.aggregators.Lookup(id, func() aggVal[N] {
		agg, err := i.aggregator(stream.Aggregation, kind, id.Temporality, id.Monotonic)
		if err != nil {
			return aggVal[N]{nil, err}
		}
		if agg == nil { // Drop aggregator.
			return aggVal[N]{nil, nil}
		}
		if stream.AttributeFilter != nil {
			agg = internal.NewFilter(agg, stream.AttributeFilter)
		}

		i.pipeline.addSync(scope, instrumentSync{
			name:        stream.Name,
			description: stream.Description,
			unit:        stream.Unit,
			aggregator:  agg,
		})
		return aggVal[N]{agg, err}
	})
	return cv.Aggregator, cv.Err
}

// logConflict validates if an instrument with the same name as id has already
// been created. If that instrument conflicts with id, a warning is logged.
func (i *inserter[N]) logConflict(id streamID) {
	existing := i.views.Lookup(id.Name, func() streamID { return id })
=======
	id := i.instID(kind, stream)
	// If there is a conflict, the specification says the view should
	// still be applied and a warning should be logged.
	i.logConflict(id)

	// If there are requests for the same instrument with different name
	// casing, the first-seen needs to be returned. Use a normalize ID for the
	// cache lookup to ensure the correct comparison.
	normID := id.normalize()
	cv := i.aggregators.Lookup(normID, func() aggVal[N] {
		b := aggregate.Builder[N]{
			Temporality: i.pipeline.reader.temporality(kind),
		}
		b.Filter = stream.AttributeFilter
		in, out, err := i.aggregateFunc(b, stream.Aggregation, kind)
		if err != nil {
			return aggVal[N]{0, nil, err}
		}
		if in == nil { // Drop aggregator.
			return aggVal[N]{0, nil, nil}
		}
		i.pipeline.addSync(scope, instrumentSync{
			// Use the first-seen name casing for this and all subsequent
			// requests of this instrument.
			name:        stream.Name,
			description: stream.Description,
			unit:        stream.Unit,
			compAgg:     out,
		})
		id := atomic.AddUint64(&aggIDCount, 1)
		return aggVal[N]{id, in, err}
	})
	return cv.Measure, cv.ID, cv.Err
}

// logConflict validates if an instrument with the same case-insensitive name
// as id has already been created. If that instrument conflicts with id, a
// warning is logged.
func (i *inserter[N]) logConflict(id instID) {
	// The API specification defines names as case-insensitive. If there is a
	// different casing of a name it needs to be a conflict.
	name := id.normalize().Name
	existing := i.views.Lookup(name, func() instID { return id })
>>>>>>> main
	if id == existing {
		return
	}

<<<<<<< HEAD
	global.Info(
		"duplicate metric stream definitions",
		"names", fmt.Sprintf("%q, %q", existing.Name, id.Name),
		"descriptions", fmt.Sprintf("%q, %q", existing.Description, id.Description),
		"units", fmt.Sprintf("%s, %s", existing.Unit, id.Unit),
		"numbers", fmt.Sprintf("%s, %s", existing.Number, id.Number),
		"aggregations", fmt.Sprintf("%s, %s", existing.Aggregation, id.Aggregation),
		"monotonics", fmt.Sprintf("%t, %t", existing.Monotonic, id.Monotonic),
		"temporalities", fmt.Sprintf("%s, %s", existing.Temporality.String(), id.Temporality.String()),
	)
}

func (i *inserter[N]) streamID(kind InstrumentKind, stream Stream) streamID {
	var zero N
	id := streamID{
		Name:        stream.Name,
		Description: stream.Description,
		Unit:        stream.Unit,
		Aggregation: fmt.Sprintf("%T", stream.Aggregation),
		Temporality: i.pipeline.reader.temporality(kind),
		Number:      fmt.Sprintf("%T", zero),
	}

	switch kind {
	case InstrumentKindObservableCounter, InstrumentKindCounter, InstrumentKindHistogram:
		id.Monotonic = true
	}

	return id
}

// aggregator returns a new Aggregator matching agg, kind, temporality, and
// monotonic. If the agg is unknown or temporality is invalid, an error is
// returned.
func (i *inserter[N]) aggregator(agg aggregation.Aggregation, kind InstrumentKind, temporality metricdata.Temporality, monotonic bool) (internal.Aggregator[N], error) {
	switch a := agg.(type) {
	case aggregation.Default:
		return i.aggregator(DefaultAggregationSelector(kind), kind, temporality, monotonic)
	case aggregation.Drop:
		return nil, nil
	case aggregation.LastValue:
		return internal.NewLastValue[N](), nil
	case aggregation.Sum:
		switch kind {
		case InstrumentKindObservableCounter, InstrumentKindObservableUpDownCounter:
			// Observable counters and up-down-counters are defined to record
			// the absolute value of the count:
			// https://github.com/open-telemetry/opentelemetry-specification/blob/v1.20.0/specification/metrics/api.md#asynchronous-counter-creation
			switch temporality {
			case metricdata.CumulativeTemporality:
				return internal.NewPrecomputedCumulativeSum[N](monotonic), nil
			case metricdata.DeltaTemporality:
				return internal.NewPrecomputedDeltaSum[N](monotonic), nil
			default:
				return nil, fmt.Errorf("%w: %s(%d)", errUnknownTemporality, temporality.String(), temporality)
			}
		}

		switch temporality {
		case metricdata.CumulativeTemporality:
			return internal.NewCumulativeSum[N](monotonic), nil
		case metricdata.DeltaTemporality:
			return internal.NewDeltaSum[N](monotonic), nil
		default:
			return nil, fmt.Errorf("%w: %s(%d)", errUnknownTemporality, temporality.String(), temporality)
		}
	case aggregation.ExplicitBucketHistogram:
		switch temporality {
		case metricdata.CumulativeTemporality:
			return internal.NewCumulativeHistogram[N](a), nil
		case metricdata.DeltaTemporality:
			return internal.NewDeltaHistogram[N](a), nil
		default:
			return nil, fmt.Errorf("%w: %s(%d)", errUnknownTemporality, temporality.String(), temporality)
		}
	}
	return nil, errUnknownAggregation
=======
	const msg = "duplicate metric stream definitions"
	args := []interface{}{
		"names", fmt.Sprintf("%q, %q", existing.Name, id.Name),
		"descriptions", fmt.Sprintf("%q, %q", existing.Description, id.Description),
		"kinds", fmt.Sprintf("%s, %s", existing.Kind, id.Kind),
		"units", fmt.Sprintf("%s, %s", existing.Unit, id.Unit),
		"numbers", fmt.Sprintf("%s, %s", existing.Number, id.Number),
	}

	// The specification recommends logging a suggested view to resolve
	// conflicts if possible.
	//
	// https://github.com/open-telemetry/opentelemetry-specification/blob/v1.21.0/specification/metrics/sdk.md#duplicate-instrument-registration
	if id.Unit != existing.Unit || id.Number != existing.Number {
		// There is no view resolution for these, don't make a suggestion.
		global.Warn(msg, args...)
		return
	}

	var stream string
	if id.Name != existing.Name || id.Kind != existing.Kind {
		stream = `Stream{Name: "{{NEW_NAME}}"}`
	} else if id.Description != existing.Description {
		stream = fmt.Sprintf("Stream{Description: %q}", existing.Description)
	}

	inst := fmt.Sprintf(
		"Instrument{Name: %q, Description: %q, Kind: %q, Unit: %q}",
		id.Name, id.Description, "InstrumentKind"+id.Kind.String(), id.Unit,
	)
	args = append(args, "suggested.view", fmt.Sprintf("NewView(%s, %s)", inst, stream))

	global.Warn(msg, args...)
}

func (i *inserter[N]) instID(kind InstrumentKind, stream Stream) instID {
	var zero N
	return instID{
		Name:        stream.Name,
		Description: stream.Description,
		Unit:        stream.Unit,
		Kind:        kind,
		Number:      fmt.Sprintf("%T", zero),
	}
}

// aggregateFunc returns new aggregate functions matching agg, kind, and
// monotonic. If the agg is unknown or temporality is invalid, an error is
// returned.
func (i *inserter[N]) aggregateFunc(b aggregate.Builder[N], agg Aggregation, kind InstrumentKind) (meas aggregate.Measure[N], comp aggregate.ComputeAggregation, err error) {
	switch a := agg.(type) {
	case AggregationDefault:
		return i.aggregateFunc(b, DefaultAggregationSelector(kind), kind)
	case AggregationDrop:
		// Return nil in and out to signify the drop aggregator.
	case AggregationLastValue:
		meas, comp = b.LastValue()
	case AggregationSum:
		switch kind {
		case InstrumentKindObservableCounter:
			meas, comp = b.PrecomputedSum(true)
		case InstrumentKindObservableUpDownCounter:
			meas, comp = b.PrecomputedSum(false)
		case InstrumentKindCounter, InstrumentKindHistogram:
			meas, comp = b.Sum(true)
		default:
			// InstrumentKindUpDownCounter, InstrumentKindObservableGauge, and
			// instrumentKindUndefined or other invalid instrument kinds.
			meas, comp = b.Sum(false)
		}
	case AggregationExplicitBucketHistogram:
		var noSum bool
		switch kind {
		case InstrumentKindUpDownCounter, InstrumentKindObservableUpDownCounter, InstrumentKindObservableGauge:
			// The sum should not be collected for any instrument that can make
			// negative measurements:
			// https://github.com/open-telemetry/opentelemetry-specification/blob/v1.21.0/specification/metrics/sdk.md#histogram-aggregations
			noSum = true
		}
		meas, comp = b.ExplicitBucketHistogram(a.Boundaries, a.NoMinMax, noSum)
	case AggregationBase2ExponentialHistogram:
		var noSum bool
		switch kind {
		case InstrumentKindUpDownCounter, InstrumentKindObservableUpDownCounter, InstrumentKindObservableGauge:
			// The sum should not be collected for any instrument that can make
			// negative measurements:
			// https://github.com/open-telemetry/opentelemetry-specification/blob/v1.21.0/specification/metrics/sdk.md#histogram-aggregations
			noSum = true
		}
		meas, comp = b.ExponentialBucketHistogram(a.MaxSize, a.MaxScale, a.NoMinMax, noSum)

	default:
		err = errUnknownAggregation
	}

	return meas, comp, err
>>>>>>> main
}

// isAggregatorCompatible checks if the aggregation can be used by the instrument.
// Current compatibility:
//
// | Instrument Kind          | Drop | LastValue | Sum | Histogram | Exponential Histogram |
// |--------------------------|------|-----------|-----|-----------|-----------------------|
<<<<<<< HEAD
// | Counter                  | X    |           | X   | X         | X                     |
// | UpDownCounter            | X    |           | X   |           |                       |
// | Histogram                | X    |           | X   | X         | X                     |
// | Observable Counter       | X    |           | X   |           |                       |
// | Observable UpDownCounter | X    |           | X   |           |                       |
// | Observable Gauge         | X    | X         |     |           |                       |.
func isAggregatorCompatible(kind InstrumentKind, agg aggregation.Aggregation) error {
	switch agg.(type) {
	case aggregation.Default:
		return nil
	case aggregation.ExplicitBucketHistogram:
		if kind == InstrumentKindCounter || kind == InstrumentKindHistogram {
			return nil
		}
		// TODO: review need for aggregation check after
		// https://github.com/open-telemetry/opentelemetry-specification/issues/2710
		return errIncompatibleAggregation
	case aggregation.Sum:
=======
// | Counter                  | ✓    |           | ✓   | ✓         | ✓                     |
// | UpDownCounter            | ✓    |           | ✓   | ✓         | ✓                     |
// | Histogram                | ✓    |           | ✓   | ✓         | ✓                     |
// | Observable Counter       | ✓    |           | ✓   | ✓         | ✓                     |
// | Observable UpDownCounter | ✓    |           | ✓   | ✓         | ✓                     |
// | Observable Gauge         | ✓    | ✓         |     | ✓         | ✓                     |.
func isAggregatorCompatible(kind InstrumentKind, agg Aggregation) error {
	switch agg.(type) {
	case AggregationDefault:
		return nil
	case AggregationExplicitBucketHistogram, AggregationBase2ExponentialHistogram:
		switch kind {
		case InstrumentKindCounter,
			InstrumentKindUpDownCounter,
			InstrumentKindHistogram,
			InstrumentKindObservableCounter,
			InstrumentKindObservableUpDownCounter,
			InstrumentKindObservableGauge:
			return nil
		default:
			return errIncompatibleAggregation
		}
	case AggregationSum:
>>>>>>> main
		switch kind {
		case InstrumentKindObservableCounter, InstrumentKindObservableUpDownCounter, InstrumentKindCounter, InstrumentKindHistogram, InstrumentKindUpDownCounter:
			return nil
		default:
			// TODO: review need for aggregation check after
			// https://github.com/open-telemetry/opentelemetry-specification/issues/2710
			return errIncompatibleAggregation
		}
<<<<<<< HEAD
	case aggregation.LastValue:
=======
	case AggregationLastValue:
>>>>>>> main
		if kind == InstrumentKindObservableGauge {
			return nil
		}
		// TODO: review need for aggregation check after
		// https://github.com/open-telemetry/opentelemetry-specification/issues/2710
		return errIncompatibleAggregation
<<<<<<< HEAD
	case aggregation.Drop:
=======
	case AggregationDrop:
>>>>>>> main
		return nil
	default:
		// This is used passed checking for default, it should be an error at this point.
		return fmt.Errorf("%w: %v", errUnknownAggregation, agg)
	}
}

// pipelines is the group of pipelines connecting Readers with instrument
// measurement.
type pipelines []*pipeline

func newPipelines(res *resource.Resource, readers []Reader, views []View) pipelines {
	pipes := make([]*pipeline, 0, len(readers))
	for _, r := range readers {
<<<<<<< HEAD
		p := &pipeline{
			resource: res,
			reader:   r,
			views:    views,
		}
=======
		p := newPipeline(res, r, views)
>>>>>>> main
		r.register(p)
		pipes = append(pipes, p)
	}
	return pipes
}

func (p pipelines) registerCallback(cback func(context.Context) error) {
	for _, pipe := range p {
		pipe.addCallback(cback)
	}
}

func (p pipelines) registerMultiCallback(c multiCallback) metric.Registration {
	unregs := make([]func(), len(p))
	for i, pipe := range p {
		unregs[i] = pipe.addMultiCallback(c)
	}
	return unregisterFuncs{f: unregs}
}

type unregisterFuncs struct {
	embedded.Registration
	f []func()
}

func (u unregisterFuncs) Unregister() error {
	for _, f := range u.f {
		f()
	}
	return nil
}

<<<<<<< HEAD
// resolver facilitates resolving Aggregators an instrument needs to aggregate
// measurements with while updating all pipelines that need to pull from those
// aggregations.
=======
// resolver facilitates resolving aggregate functions an instrument calls to
// aggregate measurements with while updating all pipelines that need to pull
// from those aggregations.
>>>>>>> main
type resolver[N int64 | float64] struct {
	inserters []*inserter[N]
}

<<<<<<< HEAD
func newResolver[N int64 | float64](p pipelines, vc *cache[string, streamID]) resolver[N] {
=======
func newResolver[N int64 | float64](p pipelines, vc *cache[string, instID]) resolver[N] {
>>>>>>> main
	in := make([]*inserter[N], len(p))
	for i := range in {
		in[i] = newInserter[N](p[i], vc)
	}
	return resolver[N]{in}
}

// Aggregators returns the Aggregators that must be updated by the instrument
// defined by key.
<<<<<<< HEAD
func (r resolver[N]) Aggregators(id Instrument) ([]internal.Aggregator[N], error) {
	var aggs []internal.Aggregator[N]

	errs := &multierror{}
	for _, i := range r.inserters {
		a, err := i.Instrument(id)
		if err != nil {
			errs.append(err)
		}
		aggs = append(aggs, a...)
	}
	return aggs, errs.errorOrNil()
=======
func (r resolver[N]) Aggregators(id Instrument) ([]aggregate.Measure[N], error) {
	var measures []aggregate.Measure[N]

	errs := &multierror{}
	for _, i := range r.inserters {
		in, err := i.Instrument(id)
		if err != nil {
			errs.append(err)
		}
		measures = append(measures, in...)
	}
	return measures, errs.errorOrNil()
>>>>>>> main
}

type multierror struct {
	wrapped error
	errors  []string
}

func (m *multierror) errorOrNil() error {
	if len(m.errors) == 0 {
		return nil
	}
	if m.wrapped == nil {
		return errors.New(strings.Join(m.errors, "; "))
	}
	return fmt.Errorf("%w: %s", m.wrapped, strings.Join(m.errors, "; "))
}

func (m *multierror) append(err error) {
	m.errors = append(m.errors, err.Error())
}
