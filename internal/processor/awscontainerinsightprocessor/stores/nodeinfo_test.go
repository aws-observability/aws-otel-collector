package stores

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGetCPUCapacity(t *testing.T) {
	nodeInfo := newNodeInfo()
	nodeInfo.setCPUCapacity(int(4))
	assert.Equal(t, int64(4), nodeInfo.getCPUCapacity())

	nodeInfo.setCPUCapacity(int32(2))
	assert.Equal(t, int64(2), nodeInfo.getCPUCapacity())

	nodeInfo.setCPUCapacity(int64(4))
	assert.Equal(t, int64(4), nodeInfo.getCPUCapacity())

	nodeInfo.setCPUCapacity(uint(2))
	assert.Equal(t, int64(2), nodeInfo.getCPUCapacity())

	nodeInfo.setCPUCapacity(uint32(4))
	assert.Equal(t, int64(4), nodeInfo.getCPUCapacity())

	nodeInfo.setCPUCapacity(uint64(2))
	assert.Equal(t, int64(2), nodeInfo.getCPUCapacity())

	//with invalid type
	nodeInfo.setCPUCapacity("2")
	assert.Equal(t, int64(0), nodeInfo.getCPUCapacity())

	//with negative value
	nodeInfo.setCPUCapacity(int64(-2))
	assert.Equal(t, int64(0), nodeInfo.getCPUCapacity())
}

func TestSetGetMemCapacity(t *testing.T) {
	nodeInfo := newNodeInfo()
	nodeInfo.setMemCapacity(int(2048))
	assert.Equal(t, int64(2048), nodeInfo.getMemCapacity())

	nodeInfo.setMemCapacity(int32(1024))
	assert.Equal(t, int64(1024), nodeInfo.getMemCapacity())

	nodeInfo.setMemCapacity(int64(2048))
	assert.Equal(t, int64(2048), nodeInfo.getMemCapacity())

	nodeInfo.setMemCapacity(uint(1024))
	assert.Equal(t, int64(1024), nodeInfo.getMemCapacity())

	nodeInfo.setMemCapacity(uint32(2048))
	assert.Equal(t, int64(2048), nodeInfo.getMemCapacity())

	nodeInfo.setMemCapacity(uint64(1024))
	assert.Equal(t, int64(1024), nodeInfo.getMemCapacity())

	//with invalid type
	nodeInfo.setMemCapacity("2")
	assert.Equal(t, int64(0), nodeInfo.getMemCapacity())

	//with negative value
	nodeInfo.setMemCapacity(int64(-2))
	assert.Equal(t, int64(0), nodeInfo.getMemCapacity())
}
