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

package prometheus // import "go.opentelemetry.io/otel/exporters/prometheus"

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"google.golang.org/protobuf/proto"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/internal/global"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
)

const (
	targetInfoMetricName  = "target_info"
	targetInfoDescription = "Target metadata"

	scopeInfoMetricName  = "otel_scope_info"
	scopeInfoDescription = "Instrumentation Scope metadata"
)

var scopeInfoKeys = [2]string{"otel_scope_name", "otel_scope_version"}

// Exporter is a Prometheus Exporter that embeds the OTel metric.Reader
// interface for easy instantiation with a MeterProvider.
type Exporter struct {
	metric.Reader
}

<<<<<<< HEAD
=======
// MarshalLog returns logging data about the Exporter.
func (e *Exporter) MarshalLog() interface{} {
	const t = "Prometheus exporter"

	if r, ok := e.Reader.(*metric.ManualReader); ok {
		under := r.MarshalLog()
		if data, ok := under.(struct {
			Type       string
			Registered bool
			Shutdown   bool
		}); ok {
			data.Type = t
			return data
		}
	}

	return struct{ Type string }{Type: t}
}

>>>>>>> main
var _ metric.Reader = &Exporter{}

// collector is used to implement prometheus.Collector.
type collector struct {
	reader metric.Reader

<<<<<<< HEAD
	disableTargetInfo    bool
	withoutUnits         bool
	targetInfo           prometheus.Metric
	disableScopeInfo     bool
	createTargetInfoOnce sync.Once
	scopeInfos           map[instrumentation.Scope]prometheus.Metric
	metricFamilies       map[string]*dto.MetricFamily
	namespace            string
}

// prometheus counters MUST have a _total suffix:
=======
	withoutUnits           bool
	withoutCounterSuffixes bool
	disableScopeInfo       bool
	namespace              string

	mu                sync.Mutex // mu protects all members below from the concurrent access.
	disableTargetInfo bool
	targetInfo        prometheus.Metric
	scopeInfos        map[instrumentation.Scope]prometheus.Metric
	metricFamilies    map[string]*dto.MetricFamily
}

// prometheus counters MUST have a _total suffix by default:
>>>>>>> main
// https://github.com/open-telemetry/opentelemetry-specification/blob/v1.20.0/specification/compatibility/prometheus_and_openmetrics.md
const counterSuffix = "_total"

// New returns a Prometheus Exporter.
func New(opts ...Option) (*Exporter, error) {
	cfg := newConfig(opts...)

	// this assumes that the default temporality selector will always return cumulative.
	// we only support cumulative temporality, so building our own reader enforces this.
	// TODO (#3244): Enable some way to configure the reader, but not change temporality.
<<<<<<< HEAD
	reader := metric.NewManualReader(cfg.manualReaderOptions()...)

	collector := &collector{
		reader:            reader,
		disableTargetInfo: cfg.disableTargetInfo,
		withoutUnits:      cfg.withoutUnits,
		disableScopeInfo:  cfg.disableScopeInfo,
		scopeInfos:        make(map[instrumentation.Scope]prometheus.Metric),
		metricFamilies:    make(map[string]*dto.MetricFamily),
		namespace:         cfg.namespace,
=======
	reader := metric.NewManualReader(cfg.readerOpts...)

	collector := &collector{
		reader:                 reader,
		disableTargetInfo:      cfg.disableTargetInfo,
		withoutUnits:           cfg.withoutUnits,
		withoutCounterSuffixes: cfg.withoutCounterSuffixes,
		disableScopeInfo:       cfg.disableScopeInfo,
		scopeInfos:             make(map[instrumentation.Scope]prometheus.Metric),
		metricFamilies:         make(map[string]*dto.MetricFamily),
		namespace:              cfg.namespace,
>>>>>>> main
	}

	if err := cfg.registerer.Register(collector); err != nil {
		return nil, fmt.Errorf("cannot register the collector: %w", err)
	}

	e := &Exporter{
		Reader: reader,
	}

	return e, nil
}

// Describe implements prometheus.Collector.
func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	// The Opentelemetry SDK doesn't have information on which will exist when the collector
	// is registered. By returning nothing we are an "unchecked" collector in Prometheus,
	// and assume responsibility for consistency of the metrics produced.
	//
	// See https://pkg.go.dev/github.com/prometheus/client_golang@v1.13.0/prometheus#hdr-Custom_Collectors_and_constant_Metrics
}

