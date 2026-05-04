package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const vkePath = "/v2/kubernetes/clusters"

// KubernetesService is the interface to interact with kubernetes endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/kubernetes
type KubernetesService interface {
	CreateCluster(ctx context.Context, createReq *ClusterReq) (*Cluster, *http.Response, error)
	GetCluster(ctx context.Context, id string) (*Cluster, *http.Response, error)
	ListClusters(ctx context.Context, options *ListOptions) ([]Cluster, *Meta, *http.Response, error)
	UpdateCluster(ctx context.Context, vkeID string, updateReq *ClusterReqUpdate) error
	DeleteCluster(ctx context.Context, id string) error
	DeleteClusterWithResources(ctx context.Context, id string) error

	CreateNodePool(ctx context.Context, vkeID string, nodePoolReq *NodePoolReq) (*NodePool, *http.Response, error)
	ListNodePools(ctx context.Context, vkeID string, options *ListOptions) ([]NodePool, *Meta, *http.Response, error)
	GetNodePool(ctx context.Context, vkeID, nodePoolID string) (*NodePool, *http.Response, error)
	UpdateNodePool(ctx context.Context, vkeID, nodePoolID string, updateReq *NodePoolReqUpdate) (*NodePool, *http.Response, error)
	DeleteNodePool(ctx context.Context, vkeID, nodePoolID string) error

	ListNodePoolLabels(ctx context.Context, vkeID, nodePoolID string) ([]NodePoolLabel, *http.Response, error)
	CreateNodePoolLabel(ctx context.Context, vkeID string, nodePoolID string, nodePoolLabelReq *NodePoolLabelReq) (*NodePoolLabel, *http.Response, error) //nolint:lll
	GetNodePoolLabel(ctx context.Context, vkeID string, nodePoolID string, nodePoolLabelID string) (*NodePoolLabel, *http.Response, error)
	DeleteNodePoolLabel(ctx context.Context, vkeID string, nodePoolID string, nodePoolLabelID string) error

	ListNodePoolTaints(ctx context.Context, vkeID, nodePoolID string) ([]NodePoolTaint, *http.Response, error)
	CreateNodePoolTaint(ctx context.Context, vkeID string, nodePoolID string, nodePoolTaintReq *NodePoolTaintReq) (*NodePoolTaint, *http.Response, error) //nolint:lll
	GetNodePoolTaint(ctx context.Context, vkeID string, nodePoolID string, nodePoolTaintID string) (*NodePoolTaint, *http.Response, error)
	DeleteNodePoolTaint(ctx context.Context, vkeID string, nodePoolID string, nodePoolTaintID string) error

	ListWorkerNodes(ctx context.Context, vkeID, nodePoolID string, options *ListOptions) ([]Node, *Meta, *http.Response, error)
	DeleteNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error
	RecycleNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error

	GetKubeConfig(ctx context.Context, vkeID string) (*KubeConfig, *http.Response, error)
	GetVersions(ctx context.Context) (*Versions, *http.Response, error)

	GetUpgrades(ctx context.Context, vkeID string) ([]string, *http.Response, error)
	Upgrade(ctx context.Context, vkeID string, body *ClusterUpgradeReq) error
}

// KubernetesHandler handles interaction with the kubernetes methods for the Vultr API
type KubernetesHandler struct {
	client *Client
}

// Cluster represents a full VKE cluster
type Cluster struct {
	ID              string            `json:"id"`
	Label           string            `json:"label"`
	DateCreated     string            `json:"date_created"`
	ClusterSubnet   string            `json:"cluster_subnet"`
	ServiceSubnet   string            `json:"service_subnet"`
	IP              string            `json:"ip"`
	Endpoint        string            `json:"endpoint"`
	Version         string            `json:"version"`
	Region          string            `json:"region"`
	Status          string            `json:"status"`
	HAControlPlanes bool              `json:"ha_controlplanes"`
	FirewallGroupID string            `json:"firewall_group_id"`
	OIDCConfig      ClusterOIDCConfig `json:"oidc"`
	NodePools       []NodePool        `json:"node_pools"`
}

