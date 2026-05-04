package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const vcrPath = "/v2/registry"
const vcrListPath = "/v2/registries"

// ContainerRegistryService is the interface to interact with the container
// registry endpoints on the Vultr API.  Link :
// https://www.vultr.com/api/#tag/Container-Registry
type ContainerRegistryService interface {
	Create(ctx context.Context, createReq *ContainerRegistryReq) (*ContainerRegistry, *http.Response, error)
	Get(ctx context.Context, vcrID string) (*ContainerRegistry, *http.Response, error)
	Update(ctx context.Context, vcrID string, updateReq *ContainerRegistryUpdateReq) (*ContainerRegistry, *http.Response, error)
	Delete(ctx context.Context, vcrID string) error
	List(ctx context.Context, options *ListOptions) ([]ContainerRegistry, *Meta, *http.Response, error)

	ListRepositories(ctx context.Context, vcrID string, options *ListOptions) ([]ContainerRegistryRepo, *Meta, *http.Response, error)
	GetRepository(ctx context.Context, vcrID, imageName string) (*ContainerRegistryRepo, *http.Response, error)
	UpdateRepository(ctx context.Context, vcrID, imageName string, updateReq *ContainerRegistryRepoUpdateReq) (*ContainerRegistryRepo, *http.Response, error) //nolint:lll
	DeleteRepository(ctx context.Context, vcrID, imageName string) error

	ListReplications(ctx context.Context, vcrID string, options *ListOptions) ([]ContainerRegistryReplication, *Meta, *http.Response, error)
	CreateReplication(ctx context.Context, vcrID, regionID string) (*ContainerRegistryReplication, *http.Response, error)
	GetReplication(ctx context.Context, vcrID, regionID string) (*ContainerRegistryReplication, *http.Response, error)
	DeleteReplication(ctx context.Context, vcrID, regionID string) error

	UpdateRetentionSchedule(ctx context.Context, vcrID, schedule string) (*ContainerRegistryRetentionSchedule, *http.Response, error)
	ExecuteRetention(ctx context.Context, vcrID string, dryRun bool) (*ContainerRegistryRetentionExecution, *http.Response, error)

	ListRetentionRules(ctx context.Context, vcrID string, options *ListOptions) ([]ContainerRegistryRetentionRule, *Meta, *http.Response, error) //nolint:lll
	GetRetentionRule(ctx context.Context, vcrID string, ruleID int) (*ContainerRegistryRetentionRule, *http.Response, error)
	CreateRetentionRule(ctx context.Context, vcrID string, ruleReq *ContainerRegistryRetentionRuleReq) (*ContainerRegistryRetentionRule, *http.Response, error) //nolint:lll
	UpdateRetentionRule(ctx context.Context, vcrID string, ruleID int, disabled bool) (*ContainerRegistryRetentionRule, *http.Response, error)
	DeleteRetentionRule(ctx context.Context, vcrID string, ruleID int) error

	ListRobots(ctx context.Context, vcrID string, options *ListOptions) ([]ContainerRegistryRobot, *Meta, *http.Response, error)
	GetRobot(ctx context.Context, vcrID, robotName string) (*ContainerRegistryRobot, *http.Response, error)
	UpdateRobot(ctx context.Context, vcrID, robotName string, updateReq *ContainerRegistryRobotReq) (*ContainerRegistryRobot, *http.Response, error) //nolint:lll
	DeleteRobot(ctx context.Context, vcrID, robotName string) error

	ListArtifacts(ctx context.Context, vcrID, imageName string, options *ListOptions) ([]ContainerRegistryArtifact, *Meta, *http.Response, error) //nolint:lll
	GetArtifact(ctx context.Context, vcrID, imageName, artifactDigest string) (*ContainerRegistryArtifact, *http.Response, error)
	DeleteArtifact(ctx context.Context, vcrID, imageName, artifactDigest string) error

	CreateDockerCredentials(ctx context.Context, vcrID string, createOptions *DockerCredentialsOpt) (*ContainerRegistryDockerCredentials, *http.Response, error) //nolint:lll

	ListRegions(ctx context.Context) ([]ContainerRegistryRegion, *Meta, *http.Response, error)
	ListPlans(ctx context.Context) (*ContainerRegistryPlans, *http.Response, error)
}

