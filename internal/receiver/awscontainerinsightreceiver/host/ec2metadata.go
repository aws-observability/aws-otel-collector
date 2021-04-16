package host

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	awsec2metadata "github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.uber.org/zap"
)

type Ec2metadata struct {
	logger          *zap.Logger
	refreshInterval time.Duration
	instanceId      string
	instanceType    string
	shutdownC       chan bool
}

func NewEc2metadata(refreshInterval time.Duration, logger *zap.Logger) *Ec2metadata {
	emd := &Ec2metadata{
		refreshInterval: refreshInterval,
		shutdownC:       make(chan bool),
		logger:          logger,
	}

	emd.refresh()
	go func() {
		refreshTicker := time.NewTicker(emd.refreshInterval)
		defer refreshTicker.Stop()
		for {
			select {
			case <-refreshTicker.C:
				//stop the refresh once we get instance id and type successfully
				if emd.instanceId != "" && emd.instanceType != "" {
					return
				}
				emd.refresh()
			case <-emd.shutdownC:
				return
			}
		}
	}()

	return emd
}

func (emd *Ec2metadata) getAwsMetadata() awsec2metadata.EC2InstanceIdentityDocument {
	emd.logger.Info("Fetch instance id and type from ec2 metadata")
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		emd.logger.Error("Failed to set up new session to call ec2 api")
		return awsec2metadata.EC2InstanceIdentityDocument{}
	}
	client := awsec2metadata.New(sess)
	doc, err := client.GetInstanceIdentityDocument()
	if err != nil {
		emd.logger.Error("Failed to get ec2 metadata", zap.Error(err))
		return awsec2metadata.EC2InstanceIdentityDocument{}
	}

	return doc
}

func (emd *Ec2metadata) refresh() {
	metadata := emd.getAwsMetadata()
	emd.instanceId = metadata.InstanceID
	emd.instanceType = metadata.InstanceType
}

func (emd *Ec2metadata) Shutdown() {
	close(emd.shutdownC)
}

func (emd *Ec2metadata) GetInstanceID() string {
	return emd.instanceId
}

func (emd *Ec2metadata) GetInstanceType() string {
	return emd.instanceType
}
