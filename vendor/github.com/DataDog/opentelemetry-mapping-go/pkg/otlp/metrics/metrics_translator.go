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

package metrics

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"

	"github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes"
	"github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes/source"
	"github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/metrics/internal/instrumentationlibrary"
	"github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/metrics/internal/instrumentationscope"
	"github.com/DataDog/opentelemetry-mapping-go/pkg/quantile"
)

const (
	metricName             string = "metric name"
	errNoBucketsNoSumCount string = "no buckets mode and no send count sum are incompatible"
)

var _ source.Provider = (*noSourceProvider)(nil)

type noSourceProvider struct{}

func (*noSourceProvider) Source(context.Context) (source.Source, error) {
	return source.Source{Kind: source.HostnameKind, Identifier: ""}, nil
}

// Translator is a metrics translator.
type Translator struct {
	prevPts *ttlCache
	logger  *zap.Logger
	cfg     translatorConfig
}

// Metadata specifies information about the outcome of the MapMetrics call.
type Metadata struct {
	// Languages specifies a list of languages for which runtime metrics were found.
	Languages []string
}

// NewTranslator creates a new translator with given options.
func NewTranslator(logger *zap.Logger, options ...TranslatorOption) (*Translator, error) {
	cfg := translatorConfig{
		HistMode:                             HistogramModeDistributions,
		SendHistogramAggregations:            false,
		Quantiles:                            false,
		NumberMode:                           NumberModeCumulativeToDelta,
		InitialCumulMonoValueMode:            InitialCumulMonoValueModeAuto,
		ResourceAttributesAsTags:             false,
		InstrumentationLibraryMetadataAsTags: false,
		sweepInterval:                        1800,
		deltaTTL:                             3600,
		fallbackSourceProvider:               &noSourceProvider{},
	}

	for _, opt := range options {
		err := opt(&cfg)
		if err != nil {
			return nil, err
		}
	}

	if cfg.HistMode == HistogramModeNoBuckets && !cfg.SendHistogramAggregations {
		return nil, errors.New(errNoBucketsNoSumCount)
	}

	cache := newTTLCache(cfg.sweepInterval, cfg.deltaTTL)
	return &Translator{
		prevPts: cache,
		logger:  logger.With(zap.String("component", "metrics translator")),
		cfg:     cfg,
	}, nil
}

// isCumulativeMonotonic checks if a metric is a cumulative monotonic metric
func isCumulativeMonotonic(md pmetric.Metric) bool {
	switch md.Type() {
	case pmetric.MetricTypeSum:
		return md.Sum().AggregationTemporality() == pmetric.AggregationTemporalityCumulative &&
			md.Sum().IsMonotonic()
	}
	return false
}

// isSkippable checks if a value can be skipped (because it is not supported by the backend).
// It logs that the value is unsupported for debugging since this sometimes means there is a bug.
func (t *Translator) isSkippable(name string, v float64) bool {
	skippable := math.IsInf(v, 0) || math.IsNaN(v)
	if skippable {
		t.logger.Debug("Unsupported metric value", zap.String(metricName, name), zap.Float64("value", v))
	}
	return skippable
}

// mapNumberMetrics maps double datapoints into Datadog metrics
func (t *Translator) mapNumberMetrics(
	ctx context.Context,
	consumer TimeSeriesConsumer,
	dims *Dimensions,
	dt DataType,
	slice pmetric.NumberDataPointSlice,
) {

	for i := 0; i < slice.Len(); i++ {
		p := slice.At(i)
		pointDims := dims.WithAttributeMap(p.Attributes())
		var val float64
		switch p.ValueType() {
		case pmetric.NumberDataPointValueTypeDouble:
			val = p.DoubleValue()
		case pmetric.NumberDataPointValueTypeInt:
			val = float64(p.IntValue())
		}

		if t.isSkippable(pointDims.name, val) {
			continue
		}

		consumer.ConsumeTimeSeries(ctx, pointDims, dt, uint64(p.Timestamp()), val)
	}
}

// TODO(songy23): consider changing this to a Translator start time that must be initialized
// if the package-level variable causes any issue.
var startTime = uint64(time.Now().Unix())

