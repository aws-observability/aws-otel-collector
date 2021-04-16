package host

import (
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.uber.org/zap"
)

const (
	ClusterNameKey          = "container-insight-eks-cluster-name"
	ClusterNameTagKeyPrefix = "kubernetes.io/cluster/"
	AutoScalingGroupNameTag = "aws:autoscaling:groupName"
)

type Ec2Tags struct {
	refreshInterval      time.Duration
	instanceId           string
	clusterName          string
	autoScalingGroupName string
	logger               *zap.Logger
	shutdownC            chan bool
}

func NewEc2Tags(instanceId string, refreshInterval time.Duration, logger *zap.Logger) *Ec2Tags {
	if instanceId == "" {
		return nil
	}

	et := &Ec2Tags{
		instanceId:      instanceId,
		refreshInterval: refreshInterval,
		shutdownC:       make(chan bool),
		logger:          logger,
	}

	et.refresh()
	go func() {
		refreshTicker := time.NewTicker(et.refreshInterval)
		defer refreshTicker.Stop()
		for {
			select {
			case <-refreshTicker.C:
				if et.clusterName != "" {
					//stop once we get the cluster name
					return
				}
				et.refresh()
			case <-et.shutdownC:
				return
			}
		}
	}()

	return et
}

func (et *Ec2Tags) fetchEc2Tags() map[string]string {
	et.logger.Info("Fetch ec2 tags to detect cluster name and auto scaling group name")
	tags := make(map[string]string)
	//add some sleep jitter to prevent a large number of receivers calling the ec2 api at the same time
	time.Sleep(hostJitter(3 * time.Second))

	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		et.logger.Warn("Fail to set up session to call ec2 api", zap.Error(err))
	}

	tagFilters := []*ec2.Filter{
		{
			Name:   aws.String("resource-type"),
			Values: aws.StringSlice([]string{"instance"}),
		},
		{
			Name:   aws.String("resource-id"),
			Values: aws.StringSlice([]string{et.instanceId}),
		},
	}

	client := ec2.New(sess)
	input := &ec2.DescribeTagsInput{
		Filters: tagFilters,
	}

	for {
		result, err := client.DescribeTags(input)
		if err != nil {
			et.logger.Warn("Fail to call ec2 DescribeTags", zap.Error(err))
		}

		for _, tag := range result.Tags {
			key := *tag.Key
			tags[key] = *tag.Value
			if strings.HasPrefix(key, ClusterNameTagKeyPrefix) && *tag.Value == "owned" {
				tags[ClusterNameKey] = key[len(ClusterNameTagKeyPrefix):]
			}
		}

		if result.NextToken == nil {
			break
		}
		input.SetNextToken(*result.NextToken)
	}

	return tags
}

func (et *Ec2Tags) GetClusterName() string {
	return et.clusterName
}

func (et *Ec2Tags) GetAutoScalingGroupName() string {
	return et.autoScalingGroupName
}

func (et *Ec2Tags) refresh() {
	tags := et.fetchEc2Tags()
	et.clusterName = tags[ClusterNameKey]
	et.autoScalingGroupName = tags[AutoScalingGroupNameTag]
}

func (et *Ec2Tags) Shutdown() {
	close(et.shutdownC)
}
