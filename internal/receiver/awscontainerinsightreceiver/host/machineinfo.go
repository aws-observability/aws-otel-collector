package host

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	cinfo "github.com/google/cadvisor/info/v1"

	"go.uber.org/zap"
)

type MachineInfo struct {
	sync.RWMutex
	logger          *zap.Logger
	manager         cadvisorManager
	info            *cinfo.MachineInfo
	refreshInterval time.Duration
	instanceId      string
	instanceType    string
	shutdownC       chan bool
}

//define this interface to remove the dependency on "github.com/google/cadvisor/manager" which only compiles on linux
type cadvisorManager interface {
	GetMachineInfo() (*cinfo.MachineInfo, error)
}

func NewMachineInfo(manager cadvisorManager, refreshInterval time.Duration, logger *zap.Logger) *MachineInfo {
	mInfo := &MachineInfo{
		manager:         manager,
		refreshInterval: refreshInterval,
		shutdownC:       make(chan bool),
		logger:          logger,
	}

	mInfo.refresh()

	go func() {
		refreshTicker := time.NewTicker(mInfo.refreshInterval)
		defer refreshTicker.Stop()
		for {
			select {
			case <-refreshTicker.C:
				mInfo.refresh()
			case <-mInfo.shutdownC:
				return
			}
		}
	}()

	return mInfo
}

func (m *MachineInfo) refresh() {
	if info, err := m.manager.GetMachineInfo(); err != nil {
		m.logger.Error("Failed to get machine info", zap.Error(err))
	} else {
		m.Lock()
		defer m.Unlock()
		m.info = info
	}

	if m.instanceId == "" || m.instanceType == "" {
		metadata := m.getAwsMetadata()
		m.instanceId = metadata.InstanceID
		m.instanceType = metadata.InstanceType
	}
}

func (m *MachineInfo) Shutdown() {
	close(m.shutdownC)
}

func (m *MachineInfo) getAwsMetadata() ec2metadata.EC2InstanceIdentityDocument {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		m.logger.Error("Failed to set up new session to call ec2 api")
		return ec2metadata.EC2InstanceIdentityDocument{}
	}
	client := ec2metadata.New(sess)
	doc, err := client.GetInstanceIdentityDocument()
	if err != nil {
		m.logger.Error("Failed to get ec2 metadata", zap.Error(err))
		return ec2metadata.EC2InstanceIdentityDocument{}
	}

	return doc
}

func (m *MachineInfo) GetInstanceID() string {
	return m.instanceId
}

func (m *MachineInfo) GetInstanceType() string {
	return m.instanceType
}

func (m *MachineInfo) GetNumCores() int {
	m.RLock()
	defer m.RUnlock()
	if m.info == nil {
		return 0
	}
	return m.info.NumCores
}

func (m *MachineInfo) GetMemoryCapacity() uint64 {
	m.RLock()
	defer m.RUnlock()
	if m.info == nil {
		return 0
	}
	return m.info.MemoryCapacity
}
