package process

// message.go is a stripped down version of the backend message processing
// with support for the same MessageVersion and MessageEncoding but with
// only a limited set of message types.

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
)

// MessageEncoding represents how messages will be encoded or decoded for
// over-the-wire transfer. Protobuf should be used for server-side messages
// (e.g. from collector <-> server) and JSON should be used for client-side.
type MessageEncoding uint8

// Message encoding constants.
const (
	MessageEncodingProtobuf     MessageEncoding = 0
	MessageEncodingJSON         MessageEncoding = 1
	MessageEncodingZstdPB       MessageEncoding = 2
	_                           MessageEncoding = 3 // This is unused
	MessageEncodingZstd1xPB     MessageEncoding = 4
	MessageEncodingZstdPBxNoCgo MessageEncoding = 5
)

// MessageVersion is the version of the message. It should always be the first
// byte in the encoded version.
type MessageVersion uint8

// Message versioning constants.
const (
	MessageV1 MessageVersion = 1
	MessageV2 MessageVersion = 2
	MessageV3 MessageVersion = 3
)

// MessageHeader is attached to all messages at the head of the message. Some
// fields are added in later versions so make sure you're only using fields that
// are available in the defined Version.
type MessageHeader struct {
	Version        MessageVersion
	Encoding       MessageEncoding
	Type           MessageType
	SubscriptionID uint8 // Unused in Agent
	OrgID          int32 // Unused in Agent
	Timestamp      int64
}

// MessageType is a string representing the type of a message.
type MessageType uint8

// Message type constants for MessageType.
// Note: Ordering may seem unusual, this is just to match the backend where there
// are additional types that aren't covered here.
const (
	TypeCollectorProc                    = 12
	TypeCollectorConnections             = 22
	TypeResCollector                     = 23
	TypeCollectorRealTime                = 27
	TypeCollectorContainer               = 39
	TypeCollectorContainerRealTime       = 40
	TypeCollectorPod                     = 41
	TypeCollectorReplicaSet              = 42
	TypeCollectorDeployment              = 43
	TypeCollectorService                 = 44
	TypeCollectorNode                    = 45
	TypeCollectorCluster                 = 46
	TypeCollectorJob                     = 47
	TypeCollectorCronJob                 = 48
	TypeCollectorDaemonSet               = 49
	TypeCollectorStatefulSet             = 50
	TypeCollectorPersistentVolume        = 51
	TypeCollectorPersistentVolumeClaim   = 52
	TypeCollectorProcDiscovery           = 53
	TypeCollectorRole                    = 54
	TypeCollectorRoleBinding             = 55
	TypeCollectorClusterRole             = 56
	TypeCollectorClusterRoleBinding      = 57
	TypeCollectorServiceAccount          = 58
	TypeCollectorIngress                 = 59
	TypeCollectorProcEvent               = 60
	TypeCollectorNamespace               = 61
	TypeCollectorManifest                = 80
	TypeCollectorManifestCRD             = 81
	TypeCollectorManifestCR              = 82
	TypeCollectorVerticalPodAutoscaler   = 83
	TypeCollectorHorizontalPodAutoscaler = 84
	TypeCollectorNetworkPolicy           = 85
	TypeCollectorLimitRange              = 86
	TypeCollectorStorageClass            = 87
	TypeCollectorPodDisruptionBudget     = 88
	TypeCollectorECSTask                 = 200
)

