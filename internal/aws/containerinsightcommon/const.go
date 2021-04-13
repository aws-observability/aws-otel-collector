package common

const (
	InstanceId         = "InstanceId"
	InstanceType       = "InstanceType"
	GoPSUtilProcDirEnv = "HOST_PROC"

	MinTimeDiff    = 50 * 1000 // We assume 50 micro-seconds is the minimal gap between two collected data sample to be valid to calculate delta
	ClusterNameKey = "ClusterName"
	NodeNameKey    = "NodeName"
	Version        = "Version"

	MetricType = "Type"
	SourcesKey = "Sources"

	// metric collected
	CpuTotal                   = "cpu_usage_total"
	CpuUser                    = "cpu_usage_user"
	CpuSystem                  = "cpu_usage_system"
	CpuLimit                   = "cpu_limit"
	CpuUtilization             = "cpu_utilization"
	CpuRequest                 = "cpu_request"
	CpuReservedCapacity        = "cpu_reserved_capacity"
	CpuUtilizationOverPodLimit = "cpu_utilization_over_pod_limit"

	MemUsage                   = "memory_usage"
	MemCache                   = "memory_cache"
	MemRss                     = "memory_rss"
	MemMaxusage                = "memory_max_usage"
	MemSwap                    = "memory_swap"
	MemFailcnt                 = "memory_failcnt"
	MemMappedfile              = "memory_mapped_file"
	MemWorkingset              = "memory_working_set"
	MemPgfault                 = "memory_pgfault"
	MemPgmajfault              = "memory_pgmajfault"
	MemHierarchicalPgfault     = "memory_hierarchical_pgfault"
	MemHierarchicalPgmajfault  = "memory_hierarchical_pgmajfault"
	MemLimit                   = "memory_limit"
	MemRequest                 = "memory_request"
	MemUtilization             = "memory_utilization"
	MemReservedCapacity        = "memory_reserved_capacity"
	MemUtilizationOverPodLimit = "memory_utilization_over_pod_limit"

	NetIfce       = "interface"
	NetRxBytes    = "network_rx_bytes"
	NetRxPackets  = "network_rx_packets"
	NetRxDropped  = "network_rx_dropped"
	NetRxErrors   = "network_rx_errors"
	NetTxBytes    = "network_tx_bytes"
	NetTxPackets  = "network_tx_packets"
	NetTxDropped  = "network_tx_dropped"
	NetTxErrors   = "network_tx_errors"
	NetTotalBytes = "network_total_bytes"

	DiskDev         = "device"
	EbsVolumeId     = "ebs_volume_id" //used by kubernetes cluster as persistent volume
	HostEbsVolumeId = "EBSVolumeId"   //used by host filesystem

	FSType        = "fstype"
	FSUsage       = "filesystem_usage"
	FSCapacity    = "filesystem_capacity"
	FSAvailable   = "filesystem_available"
	FSInodes      = "filesystem_inodes"
	FSInodesfree  = "filesystem_inodes_free"
	FSUtilization = "filesystem_utilization"

	DiskIOServiceBytesPrefix = "diskio_io_service_bytes_"
	DiskIOServicedPrefix     = "diskio_io_serviced_"
	DiskIOAsync              = "Async"
	DiskIORead               = "Read"
	DiskIOSync               = "Sync"
	DiskIOWrite              = "Write"
	DiskIOTotal              = "Total"

	TypeCluster          = "Cluster"
	TypeClusterService   = "ClusterService"
	TypeClusterNamespace = "ClusterNamespace"
	TypeService          = "Service"

	// Both TypeInstance and TypeNode mean EC2 Instance, they are used in ECS and EKS separately
	TypeInstance       = "Instance"
	TypeNode           = "Node"
	TypeInstanceFS     = "InstanceFS"
	TypeNodeFS         = "NodeFS"
	TypeInstanceNet    = "InstanceNet"
	TypeNodeNet        = "NodeNet"
	TypeInstanceDiskIO = "InstanceDiskIO"
	TypeNodeDiskIO     = "NodeDiskIO"

	TypePod             = "Pod"
	TypePodNet          = "PodNet"
	TypeContainer       = "Container"
	TypeContainerFS     = "ContainerFS"
	TypeContainerDiskIO = "ContainerDiskIO"

	//unit
	UnitBytes       = "Bytes"
	UnitMegaBytes   = "Megabytes"
	UnitNanoSecond  = "Nanoseconds"
	UnitBytesPerSec = "Bytes/Second"
	UnitCount       = "Count"
	UnitVCpu        = "vCPU"
	UnitPercent     = "Percent"
)

