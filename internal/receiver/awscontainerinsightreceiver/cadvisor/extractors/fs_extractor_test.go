package extractors

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"

	. "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	cinfo "github.com/google/cadvisor/info/v1"
	"github.com/stretchr/testify/assert"
)

const (
	testAllowList = `
{
  "/kubepods/besteffort/podaf16b540-4ae2-11e9-977b-0672b6c6fc94/573ee6cd04a6208af809b2329652c74386f1992faca8662c733d7f250014e718": {
    "id": "573ee6cd04a6208af809b2329652c74386f1992faca8662c733d7f250014e718",
    "name": "/kubepods/besteffort/podaf16b540-4ae2-11e9-977b-0672b6c6fc94/573ee6cd04a6208af809b2329652c74386f1992faca8662c733d7f250014e718",
    "aliases": [
      "k8s_ubuntu_stress-1-core-mh2pn_default_af16b540-4ae2-11e9-977b-0672b6c6fc94_0",
      "573ee6cd04a6208af809b2329652c74386f1992faca8662c733d7f250014e718"
    ],
    "namespace": "docker",
    "spec": {
      "creation_time": "2019-03-20T07:35:09.746280405Z",
      "labels": {
        "annotation.io.kubernetes.container.hash": "70bfcd85",
        "annotation.io.kubernetes.container.restartCount": "0",
        "annotation.io.kubernetes.container.terminationMessagePath": "/dev/termination-log",
        "annotation.io.kubernetes.container.terminationMessagePolicy": "File",
        "annotation.io.kubernetes.pod.terminationGracePeriod": "30",
        "io.kubernetes.container.logpath": "/var/log/pods/af16b540-4ae2-11e9-977b-0672b6c6fc94/ubuntu/0.log",
        "io.kubernetes.container.name": "ubuntu",
        "io.kubernetes.docker.type": "container",
        "io.kubernetes.pod.name": "stress-1-core-mh2pn",
        "io.kubernetes.pod.namespace": "default",
        "io.kubernetes.pod.uid": "af16b540-4ae2-11e9-977b-0672b6c6fc94",
        "io.kubernetes.sandbox.id": "a5bb552d7fb8e5014468756f165732e0c6bcd9dcbd229efc51afc014317d20d6"
      },
      "has_cpu": true,
      "cpu": {
        "limit": 2,
        "max_limit": 0,
        "mask": "0-3",
        "period": 100000
      },
      "has_memory": true,
      "memory": {
        "limit": 9223372036854771712,
        "reservation": 9223372036854771712,
        "swap_limit": 9223372036854771712
      },
      "has_network": false,
      "has_filesystem": true,
      "has_diskio": false,
      "has_custom_metrics": false,
      "image": "ubuntu@sha256:017eef0b616011647b269b5c65826e2e2ebddbe5d1f8c1e56b3599fb14fabec8"
    },
    "stats": [
      {
        "timestamp": "2019-04-09T22:26:42.984081498Z",
        "filesystem": [
          {
            "device": "tmpfs",
            "type": "vfs",
            "capacity": 21462233088,
            "usage": 25661440,
            "base_usage": 25640960,
            "available": 0,
            "has_inodes": false,
            "inodes": 67,
            "inodes_free": 0,
            "reads_completed": 0,
            "reads_merged": 0,
            "sectors_read": 0,
            "read_time": 0,
            "writes_completed": 0,
            "writes_merged": 0,
            "sectors_written": 0,
            "write_time": 0,
            "io_in_progress": 0,
            "io_time": 0,
            "weighted_io_time": 0
          },
          {
            "device": "/dev/xvda1",
            "type": "vfs",
            "capacity": 21462233088,
            "usage": 25661440,
            "base_usage": 25640960,
            "available": 0,
            "has_inodes": false,
            "inodes": 67,
            "inodes_free": 0,
            "reads_completed": 0,
            "reads_merged": 0,
            "sectors_read": 0,
            "read_time": 0,
            "writes_completed": 0,
            "writes_merged": 0,
            "sectors_written": 0,
            "write_time": 0,
            "io_in_progress": 0,
            "io_time": 0,
            "weighted_io_time": 0
          },
          {
            "device": "overlay",
            "type": "vfs",
            "capacity": 21462233088,
            "usage": 25661440,
            "base_usage": 25640960,
            "available": 0,
            "has_inodes": false,
            "inodes": 67,
            "inodes_free": 0,
            "reads_completed": 0,
            "reads_merged": 0,
            "sectors_read": 0,
            "read_time": 0,
            "writes_completed": 0,
            "writes_merged": 0,
            "sectors_written": 0,
            "write_time": 0,
            "io_in_progress": 0,
            "io_time": 0,
            "weighted_io_time": 0
          },
          {
            "device": "/dev",
            "type": "vfs",
            "capacity": 21462233088,
            "usage": 25661440,
            "base_usage": 25640960,
            "available": 0,
            "has_inodes": false,
            "inodes": 67,
            "inodes_free": 0,
            "reads_completed": 0,
            "reads_merged": 0,
            "sectors_read": 0,
            "read_time": 0,
            "writes_completed": 0,
            "writes_merged": 0,
            "sectors_written": 0,
            "write_time": 0,
            "io_in_progress": 0,
            "io_time": 0,
            "weighted_io_time": 0
          },
          {
            "device": "overlaytest",
            "type": "vfs",
            "capacity": 21462233088,
            "usage": 25661440,
            "base_usage": 25640960,
            "available": 0,
            "has_inodes": false,
            "inodes": 67,
            "inodes_free": 0,
            "reads_completed": 0,
            "reads_merged": 0,
            "sectors_read": 0,
            "read_time": 0,
            "writes_completed": 0,
            "writes_merged": 0,
            "sectors_written": 0,
            "write_time": 0,
            "io_in_progress": 0,
            "io_time": 0,
            "weighted_io_time": 0
          }
        ]
      }
    ]
  }
}
`
)

