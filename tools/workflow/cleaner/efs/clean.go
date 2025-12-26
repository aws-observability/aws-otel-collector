/*
 * Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package efs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
)

const Type = "efs"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(ctx context.Context, cfg aws.Config, expirationDate time.Time) error {
	logger.Printf("Begin to clean EFS resources")
	efsclient := efs.NewFromConfig(cfg)

	//get efs to delete
	var nextToken *string
	for {
		describeFileSystemsOutput, err := efsclient.DescribeFileSystems(ctx, &efs.DescribeFileSystemsInput{Marker: nextToken})
		if err != nil {
			return err
		}
		for _, fileSystem := range describeFileSystemsOutput.FileSystems {
			if shouldDeleteFileSystem(fileSystem, expirationDate) {
				logger.Printf("Trying to delete file system %s launch-date %v", *fileSystem.FileSystemId, fileSystem.CreationTime)
				if fileSystem.NumberOfMountTargets > 0 {
					err = deleteMountTargets(ctx, efsclient, fileSystem.FileSystemId)
				}

				if err == nil {
					_, err = efsclient.DeleteFileSystem(ctx, &efs.DeleteFileSystemInput{FileSystemId: fileSystem.FileSystemId})
					if err != nil {
						logger.Printf("Unable to delete file system %s due to %v", *fileSystem.FileSystemId, err)
					} else {
						logger.Printf("Deleted file system %s successfully", *fileSystem.FileSystemId)
					}
				} else {
					logger.Printf("Unable to delete all the mount targets for %s due to %v", *fileSystem.FileSystemId, err)
				}
			}
		}
		if describeFileSystemsOutput.NextMarker == nil {
			break
		}
		nextToken = describeFileSystemsOutput.NextMarker
	}
	return nil
}

func deleteMountTargets(ctx context.Context, client *efs.Client, fileSystemId *string) error {
	var marker *string
	for {
		dmto, err := client.DescribeMountTargets(ctx, &efs.DescribeMountTargetsInput{Marker: marker, FileSystemId: fileSystemId})
		if err != nil {
			return err
		}
		for _, mountTarget := range dmto.MountTargets {
			_, err = client.DeleteMountTarget(ctx, &efs.DeleteMountTargetInput{MountTargetId: mountTarget.MountTargetId})
			if err != nil {
				return err
			}
			log.Printf("Deleted mount target %s for %s successfully", *mountTarget.MountTargetId, *fileSystemId)
		}
		if dmto.Marker == nil {
			break
		}
		marker = dmto.Marker
	}
	return nil
}

// shouldDeleteFileSystem returns true if the file system is older than the expiration date.
func shouldDeleteFileSystem(fs types.FileSystemDescription, expirationDate time.Time) bool {
	return fs.CreationTime != nil && expirationDate.After(*fs.CreationTime)
}
