package host

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"go.uber.org/zap"
)

type EbsVolume struct {
	// device name to volumeId mapping
	dev2Vol         map[string]string
	refreshInterval time.Duration
	sync.RWMutex
	instanceId string
	logger     *zap.Logger
	shutdownC  chan bool
}

func NewEbsVolume(instanceId string, refreshInterval time.Duration, logger *zap.Logger) *EbsVolume {
	if instanceId == "" {
		return nil
	}

	ebsVolume := &EbsVolume{
		dev2Vol:         make(map[string]string),
		instanceId:      instanceId,
		refreshInterval: refreshInterval,
		shutdownC:       make(chan bool),
		logger:          logger,
	}

	//add some sleep jitter to prevent a large number of receivers calling the ec2 api at the same time
	time.Sleep(hostJitter(time.Second))
	ebsVolume.refresh()

	go func() {
		refreshTicker := time.NewTicker(ebsVolume.refreshInterval)
		defer refreshTicker.Stop()
		for {
			select {
			case <-refreshTicker.C:
				ebsVolume.refresh()
			case <-ebsVolume.shutdownC:
				return
			}
		}
	}()

	return ebsVolume
}

func (e *EbsVolume) Shutdown() {
	close(e.shutdownC)
}

func (e *EbsVolume) refresh() {
	if e.instanceId == "" {
		return
	}

	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		e.logger.Warn("Fail to set up session to call ec2 api", zap.Error(err))
	}
	client := ec2.New(sess)

	input := &ec2.DescribeVolumesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("attachment.instance-id"),
				Values: aws.StringSlice([]string{e.instanceId}),
			},
		},
	}

	devPathSet := make(map[string]bool)
	allSuccess := false
	for {
		result, err := client.DescribeVolumes(input)
		if err != nil {
			allSuccess = false
			e.logger.Warn("Fail to call ec2 DescribeVolumes", zap.Error(err))
		}
		for _, volume := range result.Volumes {
			for _, attachment := range volume.Attachments {
				devPath := e.addEbsVolumeMapping(volume.AvailabilityZone, attachment)
				devPathSet[devPath] = true
			}
		}
		allSuccess = true
		if result.NextToken == nil {
			break
		}
		input.SetNextToken(*result.NextToken)
	}

	if allSuccess {
		e.Lock()
		defer e.Unlock()
		for k := range e.dev2Vol {
			if !devPathSet[k] {
				delete(e.dev2Vol, k)
			}
		}
	}
}

func (e *EbsVolume) addEbsVolumeMapping(zone *string, attachement *ec2.VolumeAttachment) string {
	// *attachement.Device is sth like: /dev/xvda
	devPath := findNvmeBlockNameIfPresent(*attachement.Device)
	if devPath == "" {
		devPath = *attachement.Device
	}

	e.Lock()
	defer e.Unlock()
	e.dev2Vol[devPath] = fmt.Sprintf("aws://%s/%s", *zone, *attachement.VolumeId)
	return devPath
}

// find nvme block name by symlink, if symlink doesn't exist, return ""
func findNvmeBlockNameIfPresent(devName string) string {
	// for nvme(ssd), there is a symlink from devName to nvme block name, i.e. /dev/xvda -> /dev/nvme0n1
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nvme-ebs-volumes.html
	hasRootFs := true
	if _, err := os.Lstat("/rootfs/proc"); os.IsNotExist(err) {
		hasRootFs = false
	}
	nvmeName := ""

	if hasRootFs {
		devName = "/rootfs" + devName
	}

	if info, err := os.Lstat(devName); err == nil {
		if info.Mode()&os.ModeSymlink != 0 {
			if path, err := filepath.EvalSymlinks(devName); err == nil {
				nvmeName = path
			}
		}
	}

	if nvmeName != "" && hasRootFs {
		nvmeName = strings.TrimPrefix(nvmeName, "/rootfs")
	}
	return nvmeName
}

func (e *EbsVolume) GetEbsVolumeId(devName string) string {
	e.RLock()
	defer e.RUnlock()

	for k, v := range e.dev2Vol {
		// The key of dev2Vol is device name like nvme0n1, while the input devName could be a partition name like nvme0n1p1
		if strings.HasPrefix(devName, k) {
			return v
		}
	}

	return ""
}

//extract the ebs volume id used by kubernetes cluster
func ExtractEbsIdsUsedByKubernetes() map[string]string {
	ebsVolumeIds := make(map[string]string)

	file, err := os.Open("/rootfs/proc/mounts")
	if err != nil {
		log.Printf("D! cannot open /rootfs/proc/mounts %v", err)
		return ebsVolumeIds
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	ebsMountPointRegex := regexp.MustCompile(`kubernetes.io/aws-ebs/mounts/aws/(.+)/(vol-\w+)`)

	for {
		line, isPrefix, err := reader.ReadLine()

		// err could be EOF in normal case
		if err != nil {
			break
		}

		// isPrefix is set when a line exceeding 4KB which we treat it as error when reading mount file
		if isPrefix {
			break
		}

		lineStr := string(line)
		if strings.TrimSpace(lineStr) == "" {
			continue
		}

		//example line: /dev/nvme1n1 /var/lib/kubelet/plugins/kubernetes.io/aws-ebs/mounts/aws/us-west-2b/vol-0d9f0816149eb2050 ext4 rw,relatime,data=ordered 0 0
		keys := strings.Split(lineStr, " ")
		if len(keys) < 2 {
			continue
		}
		matches := ebsMountPointRegex.FindStringSubmatch(keys[1])
		if len(matches) > 0 {
			// Set {"/dev/nvme1n1": "aws://us-west-2b/vol-0d9f0816149eb2050"}
			ebsVolumeIds[keys[0]] = fmt.Sprintf("aws://%s/%s", matches[1], matches[2])
		}
	}

	return ebsVolumeIds
}

func hostJitter(max time.Duration) time.Duration {
	hostName := os.Getenv("HOST_NAME")
	if hostName == "" {
		hostName = "Unknown"
	}
	hash := fnv.New64()
	hash.Write([]byte(hostName))
	// Right shift the uint64 hash by one to make sure the jitter duration is always positive
	hostSleepJitter := time.Duration(int64(hash.Sum64()>>1)) % max
	return hostSleepJitter
}
