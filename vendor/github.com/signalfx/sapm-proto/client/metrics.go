// Copyright 2019 Splunk, Inc.
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

package client

import (
	"context"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
)

var mFailedEncodings = stats.Int64("sapm/encodings_failed", "Number of times encoding/compressing a request failed", stats.UnitDimensionless)
var mRequestsCompleted = stats.Int64("sapm/requests_completed", "Number of HTTP requests made successfully", stats.UnitDimensionless)
var mRequestsFailed = stats.Int64("sapm/requests_failed", "Number of failed HTTP requests", stats.UnitDimensionless)
var mBatchesDropped = stats.Int64("sapm/batches_dropped", "Number of batches dropped", stats.UnitDimensionless)
var mSpansDropped = stats.Int64("sapm/spans_dropped", "Number of spans dropped", stats.UnitDimensionless)
var mBatchesExported = stats.Int64("sapm/batches_exported", "Number of batches successfully exporter", stats.UnitDimensionless)
var mSpansExported = stats.Int64("sapm/spans_exported", "Number of spans successfully exporter", stats.UnitDimensionless)

func newMetricsView(m stats.Measure) *view.View {
	return &view.View{
		Name:        m.Name(),
		Measure:     m,
		Description: m.Description(),
		Aggregation: view.Sum(),
	}
}

func metricViews() []*view.View {
	return []*view.View{
		newMetricsView(mFailedEncodings),
		newMetricsView(mRequestsCompleted),
		newMetricsView(mRequestsFailed),
		newMetricsView(mBatchesDropped),
		newMetricsView(mBatchesExported),
		newMetricsView(mSpansDropped),
		newMetricsView(mSpansExported),
	}
}

func recordEncodingFailure(ctx context.Context, _ *sendRequest) {
	stats.Record(ctx, mFailedEncodings.M(1))
}

func recordSendFailure(ctx context.Context, _ *sendRequest) {
	stats.Record(ctx, mRequestsFailed.M(1))
}

func recordDrops(ctx context.Context, sr *sendRequest) {
	stats.Record(ctx, mBatchesDropped.M(sr.batches), mSpansDropped.M(sr.spans))
}

func recordSuccess(ctx context.Context, sr *sendRequest) {
	stats.Record(ctx, mBatchesExported.M(sr.batches), mSpansExported.M(sr.spans))
}
