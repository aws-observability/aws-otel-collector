// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSummaryResponse Response summarizing all usage aggregated across the months in the request for all organizations, and broken down by month and by organization.
type UsageSummaryResponse struct {
	// Shows the 99th percentile of all agent hosts over all hours in the current month for all organizations.
	AgentHostTop99pSum *int64 `json:"agent_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Azure app services using APM over all hours in the current month all organizations.
	ApmAzureAppServiceHostTop99pSum *int64 `json:"apm_azure_app_service_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all APM DevSecOps hosts over all hours in the current month for all organizations.
	ApmDevsecopsHostTop99pSum *int64 `json:"apm_devsecops_host_top99p_sum,omitempty"`
	// Shows the sum of the 99th percentile of all distinct standalone Enterprise hosts over all hours in the current month for all organizations.
	ApmEnterpriseStandaloneHostsTop99pSum *int64 `json:"apm_enterprise_standalone_hosts_top99p_sum,omitempty"`
	// Shows the average of all APM ECS Fargate tasks over all hours in the current month for all organizations.
	ApmFargateCountAvgSum *int64 `json:"apm_fargate_count_avg_sum,omitempty"`
	// Shows the 99th percentile of all distinct APM hosts over all hours in the current month for all organizations.
	ApmHostTop99pSum *int64 `json:"apm_host_top99p_sum,omitempty"`
	// Shows the sum of the 99th percentile of all distinct standalone Pro hosts over all hours in the current month for all organizations.
	ApmProStandaloneHostsTop99pSum *int64 `json:"apm_pro_standalone_hosts_top99p_sum,omitempty"`
	// Shows the average of all Application Security Monitoring ECS Fargate tasks over all hours in the current month for all organizations.
	AppsecFargateCountAvgSum *int64 `json:"appsec_fargate_count_avg_sum,omitempty"`
	// Shows the sum of all Application Security Monitoring Serverless invocations over all hours in the current months for all organizations.
	AsmServerlessAggSum *int64 `json:"asm_serverless_agg_sum,omitempty"`
	// Shows the sum of all audit logs lines indexed over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	AuditLogsLinesIndexedAggSum *int64 `json:"audit_logs_lines_indexed_agg_sum,omitempty"`
	// Shows the total number of organizations that had Audit Trail enabled over a specific number of months.
	AuditTrailEnabledHwmSum *int64 `json:"audit_trail_enabled_hwm_sum,omitempty"`
	// The average total count for Fargate Container Profiler over all hours in the current month for all organizations.
	AvgProfiledFargateTasksSum *int64 `json:"avg_profiled_fargate_tasks_sum,omitempty"`
	// Shows the 99th percentile of all AWS hosts over all hours in the current month for all organizations.
	AwsHostTop99pSum *int64 `json:"aws_host_top99p_sum,omitempty"`
	// Shows the average of the number of functions that executed 1 or more times each hour in the current month for all organizations.
	AwsLambdaFuncCount *int64 `json:"aws_lambda_func_count,omitempty"`
	// Shows the sum of all AWS Lambda invocations over all hours in the current month for all organizations.
	AwsLambdaInvocationsSum *int64 `json:"aws_lambda_invocations_sum,omitempty"`
	// Shows the 99th percentile of all Azure app services over all hours in the current month for all organizations.
	AzureAppServiceTop99pSum *int64 `json:"azure_app_service_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Azure hosts over all hours in the current month for all organizations.
	AzureHostTop99pSum *int64 `json:"azure_host_top99p_sum,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current month for all organizations.
	BillableIngestedBytesAggSum *int64 `json:"billable_ingested_bytes_agg_sum,omitempty"`
	// Shows the sum of all Bits AI Investigations over all hours in the current month for all organizations.
	BitsAiInvestigationsAggSum *int64 `json:"bits_ai_investigations_agg_sum,omitempty"`
	// Shows the sum of all browser lite sessions over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	BrowserRumLiteSessionCountAggSum *int64 `json:"browser_rum_lite_session_count_agg_sum,omitempty"`
	// Shows the sum of all browser replay sessions over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	BrowserRumReplaySessionCountAggSum *int64 `json:"browser_rum_replay_session_count_agg_sum,omitempty"`
	// Shows the sum of all browser RUM units over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	BrowserRumUnitsAggSum *int64 `json:"browser_rum_units_agg_sum,omitempty"`
	// Shows the sum of all CI pipeline indexed spans over all hours in the current month for all organizations.
	CiPipelineIndexedSpansAggSum *int64 `json:"ci_pipeline_indexed_spans_agg_sum,omitempty"`
	// Shows the sum of all CI test indexed spans over all hours in the current month for all organizations.
	CiTestIndexedSpansAggSum *int64 `json:"ci_test_indexed_spans_agg_sum,omitempty"`
	// Shows the high-water mark of all CI visibility intelligent test runner committers over all hours in the current month for all organizations.
	CiVisibilityItrCommittersHwmSum *int64 `json:"ci_visibility_itr_committers_hwm_sum,omitempty"`
	// Shows the high-water mark of all CI visibility pipeline committers over all hours in the current month for all organizations.
	CiVisibilityPipelineCommittersHwmSum *int64 `json:"ci_visibility_pipeline_committers_hwm_sum,omitempty"`
	// Shows the high-water mark of all CI visibility test committers over all hours in the current month for all organizations.
	CiVisibilityTestCommittersHwmSum *int64 `json:"ci_visibility_test_committers_hwm_sum,omitempty"`
	// Sum of the host count average for Cloud Cost Management for AWS.
	CloudCostManagementAwsHostCountAvgSum *int64 `json:"cloud_cost_management_aws_host_count_avg_sum,omitempty"`
	// Sum of the host count average for Cloud Cost Management for Azure.
	CloudCostManagementAzureHostCountAvgSum *int64 `json:"cloud_cost_management_azure_host_count_avg_sum,omitempty"`
	// Sum of the host count average for Cloud Cost Management for GCP.
	CloudCostManagementGcpHostCountAvgSum *int64 `json:"cloud_cost_management_gcp_host_count_avg_sum,omitempty"`
	// Sum of the host count average for Cloud Cost Management for all cloud providers.
	CloudCostManagementHostCountAvgSum *int64 `json:"cloud_cost_management_host_count_avg_sum,omitempty"`
	// Sum of the average host counts for Cloud Cost Management on OCI.
	CloudCostManagementOciHostCountAvgSum *int64 `json:"cloud_cost_management_oci_host_count_avg_sum,omitempty"`
	// Shows the sum of all Cloud Security Information and Event Management events over all hours in the current month for all organizations.
	CloudSiemEventsAggSum *int64 `json:"cloud_siem_events_agg_sum,omitempty"`
	// Shows the high-water mark of all Static Analysis committers over all hours in the current month for all organizations.
	CodeAnalysisSaCommittersHwmSum *int64 `json:"code_analysis_sa_committers_hwm_sum,omitempty"`
	// Shows the high-water mark of all static Software Composition Analysis committers over all hours in the current month for all organizations.
	CodeAnalysisScaCommittersHwmSum *int64 `json:"code_analysis_sca_committers_hwm_sum,omitempty"`
	// Shows the 99th percentile of all Code Security hosts over all hours in the current month for all organizations.
	CodeSecurityHostTop99pSum *int64 `json:"code_security_host_top99p_sum,omitempty"`
	// Shows the average of all distinct containers over all hours in the current month for all organizations.
	ContainerAvgSum *int64 `json:"container_avg_sum,omitempty"`
	// Shows the average of the containers without the Datadog Agent over all hours in the current month for all organizations.
	ContainerExclAgentAvgSum *int64 `json:"container_excl_agent_avg_sum,omitempty"`
	// Shows the sum of the high-water marks of all distinct containers over all hours in the current month for all organizations.
	ContainerHwmSum *int64 `json:"container_hwm_sum,omitempty"`
	// Shows the sum of all Cloud Security Management Enterprise compliance containers over all hours in the current month for all organizations.
	CsmContainerEnterpriseComplianceCountAggSum *int64 `json:"csm_container_enterprise_compliance_count_agg_sum,omitempty"`
	// Shows the sum of all Cloud Security Management Enterprise Cloud Workload Security containers over all hours in the current month for all organizations.
	CsmContainerEnterpriseCwsCountAggSum *int64 `json:"csm_container_enterprise_cws_count_agg_sum,omitempty"`
	// Shows the sum of all Cloud Security Management Enterprise containers over all hours in the current month for all organizations.
	CsmContainerEnterpriseTotalCountAggSum *int64 `json:"csm_container_enterprise_total_count_agg_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise Azure app services hosts over all hours in the current month for all organizations.
	CsmHostEnterpriseAasHostCountTop99pSum *int64 `json:"csm_host_enterprise_aas_host_count_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise AWS hosts over all hours in the current month for all organizations.
	CsmHostEnterpriseAwsHostCountTop99pSum *int64 `json:"csm_host_enterprise_aws_host_count_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise Azure hosts over all hours in the current month for all organizations.
	CsmHostEnterpriseAzureHostCountTop99pSum *int64 `json:"csm_host_enterprise_azure_host_count_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise compliance hosts over all hours in the current month for all organizations.
	CsmHostEnterpriseComplianceHostCountTop99pSum *int64 `json:"csm_host_enterprise_compliance_host_count_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise Cloud Workload Security hosts over all hours in the current month for all organizations.
	CsmHostEnterpriseCwsHostCountTop99pSum *int64 `json:"csm_host_enterprise_cws_host_count_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise GCP hosts over all hours in the current month for all organizations.
	CsmHostEnterpriseGcpHostCountTop99pSum *int64 `json:"csm_host_enterprise_gcp_host_count_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise hosts over all hours in the current month for all organizations.
	CsmHostEnterpriseTotalHostCountTop99pSum *int64 `json:"csm_host_enterprise_total_host_count_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro Azure app services hosts over all hours in the current month for all organizations.
	CspmAasHostTop99pSum *int64 `json:"cspm_aas_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro AWS hosts over all hours in the current month for all organizations.
	CspmAwsHostTop99pSum *int64 `json:"cspm_aws_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro Azure hosts over all hours in the current month for all organizations.
	CspmAzureHostTop99pSum *int64 `json:"cspm_azure_host_top99p_sum,omitempty"`
	// Shows the average number of Cloud Security Management Pro containers over all hours in the current month for all organizations.
	CspmContainerAvgSum *int64 `json:"cspm_container_avg_sum,omitempty"`
	// Shows the sum of the high-water marks of Cloud Security Management Pro containers over all hours in the current month for all organizations.
	CspmContainerHwmSum *int64 `json:"cspm_container_hwm_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro GCP hosts over all hours in the current month for all organizations.
	CspmGcpHostTop99pSum *int64 `json:"cspm_gcp_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro hosts over all hours in the current month for all organizations.
	CspmHostTop99pSum *int64 `json:"cspm_host_top99p_sum,omitempty"`
	// Shows the average number of distinct historical custom metrics over all hours in the current month for all organizations.
	CustomHistoricalTsSum *int64 `json:"custom_historical_ts_sum,omitempty"`
	// Shows the average number of distinct live custom metrics over all hours in the current month for all organizations.
	CustomLiveTsSum *int64 `json:"custom_live_ts_sum,omitempty"`
	// Shows the average number of distinct custom metrics over all hours in the current month for all organizations.
	CustomTsSum *int64 `json:"custom_ts_sum,omitempty"`
	// Shows the average of all distinct Cloud Workload Security containers over all hours in the current month for all organizations.
	CwsContainerAvgSum *int64 `json:"cws_container_avg_sum,omitempty"`
	// Shows the average of all distinct Cloud Workload Security Fargate tasks over all hours in the current month for all organizations.
	CwsFargateTaskAvgSum *int64 `json:"cws_fargate_task_avg_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Workload Security hosts over all hours in the current month for all organizations.
	CwsHostTop99pSum *int64 `json:"cws_host_top99p_sum,omitempty"`
	// Shows the sum of Data Jobs Monitoring hosts over all hours in the current months for all organizations
	DataJobsMonitoringHostHrAggSum *int64 `json:"data_jobs_monitoring_host_hr_agg_sum,omitempty"`
	// Shows the 99th percentile of all Database Monitoring hosts over all hours in the current month for all organizations.
	DbmHostTop99pSum *int64 `json:"dbm_host_top99p_sum,omitempty"`
	// Shows the average of all distinct Database Monitoring Normalized Queries over all hours in the current month for all organizations.
	DbmQueriesAvgSum *int64 `json:"dbm_queries_avg_sum,omitempty"`
	// Shows the last date of usage in the current month for all organizations.
	EndDate *time.Time `json:"end_date,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts with the Datadog Agent over all hours in the current month for all organizations.
	EphInfraHostAgentAggSum *int64 `json:"eph_infra_host_agent_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on Alibaba over all hours in the current month for all organizations.
	EphInfraHostAlibabaAggSum *int64 `json:"eph_infra_host_alibaba_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on AWS over all hours in the current month for all organizations.
	EphInfraHostAwsAggSum *int64 `json:"eph_infra_host_aws_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on Azure over all hours in the current month for all organizations.
	EphInfraHostAzureAggSum *int64 `json:"eph_infra_host_azure_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts for Enterprise over all hours in the current month for all organizations.
	EphInfraHostEntAggSum *int64 `json:"eph_infra_host_ent_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on GCP over all hours in the current month for all organizations.
	EphInfraHostGcpAggSum *int64 `json:"eph_infra_host_gcp_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on Heroku over all hours in the current month for all organizations.
	EphInfraHostHerokuAggSum *int64 `json:"eph_infra_host_heroku_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts with only Azure App Services over all hours in the current month for all organizations.
	EphInfraHostOnlyAasAggSum *int64 `json:"eph_infra_host_only_aas_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts with only vSphere over all hours in the current month for all organizations.
	EphInfraHostOnlyVsphereAggSum *int64 `json:"eph_infra_host_only_vsphere_agg_sum,omitempty"`
	// Shows the sum of all ephemeral hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current month for all organizations.
	EphInfraHostOpentelemetryAggSum *int64 `json:"eph_infra_host_opentelemetry_agg_sum,omitempty"`
	// Shows the sum of all ephemeral APM hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current month for all organizations.
	EphInfraHostOpentelemetryApmAggSum *int64 `json:"eph_infra_host_opentelemetry_apm_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts for Pro over all hours in the current month for all organizations.
	EphInfraHostProAggSum *int64 `json:"eph_infra_host_pro_agg_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts for Pro Plus over all hours in the current month for all organizations.
	EphInfraHostProplusAggSum *int64 `json:"eph_infra_host_proplus_agg_sum,omitempty"`
	// Sum of all ephemeral infrastructure hosts for Proxmox over all hours in the current month for all organizations.
	EphInfraHostProxmoxAggSum *int64 `json:"eph_infra_host_proxmox_agg_sum,omitempty"`
	// Shows the sum of all Error Tracking APM error events over all hours in the current month for all organizations.
	ErrorTrackingApmErrorEventsAggSum *int64 `json:"error_tracking_apm_error_events_agg_sum,omitempty"`
	// Shows the sum of all Error Tracking error events over all hours in the current month for all organizations.
	ErrorTrackingErrorEventsAggSum *int64 `json:"error_tracking_error_events_agg_sum,omitempty"`
	// Shows the sum of all Error Tracking events over all hours in the current months for all organizations.
	ErrorTrackingEventsAggSum *int64 `json:"error_tracking_events_agg_sum,omitempty"`
	// Shows the sum of all Error Tracking RUM error events over all hours in the current month for all organizations.
	ErrorTrackingRumErrorEventsAggSum *int64 `json:"error_tracking_rum_error_events_agg_sum,omitempty"`
	// Shows the sum of all Event Management correlations over all hours in the current month for all organizations.
	EventManagementCorrelationAggSum *int64 `json:"event_management_correlation_agg_sum,omitempty"`
	// Shows the sum of all Event Management correlated events over all hours in the current month for all organizations.
	EventManagementCorrelationCorrelatedEventsAggSum *int64 `json:"event_management_correlation_correlated_events_agg_sum,omitempty"`
	// Shows the sum of all Event Management correlated related events over all hours in the current month for all organizations.
	EventManagementCorrelationCorrelatedRelatedEventsAggSum *int64 `json:"event_management_correlation_correlated_related_events_agg_sum,omitempty"`
	// The average number of Profiling Fargate tasks over all hours in the current month for all organizations.
	FargateContainerProfilerProfilingFargateAvgSum *int64 `json:"fargate_container_profiler_profiling_fargate_avg_sum,omitempty"`
	// The average number of Profiling Fargate Elastic Kubernetes Service tasks over all hours in the current month for all organizations.
	FargateContainerProfilerProfilingFargateEksAvgSum *int64 `json:"fargate_container_profiler_profiling_fargate_eks_avg_sum,omitempty"`
	// Shows the average of all Fargate tasks over all hours in the current month for all organizations.
	FargateTasksCountAvgSum *int64 `json:"fargate_tasks_count_avg_sum,omitempty"`
	// Shows the sum of the high-water marks of all Fargate tasks over all hours in the current month for all organizations.
	FargateTasksCountHwmSum *int64 `json:"fargate_tasks_count_hwm_sum,omitempty"`
	// Shows the average number of Flex Logs Compute Large Instances over all hours in the current months for all organizations.
	FlexLogsComputeLargeAvgSum *int64 `json:"flex_logs_compute_large_avg_sum,omitempty"`
	// Shows the average number of Flex Logs Compute Medium Instances over all hours in the current months for all organizations.
	FlexLogsComputeMediumAvgSum *int64 `json:"flex_logs_compute_medium_avg_sum,omitempty"`
	// Shows the average number of Flex Logs Compute Small Instances over all hours in the current months for all organizations.
	FlexLogsComputeSmallAvgSum *int64 `json:"flex_logs_compute_small_avg_sum,omitempty"`
	// Shows the average number of Flex Logs Compute Extra Large Instances over all hours in the current months for all organizations.
	FlexLogsComputeXlargeAvgSum *int64 `json:"flex_logs_compute_xlarge_avg_sum,omitempty"`
	// Shows the average number of Flex Logs Compute Extra Small Instances over all hours in the current months for all organizations.
	FlexLogsComputeXsmallAvgSum *int64 `json:"flex_logs_compute_xsmall_avg_sum,omitempty"`
	// Shows the average number of Flex Logs Starter Instances over all hours in the current months for all organizations.
	FlexLogsStarterAvgSum *int64 `json:"flex_logs_starter_avg_sum,omitempty"`
	// Shows the average number of Flex Logs Starter Storage Index Instances over all hours in the current months for all organizations.
	FlexLogsStarterStorageIndexAvgSum *int64 `json:"flex_logs_starter_storage_index_avg_sum,omitempty"`
	// Shows the average number of Flex Logs Starter Storage Retention Adjustment Instances over all hours in the current months for all organizations.
	FlexLogsStarterStorageRetentionAdjustmentAvgSum *int64 `json:"flex_logs_starter_storage_retention_adjustment_avg_sum,omitempty"`
	// Shows the average of all Flex Stored Logs over all hours in the current months for all organizations.
	FlexStoredLogsAvgSum *int64 `json:"flex_stored_logs_avg_sum,omitempty"`
	// Shows the sum of all logs forwarding bytes over all hours in the current month for all organizations (data available as of April 1, 2023)
	ForwardingEventsBytesAggSum *int64 `json:"forwarding_events_bytes_agg_sum,omitempty"`
	// Shows the 99th percentile of all GCP hosts over all hours in the current month for all organizations.
	GcpHostTop99pSum *int64 `json:"gcp_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Heroku dynos over all hours in the current month for all organizations.
	HerokuHostTop99pSum *int64 `json:"heroku_host_top99p_sum,omitempty"`
	// Shows sum of the high-water marks of incident management monthly active users in the current month for all organizations.
	IncidentManagementMonthlyActiveUsersHwmSum *int64 `json:"incident_management_monthly_active_users_hwm_sum,omitempty"`
	// Shows the sum of the high-water marks of Incident Management seats over all hours in the current month for all organizations.
	IncidentManagementSeatsHwmSum *int64 `json:"incident_management_seats_hwm_sum,omitempty"`
	// Shows the sum of all log events indexed over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	IndexedEventsCountAggSum *int64 `json:"indexed_events_count_agg_sum,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current month for all organizations.
	InfraHostTop99pSum *int64 `json:"infra_host_top99p_sum,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current month for all organizations.
	IngestedEventsBytesAggSum *int64 `json:"ingested_events_bytes_agg_sum,omitempty"`
	// Shows the sum of all IoT devices over all hours in the current month for all organizations.
	IotDeviceAggSum *int64 `json:"iot_device_agg_sum,omitempty"`
	// Shows the 99th percentile of all IoT devices over all hours in the current month of all organizations.
	IotDeviceTop99pSum *int64 `json:"iot_device_top99p_sum,omitempty"`
	// Shows the most recent hour in the current month for all organizations for which all usages were calculated.
	LastUpdated *time.Time `json:"last_updated,omitempty"`
	// Shows the sum of all live logs indexed over all hours in the current month for all organization (To be deprecated on October 1st, 2024).
	// Deprecated
	LiveIndexedEventsAggSum *int64 `json:"live_indexed_events_agg_sum,omitempty"`
	// Shows the sum of all live logs bytes ingested over all hours in the current month for all organizations (data available as of December 1, 2020).
	LiveIngestedBytesAggSum *int64 `json:"live_ingested_bytes_agg_sum,omitempty"`
	// Sum of all LLM observability sessions for all hours in the current month for all organizations.
	LlmObservabilityAggSum *int64 `json:"llm_observability_agg_sum,omitempty"`
	// Minimum spend for LLM observability sessions for all hours in the current month for all organizations.
	LlmObservabilityMinSpendAggSum *int64 `json:"llm_observability_min_spend_agg_sum,omitempty"`
	// Object containing logs usage data broken down by retention period.
	LogsByRetention *LogsByRetention `json:"logs_by_retention,omitempty"`
	// Shows the sum of all mobile lite sessions over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumLiteSessionCountAggSum *int64 `json:"mobile_rum_lite_session_count_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountAggSum *int64 `json:"mobile_rum_session_count_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on Android over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountAndroidAggSum *int64 `json:"mobile_rum_session_count_android_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on Flutter over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountFlutterAggSum *int64 `json:"mobile_rum_session_count_flutter_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on iOS over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountIosAggSum *int64 `json:"mobile_rum_session_count_ios_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on React Native over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountReactnativeAggSum *int64 `json:"mobile_rum_session_count_reactnative_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on Roku over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountRokuAggSum *int64 `json:"mobile_rum_session_count_roku_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM units over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumUnitsAggSum *int64 `json:"mobile_rum_units_agg_sum,omitempty"`
	// Shows the sum of all Network Device Monitoring NetFlow events over all hours in the current month for all organizations.
	NdmNetflowEventsAggSum *int64 `json:"ndm_netflow_events_agg_sum,omitempty"`
	// Shows the sum of all Network flows indexed over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	NetflowIndexedEventsCountAggSum *int64 `json:"netflow_indexed_events_count_agg_sum,omitempty"`
	// Shows the 99th percentile of all Network Device Monitoring wireless devices over all hours in the current month for all organizations.
	NetworkDeviceWirelessTop99pSum *int64 `json:"network_device_wireless_top99p_sum,omitempty"`
	// Shows the 99th percentile of all distinct Cloud Network Monitoring hosts (formerly known as Network hosts) over all hours in the current month for all organizations.
	NpmHostTop99pSum *int64 `json:"npm_host_top99p_sum,omitempty"`
	// Sum of all observability pipelines bytes processed over all hours in the current month for all organizations.
	ObservabilityPipelinesBytesProcessedAggSum *int64 `json:"observability_pipelines_bytes_processed_agg_sum,omitempty"`
	// Shows the sum of Oracle Cloud Infrastructure hosts over all hours in the current months for all organizations
	OciHostAggSum *int64 `json:"oci_host_agg_sum,omitempty"`
	// Shows the 99th percentile of Oracle Cloud Infrastructure hosts over all hours in the current months for all organizations
	OciHostTop99pSum *int64 `json:"oci_host_top99p_sum,omitempty"`
	// Shows the sum of the high-water marks of On-Call seats over all hours in the current month for all organizations.
	OnCallSeatHwmSum *int64 `json:"on_call_seat_hwm_sum,omitempty"`
	// Sum of all online archived events over all hours in the current month for all organizations.
	OnlineArchiveEventsCountAggSum *int64 `json:"online_archive_events_count_agg_sum,omitempty"`
	// Shows the 99th percentile of APM hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current month for all organizations.
	OpentelemetryApmHostTop99pSum *int64 `json:"opentelemetry_apm_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current month for all organizations.
	OpentelemetryHostTop99pSum *int64 `json:"opentelemetry_host_top99p_sum,omitempty"`
	// Sum of all product analytics sessions for all hours in the current month for all organizations.
	ProductAnalyticsAggSum *int64 `json:"product_analytics_agg_sum,omitempty"`
	// Shows the 99th percentile of all profiled Azure app services over all hours in the current month for all organizations.
	ProfilingAasCountTop99pSum *int64 `json:"profiling_aas_count_top99p_sum,omitempty"`
	// Shows the average number of profiled containers over all hours in the current month for all organizations.
	ProfilingContainerAgentCountAvg *int64 `json:"profiling_container_agent_count_avg,omitempty"`
	// Shows the 99th percentile of all profiled hosts over all hours in the current month for all organizations.
	ProfilingHostCountTop99pSum *int64 `json:"profiling_host_count_top99p_sum,omitempty"`
	// Sum of all Proxmox hosts over all hours in the current month for all organizations.
	ProxmoxHostAggSum *int64 `json:"proxmox_host_agg_sum,omitempty"`
	// Sum of the 99th percentile of all Proxmox hosts over all hours in the current month for all organizations.
	ProxmoxHostTop99pSum *int64 `json:"proxmox_host_top99p_sum,omitempty"`
	// Shows the high-water mark of all published applications over all hours in the current month for all organizations.
	PublishedAppHwmSum *int64 `json:"published_app_hwm_sum,omitempty"`
	// Shows the sum of all rehydrated logs indexed over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	RehydratedIndexedEventsAggSum *int64 `json:"rehydrated_indexed_events_agg_sum,omitempty"`
	// Shows the sum of all rehydrated logs bytes ingested over all hours in the current month for all organizations (data available as of December 1, 2020).
	RehydratedIngestedBytesAggSum *int64 `json:"rehydrated_ingested_bytes_agg_sum,omitempty"`
	// Shows the sum of all mobile sessions and all browser lite and legacy sessions over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	RumBrowserAndMobileSessionCount *int64 `json:"rum_browser_and_mobile_session_count,omitempty"`
	// Shows the sum of all browser RUM legacy sessions over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumBrowserLegacySessionCountAggSum *int64 `json:"rum_browser_legacy_session_count_agg_sum,omitempty"`
	// Shows the sum of all browser RUM lite sessions over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumBrowserLiteSessionCountAggSum *int64 `json:"rum_browser_lite_session_count_agg_sum,omitempty"`
	// Shows the sum of all browser RUM Session Replay counts over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumBrowserReplaySessionCountAggSum *int64 `json:"rum_browser_replay_session_count_agg_sum,omitempty"`
	// Sum of all RUM indexed sessions for all hours in the current month for all organizations.
	RumIndexedSessionsAggSum *int64 `json:"rum_indexed_sessions_agg_sum,omitempty"`
	// Sum of all RUM ingested sessions for all hours in the current month for all organizations.
	RumIngestedSessionsAggSum *int64 `json:"rum_ingested_sessions_agg_sum,omitempty"`
	// Shows the sum of all RUM lite sessions (browser and mobile) over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumLiteSessionCountAggSum *int64 `json:"rum_lite_session_count_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy sessions on Android over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountAndroidAggSum *int64 `json:"rum_mobile_legacy_session_count_android_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy sessions on Flutter over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountFlutterAggSum *int64 `json:"rum_mobile_legacy_session_count_flutter_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy sessions on iOS over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountIosAggSum *int64 `json:"rum_mobile_legacy_session_count_ios_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy sessions on React Native over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountReactnativeAggSum *int64 `json:"rum_mobile_legacy_session_count_reactnative_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy sessions on Roku over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountRokuAggSum *int64 `json:"rum_mobile_legacy_session_count_roku_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Android over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountAndroidAggSum *int64 `json:"rum_mobile_lite_session_count_android_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Flutter over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountFlutterAggSum *int64 `json:"rum_mobile_lite_session_count_flutter_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on iOS over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountIosAggSum *int64 `json:"rum_mobile_lite_session_count_ios_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Kotlin Multiplatform over all hours within the current month for all organizations.
	RumMobileLiteSessionCountKotlinmultiplatformAggSum *int64 `json:"rum_mobile_lite_session_count_kotlinmultiplatform_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on React Native over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountReactnativeAggSum *int64 `json:"rum_mobile_lite_session_count_reactnative_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Roku over all hours within the current month for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountRokuAggSum *int64 `json:"rum_mobile_lite_session_count_roku_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Unity over all hours within the current month for all organizations.
	RumMobileLiteSessionCountUnityAggSum *int64 `json:"rum_mobile_lite_session_count_unity_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM replay sessions on Android over all hours within the current month for all organizations.
	RumMobileReplaySessionCountAndroidAggSum *int64 `json:"rum_mobile_replay_session_count_android_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM replay sessions on iOS over all hours within the current month for all organizations.
	RumMobileReplaySessionCountIosAggSum *int64 `json:"rum_mobile_replay_session_count_ios_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM replay sessions on Kotlin Multiplatform over all hours within the current month for all organizations.
	RumMobileReplaySessionCountKotlinmultiplatformAggSum *int64 `json:"rum_mobile_replay_session_count_kotlinmultiplatform_agg_sum,omitempty"`
	// Shows the sum of all mobile RUM replay sessions on React Native over all hours within the current month for all organizations.
	RumMobileReplaySessionCountReactnativeAggSum *int64 `json:"rum_mobile_replay_session_count_reactnative_agg_sum,omitempty"`
	// Shows the sum of all RUM Session Replay counts over all hours in the current month for all organizations (To be introduced on October 1st, 2024).
	RumReplaySessionCountAggSum *int64 `json:"rum_replay_session_count_agg_sum,omitempty"`
	// Shows the sum of all browser RUM lite sessions over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	RumSessionCountAggSum *int64 `json:"rum_session_count_agg_sum,omitempty"`
	// Sum of all RUM session replay add-on sessions for all hours in the current month for all organizations.
	RumSessionReplayAddOnAggSum *int64 `json:"rum_session_replay_add_on_agg_sum,omitempty"`
	// Shows the sum of RUM sessions (browser and mobile) over all hours in the current month for all organizations.
	RumTotalSessionCountAggSum *int64 `json:"rum_total_session_count_agg_sum,omitempty"`
	// Shows the sum of all browser and mobile RUM units over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	RumUnitsAggSum *int64 `json:"rum_units_agg_sum,omitempty"`
	// Shows the average of all Software Composition Analysis Fargate tasks over all hours in the current months for all organizations.
	ScaFargateCountAvgSum *int64 `json:"sca_fargate_count_avg_sum,omitempty"`
	// Shows the sum of the high-water marks of all Software Composition Analysis Fargate tasks over all hours in the current months for all organizations.
	ScaFargateCountHwmSum *int64 `json:"sca_fargate_count_hwm_sum,omitempty"`
	// Sum of all APM bytes scanned with sensitive data scanner in the current month for all organizations.
	SdsApmScannedBytesSum *int64 `json:"sds_apm_scanned_bytes_sum,omitempty"`
	// Sum of all event stream events bytes scanned with sensitive data scanner in the current month for all organizations.
	SdsEventsScannedBytesSum *int64 `json:"sds_events_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned of logs usage by the Sensitive Data Scanner over all hours in the current month for all organizations.
	SdsLogsScannedBytesSum *int64 `json:"sds_logs_scanned_bytes_sum,omitempty"`
	// Sum of all RUM bytes scanned with sensitive data scanner in the current month for all organizations.
	SdsRumScannedBytesSum *int64 `json:"sds_rum_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned across all usage types by the Sensitive Data Scanner over all hours in the current month for all organizations.
	SdsTotalScannedBytesSum *int64 `json:"sds_total_scanned_bytes_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring for Azure App Service instances in the current month for all organizations.
	ServerlessAppsApmApmAzureAppserviceInstancesAvgSum *int64 `json:"serverless_apps_apm_apm_azure_appservice_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring for Azure Function instances in the current month for all organizations.
	ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum *int64 `json:"serverless_apps_apm_apm_azure_azurefunction_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring for Azure Container App instances in the current month for all organizations.
	ServerlessAppsApmApmAzureContainerappInstancesAvgSum *int64 `json:"serverless_apps_apm_apm_azure_containerapp_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring for Fargate Elastic Container Service tasks in the current month for all organizations.
	ServerlessAppsApmApmFargateEcsTasksAvgSum *int64 `json:"serverless_apps_apm_apm_fargate_ecs_tasks_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring for Google Cloud Platform Cloud Function instances in the current month for all organizations.
	ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum *int64 `json:"serverless_apps_apm_apm_gcp_cloudfunction_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring for Google Cloud Platform Cloud Run instances in the current month for all organizations.
	ServerlessAppsApmApmGcpCloudrunInstancesAvgSum *int64 `json:"serverless_apps_apm_apm_gcp_cloudrun_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring in the current month for all organizations.
	ServerlessAppsApmAvgSum *int64 `json:"serverless_apps_apm_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Azure App Service instances in the current month for all organizations.
	ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum *int64 `json:"serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Azure Function instances in the current month for all organizations.
	ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum *int64 `json:"serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Azure Container App instances in the current month for all organizations.
	ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum *int64 `json:"serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Google Cloud Platform Cloud Function instances in the current month for all organizations.
	ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum *int64 `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Google Cloud Platform Cloud Run instances in the current month for all organizations.
	ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum *int64 `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps with Application Performance Monitoring excluding Fargate in the current month for all organizations.
	ServerlessAppsApmExclFargateAvgSum *int64 `json:"serverless_apps_apm_excl_fargate_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Azure Container App instances in the current month for all organizations.
	ServerlessAppsAzureContainerAppInstancesAvgSum *int64 `json:"serverless_apps_azure_container_app_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Azure in the current month for all organizations.
	ServerlessAppsAzureCountAvgSum *int64 `json:"serverless_apps_azure_count_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Azure Function App instances in the current month for all organizations.
	ServerlessAppsAzureFunctionAppInstancesAvgSum *int64 `json:"serverless_apps_azure_function_app_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Azure Web App instances in the current month for all organizations.
	ServerlessAppsAzureWebAppInstancesAvgSum *int64 `json:"serverless_apps_azure_web_app_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Elastic Container Service in the current month for all organizations.
	ServerlessAppsEcsAvgSum *int64 `json:"serverless_apps_ecs_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Elastic Kubernetes Service in the current month for all organizations.
	ServerlessAppsEksAvgSum *int64 `json:"serverless_apps_eks_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps excluding Fargate in the current month for all organizations.
	ServerlessAppsExclFargateAvgSum *int64 `json:"serverless_apps_excl_fargate_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps excluding Fargate for Azure Container App instances in the current month for all organizations.
	ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum *int64 `json:"serverless_apps_excl_fargate_azure_container_app_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps excluding Fargate for Azure Function App instances in the current month for all organizations.
	ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum *int64 `json:"serverless_apps_excl_fargate_azure_function_app_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps excluding Fargate for Azure Web App instances in the current month for all organizations.
	ServerlessAppsExclFargateAzureWebAppInstancesAvgSum *int64 `json:"serverless_apps_excl_fargate_azure_web_app_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps excluding Fargate for Google Cloud Platform Cloud Functions instances in the current month for all organizations.
	ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum *int64 `json:"serverless_apps_excl_fargate_google_cloud_functions_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps excluding Fargate for Google Cloud Platform Cloud Run instances in the current month for all organizations.
	ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum *int64 `json:"serverless_apps_excl_fargate_google_cloud_run_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Google Cloud Platform Cloud Functions instances in the current month for all organizations.
	ServerlessAppsGoogleCloudFunctionsInstancesAvgSum *int64 `json:"serverless_apps_google_cloud_functions_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Google Cloud Platform Cloud Run instances in the current month for all organizations.
	ServerlessAppsGoogleCloudRunInstancesAvgSum *int64 `json:"serverless_apps_google_cloud_run_instances_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Google Cloud in the current month for all organizations.
	ServerlessAppsGoogleCountAvgSum *int64 `json:"serverless_apps_google_count_avg_sum,omitempty"`
	// Sum of the average number of Serverless Apps for Azure and Google Cloud in the current month for all organizations.
	ServerlessAppsTotalCountAvgSum *int64 `json:"serverless_apps_total_count_avg_sum,omitempty"`
	// Shows the sum of all log events analyzed by Cloud SIEM over all hours in the current month for all organizations.
	SiemAnalyzedLogsAddOnCountAggSum *int64 `json:"siem_analyzed_logs_add_on_count_agg_sum,omitempty"`
	// Shows the first date of usage in the current month for all organizations.
	StartDate *time.Time `json:"start_date,omitempty"`
	// Shows the sum of all Synthetic browser tests over all hours in the current month for all organizations.
	SyntheticsBrowserCheckCallsCountAggSum *int64 `json:"synthetics_browser_check_calls_count_agg_sum,omitempty"`
	// Shows the sum of all Synthetic API tests over all hours in the current month for all organizations.
	SyntheticsCheckCallsCountAggSum *int64 `json:"synthetics_check_calls_count_agg_sum,omitempty"`
	// Shows the sum of Synthetic mobile application tests over all hours in the current month for all organizations.
	SyntheticsMobileTestRunsAggSum *int64 `json:"synthetics_mobile_test_runs_agg_sum,omitempty"`
	// Shows the sum of the high-water marks of used synthetics parallel testing slots over all hours in the current month for all organizations.
	SyntheticsParallelTestingMaxSlotsHwmSum *int64 `json:"synthetics_parallel_testing_max_slots_hwm_sum,omitempty"`
	// Shows the sum of all Indexed Spans indexed over all hours in the current month for all organizations.
	TraceSearchIndexedEventsCountAggSum *int64 `json:"trace_search_indexed_events_count_agg_sum,omitempty"`
	// Shows the sum of all ingested APM span bytes over all hours in the current month for all organizations.
	TwolIngestedEventsBytesAggSum *int64 `json:"twol_ingested_events_bytes_agg_sum,omitempty"`
	// Shows the 99th percentile of all Universal Service Monitoring hosts over all hours in the current month for all organizations.
	UniversalServiceMonitoringHostTop99pSum *int64 `json:"universal_service_monitoring_host_top99p_sum,omitempty"`
	// An array of objects regarding hourly usage.
	Usage []UsageSummaryDate `json:"usage,omitempty"`
	// Shows the 99th percentile of all vSphere hosts over all hours in the current month for all organizations.
	VsphereHostTop99pSum *int64 `json:"vsphere_host_top99p_sum,omitempty"`
	// Shows the 99th percentile of all Application Vulnerability Management hosts over all hours in the current month for all organizations.
	VulnManagementHostCountTop99pSum *int64 `json:"vuln_management_host_count_top99p_sum,omitempty"`
	// Sum of all workflows executed over all hours in the current month for all organizations.
	WorkflowExecutionsUsageAggSum *int64 `json:"workflow_executions_usage_agg_sum,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageSummaryResponse instantiates a new UsageSummaryResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSummaryResponse() *UsageSummaryResponse {
	this := UsageSummaryResponse{}
	return &this
}

// NewUsageSummaryResponseWithDefaults instantiates a new UsageSummaryResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSummaryResponseWithDefaults() *UsageSummaryResponse {
	this := UsageSummaryResponse{}
	return &this
}

// GetAgentHostTop99pSum returns the AgentHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAgentHostTop99pSum() int64 {
	if o == nil || o.AgentHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.AgentHostTop99pSum
}

// GetAgentHostTop99pSumOk returns a tuple with the AgentHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAgentHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.AgentHostTop99pSum == nil {
		return nil, false
	}
	return o.AgentHostTop99pSum, true
}

// HasAgentHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAgentHostTop99pSum() bool {
	return o != nil && o.AgentHostTop99pSum != nil
}

// SetAgentHostTop99pSum gets a reference to the given int64 and assigns it to the AgentHostTop99pSum field.
func (o *UsageSummaryResponse) SetAgentHostTop99pSum(v int64) {
	o.AgentHostTop99pSum = &v
}

// GetApmAzureAppServiceHostTop99pSum returns the ApmAzureAppServiceHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetApmAzureAppServiceHostTop99pSum() int64 {
	if o == nil || o.ApmAzureAppServiceHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.ApmAzureAppServiceHostTop99pSum
}

// GetApmAzureAppServiceHostTop99pSumOk returns a tuple with the ApmAzureAppServiceHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetApmAzureAppServiceHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.ApmAzureAppServiceHostTop99pSum == nil {
		return nil, false
	}
	return o.ApmAzureAppServiceHostTop99pSum, true
}

// HasApmAzureAppServiceHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmAzureAppServiceHostTop99pSum() bool {
	return o != nil && o.ApmAzureAppServiceHostTop99pSum != nil
}

// SetApmAzureAppServiceHostTop99pSum gets a reference to the given int64 and assigns it to the ApmAzureAppServiceHostTop99pSum field.
func (o *UsageSummaryResponse) SetApmAzureAppServiceHostTop99pSum(v int64) {
	o.ApmAzureAppServiceHostTop99pSum = &v
}

// GetApmDevsecopsHostTop99pSum returns the ApmDevsecopsHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetApmDevsecopsHostTop99pSum() int64 {
	if o == nil || o.ApmDevsecopsHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.ApmDevsecopsHostTop99pSum
}

// GetApmDevsecopsHostTop99pSumOk returns a tuple with the ApmDevsecopsHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetApmDevsecopsHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.ApmDevsecopsHostTop99pSum == nil {
		return nil, false
	}
	return o.ApmDevsecopsHostTop99pSum, true
}

// HasApmDevsecopsHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmDevsecopsHostTop99pSum() bool {
	return o != nil && o.ApmDevsecopsHostTop99pSum != nil
}

// SetApmDevsecopsHostTop99pSum gets a reference to the given int64 and assigns it to the ApmDevsecopsHostTop99pSum field.
func (o *UsageSummaryResponse) SetApmDevsecopsHostTop99pSum(v int64) {
	o.ApmDevsecopsHostTop99pSum = &v
}

// GetApmEnterpriseStandaloneHostsTop99pSum returns the ApmEnterpriseStandaloneHostsTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetApmEnterpriseStandaloneHostsTop99pSum() int64 {
	if o == nil || o.ApmEnterpriseStandaloneHostsTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.ApmEnterpriseStandaloneHostsTop99pSum
}

// GetApmEnterpriseStandaloneHostsTop99pSumOk returns a tuple with the ApmEnterpriseStandaloneHostsTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetApmEnterpriseStandaloneHostsTop99pSumOk() (*int64, bool) {
	if o == nil || o.ApmEnterpriseStandaloneHostsTop99pSum == nil {
		return nil, false
	}
	return o.ApmEnterpriseStandaloneHostsTop99pSum, true
}

// HasApmEnterpriseStandaloneHostsTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmEnterpriseStandaloneHostsTop99pSum() bool {
	return o != nil && o.ApmEnterpriseStandaloneHostsTop99pSum != nil
}

// SetApmEnterpriseStandaloneHostsTop99pSum gets a reference to the given int64 and assigns it to the ApmEnterpriseStandaloneHostsTop99pSum field.
func (o *UsageSummaryResponse) SetApmEnterpriseStandaloneHostsTop99pSum(v int64) {
	o.ApmEnterpriseStandaloneHostsTop99pSum = &v
}

// GetApmFargateCountAvgSum returns the ApmFargateCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetApmFargateCountAvgSum() int64 {
	if o == nil || o.ApmFargateCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ApmFargateCountAvgSum
}

// GetApmFargateCountAvgSumOk returns a tuple with the ApmFargateCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetApmFargateCountAvgSumOk() (*int64, bool) {
	if o == nil || o.ApmFargateCountAvgSum == nil {
		return nil, false
	}
	return o.ApmFargateCountAvgSum, true
}

// HasApmFargateCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmFargateCountAvgSum() bool {
	return o != nil && o.ApmFargateCountAvgSum != nil
}

// SetApmFargateCountAvgSum gets a reference to the given int64 and assigns it to the ApmFargateCountAvgSum field.
func (o *UsageSummaryResponse) SetApmFargateCountAvgSum(v int64) {
	o.ApmFargateCountAvgSum = &v
}

// GetApmHostTop99pSum returns the ApmHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetApmHostTop99pSum() int64 {
	if o == nil || o.ApmHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.ApmHostTop99pSum
}

// GetApmHostTop99pSumOk returns a tuple with the ApmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetApmHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.ApmHostTop99pSum == nil {
		return nil, false
	}
	return o.ApmHostTop99pSum, true
}

// HasApmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmHostTop99pSum() bool {
	return o != nil && o.ApmHostTop99pSum != nil
}

// SetApmHostTop99pSum gets a reference to the given int64 and assigns it to the ApmHostTop99pSum field.
func (o *UsageSummaryResponse) SetApmHostTop99pSum(v int64) {
	o.ApmHostTop99pSum = &v
}

// GetApmProStandaloneHostsTop99pSum returns the ApmProStandaloneHostsTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetApmProStandaloneHostsTop99pSum() int64 {
	if o == nil || o.ApmProStandaloneHostsTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.ApmProStandaloneHostsTop99pSum
}

// GetApmProStandaloneHostsTop99pSumOk returns a tuple with the ApmProStandaloneHostsTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetApmProStandaloneHostsTop99pSumOk() (*int64, bool) {
	if o == nil || o.ApmProStandaloneHostsTop99pSum == nil {
		return nil, false
	}
	return o.ApmProStandaloneHostsTop99pSum, true
}

// HasApmProStandaloneHostsTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasApmProStandaloneHostsTop99pSum() bool {
	return o != nil && o.ApmProStandaloneHostsTop99pSum != nil
}

// SetApmProStandaloneHostsTop99pSum gets a reference to the given int64 and assigns it to the ApmProStandaloneHostsTop99pSum field.
func (o *UsageSummaryResponse) SetApmProStandaloneHostsTop99pSum(v int64) {
	o.ApmProStandaloneHostsTop99pSum = &v
}

// GetAppsecFargateCountAvgSum returns the AppsecFargateCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAppsecFargateCountAvgSum() int64 {
	if o == nil || o.AppsecFargateCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.AppsecFargateCountAvgSum
}

// GetAppsecFargateCountAvgSumOk returns a tuple with the AppsecFargateCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAppsecFargateCountAvgSumOk() (*int64, bool) {
	if o == nil || o.AppsecFargateCountAvgSum == nil {
		return nil, false
	}
	return o.AppsecFargateCountAvgSum, true
}

// HasAppsecFargateCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAppsecFargateCountAvgSum() bool {
	return o != nil && o.AppsecFargateCountAvgSum != nil
}

// SetAppsecFargateCountAvgSum gets a reference to the given int64 and assigns it to the AppsecFargateCountAvgSum field.
func (o *UsageSummaryResponse) SetAppsecFargateCountAvgSum(v int64) {
	o.AppsecFargateCountAvgSum = &v
}

// GetAsmServerlessAggSum returns the AsmServerlessAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAsmServerlessAggSum() int64 {
	if o == nil || o.AsmServerlessAggSum == nil {
		var ret int64
		return ret
	}
	return *o.AsmServerlessAggSum
}

// GetAsmServerlessAggSumOk returns a tuple with the AsmServerlessAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAsmServerlessAggSumOk() (*int64, bool) {
	if o == nil || o.AsmServerlessAggSum == nil {
		return nil, false
	}
	return o.AsmServerlessAggSum, true
}

// HasAsmServerlessAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAsmServerlessAggSum() bool {
	return o != nil && o.AsmServerlessAggSum != nil
}

// SetAsmServerlessAggSum gets a reference to the given int64 and assigns it to the AsmServerlessAggSum field.
func (o *UsageSummaryResponse) SetAsmServerlessAggSum(v int64) {
	o.AsmServerlessAggSum = &v
}

// GetAuditLogsLinesIndexedAggSum returns the AuditLogsLinesIndexedAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetAuditLogsLinesIndexedAggSum() int64 {
	if o == nil || o.AuditLogsLinesIndexedAggSum == nil {
		var ret int64
		return ret
	}
	return *o.AuditLogsLinesIndexedAggSum
}

// GetAuditLogsLinesIndexedAggSumOk returns a tuple with the AuditLogsLinesIndexedAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetAuditLogsLinesIndexedAggSumOk() (*int64, bool) {
	if o == nil || o.AuditLogsLinesIndexedAggSum == nil {
		return nil, false
	}
	return o.AuditLogsLinesIndexedAggSum, true
}

// HasAuditLogsLinesIndexedAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAuditLogsLinesIndexedAggSum() bool {
	return o != nil && o.AuditLogsLinesIndexedAggSum != nil
}

// SetAuditLogsLinesIndexedAggSum gets a reference to the given int64 and assigns it to the AuditLogsLinesIndexedAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetAuditLogsLinesIndexedAggSum(v int64) {
	o.AuditLogsLinesIndexedAggSum = &v
}

// GetAuditTrailEnabledHwmSum returns the AuditTrailEnabledHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAuditTrailEnabledHwmSum() int64 {
	if o == nil || o.AuditTrailEnabledHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.AuditTrailEnabledHwmSum
}

// GetAuditTrailEnabledHwmSumOk returns a tuple with the AuditTrailEnabledHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAuditTrailEnabledHwmSumOk() (*int64, bool) {
	if o == nil || o.AuditTrailEnabledHwmSum == nil {
		return nil, false
	}
	return o.AuditTrailEnabledHwmSum, true
}

// HasAuditTrailEnabledHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAuditTrailEnabledHwmSum() bool {
	return o != nil && o.AuditTrailEnabledHwmSum != nil
}

// SetAuditTrailEnabledHwmSum gets a reference to the given int64 and assigns it to the AuditTrailEnabledHwmSum field.
func (o *UsageSummaryResponse) SetAuditTrailEnabledHwmSum(v int64) {
	o.AuditTrailEnabledHwmSum = &v
}

// GetAvgProfiledFargateTasksSum returns the AvgProfiledFargateTasksSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAvgProfiledFargateTasksSum() int64 {
	if o == nil || o.AvgProfiledFargateTasksSum == nil {
		var ret int64
		return ret
	}
	return *o.AvgProfiledFargateTasksSum
}

// GetAvgProfiledFargateTasksSumOk returns a tuple with the AvgProfiledFargateTasksSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAvgProfiledFargateTasksSumOk() (*int64, bool) {
	if o == nil || o.AvgProfiledFargateTasksSum == nil {
		return nil, false
	}
	return o.AvgProfiledFargateTasksSum, true
}

// HasAvgProfiledFargateTasksSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAvgProfiledFargateTasksSum() bool {
	return o != nil && o.AvgProfiledFargateTasksSum != nil
}

// SetAvgProfiledFargateTasksSum gets a reference to the given int64 and assigns it to the AvgProfiledFargateTasksSum field.
func (o *UsageSummaryResponse) SetAvgProfiledFargateTasksSum(v int64) {
	o.AvgProfiledFargateTasksSum = &v
}

// GetAwsHostTop99pSum returns the AwsHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAwsHostTop99pSum() int64 {
	if o == nil || o.AwsHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.AwsHostTop99pSum
}

// GetAwsHostTop99pSumOk returns a tuple with the AwsHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAwsHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.AwsHostTop99pSum == nil {
		return nil, false
	}
	return o.AwsHostTop99pSum, true
}

// HasAwsHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAwsHostTop99pSum() bool {
	return o != nil && o.AwsHostTop99pSum != nil
}

// SetAwsHostTop99pSum gets a reference to the given int64 and assigns it to the AwsHostTop99pSum field.
func (o *UsageSummaryResponse) SetAwsHostTop99pSum(v int64) {
	o.AwsHostTop99pSum = &v
}

// GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAwsLambdaFuncCount() int64 {
	if o == nil || o.AwsLambdaFuncCount == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaFuncCount
}

// GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAwsLambdaFuncCountOk() (*int64, bool) {
	if o == nil || o.AwsLambdaFuncCount == nil {
		return nil, false
	}
	return o.AwsLambdaFuncCount, true
}

// HasAwsLambdaFuncCount returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAwsLambdaFuncCount() bool {
	return o != nil && o.AwsLambdaFuncCount != nil
}

// SetAwsLambdaFuncCount gets a reference to the given int64 and assigns it to the AwsLambdaFuncCount field.
func (o *UsageSummaryResponse) SetAwsLambdaFuncCount(v int64) {
	o.AwsLambdaFuncCount = &v
}

// GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAwsLambdaInvocationsSum() int64 {
	if o == nil || o.AwsLambdaInvocationsSum == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaInvocationsSum
}

// GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAwsLambdaInvocationsSumOk() (*int64, bool) {
	if o == nil || o.AwsLambdaInvocationsSum == nil {
		return nil, false
	}
	return o.AwsLambdaInvocationsSum, true
}

// HasAwsLambdaInvocationsSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAwsLambdaInvocationsSum() bool {
	return o != nil && o.AwsLambdaInvocationsSum != nil
}

// SetAwsLambdaInvocationsSum gets a reference to the given int64 and assigns it to the AwsLambdaInvocationsSum field.
func (o *UsageSummaryResponse) SetAwsLambdaInvocationsSum(v int64) {
	o.AwsLambdaInvocationsSum = &v
}

// GetAzureAppServiceTop99pSum returns the AzureAppServiceTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAzureAppServiceTop99pSum() int64 {
	if o == nil || o.AzureAppServiceTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.AzureAppServiceTop99pSum
}

// GetAzureAppServiceTop99pSumOk returns a tuple with the AzureAppServiceTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAzureAppServiceTop99pSumOk() (*int64, bool) {
	if o == nil || o.AzureAppServiceTop99pSum == nil {
		return nil, false
	}
	return o.AzureAppServiceTop99pSum, true
}

// HasAzureAppServiceTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAzureAppServiceTop99pSum() bool {
	return o != nil && o.AzureAppServiceTop99pSum != nil
}

// SetAzureAppServiceTop99pSum gets a reference to the given int64 and assigns it to the AzureAppServiceTop99pSum field.
func (o *UsageSummaryResponse) SetAzureAppServiceTop99pSum(v int64) {
	o.AzureAppServiceTop99pSum = &v
}

// GetAzureHostTop99pSum returns the AzureHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetAzureHostTop99pSum() int64 {
	if o == nil || o.AzureHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.AzureHostTop99pSum
}

// GetAzureHostTop99pSumOk returns a tuple with the AzureHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetAzureHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.AzureHostTop99pSum == nil {
		return nil, false
	}
	return o.AzureHostTop99pSum, true
}

// HasAzureHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasAzureHostTop99pSum() bool {
	return o != nil && o.AzureHostTop99pSum != nil
}

// SetAzureHostTop99pSum gets a reference to the given int64 and assigns it to the AzureHostTop99pSum field.
func (o *UsageSummaryResponse) SetAzureHostTop99pSum(v int64) {
	o.AzureHostTop99pSum = &v
}

// GetBillableIngestedBytesAggSum returns the BillableIngestedBytesAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetBillableIngestedBytesAggSum() int64 {
	if o == nil || o.BillableIngestedBytesAggSum == nil {
		var ret int64
		return ret
	}
	return *o.BillableIngestedBytesAggSum
}

// GetBillableIngestedBytesAggSumOk returns a tuple with the BillableIngestedBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetBillableIngestedBytesAggSumOk() (*int64, bool) {
	if o == nil || o.BillableIngestedBytesAggSum == nil {
		return nil, false
	}
	return o.BillableIngestedBytesAggSum, true
}

// HasBillableIngestedBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBillableIngestedBytesAggSum() bool {
	return o != nil && o.BillableIngestedBytesAggSum != nil
}

// SetBillableIngestedBytesAggSum gets a reference to the given int64 and assigns it to the BillableIngestedBytesAggSum field.
func (o *UsageSummaryResponse) SetBillableIngestedBytesAggSum(v int64) {
	o.BillableIngestedBytesAggSum = &v
}

// GetBitsAiInvestigationsAggSum returns the BitsAiInvestigationsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetBitsAiInvestigationsAggSum() int64 {
	if o == nil || o.BitsAiInvestigationsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.BitsAiInvestigationsAggSum
}

// GetBitsAiInvestigationsAggSumOk returns a tuple with the BitsAiInvestigationsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetBitsAiInvestigationsAggSumOk() (*int64, bool) {
	if o == nil || o.BitsAiInvestigationsAggSum == nil {
		return nil, false
	}
	return o.BitsAiInvestigationsAggSum, true
}

// HasBitsAiInvestigationsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBitsAiInvestigationsAggSum() bool {
	return o != nil && o.BitsAiInvestigationsAggSum != nil
}

// SetBitsAiInvestigationsAggSum gets a reference to the given int64 and assigns it to the BitsAiInvestigationsAggSum field.
func (o *UsageSummaryResponse) SetBitsAiInvestigationsAggSum(v int64) {
	o.BitsAiInvestigationsAggSum = &v
}

// GetBrowserRumLiteSessionCountAggSum returns the BrowserRumLiteSessionCountAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetBrowserRumLiteSessionCountAggSum() int64 {
	if o == nil || o.BrowserRumLiteSessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumLiteSessionCountAggSum
}

// GetBrowserRumLiteSessionCountAggSumOk returns a tuple with the BrowserRumLiteSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetBrowserRumLiteSessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.BrowserRumLiteSessionCountAggSum == nil {
		return nil, false
	}
	return o.BrowserRumLiteSessionCountAggSum, true
}

// HasBrowserRumLiteSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBrowserRumLiteSessionCountAggSum() bool {
	return o != nil && o.BrowserRumLiteSessionCountAggSum != nil
}

// SetBrowserRumLiteSessionCountAggSum gets a reference to the given int64 and assigns it to the BrowserRumLiteSessionCountAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetBrowserRumLiteSessionCountAggSum(v int64) {
	o.BrowserRumLiteSessionCountAggSum = &v
}

// GetBrowserRumReplaySessionCountAggSum returns the BrowserRumReplaySessionCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetBrowserRumReplaySessionCountAggSum() int64 {
	if o == nil || o.BrowserRumReplaySessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumReplaySessionCountAggSum
}

// GetBrowserRumReplaySessionCountAggSumOk returns a tuple with the BrowserRumReplaySessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetBrowserRumReplaySessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.BrowserRumReplaySessionCountAggSum == nil {
		return nil, false
	}
	return o.BrowserRumReplaySessionCountAggSum, true
}

// HasBrowserRumReplaySessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBrowserRumReplaySessionCountAggSum() bool {
	return o != nil && o.BrowserRumReplaySessionCountAggSum != nil
}

// SetBrowserRumReplaySessionCountAggSum gets a reference to the given int64 and assigns it to the BrowserRumReplaySessionCountAggSum field.
func (o *UsageSummaryResponse) SetBrowserRumReplaySessionCountAggSum(v int64) {
	o.BrowserRumReplaySessionCountAggSum = &v
}

// GetBrowserRumUnitsAggSum returns the BrowserRumUnitsAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetBrowserRumUnitsAggSum() int64 {
	if o == nil || o.BrowserRumUnitsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumUnitsAggSum
}

// GetBrowserRumUnitsAggSumOk returns a tuple with the BrowserRumUnitsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetBrowserRumUnitsAggSumOk() (*int64, bool) {
	if o == nil || o.BrowserRumUnitsAggSum == nil {
		return nil, false
	}
	return o.BrowserRumUnitsAggSum, true
}

// HasBrowserRumUnitsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasBrowserRumUnitsAggSum() bool {
	return o != nil && o.BrowserRumUnitsAggSum != nil
}

// SetBrowserRumUnitsAggSum gets a reference to the given int64 and assigns it to the BrowserRumUnitsAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetBrowserRumUnitsAggSum(v int64) {
	o.BrowserRumUnitsAggSum = &v
}

// GetCiPipelineIndexedSpansAggSum returns the CiPipelineIndexedSpansAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCiPipelineIndexedSpansAggSum() int64 {
	if o == nil || o.CiPipelineIndexedSpansAggSum == nil {
		var ret int64
		return ret
	}
	return *o.CiPipelineIndexedSpansAggSum
}

// GetCiPipelineIndexedSpansAggSumOk returns a tuple with the CiPipelineIndexedSpansAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCiPipelineIndexedSpansAggSumOk() (*int64, bool) {
	if o == nil || o.CiPipelineIndexedSpansAggSum == nil {
		return nil, false
	}
	return o.CiPipelineIndexedSpansAggSum, true
}

// HasCiPipelineIndexedSpansAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCiPipelineIndexedSpansAggSum() bool {
	return o != nil && o.CiPipelineIndexedSpansAggSum != nil
}

// SetCiPipelineIndexedSpansAggSum gets a reference to the given int64 and assigns it to the CiPipelineIndexedSpansAggSum field.
func (o *UsageSummaryResponse) SetCiPipelineIndexedSpansAggSum(v int64) {
	o.CiPipelineIndexedSpansAggSum = &v
}

// GetCiTestIndexedSpansAggSum returns the CiTestIndexedSpansAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCiTestIndexedSpansAggSum() int64 {
	if o == nil || o.CiTestIndexedSpansAggSum == nil {
		var ret int64
		return ret
	}
	return *o.CiTestIndexedSpansAggSum
}

// GetCiTestIndexedSpansAggSumOk returns a tuple with the CiTestIndexedSpansAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCiTestIndexedSpansAggSumOk() (*int64, bool) {
	if o == nil || o.CiTestIndexedSpansAggSum == nil {
		return nil, false
	}
	return o.CiTestIndexedSpansAggSum, true
}

// HasCiTestIndexedSpansAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCiTestIndexedSpansAggSum() bool {
	return o != nil && o.CiTestIndexedSpansAggSum != nil
}

// SetCiTestIndexedSpansAggSum gets a reference to the given int64 and assigns it to the CiTestIndexedSpansAggSum field.
func (o *UsageSummaryResponse) SetCiTestIndexedSpansAggSum(v int64) {
	o.CiTestIndexedSpansAggSum = &v
}

// GetCiVisibilityItrCommittersHwmSum returns the CiVisibilityItrCommittersHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCiVisibilityItrCommittersHwmSum() int64 {
	if o == nil || o.CiVisibilityItrCommittersHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityItrCommittersHwmSum
}

// GetCiVisibilityItrCommittersHwmSumOk returns a tuple with the CiVisibilityItrCommittersHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCiVisibilityItrCommittersHwmSumOk() (*int64, bool) {
	if o == nil || o.CiVisibilityItrCommittersHwmSum == nil {
		return nil, false
	}
	return o.CiVisibilityItrCommittersHwmSum, true
}

// HasCiVisibilityItrCommittersHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCiVisibilityItrCommittersHwmSum() bool {
	return o != nil && o.CiVisibilityItrCommittersHwmSum != nil
}

// SetCiVisibilityItrCommittersHwmSum gets a reference to the given int64 and assigns it to the CiVisibilityItrCommittersHwmSum field.
func (o *UsageSummaryResponse) SetCiVisibilityItrCommittersHwmSum(v int64) {
	o.CiVisibilityItrCommittersHwmSum = &v
}

// GetCiVisibilityPipelineCommittersHwmSum returns the CiVisibilityPipelineCommittersHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCiVisibilityPipelineCommittersHwmSum() int64 {
	if o == nil || o.CiVisibilityPipelineCommittersHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityPipelineCommittersHwmSum
}

// GetCiVisibilityPipelineCommittersHwmSumOk returns a tuple with the CiVisibilityPipelineCommittersHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCiVisibilityPipelineCommittersHwmSumOk() (*int64, bool) {
	if o == nil || o.CiVisibilityPipelineCommittersHwmSum == nil {
		return nil, false
	}
	return o.CiVisibilityPipelineCommittersHwmSum, true
}

// HasCiVisibilityPipelineCommittersHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCiVisibilityPipelineCommittersHwmSum() bool {
	return o != nil && o.CiVisibilityPipelineCommittersHwmSum != nil
}

// SetCiVisibilityPipelineCommittersHwmSum gets a reference to the given int64 and assigns it to the CiVisibilityPipelineCommittersHwmSum field.
func (o *UsageSummaryResponse) SetCiVisibilityPipelineCommittersHwmSum(v int64) {
	o.CiVisibilityPipelineCommittersHwmSum = &v
}

// GetCiVisibilityTestCommittersHwmSum returns the CiVisibilityTestCommittersHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCiVisibilityTestCommittersHwmSum() int64 {
	if o == nil || o.CiVisibilityTestCommittersHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityTestCommittersHwmSum
}

// GetCiVisibilityTestCommittersHwmSumOk returns a tuple with the CiVisibilityTestCommittersHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCiVisibilityTestCommittersHwmSumOk() (*int64, bool) {
	if o == nil || o.CiVisibilityTestCommittersHwmSum == nil {
		return nil, false
	}
	return o.CiVisibilityTestCommittersHwmSum, true
}

// HasCiVisibilityTestCommittersHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCiVisibilityTestCommittersHwmSum() bool {
	return o != nil && o.CiVisibilityTestCommittersHwmSum != nil
}

// SetCiVisibilityTestCommittersHwmSum gets a reference to the given int64 and assigns it to the CiVisibilityTestCommittersHwmSum field.
func (o *UsageSummaryResponse) SetCiVisibilityTestCommittersHwmSum(v int64) {
	o.CiVisibilityTestCommittersHwmSum = &v
}

// GetCloudCostManagementAwsHostCountAvgSum returns the CloudCostManagementAwsHostCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCloudCostManagementAwsHostCountAvgSum() int64 {
	if o == nil || o.CloudCostManagementAwsHostCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementAwsHostCountAvgSum
}

// GetCloudCostManagementAwsHostCountAvgSumOk returns a tuple with the CloudCostManagementAwsHostCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCloudCostManagementAwsHostCountAvgSumOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementAwsHostCountAvgSum == nil {
		return nil, false
	}
	return o.CloudCostManagementAwsHostCountAvgSum, true
}

// HasCloudCostManagementAwsHostCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCloudCostManagementAwsHostCountAvgSum() bool {
	return o != nil && o.CloudCostManagementAwsHostCountAvgSum != nil
}

// SetCloudCostManagementAwsHostCountAvgSum gets a reference to the given int64 and assigns it to the CloudCostManagementAwsHostCountAvgSum field.
func (o *UsageSummaryResponse) SetCloudCostManagementAwsHostCountAvgSum(v int64) {
	o.CloudCostManagementAwsHostCountAvgSum = &v
}

// GetCloudCostManagementAzureHostCountAvgSum returns the CloudCostManagementAzureHostCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCloudCostManagementAzureHostCountAvgSum() int64 {
	if o == nil || o.CloudCostManagementAzureHostCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementAzureHostCountAvgSum
}

// GetCloudCostManagementAzureHostCountAvgSumOk returns a tuple with the CloudCostManagementAzureHostCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCloudCostManagementAzureHostCountAvgSumOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementAzureHostCountAvgSum == nil {
		return nil, false
	}
	return o.CloudCostManagementAzureHostCountAvgSum, true
}

// HasCloudCostManagementAzureHostCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCloudCostManagementAzureHostCountAvgSum() bool {
	return o != nil && o.CloudCostManagementAzureHostCountAvgSum != nil
}

// SetCloudCostManagementAzureHostCountAvgSum gets a reference to the given int64 and assigns it to the CloudCostManagementAzureHostCountAvgSum field.
func (o *UsageSummaryResponse) SetCloudCostManagementAzureHostCountAvgSum(v int64) {
	o.CloudCostManagementAzureHostCountAvgSum = &v
}

// GetCloudCostManagementGcpHostCountAvgSum returns the CloudCostManagementGcpHostCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCloudCostManagementGcpHostCountAvgSum() int64 {
	if o == nil || o.CloudCostManagementGcpHostCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementGcpHostCountAvgSum
}

// GetCloudCostManagementGcpHostCountAvgSumOk returns a tuple with the CloudCostManagementGcpHostCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCloudCostManagementGcpHostCountAvgSumOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementGcpHostCountAvgSum == nil {
		return nil, false
	}
	return o.CloudCostManagementGcpHostCountAvgSum, true
}

// HasCloudCostManagementGcpHostCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCloudCostManagementGcpHostCountAvgSum() bool {
	return o != nil && o.CloudCostManagementGcpHostCountAvgSum != nil
}

// SetCloudCostManagementGcpHostCountAvgSum gets a reference to the given int64 and assigns it to the CloudCostManagementGcpHostCountAvgSum field.
func (o *UsageSummaryResponse) SetCloudCostManagementGcpHostCountAvgSum(v int64) {
	o.CloudCostManagementGcpHostCountAvgSum = &v
}

// GetCloudCostManagementHostCountAvgSum returns the CloudCostManagementHostCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCloudCostManagementHostCountAvgSum() int64 {
	if o == nil || o.CloudCostManagementHostCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementHostCountAvgSum
}

// GetCloudCostManagementHostCountAvgSumOk returns a tuple with the CloudCostManagementHostCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCloudCostManagementHostCountAvgSumOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementHostCountAvgSum == nil {
		return nil, false
	}
	return o.CloudCostManagementHostCountAvgSum, true
}

// HasCloudCostManagementHostCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCloudCostManagementHostCountAvgSum() bool {
	return o != nil && o.CloudCostManagementHostCountAvgSum != nil
}

// SetCloudCostManagementHostCountAvgSum gets a reference to the given int64 and assigns it to the CloudCostManagementHostCountAvgSum field.
func (o *UsageSummaryResponse) SetCloudCostManagementHostCountAvgSum(v int64) {
	o.CloudCostManagementHostCountAvgSum = &v
}

// GetCloudCostManagementOciHostCountAvgSum returns the CloudCostManagementOciHostCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCloudCostManagementOciHostCountAvgSum() int64 {
	if o == nil || o.CloudCostManagementOciHostCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementOciHostCountAvgSum
}

// GetCloudCostManagementOciHostCountAvgSumOk returns a tuple with the CloudCostManagementOciHostCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCloudCostManagementOciHostCountAvgSumOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementOciHostCountAvgSum == nil {
		return nil, false
	}
	return o.CloudCostManagementOciHostCountAvgSum, true
}

// HasCloudCostManagementOciHostCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCloudCostManagementOciHostCountAvgSum() bool {
	return o != nil && o.CloudCostManagementOciHostCountAvgSum != nil
}

// SetCloudCostManagementOciHostCountAvgSum gets a reference to the given int64 and assigns it to the CloudCostManagementOciHostCountAvgSum field.
func (o *UsageSummaryResponse) SetCloudCostManagementOciHostCountAvgSum(v int64) {
	o.CloudCostManagementOciHostCountAvgSum = &v
}

// GetCloudSiemEventsAggSum returns the CloudSiemEventsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCloudSiemEventsAggSum() int64 {
	if o == nil || o.CloudSiemEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.CloudSiemEventsAggSum
}

// GetCloudSiemEventsAggSumOk returns a tuple with the CloudSiemEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCloudSiemEventsAggSumOk() (*int64, bool) {
	if o == nil || o.CloudSiemEventsAggSum == nil {
		return nil, false
	}
	return o.CloudSiemEventsAggSum, true
}

// HasCloudSiemEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCloudSiemEventsAggSum() bool {
	return o != nil && o.CloudSiemEventsAggSum != nil
}

// SetCloudSiemEventsAggSum gets a reference to the given int64 and assigns it to the CloudSiemEventsAggSum field.
func (o *UsageSummaryResponse) SetCloudSiemEventsAggSum(v int64) {
	o.CloudSiemEventsAggSum = &v
}

// GetCodeAnalysisSaCommittersHwmSum returns the CodeAnalysisSaCommittersHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCodeAnalysisSaCommittersHwmSum() int64 {
	if o == nil || o.CodeAnalysisSaCommittersHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.CodeAnalysisSaCommittersHwmSum
}

// GetCodeAnalysisSaCommittersHwmSumOk returns a tuple with the CodeAnalysisSaCommittersHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCodeAnalysisSaCommittersHwmSumOk() (*int64, bool) {
	if o == nil || o.CodeAnalysisSaCommittersHwmSum == nil {
		return nil, false
	}
	return o.CodeAnalysisSaCommittersHwmSum, true
}

// HasCodeAnalysisSaCommittersHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCodeAnalysisSaCommittersHwmSum() bool {
	return o != nil && o.CodeAnalysisSaCommittersHwmSum != nil
}

// SetCodeAnalysisSaCommittersHwmSum gets a reference to the given int64 and assigns it to the CodeAnalysisSaCommittersHwmSum field.
func (o *UsageSummaryResponse) SetCodeAnalysisSaCommittersHwmSum(v int64) {
	o.CodeAnalysisSaCommittersHwmSum = &v
}

// GetCodeAnalysisScaCommittersHwmSum returns the CodeAnalysisScaCommittersHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCodeAnalysisScaCommittersHwmSum() int64 {
	if o == nil || o.CodeAnalysisScaCommittersHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.CodeAnalysisScaCommittersHwmSum
}

// GetCodeAnalysisScaCommittersHwmSumOk returns a tuple with the CodeAnalysisScaCommittersHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCodeAnalysisScaCommittersHwmSumOk() (*int64, bool) {
	if o == nil || o.CodeAnalysisScaCommittersHwmSum == nil {
		return nil, false
	}
	return o.CodeAnalysisScaCommittersHwmSum, true
}

// HasCodeAnalysisScaCommittersHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCodeAnalysisScaCommittersHwmSum() bool {
	return o != nil && o.CodeAnalysisScaCommittersHwmSum != nil
}

// SetCodeAnalysisScaCommittersHwmSum gets a reference to the given int64 and assigns it to the CodeAnalysisScaCommittersHwmSum field.
func (o *UsageSummaryResponse) SetCodeAnalysisScaCommittersHwmSum(v int64) {
	o.CodeAnalysisScaCommittersHwmSum = &v
}

// GetCodeSecurityHostTop99pSum returns the CodeSecurityHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCodeSecurityHostTop99pSum() int64 {
	if o == nil || o.CodeSecurityHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CodeSecurityHostTop99pSum
}

// GetCodeSecurityHostTop99pSumOk returns a tuple with the CodeSecurityHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCodeSecurityHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.CodeSecurityHostTop99pSum == nil {
		return nil, false
	}
	return o.CodeSecurityHostTop99pSum, true
}

// HasCodeSecurityHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCodeSecurityHostTop99pSum() bool {
	return o != nil && o.CodeSecurityHostTop99pSum != nil
}

// SetCodeSecurityHostTop99pSum gets a reference to the given int64 and assigns it to the CodeSecurityHostTop99pSum field.
func (o *UsageSummaryResponse) SetCodeSecurityHostTop99pSum(v int64) {
	o.CodeSecurityHostTop99pSum = &v
}

// GetContainerAvgSum returns the ContainerAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetContainerAvgSum() int64 {
	if o == nil || o.ContainerAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ContainerAvgSum
}

// GetContainerAvgSumOk returns a tuple with the ContainerAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetContainerAvgSumOk() (*int64, bool) {
	if o == nil || o.ContainerAvgSum == nil {
		return nil, false
	}
	return o.ContainerAvgSum, true
}

// HasContainerAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasContainerAvgSum() bool {
	return o != nil && o.ContainerAvgSum != nil
}

// SetContainerAvgSum gets a reference to the given int64 and assigns it to the ContainerAvgSum field.
func (o *UsageSummaryResponse) SetContainerAvgSum(v int64) {
	o.ContainerAvgSum = &v
}

// GetContainerExclAgentAvgSum returns the ContainerExclAgentAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetContainerExclAgentAvgSum() int64 {
	if o == nil || o.ContainerExclAgentAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ContainerExclAgentAvgSum
}

// GetContainerExclAgentAvgSumOk returns a tuple with the ContainerExclAgentAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetContainerExclAgentAvgSumOk() (*int64, bool) {
	if o == nil || o.ContainerExclAgentAvgSum == nil {
		return nil, false
	}
	return o.ContainerExclAgentAvgSum, true
}

// HasContainerExclAgentAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasContainerExclAgentAvgSum() bool {
	return o != nil && o.ContainerExclAgentAvgSum != nil
}

// SetContainerExclAgentAvgSum gets a reference to the given int64 and assigns it to the ContainerExclAgentAvgSum field.
func (o *UsageSummaryResponse) SetContainerExclAgentAvgSum(v int64) {
	o.ContainerExclAgentAvgSum = &v
}

// GetContainerHwmSum returns the ContainerHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetContainerHwmSum() int64 {
	if o == nil || o.ContainerHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.ContainerHwmSum
}

// GetContainerHwmSumOk returns a tuple with the ContainerHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetContainerHwmSumOk() (*int64, bool) {
	if o == nil || o.ContainerHwmSum == nil {
		return nil, false
	}
	return o.ContainerHwmSum, true
}

// HasContainerHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasContainerHwmSum() bool {
	return o != nil && o.ContainerHwmSum != nil
}

// SetContainerHwmSum gets a reference to the given int64 and assigns it to the ContainerHwmSum field.
func (o *UsageSummaryResponse) SetContainerHwmSum(v int64) {
	o.ContainerHwmSum = &v
}

// GetCsmContainerEnterpriseComplianceCountAggSum returns the CsmContainerEnterpriseComplianceCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmContainerEnterpriseComplianceCountAggSum() int64 {
	if o == nil || o.CsmContainerEnterpriseComplianceCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmContainerEnterpriseComplianceCountAggSum
}

// GetCsmContainerEnterpriseComplianceCountAggSumOk returns a tuple with the CsmContainerEnterpriseComplianceCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmContainerEnterpriseComplianceCountAggSumOk() (*int64, bool) {
	if o == nil || o.CsmContainerEnterpriseComplianceCountAggSum == nil {
		return nil, false
	}
	return o.CsmContainerEnterpriseComplianceCountAggSum, true
}

// HasCsmContainerEnterpriseComplianceCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmContainerEnterpriseComplianceCountAggSum() bool {
	return o != nil && o.CsmContainerEnterpriseComplianceCountAggSum != nil
}

// SetCsmContainerEnterpriseComplianceCountAggSum gets a reference to the given int64 and assigns it to the CsmContainerEnterpriseComplianceCountAggSum field.
func (o *UsageSummaryResponse) SetCsmContainerEnterpriseComplianceCountAggSum(v int64) {
	o.CsmContainerEnterpriseComplianceCountAggSum = &v
}

// GetCsmContainerEnterpriseCwsCountAggSum returns the CsmContainerEnterpriseCwsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmContainerEnterpriseCwsCountAggSum() int64 {
	if o == nil || o.CsmContainerEnterpriseCwsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmContainerEnterpriseCwsCountAggSum
}

// GetCsmContainerEnterpriseCwsCountAggSumOk returns a tuple with the CsmContainerEnterpriseCwsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmContainerEnterpriseCwsCountAggSumOk() (*int64, bool) {
	if o == nil || o.CsmContainerEnterpriseCwsCountAggSum == nil {
		return nil, false
	}
	return o.CsmContainerEnterpriseCwsCountAggSum, true
}

// HasCsmContainerEnterpriseCwsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmContainerEnterpriseCwsCountAggSum() bool {
	return o != nil && o.CsmContainerEnterpriseCwsCountAggSum != nil
}

// SetCsmContainerEnterpriseCwsCountAggSum gets a reference to the given int64 and assigns it to the CsmContainerEnterpriseCwsCountAggSum field.
func (o *UsageSummaryResponse) SetCsmContainerEnterpriseCwsCountAggSum(v int64) {
	o.CsmContainerEnterpriseCwsCountAggSum = &v
}

// GetCsmContainerEnterpriseTotalCountAggSum returns the CsmContainerEnterpriseTotalCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmContainerEnterpriseTotalCountAggSum() int64 {
	if o == nil || o.CsmContainerEnterpriseTotalCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmContainerEnterpriseTotalCountAggSum
}

// GetCsmContainerEnterpriseTotalCountAggSumOk returns a tuple with the CsmContainerEnterpriseTotalCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmContainerEnterpriseTotalCountAggSumOk() (*int64, bool) {
	if o == nil || o.CsmContainerEnterpriseTotalCountAggSum == nil {
		return nil, false
	}
	return o.CsmContainerEnterpriseTotalCountAggSum, true
}

// HasCsmContainerEnterpriseTotalCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmContainerEnterpriseTotalCountAggSum() bool {
	return o != nil && o.CsmContainerEnterpriseTotalCountAggSum != nil
}

// SetCsmContainerEnterpriseTotalCountAggSum gets a reference to the given int64 and assigns it to the CsmContainerEnterpriseTotalCountAggSum field.
func (o *UsageSummaryResponse) SetCsmContainerEnterpriseTotalCountAggSum(v int64) {
	o.CsmContainerEnterpriseTotalCountAggSum = &v
}

// GetCsmHostEnterpriseAasHostCountTop99pSum returns the CsmHostEnterpriseAasHostCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseAasHostCountTop99pSum() int64 {
	if o == nil || o.CsmHostEnterpriseAasHostCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseAasHostCountTop99pSum
}

// GetCsmHostEnterpriseAasHostCountTop99pSumOk returns a tuple with the CsmHostEnterpriseAasHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseAasHostCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseAasHostCountTop99pSum == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseAasHostCountTop99pSum, true
}

// HasCsmHostEnterpriseAasHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmHostEnterpriseAasHostCountTop99pSum() bool {
	return o != nil && o.CsmHostEnterpriseAasHostCountTop99pSum != nil
}

// SetCsmHostEnterpriseAasHostCountTop99pSum gets a reference to the given int64 and assigns it to the CsmHostEnterpriseAasHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetCsmHostEnterpriseAasHostCountTop99pSum(v int64) {
	o.CsmHostEnterpriseAasHostCountTop99pSum = &v
}

// GetCsmHostEnterpriseAwsHostCountTop99pSum returns the CsmHostEnterpriseAwsHostCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseAwsHostCountTop99pSum() int64 {
	if o == nil || o.CsmHostEnterpriseAwsHostCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseAwsHostCountTop99pSum
}

// GetCsmHostEnterpriseAwsHostCountTop99pSumOk returns a tuple with the CsmHostEnterpriseAwsHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseAwsHostCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseAwsHostCountTop99pSum == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseAwsHostCountTop99pSum, true
}

// HasCsmHostEnterpriseAwsHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmHostEnterpriseAwsHostCountTop99pSum() bool {
	return o != nil && o.CsmHostEnterpriseAwsHostCountTop99pSum != nil
}

// SetCsmHostEnterpriseAwsHostCountTop99pSum gets a reference to the given int64 and assigns it to the CsmHostEnterpriseAwsHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetCsmHostEnterpriseAwsHostCountTop99pSum(v int64) {
	o.CsmHostEnterpriseAwsHostCountTop99pSum = &v
}

// GetCsmHostEnterpriseAzureHostCountTop99pSum returns the CsmHostEnterpriseAzureHostCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseAzureHostCountTop99pSum() int64 {
	if o == nil || o.CsmHostEnterpriseAzureHostCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseAzureHostCountTop99pSum
}

// GetCsmHostEnterpriseAzureHostCountTop99pSumOk returns a tuple with the CsmHostEnterpriseAzureHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseAzureHostCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseAzureHostCountTop99pSum == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseAzureHostCountTop99pSum, true
}

// HasCsmHostEnterpriseAzureHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmHostEnterpriseAzureHostCountTop99pSum() bool {
	return o != nil && o.CsmHostEnterpriseAzureHostCountTop99pSum != nil
}

// SetCsmHostEnterpriseAzureHostCountTop99pSum gets a reference to the given int64 and assigns it to the CsmHostEnterpriseAzureHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetCsmHostEnterpriseAzureHostCountTop99pSum(v int64) {
	o.CsmHostEnterpriseAzureHostCountTop99pSum = &v
}

// GetCsmHostEnterpriseComplianceHostCountTop99pSum returns the CsmHostEnterpriseComplianceHostCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseComplianceHostCountTop99pSum() int64 {
	if o == nil || o.CsmHostEnterpriseComplianceHostCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseComplianceHostCountTop99pSum
}

// GetCsmHostEnterpriseComplianceHostCountTop99pSumOk returns a tuple with the CsmHostEnterpriseComplianceHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseComplianceHostCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseComplianceHostCountTop99pSum == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseComplianceHostCountTop99pSum, true
}

// HasCsmHostEnterpriseComplianceHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmHostEnterpriseComplianceHostCountTop99pSum() bool {
	return o != nil && o.CsmHostEnterpriseComplianceHostCountTop99pSum != nil
}

// SetCsmHostEnterpriseComplianceHostCountTop99pSum gets a reference to the given int64 and assigns it to the CsmHostEnterpriseComplianceHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetCsmHostEnterpriseComplianceHostCountTop99pSum(v int64) {
	o.CsmHostEnterpriseComplianceHostCountTop99pSum = &v
}

// GetCsmHostEnterpriseCwsHostCountTop99pSum returns the CsmHostEnterpriseCwsHostCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseCwsHostCountTop99pSum() int64 {
	if o == nil || o.CsmHostEnterpriseCwsHostCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseCwsHostCountTop99pSum
}

// GetCsmHostEnterpriseCwsHostCountTop99pSumOk returns a tuple with the CsmHostEnterpriseCwsHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseCwsHostCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseCwsHostCountTop99pSum == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseCwsHostCountTop99pSum, true
}

// HasCsmHostEnterpriseCwsHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmHostEnterpriseCwsHostCountTop99pSum() bool {
	return o != nil && o.CsmHostEnterpriseCwsHostCountTop99pSum != nil
}

// SetCsmHostEnterpriseCwsHostCountTop99pSum gets a reference to the given int64 and assigns it to the CsmHostEnterpriseCwsHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetCsmHostEnterpriseCwsHostCountTop99pSum(v int64) {
	o.CsmHostEnterpriseCwsHostCountTop99pSum = &v
}

// GetCsmHostEnterpriseGcpHostCountTop99pSum returns the CsmHostEnterpriseGcpHostCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseGcpHostCountTop99pSum() int64 {
	if o == nil || o.CsmHostEnterpriseGcpHostCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseGcpHostCountTop99pSum
}

// GetCsmHostEnterpriseGcpHostCountTop99pSumOk returns a tuple with the CsmHostEnterpriseGcpHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseGcpHostCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseGcpHostCountTop99pSum == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseGcpHostCountTop99pSum, true
}

// HasCsmHostEnterpriseGcpHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmHostEnterpriseGcpHostCountTop99pSum() bool {
	return o != nil && o.CsmHostEnterpriseGcpHostCountTop99pSum != nil
}

// SetCsmHostEnterpriseGcpHostCountTop99pSum gets a reference to the given int64 and assigns it to the CsmHostEnterpriseGcpHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetCsmHostEnterpriseGcpHostCountTop99pSum(v int64) {
	o.CsmHostEnterpriseGcpHostCountTop99pSum = &v
}

// GetCsmHostEnterpriseTotalHostCountTop99pSum returns the CsmHostEnterpriseTotalHostCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseTotalHostCountTop99pSum() int64 {
	if o == nil || o.CsmHostEnterpriseTotalHostCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseTotalHostCountTop99pSum
}

// GetCsmHostEnterpriseTotalHostCountTop99pSumOk returns a tuple with the CsmHostEnterpriseTotalHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCsmHostEnterpriseTotalHostCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseTotalHostCountTop99pSum == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseTotalHostCountTop99pSum, true
}

// HasCsmHostEnterpriseTotalHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCsmHostEnterpriseTotalHostCountTop99pSum() bool {
	return o != nil && o.CsmHostEnterpriseTotalHostCountTop99pSum != nil
}

// SetCsmHostEnterpriseTotalHostCountTop99pSum gets a reference to the given int64 and assigns it to the CsmHostEnterpriseTotalHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetCsmHostEnterpriseTotalHostCountTop99pSum(v int64) {
	o.CsmHostEnterpriseTotalHostCountTop99pSum = &v
}

// GetCspmAasHostTop99pSum returns the CspmAasHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCspmAasHostTop99pSum() int64 {
	if o == nil || o.CspmAasHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CspmAasHostTop99pSum
}

// GetCspmAasHostTop99pSumOk returns a tuple with the CspmAasHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCspmAasHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.CspmAasHostTop99pSum == nil {
		return nil, false
	}
	return o.CspmAasHostTop99pSum, true
}

// HasCspmAasHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmAasHostTop99pSum() bool {
	return o != nil && o.CspmAasHostTop99pSum != nil
}

// SetCspmAasHostTop99pSum gets a reference to the given int64 and assigns it to the CspmAasHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmAasHostTop99pSum(v int64) {
	o.CspmAasHostTop99pSum = &v
}

// GetCspmAwsHostTop99pSum returns the CspmAwsHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCspmAwsHostTop99pSum() int64 {
	if o == nil || o.CspmAwsHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CspmAwsHostTop99pSum
}

// GetCspmAwsHostTop99pSumOk returns a tuple with the CspmAwsHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCspmAwsHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.CspmAwsHostTop99pSum == nil {
		return nil, false
	}
	return o.CspmAwsHostTop99pSum, true
}

// HasCspmAwsHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmAwsHostTop99pSum() bool {
	return o != nil && o.CspmAwsHostTop99pSum != nil
}

// SetCspmAwsHostTop99pSum gets a reference to the given int64 and assigns it to the CspmAwsHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmAwsHostTop99pSum(v int64) {
	o.CspmAwsHostTop99pSum = &v
}

// GetCspmAzureHostTop99pSum returns the CspmAzureHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCspmAzureHostTop99pSum() int64 {
	if o == nil || o.CspmAzureHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CspmAzureHostTop99pSum
}

// GetCspmAzureHostTop99pSumOk returns a tuple with the CspmAzureHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCspmAzureHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.CspmAzureHostTop99pSum == nil {
		return nil, false
	}
	return o.CspmAzureHostTop99pSum, true
}

// HasCspmAzureHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmAzureHostTop99pSum() bool {
	return o != nil && o.CspmAzureHostTop99pSum != nil
}

// SetCspmAzureHostTop99pSum gets a reference to the given int64 and assigns it to the CspmAzureHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmAzureHostTop99pSum(v int64) {
	o.CspmAzureHostTop99pSum = &v
}

// GetCspmContainerAvgSum returns the CspmContainerAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCspmContainerAvgSum() int64 {
	if o == nil || o.CspmContainerAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerAvgSum
}

// GetCspmContainerAvgSumOk returns a tuple with the CspmContainerAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCspmContainerAvgSumOk() (*int64, bool) {
	if o == nil || o.CspmContainerAvgSum == nil {
		return nil, false
	}
	return o.CspmContainerAvgSum, true
}

// HasCspmContainerAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmContainerAvgSum() bool {
	return o != nil && o.CspmContainerAvgSum != nil
}

// SetCspmContainerAvgSum gets a reference to the given int64 and assigns it to the CspmContainerAvgSum field.
func (o *UsageSummaryResponse) SetCspmContainerAvgSum(v int64) {
	o.CspmContainerAvgSum = &v
}

// GetCspmContainerHwmSum returns the CspmContainerHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCspmContainerHwmSum() int64 {
	if o == nil || o.CspmContainerHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerHwmSum
}

// GetCspmContainerHwmSumOk returns a tuple with the CspmContainerHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCspmContainerHwmSumOk() (*int64, bool) {
	if o == nil || o.CspmContainerHwmSum == nil {
		return nil, false
	}
	return o.CspmContainerHwmSum, true
}

// HasCspmContainerHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmContainerHwmSum() bool {
	return o != nil && o.CspmContainerHwmSum != nil
}

// SetCspmContainerHwmSum gets a reference to the given int64 and assigns it to the CspmContainerHwmSum field.
func (o *UsageSummaryResponse) SetCspmContainerHwmSum(v int64) {
	o.CspmContainerHwmSum = &v
}

// GetCspmGcpHostTop99pSum returns the CspmGcpHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCspmGcpHostTop99pSum() int64 {
	if o == nil || o.CspmGcpHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CspmGcpHostTop99pSum
}

// GetCspmGcpHostTop99pSumOk returns a tuple with the CspmGcpHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCspmGcpHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.CspmGcpHostTop99pSum == nil {
		return nil, false
	}
	return o.CspmGcpHostTop99pSum, true
}

// HasCspmGcpHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmGcpHostTop99pSum() bool {
	return o != nil && o.CspmGcpHostTop99pSum != nil
}

// SetCspmGcpHostTop99pSum gets a reference to the given int64 and assigns it to the CspmGcpHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmGcpHostTop99pSum(v int64) {
	o.CspmGcpHostTop99pSum = &v
}

// GetCspmHostTop99pSum returns the CspmHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCspmHostTop99pSum() int64 {
	if o == nil || o.CspmHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CspmHostTop99pSum
}

// GetCspmHostTop99pSumOk returns a tuple with the CspmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCspmHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.CspmHostTop99pSum == nil {
		return nil, false
	}
	return o.CspmHostTop99pSum, true
}

// HasCspmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCspmHostTop99pSum() bool {
	return o != nil && o.CspmHostTop99pSum != nil
}

// SetCspmHostTop99pSum gets a reference to the given int64 and assigns it to the CspmHostTop99pSum field.
func (o *UsageSummaryResponse) SetCspmHostTop99pSum(v int64) {
	o.CspmHostTop99pSum = &v
}

// GetCustomHistoricalTsSum returns the CustomHistoricalTsSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCustomHistoricalTsSum() int64 {
	if o == nil || o.CustomHistoricalTsSum == nil {
		var ret int64
		return ret
	}
	return *o.CustomHistoricalTsSum
}

// GetCustomHistoricalTsSumOk returns a tuple with the CustomHistoricalTsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCustomHistoricalTsSumOk() (*int64, bool) {
	if o == nil || o.CustomHistoricalTsSum == nil {
		return nil, false
	}
	return o.CustomHistoricalTsSum, true
}

// HasCustomHistoricalTsSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCustomHistoricalTsSum() bool {
	return o != nil && o.CustomHistoricalTsSum != nil
}

// SetCustomHistoricalTsSum gets a reference to the given int64 and assigns it to the CustomHistoricalTsSum field.
func (o *UsageSummaryResponse) SetCustomHistoricalTsSum(v int64) {
	o.CustomHistoricalTsSum = &v
}

// GetCustomLiveTsSum returns the CustomLiveTsSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCustomLiveTsSum() int64 {
	if o == nil || o.CustomLiveTsSum == nil {
		var ret int64
		return ret
	}
	return *o.CustomLiveTsSum
}

// GetCustomLiveTsSumOk returns a tuple with the CustomLiveTsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCustomLiveTsSumOk() (*int64, bool) {
	if o == nil || o.CustomLiveTsSum == nil {
		return nil, false
	}
	return o.CustomLiveTsSum, true
}

// HasCustomLiveTsSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCustomLiveTsSum() bool {
	return o != nil && o.CustomLiveTsSum != nil
}

// SetCustomLiveTsSum gets a reference to the given int64 and assigns it to the CustomLiveTsSum field.
func (o *UsageSummaryResponse) SetCustomLiveTsSum(v int64) {
	o.CustomLiveTsSum = &v
}

// GetCustomTsSum returns the CustomTsSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCustomTsSum() int64 {
	if o == nil || o.CustomTsSum == nil {
		var ret int64
		return ret
	}
	return *o.CustomTsSum
}

// GetCustomTsSumOk returns a tuple with the CustomTsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCustomTsSumOk() (*int64, bool) {
	if o == nil || o.CustomTsSum == nil {
		return nil, false
	}
	return o.CustomTsSum, true
}

// HasCustomTsSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCustomTsSum() bool {
	return o != nil && o.CustomTsSum != nil
}

// SetCustomTsSum gets a reference to the given int64 and assigns it to the CustomTsSum field.
func (o *UsageSummaryResponse) SetCustomTsSum(v int64) {
	o.CustomTsSum = &v
}

// GetCwsContainerAvgSum returns the CwsContainerAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCwsContainerAvgSum() int64 {
	if o == nil || o.CwsContainerAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.CwsContainerAvgSum
}

// GetCwsContainerAvgSumOk returns a tuple with the CwsContainerAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCwsContainerAvgSumOk() (*int64, bool) {
	if o == nil || o.CwsContainerAvgSum == nil {
		return nil, false
	}
	return o.CwsContainerAvgSum, true
}

// HasCwsContainerAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCwsContainerAvgSum() bool {
	return o != nil && o.CwsContainerAvgSum != nil
}

// SetCwsContainerAvgSum gets a reference to the given int64 and assigns it to the CwsContainerAvgSum field.
func (o *UsageSummaryResponse) SetCwsContainerAvgSum(v int64) {
	o.CwsContainerAvgSum = &v
}

// GetCwsFargateTaskAvgSum returns the CwsFargateTaskAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCwsFargateTaskAvgSum() int64 {
	if o == nil || o.CwsFargateTaskAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.CwsFargateTaskAvgSum
}

// GetCwsFargateTaskAvgSumOk returns a tuple with the CwsFargateTaskAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCwsFargateTaskAvgSumOk() (*int64, bool) {
	if o == nil || o.CwsFargateTaskAvgSum == nil {
		return nil, false
	}
	return o.CwsFargateTaskAvgSum, true
}

// HasCwsFargateTaskAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCwsFargateTaskAvgSum() bool {
	return o != nil && o.CwsFargateTaskAvgSum != nil
}

// SetCwsFargateTaskAvgSum gets a reference to the given int64 and assigns it to the CwsFargateTaskAvgSum field.
func (o *UsageSummaryResponse) SetCwsFargateTaskAvgSum(v int64) {
	o.CwsFargateTaskAvgSum = &v
}

// GetCwsHostTop99pSum returns the CwsHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetCwsHostTop99pSum() int64 {
	if o == nil || o.CwsHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.CwsHostTop99pSum
}

// GetCwsHostTop99pSumOk returns a tuple with the CwsHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetCwsHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.CwsHostTop99pSum == nil {
		return nil, false
	}
	return o.CwsHostTop99pSum, true
}

// HasCwsHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasCwsHostTop99pSum() bool {
	return o != nil && o.CwsHostTop99pSum != nil
}

// SetCwsHostTop99pSum gets a reference to the given int64 and assigns it to the CwsHostTop99pSum field.
func (o *UsageSummaryResponse) SetCwsHostTop99pSum(v int64) {
	o.CwsHostTop99pSum = &v
}

// GetDataJobsMonitoringHostHrAggSum returns the DataJobsMonitoringHostHrAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetDataJobsMonitoringHostHrAggSum() int64 {
	if o == nil || o.DataJobsMonitoringHostHrAggSum == nil {
		var ret int64
		return ret
	}
	return *o.DataJobsMonitoringHostHrAggSum
}

// GetDataJobsMonitoringHostHrAggSumOk returns a tuple with the DataJobsMonitoringHostHrAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetDataJobsMonitoringHostHrAggSumOk() (*int64, bool) {
	if o == nil || o.DataJobsMonitoringHostHrAggSum == nil {
		return nil, false
	}
	return o.DataJobsMonitoringHostHrAggSum, true
}

// HasDataJobsMonitoringHostHrAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasDataJobsMonitoringHostHrAggSum() bool {
	return o != nil && o.DataJobsMonitoringHostHrAggSum != nil
}

// SetDataJobsMonitoringHostHrAggSum gets a reference to the given int64 and assigns it to the DataJobsMonitoringHostHrAggSum field.
func (o *UsageSummaryResponse) SetDataJobsMonitoringHostHrAggSum(v int64) {
	o.DataJobsMonitoringHostHrAggSum = &v
}

// GetDbmHostTop99pSum returns the DbmHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetDbmHostTop99pSum() int64 {
	if o == nil || o.DbmHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.DbmHostTop99pSum
}

// GetDbmHostTop99pSumOk returns a tuple with the DbmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetDbmHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.DbmHostTop99pSum == nil {
		return nil, false
	}
	return o.DbmHostTop99pSum, true
}

// HasDbmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasDbmHostTop99pSum() bool {
	return o != nil && o.DbmHostTop99pSum != nil
}

// SetDbmHostTop99pSum gets a reference to the given int64 and assigns it to the DbmHostTop99pSum field.
func (o *UsageSummaryResponse) SetDbmHostTop99pSum(v int64) {
	o.DbmHostTop99pSum = &v
}

// GetDbmQueriesAvgSum returns the DbmQueriesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetDbmQueriesAvgSum() int64 {
	if o == nil || o.DbmQueriesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.DbmQueriesAvgSum
}

// GetDbmQueriesAvgSumOk returns a tuple with the DbmQueriesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetDbmQueriesAvgSumOk() (*int64, bool) {
	if o == nil || o.DbmQueriesAvgSum == nil {
		return nil, false
	}
	return o.DbmQueriesAvgSum, true
}

// HasDbmQueriesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasDbmQueriesAvgSum() bool {
	return o != nil && o.DbmQueriesAvgSum != nil
}

// SetDbmQueriesAvgSum gets a reference to the given int64 and assigns it to the DbmQueriesAvgSum field.
func (o *UsageSummaryResponse) SetDbmQueriesAvgSum(v int64) {
	o.DbmQueriesAvgSum = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEndDate() bool {
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *UsageSummaryResponse) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetEphInfraHostAgentAggSum returns the EphInfraHostAgentAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostAgentAggSum() int64 {
	if o == nil || o.EphInfraHostAgentAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostAgentAggSum
}

// GetEphInfraHostAgentAggSumOk returns a tuple with the EphInfraHostAgentAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostAgentAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostAgentAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostAgentAggSum, true
}

// HasEphInfraHostAgentAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostAgentAggSum() bool {
	return o != nil && o.EphInfraHostAgentAggSum != nil
}

// SetEphInfraHostAgentAggSum gets a reference to the given int64 and assigns it to the EphInfraHostAgentAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostAgentAggSum(v int64) {
	o.EphInfraHostAgentAggSum = &v
}

// GetEphInfraHostAlibabaAggSum returns the EphInfraHostAlibabaAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostAlibabaAggSum() int64 {
	if o == nil || o.EphInfraHostAlibabaAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostAlibabaAggSum
}

// GetEphInfraHostAlibabaAggSumOk returns a tuple with the EphInfraHostAlibabaAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostAlibabaAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostAlibabaAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostAlibabaAggSum, true
}

// HasEphInfraHostAlibabaAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostAlibabaAggSum() bool {
	return o != nil && o.EphInfraHostAlibabaAggSum != nil
}

// SetEphInfraHostAlibabaAggSum gets a reference to the given int64 and assigns it to the EphInfraHostAlibabaAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostAlibabaAggSum(v int64) {
	o.EphInfraHostAlibabaAggSum = &v
}

// GetEphInfraHostAwsAggSum returns the EphInfraHostAwsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostAwsAggSum() int64 {
	if o == nil || o.EphInfraHostAwsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostAwsAggSum
}

// GetEphInfraHostAwsAggSumOk returns a tuple with the EphInfraHostAwsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostAwsAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostAwsAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostAwsAggSum, true
}

// HasEphInfraHostAwsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostAwsAggSum() bool {
	return o != nil && o.EphInfraHostAwsAggSum != nil
}

// SetEphInfraHostAwsAggSum gets a reference to the given int64 and assigns it to the EphInfraHostAwsAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostAwsAggSum(v int64) {
	o.EphInfraHostAwsAggSum = &v
}

// GetEphInfraHostAzureAggSum returns the EphInfraHostAzureAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostAzureAggSum() int64 {
	if o == nil || o.EphInfraHostAzureAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostAzureAggSum
}

// GetEphInfraHostAzureAggSumOk returns a tuple with the EphInfraHostAzureAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostAzureAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostAzureAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostAzureAggSum, true
}

// HasEphInfraHostAzureAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostAzureAggSum() bool {
	return o != nil && o.EphInfraHostAzureAggSum != nil
}

// SetEphInfraHostAzureAggSum gets a reference to the given int64 and assigns it to the EphInfraHostAzureAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostAzureAggSum(v int64) {
	o.EphInfraHostAzureAggSum = &v
}

// GetEphInfraHostEntAggSum returns the EphInfraHostEntAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostEntAggSum() int64 {
	if o == nil || o.EphInfraHostEntAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostEntAggSum
}

// GetEphInfraHostEntAggSumOk returns a tuple with the EphInfraHostEntAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostEntAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostEntAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostEntAggSum, true
}

// HasEphInfraHostEntAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostEntAggSum() bool {
	return o != nil && o.EphInfraHostEntAggSum != nil
}

// SetEphInfraHostEntAggSum gets a reference to the given int64 and assigns it to the EphInfraHostEntAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostEntAggSum(v int64) {
	o.EphInfraHostEntAggSum = &v
}

// GetEphInfraHostGcpAggSum returns the EphInfraHostGcpAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostGcpAggSum() int64 {
	if o == nil || o.EphInfraHostGcpAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostGcpAggSum
}

// GetEphInfraHostGcpAggSumOk returns a tuple with the EphInfraHostGcpAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostGcpAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostGcpAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostGcpAggSum, true
}

// HasEphInfraHostGcpAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostGcpAggSum() bool {
	return o != nil && o.EphInfraHostGcpAggSum != nil
}

// SetEphInfraHostGcpAggSum gets a reference to the given int64 and assigns it to the EphInfraHostGcpAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostGcpAggSum(v int64) {
	o.EphInfraHostGcpAggSum = &v
}

// GetEphInfraHostHerokuAggSum returns the EphInfraHostHerokuAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostHerokuAggSum() int64 {
	if o == nil || o.EphInfraHostHerokuAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostHerokuAggSum
}

// GetEphInfraHostHerokuAggSumOk returns a tuple with the EphInfraHostHerokuAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostHerokuAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostHerokuAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostHerokuAggSum, true
}

// HasEphInfraHostHerokuAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostHerokuAggSum() bool {
	return o != nil && o.EphInfraHostHerokuAggSum != nil
}

// SetEphInfraHostHerokuAggSum gets a reference to the given int64 and assigns it to the EphInfraHostHerokuAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostHerokuAggSum(v int64) {
	o.EphInfraHostHerokuAggSum = &v
}

// GetEphInfraHostOnlyAasAggSum returns the EphInfraHostOnlyAasAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostOnlyAasAggSum() int64 {
	if o == nil || o.EphInfraHostOnlyAasAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostOnlyAasAggSum
}

// GetEphInfraHostOnlyAasAggSumOk returns a tuple with the EphInfraHostOnlyAasAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostOnlyAasAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostOnlyAasAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostOnlyAasAggSum, true
}

// HasEphInfraHostOnlyAasAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostOnlyAasAggSum() bool {
	return o != nil && o.EphInfraHostOnlyAasAggSum != nil
}

// SetEphInfraHostOnlyAasAggSum gets a reference to the given int64 and assigns it to the EphInfraHostOnlyAasAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostOnlyAasAggSum(v int64) {
	o.EphInfraHostOnlyAasAggSum = &v
}

// GetEphInfraHostOnlyVsphereAggSum returns the EphInfraHostOnlyVsphereAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostOnlyVsphereAggSum() int64 {
	if o == nil || o.EphInfraHostOnlyVsphereAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostOnlyVsphereAggSum
}

// GetEphInfraHostOnlyVsphereAggSumOk returns a tuple with the EphInfraHostOnlyVsphereAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostOnlyVsphereAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostOnlyVsphereAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostOnlyVsphereAggSum, true
}

// HasEphInfraHostOnlyVsphereAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostOnlyVsphereAggSum() bool {
	return o != nil && o.EphInfraHostOnlyVsphereAggSum != nil
}

// SetEphInfraHostOnlyVsphereAggSum gets a reference to the given int64 and assigns it to the EphInfraHostOnlyVsphereAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostOnlyVsphereAggSum(v int64) {
	o.EphInfraHostOnlyVsphereAggSum = &v
}

// GetEphInfraHostOpentelemetryAggSum returns the EphInfraHostOpentelemetryAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostOpentelemetryAggSum() int64 {
	if o == nil || o.EphInfraHostOpentelemetryAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostOpentelemetryAggSum
}

// GetEphInfraHostOpentelemetryAggSumOk returns a tuple with the EphInfraHostOpentelemetryAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostOpentelemetryAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostOpentelemetryAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostOpentelemetryAggSum, true
}

// HasEphInfraHostOpentelemetryAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostOpentelemetryAggSum() bool {
	return o != nil && o.EphInfraHostOpentelemetryAggSum != nil
}

// SetEphInfraHostOpentelemetryAggSum gets a reference to the given int64 and assigns it to the EphInfraHostOpentelemetryAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostOpentelemetryAggSum(v int64) {
	o.EphInfraHostOpentelemetryAggSum = &v
}

// GetEphInfraHostOpentelemetryApmAggSum returns the EphInfraHostOpentelemetryApmAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostOpentelemetryApmAggSum() int64 {
	if o == nil || o.EphInfraHostOpentelemetryApmAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostOpentelemetryApmAggSum
}

// GetEphInfraHostOpentelemetryApmAggSumOk returns a tuple with the EphInfraHostOpentelemetryApmAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostOpentelemetryApmAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostOpentelemetryApmAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostOpentelemetryApmAggSum, true
}

// HasEphInfraHostOpentelemetryApmAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostOpentelemetryApmAggSum() bool {
	return o != nil && o.EphInfraHostOpentelemetryApmAggSum != nil
}

// SetEphInfraHostOpentelemetryApmAggSum gets a reference to the given int64 and assigns it to the EphInfraHostOpentelemetryApmAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostOpentelemetryApmAggSum(v int64) {
	o.EphInfraHostOpentelemetryApmAggSum = &v
}

// GetEphInfraHostProAggSum returns the EphInfraHostProAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostProAggSum() int64 {
	if o == nil || o.EphInfraHostProAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostProAggSum
}

// GetEphInfraHostProAggSumOk returns a tuple with the EphInfraHostProAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostProAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostProAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostProAggSum, true
}

// HasEphInfraHostProAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostProAggSum() bool {
	return o != nil && o.EphInfraHostProAggSum != nil
}

// SetEphInfraHostProAggSum gets a reference to the given int64 and assigns it to the EphInfraHostProAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostProAggSum(v int64) {
	o.EphInfraHostProAggSum = &v
}

// GetEphInfraHostProplusAggSum returns the EphInfraHostProplusAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostProplusAggSum() int64 {
	if o == nil || o.EphInfraHostProplusAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostProplusAggSum
}

// GetEphInfraHostProplusAggSumOk returns a tuple with the EphInfraHostProplusAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostProplusAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostProplusAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostProplusAggSum, true
}

// HasEphInfraHostProplusAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostProplusAggSum() bool {
	return o != nil && o.EphInfraHostProplusAggSum != nil
}

// SetEphInfraHostProplusAggSum gets a reference to the given int64 and assigns it to the EphInfraHostProplusAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostProplusAggSum(v int64) {
	o.EphInfraHostProplusAggSum = &v
}

// GetEphInfraHostProxmoxAggSum returns the EphInfraHostProxmoxAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEphInfraHostProxmoxAggSum() int64 {
	if o == nil || o.EphInfraHostProxmoxAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostProxmoxAggSum
}

// GetEphInfraHostProxmoxAggSumOk returns a tuple with the EphInfraHostProxmoxAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEphInfraHostProxmoxAggSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostProxmoxAggSum == nil {
		return nil, false
	}
	return o.EphInfraHostProxmoxAggSum, true
}

// HasEphInfraHostProxmoxAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEphInfraHostProxmoxAggSum() bool {
	return o != nil && o.EphInfraHostProxmoxAggSum != nil
}

// SetEphInfraHostProxmoxAggSum gets a reference to the given int64 and assigns it to the EphInfraHostProxmoxAggSum field.
func (o *UsageSummaryResponse) SetEphInfraHostProxmoxAggSum(v int64) {
	o.EphInfraHostProxmoxAggSum = &v
}

// GetErrorTrackingApmErrorEventsAggSum returns the ErrorTrackingApmErrorEventsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetErrorTrackingApmErrorEventsAggSum() int64 {
	if o == nil || o.ErrorTrackingApmErrorEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.ErrorTrackingApmErrorEventsAggSum
}

// GetErrorTrackingApmErrorEventsAggSumOk returns a tuple with the ErrorTrackingApmErrorEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetErrorTrackingApmErrorEventsAggSumOk() (*int64, bool) {
	if o == nil || o.ErrorTrackingApmErrorEventsAggSum == nil {
		return nil, false
	}
	return o.ErrorTrackingApmErrorEventsAggSum, true
}

// HasErrorTrackingApmErrorEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasErrorTrackingApmErrorEventsAggSum() bool {
	return o != nil && o.ErrorTrackingApmErrorEventsAggSum != nil
}

// SetErrorTrackingApmErrorEventsAggSum gets a reference to the given int64 and assigns it to the ErrorTrackingApmErrorEventsAggSum field.
func (o *UsageSummaryResponse) SetErrorTrackingApmErrorEventsAggSum(v int64) {
	o.ErrorTrackingApmErrorEventsAggSum = &v
}

// GetErrorTrackingErrorEventsAggSum returns the ErrorTrackingErrorEventsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetErrorTrackingErrorEventsAggSum() int64 {
	if o == nil || o.ErrorTrackingErrorEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.ErrorTrackingErrorEventsAggSum
}

// GetErrorTrackingErrorEventsAggSumOk returns a tuple with the ErrorTrackingErrorEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetErrorTrackingErrorEventsAggSumOk() (*int64, bool) {
	if o == nil || o.ErrorTrackingErrorEventsAggSum == nil {
		return nil, false
	}
	return o.ErrorTrackingErrorEventsAggSum, true
}

// HasErrorTrackingErrorEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasErrorTrackingErrorEventsAggSum() bool {
	return o != nil && o.ErrorTrackingErrorEventsAggSum != nil
}

// SetErrorTrackingErrorEventsAggSum gets a reference to the given int64 and assigns it to the ErrorTrackingErrorEventsAggSum field.
func (o *UsageSummaryResponse) SetErrorTrackingErrorEventsAggSum(v int64) {
	o.ErrorTrackingErrorEventsAggSum = &v
}

// GetErrorTrackingEventsAggSum returns the ErrorTrackingEventsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetErrorTrackingEventsAggSum() int64 {
	if o == nil || o.ErrorTrackingEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.ErrorTrackingEventsAggSum
}

// GetErrorTrackingEventsAggSumOk returns a tuple with the ErrorTrackingEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetErrorTrackingEventsAggSumOk() (*int64, bool) {
	if o == nil || o.ErrorTrackingEventsAggSum == nil {
		return nil, false
	}
	return o.ErrorTrackingEventsAggSum, true
}

// HasErrorTrackingEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasErrorTrackingEventsAggSum() bool {
	return o != nil && o.ErrorTrackingEventsAggSum != nil
}

// SetErrorTrackingEventsAggSum gets a reference to the given int64 and assigns it to the ErrorTrackingEventsAggSum field.
func (o *UsageSummaryResponse) SetErrorTrackingEventsAggSum(v int64) {
	o.ErrorTrackingEventsAggSum = &v
}

// GetErrorTrackingRumErrorEventsAggSum returns the ErrorTrackingRumErrorEventsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetErrorTrackingRumErrorEventsAggSum() int64 {
	if o == nil || o.ErrorTrackingRumErrorEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.ErrorTrackingRumErrorEventsAggSum
}

// GetErrorTrackingRumErrorEventsAggSumOk returns a tuple with the ErrorTrackingRumErrorEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetErrorTrackingRumErrorEventsAggSumOk() (*int64, bool) {
	if o == nil || o.ErrorTrackingRumErrorEventsAggSum == nil {
		return nil, false
	}
	return o.ErrorTrackingRumErrorEventsAggSum, true
}

// HasErrorTrackingRumErrorEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasErrorTrackingRumErrorEventsAggSum() bool {
	return o != nil && o.ErrorTrackingRumErrorEventsAggSum != nil
}

// SetErrorTrackingRumErrorEventsAggSum gets a reference to the given int64 and assigns it to the ErrorTrackingRumErrorEventsAggSum field.
func (o *UsageSummaryResponse) SetErrorTrackingRumErrorEventsAggSum(v int64) {
	o.ErrorTrackingRumErrorEventsAggSum = &v
}

// GetEventManagementCorrelationAggSum returns the EventManagementCorrelationAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEventManagementCorrelationAggSum() int64 {
	if o == nil || o.EventManagementCorrelationAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EventManagementCorrelationAggSum
}

// GetEventManagementCorrelationAggSumOk returns a tuple with the EventManagementCorrelationAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEventManagementCorrelationAggSumOk() (*int64, bool) {
	if o == nil || o.EventManagementCorrelationAggSum == nil {
		return nil, false
	}
	return o.EventManagementCorrelationAggSum, true
}

// HasEventManagementCorrelationAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEventManagementCorrelationAggSum() bool {
	return o != nil && o.EventManagementCorrelationAggSum != nil
}

// SetEventManagementCorrelationAggSum gets a reference to the given int64 and assigns it to the EventManagementCorrelationAggSum field.
func (o *UsageSummaryResponse) SetEventManagementCorrelationAggSum(v int64) {
	o.EventManagementCorrelationAggSum = &v
}

// GetEventManagementCorrelationCorrelatedEventsAggSum returns the EventManagementCorrelationCorrelatedEventsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEventManagementCorrelationCorrelatedEventsAggSum() int64 {
	if o == nil || o.EventManagementCorrelationCorrelatedEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EventManagementCorrelationCorrelatedEventsAggSum
}

// GetEventManagementCorrelationCorrelatedEventsAggSumOk returns a tuple with the EventManagementCorrelationCorrelatedEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEventManagementCorrelationCorrelatedEventsAggSumOk() (*int64, bool) {
	if o == nil || o.EventManagementCorrelationCorrelatedEventsAggSum == nil {
		return nil, false
	}
	return o.EventManagementCorrelationCorrelatedEventsAggSum, true
}

// HasEventManagementCorrelationCorrelatedEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEventManagementCorrelationCorrelatedEventsAggSum() bool {
	return o != nil && o.EventManagementCorrelationCorrelatedEventsAggSum != nil
}

// SetEventManagementCorrelationCorrelatedEventsAggSum gets a reference to the given int64 and assigns it to the EventManagementCorrelationCorrelatedEventsAggSum field.
func (o *UsageSummaryResponse) SetEventManagementCorrelationCorrelatedEventsAggSum(v int64) {
	o.EventManagementCorrelationCorrelatedEventsAggSum = &v
}

// GetEventManagementCorrelationCorrelatedRelatedEventsAggSum returns the EventManagementCorrelationCorrelatedRelatedEventsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetEventManagementCorrelationCorrelatedRelatedEventsAggSum() int64 {
	if o == nil || o.EventManagementCorrelationCorrelatedRelatedEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.EventManagementCorrelationCorrelatedRelatedEventsAggSum
}

// GetEventManagementCorrelationCorrelatedRelatedEventsAggSumOk returns a tuple with the EventManagementCorrelationCorrelatedRelatedEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetEventManagementCorrelationCorrelatedRelatedEventsAggSumOk() (*int64, bool) {
	if o == nil || o.EventManagementCorrelationCorrelatedRelatedEventsAggSum == nil {
		return nil, false
	}
	return o.EventManagementCorrelationCorrelatedRelatedEventsAggSum, true
}

// HasEventManagementCorrelationCorrelatedRelatedEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasEventManagementCorrelationCorrelatedRelatedEventsAggSum() bool {
	return o != nil && o.EventManagementCorrelationCorrelatedRelatedEventsAggSum != nil
}

// SetEventManagementCorrelationCorrelatedRelatedEventsAggSum gets a reference to the given int64 and assigns it to the EventManagementCorrelationCorrelatedRelatedEventsAggSum field.
func (o *UsageSummaryResponse) SetEventManagementCorrelationCorrelatedRelatedEventsAggSum(v int64) {
	o.EventManagementCorrelationCorrelatedRelatedEventsAggSum = &v
}

// GetFargateContainerProfilerProfilingFargateAvgSum returns the FargateContainerProfilerProfilingFargateAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFargateContainerProfilerProfilingFargateAvgSum() int64 {
	if o == nil || o.FargateContainerProfilerProfilingFargateAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FargateContainerProfilerProfilingFargateAvgSum
}

// GetFargateContainerProfilerProfilingFargateAvgSumOk returns a tuple with the FargateContainerProfilerProfilingFargateAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFargateContainerProfilerProfilingFargateAvgSumOk() (*int64, bool) {
	if o == nil || o.FargateContainerProfilerProfilingFargateAvgSum == nil {
		return nil, false
	}
	return o.FargateContainerProfilerProfilingFargateAvgSum, true
}

// HasFargateContainerProfilerProfilingFargateAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFargateContainerProfilerProfilingFargateAvgSum() bool {
	return o != nil && o.FargateContainerProfilerProfilingFargateAvgSum != nil
}

// SetFargateContainerProfilerProfilingFargateAvgSum gets a reference to the given int64 and assigns it to the FargateContainerProfilerProfilingFargateAvgSum field.
func (o *UsageSummaryResponse) SetFargateContainerProfilerProfilingFargateAvgSum(v int64) {
	o.FargateContainerProfilerProfilingFargateAvgSum = &v
}

// GetFargateContainerProfilerProfilingFargateEksAvgSum returns the FargateContainerProfilerProfilingFargateEksAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFargateContainerProfilerProfilingFargateEksAvgSum() int64 {
	if o == nil || o.FargateContainerProfilerProfilingFargateEksAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FargateContainerProfilerProfilingFargateEksAvgSum
}

// GetFargateContainerProfilerProfilingFargateEksAvgSumOk returns a tuple with the FargateContainerProfilerProfilingFargateEksAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFargateContainerProfilerProfilingFargateEksAvgSumOk() (*int64, bool) {
	if o == nil || o.FargateContainerProfilerProfilingFargateEksAvgSum == nil {
		return nil, false
	}
	return o.FargateContainerProfilerProfilingFargateEksAvgSum, true
}

// HasFargateContainerProfilerProfilingFargateEksAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFargateContainerProfilerProfilingFargateEksAvgSum() bool {
	return o != nil && o.FargateContainerProfilerProfilingFargateEksAvgSum != nil
}

// SetFargateContainerProfilerProfilingFargateEksAvgSum gets a reference to the given int64 and assigns it to the FargateContainerProfilerProfilingFargateEksAvgSum field.
func (o *UsageSummaryResponse) SetFargateContainerProfilerProfilingFargateEksAvgSum(v int64) {
	o.FargateContainerProfilerProfilingFargateEksAvgSum = &v
}

// GetFargateTasksCountAvgSum returns the FargateTasksCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFargateTasksCountAvgSum() int64 {
	if o == nil || o.FargateTasksCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountAvgSum
}

// GetFargateTasksCountAvgSumOk returns a tuple with the FargateTasksCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFargateTasksCountAvgSumOk() (*int64, bool) {
	if o == nil || o.FargateTasksCountAvgSum == nil {
		return nil, false
	}
	return o.FargateTasksCountAvgSum, true
}

// HasFargateTasksCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFargateTasksCountAvgSum() bool {
	return o != nil && o.FargateTasksCountAvgSum != nil
}

// SetFargateTasksCountAvgSum gets a reference to the given int64 and assigns it to the FargateTasksCountAvgSum field.
func (o *UsageSummaryResponse) SetFargateTasksCountAvgSum(v int64) {
	o.FargateTasksCountAvgSum = &v
}

// GetFargateTasksCountHwmSum returns the FargateTasksCountHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFargateTasksCountHwmSum() int64 {
	if o == nil || o.FargateTasksCountHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountHwmSum
}

// GetFargateTasksCountHwmSumOk returns a tuple with the FargateTasksCountHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFargateTasksCountHwmSumOk() (*int64, bool) {
	if o == nil || o.FargateTasksCountHwmSum == nil {
		return nil, false
	}
	return o.FargateTasksCountHwmSum, true
}

// HasFargateTasksCountHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFargateTasksCountHwmSum() bool {
	return o != nil && o.FargateTasksCountHwmSum != nil
}

// SetFargateTasksCountHwmSum gets a reference to the given int64 and assigns it to the FargateTasksCountHwmSum field.
func (o *UsageSummaryResponse) SetFargateTasksCountHwmSum(v int64) {
	o.FargateTasksCountHwmSum = &v
}

// GetFlexLogsComputeLargeAvgSum returns the FlexLogsComputeLargeAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFlexLogsComputeLargeAvgSum() int64 {
	if o == nil || o.FlexLogsComputeLargeAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeLargeAvgSum
}

// GetFlexLogsComputeLargeAvgSumOk returns a tuple with the FlexLogsComputeLargeAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFlexLogsComputeLargeAvgSumOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeLargeAvgSum == nil {
		return nil, false
	}
	return o.FlexLogsComputeLargeAvgSum, true
}

// HasFlexLogsComputeLargeAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFlexLogsComputeLargeAvgSum() bool {
	return o != nil && o.FlexLogsComputeLargeAvgSum != nil
}

// SetFlexLogsComputeLargeAvgSum gets a reference to the given int64 and assigns it to the FlexLogsComputeLargeAvgSum field.
func (o *UsageSummaryResponse) SetFlexLogsComputeLargeAvgSum(v int64) {
	o.FlexLogsComputeLargeAvgSum = &v
}

// GetFlexLogsComputeMediumAvgSum returns the FlexLogsComputeMediumAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFlexLogsComputeMediumAvgSum() int64 {
	if o == nil || o.FlexLogsComputeMediumAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeMediumAvgSum
}

// GetFlexLogsComputeMediumAvgSumOk returns a tuple with the FlexLogsComputeMediumAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFlexLogsComputeMediumAvgSumOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeMediumAvgSum == nil {
		return nil, false
	}
	return o.FlexLogsComputeMediumAvgSum, true
}

// HasFlexLogsComputeMediumAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFlexLogsComputeMediumAvgSum() bool {
	return o != nil && o.FlexLogsComputeMediumAvgSum != nil
}

// SetFlexLogsComputeMediumAvgSum gets a reference to the given int64 and assigns it to the FlexLogsComputeMediumAvgSum field.
func (o *UsageSummaryResponse) SetFlexLogsComputeMediumAvgSum(v int64) {
	o.FlexLogsComputeMediumAvgSum = &v
}

// GetFlexLogsComputeSmallAvgSum returns the FlexLogsComputeSmallAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFlexLogsComputeSmallAvgSum() int64 {
	if o == nil || o.FlexLogsComputeSmallAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeSmallAvgSum
}

// GetFlexLogsComputeSmallAvgSumOk returns a tuple with the FlexLogsComputeSmallAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFlexLogsComputeSmallAvgSumOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeSmallAvgSum == nil {
		return nil, false
	}
	return o.FlexLogsComputeSmallAvgSum, true
}

// HasFlexLogsComputeSmallAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFlexLogsComputeSmallAvgSum() bool {
	return o != nil && o.FlexLogsComputeSmallAvgSum != nil
}

// SetFlexLogsComputeSmallAvgSum gets a reference to the given int64 and assigns it to the FlexLogsComputeSmallAvgSum field.
func (o *UsageSummaryResponse) SetFlexLogsComputeSmallAvgSum(v int64) {
	o.FlexLogsComputeSmallAvgSum = &v
}

// GetFlexLogsComputeXlargeAvgSum returns the FlexLogsComputeXlargeAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFlexLogsComputeXlargeAvgSum() int64 {
	if o == nil || o.FlexLogsComputeXlargeAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeXlargeAvgSum
}

// GetFlexLogsComputeXlargeAvgSumOk returns a tuple with the FlexLogsComputeXlargeAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFlexLogsComputeXlargeAvgSumOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeXlargeAvgSum == nil {
		return nil, false
	}
	return o.FlexLogsComputeXlargeAvgSum, true
}

// HasFlexLogsComputeXlargeAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFlexLogsComputeXlargeAvgSum() bool {
	return o != nil && o.FlexLogsComputeXlargeAvgSum != nil
}

// SetFlexLogsComputeXlargeAvgSum gets a reference to the given int64 and assigns it to the FlexLogsComputeXlargeAvgSum field.
func (o *UsageSummaryResponse) SetFlexLogsComputeXlargeAvgSum(v int64) {
	o.FlexLogsComputeXlargeAvgSum = &v
}

// GetFlexLogsComputeXsmallAvgSum returns the FlexLogsComputeXsmallAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFlexLogsComputeXsmallAvgSum() int64 {
	if o == nil || o.FlexLogsComputeXsmallAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeXsmallAvgSum
}

// GetFlexLogsComputeXsmallAvgSumOk returns a tuple with the FlexLogsComputeXsmallAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFlexLogsComputeXsmallAvgSumOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeXsmallAvgSum == nil {
		return nil, false
	}
	return o.FlexLogsComputeXsmallAvgSum, true
}

// HasFlexLogsComputeXsmallAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFlexLogsComputeXsmallAvgSum() bool {
	return o != nil && o.FlexLogsComputeXsmallAvgSum != nil
}

// SetFlexLogsComputeXsmallAvgSum gets a reference to the given int64 and assigns it to the FlexLogsComputeXsmallAvgSum field.
func (o *UsageSummaryResponse) SetFlexLogsComputeXsmallAvgSum(v int64) {
	o.FlexLogsComputeXsmallAvgSum = &v
}

// GetFlexLogsStarterAvgSum returns the FlexLogsStarterAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFlexLogsStarterAvgSum() int64 {
	if o == nil || o.FlexLogsStarterAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsStarterAvgSum
}

// GetFlexLogsStarterAvgSumOk returns a tuple with the FlexLogsStarterAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFlexLogsStarterAvgSumOk() (*int64, bool) {
	if o == nil || o.FlexLogsStarterAvgSum == nil {
		return nil, false
	}
	return o.FlexLogsStarterAvgSum, true
}

// HasFlexLogsStarterAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFlexLogsStarterAvgSum() bool {
	return o != nil && o.FlexLogsStarterAvgSum != nil
}

// SetFlexLogsStarterAvgSum gets a reference to the given int64 and assigns it to the FlexLogsStarterAvgSum field.
func (o *UsageSummaryResponse) SetFlexLogsStarterAvgSum(v int64) {
	o.FlexLogsStarterAvgSum = &v
}

// GetFlexLogsStarterStorageIndexAvgSum returns the FlexLogsStarterStorageIndexAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFlexLogsStarterStorageIndexAvgSum() int64 {
	if o == nil || o.FlexLogsStarterStorageIndexAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsStarterStorageIndexAvgSum
}

// GetFlexLogsStarterStorageIndexAvgSumOk returns a tuple with the FlexLogsStarterStorageIndexAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFlexLogsStarterStorageIndexAvgSumOk() (*int64, bool) {
	if o == nil || o.FlexLogsStarterStorageIndexAvgSum == nil {
		return nil, false
	}
	return o.FlexLogsStarterStorageIndexAvgSum, true
}

// HasFlexLogsStarterStorageIndexAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFlexLogsStarterStorageIndexAvgSum() bool {
	return o != nil && o.FlexLogsStarterStorageIndexAvgSum != nil
}

// SetFlexLogsStarterStorageIndexAvgSum gets a reference to the given int64 and assigns it to the FlexLogsStarterStorageIndexAvgSum field.
func (o *UsageSummaryResponse) SetFlexLogsStarterStorageIndexAvgSum(v int64) {
	o.FlexLogsStarterStorageIndexAvgSum = &v
}

// GetFlexLogsStarterStorageRetentionAdjustmentAvgSum returns the FlexLogsStarterStorageRetentionAdjustmentAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFlexLogsStarterStorageRetentionAdjustmentAvgSum() int64 {
	if o == nil || o.FlexLogsStarterStorageRetentionAdjustmentAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsStarterStorageRetentionAdjustmentAvgSum
}

// GetFlexLogsStarterStorageRetentionAdjustmentAvgSumOk returns a tuple with the FlexLogsStarterStorageRetentionAdjustmentAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFlexLogsStarterStorageRetentionAdjustmentAvgSumOk() (*int64, bool) {
	if o == nil || o.FlexLogsStarterStorageRetentionAdjustmentAvgSum == nil {
		return nil, false
	}
	return o.FlexLogsStarterStorageRetentionAdjustmentAvgSum, true
}

// HasFlexLogsStarterStorageRetentionAdjustmentAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFlexLogsStarterStorageRetentionAdjustmentAvgSum() bool {
	return o != nil && o.FlexLogsStarterStorageRetentionAdjustmentAvgSum != nil
}

// SetFlexLogsStarterStorageRetentionAdjustmentAvgSum gets a reference to the given int64 and assigns it to the FlexLogsStarterStorageRetentionAdjustmentAvgSum field.
func (o *UsageSummaryResponse) SetFlexLogsStarterStorageRetentionAdjustmentAvgSum(v int64) {
	o.FlexLogsStarterStorageRetentionAdjustmentAvgSum = &v
}

// GetFlexStoredLogsAvgSum returns the FlexStoredLogsAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetFlexStoredLogsAvgSum() int64 {
	if o == nil || o.FlexStoredLogsAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.FlexStoredLogsAvgSum
}

// GetFlexStoredLogsAvgSumOk returns a tuple with the FlexStoredLogsAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetFlexStoredLogsAvgSumOk() (*int64, bool) {
	if o == nil || o.FlexStoredLogsAvgSum == nil {
		return nil, false
	}
	return o.FlexStoredLogsAvgSum, true
}

// HasFlexStoredLogsAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasFlexStoredLogsAvgSum() bool {
	return o != nil && o.FlexStoredLogsAvgSum != nil
}

// SetFlexStoredLogsAvgSum gets a reference to the given int64 and assigns it to the FlexStoredLogsAvgSum field.
func (o *UsageSummaryResponse) SetFlexStoredLogsAvgSum(v int64) {
	o.FlexStoredLogsAvgSum = &v
}

// GetForwardingEventsBytesAggSum returns the ForwardingEventsBytesAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetForwardingEventsBytesAggSum() int64 {
	if o == nil || o.ForwardingEventsBytesAggSum == nil {
		var ret int64
		return ret
	}
	return *o.ForwardingEventsBytesAggSum
}

// GetForwardingEventsBytesAggSumOk returns a tuple with the ForwardingEventsBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetForwardingEventsBytesAggSumOk() (*int64, bool) {
	if o == nil || o.ForwardingEventsBytesAggSum == nil {
		return nil, false
	}
	return o.ForwardingEventsBytesAggSum, true
}

// HasForwardingEventsBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasForwardingEventsBytesAggSum() bool {
	return o != nil && o.ForwardingEventsBytesAggSum != nil
}

// SetForwardingEventsBytesAggSum gets a reference to the given int64 and assigns it to the ForwardingEventsBytesAggSum field.
func (o *UsageSummaryResponse) SetForwardingEventsBytesAggSum(v int64) {
	o.ForwardingEventsBytesAggSum = &v
}

// GetGcpHostTop99pSum returns the GcpHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetGcpHostTop99pSum() int64 {
	if o == nil || o.GcpHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.GcpHostTop99pSum
}

// GetGcpHostTop99pSumOk returns a tuple with the GcpHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetGcpHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.GcpHostTop99pSum == nil {
		return nil, false
	}
	return o.GcpHostTop99pSum, true
}

// HasGcpHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasGcpHostTop99pSum() bool {
	return o != nil && o.GcpHostTop99pSum != nil
}

// SetGcpHostTop99pSum gets a reference to the given int64 and assigns it to the GcpHostTop99pSum field.
func (o *UsageSummaryResponse) SetGcpHostTop99pSum(v int64) {
	o.GcpHostTop99pSum = &v
}

// GetHerokuHostTop99pSum returns the HerokuHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetHerokuHostTop99pSum() int64 {
	if o == nil || o.HerokuHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.HerokuHostTop99pSum
}

// GetHerokuHostTop99pSumOk returns a tuple with the HerokuHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetHerokuHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.HerokuHostTop99pSum == nil {
		return nil, false
	}
	return o.HerokuHostTop99pSum, true
}

// HasHerokuHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasHerokuHostTop99pSum() bool {
	return o != nil && o.HerokuHostTop99pSum != nil
}

// SetHerokuHostTop99pSum gets a reference to the given int64 and assigns it to the HerokuHostTop99pSum field.
func (o *UsageSummaryResponse) SetHerokuHostTop99pSum(v int64) {
	o.HerokuHostTop99pSum = &v
}

// GetIncidentManagementMonthlyActiveUsersHwmSum returns the IncidentManagementMonthlyActiveUsersHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetIncidentManagementMonthlyActiveUsersHwmSum() int64 {
	if o == nil || o.IncidentManagementMonthlyActiveUsersHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.IncidentManagementMonthlyActiveUsersHwmSum
}

// GetIncidentManagementMonthlyActiveUsersHwmSumOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetIncidentManagementMonthlyActiveUsersHwmSumOk() (*int64, bool) {
	if o == nil || o.IncidentManagementMonthlyActiveUsersHwmSum == nil {
		return nil, false
	}
	return o.IncidentManagementMonthlyActiveUsersHwmSum, true
}

// HasIncidentManagementMonthlyActiveUsersHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIncidentManagementMonthlyActiveUsersHwmSum() bool {
	return o != nil && o.IncidentManagementMonthlyActiveUsersHwmSum != nil
}

// SetIncidentManagementMonthlyActiveUsersHwmSum gets a reference to the given int64 and assigns it to the IncidentManagementMonthlyActiveUsersHwmSum field.
func (o *UsageSummaryResponse) SetIncidentManagementMonthlyActiveUsersHwmSum(v int64) {
	o.IncidentManagementMonthlyActiveUsersHwmSum = &v
}

// GetIncidentManagementSeatsHwmSum returns the IncidentManagementSeatsHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetIncidentManagementSeatsHwmSum() int64 {
	if o == nil || o.IncidentManagementSeatsHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.IncidentManagementSeatsHwmSum
}

// GetIncidentManagementSeatsHwmSumOk returns a tuple with the IncidentManagementSeatsHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetIncidentManagementSeatsHwmSumOk() (*int64, bool) {
	if o == nil || o.IncidentManagementSeatsHwmSum == nil {
		return nil, false
	}
	return o.IncidentManagementSeatsHwmSum, true
}

// HasIncidentManagementSeatsHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIncidentManagementSeatsHwmSum() bool {
	return o != nil && o.IncidentManagementSeatsHwmSum != nil
}

// SetIncidentManagementSeatsHwmSum gets a reference to the given int64 and assigns it to the IncidentManagementSeatsHwmSum field.
func (o *UsageSummaryResponse) SetIncidentManagementSeatsHwmSum(v int64) {
	o.IncidentManagementSeatsHwmSum = &v
}

// GetIndexedEventsCountAggSum returns the IndexedEventsCountAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetIndexedEventsCountAggSum() int64 {
	if o == nil || o.IndexedEventsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCountAggSum
}

// GetIndexedEventsCountAggSumOk returns a tuple with the IndexedEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetIndexedEventsCountAggSumOk() (*int64, bool) {
	if o == nil || o.IndexedEventsCountAggSum == nil {
		return nil, false
	}
	return o.IndexedEventsCountAggSum, true
}

// HasIndexedEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIndexedEventsCountAggSum() bool {
	return o != nil && o.IndexedEventsCountAggSum != nil
}

// SetIndexedEventsCountAggSum gets a reference to the given int64 and assigns it to the IndexedEventsCountAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetIndexedEventsCountAggSum(v int64) {
	o.IndexedEventsCountAggSum = &v
}

// GetInfraHostTop99pSum returns the InfraHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetInfraHostTop99pSum() int64 {
	if o == nil || o.InfraHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostTop99pSum
}

// GetInfraHostTop99pSumOk returns a tuple with the InfraHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetInfraHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.InfraHostTop99pSum == nil {
		return nil, false
	}
	return o.InfraHostTop99pSum, true
}

// HasInfraHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasInfraHostTop99pSum() bool {
	return o != nil && o.InfraHostTop99pSum != nil
}

// SetInfraHostTop99pSum gets a reference to the given int64 and assigns it to the InfraHostTop99pSum field.
func (o *UsageSummaryResponse) SetInfraHostTop99pSum(v int64) {
	o.InfraHostTop99pSum = &v
}

// GetIngestedEventsBytesAggSum returns the IngestedEventsBytesAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetIngestedEventsBytesAggSum() int64 {
	if o == nil || o.IngestedEventsBytesAggSum == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytesAggSum
}

// GetIngestedEventsBytesAggSumOk returns a tuple with the IngestedEventsBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetIngestedEventsBytesAggSumOk() (*int64, bool) {
	if o == nil || o.IngestedEventsBytesAggSum == nil {
		return nil, false
	}
	return o.IngestedEventsBytesAggSum, true
}

// HasIngestedEventsBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIngestedEventsBytesAggSum() bool {
	return o != nil && o.IngestedEventsBytesAggSum != nil
}

// SetIngestedEventsBytesAggSum gets a reference to the given int64 and assigns it to the IngestedEventsBytesAggSum field.
func (o *UsageSummaryResponse) SetIngestedEventsBytesAggSum(v int64) {
	o.IngestedEventsBytesAggSum = &v
}

// GetIotDeviceAggSum returns the IotDeviceAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetIotDeviceAggSum() int64 {
	if o == nil || o.IotDeviceAggSum == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceAggSum
}

// GetIotDeviceAggSumOk returns a tuple with the IotDeviceAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetIotDeviceAggSumOk() (*int64, bool) {
	if o == nil || o.IotDeviceAggSum == nil {
		return nil, false
	}
	return o.IotDeviceAggSum, true
}

// HasIotDeviceAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIotDeviceAggSum() bool {
	return o != nil && o.IotDeviceAggSum != nil
}

// SetIotDeviceAggSum gets a reference to the given int64 and assigns it to the IotDeviceAggSum field.
func (o *UsageSummaryResponse) SetIotDeviceAggSum(v int64) {
	o.IotDeviceAggSum = &v
}

// GetIotDeviceTop99pSum returns the IotDeviceTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetIotDeviceTop99pSum() int64 {
	if o == nil || o.IotDeviceTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceTop99pSum
}

// GetIotDeviceTop99pSumOk returns a tuple with the IotDeviceTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetIotDeviceTop99pSumOk() (*int64, bool) {
	if o == nil || o.IotDeviceTop99pSum == nil {
		return nil, false
	}
	return o.IotDeviceTop99pSum, true
}

// HasIotDeviceTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasIotDeviceTop99pSum() bool {
	return o != nil && o.IotDeviceTop99pSum != nil
}

// SetIotDeviceTop99pSum gets a reference to the given int64 and assigns it to the IotDeviceTop99pSum field.
func (o *UsageSummaryResponse) SetIotDeviceTop99pSum(v int64) {
	o.IotDeviceTop99pSum = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLastUpdated() bool {
	return o != nil && o.LastUpdated != nil
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *UsageSummaryResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLiveIndexedEventsAggSum returns the LiveIndexedEventsAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetLiveIndexedEventsAggSum() int64 {
	if o == nil || o.LiveIndexedEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.LiveIndexedEventsAggSum
}

// GetLiveIndexedEventsAggSumOk returns a tuple with the LiveIndexedEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetLiveIndexedEventsAggSumOk() (*int64, bool) {
	if o == nil || o.LiveIndexedEventsAggSum == nil {
		return nil, false
	}
	return o.LiveIndexedEventsAggSum, true
}

// HasLiveIndexedEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLiveIndexedEventsAggSum() bool {
	return o != nil && o.LiveIndexedEventsAggSum != nil
}

// SetLiveIndexedEventsAggSum gets a reference to the given int64 and assigns it to the LiveIndexedEventsAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetLiveIndexedEventsAggSum(v int64) {
	o.LiveIndexedEventsAggSum = &v
}

// GetLiveIngestedBytesAggSum returns the LiveIngestedBytesAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetLiveIngestedBytesAggSum() int64 {
	if o == nil || o.LiveIngestedBytesAggSum == nil {
		var ret int64
		return ret
	}
	return *o.LiveIngestedBytesAggSum
}

// GetLiveIngestedBytesAggSumOk returns a tuple with the LiveIngestedBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetLiveIngestedBytesAggSumOk() (*int64, bool) {
	if o == nil || o.LiveIngestedBytesAggSum == nil {
		return nil, false
	}
	return o.LiveIngestedBytesAggSum, true
}

// HasLiveIngestedBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLiveIngestedBytesAggSum() bool {
	return o != nil && o.LiveIngestedBytesAggSum != nil
}

// SetLiveIngestedBytesAggSum gets a reference to the given int64 and assigns it to the LiveIngestedBytesAggSum field.
func (o *UsageSummaryResponse) SetLiveIngestedBytesAggSum(v int64) {
	o.LiveIngestedBytesAggSum = &v
}

// GetLlmObservabilityAggSum returns the LlmObservabilityAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetLlmObservabilityAggSum() int64 {
	if o == nil || o.LlmObservabilityAggSum == nil {
		var ret int64
		return ret
	}
	return *o.LlmObservabilityAggSum
}

// GetLlmObservabilityAggSumOk returns a tuple with the LlmObservabilityAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetLlmObservabilityAggSumOk() (*int64, bool) {
	if o == nil || o.LlmObservabilityAggSum == nil {
		return nil, false
	}
	return o.LlmObservabilityAggSum, true
}

// HasLlmObservabilityAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLlmObservabilityAggSum() bool {
	return o != nil && o.LlmObservabilityAggSum != nil
}

// SetLlmObservabilityAggSum gets a reference to the given int64 and assigns it to the LlmObservabilityAggSum field.
func (o *UsageSummaryResponse) SetLlmObservabilityAggSum(v int64) {
	o.LlmObservabilityAggSum = &v
}

// GetLlmObservabilityMinSpendAggSum returns the LlmObservabilityMinSpendAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetLlmObservabilityMinSpendAggSum() int64 {
	if o == nil || o.LlmObservabilityMinSpendAggSum == nil {
		var ret int64
		return ret
	}
	return *o.LlmObservabilityMinSpendAggSum
}

// GetLlmObservabilityMinSpendAggSumOk returns a tuple with the LlmObservabilityMinSpendAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetLlmObservabilityMinSpendAggSumOk() (*int64, bool) {
	if o == nil || o.LlmObservabilityMinSpendAggSum == nil {
		return nil, false
	}
	return o.LlmObservabilityMinSpendAggSum, true
}

// HasLlmObservabilityMinSpendAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLlmObservabilityMinSpendAggSum() bool {
	return o != nil && o.LlmObservabilityMinSpendAggSum != nil
}

// SetLlmObservabilityMinSpendAggSum gets a reference to the given int64 and assigns it to the LlmObservabilityMinSpendAggSum field.
func (o *UsageSummaryResponse) SetLlmObservabilityMinSpendAggSum(v int64) {
	o.LlmObservabilityMinSpendAggSum = &v
}

// GetLogsByRetention returns the LogsByRetention field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetLogsByRetention() LogsByRetention {
	if o == nil || o.LogsByRetention == nil {
		var ret LogsByRetention
		return ret
	}
	return *o.LogsByRetention
}

// GetLogsByRetentionOk returns a tuple with the LogsByRetention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetLogsByRetentionOk() (*LogsByRetention, bool) {
	if o == nil || o.LogsByRetention == nil {
		return nil, false
	}
	return o.LogsByRetention, true
}

// HasLogsByRetention returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasLogsByRetention() bool {
	return o != nil && o.LogsByRetention != nil
}

// SetLogsByRetention gets a reference to the given LogsByRetention and assigns it to the LogsByRetention field.
func (o *UsageSummaryResponse) SetLogsByRetention(v LogsByRetention) {
	o.LogsByRetention = &v
}

// GetMobileRumLiteSessionCountAggSum returns the MobileRumLiteSessionCountAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumLiteSessionCountAggSum() int64 {
	if o == nil || o.MobileRumLiteSessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumLiteSessionCountAggSum
}

// GetMobileRumLiteSessionCountAggSumOk returns a tuple with the MobileRumLiteSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumLiteSessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.MobileRumLiteSessionCountAggSum == nil {
		return nil, false
	}
	return o.MobileRumLiteSessionCountAggSum, true
}

// HasMobileRumLiteSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumLiteSessionCountAggSum() bool {
	return o != nil && o.MobileRumLiteSessionCountAggSum != nil
}

// SetMobileRumLiteSessionCountAggSum gets a reference to the given int64 and assigns it to the MobileRumLiteSessionCountAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetMobileRumLiteSessionCountAggSum(v int64) {
	o.MobileRumLiteSessionCountAggSum = &v
}

// GetMobileRumSessionCountAggSum returns the MobileRumSessionCountAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountAggSum() int64 {
	if o == nil || o.MobileRumSessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountAggSum
}

// GetMobileRumSessionCountAggSumOk returns a tuple with the MobileRumSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountAggSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountAggSum, true
}

// HasMobileRumSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountAggSum() bool {
	return o != nil && o.MobileRumSessionCountAggSum != nil
}

// SetMobileRumSessionCountAggSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetMobileRumSessionCountAggSum(v int64) {
	o.MobileRumSessionCountAggSum = &v
}

// GetMobileRumSessionCountAndroidAggSum returns the MobileRumSessionCountAndroidAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountAndroidAggSum() int64 {
	if o == nil || o.MobileRumSessionCountAndroidAggSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountAndroidAggSum
}

// GetMobileRumSessionCountAndroidAggSumOk returns a tuple with the MobileRumSessionCountAndroidAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountAndroidAggSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountAndroidAggSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountAndroidAggSum, true
}

// HasMobileRumSessionCountAndroidAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountAndroidAggSum() bool {
	return o != nil && o.MobileRumSessionCountAndroidAggSum != nil
}

// SetMobileRumSessionCountAndroidAggSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountAndroidAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetMobileRumSessionCountAndroidAggSum(v int64) {
	o.MobileRumSessionCountAndroidAggSum = &v
}

// GetMobileRumSessionCountFlutterAggSum returns the MobileRumSessionCountFlutterAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountFlutterAggSum() int64 {
	if o == nil || o.MobileRumSessionCountFlutterAggSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountFlutterAggSum
}

// GetMobileRumSessionCountFlutterAggSumOk returns a tuple with the MobileRumSessionCountFlutterAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountFlutterAggSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountFlutterAggSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountFlutterAggSum, true
}

// HasMobileRumSessionCountFlutterAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountFlutterAggSum() bool {
	return o != nil && o.MobileRumSessionCountFlutterAggSum != nil
}

// SetMobileRumSessionCountFlutterAggSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountFlutterAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetMobileRumSessionCountFlutterAggSum(v int64) {
	o.MobileRumSessionCountFlutterAggSum = &v
}

// GetMobileRumSessionCountIosAggSum returns the MobileRumSessionCountIosAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountIosAggSum() int64 {
	if o == nil || o.MobileRumSessionCountIosAggSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountIosAggSum
}

// GetMobileRumSessionCountIosAggSumOk returns a tuple with the MobileRumSessionCountIosAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountIosAggSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountIosAggSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountIosAggSum, true
}

// HasMobileRumSessionCountIosAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountIosAggSum() bool {
	return o != nil && o.MobileRumSessionCountIosAggSum != nil
}

// SetMobileRumSessionCountIosAggSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountIosAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetMobileRumSessionCountIosAggSum(v int64) {
	o.MobileRumSessionCountIosAggSum = &v
}

// GetMobileRumSessionCountReactnativeAggSum returns the MobileRumSessionCountReactnativeAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountReactnativeAggSum() int64 {
	if o == nil || o.MobileRumSessionCountReactnativeAggSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountReactnativeAggSum
}

// GetMobileRumSessionCountReactnativeAggSumOk returns a tuple with the MobileRumSessionCountReactnativeAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountReactnativeAggSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountReactnativeAggSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountReactnativeAggSum, true
}

// HasMobileRumSessionCountReactnativeAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountReactnativeAggSum() bool {
	return o != nil && o.MobileRumSessionCountReactnativeAggSum != nil
}

// SetMobileRumSessionCountReactnativeAggSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountReactnativeAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetMobileRumSessionCountReactnativeAggSum(v int64) {
	o.MobileRumSessionCountReactnativeAggSum = &v
}

// GetMobileRumSessionCountRokuAggSum returns the MobileRumSessionCountRokuAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountRokuAggSum() int64 {
	if o == nil || o.MobileRumSessionCountRokuAggSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountRokuAggSum
}

// GetMobileRumSessionCountRokuAggSumOk returns a tuple with the MobileRumSessionCountRokuAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumSessionCountRokuAggSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountRokuAggSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountRokuAggSum, true
}

// HasMobileRumSessionCountRokuAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumSessionCountRokuAggSum() bool {
	return o != nil && o.MobileRumSessionCountRokuAggSum != nil
}

// SetMobileRumSessionCountRokuAggSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountRokuAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetMobileRumSessionCountRokuAggSum(v int64) {
	o.MobileRumSessionCountRokuAggSum = &v
}

// GetMobileRumUnitsAggSum returns the MobileRumUnitsAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumUnitsAggSum() int64 {
	if o == nil || o.MobileRumUnitsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumUnitsAggSum
}

// GetMobileRumUnitsAggSumOk returns a tuple with the MobileRumUnitsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetMobileRumUnitsAggSumOk() (*int64, bool) {
	if o == nil || o.MobileRumUnitsAggSum == nil {
		return nil, false
	}
	return o.MobileRumUnitsAggSum, true
}

// HasMobileRumUnitsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasMobileRumUnitsAggSum() bool {
	return o != nil && o.MobileRumUnitsAggSum != nil
}

// SetMobileRumUnitsAggSum gets a reference to the given int64 and assigns it to the MobileRumUnitsAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetMobileRumUnitsAggSum(v int64) {
	o.MobileRumUnitsAggSum = &v
}

// GetNdmNetflowEventsAggSum returns the NdmNetflowEventsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetNdmNetflowEventsAggSum() int64 {
	if o == nil || o.NdmNetflowEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.NdmNetflowEventsAggSum
}

// GetNdmNetflowEventsAggSumOk returns a tuple with the NdmNetflowEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetNdmNetflowEventsAggSumOk() (*int64, bool) {
	if o == nil || o.NdmNetflowEventsAggSum == nil {
		return nil, false
	}
	return o.NdmNetflowEventsAggSum, true
}

// HasNdmNetflowEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasNdmNetflowEventsAggSum() bool {
	return o != nil && o.NdmNetflowEventsAggSum != nil
}

// SetNdmNetflowEventsAggSum gets a reference to the given int64 and assigns it to the NdmNetflowEventsAggSum field.
func (o *UsageSummaryResponse) SetNdmNetflowEventsAggSum(v int64) {
	o.NdmNetflowEventsAggSum = &v
}

// GetNetflowIndexedEventsCountAggSum returns the NetflowIndexedEventsCountAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetNetflowIndexedEventsCountAggSum() int64 {
	if o == nil || o.NetflowIndexedEventsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.NetflowIndexedEventsCountAggSum
}

// GetNetflowIndexedEventsCountAggSumOk returns a tuple with the NetflowIndexedEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetNetflowIndexedEventsCountAggSumOk() (*int64, bool) {
	if o == nil || o.NetflowIndexedEventsCountAggSum == nil {
		return nil, false
	}
	return o.NetflowIndexedEventsCountAggSum, true
}

// HasNetflowIndexedEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasNetflowIndexedEventsCountAggSum() bool {
	return o != nil && o.NetflowIndexedEventsCountAggSum != nil
}

// SetNetflowIndexedEventsCountAggSum gets a reference to the given int64 and assigns it to the NetflowIndexedEventsCountAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetNetflowIndexedEventsCountAggSum(v int64) {
	o.NetflowIndexedEventsCountAggSum = &v
}

// GetNetworkDeviceWirelessTop99pSum returns the NetworkDeviceWirelessTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetNetworkDeviceWirelessTop99pSum() int64 {
	if o == nil || o.NetworkDeviceWirelessTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.NetworkDeviceWirelessTop99pSum
}

// GetNetworkDeviceWirelessTop99pSumOk returns a tuple with the NetworkDeviceWirelessTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetNetworkDeviceWirelessTop99pSumOk() (*int64, bool) {
	if o == nil || o.NetworkDeviceWirelessTop99pSum == nil {
		return nil, false
	}
	return o.NetworkDeviceWirelessTop99pSum, true
}

// HasNetworkDeviceWirelessTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasNetworkDeviceWirelessTop99pSum() bool {
	return o != nil && o.NetworkDeviceWirelessTop99pSum != nil
}

// SetNetworkDeviceWirelessTop99pSum gets a reference to the given int64 and assigns it to the NetworkDeviceWirelessTop99pSum field.
func (o *UsageSummaryResponse) SetNetworkDeviceWirelessTop99pSum(v int64) {
	o.NetworkDeviceWirelessTop99pSum = &v
}

// GetNpmHostTop99pSum returns the NpmHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetNpmHostTop99pSum() int64 {
	if o == nil || o.NpmHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.NpmHostTop99pSum
}

// GetNpmHostTop99pSumOk returns a tuple with the NpmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetNpmHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.NpmHostTop99pSum == nil {
		return nil, false
	}
	return o.NpmHostTop99pSum, true
}

// HasNpmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasNpmHostTop99pSum() bool {
	return o != nil && o.NpmHostTop99pSum != nil
}

// SetNpmHostTop99pSum gets a reference to the given int64 and assigns it to the NpmHostTop99pSum field.
func (o *UsageSummaryResponse) SetNpmHostTop99pSum(v int64) {
	o.NpmHostTop99pSum = &v
}

// GetObservabilityPipelinesBytesProcessedAggSum returns the ObservabilityPipelinesBytesProcessedAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetObservabilityPipelinesBytesProcessedAggSum() int64 {
	if o == nil || o.ObservabilityPipelinesBytesProcessedAggSum == nil {
		var ret int64
		return ret
	}
	return *o.ObservabilityPipelinesBytesProcessedAggSum
}

// GetObservabilityPipelinesBytesProcessedAggSumOk returns a tuple with the ObservabilityPipelinesBytesProcessedAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetObservabilityPipelinesBytesProcessedAggSumOk() (*int64, bool) {
	if o == nil || o.ObservabilityPipelinesBytesProcessedAggSum == nil {
		return nil, false
	}
	return o.ObservabilityPipelinesBytesProcessedAggSum, true
}

// HasObservabilityPipelinesBytesProcessedAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasObservabilityPipelinesBytesProcessedAggSum() bool {
	return o != nil && o.ObservabilityPipelinesBytesProcessedAggSum != nil
}

// SetObservabilityPipelinesBytesProcessedAggSum gets a reference to the given int64 and assigns it to the ObservabilityPipelinesBytesProcessedAggSum field.
func (o *UsageSummaryResponse) SetObservabilityPipelinesBytesProcessedAggSum(v int64) {
	o.ObservabilityPipelinesBytesProcessedAggSum = &v
}

// GetOciHostAggSum returns the OciHostAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetOciHostAggSum() int64 {
	if o == nil || o.OciHostAggSum == nil {
		var ret int64
		return ret
	}
	return *o.OciHostAggSum
}

// GetOciHostAggSumOk returns a tuple with the OciHostAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetOciHostAggSumOk() (*int64, bool) {
	if o == nil || o.OciHostAggSum == nil {
		return nil, false
	}
	return o.OciHostAggSum, true
}

// HasOciHostAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasOciHostAggSum() bool {
	return o != nil && o.OciHostAggSum != nil
}

// SetOciHostAggSum gets a reference to the given int64 and assigns it to the OciHostAggSum field.
func (o *UsageSummaryResponse) SetOciHostAggSum(v int64) {
	o.OciHostAggSum = &v
}

// GetOciHostTop99pSum returns the OciHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetOciHostTop99pSum() int64 {
	if o == nil || o.OciHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.OciHostTop99pSum
}

// GetOciHostTop99pSumOk returns a tuple with the OciHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetOciHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.OciHostTop99pSum == nil {
		return nil, false
	}
	return o.OciHostTop99pSum, true
}

// HasOciHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasOciHostTop99pSum() bool {
	return o != nil && o.OciHostTop99pSum != nil
}

// SetOciHostTop99pSum gets a reference to the given int64 and assigns it to the OciHostTop99pSum field.
func (o *UsageSummaryResponse) SetOciHostTop99pSum(v int64) {
	o.OciHostTop99pSum = &v
}

// GetOnCallSeatHwmSum returns the OnCallSeatHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetOnCallSeatHwmSum() int64 {
	if o == nil || o.OnCallSeatHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.OnCallSeatHwmSum
}

// GetOnCallSeatHwmSumOk returns a tuple with the OnCallSeatHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetOnCallSeatHwmSumOk() (*int64, bool) {
	if o == nil || o.OnCallSeatHwmSum == nil {
		return nil, false
	}
	return o.OnCallSeatHwmSum, true
}

// HasOnCallSeatHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasOnCallSeatHwmSum() bool {
	return o != nil && o.OnCallSeatHwmSum != nil
}

// SetOnCallSeatHwmSum gets a reference to the given int64 and assigns it to the OnCallSeatHwmSum field.
func (o *UsageSummaryResponse) SetOnCallSeatHwmSum(v int64) {
	o.OnCallSeatHwmSum = &v
}

// GetOnlineArchiveEventsCountAggSum returns the OnlineArchiveEventsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetOnlineArchiveEventsCountAggSum() int64 {
	if o == nil || o.OnlineArchiveEventsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.OnlineArchiveEventsCountAggSum
}

// GetOnlineArchiveEventsCountAggSumOk returns a tuple with the OnlineArchiveEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetOnlineArchiveEventsCountAggSumOk() (*int64, bool) {
	if o == nil || o.OnlineArchiveEventsCountAggSum == nil {
		return nil, false
	}
	return o.OnlineArchiveEventsCountAggSum, true
}

// HasOnlineArchiveEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasOnlineArchiveEventsCountAggSum() bool {
	return o != nil && o.OnlineArchiveEventsCountAggSum != nil
}

// SetOnlineArchiveEventsCountAggSum gets a reference to the given int64 and assigns it to the OnlineArchiveEventsCountAggSum field.
func (o *UsageSummaryResponse) SetOnlineArchiveEventsCountAggSum(v int64) {
	o.OnlineArchiveEventsCountAggSum = &v
}

// GetOpentelemetryApmHostTop99pSum returns the OpentelemetryApmHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetOpentelemetryApmHostTop99pSum() int64 {
	if o == nil || o.OpentelemetryApmHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryApmHostTop99pSum
}

// GetOpentelemetryApmHostTop99pSumOk returns a tuple with the OpentelemetryApmHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetOpentelemetryApmHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.OpentelemetryApmHostTop99pSum == nil {
		return nil, false
	}
	return o.OpentelemetryApmHostTop99pSum, true
}

// HasOpentelemetryApmHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasOpentelemetryApmHostTop99pSum() bool {
	return o != nil && o.OpentelemetryApmHostTop99pSum != nil
}

// SetOpentelemetryApmHostTop99pSum gets a reference to the given int64 and assigns it to the OpentelemetryApmHostTop99pSum field.
func (o *UsageSummaryResponse) SetOpentelemetryApmHostTop99pSum(v int64) {
	o.OpentelemetryApmHostTop99pSum = &v
}

// GetOpentelemetryHostTop99pSum returns the OpentelemetryHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetOpentelemetryHostTop99pSum() int64 {
	if o == nil || o.OpentelemetryHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryHostTop99pSum
}

// GetOpentelemetryHostTop99pSumOk returns a tuple with the OpentelemetryHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetOpentelemetryHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.OpentelemetryHostTop99pSum == nil {
		return nil, false
	}
	return o.OpentelemetryHostTop99pSum, true
}

// HasOpentelemetryHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasOpentelemetryHostTop99pSum() bool {
	return o != nil && o.OpentelemetryHostTop99pSum != nil
}

// SetOpentelemetryHostTop99pSum gets a reference to the given int64 and assigns it to the OpentelemetryHostTop99pSum field.
func (o *UsageSummaryResponse) SetOpentelemetryHostTop99pSum(v int64) {
	o.OpentelemetryHostTop99pSum = &v
}

// GetProductAnalyticsAggSum returns the ProductAnalyticsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetProductAnalyticsAggSum() int64 {
	if o == nil || o.ProductAnalyticsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.ProductAnalyticsAggSum
}

// GetProductAnalyticsAggSumOk returns a tuple with the ProductAnalyticsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetProductAnalyticsAggSumOk() (*int64, bool) {
	if o == nil || o.ProductAnalyticsAggSum == nil {
		return nil, false
	}
	return o.ProductAnalyticsAggSum, true
}

// HasProductAnalyticsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasProductAnalyticsAggSum() bool {
	return o != nil && o.ProductAnalyticsAggSum != nil
}

// SetProductAnalyticsAggSum gets a reference to the given int64 and assigns it to the ProductAnalyticsAggSum field.
func (o *UsageSummaryResponse) SetProductAnalyticsAggSum(v int64) {
	o.ProductAnalyticsAggSum = &v
}

// GetProfilingAasCountTop99pSum returns the ProfilingAasCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetProfilingAasCountTop99pSum() int64 {
	if o == nil || o.ProfilingAasCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.ProfilingAasCountTop99pSum
}

// GetProfilingAasCountTop99pSumOk returns a tuple with the ProfilingAasCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetProfilingAasCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.ProfilingAasCountTop99pSum == nil {
		return nil, false
	}
	return o.ProfilingAasCountTop99pSum, true
}

// HasProfilingAasCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasProfilingAasCountTop99pSum() bool {
	return o != nil && o.ProfilingAasCountTop99pSum != nil
}

// SetProfilingAasCountTop99pSum gets a reference to the given int64 and assigns it to the ProfilingAasCountTop99pSum field.
func (o *UsageSummaryResponse) SetProfilingAasCountTop99pSum(v int64) {
	o.ProfilingAasCountTop99pSum = &v
}

// GetProfilingContainerAgentCountAvg returns the ProfilingContainerAgentCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetProfilingContainerAgentCountAvg() int64 {
	if o == nil || o.ProfilingContainerAgentCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.ProfilingContainerAgentCountAvg
}

// GetProfilingContainerAgentCountAvgOk returns a tuple with the ProfilingContainerAgentCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetProfilingContainerAgentCountAvgOk() (*int64, bool) {
	if o == nil || o.ProfilingContainerAgentCountAvg == nil {
		return nil, false
	}
	return o.ProfilingContainerAgentCountAvg, true
}

// HasProfilingContainerAgentCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasProfilingContainerAgentCountAvg() bool {
	return o != nil && o.ProfilingContainerAgentCountAvg != nil
}

// SetProfilingContainerAgentCountAvg gets a reference to the given int64 and assigns it to the ProfilingContainerAgentCountAvg field.
func (o *UsageSummaryResponse) SetProfilingContainerAgentCountAvg(v int64) {
	o.ProfilingContainerAgentCountAvg = &v
}

// GetProfilingHostCountTop99pSum returns the ProfilingHostCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetProfilingHostCountTop99pSum() int64 {
	if o == nil || o.ProfilingHostCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.ProfilingHostCountTop99pSum
}

// GetProfilingHostCountTop99pSumOk returns a tuple with the ProfilingHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetProfilingHostCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.ProfilingHostCountTop99pSum == nil {
		return nil, false
	}
	return o.ProfilingHostCountTop99pSum, true
}

// HasProfilingHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasProfilingHostCountTop99pSum() bool {
	return o != nil && o.ProfilingHostCountTop99pSum != nil
}

// SetProfilingHostCountTop99pSum gets a reference to the given int64 and assigns it to the ProfilingHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetProfilingHostCountTop99pSum(v int64) {
	o.ProfilingHostCountTop99pSum = &v
}

// GetProxmoxHostAggSum returns the ProxmoxHostAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetProxmoxHostAggSum() int64 {
	if o == nil || o.ProxmoxHostAggSum == nil {
		var ret int64
		return ret
	}
	return *o.ProxmoxHostAggSum
}

// GetProxmoxHostAggSumOk returns a tuple with the ProxmoxHostAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetProxmoxHostAggSumOk() (*int64, bool) {
	if o == nil || o.ProxmoxHostAggSum == nil {
		return nil, false
	}
	return o.ProxmoxHostAggSum, true
}

// HasProxmoxHostAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasProxmoxHostAggSum() bool {
	return o != nil && o.ProxmoxHostAggSum != nil
}

// SetProxmoxHostAggSum gets a reference to the given int64 and assigns it to the ProxmoxHostAggSum field.
func (o *UsageSummaryResponse) SetProxmoxHostAggSum(v int64) {
	o.ProxmoxHostAggSum = &v
}

// GetProxmoxHostTop99pSum returns the ProxmoxHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetProxmoxHostTop99pSum() int64 {
	if o == nil || o.ProxmoxHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.ProxmoxHostTop99pSum
}

// GetProxmoxHostTop99pSumOk returns a tuple with the ProxmoxHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetProxmoxHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.ProxmoxHostTop99pSum == nil {
		return nil, false
	}
	return o.ProxmoxHostTop99pSum, true
}

// HasProxmoxHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasProxmoxHostTop99pSum() bool {
	return o != nil && o.ProxmoxHostTop99pSum != nil
}

// SetProxmoxHostTop99pSum gets a reference to the given int64 and assigns it to the ProxmoxHostTop99pSum field.
func (o *UsageSummaryResponse) SetProxmoxHostTop99pSum(v int64) {
	o.ProxmoxHostTop99pSum = &v
}

// GetPublishedAppHwmSum returns the PublishedAppHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetPublishedAppHwmSum() int64 {
	if o == nil || o.PublishedAppHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.PublishedAppHwmSum
}

// GetPublishedAppHwmSumOk returns a tuple with the PublishedAppHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetPublishedAppHwmSumOk() (*int64, bool) {
	if o == nil || o.PublishedAppHwmSum == nil {
		return nil, false
	}
	return o.PublishedAppHwmSum, true
}

// HasPublishedAppHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasPublishedAppHwmSum() bool {
	return o != nil && o.PublishedAppHwmSum != nil
}

// SetPublishedAppHwmSum gets a reference to the given int64 and assigns it to the PublishedAppHwmSum field.
func (o *UsageSummaryResponse) SetPublishedAppHwmSum(v int64) {
	o.PublishedAppHwmSum = &v
}

// GetRehydratedIndexedEventsAggSum returns the RehydratedIndexedEventsAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetRehydratedIndexedEventsAggSum() int64 {
	if o == nil || o.RehydratedIndexedEventsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RehydratedIndexedEventsAggSum
}

// GetRehydratedIndexedEventsAggSumOk returns a tuple with the RehydratedIndexedEventsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetRehydratedIndexedEventsAggSumOk() (*int64, bool) {
	if o == nil || o.RehydratedIndexedEventsAggSum == nil {
		return nil, false
	}
	return o.RehydratedIndexedEventsAggSum, true
}

// HasRehydratedIndexedEventsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRehydratedIndexedEventsAggSum() bool {
	return o != nil && o.RehydratedIndexedEventsAggSum != nil
}

// SetRehydratedIndexedEventsAggSum gets a reference to the given int64 and assigns it to the RehydratedIndexedEventsAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetRehydratedIndexedEventsAggSum(v int64) {
	o.RehydratedIndexedEventsAggSum = &v
}

// GetRehydratedIngestedBytesAggSum returns the RehydratedIngestedBytesAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRehydratedIngestedBytesAggSum() int64 {
	if o == nil || o.RehydratedIngestedBytesAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RehydratedIngestedBytesAggSum
}

// GetRehydratedIngestedBytesAggSumOk returns a tuple with the RehydratedIngestedBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRehydratedIngestedBytesAggSumOk() (*int64, bool) {
	if o == nil || o.RehydratedIngestedBytesAggSum == nil {
		return nil, false
	}
	return o.RehydratedIngestedBytesAggSum, true
}

// HasRehydratedIngestedBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRehydratedIngestedBytesAggSum() bool {
	return o != nil && o.RehydratedIngestedBytesAggSum != nil
}

// SetRehydratedIngestedBytesAggSum gets a reference to the given int64 and assigns it to the RehydratedIngestedBytesAggSum field.
func (o *UsageSummaryResponse) SetRehydratedIngestedBytesAggSum(v int64) {
	o.RehydratedIngestedBytesAggSum = &v
}

// GetRumBrowserAndMobileSessionCount returns the RumBrowserAndMobileSessionCount field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumBrowserAndMobileSessionCount() int64 {
	if o == nil || o.RumBrowserAndMobileSessionCount == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserAndMobileSessionCount
}

// GetRumBrowserAndMobileSessionCountOk returns a tuple with the RumBrowserAndMobileSessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumBrowserAndMobileSessionCountOk() (*int64, bool) {
	if o == nil || o.RumBrowserAndMobileSessionCount == nil {
		return nil, false
	}
	return o.RumBrowserAndMobileSessionCount, true
}

// HasRumBrowserAndMobileSessionCount returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumBrowserAndMobileSessionCount() bool {
	return o != nil && o.RumBrowserAndMobileSessionCount != nil
}

// SetRumBrowserAndMobileSessionCount gets a reference to the given int64 and assigns it to the RumBrowserAndMobileSessionCount field.
func (o *UsageSummaryResponse) SetRumBrowserAndMobileSessionCount(v int64) {
	o.RumBrowserAndMobileSessionCount = &v
}

// GetRumBrowserLegacySessionCountAggSum returns the RumBrowserLegacySessionCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumBrowserLegacySessionCountAggSum() int64 {
	if o == nil || o.RumBrowserLegacySessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserLegacySessionCountAggSum
}

// GetRumBrowserLegacySessionCountAggSumOk returns a tuple with the RumBrowserLegacySessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumBrowserLegacySessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.RumBrowserLegacySessionCountAggSum == nil {
		return nil, false
	}
	return o.RumBrowserLegacySessionCountAggSum, true
}

// HasRumBrowserLegacySessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumBrowserLegacySessionCountAggSum() bool {
	return o != nil && o.RumBrowserLegacySessionCountAggSum != nil
}

// SetRumBrowserLegacySessionCountAggSum gets a reference to the given int64 and assigns it to the RumBrowserLegacySessionCountAggSum field.
func (o *UsageSummaryResponse) SetRumBrowserLegacySessionCountAggSum(v int64) {
	o.RumBrowserLegacySessionCountAggSum = &v
}

// GetRumBrowserLiteSessionCountAggSum returns the RumBrowserLiteSessionCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumBrowserLiteSessionCountAggSum() int64 {
	if o == nil || o.RumBrowserLiteSessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserLiteSessionCountAggSum
}

// GetRumBrowserLiteSessionCountAggSumOk returns a tuple with the RumBrowserLiteSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumBrowserLiteSessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.RumBrowserLiteSessionCountAggSum == nil {
		return nil, false
	}
	return o.RumBrowserLiteSessionCountAggSum, true
}

// HasRumBrowserLiteSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumBrowserLiteSessionCountAggSum() bool {
	return o != nil && o.RumBrowserLiteSessionCountAggSum != nil
}

// SetRumBrowserLiteSessionCountAggSum gets a reference to the given int64 and assigns it to the RumBrowserLiteSessionCountAggSum field.
func (o *UsageSummaryResponse) SetRumBrowserLiteSessionCountAggSum(v int64) {
	o.RumBrowserLiteSessionCountAggSum = &v
}

// GetRumBrowserReplaySessionCountAggSum returns the RumBrowserReplaySessionCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumBrowserReplaySessionCountAggSum() int64 {
	if o == nil || o.RumBrowserReplaySessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserReplaySessionCountAggSum
}

// GetRumBrowserReplaySessionCountAggSumOk returns a tuple with the RumBrowserReplaySessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumBrowserReplaySessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.RumBrowserReplaySessionCountAggSum == nil {
		return nil, false
	}
	return o.RumBrowserReplaySessionCountAggSum, true
}

// HasRumBrowserReplaySessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumBrowserReplaySessionCountAggSum() bool {
	return o != nil && o.RumBrowserReplaySessionCountAggSum != nil
}

// SetRumBrowserReplaySessionCountAggSum gets a reference to the given int64 and assigns it to the RumBrowserReplaySessionCountAggSum field.
func (o *UsageSummaryResponse) SetRumBrowserReplaySessionCountAggSum(v int64) {
	o.RumBrowserReplaySessionCountAggSum = &v
}

// GetRumIndexedSessionsAggSum returns the RumIndexedSessionsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumIndexedSessionsAggSum() int64 {
	if o == nil || o.RumIndexedSessionsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumIndexedSessionsAggSum
}

// GetRumIndexedSessionsAggSumOk returns a tuple with the RumIndexedSessionsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumIndexedSessionsAggSumOk() (*int64, bool) {
	if o == nil || o.RumIndexedSessionsAggSum == nil {
		return nil, false
	}
	return o.RumIndexedSessionsAggSum, true
}

// HasRumIndexedSessionsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumIndexedSessionsAggSum() bool {
	return o != nil && o.RumIndexedSessionsAggSum != nil
}

// SetRumIndexedSessionsAggSum gets a reference to the given int64 and assigns it to the RumIndexedSessionsAggSum field.
func (o *UsageSummaryResponse) SetRumIndexedSessionsAggSum(v int64) {
	o.RumIndexedSessionsAggSum = &v
}

// GetRumIngestedSessionsAggSum returns the RumIngestedSessionsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumIngestedSessionsAggSum() int64 {
	if o == nil || o.RumIngestedSessionsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumIngestedSessionsAggSum
}

// GetRumIngestedSessionsAggSumOk returns a tuple with the RumIngestedSessionsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumIngestedSessionsAggSumOk() (*int64, bool) {
	if o == nil || o.RumIngestedSessionsAggSum == nil {
		return nil, false
	}
	return o.RumIngestedSessionsAggSum, true
}

// HasRumIngestedSessionsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumIngestedSessionsAggSum() bool {
	return o != nil && o.RumIngestedSessionsAggSum != nil
}

// SetRumIngestedSessionsAggSum gets a reference to the given int64 and assigns it to the RumIngestedSessionsAggSum field.
func (o *UsageSummaryResponse) SetRumIngestedSessionsAggSum(v int64) {
	o.RumIngestedSessionsAggSum = &v
}

// GetRumLiteSessionCountAggSum returns the RumLiteSessionCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumLiteSessionCountAggSum() int64 {
	if o == nil || o.RumLiteSessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumLiteSessionCountAggSum
}

// GetRumLiteSessionCountAggSumOk returns a tuple with the RumLiteSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumLiteSessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.RumLiteSessionCountAggSum == nil {
		return nil, false
	}
	return o.RumLiteSessionCountAggSum, true
}

// HasRumLiteSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumLiteSessionCountAggSum() bool {
	return o != nil && o.RumLiteSessionCountAggSum != nil
}

// SetRumLiteSessionCountAggSum gets a reference to the given int64 and assigns it to the RumLiteSessionCountAggSum field.
func (o *UsageSummaryResponse) SetRumLiteSessionCountAggSum(v int64) {
	o.RumLiteSessionCountAggSum = &v
}

// GetRumMobileLegacySessionCountAndroidAggSum returns the RumMobileLegacySessionCountAndroidAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountAndroidAggSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountAndroidAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountAndroidAggSum
}

// GetRumMobileLegacySessionCountAndroidAggSumOk returns a tuple with the RumMobileLegacySessionCountAndroidAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountAndroidAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountAndroidAggSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountAndroidAggSum, true
}

// HasRumMobileLegacySessionCountAndroidAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLegacySessionCountAndroidAggSum() bool {
	return o != nil && o.RumMobileLegacySessionCountAndroidAggSum != nil
}

// SetRumMobileLegacySessionCountAndroidAggSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountAndroidAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLegacySessionCountAndroidAggSum(v int64) {
	o.RumMobileLegacySessionCountAndroidAggSum = &v
}

// GetRumMobileLegacySessionCountFlutterAggSum returns the RumMobileLegacySessionCountFlutterAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountFlutterAggSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountFlutterAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountFlutterAggSum
}

// GetRumMobileLegacySessionCountFlutterAggSumOk returns a tuple with the RumMobileLegacySessionCountFlutterAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountFlutterAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountFlutterAggSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountFlutterAggSum, true
}

// HasRumMobileLegacySessionCountFlutterAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLegacySessionCountFlutterAggSum() bool {
	return o != nil && o.RumMobileLegacySessionCountFlutterAggSum != nil
}

// SetRumMobileLegacySessionCountFlutterAggSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountFlutterAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLegacySessionCountFlutterAggSum(v int64) {
	o.RumMobileLegacySessionCountFlutterAggSum = &v
}

// GetRumMobileLegacySessionCountIosAggSum returns the RumMobileLegacySessionCountIosAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountIosAggSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountIosAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountIosAggSum
}

// GetRumMobileLegacySessionCountIosAggSumOk returns a tuple with the RumMobileLegacySessionCountIosAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountIosAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountIosAggSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountIosAggSum, true
}

// HasRumMobileLegacySessionCountIosAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLegacySessionCountIosAggSum() bool {
	return o != nil && o.RumMobileLegacySessionCountIosAggSum != nil
}

// SetRumMobileLegacySessionCountIosAggSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountIosAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLegacySessionCountIosAggSum(v int64) {
	o.RumMobileLegacySessionCountIosAggSum = &v
}

// GetRumMobileLegacySessionCountReactnativeAggSum returns the RumMobileLegacySessionCountReactnativeAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountReactnativeAggSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountReactnativeAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountReactnativeAggSum
}

// GetRumMobileLegacySessionCountReactnativeAggSumOk returns a tuple with the RumMobileLegacySessionCountReactnativeAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountReactnativeAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountReactnativeAggSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountReactnativeAggSum, true
}

// HasRumMobileLegacySessionCountReactnativeAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLegacySessionCountReactnativeAggSum() bool {
	return o != nil && o.RumMobileLegacySessionCountReactnativeAggSum != nil
}

// SetRumMobileLegacySessionCountReactnativeAggSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountReactnativeAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLegacySessionCountReactnativeAggSum(v int64) {
	o.RumMobileLegacySessionCountReactnativeAggSum = &v
}

// GetRumMobileLegacySessionCountRokuAggSum returns the RumMobileLegacySessionCountRokuAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountRokuAggSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountRokuAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountRokuAggSum
}

// GetRumMobileLegacySessionCountRokuAggSumOk returns a tuple with the RumMobileLegacySessionCountRokuAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLegacySessionCountRokuAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountRokuAggSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountRokuAggSum, true
}

// HasRumMobileLegacySessionCountRokuAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLegacySessionCountRokuAggSum() bool {
	return o != nil && o.RumMobileLegacySessionCountRokuAggSum != nil
}

// SetRumMobileLegacySessionCountRokuAggSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountRokuAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLegacySessionCountRokuAggSum(v int64) {
	o.RumMobileLegacySessionCountRokuAggSum = &v
}

// GetRumMobileLiteSessionCountAndroidAggSum returns the RumMobileLiteSessionCountAndroidAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountAndroidAggSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountAndroidAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountAndroidAggSum
}

// GetRumMobileLiteSessionCountAndroidAggSumOk returns a tuple with the RumMobileLiteSessionCountAndroidAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountAndroidAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountAndroidAggSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountAndroidAggSum, true
}

// HasRumMobileLiteSessionCountAndroidAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLiteSessionCountAndroidAggSum() bool {
	return o != nil && o.RumMobileLiteSessionCountAndroidAggSum != nil
}

// SetRumMobileLiteSessionCountAndroidAggSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountAndroidAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLiteSessionCountAndroidAggSum(v int64) {
	o.RumMobileLiteSessionCountAndroidAggSum = &v
}

// GetRumMobileLiteSessionCountFlutterAggSum returns the RumMobileLiteSessionCountFlutterAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountFlutterAggSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountFlutterAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountFlutterAggSum
}

// GetRumMobileLiteSessionCountFlutterAggSumOk returns a tuple with the RumMobileLiteSessionCountFlutterAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountFlutterAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountFlutterAggSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountFlutterAggSum, true
}

// HasRumMobileLiteSessionCountFlutterAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLiteSessionCountFlutterAggSum() bool {
	return o != nil && o.RumMobileLiteSessionCountFlutterAggSum != nil
}

// SetRumMobileLiteSessionCountFlutterAggSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountFlutterAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLiteSessionCountFlutterAggSum(v int64) {
	o.RumMobileLiteSessionCountFlutterAggSum = &v
}

// GetRumMobileLiteSessionCountIosAggSum returns the RumMobileLiteSessionCountIosAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountIosAggSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountIosAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountIosAggSum
}

// GetRumMobileLiteSessionCountIosAggSumOk returns a tuple with the RumMobileLiteSessionCountIosAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountIosAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountIosAggSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountIosAggSum, true
}

// HasRumMobileLiteSessionCountIosAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLiteSessionCountIosAggSum() bool {
	return o != nil && o.RumMobileLiteSessionCountIosAggSum != nil
}

// SetRumMobileLiteSessionCountIosAggSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountIosAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLiteSessionCountIosAggSum(v int64) {
	o.RumMobileLiteSessionCountIosAggSum = &v
}

// GetRumMobileLiteSessionCountKotlinmultiplatformAggSum returns the RumMobileLiteSessionCountKotlinmultiplatformAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountKotlinmultiplatformAggSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountKotlinmultiplatformAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountKotlinmultiplatformAggSum
}

// GetRumMobileLiteSessionCountKotlinmultiplatformAggSumOk returns a tuple with the RumMobileLiteSessionCountKotlinmultiplatformAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountKotlinmultiplatformAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountKotlinmultiplatformAggSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountKotlinmultiplatformAggSum, true
}

// HasRumMobileLiteSessionCountKotlinmultiplatformAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLiteSessionCountKotlinmultiplatformAggSum() bool {
	return o != nil && o.RumMobileLiteSessionCountKotlinmultiplatformAggSum != nil
}

// SetRumMobileLiteSessionCountKotlinmultiplatformAggSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountKotlinmultiplatformAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLiteSessionCountKotlinmultiplatformAggSum(v int64) {
	o.RumMobileLiteSessionCountKotlinmultiplatformAggSum = &v
}

// GetRumMobileLiteSessionCountReactnativeAggSum returns the RumMobileLiteSessionCountReactnativeAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountReactnativeAggSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountReactnativeAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountReactnativeAggSum
}

// GetRumMobileLiteSessionCountReactnativeAggSumOk returns a tuple with the RumMobileLiteSessionCountReactnativeAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountReactnativeAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountReactnativeAggSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountReactnativeAggSum, true
}

// HasRumMobileLiteSessionCountReactnativeAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLiteSessionCountReactnativeAggSum() bool {
	return o != nil && o.RumMobileLiteSessionCountReactnativeAggSum != nil
}

// SetRumMobileLiteSessionCountReactnativeAggSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountReactnativeAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLiteSessionCountReactnativeAggSum(v int64) {
	o.RumMobileLiteSessionCountReactnativeAggSum = &v
}

// GetRumMobileLiteSessionCountRokuAggSum returns the RumMobileLiteSessionCountRokuAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountRokuAggSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountRokuAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountRokuAggSum
}

// GetRumMobileLiteSessionCountRokuAggSumOk returns a tuple with the RumMobileLiteSessionCountRokuAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountRokuAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountRokuAggSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountRokuAggSum, true
}

// HasRumMobileLiteSessionCountRokuAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLiteSessionCountRokuAggSum() bool {
	return o != nil && o.RumMobileLiteSessionCountRokuAggSum != nil
}

// SetRumMobileLiteSessionCountRokuAggSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountRokuAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLiteSessionCountRokuAggSum(v int64) {
	o.RumMobileLiteSessionCountRokuAggSum = &v
}

// GetRumMobileLiteSessionCountUnityAggSum returns the RumMobileLiteSessionCountUnityAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountUnityAggSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountUnityAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountUnityAggSum
}

// GetRumMobileLiteSessionCountUnityAggSumOk returns a tuple with the RumMobileLiteSessionCountUnityAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileLiteSessionCountUnityAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountUnityAggSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountUnityAggSum, true
}

// HasRumMobileLiteSessionCountUnityAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileLiteSessionCountUnityAggSum() bool {
	return o != nil && o.RumMobileLiteSessionCountUnityAggSum != nil
}

// SetRumMobileLiteSessionCountUnityAggSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountUnityAggSum field.
func (o *UsageSummaryResponse) SetRumMobileLiteSessionCountUnityAggSum(v int64) {
	o.RumMobileLiteSessionCountUnityAggSum = &v
}

// GetRumMobileReplaySessionCountAndroidAggSum returns the RumMobileReplaySessionCountAndroidAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileReplaySessionCountAndroidAggSum() int64 {
	if o == nil || o.RumMobileReplaySessionCountAndroidAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileReplaySessionCountAndroidAggSum
}

// GetRumMobileReplaySessionCountAndroidAggSumOk returns a tuple with the RumMobileReplaySessionCountAndroidAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileReplaySessionCountAndroidAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileReplaySessionCountAndroidAggSum == nil {
		return nil, false
	}
	return o.RumMobileReplaySessionCountAndroidAggSum, true
}

// HasRumMobileReplaySessionCountAndroidAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileReplaySessionCountAndroidAggSum() bool {
	return o != nil && o.RumMobileReplaySessionCountAndroidAggSum != nil
}

// SetRumMobileReplaySessionCountAndroidAggSum gets a reference to the given int64 and assigns it to the RumMobileReplaySessionCountAndroidAggSum field.
func (o *UsageSummaryResponse) SetRumMobileReplaySessionCountAndroidAggSum(v int64) {
	o.RumMobileReplaySessionCountAndroidAggSum = &v
}

// GetRumMobileReplaySessionCountIosAggSum returns the RumMobileReplaySessionCountIosAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileReplaySessionCountIosAggSum() int64 {
	if o == nil || o.RumMobileReplaySessionCountIosAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileReplaySessionCountIosAggSum
}

// GetRumMobileReplaySessionCountIosAggSumOk returns a tuple with the RumMobileReplaySessionCountIosAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileReplaySessionCountIosAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileReplaySessionCountIosAggSum == nil {
		return nil, false
	}
	return o.RumMobileReplaySessionCountIosAggSum, true
}

// HasRumMobileReplaySessionCountIosAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileReplaySessionCountIosAggSum() bool {
	return o != nil && o.RumMobileReplaySessionCountIosAggSum != nil
}

// SetRumMobileReplaySessionCountIosAggSum gets a reference to the given int64 and assigns it to the RumMobileReplaySessionCountIosAggSum field.
func (o *UsageSummaryResponse) SetRumMobileReplaySessionCountIosAggSum(v int64) {
	o.RumMobileReplaySessionCountIosAggSum = &v
}

// GetRumMobileReplaySessionCountKotlinmultiplatformAggSum returns the RumMobileReplaySessionCountKotlinmultiplatformAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileReplaySessionCountKotlinmultiplatformAggSum() int64 {
	if o == nil || o.RumMobileReplaySessionCountKotlinmultiplatformAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileReplaySessionCountKotlinmultiplatformAggSum
}

// GetRumMobileReplaySessionCountKotlinmultiplatformAggSumOk returns a tuple with the RumMobileReplaySessionCountKotlinmultiplatformAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileReplaySessionCountKotlinmultiplatformAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileReplaySessionCountKotlinmultiplatformAggSum == nil {
		return nil, false
	}
	return o.RumMobileReplaySessionCountKotlinmultiplatformAggSum, true
}

// HasRumMobileReplaySessionCountKotlinmultiplatformAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileReplaySessionCountKotlinmultiplatformAggSum() bool {
	return o != nil && o.RumMobileReplaySessionCountKotlinmultiplatformAggSum != nil
}

// SetRumMobileReplaySessionCountKotlinmultiplatformAggSum gets a reference to the given int64 and assigns it to the RumMobileReplaySessionCountKotlinmultiplatformAggSum field.
func (o *UsageSummaryResponse) SetRumMobileReplaySessionCountKotlinmultiplatformAggSum(v int64) {
	o.RumMobileReplaySessionCountKotlinmultiplatformAggSum = &v
}

// GetRumMobileReplaySessionCountReactnativeAggSum returns the RumMobileReplaySessionCountReactnativeAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumMobileReplaySessionCountReactnativeAggSum() int64 {
	if o == nil || o.RumMobileReplaySessionCountReactnativeAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileReplaySessionCountReactnativeAggSum
}

// GetRumMobileReplaySessionCountReactnativeAggSumOk returns a tuple with the RumMobileReplaySessionCountReactnativeAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumMobileReplaySessionCountReactnativeAggSumOk() (*int64, bool) {
	if o == nil || o.RumMobileReplaySessionCountReactnativeAggSum == nil {
		return nil, false
	}
	return o.RumMobileReplaySessionCountReactnativeAggSum, true
}

// HasRumMobileReplaySessionCountReactnativeAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumMobileReplaySessionCountReactnativeAggSum() bool {
	return o != nil && o.RumMobileReplaySessionCountReactnativeAggSum != nil
}

// SetRumMobileReplaySessionCountReactnativeAggSum gets a reference to the given int64 and assigns it to the RumMobileReplaySessionCountReactnativeAggSum field.
func (o *UsageSummaryResponse) SetRumMobileReplaySessionCountReactnativeAggSum(v int64) {
	o.RumMobileReplaySessionCountReactnativeAggSum = &v
}

// GetRumReplaySessionCountAggSum returns the RumReplaySessionCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumReplaySessionCountAggSum() int64 {
	if o == nil || o.RumReplaySessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumReplaySessionCountAggSum
}

// GetRumReplaySessionCountAggSumOk returns a tuple with the RumReplaySessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumReplaySessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.RumReplaySessionCountAggSum == nil {
		return nil, false
	}
	return o.RumReplaySessionCountAggSum, true
}

// HasRumReplaySessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumReplaySessionCountAggSum() bool {
	return o != nil && o.RumReplaySessionCountAggSum != nil
}

// SetRumReplaySessionCountAggSum gets a reference to the given int64 and assigns it to the RumReplaySessionCountAggSum field.
func (o *UsageSummaryResponse) SetRumReplaySessionCountAggSum(v int64) {
	o.RumReplaySessionCountAggSum = &v
}

// GetRumSessionCountAggSum returns the RumSessionCountAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetRumSessionCountAggSum() int64 {
	if o == nil || o.RumSessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumSessionCountAggSum
}

// GetRumSessionCountAggSumOk returns a tuple with the RumSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetRumSessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.RumSessionCountAggSum == nil {
		return nil, false
	}
	return o.RumSessionCountAggSum, true
}

// HasRumSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumSessionCountAggSum() bool {
	return o != nil && o.RumSessionCountAggSum != nil
}

// SetRumSessionCountAggSum gets a reference to the given int64 and assigns it to the RumSessionCountAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetRumSessionCountAggSum(v int64) {
	o.RumSessionCountAggSum = &v
}

// GetRumSessionReplayAddOnAggSum returns the RumSessionReplayAddOnAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumSessionReplayAddOnAggSum() int64 {
	if o == nil || o.RumSessionReplayAddOnAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumSessionReplayAddOnAggSum
}

// GetRumSessionReplayAddOnAggSumOk returns a tuple with the RumSessionReplayAddOnAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumSessionReplayAddOnAggSumOk() (*int64, bool) {
	if o == nil || o.RumSessionReplayAddOnAggSum == nil {
		return nil, false
	}
	return o.RumSessionReplayAddOnAggSum, true
}

// HasRumSessionReplayAddOnAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumSessionReplayAddOnAggSum() bool {
	return o != nil && o.RumSessionReplayAddOnAggSum != nil
}

// SetRumSessionReplayAddOnAggSum gets a reference to the given int64 and assigns it to the RumSessionReplayAddOnAggSum field.
func (o *UsageSummaryResponse) SetRumSessionReplayAddOnAggSum(v int64) {
	o.RumSessionReplayAddOnAggSum = &v
}

// GetRumTotalSessionCountAggSum returns the RumTotalSessionCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetRumTotalSessionCountAggSum() int64 {
	if o == nil || o.RumTotalSessionCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumTotalSessionCountAggSum
}

// GetRumTotalSessionCountAggSumOk returns a tuple with the RumTotalSessionCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetRumTotalSessionCountAggSumOk() (*int64, bool) {
	if o == nil || o.RumTotalSessionCountAggSum == nil {
		return nil, false
	}
	return o.RumTotalSessionCountAggSum, true
}

// HasRumTotalSessionCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumTotalSessionCountAggSum() bool {
	return o != nil && o.RumTotalSessionCountAggSum != nil
}

// SetRumTotalSessionCountAggSum gets a reference to the given int64 and assigns it to the RumTotalSessionCountAggSum field.
func (o *UsageSummaryResponse) SetRumTotalSessionCountAggSum(v int64) {
	o.RumTotalSessionCountAggSum = &v
}

// GetRumUnitsAggSum returns the RumUnitsAggSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryResponse) GetRumUnitsAggSum() int64 {
	if o == nil || o.RumUnitsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.RumUnitsAggSum
}

// GetRumUnitsAggSumOk returns a tuple with the RumUnitsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryResponse) GetRumUnitsAggSumOk() (*int64, bool) {
	if o == nil || o.RumUnitsAggSum == nil {
		return nil, false
	}
	return o.RumUnitsAggSum, true
}

// HasRumUnitsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasRumUnitsAggSum() bool {
	return o != nil && o.RumUnitsAggSum != nil
}

// SetRumUnitsAggSum gets a reference to the given int64 and assigns it to the RumUnitsAggSum field.
// Deprecated
func (o *UsageSummaryResponse) SetRumUnitsAggSum(v int64) {
	o.RumUnitsAggSum = &v
}

// GetScaFargateCountAvgSum returns the ScaFargateCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetScaFargateCountAvgSum() int64 {
	if o == nil || o.ScaFargateCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ScaFargateCountAvgSum
}

// GetScaFargateCountAvgSumOk returns a tuple with the ScaFargateCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetScaFargateCountAvgSumOk() (*int64, bool) {
	if o == nil || o.ScaFargateCountAvgSum == nil {
		return nil, false
	}
	return o.ScaFargateCountAvgSum, true
}

// HasScaFargateCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasScaFargateCountAvgSum() bool {
	return o != nil && o.ScaFargateCountAvgSum != nil
}

// SetScaFargateCountAvgSum gets a reference to the given int64 and assigns it to the ScaFargateCountAvgSum field.
func (o *UsageSummaryResponse) SetScaFargateCountAvgSum(v int64) {
	o.ScaFargateCountAvgSum = &v
}

// GetScaFargateCountHwmSum returns the ScaFargateCountHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetScaFargateCountHwmSum() int64 {
	if o == nil || o.ScaFargateCountHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.ScaFargateCountHwmSum
}

// GetScaFargateCountHwmSumOk returns a tuple with the ScaFargateCountHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetScaFargateCountHwmSumOk() (*int64, bool) {
	if o == nil || o.ScaFargateCountHwmSum == nil {
		return nil, false
	}
	return o.ScaFargateCountHwmSum, true
}

// HasScaFargateCountHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasScaFargateCountHwmSum() bool {
	return o != nil && o.ScaFargateCountHwmSum != nil
}

// SetScaFargateCountHwmSum gets a reference to the given int64 and assigns it to the ScaFargateCountHwmSum field.
func (o *UsageSummaryResponse) SetScaFargateCountHwmSum(v int64) {
	o.ScaFargateCountHwmSum = &v
}

// GetSdsApmScannedBytesSum returns the SdsApmScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSdsApmScannedBytesSum() int64 {
	if o == nil || o.SdsApmScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsApmScannedBytesSum
}

// GetSdsApmScannedBytesSumOk returns a tuple with the SdsApmScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSdsApmScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsApmScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsApmScannedBytesSum, true
}

// HasSdsApmScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsApmScannedBytesSum() bool {
	return o != nil && o.SdsApmScannedBytesSum != nil
}

// SetSdsApmScannedBytesSum gets a reference to the given int64 and assigns it to the SdsApmScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsApmScannedBytesSum(v int64) {
	o.SdsApmScannedBytesSum = &v
}

// GetSdsEventsScannedBytesSum returns the SdsEventsScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSdsEventsScannedBytesSum() int64 {
	if o == nil || o.SdsEventsScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsEventsScannedBytesSum
}

// GetSdsEventsScannedBytesSumOk returns a tuple with the SdsEventsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSdsEventsScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsEventsScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsEventsScannedBytesSum, true
}

// HasSdsEventsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsEventsScannedBytesSum() bool {
	return o != nil && o.SdsEventsScannedBytesSum != nil
}

// SetSdsEventsScannedBytesSum gets a reference to the given int64 and assigns it to the SdsEventsScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsEventsScannedBytesSum(v int64) {
	o.SdsEventsScannedBytesSum = &v
}

// GetSdsLogsScannedBytesSum returns the SdsLogsScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSdsLogsScannedBytesSum() int64 {
	if o == nil || o.SdsLogsScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsLogsScannedBytesSum
}

// GetSdsLogsScannedBytesSumOk returns a tuple with the SdsLogsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSdsLogsScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsLogsScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsLogsScannedBytesSum, true
}

// HasSdsLogsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsLogsScannedBytesSum() bool {
	return o != nil && o.SdsLogsScannedBytesSum != nil
}

// SetSdsLogsScannedBytesSum gets a reference to the given int64 and assigns it to the SdsLogsScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsLogsScannedBytesSum(v int64) {
	o.SdsLogsScannedBytesSum = &v
}

// GetSdsRumScannedBytesSum returns the SdsRumScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSdsRumScannedBytesSum() int64 {
	if o == nil || o.SdsRumScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsRumScannedBytesSum
}

// GetSdsRumScannedBytesSumOk returns a tuple with the SdsRumScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSdsRumScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsRumScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsRumScannedBytesSum, true
}

// HasSdsRumScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsRumScannedBytesSum() bool {
	return o != nil && o.SdsRumScannedBytesSum != nil
}

// SetSdsRumScannedBytesSum gets a reference to the given int64 and assigns it to the SdsRumScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsRumScannedBytesSum(v int64) {
	o.SdsRumScannedBytesSum = &v
}

// GetSdsTotalScannedBytesSum returns the SdsTotalScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSdsTotalScannedBytesSum() int64 {
	if o == nil || o.SdsTotalScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsTotalScannedBytesSum
}

// GetSdsTotalScannedBytesSumOk returns a tuple with the SdsTotalScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSdsTotalScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsTotalScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsTotalScannedBytesSum, true
}

// HasSdsTotalScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSdsTotalScannedBytesSum() bool {
	return o != nil && o.SdsTotalScannedBytesSum != nil
}

// SetSdsTotalScannedBytesSum gets a reference to the given int64 and assigns it to the SdsTotalScannedBytesSum field.
func (o *UsageSummaryResponse) SetSdsTotalScannedBytesSum(v int64) {
	o.SdsTotalScannedBytesSum = &v
}

// GetServerlessAppsApmApmAzureAppserviceInstancesAvgSum returns the ServerlessAppsApmApmAzureAppserviceInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmAzureAppserviceInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum
}

// GetServerlessAppsApmApmAzureAppserviceInstancesAvgSumOk returns a tuple with the ServerlessAppsApmApmAzureAppserviceInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmAzureAppserviceInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum, true
}

// HasServerlessAppsApmApmAzureAppserviceInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmApmAzureAppserviceInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum != nil
}

// SetServerlessAppsApmApmAzureAppserviceInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmAzureAppserviceInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmApmAzureAppserviceInstancesAvgSum(v int64) {
	o.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum = &v
}

// GetServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum returns the ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum
}

// GetServerlessAppsApmApmAzureAzurefunctionInstancesAvgSumOk returns a tuple with the ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmAzureAzurefunctionInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum, true
}

// HasServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum != nil
}

// SetServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum(v int64) {
	o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum = &v
}

// GetServerlessAppsApmApmAzureContainerappInstancesAvgSum returns the ServerlessAppsApmApmAzureContainerappInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmAzureContainerappInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmApmAzureContainerappInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmAzureContainerappInstancesAvgSum
}

// GetServerlessAppsApmApmAzureContainerappInstancesAvgSumOk returns a tuple with the ServerlessAppsApmApmAzureContainerappInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmAzureContainerappInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmAzureContainerappInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmAzureContainerappInstancesAvgSum, true
}

// HasServerlessAppsApmApmAzureContainerappInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmApmAzureContainerappInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmApmAzureContainerappInstancesAvgSum != nil
}

// SetServerlessAppsApmApmAzureContainerappInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmAzureContainerappInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmApmAzureContainerappInstancesAvgSum(v int64) {
	o.ServerlessAppsApmApmAzureContainerappInstancesAvgSum = &v
}

// GetServerlessAppsApmApmFargateEcsTasksAvgSum returns the ServerlessAppsApmApmFargateEcsTasksAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmFargateEcsTasksAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmApmFargateEcsTasksAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmFargateEcsTasksAvgSum
}

// GetServerlessAppsApmApmFargateEcsTasksAvgSumOk returns a tuple with the ServerlessAppsApmApmFargateEcsTasksAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmFargateEcsTasksAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmFargateEcsTasksAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmFargateEcsTasksAvgSum, true
}

// HasServerlessAppsApmApmFargateEcsTasksAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmApmFargateEcsTasksAvgSum() bool {
	return o != nil && o.ServerlessAppsApmApmFargateEcsTasksAvgSum != nil
}

// SetServerlessAppsApmApmFargateEcsTasksAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmFargateEcsTasksAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmApmFargateEcsTasksAvgSum(v int64) {
	o.ServerlessAppsApmApmFargateEcsTasksAvgSum = &v
}

// GetServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum returns the ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum
}

// GetServerlessAppsApmApmGcpCloudfunctionInstancesAvgSumOk returns a tuple with the ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmGcpCloudfunctionInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum, true
}

// HasServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum != nil
}

// SetServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum(v int64) {
	o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum = &v
}

// GetServerlessAppsApmApmGcpCloudrunInstancesAvgSum returns the ServerlessAppsApmApmGcpCloudrunInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmGcpCloudrunInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum
}

// GetServerlessAppsApmApmGcpCloudrunInstancesAvgSumOk returns a tuple with the ServerlessAppsApmApmGcpCloudrunInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmApmGcpCloudrunInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum, true
}

// HasServerlessAppsApmApmGcpCloudrunInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmApmGcpCloudrunInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum != nil
}

// SetServerlessAppsApmApmGcpCloudrunInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmGcpCloudrunInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmApmGcpCloudrunInstancesAvgSum(v int64) {
	o.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum = &v
}

// GetServerlessAppsApmAvgSum returns the ServerlessAppsApmAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmAvgSum
}

// GetServerlessAppsApmAvgSumOk returns a tuple with the ServerlessAppsApmAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmAvgSum, true
}

// HasServerlessAppsApmAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmAvgSum() bool {
	return o != nil && o.ServerlessAppsApmAvgSum != nil
}

// SetServerlessAppsApmAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmAvgSum(v int64) {
	o.ServerlessAppsApmAvgSum = &v
}

// GetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum returns the ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum
}

// GetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSumOk returns a tuple with the ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum, true
}

// HasServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum != nil
}

// SetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum(v int64) {
	o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum = &v
}

// GetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum returns the ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum
}

// GetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSumOk returns a tuple with the ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum, true
}

// HasServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum != nil
}

// SetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum(v int64) {
	o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum = &v
}

// GetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum returns the ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum
}

// GetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSumOk returns a tuple with the ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum, true
}

// HasServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum != nil
}

// SetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum(v int64) {
	o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum = &v
}

// GetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum returns the ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum
}

// GetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSumOk returns a tuple with the ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum, true
}

// HasServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum != nil
}

// SetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum(v int64) {
	o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum = &v
}

// GetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum returns the ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum
}

// GetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSumOk returns a tuple with the ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum, true
}

// HasServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum != nil
}

// SetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum(v int64) {
	o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum = &v
}

// GetServerlessAppsApmExclFargateAvgSum returns the ServerlessAppsApmExclFargateAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateAvgSum() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateAvgSum
}

// GetServerlessAppsApmExclFargateAvgSumOk returns a tuple with the ServerlessAppsApmExclFargateAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsApmExclFargateAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateAvgSum, true
}

// HasServerlessAppsApmExclFargateAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsApmExclFargateAvgSum() bool {
	return o != nil && o.ServerlessAppsApmExclFargateAvgSum != nil
}

// SetServerlessAppsApmExclFargateAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsApmExclFargateAvgSum(v int64) {
	o.ServerlessAppsApmExclFargateAvgSum = &v
}

// GetServerlessAppsAzureContainerAppInstancesAvgSum returns the ServerlessAppsAzureContainerAppInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsAzureContainerAppInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsAzureContainerAppInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsAzureContainerAppInstancesAvgSum
}

// GetServerlessAppsAzureContainerAppInstancesAvgSumOk returns a tuple with the ServerlessAppsAzureContainerAppInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsAzureContainerAppInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsAzureContainerAppInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsAzureContainerAppInstancesAvgSum, true
}

// HasServerlessAppsAzureContainerAppInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsAzureContainerAppInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsAzureContainerAppInstancesAvgSum != nil
}

// SetServerlessAppsAzureContainerAppInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsAzureContainerAppInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsAzureContainerAppInstancesAvgSum(v int64) {
	o.ServerlessAppsAzureContainerAppInstancesAvgSum = &v
}

// GetServerlessAppsAzureCountAvgSum returns the ServerlessAppsAzureCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsAzureCountAvgSum() int64 {
	if o == nil || o.ServerlessAppsAzureCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsAzureCountAvgSum
}

// GetServerlessAppsAzureCountAvgSumOk returns a tuple with the ServerlessAppsAzureCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsAzureCountAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsAzureCountAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsAzureCountAvgSum, true
}

// HasServerlessAppsAzureCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsAzureCountAvgSum() bool {
	return o != nil && o.ServerlessAppsAzureCountAvgSum != nil
}

// SetServerlessAppsAzureCountAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsAzureCountAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsAzureCountAvgSum(v int64) {
	o.ServerlessAppsAzureCountAvgSum = &v
}

// GetServerlessAppsAzureFunctionAppInstancesAvgSum returns the ServerlessAppsAzureFunctionAppInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsAzureFunctionAppInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsAzureFunctionAppInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsAzureFunctionAppInstancesAvgSum
}

// GetServerlessAppsAzureFunctionAppInstancesAvgSumOk returns a tuple with the ServerlessAppsAzureFunctionAppInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsAzureFunctionAppInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsAzureFunctionAppInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsAzureFunctionAppInstancesAvgSum, true
}

// HasServerlessAppsAzureFunctionAppInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsAzureFunctionAppInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsAzureFunctionAppInstancesAvgSum != nil
}

// SetServerlessAppsAzureFunctionAppInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsAzureFunctionAppInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsAzureFunctionAppInstancesAvgSum(v int64) {
	o.ServerlessAppsAzureFunctionAppInstancesAvgSum = &v
}

// GetServerlessAppsAzureWebAppInstancesAvgSum returns the ServerlessAppsAzureWebAppInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsAzureWebAppInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsAzureWebAppInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsAzureWebAppInstancesAvgSum
}

// GetServerlessAppsAzureWebAppInstancesAvgSumOk returns a tuple with the ServerlessAppsAzureWebAppInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsAzureWebAppInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsAzureWebAppInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsAzureWebAppInstancesAvgSum, true
}

// HasServerlessAppsAzureWebAppInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsAzureWebAppInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsAzureWebAppInstancesAvgSum != nil
}

// SetServerlessAppsAzureWebAppInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsAzureWebAppInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsAzureWebAppInstancesAvgSum(v int64) {
	o.ServerlessAppsAzureWebAppInstancesAvgSum = &v
}

// GetServerlessAppsEcsAvgSum returns the ServerlessAppsEcsAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsEcsAvgSum() int64 {
	if o == nil || o.ServerlessAppsEcsAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsEcsAvgSum
}

// GetServerlessAppsEcsAvgSumOk returns a tuple with the ServerlessAppsEcsAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsEcsAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsEcsAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsEcsAvgSum, true
}

// HasServerlessAppsEcsAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsEcsAvgSum() bool {
	return o != nil && o.ServerlessAppsEcsAvgSum != nil
}

// SetServerlessAppsEcsAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsEcsAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsEcsAvgSum(v int64) {
	o.ServerlessAppsEcsAvgSum = &v
}

// GetServerlessAppsEksAvgSum returns the ServerlessAppsEksAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsEksAvgSum() int64 {
	if o == nil || o.ServerlessAppsEksAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsEksAvgSum
}

// GetServerlessAppsEksAvgSumOk returns a tuple with the ServerlessAppsEksAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsEksAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsEksAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsEksAvgSum, true
}

// HasServerlessAppsEksAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsEksAvgSum() bool {
	return o != nil && o.ServerlessAppsEksAvgSum != nil
}

// SetServerlessAppsEksAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsEksAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsEksAvgSum(v int64) {
	o.ServerlessAppsEksAvgSum = &v
}

// GetServerlessAppsExclFargateAvgSum returns the ServerlessAppsExclFargateAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateAvgSum() int64 {
	if o == nil || o.ServerlessAppsExclFargateAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateAvgSum
}

// GetServerlessAppsExclFargateAvgSumOk returns a tuple with the ServerlessAppsExclFargateAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateAvgSum, true
}

// HasServerlessAppsExclFargateAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsExclFargateAvgSum() bool {
	return o != nil && o.ServerlessAppsExclFargateAvgSum != nil
}

// SetServerlessAppsExclFargateAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsExclFargateAvgSum(v int64) {
	o.ServerlessAppsExclFargateAvgSum = &v
}

// GetServerlessAppsExclFargateAzureContainerAppInstancesAvgSum returns the ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateAzureContainerAppInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum
}

// GetServerlessAppsExclFargateAzureContainerAppInstancesAvgSumOk returns a tuple with the ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateAzureContainerAppInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum, true
}

// HasServerlessAppsExclFargateAzureContainerAppInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsExclFargateAzureContainerAppInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum != nil
}

// SetServerlessAppsExclFargateAzureContainerAppInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsExclFargateAzureContainerAppInstancesAvgSum(v int64) {
	o.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum = &v
}

// GetServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum returns the ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum
}

// GetServerlessAppsExclFargateAzureFunctionAppInstancesAvgSumOk returns a tuple with the ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateAzureFunctionAppInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum, true
}

// HasServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum != nil
}

// SetServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum(v int64) {
	o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum = &v
}

// GetServerlessAppsExclFargateAzureWebAppInstancesAvgSum returns the ServerlessAppsExclFargateAzureWebAppInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateAzureWebAppInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum
}

// GetServerlessAppsExclFargateAzureWebAppInstancesAvgSumOk returns a tuple with the ServerlessAppsExclFargateAzureWebAppInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateAzureWebAppInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum, true
}

// HasServerlessAppsExclFargateAzureWebAppInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsExclFargateAzureWebAppInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum != nil
}

// SetServerlessAppsExclFargateAzureWebAppInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateAzureWebAppInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsExclFargateAzureWebAppInstancesAvgSum(v int64) {
	o.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum = &v
}

// GetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum returns the ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum
}

// GetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSumOk returns a tuple with the ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum, true
}

// HasServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum != nil
}

// SetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum(v int64) {
	o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum = &v
}

// GetServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum returns the ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum
}

// GetServerlessAppsExclFargateGoogleCloudRunInstancesAvgSumOk returns a tuple with the ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsExclFargateGoogleCloudRunInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum, true
}

// HasServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum != nil
}

// SetServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum(v int64) {
	o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum = &v
}

// GetServerlessAppsGoogleCloudFunctionsInstancesAvgSum returns the ServerlessAppsGoogleCloudFunctionsInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsGoogleCloudFunctionsInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum
}

// GetServerlessAppsGoogleCloudFunctionsInstancesAvgSumOk returns a tuple with the ServerlessAppsGoogleCloudFunctionsInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsGoogleCloudFunctionsInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum, true
}

// HasServerlessAppsGoogleCloudFunctionsInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsGoogleCloudFunctionsInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum != nil
}

// SetServerlessAppsGoogleCloudFunctionsInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsGoogleCloudFunctionsInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsGoogleCloudFunctionsInstancesAvgSum(v int64) {
	o.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum = &v
}

// GetServerlessAppsGoogleCloudRunInstancesAvgSum returns the ServerlessAppsGoogleCloudRunInstancesAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsGoogleCloudRunInstancesAvgSum() int64 {
	if o == nil || o.ServerlessAppsGoogleCloudRunInstancesAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsGoogleCloudRunInstancesAvgSum
}

// GetServerlessAppsGoogleCloudRunInstancesAvgSumOk returns a tuple with the ServerlessAppsGoogleCloudRunInstancesAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsGoogleCloudRunInstancesAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsGoogleCloudRunInstancesAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsGoogleCloudRunInstancesAvgSum, true
}

// HasServerlessAppsGoogleCloudRunInstancesAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsGoogleCloudRunInstancesAvgSum() bool {
	return o != nil && o.ServerlessAppsGoogleCloudRunInstancesAvgSum != nil
}

// SetServerlessAppsGoogleCloudRunInstancesAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsGoogleCloudRunInstancesAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsGoogleCloudRunInstancesAvgSum(v int64) {
	o.ServerlessAppsGoogleCloudRunInstancesAvgSum = &v
}

// GetServerlessAppsGoogleCountAvgSum returns the ServerlessAppsGoogleCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsGoogleCountAvgSum() int64 {
	if o == nil || o.ServerlessAppsGoogleCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsGoogleCountAvgSum
}

// GetServerlessAppsGoogleCountAvgSumOk returns a tuple with the ServerlessAppsGoogleCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsGoogleCountAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsGoogleCountAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsGoogleCountAvgSum, true
}

// HasServerlessAppsGoogleCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsGoogleCountAvgSum() bool {
	return o != nil && o.ServerlessAppsGoogleCountAvgSum != nil
}

// SetServerlessAppsGoogleCountAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsGoogleCountAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsGoogleCountAvgSum(v int64) {
	o.ServerlessAppsGoogleCountAvgSum = &v
}

// GetServerlessAppsTotalCountAvgSum returns the ServerlessAppsTotalCountAvgSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetServerlessAppsTotalCountAvgSum() int64 {
	if o == nil || o.ServerlessAppsTotalCountAvgSum == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsTotalCountAvgSum
}

// GetServerlessAppsTotalCountAvgSumOk returns a tuple with the ServerlessAppsTotalCountAvgSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetServerlessAppsTotalCountAvgSumOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsTotalCountAvgSum == nil {
		return nil, false
	}
	return o.ServerlessAppsTotalCountAvgSum, true
}

// HasServerlessAppsTotalCountAvgSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasServerlessAppsTotalCountAvgSum() bool {
	return o != nil && o.ServerlessAppsTotalCountAvgSum != nil
}

// SetServerlessAppsTotalCountAvgSum gets a reference to the given int64 and assigns it to the ServerlessAppsTotalCountAvgSum field.
func (o *UsageSummaryResponse) SetServerlessAppsTotalCountAvgSum(v int64) {
	o.ServerlessAppsTotalCountAvgSum = &v
}

// GetSiemAnalyzedLogsAddOnCountAggSum returns the SiemAnalyzedLogsAddOnCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSiemAnalyzedLogsAddOnCountAggSum() int64 {
	if o == nil || o.SiemAnalyzedLogsAddOnCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.SiemAnalyzedLogsAddOnCountAggSum
}

// GetSiemAnalyzedLogsAddOnCountAggSumOk returns a tuple with the SiemAnalyzedLogsAddOnCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSiemAnalyzedLogsAddOnCountAggSumOk() (*int64, bool) {
	if o == nil || o.SiemAnalyzedLogsAddOnCountAggSum == nil {
		return nil, false
	}
	return o.SiemAnalyzedLogsAddOnCountAggSum, true
}

// HasSiemAnalyzedLogsAddOnCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSiemAnalyzedLogsAddOnCountAggSum() bool {
	return o != nil && o.SiemAnalyzedLogsAddOnCountAggSum != nil
}

// SetSiemAnalyzedLogsAddOnCountAggSum gets a reference to the given int64 and assigns it to the SiemAnalyzedLogsAddOnCountAggSum field.
func (o *UsageSummaryResponse) SetSiemAnalyzedLogsAddOnCountAggSum(v int64) {
	o.SiemAnalyzedLogsAddOnCountAggSum = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *UsageSummaryResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetSyntheticsBrowserCheckCallsCountAggSum returns the SyntheticsBrowserCheckCallsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSyntheticsBrowserCheckCallsCountAggSum() int64 {
	if o == nil || o.SyntheticsBrowserCheckCallsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsBrowserCheckCallsCountAggSum
}

// GetSyntheticsBrowserCheckCallsCountAggSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSyntheticsBrowserCheckCallsCountAggSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsBrowserCheckCallsCountAggSum == nil {
		return nil, false
	}
	return o.SyntheticsBrowserCheckCallsCountAggSum, true
}

// HasSyntheticsBrowserCheckCallsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSyntheticsBrowserCheckCallsCountAggSum() bool {
	return o != nil && o.SyntheticsBrowserCheckCallsCountAggSum != nil
}

// SetSyntheticsBrowserCheckCallsCountAggSum gets a reference to the given int64 and assigns it to the SyntheticsBrowserCheckCallsCountAggSum field.
func (o *UsageSummaryResponse) SetSyntheticsBrowserCheckCallsCountAggSum(v int64) {
	o.SyntheticsBrowserCheckCallsCountAggSum = &v
}

// GetSyntheticsCheckCallsCountAggSum returns the SyntheticsCheckCallsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSyntheticsCheckCallsCountAggSum() int64 {
	if o == nil || o.SyntheticsCheckCallsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsCheckCallsCountAggSum
}

// GetSyntheticsCheckCallsCountAggSumOk returns a tuple with the SyntheticsCheckCallsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSyntheticsCheckCallsCountAggSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsCheckCallsCountAggSum == nil {
		return nil, false
	}
	return o.SyntheticsCheckCallsCountAggSum, true
}

// HasSyntheticsCheckCallsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSyntheticsCheckCallsCountAggSum() bool {
	return o != nil && o.SyntheticsCheckCallsCountAggSum != nil
}

// SetSyntheticsCheckCallsCountAggSum gets a reference to the given int64 and assigns it to the SyntheticsCheckCallsCountAggSum field.
func (o *UsageSummaryResponse) SetSyntheticsCheckCallsCountAggSum(v int64) {
	o.SyntheticsCheckCallsCountAggSum = &v
}

// GetSyntheticsMobileTestRunsAggSum returns the SyntheticsMobileTestRunsAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSyntheticsMobileTestRunsAggSum() int64 {
	if o == nil || o.SyntheticsMobileTestRunsAggSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsMobileTestRunsAggSum
}

// GetSyntheticsMobileTestRunsAggSumOk returns a tuple with the SyntheticsMobileTestRunsAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSyntheticsMobileTestRunsAggSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsMobileTestRunsAggSum == nil {
		return nil, false
	}
	return o.SyntheticsMobileTestRunsAggSum, true
}

// HasSyntheticsMobileTestRunsAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSyntheticsMobileTestRunsAggSum() bool {
	return o != nil && o.SyntheticsMobileTestRunsAggSum != nil
}

// SetSyntheticsMobileTestRunsAggSum gets a reference to the given int64 and assigns it to the SyntheticsMobileTestRunsAggSum field.
func (o *UsageSummaryResponse) SetSyntheticsMobileTestRunsAggSum(v int64) {
	o.SyntheticsMobileTestRunsAggSum = &v
}

// GetSyntheticsParallelTestingMaxSlotsHwmSum returns the SyntheticsParallelTestingMaxSlotsHwmSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetSyntheticsParallelTestingMaxSlotsHwmSum() int64 {
	if o == nil || o.SyntheticsParallelTestingMaxSlotsHwmSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsParallelTestingMaxSlotsHwmSum
}

// GetSyntheticsParallelTestingMaxSlotsHwmSumOk returns a tuple with the SyntheticsParallelTestingMaxSlotsHwmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetSyntheticsParallelTestingMaxSlotsHwmSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsParallelTestingMaxSlotsHwmSum == nil {
		return nil, false
	}
	return o.SyntheticsParallelTestingMaxSlotsHwmSum, true
}

// HasSyntheticsParallelTestingMaxSlotsHwmSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasSyntheticsParallelTestingMaxSlotsHwmSum() bool {
	return o != nil && o.SyntheticsParallelTestingMaxSlotsHwmSum != nil
}

// SetSyntheticsParallelTestingMaxSlotsHwmSum gets a reference to the given int64 and assigns it to the SyntheticsParallelTestingMaxSlotsHwmSum field.
func (o *UsageSummaryResponse) SetSyntheticsParallelTestingMaxSlotsHwmSum(v int64) {
	o.SyntheticsParallelTestingMaxSlotsHwmSum = &v
}

// GetTraceSearchIndexedEventsCountAggSum returns the TraceSearchIndexedEventsCountAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetTraceSearchIndexedEventsCountAggSum() int64 {
	if o == nil || o.TraceSearchIndexedEventsCountAggSum == nil {
		var ret int64
		return ret
	}
	return *o.TraceSearchIndexedEventsCountAggSum
}

// GetTraceSearchIndexedEventsCountAggSumOk returns a tuple with the TraceSearchIndexedEventsCountAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetTraceSearchIndexedEventsCountAggSumOk() (*int64, bool) {
	if o == nil || o.TraceSearchIndexedEventsCountAggSum == nil {
		return nil, false
	}
	return o.TraceSearchIndexedEventsCountAggSum, true
}

// HasTraceSearchIndexedEventsCountAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasTraceSearchIndexedEventsCountAggSum() bool {
	return o != nil && o.TraceSearchIndexedEventsCountAggSum != nil
}

// SetTraceSearchIndexedEventsCountAggSum gets a reference to the given int64 and assigns it to the TraceSearchIndexedEventsCountAggSum field.
func (o *UsageSummaryResponse) SetTraceSearchIndexedEventsCountAggSum(v int64) {
	o.TraceSearchIndexedEventsCountAggSum = &v
}

// GetTwolIngestedEventsBytesAggSum returns the TwolIngestedEventsBytesAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetTwolIngestedEventsBytesAggSum() int64 {
	if o == nil || o.TwolIngestedEventsBytesAggSum == nil {
		var ret int64
		return ret
	}
	return *o.TwolIngestedEventsBytesAggSum
}

// GetTwolIngestedEventsBytesAggSumOk returns a tuple with the TwolIngestedEventsBytesAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetTwolIngestedEventsBytesAggSumOk() (*int64, bool) {
	if o == nil || o.TwolIngestedEventsBytesAggSum == nil {
		return nil, false
	}
	return o.TwolIngestedEventsBytesAggSum, true
}

// HasTwolIngestedEventsBytesAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasTwolIngestedEventsBytesAggSum() bool {
	return o != nil && o.TwolIngestedEventsBytesAggSum != nil
}

// SetTwolIngestedEventsBytesAggSum gets a reference to the given int64 and assigns it to the TwolIngestedEventsBytesAggSum field.
func (o *UsageSummaryResponse) SetTwolIngestedEventsBytesAggSum(v int64) {
	o.TwolIngestedEventsBytesAggSum = &v
}

// GetUniversalServiceMonitoringHostTop99pSum returns the UniversalServiceMonitoringHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetUniversalServiceMonitoringHostTop99pSum() int64 {
	if o == nil || o.UniversalServiceMonitoringHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.UniversalServiceMonitoringHostTop99pSum
}

// GetUniversalServiceMonitoringHostTop99pSumOk returns a tuple with the UniversalServiceMonitoringHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetUniversalServiceMonitoringHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.UniversalServiceMonitoringHostTop99pSum == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringHostTop99pSum, true
}

// HasUniversalServiceMonitoringHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasUniversalServiceMonitoringHostTop99pSum() bool {
	return o != nil && o.UniversalServiceMonitoringHostTop99pSum != nil
}

// SetUniversalServiceMonitoringHostTop99pSum gets a reference to the given int64 and assigns it to the UniversalServiceMonitoringHostTop99pSum field.
func (o *UsageSummaryResponse) SetUniversalServiceMonitoringHostTop99pSum(v int64) {
	o.UniversalServiceMonitoringHostTop99pSum = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetUsage() []UsageSummaryDate {
	if o == nil || o.Usage == nil {
		var ret []UsageSummaryDate
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetUsageOk() (*[]UsageSummaryDate, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given []UsageSummaryDate and assigns it to the Usage field.
func (o *UsageSummaryResponse) SetUsage(v []UsageSummaryDate) {
	o.Usage = v
}

// GetVsphereHostTop99pSum returns the VsphereHostTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetVsphereHostTop99pSum() int64 {
	if o == nil || o.VsphereHostTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.VsphereHostTop99pSum
}

// GetVsphereHostTop99pSumOk returns a tuple with the VsphereHostTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetVsphereHostTop99pSumOk() (*int64, bool) {
	if o == nil || o.VsphereHostTop99pSum == nil {
		return nil, false
	}
	return o.VsphereHostTop99pSum, true
}

// HasVsphereHostTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasVsphereHostTop99pSum() bool {
	return o != nil && o.VsphereHostTop99pSum != nil
}

// SetVsphereHostTop99pSum gets a reference to the given int64 and assigns it to the VsphereHostTop99pSum field.
func (o *UsageSummaryResponse) SetVsphereHostTop99pSum(v int64) {
	o.VsphereHostTop99pSum = &v
}

// GetVulnManagementHostCountTop99pSum returns the VulnManagementHostCountTop99pSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetVulnManagementHostCountTop99pSum() int64 {
	if o == nil || o.VulnManagementHostCountTop99pSum == nil {
		var ret int64
		return ret
	}
	return *o.VulnManagementHostCountTop99pSum
}

// GetVulnManagementHostCountTop99pSumOk returns a tuple with the VulnManagementHostCountTop99pSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetVulnManagementHostCountTop99pSumOk() (*int64, bool) {
	if o == nil || o.VulnManagementHostCountTop99pSum == nil {
		return nil, false
	}
	return o.VulnManagementHostCountTop99pSum, true
}

// HasVulnManagementHostCountTop99pSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasVulnManagementHostCountTop99pSum() bool {
	return o != nil && o.VulnManagementHostCountTop99pSum != nil
}

// SetVulnManagementHostCountTop99pSum gets a reference to the given int64 and assigns it to the VulnManagementHostCountTop99pSum field.
func (o *UsageSummaryResponse) SetVulnManagementHostCountTop99pSum(v int64) {
	o.VulnManagementHostCountTop99pSum = &v
}

// GetWorkflowExecutionsUsageAggSum returns the WorkflowExecutionsUsageAggSum field value if set, zero value otherwise.
func (o *UsageSummaryResponse) GetWorkflowExecutionsUsageAggSum() int64 {
	if o == nil || o.WorkflowExecutionsUsageAggSum == nil {
		var ret int64
		return ret
	}
	return *o.WorkflowExecutionsUsageAggSum
}

// GetWorkflowExecutionsUsageAggSumOk returns a tuple with the WorkflowExecutionsUsageAggSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryResponse) GetWorkflowExecutionsUsageAggSumOk() (*int64, bool) {
	if o == nil || o.WorkflowExecutionsUsageAggSum == nil {
		return nil, false
	}
	return o.WorkflowExecutionsUsageAggSum, true
}

// HasWorkflowExecutionsUsageAggSum returns a boolean if a field has been set.
func (o *UsageSummaryResponse) HasWorkflowExecutionsUsageAggSum() bool {
	return o != nil && o.WorkflowExecutionsUsageAggSum != nil
}

// SetWorkflowExecutionsUsageAggSum gets a reference to the given int64 and assigns it to the WorkflowExecutionsUsageAggSum field.
func (o *UsageSummaryResponse) SetWorkflowExecutionsUsageAggSum(v int64) {
	o.WorkflowExecutionsUsageAggSum = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentHostTop99pSum != nil {
		toSerialize["agent_host_top99p_sum"] = o.AgentHostTop99pSum
	}
	if o.ApmAzureAppServiceHostTop99pSum != nil {
		toSerialize["apm_azure_app_service_host_top99p_sum"] = o.ApmAzureAppServiceHostTop99pSum
	}
	if o.ApmDevsecopsHostTop99pSum != nil {
		toSerialize["apm_devsecops_host_top99p_sum"] = o.ApmDevsecopsHostTop99pSum
	}
	if o.ApmEnterpriseStandaloneHostsTop99pSum != nil {
		toSerialize["apm_enterprise_standalone_hosts_top99p_sum"] = o.ApmEnterpriseStandaloneHostsTop99pSum
	}
	if o.ApmFargateCountAvgSum != nil {
		toSerialize["apm_fargate_count_avg_sum"] = o.ApmFargateCountAvgSum
	}
	if o.ApmHostTop99pSum != nil {
		toSerialize["apm_host_top99p_sum"] = o.ApmHostTop99pSum
	}
	if o.ApmProStandaloneHostsTop99pSum != nil {
		toSerialize["apm_pro_standalone_hosts_top99p_sum"] = o.ApmProStandaloneHostsTop99pSum
	}
	if o.AppsecFargateCountAvgSum != nil {
		toSerialize["appsec_fargate_count_avg_sum"] = o.AppsecFargateCountAvgSum
	}
	if o.AsmServerlessAggSum != nil {
		toSerialize["asm_serverless_agg_sum"] = o.AsmServerlessAggSum
	}
	if o.AuditLogsLinesIndexedAggSum != nil {
		toSerialize["audit_logs_lines_indexed_agg_sum"] = o.AuditLogsLinesIndexedAggSum
	}
	if o.AuditTrailEnabledHwmSum != nil {
		toSerialize["audit_trail_enabled_hwm_sum"] = o.AuditTrailEnabledHwmSum
	}
	if o.AvgProfiledFargateTasksSum != nil {
		toSerialize["avg_profiled_fargate_tasks_sum"] = o.AvgProfiledFargateTasksSum
	}
	if o.AwsHostTop99pSum != nil {
		toSerialize["aws_host_top99p_sum"] = o.AwsHostTop99pSum
	}
	if o.AwsLambdaFuncCount != nil {
		toSerialize["aws_lambda_func_count"] = o.AwsLambdaFuncCount
	}
	if o.AwsLambdaInvocationsSum != nil {
		toSerialize["aws_lambda_invocations_sum"] = o.AwsLambdaInvocationsSum
	}
	if o.AzureAppServiceTop99pSum != nil {
		toSerialize["azure_app_service_top99p_sum"] = o.AzureAppServiceTop99pSum
	}
	if o.AzureHostTop99pSum != nil {
		toSerialize["azure_host_top99p_sum"] = o.AzureHostTop99pSum
	}
	if o.BillableIngestedBytesAggSum != nil {
		toSerialize["billable_ingested_bytes_agg_sum"] = o.BillableIngestedBytesAggSum
	}
	if o.BitsAiInvestigationsAggSum != nil {
		toSerialize["bits_ai_investigations_agg_sum"] = o.BitsAiInvestigationsAggSum
	}
	if o.BrowserRumLiteSessionCountAggSum != nil {
		toSerialize["browser_rum_lite_session_count_agg_sum"] = o.BrowserRumLiteSessionCountAggSum
	}
	if o.BrowserRumReplaySessionCountAggSum != nil {
		toSerialize["browser_rum_replay_session_count_agg_sum"] = o.BrowserRumReplaySessionCountAggSum
	}
	if o.BrowserRumUnitsAggSum != nil {
		toSerialize["browser_rum_units_agg_sum"] = o.BrowserRumUnitsAggSum
	}
	if o.CiPipelineIndexedSpansAggSum != nil {
		toSerialize["ci_pipeline_indexed_spans_agg_sum"] = o.CiPipelineIndexedSpansAggSum
	}
	if o.CiTestIndexedSpansAggSum != nil {
		toSerialize["ci_test_indexed_spans_agg_sum"] = o.CiTestIndexedSpansAggSum
	}
	if o.CiVisibilityItrCommittersHwmSum != nil {
		toSerialize["ci_visibility_itr_committers_hwm_sum"] = o.CiVisibilityItrCommittersHwmSum
	}
	if o.CiVisibilityPipelineCommittersHwmSum != nil {
		toSerialize["ci_visibility_pipeline_committers_hwm_sum"] = o.CiVisibilityPipelineCommittersHwmSum
	}
	if o.CiVisibilityTestCommittersHwmSum != nil {
		toSerialize["ci_visibility_test_committers_hwm_sum"] = o.CiVisibilityTestCommittersHwmSum
	}
	if o.CloudCostManagementAwsHostCountAvgSum != nil {
		toSerialize["cloud_cost_management_aws_host_count_avg_sum"] = o.CloudCostManagementAwsHostCountAvgSum
	}
	if o.CloudCostManagementAzureHostCountAvgSum != nil {
		toSerialize["cloud_cost_management_azure_host_count_avg_sum"] = o.CloudCostManagementAzureHostCountAvgSum
	}
	if o.CloudCostManagementGcpHostCountAvgSum != nil {
		toSerialize["cloud_cost_management_gcp_host_count_avg_sum"] = o.CloudCostManagementGcpHostCountAvgSum
	}
	if o.CloudCostManagementHostCountAvgSum != nil {
		toSerialize["cloud_cost_management_host_count_avg_sum"] = o.CloudCostManagementHostCountAvgSum
	}
	if o.CloudCostManagementOciHostCountAvgSum != nil {
		toSerialize["cloud_cost_management_oci_host_count_avg_sum"] = o.CloudCostManagementOciHostCountAvgSum
	}
	if o.CloudSiemEventsAggSum != nil {
		toSerialize["cloud_siem_events_agg_sum"] = o.CloudSiemEventsAggSum
	}
	if o.CodeAnalysisSaCommittersHwmSum != nil {
		toSerialize["code_analysis_sa_committers_hwm_sum"] = o.CodeAnalysisSaCommittersHwmSum
	}
	if o.CodeAnalysisScaCommittersHwmSum != nil {
		toSerialize["code_analysis_sca_committers_hwm_sum"] = o.CodeAnalysisScaCommittersHwmSum
	}
	if o.CodeSecurityHostTop99pSum != nil {
		toSerialize["code_security_host_top99p_sum"] = o.CodeSecurityHostTop99pSum
	}
	if o.ContainerAvgSum != nil {
		toSerialize["container_avg_sum"] = o.ContainerAvgSum
	}
	if o.ContainerExclAgentAvgSum != nil {
		toSerialize["container_excl_agent_avg_sum"] = o.ContainerExclAgentAvgSum
	}
	if o.ContainerHwmSum != nil {
		toSerialize["container_hwm_sum"] = o.ContainerHwmSum
	}
	if o.CsmContainerEnterpriseComplianceCountAggSum != nil {
		toSerialize["csm_container_enterprise_compliance_count_agg_sum"] = o.CsmContainerEnterpriseComplianceCountAggSum
	}
	if o.CsmContainerEnterpriseCwsCountAggSum != nil {
		toSerialize["csm_container_enterprise_cws_count_agg_sum"] = o.CsmContainerEnterpriseCwsCountAggSum
	}
	if o.CsmContainerEnterpriseTotalCountAggSum != nil {
		toSerialize["csm_container_enterprise_total_count_agg_sum"] = o.CsmContainerEnterpriseTotalCountAggSum
	}
	if o.CsmHostEnterpriseAasHostCountTop99pSum != nil {
		toSerialize["csm_host_enterprise_aas_host_count_top99p_sum"] = o.CsmHostEnterpriseAasHostCountTop99pSum
	}
	if o.CsmHostEnterpriseAwsHostCountTop99pSum != nil {
		toSerialize["csm_host_enterprise_aws_host_count_top99p_sum"] = o.CsmHostEnterpriseAwsHostCountTop99pSum
	}
	if o.CsmHostEnterpriseAzureHostCountTop99pSum != nil {
		toSerialize["csm_host_enterprise_azure_host_count_top99p_sum"] = o.CsmHostEnterpriseAzureHostCountTop99pSum
	}
	if o.CsmHostEnterpriseComplianceHostCountTop99pSum != nil {
		toSerialize["csm_host_enterprise_compliance_host_count_top99p_sum"] = o.CsmHostEnterpriseComplianceHostCountTop99pSum
	}
	if o.CsmHostEnterpriseCwsHostCountTop99pSum != nil {
		toSerialize["csm_host_enterprise_cws_host_count_top99p_sum"] = o.CsmHostEnterpriseCwsHostCountTop99pSum
	}
	if o.CsmHostEnterpriseGcpHostCountTop99pSum != nil {
		toSerialize["csm_host_enterprise_gcp_host_count_top99p_sum"] = o.CsmHostEnterpriseGcpHostCountTop99pSum
	}
	if o.CsmHostEnterpriseTotalHostCountTop99pSum != nil {
		toSerialize["csm_host_enterprise_total_host_count_top99p_sum"] = o.CsmHostEnterpriseTotalHostCountTop99pSum
	}
	if o.CspmAasHostTop99pSum != nil {
		toSerialize["cspm_aas_host_top99p_sum"] = o.CspmAasHostTop99pSum
	}
	if o.CspmAwsHostTop99pSum != nil {
		toSerialize["cspm_aws_host_top99p_sum"] = o.CspmAwsHostTop99pSum
	}
	if o.CspmAzureHostTop99pSum != nil {
		toSerialize["cspm_azure_host_top99p_sum"] = o.CspmAzureHostTop99pSum
	}
	if o.CspmContainerAvgSum != nil {
		toSerialize["cspm_container_avg_sum"] = o.CspmContainerAvgSum
	}
	if o.CspmContainerHwmSum != nil {
		toSerialize["cspm_container_hwm_sum"] = o.CspmContainerHwmSum
	}
	if o.CspmGcpHostTop99pSum != nil {
		toSerialize["cspm_gcp_host_top99p_sum"] = o.CspmGcpHostTop99pSum
	}
	if o.CspmHostTop99pSum != nil {
		toSerialize["cspm_host_top99p_sum"] = o.CspmHostTop99pSum
	}
	if o.CustomHistoricalTsSum != nil {
		toSerialize["custom_historical_ts_sum"] = o.CustomHistoricalTsSum
	}
	if o.CustomLiveTsSum != nil {
		toSerialize["custom_live_ts_sum"] = o.CustomLiveTsSum
	}
	if o.CustomTsSum != nil {
		toSerialize["custom_ts_sum"] = o.CustomTsSum
	}
	if o.CwsContainerAvgSum != nil {
		toSerialize["cws_container_avg_sum"] = o.CwsContainerAvgSum
	}
	if o.CwsFargateTaskAvgSum != nil {
		toSerialize["cws_fargate_task_avg_sum"] = o.CwsFargateTaskAvgSum
	}
	if o.CwsHostTop99pSum != nil {
		toSerialize["cws_host_top99p_sum"] = o.CwsHostTop99pSum
	}
	if o.DataJobsMonitoringHostHrAggSum != nil {
		toSerialize["data_jobs_monitoring_host_hr_agg_sum"] = o.DataJobsMonitoringHostHrAggSum
	}
	if o.DbmHostTop99pSum != nil {
		toSerialize["dbm_host_top99p_sum"] = o.DbmHostTop99pSum
	}
	if o.DbmQueriesAvgSum != nil {
		toSerialize["dbm_queries_avg_sum"] = o.DbmQueriesAvgSum
	}
	if o.EndDate != nil {
		if o.EndDate.Nanosecond() == 0 {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EphInfraHostAgentAggSum != nil {
		toSerialize["eph_infra_host_agent_agg_sum"] = o.EphInfraHostAgentAggSum
	}
	if o.EphInfraHostAlibabaAggSum != nil {
		toSerialize["eph_infra_host_alibaba_agg_sum"] = o.EphInfraHostAlibabaAggSum
	}
	if o.EphInfraHostAwsAggSum != nil {
		toSerialize["eph_infra_host_aws_agg_sum"] = o.EphInfraHostAwsAggSum
	}
	if o.EphInfraHostAzureAggSum != nil {
		toSerialize["eph_infra_host_azure_agg_sum"] = o.EphInfraHostAzureAggSum
	}
	if o.EphInfraHostEntAggSum != nil {
		toSerialize["eph_infra_host_ent_agg_sum"] = o.EphInfraHostEntAggSum
	}
	if o.EphInfraHostGcpAggSum != nil {
		toSerialize["eph_infra_host_gcp_agg_sum"] = o.EphInfraHostGcpAggSum
	}
	if o.EphInfraHostHerokuAggSum != nil {
		toSerialize["eph_infra_host_heroku_agg_sum"] = o.EphInfraHostHerokuAggSum
	}
	if o.EphInfraHostOnlyAasAggSum != nil {
		toSerialize["eph_infra_host_only_aas_agg_sum"] = o.EphInfraHostOnlyAasAggSum
	}
	if o.EphInfraHostOnlyVsphereAggSum != nil {
		toSerialize["eph_infra_host_only_vsphere_agg_sum"] = o.EphInfraHostOnlyVsphereAggSum
	}
	if o.EphInfraHostOpentelemetryAggSum != nil {
		toSerialize["eph_infra_host_opentelemetry_agg_sum"] = o.EphInfraHostOpentelemetryAggSum
	}
	if o.EphInfraHostOpentelemetryApmAggSum != nil {
		toSerialize["eph_infra_host_opentelemetry_apm_agg_sum"] = o.EphInfraHostOpentelemetryApmAggSum
	}
	if o.EphInfraHostProAggSum != nil {
		toSerialize["eph_infra_host_pro_agg_sum"] = o.EphInfraHostProAggSum
	}
	if o.EphInfraHostProplusAggSum != nil {
		toSerialize["eph_infra_host_proplus_agg_sum"] = o.EphInfraHostProplusAggSum
	}
	if o.EphInfraHostProxmoxAggSum != nil {
		toSerialize["eph_infra_host_proxmox_agg_sum"] = o.EphInfraHostProxmoxAggSum
	}
	if o.ErrorTrackingApmErrorEventsAggSum != nil {
		toSerialize["error_tracking_apm_error_events_agg_sum"] = o.ErrorTrackingApmErrorEventsAggSum
	}
	if o.ErrorTrackingErrorEventsAggSum != nil {
		toSerialize["error_tracking_error_events_agg_sum"] = o.ErrorTrackingErrorEventsAggSum
	}
	if o.ErrorTrackingEventsAggSum != nil {
		toSerialize["error_tracking_events_agg_sum"] = o.ErrorTrackingEventsAggSum
	}
	if o.ErrorTrackingRumErrorEventsAggSum != nil {
		toSerialize["error_tracking_rum_error_events_agg_sum"] = o.ErrorTrackingRumErrorEventsAggSum
	}
	if o.EventManagementCorrelationAggSum != nil {
		toSerialize["event_management_correlation_agg_sum"] = o.EventManagementCorrelationAggSum
	}
	if o.EventManagementCorrelationCorrelatedEventsAggSum != nil {
		toSerialize["event_management_correlation_correlated_events_agg_sum"] = o.EventManagementCorrelationCorrelatedEventsAggSum
	}
	if o.EventManagementCorrelationCorrelatedRelatedEventsAggSum != nil {
		toSerialize["event_management_correlation_correlated_related_events_agg_sum"] = o.EventManagementCorrelationCorrelatedRelatedEventsAggSum
	}
	if o.FargateContainerProfilerProfilingFargateAvgSum != nil {
		toSerialize["fargate_container_profiler_profiling_fargate_avg_sum"] = o.FargateContainerProfilerProfilingFargateAvgSum
	}
	if o.FargateContainerProfilerProfilingFargateEksAvgSum != nil {
		toSerialize["fargate_container_profiler_profiling_fargate_eks_avg_sum"] = o.FargateContainerProfilerProfilingFargateEksAvgSum
	}
	if o.FargateTasksCountAvgSum != nil {
		toSerialize["fargate_tasks_count_avg_sum"] = o.FargateTasksCountAvgSum
	}
	if o.FargateTasksCountHwmSum != nil {
		toSerialize["fargate_tasks_count_hwm_sum"] = o.FargateTasksCountHwmSum
	}
	if o.FlexLogsComputeLargeAvgSum != nil {
		toSerialize["flex_logs_compute_large_avg_sum"] = o.FlexLogsComputeLargeAvgSum
	}
	if o.FlexLogsComputeMediumAvgSum != nil {
		toSerialize["flex_logs_compute_medium_avg_sum"] = o.FlexLogsComputeMediumAvgSum
	}
	if o.FlexLogsComputeSmallAvgSum != nil {
		toSerialize["flex_logs_compute_small_avg_sum"] = o.FlexLogsComputeSmallAvgSum
	}
	if o.FlexLogsComputeXlargeAvgSum != nil {
		toSerialize["flex_logs_compute_xlarge_avg_sum"] = o.FlexLogsComputeXlargeAvgSum
	}
	if o.FlexLogsComputeXsmallAvgSum != nil {
		toSerialize["flex_logs_compute_xsmall_avg_sum"] = o.FlexLogsComputeXsmallAvgSum
	}
	if o.FlexLogsStarterAvgSum != nil {
		toSerialize["flex_logs_starter_avg_sum"] = o.FlexLogsStarterAvgSum
	}
	if o.FlexLogsStarterStorageIndexAvgSum != nil {
		toSerialize["flex_logs_starter_storage_index_avg_sum"] = o.FlexLogsStarterStorageIndexAvgSum
	}
	if o.FlexLogsStarterStorageRetentionAdjustmentAvgSum != nil {
		toSerialize["flex_logs_starter_storage_retention_adjustment_avg_sum"] = o.FlexLogsStarterStorageRetentionAdjustmentAvgSum
	}
	if o.FlexStoredLogsAvgSum != nil {
		toSerialize["flex_stored_logs_avg_sum"] = o.FlexStoredLogsAvgSum
	}
	if o.ForwardingEventsBytesAggSum != nil {
		toSerialize["forwarding_events_bytes_agg_sum"] = o.ForwardingEventsBytesAggSum
	}
	if o.GcpHostTop99pSum != nil {
		toSerialize["gcp_host_top99p_sum"] = o.GcpHostTop99pSum
	}
	if o.HerokuHostTop99pSum != nil {
		toSerialize["heroku_host_top99p_sum"] = o.HerokuHostTop99pSum
	}
	if o.IncidentManagementMonthlyActiveUsersHwmSum != nil {
		toSerialize["incident_management_monthly_active_users_hwm_sum"] = o.IncidentManagementMonthlyActiveUsersHwmSum
	}
	if o.IncidentManagementSeatsHwmSum != nil {
		toSerialize["incident_management_seats_hwm_sum"] = o.IncidentManagementSeatsHwmSum
	}
	if o.IndexedEventsCountAggSum != nil {
		toSerialize["indexed_events_count_agg_sum"] = o.IndexedEventsCountAggSum
	}
	if o.InfraHostTop99pSum != nil {
		toSerialize["infra_host_top99p_sum"] = o.InfraHostTop99pSum
	}
	if o.IngestedEventsBytesAggSum != nil {
		toSerialize["ingested_events_bytes_agg_sum"] = o.IngestedEventsBytesAggSum
	}
	if o.IotDeviceAggSum != nil {
		toSerialize["iot_device_agg_sum"] = o.IotDeviceAggSum
	}
	if o.IotDeviceTop99pSum != nil {
		toSerialize["iot_device_top99p_sum"] = o.IotDeviceTop99pSum
	}
	if o.LastUpdated != nil {
		if o.LastUpdated.Nanosecond() == 0 {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_updated"] = o.LastUpdated.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LiveIndexedEventsAggSum != nil {
		toSerialize["live_indexed_events_agg_sum"] = o.LiveIndexedEventsAggSum
	}
	if o.LiveIngestedBytesAggSum != nil {
		toSerialize["live_ingested_bytes_agg_sum"] = o.LiveIngestedBytesAggSum
	}
	if o.LlmObservabilityAggSum != nil {
		toSerialize["llm_observability_agg_sum"] = o.LlmObservabilityAggSum
	}
	if o.LlmObservabilityMinSpendAggSum != nil {
		toSerialize["llm_observability_min_spend_agg_sum"] = o.LlmObservabilityMinSpendAggSum
	}
	if o.LogsByRetention != nil {
		toSerialize["logs_by_retention"] = o.LogsByRetention
	}
	if o.MobileRumLiteSessionCountAggSum != nil {
		toSerialize["mobile_rum_lite_session_count_agg_sum"] = o.MobileRumLiteSessionCountAggSum
	}
	if o.MobileRumSessionCountAggSum != nil {
		toSerialize["mobile_rum_session_count_agg_sum"] = o.MobileRumSessionCountAggSum
	}
	if o.MobileRumSessionCountAndroidAggSum != nil {
		toSerialize["mobile_rum_session_count_android_agg_sum"] = o.MobileRumSessionCountAndroidAggSum
	}
	if o.MobileRumSessionCountFlutterAggSum != nil {
		toSerialize["mobile_rum_session_count_flutter_agg_sum"] = o.MobileRumSessionCountFlutterAggSum
	}
	if o.MobileRumSessionCountIosAggSum != nil {
		toSerialize["mobile_rum_session_count_ios_agg_sum"] = o.MobileRumSessionCountIosAggSum
	}
	if o.MobileRumSessionCountReactnativeAggSum != nil {
		toSerialize["mobile_rum_session_count_reactnative_agg_sum"] = o.MobileRumSessionCountReactnativeAggSum
	}
	if o.MobileRumSessionCountRokuAggSum != nil {
		toSerialize["mobile_rum_session_count_roku_agg_sum"] = o.MobileRumSessionCountRokuAggSum
	}
	if o.MobileRumUnitsAggSum != nil {
		toSerialize["mobile_rum_units_agg_sum"] = o.MobileRumUnitsAggSum
	}
	if o.NdmNetflowEventsAggSum != nil {
		toSerialize["ndm_netflow_events_agg_sum"] = o.NdmNetflowEventsAggSum
	}
	if o.NetflowIndexedEventsCountAggSum != nil {
		toSerialize["netflow_indexed_events_count_agg_sum"] = o.NetflowIndexedEventsCountAggSum
	}
	if o.NetworkDeviceWirelessTop99pSum != nil {
		toSerialize["network_device_wireless_top99p_sum"] = o.NetworkDeviceWirelessTop99pSum
	}
	if o.NpmHostTop99pSum != nil {
		toSerialize["npm_host_top99p_sum"] = o.NpmHostTop99pSum
	}
	if o.ObservabilityPipelinesBytesProcessedAggSum != nil {
		toSerialize["observability_pipelines_bytes_processed_agg_sum"] = o.ObservabilityPipelinesBytesProcessedAggSum
	}
	if o.OciHostAggSum != nil {
		toSerialize["oci_host_agg_sum"] = o.OciHostAggSum
	}
	if o.OciHostTop99pSum != nil {
		toSerialize["oci_host_top99p_sum"] = o.OciHostTop99pSum
	}
	if o.OnCallSeatHwmSum != nil {
		toSerialize["on_call_seat_hwm_sum"] = o.OnCallSeatHwmSum
	}
	if o.OnlineArchiveEventsCountAggSum != nil {
		toSerialize["online_archive_events_count_agg_sum"] = o.OnlineArchiveEventsCountAggSum
	}
	if o.OpentelemetryApmHostTop99pSum != nil {
		toSerialize["opentelemetry_apm_host_top99p_sum"] = o.OpentelemetryApmHostTop99pSum
	}
	if o.OpentelemetryHostTop99pSum != nil {
		toSerialize["opentelemetry_host_top99p_sum"] = o.OpentelemetryHostTop99pSum
	}
	if o.ProductAnalyticsAggSum != nil {
		toSerialize["product_analytics_agg_sum"] = o.ProductAnalyticsAggSum
	}
	if o.ProfilingAasCountTop99pSum != nil {
		toSerialize["profiling_aas_count_top99p_sum"] = o.ProfilingAasCountTop99pSum
	}
	if o.ProfilingContainerAgentCountAvg != nil {
		toSerialize["profiling_container_agent_count_avg"] = o.ProfilingContainerAgentCountAvg
	}
	if o.ProfilingHostCountTop99pSum != nil {
		toSerialize["profiling_host_count_top99p_sum"] = o.ProfilingHostCountTop99pSum
	}
	if o.ProxmoxHostAggSum != nil {
		toSerialize["proxmox_host_agg_sum"] = o.ProxmoxHostAggSum
	}
	if o.ProxmoxHostTop99pSum != nil {
		toSerialize["proxmox_host_top99p_sum"] = o.ProxmoxHostTop99pSum
	}
	if o.PublishedAppHwmSum != nil {
		toSerialize["published_app_hwm_sum"] = o.PublishedAppHwmSum
	}
	if o.RehydratedIndexedEventsAggSum != nil {
		toSerialize["rehydrated_indexed_events_agg_sum"] = o.RehydratedIndexedEventsAggSum
	}
	if o.RehydratedIngestedBytesAggSum != nil {
		toSerialize["rehydrated_ingested_bytes_agg_sum"] = o.RehydratedIngestedBytesAggSum
	}
	if o.RumBrowserAndMobileSessionCount != nil {
		toSerialize["rum_browser_and_mobile_session_count"] = o.RumBrowserAndMobileSessionCount
	}
	if o.RumBrowserLegacySessionCountAggSum != nil {
		toSerialize["rum_browser_legacy_session_count_agg_sum"] = o.RumBrowserLegacySessionCountAggSum
	}
	if o.RumBrowserLiteSessionCountAggSum != nil {
		toSerialize["rum_browser_lite_session_count_agg_sum"] = o.RumBrowserLiteSessionCountAggSum
	}
	if o.RumBrowserReplaySessionCountAggSum != nil {
		toSerialize["rum_browser_replay_session_count_agg_sum"] = o.RumBrowserReplaySessionCountAggSum
	}
	if o.RumIndexedSessionsAggSum != nil {
		toSerialize["rum_indexed_sessions_agg_sum"] = o.RumIndexedSessionsAggSum
	}
	if o.RumIngestedSessionsAggSum != nil {
		toSerialize["rum_ingested_sessions_agg_sum"] = o.RumIngestedSessionsAggSum
	}
	if o.RumLiteSessionCountAggSum != nil {
		toSerialize["rum_lite_session_count_agg_sum"] = o.RumLiteSessionCountAggSum
	}
	if o.RumMobileLegacySessionCountAndroidAggSum != nil {
		toSerialize["rum_mobile_legacy_session_count_android_agg_sum"] = o.RumMobileLegacySessionCountAndroidAggSum
	}
	if o.RumMobileLegacySessionCountFlutterAggSum != nil {
		toSerialize["rum_mobile_legacy_session_count_flutter_agg_sum"] = o.RumMobileLegacySessionCountFlutterAggSum
	}
	if o.RumMobileLegacySessionCountIosAggSum != nil {
		toSerialize["rum_mobile_legacy_session_count_ios_agg_sum"] = o.RumMobileLegacySessionCountIosAggSum
	}
	if o.RumMobileLegacySessionCountReactnativeAggSum != nil {
		toSerialize["rum_mobile_legacy_session_count_reactnative_agg_sum"] = o.RumMobileLegacySessionCountReactnativeAggSum
	}
	if o.RumMobileLegacySessionCountRokuAggSum != nil {
		toSerialize["rum_mobile_legacy_session_count_roku_agg_sum"] = o.RumMobileLegacySessionCountRokuAggSum
	}
	if o.RumMobileLiteSessionCountAndroidAggSum != nil {
		toSerialize["rum_mobile_lite_session_count_android_agg_sum"] = o.RumMobileLiteSessionCountAndroidAggSum
	}
	if o.RumMobileLiteSessionCountFlutterAggSum != nil {
		toSerialize["rum_mobile_lite_session_count_flutter_agg_sum"] = o.RumMobileLiteSessionCountFlutterAggSum
	}
	if o.RumMobileLiteSessionCountIosAggSum != nil {
		toSerialize["rum_mobile_lite_session_count_ios_agg_sum"] = o.RumMobileLiteSessionCountIosAggSum
	}
	if o.RumMobileLiteSessionCountKotlinmultiplatformAggSum != nil {
		toSerialize["rum_mobile_lite_session_count_kotlinmultiplatform_agg_sum"] = o.RumMobileLiteSessionCountKotlinmultiplatformAggSum
	}
	if o.RumMobileLiteSessionCountReactnativeAggSum != nil {
		toSerialize["rum_mobile_lite_session_count_reactnative_agg_sum"] = o.RumMobileLiteSessionCountReactnativeAggSum
	}
	if o.RumMobileLiteSessionCountRokuAggSum != nil {
		toSerialize["rum_mobile_lite_session_count_roku_agg_sum"] = o.RumMobileLiteSessionCountRokuAggSum
	}
	if o.RumMobileLiteSessionCountUnityAggSum != nil {
		toSerialize["rum_mobile_lite_session_count_unity_agg_sum"] = o.RumMobileLiteSessionCountUnityAggSum
	}
	if o.RumMobileReplaySessionCountAndroidAggSum != nil {
		toSerialize["rum_mobile_replay_session_count_android_agg_sum"] = o.RumMobileReplaySessionCountAndroidAggSum
	}
	if o.RumMobileReplaySessionCountIosAggSum != nil {
		toSerialize["rum_mobile_replay_session_count_ios_agg_sum"] = o.RumMobileReplaySessionCountIosAggSum
	}
	if o.RumMobileReplaySessionCountKotlinmultiplatformAggSum != nil {
		toSerialize["rum_mobile_replay_session_count_kotlinmultiplatform_agg_sum"] = o.RumMobileReplaySessionCountKotlinmultiplatformAggSum
	}
	if o.RumMobileReplaySessionCountReactnativeAggSum != nil {
		toSerialize["rum_mobile_replay_session_count_reactnative_agg_sum"] = o.RumMobileReplaySessionCountReactnativeAggSum
	}
	if o.RumReplaySessionCountAggSum != nil {
		toSerialize["rum_replay_session_count_agg_sum"] = o.RumReplaySessionCountAggSum
	}
	if o.RumSessionCountAggSum != nil {
		toSerialize["rum_session_count_agg_sum"] = o.RumSessionCountAggSum
	}
	if o.RumSessionReplayAddOnAggSum != nil {
		toSerialize["rum_session_replay_add_on_agg_sum"] = o.RumSessionReplayAddOnAggSum
	}
	if o.RumTotalSessionCountAggSum != nil {
		toSerialize["rum_total_session_count_agg_sum"] = o.RumTotalSessionCountAggSum
	}
	if o.RumUnitsAggSum != nil {
		toSerialize["rum_units_agg_sum"] = o.RumUnitsAggSum
	}
	if o.ScaFargateCountAvgSum != nil {
		toSerialize["sca_fargate_count_avg_sum"] = o.ScaFargateCountAvgSum
	}
	if o.ScaFargateCountHwmSum != nil {
		toSerialize["sca_fargate_count_hwm_sum"] = o.ScaFargateCountHwmSum
	}
	if o.SdsApmScannedBytesSum != nil {
		toSerialize["sds_apm_scanned_bytes_sum"] = o.SdsApmScannedBytesSum
	}
	if o.SdsEventsScannedBytesSum != nil {
		toSerialize["sds_events_scanned_bytes_sum"] = o.SdsEventsScannedBytesSum
	}
	if o.SdsLogsScannedBytesSum != nil {
		toSerialize["sds_logs_scanned_bytes_sum"] = o.SdsLogsScannedBytesSum
	}
	if o.SdsRumScannedBytesSum != nil {
		toSerialize["sds_rum_scanned_bytes_sum"] = o.SdsRumScannedBytesSum
	}
	if o.SdsTotalScannedBytesSum != nil {
		toSerialize["sds_total_scanned_bytes_sum"] = o.SdsTotalScannedBytesSum
	}
	if o.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_apm_azure_appservice_instances_avg_sum"] = o.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum
	}
	if o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_apm_azure_azurefunction_instances_avg_sum"] = o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum
	}
	if o.ServerlessAppsApmApmAzureContainerappInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_apm_azure_containerapp_instances_avg_sum"] = o.ServerlessAppsApmApmAzureContainerappInstancesAvgSum
	}
	if o.ServerlessAppsApmApmFargateEcsTasksAvgSum != nil {
		toSerialize["serverless_apps_apm_apm_fargate_ecs_tasks_avg_sum"] = o.ServerlessAppsApmApmFargateEcsTasksAvgSum
	}
	if o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_apm_gcp_cloudfunction_instances_avg_sum"] = o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum
	}
	if o.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_apm_gcp_cloudrun_instances_avg_sum"] = o.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum
	}
	if o.ServerlessAppsApmAvgSum != nil {
		toSerialize["serverless_apps_apm_avg_sum"] = o.ServerlessAppsApmAvgSum
	}
	if o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg_sum"] = o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum
	}
	if o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg_sum"] = o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum
	}
	if o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg_sum"] = o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum
	}
	if o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg_sum"] = o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum
	}
	if o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg_sum"] = o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum
	}
	if o.ServerlessAppsApmExclFargateAvgSum != nil {
		toSerialize["serverless_apps_apm_excl_fargate_avg_sum"] = o.ServerlessAppsApmExclFargateAvgSum
	}
	if o.ServerlessAppsAzureContainerAppInstancesAvgSum != nil {
		toSerialize["serverless_apps_azure_container_app_instances_avg_sum"] = o.ServerlessAppsAzureContainerAppInstancesAvgSum
	}
	if o.ServerlessAppsAzureCountAvgSum != nil {
		toSerialize["serverless_apps_azure_count_avg_sum"] = o.ServerlessAppsAzureCountAvgSum
	}
	if o.ServerlessAppsAzureFunctionAppInstancesAvgSum != nil {
		toSerialize["serverless_apps_azure_function_app_instances_avg_sum"] = o.ServerlessAppsAzureFunctionAppInstancesAvgSum
	}
	if o.ServerlessAppsAzureWebAppInstancesAvgSum != nil {
		toSerialize["serverless_apps_azure_web_app_instances_avg_sum"] = o.ServerlessAppsAzureWebAppInstancesAvgSum
	}
	if o.ServerlessAppsEcsAvgSum != nil {
		toSerialize["serverless_apps_ecs_avg_sum"] = o.ServerlessAppsEcsAvgSum
	}
	if o.ServerlessAppsEksAvgSum != nil {
		toSerialize["serverless_apps_eks_avg_sum"] = o.ServerlessAppsEksAvgSum
	}
	if o.ServerlessAppsExclFargateAvgSum != nil {
		toSerialize["serverless_apps_excl_fargate_avg_sum"] = o.ServerlessAppsExclFargateAvgSum
	}
	if o.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum != nil {
		toSerialize["serverless_apps_excl_fargate_azure_container_app_instances_avg_sum"] = o.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum
	}
	if o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum != nil {
		toSerialize["serverless_apps_excl_fargate_azure_function_app_instances_avg_sum"] = o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum
	}
	if o.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum != nil {
		toSerialize["serverless_apps_excl_fargate_azure_web_app_instances_avg_sum"] = o.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum
	}
	if o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum != nil {
		toSerialize["serverless_apps_excl_fargate_google_cloud_functions_instances_avg_sum"] = o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum
	}
	if o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum != nil {
		toSerialize["serverless_apps_excl_fargate_google_cloud_run_instances_avg_sum"] = o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum
	}
	if o.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum != nil {
		toSerialize["serverless_apps_google_cloud_functions_instances_avg_sum"] = o.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum
	}
	if o.ServerlessAppsGoogleCloudRunInstancesAvgSum != nil {
		toSerialize["serverless_apps_google_cloud_run_instances_avg_sum"] = o.ServerlessAppsGoogleCloudRunInstancesAvgSum
	}
	if o.ServerlessAppsGoogleCountAvgSum != nil {
		toSerialize["serverless_apps_google_count_avg_sum"] = o.ServerlessAppsGoogleCountAvgSum
	}
	if o.ServerlessAppsTotalCountAvgSum != nil {
		toSerialize["serverless_apps_total_count_avg_sum"] = o.ServerlessAppsTotalCountAvgSum
	}
	if o.SiemAnalyzedLogsAddOnCountAggSum != nil {
		toSerialize["siem_analyzed_logs_add_on_count_agg_sum"] = o.SiemAnalyzedLogsAddOnCountAggSum
	}
	if o.StartDate != nil {
		if o.StartDate.Nanosecond() == 0 {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.SyntheticsBrowserCheckCallsCountAggSum != nil {
		toSerialize["synthetics_browser_check_calls_count_agg_sum"] = o.SyntheticsBrowserCheckCallsCountAggSum
	}
	if o.SyntheticsCheckCallsCountAggSum != nil {
		toSerialize["synthetics_check_calls_count_agg_sum"] = o.SyntheticsCheckCallsCountAggSum
	}
	if o.SyntheticsMobileTestRunsAggSum != nil {
		toSerialize["synthetics_mobile_test_runs_agg_sum"] = o.SyntheticsMobileTestRunsAggSum
	}
	if o.SyntheticsParallelTestingMaxSlotsHwmSum != nil {
		toSerialize["synthetics_parallel_testing_max_slots_hwm_sum"] = o.SyntheticsParallelTestingMaxSlotsHwmSum
	}
	if o.TraceSearchIndexedEventsCountAggSum != nil {
		toSerialize["trace_search_indexed_events_count_agg_sum"] = o.TraceSearchIndexedEventsCountAggSum
	}
	if o.TwolIngestedEventsBytesAggSum != nil {
		toSerialize["twol_ingested_events_bytes_agg_sum"] = o.TwolIngestedEventsBytesAggSum
	}
	if o.UniversalServiceMonitoringHostTop99pSum != nil {
		toSerialize["universal_service_monitoring_host_top99p_sum"] = o.UniversalServiceMonitoringHostTop99pSum
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	if o.VsphereHostTop99pSum != nil {
		toSerialize["vsphere_host_top99p_sum"] = o.VsphereHostTop99pSum
	}
	if o.VulnManagementHostCountTop99pSum != nil {
		toSerialize["vuln_management_host_count_top99p_sum"] = o.VulnManagementHostCountTop99pSum
	}
	if o.WorkflowExecutionsUsageAggSum != nil {
		toSerialize["workflow_executions_usage_agg_sum"] = o.WorkflowExecutionsUsageAggSum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageSummaryResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentHostTop99pSum                                               *int64             `json:"agent_host_top99p_sum,omitempty"`
		ApmAzureAppServiceHostTop99pSum                                  *int64             `json:"apm_azure_app_service_host_top99p_sum,omitempty"`
		ApmDevsecopsHostTop99pSum                                        *int64             `json:"apm_devsecops_host_top99p_sum,omitempty"`
		ApmEnterpriseStandaloneHostsTop99pSum                            *int64             `json:"apm_enterprise_standalone_hosts_top99p_sum,omitempty"`
		ApmFargateCountAvgSum                                            *int64             `json:"apm_fargate_count_avg_sum,omitempty"`
		ApmHostTop99pSum                                                 *int64             `json:"apm_host_top99p_sum,omitempty"`
		ApmProStandaloneHostsTop99pSum                                   *int64             `json:"apm_pro_standalone_hosts_top99p_sum,omitempty"`
		AppsecFargateCountAvgSum                                         *int64             `json:"appsec_fargate_count_avg_sum,omitempty"`
		AsmServerlessAggSum                                              *int64             `json:"asm_serverless_agg_sum,omitempty"`
		AuditLogsLinesIndexedAggSum                                      *int64             `json:"audit_logs_lines_indexed_agg_sum,omitempty"`
		AuditTrailEnabledHwmSum                                          *int64             `json:"audit_trail_enabled_hwm_sum,omitempty"`
		AvgProfiledFargateTasksSum                                       *int64             `json:"avg_profiled_fargate_tasks_sum,omitempty"`
		AwsHostTop99pSum                                                 *int64             `json:"aws_host_top99p_sum,omitempty"`
		AwsLambdaFuncCount                                               *int64             `json:"aws_lambda_func_count,omitempty"`
		AwsLambdaInvocationsSum                                          *int64             `json:"aws_lambda_invocations_sum,omitempty"`
		AzureAppServiceTop99pSum                                         *int64             `json:"azure_app_service_top99p_sum,omitempty"`
		AzureHostTop99pSum                                               *int64             `json:"azure_host_top99p_sum,omitempty"`
		BillableIngestedBytesAggSum                                      *int64             `json:"billable_ingested_bytes_agg_sum,omitempty"`
		BitsAiInvestigationsAggSum                                       *int64             `json:"bits_ai_investigations_agg_sum,omitempty"`
		BrowserRumLiteSessionCountAggSum                                 *int64             `json:"browser_rum_lite_session_count_agg_sum,omitempty"`
		BrowserRumReplaySessionCountAggSum                               *int64             `json:"browser_rum_replay_session_count_agg_sum,omitempty"`
		BrowserRumUnitsAggSum                                            *int64             `json:"browser_rum_units_agg_sum,omitempty"`
		CiPipelineIndexedSpansAggSum                                     *int64             `json:"ci_pipeline_indexed_spans_agg_sum,omitempty"`
		CiTestIndexedSpansAggSum                                         *int64             `json:"ci_test_indexed_spans_agg_sum,omitempty"`
		CiVisibilityItrCommittersHwmSum                                  *int64             `json:"ci_visibility_itr_committers_hwm_sum,omitempty"`
		CiVisibilityPipelineCommittersHwmSum                             *int64             `json:"ci_visibility_pipeline_committers_hwm_sum,omitempty"`
		CiVisibilityTestCommittersHwmSum                                 *int64             `json:"ci_visibility_test_committers_hwm_sum,omitempty"`
		CloudCostManagementAwsHostCountAvgSum                            *int64             `json:"cloud_cost_management_aws_host_count_avg_sum,omitempty"`
		CloudCostManagementAzureHostCountAvgSum                          *int64             `json:"cloud_cost_management_azure_host_count_avg_sum,omitempty"`
		CloudCostManagementGcpHostCountAvgSum                            *int64             `json:"cloud_cost_management_gcp_host_count_avg_sum,omitempty"`
		CloudCostManagementHostCountAvgSum                               *int64             `json:"cloud_cost_management_host_count_avg_sum,omitempty"`
		CloudCostManagementOciHostCountAvgSum                            *int64             `json:"cloud_cost_management_oci_host_count_avg_sum,omitempty"`
		CloudSiemEventsAggSum                                            *int64             `json:"cloud_siem_events_agg_sum,omitempty"`
		CodeAnalysisSaCommittersHwmSum                                   *int64             `json:"code_analysis_sa_committers_hwm_sum,omitempty"`
		CodeAnalysisScaCommittersHwmSum                                  *int64             `json:"code_analysis_sca_committers_hwm_sum,omitempty"`
		CodeSecurityHostTop99pSum                                        *int64             `json:"code_security_host_top99p_sum,omitempty"`
		ContainerAvgSum                                                  *int64             `json:"container_avg_sum,omitempty"`
		ContainerExclAgentAvgSum                                         *int64             `json:"container_excl_agent_avg_sum,omitempty"`
		ContainerHwmSum                                                  *int64             `json:"container_hwm_sum,omitempty"`
		CsmContainerEnterpriseComplianceCountAggSum                      *int64             `json:"csm_container_enterprise_compliance_count_agg_sum,omitempty"`
		CsmContainerEnterpriseCwsCountAggSum                             *int64             `json:"csm_container_enterprise_cws_count_agg_sum,omitempty"`
		CsmContainerEnterpriseTotalCountAggSum                           *int64             `json:"csm_container_enterprise_total_count_agg_sum,omitempty"`
		CsmHostEnterpriseAasHostCountTop99pSum                           *int64             `json:"csm_host_enterprise_aas_host_count_top99p_sum,omitempty"`
		CsmHostEnterpriseAwsHostCountTop99pSum                           *int64             `json:"csm_host_enterprise_aws_host_count_top99p_sum,omitempty"`
		CsmHostEnterpriseAzureHostCountTop99pSum                         *int64             `json:"csm_host_enterprise_azure_host_count_top99p_sum,omitempty"`
		CsmHostEnterpriseComplianceHostCountTop99pSum                    *int64             `json:"csm_host_enterprise_compliance_host_count_top99p_sum,omitempty"`
		CsmHostEnterpriseCwsHostCountTop99pSum                           *int64             `json:"csm_host_enterprise_cws_host_count_top99p_sum,omitempty"`
		CsmHostEnterpriseGcpHostCountTop99pSum                           *int64             `json:"csm_host_enterprise_gcp_host_count_top99p_sum,omitempty"`
		CsmHostEnterpriseTotalHostCountTop99pSum                         *int64             `json:"csm_host_enterprise_total_host_count_top99p_sum,omitempty"`
		CspmAasHostTop99pSum                                             *int64             `json:"cspm_aas_host_top99p_sum,omitempty"`
		CspmAwsHostTop99pSum                                             *int64             `json:"cspm_aws_host_top99p_sum,omitempty"`
		CspmAzureHostTop99pSum                                           *int64             `json:"cspm_azure_host_top99p_sum,omitempty"`
		CspmContainerAvgSum                                              *int64             `json:"cspm_container_avg_sum,omitempty"`
		CspmContainerHwmSum                                              *int64             `json:"cspm_container_hwm_sum,omitempty"`
		CspmGcpHostTop99pSum                                             *int64             `json:"cspm_gcp_host_top99p_sum,omitempty"`
		CspmHostTop99pSum                                                *int64             `json:"cspm_host_top99p_sum,omitempty"`
		CustomHistoricalTsSum                                            *int64             `json:"custom_historical_ts_sum,omitempty"`
		CustomLiveTsSum                                                  *int64             `json:"custom_live_ts_sum,omitempty"`
		CustomTsSum                                                      *int64             `json:"custom_ts_sum,omitempty"`
		CwsContainerAvgSum                                               *int64             `json:"cws_container_avg_sum,omitempty"`
		CwsFargateTaskAvgSum                                             *int64             `json:"cws_fargate_task_avg_sum,omitempty"`
		CwsHostTop99pSum                                                 *int64             `json:"cws_host_top99p_sum,omitempty"`
		DataJobsMonitoringHostHrAggSum                                   *int64             `json:"data_jobs_monitoring_host_hr_agg_sum,omitempty"`
		DbmHostTop99pSum                                                 *int64             `json:"dbm_host_top99p_sum,omitempty"`
		DbmQueriesAvgSum                                                 *int64             `json:"dbm_queries_avg_sum,omitempty"`
		EndDate                                                          *time.Time         `json:"end_date,omitempty"`
		EphInfraHostAgentAggSum                                          *int64             `json:"eph_infra_host_agent_agg_sum,omitempty"`
		EphInfraHostAlibabaAggSum                                        *int64             `json:"eph_infra_host_alibaba_agg_sum,omitempty"`
		EphInfraHostAwsAggSum                                            *int64             `json:"eph_infra_host_aws_agg_sum,omitempty"`
		EphInfraHostAzureAggSum                                          *int64             `json:"eph_infra_host_azure_agg_sum,omitempty"`
		EphInfraHostEntAggSum                                            *int64             `json:"eph_infra_host_ent_agg_sum,omitempty"`
		EphInfraHostGcpAggSum                                            *int64             `json:"eph_infra_host_gcp_agg_sum,omitempty"`
		EphInfraHostHerokuAggSum                                         *int64             `json:"eph_infra_host_heroku_agg_sum,omitempty"`
		EphInfraHostOnlyAasAggSum                                        *int64             `json:"eph_infra_host_only_aas_agg_sum,omitempty"`
		EphInfraHostOnlyVsphereAggSum                                    *int64             `json:"eph_infra_host_only_vsphere_agg_sum,omitempty"`
		EphInfraHostOpentelemetryAggSum                                  *int64             `json:"eph_infra_host_opentelemetry_agg_sum,omitempty"`
		EphInfraHostOpentelemetryApmAggSum                               *int64             `json:"eph_infra_host_opentelemetry_apm_agg_sum,omitempty"`
		EphInfraHostProAggSum                                            *int64             `json:"eph_infra_host_pro_agg_sum,omitempty"`
		EphInfraHostProplusAggSum                                        *int64             `json:"eph_infra_host_proplus_agg_sum,omitempty"`
		EphInfraHostProxmoxAggSum                                        *int64             `json:"eph_infra_host_proxmox_agg_sum,omitempty"`
		ErrorTrackingApmErrorEventsAggSum                                *int64             `json:"error_tracking_apm_error_events_agg_sum,omitempty"`
		ErrorTrackingErrorEventsAggSum                                   *int64             `json:"error_tracking_error_events_agg_sum,omitempty"`
		ErrorTrackingEventsAggSum                                        *int64             `json:"error_tracking_events_agg_sum,omitempty"`
		ErrorTrackingRumErrorEventsAggSum                                *int64             `json:"error_tracking_rum_error_events_agg_sum,omitempty"`
		EventManagementCorrelationAggSum                                 *int64             `json:"event_management_correlation_agg_sum,omitempty"`
		EventManagementCorrelationCorrelatedEventsAggSum                 *int64             `json:"event_management_correlation_correlated_events_agg_sum,omitempty"`
		EventManagementCorrelationCorrelatedRelatedEventsAggSum          *int64             `json:"event_management_correlation_correlated_related_events_agg_sum,omitempty"`
		FargateContainerProfilerProfilingFargateAvgSum                   *int64             `json:"fargate_container_profiler_profiling_fargate_avg_sum,omitempty"`
		FargateContainerProfilerProfilingFargateEksAvgSum                *int64             `json:"fargate_container_profiler_profiling_fargate_eks_avg_sum,omitempty"`
		FargateTasksCountAvgSum                                          *int64             `json:"fargate_tasks_count_avg_sum,omitempty"`
		FargateTasksCountHwmSum                                          *int64             `json:"fargate_tasks_count_hwm_sum,omitempty"`
		FlexLogsComputeLargeAvgSum                                       *int64             `json:"flex_logs_compute_large_avg_sum,omitempty"`
		FlexLogsComputeMediumAvgSum                                      *int64             `json:"flex_logs_compute_medium_avg_sum,omitempty"`
		FlexLogsComputeSmallAvgSum                                       *int64             `json:"flex_logs_compute_small_avg_sum,omitempty"`
		FlexLogsComputeXlargeAvgSum                                      *int64             `json:"flex_logs_compute_xlarge_avg_sum,omitempty"`
		FlexLogsComputeXsmallAvgSum                                      *int64             `json:"flex_logs_compute_xsmall_avg_sum,omitempty"`
		FlexLogsStarterAvgSum                                            *int64             `json:"flex_logs_starter_avg_sum,omitempty"`
		FlexLogsStarterStorageIndexAvgSum                                *int64             `json:"flex_logs_starter_storage_index_avg_sum,omitempty"`
		FlexLogsStarterStorageRetentionAdjustmentAvgSum                  *int64             `json:"flex_logs_starter_storage_retention_adjustment_avg_sum,omitempty"`
		FlexStoredLogsAvgSum                                             *int64             `json:"flex_stored_logs_avg_sum,omitempty"`
		ForwardingEventsBytesAggSum                                      *int64             `json:"forwarding_events_bytes_agg_sum,omitempty"`
		GcpHostTop99pSum                                                 *int64             `json:"gcp_host_top99p_sum,omitempty"`
		HerokuHostTop99pSum                                              *int64             `json:"heroku_host_top99p_sum,omitempty"`
		IncidentManagementMonthlyActiveUsersHwmSum                       *int64             `json:"incident_management_monthly_active_users_hwm_sum,omitempty"`
		IncidentManagementSeatsHwmSum                                    *int64             `json:"incident_management_seats_hwm_sum,omitempty"`
		IndexedEventsCountAggSum                                         *int64             `json:"indexed_events_count_agg_sum,omitempty"`
		InfraHostTop99pSum                                               *int64             `json:"infra_host_top99p_sum,omitempty"`
		IngestedEventsBytesAggSum                                        *int64             `json:"ingested_events_bytes_agg_sum,omitempty"`
		IotDeviceAggSum                                                  *int64             `json:"iot_device_agg_sum,omitempty"`
		IotDeviceTop99pSum                                               *int64             `json:"iot_device_top99p_sum,omitempty"`
		LastUpdated                                                      *time.Time         `json:"last_updated,omitempty"`
		LiveIndexedEventsAggSum                                          *int64             `json:"live_indexed_events_agg_sum,omitempty"`
		LiveIngestedBytesAggSum                                          *int64             `json:"live_ingested_bytes_agg_sum,omitempty"`
		LlmObservabilityAggSum                                           *int64             `json:"llm_observability_agg_sum,omitempty"`
		LlmObservabilityMinSpendAggSum                                   *int64             `json:"llm_observability_min_spend_agg_sum,omitempty"`
		LogsByRetention                                                  *LogsByRetention   `json:"logs_by_retention,omitempty"`
		MobileRumLiteSessionCountAggSum                                  *int64             `json:"mobile_rum_lite_session_count_agg_sum,omitempty"`
		MobileRumSessionCountAggSum                                      *int64             `json:"mobile_rum_session_count_agg_sum,omitempty"`
		MobileRumSessionCountAndroidAggSum                               *int64             `json:"mobile_rum_session_count_android_agg_sum,omitempty"`
		MobileRumSessionCountFlutterAggSum                               *int64             `json:"mobile_rum_session_count_flutter_agg_sum,omitempty"`
		MobileRumSessionCountIosAggSum                                   *int64             `json:"mobile_rum_session_count_ios_agg_sum,omitempty"`
		MobileRumSessionCountReactnativeAggSum                           *int64             `json:"mobile_rum_session_count_reactnative_agg_sum,omitempty"`
		MobileRumSessionCountRokuAggSum                                  *int64             `json:"mobile_rum_session_count_roku_agg_sum,omitempty"`
		MobileRumUnitsAggSum                                             *int64             `json:"mobile_rum_units_agg_sum,omitempty"`
		NdmNetflowEventsAggSum                                           *int64             `json:"ndm_netflow_events_agg_sum,omitempty"`
		NetflowIndexedEventsCountAggSum                                  *int64             `json:"netflow_indexed_events_count_agg_sum,omitempty"`
		NetworkDeviceWirelessTop99pSum                                   *int64             `json:"network_device_wireless_top99p_sum,omitempty"`
		NpmHostTop99pSum                                                 *int64             `json:"npm_host_top99p_sum,omitempty"`
		ObservabilityPipelinesBytesProcessedAggSum                       *int64             `json:"observability_pipelines_bytes_processed_agg_sum,omitempty"`
		OciHostAggSum                                                    *int64             `json:"oci_host_agg_sum,omitempty"`
		OciHostTop99pSum                                                 *int64             `json:"oci_host_top99p_sum,omitempty"`
		OnCallSeatHwmSum                                                 *int64             `json:"on_call_seat_hwm_sum,omitempty"`
		OnlineArchiveEventsCountAggSum                                   *int64             `json:"online_archive_events_count_agg_sum,omitempty"`
		OpentelemetryApmHostTop99pSum                                    *int64             `json:"opentelemetry_apm_host_top99p_sum,omitempty"`
		OpentelemetryHostTop99pSum                                       *int64             `json:"opentelemetry_host_top99p_sum,omitempty"`
		ProductAnalyticsAggSum                                           *int64             `json:"product_analytics_agg_sum,omitempty"`
		ProfilingAasCountTop99pSum                                       *int64             `json:"profiling_aas_count_top99p_sum,omitempty"`
		ProfilingContainerAgentCountAvg                                  *int64             `json:"profiling_container_agent_count_avg,omitempty"`
		ProfilingHostCountTop99pSum                                      *int64             `json:"profiling_host_count_top99p_sum,omitempty"`
		ProxmoxHostAggSum                                                *int64             `json:"proxmox_host_agg_sum,omitempty"`
		ProxmoxHostTop99pSum                                             *int64             `json:"proxmox_host_top99p_sum,omitempty"`
		PublishedAppHwmSum                                               *int64             `json:"published_app_hwm_sum,omitempty"`
		RehydratedIndexedEventsAggSum                                    *int64             `json:"rehydrated_indexed_events_agg_sum,omitempty"`
		RehydratedIngestedBytesAggSum                                    *int64             `json:"rehydrated_ingested_bytes_agg_sum,omitempty"`
		RumBrowserAndMobileSessionCount                                  *int64             `json:"rum_browser_and_mobile_session_count,omitempty"`
		RumBrowserLegacySessionCountAggSum                               *int64             `json:"rum_browser_legacy_session_count_agg_sum,omitempty"`
		RumBrowserLiteSessionCountAggSum                                 *int64             `json:"rum_browser_lite_session_count_agg_sum,omitempty"`
		RumBrowserReplaySessionCountAggSum                               *int64             `json:"rum_browser_replay_session_count_agg_sum,omitempty"`
		RumIndexedSessionsAggSum                                         *int64             `json:"rum_indexed_sessions_agg_sum,omitempty"`
		RumIngestedSessionsAggSum                                        *int64             `json:"rum_ingested_sessions_agg_sum,omitempty"`
		RumLiteSessionCountAggSum                                        *int64             `json:"rum_lite_session_count_agg_sum,omitempty"`
		RumMobileLegacySessionCountAndroidAggSum                         *int64             `json:"rum_mobile_legacy_session_count_android_agg_sum,omitempty"`
		RumMobileLegacySessionCountFlutterAggSum                         *int64             `json:"rum_mobile_legacy_session_count_flutter_agg_sum,omitempty"`
		RumMobileLegacySessionCountIosAggSum                             *int64             `json:"rum_mobile_legacy_session_count_ios_agg_sum,omitempty"`
		RumMobileLegacySessionCountReactnativeAggSum                     *int64             `json:"rum_mobile_legacy_session_count_reactnative_agg_sum,omitempty"`
		RumMobileLegacySessionCountRokuAggSum                            *int64             `json:"rum_mobile_legacy_session_count_roku_agg_sum,omitempty"`
		RumMobileLiteSessionCountAndroidAggSum                           *int64             `json:"rum_mobile_lite_session_count_android_agg_sum,omitempty"`
		RumMobileLiteSessionCountFlutterAggSum                           *int64             `json:"rum_mobile_lite_session_count_flutter_agg_sum,omitempty"`
		RumMobileLiteSessionCountIosAggSum                               *int64             `json:"rum_mobile_lite_session_count_ios_agg_sum,omitempty"`
		RumMobileLiteSessionCountKotlinmultiplatformAggSum               *int64             `json:"rum_mobile_lite_session_count_kotlinmultiplatform_agg_sum,omitempty"`
		RumMobileLiteSessionCountReactnativeAggSum                       *int64             `json:"rum_mobile_lite_session_count_reactnative_agg_sum,omitempty"`
		RumMobileLiteSessionCountRokuAggSum                              *int64             `json:"rum_mobile_lite_session_count_roku_agg_sum,omitempty"`
		RumMobileLiteSessionCountUnityAggSum                             *int64             `json:"rum_mobile_lite_session_count_unity_agg_sum,omitempty"`
		RumMobileReplaySessionCountAndroidAggSum                         *int64             `json:"rum_mobile_replay_session_count_android_agg_sum,omitempty"`
		RumMobileReplaySessionCountIosAggSum                             *int64             `json:"rum_mobile_replay_session_count_ios_agg_sum,omitempty"`
		RumMobileReplaySessionCountKotlinmultiplatformAggSum             *int64             `json:"rum_mobile_replay_session_count_kotlinmultiplatform_agg_sum,omitempty"`
		RumMobileReplaySessionCountReactnativeAggSum                     *int64             `json:"rum_mobile_replay_session_count_reactnative_agg_sum,omitempty"`
		RumReplaySessionCountAggSum                                      *int64             `json:"rum_replay_session_count_agg_sum,omitempty"`
		RumSessionCountAggSum                                            *int64             `json:"rum_session_count_agg_sum,omitempty"`
		RumSessionReplayAddOnAggSum                                      *int64             `json:"rum_session_replay_add_on_agg_sum,omitempty"`
		RumTotalSessionCountAggSum                                       *int64             `json:"rum_total_session_count_agg_sum,omitempty"`
		RumUnitsAggSum                                                   *int64             `json:"rum_units_agg_sum,omitempty"`
		ScaFargateCountAvgSum                                            *int64             `json:"sca_fargate_count_avg_sum,omitempty"`
		ScaFargateCountHwmSum                                            *int64             `json:"sca_fargate_count_hwm_sum,omitempty"`
		SdsApmScannedBytesSum                                            *int64             `json:"sds_apm_scanned_bytes_sum,omitempty"`
		SdsEventsScannedBytesSum                                         *int64             `json:"sds_events_scanned_bytes_sum,omitempty"`
		SdsLogsScannedBytesSum                                           *int64             `json:"sds_logs_scanned_bytes_sum,omitempty"`
		SdsRumScannedBytesSum                                            *int64             `json:"sds_rum_scanned_bytes_sum,omitempty"`
		SdsTotalScannedBytesSum                                          *int64             `json:"sds_total_scanned_bytes_sum,omitempty"`
		ServerlessAppsApmApmAzureAppserviceInstancesAvgSum               *int64             `json:"serverless_apps_apm_apm_azure_appservice_instances_avg_sum,omitempty"`
		ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum            *int64             `json:"serverless_apps_apm_apm_azure_azurefunction_instances_avg_sum,omitempty"`
		ServerlessAppsApmApmAzureContainerappInstancesAvgSum             *int64             `json:"serverless_apps_apm_apm_azure_containerapp_instances_avg_sum,omitempty"`
		ServerlessAppsApmApmFargateEcsTasksAvgSum                        *int64             `json:"serverless_apps_apm_apm_fargate_ecs_tasks_avg_sum,omitempty"`
		ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum              *int64             `json:"serverless_apps_apm_apm_gcp_cloudfunction_instances_avg_sum,omitempty"`
		ServerlessAppsApmApmGcpCloudrunInstancesAvgSum                   *int64             `json:"serverless_apps_apm_apm_gcp_cloudrun_instances_avg_sum,omitempty"`
		ServerlessAppsApmAvgSum                                          *int64             `json:"serverless_apps_apm_avg_sum,omitempty"`
		ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum    *int64             `json:"serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg_sum,omitempty"`
		ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum *int64             `json:"serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg_sum,omitempty"`
		ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum  *int64             `json:"serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg_sum,omitempty"`
		ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum   *int64             `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg_sum,omitempty"`
		ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum        *int64             `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg_sum,omitempty"`
		ServerlessAppsApmExclFargateAvgSum                               *int64             `json:"serverless_apps_apm_excl_fargate_avg_sum,omitempty"`
		ServerlessAppsAzureContainerAppInstancesAvgSum                   *int64             `json:"serverless_apps_azure_container_app_instances_avg_sum,omitempty"`
		ServerlessAppsAzureCountAvgSum                                   *int64             `json:"serverless_apps_azure_count_avg_sum,omitempty"`
		ServerlessAppsAzureFunctionAppInstancesAvgSum                    *int64             `json:"serverless_apps_azure_function_app_instances_avg_sum,omitempty"`
		ServerlessAppsAzureWebAppInstancesAvgSum                         *int64             `json:"serverless_apps_azure_web_app_instances_avg_sum,omitempty"`
		ServerlessAppsEcsAvgSum                                          *int64             `json:"serverless_apps_ecs_avg_sum,omitempty"`
		ServerlessAppsEksAvgSum                                          *int64             `json:"serverless_apps_eks_avg_sum,omitempty"`
		ServerlessAppsExclFargateAvgSum                                  *int64             `json:"serverless_apps_excl_fargate_avg_sum,omitempty"`
		ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum        *int64             `json:"serverless_apps_excl_fargate_azure_container_app_instances_avg_sum,omitempty"`
		ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum         *int64             `json:"serverless_apps_excl_fargate_azure_function_app_instances_avg_sum,omitempty"`
		ServerlessAppsExclFargateAzureWebAppInstancesAvgSum              *int64             `json:"serverless_apps_excl_fargate_azure_web_app_instances_avg_sum,omitempty"`
		ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum     *int64             `json:"serverless_apps_excl_fargate_google_cloud_functions_instances_avg_sum,omitempty"`
		ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum           *int64             `json:"serverless_apps_excl_fargate_google_cloud_run_instances_avg_sum,omitempty"`
		ServerlessAppsGoogleCloudFunctionsInstancesAvgSum                *int64             `json:"serverless_apps_google_cloud_functions_instances_avg_sum,omitempty"`
		ServerlessAppsGoogleCloudRunInstancesAvgSum                      *int64             `json:"serverless_apps_google_cloud_run_instances_avg_sum,omitempty"`
		ServerlessAppsGoogleCountAvgSum                                  *int64             `json:"serverless_apps_google_count_avg_sum,omitempty"`
		ServerlessAppsTotalCountAvgSum                                   *int64             `json:"serverless_apps_total_count_avg_sum,omitempty"`
		SiemAnalyzedLogsAddOnCountAggSum                                 *int64             `json:"siem_analyzed_logs_add_on_count_agg_sum,omitempty"`
		StartDate                                                        *time.Time         `json:"start_date,omitempty"`
		SyntheticsBrowserCheckCallsCountAggSum                           *int64             `json:"synthetics_browser_check_calls_count_agg_sum,omitempty"`
		SyntheticsCheckCallsCountAggSum                                  *int64             `json:"synthetics_check_calls_count_agg_sum,omitempty"`
		SyntheticsMobileTestRunsAggSum                                   *int64             `json:"synthetics_mobile_test_runs_agg_sum,omitempty"`
		SyntheticsParallelTestingMaxSlotsHwmSum                          *int64             `json:"synthetics_parallel_testing_max_slots_hwm_sum,omitempty"`
		TraceSearchIndexedEventsCountAggSum                              *int64             `json:"trace_search_indexed_events_count_agg_sum,omitempty"`
		TwolIngestedEventsBytesAggSum                                    *int64             `json:"twol_ingested_events_bytes_agg_sum,omitempty"`
		UniversalServiceMonitoringHostTop99pSum                          *int64             `json:"universal_service_monitoring_host_top99p_sum,omitempty"`
		Usage                                                            []UsageSummaryDate `json:"usage,omitempty"`
		VsphereHostTop99pSum                                             *int64             `json:"vsphere_host_top99p_sum,omitempty"`
		VulnManagementHostCountTop99pSum                                 *int64             `json:"vuln_management_host_count_top99p_sum,omitempty"`
		WorkflowExecutionsUsageAggSum                                    *int64             `json:"workflow_executions_usage_agg_sum,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_host_top99p_sum", "apm_azure_app_service_host_top99p_sum", "apm_devsecops_host_top99p_sum", "apm_enterprise_standalone_hosts_top99p_sum", "apm_fargate_count_avg_sum", "apm_host_top99p_sum", "apm_pro_standalone_hosts_top99p_sum", "appsec_fargate_count_avg_sum", "asm_serverless_agg_sum", "audit_logs_lines_indexed_agg_sum", "audit_trail_enabled_hwm_sum", "avg_profiled_fargate_tasks_sum", "aws_host_top99p_sum", "aws_lambda_func_count", "aws_lambda_invocations_sum", "azure_app_service_top99p_sum", "azure_host_top99p_sum", "billable_ingested_bytes_agg_sum", "bits_ai_investigations_agg_sum", "browser_rum_lite_session_count_agg_sum", "browser_rum_replay_session_count_agg_sum", "browser_rum_units_agg_sum", "ci_pipeline_indexed_spans_agg_sum", "ci_test_indexed_spans_agg_sum", "ci_visibility_itr_committers_hwm_sum", "ci_visibility_pipeline_committers_hwm_sum", "ci_visibility_test_committers_hwm_sum", "cloud_cost_management_aws_host_count_avg_sum", "cloud_cost_management_azure_host_count_avg_sum", "cloud_cost_management_gcp_host_count_avg_sum", "cloud_cost_management_host_count_avg_sum", "cloud_cost_management_oci_host_count_avg_sum", "cloud_siem_events_agg_sum", "code_analysis_sa_committers_hwm_sum", "code_analysis_sca_committers_hwm_sum", "code_security_host_top99p_sum", "container_avg_sum", "container_excl_agent_avg_sum", "container_hwm_sum", "csm_container_enterprise_compliance_count_agg_sum", "csm_container_enterprise_cws_count_agg_sum", "csm_container_enterprise_total_count_agg_sum", "csm_host_enterprise_aas_host_count_top99p_sum", "csm_host_enterprise_aws_host_count_top99p_sum", "csm_host_enterprise_azure_host_count_top99p_sum", "csm_host_enterprise_compliance_host_count_top99p_sum", "csm_host_enterprise_cws_host_count_top99p_sum", "csm_host_enterprise_gcp_host_count_top99p_sum", "csm_host_enterprise_total_host_count_top99p_sum", "cspm_aas_host_top99p_sum", "cspm_aws_host_top99p_sum", "cspm_azure_host_top99p_sum", "cspm_container_avg_sum", "cspm_container_hwm_sum", "cspm_gcp_host_top99p_sum", "cspm_host_top99p_sum", "custom_historical_ts_sum", "custom_live_ts_sum", "custom_ts_sum", "cws_container_avg_sum", "cws_fargate_task_avg_sum", "cws_host_top99p_sum", "data_jobs_monitoring_host_hr_agg_sum", "dbm_host_top99p_sum", "dbm_queries_avg_sum", "end_date", "eph_infra_host_agent_agg_sum", "eph_infra_host_alibaba_agg_sum", "eph_infra_host_aws_agg_sum", "eph_infra_host_azure_agg_sum", "eph_infra_host_ent_agg_sum", "eph_infra_host_gcp_agg_sum", "eph_infra_host_heroku_agg_sum", "eph_infra_host_only_aas_agg_sum", "eph_infra_host_only_vsphere_agg_sum", "eph_infra_host_opentelemetry_agg_sum", "eph_infra_host_opentelemetry_apm_agg_sum", "eph_infra_host_pro_agg_sum", "eph_infra_host_proplus_agg_sum", "eph_infra_host_proxmox_agg_sum", "error_tracking_apm_error_events_agg_sum", "error_tracking_error_events_agg_sum", "error_tracking_events_agg_sum", "error_tracking_rum_error_events_agg_sum", "event_management_correlation_agg_sum", "event_management_correlation_correlated_events_agg_sum", "event_management_correlation_correlated_related_events_agg_sum", "fargate_container_profiler_profiling_fargate_avg_sum", "fargate_container_profiler_profiling_fargate_eks_avg_sum", "fargate_tasks_count_avg_sum", "fargate_tasks_count_hwm_sum", "flex_logs_compute_large_avg_sum", "flex_logs_compute_medium_avg_sum", "flex_logs_compute_small_avg_sum", "flex_logs_compute_xlarge_avg_sum", "flex_logs_compute_xsmall_avg_sum", "flex_logs_starter_avg_sum", "flex_logs_starter_storage_index_avg_sum", "flex_logs_starter_storage_retention_adjustment_avg_sum", "flex_stored_logs_avg_sum", "forwarding_events_bytes_agg_sum", "gcp_host_top99p_sum", "heroku_host_top99p_sum", "incident_management_monthly_active_users_hwm_sum", "incident_management_seats_hwm_sum", "indexed_events_count_agg_sum", "infra_host_top99p_sum", "ingested_events_bytes_agg_sum", "iot_device_agg_sum", "iot_device_top99p_sum", "last_updated", "live_indexed_events_agg_sum", "live_ingested_bytes_agg_sum", "llm_observability_agg_sum", "llm_observability_min_spend_agg_sum", "logs_by_retention", "mobile_rum_lite_session_count_agg_sum", "mobile_rum_session_count_agg_sum", "mobile_rum_session_count_android_agg_sum", "mobile_rum_session_count_flutter_agg_sum", "mobile_rum_session_count_ios_agg_sum", "mobile_rum_session_count_reactnative_agg_sum", "mobile_rum_session_count_roku_agg_sum", "mobile_rum_units_agg_sum", "ndm_netflow_events_agg_sum", "netflow_indexed_events_count_agg_sum", "network_device_wireless_top99p_sum", "npm_host_top99p_sum", "observability_pipelines_bytes_processed_agg_sum", "oci_host_agg_sum", "oci_host_top99p_sum", "on_call_seat_hwm_sum", "online_archive_events_count_agg_sum", "opentelemetry_apm_host_top99p_sum", "opentelemetry_host_top99p_sum", "product_analytics_agg_sum", "profiling_aas_count_top99p_sum", "profiling_container_agent_count_avg", "profiling_host_count_top99p_sum", "proxmox_host_agg_sum", "proxmox_host_top99p_sum", "published_app_hwm_sum", "rehydrated_indexed_events_agg_sum", "rehydrated_ingested_bytes_agg_sum", "rum_browser_and_mobile_session_count", "rum_browser_legacy_session_count_agg_sum", "rum_browser_lite_session_count_agg_sum", "rum_browser_replay_session_count_agg_sum", "rum_indexed_sessions_agg_sum", "rum_ingested_sessions_agg_sum", "rum_lite_session_count_agg_sum", "rum_mobile_legacy_session_count_android_agg_sum", "rum_mobile_legacy_session_count_flutter_agg_sum", "rum_mobile_legacy_session_count_ios_agg_sum", "rum_mobile_legacy_session_count_reactnative_agg_sum", "rum_mobile_legacy_session_count_roku_agg_sum", "rum_mobile_lite_session_count_android_agg_sum", "rum_mobile_lite_session_count_flutter_agg_sum", "rum_mobile_lite_session_count_ios_agg_sum", "rum_mobile_lite_session_count_kotlinmultiplatform_agg_sum", "rum_mobile_lite_session_count_reactnative_agg_sum", "rum_mobile_lite_session_count_roku_agg_sum", "rum_mobile_lite_session_count_unity_agg_sum", "rum_mobile_replay_session_count_android_agg_sum", "rum_mobile_replay_session_count_ios_agg_sum", "rum_mobile_replay_session_count_kotlinmultiplatform_agg_sum", "rum_mobile_replay_session_count_reactnative_agg_sum", "rum_replay_session_count_agg_sum", "rum_session_count_agg_sum", "rum_session_replay_add_on_agg_sum", "rum_total_session_count_agg_sum", "rum_units_agg_sum", "sca_fargate_count_avg_sum", "sca_fargate_count_hwm_sum", "sds_apm_scanned_bytes_sum", "sds_events_scanned_bytes_sum", "sds_logs_scanned_bytes_sum", "sds_rum_scanned_bytes_sum", "sds_total_scanned_bytes_sum", "serverless_apps_apm_apm_azure_appservice_instances_avg_sum", "serverless_apps_apm_apm_azure_azurefunction_instances_avg_sum", "serverless_apps_apm_apm_azure_containerapp_instances_avg_sum", "serverless_apps_apm_apm_fargate_ecs_tasks_avg_sum", "serverless_apps_apm_apm_gcp_cloudfunction_instances_avg_sum", "serverless_apps_apm_apm_gcp_cloudrun_instances_avg_sum", "serverless_apps_apm_avg_sum", "serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg_sum", "serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg_sum", "serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg_sum", "serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg_sum", "serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg_sum", "serverless_apps_apm_excl_fargate_avg_sum", "serverless_apps_azure_container_app_instances_avg_sum", "serverless_apps_azure_count_avg_sum", "serverless_apps_azure_function_app_instances_avg_sum", "serverless_apps_azure_web_app_instances_avg_sum", "serverless_apps_ecs_avg_sum", "serverless_apps_eks_avg_sum", "serverless_apps_excl_fargate_avg_sum", "serverless_apps_excl_fargate_azure_container_app_instances_avg_sum", "serverless_apps_excl_fargate_azure_function_app_instances_avg_sum", "serverless_apps_excl_fargate_azure_web_app_instances_avg_sum", "serverless_apps_excl_fargate_google_cloud_functions_instances_avg_sum", "serverless_apps_excl_fargate_google_cloud_run_instances_avg_sum", "serverless_apps_google_cloud_functions_instances_avg_sum", "serverless_apps_google_cloud_run_instances_avg_sum", "serverless_apps_google_count_avg_sum", "serverless_apps_total_count_avg_sum", "siem_analyzed_logs_add_on_count_agg_sum", "start_date", "synthetics_browser_check_calls_count_agg_sum", "synthetics_check_calls_count_agg_sum", "synthetics_mobile_test_runs_agg_sum", "synthetics_parallel_testing_max_slots_hwm_sum", "trace_search_indexed_events_count_agg_sum", "twol_ingested_events_bytes_agg_sum", "universal_service_monitoring_host_top99p_sum", "usage", "vsphere_host_top99p_sum", "vuln_management_host_count_top99p_sum", "workflow_executions_usage_agg_sum"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AgentHostTop99pSum = all.AgentHostTop99pSum
	o.ApmAzureAppServiceHostTop99pSum = all.ApmAzureAppServiceHostTop99pSum
	o.ApmDevsecopsHostTop99pSum = all.ApmDevsecopsHostTop99pSum
	o.ApmEnterpriseStandaloneHostsTop99pSum = all.ApmEnterpriseStandaloneHostsTop99pSum
	o.ApmFargateCountAvgSum = all.ApmFargateCountAvgSum
	o.ApmHostTop99pSum = all.ApmHostTop99pSum
	o.ApmProStandaloneHostsTop99pSum = all.ApmProStandaloneHostsTop99pSum
	o.AppsecFargateCountAvgSum = all.AppsecFargateCountAvgSum
	o.AsmServerlessAggSum = all.AsmServerlessAggSum
	o.AuditLogsLinesIndexedAggSum = all.AuditLogsLinesIndexedAggSum
	o.AuditTrailEnabledHwmSum = all.AuditTrailEnabledHwmSum
	o.AvgProfiledFargateTasksSum = all.AvgProfiledFargateTasksSum
	o.AwsHostTop99pSum = all.AwsHostTop99pSum
	o.AwsLambdaFuncCount = all.AwsLambdaFuncCount
	o.AwsLambdaInvocationsSum = all.AwsLambdaInvocationsSum
	o.AzureAppServiceTop99pSum = all.AzureAppServiceTop99pSum
	o.AzureHostTop99pSum = all.AzureHostTop99pSum
	o.BillableIngestedBytesAggSum = all.BillableIngestedBytesAggSum
	o.BitsAiInvestigationsAggSum = all.BitsAiInvestigationsAggSum
	o.BrowserRumLiteSessionCountAggSum = all.BrowserRumLiteSessionCountAggSum
	o.BrowserRumReplaySessionCountAggSum = all.BrowserRumReplaySessionCountAggSum
	o.BrowserRumUnitsAggSum = all.BrowserRumUnitsAggSum
	o.CiPipelineIndexedSpansAggSum = all.CiPipelineIndexedSpansAggSum
	o.CiTestIndexedSpansAggSum = all.CiTestIndexedSpansAggSum
	o.CiVisibilityItrCommittersHwmSum = all.CiVisibilityItrCommittersHwmSum
	o.CiVisibilityPipelineCommittersHwmSum = all.CiVisibilityPipelineCommittersHwmSum
	o.CiVisibilityTestCommittersHwmSum = all.CiVisibilityTestCommittersHwmSum
	o.CloudCostManagementAwsHostCountAvgSum = all.CloudCostManagementAwsHostCountAvgSum
	o.CloudCostManagementAzureHostCountAvgSum = all.CloudCostManagementAzureHostCountAvgSum
	o.CloudCostManagementGcpHostCountAvgSum = all.CloudCostManagementGcpHostCountAvgSum
	o.CloudCostManagementHostCountAvgSum = all.CloudCostManagementHostCountAvgSum
	o.CloudCostManagementOciHostCountAvgSum = all.CloudCostManagementOciHostCountAvgSum
	o.CloudSiemEventsAggSum = all.CloudSiemEventsAggSum
	o.CodeAnalysisSaCommittersHwmSum = all.CodeAnalysisSaCommittersHwmSum
	o.CodeAnalysisScaCommittersHwmSum = all.CodeAnalysisScaCommittersHwmSum
	o.CodeSecurityHostTop99pSum = all.CodeSecurityHostTop99pSum
	o.ContainerAvgSum = all.ContainerAvgSum
	o.ContainerExclAgentAvgSum = all.ContainerExclAgentAvgSum
	o.ContainerHwmSum = all.ContainerHwmSum
	o.CsmContainerEnterpriseComplianceCountAggSum = all.CsmContainerEnterpriseComplianceCountAggSum
	o.CsmContainerEnterpriseCwsCountAggSum = all.CsmContainerEnterpriseCwsCountAggSum
	o.CsmContainerEnterpriseTotalCountAggSum = all.CsmContainerEnterpriseTotalCountAggSum
	o.CsmHostEnterpriseAasHostCountTop99pSum = all.CsmHostEnterpriseAasHostCountTop99pSum
	o.CsmHostEnterpriseAwsHostCountTop99pSum = all.CsmHostEnterpriseAwsHostCountTop99pSum
	o.CsmHostEnterpriseAzureHostCountTop99pSum = all.CsmHostEnterpriseAzureHostCountTop99pSum
	o.CsmHostEnterpriseComplianceHostCountTop99pSum = all.CsmHostEnterpriseComplianceHostCountTop99pSum
	o.CsmHostEnterpriseCwsHostCountTop99pSum = all.CsmHostEnterpriseCwsHostCountTop99pSum
	o.CsmHostEnterpriseGcpHostCountTop99pSum = all.CsmHostEnterpriseGcpHostCountTop99pSum
	o.CsmHostEnterpriseTotalHostCountTop99pSum = all.CsmHostEnterpriseTotalHostCountTop99pSum
	o.CspmAasHostTop99pSum = all.CspmAasHostTop99pSum
	o.CspmAwsHostTop99pSum = all.CspmAwsHostTop99pSum
	o.CspmAzureHostTop99pSum = all.CspmAzureHostTop99pSum
	o.CspmContainerAvgSum = all.CspmContainerAvgSum
	o.CspmContainerHwmSum = all.CspmContainerHwmSum
	o.CspmGcpHostTop99pSum = all.CspmGcpHostTop99pSum
	o.CspmHostTop99pSum = all.CspmHostTop99pSum
	o.CustomHistoricalTsSum = all.CustomHistoricalTsSum
	o.CustomLiveTsSum = all.CustomLiveTsSum
	o.CustomTsSum = all.CustomTsSum
	o.CwsContainerAvgSum = all.CwsContainerAvgSum
	o.CwsFargateTaskAvgSum = all.CwsFargateTaskAvgSum
	o.CwsHostTop99pSum = all.CwsHostTop99pSum
	o.DataJobsMonitoringHostHrAggSum = all.DataJobsMonitoringHostHrAggSum
	o.DbmHostTop99pSum = all.DbmHostTop99pSum
	o.DbmQueriesAvgSum = all.DbmQueriesAvgSum
	o.EndDate = all.EndDate
	o.EphInfraHostAgentAggSum = all.EphInfraHostAgentAggSum
	o.EphInfraHostAlibabaAggSum = all.EphInfraHostAlibabaAggSum
	o.EphInfraHostAwsAggSum = all.EphInfraHostAwsAggSum
	o.EphInfraHostAzureAggSum = all.EphInfraHostAzureAggSum
	o.EphInfraHostEntAggSum = all.EphInfraHostEntAggSum
	o.EphInfraHostGcpAggSum = all.EphInfraHostGcpAggSum
	o.EphInfraHostHerokuAggSum = all.EphInfraHostHerokuAggSum
	o.EphInfraHostOnlyAasAggSum = all.EphInfraHostOnlyAasAggSum
	o.EphInfraHostOnlyVsphereAggSum = all.EphInfraHostOnlyVsphereAggSum
	o.EphInfraHostOpentelemetryAggSum = all.EphInfraHostOpentelemetryAggSum
	o.EphInfraHostOpentelemetryApmAggSum = all.EphInfraHostOpentelemetryApmAggSum
	o.EphInfraHostProAggSum = all.EphInfraHostProAggSum
	o.EphInfraHostProplusAggSum = all.EphInfraHostProplusAggSum
	o.EphInfraHostProxmoxAggSum = all.EphInfraHostProxmoxAggSum
	o.ErrorTrackingApmErrorEventsAggSum = all.ErrorTrackingApmErrorEventsAggSum
	o.ErrorTrackingErrorEventsAggSum = all.ErrorTrackingErrorEventsAggSum
	o.ErrorTrackingEventsAggSum = all.ErrorTrackingEventsAggSum
	o.ErrorTrackingRumErrorEventsAggSum = all.ErrorTrackingRumErrorEventsAggSum
	o.EventManagementCorrelationAggSum = all.EventManagementCorrelationAggSum
	o.EventManagementCorrelationCorrelatedEventsAggSum = all.EventManagementCorrelationCorrelatedEventsAggSum
	o.EventManagementCorrelationCorrelatedRelatedEventsAggSum = all.EventManagementCorrelationCorrelatedRelatedEventsAggSum
	o.FargateContainerProfilerProfilingFargateAvgSum = all.FargateContainerProfilerProfilingFargateAvgSum
	o.FargateContainerProfilerProfilingFargateEksAvgSum = all.FargateContainerProfilerProfilingFargateEksAvgSum
	o.FargateTasksCountAvgSum = all.FargateTasksCountAvgSum
	o.FargateTasksCountHwmSum = all.FargateTasksCountHwmSum
	o.FlexLogsComputeLargeAvgSum = all.FlexLogsComputeLargeAvgSum
	o.FlexLogsComputeMediumAvgSum = all.FlexLogsComputeMediumAvgSum
	o.FlexLogsComputeSmallAvgSum = all.FlexLogsComputeSmallAvgSum
	o.FlexLogsComputeXlargeAvgSum = all.FlexLogsComputeXlargeAvgSum
	o.FlexLogsComputeXsmallAvgSum = all.FlexLogsComputeXsmallAvgSum
	o.FlexLogsStarterAvgSum = all.FlexLogsStarterAvgSum
	o.FlexLogsStarterStorageIndexAvgSum = all.FlexLogsStarterStorageIndexAvgSum
	o.FlexLogsStarterStorageRetentionAdjustmentAvgSum = all.FlexLogsStarterStorageRetentionAdjustmentAvgSum
	o.FlexStoredLogsAvgSum = all.FlexStoredLogsAvgSum
	o.ForwardingEventsBytesAggSum = all.ForwardingEventsBytesAggSum
	o.GcpHostTop99pSum = all.GcpHostTop99pSum
	o.HerokuHostTop99pSum = all.HerokuHostTop99pSum
	o.IncidentManagementMonthlyActiveUsersHwmSum = all.IncidentManagementMonthlyActiveUsersHwmSum
	o.IncidentManagementSeatsHwmSum = all.IncidentManagementSeatsHwmSum
	o.IndexedEventsCountAggSum = all.IndexedEventsCountAggSum
	o.InfraHostTop99pSum = all.InfraHostTop99pSum
	o.IngestedEventsBytesAggSum = all.IngestedEventsBytesAggSum
	o.IotDeviceAggSum = all.IotDeviceAggSum
	o.IotDeviceTop99pSum = all.IotDeviceTop99pSum
	o.LastUpdated = all.LastUpdated
	o.LiveIndexedEventsAggSum = all.LiveIndexedEventsAggSum
	o.LiveIngestedBytesAggSum = all.LiveIngestedBytesAggSum
	o.LlmObservabilityAggSum = all.LlmObservabilityAggSum
	o.LlmObservabilityMinSpendAggSum = all.LlmObservabilityMinSpendAggSum
	if all.LogsByRetention != nil && all.LogsByRetention.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsByRetention = all.LogsByRetention
	o.MobileRumLiteSessionCountAggSum = all.MobileRumLiteSessionCountAggSum
	o.MobileRumSessionCountAggSum = all.MobileRumSessionCountAggSum
	o.MobileRumSessionCountAndroidAggSum = all.MobileRumSessionCountAndroidAggSum
	o.MobileRumSessionCountFlutterAggSum = all.MobileRumSessionCountFlutterAggSum
	o.MobileRumSessionCountIosAggSum = all.MobileRumSessionCountIosAggSum
	o.MobileRumSessionCountReactnativeAggSum = all.MobileRumSessionCountReactnativeAggSum
	o.MobileRumSessionCountRokuAggSum = all.MobileRumSessionCountRokuAggSum
	o.MobileRumUnitsAggSum = all.MobileRumUnitsAggSum
	o.NdmNetflowEventsAggSum = all.NdmNetflowEventsAggSum
	o.NetflowIndexedEventsCountAggSum = all.NetflowIndexedEventsCountAggSum
	o.NetworkDeviceWirelessTop99pSum = all.NetworkDeviceWirelessTop99pSum
	o.NpmHostTop99pSum = all.NpmHostTop99pSum
	o.ObservabilityPipelinesBytesProcessedAggSum = all.ObservabilityPipelinesBytesProcessedAggSum
	o.OciHostAggSum = all.OciHostAggSum
	o.OciHostTop99pSum = all.OciHostTop99pSum
	o.OnCallSeatHwmSum = all.OnCallSeatHwmSum
	o.OnlineArchiveEventsCountAggSum = all.OnlineArchiveEventsCountAggSum
	o.OpentelemetryApmHostTop99pSum = all.OpentelemetryApmHostTop99pSum
	o.OpentelemetryHostTop99pSum = all.OpentelemetryHostTop99pSum
	o.ProductAnalyticsAggSum = all.ProductAnalyticsAggSum
	o.ProfilingAasCountTop99pSum = all.ProfilingAasCountTop99pSum
	o.ProfilingContainerAgentCountAvg = all.ProfilingContainerAgentCountAvg
	o.ProfilingHostCountTop99pSum = all.ProfilingHostCountTop99pSum
	o.ProxmoxHostAggSum = all.ProxmoxHostAggSum
	o.ProxmoxHostTop99pSum = all.ProxmoxHostTop99pSum
	o.PublishedAppHwmSum = all.PublishedAppHwmSum
	o.RehydratedIndexedEventsAggSum = all.RehydratedIndexedEventsAggSum
	o.RehydratedIngestedBytesAggSum = all.RehydratedIngestedBytesAggSum
	o.RumBrowserAndMobileSessionCount = all.RumBrowserAndMobileSessionCount
	o.RumBrowserLegacySessionCountAggSum = all.RumBrowserLegacySessionCountAggSum
	o.RumBrowserLiteSessionCountAggSum = all.RumBrowserLiteSessionCountAggSum
	o.RumBrowserReplaySessionCountAggSum = all.RumBrowserReplaySessionCountAggSum
	o.RumIndexedSessionsAggSum = all.RumIndexedSessionsAggSum
	o.RumIngestedSessionsAggSum = all.RumIngestedSessionsAggSum
	o.RumLiteSessionCountAggSum = all.RumLiteSessionCountAggSum
	o.RumMobileLegacySessionCountAndroidAggSum = all.RumMobileLegacySessionCountAndroidAggSum
	o.RumMobileLegacySessionCountFlutterAggSum = all.RumMobileLegacySessionCountFlutterAggSum
	o.RumMobileLegacySessionCountIosAggSum = all.RumMobileLegacySessionCountIosAggSum
	o.RumMobileLegacySessionCountReactnativeAggSum = all.RumMobileLegacySessionCountReactnativeAggSum
	o.RumMobileLegacySessionCountRokuAggSum = all.RumMobileLegacySessionCountRokuAggSum
	o.RumMobileLiteSessionCountAndroidAggSum = all.RumMobileLiteSessionCountAndroidAggSum
	o.RumMobileLiteSessionCountFlutterAggSum = all.RumMobileLiteSessionCountFlutterAggSum
	o.RumMobileLiteSessionCountIosAggSum = all.RumMobileLiteSessionCountIosAggSum
	o.RumMobileLiteSessionCountKotlinmultiplatformAggSum = all.RumMobileLiteSessionCountKotlinmultiplatformAggSum
	o.RumMobileLiteSessionCountReactnativeAggSum = all.RumMobileLiteSessionCountReactnativeAggSum
	o.RumMobileLiteSessionCountRokuAggSum = all.RumMobileLiteSessionCountRokuAggSum
	o.RumMobileLiteSessionCountUnityAggSum = all.RumMobileLiteSessionCountUnityAggSum
	o.RumMobileReplaySessionCountAndroidAggSum = all.RumMobileReplaySessionCountAndroidAggSum
	o.RumMobileReplaySessionCountIosAggSum = all.RumMobileReplaySessionCountIosAggSum
	o.RumMobileReplaySessionCountKotlinmultiplatformAggSum = all.RumMobileReplaySessionCountKotlinmultiplatformAggSum
	o.RumMobileReplaySessionCountReactnativeAggSum = all.RumMobileReplaySessionCountReactnativeAggSum
	o.RumReplaySessionCountAggSum = all.RumReplaySessionCountAggSum
	o.RumSessionCountAggSum = all.RumSessionCountAggSum
	o.RumSessionReplayAddOnAggSum = all.RumSessionReplayAddOnAggSum
	o.RumTotalSessionCountAggSum = all.RumTotalSessionCountAggSum
	o.RumUnitsAggSum = all.RumUnitsAggSum
	o.ScaFargateCountAvgSum = all.ScaFargateCountAvgSum
	o.ScaFargateCountHwmSum = all.ScaFargateCountHwmSum
	o.SdsApmScannedBytesSum = all.SdsApmScannedBytesSum
	o.SdsEventsScannedBytesSum = all.SdsEventsScannedBytesSum
	o.SdsLogsScannedBytesSum = all.SdsLogsScannedBytesSum
	o.SdsRumScannedBytesSum = all.SdsRumScannedBytesSum
	o.SdsTotalScannedBytesSum = all.SdsTotalScannedBytesSum
	o.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum = all.ServerlessAppsApmApmAzureAppserviceInstancesAvgSum
	o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum = all.ServerlessAppsApmApmAzureAzurefunctionInstancesAvgSum
	o.ServerlessAppsApmApmAzureContainerappInstancesAvgSum = all.ServerlessAppsApmApmAzureContainerappInstancesAvgSum
	o.ServerlessAppsApmApmFargateEcsTasksAvgSum = all.ServerlessAppsApmApmFargateEcsTasksAvgSum
	o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum = all.ServerlessAppsApmApmGcpCloudfunctionInstancesAvgSum
	o.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum = all.ServerlessAppsApmApmGcpCloudrunInstancesAvgSum
	o.ServerlessAppsApmAvgSum = all.ServerlessAppsApmAvgSum
	o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum = all.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgSum
	o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum = all.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgSum
	o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum = all.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgSum
	o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum = all.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgSum
	o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum = all.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgSum
	o.ServerlessAppsApmExclFargateAvgSum = all.ServerlessAppsApmExclFargateAvgSum
	o.ServerlessAppsAzureContainerAppInstancesAvgSum = all.ServerlessAppsAzureContainerAppInstancesAvgSum
	o.ServerlessAppsAzureCountAvgSum = all.ServerlessAppsAzureCountAvgSum
	o.ServerlessAppsAzureFunctionAppInstancesAvgSum = all.ServerlessAppsAzureFunctionAppInstancesAvgSum
	o.ServerlessAppsAzureWebAppInstancesAvgSum = all.ServerlessAppsAzureWebAppInstancesAvgSum
	o.ServerlessAppsEcsAvgSum = all.ServerlessAppsEcsAvgSum
	o.ServerlessAppsEksAvgSum = all.ServerlessAppsEksAvgSum
	o.ServerlessAppsExclFargateAvgSum = all.ServerlessAppsExclFargateAvgSum
	o.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum = all.ServerlessAppsExclFargateAzureContainerAppInstancesAvgSum
	o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum = all.ServerlessAppsExclFargateAzureFunctionAppInstancesAvgSum
	o.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum = all.ServerlessAppsExclFargateAzureWebAppInstancesAvgSum
	o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum = all.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgSum
	o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum = all.ServerlessAppsExclFargateGoogleCloudRunInstancesAvgSum
	o.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum = all.ServerlessAppsGoogleCloudFunctionsInstancesAvgSum
	o.ServerlessAppsGoogleCloudRunInstancesAvgSum = all.ServerlessAppsGoogleCloudRunInstancesAvgSum
	o.ServerlessAppsGoogleCountAvgSum = all.ServerlessAppsGoogleCountAvgSum
	o.ServerlessAppsTotalCountAvgSum = all.ServerlessAppsTotalCountAvgSum
	o.SiemAnalyzedLogsAddOnCountAggSum = all.SiemAnalyzedLogsAddOnCountAggSum
	o.StartDate = all.StartDate
	o.SyntheticsBrowserCheckCallsCountAggSum = all.SyntheticsBrowserCheckCallsCountAggSum
	o.SyntheticsCheckCallsCountAggSum = all.SyntheticsCheckCallsCountAggSum
	o.SyntheticsMobileTestRunsAggSum = all.SyntheticsMobileTestRunsAggSum
	o.SyntheticsParallelTestingMaxSlotsHwmSum = all.SyntheticsParallelTestingMaxSlotsHwmSum
	o.TraceSearchIndexedEventsCountAggSum = all.TraceSearchIndexedEventsCountAggSum
	o.TwolIngestedEventsBytesAggSum = all.TwolIngestedEventsBytesAggSum
	o.UniversalServiceMonitoringHostTop99pSum = all.UniversalServiceMonitoringHostTop99pSum
	o.Usage = all.Usage
	o.VsphereHostTop99pSum = all.VsphereHostTop99pSum
	o.VulnManagementHostCountTop99pSum = all.VulnManagementHostCountTop99pSum
	o.WorkflowExecutionsUsageAggSum = all.WorkflowExecutionsUsageAggSum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
