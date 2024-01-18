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
	"strings"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

const (
	// divMebibytes specifies the number of bytes in a mebibyte.
	divMebibytes = 1024 * 1024
	// divPercentage specifies the division necessary for converting fractions to percentages.
	divPercentage = 0.01
)

// remapMetrics extracts any Datadog specific metrics from m and appends them to all.
func remapMetrics(all pmetric.MetricSlice, m pmetric.Metric) {
	remapSystemMetrics(all, m)
	remapContainerMetrics(all, m)
}

// remapSystemMetrics extracts system metrics from m and appends them to all.
func remapSystemMetrics(all pmetric.MetricSlice, m pmetric.Metric) {
	name := m.Name()
	if !strings.HasPrefix(name, "process.") && !strings.HasPrefix(name, "system.") {
		// not a system metric
		return
	}
	switch name {
	case "system.cpu.load_average.1m":
		copyMetric(all, m, "system.load.1", 1)
	case "system.cpu.load_average.5m":
		copyMetric(all, m, "system.load.5", 1)
	case "system.cpu.load_average.15m":
		copyMetric(all, m, "system.load.15", 1)
	case "system.cpu.utilization":
		copyMetric(all, m, "system.cpu.idle", divPercentage, kv{"state", "idle"})
		copyMetric(all, m, "system.cpu.user", divPercentage, kv{"state", "user"})
		copyMetric(all, m, "system.cpu.system", divPercentage, kv{"state", "system"})
		copyMetric(all, m, "system.cpu.iowait", divPercentage, kv{"state", "wait"})
		copyMetric(all, m, "system.cpu.stolen", divPercentage, kv{"state", "steal"})
	case "system.memory.usage":
		copyMetric(all, m, "system.mem.total", divMebibytes)
		copyMetric(all, m, "system.mem.usable", divMebibytes,
			kv{"state", "free"},
			kv{"state", "cached"},
			kv{"state", "buffered"},
		)
	case "system.network.io":
		copyMetric(all, m, "system.net.bytes_rcvd", 1, kv{"direction", "receive"})
		copyMetric(all, m, "system.net.bytes_sent", 1, kv{"direction", "transmit"})
	case "system.paging.usage":
		copyMetric(all, m, "system.swap.free", divMebibytes, kv{"state", "free"})
		copyMetric(all, m, "system.swap.used", divMebibytes, kv{"state", "used"})
	case "system.filesystem.utilization":
		copyMetric(all, m, "system.disk.in_use", 1)
	}
	// process.* and system.* metrics need to be prepended with the otel.* namespace
	m.SetName("otel." + m.Name())
}

// remapContainerMetrics extracts system metrics from m and appends them to all.
func remapContainerMetrics(all pmetric.MetricSlice, m pmetric.Metric) {
	name := m.Name()
	if !strings.HasPrefix(name, "container.") {
		// not a container metric
		return
	}
	switch name {
	case "container.cpu.usage.total":
		if addm, ok := copyMetric(all, m, "container.cpu.usage", 1); ok {
			addm.SetUnit("nanocore")
		}
	case "container.cpu.usage.usermode":
		if addm, ok := copyMetric(all, m, "container.cpu.user", 1); ok {
			addm.SetUnit("nanocore")
		}
	case "container.cpu.usage.system":
		if addm, ok := copyMetric(all, m, "container.cpu.system", 1); ok {
			addm.SetUnit("nanocore")
		}
	case "container.cpu.throttling_data.throttled_time":
		copyMetric(all, m, "container.cpu.throttled", 1)
	case "container.cpu.throttling_data.throttled_periods":
		copyMetric(all, m, "container.cpu.throttled.periods", 1)
	case "container.memory.usage.total":
		copyMetric(all, m, "container.memory.usage", 1)
	case "container.memory.active_anon":
		copyMetric(all, m, "container.memory.kernel", 1)
	case "container.memory.hierarchical_memory_limit":
		copyMetric(all, m, "container.memory.limit", 1)
	case "container.memory.usage.limit":
		copyMetric(all, m, "container.memory.soft_limit", 1)
	case "container.memory.total_cache":
		copyMetric(all, m, "container.memory.cache", 1)
	case "container.memory.total_swap":
		copyMetric(all, m, "container.memory.swap", 1)
	case "container.blockio.io_service_bytes_recursive":
		copyMetric(all, m, "container.io.write", 1, kv{"operation", "write"})
		copyMetric(all, m, "container.io.read", 1, kv{"operation", "read"})
	case "container.blockio.io_serviced_recursive":
		copyMetric(all, m, "container.io.write.operations", 1, kv{"operation", "write"})
		copyMetric(all, m, "container.io.read.operations", 1, kv{"operation", "read"})
	case "container.network.io.usage.tx_bytes":
		copyMetric(all, m, "container.net.sent", 1)
	case "container.network.io.usage.tx_packets":
		copyMetric(all, m, "container.net.sent.packets", 1)
	case "container.network.io.usage.rx_bytes":
		copyMetric(all, m, "container.net.rcvd", 1)
	case "container.network.io.usage.rx_packets":
		copyMetric(all, m, "container.net.rcvd.packets", 1)
	}
}

// kv represents a key/value pair.
type kv struct{ K, V string }

// copyMetric copies metric m to dest. The new metric's name will be newname, and all of its datapoints will
// be divided by div. If filter is provided, only the data points that have *either* of the specified string
// attributes will be copied over. If the filtering results in no datapoints, no new metric is added to dest.
//
// copyMetric returns the new metric and reports whether it was added to dest.
//
// Please note that copyMetric is restricted to the metric types Sum and Gauge.
func copyMetric(dest pmetric.MetricSlice, m pmetric.Metric, newname string, div float64, filter ...kv) (pmetric.Metric, bool) {
	newm := pmetric.NewMetric()
	m.CopyTo(newm)
	newm.SetName(newname)
	var dps pmetric.NumberDataPointSlice
	switch newm.Type() {
	case pmetric.MetricTypeGauge:
		dps = newm.Gauge().DataPoints()
	case pmetric.MetricTypeSum:
		dps = newm.Sum().DataPoints()
	default:
		// invalid metric type
		return newm, false
	}
	dps.RemoveIf(func(dp pmetric.NumberDataPoint) bool {
		if !hasAny(dp, filter...) {
			return true
		}
		switch dp.ValueType() {
		case pmetric.NumberDataPointValueTypeInt:
			if div >= 1 {
				// avoid division by zero
				dp.SetIntValue(dp.IntValue() / int64(div))
			}
		case pmetric.NumberDataPointValueTypeDouble:
			if div != 0 {
				dp.SetDoubleValue(dp.DoubleValue() / div)
			}
		}
		return false
	})
	if dps.Len() > 0 {
		// if we have datapoints, copy it
		addm := dest.AppendEmpty()
		newm.CopyTo(addm)
		return addm, true
	}
	return newm, false
}

// hasAny reports whether point has any of the given string tags.
// If no tags are provided it returns true.
func hasAny(point pmetric.NumberDataPoint, tags ...kv) bool {
	if len(tags) == 0 {
		return true
	}
	attr := point.Attributes()
	for _, tag := range tags {
		v, ok := attr.Get(tag.K)
		if !ok {
			continue
		}
		if v.Str() == tag.V {
			return true
		}
	}
	return false
}