var metricToUnitMap = make(map[string]string)

func init() {
	//cpu metrics
	metricToUnitMap[CpuTotal] = UnitNanoSecond
	metricToUnitMap[CpuUser] = UnitNanoSecond
	metricToUnitMap[CpuSystem] = UnitNanoSecond
	metricToUnitMap[CpuLimit] = UnitCount
	metricToUnitMap[CpuUtilization] = UnitPercent

	//memory metrics
	metricToUnitMap[MemUsage] = UnitBytes
	metricToUnitMap[MemCache] = UnitBytes
	metricToUnitMap[MemRss] = UnitBytes
	metricToUnitMap[MemMaxusage] = UnitBytes
	metricToUnitMap[MemSwap] = UnitBytes
	metricToUnitMap[MemFailcnt] = UnitCount
	metricToUnitMap[MemMappedfile] = UnitBytes
	metricToUnitMap[MemWorkingset] = UnitBytes
	metricToUnitMap[MemLimit] = UnitBytes
	metricToUnitMap[MemUtilization] = UnitPercent
	//need to compute the rate of following metrics later in processor
	metricToUnitMap[MemPgfault] = UnitCount
	metricToUnitMap[MemPgmajfault] = UnitCount
	metricToUnitMap[MemHierarchicalPgfault] = UnitCount
	metricToUnitMap[MemHierarchicalPgmajfault] = UnitCount

	//disk io metrics
	//need to compute the rate of following metrics later in processor
	metricToUnitMap[DiskIOServiceBytesPrefix+DiskIOAsync] = UnitBytes
	metricToUnitMap[DiskIOServiceBytesPrefix+DiskIORead] = UnitBytes
	metricToUnitMap[DiskIOServiceBytesPrefix+DiskIOSync] = UnitBytes
	metricToUnitMap[DiskIOServiceBytesPrefix+DiskIOWrite] = UnitBytes
	metricToUnitMap[DiskIOServiceBytesPrefix+DiskIOTotal] = UnitBytes
	metricToUnitMap[DiskIOServicedPrefix+DiskIOAsync] = UnitCount
	metricToUnitMap[DiskIOServicedPrefix+DiskIORead] = UnitCount
	metricToUnitMap[DiskIOServicedPrefix+DiskIOSync] = UnitCount
	metricToUnitMap[DiskIOServicedPrefix+DiskIOWrite] = UnitCount
	metricToUnitMap[DiskIOServicedPrefix+DiskIOTotal] = UnitCount

	//network metrics
	//need to compute the rate of following metrics later in processor
	metricToUnitMap[NetRxBytes] = UnitBytes
	metricToUnitMap[NetRxPackets] = UnitCount
	metricToUnitMap[NetRxDropped] = UnitCount
	metricToUnitMap[NetRxErrors] = UnitCount
	metricToUnitMap[NetTxBytes] = UnitBytes
	metricToUnitMap[NetTxPackets] = UnitCount
	metricToUnitMap[NetTxDropped] = UnitCount
	metricToUnitMap[NetTxErrors] = UnitCount
	metricToUnitMap[NetTotalBytes] = UnitBytes

	//filesystem metrics
	metricToUnitMap[FSUsage] = UnitBytes
	metricToUnitMap[FSCapacity] = UnitBytes
	metricToUnitMap[FSAvailable] = UnitBytes
	metricToUnitMap[FSUtilization] = UnitPercent
	metricToUnitMap[FSInodes] = UnitCount
	metricToUnitMap[FSInodesfree] = UnitCount

	//cluster metrics
	metricToUnitMap[RunningPodCount] = UnitCount
	metricToUnitMap[NodeCount] = UnitCount
	metricToUnitMap[FailedNodeCount] = UnitCount
}
