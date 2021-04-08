package stores

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils_parseDeploymentFromReplicaSet(t *testing.T) {
	assert.Equal(t, "", parseDeploymentFromReplicaSet("cloudwatch-agent"))
	assert.Equal(t, "cloudwatch-agent", parseDeploymentFromReplicaSet("cloudwatch-agent-42kcz"))
}

func TestUtils_parseCronJobFromJob(t *testing.T) {
	assert.Equal(t, "", parseCronJobFromJob("hello-123"))
	assert.Equal(t, "hello", parseCronJobFromJob("hello-1234567890"))
	assert.Equal(t, "", parseCronJobFromJob("hello-123456789a"))
}
