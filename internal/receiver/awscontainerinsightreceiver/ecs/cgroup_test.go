package ecs

import (
	"go.uber.org/zap"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var logger *zap.Logger

func TestGetCGroupMountPoint(t *testing.T) {
	result, _ := getCGroupMountPoint("test/mountinfo")
	assert.Equal(t, "test", result, "Expected to be equal")
}

func TestGetCPUReservedFromShares(t *testing.T) {
	cgroup := newCGroupScanner("test/mountinfo", logger)

	assert.Equal(t, int64(128), cgroup.getCPUReserved("test1", "", logger))
	assert.Equal(t, int64(128), cgroup.getCPUReserved("test4", "myCluster", logger))
}

func TestGetCPUReservedFromQuota(t *testing.T) {
	cgroup := newCGroupScanner("test/mountinfo", logger)
	assert.Equal(t, int64(256), cgroup.getCPUReserved("test2", "", logger))
}

func TestGetCPUReservedFromBoth(t *testing.T) {
	cgroup := newCGroupScanner("test/mountinfo", logger)
	assert.Equal(t, int64(256), cgroup.getCPUReserved("test3", "", logger))
}

func TestGetCPUReservedFromFalseTaskID(t *testing.T) {
	cgroup := newCGroupScanner("test/mountinfo", logger)
	assert.Equal(t, int64(0), cgroup.getCPUReserved("fake", "", logger))
}

func TestGetMEMReservedFromTask(t *testing.T) {
	cgroup := newCGroupScanner("test/mountinfo", logger)
	containers := []ECSContainer{}
	assert.Equal(t, int64(256), cgroup.getMEMReserved("test1", "", containers, logger))
	assert.Equal(t, int64(256), cgroup.getMEMReserved("test3", "myCluster", containers, logger))
}

func TestGetMEMReservedFromContainers(t *testing.T) {
	cgroup := newCGroupScanner("test/mountinfo", logger)
	containers := []ECSContainer{ECSContainer{DockerId: "container1"}, ECSContainer{DockerId: "container2"}}
	assert.Equal(t, int64(384), cgroup.getMEMReserved("test2", "", containers, logger))
}

func TestGetMEMReservedFromFalseTaskID(t *testing.T) {
	cgroup := newCGroupScanner("test/mountinfo", logger)
	containers := []ECSContainer{ECSContainer{DockerId: "container1"}, ECSContainer{DockerId: "container2"}}
	assert.Equal(t, int64(0), cgroup.getMEMReserved("fake", "", containers, logger))
}

func TestGetCGroupPathForTask(t *testing.T) {
	cgroupMount := "test"
	controller := "cpu"
	taskID := "test1"
	clusterName := "myCluster"
	result, _ := getCGroupPathForTask(cgroupMount, controller, taskID, clusterName)
	assert.Equal(t, path.Join(cgroupMount, controller, "ecs", taskID), result)

	taskID = "test4"
	result, _ = getCGroupPathForTask(cgroupMount, controller, taskID, clusterName)
	assert.Equal(t, path.Join(cgroupMount, controller, "ecs", clusterName, taskID), result)
}