func (m MessageType) String() string {
	switch m {
	case TypeCollectorProc:
		return "process"
	case TypeCollectorConnections:
		return "network"
	case TypeCollectorRealTime:
		return "process-rt"
	case TypeCollectorContainer:
		return "container"
	case TypeCollectorContainerRealTime:
		return "container-rt"
	case TypeCollectorPod:
		return "pod"
	case TypeCollectorReplicaSet:
		return "replica-set"
	case TypeCollectorDeployment:
		return "deployment"
	case TypeCollectorService:
		return "service"
	case TypeCollectorNode:
		return "node"
	case TypeCollectorCluster:
		return "cluster"
	case TypeCollectorJob:
		return "job"
	case TypeCollectorCronJob:
		return "cron-job"
	case TypeCollectorDaemonSet:
		return "daemon-set"
	case TypeCollectorStatefulSet:
		return "stateful-set"
	case TypeCollectorPersistentVolume:
		return "persistent-volume"
	case TypeCollectorPersistentVolumeClaim:
		return "persistent-volume-claim"
	case TypeCollectorProcDiscovery:
		return "process-discovery"
	case TypeCollectorRole:
		return "role"
	case TypeCollectorRoleBinding:
		return "role-binding"
	case TypeCollectorClusterRole:
		return "cluster-role"
	case TypeCollectorClusterRoleBinding:
		return "cluster-role-binding"
	case TypeCollectorServiceAccount:
		return "service-account"
	case TypeCollectorIngress:
		return "ingress"
	case TypeCollectorProcEvent:
		return "process-event"
	case TypeCollectorNamespace:
		return "namespace"
	case TypeCollectorManifest:
		return "manifest"
	case TypeCollectorManifestCRD:
		return "manifest-crd"
	case TypeCollectorManifestCR:
		return "manifest-cr"
	case TypeCollectorVerticalPodAutoscaler:
		return "vertical-pod-autoscaler"
	case TypeCollectorHorizontalPodAutoscaler:
		return "horizontal-pod-autoscaler"
	case TypeCollectorNetworkPolicy:
		return "network-policy"
	case TypeCollectorLimitRange:
		return "limit-range"
	case TypeCollectorStorageClass:
		return "storage-class"
	case TypeCollectorECSTask:
		return "ecs-task"
	case TypeCollectorPodDisruptionBudget:
		return "pod-disruption-budget"
	default:
		// otherwise convert the type identifier
		return strconv.Itoa(int(m))
	}
}

// Message is a generic type for all messages with a Header and Body.
type Message struct {
	Header MessageHeader
	Body   MessageBody
}

// MessageBody is a common interface used by all message types.
type MessageBody interface {
	ProtoMessage()
	Reset()
	String() string
	Size() int
}

// DecodeMessage decodes raw message bytes into a specific type that satisfies
// the Message interface. If we can't decode, an error is returned.
func DecodeMessage(data []byte) (Message, error) {
	header, offset, err := ReadHeader(data)
	if err != nil {
		return Message{}, err
	}
	body := data[offset:]
	var m MessageBody
	switch header.Type {
	case TypeCollectorProc:
		m = &CollectorProc{}
	case TypeCollectorConnections:
		m = &CollectorConnections{}
	case TypeCollectorRealTime:
		m = &CollectorRealTime{}
	case TypeResCollector:
		m = &ResCollector{}
	case TypeCollectorContainer:
		m = &CollectorContainer{}
	case TypeCollectorContainerRealTime:
		m = &CollectorContainerRealTime{}
	case TypeCollectorPod:
		m = &CollectorPod{}
	case TypeCollectorReplicaSet:
		m = &CollectorReplicaSet{}
	case TypeCollectorDeployment:
		m = &CollectorDeployment{}
	case TypeCollectorService:
		m = &CollectorService{}
	case TypeCollectorNode:
		m = &CollectorNode{}
	case TypeCollectorCluster:
		m = &CollectorCluster{}
	case TypeCollectorJob:
		m = &CollectorJob{}
	case TypeCollectorCronJob:
		m = &CollectorCronJob{}
	case TypeCollectorDaemonSet:
		m = &CollectorDaemonSet{}
	case TypeCollectorStatefulSet:
		m = &CollectorStatefulSet{}
	case TypeCollectorPersistentVolume:
		m = &CollectorPersistentVolume{}
	case TypeCollectorPersistentVolumeClaim:
		m = &CollectorPersistentVolumeClaim{}
	case TypeCollectorProcDiscovery:
		m = &CollectorProcDiscovery{}
	case TypeCollectorRole:
		m = &CollectorRole{}
	case TypeCollectorRoleBinding:
		m = &CollectorRoleBinding{}
	case TypeCollectorClusterRole:
		m = &CollectorClusterRole{}
	case TypeCollectorClusterRoleBinding:
		m = &CollectorClusterRoleBinding{}
	case TypeCollectorServiceAccount:
		m = &CollectorServiceAccount{}
	case TypeCollectorIngress:
		m = &CollectorIngress{}
	case TypeCollectorProcEvent:
		m = &CollectorProcEvent{}
	case TypeCollectorNamespace:
		m = &CollectorNamespace{}
	case TypeCollectorManifest:
		m = &CollectorManifest{}
	case TypeCollectorManifestCRD:
		m = &CollectorManifestCRD{}
	case TypeCollectorManifestCR:
		m = &CollectorManifestCR{}
	case TypeCollectorVerticalPodAutoscaler:
		m = &CollectorVerticalPodAutoscaler{}
	case TypeCollectorHorizontalPodAutoscaler:
		m = &CollectorHorizontalPodAutoscaler{}
	case TypeCollectorNetworkPolicy:
		m = &CollectorNetworkPolicy{}
	case TypeCollectorLimitRange:
		m = &CollectorLimitRange{}
	case TypeCollectorStorageClass:
		m = &CollectorStorageClass{}
	case TypeCollectorECSTask:
		m = &CollectorECSTask{}
	case TypeCollectorPodDisruptionBudget:
		m = &CollectorPodDisruptionBudget{}
	default:
		return Message{}, fmt.Errorf("unhandled message type: %d", header.Type)
	}
	if err = unmarshal(header.Encoding, body, m); err != nil {
		return Message{}, err
	}
	return Message{header, m}, nil
}

