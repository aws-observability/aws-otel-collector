// Copyright 2021 Dynatrace LLC
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

package serialize

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dynatrace-oss/dynatrace-metric-utils-go/metric/dimensions"
	"github.com/dynatrace-oss/dynatrace-metric-utils-go/normalize"
)

func joinPrefix(metricKey, prefix string) string {
	if prefix != "" {
		return fmt.Sprintf("%s.%s", prefix, metricKey)
	}
	return metricKey
}

// MetricKey joins and normalizes a metric key. Skips the prefix if empty.
func MetricKey(metricKey, prefix string) (string, error) {
	return normalize.MetricKey(joinPrefix(metricKey, prefix))
}

func formatDimensions(dims []dimensions.Dimension) string {
	var sb strings.Builder
	firstIter := true

	for _, dim := range dims {
		if firstIter {
			firstIter = false
		} else {
			sb.WriteString(",")
		}

		sb.WriteString(fmt.Sprintf("%s=%s", dim.Key, dim.Value))
	}

	return sb.String()
}

// Dimensions combines the individual dimensions into one string, separated by a comma.
func Dimensions(dims dimensions.NormalizedDimensionList) string {
	return dims.Format(formatDimensions)
}

// IntSummaryValue returns the value part of an metrics ingestion line for the given integers
func IntSummaryValue(min, max, sum, count int64) string {
	return fmt.Sprintf("gauge,min=%d,max=%d,sum=%d,count=%d", min, max, sum, count)
}

// IntCountValue transforms the integer given integer into a valid ingestion line value part.
func IntCountValue(value int64) string {
	return fmt.Sprintf("count,delta=%d", value)
}

// FloatSummaryValue returns the value part of an metrics ingestion line for the given floats, and an integer count
func FloatSummaryValue(min, max, sum float64, count int64) string {
	return fmt.Sprintf("gauge,min=%s,max=%s,sum=%s,count=%d", SerializeFloat64(min), SerializeFloat64(max), SerializeFloat64(sum), count)
}

// FloatCountValue transforms the float given integer into a valid ingestion line value part.
func FloatCountValue(value float64) string {
	return fmt.Sprintf("count,delta=%s", SerializeFloat64(value))
}

// IntGaugeValue transforms the given value to a gauge value that can be sent to the ingestion endpoint.
func IntGaugeValue(value int64) string {
	return fmt.Sprintf("gauge,%d", value)
}

// FloatGaugeValue transforms the given value to a gauge value that can be sent to the ingestion endpoint.
func FloatGaugeValue(value float64) string {
	return fmt.Sprintf("gauge,%s", SerializeFloat64(value))
}

func SerializeFloat64(n float64) string {
	formatted := strconv.FormatFloat(n, 'g', -1, 64)
	if strings.Contains(formatted, "e") && !strings.Contains(formatted, ".") {
		// e. g. 1e+10, which is not valid as per the API spec which requires a decimal point for scientific notation.
		eIndex := strings.Index(formatted, "e")
		formatted = formatted[:eIndex] + ".0" + formatted[eIndex:]
	}
	return formatted
}

// Timestamp retruns the current timestamp as Unix time or an empty string if time has not been set.
func Timestamp(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	milliseconds := t.UnixNano() / 1_000_000
	return strconv.FormatInt(milliseconds, 10)
}