// getProcessStartTime returns the start time of the Agent process in seconds since epoch
func getProcessStartTime() uint64 {
	return startTime
}

// shouldConsumeInitialValue checks if the initial value of a cumulative monotonic metric
// should be consumed or dropped.
func (t *Translator) shouldConsumeInitialValue(startTs, ts uint64) bool {
	switch t.cfg.InitialCumulMonoValueMode {
	case InitialCumulMonoValueModeAuto:
		if getProcessStartTime() < startTs && startTs != ts {
			// Report the first value if the timeseries started after the Datadog Agent process started.
			return true
		}
	case InitialCumulMonoValueModeKeep:
		return true
	case InitialCumulMonoValueModeDrop:
		// do nothing, drop the point
	}
	return false
}

// mapNumberMonotonicMetrics maps monotonic datapoints into Datadog metrics
func (t *Translator) mapNumberMonotonicMetrics(
	ctx context.Context,
	consumer TimeSeriesConsumer,
	dims *Dimensions,
	slice pmetric.NumberDataPointSlice,
) {
	for i := 0; i < slice.Len(); i++ {
		p := slice.At(i)
		ts := uint64(p.Timestamp())
		startTs := uint64(p.StartTimestamp())
		pointDims := dims.WithAttributeMap(p.Attributes())

		var val float64
		switch p.ValueType() {
		case pmetric.NumberDataPointValueTypeDouble:
			val = p.DoubleValue()
		case pmetric.NumberDataPointValueTypeInt:
			val = float64(p.IntValue())
		}

		if t.isSkippable(pointDims.name, val) {
			continue
		}

		if dx, ok := t.prevPts.MonotonicDiff(pointDims, startTs, ts, val); ok {
			consumer.ConsumeTimeSeries(ctx, pointDims, Count, ts, dx)
		} else if i == 0 && t.shouldConsumeInitialValue(startTs, ts) {
			consumer.ConsumeTimeSeries(ctx, pointDims, Count, ts, val)
		}
	}
}

func getBounds(p pmetric.HistogramDataPoint, idx int) (lowerBound float64, upperBound float64) {
	// See https://github.com/open-telemetry/opentelemetry-proto/blob/v0.10.0/opentelemetry/proto/metrics/v1/metrics.proto#L427-L439
	lowerBound = math.Inf(-1)
	upperBound = math.Inf(1)
	if idx > 0 {
		lowerBound = p.ExplicitBounds().At(idx - 1)
	}
	if idx < p.ExplicitBounds().Len() {
		upperBound = p.ExplicitBounds().At(idx)
	}
	return
}

type histogramInfo struct {
	// sum of histogram (exact)
	sum float64
	// count of histogram (exact)
	count uint64

	// hasMinFromLastTimeWindow indicates whether the minimum was reached in the last time window.
	// If the minimum is NOT available, its value is false.
	hasMinFromLastTimeWindow bool

	// hasMaxFromLastTimeWindow indicates whether the maximum was reached in the last time window.
	// If the maximum is NOT available, its value is false.
	hasMaxFromLastTimeWindow bool

	// ok to use sum/count.
	ok bool
}

