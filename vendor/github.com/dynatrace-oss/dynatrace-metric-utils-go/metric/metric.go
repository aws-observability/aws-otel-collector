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

package metric

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
	"sync/atomic"
	"time"

	"github.com/dynatrace-oss/dynatrace-metric-utils-go/metric/dimensions"
	"github.com/dynatrace-oss/dynatrace-metric-utils-go/serialize"
)

const (
	timestampWarningThrottleFactor = 1000
	// The maximum number of characters per serialized line accepted by the ingest API.
	// Lines exceeding this threshold should be dropped.
	metricLineMaxLength = 50_000
)

var timestampWarningCounter uint32 = 0

// Metric contains all information needed to create a string representation of the accumulated metric data.
type Metric struct {
	metricKey  string
	prefix     string
	value      metricValue
	dimensions dimensions.NormalizedDimensionList
	timestamp  time.Time
}

// MetricOption represents the function interface used to set options on the metric object.
type MetricOption func(m *Metric) error

func joinStrings(key, dim, value, timestamp string) (string, error) {
	var sb strings.Builder

	sb.WriteString(key)
	if dim != "" {
		sb.WriteString(",")
		sb.WriteString(dim)
	}

	sb.WriteString(" ")
	sb.WriteString(value)

	if timestamp != "" {
		sb.WriteString(" ")
		sb.WriteString(timestamp)
	}

	return sb.String(), nil

}

func (m Metric) ensureRequiredFieldsSet() error {
	if m.metricKey == "" && m.prefix == "" {
		return errors.New("metric key and prefix empty, cannot create metric name")
	}

	if m.value == nil {
		return errors.New("metric value not set, cannot create metric")
	}

	return nil
}

func ensureFloatsAreValid(values ...float64) error {
	for _, value := range values {
		if math.IsNaN(value) {
			return errors.New("value is NaN.")
		}

		if math.IsInf(value, 0) {
			return errors.New("value is infinite")
		}
	}

	return nil
}

// Serialize creates the string representation of the Metric object.
// Will return an error if the serialization fails, or if the line length after serialization exceeds the maximum line length accepted by the ingest API.
func (m Metric) Serialize() (string, error) {
	keyString, err := serialize.MetricKey(m.metricKey, m.prefix)
	if err != nil {
		return "", err
	}
	if m.value == nil {
		return "", errors.New("cannot serialize nil value")
	}

	dimString := serialize.Dimensions(m.dimensions)
	valueString := m.value.serialize()
	timeString := serialize.Timestamp(m.timestamp)

	metricLine, err := joinStrings(keyString, dimString, valueString, timeString)
	if err != nil {
		return "", err
	}

	if len(metricLine) > metricLineMaxLength {
		return "", fmt.Errorf("serialized line exceeds limit of %d characters accepted by the ingest API. Metric name: '%s'", metricLineMaxLength, keyString)
	}

	return metricLine, nil
}

