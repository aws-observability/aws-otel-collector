package ebs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const Type = "ebs"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

// delete ebs volumes that are not actively attached to an instance and are older than the expiration date.
func Clean(sess *session.Session, expirationDate time.Time) error {
	ec2client := ec2.New(sess)

	var nextToken *string
	for {
		describeVolumesInput := ec2.DescribeVolumesInput{MaxResults: aws.Int64(100), NextToken: nextToken}
		describeVolumesOutput, err := ec2client.DescribeVolumes(&describeVolumesInput)

		if err != nil {
			return err
		}

		for _, volume := range describeVolumesOutput.Volumes {
			if expirationDate.After(*volume.CreateTime) && len(volume.Attachments) == 0 {
				logger.Printf("Try to delete volume %s launch-date %v", *volume.VolumeId, volume.CreateTime)
				deleteVolumeInput := ec2.DeleteVolumeInput{VolumeId: volume.VolumeId}
				if _, err := ec2client.DeleteVolume(&deleteVolumeInput); err != nil {
					return err
				}
			}
		}
		if describeVolumesOutput.NextToken == nil {
			break
		}
		nextToken = describeVolumesOutput.NextToken
	}

	return nil
}
