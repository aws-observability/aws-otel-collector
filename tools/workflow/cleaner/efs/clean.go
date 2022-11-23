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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"
)

const Type = "efs"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(sess *session.Session, expirationDate time.Time) error {
	logger.Printf("Begin to clean EFS resources")
	efsclient := efs.New(sess)

	//get efs to delete
	var nextToken *string
	for {
		describeFileSystemsInput := efs.DescribeFileSystemsInput{Marker: nextToken}
		describeFileSystemsOutput, err := efsclient.DescribeFileSystems(&describeFileSystemsInput)
		if err != nil {
			return err
		}
		for _, fileSystem := range describeFileSystemsOutput.FileSystems {
			if expirationDate.After(*fileSystem.CreationTime) {
				logger.Printf("Trying to delete file system %s launch-date %v", *fileSystem.FileSystemId, fileSystem.CreationTime)
				if *fileSystem.NumberOfMountTargets > 0 {
					err = deleteMountTargets(efsclient, fileSystem.FileSystemId)
				}

				if err == nil {
					terminateFileSystemsInput := efs.DeleteFileSystemInput{FileSystemId: fileSystem.FileSystemId}
					if _, err = efsclient.DeleteFileSystem(&terminateFileSystemsInput); err != nil {
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

func deleteMountTargets(client *efs.EFS, fileSystemId *string) error {
	var marker *string
	for {
		dmti := &efs.DescribeMountTargetsInput{Marker: marker, FileSystemId: fileSystemId}
		dmto, err := client.DescribeMountTargets(dmti)
		if err != nil {
			return err
		}
		for _, mountTarget := range dmto.MountTargets {
			dlmti := &efs.DeleteMountTargetInput{MountTargetId: mountTarget.MountTargetId}
			if _, err = client.DeleteMountTarget(dlmti); err != nil {
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
