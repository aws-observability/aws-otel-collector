// +build linux

package cadvisor

import (
	"errors"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/google/cadvisor/cache/memory"
	cadvisormetrics "github.com/google/cadvisor/container"
	"github.com/google/cadvisor/container/containerd"
	"github.com/google/cadvisor/container/crio"
	"github.com/google/cadvisor/container/docker"

	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver/cadvisor/extractors"
	"github.com/aws-observability/aws-otel-collector/internal/receiver/awscontainerinsightreceiver/host"
	"go.uber.org/zap"

	"github.com/google/cadvisor/container/systemd"
	cinfo "github.com/google/cadvisor/info/v1"
	"github.com/google/cadvisor/manager"
	"github.com/google/cadvisor/utils/sysfs"
	"go.opentelemetry.io/collector/consumer/pdata"
)

// The amount of time for which to keep stats in memory.
const statsCacheDuration = 2 * time.Minute

// Max collection interval, it is not meaningful if allowDynamicHousekeeping = false
var maxHousekeepingInterval = 15 * time.Second

// When allowDynamicHousekeeping is true, the collection interval is floating between 1s(default) to maxHousekeepingInterval
var allowDynamicHousekeeping = true

const defaultHousekeepingInterval = 10 * time.Second

type Cadvisor struct {
	logger                *zap.Logger
	nodeName              string //get the value from downward API
	manager               manager.Manager
	ebsVolume             *host.EbsVolume
	version               string
	machineInfo           *host.MachineInfo
	containerOrchestrator string
}

func overrideCadvisorFlagDefault(logger *zap.Logger) {
	flagOverrides := map[string]string{
		// Override the default cAdvisor housekeeping interval.
		"housekeeping_interval": defaultHousekeepingInterval.String(),
		// override other defaults (in future)
	}
	for name, defaultValue := range flagOverrides {
		if f := flag.Lookup(name); f != nil {
			f.DefValue = defaultValue
			f.Value.Set(defaultValue)
		} else {
			logger.Warn("Expected cAdvisor flag not found", zap.String("name", name))
		}
	}
}

func New(containerOrchestrator string, logger *zap.Logger) *Cadvisor {
	nodeName := os.Getenv("HOST_NAME")
	if nodeName == "" {
		err := errors.New("Missing environment variable HOST_NAME. Please check your YAML config.")
		logger.Warn("Error in Cadvisor initialization: ", zap.Error(err))
	}

	c := &Cadvisor{
		logger:                logger,
		nodeName:              nodeName,
		version:               "0",
		containerOrchestrator: containerOrchestrator,
	}
	overrideCadvisorFlagDefault(logger)

	if err := c.initManager(); err != nil {
		c.logger.Warn("Fail to initialize cadvisor", zap.Error(err))
		return nil
	}

	c.machineInfo = host.NewMachineInfo(c.manager, time.Minute, c.logger)
	return c
}

var metricsExtractors = []extractors.MetricExtractor{}

func GetMetricsExtractors() []extractors.MetricExtractor {
	return metricsExtractors
}

func (c *Cadvisor) addEbsVolumeInfo(tags map[string]string, ebsVolumeIdsUsedAsPV map[string]string) {
	deviceName, ok := tags[common.DiskDev]
	if !ok {
		return
	}

	if c.ebsVolume != nil {
		if volId := c.ebsVolume.GetEbsVolumeId(deviceName); volId != "" {
			tags[common.HostEbsVolumeId] = volId
		}
	}

	if tags[common.MetricType] == common.TypeContainerFS || tags[common.MetricType] == common.TypeNodeFS ||
		tags[common.MetricType] == common.TypeNodeDiskIO || tags[common.MetricType] == common.TypeContainerDiskIO {
		if volId := ebsVolumeIdsUsedAsPV[deviceName]; volId != "" {
			tags[common.EbsVolumeId] = volId
		}
	}
}