// NodePool represents a pool of nodes that are grouped by their label and plan type
type NodePool struct {
	ID           string            `json:"id"`
	DateCreated  string            `json:"date_created"`
	DateUpdated  string            `json:"date_updated"`
	Label        string            `json:"label"`
	Plan         string            `json:"plan"`
	Status       string            `json:"status"`
	NodeQuantity int               `json:"node_quantity"`
	MinNodes     int               `json:"min_nodes"`
	MaxNodes     int               `json:"max_nodes"`
	AutoScaler   bool              `json:"auto_scaler"`
	UserData     string            `json:"user_data"`
	Tag          string            `json:"tag"`
	Labels       map[string]string `json:"labels"`
	Taints       []Taint           `json:"taints"`
	Nodes        []Node            `json:"nodes"`
}

// Node represents a node that will live within a nodepool
type Node struct {
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Label       string `json:"label"`
	IP          string `json:"ip,omitempty"` // Optional, may not be present in older API responses
	Status      string `json:"status"`
}

// KubeConfig will contain the kubeconfig b64 encoded
type KubeConfig struct {
	KubeConfig string `json:"kube_config"`
}

// ClusterReq struct used to create a cluster
type ClusterReq struct {
	Label           string             `json:"label"`
	Region          string             `json:"region"`
	Version         string             `json:"version"`
	HAControlPlanes bool               `json:"ha_controlplanes,omitempty"`
	EnableFirewall  bool               `json:"enable_firewall,omitempty"`
	VPCID           string             `json:"vpc_id,omitempty"`
	OIDCConfig      *ClusterOIDCConfig `json:"oidc,omitempty"`
	NodePools       []NodePoolReq      `json:"node_pools"`
}

// ClusterReqUpdate struct used to update update a cluster
type ClusterReqUpdate struct {
	Label      string             `json:"label"`
	OIDCConfig *ClusterOIDCConfig `json:"oidc,omitempty"`
}

// NodePoolReq struct used to create a node pool
type NodePoolReq struct {
	NodeQuantity int               `json:"node_quantity"`
	Label        string            `json:"label"`
	Plan         string            `json:"plan"`
	Tag          string            `json:"tag"`
	MinNodes     int               `json:"min_nodes,omitempty"`
	MaxNodes     int               `json:"max_nodes,omitempty"`
	AutoScaler   *bool             `json:"auto_scaler"`
	Labels       map[string]string `json:"labels,omitempty"`
	Taints       []Taint           `json:"taints,omitempty"`
	UserData     string            `json:"user_data"`
}

// NodePoolReqUpdate struct used to update a node pool
type NodePoolReqUpdate struct {
	NodeQuantity int     `json:"node_quantity,omitempty"`
	Tag          *string `json:"tag,omitempty"`
	MinNodes     int     `json:"min_nodes,omitempty"`
	MaxNodes     int     `json:"max_nodes,omitempty"`
	AutoScaler   *bool   `json:"auto_scaler,omitempty"`

	// Deprecated: Use CreateNodePoolLabel, DeleteNodePoolLabel instead
	Labels map[string]string `json:"labels,omitempty"`

	// Deprecated: Use CreateNodePoolTaint, DeleteNodePoolTaint instead
	Taints []Taint `json:"taints,omitempty"`

	UserData *string `json:"user_data,omitempty"`
}

// NodePoolLabel struct used to define a NodePool Label
type NodePoolLabel struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NodePoolLabelReq struct used to create a NodePool Label
type NodePoolLabelReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// NodePoolTaint struct used to define a NodePool taint
type NodePoolTaint struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

// NodePoolTaintReq struct used to create a NodePool taint
type NodePoolTaintReq struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