// ContainerRegistryServiceHandler handles interaction between the container
// registry service and the Vultr API.
type ContainerRegistryServiceHandler struct {
	client *Client
}

// ContainerRegistry represents a Vultr container registry subscription.
type ContainerRegistry struct {
	ID          string                    `json:"id"`
	Name        string                    `json:"name"`
	URN         string                    `json:"urn"`
	Storage     ContainerRegistryStorage  `json:"storage"`
	DateCreated string                    `json:"date_created"`
	Public      bool                      `json:"public"`
	RootUser    ContainerRegistryUser     `json:"root_user"`
	Metadata    ContainerRegistryMetadata `json:"metadata"`
}

type containerRegistries struct {
	ContainerRegistries []ContainerRegistry `json:"registries"`
	Meta                *Meta               `json:"meta"`
}

// ContainerRegistryStorage represents the storage usage and limit.
type ContainerRegistryStorage struct {
	Used    ContainerRegistryStorageCount `json:"used"`
	Allowed ContainerRegistryStorageCount `json:"allowed"`
}

// ContainerRegistryStorageCount represents the different storage usage counts.
type ContainerRegistryStorageCount struct {
	Bytes        float32 `json:"bytes"`
	MegaBytes    float32 `json:"mb"`
	GigaBytes    float32 `json:"gb"`
	TeraBytes    float32 `json:"tb"`
	DateModified string  `json:"updated_at"`
}

// ContainerRegistryUser contains the user data.
type ContainerRegistryUser struct {
	ID           int    `json:"id"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	Root         bool   `json:"root"`
	DateCreated  string `json:"added_at"`
	DateModified string `json:"updated_at"`
}

// ContainerRegistryMetadata contains the meta data for the registry.
type ContainerRegistryMetadata struct {
	Region       ContainerRegistryRegion       `json:"region"`
	Subscription ContainerRegistrySubscription `json:"subscription"`
}

// ContainerRegistrySubscription contains the subscription information for the
// registry.
type ContainerRegistrySubscription struct {
	Billing ContainerRegistrySubscriptionBilling `json:"billing"`
}

// ContainerRegistrySubscriptionBilling represents the subscription billing
// data on the registry.
type ContainerRegistrySubscriptionBilling struct {
	MonthlyPrice   float32 `json:"monthly_price"`
	PendingCharges float32 `json:"pending_charges"`
}

// ContainerRegistryReq represents the data used to create a registry.
type ContainerRegistryReq struct {
	Name   string `json:"name"`
	Public bool   `json:"public"`
	Region string `json:"region"`
	Plan   string `json:"plan"`
}

// ContainerRegistryUpdateReq represents the data used to update a registry.
type ContainerRegistryUpdateReq struct {
	Public *bool   `json:"public,omitempty"`
	Plan   *string `json:"plan,omitempty"`
}

// ContainerRegistryRepo represents the data of a registry repository.
type ContainerRegistryRepo struct {
	Name          string `json:"name"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	DateCreated   string `json:"added_at"`
	DateModified  string `json:"updated_at"`
	PullCount     int    `json:"pull_count"`
	ArtifactCount int    `json:"artifact_count"`
}

type containerRegistryRepos struct {
	Repositories []ContainerRegistryRepo `json:"repositories"`
	Meta         *Meta                   `json:"meta"`
}

// ContainerRegistryRepoUpdateReq is the data to update a registry repository.
type ContainerRegistryRepoUpdateReq struct {
	Description string `json:"description"`
}

// ContainerRegistryReplication represents a container registry replication.
type ContainerRegistryReplication struct {
	Region    string `json:"region"`
	Namespace string `json:"namespace"`
	URN       string `json:"urn"`
}

// ContainerRegistryReplicationReq represents a container registry replication.
// request
type ContainerRegistryReplicationReq struct {
	Region string `json:"region"`
}

type containerRegistryReplicationsBase struct {
	Replications []ContainerRegistryReplication `json:"replications"`
	Meta         *Meta                          `json:"meta"`
}

// ContainerRegistryRetentionSchedule represents a container registry retention
// schedule.
type ContainerRegistryRetentionSchedule struct {
	Schedule          string `json:"schedule"`
	NextScheduledTime string `json:"next_scheduled_time"`
}

// ContainerRegistryRetentionScheduleReq represents a container registry
// schedule request.
type ContainerRegistryRetentionScheduleReq struct {
	Cron string `json:"cron"`
}

