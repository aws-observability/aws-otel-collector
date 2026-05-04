// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	client "github.com/DataDog/datadog-api-client-go/v2"
)

//go:generate mockgen -source=./configuration.go -destination=../../tests/api/mocks/configuration.go -package=mocks

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKeys takes a string apikey as authentication for the request
	ContextAPIKeys = contextKey("apiKeys")

	// ContextDelegatedToken takes a string delegated token as authentication for the request
	ContextDelegatedToken = contextKey("delegatedToken")

	// ContextAWSVariables takes AWS credentials as authentication for the request.
	ContextAWSVariables = contextKey("awsVariables")

	// ContextHttpSignatureAuth takes HttpSignatureAuth as authentication for the request.
	ContextHttpSignatureAuth = contextKey("httpsignature")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth.
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey.
type APIKey struct {
	Key    string
	Prefix string
}

// DelegatedTokenCredentials delegated token authentication to a request passed via context using ContextDelegatedToken.
type DelegatedTokenCredentials struct {
	OrgUUID        string
	DelegatedToken string
	DelegatedProof string
	Expiration     time.Time
}

// DelegatedTokenConfig provides cloud provider based authentication configuration.
type DelegatedTokenConfig struct {
	OrgUUID      string
	Provider     string
	ProviderAuth DelegatedTokenProvider
}

// DelegatedTokenProvider is an interface for getting a delegated token utilizing different methods.
type DelegatedTokenProvider interface {
	Authenticate(ctx context.Context, config *DelegatedTokenConfig) (*DelegatedTokenCredentials, error)
}

// ServerVariable stores the information about a server variable.
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server.
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items.
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host                 string            `json:"host,omitempty"`
	Scheme               string            `json:"scheme,omitempty"`
	DefaultHeader        map[string]string `json:"defaultHeader,omitempty"`
	UserAgent            string            `json:"userAgent,omitempty"`
	Debug                bool              `json:"debug,omitempty"`
	Compress             bool              `json:"compress,omitempty"`
	Servers              ServerConfigurations
	OperationServers     map[string]ServerConfigurations
	HTTPClient           *http.Client
	unstableOperations   map[string]bool
	RetryConfiguration   RetryConfiguration
	DelegatedTokenConfig *DelegatedTokenConfig
}

// RetryConfiguration stores the configuration of the retry behavior of the api client
type RetryConfiguration struct {
	EnableRetry       bool
	BackOffMultiplier float64
	BackOffBase       float64
	HTTPRetryTimeout  time.Duration
	MaxRetries        int
}