func (c *Cadvisor) decorateMetrics(cadvisormetrics []*extractors.CAdvisorMetric) {
	if c.ebsVolume == nil {
		//delay the initialization. If instance id is not available, c.ebsVolume is set to nil
		c.ebsVolume = host.NewEbsVolume(c.machineInfo.GetInstanceID(), time.Minute, c.logger)
	}

	ebsVolumeIdsUsedAsPV := host.ExtractEbsIdsUsedByKubernetes()

	for _, m := range cadvisormetrics {
		tags := m.GetTags()
		c.addEbsVolumeInfo(tags, ebsVolumeIdsUsedAsPV)

		//add version
		tags[common.Version] = c.version

		//add NodeName for node, pod and container
		metricType := tags[common.MetricType]
		if c.nodeName != "" && (common.IsNode(metricType) || common.IsInstance(metricType) ||
			common.IsPod(metricType) || common.IsContainer(metricType)) {
			tags[common.NodeNameKey] = c.nodeName
		}

		//add instance id and type
		if instanceId := c.machineInfo.GetInstanceID(); instanceId != "" {
			tags[common.InstanceId] = instanceId
		}
		if instanceType := c.machineInfo.GetInstanceType(); instanceType != "" {
			tags[common.InstanceType] = instanceType
		}
	}
}

func (c *Cadvisor) GetMetrics() []pdata.Metrics {
	c.logger.Debug("collect data from cadvisor...")
	var result []pdata.Metrics
	var containerinfos []*cinfo.ContainerInfo
	var err error

	if c.manager == nil && c.initManager() != nil {
		panic("Cannot initiate manager")
	}

	req := &cinfo.ContainerInfoRequest{
		NumStats: 1,
	}

	containerinfos, err = c.manager.SubcontainersInfo("/", req)
	if err != nil {
		c.logger.Warn("GetContainerInfo failed", zap.Error(err))
		return result
	}

	c.logger.Debug("cadvisor containers stats", zap.Int("size", len(containerinfos)))
	results := processContainers(containerinfos, c.machineInfo, c.containerOrchestrator, c.logger)

	c.decorateMetrics(results)

	for _, cadvisorMetric := range results {
		md := common.ConvertToOTLPMetrics(cadvisorMetric.GetFields(), cadvisorMetric.GetTags(), c.logger)
		result = append(result, md)
	}

	return result
}

func (c *Cadvisor) initManager() error {
	sysFs := sysfs.NewRealSysFs()
	includedMetrics := cadvisormetrics.MetricSet{
		cadvisormetrics.CpuUsageMetrics:     struct{}{},
		cadvisormetrics.MemoryUsageMetrics:  struct{}{},
		cadvisormetrics.DiskIOMetrics:       struct{}{},
		cadvisormetrics.NetworkUsageMetrics: struct{}{},
		cadvisormetrics.DiskUsageMetrics:    struct{}{},
	}
	var cgroupRoots []string
	if c.containerOrchestrator == common.EKS {
		cgroupRoots = []string{"/kubepods"}
	}

	houseKeepingConfig := manager.HouskeepingConfig{
		Interval:     &maxHousekeepingInterval,
		AllowDynamic: &allowDynamicHousekeeping,
	}
	// Create and start the cAdvisor container manager.
	m, err := manager.New(memory.New(statsCacheDuration, nil), sysFs, houseKeepingConfig, includedMetrics, http.DefaultClient, cgroupRoots, "")
	if err != nil {
		c.logger.Error("cadvisor manager allocate failed, ", zap.Error(err))
		return err
	}
	cadvisormetrics.RegisterPlugin("containerd", containerd.NewPlugin())
	cadvisormetrics.RegisterPlugin("crio", crio.NewPlugin())
	cadvisormetrics.RegisterPlugin("docker", docker.NewPlugin())
	cadvisormetrics.RegisterPlugin("systemd", systemd.NewPlugin())
	c.manager = m
	err = c.manager.Start()
	if err != nil {
		c.logger.Error("cadvisor manager start failed", zap.Error(err))
		return err
	}

	metricsExtractors = append(metricsExtractors, extractors.NewCpuMetricExtractor(c.logger))
	metricsExtractors = append(metricsExtractors, extractors.NewMemMetricExtractor(c.logger))
	metricsExtractors = append(metricsExtractors, extractors.NewDiskIOMetricExtractor(c.logger))
	metricsExtractors = append(metricsExtractors, extractors.NewNetMetricExtractor(c.logger))
	metricsExtractors = append(metricsExtractors, extractors.NewFileSystemMetricExtractor(c.logger))

	return nil
}
