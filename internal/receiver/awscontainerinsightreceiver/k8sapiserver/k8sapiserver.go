package k8sapiserver

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	common "github.com/aws-observability/aws-otel-collector/internal/aws/containerinsightcommon"
	"github.com/aws-observability/aws-otel-collector/internal/aws/k8scommon/k8sclient"
	"go.opentelemetry.io/collector/consumer/pdata"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
)

const (
	lockName = "aoc-clusterleader"
)

type K8sAPIServer struct {
	nodeName string //get the value from downward API
	logger   *zap.Logger
	cancel   context.CancelFunc
	leading  bool
}

func New(logger *zap.Logger) *K8sAPIServer {
	k := &K8sAPIServer{
		logger: logger,
	}

	if err := k.start(); err != nil {
		k.logger.Warn("Fail to start k8sapiserver", zap.Error(err))
		return nil
	}

	return k
}

func (k *K8sAPIServer) GetMetrics() []pdata.Metrics {
	var result []pdata.Metrics
	if k.leading {
		k.logger.Info("collect data from K8s API Server...")
		timestampNs := strconv.FormatInt(time.Now().UnixNano(), 10)
		client := k8sclient.Get()

		fields := map[string]interface{}{
			"cluster_failed_node_count": client.Node.ClusterFailedNodeCount(),
			"cluster_node_count":        client.Node.ClusterNodeCount(),
		}
		attributes := map[string]string{
			common.MetricType: common.TypeCluster,
			common.Timestamp:  timestampNs,
			common.Version:    "0",
		}
		if k.nodeName != "" {
			attributes["NodeName"] = k.nodeName
		}
		md := common.ConvertToOTLPMetrics(fields, attributes, k.logger)
		result = append(result, md)

		for service, podNum := range client.Ep.ServiceToPodNum() {
			fields := map[string]interface{}{
				"service_number_of_running_pods": podNum,
			}
			attributes := map[string]string{
				common.MetricType:   common.TypeClusterService,
				common.Timestamp:    timestampNs,
				common.TypeService:  service.ServiceName,
				common.K8sNamespace: service.Namespace,
				common.Version:      "0",
			}
			if k.nodeName != "" {
				attributes["NodeName"] = k.nodeName
			}
			md := common.ConvertToOTLPMetrics(fields, attributes, k.logger)
			result = append(result, md)
		}

		for namespace, podNum := range client.Pod.NamespaceToRunningPodNum() {
			fields := map[string]interface{}{
				"namespace_number_of_running_pods": podNum,
			}
			attributes := map[string]string{
				common.MetricType:   common.TypeClusterNamespace,
				common.Timestamp:    timestampNs,
				common.K8sNamespace: namespace,
				common.Version:      "0",
			}
			if k.nodeName != "" {
				attributes["NodeName"] = k.nodeName
			}
			md := common.ConvertToOTLPMetrics(fields, attributes, k.logger)
			result = append(result, md)
		}
	}
	return result
}

func (k *K8sAPIServer) start() error {
	var ctx context.Context
	ctx, k.cancel = context.WithCancel(context.Background())

	k.nodeName = os.Getenv("HOST_NAME")
	if k.nodeName == "" {
		return errors.New("missing environment variable HOST_NAME. Please check your YAML config")
	}

	lockNamespace := os.Getenv("K8S_NAMESPACE")
	if lockNamespace == "" {
		return errors.New("missing environment variable K8S_NAMESPACE. Please check your YAML config")
	}

	configMapInterface := k8sclient.Get().ClientSet.CoreV1().ConfigMaps(lockNamespace)
	if configMap, err := configMapInterface.Get(ctx, lockName, metav1.GetOptions{}); configMap == nil || err != nil {
		k.logger.Info(fmt.Sprintf("Cannot get the leader config map: %v, try to create the config map...", err))
		configMap, err = configMapInterface.Create(ctx,
			&v1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: lockNamespace,
					Name:      lockName,
				},
			}, metav1.CreateOptions{})
		k.logger.Info(fmt.Sprintf("configMap: %v, err: %v", configMap, err))
	}

	lock, err := resourcelock.New(
		resourcelock.ConfigMapsResourceLock,
		lockNamespace, lockName,
		k8sclient.Get().ClientSet.CoreV1(),
		k8sclient.Get().ClientSet.CoordinationV1(),
		resourcelock.ResourceLockConfig{
			Identity:      k.nodeName,
			EventRecorder: createRecorder(k8sclient.Get().ClientSet, lockName, lockNamespace),
		})
	if err != nil {
		k.logger.Warn("Failed to create resource lock", zap.Error(err))
		return err
	}

	go k.startLeaderElection(ctx, lock)

	return nil
}

func (k *K8sAPIServer) startLeaderElection(ctx context.Context, lock resourcelock.Interface) {

	for {
		leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
			Lock: lock,
			// IMPORTANT: you MUST ensure that any code you have that
			// is protected by the lease must terminate **before**
			// you call cancel. Otherwise, you could have a background
			// loop still running and another process could
			// get elected before your background loop finished, violating
			// the stated goal of the lease.
			LeaseDuration: 60 * time.Second,
			RenewDeadline: 15 * time.Second,
			RetryPeriod:   5 * time.Second,
			Callbacks: leaderelection.LeaderCallbacks{
				OnStartedLeading: func(ctx context.Context) {
					k.logger.Info(fmt.Sprintf("k8sapiserver OnStartedLeading: %s", k.nodeName))
					// we're notified when we start
					k.leading = true
				},
				OnStoppedLeading: func() {
					k.logger.Info(fmt.Sprintf("k8sapiserver OnStoppedLeading: %s", k.nodeName))
					// we can do cleanup here, or after the RunOrDie method returns
					k.leading = false
					//node and pod are only used for cluster level metrics, endpoint is used for decorator too.
					k8sclient.Get().Node.Shutdown()
					k8sclient.Get().Pod.Shutdown()
				},
				OnNewLeader: func(identity string) {
					k.logger.Info(fmt.Sprintf("k8sapiserver Switch New Leader: %s", identity))
				},
			},
		})

		select {
		case <-ctx.Done(): //when leader election ends, the channel ctx.Done() will be closed
			k.logger.Info(fmt.Sprintf("k8sapiserver shutdown Leader Election: %s", k.nodeName))
			return
		default:
		}
	}
}

func (k *K8sAPIServer) Stop() {
	if k.cancel != nil {
		k.cancel()
	}
}

func createRecorder(clientSet kubernetes.Interface, name, namespace string) record.EventRecorder {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(klog.Infof)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: v1core.New(clientSet.CoreV1().RESTClient()).Events(namespace)})
	return eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: name})
}
