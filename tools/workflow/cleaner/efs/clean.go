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
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/efs"
)

const Type = "efs"

func Clean(sess *session.Session, keepDuration time.Duration) {
	efsclient := efs.New(sess)

	//get efs to delete
	var nextToken *string
	for {
		describeFileSystemsInput := efs.DescribeFileSystemsInput{Marker: nextToken}
		describeFileSystemsOutput, err := efsclient.DescribeFileSystems(&describeFileSystemsInput)

		if err != nil {
			log.Printf("Failed to get efs for error %v", err)
		}

		for _, fileSystem := range describeFileSystemsOutput.FileSystems {
			//only delete system if older than 5 days and not mounted
			if time.Now().UTC().Add(keepDuration).After(*fileSystem.CreationTime) && *fileSystem.NumberOfMountTargets == 0 {
				log.Printf("Trying to delete file system %v launch-date %v", *fileSystem.FileSystemId, fileSystem.CreationTime)
				terminateFileSystemsInput := efs.DeleteFileSystemInput{FileSystemId: fileSystem.FileSystemId}
				_, err = efsclient.DeleteFileSystem(&terminateFileSystemsInput)
				if err != nil {
					log.Printf("Failed to terminate file system %v because of %v", terminateFileSystemsInput, err)
				}
			}
		}
		if describeFileSystemsOutput.NextMarker == nil {
			break
		}
		nextToken = describeFileSystemsOutput.NextMarker
	}
}
