package common

import "time"

const (
	InstanceId         = "InstanceId"
	InstanceType       = "InstanceType"
	GoPSUtilProcDirEnv = "HOST_PROC"

	MinTimeDiff             = 50 * time.Microsecond // We assume 50 micro-seconds is the minimal gap between two collected data sample to be valid to calculate delta
	ClusterNameKey          = "ClusterName"
	AutoScalingGroupNameKey = "AutoScalingGroupName"
	NodeNameKey             = "NodeName"
	Version                 = "Version"

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

	RunningPodCount       = "number_of_running_pods"
	RunningContainerCount = "number_of_running_containers"
	ContainerCount        = "number_of_containers"
	NodeCount             = "node_count"
	FailedNodeCount       = "failed_node_count"
	ContainerRestartCount = "number_of_container_restarts"

	//unit
	UnitBytes       = "Bytes"
	UnitMegaBytes   = "Megabytes"
	UnitNanoSecond  = "Nanoseconds"
	UnitBytesPerSec = "Bytes/Second"
	UnitCount       = "Count"
	UnitCountPerSec = "Count/Second"
	UnitVCpu        = "vCPU"
	UnitPercent     = "Percent"
)

var metricToUnitMap map[string]string

func init() {
	metricToUnitMap = map[string]string{
		//cpu metrics
		//The following metrics are reported in unit of millicores, but cloudwatch doesn't support it
		// CpuTotal
		// CpuUser
		// CpuSystem
		// CpuLimit
		// CpuRequest
		CpuUtilization:             UnitPercent,
		CpuReservedCapacity:        UnitPercent,
		CpuUtilizationOverPodLimit: UnitPercent,

		//memory metrics
		MemUsage:                   UnitBytes,
		MemCache:                   UnitBytes,
		MemRss:                     UnitBytes,
		MemMaxusage:                UnitBytes,
		MemSwap:                    UnitBytes,
		MemFailcnt:                 UnitCount,
		MemMappedfile:              UnitBytes,
		MemWorkingset:              UnitBytes,
		MemLimit:                   UnitBytes,
		MemUtilization:             UnitPercent,
		MemReservedCapacity:        UnitPercent,
		MemUtilizationOverPodLimit: UnitPercent,

		MemPgfault:                UnitCountPerSec,
		MemPgmajfault:             UnitCountPerSec,
		MemHierarchicalPgfault:    UnitCountPerSec,
		MemHierarchicalPgmajfault: UnitCountPerSec,

		//disk io metrics
		DiskIOServiceBytesPrefix + DiskIOAsync: UnitBytesPerSec,
		DiskIOServiceBytesPrefix + DiskIORead:  UnitBytesPerSec,
		DiskIOServiceBytesPrefix + DiskIOSync:  UnitBytesPerSec,
		DiskIOServiceBytesPrefix + DiskIOWrite: UnitBytesPerSec,
		DiskIOServiceBytesPrefix + DiskIOTotal: UnitBytesPerSec,
		DiskIOServicedPrefix + DiskIOAsync:     UnitCountPerSec,
		DiskIOServicedPrefix + DiskIORead:      UnitCountPerSec,
		DiskIOServicedPrefix + DiskIOSync:      UnitCountPerSec,
		DiskIOServicedPrefix + DiskIOWrite:     UnitCountPerSec,
		DiskIOServicedPrefix + DiskIOTotal:     UnitCountPerSec,

		//network metrics
		NetRxBytes:    UnitBytesPerSec,
		NetRxPackets:  UnitCountPerSec,
		NetRxDropped:  UnitCountPerSec,
		NetRxErrors:   UnitCountPerSec,
		NetTxBytes:    UnitBytesPerSec,
		NetTxPackets:  UnitCountPerSec,
		NetTxDropped:  UnitCountPerSec,
		NetTxErrors:   UnitCountPerSec,
		NetTotalBytes: UnitBytesPerSec,

		//filesystem metrics
		FSUsage:       UnitBytes,
		FSCapacity:    UnitBytes,
		FSAvailable:   UnitBytes,
		FSInodes:      UnitCount,
		FSInodesfree:  UnitCount,
		FSUtilization: UnitPercent,

		//cluster metrics
		NodeCount:       UnitCount,
		FailedNodeCount: UnitCount,

		//others
		RunningPodCount:       UnitCount,
		RunningContainerCount: UnitCount,
		ContainerCount:        UnitCount,
		ContainerRestartCount: UnitCount,
	}
}
