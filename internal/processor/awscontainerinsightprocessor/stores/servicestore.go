package stores

import (
	"sync"
	"time"

	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/aws/k8scommon/k8sclient"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
)

const (
	refreshIntervalService = 10 //10s
)

type ServiceStore struct {
	podKeyToServiceNamesMap map[string][]string
	sync.Mutex
	lastRefreshed time.Time
	logger        *zap.Logger
}

func NewServiceStore(logger *zap.Logger) *ServiceStore {
	serviceStore := &ServiceStore{
		podKeyToServiceNamesMap: make(map[string][]string),
		logger:                  logger,
	}
	return serviceStore
}

func (s *ServiceStore) RefreshTick() {
	now := time.Now()
	if now.Sub(s.lastRefreshed).Seconds() >= refreshIntervalService {
		s.refresh()
		s.lastRefreshed = now
	}
}

// service info is not mandatory
func (s *ServiceStore) Decorate(rm pdata.ResourceMetrics, kubernetesBlob map[string]interface{}) bool {
	attributes := rm.Resource().Attributes()
	if val := GetStringValFromAttributes(K8sPodNameKey, attributes); val != "" {
		podKey := createPodKeyFromMetric(attributes)
		if podKey == "" {
			s.logger.Error("podKey is unavailable when decorating service.", zap.Any("podKey", podKey))
			return false
		}
		if serviceList, ok := s.podKeyToServiceNamesMap[podKey]; ok {
			if len(serviceList) > 0 {
				addServiceNameTag(attributes, serviceList)
			}
		}
	}

	return true
}

func (s *ServiceStore) refresh() {
	s.podKeyToServiceNamesMap = k8sclient.Get().Ep.PodKeyToServiceNames()
	s.logger.Debug("pod to service name map", zap.Any("podKeyToServiceNamesMap", s.podKeyToServiceNamesMap))
}

func addServiceNameTag(attributes pdata.AttributeMap, serviceNames []string) {
	// TODO handle serviceNames len is larger than 1. We need to duplicate the metric object
	AddAttribute(TypeService, serviceNames[0], attributes)
}
