// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyUsageAttributionValues Fields in Usage Summary by tag(s).
// The following values have been **deprecated**: `estimated_indexed_spans_usage`, `estimated_indexed_spans_percentage`, `estimated_ingested_spans_usage`, `estimated_ingested_spans_percentage`.
type MonthlyUsageAttributionValues struct {
	// The percentage of synthetic API test usage by tag(s).
	ApiPercentage *float64 `json:"api_percentage,omitempty"`
	// The synthetic API test usage by tag(s).
	ApiUsage *float64 `json:"api_usage,omitempty"`
	// The percentage of APM ECS Fargate task usage by tag(s).
	ApmFargatePercentage *float64 `json:"apm_fargate_percentage,omitempty"`
	// The APM ECS Fargate task usage by tag(s).
	ApmFargateUsage *float64 `json:"apm_fargate_usage,omitempty"`
	// The percentage of APM host usage by tag(s).
	ApmHostPercentage *float64 `json:"apm_host_percentage,omitempty"`
	// The APM host usage by tag(s).
	ApmHostUsage *float64 `json:"apm_host_usage,omitempty"`
	// The percentage of APM and Universal Service Monitoring host usage by tag(s).
	ApmUsmPercentage *float64 `json:"apm_usm_percentage,omitempty"`
	// The APM and Universal Service Monitoring host usage by tag(s).
	ApmUsmUsage *float64 `json:"apm_usm_usage,omitempty"`
	// The percentage of Application Security Monitoring ECS Fargate task usage by tag(s).
	AppsecFargatePercentage *float64 `json:"appsec_fargate_percentage,omitempty"`
	// The Application Security Monitoring ECS Fargate task usage by tag(s).
	AppsecFargateUsage *float64 `json:"appsec_fargate_usage,omitempty"`
	// The percentage of Application Security Monitoring host usage by tag(s).
	AppsecPercentage *float64 `json:"appsec_percentage,omitempty"`
	// The Application Security Monitoring host usage by tag(s).
	AppsecUsage *float64 `json:"appsec_usage,omitempty"`
	// The percentage of Application Security Monitoring Serverless traced invocations usage by tag(s).
	AsmServerlessTracedInvocationsPercentage *float64 `json:"asm_serverless_traced_invocations_percentage,omitempty"`
	// The Application Security Monitoring Serverless traced invocations usage by tag(s).
	AsmServerlessTracedInvocationsUsage *float64 `json:"asm_serverless_traced_invocations_usage,omitempty"`
	// The percentage of synthetic browser test usage by tag(s).
	BrowserPercentage *float64 `json:"browser_percentage,omitempty"`
	// The synthetic browser test usage by tag(s).
	BrowserUsage *float64 `json:"browser_usage,omitempty"`
	// The percentage of CI Pipeline Indexed Spans usage by tag(s).
	CiPipelineIndexedSpansPercentage *float64 `json:"ci_pipeline_indexed_spans_percentage,omitempty"`
	// The total CI Pipeline Indexed Spans usage by tag(s).
	CiPipelineIndexedSpansUsage *float64 `json:"ci_pipeline_indexed_spans_usage,omitempty"`
	// The percentage of CI Test Indexed Spans usage by tag(s).
	CiTestIndexedSpansPercentage *float64 `json:"ci_test_indexed_spans_percentage,omitempty"`
	// The total CI Test Indexed Spans usage by tag(s).
	CiTestIndexedSpansUsage *float64 `json:"ci_test_indexed_spans_usage,omitempty"`
	// The percentage of Git committers for Intelligent Test Runner usage by tag(s).
	CiVisibilityItrPercentage *float64 `json:"ci_visibility_itr_percentage,omitempty"`
	// The Git committers for Intelligent Test Runner usage by tag(s).
	CiVisibilityItrUsage *float64 `json:"ci_visibility_itr_usage,omitempty"`
	// The percentage of Cloud Security Information and Event Management usage by tag(s).
	CloudSiemPercentage *float64 `json:"cloud_siem_percentage,omitempty"`
	// The Cloud Security Information and Event Management usage by tag(s).
	CloudSiemUsage *float64 `json:"cloud_siem_usage,omitempty"`
	// The percentage of Code Security host usage by tags.
	CodeSecurityHostPercentage *float64 `json:"code_security_host_percentage,omitempty"`
	// The Code Security host usage by tags.
	CodeSecurityHostUsage *float64 `json:"code_security_host_usage,omitempty"`
	// The percentage of container usage without the Datadog Agent by tag(s).
	ContainerExclAgentPercentage *float64 `json:"container_excl_agent_percentage,omitempty"`
	// The container usage without the Datadog Agent by tag(s).
	ContainerExclAgentUsage *float64 `json:"container_excl_agent_usage,omitempty"`
	// The percentage of container usage by tag(s).
	ContainerPercentage *float64 `json:"container_percentage,omitempty"`
	// The container usage by tag(s).
	ContainerUsage *float64 `json:"container_usage,omitempty"`
	// The percentage of Cloud Security Management Pro container usage by tag(s).
	CspmContainersPercentage *float64 `json:"cspm_containers_percentage,omitempty"`
	// The Cloud Security Management Pro container usage by tag(s).
	CspmContainersUsage *float64 `json:"cspm_containers_usage,omitempty"`
	// The percentage of Cloud Security Management Pro host usage by tag(s).
	CspmHostsPercentage *float64 `json:"cspm_hosts_percentage,omitempty"`
	// The Cloud Security Management Pro host usage by tag(s).
	CspmHostsUsage *float64 `json:"cspm_hosts_usage,omitempty"`
	// The percentage of Custom Events usage by tag(s).
	CustomEventPercentage *float64 `json:"custom_event_percentage,omitempty"`
	// The total Custom Events usage by tag(s).
	CustomEventUsage *float64 `json:"custom_event_usage,omitempty"`
	// The percentage of ingested custom metrics usage by tag(s).
	CustomIngestedTimeseriesPercentage *float64 `json:"custom_ingested_timeseries_percentage,omitempty"`
	// The ingested custom metrics usage by tag(s).
	CustomIngestedTimeseriesUsage *float64 `json:"custom_ingested_timeseries_usage,omitempty"`
	// The percentage of indexed custom metrics usage by tag(s).
	CustomTimeseriesPercentage *float64 `json:"custom_timeseries_percentage,omitempty"`
	// The indexed custom metrics usage by tag(s).
	CustomTimeseriesUsage *float64 `json:"custom_timeseries_usage,omitempty"`
	// The percentage of Cloud Workload Security container usage by tag(s).
	CwsContainersPercentage *float64 `json:"cws_containers_percentage,omitempty"`
	// The Cloud Workload Security container usage by tag(s).
	CwsContainersUsage *float64 `json:"cws_containers_usage,omitempty"`
	// The percentage of Cloud Workload Security Fargate task usage by tag(s).
	CwsFargateTaskPercentage *float64 `json:"cws_fargate_task_percentage,omitempty"`
	// The Cloud Workload Security Fargate task usage by tag(s).
	CwsFargateTaskUsage *float64 `json:"cws_fargate_task_usage,omitempty"`
	// The percentage of Cloud Workload Security host usage by tag(s).
	CwsHostsPercentage *float64 `json:"cws_hosts_percentage,omitempty"`
	// The Cloud Workload Security host usage by tag(s).
	CwsHostsUsage *float64 `json:"cws_hosts_usage,omitempty"`
	// The Data Jobs Monitoring usage by tag(s).
	DataJobsMonitoringUsage *float64 `json:"data_jobs_monitoring_usage,omitempty"`
	// The Data Stream Monitoring usage by tag(s).
	DataStreamMonitoringUsage *float64 `json:"data_stream_monitoring_usage,omitempty"`
	// The percentage of Database Monitoring host usage by tag(s).
	DbmHostsPercentage *float64 `json:"dbm_hosts_percentage,omitempty"`
	// The Database Monitoring host usage by tag(s).
	DbmHostsUsage *float64 `json:"dbm_hosts_usage,omitempty"`
	// The percentage of Database Monitoring queries usage by tag(s).
	DbmQueriesPercentage *float64 `json:"dbm_queries_percentage,omitempty"`
	// The Database Monitoring queries usage by tag(s).
	DbmQueriesUsage *float64 `json:"dbm_queries_usage,omitempty"`
	// The percentage of error tracking events usage by tag(s).
	ErrorTrackingPercentage *float64 `json:"error_tracking_percentage,omitempty"`
	// The error tracking events usage by tag(s).
	ErrorTrackingUsage *float64 `json:"error_tracking_usage,omitempty"`
	// The percentage of estimated indexed spans usage by tag(s).
	EstimatedIndexedSpansPercentage *float64 `json:"estimated_indexed_spans_percentage,omitempty"`
	// The estimated indexed spans usage by tag(s).
	EstimatedIndexedSpansUsage *float64 `json:"estimated_indexed_spans_usage,omitempty"`
	// The percentage of estimated ingested spans usage by tag(s).
	EstimatedIngestedSpansPercentage *float64 `json:"estimated_ingested_spans_percentage,omitempty"`
	// The estimated ingested spans usage by tag(s).
	EstimatedIngestedSpansUsage *float64 `json:"estimated_ingested_spans_usage,omitempty"`
	// The percentage of Fargate usage by tags.
	FargatePercentage *float64 `json:"fargate_percentage,omitempty"`
	// The Fargate usage by tags.
	FargateUsage *float64 `json:"fargate_usage,omitempty"`
	// The percentage of Lambda function usage by tag(s).
	FunctionsPercentage *float64 `json:"functions_percentage,omitempty"`
	// The Lambda function usage by tag(s).
	FunctionsUsage *float64 `json:"functions_usage,omitempty"`
	// The percentage of Incident Management monthly active users usage by tag(s).
	IncidentManagementMonthlyActiveUsersPercentage *float64 `json:"incident_management_monthly_active_users_percentage,omitempty"`
	// The Incident Management monthly active users usage by tag(s).
	IncidentManagementMonthlyActiveUsersUsage *float64 `json:"incident_management_monthly_active_users_usage,omitempty"`
	// The percentage of APM Indexed Spans usage by tag(s).
	IndexedSpansPercentage *float64 `json:"indexed_spans_percentage,omitempty"`
	// The total APM Indexed Spans usage by tag(s).
	IndexedSpansUsage *float64 `json:"indexed_spans_usage,omitempty"`
	// The percentage of infrastructure host usage by tag(s).
	InfraHostPercentage *float64 `json:"infra_host_percentage,omitempty"`
	// The infrastructure host usage by tag(s).
	InfraHostUsage *float64 `json:"infra_host_usage,omitempty"`
	// The percentage of Ingested Logs usage by tag(s).
	IngestedLogsBytesPercentage *float64 `json:"ingested_logs_bytes_percentage,omitempty"`
	// The total Ingested Logs usage by tag(s).
	IngestedLogsBytesUsage *float64 `json:"ingested_logs_bytes_usage,omitempty"`
	// The percentage of APM Ingested Spans usage by tag(s).
	IngestedSpansBytesPercentage *float64 `json:"ingested_spans_bytes_percentage,omitempty"`
	// The total APM Ingested Spans usage by tag(s).
	IngestedSpansBytesUsage *float64 `json:"ingested_spans_bytes_usage,omitempty"`
	// The percentage of Lambda invocation usage by tag(s).
	InvocationsPercentage *float64 `json:"invocations_percentage,omitempty"`
	// The Lambda invocation usage by tag(s).
	InvocationsUsage *float64 `json:"invocations_usage,omitempty"`
	// The percentage of Serverless APM usage by tag(s).
	LambdaTracedInvocationsPercentage *float64 `json:"lambda_traced_invocations_percentage,omitempty"`
	// The Serverless APM usage by tag(s).
	LambdaTracedInvocationsUsage *float64 `json:"lambda_traced_invocations_usage,omitempty"`
	// The percentage of LLM Observability usage by tag(s).
	LlmObservabilityPercentage *float64 `json:"llm_observability_percentage,omitempty"`
	// The LLM Observability usage by tag(s).
	LlmObservabilityUsage *float64 `json:"llm_observability_usage,omitempty"`
	// The percentage of Indexed Logs (15-day Retention) usage by tag(s).
	LogsIndexed15dayPercentage *float64 `json:"logs_indexed_15day_percentage,omitempty"`
	// The total Indexed Logs (15-day Retention) usage by tag(s).
	LogsIndexed15dayUsage *float64 `json:"logs_indexed_15day_usage,omitempty"`
	// The percentage of Indexed Logs (180-day Retention) usage by tag(s).
	LogsIndexed180dayPercentage *float64 `json:"logs_indexed_180day_percentage,omitempty"`
	// The total Indexed Logs (180-day Retention) usage by tag(s).
	LogsIndexed180dayUsage *float64 `json:"logs_indexed_180day_usage,omitempty"`
	// The percentage of Indexed Logs (1-day Retention) usage by tag(s).
	LogsIndexed1dayPercentage *float64 `json:"logs_indexed_1day_percentage,omitempty"`
	// The total Indexed Logs (1-day Retention) usage by tag(s).
	LogsIndexed1dayUsage *float64 `json:"logs_indexed_1day_usage,omitempty"`
	// The percentage of Indexed Logs (30-day Retention) usage by tag(s).
	LogsIndexed30dayPercentage *float64 `json:"logs_indexed_30day_percentage,omitempty"`
	// The total Indexed Logs (30-day Retention) usage by tag(s).
	LogsIndexed30dayUsage *float64 `json:"logs_indexed_30day_usage,omitempty"`
	// The percentage of Indexed Logs (360-day Retention) usage by tag(s).
	LogsIndexed360dayPercentage *float64 `json:"logs_indexed_360day_percentage,omitempty"`
	// The total Indexed Logs (360-day Retention) usage by tag(s).
	LogsIndexed360dayUsage *float64 `json:"logs_indexed_360day_usage,omitempty"`
	// The percentage of Indexed Logs (3-day Retention) usage by tag(s).
	LogsIndexed3dayPercentage *float64 `json:"logs_indexed_3day_percentage,omitempty"`
	// The total Indexed Logs (3-day Retention) usage by tag(s).
	LogsIndexed3dayUsage *float64 `json:"logs_indexed_3day_usage,omitempty"`
	// The percentage of Indexed Logs (45-day Retention) usage by tag(s).
	LogsIndexed45dayPercentage *float64 `json:"logs_indexed_45day_percentage,omitempty"`
	// The total Indexed Logs (45-day Retention) usage by tag(s).
	LogsIndexed45dayUsage *float64 `json:"logs_indexed_45day_usage,omitempty"`
	// The percentage of Indexed Logs (60-day Retention) usage by tag(s).
	LogsIndexed60dayPercentage *float64 `json:"logs_indexed_60day_percentage,omitempty"`
	// The total Indexed Logs (60-day Retention) usage by tag(s).
	LogsIndexed60dayUsage *float64 `json:"logs_indexed_60day_usage,omitempty"`
	// The percentage of Indexed Logs (7-day Retention) usage by tag(s).
	LogsIndexed7dayPercentage *float64 `json:"logs_indexed_7day_percentage,omitempty"`
	// The total Indexed Logs (7-day Retention) usage by tag(s).
	LogsIndexed7dayUsage *float64 `json:"logs_indexed_7day_usage,omitempty"`
	// The percentage of Indexed Logs (90-day Retention) usage by tag(s).
	LogsIndexed90dayPercentage *float64 `json:"logs_indexed_90day_percentage,omitempty"`
	// The total Indexed Logs (90-day Retention) usage by tag(s).
	LogsIndexed90dayUsage *float64 `json:"logs_indexed_90day_usage,omitempty"`
	// The percentage of Indexed Logs (Custom Retention) usage by tag(s).
	LogsIndexedCustomRetentionPercentage *float64 `json:"logs_indexed_custom_retention_percentage,omitempty"`
	// The total Indexed Logs (Custom Retention) usage by tag(s).
	LogsIndexedCustomRetentionUsage *float64 `json:"logs_indexed_custom_retention_usage,omitempty"`
	// The percentage of Synthetic mobile application test usage by tag(s).
	MobileAppTestingPercentage *float64 `json:"mobile_app_testing_percentage,omitempty"`
	// The Synthetic mobile application test usage by tag(s).
	MobileAppTestingUsage *float64 `json:"mobile_app_testing_usage,omitempty"`
	// The percentage of Network Device Monitoring NetFlow usage by tag(s).
	NdmNetflowPercentage *float64 `json:"ndm_netflow_percentage,omitempty"`
	// The Network Device Monitoring NetFlow usage by tag(s).
	NdmNetflowUsage *float64 `json:"ndm_netflow_usage,omitempty"`
	// The percentage of network device wireless usage by tag(s).
	NetworkDeviceWirelessPercentage *float64 `json:"network_device_wireless_percentage,omitempty"`
	// The network device wireless usage by tag(s).
	NetworkDeviceWirelessUsage *float64 `json:"network_device_wireless_usage,omitempty"`
	// The percentage of network host usage by tag(s).
	NpmHostPercentage *float64 `json:"npm_host_percentage,omitempty"`
	// The network host usage by tag(s).
	NpmHostUsage *float64 `json:"npm_host_usage,omitempty"`
	// The percentage of observability pipeline bytes usage by tag(s).
	ObsPipelineBytesPercentage *float64 `json:"obs_pipeline_bytes_percentage,omitempty"`
	// The observability pipeline bytes usage by tag(s).
	ObsPipelineBytesUsage *float64 `json:"obs_pipeline_bytes_usage,omitempty"`
	// The percentage of observability pipeline per core usage by tag(s).
	ObsPipelinesVcpuPercentage *float64 `json:"obs_pipelines_vcpu_percentage,omitempty"`
	// The observability pipeline per core usage by tag(s).
	ObsPipelinesVcpuUsage *float64 `json:"obs_pipelines_vcpu_usage,omitempty"`
	// The percentage of online archive usage by tag(s).
	OnlineArchivePercentage *float64 `json:"online_archive_percentage,omitempty"`
	// The online archive usage by tag(s).
	OnlineArchiveUsage *float64 `json:"online_archive_usage,omitempty"`
	// The percentage of Product Analytics session usage by tag(s).
	ProductAnalyticsSessionPercentage *float64 `json:"product_analytics_session_percentage,omitempty"`
	// The Product Analytics session usage by tag(s).
	ProductAnalyticsSessionUsage *float64 `json:"product_analytics_session_usage,omitempty"`
	// The percentage of profiled container usage by tag(s).
	ProfiledContainerPercentage *float64 `json:"profiled_container_percentage,omitempty"`
	// The profiled container usage by tag(s).
	ProfiledContainerUsage *float64 `json:"profiled_container_usage,omitempty"`
	// The percentage of profiled Fargate task usage by tag(s).
	ProfiledFargatePercentage *float64 `json:"profiled_fargate_percentage,omitempty"`
	// The profiled Fargate task usage by tag(s).
	ProfiledFargateUsage *float64 `json:"profiled_fargate_usage,omitempty"`
	// The percentage of profiled hosts usage by tag(s).
	ProfiledHostPercentage *float64 `json:"profiled_host_percentage,omitempty"`
	// The profiled hosts usage by tag(s).
	ProfiledHostUsage *float64 `json:"profiled_host_usage,omitempty"`
	// The percentage of published application usage by tag(s).
	PublishedAppPercentage *float64 `json:"published_app_percentage,omitempty"`
	// The published application usage by tag(s).
	PublishedAppUsage *float64 `json:"published_app_usage,omitempty"`
	// The percentage of RUM Browser and Mobile usage by tag(s).
	RumBrowserMobileSessionsPercentage *float64 `json:"rum_browser_mobile_sessions_percentage,omitempty"`
	// The total RUM Browser and Mobile usage by tag(s).
	RumBrowserMobileSessionsUsage *float64 `json:"rum_browser_mobile_sessions_usage,omitempty"`
	// The percentage of RUM Ingested usage by tag(s).
	RumIngestedPercentage *float64 `json:"rum_ingested_percentage,omitempty"`
	// The total RUM Ingested usage by tag(s).
	RumIngestedUsage *float64 `json:"rum_ingested_usage,omitempty"`
	// The percentage of RUM Investigate usage by tag(s).
	RumInvestigatePercentage *float64 `json:"rum_investigate_percentage,omitempty"`
	// The total RUM Investigate usage by tag(s).
	RumInvestigateUsage *float64 `json:"rum_investigate_usage,omitempty"`
	// The percentage of RUM Session Replay usage by tag(s).
	RumReplaySessionsPercentage *float64 `json:"rum_replay_sessions_percentage,omitempty"`
	// The total RUM Session Replay usage by tag(s).
	RumReplaySessionsUsage *float64 `json:"rum_replay_sessions_usage,omitempty"`
	// The percentage of RUM Session Replay Add-On usage by tag(s).
	RumSessionReplayAddOnPercentage *float64 `json:"rum_session_replay_add_on_percentage,omitempty"`
	// The total RUM Session Replay Add-On usage by tag(s).
	RumSessionReplayAddOnUsage *float64 `json:"rum_session_replay_add_on_usage,omitempty"`
	// The percentage of Software Composition Analysis Fargate task usage by tag(s).
	ScaFargatePercentage *float64 `json:"sca_fargate_percentage,omitempty"`
	// The total Software Composition Analysis Fargate task usage by tag(s).
	ScaFargateUsage *float64 `json:"sca_fargate_usage,omitempty"`
	// The percentage of Sensitive Data Scanner usage by tag(s).
	SdsScannedBytesPercentage *float64 `json:"sds_scanned_bytes_percentage,omitempty"`
	// The total Sensitive Data Scanner usage by tag(s).
	SdsScannedBytesUsage *float64 `json:"sds_scanned_bytes_usage,omitempty"`
	// The percentage of Serverless Apps usage by tag(s).
	ServerlessAppsPercentage *float64 `json:"serverless_apps_percentage,omitempty"`
	// The total Serverless Apps usage by tag(s).
	ServerlessAppsUsage *float64 `json:"serverless_apps_usage,omitempty"`
	// The percentage of log events analyzed by Cloud SIEM usage by tag(s).
	SiemAnalyzedLogsAddOnPercentage *float64 `json:"siem_analyzed_logs_add_on_percentage,omitempty"`
	// The log events analyzed by Cloud SIEM usage by tag(s).
	SiemAnalyzedLogsAddOnUsage *float64 `json:"siem_analyzed_logs_add_on_usage,omitempty"`
	// The percentage of SIEM usage by tag(s).
	SiemIngestedBytesPercentage *float64 `json:"siem_ingested_bytes_percentage,omitempty"`
	// The total SIEM usage by tag(s).
	SiemIngestedBytesUsage *float64 `json:"siem_ingested_bytes_usage,omitempty"`
	// The percentage of network device usage by tag(s).
	SnmpPercentage *float64 `json:"snmp_percentage,omitempty"`
	// The network device usage by tag(s).
	SnmpUsage *float64 `json:"snmp_usage,omitempty"`
	// The percentage of universal service monitoring usage by tag(s).
	UniversalServiceMonitoringPercentage *float64 `json:"universal_service_monitoring_percentage,omitempty"`
	// The universal service monitoring usage by tag(s).
	UniversalServiceMonitoringUsage *float64 `json:"universal_service_monitoring_usage,omitempty"`
	// The percentage of Application Vulnerability Management usage by tag(s).
	VulnManagementHostsPercentage *float64 `json:"vuln_management_hosts_percentage,omitempty"`
	// The Application Vulnerability Management usage by tag(s).
	VulnManagementHostsUsage *float64 `json:"vuln_management_hosts_usage,omitempty"`
	// The percentage of workflow executions usage by tag(s).
	WorkflowExecutionsPercentage *float64 `json:"workflow_executions_percentage,omitempty"`
	// The total workflow executions usage by tag(s).
	WorkflowExecutionsUsage *float64 `json:"workflow_executions_usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonthlyUsageAttributionValues instantiates a new MonthlyUsageAttributionValues object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyUsageAttributionValues() *MonthlyUsageAttributionValues {
	this := MonthlyUsageAttributionValues{}
	return &this
}

// NewMonthlyUsageAttributionValuesWithDefaults instantiates a new MonthlyUsageAttributionValues object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyUsageAttributionValuesWithDefaults() *MonthlyUsageAttributionValues {
	this := MonthlyUsageAttributionValues{}
	return &this
}

// GetApiPercentage returns the ApiPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApiPercentage() float64 {
	if o == nil || o.ApiPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ApiPercentage
}

// GetApiPercentageOk returns a tuple with the ApiPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApiPercentageOk() (*float64, bool) {
	if o == nil || o.ApiPercentage == nil {
		return nil, false
	}
	return o.ApiPercentage, true
}

// HasApiPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApiPercentage() bool {
	return o != nil && o.ApiPercentage != nil
}

// SetApiPercentage gets a reference to the given float64 and assigns it to the ApiPercentage field.
func (o *MonthlyUsageAttributionValues) SetApiPercentage(v float64) {
	o.ApiPercentage = &v
}

// GetApiUsage returns the ApiUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApiUsage() float64 {
	if o == nil || o.ApiUsage == nil {
		var ret float64
		return ret
	}
	return *o.ApiUsage
}

// GetApiUsageOk returns a tuple with the ApiUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApiUsageOk() (*float64, bool) {
	if o == nil || o.ApiUsage == nil {
		return nil, false
	}
	return o.ApiUsage, true
}

// HasApiUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApiUsage() bool {
	return o != nil && o.ApiUsage != nil
}

// SetApiUsage gets a reference to the given float64 and assigns it to the ApiUsage field.
func (o *MonthlyUsageAttributionValues) SetApiUsage(v float64) {
	o.ApiUsage = &v
}

// GetApmFargatePercentage returns the ApmFargatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmFargatePercentage() float64 {
	if o == nil || o.ApmFargatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.ApmFargatePercentage
}

// GetApmFargatePercentageOk returns a tuple with the ApmFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmFargatePercentageOk() (*float64, bool) {
	if o == nil || o.ApmFargatePercentage == nil {
		return nil, false
	}
	return o.ApmFargatePercentage, true
}

// HasApmFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmFargatePercentage() bool {
	return o != nil && o.ApmFargatePercentage != nil
}

// SetApmFargatePercentage gets a reference to the given float64 and assigns it to the ApmFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetApmFargatePercentage(v float64) {
	o.ApmFargatePercentage = &v
}

// GetApmFargateUsage returns the ApmFargateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmFargateUsage() float64 {
	if o == nil || o.ApmFargateUsage == nil {
		var ret float64
		return ret
	}
	return *o.ApmFargateUsage
}

// GetApmFargateUsageOk returns a tuple with the ApmFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmFargateUsageOk() (*float64, bool) {
	if o == nil || o.ApmFargateUsage == nil {
		return nil, false
	}
	return o.ApmFargateUsage, true
}

// HasApmFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmFargateUsage() bool {
	return o != nil && o.ApmFargateUsage != nil
}

// SetApmFargateUsage gets a reference to the given float64 and assigns it to the ApmFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetApmFargateUsage(v float64) {
	o.ApmFargateUsage = &v
}

// GetApmHostPercentage returns the ApmHostPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmHostPercentage() float64 {
	if o == nil || o.ApmHostPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ApmHostPercentage
}

// GetApmHostPercentageOk returns a tuple with the ApmHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmHostPercentageOk() (*float64, bool) {
	if o == nil || o.ApmHostPercentage == nil {
		return nil, false
	}
	return o.ApmHostPercentage, true
}

// HasApmHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmHostPercentage() bool {
	return o != nil && o.ApmHostPercentage != nil
}

// SetApmHostPercentage gets a reference to the given float64 and assigns it to the ApmHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetApmHostPercentage(v float64) {
	o.ApmHostPercentage = &v
}

// GetApmHostUsage returns the ApmHostUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmHostUsage() float64 {
	if o == nil || o.ApmHostUsage == nil {
		var ret float64
		return ret
	}
	return *o.ApmHostUsage
}

// GetApmHostUsageOk returns a tuple with the ApmHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmHostUsageOk() (*float64, bool) {
	if o == nil || o.ApmHostUsage == nil {
		return nil, false
	}
	return o.ApmHostUsage, true
}

// HasApmHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmHostUsage() bool {
	return o != nil && o.ApmHostUsage != nil
}

// SetApmHostUsage gets a reference to the given float64 and assigns it to the ApmHostUsage field.
func (o *MonthlyUsageAttributionValues) SetApmHostUsage(v float64) {
	o.ApmHostUsage = &v
}

// GetApmUsmPercentage returns the ApmUsmPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmUsmPercentage() float64 {
	if o == nil || o.ApmUsmPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ApmUsmPercentage
}

// GetApmUsmPercentageOk returns a tuple with the ApmUsmPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmUsmPercentageOk() (*float64, bool) {
	if o == nil || o.ApmUsmPercentage == nil {
		return nil, false
	}
	return o.ApmUsmPercentage, true
}

// HasApmUsmPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmUsmPercentage() bool {
	return o != nil && o.ApmUsmPercentage != nil
}

// SetApmUsmPercentage gets a reference to the given float64 and assigns it to the ApmUsmPercentage field.
func (o *MonthlyUsageAttributionValues) SetApmUsmPercentage(v float64) {
	o.ApmUsmPercentage = &v
}

// GetApmUsmUsage returns the ApmUsmUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetApmUsmUsage() float64 {
	if o == nil || o.ApmUsmUsage == nil {
		var ret float64
		return ret
	}
	return *o.ApmUsmUsage
}

// GetApmUsmUsageOk returns a tuple with the ApmUsmUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetApmUsmUsageOk() (*float64, bool) {
	if o == nil || o.ApmUsmUsage == nil {
		return nil, false
	}
	return o.ApmUsmUsage, true
}

// HasApmUsmUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasApmUsmUsage() bool {
	return o != nil && o.ApmUsmUsage != nil
}

// SetApmUsmUsage gets a reference to the given float64 and assigns it to the ApmUsmUsage field.
func (o *MonthlyUsageAttributionValues) SetApmUsmUsage(v float64) {
	o.ApmUsmUsage = &v
}

// GetAppsecFargatePercentage returns the AppsecFargatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAppsecFargatePercentage() float64 {
	if o == nil || o.AppsecFargatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.AppsecFargatePercentage
}

// GetAppsecFargatePercentageOk returns a tuple with the AppsecFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAppsecFargatePercentageOk() (*float64, bool) {
	if o == nil || o.AppsecFargatePercentage == nil {
		return nil, false
	}
	return o.AppsecFargatePercentage, true
}

// HasAppsecFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecFargatePercentage() bool {
	return o != nil && o.AppsecFargatePercentage != nil
}

// SetAppsecFargatePercentage gets a reference to the given float64 and assigns it to the AppsecFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetAppsecFargatePercentage(v float64) {
	o.AppsecFargatePercentage = &v
}

// GetAppsecFargateUsage returns the AppsecFargateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAppsecFargateUsage() float64 {
	if o == nil || o.AppsecFargateUsage == nil {
		var ret float64
		return ret
	}
	return *o.AppsecFargateUsage
}

// GetAppsecFargateUsageOk returns a tuple with the AppsecFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAppsecFargateUsageOk() (*float64, bool) {
	if o == nil || o.AppsecFargateUsage == nil {
		return nil, false
	}
	return o.AppsecFargateUsage, true
}

// HasAppsecFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecFargateUsage() bool {
	return o != nil && o.AppsecFargateUsage != nil
}

// SetAppsecFargateUsage gets a reference to the given float64 and assigns it to the AppsecFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetAppsecFargateUsage(v float64) {
	o.AppsecFargateUsage = &v
}

// GetAppsecPercentage returns the AppsecPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAppsecPercentage() float64 {
	if o == nil || o.AppsecPercentage == nil {
		var ret float64
		return ret
	}
	return *o.AppsecPercentage
}

// GetAppsecPercentageOk returns a tuple with the AppsecPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAppsecPercentageOk() (*float64, bool) {
	if o == nil || o.AppsecPercentage == nil {
		return nil, false
	}
	return o.AppsecPercentage, true
}

// HasAppsecPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecPercentage() bool {
	return o != nil && o.AppsecPercentage != nil
}

// SetAppsecPercentage gets a reference to the given float64 and assigns it to the AppsecPercentage field.
func (o *MonthlyUsageAttributionValues) SetAppsecPercentage(v float64) {
	o.AppsecPercentage = &v
}

// GetAppsecUsage returns the AppsecUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAppsecUsage() float64 {
	if o == nil || o.AppsecUsage == nil {
		var ret float64
		return ret
	}
	return *o.AppsecUsage
}

// GetAppsecUsageOk returns a tuple with the AppsecUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAppsecUsageOk() (*float64, bool) {
	if o == nil || o.AppsecUsage == nil {
		return nil, false
	}
	return o.AppsecUsage, true
}

// HasAppsecUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAppsecUsage() bool {
	return o != nil && o.AppsecUsage != nil
}

// SetAppsecUsage gets a reference to the given float64 and assigns it to the AppsecUsage field.
func (o *MonthlyUsageAttributionValues) SetAppsecUsage(v float64) {
	o.AppsecUsage = &v
}

// GetAsmServerlessTracedInvocationsPercentage returns the AsmServerlessTracedInvocationsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAsmServerlessTracedInvocationsPercentage() float64 {
	if o == nil || o.AsmServerlessTracedInvocationsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.AsmServerlessTracedInvocationsPercentage
}

// GetAsmServerlessTracedInvocationsPercentageOk returns a tuple with the AsmServerlessTracedInvocationsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAsmServerlessTracedInvocationsPercentageOk() (*float64, bool) {
	if o == nil || o.AsmServerlessTracedInvocationsPercentage == nil {
		return nil, false
	}
	return o.AsmServerlessTracedInvocationsPercentage, true
}

// HasAsmServerlessTracedInvocationsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAsmServerlessTracedInvocationsPercentage() bool {
	return o != nil && o.AsmServerlessTracedInvocationsPercentage != nil
}

// SetAsmServerlessTracedInvocationsPercentage gets a reference to the given float64 and assigns it to the AsmServerlessTracedInvocationsPercentage field.
func (o *MonthlyUsageAttributionValues) SetAsmServerlessTracedInvocationsPercentage(v float64) {
	o.AsmServerlessTracedInvocationsPercentage = &v
}

// GetAsmServerlessTracedInvocationsUsage returns the AsmServerlessTracedInvocationsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetAsmServerlessTracedInvocationsUsage() float64 {
	if o == nil || o.AsmServerlessTracedInvocationsUsage == nil {
		var ret float64
		return ret
	}
	return *o.AsmServerlessTracedInvocationsUsage
}

// GetAsmServerlessTracedInvocationsUsageOk returns a tuple with the AsmServerlessTracedInvocationsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetAsmServerlessTracedInvocationsUsageOk() (*float64, bool) {
	if o == nil || o.AsmServerlessTracedInvocationsUsage == nil {
		return nil, false
	}
	return o.AsmServerlessTracedInvocationsUsage, true
}

// HasAsmServerlessTracedInvocationsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasAsmServerlessTracedInvocationsUsage() bool {
	return o != nil && o.AsmServerlessTracedInvocationsUsage != nil
}

// SetAsmServerlessTracedInvocationsUsage gets a reference to the given float64 and assigns it to the AsmServerlessTracedInvocationsUsage field.
func (o *MonthlyUsageAttributionValues) SetAsmServerlessTracedInvocationsUsage(v float64) {
	o.AsmServerlessTracedInvocationsUsage = &v
}

// GetBrowserPercentage returns the BrowserPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetBrowserPercentage() float64 {
	if o == nil || o.BrowserPercentage == nil {
		var ret float64
		return ret
	}
	return *o.BrowserPercentage
}

// GetBrowserPercentageOk returns a tuple with the BrowserPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetBrowserPercentageOk() (*float64, bool) {
	if o == nil || o.BrowserPercentage == nil {
		return nil, false
	}
	return o.BrowserPercentage, true
}

// HasBrowserPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasBrowserPercentage() bool {
	return o != nil && o.BrowserPercentage != nil
}

// SetBrowserPercentage gets a reference to the given float64 and assigns it to the BrowserPercentage field.
func (o *MonthlyUsageAttributionValues) SetBrowserPercentage(v float64) {
	o.BrowserPercentage = &v
}

// GetBrowserUsage returns the BrowserUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetBrowserUsage() float64 {
	if o == nil || o.BrowserUsage == nil {
		var ret float64
		return ret
	}
	return *o.BrowserUsage
}

// GetBrowserUsageOk returns a tuple with the BrowserUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetBrowserUsageOk() (*float64, bool) {
	if o == nil || o.BrowserUsage == nil {
		return nil, false
	}
	return o.BrowserUsage, true
}

// HasBrowserUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasBrowserUsage() bool {
	return o != nil && o.BrowserUsage != nil
}

// SetBrowserUsage gets a reference to the given float64 and assigns it to the BrowserUsage field.
func (o *MonthlyUsageAttributionValues) SetBrowserUsage(v float64) {
	o.BrowserUsage = &v
}

// GetCiPipelineIndexedSpansPercentage returns the CiPipelineIndexedSpansPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCiPipelineIndexedSpansPercentage() float64 {
	if o == nil || o.CiPipelineIndexedSpansPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CiPipelineIndexedSpansPercentage
}

// GetCiPipelineIndexedSpansPercentageOk returns a tuple with the CiPipelineIndexedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCiPipelineIndexedSpansPercentageOk() (*float64, bool) {
	if o == nil || o.CiPipelineIndexedSpansPercentage == nil {
		return nil, false
	}
	return o.CiPipelineIndexedSpansPercentage, true
}

// HasCiPipelineIndexedSpansPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCiPipelineIndexedSpansPercentage() bool {
	return o != nil && o.CiPipelineIndexedSpansPercentage != nil
}

// SetCiPipelineIndexedSpansPercentage gets a reference to the given float64 and assigns it to the CiPipelineIndexedSpansPercentage field.
func (o *MonthlyUsageAttributionValues) SetCiPipelineIndexedSpansPercentage(v float64) {
	o.CiPipelineIndexedSpansPercentage = &v
}

// GetCiPipelineIndexedSpansUsage returns the CiPipelineIndexedSpansUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCiPipelineIndexedSpansUsage() float64 {
	if o == nil || o.CiPipelineIndexedSpansUsage == nil {
		var ret float64
		return ret
	}
	return *o.CiPipelineIndexedSpansUsage
}

// GetCiPipelineIndexedSpansUsageOk returns a tuple with the CiPipelineIndexedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCiPipelineIndexedSpansUsageOk() (*float64, bool) {
	if o == nil || o.CiPipelineIndexedSpansUsage == nil {
		return nil, false
	}
	return o.CiPipelineIndexedSpansUsage, true
}

// HasCiPipelineIndexedSpansUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCiPipelineIndexedSpansUsage() bool {
	return o != nil && o.CiPipelineIndexedSpansUsage != nil
}

// SetCiPipelineIndexedSpansUsage gets a reference to the given float64 and assigns it to the CiPipelineIndexedSpansUsage field.
func (o *MonthlyUsageAttributionValues) SetCiPipelineIndexedSpansUsage(v float64) {
	o.CiPipelineIndexedSpansUsage = &v
}

// GetCiTestIndexedSpansPercentage returns the CiTestIndexedSpansPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCiTestIndexedSpansPercentage() float64 {
	if o == nil || o.CiTestIndexedSpansPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CiTestIndexedSpansPercentage
}

// GetCiTestIndexedSpansPercentageOk returns a tuple with the CiTestIndexedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCiTestIndexedSpansPercentageOk() (*float64, bool) {
	if o == nil || o.CiTestIndexedSpansPercentage == nil {
		return nil, false
	}
	return o.CiTestIndexedSpansPercentage, true
}

// HasCiTestIndexedSpansPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCiTestIndexedSpansPercentage() bool {
	return o != nil && o.CiTestIndexedSpansPercentage != nil
}

// SetCiTestIndexedSpansPercentage gets a reference to the given float64 and assigns it to the CiTestIndexedSpansPercentage field.
func (o *MonthlyUsageAttributionValues) SetCiTestIndexedSpansPercentage(v float64) {
	o.CiTestIndexedSpansPercentage = &v
}

// GetCiTestIndexedSpansUsage returns the CiTestIndexedSpansUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCiTestIndexedSpansUsage() float64 {
	if o == nil || o.CiTestIndexedSpansUsage == nil {
		var ret float64
		return ret
	}
	return *o.CiTestIndexedSpansUsage
}

// GetCiTestIndexedSpansUsageOk returns a tuple with the CiTestIndexedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCiTestIndexedSpansUsageOk() (*float64, bool) {
	if o == nil || o.CiTestIndexedSpansUsage == nil {
		return nil, false
	}
	return o.CiTestIndexedSpansUsage, true
}

// HasCiTestIndexedSpansUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCiTestIndexedSpansUsage() bool {
	return o != nil && o.CiTestIndexedSpansUsage != nil
}

// SetCiTestIndexedSpansUsage gets a reference to the given float64 and assigns it to the CiTestIndexedSpansUsage field.
func (o *MonthlyUsageAttributionValues) SetCiTestIndexedSpansUsage(v float64) {
	o.CiTestIndexedSpansUsage = &v
}

// GetCiVisibilityItrPercentage returns the CiVisibilityItrPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCiVisibilityItrPercentage() float64 {
	if o == nil || o.CiVisibilityItrPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CiVisibilityItrPercentage
}

// GetCiVisibilityItrPercentageOk returns a tuple with the CiVisibilityItrPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCiVisibilityItrPercentageOk() (*float64, bool) {
	if o == nil || o.CiVisibilityItrPercentage == nil {
		return nil, false
	}
	return o.CiVisibilityItrPercentage, true
}

// HasCiVisibilityItrPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCiVisibilityItrPercentage() bool {
	return o != nil && o.CiVisibilityItrPercentage != nil
}

// SetCiVisibilityItrPercentage gets a reference to the given float64 and assigns it to the CiVisibilityItrPercentage field.
func (o *MonthlyUsageAttributionValues) SetCiVisibilityItrPercentage(v float64) {
	o.CiVisibilityItrPercentage = &v
}

// GetCiVisibilityItrUsage returns the CiVisibilityItrUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCiVisibilityItrUsage() float64 {
	if o == nil || o.CiVisibilityItrUsage == nil {
		var ret float64
		return ret
	}
	return *o.CiVisibilityItrUsage
}

// GetCiVisibilityItrUsageOk returns a tuple with the CiVisibilityItrUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCiVisibilityItrUsageOk() (*float64, bool) {
	if o == nil || o.CiVisibilityItrUsage == nil {
		return nil, false
	}
	return o.CiVisibilityItrUsage, true
}

// HasCiVisibilityItrUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCiVisibilityItrUsage() bool {
	return o != nil && o.CiVisibilityItrUsage != nil
}

// SetCiVisibilityItrUsage gets a reference to the given float64 and assigns it to the CiVisibilityItrUsage field.
func (o *MonthlyUsageAttributionValues) SetCiVisibilityItrUsage(v float64) {
	o.CiVisibilityItrUsage = &v
}

// GetCloudSiemPercentage returns the CloudSiemPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCloudSiemPercentage() float64 {
	if o == nil || o.CloudSiemPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CloudSiemPercentage
}

// GetCloudSiemPercentageOk returns a tuple with the CloudSiemPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCloudSiemPercentageOk() (*float64, bool) {
	if o == nil || o.CloudSiemPercentage == nil {
		return nil, false
	}
	return o.CloudSiemPercentage, true
}

// HasCloudSiemPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCloudSiemPercentage() bool {
	return o != nil && o.CloudSiemPercentage != nil
}

// SetCloudSiemPercentage gets a reference to the given float64 and assigns it to the CloudSiemPercentage field.
func (o *MonthlyUsageAttributionValues) SetCloudSiemPercentage(v float64) {
	o.CloudSiemPercentage = &v
}

// GetCloudSiemUsage returns the CloudSiemUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCloudSiemUsage() float64 {
	if o == nil || o.CloudSiemUsage == nil {
		var ret float64
		return ret
	}
	return *o.CloudSiemUsage
}

// GetCloudSiemUsageOk returns a tuple with the CloudSiemUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCloudSiemUsageOk() (*float64, bool) {
	if o == nil || o.CloudSiemUsage == nil {
		return nil, false
	}
	return o.CloudSiemUsage, true
}

// HasCloudSiemUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCloudSiemUsage() bool {
	return o != nil && o.CloudSiemUsage != nil
}

// SetCloudSiemUsage gets a reference to the given float64 and assigns it to the CloudSiemUsage field.
func (o *MonthlyUsageAttributionValues) SetCloudSiemUsage(v float64) {
	o.CloudSiemUsage = &v
}

// GetCodeSecurityHostPercentage returns the CodeSecurityHostPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCodeSecurityHostPercentage() float64 {
	if o == nil || o.CodeSecurityHostPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CodeSecurityHostPercentage
}

// GetCodeSecurityHostPercentageOk returns a tuple with the CodeSecurityHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCodeSecurityHostPercentageOk() (*float64, bool) {
	if o == nil || o.CodeSecurityHostPercentage == nil {
		return nil, false
	}
	return o.CodeSecurityHostPercentage, true
}

// HasCodeSecurityHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCodeSecurityHostPercentage() bool {
	return o != nil && o.CodeSecurityHostPercentage != nil
}

// SetCodeSecurityHostPercentage gets a reference to the given float64 and assigns it to the CodeSecurityHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetCodeSecurityHostPercentage(v float64) {
	o.CodeSecurityHostPercentage = &v
}

// GetCodeSecurityHostUsage returns the CodeSecurityHostUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCodeSecurityHostUsage() float64 {
	if o == nil || o.CodeSecurityHostUsage == nil {
		var ret float64
		return ret
	}
	return *o.CodeSecurityHostUsage
}

// GetCodeSecurityHostUsageOk returns a tuple with the CodeSecurityHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCodeSecurityHostUsageOk() (*float64, bool) {
	if o == nil || o.CodeSecurityHostUsage == nil {
		return nil, false
	}
	return o.CodeSecurityHostUsage, true
}

// HasCodeSecurityHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCodeSecurityHostUsage() bool {
	return o != nil && o.CodeSecurityHostUsage != nil
}

// SetCodeSecurityHostUsage gets a reference to the given float64 and assigns it to the CodeSecurityHostUsage field.
func (o *MonthlyUsageAttributionValues) SetCodeSecurityHostUsage(v float64) {
	o.CodeSecurityHostUsage = &v
}

// GetContainerExclAgentPercentage returns the ContainerExclAgentPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentPercentage() float64 {
	if o == nil || o.ContainerExclAgentPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ContainerExclAgentPercentage
}

// GetContainerExclAgentPercentageOk returns a tuple with the ContainerExclAgentPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentPercentageOk() (*float64, bool) {
	if o == nil || o.ContainerExclAgentPercentage == nil {
		return nil, false
	}
	return o.ContainerExclAgentPercentage, true
}

// HasContainerExclAgentPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerExclAgentPercentage() bool {
	return o != nil && o.ContainerExclAgentPercentage != nil
}

// SetContainerExclAgentPercentage gets a reference to the given float64 and assigns it to the ContainerExclAgentPercentage field.
func (o *MonthlyUsageAttributionValues) SetContainerExclAgentPercentage(v float64) {
	o.ContainerExclAgentPercentage = &v
}

// GetContainerExclAgentUsage returns the ContainerExclAgentUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentUsage() float64 {
	if o == nil || o.ContainerExclAgentUsage == nil {
		var ret float64
		return ret
	}
	return *o.ContainerExclAgentUsage
}

// GetContainerExclAgentUsageOk returns a tuple with the ContainerExclAgentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetContainerExclAgentUsageOk() (*float64, bool) {
	if o == nil || o.ContainerExclAgentUsage == nil {
		return nil, false
	}
	return o.ContainerExclAgentUsage, true
}

// HasContainerExclAgentUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerExclAgentUsage() bool {
	return o != nil && o.ContainerExclAgentUsage != nil
}

// SetContainerExclAgentUsage gets a reference to the given float64 and assigns it to the ContainerExclAgentUsage field.
func (o *MonthlyUsageAttributionValues) SetContainerExclAgentUsage(v float64) {
	o.ContainerExclAgentUsage = &v
}

// GetContainerPercentage returns the ContainerPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetContainerPercentage() float64 {
	if o == nil || o.ContainerPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ContainerPercentage
}

// GetContainerPercentageOk returns a tuple with the ContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetContainerPercentageOk() (*float64, bool) {
	if o == nil || o.ContainerPercentage == nil {
		return nil, false
	}
	return o.ContainerPercentage, true
}

// HasContainerPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerPercentage() bool {
	return o != nil && o.ContainerPercentage != nil
}

// SetContainerPercentage gets a reference to the given float64 and assigns it to the ContainerPercentage field.
func (o *MonthlyUsageAttributionValues) SetContainerPercentage(v float64) {
	o.ContainerPercentage = &v
}

// GetContainerUsage returns the ContainerUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetContainerUsage() float64 {
	if o == nil || o.ContainerUsage == nil {
		var ret float64
		return ret
	}
	return *o.ContainerUsage
}

// GetContainerUsageOk returns a tuple with the ContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetContainerUsageOk() (*float64, bool) {
	if o == nil || o.ContainerUsage == nil {
		return nil, false
	}
	return o.ContainerUsage, true
}

// HasContainerUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasContainerUsage() bool {
	return o != nil && o.ContainerUsage != nil
}

// SetContainerUsage gets a reference to the given float64 and assigns it to the ContainerUsage field.
func (o *MonthlyUsageAttributionValues) SetContainerUsage(v float64) {
	o.ContainerUsage = &v
}

// GetCspmContainersPercentage returns the CspmContainersPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCspmContainersPercentage() float64 {
	if o == nil || o.CspmContainersPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CspmContainersPercentage
}

// GetCspmContainersPercentageOk returns a tuple with the CspmContainersPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCspmContainersPercentageOk() (*float64, bool) {
	if o == nil || o.CspmContainersPercentage == nil {
		return nil, false
	}
	return o.CspmContainersPercentage, true
}

// HasCspmContainersPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmContainersPercentage() bool {
	return o != nil && o.CspmContainersPercentage != nil
}

// SetCspmContainersPercentage gets a reference to the given float64 and assigns it to the CspmContainersPercentage field.
func (o *MonthlyUsageAttributionValues) SetCspmContainersPercentage(v float64) {
	o.CspmContainersPercentage = &v
}

// GetCspmContainersUsage returns the CspmContainersUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCspmContainersUsage() float64 {
	if o == nil || o.CspmContainersUsage == nil {
		var ret float64
		return ret
	}
	return *o.CspmContainersUsage
}

// GetCspmContainersUsageOk returns a tuple with the CspmContainersUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCspmContainersUsageOk() (*float64, bool) {
	if o == nil || o.CspmContainersUsage == nil {
		return nil, false
	}
	return o.CspmContainersUsage, true
}

// HasCspmContainersUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmContainersUsage() bool {
	return o != nil && o.CspmContainersUsage != nil
}

// SetCspmContainersUsage gets a reference to the given float64 and assigns it to the CspmContainersUsage field.
func (o *MonthlyUsageAttributionValues) SetCspmContainersUsage(v float64) {
	o.CspmContainersUsage = &v
}

// GetCspmHostsPercentage returns the CspmHostsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCspmHostsPercentage() float64 {
	if o == nil || o.CspmHostsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CspmHostsPercentage
}

// GetCspmHostsPercentageOk returns a tuple with the CspmHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCspmHostsPercentageOk() (*float64, bool) {
	if o == nil || o.CspmHostsPercentage == nil {
		return nil, false
	}
	return o.CspmHostsPercentage, true
}

// HasCspmHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmHostsPercentage() bool {
	return o != nil && o.CspmHostsPercentage != nil
}

// SetCspmHostsPercentage gets a reference to the given float64 and assigns it to the CspmHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetCspmHostsPercentage(v float64) {
	o.CspmHostsPercentage = &v
}

// GetCspmHostsUsage returns the CspmHostsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCspmHostsUsage() float64 {
	if o == nil || o.CspmHostsUsage == nil {
		var ret float64
		return ret
	}
	return *o.CspmHostsUsage
}

// GetCspmHostsUsageOk returns a tuple with the CspmHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCspmHostsUsageOk() (*float64, bool) {
	if o == nil || o.CspmHostsUsage == nil {
		return nil, false
	}
	return o.CspmHostsUsage, true
}

// HasCspmHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCspmHostsUsage() bool {
	return o != nil && o.CspmHostsUsage != nil
}

// SetCspmHostsUsage gets a reference to the given float64 and assigns it to the CspmHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetCspmHostsUsage(v float64) {
	o.CspmHostsUsage = &v
}

// GetCustomEventPercentage returns the CustomEventPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomEventPercentage() float64 {
	if o == nil || o.CustomEventPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CustomEventPercentage
}

// GetCustomEventPercentageOk returns a tuple with the CustomEventPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomEventPercentageOk() (*float64, bool) {
	if o == nil || o.CustomEventPercentage == nil {
		return nil, false
	}
	return o.CustomEventPercentage, true
}

// HasCustomEventPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomEventPercentage() bool {
	return o != nil && o.CustomEventPercentage != nil
}

// SetCustomEventPercentage gets a reference to the given float64 and assigns it to the CustomEventPercentage field.
func (o *MonthlyUsageAttributionValues) SetCustomEventPercentage(v float64) {
	o.CustomEventPercentage = &v
}

// GetCustomEventUsage returns the CustomEventUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomEventUsage() float64 {
	if o == nil || o.CustomEventUsage == nil {
		var ret float64
		return ret
	}
	return *o.CustomEventUsage
}

// GetCustomEventUsageOk returns a tuple with the CustomEventUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomEventUsageOk() (*float64, bool) {
	if o == nil || o.CustomEventUsage == nil {
		return nil, false
	}
	return o.CustomEventUsage, true
}

// HasCustomEventUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomEventUsage() bool {
	return o != nil && o.CustomEventUsage != nil
}

// SetCustomEventUsage gets a reference to the given float64 and assigns it to the CustomEventUsage field.
func (o *MonthlyUsageAttributionValues) SetCustomEventUsage(v float64) {
	o.CustomEventUsage = &v
}

// GetCustomIngestedTimeseriesPercentage returns the CustomIngestedTimeseriesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesPercentage() float64 {
	if o == nil || o.CustomIngestedTimeseriesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CustomIngestedTimeseriesPercentage
}

// GetCustomIngestedTimeseriesPercentageOk returns a tuple with the CustomIngestedTimeseriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesPercentageOk() (*float64, bool) {
	if o == nil || o.CustomIngestedTimeseriesPercentage == nil {
		return nil, false
	}
	return o.CustomIngestedTimeseriesPercentage, true
}

// HasCustomIngestedTimeseriesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomIngestedTimeseriesPercentage() bool {
	return o != nil && o.CustomIngestedTimeseriesPercentage != nil
}

// SetCustomIngestedTimeseriesPercentage gets a reference to the given float64 and assigns it to the CustomIngestedTimeseriesPercentage field.
func (o *MonthlyUsageAttributionValues) SetCustomIngestedTimeseriesPercentage(v float64) {
	o.CustomIngestedTimeseriesPercentage = &v
}

// GetCustomIngestedTimeseriesUsage returns the CustomIngestedTimeseriesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesUsage() float64 {
	if o == nil || o.CustomIngestedTimeseriesUsage == nil {
		var ret float64
		return ret
	}
	return *o.CustomIngestedTimeseriesUsage
}

// GetCustomIngestedTimeseriesUsageOk returns a tuple with the CustomIngestedTimeseriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomIngestedTimeseriesUsageOk() (*float64, bool) {
	if o == nil || o.CustomIngestedTimeseriesUsage == nil {
		return nil, false
	}
	return o.CustomIngestedTimeseriesUsage, true
}

// HasCustomIngestedTimeseriesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomIngestedTimeseriesUsage() bool {
	return o != nil && o.CustomIngestedTimeseriesUsage != nil
}

// SetCustomIngestedTimeseriesUsage gets a reference to the given float64 and assigns it to the CustomIngestedTimeseriesUsage field.
func (o *MonthlyUsageAttributionValues) SetCustomIngestedTimeseriesUsage(v float64) {
	o.CustomIngestedTimeseriesUsage = &v
}

// GetCustomTimeseriesPercentage returns the CustomTimeseriesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesPercentage() float64 {
	if o == nil || o.CustomTimeseriesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CustomTimeseriesPercentage
}

// GetCustomTimeseriesPercentageOk returns a tuple with the CustomTimeseriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesPercentageOk() (*float64, bool) {
	if o == nil || o.CustomTimeseriesPercentage == nil {
		return nil, false
	}
	return o.CustomTimeseriesPercentage, true
}

// HasCustomTimeseriesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomTimeseriesPercentage() bool {
	return o != nil && o.CustomTimeseriesPercentage != nil
}

// SetCustomTimeseriesPercentage gets a reference to the given float64 and assigns it to the CustomTimeseriesPercentage field.
func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesPercentage(v float64) {
	o.CustomTimeseriesPercentage = &v
}

// GetCustomTimeseriesUsage returns the CustomTimeseriesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesUsage() float64 {
	if o == nil || o.CustomTimeseriesUsage == nil {
		var ret float64
		return ret
	}
	return *o.CustomTimeseriesUsage
}

// GetCustomTimeseriesUsageOk returns a tuple with the CustomTimeseriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCustomTimeseriesUsageOk() (*float64, bool) {
	if o == nil || o.CustomTimeseriesUsage == nil {
		return nil, false
	}
	return o.CustomTimeseriesUsage, true
}

// HasCustomTimeseriesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCustomTimeseriesUsage() bool {
	return o != nil && o.CustomTimeseriesUsage != nil
}

// SetCustomTimeseriesUsage gets a reference to the given float64 and assigns it to the CustomTimeseriesUsage field.
func (o *MonthlyUsageAttributionValues) SetCustomTimeseriesUsage(v float64) {
	o.CustomTimeseriesUsage = &v
}

// GetCwsContainersPercentage returns the CwsContainersPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsContainersPercentage() float64 {
	if o == nil || o.CwsContainersPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CwsContainersPercentage
}

// GetCwsContainersPercentageOk returns a tuple with the CwsContainersPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsContainersPercentageOk() (*float64, bool) {
	if o == nil || o.CwsContainersPercentage == nil {
		return nil, false
	}
	return o.CwsContainersPercentage, true
}

// HasCwsContainersPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsContainersPercentage() bool {
	return o != nil && o.CwsContainersPercentage != nil
}

// SetCwsContainersPercentage gets a reference to the given float64 and assigns it to the CwsContainersPercentage field.
func (o *MonthlyUsageAttributionValues) SetCwsContainersPercentage(v float64) {
	o.CwsContainersPercentage = &v
}

// GetCwsContainersUsage returns the CwsContainersUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsContainersUsage() float64 {
	if o == nil || o.CwsContainersUsage == nil {
		var ret float64
		return ret
	}
	return *o.CwsContainersUsage
}

// GetCwsContainersUsageOk returns a tuple with the CwsContainersUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsContainersUsageOk() (*float64, bool) {
	if o == nil || o.CwsContainersUsage == nil {
		return nil, false
	}
	return o.CwsContainersUsage, true
}

// HasCwsContainersUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsContainersUsage() bool {
	return o != nil && o.CwsContainersUsage != nil
}

// SetCwsContainersUsage gets a reference to the given float64 and assigns it to the CwsContainersUsage field.
func (o *MonthlyUsageAttributionValues) SetCwsContainersUsage(v float64) {
	o.CwsContainersUsage = &v
}

// GetCwsFargateTaskPercentage returns the CwsFargateTaskPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsFargateTaskPercentage() float64 {
	if o == nil || o.CwsFargateTaskPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CwsFargateTaskPercentage
}

// GetCwsFargateTaskPercentageOk returns a tuple with the CwsFargateTaskPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsFargateTaskPercentageOk() (*float64, bool) {
	if o == nil || o.CwsFargateTaskPercentage == nil {
		return nil, false
	}
	return o.CwsFargateTaskPercentage, true
}

// HasCwsFargateTaskPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsFargateTaskPercentage() bool {
	return o != nil && o.CwsFargateTaskPercentage != nil
}

// SetCwsFargateTaskPercentage gets a reference to the given float64 and assigns it to the CwsFargateTaskPercentage field.
func (o *MonthlyUsageAttributionValues) SetCwsFargateTaskPercentage(v float64) {
	o.CwsFargateTaskPercentage = &v
}

// GetCwsFargateTaskUsage returns the CwsFargateTaskUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsFargateTaskUsage() float64 {
	if o == nil || o.CwsFargateTaskUsage == nil {
		var ret float64
		return ret
	}
	return *o.CwsFargateTaskUsage
}

// GetCwsFargateTaskUsageOk returns a tuple with the CwsFargateTaskUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsFargateTaskUsageOk() (*float64, bool) {
	if o == nil || o.CwsFargateTaskUsage == nil {
		return nil, false
	}
	return o.CwsFargateTaskUsage, true
}

// HasCwsFargateTaskUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsFargateTaskUsage() bool {
	return o != nil && o.CwsFargateTaskUsage != nil
}

// SetCwsFargateTaskUsage gets a reference to the given float64 and assigns it to the CwsFargateTaskUsage field.
func (o *MonthlyUsageAttributionValues) SetCwsFargateTaskUsage(v float64) {
	o.CwsFargateTaskUsage = &v
}

// GetCwsHostsPercentage returns the CwsHostsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsHostsPercentage() float64 {
	if o == nil || o.CwsHostsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.CwsHostsPercentage
}

// GetCwsHostsPercentageOk returns a tuple with the CwsHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsHostsPercentageOk() (*float64, bool) {
	if o == nil || o.CwsHostsPercentage == nil {
		return nil, false
	}
	return o.CwsHostsPercentage, true
}

// HasCwsHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsHostsPercentage() bool {
	return o != nil && o.CwsHostsPercentage != nil
}

// SetCwsHostsPercentage gets a reference to the given float64 and assigns it to the CwsHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetCwsHostsPercentage(v float64) {
	o.CwsHostsPercentage = &v
}

// GetCwsHostsUsage returns the CwsHostsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetCwsHostsUsage() float64 {
	if o == nil || o.CwsHostsUsage == nil {
		var ret float64
		return ret
	}
	return *o.CwsHostsUsage
}

// GetCwsHostsUsageOk returns a tuple with the CwsHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetCwsHostsUsageOk() (*float64, bool) {
	if o == nil || o.CwsHostsUsage == nil {
		return nil, false
	}
	return o.CwsHostsUsage, true
}

// HasCwsHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasCwsHostsUsage() bool {
	return o != nil && o.CwsHostsUsage != nil
}

// SetCwsHostsUsage gets a reference to the given float64 and assigns it to the CwsHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetCwsHostsUsage(v float64) {
	o.CwsHostsUsage = &v
}

// GetDataJobsMonitoringUsage returns the DataJobsMonitoringUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDataJobsMonitoringUsage() float64 {
	if o == nil || o.DataJobsMonitoringUsage == nil {
		var ret float64
		return ret
	}
	return *o.DataJobsMonitoringUsage
}

// GetDataJobsMonitoringUsageOk returns a tuple with the DataJobsMonitoringUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDataJobsMonitoringUsageOk() (*float64, bool) {
	if o == nil || o.DataJobsMonitoringUsage == nil {
		return nil, false
	}
	return o.DataJobsMonitoringUsage, true
}

// HasDataJobsMonitoringUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDataJobsMonitoringUsage() bool {
	return o != nil && o.DataJobsMonitoringUsage != nil
}

// SetDataJobsMonitoringUsage gets a reference to the given float64 and assigns it to the DataJobsMonitoringUsage field.
func (o *MonthlyUsageAttributionValues) SetDataJobsMonitoringUsage(v float64) {
	o.DataJobsMonitoringUsage = &v
}

// GetDataStreamMonitoringUsage returns the DataStreamMonitoringUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDataStreamMonitoringUsage() float64 {
	if o == nil || o.DataStreamMonitoringUsage == nil {
		var ret float64
		return ret
	}
	return *o.DataStreamMonitoringUsage
}

// GetDataStreamMonitoringUsageOk returns a tuple with the DataStreamMonitoringUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDataStreamMonitoringUsageOk() (*float64, bool) {
	if o == nil || o.DataStreamMonitoringUsage == nil {
		return nil, false
	}
	return o.DataStreamMonitoringUsage, true
}

// HasDataStreamMonitoringUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDataStreamMonitoringUsage() bool {
	return o != nil && o.DataStreamMonitoringUsage != nil
}

// SetDataStreamMonitoringUsage gets a reference to the given float64 and assigns it to the DataStreamMonitoringUsage field.
func (o *MonthlyUsageAttributionValues) SetDataStreamMonitoringUsage(v float64) {
	o.DataStreamMonitoringUsage = &v
}

// GetDbmHostsPercentage returns the DbmHostsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDbmHostsPercentage() float64 {
	if o == nil || o.DbmHostsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.DbmHostsPercentage
}

// GetDbmHostsPercentageOk returns a tuple with the DbmHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDbmHostsPercentageOk() (*float64, bool) {
	if o == nil || o.DbmHostsPercentage == nil {
		return nil, false
	}
	return o.DbmHostsPercentage, true
}

// HasDbmHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmHostsPercentage() bool {
	return o != nil && o.DbmHostsPercentage != nil
}

// SetDbmHostsPercentage gets a reference to the given float64 and assigns it to the DbmHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetDbmHostsPercentage(v float64) {
	o.DbmHostsPercentage = &v
}

// GetDbmHostsUsage returns the DbmHostsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDbmHostsUsage() float64 {
	if o == nil || o.DbmHostsUsage == nil {
		var ret float64
		return ret
	}
	return *o.DbmHostsUsage
}

// GetDbmHostsUsageOk returns a tuple with the DbmHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDbmHostsUsageOk() (*float64, bool) {
	if o == nil || o.DbmHostsUsage == nil {
		return nil, false
	}
	return o.DbmHostsUsage, true
}

// HasDbmHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmHostsUsage() bool {
	return o != nil && o.DbmHostsUsage != nil
}

// SetDbmHostsUsage gets a reference to the given float64 and assigns it to the DbmHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetDbmHostsUsage(v float64) {
	o.DbmHostsUsage = &v
}

// GetDbmQueriesPercentage returns the DbmQueriesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesPercentage() float64 {
	if o == nil || o.DbmQueriesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.DbmQueriesPercentage
}

// GetDbmQueriesPercentageOk returns a tuple with the DbmQueriesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesPercentageOk() (*float64, bool) {
	if o == nil || o.DbmQueriesPercentage == nil {
		return nil, false
	}
	return o.DbmQueriesPercentage, true
}

// HasDbmQueriesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmQueriesPercentage() bool {
	return o != nil && o.DbmQueriesPercentage != nil
}

// SetDbmQueriesPercentage gets a reference to the given float64 and assigns it to the DbmQueriesPercentage field.
func (o *MonthlyUsageAttributionValues) SetDbmQueriesPercentage(v float64) {
	o.DbmQueriesPercentage = &v
}

// GetDbmQueriesUsage returns the DbmQueriesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesUsage() float64 {
	if o == nil || o.DbmQueriesUsage == nil {
		var ret float64
		return ret
	}
	return *o.DbmQueriesUsage
}

// GetDbmQueriesUsageOk returns a tuple with the DbmQueriesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetDbmQueriesUsageOk() (*float64, bool) {
	if o == nil || o.DbmQueriesUsage == nil {
		return nil, false
	}
	return o.DbmQueriesUsage, true
}

// HasDbmQueriesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasDbmQueriesUsage() bool {
	return o != nil && o.DbmQueriesUsage != nil
}

// SetDbmQueriesUsage gets a reference to the given float64 and assigns it to the DbmQueriesUsage field.
func (o *MonthlyUsageAttributionValues) SetDbmQueriesUsage(v float64) {
	o.DbmQueriesUsage = &v
}

// GetErrorTrackingPercentage returns the ErrorTrackingPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetErrorTrackingPercentage() float64 {
	if o == nil || o.ErrorTrackingPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ErrorTrackingPercentage
}

// GetErrorTrackingPercentageOk returns a tuple with the ErrorTrackingPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetErrorTrackingPercentageOk() (*float64, bool) {
	if o == nil || o.ErrorTrackingPercentage == nil {
		return nil, false
	}
	return o.ErrorTrackingPercentage, true
}

// HasErrorTrackingPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasErrorTrackingPercentage() bool {
	return o != nil && o.ErrorTrackingPercentage != nil
}

// SetErrorTrackingPercentage gets a reference to the given float64 and assigns it to the ErrorTrackingPercentage field.
func (o *MonthlyUsageAttributionValues) SetErrorTrackingPercentage(v float64) {
	o.ErrorTrackingPercentage = &v
}

// GetErrorTrackingUsage returns the ErrorTrackingUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetErrorTrackingUsage() float64 {
	if o == nil || o.ErrorTrackingUsage == nil {
		var ret float64
		return ret
	}
	return *o.ErrorTrackingUsage
}

// GetErrorTrackingUsageOk returns a tuple with the ErrorTrackingUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetErrorTrackingUsageOk() (*float64, bool) {
	if o == nil || o.ErrorTrackingUsage == nil {
		return nil, false
	}
	return o.ErrorTrackingUsage, true
}

// HasErrorTrackingUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasErrorTrackingUsage() bool {
	return o != nil && o.ErrorTrackingUsage != nil
}

// SetErrorTrackingUsage gets a reference to the given float64 and assigns it to the ErrorTrackingUsage field.
func (o *MonthlyUsageAttributionValues) SetErrorTrackingUsage(v float64) {
	o.ErrorTrackingUsage = &v
}

// GetEstimatedIndexedSpansPercentage returns the EstimatedIndexedSpansPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansPercentage() float64 {
	if o == nil || o.EstimatedIndexedSpansPercentage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedSpansPercentage
}

// GetEstimatedIndexedSpansPercentageOk returns a tuple with the EstimatedIndexedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansPercentageOk() (*float64, bool) {
	if o == nil || o.EstimatedIndexedSpansPercentage == nil {
		return nil, false
	}
	return o.EstimatedIndexedSpansPercentage, true
}

// HasEstimatedIndexedSpansPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedSpansPercentage() bool {
	return o != nil && o.EstimatedIndexedSpansPercentage != nil
}

// SetEstimatedIndexedSpansPercentage gets a reference to the given float64 and assigns it to the EstimatedIndexedSpansPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedSpansPercentage(v float64) {
	o.EstimatedIndexedSpansPercentage = &v
}

// GetEstimatedIndexedSpansUsage returns the EstimatedIndexedSpansUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansUsage() float64 {
	if o == nil || o.EstimatedIndexedSpansUsage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIndexedSpansUsage
}

// GetEstimatedIndexedSpansUsageOk returns a tuple with the EstimatedIndexedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIndexedSpansUsageOk() (*float64, bool) {
	if o == nil || o.EstimatedIndexedSpansUsage == nil {
		return nil, false
	}
	return o.EstimatedIndexedSpansUsage, true
}

// HasEstimatedIndexedSpansUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIndexedSpansUsage() bool {
	return o != nil && o.EstimatedIndexedSpansUsage != nil
}

// SetEstimatedIndexedSpansUsage gets a reference to the given float64 and assigns it to the EstimatedIndexedSpansUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIndexedSpansUsage(v float64) {
	o.EstimatedIndexedSpansUsage = &v
}

// GetEstimatedIngestedSpansPercentage returns the EstimatedIngestedSpansPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansPercentage() float64 {
	if o == nil || o.EstimatedIngestedSpansPercentage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedSpansPercentage
}

// GetEstimatedIngestedSpansPercentageOk returns a tuple with the EstimatedIngestedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansPercentageOk() (*float64, bool) {
	if o == nil || o.EstimatedIngestedSpansPercentage == nil {
		return nil, false
	}
	return o.EstimatedIngestedSpansPercentage, true
}

// HasEstimatedIngestedSpansPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedSpansPercentage() bool {
	return o != nil && o.EstimatedIngestedSpansPercentage != nil
}

// SetEstimatedIngestedSpansPercentage gets a reference to the given float64 and assigns it to the EstimatedIngestedSpansPercentage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedSpansPercentage(v float64) {
	o.EstimatedIngestedSpansPercentage = &v
}

// GetEstimatedIngestedSpansUsage returns the EstimatedIngestedSpansUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansUsage() float64 {
	if o == nil || o.EstimatedIngestedSpansUsage == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedIngestedSpansUsage
}

// GetEstimatedIngestedSpansUsageOk returns a tuple with the EstimatedIngestedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetEstimatedIngestedSpansUsageOk() (*float64, bool) {
	if o == nil || o.EstimatedIngestedSpansUsage == nil {
		return nil, false
	}
	return o.EstimatedIngestedSpansUsage, true
}

// HasEstimatedIngestedSpansUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasEstimatedIngestedSpansUsage() bool {
	return o != nil && o.EstimatedIngestedSpansUsage != nil
}

// SetEstimatedIngestedSpansUsage gets a reference to the given float64 and assigns it to the EstimatedIngestedSpansUsage field.
func (o *MonthlyUsageAttributionValues) SetEstimatedIngestedSpansUsage(v float64) {
	o.EstimatedIngestedSpansUsage = &v
}

// GetFargatePercentage returns the FargatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetFargatePercentage() float64 {
	if o == nil || o.FargatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.FargatePercentage
}

// GetFargatePercentageOk returns a tuple with the FargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetFargatePercentageOk() (*float64, bool) {
	if o == nil || o.FargatePercentage == nil {
		return nil, false
	}
	return o.FargatePercentage, true
}

// HasFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFargatePercentage() bool {
	return o != nil && o.FargatePercentage != nil
}

// SetFargatePercentage gets a reference to the given float64 and assigns it to the FargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetFargatePercentage(v float64) {
	o.FargatePercentage = &v
}

// GetFargateUsage returns the FargateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetFargateUsage() float64 {
	if o == nil || o.FargateUsage == nil {
		var ret float64
		return ret
	}
	return *o.FargateUsage
}

// GetFargateUsageOk returns a tuple with the FargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetFargateUsageOk() (*float64, bool) {
	if o == nil || o.FargateUsage == nil {
		return nil, false
	}
	return o.FargateUsage, true
}

// HasFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFargateUsage() bool {
	return o != nil && o.FargateUsage != nil
}

// SetFargateUsage gets a reference to the given float64 and assigns it to the FargateUsage field.
func (o *MonthlyUsageAttributionValues) SetFargateUsage(v float64) {
	o.FargateUsage = &v
}

// GetFunctionsPercentage returns the FunctionsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetFunctionsPercentage() float64 {
	if o == nil || o.FunctionsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.FunctionsPercentage
}

// GetFunctionsPercentageOk returns a tuple with the FunctionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetFunctionsPercentageOk() (*float64, bool) {
	if o == nil || o.FunctionsPercentage == nil {
		return nil, false
	}
	return o.FunctionsPercentage, true
}

// HasFunctionsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFunctionsPercentage() bool {
	return o != nil && o.FunctionsPercentage != nil
}

// SetFunctionsPercentage gets a reference to the given float64 and assigns it to the FunctionsPercentage field.
func (o *MonthlyUsageAttributionValues) SetFunctionsPercentage(v float64) {
	o.FunctionsPercentage = &v
}

// GetFunctionsUsage returns the FunctionsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetFunctionsUsage() float64 {
	if o == nil || o.FunctionsUsage == nil {
		var ret float64
		return ret
	}
	return *o.FunctionsUsage
}

// GetFunctionsUsageOk returns a tuple with the FunctionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetFunctionsUsageOk() (*float64, bool) {
	if o == nil || o.FunctionsUsage == nil {
		return nil, false
	}
	return o.FunctionsUsage, true
}

// HasFunctionsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasFunctionsUsage() bool {
	return o != nil && o.FunctionsUsage != nil
}

// SetFunctionsUsage gets a reference to the given float64 and assigns it to the FunctionsUsage field.
func (o *MonthlyUsageAttributionValues) SetFunctionsUsage(v float64) {
	o.FunctionsUsage = &v
}

// GetIncidentManagementMonthlyActiveUsersPercentage returns the IncidentManagementMonthlyActiveUsersPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetIncidentManagementMonthlyActiveUsersPercentage() float64 {
	if o == nil || o.IncidentManagementMonthlyActiveUsersPercentage == nil {
		var ret float64
		return ret
	}
	return *o.IncidentManagementMonthlyActiveUsersPercentage
}

// GetIncidentManagementMonthlyActiveUsersPercentageOk returns a tuple with the IncidentManagementMonthlyActiveUsersPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetIncidentManagementMonthlyActiveUsersPercentageOk() (*float64, bool) {
	if o == nil || o.IncidentManagementMonthlyActiveUsersPercentage == nil {
		return nil, false
	}
	return o.IncidentManagementMonthlyActiveUsersPercentage, true
}

// HasIncidentManagementMonthlyActiveUsersPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasIncidentManagementMonthlyActiveUsersPercentage() bool {
	return o != nil && o.IncidentManagementMonthlyActiveUsersPercentage != nil
}

// SetIncidentManagementMonthlyActiveUsersPercentage gets a reference to the given float64 and assigns it to the IncidentManagementMonthlyActiveUsersPercentage field.
func (o *MonthlyUsageAttributionValues) SetIncidentManagementMonthlyActiveUsersPercentage(v float64) {
	o.IncidentManagementMonthlyActiveUsersPercentage = &v
}

// GetIncidentManagementMonthlyActiveUsersUsage returns the IncidentManagementMonthlyActiveUsersUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetIncidentManagementMonthlyActiveUsersUsage() float64 {
	if o == nil || o.IncidentManagementMonthlyActiveUsersUsage == nil {
		var ret float64
		return ret
	}
	return *o.IncidentManagementMonthlyActiveUsersUsage
}

// GetIncidentManagementMonthlyActiveUsersUsageOk returns a tuple with the IncidentManagementMonthlyActiveUsersUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetIncidentManagementMonthlyActiveUsersUsageOk() (*float64, bool) {
	if o == nil || o.IncidentManagementMonthlyActiveUsersUsage == nil {
		return nil, false
	}
	return o.IncidentManagementMonthlyActiveUsersUsage, true
}

// HasIncidentManagementMonthlyActiveUsersUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasIncidentManagementMonthlyActiveUsersUsage() bool {
	return o != nil && o.IncidentManagementMonthlyActiveUsersUsage != nil
}

// SetIncidentManagementMonthlyActiveUsersUsage gets a reference to the given float64 and assigns it to the IncidentManagementMonthlyActiveUsersUsage field.
func (o *MonthlyUsageAttributionValues) SetIncidentManagementMonthlyActiveUsersUsage(v float64) {
	o.IncidentManagementMonthlyActiveUsersUsage = &v
}

// GetIndexedSpansPercentage returns the IndexedSpansPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetIndexedSpansPercentage() float64 {
	if o == nil || o.IndexedSpansPercentage == nil {
		var ret float64
		return ret
	}
	return *o.IndexedSpansPercentage
}

// GetIndexedSpansPercentageOk returns a tuple with the IndexedSpansPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetIndexedSpansPercentageOk() (*float64, bool) {
	if o == nil || o.IndexedSpansPercentage == nil {
		return nil, false
	}
	return o.IndexedSpansPercentage, true
}

// HasIndexedSpansPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasIndexedSpansPercentage() bool {
	return o != nil && o.IndexedSpansPercentage != nil
}

// SetIndexedSpansPercentage gets a reference to the given float64 and assigns it to the IndexedSpansPercentage field.
func (o *MonthlyUsageAttributionValues) SetIndexedSpansPercentage(v float64) {
	o.IndexedSpansPercentage = &v
}

// GetIndexedSpansUsage returns the IndexedSpansUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetIndexedSpansUsage() float64 {
	if o == nil || o.IndexedSpansUsage == nil {
		var ret float64
		return ret
	}
	return *o.IndexedSpansUsage
}

// GetIndexedSpansUsageOk returns a tuple with the IndexedSpansUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetIndexedSpansUsageOk() (*float64, bool) {
	if o == nil || o.IndexedSpansUsage == nil {
		return nil, false
	}
	return o.IndexedSpansUsage, true
}

// HasIndexedSpansUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasIndexedSpansUsage() bool {
	return o != nil && o.IndexedSpansUsage != nil
}

// SetIndexedSpansUsage gets a reference to the given float64 and assigns it to the IndexedSpansUsage field.
func (o *MonthlyUsageAttributionValues) SetIndexedSpansUsage(v float64) {
	o.IndexedSpansUsage = &v
}

// GetInfraHostPercentage returns the InfraHostPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetInfraHostPercentage() float64 {
	if o == nil || o.InfraHostPercentage == nil {
		var ret float64
		return ret
	}
	return *o.InfraHostPercentage
}

// GetInfraHostPercentageOk returns a tuple with the InfraHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetInfraHostPercentageOk() (*float64, bool) {
	if o == nil || o.InfraHostPercentage == nil {
		return nil, false
	}
	return o.InfraHostPercentage, true
}

// HasInfraHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInfraHostPercentage() bool {
	return o != nil && o.InfraHostPercentage != nil
}

// SetInfraHostPercentage gets a reference to the given float64 and assigns it to the InfraHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetInfraHostPercentage(v float64) {
	o.InfraHostPercentage = &v
}

// GetInfraHostUsage returns the InfraHostUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetInfraHostUsage() float64 {
	if o == nil || o.InfraHostUsage == nil {
		var ret float64
		return ret
	}
	return *o.InfraHostUsage
}

// GetInfraHostUsageOk returns a tuple with the InfraHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetInfraHostUsageOk() (*float64, bool) {
	if o == nil || o.InfraHostUsage == nil {
		return nil, false
	}
	return o.InfraHostUsage, true
}

// HasInfraHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInfraHostUsage() bool {
	return o != nil && o.InfraHostUsage != nil
}

// SetInfraHostUsage gets a reference to the given float64 and assigns it to the InfraHostUsage field.
func (o *MonthlyUsageAttributionValues) SetInfraHostUsage(v float64) {
	o.InfraHostUsage = &v
}

// GetIngestedLogsBytesPercentage returns the IngestedLogsBytesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetIngestedLogsBytesPercentage() float64 {
	if o == nil || o.IngestedLogsBytesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.IngestedLogsBytesPercentage
}

// GetIngestedLogsBytesPercentageOk returns a tuple with the IngestedLogsBytesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetIngestedLogsBytesPercentageOk() (*float64, bool) {
	if o == nil || o.IngestedLogsBytesPercentage == nil {
		return nil, false
	}
	return o.IngestedLogsBytesPercentage, true
}

// HasIngestedLogsBytesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasIngestedLogsBytesPercentage() bool {
	return o != nil && o.IngestedLogsBytesPercentage != nil
}

// SetIngestedLogsBytesPercentage gets a reference to the given float64 and assigns it to the IngestedLogsBytesPercentage field.
func (o *MonthlyUsageAttributionValues) SetIngestedLogsBytesPercentage(v float64) {
	o.IngestedLogsBytesPercentage = &v
}

// GetIngestedLogsBytesUsage returns the IngestedLogsBytesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetIngestedLogsBytesUsage() float64 {
	if o == nil || o.IngestedLogsBytesUsage == nil {
		var ret float64
		return ret
	}
	return *o.IngestedLogsBytesUsage
}

// GetIngestedLogsBytesUsageOk returns a tuple with the IngestedLogsBytesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetIngestedLogsBytesUsageOk() (*float64, bool) {
	if o == nil || o.IngestedLogsBytesUsage == nil {
		return nil, false
	}
	return o.IngestedLogsBytesUsage, true
}

// HasIngestedLogsBytesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasIngestedLogsBytesUsage() bool {
	return o != nil && o.IngestedLogsBytesUsage != nil
}

// SetIngestedLogsBytesUsage gets a reference to the given float64 and assigns it to the IngestedLogsBytesUsage field.
func (o *MonthlyUsageAttributionValues) SetIngestedLogsBytesUsage(v float64) {
	o.IngestedLogsBytesUsage = &v
}

// GetIngestedSpansBytesPercentage returns the IngestedSpansBytesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetIngestedSpansBytesPercentage() float64 {
	if o == nil || o.IngestedSpansBytesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.IngestedSpansBytesPercentage
}

// GetIngestedSpansBytesPercentageOk returns a tuple with the IngestedSpansBytesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetIngestedSpansBytesPercentageOk() (*float64, bool) {
	if o == nil || o.IngestedSpansBytesPercentage == nil {
		return nil, false
	}
	return o.IngestedSpansBytesPercentage, true
}

// HasIngestedSpansBytesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasIngestedSpansBytesPercentage() bool {
	return o != nil && o.IngestedSpansBytesPercentage != nil
}

// SetIngestedSpansBytesPercentage gets a reference to the given float64 and assigns it to the IngestedSpansBytesPercentage field.
func (o *MonthlyUsageAttributionValues) SetIngestedSpansBytesPercentage(v float64) {
	o.IngestedSpansBytesPercentage = &v
}

// GetIngestedSpansBytesUsage returns the IngestedSpansBytesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetIngestedSpansBytesUsage() float64 {
	if o == nil || o.IngestedSpansBytesUsage == nil {
		var ret float64
		return ret
	}
	return *o.IngestedSpansBytesUsage
}

// GetIngestedSpansBytesUsageOk returns a tuple with the IngestedSpansBytesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetIngestedSpansBytesUsageOk() (*float64, bool) {
	if o == nil || o.IngestedSpansBytesUsage == nil {
		return nil, false
	}
	return o.IngestedSpansBytesUsage, true
}

// HasIngestedSpansBytesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasIngestedSpansBytesUsage() bool {
	return o != nil && o.IngestedSpansBytesUsage != nil
}

// SetIngestedSpansBytesUsage gets a reference to the given float64 and assigns it to the IngestedSpansBytesUsage field.
func (o *MonthlyUsageAttributionValues) SetIngestedSpansBytesUsage(v float64) {
	o.IngestedSpansBytesUsage = &v
}

// GetInvocationsPercentage returns the InvocationsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetInvocationsPercentage() float64 {
	if o == nil || o.InvocationsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.InvocationsPercentage
}

// GetInvocationsPercentageOk returns a tuple with the InvocationsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetInvocationsPercentageOk() (*float64, bool) {
	if o == nil || o.InvocationsPercentage == nil {
		return nil, false
	}
	return o.InvocationsPercentage, true
}

// HasInvocationsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInvocationsPercentage() bool {
	return o != nil && o.InvocationsPercentage != nil
}

// SetInvocationsPercentage gets a reference to the given float64 and assigns it to the InvocationsPercentage field.
func (o *MonthlyUsageAttributionValues) SetInvocationsPercentage(v float64) {
	o.InvocationsPercentage = &v
}

// GetInvocationsUsage returns the InvocationsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetInvocationsUsage() float64 {
	if o == nil || o.InvocationsUsage == nil {
		var ret float64
		return ret
	}
	return *o.InvocationsUsage
}

// GetInvocationsUsageOk returns a tuple with the InvocationsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetInvocationsUsageOk() (*float64, bool) {
	if o == nil || o.InvocationsUsage == nil {
		return nil, false
	}
	return o.InvocationsUsage, true
}

// HasInvocationsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasInvocationsUsage() bool {
	return o != nil && o.InvocationsUsage != nil
}

// SetInvocationsUsage gets a reference to the given float64 and assigns it to the InvocationsUsage field.
func (o *MonthlyUsageAttributionValues) SetInvocationsUsage(v float64) {
	o.InvocationsUsage = &v
}

// GetLambdaTracedInvocationsPercentage returns the LambdaTracedInvocationsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLambdaTracedInvocationsPercentage() float64 {
	if o == nil || o.LambdaTracedInvocationsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LambdaTracedInvocationsPercentage
}

// GetLambdaTracedInvocationsPercentageOk returns a tuple with the LambdaTracedInvocationsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLambdaTracedInvocationsPercentageOk() (*float64, bool) {
	if o == nil || o.LambdaTracedInvocationsPercentage == nil {
		return nil, false
	}
	return o.LambdaTracedInvocationsPercentage, true
}

// HasLambdaTracedInvocationsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLambdaTracedInvocationsPercentage() bool {
	return o != nil && o.LambdaTracedInvocationsPercentage != nil
}

// SetLambdaTracedInvocationsPercentage gets a reference to the given float64 and assigns it to the LambdaTracedInvocationsPercentage field.
func (o *MonthlyUsageAttributionValues) SetLambdaTracedInvocationsPercentage(v float64) {
	o.LambdaTracedInvocationsPercentage = &v
}

// GetLambdaTracedInvocationsUsage returns the LambdaTracedInvocationsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLambdaTracedInvocationsUsage() float64 {
	if o == nil || o.LambdaTracedInvocationsUsage == nil {
		var ret float64
		return ret
	}
	return *o.LambdaTracedInvocationsUsage
}

// GetLambdaTracedInvocationsUsageOk returns a tuple with the LambdaTracedInvocationsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLambdaTracedInvocationsUsageOk() (*float64, bool) {
	if o == nil || o.LambdaTracedInvocationsUsage == nil {
		return nil, false
	}
	return o.LambdaTracedInvocationsUsage, true
}

// HasLambdaTracedInvocationsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLambdaTracedInvocationsUsage() bool {
	return o != nil && o.LambdaTracedInvocationsUsage != nil
}

// SetLambdaTracedInvocationsUsage gets a reference to the given float64 and assigns it to the LambdaTracedInvocationsUsage field.
func (o *MonthlyUsageAttributionValues) SetLambdaTracedInvocationsUsage(v float64) {
	o.LambdaTracedInvocationsUsage = &v
}

// GetLlmObservabilityPercentage returns the LlmObservabilityPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLlmObservabilityPercentage() float64 {
	if o == nil || o.LlmObservabilityPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LlmObservabilityPercentage
}

// GetLlmObservabilityPercentageOk returns a tuple with the LlmObservabilityPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLlmObservabilityPercentageOk() (*float64, bool) {
	if o == nil || o.LlmObservabilityPercentage == nil {
		return nil, false
	}
	return o.LlmObservabilityPercentage, true
}

// HasLlmObservabilityPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLlmObservabilityPercentage() bool {
	return o != nil && o.LlmObservabilityPercentage != nil
}

// SetLlmObservabilityPercentage gets a reference to the given float64 and assigns it to the LlmObservabilityPercentage field.
func (o *MonthlyUsageAttributionValues) SetLlmObservabilityPercentage(v float64) {
	o.LlmObservabilityPercentage = &v
}

// GetLlmObservabilityUsage returns the LlmObservabilityUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLlmObservabilityUsage() float64 {
	if o == nil || o.LlmObservabilityUsage == nil {
		var ret float64
		return ret
	}
	return *o.LlmObservabilityUsage
}

// GetLlmObservabilityUsageOk returns a tuple with the LlmObservabilityUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLlmObservabilityUsageOk() (*float64, bool) {
	if o == nil || o.LlmObservabilityUsage == nil {
		return nil, false
	}
	return o.LlmObservabilityUsage, true
}

// HasLlmObservabilityUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLlmObservabilityUsage() bool {
	return o != nil && o.LlmObservabilityUsage != nil
}

// SetLlmObservabilityUsage gets a reference to the given float64 and assigns it to the LlmObservabilityUsage field.
func (o *MonthlyUsageAttributionValues) SetLlmObservabilityUsage(v float64) {
	o.LlmObservabilityUsage = &v
}

// GetLogsIndexed15dayPercentage returns the LogsIndexed15dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed15dayPercentage() float64 {
	if o == nil || o.LogsIndexed15dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed15dayPercentage
}

// GetLogsIndexed15dayPercentageOk returns a tuple with the LogsIndexed15dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed15dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed15dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed15dayPercentage, true
}

// HasLogsIndexed15dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed15dayPercentage() bool {
	return o != nil && o.LogsIndexed15dayPercentage != nil
}

// SetLogsIndexed15dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed15dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed15dayPercentage(v float64) {
	o.LogsIndexed15dayPercentage = &v
}

// GetLogsIndexed15dayUsage returns the LogsIndexed15dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed15dayUsage() float64 {
	if o == nil || o.LogsIndexed15dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed15dayUsage
}

// GetLogsIndexed15dayUsageOk returns a tuple with the LogsIndexed15dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed15dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed15dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed15dayUsage, true
}

// HasLogsIndexed15dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed15dayUsage() bool {
	return o != nil && o.LogsIndexed15dayUsage != nil
}

// SetLogsIndexed15dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed15dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed15dayUsage(v float64) {
	o.LogsIndexed15dayUsage = &v
}

// GetLogsIndexed180dayPercentage returns the LogsIndexed180dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed180dayPercentage() float64 {
	if o == nil || o.LogsIndexed180dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed180dayPercentage
}

// GetLogsIndexed180dayPercentageOk returns a tuple with the LogsIndexed180dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed180dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed180dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed180dayPercentage, true
}

// HasLogsIndexed180dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed180dayPercentage() bool {
	return o != nil && o.LogsIndexed180dayPercentage != nil
}

// SetLogsIndexed180dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed180dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed180dayPercentage(v float64) {
	o.LogsIndexed180dayPercentage = &v
}

// GetLogsIndexed180dayUsage returns the LogsIndexed180dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed180dayUsage() float64 {
	if o == nil || o.LogsIndexed180dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed180dayUsage
}

// GetLogsIndexed180dayUsageOk returns a tuple with the LogsIndexed180dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed180dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed180dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed180dayUsage, true
}

// HasLogsIndexed180dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed180dayUsage() bool {
	return o != nil && o.LogsIndexed180dayUsage != nil
}

// SetLogsIndexed180dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed180dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed180dayUsage(v float64) {
	o.LogsIndexed180dayUsage = &v
}

// GetLogsIndexed1dayPercentage returns the LogsIndexed1dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed1dayPercentage() float64 {
	if o == nil || o.LogsIndexed1dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed1dayPercentage
}

// GetLogsIndexed1dayPercentageOk returns a tuple with the LogsIndexed1dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed1dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed1dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed1dayPercentage, true
}

// HasLogsIndexed1dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed1dayPercentage() bool {
	return o != nil && o.LogsIndexed1dayPercentage != nil
}

// SetLogsIndexed1dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed1dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed1dayPercentage(v float64) {
	o.LogsIndexed1dayPercentage = &v
}

// GetLogsIndexed1dayUsage returns the LogsIndexed1dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed1dayUsage() float64 {
	if o == nil || o.LogsIndexed1dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed1dayUsage
}

// GetLogsIndexed1dayUsageOk returns a tuple with the LogsIndexed1dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed1dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed1dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed1dayUsage, true
}

// HasLogsIndexed1dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed1dayUsage() bool {
	return o != nil && o.LogsIndexed1dayUsage != nil
}

// SetLogsIndexed1dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed1dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed1dayUsage(v float64) {
	o.LogsIndexed1dayUsage = &v
}

// GetLogsIndexed30dayPercentage returns the LogsIndexed30dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed30dayPercentage() float64 {
	if o == nil || o.LogsIndexed30dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed30dayPercentage
}

// GetLogsIndexed30dayPercentageOk returns a tuple with the LogsIndexed30dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed30dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed30dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed30dayPercentage, true
}

// HasLogsIndexed30dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed30dayPercentage() bool {
	return o != nil && o.LogsIndexed30dayPercentage != nil
}

// SetLogsIndexed30dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed30dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed30dayPercentage(v float64) {
	o.LogsIndexed30dayPercentage = &v
}

// GetLogsIndexed30dayUsage returns the LogsIndexed30dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed30dayUsage() float64 {
	if o == nil || o.LogsIndexed30dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed30dayUsage
}

// GetLogsIndexed30dayUsageOk returns a tuple with the LogsIndexed30dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed30dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed30dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed30dayUsage, true
}

// HasLogsIndexed30dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed30dayUsage() bool {
	return o != nil && o.LogsIndexed30dayUsage != nil
}

// SetLogsIndexed30dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed30dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed30dayUsage(v float64) {
	o.LogsIndexed30dayUsage = &v
}

// GetLogsIndexed360dayPercentage returns the LogsIndexed360dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed360dayPercentage() float64 {
	if o == nil || o.LogsIndexed360dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed360dayPercentage
}

// GetLogsIndexed360dayPercentageOk returns a tuple with the LogsIndexed360dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed360dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed360dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed360dayPercentage, true
}

// HasLogsIndexed360dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed360dayPercentage() bool {
	return o != nil && o.LogsIndexed360dayPercentage != nil
}

// SetLogsIndexed360dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed360dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed360dayPercentage(v float64) {
	o.LogsIndexed360dayPercentage = &v
}

// GetLogsIndexed360dayUsage returns the LogsIndexed360dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed360dayUsage() float64 {
	if o == nil || o.LogsIndexed360dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed360dayUsage
}

// GetLogsIndexed360dayUsageOk returns a tuple with the LogsIndexed360dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed360dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed360dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed360dayUsage, true
}

// HasLogsIndexed360dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed360dayUsage() bool {
	return o != nil && o.LogsIndexed360dayUsage != nil
}

// SetLogsIndexed360dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed360dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed360dayUsage(v float64) {
	o.LogsIndexed360dayUsage = &v
}

// GetLogsIndexed3dayPercentage returns the LogsIndexed3dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed3dayPercentage() float64 {
	if o == nil || o.LogsIndexed3dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed3dayPercentage
}

// GetLogsIndexed3dayPercentageOk returns a tuple with the LogsIndexed3dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed3dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed3dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed3dayPercentage, true
}

// HasLogsIndexed3dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed3dayPercentage() bool {
	return o != nil && o.LogsIndexed3dayPercentage != nil
}

// SetLogsIndexed3dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed3dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed3dayPercentage(v float64) {
	o.LogsIndexed3dayPercentage = &v
}

// GetLogsIndexed3dayUsage returns the LogsIndexed3dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed3dayUsage() float64 {
	if o == nil || o.LogsIndexed3dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed3dayUsage
}

// GetLogsIndexed3dayUsageOk returns a tuple with the LogsIndexed3dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed3dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed3dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed3dayUsage, true
}

// HasLogsIndexed3dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed3dayUsage() bool {
	return o != nil && o.LogsIndexed3dayUsage != nil
}

// SetLogsIndexed3dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed3dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed3dayUsage(v float64) {
	o.LogsIndexed3dayUsage = &v
}

// GetLogsIndexed45dayPercentage returns the LogsIndexed45dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed45dayPercentage() float64 {
	if o == nil || o.LogsIndexed45dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed45dayPercentage
}

// GetLogsIndexed45dayPercentageOk returns a tuple with the LogsIndexed45dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed45dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed45dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed45dayPercentage, true
}

// HasLogsIndexed45dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed45dayPercentage() bool {
	return o != nil && o.LogsIndexed45dayPercentage != nil
}

// SetLogsIndexed45dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed45dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed45dayPercentage(v float64) {
	o.LogsIndexed45dayPercentage = &v
}

// GetLogsIndexed45dayUsage returns the LogsIndexed45dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed45dayUsage() float64 {
	if o == nil || o.LogsIndexed45dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed45dayUsage
}

// GetLogsIndexed45dayUsageOk returns a tuple with the LogsIndexed45dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed45dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed45dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed45dayUsage, true
}

// HasLogsIndexed45dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed45dayUsage() bool {
	return o != nil && o.LogsIndexed45dayUsage != nil
}

// SetLogsIndexed45dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed45dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed45dayUsage(v float64) {
	o.LogsIndexed45dayUsage = &v
}

// GetLogsIndexed60dayPercentage returns the LogsIndexed60dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed60dayPercentage() float64 {
	if o == nil || o.LogsIndexed60dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed60dayPercentage
}

// GetLogsIndexed60dayPercentageOk returns a tuple with the LogsIndexed60dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed60dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed60dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed60dayPercentage, true
}

// HasLogsIndexed60dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed60dayPercentage() bool {
	return o != nil && o.LogsIndexed60dayPercentage != nil
}

// SetLogsIndexed60dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed60dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed60dayPercentage(v float64) {
	o.LogsIndexed60dayPercentage = &v
}

// GetLogsIndexed60dayUsage returns the LogsIndexed60dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed60dayUsage() float64 {
	if o == nil || o.LogsIndexed60dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed60dayUsage
}

// GetLogsIndexed60dayUsageOk returns a tuple with the LogsIndexed60dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed60dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed60dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed60dayUsage, true
}

// HasLogsIndexed60dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed60dayUsage() bool {
	return o != nil && o.LogsIndexed60dayUsage != nil
}

// SetLogsIndexed60dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed60dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed60dayUsage(v float64) {
	o.LogsIndexed60dayUsage = &v
}

// GetLogsIndexed7dayPercentage returns the LogsIndexed7dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed7dayPercentage() float64 {
	if o == nil || o.LogsIndexed7dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed7dayPercentage
}

// GetLogsIndexed7dayPercentageOk returns a tuple with the LogsIndexed7dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed7dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed7dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed7dayPercentage, true
}

// HasLogsIndexed7dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed7dayPercentage() bool {
	return o != nil && o.LogsIndexed7dayPercentage != nil
}

// SetLogsIndexed7dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed7dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed7dayPercentage(v float64) {
	o.LogsIndexed7dayPercentage = &v
}

// GetLogsIndexed7dayUsage returns the LogsIndexed7dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed7dayUsage() float64 {
	if o == nil || o.LogsIndexed7dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed7dayUsage
}

// GetLogsIndexed7dayUsageOk returns a tuple with the LogsIndexed7dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed7dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed7dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed7dayUsage, true
}

// HasLogsIndexed7dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed7dayUsage() bool {
	return o != nil && o.LogsIndexed7dayUsage != nil
}

// SetLogsIndexed7dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed7dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed7dayUsage(v float64) {
	o.LogsIndexed7dayUsage = &v
}

// GetLogsIndexed90dayPercentage returns the LogsIndexed90dayPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed90dayPercentage() float64 {
	if o == nil || o.LogsIndexed90dayPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed90dayPercentage
}

// GetLogsIndexed90dayPercentageOk returns a tuple with the LogsIndexed90dayPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed90dayPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed90dayPercentage == nil {
		return nil, false
	}
	return o.LogsIndexed90dayPercentage, true
}

// HasLogsIndexed90dayPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed90dayPercentage() bool {
	return o != nil && o.LogsIndexed90dayPercentage != nil
}

// SetLogsIndexed90dayPercentage gets a reference to the given float64 and assigns it to the LogsIndexed90dayPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed90dayPercentage(v float64) {
	o.LogsIndexed90dayPercentage = &v
}

// GetLogsIndexed90dayUsage returns the LogsIndexed90dayUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed90dayUsage() float64 {
	if o == nil || o.LogsIndexed90dayUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexed90dayUsage
}

// GetLogsIndexed90dayUsageOk returns a tuple with the LogsIndexed90dayUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexed90dayUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexed90dayUsage == nil {
		return nil, false
	}
	return o.LogsIndexed90dayUsage, true
}

// HasLogsIndexed90dayUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexed90dayUsage() bool {
	return o != nil && o.LogsIndexed90dayUsage != nil
}

// SetLogsIndexed90dayUsage gets a reference to the given float64 and assigns it to the LogsIndexed90dayUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexed90dayUsage(v float64) {
	o.LogsIndexed90dayUsage = &v
}

// GetLogsIndexedCustomRetentionPercentage returns the LogsIndexedCustomRetentionPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexedCustomRetentionPercentage() float64 {
	if o == nil || o.LogsIndexedCustomRetentionPercentage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexedCustomRetentionPercentage
}

// GetLogsIndexedCustomRetentionPercentageOk returns a tuple with the LogsIndexedCustomRetentionPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexedCustomRetentionPercentageOk() (*float64, bool) {
	if o == nil || o.LogsIndexedCustomRetentionPercentage == nil {
		return nil, false
	}
	return o.LogsIndexedCustomRetentionPercentage, true
}

// HasLogsIndexedCustomRetentionPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexedCustomRetentionPercentage() bool {
	return o != nil && o.LogsIndexedCustomRetentionPercentage != nil
}

// SetLogsIndexedCustomRetentionPercentage gets a reference to the given float64 and assigns it to the LogsIndexedCustomRetentionPercentage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexedCustomRetentionPercentage(v float64) {
	o.LogsIndexedCustomRetentionPercentage = &v
}

// GetLogsIndexedCustomRetentionUsage returns the LogsIndexedCustomRetentionUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetLogsIndexedCustomRetentionUsage() float64 {
	if o == nil || o.LogsIndexedCustomRetentionUsage == nil {
		var ret float64
		return ret
	}
	return *o.LogsIndexedCustomRetentionUsage
}

// GetLogsIndexedCustomRetentionUsageOk returns a tuple with the LogsIndexedCustomRetentionUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetLogsIndexedCustomRetentionUsageOk() (*float64, bool) {
	if o == nil || o.LogsIndexedCustomRetentionUsage == nil {
		return nil, false
	}
	return o.LogsIndexedCustomRetentionUsage, true
}

// HasLogsIndexedCustomRetentionUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasLogsIndexedCustomRetentionUsage() bool {
	return o != nil && o.LogsIndexedCustomRetentionUsage != nil
}

// SetLogsIndexedCustomRetentionUsage gets a reference to the given float64 and assigns it to the LogsIndexedCustomRetentionUsage field.
func (o *MonthlyUsageAttributionValues) SetLogsIndexedCustomRetentionUsage(v float64) {
	o.LogsIndexedCustomRetentionUsage = &v
}

// GetMobileAppTestingPercentage returns the MobileAppTestingPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetMobileAppTestingPercentage() float64 {
	if o == nil || o.MobileAppTestingPercentage == nil {
		var ret float64
		return ret
	}
	return *o.MobileAppTestingPercentage
}

// GetMobileAppTestingPercentageOk returns a tuple with the MobileAppTestingPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetMobileAppTestingPercentageOk() (*float64, bool) {
	if o == nil || o.MobileAppTestingPercentage == nil {
		return nil, false
	}
	return o.MobileAppTestingPercentage, true
}

// HasMobileAppTestingPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasMobileAppTestingPercentage() bool {
	return o != nil && o.MobileAppTestingPercentage != nil
}

// SetMobileAppTestingPercentage gets a reference to the given float64 and assigns it to the MobileAppTestingPercentage field.
func (o *MonthlyUsageAttributionValues) SetMobileAppTestingPercentage(v float64) {
	o.MobileAppTestingPercentage = &v
}

// GetMobileAppTestingUsage returns the MobileAppTestingUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetMobileAppTestingUsage() float64 {
	if o == nil || o.MobileAppTestingUsage == nil {
		var ret float64
		return ret
	}
	return *o.MobileAppTestingUsage
}

// GetMobileAppTestingUsageOk returns a tuple with the MobileAppTestingUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetMobileAppTestingUsageOk() (*float64, bool) {
	if o == nil || o.MobileAppTestingUsage == nil {
		return nil, false
	}
	return o.MobileAppTestingUsage, true
}

// HasMobileAppTestingUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasMobileAppTestingUsage() bool {
	return o != nil && o.MobileAppTestingUsage != nil
}

// SetMobileAppTestingUsage gets a reference to the given float64 and assigns it to the MobileAppTestingUsage field.
func (o *MonthlyUsageAttributionValues) SetMobileAppTestingUsage(v float64) {
	o.MobileAppTestingUsage = &v
}

// GetNdmNetflowPercentage returns the NdmNetflowPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNdmNetflowPercentage() float64 {
	if o == nil || o.NdmNetflowPercentage == nil {
		var ret float64
		return ret
	}
	return *o.NdmNetflowPercentage
}

// GetNdmNetflowPercentageOk returns a tuple with the NdmNetflowPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNdmNetflowPercentageOk() (*float64, bool) {
	if o == nil || o.NdmNetflowPercentage == nil {
		return nil, false
	}
	return o.NdmNetflowPercentage, true
}

// HasNdmNetflowPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNdmNetflowPercentage() bool {
	return o != nil && o.NdmNetflowPercentage != nil
}

// SetNdmNetflowPercentage gets a reference to the given float64 and assigns it to the NdmNetflowPercentage field.
func (o *MonthlyUsageAttributionValues) SetNdmNetflowPercentage(v float64) {
	o.NdmNetflowPercentage = &v
}

// GetNdmNetflowUsage returns the NdmNetflowUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNdmNetflowUsage() float64 {
	if o == nil || o.NdmNetflowUsage == nil {
		var ret float64
		return ret
	}
	return *o.NdmNetflowUsage
}

// GetNdmNetflowUsageOk returns a tuple with the NdmNetflowUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNdmNetflowUsageOk() (*float64, bool) {
	if o == nil || o.NdmNetflowUsage == nil {
		return nil, false
	}
	return o.NdmNetflowUsage, true
}

// HasNdmNetflowUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNdmNetflowUsage() bool {
	return o != nil && o.NdmNetflowUsage != nil
}

// SetNdmNetflowUsage gets a reference to the given float64 and assigns it to the NdmNetflowUsage field.
func (o *MonthlyUsageAttributionValues) SetNdmNetflowUsage(v float64) {
	o.NdmNetflowUsage = &v
}

// GetNetworkDeviceWirelessPercentage returns the NetworkDeviceWirelessPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNetworkDeviceWirelessPercentage() float64 {
	if o == nil || o.NetworkDeviceWirelessPercentage == nil {
		var ret float64
		return ret
	}
	return *o.NetworkDeviceWirelessPercentage
}

// GetNetworkDeviceWirelessPercentageOk returns a tuple with the NetworkDeviceWirelessPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNetworkDeviceWirelessPercentageOk() (*float64, bool) {
	if o == nil || o.NetworkDeviceWirelessPercentage == nil {
		return nil, false
	}
	return o.NetworkDeviceWirelessPercentage, true
}

// HasNetworkDeviceWirelessPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNetworkDeviceWirelessPercentage() bool {
	return o != nil && o.NetworkDeviceWirelessPercentage != nil
}

// SetNetworkDeviceWirelessPercentage gets a reference to the given float64 and assigns it to the NetworkDeviceWirelessPercentage field.
func (o *MonthlyUsageAttributionValues) SetNetworkDeviceWirelessPercentage(v float64) {
	o.NetworkDeviceWirelessPercentage = &v
}

// GetNetworkDeviceWirelessUsage returns the NetworkDeviceWirelessUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNetworkDeviceWirelessUsage() float64 {
	if o == nil || o.NetworkDeviceWirelessUsage == nil {
		var ret float64
		return ret
	}
	return *o.NetworkDeviceWirelessUsage
}

// GetNetworkDeviceWirelessUsageOk returns a tuple with the NetworkDeviceWirelessUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNetworkDeviceWirelessUsageOk() (*float64, bool) {
	if o == nil || o.NetworkDeviceWirelessUsage == nil {
		return nil, false
	}
	return o.NetworkDeviceWirelessUsage, true
}

// HasNetworkDeviceWirelessUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNetworkDeviceWirelessUsage() bool {
	return o != nil && o.NetworkDeviceWirelessUsage != nil
}

// SetNetworkDeviceWirelessUsage gets a reference to the given float64 and assigns it to the NetworkDeviceWirelessUsage field.
func (o *MonthlyUsageAttributionValues) SetNetworkDeviceWirelessUsage(v float64) {
	o.NetworkDeviceWirelessUsage = &v
}

// GetNpmHostPercentage returns the NpmHostPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNpmHostPercentage() float64 {
	if o == nil || o.NpmHostPercentage == nil {
		var ret float64
		return ret
	}
	return *o.NpmHostPercentage
}

// GetNpmHostPercentageOk returns a tuple with the NpmHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNpmHostPercentageOk() (*float64, bool) {
	if o == nil || o.NpmHostPercentage == nil {
		return nil, false
	}
	return o.NpmHostPercentage, true
}

// HasNpmHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNpmHostPercentage() bool {
	return o != nil && o.NpmHostPercentage != nil
}

// SetNpmHostPercentage gets a reference to the given float64 and assigns it to the NpmHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetNpmHostPercentage(v float64) {
	o.NpmHostPercentage = &v
}

// GetNpmHostUsage returns the NpmHostUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetNpmHostUsage() float64 {
	if o == nil || o.NpmHostUsage == nil {
		var ret float64
		return ret
	}
	return *o.NpmHostUsage
}

// GetNpmHostUsageOk returns a tuple with the NpmHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetNpmHostUsageOk() (*float64, bool) {
	if o == nil || o.NpmHostUsage == nil {
		return nil, false
	}
	return o.NpmHostUsage, true
}

// HasNpmHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasNpmHostUsage() bool {
	return o != nil && o.NpmHostUsage != nil
}

// SetNpmHostUsage gets a reference to the given float64 and assigns it to the NpmHostUsage field.
func (o *MonthlyUsageAttributionValues) SetNpmHostUsage(v float64) {
	o.NpmHostUsage = &v
}

// GetObsPipelineBytesPercentage returns the ObsPipelineBytesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetObsPipelineBytesPercentage() float64 {
	if o == nil || o.ObsPipelineBytesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ObsPipelineBytesPercentage
}

// GetObsPipelineBytesPercentageOk returns a tuple with the ObsPipelineBytesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetObsPipelineBytesPercentageOk() (*float64, bool) {
	if o == nil || o.ObsPipelineBytesPercentage == nil {
		return nil, false
	}
	return o.ObsPipelineBytesPercentage, true
}

// HasObsPipelineBytesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasObsPipelineBytesPercentage() bool {
	return o != nil && o.ObsPipelineBytesPercentage != nil
}

// SetObsPipelineBytesPercentage gets a reference to the given float64 and assigns it to the ObsPipelineBytesPercentage field.
func (o *MonthlyUsageAttributionValues) SetObsPipelineBytesPercentage(v float64) {
	o.ObsPipelineBytesPercentage = &v
}

// GetObsPipelineBytesUsage returns the ObsPipelineBytesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetObsPipelineBytesUsage() float64 {
	if o == nil || o.ObsPipelineBytesUsage == nil {
		var ret float64
		return ret
	}
	return *o.ObsPipelineBytesUsage
}

// GetObsPipelineBytesUsageOk returns a tuple with the ObsPipelineBytesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetObsPipelineBytesUsageOk() (*float64, bool) {
	if o == nil || o.ObsPipelineBytesUsage == nil {
		return nil, false
	}
	return o.ObsPipelineBytesUsage, true
}

// HasObsPipelineBytesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasObsPipelineBytesUsage() bool {
	return o != nil && o.ObsPipelineBytesUsage != nil
}

// SetObsPipelineBytesUsage gets a reference to the given float64 and assigns it to the ObsPipelineBytesUsage field.
func (o *MonthlyUsageAttributionValues) SetObsPipelineBytesUsage(v float64) {
	o.ObsPipelineBytesUsage = &v
}

// GetObsPipelinesVcpuPercentage returns the ObsPipelinesVcpuPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetObsPipelinesVcpuPercentage() float64 {
	if o == nil || o.ObsPipelinesVcpuPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ObsPipelinesVcpuPercentage
}

// GetObsPipelinesVcpuPercentageOk returns a tuple with the ObsPipelinesVcpuPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetObsPipelinesVcpuPercentageOk() (*float64, bool) {
	if o == nil || o.ObsPipelinesVcpuPercentage == nil {
		return nil, false
	}
	return o.ObsPipelinesVcpuPercentage, true
}

// HasObsPipelinesVcpuPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasObsPipelinesVcpuPercentage() bool {
	return o != nil && o.ObsPipelinesVcpuPercentage != nil
}

// SetObsPipelinesVcpuPercentage gets a reference to the given float64 and assigns it to the ObsPipelinesVcpuPercentage field.
func (o *MonthlyUsageAttributionValues) SetObsPipelinesVcpuPercentage(v float64) {
	o.ObsPipelinesVcpuPercentage = &v
}

// GetObsPipelinesVcpuUsage returns the ObsPipelinesVcpuUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetObsPipelinesVcpuUsage() float64 {
	if o == nil || o.ObsPipelinesVcpuUsage == nil {
		var ret float64
		return ret
	}
	return *o.ObsPipelinesVcpuUsage
}

// GetObsPipelinesVcpuUsageOk returns a tuple with the ObsPipelinesVcpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetObsPipelinesVcpuUsageOk() (*float64, bool) {
	if o == nil || o.ObsPipelinesVcpuUsage == nil {
		return nil, false
	}
	return o.ObsPipelinesVcpuUsage, true
}

// HasObsPipelinesVcpuUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasObsPipelinesVcpuUsage() bool {
	return o != nil && o.ObsPipelinesVcpuUsage != nil
}

// SetObsPipelinesVcpuUsage gets a reference to the given float64 and assigns it to the ObsPipelinesVcpuUsage field.
func (o *MonthlyUsageAttributionValues) SetObsPipelinesVcpuUsage(v float64) {
	o.ObsPipelinesVcpuUsage = &v
}

// GetOnlineArchivePercentage returns the OnlineArchivePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetOnlineArchivePercentage() float64 {
	if o == nil || o.OnlineArchivePercentage == nil {
		var ret float64
		return ret
	}
	return *o.OnlineArchivePercentage
}

// GetOnlineArchivePercentageOk returns a tuple with the OnlineArchivePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetOnlineArchivePercentageOk() (*float64, bool) {
	if o == nil || o.OnlineArchivePercentage == nil {
		return nil, false
	}
	return o.OnlineArchivePercentage, true
}

// HasOnlineArchivePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasOnlineArchivePercentage() bool {
	return o != nil && o.OnlineArchivePercentage != nil
}

// SetOnlineArchivePercentage gets a reference to the given float64 and assigns it to the OnlineArchivePercentage field.
func (o *MonthlyUsageAttributionValues) SetOnlineArchivePercentage(v float64) {
	o.OnlineArchivePercentage = &v
}

// GetOnlineArchiveUsage returns the OnlineArchiveUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetOnlineArchiveUsage() float64 {
	if o == nil || o.OnlineArchiveUsage == nil {
		var ret float64
		return ret
	}
	return *o.OnlineArchiveUsage
}

// GetOnlineArchiveUsageOk returns a tuple with the OnlineArchiveUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetOnlineArchiveUsageOk() (*float64, bool) {
	if o == nil || o.OnlineArchiveUsage == nil {
		return nil, false
	}
	return o.OnlineArchiveUsage, true
}

// HasOnlineArchiveUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasOnlineArchiveUsage() bool {
	return o != nil && o.OnlineArchiveUsage != nil
}

// SetOnlineArchiveUsage gets a reference to the given float64 and assigns it to the OnlineArchiveUsage field.
func (o *MonthlyUsageAttributionValues) SetOnlineArchiveUsage(v float64) {
	o.OnlineArchiveUsage = &v
}

// GetProductAnalyticsSessionPercentage returns the ProductAnalyticsSessionPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProductAnalyticsSessionPercentage() float64 {
	if o == nil || o.ProductAnalyticsSessionPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ProductAnalyticsSessionPercentage
}

// GetProductAnalyticsSessionPercentageOk returns a tuple with the ProductAnalyticsSessionPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProductAnalyticsSessionPercentageOk() (*float64, bool) {
	if o == nil || o.ProductAnalyticsSessionPercentage == nil {
		return nil, false
	}
	return o.ProductAnalyticsSessionPercentage, true
}

// HasProductAnalyticsSessionPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProductAnalyticsSessionPercentage() bool {
	return o != nil && o.ProductAnalyticsSessionPercentage != nil
}

// SetProductAnalyticsSessionPercentage gets a reference to the given float64 and assigns it to the ProductAnalyticsSessionPercentage field.
func (o *MonthlyUsageAttributionValues) SetProductAnalyticsSessionPercentage(v float64) {
	o.ProductAnalyticsSessionPercentage = &v
}

// GetProductAnalyticsSessionUsage returns the ProductAnalyticsSessionUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProductAnalyticsSessionUsage() float64 {
	if o == nil || o.ProductAnalyticsSessionUsage == nil {
		var ret float64
		return ret
	}
	return *o.ProductAnalyticsSessionUsage
}

// GetProductAnalyticsSessionUsageOk returns a tuple with the ProductAnalyticsSessionUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProductAnalyticsSessionUsageOk() (*float64, bool) {
	if o == nil || o.ProductAnalyticsSessionUsage == nil {
		return nil, false
	}
	return o.ProductAnalyticsSessionUsage, true
}

// HasProductAnalyticsSessionUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProductAnalyticsSessionUsage() bool {
	return o != nil && o.ProductAnalyticsSessionUsage != nil
}

// SetProductAnalyticsSessionUsage gets a reference to the given float64 and assigns it to the ProductAnalyticsSessionUsage field.
func (o *MonthlyUsageAttributionValues) SetProductAnalyticsSessionUsage(v float64) {
	o.ProductAnalyticsSessionUsage = &v
}

// GetProfiledContainerPercentage returns the ProfiledContainerPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerPercentage() float64 {
	if o == nil || o.ProfiledContainerPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledContainerPercentage
}

// GetProfiledContainerPercentageOk returns a tuple with the ProfiledContainerPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerPercentageOk() (*float64, bool) {
	if o == nil || o.ProfiledContainerPercentage == nil {
		return nil, false
	}
	return o.ProfiledContainerPercentage, true
}

// HasProfiledContainerPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledContainerPercentage() bool {
	return o != nil && o.ProfiledContainerPercentage != nil
}

// SetProfiledContainerPercentage gets a reference to the given float64 and assigns it to the ProfiledContainerPercentage field.
func (o *MonthlyUsageAttributionValues) SetProfiledContainerPercentage(v float64) {
	o.ProfiledContainerPercentage = &v
}

// GetProfiledContainerUsage returns the ProfiledContainerUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerUsage() float64 {
	if o == nil || o.ProfiledContainerUsage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledContainerUsage
}

// GetProfiledContainerUsageOk returns a tuple with the ProfiledContainerUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledContainerUsageOk() (*float64, bool) {
	if o == nil || o.ProfiledContainerUsage == nil {
		return nil, false
	}
	return o.ProfiledContainerUsage, true
}

// HasProfiledContainerUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledContainerUsage() bool {
	return o != nil && o.ProfiledContainerUsage != nil
}

// SetProfiledContainerUsage gets a reference to the given float64 and assigns it to the ProfiledContainerUsage field.
func (o *MonthlyUsageAttributionValues) SetProfiledContainerUsage(v float64) {
	o.ProfiledContainerUsage = &v
}

// GetProfiledFargatePercentage returns the ProfiledFargatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledFargatePercentage() float64 {
	if o == nil || o.ProfiledFargatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledFargatePercentage
}

// GetProfiledFargatePercentageOk returns a tuple with the ProfiledFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledFargatePercentageOk() (*float64, bool) {
	if o == nil || o.ProfiledFargatePercentage == nil {
		return nil, false
	}
	return o.ProfiledFargatePercentage, true
}

// HasProfiledFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledFargatePercentage() bool {
	return o != nil && o.ProfiledFargatePercentage != nil
}

// SetProfiledFargatePercentage gets a reference to the given float64 and assigns it to the ProfiledFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetProfiledFargatePercentage(v float64) {
	o.ProfiledFargatePercentage = &v
}

// GetProfiledFargateUsage returns the ProfiledFargateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledFargateUsage() float64 {
	if o == nil || o.ProfiledFargateUsage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledFargateUsage
}

// GetProfiledFargateUsageOk returns a tuple with the ProfiledFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledFargateUsageOk() (*float64, bool) {
	if o == nil || o.ProfiledFargateUsage == nil {
		return nil, false
	}
	return o.ProfiledFargateUsage, true
}

// HasProfiledFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledFargateUsage() bool {
	return o != nil && o.ProfiledFargateUsage != nil
}

// SetProfiledFargateUsage gets a reference to the given float64 and assigns it to the ProfiledFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetProfiledFargateUsage(v float64) {
	o.ProfiledFargateUsage = &v
}

// GetProfiledHostPercentage returns the ProfiledHostPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledHostPercentage() float64 {
	if o == nil || o.ProfiledHostPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledHostPercentage
}

// GetProfiledHostPercentageOk returns a tuple with the ProfiledHostPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledHostPercentageOk() (*float64, bool) {
	if o == nil || o.ProfiledHostPercentage == nil {
		return nil, false
	}
	return o.ProfiledHostPercentage, true
}

// HasProfiledHostPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledHostPercentage() bool {
	return o != nil && o.ProfiledHostPercentage != nil
}

// SetProfiledHostPercentage gets a reference to the given float64 and assigns it to the ProfiledHostPercentage field.
func (o *MonthlyUsageAttributionValues) SetProfiledHostPercentage(v float64) {
	o.ProfiledHostPercentage = &v
}

// GetProfiledHostUsage returns the ProfiledHostUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetProfiledHostUsage() float64 {
	if o == nil || o.ProfiledHostUsage == nil {
		var ret float64
		return ret
	}
	return *o.ProfiledHostUsage
}

// GetProfiledHostUsageOk returns a tuple with the ProfiledHostUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetProfiledHostUsageOk() (*float64, bool) {
	if o == nil || o.ProfiledHostUsage == nil {
		return nil, false
	}
	return o.ProfiledHostUsage, true
}

// HasProfiledHostUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasProfiledHostUsage() bool {
	return o != nil && o.ProfiledHostUsage != nil
}

// SetProfiledHostUsage gets a reference to the given float64 and assigns it to the ProfiledHostUsage field.
func (o *MonthlyUsageAttributionValues) SetProfiledHostUsage(v float64) {
	o.ProfiledHostUsage = &v
}

// GetPublishedAppPercentage returns the PublishedAppPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetPublishedAppPercentage() float64 {
	if o == nil || o.PublishedAppPercentage == nil {
		var ret float64
		return ret
	}
	return *o.PublishedAppPercentage
}

// GetPublishedAppPercentageOk returns a tuple with the PublishedAppPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetPublishedAppPercentageOk() (*float64, bool) {
	if o == nil || o.PublishedAppPercentage == nil {
		return nil, false
	}
	return o.PublishedAppPercentage, true
}

// HasPublishedAppPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasPublishedAppPercentage() bool {
	return o != nil && o.PublishedAppPercentage != nil
}

// SetPublishedAppPercentage gets a reference to the given float64 and assigns it to the PublishedAppPercentage field.
func (o *MonthlyUsageAttributionValues) SetPublishedAppPercentage(v float64) {
	o.PublishedAppPercentage = &v
}

// GetPublishedAppUsage returns the PublishedAppUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetPublishedAppUsage() float64 {
	if o == nil || o.PublishedAppUsage == nil {
		var ret float64
		return ret
	}
	return *o.PublishedAppUsage
}

// GetPublishedAppUsageOk returns a tuple with the PublishedAppUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetPublishedAppUsageOk() (*float64, bool) {
	if o == nil || o.PublishedAppUsage == nil {
		return nil, false
	}
	return o.PublishedAppUsage, true
}

// HasPublishedAppUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasPublishedAppUsage() bool {
	return o != nil && o.PublishedAppUsage != nil
}

// SetPublishedAppUsage gets a reference to the given float64 and assigns it to the PublishedAppUsage field.
func (o *MonthlyUsageAttributionValues) SetPublishedAppUsage(v float64) {
	o.PublishedAppUsage = &v
}

// GetRumBrowserMobileSessionsPercentage returns the RumBrowserMobileSessionsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumBrowserMobileSessionsPercentage() float64 {
	if o == nil || o.RumBrowserMobileSessionsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.RumBrowserMobileSessionsPercentage
}

// GetRumBrowserMobileSessionsPercentageOk returns a tuple with the RumBrowserMobileSessionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumBrowserMobileSessionsPercentageOk() (*float64, bool) {
	if o == nil || o.RumBrowserMobileSessionsPercentage == nil {
		return nil, false
	}
	return o.RumBrowserMobileSessionsPercentage, true
}

// HasRumBrowserMobileSessionsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumBrowserMobileSessionsPercentage() bool {
	return o != nil && o.RumBrowserMobileSessionsPercentage != nil
}

// SetRumBrowserMobileSessionsPercentage gets a reference to the given float64 and assigns it to the RumBrowserMobileSessionsPercentage field.
func (o *MonthlyUsageAttributionValues) SetRumBrowserMobileSessionsPercentage(v float64) {
	o.RumBrowserMobileSessionsPercentage = &v
}

// GetRumBrowserMobileSessionsUsage returns the RumBrowserMobileSessionsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumBrowserMobileSessionsUsage() float64 {
	if o == nil || o.RumBrowserMobileSessionsUsage == nil {
		var ret float64
		return ret
	}
	return *o.RumBrowserMobileSessionsUsage
}

// GetRumBrowserMobileSessionsUsageOk returns a tuple with the RumBrowserMobileSessionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumBrowserMobileSessionsUsageOk() (*float64, bool) {
	if o == nil || o.RumBrowserMobileSessionsUsage == nil {
		return nil, false
	}
	return o.RumBrowserMobileSessionsUsage, true
}

// HasRumBrowserMobileSessionsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumBrowserMobileSessionsUsage() bool {
	return o != nil && o.RumBrowserMobileSessionsUsage != nil
}

// SetRumBrowserMobileSessionsUsage gets a reference to the given float64 and assigns it to the RumBrowserMobileSessionsUsage field.
func (o *MonthlyUsageAttributionValues) SetRumBrowserMobileSessionsUsage(v float64) {
	o.RumBrowserMobileSessionsUsage = &v
}

// GetRumIngestedPercentage returns the RumIngestedPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumIngestedPercentage() float64 {
	if o == nil || o.RumIngestedPercentage == nil {
		var ret float64
		return ret
	}
	return *o.RumIngestedPercentage
}

// GetRumIngestedPercentageOk returns a tuple with the RumIngestedPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumIngestedPercentageOk() (*float64, bool) {
	if o == nil || o.RumIngestedPercentage == nil {
		return nil, false
	}
	return o.RumIngestedPercentage, true
}

// HasRumIngestedPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumIngestedPercentage() bool {
	return o != nil && o.RumIngestedPercentage != nil
}

// SetRumIngestedPercentage gets a reference to the given float64 and assigns it to the RumIngestedPercentage field.
func (o *MonthlyUsageAttributionValues) SetRumIngestedPercentage(v float64) {
	o.RumIngestedPercentage = &v
}

// GetRumIngestedUsage returns the RumIngestedUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumIngestedUsage() float64 {
	if o == nil || o.RumIngestedUsage == nil {
		var ret float64
		return ret
	}
	return *o.RumIngestedUsage
}

// GetRumIngestedUsageOk returns a tuple with the RumIngestedUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumIngestedUsageOk() (*float64, bool) {
	if o == nil || o.RumIngestedUsage == nil {
		return nil, false
	}
	return o.RumIngestedUsage, true
}

// HasRumIngestedUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumIngestedUsage() bool {
	return o != nil && o.RumIngestedUsage != nil
}

// SetRumIngestedUsage gets a reference to the given float64 and assigns it to the RumIngestedUsage field.
func (o *MonthlyUsageAttributionValues) SetRumIngestedUsage(v float64) {
	o.RumIngestedUsage = &v
}

// GetRumInvestigatePercentage returns the RumInvestigatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumInvestigatePercentage() float64 {
	if o == nil || o.RumInvestigatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.RumInvestigatePercentage
}

// GetRumInvestigatePercentageOk returns a tuple with the RumInvestigatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumInvestigatePercentageOk() (*float64, bool) {
	if o == nil || o.RumInvestigatePercentage == nil {
		return nil, false
	}
	return o.RumInvestigatePercentage, true
}

// HasRumInvestigatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumInvestigatePercentage() bool {
	return o != nil && o.RumInvestigatePercentage != nil
}

// SetRumInvestigatePercentage gets a reference to the given float64 and assigns it to the RumInvestigatePercentage field.
func (o *MonthlyUsageAttributionValues) SetRumInvestigatePercentage(v float64) {
	o.RumInvestigatePercentage = &v
}

// GetRumInvestigateUsage returns the RumInvestigateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumInvestigateUsage() float64 {
	if o == nil || o.RumInvestigateUsage == nil {
		var ret float64
		return ret
	}
	return *o.RumInvestigateUsage
}

// GetRumInvestigateUsageOk returns a tuple with the RumInvestigateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumInvestigateUsageOk() (*float64, bool) {
	if o == nil || o.RumInvestigateUsage == nil {
		return nil, false
	}
	return o.RumInvestigateUsage, true
}

// HasRumInvestigateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumInvestigateUsage() bool {
	return o != nil && o.RumInvestigateUsage != nil
}

// SetRumInvestigateUsage gets a reference to the given float64 and assigns it to the RumInvestigateUsage field.
func (o *MonthlyUsageAttributionValues) SetRumInvestigateUsage(v float64) {
	o.RumInvestigateUsage = &v
}

// GetRumReplaySessionsPercentage returns the RumReplaySessionsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumReplaySessionsPercentage() float64 {
	if o == nil || o.RumReplaySessionsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.RumReplaySessionsPercentage
}

// GetRumReplaySessionsPercentageOk returns a tuple with the RumReplaySessionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumReplaySessionsPercentageOk() (*float64, bool) {
	if o == nil || o.RumReplaySessionsPercentage == nil {
		return nil, false
	}
	return o.RumReplaySessionsPercentage, true
}

// HasRumReplaySessionsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumReplaySessionsPercentage() bool {
	return o != nil && o.RumReplaySessionsPercentage != nil
}

// SetRumReplaySessionsPercentage gets a reference to the given float64 and assigns it to the RumReplaySessionsPercentage field.
func (o *MonthlyUsageAttributionValues) SetRumReplaySessionsPercentage(v float64) {
	o.RumReplaySessionsPercentage = &v
}

// GetRumReplaySessionsUsage returns the RumReplaySessionsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumReplaySessionsUsage() float64 {
	if o == nil || o.RumReplaySessionsUsage == nil {
		var ret float64
		return ret
	}
	return *o.RumReplaySessionsUsage
}

// GetRumReplaySessionsUsageOk returns a tuple with the RumReplaySessionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumReplaySessionsUsageOk() (*float64, bool) {
	if o == nil || o.RumReplaySessionsUsage == nil {
		return nil, false
	}
	return o.RumReplaySessionsUsage, true
}

// HasRumReplaySessionsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumReplaySessionsUsage() bool {
	return o != nil && o.RumReplaySessionsUsage != nil
}

// SetRumReplaySessionsUsage gets a reference to the given float64 and assigns it to the RumReplaySessionsUsage field.
func (o *MonthlyUsageAttributionValues) SetRumReplaySessionsUsage(v float64) {
	o.RumReplaySessionsUsage = &v
}

// GetRumSessionReplayAddOnPercentage returns the RumSessionReplayAddOnPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumSessionReplayAddOnPercentage() float64 {
	if o == nil || o.RumSessionReplayAddOnPercentage == nil {
		var ret float64
		return ret
	}
	return *o.RumSessionReplayAddOnPercentage
}

// GetRumSessionReplayAddOnPercentageOk returns a tuple with the RumSessionReplayAddOnPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumSessionReplayAddOnPercentageOk() (*float64, bool) {
	if o == nil || o.RumSessionReplayAddOnPercentage == nil {
		return nil, false
	}
	return o.RumSessionReplayAddOnPercentage, true
}

// HasRumSessionReplayAddOnPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumSessionReplayAddOnPercentage() bool {
	return o != nil && o.RumSessionReplayAddOnPercentage != nil
}

// SetRumSessionReplayAddOnPercentage gets a reference to the given float64 and assigns it to the RumSessionReplayAddOnPercentage field.
func (o *MonthlyUsageAttributionValues) SetRumSessionReplayAddOnPercentage(v float64) {
	o.RumSessionReplayAddOnPercentage = &v
}

// GetRumSessionReplayAddOnUsage returns the RumSessionReplayAddOnUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetRumSessionReplayAddOnUsage() float64 {
	if o == nil || o.RumSessionReplayAddOnUsage == nil {
		var ret float64
		return ret
	}
	return *o.RumSessionReplayAddOnUsage
}

// GetRumSessionReplayAddOnUsageOk returns a tuple with the RumSessionReplayAddOnUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetRumSessionReplayAddOnUsageOk() (*float64, bool) {
	if o == nil || o.RumSessionReplayAddOnUsage == nil {
		return nil, false
	}
	return o.RumSessionReplayAddOnUsage, true
}

// HasRumSessionReplayAddOnUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasRumSessionReplayAddOnUsage() bool {
	return o != nil && o.RumSessionReplayAddOnUsage != nil
}

// SetRumSessionReplayAddOnUsage gets a reference to the given float64 and assigns it to the RumSessionReplayAddOnUsage field.
func (o *MonthlyUsageAttributionValues) SetRumSessionReplayAddOnUsage(v float64) {
	o.RumSessionReplayAddOnUsage = &v
}

// GetScaFargatePercentage returns the ScaFargatePercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetScaFargatePercentage() float64 {
	if o == nil || o.ScaFargatePercentage == nil {
		var ret float64
		return ret
	}
	return *o.ScaFargatePercentage
}

// GetScaFargatePercentageOk returns a tuple with the ScaFargatePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetScaFargatePercentageOk() (*float64, bool) {
	if o == nil || o.ScaFargatePercentage == nil {
		return nil, false
	}
	return o.ScaFargatePercentage, true
}

// HasScaFargatePercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasScaFargatePercentage() bool {
	return o != nil && o.ScaFargatePercentage != nil
}

// SetScaFargatePercentage gets a reference to the given float64 and assigns it to the ScaFargatePercentage field.
func (o *MonthlyUsageAttributionValues) SetScaFargatePercentage(v float64) {
	o.ScaFargatePercentage = &v
}

// GetScaFargateUsage returns the ScaFargateUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetScaFargateUsage() float64 {
	if o == nil || o.ScaFargateUsage == nil {
		var ret float64
		return ret
	}
	return *o.ScaFargateUsage
}

// GetScaFargateUsageOk returns a tuple with the ScaFargateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetScaFargateUsageOk() (*float64, bool) {
	if o == nil || o.ScaFargateUsage == nil {
		return nil, false
	}
	return o.ScaFargateUsage, true
}

// HasScaFargateUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasScaFargateUsage() bool {
	return o != nil && o.ScaFargateUsage != nil
}

// SetScaFargateUsage gets a reference to the given float64 and assigns it to the ScaFargateUsage field.
func (o *MonthlyUsageAttributionValues) SetScaFargateUsage(v float64) {
	o.ScaFargateUsage = &v
}

// GetSdsScannedBytesPercentage returns the SdsScannedBytesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSdsScannedBytesPercentage() float64 {
	if o == nil || o.SdsScannedBytesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.SdsScannedBytesPercentage
}

// GetSdsScannedBytesPercentageOk returns a tuple with the SdsScannedBytesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSdsScannedBytesPercentageOk() (*float64, bool) {
	if o == nil || o.SdsScannedBytesPercentage == nil {
		return nil, false
	}
	return o.SdsScannedBytesPercentage, true
}

// HasSdsScannedBytesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSdsScannedBytesPercentage() bool {
	return o != nil && o.SdsScannedBytesPercentage != nil
}

// SetSdsScannedBytesPercentage gets a reference to the given float64 and assigns it to the SdsScannedBytesPercentage field.
func (o *MonthlyUsageAttributionValues) SetSdsScannedBytesPercentage(v float64) {
	o.SdsScannedBytesPercentage = &v
}

// GetSdsScannedBytesUsage returns the SdsScannedBytesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSdsScannedBytesUsage() float64 {
	if o == nil || o.SdsScannedBytesUsage == nil {
		var ret float64
		return ret
	}
	return *o.SdsScannedBytesUsage
}

// GetSdsScannedBytesUsageOk returns a tuple with the SdsScannedBytesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSdsScannedBytesUsageOk() (*float64, bool) {
	if o == nil || o.SdsScannedBytesUsage == nil {
		return nil, false
	}
	return o.SdsScannedBytesUsage, true
}

// HasSdsScannedBytesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSdsScannedBytesUsage() bool {
	return o != nil && o.SdsScannedBytesUsage != nil
}

// SetSdsScannedBytesUsage gets a reference to the given float64 and assigns it to the SdsScannedBytesUsage field.
func (o *MonthlyUsageAttributionValues) SetSdsScannedBytesUsage(v float64) {
	o.SdsScannedBytesUsage = &v
}

// GetServerlessAppsPercentage returns the ServerlessAppsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetServerlessAppsPercentage() float64 {
	if o == nil || o.ServerlessAppsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.ServerlessAppsPercentage
}

// GetServerlessAppsPercentageOk returns a tuple with the ServerlessAppsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetServerlessAppsPercentageOk() (*float64, bool) {
	if o == nil || o.ServerlessAppsPercentage == nil {
		return nil, false
	}
	return o.ServerlessAppsPercentage, true
}

// HasServerlessAppsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasServerlessAppsPercentage() bool {
	return o != nil && o.ServerlessAppsPercentage != nil
}

// SetServerlessAppsPercentage gets a reference to the given float64 and assigns it to the ServerlessAppsPercentage field.
func (o *MonthlyUsageAttributionValues) SetServerlessAppsPercentage(v float64) {
	o.ServerlessAppsPercentage = &v
}

// GetServerlessAppsUsage returns the ServerlessAppsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetServerlessAppsUsage() float64 {
	if o == nil || o.ServerlessAppsUsage == nil {
		var ret float64
		return ret
	}
	return *o.ServerlessAppsUsage
}

// GetServerlessAppsUsageOk returns a tuple with the ServerlessAppsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetServerlessAppsUsageOk() (*float64, bool) {
	if o == nil || o.ServerlessAppsUsage == nil {
		return nil, false
	}
	return o.ServerlessAppsUsage, true
}

// HasServerlessAppsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasServerlessAppsUsage() bool {
	return o != nil && o.ServerlessAppsUsage != nil
}

// SetServerlessAppsUsage gets a reference to the given float64 and assigns it to the ServerlessAppsUsage field.
func (o *MonthlyUsageAttributionValues) SetServerlessAppsUsage(v float64) {
	o.ServerlessAppsUsage = &v
}

// GetSiemAnalyzedLogsAddOnPercentage returns the SiemAnalyzedLogsAddOnPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSiemAnalyzedLogsAddOnPercentage() float64 {
	if o == nil || o.SiemAnalyzedLogsAddOnPercentage == nil {
		var ret float64
		return ret
	}
	return *o.SiemAnalyzedLogsAddOnPercentage
}

// GetSiemAnalyzedLogsAddOnPercentageOk returns a tuple with the SiemAnalyzedLogsAddOnPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSiemAnalyzedLogsAddOnPercentageOk() (*float64, bool) {
	if o == nil || o.SiemAnalyzedLogsAddOnPercentage == nil {
		return nil, false
	}
	return o.SiemAnalyzedLogsAddOnPercentage, true
}

// HasSiemAnalyzedLogsAddOnPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSiemAnalyzedLogsAddOnPercentage() bool {
	return o != nil && o.SiemAnalyzedLogsAddOnPercentage != nil
}

// SetSiemAnalyzedLogsAddOnPercentage gets a reference to the given float64 and assigns it to the SiemAnalyzedLogsAddOnPercentage field.
func (o *MonthlyUsageAttributionValues) SetSiemAnalyzedLogsAddOnPercentage(v float64) {
	o.SiemAnalyzedLogsAddOnPercentage = &v
}

// GetSiemAnalyzedLogsAddOnUsage returns the SiemAnalyzedLogsAddOnUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSiemAnalyzedLogsAddOnUsage() float64 {
	if o == nil || o.SiemAnalyzedLogsAddOnUsage == nil {
		var ret float64
		return ret
	}
	return *o.SiemAnalyzedLogsAddOnUsage
}

// GetSiemAnalyzedLogsAddOnUsageOk returns a tuple with the SiemAnalyzedLogsAddOnUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSiemAnalyzedLogsAddOnUsageOk() (*float64, bool) {
	if o == nil || o.SiemAnalyzedLogsAddOnUsage == nil {
		return nil, false
	}
	return o.SiemAnalyzedLogsAddOnUsage, true
}

// HasSiemAnalyzedLogsAddOnUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSiemAnalyzedLogsAddOnUsage() bool {
	return o != nil && o.SiemAnalyzedLogsAddOnUsage != nil
}

// SetSiemAnalyzedLogsAddOnUsage gets a reference to the given float64 and assigns it to the SiemAnalyzedLogsAddOnUsage field.
func (o *MonthlyUsageAttributionValues) SetSiemAnalyzedLogsAddOnUsage(v float64) {
	o.SiemAnalyzedLogsAddOnUsage = &v
}

// GetSiemIngestedBytesPercentage returns the SiemIngestedBytesPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSiemIngestedBytesPercentage() float64 {
	if o == nil || o.SiemIngestedBytesPercentage == nil {
		var ret float64
		return ret
	}
	return *o.SiemIngestedBytesPercentage
}

// GetSiemIngestedBytesPercentageOk returns a tuple with the SiemIngestedBytesPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSiemIngestedBytesPercentageOk() (*float64, bool) {
	if o == nil || o.SiemIngestedBytesPercentage == nil {
		return nil, false
	}
	return o.SiemIngestedBytesPercentage, true
}

// HasSiemIngestedBytesPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSiemIngestedBytesPercentage() bool {
	return o != nil && o.SiemIngestedBytesPercentage != nil
}

// SetSiemIngestedBytesPercentage gets a reference to the given float64 and assigns it to the SiemIngestedBytesPercentage field.
func (o *MonthlyUsageAttributionValues) SetSiemIngestedBytesPercentage(v float64) {
	o.SiemIngestedBytesPercentage = &v
}

// GetSiemIngestedBytesUsage returns the SiemIngestedBytesUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSiemIngestedBytesUsage() float64 {
	if o == nil || o.SiemIngestedBytesUsage == nil {
		var ret float64
		return ret
	}
	return *o.SiemIngestedBytesUsage
}

// GetSiemIngestedBytesUsageOk returns a tuple with the SiemIngestedBytesUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSiemIngestedBytesUsageOk() (*float64, bool) {
	if o == nil || o.SiemIngestedBytesUsage == nil {
		return nil, false
	}
	return o.SiemIngestedBytesUsage, true
}

// HasSiemIngestedBytesUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSiemIngestedBytesUsage() bool {
	return o != nil && o.SiemIngestedBytesUsage != nil
}

// SetSiemIngestedBytesUsage gets a reference to the given float64 and assigns it to the SiemIngestedBytesUsage field.
func (o *MonthlyUsageAttributionValues) SetSiemIngestedBytesUsage(v float64) {
	o.SiemIngestedBytesUsage = &v
}

// GetSnmpPercentage returns the SnmpPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSnmpPercentage() float64 {
	if o == nil || o.SnmpPercentage == nil {
		var ret float64
		return ret
	}
	return *o.SnmpPercentage
}

// GetSnmpPercentageOk returns a tuple with the SnmpPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSnmpPercentageOk() (*float64, bool) {
	if o == nil || o.SnmpPercentage == nil {
		return nil, false
	}
	return o.SnmpPercentage, true
}

// HasSnmpPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSnmpPercentage() bool {
	return o != nil && o.SnmpPercentage != nil
}

// SetSnmpPercentage gets a reference to the given float64 and assigns it to the SnmpPercentage field.
func (o *MonthlyUsageAttributionValues) SetSnmpPercentage(v float64) {
	o.SnmpPercentage = &v
}

// GetSnmpUsage returns the SnmpUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetSnmpUsage() float64 {
	if o == nil || o.SnmpUsage == nil {
		var ret float64
		return ret
	}
	return *o.SnmpUsage
}

// GetSnmpUsageOk returns a tuple with the SnmpUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetSnmpUsageOk() (*float64, bool) {
	if o == nil || o.SnmpUsage == nil {
		return nil, false
	}
	return o.SnmpUsage, true
}

// HasSnmpUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasSnmpUsage() bool {
	return o != nil && o.SnmpUsage != nil
}

// SetSnmpUsage gets a reference to the given float64 and assigns it to the SnmpUsage field.
func (o *MonthlyUsageAttributionValues) SetSnmpUsage(v float64) {
	o.SnmpUsage = &v
}

// GetUniversalServiceMonitoringPercentage returns the UniversalServiceMonitoringPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringPercentage() float64 {
	if o == nil || o.UniversalServiceMonitoringPercentage == nil {
		var ret float64
		return ret
	}
	return *o.UniversalServiceMonitoringPercentage
}

// GetUniversalServiceMonitoringPercentageOk returns a tuple with the UniversalServiceMonitoringPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringPercentageOk() (*float64, bool) {
	if o == nil || o.UniversalServiceMonitoringPercentage == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringPercentage, true
}

// HasUniversalServiceMonitoringPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasUniversalServiceMonitoringPercentage() bool {
	return o != nil && o.UniversalServiceMonitoringPercentage != nil
}

// SetUniversalServiceMonitoringPercentage gets a reference to the given float64 and assigns it to the UniversalServiceMonitoringPercentage field.
func (o *MonthlyUsageAttributionValues) SetUniversalServiceMonitoringPercentage(v float64) {
	o.UniversalServiceMonitoringPercentage = &v
}

// GetUniversalServiceMonitoringUsage returns the UniversalServiceMonitoringUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringUsage() float64 {
	if o == nil || o.UniversalServiceMonitoringUsage == nil {
		var ret float64
		return ret
	}
	return *o.UniversalServiceMonitoringUsage
}

// GetUniversalServiceMonitoringUsageOk returns a tuple with the UniversalServiceMonitoringUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetUniversalServiceMonitoringUsageOk() (*float64, bool) {
	if o == nil || o.UniversalServiceMonitoringUsage == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringUsage, true
}

// HasUniversalServiceMonitoringUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasUniversalServiceMonitoringUsage() bool {
	return o != nil && o.UniversalServiceMonitoringUsage != nil
}

// SetUniversalServiceMonitoringUsage gets a reference to the given float64 and assigns it to the UniversalServiceMonitoringUsage field.
func (o *MonthlyUsageAttributionValues) SetUniversalServiceMonitoringUsage(v float64) {
	o.UniversalServiceMonitoringUsage = &v
}

// GetVulnManagementHostsPercentage returns the VulnManagementHostsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetVulnManagementHostsPercentage() float64 {
	if o == nil || o.VulnManagementHostsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.VulnManagementHostsPercentage
}

// GetVulnManagementHostsPercentageOk returns a tuple with the VulnManagementHostsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetVulnManagementHostsPercentageOk() (*float64, bool) {
	if o == nil || o.VulnManagementHostsPercentage == nil {
		return nil, false
	}
	return o.VulnManagementHostsPercentage, true
}

// HasVulnManagementHostsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasVulnManagementHostsPercentage() bool {
	return o != nil && o.VulnManagementHostsPercentage != nil
}

// SetVulnManagementHostsPercentage gets a reference to the given float64 and assigns it to the VulnManagementHostsPercentage field.
func (o *MonthlyUsageAttributionValues) SetVulnManagementHostsPercentage(v float64) {
	o.VulnManagementHostsPercentage = &v
}

// GetVulnManagementHostsUsage returns the VulnManagementHostsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetVulnManagementHostsUsage() float64 {
	if o == nil || o.VulnManagementHostsUsage == nil {
		var ret float64
		return ret
	}
	return *o.VulnManagementHostsUsage
}

// GetVulnManagementHostsUsageOk returns a tuple with the VulnManagementHostsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetVulnManagementHostsUsageOk() (*float64, bool) {
	if o == nil || o.VulnManagementHostsUsage == nil {
		return nil, false
	}
	return o.VulnManagementHostsUsage, true
}

// HasVulnManagementHostsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasVulnManagementHostsUsage() bool {
	return o != nil && o.VulnManagementHostsUsage != nil
}

// SetVulnManagementHostsUsage gets a reference to the given float64 and assigns it to the VulnManagementHostsUsage field.
func (o *MonthlyUsageAttributionValues) SetVulnManagementHostsUsage(v float64) {
	o.VulnManagementHostsUsage = &v
}

// GetWorkflowExecutionsPercentage returns the WorkflowExecutionsPercentage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetWorkflowExecutionsPercentage() float64 {
	if o == nil || o.WorkflowExecutionsPercentage == nil {
		var ret float64
		return ret
	}
	return *o.WorkflowExecutionsPercentage
}

// GetWorkflowExecutionsPercentageOk returns a tuple with the WorkflowExecutionsPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetWorkflowExecutionsPercentageOk() (*float64, bool) {
	if o == nil || o.WorkflowExecutionsPercentage == nil {
		return nil, false
	}
	return o.WorkflowExecutionsPercentage, true
}

// HasWorkflowExecutionsPercentage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasWorkflowExecutionsPercentage() bool {
	return o != nil && o.WorkflowExecutionsPercentage != nil
}

// SetWorkflowExecutionsPercentage gets a reference to the given float64 and assigns it to the WorkflowExecutionsPercentage field.
func (o *MonthlyUsageAttributionValues) SetWorkflowExecutionsPercentage(v float64) {
	o.WorkflowExecutionsPercentage = &v
}

// GetWorkflowExecutionsUsage returns the WorkflowExecutionsUsage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionValues) GetWorkflowExecutionsUsage() float64 {
	if o == nil || o.WorkflowExecutionsUsage == nil {
		var ret float64
		return ret
	}
	return *o.WorkflowExecutionsUsage
}

// GetWorkflowExecutionsUsageOk returns a tuple with the WorkflowExecutionsUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionValues) GetWorkflowExecutionsUsageOk() (*float64, bool) {
	if o == nil || o.WorkflowExecutionsUsage == nil {
		return nil, false
	}
	return o.WorkflowExecutionsUsage, true
}

// HasWorkflowExecutionsUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionValues) HasWorkflowExecutionsUsage() bool {
	return o != nil && o.WorkflowExecutionsUsage != nil
}

// SetWorkflowExecutionsUsage gets a reference to the given float64 and assigns it to the WorkflowExecutionsUsage field.
func (o *MonthlyUsageAttributionValues) SetWorkflowExecutionsUsage(v float64) {
	o.WorkflowExecutionsUsage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyUsageAttributionValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiPercentage != nil {
		toSerialize["api_percentage"] = o.ApiPercentage
	}
	if o.ApiUsage != nil {
		toSerialize["api_usage"] = o.ApiUsage
	}
	if o.ApmFargatePercentage != nil {
		toSerialize["apm_fargate_percentage"] = o.ApmFargatePercentage
	}
	if o.ApmFargateUsage != nil {
		toSerialize["apm_fargate_usage"] = o.ApmFargateUsage
	}
	if o.ApmHostPercentage != nil {
		toSerialize["apm_host_percentage"] = o.ApmHostPercentage
	}
	if o.ApmHostUsage != nil {
		toSerialize["apm_host_usage"] = o.ApmHostUsage
	}
	if o.ApmUsmPercentage != nil {
		toSerialize["apm_usm_percentage"] = o.ApmUsmPercentage
	}
	if o.ApmUsmUsage != nil {
		toSerialize["apm_usm_usage"] = o.ApmUsmUsage
	}
	if o.AppsecFargatePercentage != nil {
		toSerialize["appsec_fargate_percentage"] = o.AppsecFargatePercentage
	}
	if o.AppsecFargateUsage != nil {
		toSerialize["appsec_fargate_usage"] = o.AppsecFargateUsage
	}
	if o.AppsecPercentage != nil {
		toSerialize["appsec_percentage"] = o.AppsecPercentage
	}
	if o.AppsecUsage != nil {
		toSerialize["appsec_usage"] = o.AppsecUsage
	}
	if o.AsmServerlessTracedInvocationsPercentage != nil {
		toSerialize["asm_serverless_traced_invocations_percentage"] = o.AsmServerlessTracedInvocationsPercentage
	}
	if o.AsmServerlessTracedInvocationsUsage != nil {
		toSerialize["asm_serverless_traced_invocations_usage"] = o.AsmServerlessTracedInvocationsUsage
	}
	if o.BrowserPercentage != nil {
		toSerialize["browser_percentage"] = o.BrowserPercentage
	}
	if o.BrowserUsage != nil {
		toSerialize["browser_usage"] = o.BrowserUsage
	}
	if o.CiPipelineIndexedSpansPercentage != nil {
		toSerialize["ci_pipeline_indexed_spans_percentage"] = o.CiPipelineIndexedSpansPercentage
	}
	if o.CiPipelineIndexedSpansUsage != nil {
		toSerialize["ci_pipeline_indexed_spans_usage"] = o.CiPipelineIndexedSpansUsage
	}
	if o.CiTestIndexedSpansPercentage != nil {
		toSerialize["ci_test_indexed_spans_percentage"] = o.CiTestIndexedSpansPercentage
	}
	if o.CiTestIndexedSpansUsage != nil {
		toSerialize["ci_test_indexed_spans_usage"] = o.CiTestIndexedSpansUsage
	}
	if o.CiVisibilityItrPercentage != nil {
		toSerialize["ci_visibility_itr_percentage"] = o.CiVisibilityItrPercentage
	}
	if o.CiVisibilityItrUsage != nil {
		toSerialize["ci_visibility_itr_usage"] = o.CiVisibilityItrUsage
	}
	if o.CloudSiemPercentage != nil {
		toSerialize["cloud_siem_percentage"] = o.CloudSiemPercentage
	}
	if o.CloudSiemUsage != nil {
		toSerialize["cloud_siem_usage"] = o.CloudSiemUsage
	}
	if o.CodeSecurityHostPercentage != nil {
		toSerialize["code_security_host_percentage"] = o.CodeSecurityHostPercentage
	}
	if o.CodeSecurityHostUsage != nil {
		toSerialize["code_security_host_usage"] = o.CodeSecurityHostUsage
	}
	if o.ContainerExclAgentPercentage != nil {
		toSerialize["container_excl_agent_percentage"] = o.ContainerExclAgentPercentage
	}
	if o.ContainerExclAgentUsage != nil {
		toSerialize["container_excl_agent_usage"] = o.ContainerExclAgentUsage
	}
	if o.ContainerPercentage != nil {
		toSerialize["container_percentage"] = o.ContainerPercentage
	}
	if o.ContainerUsage != nil {
		toSerialize["container_usage"] = o.ContainerUsage
	}
	if o.CspmContainersPercentage != nil {
		toSerialize["cspm_containers_percentage"] = o.CspmContainersPercentage
	}
	if o.CspmContainersUsage != nil {
		toSerialize["cspm_containers_usage"] = o.CspmContainersUsage
	}
	if o.CspmHostsPercentage != nil {
		toSerialize["cspm_hosts_percentage"] = o.CspmHostsPercentage
	}
	if o.CspmHostsUsage != nil {
		toSerialize["cspm_hosts_usage"] = o.CspmHostsUsage
	}
	if o.CustomEventPercentage != nil {
		toSerialize["custom_event_percentage"] = o.CustomEventPercentage
	}
	if o.CustomEventUsage != nil {
		toSerialize["custom_event_usage"] = o.CustomEventUsage
	}
	if o.CustomIngestedTimeseriesPercentage != nil {
		toSerialize["custom_ingested_timeseries_percentage"] = o.CustomIngestedTimeseriesPercentage
	}
	if o.CustomIngestedTimeseriesUsage != nil {
		toSerialize["custom_ingested_timeseries_usage"] = o.CustomIngestedTimeseriesUsage
	}
	if o.CustomTimeseriesPercentage != nil {
		toSerialize["custom_timeseries_percentage"] = o.CustomTimeseriesPercentage
	}
	if o.CustomTimeseriesUsage != nil {
		toSerialize["custom_timeseries_usage"] = o.CustomTimeseriesUsage
	}
	if o.CwsContainersPercentage != nil {
		toSerialize["cws_containers_percentage"] = o.CwsContainersPercentage
	}
	if o.CwsContainersUsage != nil {
		toSerialize["cws_containers_usage"] = o.CwsContainersUsage
	}
	if o.CwsFargateTaskPercentage != nil {
		toSerialize["cws_fargate_task_percentage"] = o.CwsFargateTaskPercentage
	}
	if o.CwsFargateTaskUsage != nil {
		toSerialize["cws_fargate_task_usage"] = o.CwsFargateTaskUsage
	}
	if o.CwsHostsPercentage != nil {
		toSerialize["cws_hosts_percentage"] = o.CwsHostsPercentage
	}
	if o.CwsHostsUsage != nil {
		toSerialize["cws_hosts_usage"] = o.CwsHostsUsage
	}
	if o.DataJobsMonitoringUsage != nil {
		toSerialize["data_jobs_monitoring_usage"] = o.DataJobsMonitoringUsage
	}
	if o.DataStreamMonitoringUsage != nil {
		toSerialize["data_stream_monitoring_usage"] = o.DataStreamMonitoringUsage
	}
	if o.DbmHostsPercentage != nil {
		toSerialize["dbm_hosts_percentage"] = o.DbmHostsPercentage
	}
	if o.DbmHostsUsage != nil {
		toSerialize["dbm_hosts_usage"] = o.DbmHostsUsage
	}
	if o.DbmQueriesPercentage != nil {
		toSerialize["dbm_queries_percentage"] = o.DbmQueriesPercentage
	}
	if o.DbmQueriesUsage != nil {
		toSerialize["dbm_queries_usage"] = o.DbmQueriesUsage
	}
	if o.ErrorTrackingPercentage != nil {
		toSerialize["error_tracking_percentage"] = o.ErrorTrackingPercentage
	}
	if o.ErrorTrackingUsage != nil {
		toSerialize["error_tracking_usage"] = o.ErrorTrackingUsage
	}
	if o.EstimatedIndexedSpansPercentage != nil {
		toSerialize["estimated_indexed_spans_percentage"] = o.EstimatedIndexedSpansPercentage
	}
	if o.EstimatedIndexedSpansUsage != nil {
		toSerialize["estimated_indexed_spans_usage"] = o.EstimatedIndexedSpansUsage
	}
	if o.EstimatedIngestedSpansPercentage != nil {
		toSerialize["estimated_ingested_spans_percentage"] = o.EstimatedIngestedSpansPercentage
	}
	if o.EstimatedIngestedSpansUsage != nil {
		toSerialize["estimated_ingested_spans_usage"] = o.EstimatedIngestedSpansUsage
	}
	if o.FargatePercentage != nil {
		toSerialize["fargate_percentage"] = o.FargatePercentage
	}
	if o.FargateUsage != nil {
		toSerialize["fargate_usage"] = o.FargateUsage
	}
	if o.FunctionsPercentage != nil {
		toSerialize["functions_percentage"] = o.FunctionsPercentage
	}
	if o.FunctionsUsage != nil {
		toSerialize["functions_usage"] = o.FunctionsUsage
	}
	if o.IncidentManagementMonthlyActiveUsersPercentage != nil {
		toSerialize["incident_management_monthly_active_users_percentage"] = o.IncidentManagementMonthlyActiveUsersPercentage
	}
	if o.IncidentManagementMonthlyActiveUsersUsage != nil {
		toSerialize["incident_management_monthly_active_users_usage"] = o.IncidentManagementMonthlyActiveUsersUsage
	}
	if o.IndexedSpansPercentage != nil {
		toSerialize["indexed_spans_percentage"] = o.IndexedSpansPercentage
	}
	if o.IndexedSpansUsage != nil {
		toSerialize["indexed_spans_usage"] = o.IndexedSpansUsage
	}
	if o.InfraHostPercentage != nil {
		toSerialize["infra_host_percentage"] = o.InfraHostPercentage
	}
	if o.InfraHostUsage != nil {
		toSerialize["infra_host_usage"] = o.InfraHostUsage
	}
	if o.IngestedLogsBytesPercentage != nil {
		toSerialize["ingested_logs_bytes_percentage"] = o.IngestedLogsBytesPercentage
	}
	if o.IngestedLogsBytesUsage != nil {
		toSerialize["ingested_logs_bytes_usage"] = o.IngestedLogsBytesUsage
	}
	if o.IngestedSpansBytesPercentage != nil {
		toSerialize["ingested_spans_bytes_percentage"] = o.IngestedSpansBytesPercentage
	}
	if o.IngestedSpansBytesUsage != nil {
		toSerialize["ingested_spans_bytes_usage"] = o.IngestedSpansBytesUsage
	}
	if o.InvocationsPercentage != nil {
		toSerialize["invocations_percentage"] = o.InvocationsPercentage
	}
	if o.InvocationsUsage != nil {
		toSerialize["invocations_usage"] = o.InvocationsUsage
	}
	if o.LambdaTracedInvocationsPercentage != nil {
		toSerialize["lambda_traced_invocations_percentage"] = o.LambdaTracedInvocationsPercentage
	}
	if o.LambdaTracedInvocationsUsage != nil {
		toSerialize["lambda_traced_invocations_usage"] = o.LambdaTracedInvocationsUsage
	}
	if o.LlmObservabilityPercentage != nil {
		toSerialize["llm_observability_percentage"] = o.LlmObservabilityPercentage
	}
	if o.LlmObservabilityUsage != nil {
		toSerialize["llm_observability_usage"] = o.LlmObservabilityUsage
	}
	if o.LogsIndexed15dayPercentage != nil {
		toSerialize["logs_indexed_15day_percentage"] = o.LogsIndexed15dayPercentage
	}
	if o.LogsIndexed15dayUsage != nil {
		toSerialize["logs_indexed_15day_usage"] = o.LogsIndexed15dayUsage
	}
	if o.LogsIndexed180dayPercentage != nil {
		toSerialize["logs_indexed_180day_percentage"] = o.LogsIndexed180dayPercentage
	}
	if o.LogsIndexed180dayUsage != nil {
		toSerialize["logs_indexed_180day_usage"] = o.LogsIndexed180dayUsage
	}
	if o.LogsIndexed1dayPercentage != nil {
		toSerialize["logs_indexed_1day_percentage"] = o.LogsIndexed1dayPercentage
	}
	if o.LogsIndexed1dayUsage != nil {
		toSerialize["logs_indexed_1day_usage"] = o.LogsIndexed1dayUsage
	}
	if o.LogsIndexed30dayPercentage != nil {
		toSerialize["logs_indexed_30day_percentage"] = o.LogsIndexed30dayPercentage
	}
	if o.LogsIndexed30dayUsage != nil {
		toSerialize["logs_indexed_30day_usage"] = o.LogsIndexed30dayUsage
	}
	if o.LogsIndexed360dayPercentage != nil {
		toSerialize["logs_indexed_360day_percentage"] = o.LogsIndexed360dayPercentage
	}
	if o.LogsIndexed360dayUsage != nil {
		toSerialize["logs_indexed_360day_usage"] = o.LogsIndexed360dayUsage
	}
	if o.LogsIndexed3dayPercentage != nil {
		toSerialize["logs_indexed_3day_percentage"] = o.LogsIndexed3dayPercentage
	}
	if o.LogsIndexed3dayUsage != nil {
		toSerialize["logs_indexed_3day_usage"] = o.LogsIndexed3dayUsage
	}
	if o.LogsIndexed45dayPercentage != nil {
		toSerialize["logs_indexed_45day_percentage"] = o.LogsIndexed45dayPercentage
	}
	if o.LogsIndexed45dayUsage != nil {
		toSerialize["logs_indexed_45day_usage"] = o.LogsIndexed45dayUsage
	}
	if o.LogsIndexed60dayPercentage != nil {
		toSerialize["logs_indexed_60day_percentage"] = o.LogsIndexed60dayPercentage
	}
	if o.LogsIndexed60dayUsage != nil {
		toSerialize["logs_indexed_60day_usage"] = o.LogsIndexed60dayUsage
	}
	if o.LogsIndexed7dayPercentage != nil {
		toSerialize["logs_indexed_7day_percentage"] = o.LogsIndexed7dayPercentage
	}
	if o.LogsIndexed7dayUsage != nil {
		toSerialize["logs_indexed_7day_usage"] = o.LogsIndexed7dayUsage
	}
	if o.LogsIndexed90dayPercentage != nil {
		toSerialize["logs_indexed_90day_percentage"] = o.LogsIndexed90dayPercentage
	}
	if o.LogsIndexed90dayUsage != nil {
		toSerialize["logs_indexed_90day_usage"] = o.LogsIndexed90dayUsage
	}
	if o.LogsIndexedCustomRetentionPercentage != nil {
		toSerialize["logs_indexed_custom_retention_percentage"] = o.LogsIndexedCustomRetentionPercentage
	}
	if o.LogsIndexedCustomRetentionUsage != nil {
		toSerialize["logs_indexed_custom_retention_usage"] = o.LogsIndexedCustomRetentionUsage
	}
	if o.MobileAppTestingPercentage != nil {
		toSerialize["mobile_app_testing_percentage"] = o.MobileAppTestingPercentage
	}
	if o.MobileAppTestingUsage != nil {
		toSerialize["mobile_app_testing_usage"] = o.MobileAppTestingUsage
	}
	if o.NdmNetflowPercentage != nil {
		toSerialize["ndm_netflow_percentage"] = o.NdmNetflowPercentage
	}
	if o.NdmNetflowUsage != nil {
		toSerialize["ndm_netflow_usage"] = o.NdmNetflowUsage
	}
	if o.NetworkDeviceWirelessPercentage != nil {
		toSerialize["network_device_wireless_percentage"] = o.NetworkDeviceWirelessPercentage
	}
	if o.NetworkDeviceWirelessUsage != nil {
		toSerialize["network_device_wireless_usage"] = o.NetworkDeviceWirelessUsage
	}
	if o.NpmHostPercentage != nil {
		toSerialize["npm_host_percentage"] = o.NpmHostPercentage
	}
	if o.NpmHostUsage != nil {
		toSerialize["npm_host_usage"] = o.NpmHostUsage
	}
	if o.ObsPipelineBytesPercentage != nil {
		toSerialize["obs_pipeline_bytes_percentage"] = o.ObsPipelineBytesPercentage
	}
	if o.ObsPipelineBytesUsage != nil {
		toSerialize["obs_pipeline_bytes_usage"] = o.ObsPipelineBytesUsage
	}
	if o.ObsPipelinesVcpuPercentage != nil {
		toSerialize["obs_pipelines_vcpu_percentage"] = o.ObsPipelinesVcpuPercentage
	}
	if o.ObsPipelinesVcpuUsage != nil {
		toSerialize["obs_pipelines_vcpu_usage"] = o.ObsPipelinesVcpuUsage
	}
	if o.OnlineArchivePercentage != nil {
		toSerialize["online_archive_percentage"] = o.OnlineArchivePercentage
	}
	if o.OnlineArchiveUsage != nil {
		toSerialize["online_archive_usage"] = o.OnlineArchiveUsage
	}
	if o.ProductAnalyticsSessionPercentage != nil {
		toSerialize["product_analytics_session_percentage"] = o.ProductAnalyticsSessionPercentage
	}
	if o.ProductAnalyticsSessionUsage != nil {
		toSerialize["product_analytics_session_usage"] = o.ProductAnalyticsSessionUsage
	}
	if o.ProfiledContainerPercentage != nil {
		toSerialize["profiled_container_percentage"] = o.ProfiledContainerPercentage
	}
	if o.ProfiledContainerUsage != nil {
		toSerialize["profiled_container_usage"] = o.ProfiledContainerUsage
	}
	if o.ProfiledFargatePercentage != nil {
		toSerialize["profiled_fargate_percentage"] = o.ProfiledFargatePercentage
	}
	if o.ProfiledFargateUsage != nil {
		toSerialize["profiled_fargate_usage"] = o.ProfiledFargateUsage
	}
	if o.ProfiledHostPercentage != nil {
		toSerialize["profiled_host_percentage"] = o.ProfiledHostPercentage
	}
	if o.ProfiledHostUsage != nil {
		toSerialize["profiled_host_usage"] = o.ProfiledHostUsage
	}
	if o.PublishedAppPercentage != nil {
		toSerialize["published_app_percentage"] = o.PublishedAppPercentage
	}
	if o.PublishedAppUsage != nil {
		toSerialize["published_app_usage"] = o.PublishedAppUsage
	}
	if o.RumBrowserMobileSessionsPercentage != nil {
		toSerialize["rum_browser_mobile_sessions_percentage"] = o.RumBrowserMobileSessionsPercentage
	}
	if o.RumBrowserMobileSessionsUsage != nil {
		toSerialize["rum_browser_mobile_sessions_usage"] = o.RumBrowserMobileSessionsUsage
	}
	if o.RumIngestedPercentage != nil {
		toSerialize["rum_ingested_percentage"] = o.RumIngestedPercentage
	}
	if o.RumIngestedUsage != nil {
		toSerialize["rum_ingested_usage"] = o.RumIngestedUsage
	}
	if o.RumInvestigatePercentage != nil {
		toSerialize["rum_investigate_percentage"] = o.RumInvestigatePercentage
	}
	if o.RumInvestigateUsage != nil {
		toSerialize["rum_investigate_usage"] = o.RumInvestigateUsage
	}
	if o.RumReplaySessionsPercentage != nil {
		toSerialize["rum_replay_sessions_percentage"] = o.RumReplaySessionsPercentage
	}
	if o.RumReplaySessionsUsage != nil {
		toSerialize["rum_replay_sessions_usage"] = o.RumReplaySessionsUsage
	}
	if o.RumSessionReplayAddOnPercentage != nil {
		toSerialize["rum_session_replay_add_on_percentage"] = o.RumSessionReplayAddOnPercentage
	}
	if o.RumSessionReplayAddOnUsage != nil {
		toSerialize["rum_session_replay_add_on_usage"] = o.RumSessionReplayAddOnUsage
	}
	if o.ScaFargatePercentage != nil {
		toSerialize["sca_fargate_percentage"] = o.ScaFargatePercentage
	}
	if o.ScaFargateUsage != nil {
		toSerialize["sca_fargate_usage"] = o.ScaFargateUsage
	}
	if o.SdsScannedBytesPercentage != nil {
		toSerialize["sds_scanned_bytes_percentage"] = o.SdsScannedBytesPercentage
	}
	if o.SdsScannedBytesUsage != nil {
		toSerialize["sds_scanned_bytes_usage"] = o.SdsScannedBytesUsage
	}
	if o.ServerlessAppsPercentage != nil {
		toSerialize["serverless_apps_percentage"] = o.ServerlessAppsPercentage
	}
	if o.ServerlessAppsUsage != nil {
		toSerialize["serverless_apps_usage"] = o.ServerlessAppsUsage
	}
	if o.SiemAnalyzedLogsAddOnPercentage != nil {
		toSerialize["siem_analyzed_logs_add_on_percentage"] = o.SiemAnalyzedLogsAddOnPercentage
	}
	if o.SiemAnalyzedLogsAddOnUsage != nil {
		toSerialize["siem_analyzed_logs_add_on_usage"] = o.SiemAnalyzedLogsAddOnUsage
	}
	if o.SiemIngestedBytesPercentage != nil {
		toSerialize["siem_ingested_bytes_percentage"] = o.SiemIngestedBytesPercentage
	}
	if o.SiemIngestedBytesUsage != nil {
		toSerialize["siem_ingested_bytes_usage"] = o.SiemIngestedBytesUsage
	}
	if o.SnmpPercentage != nil {
		toSerialize["snmp_percentage"] = o.SnmpPercentage
	}
	if o.SnmpUsage != nil {
		toSerialize["snmp_usage"] = o.SnmpUsage
	}
	if o.UniversalServiceMonitoringPercentage != nil {
		toSerialize["universal_service_monitoring_percentage"] = o.UniversalServiceMonitoringPercentage
	}
	if o.UniversalServiceMonitoringUsage != nil {
		toSerialize["universal_service_monitoring_usage"] = o.UniversalServiceMonitoringUsage
	}
	if o.VulnManagementHostsPercentage != nil {
		toSerialize["vuln_management_hosts_percentage"] = o.VulnManagementHostsPercentage
	}
	if o.VulnManagementHostsUsage != nil {
		toSerialize["vuln_management_hosts_usage"] = o.VulnManagementHostsUsage
	}
	if o.WorkflowExecutionsPercentage != nil {
		toSerialize["workflow_executions_percentage"] = o.WorkflowExecutionsPercentage
	}
	if o.WorkflowExecutionsUsage != nil {
		toSerialize["workflow_executions_usage"] = o.WorkflowExecutionsUsage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonthlyUsageAttributionValues) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiPercentage                                  *float64 `json:"api_percentage,omitempty"`
		ApiUsage                                       *float64 `json:"api_usage,omitempty"`
		ApmFargatePercentage                           *float64 `json:"apm_fargate_percentage,omitempty"`
		ApmFargateUsage                                *float64 `json:"apm_fargate_usage,omitempty"`
		ApmHostPercentage                              *float64 `json:"apm_host_percentage,omitempty"`
		ApmHostUsage                                   *float64 `json:"apm_host_usage,omitempty"`
		ApmUsmPercentage                               *float64 `json:"apm_usm_percentage,omitempty"`
		ApmUsmUsage                                    *float64 `json:"apm_usm_usage,omitempty"`
		AppsecFargatePercentage                        *float64 `json:"appsec_fargate_percentage,omitempty"`
		AppsecFargateUsage                             *float64 `json:"appsec_fargate_usage,omitempty"`
		AppsecPercentage                               *float64 `json:"appsec_percentage,omitempty"`
		AppsecUsage                                    *float64 `json:"appsec_usage,omitempty"`
		AsmServerlessTracedInvocationsPercentage       *float64 `json:"asm_serverless_traced_invocations_percentage,omitempty"`
		AsmServerlessTracedInvocationsUsage            *float64 `json:"asm_serverless_traced_invocations_usage,omitempty"`
		BrowserPercentage                              *float64 `json:"browser_percentage,omitempty"`
		BrowserUsage                                   *float64 `json:"browser_usage,omitempty"`
		CiPipelineIndexedSpansPercentage               *float64 `json:"ci_pipeline_indexed_spans_percentage,omitempty"`
		CiPipelineIndexedSpansUsage                    *float64 `json:"ci_pipeline_indexed_spans_usage,omitempty"`
		CiTestIndexedSpansPercentage                   *float64 `json:"ci_test_indexed_spans_percentage,omitempty"`
		CiTestIndexedSpansUsage                        *float64 `json:"ci_test_indexed_spans_usage,omitempty"`
		CiVisibilityItrPercentage                      *float64 `json:"ci_visibility_itr_percentage,omitempty"`
		CiVisibilityItrUsage                           *float64 `json:"ci_visibility_itr_usage,omitempty"`
		CloudSiemPercentage                            *float64 `json:"cloud_siem_percentage,omitempty"`
		CloudSiemUsage                                 *float64 `json:"cloud_siem_usage,omitempty"`
		CodeSecurityHostPercentage                     *float64 `json:"code_security_host_percentage,omitempty"`
		CodeSecurityHostUsage                          *float64 `json:"code_security_host_usage,omitempty"`
		ContainerExclAgentPercentage                   *float64 `json:"container_excl_agent_percentage,omitempty"`
		ContainerExclAgentUsage                        *float64 `json:"container_excl_agent_usage,omitempty"`
		ContainerPercentage                            *float64 `json:"container_percentage,omitempty"`
		ContainerUsage                                 *float64 `json:"container_usage,omitempty"`
		CspmContainersPercentage                       *float64 `json:"cspm_containers_percentage,omitempty"`
		CspmContainersUsage                            *float64 `json:"cspm_containers_usage,omitempty"`
		CspmHostsPercentage                            *float64 `json:"cspm_hosts_percentage,omitempty"`
		CspmHostsUsage                                 *float64 `json:"cspm_hosts_usage,omitempty"`
		CustomEventPercentage                          *float64 `json:"custom_event_percentage,omitempty"`
		CustomEventUsage                               *float64 `json:"custom_event_usage,omitempty"`
		CustomIngestedTimeseriesPercentage             *float64 `json:"custom_ingested_timeseries_percentage,omitempty"`
		CustomIngestedTimeseriesUsage                  *float64 `json:"custom_ingested_timeseries_usage,omitempty"`
		CustomTimeseriesPercentage                     *float64 `json:"custom_timeseries_percentage,omitempty"`
		CustomTimeseriesUsage                          *float64 `json:"custom_timeseries_usage,omitempty"`
		CwsContainersPercentage                        *float64 `json:"cws_containers_percentage,omitempty"`
		CwsContainersUsage                             *float64 `json:"cws_containers_usage,omitempty"`
		CwsFargateTaskPercentage                       *float64 `json:"cws_fargate_task_percentage,omitempty"`
		CwsFargateTaskUsage                            *float64 `json:"cws_fargate_task_usage,omitempty"`
		CwsHostsPercentage                             *float64 `json:"cws_hosts_percentage,omitempty"`
		CwsHostsUsage                                  *float64 `json:"cws_hosts_usage,omitempty"`
		DataJobsMonitoringUsage                        *float64 `json:"data_jobs_monitoring_usage,omitempty"`
		DataStreamMonitoringUsage                      *float64 `json:"data_stream_monitoring_usage,omitempty"`
		DbmHostsPercentage                             *float64 `json:"dbm_hosts_percentage,omitempty"`
		DbmHostsUsage                                  *float64 `json:"dbm_hosts_usage,omitempty"`
		DbmQueriesPercentage                           *float64 `json:"dbm_queries_percentage,omitempty"`
		DbmQueriesUsage                                *float64 `json:"dbm_queries_usage,omitempty"`
		ErrorTrackingPercentage                        *float64 `json:"error_tracking_percentage,omitempty"`
		ErrorTrackingUsage                             *float64 `json:"error_tracking_usage,omitempty"`
		EstimatedIndexedSpansPercentage                *float64 `json:"estimated_indexed_spans_percentage,omitempty"`
		EstimatedIndexedSpansUsage                     *float64 `json:"estimated_indexed_spans_usage,omitempty"`
		EstimatedIngestedSpansPercentage               *float64 `json:"estimated_ingested_spans_percentage,omitempty"`
		EstimatedIngestedSpansUsage                    *float64 `json:"estimated_ingested_spans_usage,omitempty"`
		FargatePercentage                              *float64 `json:"fargate_percentage,omitempty"`
		FargateUsage                                   *float64 `json:"fargate_usage,omitempty"`
		FunctionsPercentage                            *float64 `json:"functions_percentage,omitempty"`
		FunctionsUsage                                 *float64 `json:"functions_usage,omitempty"`
		IncidentManagementMonthlyActiveUsersPercentage *float64 `json:"incident_management_monthly_active_users_percentage,omitempty"`
		IncidentManagementMonthlyActiveUsersUsage      *float64 `json:"incident_management_monthly_active_users_usage,omitempty"`
		IndexedSpansPercentage                         *float64 `json:"indexed_spans_percentage,omitempty"`
		IndexedSpansUsage                              *float64 `json:"indexed_spans_usage,omitempty"`
		InfraHostPercentage                            *float64 `json:"infra_host_percentage,omitempty"`
		InfraHostUsage                                 *float64 `json:"infra_host_usage,omitempty"`
		IngestedLogsBytesPercentage                    *float64 `json:"ingested_logs_bytes_percentage,omitempty"`
		IngestedLogsBytesUsage                         *float64 `json:"ingested_logs_bytes_usage,omitempty"`
		IngestedSpansBytesPercentage                   *float64 `json:"ingested_spans_bytes_percentage,omitempty"`
		IngestedSpansBytesUsage                        *float64 `json:"ingested_spans_bytes_usage,omitempty"`
		InvocationsPercentage                          *float64 `json:"invocations_percentage,omitempty"`
		InvocationsUsage                               *float64 `json:"invocations_usage,omitempty"`
		LambdaTracedInvocationsPercentage              *float64 `json:"lambda_traced_invocations_percentage,omitempty"`
		LambdaTracedInvocationsUsage                   *float64 `json:"lambda_traced_invocations_usage,omitempty"`
		LlmObservabilityPercentage                     *float64 `json:"llm_observability_percentage,omitempty"`
		LlmObservabilityUsage                          *float64 `json:"llm_observability_usage,omitempty"`
		LogsIndexed15dayPercentage                     *float64 `json:"logs_indexed_15day_percentage,omitempty"`
		LogsIndexed15dayUsage                          *float64 `json:"logs_indexed_15day_usage,omitempty"`
		LogsIndexed180dayPercentage                    *float64 `json:"logs_indexed_180day_percentage,omitempty"`
		LogsIndexed180dayUsage                         *float64 `json:"logs_indexed_180day_usage,omitempty"`
		LogsIndexed1dayPercentage                      *float64 `json:"logs_indexed_1day_percentage,omitempty"`
		LogsIndexed1dayUsage                           *float64 `json:"logs_indexed_1day_usage,omitempty"`
		LogsIndexed30dayPercentage                     *float64 `json:"logs_indexed_30day_percentage,omitempty"`
		LogsIndexed30dayUsage                          *float64 `json:"logs_indexed_30day_usage,omitempty"`
		LogsIndexed360dayPercentage                    *float64 `json:"logs_indexed_360day_percentage,omitempty"`
		LogsIndexed360dayUsage                         *float64 `json:"logs_indexed_360day_usage,omitempty"`
		LogsIndexed3dayPercentage                      *float64 `json:"logs_indexed_3day_percentage,omitempty"`
		LogsIndexed3dayUsage                           *float64 `json:"logs_indexed_3day_usage,omitempty"`
		LogsIndexed45dayPercentage                     *float64 `json:"logs_indexed_45day_percentage,omitempty"`
		LogsIndexed45dayUsage                          *float64 `json:"logs_indexed_45day_usage,omitempty"`
		LogsIndexed60dayPercentage                     *float64 `json:"logs_indexed_60day_percentage,omitempty"`
		LogsIndexed60dayUsage                          *float64 `json:"logs_indexed_60day_usage,omitempty"`
		LogsIndexed7dayPercentage                      *float64 `json:"logs_indexed_7day_percentage,omitempty"`
		LogsIndexed7dayUsage                           *float64 `json:"logs_indexed_7day_usage,omitempty"`
		LogsIndexed90dayPercentage                     *float64 `json:"logs_indexed_90day_percentage,omitempty"`
		LogsIndexed90dayUsage                          *float64 `json:"logs_indexed_90day_usage,omitempty"`
		LogsIndexedCustomRetentionPercentage           *float64 `json:"logs_indexed_custom_retention_percentage,omitempty"`
		LogsIndexedCustomRetentionUsage                *float64 `json:"logs_indexed_custom_retention_usage,omitempty"`
		MobileAppTestingPercentage                     *float64 `json:"mobile_app_testing_percentage,omitempty"`
		MobileAppTestingUsage                          *float64 `json:"mobile_app_testing_usage,omitempty"`
		NdmNetflowPercentage                           *float64 `json:"ndm_netflow_percentage,omitempty"`
		NdmNetflowUsage                                *float64 `json:"ndm_netflow_usage,omitempty"`
		NetworkDeviceWirelessPercentage                *float64 `json:"network_device_wireless_percentage,omitempty"`
		NetworkDeviceWirelessUsage                     *float64 `json:"network_device_wireless_usage,omitempty"`
		NpmHostPercentage                              *float64 `json:"npm_host_percentage,omitempty"`
		NpmHostUsage                                   *float64 `json:"npm_host_usage,omitempty"`
		ObsPipelineBytesPercentage                     *float64 `json:"obs_pipeline_bytes_percentage,omitempty"`
		ObsPipelineBytesUsage                          *float64 `json:"obs_pipeline_bytes_usage,omitempty"`
		ObsPipelinesVcpuPercentage                     *float64 `json:"obs_pipelines_vcpu_percentage,omitempty"`
		ObsPipelinesVcpuUsage                          *float64 `json:"obs_pipelines_vcpu_usage,omitempty"`
		OnlineArchivePercentage                        *float64 `json:"online_archive_percentage,omitempty"`
		OnlineArchiveUsage                             *float64 `json:"online_archive_usage,omitempty"`
		ProductAnalyticsSessionPercentage              *float64 `json:"product_analytics_session_percentage,omitempty"`
		ProductAnalyticsSessionUsage                   *float64 `json:"product_analytics_session_usage,omitempty"`
		ProfiledContainerPercentage                    *float64 `json:"profiled_container_percentage,omitempty"`
		ProfiledContainerUsage                         *float64 `json:"profiled_container_usage,omitempty"`
		ProfiledFargatePercentage                      *float64 `json:"profiled_fargate_percentage,omitempty"`
		ProfiledFargateUsage                           *float64 `json:"profiled_fargate_usage,omitempty"`
		ProfiledHostPercentage                         *float64 `json:"profiled_host_percentage,omitempty"`
		ProfiledHostUsage                              *float64 `json:"profiled_host_usage,omitempty"`
		PublishedAppPercentage                         *float64 `json:"published_app_percentage,omitempty"`
		PublishedAppUsage                              *float64 `json:"published_app_usage,omitempty"`
		RumBrowserMobileSessionsPercentage             *float64 `json:"rum_browser_mobile_sessions_percentage,omitempty"`
		RumBrowserMobileSessionsUsage                  *float64 `json:"rum_browser_mobile_sessions_usage,omitempty"`
		RumIngestedPercentage                          *float64 `json:"rum_ingested_percentage,omitempty"`
		RumIngestedUsage                               *float64 `json:"rum_ingested_usage,omitempty"`
		RumInvestigatePercentage                       *float64 `json:"rum_investigate_percentage,omitempty"`
		RumInvestigateUsage                            *float64 `json:"rum_investigate_usage,omitempty"`
		RumReplaySessionsPercentage                    *float64 `json:"rum_replay_sessions_percentage,omitempty"`
		RumReplaySessionsUsage                         *float64 `json:"rum_replay_sessions_usage,omitempty"`
		RumSessionReplayAddOnPercentage                *float64 `json:"rum_session_replay_add_on_percentage,omitempty"`
		RumSessionReplayAddOnUsage                     *float64 `json:"rum_session_replay_add_on_usage,omitempty"`
		ScaFargatePercentage                           *float64 `json:"sca_fargate_percentage,omitempty"`
		ScaFargateUsage                                *float64 `json:"sca_fargate_usage,omitempty"`
		SdsScannedBytesPercentage                      *float64 `json:"sds_scanned_bytes_percentage,omitempty"`
		SdsScannedBytesUsage                           *float64 `json:"sds_scanned_bytes_usage,omitempty"`
		ServerlessAppsPercentage                       *float64 `json:"serverless_apps_percentage,omitempty"`
		ServerlessAppsUsage                            *float64 `json:"serverless_apps_usage,omitempty"`
		SiemAnalyzedLogsAddOnPercentage                *float64 `json:"siem_analyzed_logs_add_on_percentage,omitempty"`
		SiemAnalyzedLogsAddOnUsage                     *float64 `json:"siem_analyzed_logs_add_on_usage,omitempty"`
		SiemIngestedBytesPercentage                    *float64 `json:"siem_ingested_bytes_percentage,omitempty"`
		SiemIngestedBytesUsage                         *float64 `json:"siem_ingested_bytes_usage,omitempty"`
		SnmpPercentage                                 *float64 `json:"snmp_percentage,omitempty"`
		SnmpUsage                                      *float64 `json:"snmp_usage,omitempty"`
		UniversalServiceMonitoringPercentage           *float64 `json:"universal_service_monitoring_percentage,omitempty"`
		UniversalServiceMonitoringUsage                *float64 `json:"universal_service_monitoring_usage,omitempty"`
		VulnManagementHostsPercentage                  *float64 `json:"vuln_management_hosts_percentage,omitempty"`
		VulnManagementHostsUsage                       *float64 `json:"vuln_management_hosts_usage,omitempty"`
		WorkflowExecutionsPercentage                   *float64 `json:"workflow_executions_percentage,omitempty"`
		WorkflowExecutionsUsage                        *float64 `json:"workflow_executions_usage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_percentage", "api_usage", "apm_fargate_percentage", "apm_fargate_usage", "apm_host_percentage", "apm_host_usage", "apm_usm_percentage", "apm_usm_usage", "appsec_fargate_percentage", "appsec_fargate_usage", "appsec_percentage", "appsec_usage", "asm_serverless_traced_invocations_percentage", "asm_serverless_traced_invocations_usage", "browser_percentage", "browser_usage", "ci_pipeline_indexed_spans_percentage", "ci_pipeline_indexed_spans_usage", "ci_test_indexed_spans_percentage", "ci_test_indexed_spans_usage", "ci_visibility_itr_percentage", "ci_visibility_itr_usage", "cloud_siem_percentage", "cloud_siem_usage", "code_security_host_percentage", "code_security_host_usage", "container_excl_agent_percentage", "container_excl_agent_usage", "container_percentage", "container_usage", "cspm_containers_percentage", "cspm_containers_usage", "cspm_hosts_percentage", "cspm_hosts_usage", "custom_event_percentage", "custom_event_usage", "custom_ingested_timeseries_percentage", "custom_ingested_timeseries_usage", "custom_timeseries_percentage", "custom_timeseries_usage", "cws_containers_percentage", "cws_containers_usage", "cws_fargate_task_percentage", "cws_fargate_task_usage", "cws_hosts_percentage", "cws_hosts_usage", "data_jobs_monitoring_usage", "data_stream_monitoring_usage", "dbm_hosts_percentage", "dbm_hosts_usage", "dbm_queries_percentage", "dbm_queries_usage", "error_tracking_percentage", "error_tracking_usage", "estimated_indexed_spans_percentage", "estimated_indexed_spans_usage", "estimated_ingested_spans_percentage", "estimated_ingested_spans_usage", "fargate_percentage", "fargate_usage", "functions_percentage", "functions_usage", "incident_management_monthly_active_users_percentage", "incident_management_monthly_active_users_usage", "indexed_spans_percentage", "indexed_spans_usage", "infra_host_percentage", "infra_host_usage", "ingested_logs_bytes_percentage", "ingested_logs_bytes_usage", "ingested_spans_bytes_percentage", "ingested_spans_bytes_usage", "invocations_percentage", "invocations_usage", "lambda_traced_invocations_percentage", "lambda_traced_invocations_usage", "llm_observability_percentage", "llm_observability_usage", "logs_indexed_15day_percentage", "logs_indexed_15day_usage", "logs_indexed_180day_percentage", "logs_indexed_180day_usage", "logs_indexed_1day_percentage", "logs_indexed_1day_usage", "logs_indexed_30day_percentage", "logs_indexed_30day_usage", "logs_indexed_360day_percentage", "logs_indexed_360day_usage", "logs_indexed_3day_percentage", "logs_indexed_3day_usage", "logs_indexed_45day_percentage", "logs_indexed_45day_usage", "logs_indexed_60day_percentage", "logs_indexed_60day_usage", "logs_indexed_7day_percentage", "logs_indexed_7day_usage", "logs_indexed_90day_percentage", "logs_indexed_90day_usage", "logs_indexed_custom_retention_percentage", "logs_indexed_custom_retention_usage", "mobile_app_testing_percentage", "mobile_app_testing_usage", "ndm_netflow_percentage", "ndm_netflow_usage", "network_device_wireless_percentage", "network_device_wireless_usage", "npm_host_percentage", "npm_host_usage", "obs_pipeline_bytes_percentage", "obs_pipeline_bytes_usage", "obs_pipelines_vcpu_percentage", "obs_pipelines_vcpu_usage", "online_archive_percentage", "online_archive_usage", "product_analytics_session_percentage", "product_analytics_session_usage", "profiled_container_percentage", "profiled_container_usage", "profiled_fargate_percentage", "profiled_fargate_usage", "profiled_host_percentage", "profiled_host_usage", "published_app_percentage", "published_app_usage", "rum_browser_mobile_sessions_percentage", "rum_browser_mobile_sessions_usage", "rum_ingested_percentage", "rum_ingested_usage", "rum_investigate_percentage", "rum_investigate_usage", "rum_replay_sessions_percentage", "rum_replay_sessions_usage", "rum_session_replay_add_on_percentage", "rum_session_replay_add_on_usage", "sca_fargate_percentage", "sca_fargate_usage", "sds_scanned_bytes_percentage", "sds_scanned_bytes_usage", "serverless_apps_percentage", "serverless_apps_usage", "siem_analyzed_logs_add_on_percentage", "siem_analyzed_logs_add_on_usage", "siem_ingested_bytes_percentage", "siem_ingested_bytes_usage", "snmp_percentage", "snmp_usage", "universal_service_monitoring_percentage", "universal_service_monitoring_usage", "vuln_management_hosts_percentage", "vuln_management_hosts_usage", "workflow_executions_percentage", "workflow_executions_usage"})
	} else {
		return err
	}
	o.ApiPercentage = all.ApiPercentage
	o.ApiUsage = all.ApiUsage
	o.ApmFargatePercentage = all.ApmFargatePercentage
	o.ApmFargateUsage = all.ApmFargateUsage
	o.ApmHostPercentage = all.ApmHostPercentage
	o.ApmHostUsage = all.ApmHostUsage
	o.ApmUsmPercentage = all.ApmUsmPercentage
	o.ApmUsmUsage = all.ApmUsmUsage
	o.AppsecFargatePercentage = all.AppsecFargatePercentage
	o.AppsecFargateUsage = all.AppsecFargateUsage
	o.AppsecPercentage = all.AppsecPercentage
	o.AppsecUsage = all.AppsecUsage
	o.AsmServerlessTracedInvocationsPercentage = all.AsmServerlessTracedInvocationsPercentage
	o.AsmServerlessTracedInvocationsUsage = all.AsmServerlessTracedInvocationsUsage
	o.BrowserPercentage = all.BrowserPercentage
	o.BrowserUsage = all.BrowserUsage
	o.CiPipelineIndexedSpansPercentage = all.CiPipelineIndexedSpansPercentage
	o.CiPipelineIndexedSpansUsage = all.CiPipelineIndexedSpansUsage
	o.CiTestIndexedSpansPercentage = all.CiTestIndexedSpansPercentage
	o.CiTestIndexedSpansUsage = all.CiTestIndexedSpansUsage
	o.CiVisibilityItrPercentage = all.CiVisibilityItrPercentage
	o.CiVisibilityItrUsage = all.CiVisibilityItrUsage
	o.CloudSiemPercentage = all.CloudSiemPercentage
	o.CloudSiemUsage = all.CloudSiemUsage
	o.CodeSecurityHostPercentage = all.CodeSecurityHostPercentage
	o.CodeSecurityHostUsage = all.CodeSecurityHostUsage
	o.ContainerExclAgentPercentage = all.ContainerExclAgentPercentage
	o.ContainerExclAgentUsage = all.ContainerExclAgentUsage
	o.ContainerPercentage = all.ContainerPercentage
	o.ContainerUsage = all.ContainerUsage
	o.CspmContainersPercentage = all.CspmContainersPercentage
	o.CspmContainersUsage = all.CspmContainersUsage
	o.CspmHostsPercentage = all.CspmHostsPercentage
	o.CspmHostsUsage = all.CspmHostsUsage
	o.CustomEventPercentage = all.CustomEventPercentage
	o.CustomEventUsage = all.CustomEventUsage
	o.CustomIngestedTimeseriesPercentage = all.CustomIngestedTimeseriesPercentage
	o.CustomIngestedTimeseriesUsage = all.CustomIngestedTimeseriesUsage
	o.CustomTimeseriesPercentage = all.CustomTimeseriesPercentage
	o.CustomTimeseriesUsage = all.CustomTimeseriesUsage
	o.CwsContainersPercentage = all.CwsContainersPercentage
	o.CwsContainersUsage = all.CwsContainersUsage
	o.CwsFargateTaskPercentage = all.CwsFargateTaskPercentage
	o.CwsFargateTaskUsage = all.CwsFargateTaskUsage
	o.CwsHostsPercentage = all.CwsHostsPercentage
	o.CwsHostsUsage = all.CwsHostsUsage
	o.DataJobsMonitoringUsage = all.DataJobsMonitoringUsage
	o.DataStreamMonitoringUsage = all.DataStreamMonitoringUsage
	o.DbmHostsPercentage = all.DbmHostsPercentage
	o.DbmHostsUsage = all.DbmHostsUsage
	o.DbmQueriesPercentage = all.DbmQueriesPercentage
	o.DbmQueriesUsage = all.DbmQueriesUsage
	o.ErrorTrackingPercentage = all.ErrorTrackingPercentage
	o.ErrorTrackingUsage = all.ErrorTrackingUsage
	o.EstimatedIndexedSpansPercentage = all.EstimatedIndexedSpansPercentage
	o.EstimatedIndexedSpansUsage = all.EstimatedIndexedSpansUsage
	o.EstimatedIngestedSpansPercentage = all.EstimatedIngestedSpansPercentage
	o.EstimatedIngestedSpansUsage = all.EstimatedIngestedSpansUsage
	o.FargatePercentage = all.FargatePercentage
	o.FargateUsage = all.FargateUsage
	o.FunctionsPercentage = all.FunctionsPercentage
	o.FunctionsUsage = all.FunctionsUsage
	o.IncidentManagementMonthlyActiveUsersPercentage = all.IncidentManagementMonthlyActiveUsersPercentage
	o.IncidentManagementMonthlyActiveUsersUsage = all.IncidentManagementMonthlyActiveUsersUsage
	o.IndexedSpansPercentage = all.IndexedSpansPercentage
	o.IndexedSpansUsage = all.IndexedSpansUsage
	o.InfraHostPercentage = all.InfraHostPercentage
	o.InfraHostUsage = all.InfraHostUsage
	o.IngestedLogsBytesPercentage = all.IngestedLogsBytesPercentage
	o.IngestedLogsBytesUsage = all.IngestedLogsBytesUsage
	o.IngestedSpansBytesPercentage = all.IngestedSpansBytesPercentage
	o.IngestedSpansBytesUsage = all.IngestedSpansBytesUsage
	o.InvocationsPercentage = all.InvocationsPercentage
	o.InvocationsUsage = all.InvocationsUsage
	o.LambdaTracedInvocationsPercentage = all.LambdaTracedInvocationsPercentage
	o.LambdaTracedInvocationsUsage = all.LambdaTracedInvocationsUsage
	o.LlmObservabilityPercentage = all.LlmObservabilityPercentage
	o.LlmObservabilityUsage = all.LlmObservabilityUsage
	o.LogsIndexed15dayPercentage = all.LogsIndexed15dayPercentage
	o.LogsIndexed15dayUsage = all.LogsIndexed15dayUsage
	o.LogsIndexed180dayPercentage = all.LogsIndexed180dayPercentage
	o.LogsIndexed180dayUsage = all.LogsIndexed180dayUsage
	o.LogsIndexed1dayPercentage = all.LogsIndexed1dayPercentage
	o.LogsIndexed1dayUsage = all.LogsIndexed1dayUsage
	o.LogsIndexed30dayPercentage = all.LogsIndexed30dayPercentage
	o.LogsIndexed30dayUsage = all.LogsIndexed30dayUsage
	o.LogsIndexed360dayPercentage = all.LogsIndexed360dayPercentage
	o.LogsIndexed360dayUsage = all.LogsIndexed360dayUsage
	o.LogsIndexed3dayPercentage = all.LogsIndexed3dayPercentage
	o.LogsIndexed3dayUsage = all.LogsIndexed3dayUsage
	o.LogsIndexed45dayPercentage = all.LogsIndexed45dayPercentage
	o.LogsIndexed45dayUsage = all.LogsIndexed45dayUsage
	o.LogsIndexed60dayPercentage = all.LogsIndexed60dayPercentage
	o.LogsIndexed60dayUsage = all.LogsIndexed60dayUsage
	o.LogsIndexed7dayPercentage = all.LogsIndexed7dayPercentage
	o.LogsIndexed7dayUsage = all.LogsIndexed7dayUsage
	o.LogsIndexed90dayPercentage = all.LogsIndexed90dayPercentage
	o.LogsIndexed90dayUsage = all.LogsIndexed90dayUsage
	o.LogsIndexedCustomRetentionPercentage = all.LogsIndexedCustomRetentionPercentage
	o.LogsIndexedCustomRetentionUsage = all.LogsIndexedCustomRetentionUsage
	o.MobileAppTestingPercentage = all.MobileAppTestingPercentage
	o.MobileAppTestingUsage = all.MobileAppTestingUsage
	o.NdmNetflowPercentage = all.NdmNetflowPercentage
	o.NdmNetflowUsage = all.NdmNetflowUsage
	o.NetworkDeviceWirelessPercentage = all.NetworkDeviceWirelessPercentage
	o.NetworkDeviceWirelessUsage = all.NetworkDeviceWirelessUsage
	o.NpmHostPercentage = all.NpmHostPercentage
	o.NpmHostUsage = all.NpmHostUsage
	o.ObsPipelineBytesPercentage = all.ObsPipelineBytesPercentage
	o.ObsPipelineBytesUsage = all.ObsPipelineBytesUsage
	o.ObsPipelinesVcpuPercentage = all.ObsPipelinesVcpuPercentage
	o.ObsPipelinesVcpuUsage = all.ObsPipelinesVcpuUsage
	o.OnlineArchivePercentage = all.OnlineArchivePercentage
	o.OnlineArchiveUsage = all.OnlineArchiveUsage
	o.ProductAnalyticsSessionPercentage = all.ProductAnalyticsSessionPercentage
	o.ProductAnalyticsSessionUsage = all.ProductAnalyticsSessionUsage
	o.ProfiledContainerPercentage = all.ProfiledContainerPercentage
	o.ProfiledContainerUsage = all.ProfiledContainerUsage
	o.ProfiledFargatePercentage = all.ProfiledFargatePercentage
	o.ProfiledFargateUsage = all.ProfiledFargateUsage
	o.ProfiledHostPercentage = all.ProfiledHostPercentage
	o.ProfiledHostUsage = all.ProfiledHostUsage
	o.PublishedAppPercentage = all.PublishedAppPercentage
	o.PublishedAppUsage = all.PublishedAppUsage
	o.RumBrowserMobileSessionsPercentage = all.RumBrowserMobileSessionsPercentage
	o.RumBrowserMobileSessionsUsage = all.RumBrowserMobileSessionsUsage
	o.RumIngestedPercentage = all.RumIngestedPercentage
	o.RumIngestedUsage = all.RumIngestedUsage
	o.RumInvestigatePercentage = all.RumInvestigatePercentage
	o.RumInvestigateUsage = all.RumInvestigateUsage
	o.RumReplaySessionsPercentage = all.RumReplaySessionsPercentage
	o.RumReplaySessionsUsage = all.RumReplaySessionsUsage
	o.RumSessionReplayAddOnPercentage = all.RumSessionReplayAddOnPercentage
	o.RumSessionReplayAddOnUsage = all.RumSessionReplayAddOnUsage
	o.ScaFargatePercentage = all.ScaFargatePercentage
	o.ScaFargateUsage = all.ScaFargateUsage
	o.SdsScannedBytesPercentage = all.SdsScannedBytesPercentage
	o.SdsScannedBytesUsage = all.SdsScannedBytesUsage
	o.ServerlessAppsPercentage = all.ServerlessAppsPercentage
	o.ServerlessAppsUsage = all.ServerlessAppsUsage
	o.SiemAnalyzedLogsAddOnPercentage = all.SiemAnalyzedLogsAddOnPercentage
	o.SiemAnalyzedLogsAddOnUsage = all.SiemAnalyzedLogsAddOnUsage
	o.SiemIngestedBytesPercentage = all.SiemIngestedBytesPercentage
	o.SiemIngestedBytesUsage = all.SiemIngestedBytesUsage
	o.SnmpPercentage = all.SnmpPercentage
	o.SnmpUsage = all.SnmpUsage
	o.UniversalServiceMonitoringPercentage = all.UniversalServiceMonitoringPercentage
	o.UniversalServiceMonitoringUsage = all.UniversalServiceMonitoringUsage
	o.VulnManagementHostsPercentage = all.VulnManagementHostsPercentage
	o.VulnManagementHostsUsage = all.VulnManagementHostsUsage
	o.WorkflowExecutionsPercentage = all.WorkflowExecutionsPercentage
	o.WorkflowExecutionsUsage = all.WorkflowExecutionsUsage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