func (t *Translator) getSketchBuckets(
	ctx context.Context,
	consumer SketchConsumer,
	pointDims *Dimensions,
	p pmetric.HistogramDataPoint,
	histInfo histogramInfo,
	delta bool,
) {
	startTs := uint64(p.StartTimestamp())
	ts := uint64(p.Timestamp())
	as := &quantile.Agent{}
	for j := 0; j < p.BucketCounts().Len(); j++ {
		lowerBound, upperBound := getBounds(p, j)

		// Compute temporary bucketTags to have unique keys in the t.prevPts cache for each bucket
		// The bucketTags are computed from the bounds before the InsertInterpolate fix is done,
		// otherwise in the case where p.MExplicitBounds() has a size of 1 (eg. [0]), the two buckets
		// would have the same bucketTags (lower_bound:0 and upper_bound:0), resulting in a buggy behavior.
		bucketDims := pointDims.AddTags(
			fmt.Sprintf("lower_bound:%s", formatFloat(lowerBound)),
			fmt.Sprintf("upper_bound:%s", formatFloat(upperBound)),
		)

		// InsertInterpolate doesn't work with an infinite bound; insert in to the bucket that contains the non-infinite bound
		// https://github.com/DataDog/datadog-agent/blob/7.31.0/pkg/aggregator/check_sampler.go#L107-L111
		if math.IsInf(upperBound, 1) {
			upperBound = lowerBound
		} else if math.IsInf(lowerBound, -1) {
			lowerBound = upperBound
		}

		count := p.BucketCounts().At(j)
		if delta {
			as.InsertInterpolate(lowerBound, upperBound, uint(count))
		} else if dx, ok := t.prevPts.Diff(bucketDims, startTs, ts, float64(count)); ok {
			as.InsertInterpolate(lowerBound, upperBound, uint(dx))
		}

	}

	sketch := as.Finish()
	if sketch != nil {
		if histInfo.ok {
			// override approximate sum, count and average in sketch with exact values if available.
			sketch.Basic.Cnt = int64(histInfo.count)
			sketch.Basic.Sum = histInfo.sum
			sketch.Basic.Avg = sketch.Basic.Sum / float64(sketch.Basic.Cnt)
		}

		if histInfo.hasMinFromLastTimeWindow {
			// We know exact minimum for the last time window.
			sketch.Basic.Min = p.Min()
		} else if p.HasMin() {
			// Clamp minimum with the global minimum (p.Min()) to account for sketch mapping error.
			sketch.Basic.Min = math.Max(p.Min(), sketch.Basic.Min)
		}

		if histInfo.hasMaxFromLastTimeWindow {
			// We know exact maximum for the last time window.
			sketch.Basic.Max = p.Max()
		} else if p.HasMax() {
			// Clamp maximum with global maximum (p.Max()) to account for sketch mapping error.
			sketch.Basic.Max = math.Min(p.Max(), sketch.Basic.Max)
		}

		consumer.ConsumeSketch(ctx, pointDims, ts, sketch)
	}
}

func (t *Translator) getLegacyBuckets(
	ctx context.Context,
	consumer TimeSeriesConsumer,
	pointDims *Dimensions,
	p pmetric.HistogramDataPoint,
	delta bool,
) {
	startTs := uint64(p.StartTimestamp())
	ts := uint64(p.Timestamp())
	// We have a single metric, 'bucket', which is tagged with the bucket bounds. See:
	// https://github.com/DataDog/integrations-core/blob/7.30.1/datadog_checks_base/datadog_checks/base/checks/openmetrics/v2/transformers/histogram.py
	baseBucketDims := pointDims.WithSuffix("bucket")
	for idx := 0; idx < p.BucketCounts().Len(); idx++ {
		lowerBound, upperBound := getBounds(p, idx)
		bucketDims := baseBucketDims.AddTags(
			fmt.Sprintf("lower_bound:%s", formatFloat(lowerBound)),
			fmt.Sprintf("upper_bound:%s", formatFloat(upperBound)),
		)

		count := float64(p.BucketCounts().At(idx))
		if delta {
			consumer.ConsumeTimeSeries(ctx, bucketDims, Count, ts, count)
		} else if dx, ok := t.prevPts.Diff(bucketDims, startTs, ts, count); ok {
			consumer.ConsumeTimeSeries(ctx, bucketDims, Count, ts, dx)
		}
	}
}

