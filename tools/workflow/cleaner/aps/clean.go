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

package aps

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/prometheusservice"
)

const Type = "aps"

var logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", Type), log.LstdFlags)

func Clean(sess *session.Session, expirationDate time.Time) error {
	logger.Printf("Begin to clean AMP workspaces")

	apsclient := prometheusservice.New(sess)

	//get Workspaces to delete
	var deleteWorkspaceIds []*string
	var nextToken *string
	for {
		listWorkspacesInput := &prometheusservice.ListWorkspacesInput{NextToken: nextToken}
		listWorkspacesOutput, err := apsclient.ListWorkspaces(listWorkspacesInput)

		if err != nil {
			return fmt.Errorf("unable to list workspaces: %w", err)
		}

		for _, ws := range listWorkspacesOutput.Workspaces {
			if expirationDate.After(*ws.CreatedAt) && shouldDelete(ws) {
				logger.Printf("Try to delete workspace %s tags %v created-at %v", *ws.WorkspaceId, ws.Tags, ws.CreatedAt)
				deleteWorkspaceIds = append(deleteWorkspaceIds, ws.WorkspaceId)
			}
		}

		if listWorkspacesOutput.NextToken == nil {
			break
		}
		nextToken = listWorkspacesOutput.NextToken
	}

	if len(deleteWorkspaceIds) < 1 {
		logger.Printf("No Workspaces to delete")
		return nil
	}

	for _, id := range deleteWorkspaceIds {
		terminateWorkspacesInput := &prometheusservice.DeleteWorkspaceInput{WorkspaceId: id}
		if _, err := apsclient.DeleteWorkspace(terminateWorkspacesInput); err != nil {
			return fmt.Errorf("unable to delete workspace: %w", err)
		}
	}
	return nil
}

func shouldDelete(ws *prometheusservice.WorkspaceSummary) bool {
	et, ok := ws.Tags["ephemeral"]
	if ok {
		return *et == "true"
	}
	return len(ws.Tags) == 0 && (ws.Alias == nil || *ws.Alias == "")
}