// ContainerRegistryRetentionExecution represents a container registry
// execution.
type ContainerRegistryRetentionExecution struct {
	Start   string `json:"start_time"`
	End     string `json:"end_time"`
	Trigger string `json:"trigger"`
	DryRun  bool   `json:"dry_run"`
}

// ContainerRegistryRetentionExecutionReq represents a container registry
// execution request.
type ContainerRegistryRetentionExecutionReq struct {
	DryRun bool `json:"dry_run"`
}

// ContainerRegistryRetentionRule represents a container registry retention
// rule.
type ContainerRegistryRetentionRule struct {
	ID            int                                      `json:"id"`
	Disabled      bool                                     `json:"disabled"`
	Action        string                                   `json:"action"`
	Parameters    ContainerRegistryRetentionRuleParameter  `json:"params"`
	ScopeSelector ContainerRegistryScopeSelector           `json:"scope_selector"`
	TagSelectors  []ContainerRegistryRetentionRuleSelector `json:"tag_selectors"`
	Template      string                                   `json:"template"`
}

// ContainerRegistryRetentionRuleParameter represents a container registry rule
// parameter.
type ContainerRegistryRetentionRuleParameter struct {
	AdditionalProperty int `json:"additional_prop"`
}

// ContainerRegistryScopeSelector represents a container registry retention
// rule scope selector.
type ContainerRegistryScopeSelector struct {
	Repository []ContainerRegistryRetentionRuleSelector `json:"repository"`
}

// ContainerRegistryRetentionRuleSelector represents a container registry
// retention rule selector.
type ContainerRegistryRetentionRuleSelector struct {
	Decoration string `json:"decoration"`
	Kind       string `json:"kind"`
	Pattern    string `json:"pattern"`
}

type containerRegistryRetentionRulesBase struct {
	Rules []ContainerRegistryRetentionRule `json:"retention_rules"`
	Meta  *Meta                            `json:"meta"`
}

// ContainerRegistryRetentionRuleReq represents the options for creating a
// container registry retention rule.
type ContainerRegistryRetentionRuleReq struct {
	Type             string `json:"rule_type"`
	Count            int    `json:"count,omitempty"`
	RepositoryAction string `json:"repository_action"`
	RepositoryMatch  string `json:"repository_match"`
	TagAction        string `json:"tag_action"`
	TagMatch         string `json:"tag_match"`
	Untagged         bool   `json:"untagged"`
}

// ContainerRegistryRetentionRuleUpdateReq represents the options for updating
// a container registry retention rule.
type ContainerRegistryRetentionRuleUpdateReq struct {
	Disabled bool `json:"disabled"`
}

// ContainerRegistryRobot represents the details of a container registry robot.
type ContainerRegistryRobot struct {
	Name        string                             `json:"name"`
	Description string                             `json:"description"`
	Secret      string                             `json:"secret"`
	Disable     bool                               `json:"disable"`
	Duration    int                                `json:"duration"`
	Permissions []ContainerRegistryRobotPermission `json:"permissions"`
	DateCreated string                             `json:"creation_time"`
}

// ContainerRegistryRobotPermission represent container registry robot
// permission details.
type ContainerRegistryRobotPermission struct {
	Kind      string                         `json:"kind"`
	Namespace string                         `json:"namespace"`
	Access    []ContainerRegistryRobotAccess `json:"access"`
}

// ContainerRegistryRobotAccess represents container registry robot access
// details.
type ContainerRegistryRobotAccess struct {
	Action   string `json:"action"`
	Resource string `json:"resource"`
	Effect   string `json:"effect"`
}

type containerRegistryRobotsBase struct {
	Robots []ContainerRegistryRobot `json:"robots"`
	Meta   *Meta                    `json:"meta"`
}

type ContainerRegistryRobotReq struct {
	Description string                       `json:"description,omitempty"`
	Disable     bool                         `json:"disable,omitempty"`
	Duration    int                          `json:"duration,omitempty"`
	Access      ContainerRegistryRobotAccess `json:"access,omitempty"`
}

