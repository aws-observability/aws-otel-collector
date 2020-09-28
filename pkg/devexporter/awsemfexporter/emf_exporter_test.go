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

package awsemfexporter

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"
	commonpb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	resourcepb "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumerdata"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/translator/internaldata"
	"go.uber.org/zap"
)

type mockPusher struct {
	mock.Mock
}

func (p *mockPusher) AddLogEntry(logEvent *LogEvent) error {
	args := p.Called(nil)
	errorStr := args.String(0)
	if errorStr != "" {
		return awserr.NewRequestFailure(nil, 400, "").(error)
	}
	return nil
}

func (p *mockPusher) ForceFlush() error {
	args := p.Called(nil)
	errorStr := args.String(0)
	if errorStr != "" {
		return awserr.NewRequestFailure(nil, 400, "").(error)
	}
	return nil
}

func TestConsumeMetrics(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	factory := NewFactory()
	expCfg := factory.CreateDefaultConfig().(*Config)
	expCfg.Region = "us-west-2"
	expCfg.MaxRetries = 0
	exp, err := New(expCfg, component.ExporterCreateParams{Logger: zap.NewNop()})
	assert.Nil(t, err)
	assert.NotNil(t, exp)

	mdata := consumerdata.MetricsData{
		Node: &commonpb.Node{
			ServiceInfo: &commonpb.ServiceInfo{Name: "test-emf"},
			LibraryInfo: &commonpb.LibraryInfo{ExporterVersion: "SomeVersion"},
		},
		Resource: &resourcepb.Resource{
			Labels: map[string]string{
				"resource": "R1",
			},
		},
		Metrics: []*metricspb.Metric{
			{
				MetricDescriptor: &metricspb.MetricDescriptor{
					Name:        "spanCounter",
					Description: "Counting all the spans",
					Unit:        "Count",
					Type:        metricspb.MetricDescriptor_CUMULATIVE_INT64,
					LabelKeys: []*metricspb.LabelKey{
						{Key: "spanName"},
						{Key: "isItAnError"},
					},
				},
				Timeseries: []*metricspb.TimeSeries{
					{
						LabelValues: []*metricspb.LabelValue{
							{Value: "testSpan"},
							{Value: "false"},
						},
						Points: []*metricspb.Point{
							{
								Timestamp: &timestamp.Timestamp{
									Seconds: 100,
								},
								Value: &metricspb.Point_Int64Value{
									Int64Value: 1,
								},
							},
						},
					},
				},
			},
		},
	}
	md := internaldata.OCToMetrics(mdata)
	require.Error(t, exp.ConsumeMetrics(ctx, md))
	require.NoError(t, exp.Shutdown(ctx))
}

func TestConsumeMetricsWithLogGroupStreamConfig(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	factory := NewFactory()
	expCfg := factory.CreateDefaultConfig().(*Config)
	expCfg.Region = "us-west-2"
	expCfg.MaxRetries = defaultRetryCount
	expCfg.LogGroupName = "test-logGroupName"
	expCfg.LogStreamName = "test-logStreamName"
	exp, err := New(expCfg, component.ExporterCreateParams{Logger: zap.NewNop()})
	assert.Nil(t, err)
	assert.NotNil(t, exp)

	mdata := consumerdata.MetricsData{}
	md := internaldata.OCToMetrics(mdata)
	require.NoError(t, exp.Start(ctx, nil))
	require.NoError(t, exp.ConsumeMetrics(ctx, md))
	require.NoError(t, exp.Shutdown(ctx))
	streamToPusherMap, ok := exp.(*emfExporter).groupStreamToPusherMap["test-logGroupName"]
	assert.True(t, ok)
	pusher, ok := streamToPusherMap["test-logStreamName"]
	assert.True(t, ok)
	assert.NotNil(t, pusher)
}