// NewMetric creates a new metric with a mandatory name and options. At least one value option must be set.
func NewMetric(name string, options ...MetricOption) (*Metric, error) {
	m := &Metric{
		metricKey: name,
	}

	for _, option := range options {
		err := option(m)
		if err != nil {
			return nil, err
		}
	}

	err := m.ensureRequiredFieldsSet()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func checkValueAlreadySet(m *Metric) error {
	if m.value != nil {
		return errors.New("cannot set two values on one metric.")
	}
	return nil
}

// WithPrefix sets the prefix for Metric creation.
func WithPrefix(prefix string) MetricOption {
	return func(m *Metric) error {
		m.prefix = prefix
		return nil
	}
}

// WithDimensions sets the passed dimension list as the dimensions to export.
// Pass only dimension lists that have been de-duplicated (by dimensions.MergeLists)
func WithDimensions(dims dimensions.NormalizedDimensionList) MetricOption {
	return func(m *Metric) error {
		m.dimensions = dims

		return nil
	}
}

func trySetValue(m *Metric, val metricValue) error {
	if err := checkValueAlreadySet(m); err != nil {
		return err
	}
	m.value = val
	return nil
}

// WithIntCounterValueDelta sets a value on the metric that will be formatted as "count,delta=<value>"
// Returns an error if a value is already set.
func WithIntCounterValueDelta(val int64) MetricOption {
	return func(m *Metric) error {
		return trySetValue(m, intCounterValue{value: val})
	}
}

// WithFloatCounterValueDelta sets a value on the metric that will be formatted as "count,delta=<value>"
// Returns an error if a value is already set, or if the value is NaN or Infinity.
func WithFloatCounterValueDelta(val float64) MetricOption {
	return func(m *Metric) error {
		if err := ensureFloatsAreValid(val); err != nil {
			return err
		}

		return trySetValue(m, floatCounterValue{value: val})
	}
}

// WithIntSummaryValue sets a summary statistic on the metric that will be formatted as
// "gauge,min=<min>,max=<max>,sum=<sum>,count=<count>".
// Returns an error if the count is below 0, if a value is already set, or if min is greater than max.
func WithIntSummaryValue(min, max, sum, count int64) MetricOption {
	return func(m *Metric) error {
		if count < 0 {
			return fmt.Errorf("count cannot be smaller than 0, was %v", count)
		}
		if min > max {
			return fmt.Errorf("min (%d) cannot be greater than max (%d)", min, max)
		}

		return trySetValue(m, intSummaryValue{min: min, max: max, sum: sum, count: count})
	}
}

//  WithFloatSummaryValue sets a summary statistic on the metric that will be formatted as
// "gauge,min=<min>,max=<max>,sum=<sum>,count=<count>".
// Returns an error if the count is below 0, if min is greater than max,
// if a value is already set, or if any of min, max, or sum are NaN or infinite.
func WithFloatSummaryValue(min, max, sum float64, count int64) MetricOption {
	return func(m *Metric) error {
		if count < 0 {
			return fmt.Errorf("count cannot be smaller than 0, was %v", count)
		}
		if min > max {
			return fmt.Errorf("min (%.3f) cannot be greater than max (%.3f)", min, max)
		}

		if err := ensureFloatsAreValid(min, max, sum); err != nil {
			return err
		}

		return trySetValue(m, floatSummaryValue{min: min, max: max, sum: sum, count: count})
	}
}

// WithIntGaugeValue sets a gauge value on the metric that will be formatted as "gauge,<value>"
// Returns an error if a value is already set.
func WithIntGaugeValue(val int64) MetricOption {
	return func(m *Metric) error {
		return trySetValue(m, intGaugeValue{value: val})
	}
}

// WithFloatGaugeValue sets a gauge value on the metric that will be formatted as "gauge,<value>"
// Retruns an error if the value is already set or if the value is NaN or Infinity.
func WithFloatGaugeValue(val float64) MetricOption {
	return func(m *Metric) error {

		if err := ensureFloatsAreValid(val); err != nil {
			return err
		}

		return trySetValue(m, floatGaugeValue{value: val})
	}
}

// WithTimestamp sets a specific timestamp for the metric.
// If the timestamp represents a time before 2000 or after 3000, no timestamp is set and the
// server timestamp will be used upon ingestion. This might be case if seconds or microseconds
// are set as milliseconds upon timestamp creation.
func WithTimestamp(t time.Time) MetricOption {
	return func(m *Metric) error {
		if t.Year() < 2000 || t.Year() > 3000 {
			currentTimestampWarningCounter := atomic.AddUint32(&timestampWarningCounter, 1)
			if currentTimestampWarningCounter == 1 {
				log.Printf("Order of magnitude of the timestamp seems off (%s). "+
					"The timestamp represents a time before the year 2000 or after the year 3000. "+
					"Skipping setting timestamp, the current server time will be added upon ingestion. "+
					"Only one out of every %d of these messages will be printed.", t.String(), timestampWarningThrottleFactor)
			}
			if currentTimestampWarningCounter == timestampWarningThrottleFactor {
				atomic.StoreUint32(&timestampWarningCounter, 0)
			}

			m.timestamp = time.Time{}
			return nil
		}

		m.timestamp = t
		return nil
	}
}

// WithCurrentTime sets the current time as timestamp for the metric.
func WithCurrentTime() MetricOption {
	return WithTimestamp(time.Now())
}