// ContainerRegistryArtifact represents a container registry artifact.
type ContainerRegistryArtifact struct {
	ArtifactType      string                         `json:"artifact_type"`
	Digest            string                         `json:"digest"`
	ManifestMediaType string                         `json:"manifest_media_type"`
	MediaType         string                         `json:"media_type"`
	RepositoryName    string                         `json:"repository_name"`
	Size              int                            `json:"size"`
	Type              string                         `json:"type"`
	Tags              []ContainerRegistryArtifactTag `json:"tags"`
	DatePulled        string                         `json:"pull_time"`
	DatePushed        string                         `json:"push_time"`
}

// ContainerRegistryArtifactTag represents tags on an artifact.
type ContainerRegistryArtifactTag struct {
	Name       string `json:"name"`
	Immutable  bool   `json:"immutable"`
	DatePulled string `json:"pull_time"`
	DatePushed string `json:"push_time"`
}

type containerRegistryArtifactsBase struct {
	Artifacts []ContainerRegistryArtifact `json:"artifacts"`
	Meta      *Meta                       `json:"meta"`
}

// DockerCredentialsOpt contains the options used to create Docker credentials.
type DockerCredentialsOpt struct {
	ExpirySeconds *int
	WriteAccess   *bool
}

// ContainerRegistryDockerCredentials represents the byte array of character
// data returned after creating a Docker credential.
type ContainerRegistryDockerCredentials []byte

// UnmarshalJSON is a custom unmarshal function for
// ContainerRegistryDockerCredentials.
func (c *ContainerRegistryDockerCredentials) UnmarshalJSON(b []byte) error {
	*c = b
	return nil
}

// String converts the ContainerRegistryDockerCredentials to a string.
func (c *ContainerRegistryDockerCredentials) String() string {
	return string(*c)
}

// ContainerRegistryRegion represents the region data.
type ContainerRegistryRegion struct {
	ID           int                               `json:"id"`
	Name         string                            `json:"name"`
	URN          string                            `json:"urn"`
	BaseURL      string                            `json:"base_url"`
	Public       bool                              `json:"public"`
	DateCreated  string                            `json:"added_at"`
	DateModified string                            `json:"updated_at"`
	DataCenter   ContainerRegistryRegionDataCenter `json:"data_center"`
}

// ContainerRegistryRegionDataCenter is the datacenter info for a given region.
type ContainerRegistryRegionDataCenter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SiteCode    string `json:"site_code"`
	Region      string `json:"region"`
	Country     string `json:"country"`
	Continent   string `json:"continent"`
	Description string `json:"description"`
	Airport     string `json:"airport"`
}

type containerRegistryRegions struct {
	Regions []ContainerRegistryRegion `json:"regions"`
	Meta    *Meta                     `json:"meta"`
}

// ContainerRegistryPlans contains all plan types.
type ContainerRegistryPlans struct {
	Plans ContainerRegistryPlanTypes `json:"plans"`
}

// ContainerRegistryPlanTypes represent the different plan types.
type ContainerRegistryPlanTypes struct {
	StartUp    ContainerRegistryPlan `json:"start_up"`
	Business   ContainerRegistryPlan `json:"business"`
	Premium    ContainerRegistryPlan `json:"premium"`
	Enterprise ContainerRegistryPlan `json:"enterprise"`
}

// ContainerRegistryPlan represent the plan data.
type ContainerRegistryPlan struct {
	VanityName   string `json:"vanity_name"`
	MaxStorageMB int    `json:"max_storage_mb"`
	MonthlyPrice int    `json:"monthly_price"`
}

// Get retrieves a contrainer registry by ID
func (h *ContainerRegistryServiceHandler) Get(ctx context.Context, id string) (*ContainerRegistry, *http.Response, error) {
	req, errReq := h.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s", vcrPath, id), nil)
	if errReq != nil {
		return nil, nil, errReq
	}

	vcr := new(ContainerRegistry)
	resp, errResp := h.client.DoWithContext(ctx, req, &vcr)
	if errResp != nil {
		return nil, resp, errResp
	}

	return vcr, resp, nil
}

// List retrieves the list of all container registries
func (h *ContainerRegistryServiceHandler) List(ctx context.Context, options *ListOptions) ([]ContainerRegistry, *Meta, *http.Response, error) { //nolint:lll,dupl
	req, errReq := h.client.NewRequest(ctx, http.MethodGet, vcrListPath, nil)
	if errReq != nil {
		return nil, nil, nil, errReq
	}

	qStrings, errQ := query.Values(options)
	if errQ != nil {
		return nil, nil, nil, errQ
	}

	req.URL.RawQuery = qStrings.Encode()

	vcrs := new(containerRegistries)
	resp, errResp := h.client.DoWithContext(ctx, req, &vcrs)
	if errResp != nil {
		return nil, nil, resp, errResp
	}

	return vcrs.ContainerRegistries, vcrs.Meta, resp, nil
}

