package ebs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const Type = "ebs"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

// delete ebs volumes that are not actively attached to an instance and are older than the expiration date.
func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	ec2client := ec2.NewFromConfig(cfg)

	paginator := ec2.NewDescribeVolumesPaginator(ec2client, &ec2.DescribeVolumesInput{
		MaxResults: aws.Int32(100),
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		for _, volume := range output.Volumes {
			if shouldDeleteVolume(volume, expirationDate) {
				logger.Printf("Try to delete volume %s launch-date %v", *volume.VolumeId, volume.CreateTime)
				_, err := ec2client.DeleteVolume(ctx, &ec2.DeleteVolumeInput{VolumeId: volume.VolumeId})
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// shouldDeleteVolume returns true if the volume is unattached and older than the expiration date.
func shouldDeleteVolume(volume types.Volume, expirationDate time.Time) bool {
	return volume.CreateTime != nil && expirationDate.After(*volume.CreateTime) && len(volume.Attachments) == 0
}
