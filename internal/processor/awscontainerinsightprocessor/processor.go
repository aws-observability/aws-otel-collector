package awscontainerinsightprocessor

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/aws-observability/aws-otel-collector/internal/processor/awscontainerinsightprocessor/stores"
	"github.com/aws-observability/aws-otel-collector/internal/processor/awscontainerinsightprocessor/structuredlogsadapter"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

type awscontainerinsightprocessor struct {
	logger          *zap.Logger
	started         bool
	stores          []stores.K8sStore
	shutdownC       chan bool
	TagService      bool   `toml:"tag_service"`
	ClusterName     string `toml:"cluster_name"`
	HostIP          string `toml:"host_ip"`
	NodeName        string `toml:"node_name"`
	PrefFullPodName bool   `toml:"prefer_full_pod_name"`
}

func (acip *awscontainerinsightprocessor) ProcessMetrics(_ context.Context, md pdata.Metrics) (pdata.Metrics, error) {
	resourceMetricsSlice := md.ResourceMetrics()

OUTER:
	for i := 0; i < resourceMetricsSlice.Len(); i++ {
		rm := resourceMetricsSlice.At(i)
		attributes := rm.Resource().Attributes()
		kubernetesBlob := make(map[string]interface{})
		for _, store := range acip.stores {
			if !store.Decorate(rm, kubernetesBlob) {
				// drop the unexpected metric
				continue OUTER
			}
		}
		structuredlogsadapter.AddKubernetesInfo(attributes, kubernetesBlob)
		structuredlogsadapter.TagMetricSource(attributes)

	}

	return md, nil
}

func (acip *awscontainerinsightprocessor) Shutdown() {
	close(acip.shutdownC)
}

func (acip *awscontainerinsightprocessor) start() error {
	acip.NodeName = os.Getenv("HOST_NAME")
	if acip.NodeName == "" {
		return errors.New("missing environment variable HOST_NAME. Please check your YAML config")
	}

	acip.HostIP = os.Getenv("HOST_IP")
	if acip.HostIP == "" {
		return errors.New("missing environment variable HOST_IP. Please check your YAML config")
	}

	acip.shutdownC = make(chan bool)

	acip.stores = append(acip.stores, stores.NewPodStore(acip.HostIP, acip.PrefFullPodName))
	if acip.TagService {
		acip.stores = append(acip.stores, stores.NewServiceStore(acip.logger))
	}

	for _, store := range acip.stores {
		store.RefreshTick()
	}

	go func() {
		refreshTicker := time.NewTicker(time.Second)
		defer refreshTicker.Stop()
		for {
			select {
			case <-refreshTicker.C:
				for _, store := range acip.stores {
					store.RefreshTick()
				}
			case <-acip.shutdownC:
				refreshTicker.Stop()
				return
			}
		}
	}()
	acip.started = true
	return nil
}

func newAwsContainerInsightProcessor(logger *zap.Logger, config *Config) *awscontainerinsightprocessor {
	processor := &awscontainerinsightprocessor{
		logger:          logger,
		TagService:      config.TagService,
		PrefFullPodName: config.PrefFullPodName,
	}
	if err := processor.start(); err != nil {
		processor.logger.Error("Fail to start k8sapiserver", zap.Error(err))
		return nil
	}
	return processor
}