// Create creates a container registry
func (h *ContainerRegistryServiceHandler) Create(ctx context.Context, createReq *ContainerRegistryReq) (*ContainerRegistry, *http.Response, error) { //nolint:lll
	req, errReq := h.client.NewRequest(ctx, http.MethodPost, vcrPath, createReq)
	if errReq != nil {
		return nil, nil, errReq
	}

	vcr := new(ContainerRegistry)
	resp, errResp := h.client.DoWithContext(ctx, req, &vcr)
	if errResp != nil {
		return nil, resp, errResp
	}

	return vcr, resp, nil
}

// Update will update an existing container registry
func (h *ContainerRegistryServiceHandler) Update(ctx context.Context, vcrID string, updateReq *ContainerRegistryUpdateReq) (*ContainerRegistry, *http.Response, error) { //nolint:lll
	req, errReq := h.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("%s/%s", vcrPath, vcrID), updateReq)
	if errReq != nil {
		return nil, nil, errReq
	}

	vcr := new(ContainerRegistry)
	resp, errResp := h.client.DoWithContext(ctx, req, &vcr)
	if errResp != nil {
		return nil, resp, errResp
	}

	return vcr, resp, nil
}

// Delete will delete a container registry
func (h *ContainerRegistryServiceHandler) Delete(ctx context.Context, vcrID string) error {
	req, errReq := h.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s", vcrPath, vcrID), nil)
	if errReq != nil {
		return errReq
	}

	_, errResp := h.client.DoWithContext(ctx, req, nil)
	if errResp != nil {
		return errResp
	}

	return nil
}

// ListRepositories will get a list of the repositories for a existing
// container registry
func (h *ContainerRegistryServiceHandler) ListRepositories(ctx context.Context, vcrID string, options *ListOptions) ([]ContainerRegistryRepo, *Meta, *http.Response, error) { //nolint:lll,dupl
	req, errReq := h.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/repositories", vcrPath, vcrID), nil)
	if errReq != nil {
		return nil, nil, nil, errReq
	}

	qStrings, errQ := query.Values(options)
	if errQ != nil {
		return nil, nil, nil, errQ
	}

	req.URL.RawQuery = qStrings.Encode()

	vcrRepos := new(containerRegistryRepos)
	resp, errResp := h.client.DoWithContext(ctx, req, &vcrRepos)
	if errResp != nil {
		return nil, nil, resp, errResp
	}

	return vcrRepos.Repositories, vcrRepos.Meta, resp, nil
}

// GetRepository will return an existing repository of the requested registry
// ID and image name
func (h *ContainerRegistryServiceHandler) GetRepository(ctx context.Context, vcrID, imageName string) (*ContainerRegistryRepo, *http.Response, error) { //nolint:lll
	req, errReq := h.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/%s/repository/%s", vcrPath, vcrID, imageName), nil)
	if errReq != nil {
		return nil, nil, errReq
	}

	vcrRepo := new(ContainerRegistryRepo)
	resp, errResp := h.client.DoWithContext(ctx, req, &vcrRepo)
	if errResp != nil {
		return nil, resp, errResp
	}

	return vcrRepo, resp, nil
}

// UpdateRepository allows updating the repository with the specified registry
// ID and image name
func (h *ContainerRegistryServiceHandler) UpdateRepository(ctx context.Context, vcrID, imageName string, updateReq *ContainerRegistryRepoUpdateReq) (*ContainerRegistryRepo, *http.Response, error) { //nolint: lll
	req, errReq := h.client.NewRequest(ctx, http.MethodPut, fmt.Sprintf("%s/%s/repository/%s", vcrPath, vcrID, imageName), updateReq)
	if errReq != nil {
		return nil, nil, errReq
	}

	vcrRepo := new(ContainerRegistryRepo)
	resp, errResp := h.client.DoWithContext(ctx, req, &vcrRepo)
	if errResp != nil {
		return nil, resp, errResp
	}

	return vcrRepo, resp, nil
}

