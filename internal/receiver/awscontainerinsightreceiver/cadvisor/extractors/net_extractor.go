package extractors

import (
	"time"

	aws "github.com/aws-observability/aws-otel-collector/internal/aws"
	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	cinfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"
)

type NetMetricExtractor struct {
	logger         *zap.Logger
	rateCalculator aws.MetricCalculator
}

func getInterfacesStats(stats *cinfo.ContainerStats) []cinfo.InterfaceStats {
	ifceStats := stats.Network.Interfaces
	if len(ifceStats) == 0 {
		ifceStats = []cinfo.InterfaceStats{stats.Network.InterfaceStats}
	}
	return ifceStats
}

func (n *NetMetricExtractor) HasValue(info *cinfo.ContainerInfo) bool {
	return info.Spec.HasNetwork
}

func (n *NetMetricExtractor) GetValue(info *cinfo.ContainerInfo, _ MachineInfoProvider, containerType string) []*CAdvisorMetric {
	var metrics []*CAdvisorMetric

	// Just a protection here, there is no Container level Net metrics
	if (containerType == TypePod && info.Spec.Labels[containerNameLable] != infraContainerName) || containerType == TypeContainer {
		return metrics
	}

	curStats := GetStats(info)
	curIfceStats := getInterfacesStats(curStats)

	// used for aggregation
	var netIfceMetrics []map[string]interface{}

	for _, cur := range curIfceStats {
		mType := getNetMetricType(containerType, n.logger)
		netIfceMetric := make(map[string]interface{})

		infoName := info.Name + containerType + cur.Name //used to identify the network interface
		multiplier := float64(time.Second)
		assignRateValueToField(&n.rateCalculator, netIfceMetric, NetRxBytes, infoName, float64(cur.RxBytes), curStats.Timestamp, multiplier)
		assignRateValueToField(&n.rateCalculator, netIfceMetric, NetRxPackets, infoName, float64(cur.RxPackets), curStats.Timestamp, multiplier)
		assignRateValueToField(&n.rateCalculator, netIfceMetric, NetRxDropped, infoName, float64(cur.RxDropped), curStats.Timestamp, multiplier)
		assignRateValueToField(&n.rateCalculator, netIfceMetric, NetRxErrors, infoName, float64(cur.RxErrors), curStats.Timestamp, multiplier)
		assignRateValueToField(&n.rateCalculator, netIfceMetric, NetTxBytes, infoName, float64(cur.TxBytes), curStats.Timestamp, multiplier)
		assignRateValueToField(&n.rateCalculator, netIfceMetric, NetTxPackets, infoName, float64(cur.TxPackets), curStats.Timestamp, multiplier)
		assignRateValueToField(&n.rateCalculator, netIfceMetric, NetTxDropped, infoName, float64(cur.TxDropped), curStats.Timestamp, multiplier)
		assignRateValueToField(&n.rateCalculator, netIfceMetric, NetTxErrors, infoName, float64(cur.TxErrors), curStats.Timestamp, multiplier)

		if netIfceMetric[NetRxBytes] != nil && netIfceMetric[NetTxBytes] != nil {
			netIfceMetric[NetTotalBytes] = netIfceMetric[NetRxBytes].(float64) + netIfceMetric[NetTxBytes].(float64)
		}

		netIfceMetrics = append(netIfceMetrics, netIfceMetric)

		metric := newCadvisorMetric(mType)
		metric.tags[NetIfce] = cur.Name
		for k, v := range netIfceMetric {
			metric.fields[MetricName(mType, k)] = v
		}

		metrics = append(metrics, metric)
	}

	aggregatedFields := AggregateFields(netIfceMetrics)
	if len(aggregatedFields) > 0 {
		metric := newCadvisorMetric(containerType)
		for k, v := range aggregatedFields {
			metric.fields[MetricName(containerType, k)] = v
		}
		metrics = append(metrics, metric)
	}

	return metrics
}

func NewNetMetricExtractor(logger *zap.Logger) *NetMetricExtractor {
	return &NetMetricExtractor{
		logger:         logger,
		rateCalculator: newFloat64RateCalculator(),
	}
}

func getNetMetricType(containerType string, logger *zap.Logger) string {
	metricType := ""
	switch containerType {
	case TypeNode:
		metricType = TypeNodeNet
	case TypeInstance:
		metricType = TypeInstanceNet
	case TypePod:
		metricType = TypePodNet
	default:
		logger.Warn("net_extractor: net metric extractor is parsing unexpected containerType", zap.String("containerType", containerType))
	}
	return metricType
}