// Taint represents a kubernetes taint that can be applied to nodes in a node pool
type Taint struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}

type vkeClustersBase struct {
	VKEClusters []Cluster `json:"vke_clusters"`
	Meta        *Meta     `json:"meta"`
}

type vkeClusterBase struct {
	VKECluster *Cluster `json:"vke_cluster"`
}

type vkeNodePoolsBase struct {
	NodePools []NodePool `json:"node_pools"`
	Meta      *Meta      `json:"meta"`
}

type vkeNodePoolLabelsBase struct {
	Labels []NodePoolLabel `json:"labels"`
}

type vkeNodePoolLabelBase struct {
	Label *NodePoolLabel `json:"label"`
}

type vkeNodePoolTaintsBase struct {
	Taints []NodePoolTaint `json:"taints"`
}

type vkeNodePoolTaintBase struct {
	Taint *NodePoolTaint `json:"taint"`
}

type vkeWorkerNodesBase struct {
	WorkerNodes []Node `json:"worker_nodes"`
	Meta        *Meta  `json:"meta"`
}

type vkeNodePoolBase struct {
	NodePool *NodePool `json:"node_pool"`
}

// ClusterOIDCConfig represents the OIDC config used in the kubernetes cluster
type ClusterOIDCConfig struct {
	IssuerURL     string `json:"issuer_url"`
	ClientID      string `json:"client_id"`
	UserNameClaim string `json:"username_claim"`
	GroupsClaim   string `json:"groups_claim"`
}

// Versions that are supported for VKE
type Versions struct {
	Versions []string `json:"versions"`
}

// AvailableUpgrades for a given VKE cluster
type availableUpgrades struct {
	AvailableUpgrades []string `json:"available_upgrades"`
}

// ClusterUpgradeReq struct for vke upgradse
type ClusterUpgradeReq struct {
	UpgradeVersion string `json:"upgrade_version,omitempty"`
}

// CreateCluster will create a Kubernetes cluster.
func (k *KubernetesHandler) CreateCluster(ctx context.Context, createReq *ClusterReq) (*Cluster, *http.Response, error) {
	req, err := k.client.NewRequest(ctx, http.MethodPost, vkePath, createReq)
	if err != nil {
		return nil, nil, err
	}

	var k8 = new(vkeClusterBase)
	resp, err := k.client.DoWithContext(ctx, req, &k8)
	if err != nil {
		return nil, resp, err
	}

	return k8.VKECluster, resp, nil
}

// GetCluster will return a Kubernetes cluster.
func (k *KubernetesHandler) GetCluster(ctx context.Context, id string) (*Cluster, *http.Response, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", vkePath, id), nil)
	if err != nil {
		return nil, nil, err
	}

	k8 := new(vkeClusterBase)
	resp, err := k.client.DoWithContext(ctx, req, &k8)
	if err != nil {
		return nil, resp, err
	}

	return k8.VKECluster, resp, nil
}

// ListClusters will return all kubernetes clusters.
func (k *KubernetesHandler) ListClusters(ctx context.Context, options *ListOptions) ([]Cluster, *Meta, *http.Response, error) { //nolint:dupl,lll
	req, err := k.client.NewRequest(ctx, http.MethodGet, vkePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	k8s := new(vkeClustersBase)
	resp, err := k.client.DoWithContext(ctx, req, &k8s)
	if err != nil {
		return nil, nil, resp, err
	}

	return k8s.VKEClusters, k8s.Meta, resp, nil
}

// UpdateCluster updates label on VKE cluster
func (k *KubernetesHandler) UpdateCluster(ctx context.Context, vkeID string, updateReq *ClusterReqUpdate) error {
	req, err := k.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("%s/%s", vkePath, vkeID), updateReq)
	if err != nil {
		return err
	}

	_, err = k.client.DoWithContext(ctx, req, nil)
	return err
}