func TestFSStats(t *testing.T) {
	result := loadContainerInfo("./testdata/CurInfoContainer.json")
	//container type
	containerType := TypeContainer
	extractor := NewFileSystemMetricExtractor(nil)

	var cMetrics []*CAdvisorMetric
	if extractor.HasValue(result[0]) {
		cMetrics = extractor.GetValue(result[0], nil, containerType)
	}
	for _, cadvisorMetric := range cMetrics {
		log.Printf("cadvisor Metrics received:\n %v \n", *cadvisorMetric)
	}
	expectedFields := map[string]interface{}{
		"container_filesystem_usage":       uint64(25661440),
		"container_filesystem_capacity":    uint64(21462233088),
		"container_filesystem_available":   uint64(0),
		"container_filesystem_utilization": float64(0.11956556381986117),
	}
	expectedTags := map[string]string{
		"device": "/dev/xvda1",
		"fstype": "vfs",
		"Type":   "ContainerFS",
	}
	AssertContainsTaggedField(t, cMetrics[0], expectedFields, expectedTags)

	//pod type
	containerType = TypePod
	extractor = NewFileSystemMetricExtractor(nil)

	if extractor.HasValue(result[0]) {
		cMetrics = extractor.GetValue(result[0], nil, containerType)
	}

	assert.Equal(t, len(cMetrics), 0)

	//node type for eks

	result2 := loadContainerInfo("./testdata/CurInfoNode.json")
	containerType = TypeNode
	extractor = NewFileSystemMetricExtractor(nil)

	if extractor.HasValue(result2[0]) {
		cMetrics = extractor.GetValue(result2[0], nil, containerType)
	}

	for _, cadvisorMetric := range cMetrics {
		log.Printf("cadvisor Metrics received:\n %v \n", *cadvisorMetric)
	}
	expectedFields = map[string]interface{}{
		"node_filesystem_available":   uint64(67108864),
		"node_filesystem_capacity":    uint64(67108864),
		"node_filesystem_inodes":      uint64(2052980),
		"node_filesystem_inodes_free": uint64(2052979),
		"node_filesystem_usage":       uint64(0),
		"node_filesystem_utilization": float64(0),
	}
	expectedTags = map[string]string{
		"device": "/dev/shm",
		"fstype": "vfs",
		"Type":   "NodeFS",
	}
	AssertContainsTaggedField(t, cMetrics[0], expectedFields, expectedTags)

	expectedFields = map[string]interface{}{
		"node_filesystem_available":   uint64(6925574144),
		"node_filesystem_capacity":    uint64(21462233088),
		"node_filesystem_inodes":      uint64(10484672),
		"node_filesystem_inodes_free": uint64(10387672),
		"node_filesystem_usage":       uint64(14536658944),
		"node_filesystem_utilization": float64(67.73134409824186),
	}
	expectedTags = map[string]string{
		"device": "/dev/xvda1",
		"fstype": "vfs",
		"Type":   "NodeFS",
	}
	AssertContainsTaggedField(t, cMetrics[1], expectedFields, expectedTags)

	expectedFields = map[string]interface{}{
		"node_filesystem_available":   uint64(10682417152),
		"node_filesystem_capacity":    uint64(10726932480),
		"node_filesystem_inodes":      uint64(5242880),
		"node_filesystem_inodes_free": uint64(5242877),
		"node_filesystem_usage":       uint64(44515328),
		"node_filesystem_utilization": float64(0.4149865591397849),
	}
	expectedTags = map[string]string{
		"device": "/dev/xvdce",
		"fstype": "vfs",
		"Type":   "NodeFS",
	}
	AssertContainsTaggedField(t, cMetrics[2], expectedFields, expectedTags)
}

func TestAllowList(t *testing.T) {
	extractor := NewFileSystemMetricExtractor(nil)
	assert.Equal(t, true, extractor.allowListRegexP.MatchString("/dev/shm"))
	assert.Equal(t, true, extractor.allowListRegexP.MatchString("tmpfs"))
	assert.Equal(t, true, extractor.allowListRegexP.MatchString("overlay"))
	assert.Equal(t, false, extractor.allowListRegexP.MatchString("overlaytest"))
	assert.Equal(t, false, extractor.allowListRegexP.MatchString("/dev"))
}

func TestFSStatsWithAllowList(t *testing.T) {
	var result []*cinfo.ContainerInfo
	containers := map[string]*cinfo.ContainerInfo{}
	err := json.Unmarshal([]byte(testAllowList), &containers)

	if err != nil {
		log.Printf("Fail to read request body: %s", err)
	}

	for _, containerInfo := range containers {
		result = append(result, containerInfo)
	}

	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(result)
	containerType := TypeContainer
	extractor := NewFileSystemMetricExtractor(nil)

	var cMetrics []*CAdvisorMetric
	if extractor.HasValue(result[0]) {
		cMetrics = extractor.GetValue(result[0], nil, containerType)
	}
	for _, cadvisorMetric := range cMetrics {
		log.Printf("cadvisor Metrics received:\n %v \n", *cadvisorMetric)
	}
	// There are 3 valid device names which pass the allowlist in testAllowList json.
	assert.Equal(t, 3, len(cMetrics))
	assert.Equal(t, "tmpfs", cMetrics[0].tags["device"])
	assert.Equal(t, "/dev/xvda1", cMetrics[1].tags["device"])
	assert.Equal(t, "overlay", cMetrics[2].tags["device"])

}