// DetectMessageType returns the message type for the given MessageBody
func DetectMessageType(b MessageBody) (MessageType, error) {
	var t MessageType
	switch b.(type) {
	case *CollectorProc:
		t = TypeCollectorProc
	case *CollectorConnections:
		t = TypeCollectorConnections
	case *CollectorRealTime:
		t = TypeCollectorRealTime
	case *ResCollector:
		t = TypeResCollector
	case *CollectorContainer:
		t = TypeCollectorContainer
	case *CollectorContainerRealTime:
		t = TypeCollectorContainerRealTime
	case *CollectorPod:
		t = TypeCollectorPod
	case *CollectorReplicaSet:
		t = TypeCollectorReplicaSet
	case *CollectorDeployment:
		t = TypeCollectorDeployment
	case *CollectorService:
		t = TypeCollectorService
	case *CollectorNode:
		t = TypeCollectorNode
	case *CollectorManifest:
		t = TypeCollectorManifest
	case *CollectorManifestCRD:
		t = TypeCollectorManifestCRD
	case *CollectorManifestCR:
		t = TypeCollectorManifestCR
	case *CollectorCluster:
		t = TypeCollectorCluster
	case *CollectorJob:
		t = TypeCollectorJob
	case *CollectorCronJob:
		t = TypeCollectorCronJob
	case *CollectorDaemonSet:
		t = TypeCollectorDaemonSet
	case *CollectorStatefulSet:
		t = TypeCollectorStatefulSet
	case *CollectorPersistentVolume:
		t = TypeCollectorPersistentVolume
	case *CollectorPersistentVolumeClaim:
		t = TypeCollectorPersistentVolumeClaim
	case *CollectorProcDiscovery:
		t = TypeCollectorProcDiscovery
	case *CollectorRole:
		t = TypeCollectorRole
	case *CollectorRoleBinding:
		t = TypeCollectorRoleBinding
	case *CollectorClusterRole:
		t = TypeCollectorClusterRole
	case *CollectorClusterRoleBinding:
		t = TypeCollectorClusterRoleBinding
	case *CollectorServiceAccount:
		t = TypeCollectorServiceAccount
	case *CollectorIngress:
		t = TypeCollectorIngress
	case *CollectorProcEvent:
		t = TypeCollectorProcEvent
	case *CollectorNamespace:
		t = TypeCollectorNamespace
	case *CollectorVerticalPodAutoscaler:
		t = TypeCollectorVerticalPodAutoscaler
	case *CollectorHorizontalPodAutoscaler:
		t = TypeCollectorHorizontalPodAutoscaler
	case *CollectorNetworkPolicy:
		t = TypeCollectorNetworkPolicy
	case *CollectorLimitRange:
		t = TypeCollectorLimitRange
	case *CollectorStorageClass:
		t = TypeCollectorStorageClass
	case *CollectorECSTask:
		t = TypeCollectorECSTask
	case *CollectorPodDisruptionBudget:
		t = TypeCollectorPodDisruptionBudget

	default:
		return 0, fmt.Errorf("unknown message body type: %s", reflect.TypeOf(b))
	}
	return t, nil
}