// DeleteCluster will delete a Kubernetes cluster.
func (k *KubernetesHandler) DeleteCluster(ctx context.Context, id string) error {
	req, err := k.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", vkePath, id), nil)
	if err != nil {
		return err
	}
	_, err = k.client.DoWithContext(ctx, req, nil)
	return err
}

// DeleteClusterWithResources will delete a Kubernetes cluster and all related resources.
func (k *KubernetesHandler) DeleteClusterWithResources(ctx context.Context, id string) error {
	req, err := k.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s/delete-with-linked-resources", vkePath, id), nil)
	if err != nil {
		return err
	}
	_, err = k.client.DoWithContext(ctx, req, nil)
	return err
}

// CreateNodePool creates a nodepool on a VKE cluster
func (k *KubernetesHandler) CreateNodePool(ctx context.Context, vkeID string, nodePoolReq *NodePoolReq) (*NodePool, *http.Response, error) {
	req, err := k.client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s/node-pools", vkePath, vkeID), nodePoolReq)
	if err != nil {
		return nil, nil, err
	}

	n := new(vkeNodePoolBase)
	resp, err := k.client.DoWithContext(ctx, req, n)
	if err != nil {
		return nil, resp, err
	}

	return n.NodePool, resp, nil
}

// ListNodePools will return all nodepools for a given VKE cluster
func (k *KubernetesHandler) ListNodePools(ctx context.Context, vkeID string, options *ListOptions) ([]NodePool, *Meta, *http.Response, error) { //nolint:lll,dupl
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/node-pools", vkePath, vkeID), nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	n := new(vkeNodePoolsBase)
	resp, err := k.client.DoWithContext(ctx, req, &n)
	if err != nil {
		return nil, nil, resp, err
	}

	return n.NodePools, n.Meta, resp, nil
}

// GetNodePool will return a single nodepool
func (k *KubernetesHandler) GetNodePool(ctx context.Context, vkeID, nodePoolID string) (*NodePool, *http.Response, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/node-pools/%s", vkePath, vkeID, nodePoolID), nil)
	if err != nil {
		return nil, nil, err
	}

	n := new(vkeNodePoolBase)
	resp, err := k.client.DoWithContext(ctx, req, &n)
	if err != nil {
		return nil, resp, err
	}

	return n.NodePool, resp, nil
}

// UpdateNodePool will allow you change the quantity of nodes within a nodepool
func (k *KubernetesHandler) UpdateNodePool(ctx context.Context, vkeID, nodePoolID string, updateReq *NodePoolReqUpdate) (*NodePool, *http.Response, error) { //nolint:lll
	req, err := k.client.NewRequest(ctx, http.MethodPatch, fmt.Sprintf("%s/%s/node-pools/%s", vkePath, vkeID, nodePoolID), updateReq)
	if err != nil {
		return nil, nil, err
	}

	np := new(vkeNodePoolBase)
	resp, err := k.client.DoWithContext(ctx, req, np)
	if err != nil {
		return nil, resp, err
	}

	return np.NodePool, resp, nil
}

// DeleteNodePool will remove a nodepool from a VKE cluster
func (k *KubernetesHandler) DeleteNodePool(ctx context.Context, vkeID, nodePoolID string) error {
	req, err := k.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s/node-pools/%s", vkePath, vkeID, nodePoolID), nil)
	if err != nil {
		return err
	}

	_, err = k.client.DoWithContext(ctx, req, nil)
	return err
}

// ListNodePoolLabels retrieves a list of labels from a node pool
func (k *KubernetesHandler) ListNodePoolLabels(ctx context.Context, vkeID, nodePoolID string) ([]NodePoolLabel, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/node-pools/%s/labels", vkePath, vkeID, nodePoolID)

	req, err := k.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	labels := new(vkeNodePoolLabelsBase)
	resp, err := k.client.DoWithContext(ctx, req, labels)
	if err != nil {
		return nil, resp, err
	}

	return labels.Labels, resp, nil
}

