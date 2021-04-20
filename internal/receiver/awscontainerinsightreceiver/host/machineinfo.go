package host

import (
	"sync"
	"time"

	"go.uber.org/zap"
)

type MachineInfo struct {
	sync.RWMutex
	logger          *zap.Logger
	nodeCapacity    *NodeCapacity
	ec2metadata     *Ec2metadata
	ebsVolume       *EbsVolume
	ec2Tags         *Ec2Tags
	refreshInterval time.Duration
	shutdownC       chan bool
}

func NewMachineInfo(refreshInterval time.Duration, logger *zap.Logger) *MachineInfo {
	nodeCapacity, err := NewNodeCapacity(logger)
	if err != nil {
		logger.Error("Failed to initialize NodeCapacity", zap.Error(err))
	}

	mInfo := &MachineInfo{
		nodeCapacity:    nodeCapacity,
		ec2metadata:     NewEc2metadata(refreshInterval, logger),
		refreshInterval: refreshInterval,
		shutdownC:       make(chan bool),
		logger:          logger,
	}

	mInfo.lazyInitEbsVolume()
	mInfo.lazyInitEc2Tags()
	return mInfo
}

func (m *MachineInfo) Shutdown() {
	close(m.shutdownC)
}

func (m *MachineInfo) GetInstanceID() string {
	return m.ec2metadata.GetInstanceID()
}

func (m *MachineInfo) GetInstanceType() string {
	return m.ec2metadata.GetInstanceType()
}

func (m *MachineInfo) GetNumCores() int64 {
	return m.nodeCapacity.CPUCapacity
}

func (m *MachineInfo) GetMemoryCapacity() int64 {
	return m.nodeCapacity.MemCapacity
}

func (m *MachineInfo) GetEbsVolumeId(devName string) string {
	if m.ebsVolume != nil {
		return m.ebsVolume.GetEbsVolumeId(devName)
	}

	return ""
}

func (m *MachineInfo) lazyInitEbsVolume() {
	if m.ebsVolume == nil {
		//delay the initialization. If instance id is not available, ebsVolume is set to nil
		//Because ebs volumes only change occasionally, we refresh every 5 collection intervals to reduce ec2 api calls
		m.ebsVolume = NewEbsVolume(m.GetInstanceID(), 5*m.refreshInterval, m.logger)
	}

	go func() {
		refreshTicker := time.NewTicker(m.refreshInterval)
		defer refreshTicker.Stop()
		for {
			select {
			case <-refreshTicker.C:
				if m.ebsVolume != nil {
					return
				}
				m.logger.Info("refresh to initialize ebsVolume")
				m.ebsVolume = NewEbsVolume(m.GetInstanceID(), m.refreshInterval, m.logger)
			case <-m.shutdownC:
				return
			}
		}
	}()
}

func (m *MachineInfo) lazyInitEc2Tags() {
	if m.ec2Tags == nil {
		//delay the initialization. If instance id is not available, c2Tags is set to nil
		m.ec2Tags = NewEc2Tags(m.GetInstanceID(), m.refreshInterval, m.logger)
	}

	go func() {
		refreshTicker := time.NewTicker(m.refreshInterval)
		defer refreshTicker.Stop()
		for {
			select {
			case <-refreshTicker.C:
				if m.ec2Tags != nil {
					return
				}
				m.logger.Info("refresh to initialize ec2Tags")
				m.ec2Tags = NewEc2Tags(m.GetInstanceID(), m.refreshInterval, m.logger)
			case <-m.shutdownC:
				return
			}
		}
	}()
}

func (m *MachineInfo) GetClusterName() string {
	if m.ec2Tags != nil {
		return m.ec2Tags.GetClusterName()
	}

	return ""
}

func (m *MachineInfo) GetAutoScalingGroupName() string {
	if m.ec2Tags != nil {
		return m.ec2Tags.GetAutoScalingGroupName()
	}

	return ""
}