func TestPushMetricsDataWithErr(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	factory := NewFactory()
	expCfg := factory.CreateDefaultConfig().(*Config)
	expCfg.Region = "us-west-2"
	expCfg.MaxRetries = 0
	expCfg.LogGroupName = "test-logGroupName"
	expCfg.LogStreamName = "test-logStreamName"
	exp, err := New(expCfg, component.ExporterCreateParams{Logger: zap.NewNop()})
	assert.Nil(t, err)
	assert.NotNil(t, exp)

	pusher := new(mockPusher)
	pusher.On("AddLogEntry", nil).Return("some error").Once()
	pusher.On("AddLogEntry", nil).Return("").Twice()
	pusher.On("ForceFlush", nil).Return("some error").Once()
	pusher.On("ForceFlush", nil).Return("").Once()
	pusher.On("ForceFlush", nil).Return("some error").Once()
	streamToPusherMap := map[string]Pusher{"test-logStreamName": pusher}
	exp.(*emfExporter).groupStreamToPusherMap = map[string]map[string]Pusher{}
	exp.(*emfExporter).groupStreamToPusherMap["test-logGroupName"] = streamToPusherMap

	mdata := consumerdata.MetricsData{
		Node: &commonpb.Node{
			ServiceInfo: &commonpb.ServiceInfo{Name: "test-emf"},
			LibraryInfo: &commonpb.LibraryInfo{ExporterVersion: "SomeVersion"},
		},
		Resource: &resourcepb.Resource{
			Labels: map[string]string{
				"resource": "R1",
			},
		},
		Metrics: []*metricspb.Metric{
			{
				MetricDescriptor: &metricspb.MetricDescriptor{
					Name:        "spanCounter",
					Description: "Counting all the spans",
					Unit:        "Count",
					Type:        metricspb.MetricDescriptor_CUMULATIVE_INT64,
					LabelKeys: []*metricspb.LabelKey{
						{Key: "spanName"},
						{Key: "isItAnError"},
					},
				},
				Timeseries: []*metricspb.TimeSeries{
					{
						LabelValues: []*metricspb.LabelValue{
							{Value: "testSpan"},
							{Value: "false"},
						},
						Points: []*metricspb.Point{
							{
								Timestamp: &timestamp.Timestamp{
									Seconds: 100,
								},
								Value: &metricspb.Point_Int64Value{
									Int64Value: 1,
								},
							},
						},
					},
				},
			},
		},
	}
	md := internaldata.OCToMetrics(mdata)
	_, err = exp.(*emfExporter).pushMetricsData(ctx, md)
	assert.NotNil(t, err)
	_, err = exp.(*emfExporter).pushMetricsData(ctx, md)
	assert.NotNil(t, err)
	_, err = exp.(*emfExporter).pushMetricsData(ctx, md)
	assert.Nil(t, err)
	err = exp.(*emfExporter).Shutdown(ctx)
	assert.Nil(t, err)
}

func TestNewExporterWithoutConfig(t *testing.T) {
	factory := NewFactory()
	expCfg := factory.CreateDefaultConfig().(*Config)
	env := stashEnv()
	defer popEnv(env)
	os.Setenv("AWS_STS_REGIONAL_ENDPOINTS", "fake")

	exp, err := New(expCfg, component.ExporterCreateParams{Logger: zap.NewNop()})
	assert.NotNil(t, err)
	assert.Nil(t, exp)
}

func TestNewExporterWithoutSession(t *testing.T) {
	exp, err := New(nil, component.ExporterCreateParams{Logger: zap.NewNop()})
	assert.NotNil(t, err)
	assert.Nil(t, exp)
}

func TestWrapErrorIfBadRequest(t *testing.T) {
	awsErr := awserr.NewRequestFailure(nil, 400, "").(error)
	err := wrapErrorIfBadRequest(&awsErr)
	assert.True(t, consumererror.IsPermanent(err))
	awsErr = awserr.NewRequestFailure(nil, 500, "").(error)
	err = wrapErrorIfBadRequest(&awsErr)
	assert.False(t, consumererror.IsPermanent(err))
}