// Collect implements prometheus.Collector.
<<<<<<< HEAD
=======
//
// This method is safe to call concurrently.
>>>>>>> main
func (c *collector) Collect(ch chan<- prometheus.Metric) {
	// TODO (#3047): Use a sync.Pool instead of allocating metrics every Collect.
	metrics := metricdata.ResourceMetrics{}
	err := c.reader.Collect(context.TODO(), &metrics)
	if err != nil {
		otel.Handle(err)
		if err == metric.ErrReaderNotRegistered {
			return
		}
	}

<<<<<<< HEAD
	c.createTargetInfoOnce.Do(func() {
		// Resource should be immutable, we don't need to compute again
		targetInfo, err := c.createInfoMetric(targetInfoMetricName, targetInfoDescription, metrics.Resource)
		if err != nil {
			// If the target info metric is invalid, disable sending it.
			otel.Handle(err)
			c.disableTargetInfo = true
		}
		c.targetInfo = targetInfo
	})
=======
	global.Debug("Prometheus exporter export", "Data", metrics)

	// Initialize (once) targetInfo and disableTargetInfo.
	func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		if c.targetInfo == nil && !c.disableTargetInfo {
			targetInfo, err := createInfoMetric(targetInfoMetricName, targetInfoDescription, metrics.Resource)
			if err != nil {
				// If the target info metric is invalid, disable sending it.
				c.disableTargetInfo = true
				otel.Handle(err)
				return
			}

			c.targetInfo = targetInfo
		}
	}()

>>>>>>> main
	if !c.disableTargetInfo {
		ch <- c.targetInfo
	}

	for _, scopeMetrics := range metrics.ScopeMetrics {
		var keys, values [2]string

		if !c.disableScopeInfo {
<<<<<<< HEAD
			scopeInfo, ok := c.scopeInfos[scopeMetrics.Scope]
			if !ok {
				scopeInfo, err = createScopeInfoMetric(scopeMetrics.Scope)
				if err != nil {
					otel.Handle(err)
				}
				c.scopeInfos[scopeMetrics.Scope] = scopeInfo
			}
			ch <- scopeInfo
=======
			scopeInfo, err := c.scopeInfo(scopeMetrics.Scope)
			if err != nil {
				otel.Handle(err)
				continue
			}

			ch <- scopeInfo

>>>>>>> main
			keys = scopeInfoKeys
			values = [2]string{scopeMetrics.Scope.Name, scopeMetrics.Scope.Version}
		}

		for _, m := range scopeMetrics.Metrics {
<<<<<<< HEAD
			switch v := m.Data.(type) {
			case metricdata.Histogram[int64]:
				addHistogramMetric(ch, v, m, keys, values, c.getName(m), c.metricFamilies)
			case metricdata.Histogram[float64]:
				addHistogramMetric(ch, v, m, keys, values, c.getName(m), c.metricFamilies)
			case metricdata.Sum[int64]:
				addSumMetric(ch, v, m, keys, values, c.getName(m), c.metricFamilies)
			case metricdata.Sum[float64]:
				addSumMetric(ch, v, m, keys, values, c.getName(m), c.metricFamilies)
			case metricdata.Gauge[int64]:
				addGaugeMetric(ch, v, m, keys, values, c.getName(m), c.metricFamilies)
			case metricdata.Gauge[float64]:
				addGaugeMetric(ch, v, m, keys, values, c.getName(m), c.metricFamilies)
=======
			typ := c.metricType(m)
			if typ == nil {
				continue
			}
			name := c.getName(m, typ)

			drop, help := c.validateMetrics(name, m.Description, typ)
			if drop {
				continue
			}

			if help != "" {
				m.Description = help
			}

			switch v := m.Data.(type) {
			case metricdata.Histogram[int64]:
				addHistogramMetric(ch, v, m, keys, values, name)
			case metricdata.Histogram[float64]:
				addHistogramMetric(ch, v, m, keys, values, name)
			case metricdata.Sum[int64]:
				addSumMetric(ch, v, m, keys, values, name)
			case metricdata.Sum[float64]:
				addSumMetric(ch, v, m, keys, values, name)
			case metricdata.Gauge[int64]:
				addGaugeMetric(ch, v, m, keys, values, name)
			case metricdata.Gauge[float64]:
				addGaugeMetric(ch, v, m, keys, values, name)
>>>>>>> main
			}
		}
	}
}