// ReadHeader reads the header off raw message bytes.
func ReadHeader(data []byte) (MessageHeader, int, error) {
	if len(data) <= 4 {
		return MessageHeader{}, 0, fmt.Errorf("invalid message length: %d", len(data))
	}
	switch MessageVersion(uint8(data[0])) {
	case MessageV1:
		return readHeaderV1(data)
	case MessageV2:
		return readHeaderV2(data)
	case MessageV3:
		return readHeaderV3(data)
	default:
		return MessageHeader{}, 0, fmt.Errorf("invalid message version: %d", uint8(data[0]))
	}
}

func readHeaderV1(data []byte) (MessageHeader, int, error) {
	b := bytes.NewBuffer(data[1:])
	var msgEnc uint8
	if err := binary.Read(b, binary.LittleEndian, &msgEnc); err != nil {
		return MessageHeader{}, 0, err
	}
	var msgType uint8
	if err := binary.Read(b, binary.LittleEndian, &msgType); err != nil {
		return MessageHeader{}, 0, err
	}
	var subID uint8
	if err := binary.Read(b, binary.LittleEndian, &subID); err != nil {
		return MessageHeader{}, 0, err
	}
	return MessageHeader{
		Version:        MessageV1,
		Encoding:       MessageEncoding(msgEnc),
		Type:           MessageType(msgType),
		SubscriptionID: subID,
		OrgID:          0,
	}, 4, nil
}

func readHeaderV2(data []byte) (MessageHeader, int, error) {
	b := bytes.NewBuffer(data[1:])
	var msgEnc uint8
	if err := binary.Read(b, binary.LittleEndian, &msgEnc); err != nil {
		return MessageHeader{}, 0, err
	}
	var msgType uint8
	if err := binary.Read(b, binary.LittleEndian, &msgType); err != nil {
		return MessageHeader{}, 0, err
	}
	var subID uint8
	if err := binary.Read(b, binary.LittleEndian, &subID); err != nil {
		return MessageHeader{}, 0, err
	}
	var orgID int32
	if err := binary.Read(b, binary.LittleEndian, &orgID); err != nil {
		return MessageHeader{}, 0, err
	}
	return MessageHeader{
		Version:        MessageV2,
		Encoding:       MessageEncoding(msgEnc),
		Type:           MessageType(msgType),
		SubscriptionID: subID,
		OrgID:          orgID,
	}, 8, nil
}

func readHeaderV3(data []byte) (MessageHeader, int, error) {
	b := bytes.NewBuffer(data[1:])
	var msgEnc uint8
	if err := binary.Read(b, binary.LittleEndian, &msgEnc); err != nil {
		return MessageHeader{}, 0, err
	}
	var msgType uint8
	if err := binary.Read(b, binary.LittleEndian, &msgType); err != nil {
		return MessageHeader{}, 0, err
	}
	var subID uint8
	if err := binary.Read(b, binary.LittleEndian, &subID); err != nil {
		return MessageHeader{}, 0, err
	}
	var orgID int32
	if err := binary.Read(b, binary.LittleEndian, &orgID); err != nil {
		return MessageHeader{}, 0, err
	}
	var timestamp int64
	if err := binary.Read(b, binary.LittleEndian, &timestamp); err != nil {
		return MessageHeader{}, 0, err
	}
	return MessageHeader{
		Version:        MessageV3,
		Encoding:       MessageEncoding(msgEnc),
		Type:           MessageType(msgType),
		SubscriptionID: subID,
		OrgID:          orgID,
		Timestamp:      timestamp,
	}, 16, nil
}

func encodeHeader(h MessageHeader) ([]byte, error) {
	switch h.Version {
	case MessageV3:
		return encodeHeaderV3(h)
	default:
		return nil, fmt.Errorf("invalid message version: %d", h.Version)
	}
}

func encodeHeaderV3(h MessageHeader) ([]byte, error) {
	b := new(bytes.Buffer)
	err := binary.Write(b, binary.LittleEndian, uint8(h.Version))
	if err != nil {
		return nil, err
	}
	err = binary.Write(b, binary.LittleEndian, uint8(h.Encoding))
	if err != nil {
		return nil, err
	}
	err = binary.Write(b, binary.LittleEndian, uint8(h.Type))
	if err != nil {
		return nil, err
	}
	err = binary.Write(b, binary.LittleEndian, uint8(h.SubscriptionID))
	if err != nil {
		return nil, err
	}
	err = binary.Write(b, binary.LittleEndian, h.OrgID)
	if err != nil {
		return nil, err
	}
	err = binary.Write(b, binary.LittleEndian, h.Timestamp)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