// mapHistogramMetrics maps double histogram metrics slices to Datadog metrics
//
// A Histogram metric has:
// - The count of values in the population
// - The sum of values in the population
// - A number of buckets, each of them having
//   - the bounds that define the bucket
//   - the count of the number of items in that bucket
//   - a sample value from each bucket
//
// We follow a similar approach to our OpenMetrics check:
// we report sum and count by default; buckets count can also
// be reported (opt-in) tagged by lower bound.
func (t *Translator) mapHistogramMetrics(
	ctx context.Context,
	consumer Consumer,
	dims *Dimensions,
	slice pmetric.HistogramDataPointSlice,
	delta bool,
) {
	for i := 0; i < slice.Len(); i++ {
		p := slice.At(i)
		startTs := uint64(p.StartTimestamp())
		ts := uint64(p.Timestamp())
		pointDims := dims.WithAttributeMap(p.Attributes())

		histInfo := histogramInfo{ok: true}

		countDims := pointDims.WithSuffix("count")
		if delta {
			histInfo.count = p.Count()
		} else if dx, ok := t.prevPts.Diff(countDims, startTs, ts, float64(p.Count())); ok {
			histInfo.count = uint64(dx)
		} else { // not ok
			histInfo.ok = false
		}

		sumDims := pointDims.WithSuffix("sum")
		if !t.isSkippable(sumDims.name, p.Sum()) {
			if delta {
				histInfo.sum = p.Sum()
			} else if dx, ok := t.prevPts.Diff(sumDims, startTs, ts, p.Sum()); ok {
				histInfo.sum = dx
			} else { // not ok
				histInfo.ok = false
			}
		} else { // skippable
			histInfo.ok = false
		}

		minDims := pointDims.WithSuffix("min")
		if p.HasMin() {
			histInfo.hasMinFromLastTimeWindow = delta || t.prevPts.PutAndCheckMin(minDims, startTs, ts, p.Min())
		}

		maxDims := pointDims.WithSuffix("max")
		if p.HasMax() {
			histInfo.hasMaxFromLastTimeWindow = delta || t.prevPts.PutAndCheckMax(maxDims, startTs, ts, p.Max())
		}

		if t.cfg.SendHistogramAggregations && histInfo.ok {
			// We only send the sum and count if both values were ok.
			consumer.ConsumeTimeSeries(ctx, countDims, Count, ts, float64(histInfo.count))
			consumer.ConsumeTimeSeries(ctx, sumDims, Count, ts, histInfo.sum)

			if delta {
				// We could check is[Min/Max]FromLastTimeWindow here, and report the minimum/maximum
				// for cumulative timeseries when we know it. These would be metrics with progressively
				// less frequency which would be confusing, so we limit reporting these metrics to delta points,
				// where the min/max is (pressumably) available in either all or none of the points.

				if p.HasMin() {
					consumer.ConsumeTimeSeries(ctx, minDims, Gauge, ts, p.Min())
				}
				if p.HasMax() {
					consumer.ConsumeTimeSeries(ctx, maxDims, Gauge, ts, p.Max())
				}
			}
		}

		switch t.cfg.HistMode {
		case HistogramModeCounters:
			t.getLegacyBuckets(ctx, consumer, pointDims, p, delta)
		case HistogramModeDistributions:
			t.getSketchBuckets(ctx, consumer, pointDims, p, histInfo, delta)
		}
	}
}

// formatFloat formats a float number as close as possible to what
// we do on the Datadog Agent Python OpenMetrics check, which, in turn, tries to
// follow https://github.com/OpenObservability/OpenMetrics/blob/v1.0.0/specification/OpenMetrics.md#considerations-canonical-numbers
func formatFloat(f float64) string {
	if math.IsInf(f, 1) {
		return "inf"
	} else if math.IsInf(f, -1) {
		return "-inf"
	} else if math.IsNaN(f) {
		return "nan"
	} else if f == 0 {
		return "0"
	}

	// Add .0 to whole numbers
	s := strconv.FormatFloat(f, 'g', -1, 64)
	if f == math.Floor(f) {
		s = s + ".0"
	}
	return s
}

// getQuantileTag returns the quantile tag for summary types.
func getQuantileTag(quantile float64) string {
	return fmt.Sprintf("quantile:%s", formatFloat(quantile))
}

