/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package config

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor/processortest"
)

const (
	accountID  = "111122223333"
	instanceID = "i-0123456789abcdef0"
	clusterARN = "arn:aws:ecs:us-west-2:111122223333:cluster/demo"
	lambdaARN  = "arn:aws:lambda:us-west-2:111122223333:function:mine"
)

func TestCloudResourceID(t *testing.T) {
	t.Setenv("AWS_REGION", "us-west-2")

	ec2 := func(region string) map[string]string {
		return map[string]string{
			"cloud.provider":   "aws",
			"cloud.platform":   "aws_ec2",
			"cloud.region":     region,
			"cloud.account.id": accountID,
			"host.id":          instanceID,
		}
	}
	withAttr := func(attrs map[string]string, key, value string) map[string]string {
		attrs[key] = value
		return attrs
	}

	tests := []struct {
		name     string
		config   string
		resource map[string]string
		want     string // empty means cloud.resource_id stays unset
	}{
		{
			name:     "EC2 instance",
			config:   "cloudwatch-otlp.yaml",
			resource: ec2("us-west-2"),
			want:     "arn:aws:ec2:us-west-2:" + accountID + ":instance/" + instanceID,
		},
		{
			name:     "EC2 instance in AWS GovCloud (US)",
			config:   "cloudwatch-otlp.yaml",
			resource: ec2("us-gov-west-1"),
			want:     "arn:aws-us-gov:ec2:us-gov-west-1:" + accountID + ":instance/" + instanceID,
		},
		{
			name:     "EC2 instance in China",
			config:   "cloudwatch-otlp.yaml",
			resource: ec2("cn-north-1"),
			want:     "arn:aws-cn:ec2:cn-north-1:" + accountID + ":instance/" + instanceID,
		},
		{
			name:     "not on EC2",
			config:   "cloudwatch-otlp.yaml",
			resource: map[string]string{"cloud.region": "us-west-2"},
		},
		{
			name:     "EC2 instance without an account ID",
			config:   "cloudwatch-otlp.yaml",
			resource: map[string]string{"cloud.platform": "aws_ec2", "cloud.region": "us-west-2", "host.id": instanceID},
		},
		{
			name:     "application-provided value on EC2",
			config:   "cloudwatch-otlp.yaml",
			resource: withAttr(ec2("us-west-2"), "cloud.resource_id", lambdaARN),
			want:     lambdaARN,
		},
		{
			name:     "ECS cluster",
			config:   filepath.Join("ecs", "ecs-cloudwatch-otlp.yaml"),
			resource: map[string]string{"cloud.platform": "aws_ecs", "aws.ecs.cluster.arn": clusterARN},
			want:     clusterARN,
		},
		{
			name:     "application-provided value on ECS",
			config:   filepath.Join("ecs", "ecs-cloudwatch-otlp.yaml"),
			resource: map[string]string{"aws.ecs.cluster.arn": clusterARN, "cloud.resource_id": lambdaARN},
			want:     lambdaARN,
		},
		{
			name:     "ECS without a cluster ARN",
			config:   filepath.Join("ecs", "ecs-cloudwatch-otlp.yaml"),
			resource: map[string]string{"cloud.platform": "aws_ecs"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := loadBundledConfig(t, tt.config)
			procCfg := cfg.Processors[component.MustNewIDWithName("transform", "cloud_resource_id")]
			require.NotNil(t, procCfg)

			assertResourceID(t, tt.want, runTraces(t, procCfg, tt.resource))
			assertResourceID(t, tt.want, runMetrics(t, procCfg, tt.resource))
		})
	}
}

func runTraces(t *testing.T, cfg component.Config, attrs map[string]string) pcommon.Map {
	t.Helper()
	ctx := context.Background()
	factory := transformprocessor.NewFactory()
	sink := new(consumertest.TracesSink)
	proc, err := factory.CreateTraces(ctx, processortest.NewNopSettings(factory.Type()), cfg, sink)
	require.NoError(t, err)
	require.NoError(t, proc.Start(ctx, componenttest.NewNopHost()))
	defer func() { require.NoError(t, proc.Shutdown(ctx)) }()

	td := ptrace.NewTraces()
	rs := td.ResourceSpans().AppendEmpty()
	putAll(rs.Resource().Attributes(), attrs)
	rs.ScopeSpans().AppendEmpty().Spans().AppendEmpty().SetName("span")
	require.NoError(t, proc.ConsumeTraces(ctx, td))

	require.Len(t, sink.AllTraces(), 1)
	return sink.AllTraces()[0].ResourceSpans().At(0).Resource().Attributes()
}

func runMetrics(t *testing.T, cfg component.Config, attrs map[string]string) pcommon.Map {
	t.Helper()
	ctx := context.Background()
	factory := transformprocessor.NewFactory()
	sink := new(consumertest.MetricsSink)
	proc, err := factory.CreateMetrics(ctx, processortest.NewNopSettings(factory.Type()), cfg, sink)
	require.NoError(t, err)
	require.NoError(t, proc.Start(ctx, componenttest.NewNopHost()))
	defer func() { require.NoError(t, proc.Shutdown(ctx)) }()

	md := pmetric.NewMetrics()
	rm := md.ResourceMetrics().AppendEmpty()
	putAll(rm.Resource().Attributes(), attrs)
	m := rm.ScopeMetrics().AppendEmpty().Metrics().AppendEmpty()
	m.SetName("metric")
	m.SetEmptyGauge().DataPoints().AppendEmpty().SetIntValue(1)
	require.NoError(t, proc.ConsumeMetrics(ctx, md))

	require.Len(t, sink.AllMetrics(), 1)
	return sink.AllMetrics()[0].ResourceMetrics().At(0).Resource().Attributes()
}

func putAll(m pcommon.Map, attrs map[string]string) {
	for k, v := range attrs {
		m.PutStr(k, v)
	}
}

func assertResourceID(t *testing.T, want string, attrs pcommon.Map) {
	t.Helper()
	got, ok := attrs.Get("cloud.resource_id")
	if want == "" {
		if ok {
			assert.Fail(t, "cloud.resource_id should not be set", "got %q", got.AsString())
		}
		return
	}
	require.True(t, ok, "cloud.resource_id is not set")
	assert.Equal(t, want, got.Str())
}