// DeleteRepository remove a repository from the container registry
func (h *ContainerRegistryServiceHandler) DeleteRepository(ctx context.Context, vcrID, imageName string) error {
	req, errReq := h.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf("%s/%s/repository/%s", vcrPath, vcrID, imageName), nil)
	if errReq != nil {
		return errReq
	}

	_, errResp := h.client.DoWithContext(ctx, req, nil)
	if errResp != nil {
		return errResp
	}

	return nil
}

// ListReplications retrieves a list of container registry replications
func (h *ContainerRegistryServiceHandler) ListReplications(ctx context.Context, vcrID string, options *ListOptions) ([]ContainerRegistryReplication, *Meta, *http.Response, error) { //nolint:lll,dupl
	url := fmt.Sprintf("%s/%s/replications", vcrPath, vcrID)

	req, err := h.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	qStrings, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = qStrings.Encode()

	replications := new(containerRegistryReplicationsBase)

	resp, err := h.client.DoWithContext(ctx, req, &replications)
	if err != nil {
		return nil, nil, resp, err
	}

	return replications.Replications, replications.Meta, resp, nil
}

// CreateReplication creates a container registry replication
func (h *ContainerRegistryServiceHandler) CreateReplication(ctx context.Context, vcrID, regionID string) (*ContainerRegistryReplication, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/replication", vcrPath, vcrID)

	req, err := h.client.NewRequest(ctx, http.MethodPost, url, &ContainerRegistryReplicationReq{Region: regionID})
	if err != nil {
		return nil, nil, err
	}

	replication := new(ContainerRegistryReplication)

	resp, err := h.client.DoWithContext(ctx, req, &replication)
	if err != nil {
		return nil, resp, err
	}

	return replication, resp, nil
}

// GetReplication retrieves a container registry replication
func (h *ContainerRegistryServiceHandler) GetReplication(ctx context.Context, vcrID, regionID string) (*ContainerRegistryReplication, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/replication/%s", vcrPath, vcrID, regionID)

	req, errReq := h.client.NewRequest(ctx, http.MethodGet, url, nil)
	if errReq != nil {
		return nil, nil, errReq
	}

	replication := new(ContainerRegistryReplication)

	resp, errResp := h.client.DoWithContext(ctx, req, &replication)
	if errResp != nil {
		return nil, resp, errResp
	}

	return replication, resp, nil
}