// mapSummaryMetrics maps summary datapoints into Datadog metrics
func (t *Translator) mapSummaryMetrics(
	ctx context.Context,
	consumer TimeSeriesConsumer,
	dims *Dimensions,
	slice pmetric.SummaryDataPointSlice,
) {

	for i := 0; i < slice.Len(); i++ {
		p := slice.At(i)
		startTs := uint64(p.StartTimestamp())
		ts := uint64(p.Timestamp())
		pointDims := dims.WithAttributeMap(p.Attributes())

		// count and sum are increasing; we treat them as cumulative monotonic sums.
		{
			countDims := pointDims.WithSuffix("count")
			if dx, ok := t.prevPts.Diff(countDims, startTs, ts, float64(p.Count())); ok && !t.isSkippable(countDims.name, dx) {
				consumer.ConsumeTimeSeries(ctx, countDims, Count, ts, dx)
			}
		}

		{
			sumDims := pointDims.WithSuffix("sum")
			if !t.isSkippable(sumDims.name, p.Sum()) {
				if dx, ok := t.prevPts.Diff(sumDims, startTs, ts, p.Sum()); ok {
					consumer.ConsumeTimeSeries(ctx, sumDims, Count, ts, dx)
				}
			}
		}

		if t.cfg.Quantiles {
			baseQuantileDims := pointDims.WithSuffix("quantile")
			quantiles := p.QuantileValues()
			for i := 0; i < quantiles.Len(); i++ {
				q := quantiles.At(i)

				if t.isSkippable(baseQuantileDims.name, q.Value()) {
					continue
				}

				quantileDims := baseQuantileDims.AddTags(getQuantileTag(q.Quantile()))
				consumer.ConsumeTimeSeries(ctx, quantileDims, Gauge, ts, q.Value())
			}
		}
	}
}

func (t *Translator) source(m pcommon.Map) (source.Source, error) {
	src, ok := attributes.SourceFromAttrs(m)
	if !ok {
		var err error
		src, err = t.cfg.fallbackSourceProvider.Source(context.Background())
		if err != nil {
			return source.Source{}, fmt.Errorf("failed to get fallback source: %w", err)
		}
	}
	return src, nil
}

// extractLanguageTag appends a new language tag to languageTags if a new language tag is found from the given name
func extractLanguageTag(name string, languageTags []string) []string {
	for prefix, lang := range runtimeMetricPrefixLanguageMap {
		if !slices.Contains(languageTags, lang) && strings.HasPrefix(name, prefix) {
			return append(languageTags, lang)
		}
	}
	return languageTags
}

// mapGaugeRuntimeMetricWithAttributes maps the specified runtime metric from metric attributes into a new Gauge metric
func mapGaugeRuntimeMetricWithAttributes(md pmetric.Metric, metricsArray pmetric.MetricSlice, mp runtimeMetricMapping) {
	for i := 0; i < md.Gauge().DataPoints().Len(); i++ {
		matchesAttributes := true
		for _, attribute := range mp.attributes {
			attributeValue, res := md.Gauge().DataPoints().At(i).Attributes().Get(attribute.key)
			if !res || !slices.Contains(attribute.values, attributeValue.AsString()) {
				matchesAttributes = false
				break
			}
		}
		if matchesAttributes {
			cp := metricsArray.AppendEmpty()
			cp.SetEmptyGauge()
			dataPoint := cp.Gauge().DataPoints().AppendEmpty()
			md.Gauge().DataPoints().At(i).CopyTo(dataPoint)
			dataPoint.Attributes().RemoveIf(func(s string, value pcommon.Value) bool {
				for _, attribute := range mp.attributes {
					if s == attribute.key {
						return true
					}
				}
				return false
			})
			cp.SetName(mp.mappedName)
		}
	}
}

// mapSumRuntimeMetricWithAttributes maps the specified runtime metric from metric attributes into a new Sum metric
func mapSumRuntimeMetricWithAttributes(md pmetric.Metric, metricsArray pmetric.MetricSlice, mp runtimeMetricMapping) {
	for i := 0; i < md.Sum().DataPoints().Len(); i++ {
		matchesAttributes := true
		for _, attribute := range mp.attributes {
			attributeValue, res := md.Sum().DataPoints().At(i).Attributes().Get(attribute.key)
			if !res || !slices.Contains(attribute.values, attributeValue.AsString()) {
				matchesAttributes = false
				break
			}
		}
		if matchesAttributes {
			cp := metricsArray.AppendEmpty()
			cp.SetEmptySum()
			cp.Sum().SetAggregationTemporality(md.Sum().AggregationTemporality())
			cp.Sum().SetIsMonotonic(md.Sum().IsMonotonic())
			dataPoint := cp.Sum().DataPoints().AppendEmpty()
			md.Sum().DataPoints().At(i).CopyTo(dataPoint)
			dataPoint.Attributes().RemoveIf(func(s string, value pcommon.Value) bool {
				for _, attribute := range mp.attributes {
					if s == attribute.key {
						return true
					}
				}
				return false
			})
			cp.SetName(mp.mappedName)
		}
	}
}

