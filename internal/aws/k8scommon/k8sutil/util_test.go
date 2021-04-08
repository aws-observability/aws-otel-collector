package k8sutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePodKey(t *testing.T) {
	assert.Equal(t, "namespace:default,podName:testPod", CreatePodKey("default", "testPod"))
}
