package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTaskCgroupPathFromARN(t *testing.T) {
	oldFormatARN := "arn:aws:ecs:region:aws_account_id:task/task-id"
	newFormatARN := "arn:aws:ecs:region:aws_account_id:task/cluster-name/task-id"
	result, _ := getTaskCgroupPathFromARN(oldFormatARN)
	assert.Equal(t, "task-id", result, "Expected to be equal")
	result, _ = getTaskCgroupPathFromARN(newFormatARN)
	assert.Equal(t, "task-id", result, "Expected to be equal")
}