// NewConfiguration returns a new Configuration object.
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     GetUserAgent(),
		Debug:         false,
		Compress:      true,
		Servers: ServerConfigurations{
			{
				URL:         "https://{subdomain}.{site}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"site": {
						Description:  "The regional site for Datadog customers.",
						DefaultValue: "datadoghq.com",
						EnumValues: []string{
							"datadoghq.com",
							"us3.datadoghq.com",
							"us5.datadoghq.com",
							"ap1.datadoghq.com",
							"ap2.datadoghq.com",
							"datadoghq.eu",
							"ddog-gov.com",
							"us2.ddog-gov.com",
						},
					},
					"subdomain": {
						Description:  "The subdomain where the API is deployed.",
						DefaultValue: "api",
					},
				},
			},
			{
				URL:         "{protocol}://{name}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"name": {
						Description:  "Full site DNS name.",
						DefaultValue: "api.datadoghq.com",
					},
					"protocol": {
						Description:  "The protocol for accessing the API.",
						DefaultValue: "https",
					},
				},
			},
			{
				URL:         "https://{subdomain}.{site}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"site": {
						Description:  "Any Datadog deployment.",
						DefaultValue: "datadoghq.com",
					},
					"subdomain": {
						Description:  "The subdomain where the API is deployed.",
						DefaultValue: "api",
					},
				},
			},
		},
		OperationServers: map[string]ServerConfigurations{
			"v1.IPRangesApi.GetIPRanges": {
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "The regional site for Datadog customers.",
							DefaultValue: "datadoghq.com",
							EnumValues: []string{
								"datadoghq.com",
								"us3.datadoghq.com",
								"us5.datadoghq.com",
								"ap1.datadoghq.com",
								"ap2.datadoghq.com",
								"datadoghq.eu",
								"ddog-gov.com",
								"us2.ddog-gov.com",
							},
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "ip-ranges",
						},
					},
				},
				{
					URL:         "{protocol}://{name}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"name": {
							Description:  "Full site DNS name.",
							DefaultValue: "ip-ranges.datadoghq.com",
						},
						"protocol": {
							Description:  "The protocol for accessing the API.",
							DefaultValue: "https",
						},
					},
				},
				{
					URL:         "https://{subdomain}.datadoghq.com",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "ip-ranges",
						},
					},
				},
			},
			"v1.LogsApi.SubmitLog": {
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "The regional site for Datadog customers.",
							DefaultValue: "datadoghq.com",
							EnumValues: []string{
								"datadoghq.com",
								"us3.datadoghq.com",
								"us5.datadoghq.com",
								"ap1.datadoghq.com",
								"ap2.datadoghq.com",
								"datadoghq.eu",
								"ddog-gov.com",
								"us2.ddog-gov.com",
							},
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "http-intake.logs",
						},
					},
				},
				{
					URL:         "{protocol}://{name}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"name": {
							Description:  "Full site DNS name.",
							DefaultValue: "http-intake.logs.datadoghq.com",
						},
						"protocol": {
							Description:  "The protocol for accessing the API.",
							DefaultValue: "https",
						},
					},
				},
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "Any Datadog deployment.",
							DefaultValue: "datadoghq.com",
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "http-intake.logs",
						},
					},
				},
			},
			"v2.EventsApi.CreateEvent": {
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "The regional site for customers.",
							DefaultValue: "datadoghq.com",
							EnumValues: []string{
								"datadoghq.com",
								"us3.datadoghq.com",
								"us5.datadoghq.com",
								"ap1.datadoghq.com",
								"ap2.datadoghq.com",
								"datadoghq.eu",
								"ddog-gov.com",
								"us2.ddog-gov.com",
							},
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "event-management-intake",
						},
					},
				},
				{
					URL:         "{protocol}://{name}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"name": {
							Description:  "Full site DNS name.",
							DefaultValue: "event-management-intake.datadoghq.com",
						},
						"protocol": {
							Description:  "The protocol for accessing the API.",
							DefaultValue: "https",
						},
					},
				},
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "Any Datadog deployment.",
							DefaultValue: "datadoghq.com",
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "event-management-intake",
						},
					},
				},
			},
			"v2.LogsApi.SubmitLog": {
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "The regional site for customers.",
							DefaultValue: "datadoghq.com",
							EnumValues: []string{
								"datadoghq.com",
								"us3.datadoghq.com",
								"us5.datadoghq.com",
								"ap1.datadoghq.com",
								"ap2.datadoghq.com",
								"datadoghq.eu",
								"ddog-gov.com",
								"us2.ddog-gov.com",
							},
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "http-intake.logs",
						},
					},
				},
				{
					URL:         "{protocol}://{name}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"name": {
							Description:  "Full site DNS name.",
							DefaultValue: "http-intake.logs.datadoghq.com",
						},
						"protocol": {
							Description:  "The protocol for accessing the API.",
							DefaultValue: "https",
						},
					},
				},
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "Any Datadog deployment.",
							DefaultValue: "datadoghq.com",
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "http-intake.logs",
						},
					},
				},
			},
			"v2.OnCallPagingApi.CreateOnCallPage": {
				{
					URL:         "https://{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "The globally available endpoint for On-Call.",
							DefaultValue: "navy.oncall.datadoghq.com",
							EnumValues: []string{
								"lava.oncall.datadoghq.com",
								"saffron.oncall.datadoghq.com",
								"navy.oncall.datadoghq.com",
								"coral.oncall.datadoghq.com",
								"teal.oncall.datadoghq.com",
								"beige.oncall.datadoghq.eu",
							},
						},
					},
				},
				{
					URL:         "{protocol}://{name}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"name": {
							Description:  "Full site DNS name.",
							DefaultValue: "api.datadoghq.com",
						},
						"protocol": {
							Description:  "The protocol for accessing the API.",
							DefaultValue: "https",
						},
					},
				},
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "Any Datadog deployment.",
							DefaultValue: "datadoghq.com",
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "api",
						},
					},
				},
			},
			"v2.OnCallPagingApi.AcknowledgeOnCallPage": {
				{
					URL:         "https://{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "The globally available endpoint for On-Call.",
							DefaultValue: "navy.oncall.datadoghq.com",
							EnumValues: []string{
								"lava.oncall.datadoghq.com",
								"saffron.oncall.datadoghq.com",
								"navy.oncall.datadoghq.com",
								"coral.oncall.datadoghq.com",
								"teal.oncall.datadoghq.com",
								"beige.oncall.datadoghq.eu",
							},
						},
					},
				},
				{
					URL:         "{protocol}://{name}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"name": {
							Description:  "Full site DNS name.",
							DefaultValue: "api.datadoghq.com",
						},
						"protocol": {
							Description:  "The protocol for accessing the API.",
							DefaultValue: "https",
						},
					},
				},
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "Any Datadog deployment.",
							DefaultValue: "datadoghq.com",
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "api",
						},
					},
				},
			},
			"v2.OnCallPagingApi.EscalateOnCallPage": {
				{
					URL:         "https://{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "The globally available endpoint for On-Call.",
							DefaultValue: "navy.oncall.datadoghq.com",
							EnumValues: []string{
								"lava.oncall.datadoghq.com",
								"saffron.oncall.datadoghq.com",
								"navy.oncall.datadoghq.com",
								"coral.oncall.datadoghq.com",
								"teal.oncall.datadoghq.com",
								"beige.oncall.datadoghq.eu",
							},
						},
					},
				},
				{
					URL:         "{protocol}://{name}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"name": {
							Description:  "Full site DNS name.",
							DefaultValue: "api.datadoghq.com",
						},
						"protocol": {
							Description:  "The protocol for accessing the API.",
							DefaultValue: "https",
						},
					},
				},
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "Any Datadog deployment.",
							DefaultValue: "datadoghq.com",
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "api",
						},
					},
				},
			},
			"v2.OnCallPagingApi.ResolveOnCallPage": {
				{
					URL:         "https://{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "The globally available endpoint for On-Call.",
							DefaultValue: "navy.oncall.datadoghq.com",
							EnumValues: []string{
								"lava.oncall.datadoghq.com",
								"saffron.oncall.datadoghq.com",
								"navy.oncall.datadoghq.com",
								"coral.oncall.datadoghq.com",
								"teal.oncall.datadoghq.com",
								"beige.oncall.datadoghq.eu",
							},
						},
					},
				},
				{
					URL:         "{protocol}://{name}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"name": {
							Description:  "Full site DNS name.",
							DefaultValue: "api.datadoghq.com",
						},
						"protocol": {
							Description:  "The protocol for accessing the API.",
							DefaultValue: "https",
						},
					},
				},
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "Any Datadog deployment.",
							DefaultValue: "datadoghq.com",
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "api",
						},
					},
				},
			},
			"v2.ProductAnalyticsApi.SubmitProductAnalyticsEvent": {
				{
					URL:         "https://{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "The intake domain for the regional site.",
							DefaultValue: "browser-intake-datadoghq.com",
							EnumValues: []string{
								"browser-intake-datadoghq.com",
								"browser-intake-us3-datadoghq.com",
								"browser-intake-us5-datadoghq.com",
								"browser-intake-ap1-datadoghq.com",
								"browser-intake-ap2-datadoghq.com",
								"browser-intake-datadoghq.eu",
							},
						},
					},
				},
				{
					URL:         "{protocol}://{name}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"name": {
							Description:  "Full site DNS name.",
							DefaultValue: "browser-intake-datadoghq.com",
						},
						"protocol": {
							Description:  "The protocol for accessing the API.",
							DefaultValue: "https",
						},
					},
				},
				{
					URL:         "https://{subdomain}.{site}",
					Description: "No description provided",
					Variables: map[string]ServerVariable{
						"site": {
							Description:  "Any Datadog deployment.",
							DefaultValue: "datadoghq.com",
						},
						"subdomain": {
							Description:  "The subdomain where the API is deployed.",
							DefaultValue: "api",
						},
					},
				},
			},
		},
		unstableOperations: map[string]bool{
			"v2.CancelFleetDeployment":                          false,
			"v2.CreateFleetDeploymentConfigure":                 false,
			"v2.CreateFleetDeploymentUpgrade":                   false,
			"v2.CreateFleetSchedule":                            false,
			"v2.DeleteFleetSchedule":                            false,
			"v2.GetFleetAgentInfo":                              false,
			"v2.GetFleetDeployment":                             false,
			"v2.GetFleetSchedule":                               false,
			"v2.ListFleetAgents":                                false,
			"v2.ListFleetAgentTracers":                          false,
			"v2.ListFleetAgentVersions":                         false,
			"v2.ListFleetClusters":                              false,
			"v2.ListFleetDeployments":                           false,
			"v2.ListFleetInstrumentedPods":                      false,
			"v2.ListFleetSchedules":                             false,
			"v2.ListFleetTracers":                               false,
			"v2.TriggerFleetSchedule":                           false,
			"v2.UpdateFleetSchedule":                            false,
			"v2.CreateLLMObsAnnotationQueue":                    false,
			"v2.CreateLLMObsAnnotationQueueInteractions":        false,
			"v2.CreateLLMObsDataset":                            false,
			"v2.CreateLLMObsDatasetRecords":                     false,
			"v2.CreateLLMObsExperiment":                         false,
			"v2.CreateLLMObsExperimentEvents":                   false,
			"v2.CreateLLMObsProject":                            false,
			"v2.DeleteLLMObsAnnotationQueue":                    false,
			"v2.DeleteLLMObsAnnotationQueueInteractions":        false,
			"v2.DeleteLLMObsCustomEvalConfig":                   false,
			"v2.DeleteLLMObsDatasetRecords":                     false,
			"v2.DeleteLLMObsDatasets":                           false,
			"v2.DeleteLLMObsExperiments":                        false,
			"v2.DeleteLLMObsProjects":                           false,
			"v2.GetLLMObsAnnotatedInteractions":                 false,
			"v2.GetLLMObsCustomEvalConfig":                      false,
			"v2.ListLLMObsAnnotationQueues":                     false,
			"v2.ListLLMObsDatasetRecords":                       false,
			"v2.ListLLMObsDatasets":                             false,
			"v2.ListLLMObsExperiments":                          false,
			"v2.ListLLMObsProjects":                             false,
			"v2.UpdateLLMObsAnnotationQueue":                    false,
			"v2.UpdateLLMObsCustomEvalConfig":                   false,
			"v2.UpdateLLMObsDataset":                            false,
			"v2.UpdateLLMObsDatasetRecords":                     false,
			"v2.UpdateLLMObsExperiment":                         false,
			"v2.UpdateLLMObsProject":                            false,
			"v2.AnonymizeUsers":                                 false,
			"v2.CreateOpenAPI":                                  false,
			"v2.DeleteOpenAPI":                                  false,
			"v2.GetOpenAPI":                                     false,
			"v2.ListAPIs":                                       false,
			"v2.UpdateOpenAPI":                                  false,
			"v2.GetInvestigation":                               false,
			"v2.ListInvestigations":                             false,
			"v2.TriggerInvestigation":                           false,
			"v2.CreateCaseJiraIssue":                            false,
			"v2.CreateCaseNotebook":                             false,
			"v2.CreateCaseServiceNowTicket":                     false,
			"v2.LinkIncident":                                   false,
			"v2.LinkJiraIssueToCase":                            false,
			"v2.MoveCaseToProject":                              false,
			"v2.UnlinkJiraIssue":                                false,
			"v2.CreateChangeRequest":                            false,
			"v2.CreateChangeRequestBranch":                      false,
			"v2.DeleteChangeRequestDecision":                    false,
			"v2.GetChangeRequest":                               false,
			"v2.UpdateChangeRequest":                            false,
			"v2.UpdateChangeRequestDecision":                    false,
			"v2.DeleteTestOptimizationServiceSettings":          false,
			"v2.GetFlakyTestsManagementPolicies":                false,
			"v2.GetTestOptimizationServiceSettings":             false,
			"v2.SearchFlakyTests":                               false,
			"v2.UpdateFlakyTests":                               false,
			"v2.UpdateFlakyTestsManagementPolicies":             false,
			"v2.UpdateTestOptimizationServiceSettings":          false,
			"v2.CreateAWSCloudAuthPersonaMapping":               false,
			"v2.DeleteAWSCloudAuthPersonaMapping":               false,
			"v2.GetAWSCloudAuthPersonaMapping":                  false,
			"v2.ListAWSCloudAuthPersonaMappings":                false,
			"v2.ActivateContentPack":                            false,
			"v2.BulkExportSecurityMonitoringTerraformResources": false,
			"v2.CancelHistoricalJob":                            false,
			"v2.ConvertJobResultToSignal":                       false,
			"v2.ConvertSecurityMonitoringTerraformResource":     false,
			"v2.DeactivateContentPack":                          false,
			"v2.DeleteHistoricalJob":                            false,
			"v2.ExportSecurityMonitoringTerraformResource":      false,
			"v2.GetContentPacksStates":                          false,
			"v2.GetFinding":                                     false,
			"v2.GetHistoricalJob":                               false,
			"v2.GetIndicatorOfCompromise":                       false,
			"v2.GetRuleVersionHistory":                          false,
			"v2.GetSecretsRules":                                false,
			"v2.GetSecurityMonitoringHistsignal":                false,
			"v2.GetSecurityMonitoringHistsignalsByJobId":        false,
			"v2.ListFindings":                                   false,
			"v2.ListHistoricalJobs":                             false,
			"v2.ListIndicatorsOfCompromise":                     false,
			"v2.ListMultipleRulesets":                           false,
			"v2.ListScannedAssetsMetadata":                      false,
			"v2.ListSecurityMonitoringHistsignals":              false,
			"v2.ListVulnerabilities":                            false,
			"v2.ListVulnerableAssets":                           false,
			"v2.MuteFindings":                                   false,
			"v2.RunHistoricalJob":                               false,
			"v2.SearchSecurityMonitoringHistsignals":            false,
			"v2.GetCodeCoverageBranchSummary":                   false,
			"v2.GetCodeCoverageCommitSummary":                   false,
			"v2.CreateDashboardSecureEmbed":                     false,
			"v2.DeleteDashboardSecureEmbed":                     false,
			"v2.GetDashboardSecureEmbed":                        false,
			"v2.UpdateDashboardSecureEmbed":                     false,
			"v2.CreateDataset":                                  false,
			"v2.DeleteDataset":                                  false,
			"v2.GetAllDatasets":                                 false,
			"v2.GetDataset":                                     false,
			"v2.UpdateDataset":                                  false,
			"v2.CancelDataDeletionRequest":                      false,
			"v2.CreateDataDeletionRequest":                      false,
			"v2.GetDataDeletionRequests":                        false,
			"v2.CreateDeploymentGate":                           false,
			"v2.CreateDeploymentRule":                           false,
			"v2.DeleteDeploymentGate":                           false,
			"v2.DeleteDeploymentRule":                           false,
			"v2.GetDeploymentGate":                              false,
			"v2.GetDeploymentGateRules":                         false,
			"v2.GetDeploymentGatesEvaluationResult":             false,
			"v2.GetDeploymentRule":                              false,
			"v2.ListDeploymentGates":                            false,
			"v2.TriggerDeploymentGatesEvaluation":               false,
			"v2.UpdateDeploymentGate":                           false,
			"v2.UpdateDeploymentRule":                           false,
			"v2.CreateHamrOrgConnection":                        false,
			"v2.GetHamrOrgConnection":                           false,
			"v2.CreateGlobalIncidentHandle":                     false,
			"v2.CreateIncident":                                 false,
			"v2.CreateIncidentAttachment":                       false,
			"v2.CreateIncidentIntegration":                      false,
			"v2.CreateIncidentNotificationRule":                 false,
			"v2.CreateIncidentNotificationTemplate":             false,
			"v2.CreateIncidentPostmortemAttachment":             false,
			"v2.CreateIncidentPostmortemTemplate":               false,
			"v2.CreateIncidentTodo":                             false,
			"v2.CreateIncidentType":                             false,
			"v2.CreateIncidentUserDefinedField":                 false,
			"v2.DeleteGlobalIncidentHandle":                     false,
			"v2.DeleteIncident":                                 false,
			"v2.DeleteIncidentAttachment":                       false,
			"v2.DeleteIncidentIntegration":                      false,
			"v2.DeleteIncidentNotificationRule":                 false,
			"v2.DeleteIncidentNotificationTemplate":             false,
			"v2.DeleteIncidentPostmortemTemplate":               false,
			"v2.DeleteIncidentTodo":                             false,
			"v2.DeleteIncidentType":                             false,
			"v2.DeleteIncidentUserDefinedField":                 false,
			"v2.GetGlobalIncidentSettings":                      false,
			"v2.GetIncident":                                    false,
			"v2.GetIncidentIntegration":                         false,
			"v2.GetIncidentNotificationRule":                    false,
			"v2.GetIncidentNotificationTemplate":                false,
			"v2.GetIncidentPostmortemTemplate":                  false,
			"v2.GetIncidentTodo":                                false,
			"v2.GetIncidentType":                                false,
			"v2.GetIncidentUserDefinedField":                    false,
			"v2.ImportIncident":                                 false,
			"v2.ListGlobalIncidentHandles":                      false,
			"v2.ListIncidentAttachments":                        false,
			"v2.ListIncidentIntegrations":                       false,
			"v2.ListIncidentNotificationRules":                  false,
			"v2.ListIncidentNotificationTemplates":              false,
			"v2.ListIncidentPostmortemTemplates":                false,
			"v2.ListIncidents":                                  false,
			"v2.ListIncidentTodos":                              false,
			"v2.ListIncidentTypes":                              false,
			"v2.ListIncidentUserDefinedFields":                  false,
			"v2.SearchIncidents":                                false,
			"v2.UpdateGlobalIncidentHandle":                     false,
			"v2.UpdateGlobalIncidentSettings":                   false,
			"v2.UpdateIncident":                                 false,
			"v2.UpdateIncidentAttachment":                       false,
			"v2.UpdateIncidentIntegration":                      false,
			"v2.UpdateIncidentNotificationRule":                 false,
			"v2.UpdateIncidentNotificationTemplate":             false,
			"v2.UpdateIncidentPostmortemTemplate":               false,
			"v2.UpdateIncidentTodo":                             false,
			"v2.UpdateIncidentType":                             false,
			"v2.UpdateIncidentUserDefinedField":                 false,
			"v2.CreateAWSAccountCCMConfig":                      false,
			"v2.DeleteAWSAccountCCMConfig":                      false,
			"v2.GetAWSAccountCCMConfig":                         false,
			"v2.UpdateAWSAccountCCMConfig":                      false,
			"v2.CreateJiraIssueTemplate":                        false,
			"v2.DeleteJiraAccount":                              false,
			"v2.DeleteJiraIssueTemplate":                        false,
			"v2.GetJiraIssueTemplate":                           false,
			"v2.ListJiraAccounts":                               false,
			"v2.ListJiraIssueTemplates":                         false,
			"v2.UpdateJiraIssueTemplate":                        false,
			"v2.CreateTenancyConfig":                            false,
			"v2.GetTenancyConfigs":                              false,
			"v2.AddRoleToRestrictionQuery":                      false,
			"v2.CreateRestrictionQuery":                         false,
			"v2.DeleteRestrictionQuery":                         false,
			"v2.GetRestrictionQuery":                            false,
			"v2.GetRoleRestrictionQuery":                        false,
			"v2.ListRestrictionQueries":                         false,
			"v2.ListRestrictionQueryRoles":                      false,
			"v2.ListUserRestrictionQueries":                     false,
			"v2.RemoveRoleFromRestrictionQuery":                 false,
			"v2.ReplaceRestrictionQuery":                        false,
			"v2.UpdateRestrictionQuery":                         false,
			"v2.CreateMonitorUserTemplate":                      false,
			"v2.DeleteMonitorUserTemplate":                      false,
			"v2.GetMonitorUserTemplate":                         false,
			"v2.ListMonitorUserTemplates":                       false,
			"v2.UpdateMonitorUserTemplate":                      false,
			"v2.ValidateExistingMonitorUserTemplate":            false,
			"v2.ValidateMonitorUserTemplate":                    false,
			"v2.BulkUpdateOrgGroupMemberships":                  false,
			"v2.CreateOrgGroup":                                 false,
			"v2.CreateOrgGroupPolicy":                           false,
			"v2.CreateOrgGroupPolicyOverride":                   false,
			"v2.DeleteOrgGroup":                                 false,
			"v2.DeleteOrgGroupPolicy":                           false,
			"v2.DeleteOrgGroupPolicyOverride":                   false,
			"v2.GetOrgGroup":                                    false,
			"v2.GetOrgGroupMembership":                          false,
			"v2.GetOrgGroupPolicy":                              false,
			"v2.GetOrgGroupPolicyOverride":                      false,
			"v2.ListOrgGroupMemberships":                        false,
			"v2.ListOrgGroupPolicies":                           false,
			"v2.ListOrgGroupPolicyConfigs":                      false,
			"v2.ListOrgGroupPolicyOverrides":                    false,
			"v2.ListOrgGroups":                                  false,
			"v2.UpdateOrgGroup":                                 false,
			"v2.UpdateOrgGroupMembership":                       false,
			"v2.UpdateOrgGroupPolicy":                           false,
			"v2.UpdateOrgGroupPolicyOverride":                   false,
			"v2.ListRoleTemplates":                              false,
			"v2.CreateConnection":                               false,
			"v2.DeleteConnection":                               false,
			"v2.GetAccountFacetInfo":                            false,
			"v2.GetMapping":                                     false,
			"v2.GetUserFacetInfo":                               false,
			"v2.ListConnections":                                false,
			"v2.QueryAccounts":                                  false,
			"v2.QueryEventFilteredUsers":                        false,
			"v2.QueryUsers":                                     false,
			"v2.UpdateConnection":                               false,
			"v2.CreateScorecardOutcomesBatch":                   false,
			"v2.ListEntityRiskScores":                           false,
			"v2.CreateIncidentService":                          false,
			"v2.DeleteIncidentService":                          false,
			"v2.GetIncidentService":                             false,
			"v2.ListIncidentServices":                           false,
			"v2.UpdateIncidentService":                          false,
			"v2.CreateSLOReportJob":                             false,
			"v2.GetSLOReport":                                   false,
			"v2.GetSLOReportJobStatus":                          false,
			"v2.GetSloStatus":                                   false,
			"v2.GetSPARecommendations":                          false,
			"v2.GetSPARecommendationsWithShard":                 false,
			"v2.CreateCustomRule":                               false,
			"v2.CreateCustomRuleRevision":                       false,
			"v2.CreateSCAResolveVulnerableSymbols":              false,
			"v2.CreateSCAResult":                                false,
			"v2.DeleteCustomRule":                               false,
			"v2.DeleteCustomRuleset":                            false,
			"v2.GetCustomRule":                                  false,
			"v2.GetCustomRuleRevision":                          false,
			"v2.GetCustomRuleset":                               false,
			"v2.ListCustomRuleRevisions":                        false,
			"v2.RevertCustomRuleRevision":                       false,
			"v2.UpdateCustomRuleset":                            false,
			"v2.AddMemberTeam":                                  false,
			"v2.ListMemberTeams":                                false,
			"v2.RemoveMemberTeam":                               false,
			"v2.CreateIncidentTeam":                             false,
			"v2.DeleteIncidentTeam":                             false,
			"v2.GetIncidentTeam":                                false,
			"v2.ListIncidentTeams":                              false,
			"v2.UpdateIncidentTeam":                             false,
		},
		RetryConfiguration: RetryConfiguration{
			EnableRetry:       false,
			BackOffMultiplier: 2,
			BackOffBase:       2,
			HTTPRetryTimeout:  60 * time.Second,
			MaxRetries:        3,
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request.
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables.
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings.
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, ReportError("invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		operationIndices, ok := osi.(map[string]int)
		if !ok {
			return 0, ReportError("invalid type %T should be map[string]int", osi)
		}
		index, ok := operationIndices[endpoint]
		if ok {
			return index, nil
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, ReportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		operationVariables, ok := osv.(map[string]map[string]string)
		if !ok {
			return nil, ReportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		}
		variables, ok := operationVariables[endpoint]
		if ok {
			return variables, nil
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint.
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}

// GetUnstableOperations returns a slice with all unstable operation Ids.
func (c *Configuration) GetUnstableOperations() []string {
	ids := make([]string, len(c.unstableOperations))
	for id := range c.unstableOperations {
		ids = append(ids, id)
	}
	return ids
}

// SetUnstableOperationEnabled sets an unstable operation as enabled (true) or disabled (false).
// This function accepts operation ID as an argument - this is the name of the method on the API class, e.g. "CreateFoo"
// Returns true if the operation is marked as unstable and thus was enabled/disabled, false otherwise.
func (c *Configuration) SetUnstableOperationEnabled(operation string, enabled bool) bool {
	if _, ok := c.unstableOperations[operation]; ok {
		c.unstableOperations[operation] = enabled
		return true
	}
	log.Printf("WARNING: '%s' is not an unstable operation, can't enable/disable", operation)
	return false
}

// IsUnstableOperation determines whether an operation is an unstable operation.
// This function accepts operation ID as an argument - this is the name of the method on the API class, e.g. "CreateFoo".
func (c *Configuration) IsUnstableOperation(operation string) bool {
	_, present := c.unstableOperations[operation]
	return present
}

// IsUnstableOperationEnabled determines whether an unstable operation is enabled.
// This function accepts operation ID as an argument - this is the name of the method on the API class, e.g. "CreateFoo"
// Returns true if the operation is unstable and it is enabled, false otherwise.
func (c *Configuration) IsUnstableOperationEnabled(operation string) bool {
	if enabled, present := c.unstableOperations[operation]; present {
		return enabled
	}
	log.Printf("WARNING: '%s' is not an unstable operation, is always enabled", operation)
	return false
}

func GetUserAgent() string {
	return fmt.Sprintf(
		"datadog-api-client-go/%s (go %s; os %s; arch %s)",
		client.Version,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
}

// NewDefaultContext returns a new context setup with environment variables.
func NewDefaultContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	if site, ok := os.LookupEnv("DD_SITE"); ok {
		ctx = context.WithValue(
			ctx,
			ContextServerVariables,
			map[string]string{"site": site},
		)
	}

	keys := make(map[string]APIKey)
	if apiKey, ok := os.LookupEnv("DD_API_KEY"); ok {
		keys["apiKeyAuth"] = APIKey{Key: apiKey}
	}
	if apiKey, ok := os.LookupEnv("DD_APP_KEY"); ok {
		keys["appKeyAuth"] = APIKey{Key: apiKey}
	}
	ctx = context.WithValue(
		ctx,
		ContextAPIKeys,
		keys,
	)

	awsKeys := make(map[string]string)
	if accessKey, ok := os.LookupEnv(AWSAccessKeyIdName); ok {
		awsKeys[AWSAccessKeyIdName] = accessKey
	}
	if secretKey, ok := os.LookupEnv(AWSSecretAccessKeyName); ok {
		awsKeys[AWSSecretAccessKeyName] = secretKey
	}
	if sessionToken, ok := os.LookupEnv(AWSSessionTokenName); ok {
		awsKeys[AWSSessionTokenName] = sessionToken
	}
	ctx = context.WithValue(
		ctx,
		ContextAWSVariables,
		awsKeys,
	)

	ctx = context.WithValue(
		ctx,
		ContextDelegatedToken,
		&DelegatedTokenCredentials{},
	)
	return ctx
}
