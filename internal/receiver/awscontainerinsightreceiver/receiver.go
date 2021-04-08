// Copyright 2019, OpenTelemetry Authors
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

package awscontainerinsightreceiver

import (
	"context"
	"errors"
	"time"

	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver/cadvisor"
	"github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver/k8sapiserver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.opentelemetry.io/collector/obsreport"
	"go.uber.org/zap"
)

var _ component.MetricsReceiver = (*awsContainerInsightReceiver)(nil)

// awsContainerInsightReceiver implements the component.MetricsReceiver
type awsContainerInsightReceiver struct {
	logger       *zap.Logger
	nextConsumer consumer.Metrics
	config       *Config
	cancel       context.CancelFunc
	cadvisor     *cadvisor.Cadvisor
	k8sapiserver *k8sapiserver.K8sAPIServer
}

// New creates the aws container insight receiver with the given parameters.
func New(
	logger *zap.Logger,
	config *Config,
	nextConsumer consumer.Metrics) (component.MetricsReceiver, error) {
	if nextConsumer == nil {
		return nil, componenterror.ErrNilNextConsumer
	}

	r := &awsContainerInsightReceiver{
		logger:       logger,
		nextConsumer: nextConsumer,
		config:       config,
	}
	return r, nil
}

// Start collecting metrics from cadvisor and k8s api server (if it is an elected leader)
func (acir *awsContainerInsightReceiver) Start(ctx context.Context, host component.Host) error {
	ctx, acir.cancel = context.WithCancel(obsreport.ReceiverContext(ctx, acir.config.Name(), "http"))
	acir.cadvisor = cadvisor.New(acir.config.ContainerOrchestrator, acir.logger)
	if acir.config.ContainerOrchestrator == common.EKS {
		acir.k8sapiserver = k8sapiserver.New(acir.logger)
	}

	go func() {
		ticker := time.NewTicker(acir.config.CollectionInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				acir.collectData(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

// Shutdown stops the awsContainerInsightReceiver receiver.
func (acir *awsContainerInsightReceiver) Shutdown(context.Context) error {
	acir.cancel()
	return nil
}

// collectData collects container stats from Amazon ECS Task Metadata Endpoint
func (acir *awsContainerInsightReceiver) collectData(ctx context.Context) error {
	var mds []pdata.Metrics
	if acir.cadvisor == nil &&
		(acir.config.ContainerOrchestrator == common.EKS && acir.k8sapiserver == nil) {
		err := errors.New("both cadvisor and k8sapiserver failed to start")
		acir.logger.Error("Failed to collect stats", zap.Error(err))
		return err
	}

	if acir.cadvisor != nil {
		mds = append(mds, acir.cadvisor.GetMetrics()...)
	}

	if acir.k8sapiserver != nil {
		mds = append(mds, acir.k8sapiserver.GetMetrics()...)
	}

	for _, md := range mds {
		err := acir.nextConsumer.ConsumeMetrics(ctx, md)
		if err != nil {
			return err
		}
	}

	return nil
}