<<<<<<< HEAD
func addHistogramMetric[N int64 | float64](ch chan<- prometheus.Metric, histogram metricdata.Histogram[N], m metricdata.Metrics, ks, vs [2]string, name string, mfs map[string]*dto.MetricFamily) {
	// TODO(https://github.com/open-telemetry/opentelemetry-go/issues/3163): support exemplars
	drop, help := validateMetrics(name, m.Description, dto.MetricType_HISTOGRAM.Enum(), mfs)
	if drop {
		return
	}
	if help != "" {
		m.Description = help
	}

=======
func addHistogramMetric[N int64 | float64](ch chan<- prometheus.Metric, histogram metricdata.Histogram[N], m metricdata.Metrics, ks, vs [2]string, name string) {
	// TODO(https://github.com/open-telemetry/opentelemetry-go/issues/3163): support exemplars
>>>>>>> main
	for _, dp := range histogram.DataPoints {
		keys, values := getAttrs(dp.Attributes, ks, vs)

		desc := prometheus.NewDesc(name, m.Description, keys, nil)
		buckets := make(map[float64]uint64, len(dp.Bounds))

		cumulativeCount := uint64(0)
		for i, bound := range dp.Bounds {
			cumulativeCount += dp.BucketCounts[i]
			buckets[bound] = cumulativeCount
		}
		m, err := prometheus.NewConstHistogram(desc, dp.Count, float64(dp.Sum), buckets, values...)
		if err != nil {
			otel.Handle(err)
			continue
		}
		ch <- m
	}
}

<<<<<<< HEAD
func addSumMetric[N int64 | float64](ch chan<- prometheus.Metric, sum metricdata.Sum[N], m metricdata.Metrics, ks, vs [2]string, name string, mfs map[string]*dto.MetricFamily) {
	valueType := prometheus.CounterValue
	metricType := dto.MetricType_COUNTER
	if !sum.IsMonotonic {
		valueType = prometheus.GaugeValue
		metricType = dto.MetricType_GAUGE
	}
	if sum.IsMonotonic {
		// Add _total suffix for counters
		name += counterSuffix
	}

	drop, help := validateMetrics(name, m.Description, metricType.Enum(), mfs)
	if drop {
		return
	}
	if help != "" {
		m.Description = help
=======
func addSumMetric[N int64 | float64](ch chan<- prometheus.Metric, sum metricdata.Sum[N], m metricdata.Metrics, ks, vs [2]string, name string) {
	valueType := prometheus.CounterValue
	if !sum.IsMonotonic {
		valueType = prometheus.GaugeValue
>>>>>>> main
	}

	for _, dp := range sum.DataPoints {
		keys, values := getAttrs(dp.Attributes, ks, vs)

		desc := prometheus.NewDesc(name, m.Description, keys, nil)
		m, err := prometheus.NewConstMetric(desc, valueType, float64(dp.Value), values...)
		if err != nil {
			otel.Handle(err)
			continue
		}
		ch <- m
	}
}

<<<<<<< HEAD
func addGaugeMetric[N int64 | float64](ch chan<- prometheus.Metric, gauge metricdata.Gauge[N], m metricdata.Metrics, ks, vs [2]string, name string, mfs map[string]*dto.MetricFamily) {
	drop, help := validateMetrics(name, m.Description, dto.MetricType_GAUGE.Enum(), mfs)
	if drop {
		return
	}
	if help != "" {
		m.Description = help
	}

=======
func addGaugeMetric[N int64 | float64](ch chan<- prometheus.Metric, gauge metricdata.Gauge[N], m metricdata.Metrics, ks, vs [2]string, name string) {
>>>>>>> main
	for _, dp := range gauge.DataPoints {
		keys, values := getAttrs(dp.Attributes, ks, vs)

		desc := prometheus.NewDesc(name, m.Description, keys, nil)
		m, err := prometheus.NewConstMetric(desc, prometheus.GaugeValue, float64(dp.Value), values...)
		if err != nil {
			otel.Handle(err)
			continue
		}
		ch <- m
	}
}

// getAttrs parses the attribute.Set to two lists of matching Prometheus-style
// keys and values. It sanitizes invalid characters and handles duplicate keys
// (due to sanitization) by sorting and concatenating the values following the spec.
func getAttrs(attrs attribute.Set, ks, vs [2]string) ([]string, []string) {
	keysMap := make(map[string][]string)
	itr := attrs.Iter()
	for itr.Next() {
		kv := itr.Attribute()
		key := strings.Map(sanitizeRune, string(kv.Key))
		if _, ok := keysMap[key]; !ok {
			keysMap[key] = []string{kv.Value.Emit()}
		} else {
			// if the sanitized key is a duplicate, append to the list of keys
			keysMap[key] = append(keysMap[key], kv.Value.Emit())
		}
	}

	keys := make([]string, 0, attrs.Len())
	values := make([]string, 0, attrs.Len())
	for key, vals := range keysMap {
		keys = append(keys, key)
		sort.Slice(vals, func(i, j int) bool {
			return i < j
		})
		values = append(values, strings.Join(vals, ";"))
	}

	if ks[0] != "" {
		keys = append(keys, ks[:]...)
		values = append(values, vs[:]...)
	}
	return keys, values
}

<<<<<<< HEAD
func (c *collector) createInfoMetric(name, description string, res *resource.Resource) (prometheus.Metric, error) {
=======
func createInfoMetric(name, description string, res *resource.Resource) (prometheus.Metric, error) {
>>>>>>> main
	keys, values := getAttrs(*res.Set(), [2]string{}, [2]string{})
	desc := prometheus.NewDesc(name, description, keys, nil)
	return prometheus.NewConstMetric(desc, prometheus.GaugeValue, float64(1), values...)
}

func createScopeInfoMetric(scope instrumentation.Scope) (prometheus.Metric, error) {
	keys := scopeInfoKeys[:]
	desc := prometheus.NewDesc(scopeInfoMetricName, scopeInfoDescription, keys, nil)
	return prometheus.NewConstMetric(desc, prometheus.GaugeValue, float64(1), scope.Name, scope.Version)
}

func sanitizeRune(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ':' || r == '_' {
		return r
	}
	return '_'
}

var unitSuffixes = map[string]string{
<<<<<<< HEAD
	"1":  "_ratio",
	"By": "_bytes",
	"ms": "_milliseconds",
}

// getName returns the sanitized name, prefixed with the namespace and suffixed with unit.
func (c *collector) getName(m metricdata.Metrics) string {
	name := sanitizeName(m.Name)
	if c.namespace != "" {
		name = c.namespace + name
	}
	if c.withoutUnits {
		return name
	}
	if suffix, ok := unitSuffixes[m.Unit]; ok {
		name += suffix
	}
=======
	// Time
	"d":   "_days",
	"h":   "_hours",
	"min": "_minutes",
	"s":   "_seconds",
	"ms":  "_milliseconds",
	"us":  "_microseconds",
	"ns":  "_nanoseconds",

	// Bytes
	"By":   "_bytes",
	"KiBy": "_kibibytes",
	"MiBy": "_mebibytes",
	"GiBy": "_gibibytes",
	"TiBy": "_tibibytes",
	"KBy":  "_kilobytes",
	"MBy":  "_megabytes",
	"GBy":  "_gigabytes",
	"TBy":  "_terabytes",

	// SI
	"m": "_meters",
	"V": "_volts",
	"A": "_amperes",
	"J": "_joules",
	"W": "_watts",
	"g": "_grams",

	// Misc
	"Cel": "_celsius",
	"Hz":  "_hertz",
	"1":   "_ratio",
	"%":   "_percent",
}

// getName returns the sanitized name, prefixed with the namespace and suffixed with unit.
func (c *collector) getName(m metricdata.Metrics, typ *dto.MetricType) string {
	name := sanitizeName(m.Name)
	addCounterSuffix := !c.withoutCounterSuffixes && *typ == dto.MetricType_COUNTER
	if addCounterSuffix {
		// Remove the _total suffix here, as we will re-add the total suffix
		// later, and it needs to come after the unit suffix.
		name = strings.TrimSuffix(name, counterSuffix)
	}
	if c.namespace != "" {
		name = c.namespace + name
	}
	if suffix, ok := unitSuffixes[m.Unit]; ok && !c.withoutUnits && !strings.HasSuffix(name, suffix) {
		name += suffix
	}
	if addCounterSuffix {
		name += counterSuffix
	}
>>>>>>> main
	return name
}

func sanitizeName(n string) string {
	// This algorithm is based on strings.Map from Go 1.19.
	const replacement = '_'

	valid := func(i int, r rune) bool {
		// Taken from
		// https://github.com/prometheus/common/blob/dfbc25bd00225c70aca0d94c3c4bb7744f28ace0/model/metric.go#L92-L102
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' || r == ':' || (r >= '0' && r <= '9' && i > 0) {
			return true
		}
		return false
	}

	// This output buffer b is initialized on demand, the first time a
	// character needs to be replaced.
	var b strings.Builder
	for i, c := range n {
		if valid(i, c) {
			continue
		}

		if i == 0 && c >= '0' && c <= '9' {
			// Prefix leading number with replacement character.
			b.Grow(len(n) + 1)
			_ = b.WriteByte(byte(replacement))
			break
		}
		b.Grow(len(n))
		_, _ = b.WriteString(n[:i])
		_ = b.WriteByte(byte(replacement))
		width := utf8.RuneLen(c)
		n = n[i+width:]
		break
	}

	// Fast path for unchanged input.
	if b.Cap() == 0 { // b.Grow was not called above.
		return n
	}

	for _, c := range n {
		// Due to inlining, it is more performant to invoke WriteByte rather then
		// WriteRune.
		if valid(1, c) { // We are guaranteed to not be at the start.
			_ = b.WriteByte(byte(c))
		} else {
			_ = b.WriteByte(byte(replacement))
		}
	}

	return b.String()
}

<<<<<<< HEAD
func validateMetrics(name, description string, metricType *dto.MetricType, mfs map[string]*dto.MetricFamily) (drop bool, help string) {
	emf, exist := mfs[name]
	if !exist {
		mfs[name] = &dto.MetricFamily{
=======
func (c *collector) metricType(m metricdata.Metrics) *dto.MetricType {
	switch v := m.Data.(type) {
	case metricdata.Histogram[int64], metricdata.Histogram[float64]:
		return dto.MetricType_HISTOGRAM.Enum()
	case metricdata.Sum[float64]:
		if v.IsMonotonic {
			return dto.MetricType_COUNTER.Enum()
		}
		return dto.MetricType_GAUGE.Enum()
	case metricdata.Sum[int64]:
		if v.IsMonotonic {
			return dto.MetricType_COUNTER.Enum()
		}
		return dto.MetricType_GAUGE.Enum()
	case metricdata.Gauge[int64], metricdata.Gauge[float64]:
		return dto.MetricType_GAUGE.Enum()
	}
	return nil
}

func (c *collector) scopeInfo(scope instrumentation.Scope) (prometheus.Metric, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	scopeInfo, ok := c.scopeInfos[scope]
	if ok {
		return scopeInfo, nil
	}

	scopeInfo, err := createScopeInfoMetric(scope)
	if err != nil {
		return nil, fmt.Errorf("cannot create scope info metric: %w", err)
	}

	c.scopeInfos[scope] = scopeInfo

	return scopeInfo, nil
}

func (c *collector) validateMetrics(name, description string, metricType *dto.MetricType) (drop bool, help string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	emf, exist := c.metricFamilies[name]

	if !exist {
		c.metricFamilies[name] = &dto.MetricFamily{
>>>>>>> main
			Name: proto.String(name),
			Help: proto.String(description),
			Type: metricType,
		}
		return false, ""
	}
<<<<<<< HEAD
=======

>>>>>>> main
	if emf.GetType() != *metricType {
		global.Error(
			errors.New("instrument type conflict"),
			"Using existing type definition.",
			"instrument", name,
			"existing", emf.GetType(),
			"dropped", *metricType,
		)
		return true, ""
	}
	if emf.GetHelp() != description {
		global.Info(
			"Instrument description conflict, using existing",
			"instrument", name,
			"existing", emf.GetHelp(),
			"dropped", description,
		)
		return false, emf.GetHelp()
	}

	return false, ""
}