// mapHistogramRuntimeMetricWithAttributes maps the specified runtime metric from metric attributes into a new Histogram metric
func mapHistogramRuntimeMetricWithAttributes(md pmetric.Metric, metricsArray pmetric.MetricSlice, mp runtimeMetricMapping) {
	for i := 0; i < md.Histogram().DataPoints().Len(); i++ {
		matchesAttributes := true
		for _, attribute := range mp.attributes {
			attributeValue, res := md.Histogram().DataPoints().At(i).Attributes().Get(attribute.key)
			if !res || !slices.Contains(attribute.values, attributeValue.AsString()) {
				matchesAttributes = false
				break
			}
		}
		if matchesAttributes {
			cp := metricsArray.AppendEmpty()
			cp.SetEmptyHistogram()
			cp.Histogram().SetAggregationTemporality(md.Histogram().AggregationTemporality())
			dataPoint := cp.Histogram().DataPoints().AppendEmpty()
			md.Histogram().DataPoints().At(i).CopyTo(dataPoint)
			dataPoint.Attributes().RemoveIf(func(s string, value pcommon.Value) bool {
				for _, attribute := range mp.attributes {
					if s == attribute.key {
						return true
					}
				}
				return false
			})
			cp.SetName(mp.mappedName)
			break
		}
	}
}

