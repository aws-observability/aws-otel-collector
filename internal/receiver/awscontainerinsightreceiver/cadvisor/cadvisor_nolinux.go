// +build !linux

package cadvisor

import (
	"github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver/cadvisor/extractors"
	"github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver/host"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

// cadvisor doesn't support windows, define the dummy functions

type Cadvisor struct {
}

func New(containerOrchestrator string, machineInfo *host.MachineInfo, logger *zap.Logger) *Cadvisor {
	return nil
}

func (c *Cadvisor) GetMetrics() []pdata.Metrics {
	return nil
}

func GetMetricsExtractors() []extractors.MetricExtractor {
	return nil
}