// CreateNodePoolLabel creates a label on a node pool
func (k *KubernetesHandler) CreateNodePoolLabel(ctx context.Context, vkeID, nodePoolID string, nodePoolLabelReq *NodePoolLabelReq) (*NodePoolLabel, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/node-pools/%s/labels", vkePath, vkeID, nodePoolID)

	req, err := k.client.NewRequest(ctx, http.MethodPost, uri, nodePoolLabelReq)
	if err != nil {
		return nil, nil, err
	}

	label := new(vkeNodePoolLabelBase)
	resp, err := k.client.DoWithContext(ctx, req, label)
	if err != nil {
		return nil, resp, err
	}

	return label.Label, resp, nil
}

// GetNodePoolLabel retrieves a label from a node pool
func (k *KubernetesHandler) GetNodePoolLabel(ctx context.Context, vkeID, nodePoolID, nodePoolLabelID string) (*NodePoolLabel, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/node-pools/%s/labels/%s", vkePath, vkeID, nodePoolID, nodePoolLabelID)

	req, err := k.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	label := new(vkeNodePoolLabelBase)
	resp, err := k.client.DoWithContext(ctx, req, label)
	if err != nil {
		return nil, resp, err
	}

	return label.Label, resp, nil
}

// DeleteNodePoolLabel deletes a label from a node pool
func (k *KubernetesHandler) DeleteNodePoolLabel(ctx context.Context, vkeID, nodePoolID, nodePoolLabelID string) error {
	uri := fmt.Sprintf("%s/%s/node-pools/%s/labels/%s", vkePath, vkeID, nodePoolID, nodePoolLabelID)

	req, err := k.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = k.client.DoWithContext(ctx, req, nil)
	return err
}

// ListNodePoolTaints retrieves a list of taints from a node pool
func (k *KubernetesHandler) ListNodePoolTaints(ctx context.Context, vkeID, nodePoolID string) ([]NodePoolTaint, *http.Response, error) {
	uri := fmt.Sprintf("%s/%s/node-pools/%s/taints", vkePath, vkeID, nodePoolID)

	req, err := k.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	taints := new(vkeNodePoolTaintsBase)
	resp, err := k.client.DoWithContext(ctx, req, taints)
	if err != nil {
		return nil, resp, err
	}

	return taints.Taints, resp, nil
}

// CreateNodePoolTaint creates a taint on a node pool
func (k *KubernetesHandler) CreateNodePoolTaint(ctx context.Context, vkeID, nodePoolID string, nodePoolTaintReq *NodePoolTaintReq) (*NodePoolTaint, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/node-pools/%s/taints", vkePath, vkeID, nodePoolID)

	req, err := k.client.NewRequest(ctx, http.MethodPost, uri, nodePoolTaintReq)
	if err != nil {
		return nil, nil, err
	}

	taint := new(vkeNodePoolTaintBase)
	resp, err := k.client.DoWithContext(ctx, req, taint)
	if err != nil {
		return nil, resp, err
	}

	return taint.Taint, resp, nil
}

// GetNodePoolTaint retrieves a taint from a node pool
func (k *KubernetesHandler) GetNodePoolTaint(ctx context.Context, vkeID, nodePoolID, nodePoolTaintID string) (*NodePoolTaint, *http.Response, error) { //nolint:lll
	uri := fmt.Sprintf("%s/%s/node-pools/%s/taints/%s", vkePath, vkeID, nodePoolID, nodePoolTaintID)
	req, err := k.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	taint := new(vkeNodePoolTaintBase)
	resp, err := k.client.DoWithContext(ctx, req, taint)
	if err != nil {
		return nil, resp, err
	}

	return taint.Taint, resp, nil
}

