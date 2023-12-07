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

import "github.com/dynatrace-oss/dynatrace-metric-utils-go/serialize"

// everything that has a serialize method can be used as metric value
type metricValue interface {
	serialize() string
}

type intCounterValue struct {
	value int64
}

func (i intCounterValue) serialize() string {
	return serialize.IntCountValue(i.value)
}

type floatCounterValue struct {
	value float64
}

func (f floatCounterValue) serialize() string {
	return serialize.FloatCountValue(f.value)
}

type intSummaryValue struct {
	min, max, sum, count int64
}

func (i intSummaryValue) serialize() string {
	return serialize.IntSummaryValue(i.min, i.max, i.sum, i.count)
}

type floatSummaryValue struct {
	min, max, sum float64
	count         int64
}

func (f floatSummaryValue) serialize() string {
	return serialize.FloatSummaryValue(f.min, f.max, f.sum, f.count)
}

type intGaugeValue struct {
	value int64
}

func (i intGaugeValue) serialize() string {
	return serialize.IntGaugeValue(i.value)
}

type floatGaugeValue struct {
	value float64
}

func (f floatGaugeValue) serialize() string {
	return serialize.FloatGaugeValue(f.value)
}
