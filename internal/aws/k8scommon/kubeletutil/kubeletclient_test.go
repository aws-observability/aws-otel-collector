package kubeletutil

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Test Data
const (
	podJson = `
{
  "kind": "PodList",
  "apiVersion": "v1",
  "metadata": {

  },
  "items": [
    {
      "metadata": {
        "name": "cpu-limit",
        "namespace": "default",
        "ownerReferences": [
            {
                "apiVersion": "apps/v1",
                "blockOwnerDeletion": true,
                "controller": true,
                "kind": "DaemonSet",
                "name": "DaemonSetTest",
                "uid": "36779a62-4aca-11e9-977b-0672b6c6fc94"
            }
        ],
        "selfLink": "/api/v1/namespaces/default/pods/cpu-limit",
        "uid": "764d01e1-2a2f-11e9-95ea-0a695d7ce286",
        "resourceVersion": "5671573",
        "creationTimestamp": "2019-02-06T16:51:34Z",
        "labels": {
          "app": "hello_test"
        },
        "annotations": {
          "kubernetes.io/config.seen": "2019-02-19T00:06:56.109155665Z",
          "kubernetes.io/config.source": "api"
        }
      },
      "spec": {
        "volumes": [
          {
            "name": "default-token-tlgw7",
            "secret": {
              "secretName": "default-token-tlgw7",
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "ubuntu",
            "image": "ubuntu",
            "command": [
              "/bin/bash"
            ],
            "args": [
              "-c",
              "sleep 300000000"
            ],
            "resources": {
              "limits": {
                "cpu": "10m",
                "memory": "50Mi"
              },
              "requests": {
                "cpu": "10m",
                "memory": "50Mi"
              }
            },
            "volumeMounts": [
              {
                "name": "default-token-tlgw7",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "ip-192-168-67-127.us-west-2.compute.internal",
        "securityContext": {

        },
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0
      },
      "status": {
        "phase": "Running",
        "conditions": [
          {
            "type": "Initialized",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-02-06T16:51:34Z"
          },
          {
            "type": "Ready",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-02-06T16:51:43Z"
          },
          {
            "type": "ContainersReady",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": null
          },
          {
            "type": "PodScheduled",
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2019-02-06T16:51:34Z"
          }
        ],
        "hostIP": "192.168.67.127",
        "podIP": "192.168.76.93",
        "startTime": "2019-02-06T16:51:34Z",
        "containerStatuses": [
          {
            "name": "ubuntu",
            "state": {
              "running": {
                "startedAt": "2019-02-06T16:51:42Z"
              }
            },
            "lastState": {

            },
            "ready": true,
            "restartCount": 0,
            "image": "ubuntu:latest",
            "imageID": "docker-pullable://ubuntu@sha256:7a47ccc3bbe8a451b500d2b53104868b46d60ee8f5b35a24b41a86077c650210",
            "containerID": "docker://637631e2634ea92c0c1aa5d24734cfe794f09c57933026592c12acafbaf6972c"
          }
        ],
        "qosClass": "Guaranteed"
      }
    }
  ]
}`
)

type MockHttpRoundTripper struct {
	http.RoundTripper
	mock.Mock
}

func (m *MockHttpRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	args := m.Called()
	return args.Get(0).(*http.Response), nil
}

// Test
func TestKubeClient_ListPods(t *testing.T) {
	mockRoundTripper := new(MockHttpRoundTripper)
	mockRoundTripper.On("RoundTrip", mock.Anything).Return(&http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(strings.NewReader(podJson))})
	client := KubeClient{roundTripper: mockRoundTripper}
	pods, err := client.ListPods()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(pods))
}