// MapMetrics maps OTLP metrics into the Datadog format
func (t *Translator) MapMetrics(ctx context.Context, md pmetric.Metrics, consumer Consumer) (Metadata, error) {
	metadata := Metadata{
		Languages: []string{},
	}
	rms := md.ResourceMetrics()
	for i := 0; i < rms.Len(); i++ {
		rm := rms.At(i)
		if v, ok := rm.Resource().Attributes().Get(keyAPMStats); ok && v.Bool() {
			// these resource metrics are an APM Stats payload; consume it as such
			sp, err := t.statsPayloadFromMetrics(rm)
			if err != nil {
				return metadata, fmt.Errorf("error extracting APM Stats from Metrics: %w", err)
			}
			consumer.ConsumeAPMStats(sp)
			continue
		}
		src, err := t.source(rm.Resource().Attributes())
		if err != nil {
			return metadata, err
		}
		var host string
		switch src.Kind {
		case source.HostnameKind:
			host = src.Identifier
			if c, ok := consumer.(HostConsumer); ok {
				c.ConsumeHost(host)
			}
		case source.AWSECSFargateKind:
			if c, ok := consumer.(TagsConsumer); ok {
				c.ConsumeTag(src.Tag())
			}
		}

		// Fetch tags from attributes.
		attributeTags := attributes.TagsFromAttributes(rm.Resource().Attributes())
		ilms := rm.ScopeMetrics()
		for j := 0; j < ilms.Len(); j++ {
			ilm := ilms.At(j)
			metricsArray := ilm.Metrics()

			var additionalTags []string
			if t.cfg.InstrumentationScopeMetadataAsTags {
				additionalTags = append(attributeTags, instrumentationscope.TagsFromInstrumentationScopeMetadata(ilm.Scope())...)
			} else if t.cfg.InstrumentationLibraryMetadataAsTags {
				additionalTags = append(attributeTags, instrumentationlibrary.TagsFromInstrumentationLibraryMetadata(ilm.Scope())...)
			} else {
				additionalTags = attributeTags
			}

			for k := 0; k < metricsArray.Len(); k++ {
				md := metricsArray.At(k)
				if v, ok := runtimeMetricsMappings[md.Name()]; ok {
					metadata.Languages = extractLanguageTag(md.Name(), metadata.Languages)
					for _, mp := range v {
						if mp.attributes == nil {
							// duplicate runtime metrics as Datadog runtime metrics
							cp := metricsArray.AppendEmpty()
							md.CopyTo(cp)
							cp.SetName(mp.mappedName)
							break
						}
						if md.Type() == pmetric.MetricTypeSum {
							mapSumRuntimeMetricWithAttributes(md, metricsArray, mp)
						} else if md.Type() == pmetric.MetricTypeGauge {
							mapGaugeRuntimeMetricWithAttributes(md, metricsArray, mp)
						} else if md.Type() == pmetric.MetricTypeHistogram {
							mapHistogramRuntimeMetricWithAttributes(md, metricsArray, mp)
						}
					}
				}
				if t.cfg.withRemapping {
					remapMetrics(metricsArray, md)
				}
				baseDims := &Dimensions{
					name:     md.Name(),
					tags:     additionalTags,
					host:     host,
					originID: attributes.OriginIDFromAttributes(rm.Resource().Attributes()),
				}
				switch md.Type() {
				case pmetric.MetricTypeGauge:
					t.mapNumberMetrics(ctx, consumer, baseDims, Gauge, md.Gauge().DataPoints())
				case pmetric.MetricTypeSum:
					switch md.Sum().AggregationTemporality() {
					case pmetric.AggregationTemporalityCumulative:
						if isCumulativeMonotonic(md) {
							switch t.cfg.NumberMode {
							case NumberModeCumulativeToDelta:
								t.mapNumberMonotonicMetrics(ctx, consumer, baseDims, md.Sum().DataPoints())
							case NumberModeRawValue:
								t.mapNumberMetrics(ctx, consumer, baseDims, Gauge, md.Sum().DataPoints())
							}
						} else { // delta and cumulative non-monotonic sums
							t.mapNumberMetrics(ctx, consumer, baseDims, Gauge, md.Sum().DataPoints())
						}
					case pmetric.AggregationTemporalityDelta:
						t.mapNumberMetrics(ctx, consumer, baseDims, Count, md.Sum().DataPoints())
					default: // pmetric.AggregationTemporalityUnspecified or any other not supported type
						t.logger.Debug("Unknown or unsupported aggregation temporality",
							zap.String(metricName, md.Name()),
							zap.Any("aggregation temporality", md.Sum().AggregationTemporality()),
						)
						continue
					}
				case pmetric.MetricTypeHistogram:
					switch md.Histogram().AggregationTemporality() {
					case pmetric.AggregationTemporalityCumulative, pmetric.AggregationTemporalityDelta:
						delta := md.Histogram().AggregationTemporality() == pmetric.AggregationTemporalityDelta
						t.mapHistogramMetrics(ctx, consumer, baseDims, md.Histogram().DataPoints(), delta)
					default: // pmetric.AggregationTemporalityUnspecified or any other not supported type
						t.logger.Debug("Unknown or unsupported aggregation temporality",
							zap.String("metric name", md.Name()),
							zap.Any("aggregation temporality", md.Histogram().AggregationTemporality()),
						)
						continue
					}
				case pmetric.MetricTypeExponentialHistogram:
					switch md.ExponentialHistogram().AggregationTemporality() {
					case pmetric.AggregationTemporalityDelta:
						delta := md.ExponentialHistogram().AggregationTemporality() == pmetric.AggregationTemporalityDelta
						t.mapExponentialHistogramMetrics(ctx, consumer, baseDims, md.ExponentialHistogram().DataPoints(), delta)
					default: // pmetric.AggregationTemporalityCumulative, pmetric.AggregationTemporalityUnspecified or any other not supported type
						t.logger.Debug("Unknown or unsupported aggregation temporality",
							zap.String("metric name", md.Name()),
							zap.Any("aggregation temporality", md.ExponentialHistogram().AggregationTemporality()),
						)
						continue
					}
				case pmetric.MetricTypeSummary:
					t.mapSummaryMetrics(ctx, consumer, baseDims, md.Summary().DataPoints())
				default: // pmetric.MetricDataTypeNone or any other not supported type
					t.logger.Debug("Unknown or unsupported metric type", zap.String(metricName, md.Name()), zap.Any("data type", md.Type()))
					continue
				}
			}
		}
	}
	return metadata, nil
}