// DeleteReplication deletes a container registry replication
func (h *ContainerRegistryServiceHandler) DeleteReplication(ctx context.Context, vcrID, regionID string) error {
	url := fmt.Sprintf("%s/%s/replication/%s", vcrPath, vcrID, regionID)

	req, err := h.client.NewRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	if _, err := h.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// UpdateRetentionSchedule updates a container registry retention schedule
func (h *ContainerRegistryServiceHandler) UpdateRetentionSchedule(ctx context.Context, vcrID, schedule string) (*ContainerRegistryRetentionSchedule, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/retention/schedule", vcrPath, vcrID)

	req, err := h.client.NewRequest(ctx, http.MethodPut, url, &ContainerRegistryRetentionScheduleReq{Cron: schedule})
	if err != nil {
		return nil, nil, err
	}

	sched := new(ContainerRegistryRetentionSchedule)

	resp, errResp := h.client.DoWithContext(ctx, req, &sched)
	if errResp != nil {
		return nil, resp, errResp
	}

	return sched, resp, nil
}

// ExecuteRetention triggers a container registry retention execution
func (h *ContainerRegistryServiceHandler) ExecuteRetention(ctx context.Context, vcrID string, dryRun bool) (*ContainerRegistryRetentionExecution, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/retention/executions", vcrPath, vcrID)

	req, err := h.client.NewRequest(ctx, http.MethodPost, url, &ContainerRegistryRetentionExecutionReq{DryRun: dryRun})
	if err != nil {
		return nil, nil, err
	}

	execution := new(ContainerRegistryRetentionExecution)

	resp, errResp := h.client.DoWithContext(ctx, req, &execution)
	if errResp != nil {
		return nil, resp, errResp
	}

	return execution, resp, nil
}

// ListRetentionRules retrieves a list of container registry retention rules
func (h *ContainerRegistryServiceHandler) ListRetentionRules(ctx context.Context, vcrID string, options *ListOptions) ([]ContainerRegistryRetentionRule, *Meta, *http.Response, error) { //nolint:lll,dupl
	url := fmt.Sprintf("%s/%s/retention/rules", vcrPath, vcrID)

	req, err := h.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	qStrings, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = qStrings.Encode()

	rules := new(containerRegistryRetentionRulesBase)

	resp, err := h.client.DoWithContext(ctx, req, &rules)
	if err != nil {
		return nil, nil, resp, err
	}

	return rules.Rules, rules.Meta, resp, nil
}

// CreateRetentionRule creates a new container registry retention rule
func (h *ContainerRegistryServiceHandler) CreateRetentionRule(ctx context.Context, vcrID string, ruleReq *ContainerRegistryRetentionRuleReq) (*ContainerRegistryRetentionRule, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/retention/rules", vcrPath, vcrID)

	req, err := h.client.NewRequest(ctx, http.MethodPost, url, ruleReq)
	if err != nil {
		return nil, nil, err
	}

	rule := new(ContainerRegistryRetentionRule)

	resp, err := h.client.DoWithContext(ctx, req, &rule)
	if err != nil {
		return nil, resp, err
	}

	return rule, resp, nil
}

// GetRetentionRule retrieves a container registry retention rule
func (h *ContainerRegistryServiceHandler) GetRetentionRule(ctx context.Context, vcrID string, ruleID int) (*ContainerRegistryRetentionRule, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/retention/rules/%d", vcrPath, vcrID, ruleID)

	req, err := h.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	rule := new(ContainerRegistryRetentionRule)

	resp, err := h.client.DoWithContext(ctx, req, &rule)
	if err != nil {
		return nil, resp, err
	}

	return rule, resp, nil
}

// UpdateRetentionRule updates a container registry retention rule
func (h *ContainerRegistryServiceHandler) UpdateRetentionRule(ctx context.Context, vcrID string, ruleID int, disabled bool) (*ContainerRegistryRetentionRule, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/retention/rules/%d", vcrPath, vcrID, ruleID)

	req, err := h.client.NewRequest(ctx, http.MethodPut, url, &ContainerRegistryRetentionRuleUpdateReq{Disabled: disabled})
	if err != nil {
		return nil, nil, err
	}

	rule := new(ContainerRegistryRetentionRule)

	resp, err := h.client.DoWithContext(ctx, req, &rule)
	if err != nil {
		return nil, resp, err
	}

	return rule, resp, nil
}

// DeleteRetentionRule deletes a container registry retention rule
func (h *ContainerRegistryServiceHandler) DeleteRetentionRule(ctx context.Context, vcrID string, ruleID int) error {
	url := fmt.Sprintf("%s/%s/retention/rules/%d", vcrPath, vcrID, ruleID)

	req, err := h.client.NewRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	if _, err := h.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (h *ContainerRegistryServiceHandler) ListRobots(ctx context.Context, vcrID string, options *ListOptions) ([]ContainerRegistryRobot, *Meta, *http.Response, error) { //nolint:lll,dupl
	url := fmt.Sprintf("%s/%s/robots", vcrPath, vcrID)

	req, err := h.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	qStrings, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = qStrings.Encode()

	bots := new(containerRegistryRobotsBase)

	resp, err := h.client.DoWithContext(ctx, req, &bots)
	if err != nil {
		return nil, nil, resp, err
	}

	return bots.Robots, bots.Meta, resp, nil
}

func (h *ContainerRegistryServiceHandler) GetRobot(ctx context.Context, vcrID, robotName string) (*ContainerRegistryRobot, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/robot/%s", vcrPath, vcrID, robotName)

	req, err := h.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	bot := new(ContainerRegistryRobot)

	resp, err := h.client.DoWithContext(ctx, req, &bot)
	if err != nil {
		return nil, resp, err
	}

	return bot, resp, nil
}

func (h *ContainerRegistryServiceHandler) UpdateRobot(ctx context.Context, vcrID, robotName string, updateReq *ContainerRegistryRobotReq) (*ContainerRegistryRobot, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/robot/%s", vcrPath, vcrID, robotName)

	req, err := h.client.NewRequest(ctx, http.MethodPut, url, updateReq)
	if err != nil {
		return nil, nil, err
	}

	bot := new(ContainerRegistryRobot)

	resp, err := h.client.DoWithContext(ctx, req, &bot)
	if err != nil {
		return nil, resp, err
	}

	return bot, resp, nil
}

func (h *ContainerRegistryServiceHandler) DeleteRobot(ctx context.Context, vcrID, robotName string) error {
	url := fmt.Sprintf("%s/%s/robot/%s", vcrPath, vcrID, robotName)

	req, err := h.client.NewRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	if _, err := h.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (h *ContainerRegistryServiceHandler) ListArtifacts(ctx context.Context, vcrID, imageName string, options *ListOptions) ([]ContainerRegistryArtifact, *Meta, *http.Response, error) { //nolint:lll,dupl
	url := fmt.Sprintf("%s/%s/repository/%s/artifacts", vcrPath, vcrID, imageName)

	req, err := h.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	qStrings, err := query.Values(options)
	if err != nil {
		return nil, nil, nil, err
	}

	req.URL.RawQuery = qStrings.Encode()

	artifacts := new(containerRegistryArtifactsBase)

	resp, err := h.client.DoWithContext(ctx, req, &artifacts)
	if err != nil {
		return nil, nil, resp, err
	}

	return artifacts.Artifacts, artifacts.Meta, resp, nil
}

func (h *ContainerRegistryServiceHandler) GetArtifact(ctx context.Context, vcrID, imageName, artifactDigest string) (*ContainerRegistryArtifact, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/repository/%s/artifact/%s", vcrPath, vcrID, imageName, artifactDigest)

	req, err := h.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	artifact := new(ContainerRegistryArtifact)

	resp, err := h.client.DoWithContext(ctx, req, &artifact)
	if err != nil {
		return nil, resp, err
	}

	return artifact, resp, nil
}

func (h *ContainerRegistryServiceHandler) DeleteArtifact(ctx context.Context, vcrID, imageName, artifactDigest string) error {
	url := fmt.Sprintf("%s/%s/repository/%s/artifact/%s", vcrPath, vcrID, imageName, artifactDigest)

	req, err := h.client.NewRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	if _, err := h.client.DoWithContext(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// CreateDockerCredentials will create new Docker credentials used by the
// Docker CLI
func (h *ContainerRegistryServiceHandler) CreateDockerCredentials(ctx context.Context, vcrID string, createOptions *DockerCredentialsOpt) (*ContainerRegistryDockerCredentials, *http.Response, error) { //nolint:lll
	url := fmt.Sprintf("%s/%s/docker-credentials", vcrPath, vcrID)
	req, errReq := h.client.NewRequest(ctx, http.MethodOptions, url, nil)
	if errReq != nil {
		return nil, nil, errReq
	}

	queryParam := req.URL.Query()
	if createOptions.ExpirySeconds != nil {
		queryParam.Add("expiry_seconds", fmt.Sprintf("%d", createOptions.ExpirySeconds))
	}

	if createOptions.WriteAccess != nil {
		queryParam.Add("read_write", fmt.Sprintf("%t", *createOptions.WriteAccess))
	}

	req.URL.RawQuery = queryParam.Encode()

	creds := new(ContainerRegistryDockerCredentials)
	resp, errResp := h.client.DoWithContext(ctx, req, &creds)
	if errResp != nil {
		return nil, nil, errResp
	}

	return creds, resp, nil
}

// ListRegions will return a list of regions relevant to the container registry
// API operations
func (h *ContainerRegistryServiceHandler) ListRegions(ctx context.Context) ([]ContainerRegistryRegion, *Meta, *http.Response, error) {
	req, errReq := h.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/region/list", vcrPath), nil)
	if errReq != nil {
		return nil, nil, nil, errReq
	}

	vcrRegions := new(containerRegistryRegions)
	resp, errResp := h.client.DoWithContext(ctx, req, &vcrRegions)
	if errResp != nil {
		return nil, nil, resp, errResp
	}

	return vcrRegions.Regions, vcrRegions.Meta, resp, nil
}

// ListPlans returns a list of plans relevant to the container registry
// offerings
func (h *ContainerRegistryServiceHandler) ListPlans(ctx context.Context) (*ContainerRegistryPlans, *http.Response, error) {
	req, errReq := h.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/plan/list", vcrPath), nil)
	if errReq != nil {
		return nil, nil, errReq
	}

	vcrPlans := new(ContainerRegistryPlans)
	resp, errResp := h.client.DoWithContext(ctx, req, &vcrPlans)
	if errResp != nil {
		return nil, resp, errResp
	}

	return vcrPlans, resp, nil
}