// DeleteNodePoolTaint deletes a taint on a node pool
func (k *KubernetesHandler) DeleteNodePoolTaint(ctx context.Context, vkeID, nodePoolID, nodePoolTaintID string) error {
	uri := fmt.Sprintf("%s/%s/node-pools/%s/taints/%s", vkePath, vkeID, nodePoolID, nodePoolTaintID)

	req, err := k.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	_, err = k.client.DoWithContext(ctx, req, nil)
	return err
}

// ListWorkerNodes retrieves a list of all worker nodes for a given node pool
func (k *KubernetesHandler) ListWorkerNodes(ctx context.Context, vkeID, nodePoolID string, options *ListOptions) ([]Node, *Meta, *http.Response, error) { //nolint:lll,dupl
	uri := fmt.Sprintf("%s/%s/node-pools/%s/nodes", vkePath, vkeID, nodePoolID)

	req, err := k.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	nodes := new(vkeWorkerNodesBase)
	resp, err := k.client.DoWithContext(ctx, req, &nodes)
	if err != nil {
		return nil, nil, resp, err
	}

	return nodes.WorkerNodes, nodes.Meta, resp, nil
}

// DeleteNodePoolInstance will remove a specified node from a nodepool
func (k *KubernetesHandler) DeleteNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error {
	req, err := k.client.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/%s/node-pools/%s/nodes/%s", vkePath, vkeID, nodePoolID, nodeID),
		nil,
	)
	if err != nil {
		return err
	}

	_, err = k.client.DoWithContext(ctx, req, nil)
	return err
}

// RecycleNodePoolInstance will recycle (destroy + redeploy) a given node on a nodepool
func (k *KubernetesHandler) RecycleNodePoolInstance(ctx context.Context, vkeID, nodePoolID, nodeID string) error {
	req, err := k.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/%s/node-pools/%s/nodes/%s/recycle", vkePath, vkeID, nodePoolID, nodeID),
		nil,
	)
	if err != nil {
		return err
	}

	_, err = k.client.DoWithContext(ctx, req, nil)
	return err
}

// GetKubeConfig returns the kubeconfig for the specified VKE cluster
func (k *KubernetesHandler) GetKubeConfig(ctx context.Context, vkeID string) (*KubeConfig, *http.Response, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/config", vkePath, vkeID), nil)
	if err != nil {
		return nil, nil, err
	}

	kc := new(KubeConfig)
	resp, err := k.client.DoWithContext(ctx, req, &kc)
	if err != nil {
		return nil, resp, err
	}

	return kc, resp, nil
}

// GetVersions returns the supported kubernetes versions
func (k *KubernetesHandler) GetVersions(ctx context.Context) (*Versions, *http.Response, error) {
	uri := "/v2/kubernetes/versions"
	req, err := k.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	versions := new(Versions)
	resp, err := k.client.DoWithContext(ctx, req, &versions)
	if err != nil {
		return nil, resp, err
	}

	return versions, resp, nil
}

// GetUpgrades returns all version a VKE cluster can upgrade to
func (k *KubernetesHandler) GetUpgrades(ctx context.Context, vkeID string) ([]string, *http.Response, error) {
	req, err := k.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/available-upgrades", vkePath, vkeID), nil)
	if err != nil {
		return nil, nil, err
	}

	upgrades := new(availableUpgrades)
	resp, err := k.client.DoWithContext(ctx, req, &upgrades)
	if err != nil {
		return nil, resp, err
	}

	return upgrades.AvailableUpgrades, resp, nil
}

// Upgrade beings a VKE cluster upgrade
func (k *KubernetesHandler) Upgrade(ctx context.Context, vkeID string, body *ClusterUpgradeReq) error {
	req, err := k.client.NewRequest(ctx, http.MethodPost, fmt.Sprintf("%s/%s/upgrades", vkePath, vkeID), body)
	if err != nil {
		return err
	}

	_, err = k.client.DoWithContext(ctx, req, nil)
	return err
}
