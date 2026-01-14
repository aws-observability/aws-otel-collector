// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSummaryDate Response with hourly report of all data billed by Datadog all organizations.
type UsageSummaryDate struct {
	// Shows the 99th percentile of all agent hosts over all hours in the current date for all organizations.
	AgentHostTop99p *int64 `json:"agent_host_top99p,omitempty"`
	// Shows the 99th percentile of all Azure app services using APM over all hours in the current date all organizations.
	ApmAzureAppServiceHostTop99p *int64 `json:"apm_azure_app_service_host_top99p,omitempty"`
	// Shows the 99th percentile of all APM DevSecOps hosts over all hours in the current date for the given org.
	ApmDevsecopsHostTop99p *int64 `json:"apm_devsecops_host_top99p,omitempty"`
	// Shows the 99th percentile of all distinct standalone Enterprise hosts over all hours in the current date for all organizations.
	ApmEnterpriseStandaloneHostsTop99p *int64 `json:"apm_enterprise_standalone_hosts_top99p,omitempty"`
	// Shows the average of all APM ECS Fargate tasks over all hours in the current date for all organizations.
	ApmFargateCountAvg *int64 `json:"apm_fargate_count_avg,omitempty"`
	// Shows the 99th percentile of all distinct APM hosts over all hours in the current date for all organizations.
	ApmHostTop99p *int64 `json:"apm_host_top99p,omitempty"`
	// Shows the 99th percentile of all distinct standalone Pro hosts over all hours in the current date for all organizations.
	ApmProStandaloneHostsTop99p *int64 `json:"apm_pro_standalone_hosts_top99p,omitempty"`
	// Shows the average of all Application Security Monitoring ECS Fargate tasks over all hours in the current date for all organizations.
	AppsecFargateCountAvg *int64 `json:"appsec_fargate_count_avg,omitempty"`
	// Shows the sum of all Application Security Monitoring Serverless invocations over all hours in the current date for all organizations.
	AsmServerlessSum *int64 `json:"asm_serverless_sum,omitempty"`
	// Shows the sum of audit logs lines indexed over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	AuditLogsLinesIndexedSum *int64 `json:"audit_logs_lines_indexed_sum,omitempty"`
	// Shows the number of organizations that had Audit Trail enabled in the current date.
	AuditTrailEnabledHwm *int64 `json:"audit_trail_enabled_hwm,omitempty"`
	// The average total count for Fargate Container Profiler over all hours in the current date for all organizations.
	AvgProfiledFargateTasks *int64 `json:"avg_profiled_fargate_tasks,omitempty"`
	// Shows the 99th percentile of all AWS hosts over all hours in the current date for all organizations.
	AwsHostTop99p *int64 `json:"aws_host_top99p,omitempty"`
	// Shows the average of the number of functions that executed 1 or more times each hour in the current date for all organizations.
	AwsLambdaFuncCount *int64 `json:"aws_lambda_func_count,omitempty"`
	// Shows the sum of all AWS Lambda invocations over all hours in the current date for all organizations.
	AwsLambdaInvocationsSum *int64 `json:"aws_lambda_invocations_sum,omitempty"`
	// Shows the 99th percentile of all Azure app services over all hours in the current date for all organizations.
	AzureAppServiceTop99p *int64 `json:"azure_app_service_top99p,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current date for all organizations.
	BillableIngestedBytesSum *int64 `json:"billable_ingested_bytes_sum,omitempty"`
	// Shows the sum of all Bits AI Investigations over all hours in the current date for all organizations.
	BitsAiInvestigationsSum *int64 `json:"bits_ai_investigations_sum,omitempty"`
	// Shows the sum of all browser lite sessions over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	BrowserRumLiteSessionCountSum *int64 `json:"browser_rum_lite_session_count_sum,omitempty"`
	// Shows the sum of all browser replay sessions over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	BrowserRumReplaySessionCountSum *int64 `json:"browser_rum_replay_session_count_sum,omitempty"`
	// Shows the sum of all browser RUM units over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	BrowserRumUnitsSum *int64 `json:"browser_rum_units_sum,omitempty"`
	// Shows the sum of all CI pipeline indexed spans over all hours in the current month for all organizations.
	CiPipelineIndexedSpansSum *int64 `json:"ci_pipeline_indexed_spans_sum,omitempty"`
	// Shows the sum of all CI test indexed spans over all hours in the current month for all organizations.
	CiTestIndexedSpansSum *int64 `json:"ci_test_indexed_spans_sum,omitempty"`
	// Shows the high-water mark of all CI visibility intelligent test runner committers over all hours in the current month for all organizations.
	CiVisibilityItrCommittersHwm *int64 `json:"ci_visibility_itr_committers_hwm,omitempty"`
	// Shows the high-water mark of all CI visibility pipeline committers over all hours in the current month for all organizations.
	CiVisibilityPipelineCommittersHwm *int64 `json:"ci_visibility_pipeline_committers_hwm,omitempty"`
	// Shows the high-water mark of all CI visibility test committers over all hours in the current month for all organizations.
	CiVisibilityTestCommittersHwm *int64 `json:"ci_visibility_test_committers_hwm,omitempty"`
	// Host count average of Cloud Cost Management for AWS for the given date and given organization.
	CloudCostManagementAwsHostCountAvg *int64 `json:"cloud_cost_management_aws_host_count_avg,omitempty"`
	// Host count average of Cloud Cost Management for Azure for the given date and given organization.
	CloudCostManagementAzureHostCountAvg *int64 `json:"cloud_cost_management_azure_host_count_avg,omitempty"`
	// Host count average of Cloud Cost Management for GCP for the given date and given organization.
	CloudCostManagementGcpHostCountAvg *int64 `json:"cloud_cost_management_gcp_host_count_avg,omitempty"`
	// Host count average of Cloud Cost Management for all cloud providers for the given date and given organization.
	CloudCostManagementHostCountAvg *int64 `json:"cloud_cost_management_host_count_avg,omitempty"`
	// Average host count for Cloud Cost Management on OCI for the given date and organization.
	CloudCostManagementOciHostCountAvg *int64 `json:"cloud_cost_management_oci_host_count_avg,omitempty"`
	// Shows the sum of all Cloud Security Information and Event Management events over all hours in the current date for the given org.
	CloudSiemEventsSum *int64 `json:"cloud_siem_events_sum,omitempty"`
	// Shows the high-water mark of all Static Analysis committers over all hours in the current date for the given org.
	CodeAnalysisSaCommittersHwm *int64 `json:"code_analysis_sa_committers_hwm,omitempty"`
	// Shows the high-water mark of all static Software Composition Analysis committers over all hours in the current date for the given org.
	CodeAnalysisScaCommittersHwm *int64 `json:"code_analysis_sca_committers_hwm,omitempty"`
	// Shows the 99th percentile of all Code Security hosts over all hours in the current date for the given org.
	CodeSecurityHostTop99p *int64 `json:"code_security_host_top99p,omitempty"`
	// Shows the average of all distinct containers over all hours in the current date for all organizations.
	ContainerAvg *int64 `json:"container_avg,omitempty"`
	// Shows the average of containers without the Datadog Agent over all hours in the current date for all organizations.
	ContainerExclAgentAvg *int64 `json:"container_excl_agent_avg,omitempty"`
	// Shows the high-water mark of all distinct containers over all hours in the current date for all organizations.
	ContainerHwm *int64 `json:"container_hwm,omitempty"`
	// Shows the sum of all Cloud Security Management Enterprise compliance containers over all hours in the current date for the given org.
	CsmContainerEnterpriseComplianceCountSum *int64 `json:"csm_container_enterprise_compliance_count_sum,omitempty"`
	// Shows the sum of all Cloud Security Management Enterprise Cloud Workload Security containers over all hours in the current date for the given org.
	CsmContainerEnterpriseCwsCountSum *int64 `json:"csm_container_enterprise_cws_count_sum,omitempty"`
	// Shows the sum of all Cloud Security Management Enterprise containers over all hours in the current date for the given org.
	CsmContainerEnterpriseTotalCountSum *int64 `json:"csm_container_enterprise_total_count_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise Azure app services hosts over all hours in the current date for the given org.
	CsmHostEnterpriseAasHostCountTop99p *int64 `json:"csm_host_enterprise_aas_host_count_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise AWS hosts over all hours in the current date for the given org.
	CsmHostEnterpriseAwsHostCountTop99p *int64 `json:"csm_host_enterprise_aws_host_count_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise Azure hosts over all hours in the current date for the given org.
	CsmHostEnterpriseAzureHostCountTop99p *int64 `json:"csm_host_enterprise_azure_host_count_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise compliance hosts over all hours in the current date for the given org.
	CsmHostEnterpriseComplianceHostCountTop99p *int64 `json:"csm_host_enterprise_compliance_host_count_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise Cloud Workload Security hosts over all hours in the current date for the given org.
	CsmHostEnterpriseCwsHostCountTop99p *int64 `json:"csm_host_enterprise_cws_host_count_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise GCP hosts over all hours in the current date for the given org.
	CsmHostEnterpriseGcpHostCountTop99p *int64 `json:"csm_host_enterprise_gcp_host_count_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise hosts over all hours in the current date for the given org.
	CsmHostEnterpriseTotalHostCountTop99p *int64 `json:"csm_host_enterprise_total_host_count_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro Azure app services hosts over all hours in the current date for all organizations.
	CspmAasHostTop99p *int64 `json:"cspm_aas_host_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro AWS hosts over all hours in the current date for all organizations.
	CspmAwsHostTop99p *int64 `json:"cspm_aws_host_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro Azure hosts over all hours in the current date for all organizations.
	CspmAzureHostTop99p *int64 `json:"cspm_azure_host_top99p,omitempty"`
	// Shows the average number of Cloud Security Management Pro containers over all hours in the current date for all organizations.
	CspmContainerAvg *int64 `json:"cspm_container_avg,omitempty"`
	// Shows the high-water mark of Cloud Security Management Pro containers over all hours in the current date for all organizations.
	CspmContainerHwm *int64 `json:"cspm_container_hwm,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro GCP hosts over all hours in the current date for all organizations.
	CspmGcpHostTop99p *int64 `json:"cspm_gcp_host_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro hosts over all hours in the current date for all organizations.
	CspmHostTop99p *int64 `json:"cspm_host_top99p,omitempty"`
	// Shows the average number of distinct custom metrics over all hours in the current date for all organizations.
	CustomTsAvg *int64 `json:"custom_ts_avg,omitempty"`
	// Shows the average of all distinct Cloud Workload Security containers over all hours in the current date for all organizations.
	CwsContainerCountAvg *int64 `json:"cws_container_count_avg,omitempty"`
	// Shows the average of all distinct Cloud Workload Security Fargate tasks over all hours in the current date for all organizations.
	CwsFargateTaskAvg *int64 `json:"cws_fargate_task_avg,omitempty"`
	// Shows the 99th percentile of all Cloud Workload Security hosts over all hours in the current date for all organizations.
	CwsHostTop99p *int64 `json:"cws_host_top99p,omitempty"`
	// Shows the sum of all Data Jobs Monitoring hosts over all hours in the current date for the given org.
	DataJobsMonitoringHostHrSum *int64 `json:"data_jobs_monitoring_host_hr_sum,omitempty"`
	// The date for the usage.
	Date *time.Time `json:"date,omitempty"`
	// Shows the 99th percentile of all Database Monitoring hosts over all hours in the current date for all organizations.
	DbmHostTop99p *int64 `json:"dbm_host_top99p,omitempty"`
	// Shows the average of all normalized Database Monitoring queries over all hours in the current date for all organizations.
	DbmQueriesCountAvg *int64 `json:"dbm_queries_count_avg,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts with the Datadog Agent over all hours in the current date for the given org.
	EphInfraHostAgentSum *int64 `json:"eph_infra_host_agent_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on Alibaba over all hours in the current date for the given org.
	EphInfraHostAlibabaSum *int64 `json:"eph_infra_host_alibaba_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on AWS over all hours in the current date for the given org.
	EphInfraHostAwsSum *int64 `json:"eph_infra_host_aws_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on Azure over all hours in the current date for the given org.
	EphInfraHostAzureSum *int64 `json:"eph_infra_host_azure_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts for Enterprise over all hours in the current date for the given org.
	EphInfraHostEntSum *int64 `json:"eph_infra_host_ent_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on GCP over all hours in the current date for the given org.
	EphInfraHostGcpSum *int64 `json:"eph_infra_host_gcp_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on Heroku over all hours in the current date for the given org.
	EphInfraHostHerokuSum *int64 `json:"eph_infra_host_heroku_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts with only Azure App Services over all hours in the current date for the given org.
	EphInfraHostOnlyAasSum *int64 `json:"eph_infra_host_only_aas_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts with only vSphere over all hours in the current date for the given org.
	EphInfraHostOnlyVsphereSum *int64 `json:"eph_infra_host_only_vsphere_sum,omitempty"`
	// Shows the sum of all ephemeral APM hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for the given org.
	EphInfraHostOpentelemetryApmSum *int64 `json:"eph_infra_host_opentelemetry_apm_sum,omitempty"`
	// Shows the sum of all ephemeral hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for the given org.
	EphInfraHostOpentelemetrySum *int64 `json:"eph_infra_host_opentelemetry_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts for Pro over all hours in the current date for the given org.
	EphInfraHostProSum *int64 `json:"eph_infra_host_pro_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts for Pro Plus over all hours in the current date for the given org.
	EphInfraHostProplusSum *int64 `json:"eph_infra_host_proplus_sum,omitempty"`
	// Sum of all ephemeral infrastructure hosts for Proxmox over all hours in the current date for all organizations.
	EphInfraHostProxmoxSum *int64 `json:"eph_infra_host_proxmox_sum,omitempty"`
	// Shows the sum of all Error Tracking APM error events over all hours in the current date for the given org.
	ErrorTrackingApmErrorEventsSum *int64 `json:"error_tracking_apm_error_events_sum,omitempty"`
	// Shows the sum of all Error Tracking error events over all hours in the current date for the given org.
	ErrorTrackingErrorEventsSum *int64 `json:"error_tracking_error_events_sum,omitempty"`
	// Shows the sum of all Error Tracking events over all hours in the current date for the given org.
	ErrorTrackingEventsSum *int64 `json:"error_tracking_events_sum,omitempty"`
	// Shows the sum of all Error Tracking RUM error events over all hours in the current date for the given org.
	ErrorTrackingRumErrorEventsSum *int64 `json:"error_tracking_rum_error_events_sum,omitempty"`
	// Shows the sum of all Event Management correlated events over all hours in the current date for all organizations.
	EventManagementCorrelationCorrelatedEventsSum *int64 `json:"event_management_correlation_correlated_events_sum,omitempty"`
	// Shows the sum of all Event Management correlated related events over all hours in the current date for all organizations.
	EventManagementCorrelationCorrelatedRelatedEventsSum *int64 `json:"event_management_correlation_correlated_related_events_sum,omitempty"`
	// Shows the sum of all Event Management correlations over all hours in the current date for all organizations.
	EventManagementCorrelationSum *int64 `json:"event_management_correlation_sum,omitempty"`
	// The average number of Profiling Fargate tasks over all hours in the current date for all organizations.
	FargateContainerProfilerProfilingFargateAvg *int64 `json:"fargate_container_profiler_profiling_fargate_avg,omitempty"`
	// The average number of Profiling Fargate Elastic Kubernetes Service tasks over all hours in the current date for all organizations.
	FargateContainerProfilerProfilingFargateEksAvg *int64 `json:"fargate_container_profiler_profiling_fargate_eks_avg,omitempty"`
	// Shows the high-watermark of all Fargate tasks over all hours in the current date for all organizations.
	FargateTasksCountAvg *int64 `json:"fargate_tasks_count_avg,omitempty"`
	// Shows the average of all Fargate tasks over all hours in the current date for all organizations.
	FargateTasksCountHwm *int64 `json:"fargate_tasks_count_hwm,omitempty"`
	// Shows the average number of Flex Logs Compute Large Instances over all hours in the current date for the given org.
	FlexLogsComputeLargeAvg *int64 `json:"flex_logs_compute_large_avg,omitempty"`
	// Shows the average number of Flex Logs Compute Medium Instances over all hours in the current date for the given org.
	FlexLogsComputeMediumAvg *int64 `json:"flex_logs_compute_medium_avg,omitempty"`
	// Shows the average number of Flex Logs Compute Small Instances over all hours in the current date for the given org.
	FlexLogsComputeSmallAvg *int64 `json:"flex_logs_compute_small_avg,omitempty"`
	// Shows the average number of Flex Logs Compute Extra Large Instances over all hours in the current date for the given org.
	FlexLogsComputeXlargeAvg *int64 `json:"flex_logs_compute_xlarge_avg,omitempty"`
	// Shows the average number of Flex Logs Compute Extra Small Instances over all hours in the current date for the given org.
	FlexLogsComputeXsmallAvg *int64 `json:"flex_logs_compute_xsmall_avg,omitempty"`
	// Shows the average number of Flex Logs Starter Instances over all hours in the current date for the given org.
	FlexLogsStarterAvg *int64 `json:"flex_logs_starter_avg,omitempty"`
	// Shows the average number of Flex Logs Starter Storage Index Instances over all hours in the current date for the given org.
	FlexLogsStarterStorageIndexAvg *int64 `json:"flex_logs_starter_storage_index_avg,omitempty"`
	// Shows the average number of Flex Logs Starter Storage Retention Adjustment Instances over all hours in the current date for the given org.
	FlexLogsStarterStorageRetentionAdjustmentAvg *int64 `json:"flex_logs_starter_storage_retention_adjustment_avg,omitempty"`
	// Shows the average of all Flex Stored Logs over all hours in the current date for the given org.
	FlexStoredLogsAvg *int64 `json:"flex_stored_logs_avg,omitempty"`
	// Shows the sum of all log bytes forwarded over all hours in the current date for all organizations.
	ForwardingEventsBytesSum *int64 `json:"forwarding_events_bytes_sum,omitempty"`
	// Shows the 99th percentile of all GCP hosts over all hours in the current date for all organizations.
	GcpHostTop99p *int64 `json:"gcp_host_top99p,omitempty"`
	// Shows the 99th percentile of all Heroku dynos over all hours in the current date for all organizations.
	HerokuHostTop99p *int64 `json:"heroku_host_top99p,omitempty"`
	// Shows the high-water mark of incident management monthly active users over all hours in the current date for all organizations.
	IncidentManagementMonthlyActiveUsersHwm *int64 `json:"incident_management_monthly_active_users_hwm,omitempty"`
	// Shows the high-water mark of Incident Management seats over all hours on the current date for all organizations.
	IncidentManagementSeatsHwm *int64 `json:"incident_management_seats_hwm,omitempty"`
	// Shows the sum of all log events indexed over all hours in the current date for all organizations.
	IndexedEventsCountSum *int64 `json:"indexed_events_count_sum,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for all organizations.
	InfraHostTop99p *int64 `json:"infra_host_top99p,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current date for all organizations.
	IngestedEventsBytesSum *int64 `json:"ingested_events_bytes_sum,omitempty"`
	// Shows the sum of all IoT devices over all hours in the current date for all organizations.
	IotDeviceSum *int64 `json:"iot_device_sum,omitempty"`
	// Shows the 99th percentile of all IoT devices over all hours in the current date all organizations.
	IotDeviceTop99p *int64 `json:"iot_device_top99p,omitempty"`
	// Sum of all LLM observability minimum spend over all hours in the current date for all organizations.
	LlmObservabilityMinSpendSum *int64 `json:"llm_observability_min_spend_sum,omitempty"`
	// Sum of all LLM observability sessions over all hours in the current date for all organizations.
	LlmObservabilitySum *int64 `json:"llm_observability_sum,omitempty"`
	// Shows the sum of all mobile lite sessions over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumLiteSessionCountSum *int64 `json:"mobile_rum_lite_session_count_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on Android over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountAndroidSum *int64 `json:"mobile_rum_session_count_android_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on Flutter over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountFlutterSum *int64 `json:"mobile_rum_session_count_flutter_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on iOS over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountIosSum *int64 `json:"mobile_rum_session_count_ios_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on React Native over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountReactnativeSum *int64 `json:"mobile_rum_session_count_reactnative_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions on Roku over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountRokuSum *int64 `json:"mobile_rum_session_count_roku_sum,omitempty"`
	// Shows the sum of all mobile RUM sessions over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumSessionCountSum *int64 `json:"mobile_rum_session_count_sum,omitempty"`
	// Shows the sum of all mobile RUM units over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	MobileRumUnitsSum *int64 `json:"mobile_rum_units_sum,omitempty"`
	// Shows the sum of all Network Device Monitoring NetFlow events over all hours in the current date for the given org.
	NdmNetflowEventsSum *int64 `json:"ndm_netflow_events_sum,omitempty"`
	// Shows the sum of all Network flows indexed over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	NetflowIndexedEventsCountSum *int64 `json:"netflow_indexed_events_count_sum,omitempty"`
	// Shows the 99th percentile of all Network Device Monitoring wireless devices over all hours in the current date for all organizations.
	NetworkDeviceWirelessTop99p *int64 `json:"network_device_wireless_top99p,omitempty"`
	// Shows the 99th percentile of all distinct Cloud Network Monitoring hosts (formerly known as Network hosts) over all hours in the current date for all organizations.
	NpmHostTop99p *int64 `json:"npm_host_top99p,omitempty"`
	// Sum of all observability pipelines bytes processed over all hours in the current date for the given org.
	ObservabilityPipelinesBytesProcessedSum *int64 `json:"observability_pipelines_bytes_processed_sum,omitempty"`
	// Shows the sum of all Oracle Cloud Infrastructure hosts over all hours in the current date for the given org.
	OciHostSum *int64 `json:"oci_host_sum,omitempty"`
	// Shows the 99th percentile of all Oracle Cloud Infrastructure hosts over all hours in the current date for the given org.
	OciHostTop99p *int64 `json:"oci_host_top99p,omitempty"`
	// Shows the high-water mark of On-Call seats over all hours in the current date for all organizations.
	OnCallSeatHwm *int64 `json:"on_call_seat_hwm,omitempty"`
	// Sum of all online archived events over all hours in the current date for all organizations.
	OnlineArchiveEventsCountSum *int64 `json:"online_archive_events_count_sum,omitempty"`
	// Shows the 99th percentile of APM hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for all organizations.
	OpentelemetryApmHostTop99p *int64 `json:"opentelemetry_apm_host_top99p,omitempty"`
	// Shows the 99th percentile of all hosts reported by the Datadog exporter for the OpenTelemetry Collector over all hours in the current date for all organizations.
	OpentelemetryHostTop99p *int64 `json:"opentelemetry_host_top99p,omitempty"`
	// Organizations associated with a user.
	Orgs []UsageSummaryDateOrg `json:"orgs,omitempty"`
	// Sum of all product analytics sessions over all hours in the current date for all organizations.
	ProductAnalyticsSum *int64 `json:"product_analytics_sum,omitempty"`
	// Shows the 99th percentile of all profiled Azure app services over all hours in the current date for all organizations.
	ProfilingAasCountTop99p *int64 `json:"profiling_aas_count_top99p,omitempty"`
	// Shows the 99th percentile of all profiled hosts over all hours within the current date for all organizations.
	ProfilingHostTop99p *int64 `json:"profiling_host_top99p,omitempty"`
	// Sum of all Proxmox hosts over all hours in the current date for all organizations.
	ProxmoxHostSum *int64 `json:"proxmox_host_sum,omitempty"`
	// 99th percentile of all Proxmox hosts over all hours in the current date for all organizations.
	ProxmoxHostTop99p *int64 `json:"proxmox_host_top99p,omitempty"`
	// Shows the high-water mark of all published applications over all hours in the current date for all organizations.
	PublishedAppHwm *int64 `json:"published_app_hwm,omitempty"`
	// Shows the sum of all mobile sessions and all browser lite and legacy sessions over all hours in the current month for all organizations (To be deprecated on October 1st, 2024).
	RumBrowserAndMobileSessionCount *int64 `json:"rum_browser_and_mobile_session_count,omitempty"`
	// Shows the sum of all browser RUM legacy sessions over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumBrowserLegacySessionCountSum *int64 `json:"rum_browser_legacy_session_count_sum,omitempty"`
	// Shows the sum of all browser RUM lite sessions over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumBrowserLiteSessionCountSum *int64 `json:"rum_browser_lite_session_count_sum,omitempty"`
	// Shows the sum of all browser RUM Session Replay counts over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumBrowserReplaySessionCountSum *int64 `json:"rum_browser_replay_session_count_sum,omitempty"`
	// Sum of all RUM indexed sessions over all hours in the current date for all organizations.
	RumIndexedSessionsSum *int64 `json:"rum_indexed_sessions_sum,omitempty"`
	// Sum of all RUM ingested sessions over all hours in the current date for all organizations.
	RumIngestedSessionsSum *int64 `json:"rum_ingested_sessions_sum,omitempty"`
	// Shows the sum of all RUM lite sessions (browser and mobile) over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumLiteSessionCountSum *int64 `json:"rum_lite_session_count_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy sessions on Android over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountAndroidSum *int64 `json:"rum_mobile_legacy_session_count_android_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy Sessions on Flutter over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountFlutterSum *int64 `json:"rum_mobile_legacy_session_count_flutter_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy sessions on iOS over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountIosSum *int64 `json:"rum_mobile_legacy_session_count_ios_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy sessions on React Native over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountReactnativeSum *int64 `json:"rum_mobile_legacy_session_count_reactnative_sum,omitempty"`
	// Shows the sum of all mobile RUM legacy sessions on Roku over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLegacySessionCountRokuSum *int64 `json:"rum_mobile_legacy_session_count_roku_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Android over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountAndroidSum *int64 `json:"rum_mobile_lite_session_count_android_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Flutter over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountFlutterSum *int64 `json:"rum_mobile_lite_session_count_flutter_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on iOS over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountIosSum *int64 `json:"rum_mobile_lite_session_count_ios_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Kotlin Multiplatform over all hours within the current date for all organizations.
	RumMobileLiteSessionCountKotlinmultiplatformSum *int64 `json:"rum_mobile_lite_session_count_kotlinmultiplatform_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on React Native over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountReactnativeSum *int64 `json:"rum_mobile_lite_session_count_reactnative_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Roku over all hours within the current date for all organizations (To be introduced on October 1st, 2024).
	RumMobileLiteSessionCountRokuSum *int64 `json:"rum_mobile_lite_session_count_roku_sum,omitempty"`
	// Shows the sum of all mobile RUM lite sessions on Unity over all hours within the current date for all organizations.
	RumMobileLiteSessionCountUnitySum *int64 `json:"rum_mobile_lite_session_count_unity_sum,omitempty"`
	// Shows the sum of all mobile RUM replay sessions on Android over all hours within the current date for the given org.
	RumMobileReplaySessionCountAndroidSum *int64 `json:"rum_mobile_replay_session_count_android_sum,omitempty"`
	// Shows the sum of all mobile RUM replay sessions on iOS over all hours within the current date for the given org.
	RumMobileReplaySessionCountIosSum *int64 `json:"rum_mobile_replay_session_count_ios_sum,omitempty"`
	// Shows the sum of all mobile RUM replay sessions on Kotlin Multiplatform over all hours within the current date for all organizations.
	RumMobileReplaySessionCountKotlinmultiplatformSum *int64 `json:"rum_mobile_replay_session_count_kotlinmultiplatform_sum,omitempty"`
	// Shows the sum of all mobile RUM replay sessions on React Native over all hours within the current date for the given org.
	RumMobileReplaySessionCountReactnativeSum *int64 `json:"rum_mobile_replay_session_count_reactnative_sum,omitempty"`
	// Shows the sum of all RUM Session Replay counts over all hours in the current date for all organizations (To be introduced on October 1st, 2024).
	RumReplaySessionCountSum *int64 `json:"rum_replay_session_count_sum,omitempty"`
	// Shows the sum of all browser RUM lite sessions over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	RumSessionCountSum *int64 `json:"rum_session_count_sum,omitempty"`
	// Sum of all RUM session replay add-on sessions over all hours in the current date for all organizations.
	RumSessionReplayAddOnSum *int64 `json:"rum_session_replay_add_on_sum,omitempty"`
	// Shows the sum of RUM sessions (browser and mobile) over all hours in the current date for all organizations.
	RumTotalSessionCountSum *int64 `json:"rum_total_session_count_sum,omitempty"`
	// Shows the sum of all browser and mobile RUM units over all hours in the current date for all organizations (To be deprecated on October 1st, 2024).
	// Deprecated
	RumUnitsSum *int64 `json:"rum_units_sum,omitempty"`
	// Shows the average of all Software Composition Analysis Fargate tasks over all hours in the current date for the given org.
	ScaFargateCountAvg *int64 `json:"sca_fargate_count_avg,omitempty"`
	// Shows the sum of the high-water marks of all Software Composition Analysis Fargate tasks over all hours in the current date for the given org.
	ScaFargateCountHwm *int64 `json:"sca_fargate_count_hwm,omitempty"`
	// Sum of all APM bytes scanned with sensitive data scanner over all hours in the current date for all organizations.
	SdsApmScannedBytesSum *int64 `json:"sds_apm_scanned_bytes_sum,omitempty"`
	// Sum of all event stream events bytes scanned with sensitive data scanner over all hours in the current date for all organizations.
	SdsEventsScannedBytesSum *int64 `json:"sds_events_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned of logs usage by the Sensitive Data Scanner over all hours in the current month for all organizations.
	SdsLogsScannedBytesSum *int64 `json:"sds_logs_scanned_bytes_sum,omitempty"`
	// Sum of all RUM bytes scanned with sensitive data scanner over all hours in the current date for all organizations.
	SdsRumScannedBytesSum *int64 `json:"sds_rum_scanned_bytes_sum,omitempty"`
	// Shows the sum of all bytes scanned across all usage types by the Sensitive Data Scanner over all hours in the current month for all organizations.
	SdsTotalScannedBytesSum *int64 `json:"sds_total_scanned_bytes_sum,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring for Azure App Service instances for the current date for all organizations.
	ServerlessAppsApmApmAzureAppserviceInstancesAvg *int64 `json:"serverless_apps_apm_apm_azure_appservice_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring for Azure Function instances for the current date for all organizations.
	ServerlessAppsApmApmAzureAzurefunctionInstancesAvg *int64 `json:"serverless_apps_apm_apm_azure_azurefunction_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring for Azure Container App instances for the current date for all organizations.
	ServerlessAppsApmApmAzureContainerappInstancesAvg *int64 `json:"serverless_apps_apm_apm_azure_containerapp_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring for Fargate Elastic Container Service tasks for the current date for all organizations.
	ServerlessAppsApmApmFargateEcsTasksAvg *int64 `json:"serverless_apps_apm_apm_fargate_ecs_tasks_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring for Google Cloud Platform Cloud Function instances for the current date for all organizations.
	ServerlessAppsApmApmGcpCloudfunctionInstancesAvg *int64 `json:"serverless_apps_apm_apm_gcp_cloudfunction_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring for Google Cloud Platform Cloud Run instances for the current date for all organizations.
	ServerlessAppsApmApmGcpCloudrunInstancesAvg *int64 `json:"serverless_apps_apm_apm_gcp_cloudrun_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring for the current date for all organizations.
	ServerlessAppsApmAvg *int64 `json:"serverless_apps_apm_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Azure App Service instances for the current date for all organizations.
	ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg *int64 `json:"serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Azure Function instances for the current date for all organizations.
	ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg *int64 `json:"serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Azure Container App instances for the current date for all organizations.
	ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg *int64 `json:"serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Google Cloud Platform Cloud Function instances for the current date for all organizations.
	ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg *int64 `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Google Cloud Platform Cloud Run instances for the current date for all organizations.
	ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg *int64 `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for the current date for all organizations.
	ServerlessAppsApmExclFargateAvg *int64 `json:"serverless_apps_apm_excl_fargate_avg,omitempty"`
	// Shows the average number of Serverless Apps for Azure Container App instances for the current date for all organizations.
	ServerlessAppsAzureContainerAppInstancesAvg *int64 `json:"serverless_apps_azure_container_app_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps for Azure for the given date and given org.
	ServerlessAppsAzureCountAvg *int64 `json:"serverless_apps_azure_count_avg,omitempty"`
	// Shows the average number of Serverless Apps for Azure Function App instances for the current date for all organizations.
	ServerlessAppsAzureFunctionAppInstancesAvg *int64 `json:"serverless_apps_azure_function_app_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps for Azure Web App instances for the current date for all organizations.
	ServerlessAppsAzureWebAppInstancesAvg *int64 `json:"serverless_apps_azure_web_app_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps for Elastic Container Service for the current date for all organizations.
	ServerlessAppsEcsAvg *int64 `json:"serverless_apps_ecs_avg,omitempty"`
	// Shows the average number of Serverless Apps for Elastic Kubernetes Service for the current date for all organizations.
	ServerlessAppsEksAvg *int64 `json:"serverless_apps_eks_avg,omitempty"`
	// Shows the average number of Serverless Apps excluding Fargate for the current date for all organizations.
	ServerlessAppsExclFargateAvg *int64 `json:"serverless_apps_excl_fargate_avg,omitempty"`
	// Shows the average number of Serverless Apps excluding Fargate for Azure Container App instances for the current date for all organizations.
	ServerlessAppsExclFargateAzureContainerAppInstancesAvg *int64 `json:"serverless_apps_excl_fargate_azure_container_app_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps excluding Fargate for Azure Function App instances for the current date for all organizations.
	ServerlessAppsExclFargateAzureFunctionAppInstancesAvg *int64 `json:"serverless_apps_excl_fargate_azure_function_app_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps excluding Fargate for Azure Web App instances for the current date for all organizations.
	ServerlessAppsExclFargateAzureWebAppInstancesAvg *int64 `json:"serverless_apps_excl_fargate_azure_web_app_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps excluding Fargate for Google Cloud Platform Cloud Functions instances for the current date for all organizations.
	ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg *int64 `json:"serverless_apps_excl_fargate_google_cloud_functions_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps excluding Fargate for Google Cloud Platform Cloud Run instances for the current date for all organizations.
	ServerlessAppsExclFargateGoogleCloudRunInstancesAvg *int64 `json:"serverless_apps_excl_fargate_google_cloud_run_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps for Google Cloud Platform Cloud Functions instances for the current date for all organizations.
	ServerlessAppsGoogleCloudFunctionsInstancesAvg *int64 `json:"serverless_apps_google_cloud_functions_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps for Google Cloud Platform Cloud Run instances for the current date for all organizations.
	ServerlessAppsGoogleCloudRunInstancesAvg *int64 `json:"serverless_apps_google_cloud_run_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps for Google Cloud for the given date and given org.
	ServerlessAppsGoogleCountAvg *int64 `json:"serverless_apps_google_count_avg,omitempty"`
	// Shows the average number of Serverless Apps for Azure and Google Cloud for the given date and given org.
	ServerlessAppsTotalCountAvg *int64 `json:"serverless_apps_total_count_avg,omitempty"`
	// Shows the sum of all log events analyzed by Cloud SIEM over all hours in the current date for the given org.
	SiemAnalyzedLogsAddOnCountSum *int64 `json:"siem_analyzed_logs_add_on_count_sum,omitempty"`
	// Shows the sum of all Synthetic browser tests over all hours in the current date for all organizations.
	SyntheticsBrowserCheckCallsCountSum *int64 `json:"synthetics_browser_check_calls_count_sum,omitempty"`
	// Shows the sum of all Synthetic API tests over all hours in the current date for all organizations.
	SyntheticsCheckCallsCountSum *int64 `json:"synthetics_check_calls_count_sum,omitempty"`
	// Shows the sum of all Synthetic mobile application tests over all hours in the current date for all organizations.
	SyntheticsMobileTestRunsSum *int64 `json:"synthetics_mobile_test_runs_sum,omitempty"`
	// Shows the high-water mark of used synthetics parallel testing slots over all hours in the current date for all organizations.
	SyntheticsParallelTestingMaxSlotsHwm *int64 `json:"synthetics_parallel_testing_max_slots_hwm,omitempty"`
	// Shows the sum of all Indexed Spans indexed over all hours in the current date for all organizations.
	TraceSearchIndexedEventsCountSum *int64 `json:"trace_search_indexed_events_count_sum,omitempty"`
	// Shows the sum of all ingested APM span bytes over all hours in the current date for all organizations.
	TwolIngestedEventsBytesSum *int64 `json:"twol_ingested_events_bytes_sum,omitempty"`
	// Shows the 99th percentile of all universal service management hosts over all hours in the current date for the given org.
	UniversalServiceMonitoringHostTop99p *int64 `json:"universal_service_monitoring_host_top99p,omitempty"`
	// Shows the 99th percentile of all vSphere hosts over all hours in the current date for all organizations.
	VsphereHostTop99p *int64 `json:"vsphere_host_top99p,omitempty"`
	// Shows the 99th percentile of all Application Vulnerability Management hosts over all hours in the current date for the given org.
	VulnManagementHostCountTop99p *int64 `json:"vuln_management_host_count_top99p,omitempty"`
	// Sum of all workflows executed over all hours in the current date for all organizations.
	WorkflowExecutionsUsageSum *int64 `json:"workflow_executions_usage_sum,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageSummaryDate instantiates a new UsageSummaryDate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSummaryDate() *UsageSummaryDate {
	this := UsageSummaryDate{}
	return &this
}

// NewUsageSummaryDateWithDefaults instantiates a new UsageSummaryDate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSummaryDateWithDefaults() *UsageSummaryDate {
	this := UsageSummaryDate{}
	return &this
}

// GetAgentHostTop99p returns the AgentHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAgentHostTop99p() int64 {
	if o == nil || o.AgentHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.AgentHostTop99p
}

// GetAgentHostTop99pOk returns a tuple with the AgentHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAgentHostTop99pOk() (*int64, bool) {
	if o == nil || o.AgentHostTop99p == nil {
		return nil, false
	}
	return o.AgentHostTop99p, true
}

// HasAgentHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAgentHostTop99p() bool {
	return o != nil && o.AgentHostTop99p != nil
}

// SetAgentHostTop99p gets a reference to the given int64 and assigns it to the AgentHostTop99p field.
func (o *UsageSummaryDate) SetAgentHostTop99p(v int64) {
	o.AgentHostTop99p = &v
}

// GetApmAzureAppServiceHostTop99p returns the ApmAzureAppServiceHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetApmAzureAppServiceHostTop99p() int64 {
	if o == nil || o.ApmAzureAppServiceHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.ApmAzureAppServiceHostTop99p
}

// GetApmAzureAppServiceHostTop99pOk returns a tuple with the ApmAzureAppServiceHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetApmAzureAppServiceHostTop99pOk() (*int64, bool) {
	if o == nil || o.ApmAzureAppServiceHostTop99p == nil {
		return nil, false
	}
	return o.ApmAzureAppServiceHostTop99p, true
}

// HasApmAzureAppServiceHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmAzureAppServiceHostTop99p() bool {
	return o != nil && o.ApmAzureAppServiceHostTop99p != nil
}

// SetApmAzureAppServiceHostTop99p gets a reference to the given int64 and assigns it to the ApmAzureAppServiceHostTop99p field.
func (o *UsageSummaryDate) SetApmAzureAppServiceHostTop99p(v int64) {
	o.ApmAzureAppServiceHostTop99p = &v
}

// GetApmDevsecopsHostTop99p returns the ApmDevsecopsHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetApmDevsecopsHostTop99p() int64 {
	if o == nil || o.ApmDevsecopsHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.ApmDevsecopsHostTop99p
}

// GetApmDevsecopsHostTop99pOk returns a tuple with the ApmDevsecopsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetApmDevsecopsHostTop99pOk() (*int64, bool) {
	if o == nil || o.ApmDevsecopsHostTop99p == nil {
		return nil, false
	}
	return o.ApmDevsecopsHostTop99p, true
}

// HasApmDevsecopsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmDevsecopsHostTop99p() bool {
	return o != nil && o.ApmDevsecopsHostTop99p != nil
}

// SetApmDevsecopsHostTop99p gets a reference to the given int64 and assigns it to the ApmDevsecopsHostTop99p field.
func (o *UsageSummaryDate) SetApmDevsecopsHostTop99p(v int64) {
	o.ApmDevsecopsHostTop99p = &v
}

// GetApmEnterpriseStandaloneHostsTop99p returns the ApmEnterpriseStandaloneHostsTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetApmEnterpriseStandaloneHostsTop99p() int64 {
	if o == nil || o.ApmEnterpriseStandaloneHostsTop99p == nil {
		var ret int64
		return ret
	}
	return *o.ApmEnterpriseStandaloneHostsTop99p
}

// GetApmEnterpriseStandaloneHostsTop99pOk returns a tuple with the ApmEnterpriseStandaloneHostsTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetApmEnterpriseStandaloneHostsTop99pOk() (*int64, bool) {
	if o == nil || o.ApmEnterpriseStandaloneHostsTop99p == nil {
		return nil, false
	}
	return o.ApmEnterpriseStandaloneHostsTop99p, true
}

// HasApmEnterpriseStandaloneHostsTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmEnterpriseStandaloneHostsTop99p() bool {
	return o != nil && o.ApmEnterpriseStandaloneHostsTop99p != nil
}

// SetApmEnterpriseStandaloneHostsTop99p gets a reference to the given int64 and assigns it to the ApmEnterpriseStandaloneHostsTop99p field.
func (o *UsageSummaryDate) SetApmEnterpriseStandaloneHostsTop99p(v int64) {
	o.ApmEnterpriseStandaloneHostsTop99p = &v
}

// GetApmFargateCountAvg returns the ApmFargateCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetApmFargateCountAvg() int64 {
	if o == nil || o.ApmFargateCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.ApmFargateCountAvg
}

// GetApmFargateCountAvgOk returns a tuple with the ApmFargateCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetApmFargateCountAvgOk() (*int64, bool) {
	if o == nil || o.ApmFargateCountAvg == nil {
		return nil, false
	}
	return o.ApmFargateCountAvg, true
}

// HasApmFargateCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmFargateCountAvg() bool {
	return o != nil && o.ApmFargateCountAvg != nil
}

// SetApmFargateCountAvg gets a reference to the given int64 and assigns it to the ApmFargateCountAvg field.
func (o *UsageSummaryDate) SetApmFargateCountAvg(v int64) {
	o.ApmFargateCountAvg = &v
}

// GetApmHostTop99p returns the ApmHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetApmHostTop99p() int64 {
	if o == nil || o.ApmHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.ApmHostTop99p
}

// GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetApmHostTop99pOk() (*int64, bool) {
	if o == nil || o.ApmHostTop99p == nil {
		return nil, false
	}
	return o.ApmHostTop99p, true
}

// HasApmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmHostTop99p() bool {
	return o != nil && o.ApmHostTop99p != nil
}

// SetApmHostTop99p gets a reference to the given int64 and assigns it to the ApmHostTop99p field.
func (o *UsageSummaryDate) SetApmHostTop99p(v int64) {
	o.ApmHostTop99p = &v
}

// GetApmProStandaloneHostsTop99p returns the ApmProStandaloneHostsTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetApmProStandaloneHostsTop99p() int64 {
	if o == nil || o.ApmProStandaloneHostsTop99p == nil {
		var ret int64
		return ret
	}
	return *o.ApmProStandaloneHostsTop99p
}

// GetApmProStandaloneHostsTop99pOk returns a tuple with the ApmProStandaloneHostsTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetApmProStandaloneHostsTop99pOk() (*int64, bool) {
	if o == nil || o.ApmProStandaloneHostsTop99p == nil {
		return nil, false
	}
	return o.ApmProStandaloneHostsTop99p, true
}

// HasApmProStandaloneHostsTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasApmProStandaloneHostsTop99p() bool {
	return o != nil && o.ApmProStandaloneHostsTop99p != nil
}

// SetApmProStandaloneHostsTop99p gets a reference to the given int64 and assigns it to the ApmProStandaloneHostsTop99p field.
func (o *UsageSummaryDate) SetApmProStandaloneHostsTop99p(v int64) {
	o.ApmProStandaloneHostsTop99p = &v
}

// GetAppsecFargateCountAvg returns the AppsecFargateCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAppsecFargateCountAvg() int64 {
	if o == nil || o.AppsecFargateCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.AppsecFargateCountAvg
}

// GetAppsecFargateCountAvgOk returns a tuple with the AppsecFargateCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAppsecFargateCountAvgOk() (*int64, bool) {
	if o == nil || o.AppsecFargateCountAvg == nil {
		return nil, false
	}
	return o.AppsecFargateCountAvg, true
}

// HasAppsecFargateCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAppsecFargateCountAvg() bool {
	return o != nil && o.AppsecFargateCountAvg != nil
}

// SetAppsecFargateCountAvg gets a reference to the given int64 and assigns it to the AppsecFargateCountAvg field.
func (o *UsageSummaryDate) SetAppsecFargateCountAvg(v int64) {
	o.AppsecFargateCountAvg = &v
}

// GetAsmServerlessSum returns the AsmServerlessSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAsmServerlessSum() int64 {
	if o == nil || o.AsmServerlessSum == nil {
		var ret int64
		return ret
	}
	return *o.AsmServerlessSum
}

// GetAsmServerlessSumOk returns a tuple with the AsmServerlessSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAsmServerlessSumOk() (*int64, bool) {
	if o == nil || o.AsmServerlessSum == nil {
		return nil, false
	}
	return o.AsmServerlessSum, true
}

// HasAsmServerlessSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAsmServerlessSum() bool {
	return o != nil && o.AsmServerlessSum != nil
}

// SetAsmServerlessSum gets a reference to the given int64 and assigns it to the AsmServerlessSum field.
func (o *UsageSummaryDate) SetAsmServerlessSum(v int64) {
	o.AsmServerlessSum = &v
}

// GetAuditLogsLinesIndexedSum returns the AuditLogsLinesIndexedSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetAuditLogsLinesIndexedSum() int64 {
	if o == nil || o.AuditLogsLinesIndexedSum == nil {
		var ret int64
		return ret
	}
	return *o.AuditLogsLinesIndexedSum
}

// GetAuditLogsLinesIndexedSumOk returns a tuple with the AuditLogsLinesIndexedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetAuditLogsLinesIndexedSumOk() (*int64, bool) {
	if o == nil || o.AuditLogsLinesIndexedSum == nil {
		return nil, false
	}
	return o.AuditLogsLinesIndexedSum, true
}

// HasAuditLogsLinesIndexedSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAuditLogsLinesIndexedSum() bool {
	return o != nil && o.AuditLogsLinesIndexedSum != nil
}

// SetAuditLogsLinesIndexedSum gets a reference to the given int64 and assigns it to the AuditLogsLinesIndexedSum field.
// Deprecated
func (o *UsageSummaryDate) SetAuditLogsLinesIndexedSum(v int64) {
	o.AuditLogsLinesIndexedSum = &v
}

// GetAuditTrailEnabledHwm returns the AuditTrailEnabledHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAuditTrailEnabledHwm() int64 {
	if o == nil || o.AuditTrailEnabledHwm == nil {
		var ret int64
		return ret
	}
	return *o.AuditTrailEnabledHwm
}

// GetAuditTrailEnabledHwmOk returns a tuple with the AuditTrailEnabledHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAuditTrailEnabledHwmOk() (*int64, bool) {
	if o == nil || o.AuditTrailEnabledHwm == nil {
		return nil, false
	}
	return o.AuditTrailEnabledHwm, true
}

// HasAuditTrailEnabledHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAuditTrailEnabledHwm() bool {
	return o != nil && o.AuditTrailEnabledHwm != nil
}

// SetAuditTrailEnabledHwm gets a reference to the given int64 and assigns it to the AuditTrailEnabledHwm field.
func (o *UsageSummaryDate) SetAuditTrailEnabledHwm(v int64) {
	o.AuditTrailEnabledHwm = &v
}

// GetAvgProfiledFargateTasks returns the AvgProfiledFargateTasks field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAvgProfiledFargateTasks() int64 {
	if o == nil || o.AvgProfiledFargateTasks == nil {
		var ret int64
		return ret
	}
	return *o.AvgProfiledFargateTasks
}

// GetAvgProfiledFargateTasksOk returns a tuple with the AvgProfiledFargateTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAvgProfiledFargateTasksOk() (*int64, bool) {
	if o == nil || o.AvgProfiledFargateTasks == nil {
		return nil, false
	}
	return o.AvgProfiledFargateTasks, true
}

// HasAvgProfiledFargateTasks returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAvgProfiledFargateTasks() bool {
	return o != nil && o.AvgProfiledFargateTasks != nil
}

// SetAvgProfiledFargateTasks gets a reference to the given int64 and assigns it to the AvgProfiledFargateTasks field.
func (o *UsageSummaryDate) SetAvgProfiledFargateTasks(v int64) {
	o.AvgProfiledFargateTasks = &v
}

// GetAwsHostTop99p returns the AwsHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAwsHostTop99p() int64 {
	if o == nil || o.AwsHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.AwsHostTop99p
}

// GetAwsHostTop99pOk returns a tuple with the AwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAwsHostTop99pOk() (*int64, bool) {
	if o == nil || o.AwsHostTop99p == nil {
		return nil, false
	}
	return o.AwsHostTop99p, true
}

// HasAwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAwsHostTop99p() bool {
	return o != nil && o.AwsHostTop99p != nil
}

// SetAwsHostTop99p gets a reference to the given int64 and assigns it to the AwsHostTop99p field.
func (o *UsageSummaryDate) SetAwsHostTop99p(v int64) {
	o.AwsHostTop99p = &v
}

// GetAwsLambdaFuncCount returns the AwsLambdaFuncCount field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAwsLambdaFuncCount() int64 {
	if o == nil || o.AwsLambdaFuncCount == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaFuncCount
}

// GetAwsLambdaFuncCountOk returns a tuple with the AwsLambdaFuncCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAwsLambdaFuncCountOk() (*int64, bool) {
	if o == nil || o.AwsLambdaFuncCount == nil {
		return nil, false
	}
	return o.AwsLambdaFuncCount, true
}

// HasAwsLambdaFuncCount returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAwsLambdaFuncCount() bool {
	return o != nil && o.AwsLambdaFuncCount != nil
}

// SetAwsLambdaFuncCount gets a reference to the given int64 and assigns it to the AwsLambdaFuncCount field.
func (o *UsageSummaryDate) SetAwsLambdaFuncCount(v int64) {
	o.AwsLambdaFuncCount = &v
}

// GetAwsLambdaInvocationsSum returns the AwsLambdaInvocationsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAwsLambdaInvocationsSum() int64 {
	if o == nil || o.AwsLambdaInvocationsSum == nil {
		var ret int64
		return ret
	}
	return *o.AwsLambdaInvocationsSum
}

// GetAwsLambdaInvocationsSumOk returns a tuple with the AwsLambdaInvocationsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAwsLambdaInvocationsSumOk() (*int64, bool) {
	if o == nil || o.AwsLambdaInvocationsSum == nil {
		return nil, false
	}
	return o.AwsLambdaInvocationsSum, true
}

// HasAwsLambdaInvocationsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAwsLambdaInvocationsSum() bool {
	return o != nil && o.AwsLambdaInvocationsSum != nil
}

// SetAwsLambdaInvocationsSum gets a reference to the given int64 and assigns it to the AwsLambdaInvocationsSum field.
func (o *UsageSummaryDate) SetAwsLambdaInvocationsSum(v int64) {
	o.AwsLambdaInvocationsSum = &v
}

// GetAzureAppServiceTop99p returns the AzureAppServiceTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAzureAppServiceTop99p() int64 {
	if o == nil || o.AzureAppServiceTop99p == nil {
		var ret int64
		return ret
	}
	return *o.AzureAppServiceTop99p
}

// GetAzureAppServiceTop99pOk returns a tuple with the AzureAppServiceTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAzureAppServiceTop99pOk() (*int64, bool) {
	if o == nil || o.AzureAppServiceTop99p == nil {
		return nil, false
	}
	return o.AzureAppServiceTop99p, true
}

// HasAzureAppServiceTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAzureAppServiceTop99p() bool {
	return o != nil && o.AzureAppServiceTop99p != nil
}

// SetAzureAppServiceTop99p gets a reference to the given int64 and assigns it to the AzureAppServiceTop99p field.
func (o *UsageSummaryDate) SetAzureAppServiceTop99p(v int64) {
	o.AzureAppServiceTop99p = &v
}

// GetBillableIngestedBytesSum returns the BillableIngestedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetBillableIngestedBytesSum() int64 {
	if o == nil || o.BillableIngestedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.BillableIngestedBytesSum
}

// GetBillableIngestedBytesSumOk returns a tuple with the BillableIngestedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetBillableIngestedBytesSumOk() (*int64, bool) {
	if o == nil || o.BillableIngestedBytesSum == nil {
		return nil, false
	}
	return o.BillableIngestedBytesSum, true
}

// HasBillableIngestedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBillableIngestedBytesSum() bool {
	return o != nil && o.BillableIngestedBytesSum != nil
}

// SetBillableIngestedBytesSum gets a reference to the given int64 and assigns it to the BillableIngestedBytesSum field.
func (o *UsageSummaryDate) SetBillableIngestedBytesSum(v int64) {
	o.BillableIngestedBytesSum = &v
}

// GetBitsAiInvestigationsSum returns the BitsAiInvestigationsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetBitsAiInvestigationsSum() int64 {
	if o == nil || o.BitsAiInvestigationsSum == nil {
		var ret int64
		return ret
	}
	return *o.BitsAiInvestigationsSum
}

// GetBitsAiInvestigationsSumOk returns a tuple with the BitsAiInvestigationsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetBitsAiInvestigationsSumOk() (*int64, bool) {
	if o == nil || o.BitsAiInvestigationsSum == nil {
		return nil, false
	}
	return o.BitsAiInvestigationsSum, true
}

// HasBitsAiInvestigationsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBitsAiInvestigationsSum() bool {
	return o != nil && o.BitsAiInvestigationsSum != nil
}

// SetBitsAiInvestigationsSum gets a reference to the given int64 and assigns it to the BitsAiInvestigationsSum field.
func (o *UsageSummaryDate) SetBitsAiInvestigationsSum(v int64) {
	o.BitsAiInvestigationsSum = &v
}

// GetBrowserRumLiteSessionCountSum returns the BrowserRumLiteSessionCountSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetBrowserRumLiteSessionCountSum() int64 {
	if o == nil || o.BrowserRumLiteSessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumLiteSessionCountSum
}

// GetBrowserRumLiteSessionCountSumOk returns a tuple with the BrowserRumLiteSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetBrowserRumLiteSessionCountSumOk() (*int64, bool) {
	if o == nil || o.BrowserRumLiteSessionCountSum == nil {
		return nil, false
	}
	return o.BrowserRumLiteSessionCountSum, true
}

// HasBrowserRumLiteSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBrowserRumLiteSessionCountSum() bool {
	return o != nil && o.BrowserRumLiteSessionCountSum != nil
}

// SetBrowserRumLiteSessionCountSum gets a reference to the given int64 and assigns it to the BrowserRumLiteSessionCountSum field.
// Deprecated
func (o *UsageSummaryDate) SetBrowserRumLiteSessionCountSum(v int64) {
	o.BrowserRumLiteSessionCountSum = &v
}

// GetBrowserRumReplaySessionCountSum returns the BrowserRumReplaySessionCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetBrowserRumReplaySessionCountSum() int64 {
	if o == nil || o.BrowserRumReplaySessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumReplaySessionCountSum
}

// GetBrowserRumReplaySessionCountSumOk returns a tuple with the BrowserRumReplaySessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetBrowserRumReplaySessionCountSumOk() (*int64, bool) {
	if o == nil || o.BrowserRumReplaySessionCountSum == nil {
		return nil, false
	}
	return o.BrowserRumReplaySessionCountSum, true
}

// HasBrowserRumReplaySessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBrowserRumReplaySessionCountSum() bool {
	return o != nil && o.BrowserRumReplaySessionCountSum != nil
}

// SetBrowserRumReplaySessionCountSum gets a reference to the given int64 and assigns it to the BrowserRumReplaySessionCountSum field.
func (o *UsageSummaryDate) SetBrowserRumReplaySessionCountSum(v int64) {
	o.BrowserRumReplaySessionCountSum = &v
}

// GetBrowserRumUnitsSum returns the BrowserRumUnitsSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetBrowserRumUnitsSum() int64 {
	if o == nil || o.BrowserRumUnitsSum == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumUnitsSum
}

// GetBrowserRumUnitsSumOk returns a tuple with the BrowserRumUnitsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetBrowserRumUnitsSumOk() (*int64, bool) {
	if o == nil || o.BrowserRumUnitsSum == nil {
		return nil, false
	}
	return o.BrowserRumUnitsSum, true
}

// HasBrowserRumUnitsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasBrowserRumUnitsSum() bool {
	return o != nil && o.BrowserRumUnitsSum != nil
}

// SetBrowserRumUnitsSum gets a reference to the given int64 and assigns it to the BrowserRumUnitsSum field.
// Deprecated
func (o *UsageSummaryDate) SetBrowserRumUnitsSum(v int64) {
	o.BrowserRumUnitsSum = &v
}

// GetCiPipelineIndexedSpansSum returns the CiPipelineIndexedSpansSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCiPipelineIndexedSpansSum() int64 {
	if o == nil || o.CiPipelineIndexedSpansSum == nil {
		var ret int64
		return ret
	}
	return *o.CiPipelineIndexedSpansSum
}

// GetCiPipelineIndexedSpansSumOk returns a tuple with the CiPipelineIndexedSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCiPipelineIndexedSpansSumOk() (*int64, bool) {
	if o == nil || o.CiPipelineIndexedSpansSum == nil {
		return nil, false
	}
	return o.CiPipelineIndexedSpansSum, true
}

// HasCiPipelineIndexedSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCiPipelineIndexedSpansSum() bool {
	return o != nil && o.CiPipelineIndexedSpansSum != nil
}

// SetCiPipelineIndexedSpansSum gets a reference to the given int64 and assigns it to the CiPipelineIndexedSpansSum field.
func (o *UsageSummaryDate) SetCiPipelineIndexedSpansSum(v int64) {
	o.CiPipelineIndexedSpansSum = &v
}

// GetCiTestIndexedSpansSum returns the CiTestIndexedSpansSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCiTestIndexedSpansSum() int64 {
	if o == nil || o.CiTestIndexedSpansSum == nil {
		var ret int64
		return ret
	}
	return *o.CiTestIndexedSpansSum
}

// GetCiTestIndexedSpansSumOk returns a tuple with the CiTestIndexedSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCiTestIndexedSpansSumOk() (*int64, bool) {
	if o == nil || o.CiTestIndexedSpansSum == nil {
		return nil, false
	}
	return o.CiTestIndexedSpansSum, true
}

// HasCiTestIndexedSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCiTestIndexedSpansSum() bool {
	return o != nil && o.CiTestIndexedSpansSum != nil
}

// SetCiTestIndexedSpansSum gets a reference to the given int64 and assigns it to the CiTestIndexedSpansSum field.
func (o *UsageSummaryDate) SetCiTestIndexedSpansSum(v int64) {
	o.CiTestIndexedSpansSum = &v
}

// GetCiVisibilityItrCommittersHwm returns the CiVisibilityItrCommittersHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCiVisibilityItrCommittersHwm() int64 {
	if o == nil || o.CiVisibilityItrCommittersHwm == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityItrCommittersHwm
}

// GetCiVisibilityItrCommittersHwmOk returns a tuple with the CiVisibilityItrCommittersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCiVisibilityItrCommittersHwmOk() (*int64, bool) {
	if o == nil || o.CiVisibilityItrCommittersHwm == nil {
		return nil, false
	}
	return o.CiVisibilityItrCommittersHwm, true
}

// HasCiVisibilityItrCommittersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCiVisibilityItrCommittersHwm() bool {
	return o != nil && o.CiVisibilityItrCommittersHwm != nil
}

// SetCiVisibilityItrCommittersHwm gets a reference to the given int64 and assigns it to the CiVisibilityItrCommittersHwm field.
func (o *UsageSummaryDate) SetCiVisibilityItrCommittersHwm(v int64) {
	o.CiVisibilityItrCommittersHwm = &v
}

// GetCiVisibilityPipelineCommittersHwm returns the CiVisibilityPipelineCommittersHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCiVisibilityPipelineCommittersHwm() int64 {
	if o == nil || o.CiVisibilityPipelineCommittersHwm == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityPipelineCommittersHwm
}

// GetCiVisibilityPipelineCommittersHwmOk returns a tuple with the CiVisibilityPipelineCommittersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCiVisibilityPipelineCommittersHwmOk() (*int64, bool) {
	if o == nil || o.CiVisibilityPipelineCommittersHwm == nil {
		return nil, false
	}
	return o.CiVisibilityPipelineCommittersHwm, true
}

// HasCiVisibilityPipelineCommittersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCiVisibilityPipelineCommittersHwm() bool {
	return o != nil && o.CiVisibilityPipelineCommittersHwm != nil
}

// SetCiVisibilityPipelineCommittersHwm gets a reference to the given int64 and assigns it to the CiVisibilityPipelineCommittersHwm field.
func (o *UsageSummaryDate) SetCiVisibilityPipelineCommittersHwm(v int64) {
	o.CiVisibilityPipelineCommittersHwm = &v
}

// GetCiVisibilityTestCommittersHwm returns the CiVisibilityTestCommittersHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCiVisibilityTestCommittersHwm() int64 {
	if o == nil || o.CiVisibilityTestCommittersHwm == nil {
		var ret int64
		return ret
	}
	return *o.CiVisibilityTestCommittersHwm
}

// GetCiVisibilityTestCommittersHwmOk returns a tuple with the CiVisibilityTestCommittersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCiVisibilityTestCommittersHwmOk() (*int64, bool) {
	if o == nil || o.CiVisibilityTestCommittersHwm == nil {
		return nil, false
	}
	return o.CiVisibilityTestCommittersHwm, true
}

// HasCiVisibilityTestCommittersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCiVisibilityTestCommittersHwm() bool {
	return o != nil && o.CiVisibilityTestCommittersHwm != nil
}

// SetCiVisibilityTestCommittersHwm gets a reference to the given int64 and assigns it to the CiVisibilityTestCommittersHwm field.
func (o *UsageSummaryDate) SetCiVisibilityTestCommittersHwm(v int64) {
	o.CiVisibilityTestCommittersHwm = &v
}

// GetCloudCostManagementAwsHostCountAvg returns the CloudCostManagementAwsHostCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCloudCostManagementAwsHostCountAvg() int64 {
	if o == nil || o.CloudCostManagementAwsHostCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementAwsHostCountAvg
}

// GetCloudCostManagementAwsHostCountAvgOk returns a tuple with the CloudCostManagementAwsHostCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCloudCostManagementAwsHostCountAvgOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementAwsHostCountAvg == nil {
		return nil, false
	}
	return o.CloudCostManagementAwsHostCountAvg, true
}

// HasCloudCostManagementAwsHostCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCloudCostManagementAwsHostCountAvg() bool {
	return o != nil && o.CloudCostManagementAwsHostCountAvg != nil
}

// SetCloudCostManagementAwsHostCountAvg gets a reference to the given int64 and assigns it to the CloudCostManagementAwsHostCountAvg field.
func (o *UsageSummaryDate) SetCloudCostManagementAwsHostCountAvg(v int64) {
	o.CloudCostManagementAwsHostCountAvg = &v
}

// GetCloudCostManagementAzureHostCountAvg returns the CloudCostManagementAzureHostCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCloudCostManagementAzureHostCountAvg() int64 {
	if o == nil || o.CloudCostManagementAzureHostCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementAzureHostCountAvg
}

// GetCloudCostManagementAzureHostCountAvgOk returns a tuple with the CloudCostManagementAzureHostCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCloudCostManagementAzureHostCountAvgOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementAzureHostCountAvg == nil {
		return nil, false
	}
	return o.CloudCostManagementAzureHostCountAvg, true
}

// HasCloudCostManagementAzureHostCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCloudCostManagementAzureHostCountAvg() bool {
	return o != nil && o.CloudCostManagementAzureHostCountAvg != nil
}

// SetCloudCostManagementAzureHostCountAvg gets a reference to the given int64 and assigns it to the CloudCostManagementAzureHostCountAvg field.
func (o *UsageSummaryDate) SetCloudCostManagementAzureHostCountAvg(v int64) {
	o.CloudCostManagementAzureHostCountAvg = &v
}

// GetCloudCostManagementGcpHostCountAvg returns the CloudCostManagementGcpHostCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCloudCostManagementGcpHostCountAvg() int64 {
	if o == nil || o.CloudCostManagementGcpHostCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementGcpHostCountAvg
}

// GetCloudCostManagementGcpHostCountAvgOk returns a tuple with the CloudCostManagementGcpHostCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCloudCostManagementGcpHostCountAvgOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementGcpHostCountAvg == nil {
		return nil, false
	}
	return o.CloudCostManagementGcpHostCountAvg, true
}

// HasCloudCostManagementGcpHostCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCloudCostManagementGcpHostCountAvg() bool {
	return o != nil && o.CloudCostManagementGcpHostCountAvg != nil
}

// SetCloudCostManagementGcpHostCountAvg gets a reference to the given int64 and assigns it to the CloudCostManagementGcpHostCountAvg field.
func (o *UsageSummaryDate) SetCloudCostManagementGcpHostCountAvg(v int64) {
	o.CloudCostManagementGcpHostCountAvg = &v
}

// GetCloudCostManagementHostCountAvg returns the CloudCostManagementHostCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCloudCostManagementHostCountAvg() int64 {
	if o == nil || o.CloudCostManagementHostCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementHostCountAvg
}

// GetCloudCostManagementHostCountAvgOk returns a tuple with the CloudCostManagementHostCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCloudCostManagementHostCountAvgOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementHostCountAvg == nil {
		return nil, false
	}
	return o.CloudCostManagementHostCountAvg, true
}

// HasCloudCostManagementHostCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCloudCostManagementHostCountAvg() bool {
	return o != nil && o.CloudCostManagementHostCountAvg != nil
}

// SetCloudCostManagementHostCountAvg gets a reference to the given int64 and assigns it to the CloudCostManagementHostCountAvg field.
func (o *UsageSummaryDate) SetCloudCostManagementHostCountAvg(v int64) {
	o.CloudCostManagementHostCountAvg = &v
}

// GetCloudCostManagementOciHostCountAvg returns the CloudCostManagementOciHostCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCloudCostManagementOciHostCountAvg() int64 {
	if o == nil || o.CloudCostManagementOciHostCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.CloudCostManagementOciHostCountAvg
}

// GetCloudCostManagementOciHostCountAvgOk returns a tuple with the CloudCostManagementOciHostCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCloudCostManagementOciHostCountAvgOk() (*int64, bool) {
	if o == nil || o.CloudCostManagementOciHostCountAvg == nil {
		return nil, false
	}
	return o.CloudCostManagementOciHostCountAvg, true
}

// HasCloudCostManagementOciHostCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCloudCostManagementOciHostCountAvg() bool {
	return o != nil && o.CloudCostManagementOciHostCountAvg != nil
}

// SetCloudCostManagementOciHostCountAvg gets a reference to the given int64 and assigns it to the CloudCostManagementOciHostCountAvg field.
func (o *UsageSummaryDate) SetCloudCostManagementOciHostCountAvg(v int64) {
	o.CloudCostManagementOciHostCountAvg = &v
}

// GetCloudSiemEventsSum returns the CloudSiemEventsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCloudSiemEventsSum() int64 {
	if o == nil || o.CloudSiemEventsSum == nil {
		var ret int64
		return ret
	}
	return *o.CloudSiemEventsSum
}

// GetCloudSiemEventsSumOk returns a tuple with the CloudSiemEventsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCloudSiemEventsSumOk() (*int64, bool) {
	if o == nil || o.CloudSiemEventsSum == nil {
		return nil, false
	}
	return o.CloudSiemEventsSum, true
}

// HasCloudSiemEventsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCloudSiemEventsSum() bool {
	return o != nil && o.CloudSiemEventsSum != nil
}

// SetCloudSiemEventsSum gets a reference to the given int64 and assigns it to the CloudSiemEventsSum field.
func (o *UsageSummaryDate) SetCloudSiemEventsSum(v int64) {
	o.CloudSiemEventsSum = &v
}

// GetCodeAnalysisSaCommittersHwm returns the CodeAnalysisSaCommittersHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCodeAnalysisSaCommittersHwm() int64 {
	if o == nil || o.CodeAnalysisSaCommittersHwm == nil {
		var ret int64
		return ret
	}
	return *o.CodeAnalysisSaCommittersHwm
}

// GetCodeAnalysisSaCommittersHwmOk returns a tuple with the CodeAnalysisSaCommittersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCodeAnalysisSaCommittersHwmOk() (*int64, bool) {
	if o == nil || o.CodeAnalysisSaCommittersHwm == nil {
		return nil, false
	}
	return o.CodeAnalysisSaCommittersHwm, true
}

// HasCodeAnalysisSaCommittersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCodeAnalysisSaCommittersHwm() bool {
	return o != nil && o.CodeAnalysisSaCommittersHwm != nil
}

// SetCodeAnalysisSaCommittersHwm gets a reference to the given int64 and assigns it to the CodeAnalysisSaCommittersHwm field.
func (o *UsageSummaryDate) SetCodeAnalysisSaCommittersHwm(v int64) {
	o.CodeAnalysisSaCommittersHwm = &v
}

// GetCodeAnalysisScaCommittersHwm returns the CodeAnalysisScaCommittersHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCodeAnalysisScaCommittersHwm() int64 {
	if o == nil || o.CodeAnalysisScaCommittersHwm == nil {
		var ret int64
		return ret
	}
	return *o.CodeAnalysisScaCommittersHwm
}

// GetCodeAnalysisScaCommittersHwmOk returns a tuple with the CodeAnalysisScaCommittersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCodeAnalysisScaCommittersHwmOk() (*int64, bool) {
	if o == nil || o.CodeAnalysisScaCommittersHwm == nil {
		return nil, false
	}
	return o.CodeAnalysisScaCommittersHwm, true
}

// HasCodeAnalysisScaCommittersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCodeAnalysisScaCommittersHwm() bool {
	return o != nil && o.CodeAnalysisScaCommittersHwm != nil
}

// SetCodeAnalysisScaCommittersHwm gets a reference to the given int64 and assigns it to the CodeAnalysisScaCommittersHwm field.
func (o *UsageSummaryDate) SetCodeAnalysisScaCommittersHwm(v int64) {
	o.CodeAnalysisScaCommittersHwm = &v
}

// GetCodeSecurityHostTop99p returns the CodeSecurityHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCodeSecurityHostTop99p() int64 {
	if o == nil || o.CodeSecurityHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CodeSecurityHostTop99p
}

// GetCodeSecurityHostTop99pOk returns a tuple with the CodeSecurityHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCodeSecurityHostTop99pOk() (*int64, bool) {
	if o == nil || o.CodeSecurityHostTop99p == nil {
		return nil, false
	}
	return o.CodeSecurityHostTop99p, true
}

// HasCodeSecurityHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCodeSecurityHostTop99p() bool {
	return o != nil && o.CodeSecurityHostTop99p != nil
}

// SetCodeSecurityHostTop99p gets a reference to the given int64 and assigns it to the CodeSecurityHostTop99p field.
func (o *UsageSummaryDate) SetCodeSecurityHostTop99p(v int64) {
	o.CodeSecurityHostTop99p = &v
}

// GetContainerAvg returns the ContainerAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetContainerAvg() int64 {
	if o == nil || o.ContainerAvg == nil {
		var ret int64
		return ret
	}
	return *o.ContainerAvg
}

// GetContainerAvgOk returns a tuple with the ContainerAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetContainerAvgOk() (*int64, bool) {
	if o == nil || o.ContainerAvg == nil {
		return nil, false
	}
	return o.ContainerAvg, true
}

// HasContainerAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasContainerAvg() bool {
	return o != nil && o.ContainerAvg != nil
}

// SetContainerAvg gets a reference to the given int64 and assigns it to the ContainerAvg field.
func (o *UsageSummaryDate) SetContainerAvg(v int64) {
	o.ContainerAvg = &v
}

// GetContainerExclAgentAvg returns the ContainerExclAgentAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetContainerExclAgentAvg() int64 {
	if o == nil || o.ContainerExclAgentAvg == nil {
		var ret int64
		return ret
	}
	return *o.ContainerExclAgentAvg
}

// GetContainerExclAgentAvgOk returns a tuple with the ContainerExclAgentAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetContainerExclAgentAvgOk() (*int64, bool) {
	if o == nil || o.ContainerExclAgentAvg == nil {
		return nil, false
	}
	return o.ContainerExclAgentAvg, true
}

// HasContainerExclAgentAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasContainerExclAgentAvg() bool {
	return o != nil && o.ContainerExclAgentAvg != nil
}

// SetContainerExclAgentAvg gets a reference to the given int64 and assigns it to the ContainerExclAgentAvg field.
func (o *UsageSummaryDate) SetContainerExclAgentAvg(v int64) {
	o.ContainerExclAgentAvg = &v
}

// GetContainerHwm returns the ContainerHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetContainerHwm() int64 {
	if o == nil || o.ContainerHwm == nil {
		var ret int64
		return ret
	}
	return *o.ContainerHwm
}

// GetContainerHwmOk returns a tuple with the ContainerHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetContainerHwmOk() (*int64, bool) {
	if o == nil || o.ContainerHwm == nil {
		return nil, false
	}
	return o.ContainerHwm, true
}

// HasContainerHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasContainerHwm() bool {
	return o != nil && o.ContainerHwm != nil
}

// SetContainerHwm gets a reference to the given int64 and assigns it to the ContainerHwm field.
func (o *UsageSummaryDate) SetContainerHwm(v int64) {
	o.ContainerHwm = &v
}

// GetCsmContainerEnterpriseComplianceCountSum returns the CsmContainerEnterpriseComplianceCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmContainerEnterpriseComplianceCountSum() int64 {
	if o == nil || o.CsmContainerEnterpriseComplianceCountSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmContainerEnterpriseComplianceCountSum
}

// GetCsmContainerEnterpriseComplianceCountSumOk returns a tuple with the CsmContainerEnterpriseComplianceCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmContainerEnterpriseComplianceCountSumOk() (*int64, bool) {
	if o == nil || o.CsmContainerEnterpriseComplianceCountSum == nil {
		return nil, false
	}
	return o.CsmContainerEnterpriseComplianceCountSum, true
}

// HasCsmContainerEnterpriseComplianceCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmContainerEnterpriseComplianceCountSum() bool {
	return o != nil && o.CsmContainerEnterpriseComplianceCountSum != nil
}

// SetCsmContainerEnterpriseComplianceCountSum gets a reference to the given int64 and assigns it to the CsmContainerEnterpriseComplianceCountSum field.
func (o *UsageSummaryDate) SetCsmContainerEnterpriseComplianceCountSum(v int64) {
	o.CsmContainerEnterpriseComplianceCountSum = &v
}

// GetCsmContainerEnterpriseCwsCountSum returns the CsmContainerEnterpriseCwsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmContainerEnterpriseCwsCountSum() int64 {
	if o == nil || o.CsmContainerEnterpriseCwsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmContainerEnterpriseCwsCountSum
}

// GetCsmContainerEnterpriseCwsCountSumOk returns a tuple with the CsmContainerEnterpriseCwsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmContainerEnterpriseCwsCountSumOk() (*int64, bool) {
	if o == nil || o.CsmContainerEnterpriseCwsCountSum == nil {
		return nil, false
	}
	return o.CsmContainerEnterpriseCwsCountSum, true
}

// HasCsmContainerEnterpriseCwsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmContainerEnterpriseCwsCountSum() bool {
	return o != nil && o.CsmContainerEnterpriseCwsCountSum != nil
}

// SetCsmContainerEnterpriseCwsCountSum gets a reference to the given int64 and assigns it to the CsmContainerEnterpriseCwsCountSum field.
func (o *UsageSummaryDate) SetCsmContainerEnterpriseCwsCountSum(v int64) {
	o.CsmContainerEnterpriseCwsCountSum = &v
}

// GetCsmContainerEnterpriseTotalCountSum returns the CsmContainerEnterpriseTotalCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmContainerEnterpriseTotalCountSum() int64 {
	if o == nil || o.CsmContainerEnterpriseTotalCountSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmContainerEnterpriseTotalCountSum
}

// GetCsmContainerEnterpriseTotalCountSumOk returns a tuple with the CsmContainerEnterpriseTotalCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmContainerEnterpriseTotalCountSumOk() (*int64, bool) {
	if o == nil || o.CsmContainerEnterpriseTotalCountSum == nil {
		return nil, false
	}
	return o.CsmContainerEnterpriseTotalCountSum, true
}

// HasCsmContainerEnterpriseTotalCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmContainerEnterpriseTotalCountSum() bool {
	return o != nil && o.CsmContainerEnterpriseTotalCountSum != nil
}

// SetCsmContainerEnterpriseTotalCountSum gets a reference to the given int64 and assigns it to the CsmContainerEnterpriseTotalCountSum field.
func (o *UsageSummaryDate) SetCsmContainerEnterpriseTotalCountSum(v int64) {
	o.CsmContainerEnterpriseTotalCountSum = &v
}

// GetCsmHostEnterpriseAasHostCountTop99p returns the CsmHostEnterpriseAasHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostEnterpriseAasHostCountTop99p() int64 {
	if o == nil || o.CsmHostEnterpriseAasHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseAasHostCountTop99p
}

// GetCsmHostEnterpriseAasHostCountTop99pOk returns a tuple with the CsmHostEnterpriseAasHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostEnterpriseAasHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseAasHostCountTop99p == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseAasHostCountTop99p, true
}

// HasCsmHostEnterpriseAasHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostEnterpriseAasHostCountTop99p() bool {
	return o != nil && o.CsmHostEnterpriseAasHostCountTop99p != nil
}

// SetCsmHostEnterpriseAasHostCountTop99p gets a reference to the given int64 and assigns it to the CsmHostEnterpriseAasHostCountTop99p field.
func (o *UsageSummaryDate) SetCsmHostEnterpriseAasHostCountTop99p(v int64) {
	o.CsmHostEnterpriseAasHostCountTop99p = &v
}

// GetCsmHostEnterpriseAwsHostCountTop99p returns the CsmHostEnterpriseAwsHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostEnterpriseAwsHostCountTop99p() int64 {
	if o == nil || o.CsmHostEnterpriseAwsHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseAwsHostCountTop99p
}

// GetCsmHostEnterpriseAwsHostCountTop99pOk returns a tuple with the CsmHostEnterpriseAwsHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostEnterpriseAwsHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseAwsHostCountTop99p == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseAwsHostCountTop99p, true
}

// HasCsmHostEnterpriseAwsHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostEnterpriseAwsHostCountTop99p() bool {
	return o != nil && o.CsmHostEnterpriseAwsHostCountTop99p != nil
}

// SetCsmHostEnterpriseAwsHostCountTop99p gets a reference to the given int64 and assigns it to the CsmHostEnterpriseAwsHostCountTop99p field.
func (o *UsageSummaryDate) SetCsmHostEnterpriseAwsHostCountTop99p(v int64) {
	o.CsmHostEnterpriseAwsHostCountTop99p = &v
}

// GetCsmHostEnterpriseAzureHostCountTop99p returns the CsmHostEnterpriseAzureHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostEnterpriseAzureHostCountTop99p() int64 {
	if o == nil || o.CsmHostEnterpriseAzureHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseAzureHostCountTop99p
}

// GetCsmHostEnterpriseAzureHostCountTop99pOk returns a tuple with the CsmHostEnterpriseAzureHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostEnterpriseAzureHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseAzureHostCountTop99p == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseAzureHostCountTop99p, true
}

// HasCsmHostEnterpriseAzureHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostEnterpriseAzureHostCountTop99p() bool {
	return o != nil && o.CsmHostEnterpriseAzureHostCountTop99p != nil
}

// SetCsmHostEnterpriseAzureHostCountTop99p gets a reference to the given int64 and assigns it to the CsmHostEnterpriseAzureHostCountTop99p field.
func (o *UsageSummaryDate) SetCsmHostEnterpriseAzureHostCountTop99p(v int64) {
	o.CsmHostEnterpriseAzureHostCountTop99p = &v
}

// GetCsmHostEnterpriseComplianceHostCountTop99p returns the CsmHostEnterpriseComplianceHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostEnterpriseComplianceHostCountTop99p() int64 {
	if o == nil || o.CsmHostEnterpriseComplianceHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseComplianceHostCountTop99p
}

// GetCsmHostEnterpriseComplianceHostCountTop99pOk returns a tuple with the CsmHostEnterpriseComplianceHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostEnterpriseComplianceHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseComplianceHostCountTop99p == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseComplianceHostCountTop99p, true
}

// HasCsmHostEnterpriseComplianceHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostEnterpriseComplianceHostCountTop99p() bool {
	return o != nil && o.CsmHostEnterpriseComplianceHostCountTop99p != nil
}

// SetCsmHostEnterpriseComplianceHostCountTop99p gets a reference to the given int64 and assigns it to the CsmHostEnterpriseComplianceHostCountTop99p field.
func (o *UsageSummaryDate) SetCsmHostEnterpriseComplianceHostCountTop99p(v int64) {
	o.CsmHostEnterpriseComplianceHostCountTop99p = &v
}

// GetCsmHostEnterpriseCwsHostCountTop99p returns the CsmHostEnterpriseCwsHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostEnterpriseCwsHostCountTop99p() int64 {
	if o == nil || o.CsmHostEnterpriseCwsHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseCwsHostCountTop99p
}

// GetCsmHostEnterpriseCwsHostCountTop99pOk returns a tuple with the CsmHostEnterpriseCwsHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostEnterpriseCwsHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseCwsHostCountTop99p == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseCwsHostCountTop99p, true
}

// HasCsmHostEnterpriseCwsHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostEnterpriseCwsHostCountTop99p() bool {
	return o != nil && o.CsmHostEnterpriseCwsHostCountTop99p != nil
}

// SetCsmHostEnterpriseCwsHostCountTop99p gets a reference to the given int64 and assigns it to the CsmHostEnterpriseCwsHostCountTop99p field.
func (o *UsageSummaryDate) SetCsmHostEnterpriseCwsHostCountTop99p(v int64) {
	o.CsmHostEnterpriseCwsHostCountTop99p = &v
}

// GetCsmHostEnterpriseGcpHostCountTop99p returns the CsmHostEnterpriseGcpHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostEnterpriseGcpHostCountTop99p() int64 {
	if o == nil || o.CsmHostEnterpriseGcpHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseGcpHostCountTop99p
}

// GetCsmHostEnterpriseGcpHostCountTop99pOk returns a tuple with the CsmHostEnterpriseGcpHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostEnterpriseGcpHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseGcpHostCountTop99p == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseGcpHostCountTop99p, true
}

// HasCsmHostEnterpriseGcpHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostEnterpriseGcpHostCountTop99p() bool {
	return o != nil && o.CsmHostEnterpriseGcpHostCountTop99p != nil
}

// SetCsmHostEnterpriseGcpHostCountTop99p gets a reference to the given int64 and assigns it to the CsmHostEnterpriseGcpHostCountTop99p field.
func (o *UsageSummaryDate) SetCsmHostEnterpriseGcpHostCountTop99p(v int64) {
	o.CsmHostEnterpriseGcpHostCountTop99p = &v
}

// GetCsmHostEnterpriseTotalHostCountTop99p returns the CsmHostEnterpriseTotalHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostEnterpriseTotalHostCountTop99p() int64 {
	if o == nil || o.CsmHostEnterpriseTotalHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseTotalHostCountTop99p
}

// GetCsmHostEnterpriseTotalHostCountTop99pOk returns a tuple with the CsmHostEnterpriseTotalHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostEnterpriseTotalHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseTotalHostCountTop99p == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseTotalHostCountTop99p, true
}

// HasCsmHostEnterpriseTotalHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostEnterpriseTotalHostCountTop99p() bool {
	return o != nil && o.CsmHostEnterpriseTotalHostCountTop99p != nil
}

// SetCsmHostEnterpriseTotalHostCountTop99p gets a reference to the given int64 and assigns it to the CsmHostEnterpriseTotalHostCountTop99p field.
func (o *UsageSummaryDate) SetCsmHostEnterpriseTotalHostCountTop99p(v int64) {
	o.CsmHostEnterpriseTotalHostCountTop99p = &v
}

// GetCspmAasHostTop99p returns the CspmAasHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCspmAasHostTop99p() int64 {
	if o == nil || o.CspmAasHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CspmAasHostTop99p
}

// GetCspmAasHostTop99pOk returns a tuple with the CspmAasHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCspmAasHostTop99pOk() (*int64, bool) {
	if o == nil || o.CspmAasHostTop99p == nil {
		return nil, false
	}
	return o.CspmAasHostTop99p, true
}

// HasCspmAasHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmAasHostTop99p() bool {
	return o != nil && o.CspmAasHostTop99p != nil
}

// SetCspmAasHostTop99p gets a reference to the given int64 and assigns it to the CspmAasHostTop99p field.
func (o *UsageSummaryDate) SetCspmAasHostTop99p(v int64) {
	o.CspmAasHostTop99p = &v
}

// GetCspmAwsHostTop99p returns the CspmAwsHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCspmAwsHostTop99p() int64 {
	if o == nil || o.CspmAwsHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CspmAwsHostTop99p
}

// GetCspmAwsHostTop99pOk returns a tuple with the CspmAwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCspmAwsHostTop99pOk() (*int64, bool) {
	if o == nil || o.CspmAwsHostTop99p == nil {
		return nil, false
	}
	return o.CspmAwsHostTop99p, true
}

// HasCspmAwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmAwsHostTop99p() bool {
	return o != nil && o.CspmAwsHostTop99p != nil
}

// SetCspmAwsHostTop99p gets a reference to the given int64 and assigns it to the CspmAwsHostTop99p field.
func (o *UsageSummaryDate) SetCspmAwsHostTop99p(v int64) {
	o.CspmAwsHostTop99p = &v
}

// GetCspmAzureHostTop99p returns the CspmAzureHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCspmAzureHostTop99p() int64 {
	if o == nil || o.CspmAzureHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CspmAzureHostTop99p
}

// GetCspmAzureHostTop99pOk returns a tuple with the CspmAzureHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCspmAzureHostTop99pOk() (*int64, bool) {
	if o == nil || o.CspmAzureHostTop99p == nil {
		return nil, false
	}
	return o.CspmAzureHostTop99p, true
}

// HasCspmAzureHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmAzureHostTop99p() bool {
	return o != nil && o.CspmAzureHostTop99p != nil
}

// SetCspmAzureHostTop99p gets a reference to the given int64 and assigns it to the CspmAzureHostTop99p field.
func (o *UsageSummaryDate) SetCspmAzureHostTop99p(v int64) {
	o.CspmAzureHostTop99p = &v
}

// GetCspmContainerAvg returns the CspmContainerAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCspmContainerAvg() int64 {
	if o == nil || o.CspmContainerAvg == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerAvg
}

// GetCspmContainerAvgOk returns a tuple with the CspmContainerAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCspmContainerAvgOk() (*int64, bool) {
	if o == nil || o.CspmContainerAvg == nil {
		return nil, false
	}
	return o.CspmContainerAvg, true
}

// HasCspmContainerAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmContainerAvg() bool {
	return o != nil && o.CspmContainerAvg != nil
}

// SetCspmContainerAvg gets a reference to the given int64 and assigns it to the CspmContainerAvg field.
func (o *UsageSummaryDate) SetCspmContainerAvg(v int64) {
	o.CspmContainerAvg = &v
}

// GetCspmContainerHwm returns the CspmContainerHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCspmContainerHwm() int64 {
	if o == nil || o.CspmContainerHwm == nil {
		var ret int64
		return ret
	}
	return *o.CspmContainerHwm
}

// GetCspmContainerHwmOk returns a tuple with the CspmContainerHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCspmContainerHwmOk() (*int64, bool) {
	if o == nil || o.CspmContainerHwm == nil {
		return nil, false
	}
	return o.CspmContainerHwm, true
}

// HasCspmContainerHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmContainerHwm() bool {
	return o != nil && o.CspmContainerHwm != nil
}

// SetCspmContainerHwm gets a reference to the given int64 and assigns it to the CspmContainerHwm field.
func (o *UsageSummaryDate) SetCspmContainerHwm(v int64) {
	o.CspmContainerHwm = &v
}

// GetCspmGcpHostTop99p returns the CspmGcpHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCspmGcpHostTop99p() int64 {
	if o == nil || o.CspmGcpHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CspmGcpHostTop99p
}

// GetCspmGcpHostTop99pOk returns a tuple with the CspmGcpHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCspmGcpHostTop99pOk() (*int64, bool) {
	if o == nil || o.CspmGcpHostTop99p == nil {
		return nil, false
	}
	return o.CspmGcpHostTop99p, true
}

// HasCspmGcpHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmGcpHostTop99p() bool {
	return o != nil && o.CspmGcpHostTop99p != nil
}

// SetCspmGcpHostTop99p gets a reference to the given int64 and assigns it to the CspmGcpHostTop99p field.
func (o *UsageSummaryDate) SetCspmGcpHostTop99p(v int64) {
	o.CspmGcpHostTop99p = &v
}

// GetCspmHostTop99p returns the CspmHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCspmHostTop99p() int64 {
	if o == nil || o.CspmHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CspmHostTop99p
}

// GetCspmHostTop99pOk returns a tuple with the CspmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCspmHostTop99pOk() (*int64, bool) {
	if o == nil || o.CspmHostTop99p == nil {
		return nil, false
	}
	return o.CspmHostTop99p, true
}

// HasCspmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmHostTop99p() bool {
	return o != nil && o.CspmHostTop99p != nil
}

// SetCspmHostTop99p gets a reference to the given int64 and assigns it to the CspmHostTop99p field.
func (o *UsageSummaryDate) SetCspmHostTop99p(v int64) {
	o.CspmHostTop99p = &v
}

// GetCustomTsAvg returns the CustomTsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCustomTsAvg() int64 {
	if o == nil || o.CustomTsAvg == nil {
		var ret int64
		return ret
	}
	return *o.CustomTsAvg
}

// GetCustomTsAvgOk returns a tuple with the CustomTsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCustomTsAvgOk() (*int64, bool) {
	if o == nil || o.CustomTsAvg == nil {
		return nil, false
	}
	return o.CustomTsAvg, true
}

// HasCustomTsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCustomTsAvg() bool {
	return o != nil && o.CustomTsAvg != nil
}

// SetCustomTsAvg gets a reference to the given int64 and assigns it to the CustomTsAvg field.
func (o *UsageSummaryDate) SetCustomTsAvg(v int64) {
	o.CustomTsAvg = &v
}

// GetCwsContainerCountAvg returns the CwsContainerCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCwsContainerCountAvg() int64 {
	if o == nil || o.CwsContainerCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.CwsContainerCountAvg
}

// GetCwsContainerCountAvgOk returns a tuple with the CwsContainerCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCwsContainerCountAvgOk() (*int64, bool) {
	if o == nil || o.CwsContainerCountAvg == nil {
		return nil, false
	}
	return o.CwsContainerCountAvg, true
}

// HasCwsContainerCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCwsContainerCountAvg() bool {
	return o != nil && o.CwsContainerCountAvg != nil
}

// SetCwsContainerCountAvg gets a reference to the given int64 and assigns it to the CwsContainerCountAvg field.
func (o *UsageSummaryDate) SetCwsContainerCountAvg(v int64) {
	o.CwsContainerCountAvg = &v
}

// GetCwsFargateTaskAvg returns the CwsFargateTaskAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCwsFargateTaskAvg() int64 {
	if o == nil || o.CwsFargateTaskAvg == nil {
		var ret int64
		return ret
	}
	return *o.CwsFargateTaskAvg
}

// GetCwsFargateTaskAvgOk returns a tuple with the CwsFargateTaskAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCwsFargateTaskAvgOk() (*int64, bool) {
	if o == nil || o.CwsFargateTaskAvg == nil {
		return nil, false
	}
	return o.CwsFargateTaskAvg, true
}

// HasCwsFargateTaskAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCwsFargateTaskAvg() bool {
	return o != nil && o.CwsFargateTaskAvg != nil
}

// SetCwsFargateTaskAvg gets a reference to the given int64 and assigns it to the CwsFargateTaskAvg field.
func (o *UsageSummaryDate) SetCwsFargateTaskAvg(v int64) {
	o.CwsFargateTaskAvg = &v
}

// GetCwsHostTop99p returns the CwsHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCwsHostTop99p() int64 {
	if o == nil || o.CwsHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CwsHostTop99p
}

// GetCwsHostTop99pOk returns a tuple with the CwsHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCwsHostTop99pOk() (*int64, bool) {
	if o == nil || o.CwsHostTop99p == nil {
		return nil, false
	}
	return o.CwsHostTop99p, true
}

// HasCwsHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCwsHostTop99p() bool {
	return o != nil && o.CwsHostTop99p != nil
}

// SetCwsHostTop99p gets a reference to the given int64 and assigns it to the CwsHostTop99p field.
func (o *UsageSummaryDate) SetCwsHostTop99p(v int64) {
	o.CwsHostTop99p = &v
}

// GetDataJobsMonitoringHostHrSum returns the DataJobsMonitoringHostHrSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetDataJobsMonitoringHostHrSum() int64 {
	if o == nil || o.DataJobsMonitoringHostHrSum == nil {
		var ret int64
		return ret
	}
	return *o.DataJobsMonitoringHostHrSum
}

// GetDataJobsMonitoringHostHrSumOk returns a tuple with the DataJobsMonitoringHostHrSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetDataJobsMonitoringHostHrSumOk() (*int64, bool) {
	if o == nil || o.DataJobsMonitoringHostHrSum == nil {
		return nil, false
	}
	return o.DataJobsMonitoringHostHrSum, true
}

// HasDataJobsMonitoringHostHrSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDataJobsMonitoringHostHrSum() bool {
	return o != nil && o.DataJobsMonitoringHostHrSum != nil
}

// SetDataJobsMonitoringHostHrSum gets a reference to the given int64 and assigns it to the DataJobsMonitoringHostHrSum field.
func (o *UsageSummaryDate) SetDataJobsMonitoringHostHrSum(v int64) {
	o.DataJobsMonitoringHostHrSum = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDate() bool {
	return o != nil && o.Date != nil
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *UsageSummaryDate) SetDate(v time.Time) {
	o.Date = &v
}

// GetDbmHostTop99p returns the DbmHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetDbmHostTop99p() int64 {
	if o == nil || o.DbmHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.DbmHostTop99p
}

// GetDbmHostTop99pOk returns a tuple with the DbmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetDbmHostTop99pOk() (*int64, bool) {
	if o == nil || o.DbmHostTop99p == nil {
		return nil, false
	}
	return o.DbmHostTop99p, true
}

// HasDbmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDbmHostTop99p() bool {
	return o != nil && o.DbmHostTop99p != nil
}

// SetDbmHostTop99p gets a reference to the given int64 and assigns it to the DbmHostTop99p field.
func (o *UsageSummaryDate) SetDbmHostTop99p(v int64) {
	o.DbmHostTop99p = &v
}

// GetDbmQueriesCountAvg returns the DbmQueriesCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetDbmQueriesCountAvg() int64 {
	if o == nil || o.DbmQueriesCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.DbmQueriesCountAvg
}

// GetDbmQueriesCountAvgOk returns a tuple with the DbmQueriesCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetDbmQueriesCountAvgOk() (*int64, bool) {
	if o == nil || o.DbmQueriesCountAvg == nil {
		return nil, false
	}
	return o.DbmQueriesCountAvg, true
}

// HasDbmQueriesCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDbmQueriesCountAvg() bool {
	return o != nil && o.DbmQueriesCountAvg != nil
}

// SetDbmQueriesCountAvg gets a reference to the given int64 and assigns it to the DbmQueriesCountAvg field.
func (o *UsageSummaryDate) SetDbmQueriesCountAvg(v int64) {
	o.DbmQueriesCountAvg = &v
}

// GetEphInfraHostAgentSum returns the EphInfraHostAgentSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostAgentSum() int64 {
	if o == nil || o.EphInfraHostAgentSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostAgentSum
}

// GetEphInfraHostAgentSumOk returns a tuple with the EphInfraHostAgentSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostAgentSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostAgentSum == nil {
		return nil, false
	}
	return o.EphInfraHostAgentSum, true
}

// HasEphInfraHostAgentSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostAgentSum() bool {
	return o != nil && o.EphInfraHostAgentSum != nil
}

// SetEphInfraHostAgentSum gets a reference to the given int64 and assigns it to the EphInfraHostAgentSum field.
func (o *UsageSummaryDate) SetEphInfraHostAgentSum(v int64) {
	o.EphInfraHostAgentSum = &v
}

// GetEphInfraHostAlibabaSum returns the EphInfraHostAlibabaSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostAlibabaSum() int64 {
	if o == nil || o.EphInfraHostAlibabaSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostAlibabaSum
}

// GetEphInfraHostAlibabaSumOk returns a tuple with the EphInfraHostAlibabaSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostAlibabaSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostAlibabaSum == nil {
		return nil, false
	}
	return o.EphInfraHostAlibabaSum, true
}

// HasEphInfraHostAlibabaSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostAlibabaSum() bool {
	return o != nil && o.EphInfraHostAlibabaSum != nil
}

// SetEphInfraHostAlibabaSum gets a reference to the given int64 and assigns it to the EphInfraHostAlibabaSum field.
func (o *UsageSummaryDate) SetEphInfraHostAlibabaSum(v int64) {
	o.EphInfraHostAlibabaSum = &v
}

// GetEphInfraHostAwsSum returns the EphInfraHostAwsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostAwsSum() int64 {
	if o == nil || o.EphInfraHostAwsSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostAwsSum
}

// GetEphInfraHostAwsSumOk returns a tuple with the EphInfraHostAwsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostAwsSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostAwsSum == nil {
		return nil, false
	}
	return o.EphInfraHostAwsSum, true
}

// HasEphInfraHostAwsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostAwsSum() bool {
	return o != nil && o.EphInfraHostAwsSum != nil
}

// SetEphInfraHostAwsSum gets a reference to the given int64 and assigns it to the EphInfraHostAwsSum field.
func (o *UsageSummaryDate) SetEphInfraHostAwsSum(v int64) {
	o.EphInfraHostAwsSum = &v
}

// GetEphInfraHostAzureSum returns the EphInfraHostAzureSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostAzureSum() int64 {
	if o == nil || o.EphInfraHostAzureSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostAzureSum
}

// GetEphInfraHostAzureSumOk returns a tuple with the EphInfraHostAzureSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostAzureSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostAzureSum == nil {
		return nil, false
	}
	return o.EphInfraHostAzureSum, true
}

// HasEphInfraHostAzureSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostAzureSum() bool {
	return o != nil && o.EphInfraHostAzureSum != nil
}

// SetEphInfraHostAzureSum gets a reference to the given int64 and assigns it to the EphInfraHostAzureSum field.
func (o *UsageSummaryDate) SetEphInfraHostAzureSum(v int64) {
	o.EphInfraHostAzureSum = &v
}

// GetEphInfraHostEntSum returns the EphInfraHostEntSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostEntSum() int64 {
	if o == nil || o.EphInfraHostEntSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostEntSum
}

// GetEphInfraHostEntSumOk returns a tuple with the EphInfraHostEntSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostEntSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostEntSum == nil {
		return nil, false
	}
	return o.EphInfraHostEntSum, true
}

// HasEphInfraHostEntSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostEntSum() bool {
	return o != nil && o.EphInfraHostEntSum != nil
}

// SetEphInfraHostEntSum gets a reference to the given int64 and assigns it to the EphInfraHostEntSum field.
func (o *UsageSummaryDate) SetEphInfraHostEntSum(v int64) {
	o.EphInfraHostEntSum = &v
}

// GetEphInfraHostGcpSum returns the EphInfraHostGcpSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostGcpSum() int64 {
	if o == nil || o.EphInfraHostGcpSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostGcpSum
}

// GetEphInfraHostGcpSumOk returns a tuple with the EphInfraHostGcpSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostGcpSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostGcpSum == nil {
		return nil, false
	}
	return o.EphInfraHostGcpSum, true
}

// HasEphInfraHostGcpSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostGcpSum() bool {
	return o != nil && o.EphInfraHostGcpSum != nil
}

// SetEphInfraHostGcpSum gets a reference to the given int64 and assigns it to the EphInfraHostGcpSum field.
func (o *UsageSummaryDate) SetEphInfraHostGcpSum(v int64) {
	o.EphInfraHostGcpSum = &v
}

// GetEphInfraHostHerokuSum returns the EphInfraHostHerokuSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostHerokuSum() int64 {
	if o == nil || o.EphInfraHostHerokuSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostHerokuSum
}

// GetEphInfraHostHerokuSumOk returns a tuple with the EphInfraHostHerokuSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostHerokuSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostHerokuSum == nil {
		return nil, false
	}
	return o.EphInfraHostHerokuSum, true
}

// HasEphInfraHostHerokuSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostHerokuSum() bool {
	return o != nil && o.EphInfraHostHerokuSum != nil
}

// SetEphInfraHostHerokuSum gets a reference to the given int64 and assigns it to the EphInfraHostHerokuSum field.
func (o *UsageSummaryDate) SetEphInfraHostHerokuSum(v int64) {
	o.EphInfraHostHerokuSum = &v
}

// GetEphInfraHostOnlyAasSum returns the EphInfraHostOnlyAasSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostOnlyAasSum() int64 {
	if o == nil || o.EphInfraHostOnlyAasSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostOnlyAasSum
}

// GetEphInfraHostOnlyAasSumOk returns a tuple with the EphInfraHostOnlyAasSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostOnlyAasSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostOnlyAasSum == nil {
		return nil, false
	}
	return o.EphInfraHostOnlyAasSum, true
}

// HasEphInfraHostOnlyAasSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostOnlyAasSum() bool {
	return o != nil && o.EphInfraHostOnlyAasSum != nil
}

// SetEphInfraHostOnlyAasSum gets a reference to the given int64 and assigns it to the EphInfraHostOnlyAasSum field.
func (o *UsageSummaryDate) SetEphInfraHostOnlyAasSum(v int64) {
	o.EphInfraHostOnlyAasSum = &v
}

// GetEphInfraHostOnlyVsphereSum returns the EphInfraHostOnlyVsphereSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostOnlyVsphereSum() int64 {
	if o == nil || o.EphInfraHostOnlyVsphereSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostOnlyVsphereSum
}

// GetEphInfraHostOnlyVsphereSumOk returns a tuple with the EphInfraHostOnlyVsphereSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostOnlyVsphereSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostOnlyVsphereSum == nil {
		return nil, false
	}
	return o.EphInfraHostOnlyVsphereSum, true
}

// HasEphInfraHostOnlyVsphereSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostOnlyVsphereSum() bool {
	return o != nil && o.EphInfraHostOnlyVsphereSum != nil
}

// SetEphInfraHostOnlyVsphereSum gets a reference to the given int64 and assigns it to the EphInfraHostOnlyVsphereSum field.
func (o *UsageSummaryDate) SetEphInfraHostOnlyVsphereSum(v int64) {
	o.EphInfraHostOnlyVsphereSum = &v
}

// GetEphInfraHostOpentelemetryApmSum returns the EphInfraHostOpentelemetryApmSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostOpentelemetryApmSum() int64 {
	if o == nil || o.EphInfraHostOpentelemetryApmSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostOpentelemetryApmSum
}

// GetEphInfraHostOpentelemetryApmSumOk returns a tuple with the EphInfraHostOpentelemetryApmSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostOpentelemetryApmSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostOpentelemetryApmSum == nil {
		return nil, false
	}
	return o.EphInfraHostOpentelemetryApmSum, true
}

// HasEphInfraHostOpentelemetryApmSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostOpentelemetryApmSum() bool {
	return o != nil && o.EphInfraHostOpentelemetryApmSum != nil
}

// SetEphInfraHostOpentelemetryApmSum gets a reference to the given int64 and assigns it to the EphInfraHostOpentelemetryApmSum field.
func (o *UsageSummaryDate) SetEphInfraHostOpentelemetryApmSum(v int64) {
	o.EphInfraHostOpentelemetryApmSum = &v
}

// GetEphInfraHostOpentelemetrySum returns the EphInfraHostOpentelemetrySum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostOpentelemetrySum() int64 {
	if o == nil || o.EphInfraHostOpentelemetrySum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostOpentelemetrySum
}

// GetEphInfraHostOpentelemetrySumOk returns a tuple with the EphInfraHostOpentelemetrySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostOpentelemetrySumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostOpentelemetrySum == nil {
		return nil, false
	}
	return o.EphInfraHostOpentelemetrySum, true
}

// HasEphInfraHostOpentelemetrySum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostOpentelemetrySum() bool {
	return o != nil && o.EphInfraHostOpentelemetrySum != nil
}

// SetEphInfraHostOpentelemetrySum gets a reference to the given int64 and assigns it to the EphInfraHostOpentelemetrySum field.
func (o *UsageSummaryDate) SetEphInfraHostOpentelemetrySum(v int64) {
	o.EphInfraHostOpentelemetrySum = &v
}

// GetEphInfraHostProSum returns the EphInfraHostProSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostProSum() int64 {
	if o == nil || o.EphInfraHostProSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostProSum
}

// GetEphInfraHostProSumOk returns a tuple with the EphInfraHostProSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostProSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostProSum == nil {
		return nil, false
	}
	return o.EphInfraHostProSum, true
}

// HasEphInfraHostProSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostProSum() bool {
	return o != nil && o.EphInfraHostProSum != nil
}

// SetEphInfraHostProSum gets a reference to the given int64 and assigns it to the EphInfraHostProSum field.
func (o *UsageSummaryDate) SetEphInfraHostProSum(v int64) {
	o.EphInfraHostProSum = &v
}

// GetEphInfraHostProplusSum returns the EphInfraHostProplusSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostProplusSum() int64 {
	if o == nil || o.EphInfraHostProplusSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostProplusSum
}

// GetEphInfraHostProplusSumOk returns a tuple with the EphInfraHostProplusSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostProplusSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostProplusSum == nil {
		return nil, false
	}
	return o.EphInfraHostProplusSum, true
}

// HasEphInfraHostProplusSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostProplusSum() bool {
	return o != nil && o.EphInfraHostProplusSum != nil
}

// SetEphInfraHostProplusSum gets a reference to the given int64 and assigns it to the EphInfraHostProplusSum field.
func (o *UsageSummaryDate) SetEphInfraHostProplusSum(v int64) {
	o.EphInfraHostProplusSum = &v
}

// GetEphInfraHostProxmoxSum returns the EphInfraHostProxmoxSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostProxmoxSum() int64 {
	if o == nil || o.EphInfraHostProxmoxSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostProxmoxSum
}

// GetEphInfraHostProxmoxSumOk returns a tuple with the EphInfraHostProxmoxSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostProxmoxSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostProxmoxSum == nil {
		return nil, false
	}
	return o.EphInfraHostProxmoxSum, true
}

// HasEphInfraHostProxmoxSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostProxmoxSum() bool {
	return o != nil && o.EphInfraHostProxmoxSum != nil
}

// SetEphInfraHostProxmoxSum gets a reference to the given int64 and assigns it to the EphInfraHostProxmoxSum field.
func (o *UsageSummaryDate) SetEphInfraHostProxmoxSum(v int64) {
	o.EphInfraHostProxmoxSum = &v
}

// GetErrorTrackingApmErrorEventsSum returns the ErrorTrackingApmErrorEventsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetErrorTrackingApmErrorEventsSum() int64 {
	if o == nil || o.ErrorTrackingApmErrorEventsSum == nil {
		var ret int64
		return ret
	}
	return *o.ErrorTrackingApmErrorEventsSum
}

// GetErrorTrackingApmErrorEventsSumOk returns a tuple with the ErrorTrackingApmErrorEventsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetErrorTrackingApmErrorEventsSumOk() (*int64, bool) {
	if o == nil || o.ErrorTrackingApmErrorEventsSum == nil {
		return nil, false
	}
	return o.ErrorTrackingApmErrorEventsSum, true
}

// HasErrorTrackingApmErrorEventsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasErrorTrackingApmErrorEventsSum() bool {
	return o != nil && o.ErrorTrackingApmErrorEventsSum != nil
}

// SetErrorTrackingApmErrorEventsSum gets a reference to the given int64 and assigns it to the ErrorTrackingApmErrorEventsSum field.
func (o *UsageSummaryDate) SetErrorTrackingApmErrorEventsSum(v int64) {
	o.ErrorTrackingApmErrorEventsSum = &v
}

// GetErrorTrackingErrorEventsSum returns the ErrorTrackingErrorEventsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetErrorTrackingErrorEventsSum() int64 {
	if o == nil || o.ErrorTrackingErrorEventsSum == nil {
		var ret int64
		return ret
	}
	return *o.ErrorTrackingErrorEventsSum
}

// GetErrorTrackingErrorEventsSumOk returns a tuple with the ErrorTrackingErrorEventsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetErrorTrackingErrorEventsSumOk() (*int64, bool) {
	if o == nil || o.ErrorTrackingErrorEventsSum == nil {
		return nil, false
	}
	return o.ErrorTrackingErrorEventsSum, true
}

// HasErrorTrackingErrorEventsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasErrorTrackingErrorEventsSum() bool {
	return o != nil && o.ErrorTrackingErrorEventsSum != nil
}

// SetErrorTrackingErrorEventsSum gets a reference to the given int64 and assigns it to the ErrorTrackingErrorEventsSum field.
func (o *UsageSummaryDate) SetErrorTrackingErrorEventsSum(v int64) {
	o.ErrorTrackingErrorEventsSum = &v
}

// GetErrorTrackingEventsSum returns the ErrorTrackingEventsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetErrorTrackingEventsSum() int64 {
	if o == nil || o.ErrorTrackingEventsSum == nil {
		var ret int64
		return ret
	}
	return *o.ErrorTrackingEventsSum
}

// GetErrorTrackingEventsSumOk returns a tuple with the ErrorTrackingEventsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetErrorTrackingEventsSumOk() (*int64, bool) {
	if o == nil || o.ErrorTrackingEventsSum == nil {
		return nil, false
	}
	return o.ErrorTrackingEventsSum, true
}

// HasErrorTrackingEventsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasErrorTrackingEventsSum() bool {
	return o != nil && o.ErrorTrackingEventsSum != nil
}

// SetErrorTrackingEventsSum gets a reference to the given int64 and assigns it to the ErrorTrackingEventsSum field.
func (o *UsageSummaryDate) SetErrorTrackingEventsSum(v int64) {
	o.ErrorTrackingEventsSum = &v
}

// GetErrorTrackingRumErrorEventsSum returns the ErrorTrackingRumErrorEventsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetErrorTrackingRumErrorEventsSum() int64 {
	if o == nil || o.ErrorTrackingRumErrorEventsSum == nil {
		var ret int64
		return ret
	}
	return *o.ErrorTrackingRumErrorEventsSum
}

// GetErrorTrackingRumErrorEventsSumOk returns a tuple with the ErrorTrackingRumErrorEventsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetErrorTrackingRumErrorEventsSumOk() (*int64, bool) {
	if o == nil || o.ErrorTrackingRumErrorEventsSum == nil {
		return nil, false
	}
	return o.ErrorTrackingRumErrorEventsSum, true
}

// HasErrorTrackingRumErrorEventsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasErrorTrackingRumErrorEventsSum() bool {
	return o != nil && o.ErrorTrackingRumErrorEventsSum != nil
}

// SetErrorTrackingRumErrorEventsSum gets a reference to the given int64 and assigns it to the ErrorTrackingRumErrorEventsSum field.
func (o *UsageSummaryDate) SetErrorTrackingRumErrorEventsSum(v int64) {
	o.ErrorTrackingRumErrorEventsSum = &v
}

// GetEventManagementCorrelationCorrelatedEventsSum returns the EventManagementCorrelationCorrelatedEventsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEventManagementCorrelationCorrelatedEventsSum() int64 {
	if o == nil || o.EventManagementCorrelationCorrelatedEventsSum == nil {
		var ret int64
		return ret
	}
	return *o.EventManagementCorrelationCorrelatedEventsSum
}

// GetEventManagementCorrelationCorrelatedEventsSumOk returns a tuple with the EventManagementCorrelationCorrelatedEventsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEventManagementCorrelationCorrelatedEventsSumOk() (*int64, bool) {
	if o == nil || o.EventManagementCorrelationCorrelatedEventsSum == nil {
		return nil, false
	}
	return o.EventManagementCorrelationCorrelatedEventsSum, true
}

// HasEventManagementCorrelationCorrelatedEventsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEventManagementCorrelationCorrelatedEventsSum() bool {
	return o != nil && o.EventManagementCorrelationCorrelatedEventsSum != nil
}

// SetEventManagementCorrelationCorrelatedEventsSum gets a reference to the given int64 and assigns it to the EventManagementCorrelationCorrelatedEventsSum field.
func (o *UsageSummaryDate) SetEventManagementCorrelationCorrelatedEventsSum(v int64) {
	o.EventManagementCorrelationCorrelatedEventsSum = &v
}

// GetEventManagementCorrelationCorrelatedRelatedEventsSum returns the EventManagementCorrelationCorrelatedRelatedEventsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEventManagementCorrelationCorrelatedRelatedEventsSum() int64 {
	if o == nil || o.EventManagementCorrelationCorrelatedRelatedEventsSum == nil {
		var ret int64
		return ret
	}
	return *o.EventManagementCorrelationCorrelatedRelatedEventsSum
}

// GetEventManagementCorrelationCorrelatedRelatedEventsSumOk returns a tuple with the EventManagementCorrelationCorrelatedRelatedEventsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEventManagementCorrelationCorrelatedRelatedEventsSumOk() (*int64, bool) {
	if o == nil || o.EventManagementCorrelationCorrelatedRelatedEventsSum == nil {
		return nil, false
	}
	return o.EventManagementCorrelationCorrelatedRelatedEventsSum, true
}

// HasEventManagementCorrelationCorrelatedRelatedEventsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEventManagementCorrelationCorrelatedRelatedEventsSum() bool {
	return o != nil && o.EventManagementCorrelationCorrelatedRelatedEventsSum != nil
}

// SetEventManagementCorrelationCorrelatedRelatedEventsSum gets a reference to the given int64 and assigns it to the EventManagementCorrelationCorrelatedRelatedEventsSum field.
func (o *UsageSummaryDate) SetEventManagementCorrelationCorrelatedRelatedEventsSum(v int64) {
	o.EventManagementCorrelationCorrelatedRelatedEventsSum = &v
}

// GetEventManagementCorrelationSum returns the EventManagementCorrelationSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEventManagementCorrelationSum() int64 {
	if o == nil || o.EventManagementCorrelationSum == nil {
		var ret int64
		return ret
	}
	return *o.EventManagementCorrelationSum
}

// GetEventManagementCorrelationSumOk returns a tuple with the EventManagementCorrelationSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEventManagementCorrelationSumOk() (*int64, bool) {
	if o == nil || o.EventManagementCorrelationSum == nil {
		return nil, false
	}
	return o.EventManagementCorrelationSum, true
}

// HasEventManagementCorrelationSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEventManagementCorrelationSum() bool {
	return o != nil && o.EventManagementCorrelationSum != nil
}

// SetEventManagementCorrelationSum gets a reference to the given int64 and assigns it to the EventManagementCorrelationSum field.
func (o *UsageSummaryDate) SetEventManagementCorrelationSum(v int64) {
	o.EventManagementCorrelationSum = &v
}

// GetFargateContainerProfilerProfilingFargateAvg returns the FargateContainerProfilerProfilingFargateAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFargateContainerProfilerProfilingFargateAvg() int64 {
	if o == nil || o.FargateContainerProfilerProfilingFargateAvg == nil {
		var ret int64
		return ret
	}
	return *o.FargateContainerProfilerProfilingFargateAvg
}

// GetFargateContainerProfilerProfilingFargateAvgOk returns a tuple with the FargateContainerProfilerProfilingFargateAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFargateContainerProfilerProfilingFargateAvgOk() (*int64, bool) {
	if o == nil || o.FargateContainerProfilerProfilingFargateAvg == nil {
		return nil, false
	}
	return o.FargateContainerProfilerProfilingFargateAvg, true
}

// HasFargateContainerProfilerProfilingFargateAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFargateContainerProfilerProfilingFargateAvg() bool {
	return o != nil && o.FargateContainerProfilerProfilingFargateAvg != nil
}

// SetFargateContainerProfilerProfilingFargateAvg gets a reference to the given int64 and assigns it to the FargateContainerProfilerProfilingFargateAvg field.
func (o *UsageSummaryDate) SetFargateContainerProfilerProfilingFargateAvg(v int64) {
	o.FargateContainerProfilerProfilingFargateAvg = &v
}

// GetFargateContainerProfilerProfilingFargateEksAvg returns the FargateContainerProfilerProfilingFargateEksAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFargateContainerProfilerProfilingFargateEksAvg() int64 {
	if o == nil || o.FargateContainerProfilerProfilingFargateEksAvg == nil {
		var ret int64
		return ret
	}
	return *o.FargateContainerProfilerProfilingFargateEksAvg
}

// GetFargateContainerProfilerProfilingFargateEksAvgOk returns a tuple with the FargateContainerProfilerProfilingFargateEksAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFargateContainerProfilerProfilingFargateEksAvgOk() (*int64, bool) {
	if o == nil || o.FargateContainerProfilerProfilingFargateEksAvg == nil {
		return nil, false
	}
	return o.FargateContainerProfilerProfilingFargateEksAvg, true
}

// HasFargateContainerProfilerProfilingFargateEksAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFargateContainerProfilerProfilingFargateEksAvg() bool {
	return o != nil && o.FargateContainerProfilerProfilingFargateEksAvg != nil
}

// SetFargateContainerProfilerProfilingFargateEksAvg gets a reference to the given int64 and assigns it to the FargateContainerProfilerProfilingFargateEksAvg field.
func (o *UsageSummaryDate) SetFargateContainerProfilerProfilingFargateEksAvg(v int64) {
	o.FargateContainerProfilerProfilingFargateEksAvg = &v
}

// GetFargateTasksCountAvg returns the FargateTasksCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFargateTasksCountAvg() int64 {
	if o == nil || o.FargateTasksCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountAvg
}

// GetFargateTasksCountAvgOk returns a tuple with the FargateTasksCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFargateTasksCountAvgOk() (*int64, bool) {
	if o == nil || o.FargateTasksCountAvg == nil {
		return nil, false
	}
	return o.FargateTasksCountAvg, true
}

// HasFargateTasksCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFargateTasksCountAvg() bool {
	return o != nil && o.FargateTasksCountAvg != nil
}

// SetFargateTasksCountAvg gets a reference to the given int64 and assigns it to the FargateTasksCountAvg field.
func (o *UsageSummaryDate) SetFargateTasksCountAvg(v int64) {
	o.FargateTasksCountAvg = &v
}

// GetFargateTasksCountHwm returns the FargateTasksCountHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFargateTasksCountHwm() int64 {
	if o == nil || o.FargateTasksCountHwm == nil {
		var ret int64
		return ret
	}
	return *o.FargateTasksCountHwm
}

// GetFargateTasksCountHwmOk returns a tuple with the FargateTasksCountHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFargateTasksCountHwmOk() (*int64, bool) {
	if o == nil || o.FargateTasksCountHwm == nil {
		return nil, false
	}
	return o.FargateTasksCountHwm, true
}

// HasFargateTasksCountHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFargateTasksCountHwm() bool {
	return o != nil && o.FargateTasksCountHwm != nil
}

// SetFargateTasksCountHwm gets a reference to the given int64 and assigns it to the FargateTasksCountHwm field.
func (o *UsageSummaryDate) SetFargateTasksCountHwm(v int64) {
	o.FargateTasksCountHwm = &v
}

// GetFlexLogsComputeLargeAvg returns the FlexLogsComputeLargeAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFlexLogsComputeLargeAvg() int64 {
	if o == nil || o.FlexLogsComputeLargeAvg == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeLargeAvg
}

// GetFlexLogsComputeLargeAvgOk returns a tuple with the FlexLogsComputeLargeAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFlexLogsComputeLargeAvgOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeLargeAvg == nil {
		return nil, false
	}
	return o.FlexLogsComputeLargeAvg, true
}

// HasFlexLogsComputeLargeAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFlexLogsComputeLargeAvg() bool {
	return o != nil && o.FlexLogsComputeLargeAvg != nil
}

// SetFlexLogsComputeLargeAvg gets a reference to the given int64 and assigns it to the FlexLogsComputeLargeAvg field.
func (o *UsageSummaryDate) SetFlexLogsComputeLargeAvg(v int64) {
	o.FlexLogsComputeLargeAvg = &v
}

// GetFlexLogsComputeMediumAvg returns the FlexLogsComputeMediumAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFlexLogsComputeMediumAvg() int64 {
	if o == nil || o.FlexLogsComputeMediumAvg == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeMediumAvg
}

// GetFlexLogsComputeMediumAvgOk returns a tuple with the FlexLogsComputeMediumAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFlexLogsComputeMediumAvgOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeMediumAvg == nil {
		return nil, false
	}
	return o.FlexLogsComputeMediumAvg, true
}

// HasFlexLogsComputeMediumAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFlexLogsComputeMediumAvg() bool {
	return o != nil && o.FlexLogsComputeMediumAvg != nil
}

// SetFlexLogsComputeMediumAvg gets a reference to the given int64 and assigns it to the FlexLogsComputeMediumAvg field.
func (o *UsageSummaryDate) SetFlexLogsComputeMediumAvg(v int64) {
	o.FlexLogsComputeMediumAvg = &v
}

// GetFlexLogsComputeSmallAvg returns the FlexLogsComputeSmallAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFlexLogsComputeSmallAvg() int64 {
	if o == nil || o.FlexLogsComputeSmallAvg == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeSmallAvg
}

// GetFlexLogsComputeSmallAvgOk returns a tuple with the FlexLogsComputeSmallAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFlexLogsComputeSmallAvgOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeSmallAvg == nil {
		return nil, false
	}
	return o.FlexLogsComputeSmallAvg, true
}

// HasFlexLogsComputeSmallAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFlexLogsComputeSmallAvg() bool {
	return o != nil && o.FlexLogsComputeSmallAvg != nil
}

// SetFlexLogsComputeSmallAvg gets a reference to the given int64 and assigns it to the FlexLogsComputeSmallAvg field.
func (o *UsageSummaryDate) SetFlexLogsComputeSmallAvg(v int64) {
	o.FlexLogsComputeSmallAvg = &v
}

// GetFlexLogsComputeXlargeAvg returns the FlexLogsComputeXlargeAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFlexLogsComputeXlargeAvg() int64 {
	if o == nil || o.FlexLogsComputeXlargeAvg == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeXlargeAvg
}

// GetFlexLogsComputeXlargeAvgOk returns a tuple with the FlexLogsComputeXlargeAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFlexLogsComputeXlargeAvgOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeXlargeAvg == nil {
		return nil, false
	}
	return o.FlexLogsComputeXlargeAvg, true
}

// HasFlexLogsComputeXlargeAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFlexLogsComputeXlargeAvg() bool {
	return o != nil && o.FlexLogsComputeXlargeAvg != nil
}

// SetFlexLogsComputeXlargeAvg gets a reference to the given int64 and assigns it to the FlexLogsComputeXlargeAvg field.
func (o *UsageSummaryDate) SetFlexLogsComputeXlargeAvg(v int64) {
	o.FlexLogsComputeXlargeAvg = &v
}

// GetFlexLogsComputeXsmallAvg returns the FlexLogsComputeXsmallAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFlexLogsComputeXsmallAvg() int64 {
	if o == nil || o.FlexLogsComputeXsmallAvg == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsComputeXsmallAvg
}

// GetFlexLogsComputeXsmallAvgOk returns a tuple with the FlexLogsComputeXsmallAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFlexLogsComputeXsmallAvgOk() (*int64, bool) {
	if o == nil || o.FlexLogsComputeXsmallAvg == nil {
		return nil, false
	}
	return o.FlexLogsComputeXsmallAvg, true
}

// HasFlexLogsComputeXsmallAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFlexLogsComputeXsmallAvg() bool {
	return o != nil && o.FlexLogsComputeXsmallAvg != nil
}

// SetFlexLogsComputeXsmallAvg gets a reference to the given int64 and assigns it to the FlexLogsComputeXsmallAvg field.
func (o *UsageSummaryDate) SetFlexLogsComputeXsmallAvg(v int64) {
	o.FlexLogsComputeXsmallAvg = &v
}

// GetFlexLogsStarterAvg returns the FlexLogsStarterAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFlexLogsStarterAvg() int64 {
	if o == nil || o.FlexLogsStarterAvg == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsStarterAvg
}

// GetFlexLogsStarterAvgOk returns a tuple with the FlexLogsStarterAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFlexLogsStarterAvgOk() (*int64, bool) {
	if o == nil || o.FlexLogsStarterAvg == nil {
		return nil, false
	}
	return o.FlexLogsStarterAvg, true
}

// HasFlexLogsStarterAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFlexLogsStarterAvg() bool {
	return o != nil && o.FlexLogsStarterAvg != nil
}

// SetFlexLogsStarterAvg gets a reference to the given int64 and assigns it to the FlexLogsStarterAvg field.
func (o *UsageSummaryDate) SetFlexLogsStarterAvg(v int64) {
	o.FlexLogsStarterAvg = &v
}

// GetFlexLogsStarterStorageIndexAvg returns the FlexLogsStarterStorageIndexAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFlexLogsStarterStorageIndexAvg() int64 {
	if o == nil || o.FlexLogsStarterStorageIndexAvg == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsStarterStorageIndexAvg
}

// GetFlexLogsStarterStorageIndexAvgOk returns a tuple with the FlexLogsStarterStorageIndexAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFlexLogsStarterStorageIndexAvgOk() (*int64, bool) {
	if o == nil || o.FlexLogsStarterStorageIndexAvg == nil {
		return nil, false
	}
	return o.FlexLogsStarterStorageIndexAvg, true
}

// HasFlexLogsStarterStorageIndexAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFlexLogsStarterStorageIndexAvg() bool {
	return o != nil && o.FlexLogsStarterStorageIndexAvg != nil
}

// SetFlexLogsStarterStorageIndexAvg gets a reference to the given int64 and assigns it to the FlexLogsStarterStorageIndexAvg field.
func (o *UsageSummaryDate) SetFlexLogsStarterStorageIndexAvg(v int64) {
	o.FlexLogsStarterStorageIndexAvg = &v
}

// GetFlexLogsStarterStorageRetentionAdjustmentAvg returns the FlexLogsStarterStorageRetentionAdjustmentAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFlexLogsStarterStorageRetentionAdjustmentAvg() int64 {
	if o == nil || o.FlexLogsStarterStorageRetentionAdjustmentAvg == nil {
		var ret int64
		return ret
	}
	return *o.FlexLogsStarterStorageRetentionAdjustmentAvg
}

// GetFlexLogsStarterStorageRetentionAdjustmentAvgOk returns a tuple with the FlexLogsStarterStorageRetentionAdjustmentAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFlexLogsStarterStorageRetentionAdjustmentAvgOk() (*int64, bool) {
	if o == nil || o.FlexLogsStarterStorageRetentionAdjustmentAvg == nil {
		return nil, false
	}
	return o.FlexLogsStarterStorageRetentionAdjustmentAvg, true
}

// HasFlexLogsStarterStorageRetentionAdjustmentAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFlexLogsStarterStorageRetentionAdjustmentAvg() bool {
	return o != nil && o.FlexLogsStarterStorageRetentionAdjustmentAvg != nil
}

// SetFlexLogsStarterStorageRetentionAdjustmentAvg gets a reference to the given int64 and assigns it to the FlexLogsStarterStorageRetentionAdjustmentAvg field.
func (o *UsageSummaryDate) SetFlexLogsStarterStorageRetentionAdjustmentAvg(v int64) {
	o.FlexLogsStarterStorageRetentionAdjustmentAvg = &v
}

// GetFlexStoredLogsAvg returns the FlexStoredLogsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFlexStoredLogsAvg() int64 {
	if o == nil || o.FlexStoredLogsAvg == nil {
		var ret int64
		return ret
	}
	return *o.FlexStoredLogsAvg
}

// GetFlexStoredLogsAvgOk returns a tuple with the FlexStoredLogsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFlexStoredLogsAvgOk() (*int64, bool) {
	if o == nil || o.FlexStoredLogsAvg == nil {
		return nil, false
	}
	return o.FlexStoredLogsAvg, true
}

// HasFlexStoredLogsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFlexStoredLogsAvg() bool {
	return o != nil && o.FlexStoredLogsAvg != nil
}

// SetFlexStoredLogsAvg gets a reference to the given int64 and assigns it to the FlexStoredLogsAvg field.
func (o *UsageSummaryDate) SetFlexStoredLogsAvg(v int64) {
	o.FlexStoredLogsAvg = &v
}

// GetForwardingEventsBytesSum returns the ForwardingEventsBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetForwardingEventsBytesSum() int64 {
	if o == nil || o.ForwardingEventsBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.ForwardingEventsBytesSum
}

// GetForwardingEventsBytesSumOk returns a tuple with the ForwardingEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetForwardingEventsBytesSumOk() (*int64, bool) {
	if o == nil || o.ForwardingEventsBytesSum == nil {
		return nil, false
	}
	return o.ForwardingEventsBytesSum, true
}

// HasForwardingEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasForwardingEventsBytesSum() bool {
	return o != nil && o.ForwardingEventsBytesSum != nil
}

// SetForwardingEventsBytesSum gets a reference to the given int64 and assigns it to the ForwardingEventsBytesSum field.
func (o *UsageSummaryDate) SetForwardingEventsBytesSum(v int64) {
	o.ForwardingEventsBytesSum = &v
}

// GetGcpHostTop99p returns the GcpHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetGcpHostTop99p() int64 {
	if o == nil || o.GcpHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.GcpHostTop99p
}

// GetGcpHostTop99pOk returns a tuple with the GcpHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetGcpHostTop99pOk() (*int64, bool) {
	if o == nil || o.GcpHostTop99p == nil {
		return nil, false
	}
	return o.GcpHostTop99p, true
}

// HasGcpHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasGcpHostTop99p() bool {
	return o != nil && o.GcpHostTop99p != nil
}

// SetGcpHostTop99p gets a reference to the given int64 and assigns it to the GcpHostTop99p field.
func (o *UsageSummaryDate) SetGcpHostTop99p(v int64) {
	o.GcpHostTop99p = &v
}

// GetHerokuHostTop99p returns the HerokuHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetHerokuHostTop99p() int64 {
	if o == nil || o.HerokuHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.HerokuHostTop99p
}

// GetHerokuHostTop99pOk returns a tuple with the HerokuHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetHerokuHostTop99pOk() (*int64, bool) {
	if o == nil || o.HerokuHostTop99p == nil {
		return nil, false
	}
	return o.HerokuHostTop99p, true
}

// HasHerokuHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasHerokuHostTop99p() bool {
	return o != nil && o.HerokuHostTop99p != nil
}

// SetHerokuHostTop99p gets a reference to the given int64 and assigns it to the HerokuHostTop99p field.
func (o *UsageSummaryDate) SetHerokuHostTop99p(v int64) {
	o.HerokuHostTop99p = &v
}

// GetIncidentManagementMonthlyActiveUsersHwm returns the IncidentManagementMonthlyActiveUsersHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIncidentManagementMonthlyActiveUsersHwm() int64 {
	if o == nil || o.IncidentManagementMonthlyActiveUsersHwm == nil {
		var ret int64
		return ret
	}
	return *o.IncidentManagementMonthlyActiveUsersHwm
}

// GetIncidentManagementMonthlyActiveUsersHwmOk returns a tuple with the IncidentManagementMonthlyActiveUsersHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIncidentManagementMonthlyActiveUsersHwmOk() (*int64, bool) {
	if o == nil || o.IncidentManagementMonthlyActiveUsersHwm == nil {
		return nil, false
	}
	return o.IncidentManagementMonthlyActiveUsersHwm, true
}

// HasIncidentManagementMonthlyActiveUsersHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIncidentManagementMonthlyActiveUsersHwm() bool {
	return o != nil && o.IncidentManagementMonthlyActiveUsersHwm != nil
}

// SetIncidentManagementMonthlyActiveUsersHwm gets a reference to the given int64 and assigns it to the IncidentManagementMonthlyActiveUsersHwm field.
func (o *UsageSummaryDate) SetIncidentManagementMonthlyActiveUsersHwm(v int64) {
	o.IncidentManagementMonthlyActiveUsersHwm = &v
}

// GetIncidentManagementSeatsHwm returns the IncidentManagementSeatsHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIncidentManagementSeatsHwm() int64 {
	if o == nil || o.IncidentManagementSeatsHwm == nil {
		var ret int64
		return ret
	}
	return *o.IncidentManagementSeatsHwm
}

// GetIncidentManagementSeatsHwmOk returns a tuple with the IncidentManagementSeatsHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIncidentManagementSeatsHwmOk() (*int64, bool) {
	if o == nil || o.IncidentManagementSeatsHwm == nil {
		return nil, false
	}
	return o.IncidentManagementSeatsHwm, true
}

// HasIncidentManagementSeatsHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIncidentManagementSeatsHwm() bool {
	return o != nil && o.IncidentManagementSeatsHwm != nil
}

// SetIncidentManagementSeatsHwm gets a reference to the given int64 and assigns it to the IncidentManagementSeatsHwm field.
func (o *UsageSummaryDate) SetIncidentManagementSeatsHwm(v int64) {
	o.IncidentManagementSeatsHwm = &v
}

// GetIndexedEventsCountSum returns the IndexedEventsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIndexedEventsCountSum() int64 {
	if o == nil || o.IndexedEventsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCountSum
}

// GetIndexedEventsCountSumOk returns a tuple with the IndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil || o.IndexedEventsCountSum == nil {
		return nil, false
	}
	return o.IndexedEventsCountSum, true
}

// HasIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIndexedEventsCountSum() bool {
	return o != nil && o.IndexedEventsCountSum != nil
}

// SetIndexedEventsCountSum gets a reference to the given int64 and assigns it to the IndexedEventsCountSum field.
func (o *UsageSummaryDate) SetIndexedEventsCountSum(v int64) {
	o.IndexedEventsCountSum = &v
}

// GetInfraHostTop99p returns the InfraHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraHostTop99p() int64 {
	if o == nil || o.InfraHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostTop99p
}

// GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraHostTop99pOk() (*int64, bool) {
	if o == nil || o.InfraHostTop99p == nil {
		return nil, false
	}
	return o.InfraHostTop99p, true
}

// HasInfraHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraHostTop99p() bool {
	return o != nil && o.InfraHostTop99p != nil
}

// SetInfraHostTop99p gets a reference to the given int64 and assigns it to the InfraHostTop99p field.
func (o *UsageSummaryDate) SetInfraHostTop99p(v int64) {
	o.InfraHostTop99p = &v
}

// GetIngestedEventsBytesSum returns the IngestedEventsBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIngestedEventsBytesSum() int64 {
	if o == nil || o.IngestedEventsBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytesSum
}

// GetIngestedEventsBytesSumOk returns a tuple with the IngestedEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIngestedEventsBytesSumOk() (*int64, bool) {
	if o == nil || o.IngestedEventsBytesSum == nil {
		return nil, false
	}
	return o.IngestedEventsBytesSum, true
}

// HasIngestedEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIngestedEventsBytesSum() bool {
	return o != nil && o.IngestedEventsBytesSum != nil
}

// SetIngestedEventsBytesSum gets a reference to the given int64 and assigns it to the IngestedEventsBytesSum field.
func (o *UsageSummaryDate) SetIngestedEventsBytesSum(v int64) {
	o.IngestedEventsBytesSum = &v
}

// GetIotDeviceSum returns the IotDeviceSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIotDeviceSum() int64 {
	if o == nil || o.IotDeviceSum == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceSum
}

// GetIotDeviceSumOk returns a tuple with the IotDeviceSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIotDeviceSumOk() (*int64, bool) {
	if o == nil || o.IotDeviceSum == nil {
		return nil, false
	}
	return o.IotDeviceSum, true
}

// HasIotDeviceSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIotDeviceSum() bool {
	return o != nil && o.IotDeviceSum != nil
}

// SetIotDeviceSum gets a reference to the given int64 and assigns it to the IotDeviceSum field.
func (o *UsageSummaryDate) SetIotDeviceSum(v int64) {
	o.IotDeviceSum = &v
}

// GetIotDeviceTop99p returns the IotDeviceTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIotDeviceTop99p() int64 {
	if o == nil || o.IotDeviceTop99p == nil {
		var ret int64
		return ret
	}
	return *o.IotDeviceTop99p
}

// GetIotDeviceTop99pOk returns a tuple with the IotDeviceTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIotDeviceTop99pOk() (*int64, bool) {
	if o == nil || o.IotDeviceTop99p == nil {
		return nil, false
	}
	return o.IotDeviceTop99p, true
}

// HasIotDeviceTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIotDeviceTop99p() bool {
	return o != nil && o.IotDeviceTop99p != nil
}

// SetIotDeviceTop99p gets a reference to the given int64 and assigns it to the IotDeviceTop99p field.
func (o *UsageSummaryDate) SetIotDeviceTop99p(v int64) {
	o.IotDeviceTop99p = &v
}

// GetLlmObservabilityMinSpendSum returns the LlmObservabilityMinSpendSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetLlmObservabilityMinSpendSum() int64 {
	if o == nil || o.LlmObservabilityMinSpendSum == nil {
		var ret int64
		return ret
	}
	return *o.LlmObservabilityMinSpendSum
}

// GetLlmObservabilityMinSpendSumOk returns a tuple with the LlmObservabilityMinSpendSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetLlmObservabilityMinSpendSumOk() (*int64, bool) {
	if o == nil || o.LlmObservabilityMinSpendSum == nil {
		return nil, false
	}
	return o.LlmObservabilityMinSpendSum, true
}

// HasLlmObservabilityMinSpendSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasLlmObservabilityMinSpendSum() bool {
	return o != nil && o.LlmObservabilityMinSpendSum != nil
}

// SetLlmObservabilityMinSpendSum gets a reference to the given int64 and assigns it to the LlmObservabilityMinSpendSum field.
func (o *UsageSummaryDate) SetLlmObservabilityMinSpendSum(v int64) {
	o.LlmObservabilityMinSpendSum = &v
}

// GetLlmObservabilitySum returns the LlmObservabilitySum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetLlmObservabilitySum() int64 {
	if o == nil || o.LlmObservabilitySum == nil {
		var ret int64
		return ret
	}
	return *o.LlmObservabilitySum
}

// GetLlmObservabilitySumOk returns a tuple with the LlmObservabilitySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetLlmObservabilitySumOk() (*int64, bool) {
	if o == nil || o.LlmObservabilitySum == nil {
		return nil, false
	}
	return o.LlmObservabilitySum, true
}

// HasLlmObservabilitySum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasLlmObservabilitySum() bool {
	return o != nil && o.LlmObservabilitySum != nil
}

// SetLlmObservabilitySum gets a reference to the given int64 and assigns it to the LlmObservabilitySum field.
func (o *UsageSummaryDate) SetLlmObservabilitySum(v int64) {
	o.LlmObservabilitySum = &v
}

// GetMobileRumLiteSessionCountSum returns the MobileRumLiteSessionCountSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumLiteSessionCountSum() int64 {
	if o == nil || o.MobileRumLiteSessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumLiteSessionCountSum
}

// GetMobileRumLiteSessionCountSumOk returns a tuple with the MobileRumLiteSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumLiteSessionCountSumOk() (*int64, bool) {
	if o == nil || o.MobileRumLiteSessionCountSum == nil {
		return nil, false
	}
	return o.MobileRumLiteSessionCountSum, true
}

// HasMobileRumLiteSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumLiteSessionCountSum() bool {
	return o != nil && o.MobileRumLiteSessionCountSum != nil
}

// SetMobileRumLiteSessionCountSum gets a reference to the given int64 and assigns it to the MobileRumLiteSessionCountSum field.
// Deprecated
func (o *UsageSummaryDate) SetMobileRumLiteSessionCountSum(v int64) {
	o.MobileRumLiteSessionCountSum = &v
}

// GetMobileRumSessionCountAndroidSum returns the MobileRumSessionCountAndroidSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountAndroidSum() int64 {
	if o == nil || o.MobileRumSessionCountAndroidSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountAndroidSum
}

// GetMobileRumSessionCountAndroidSumOk returns a tuple with the MobileRumSessionCountAndroidSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountAndroidSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountAndroidSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountAndroidSum, true
}

// HasMobileRumSessionCountAndroidSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountAndroidSum() bool {
	return o != nil && o.MobileRumSessionCountAndroidSum != nil
}

// SetMobileRumSessionCountAndroidSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountAndroidSum field.
// Deprecated
func (o *UsageSummaryDate) SetMobileRumSessionCountAndroidSum(v int64) {
	o.MobileRumSessionCountAndroidSum = &v
}

// GetMobileRumSessionCountFlutterSum returns the MobileRumSessionCountFlutterSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountFlutterSum() int64 {
	if o == nil || o.MobileRumSessionCountFlutterSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountFlutterSum
}

// GetMobileRumSessionCountFlutterSumOk returns a tuple with the MobileRumSessionCountFlutterSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountFlutterSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountFlutterSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountFlutterSum, true
}

// HasMobileRumSessionCountFlutterSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountFlutterSum() bool {
	return o != nil && o.MobileRumSessionCountFlutterSum != nil
}

// SetMobileRumSessionCountFlutterSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountFlutterSum field.
// Deprecated
func (o *UsageSummaryDate) SetMobileRumSessionCountFlutterSum(v int64) {
	o.MobileRumSessionCountFlutterSum = &v
}

// GetMobileRumSessionCountIosSum returns the MobileRumSessionCountIosSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountIosSum() int64 {
	if o == nil || o.MobileRumSessionCountIosSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountIosSum
}

// GetMobileRumSessionCountIosSumOk returns a tuple with the MobileRumSessionCountIosSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountIosSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountIosSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountIosSum, true
}

// HasMobileRumSessionCountIosSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountIosSum() bool {
	return o != nil && o.MobileRumSessionCountIosSum != nil
}

// SetMobileRumSessionCountIosSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountIosSum field.
// Deprecated
func (o *UsageSummaryDate) SetMobileRumSessionCountIosSum(v int64) {
	o.MobileRumSessionCountIosSum = &v
}

// GetMobileRumSessionCountReactnativeSum returns the MobileRumSessionCountReactnativeSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountReactnativeSum() int64 {
	if o == nil || o.MobileRumSessionCountReactnativeSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountReactnativeSum
}

// GetMobileRumSessionCountReactnativeSumOk returns a tuple with the MobileRumSessionCountReactnativeSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountReactnativeSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountReactnativeSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountReactnativeSum, true
}

// HasMobileRumSessionCountReactnativeSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountReactnativeSum() bool {
	return o != nil && o.MobileRumSessionCountReactnativeSum != nil
}

// SetMobileRumSessionCountReactnativeSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountReactnativeSum field.
// Deprecated
func (o *UsageSummaryDate) SetMobileRumSessionCountReactnativeSum(v int64) {
	o.MobileRumSessionCountReactnativeSum = &v
}

// GetMobileRumSessionCountRokuSum returns the MobileRumSessionCountRokuSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountRokuSum() int64 {
	if o == nil || o.MobileRumSessionCountRokuSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountRokuSum
}

// GetMobileRumSessionCountRokuSumOk returns a tuple with the MobileRumSessionCountRokuSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountRokuSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountRokuSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountRokuSum, true
}

// HasMobileRumSessionCountRokuSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountRokuSum() bool {
	return o != nil && o.MobileRumSessionCountRokuSum != nil
}

// SetMobileRumSessionCountRokuSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountRokuSum field.
// Deprecated
func (o *UsageSummaryDate) SetMobileRumSessionCountRokuSum(v int64) {
	o.MobileRumSessionCountRokuSum = &v
}

// GetMobileRumSessionCountSum returns the MobileRumSessionCountSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountSum() int64 {
	if o == nil || o.MobileRumSessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumSessionCountSum
}

// GetMobileRumSessionCountSumOk returns a tuple with the MobileRumSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumSessionCountSumOk() (*int64, bool) {
	if o == nil || o.MobileRumSessionCountSum == nil {
		return nil, false
	}
	return o.MobileRumSessionCountSum, true
}

// HasMobileRumSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumSessionCountSum() bool {
	return o != nil && o.MobileRumSessionCountSum != nil
}

// SetMobileRumSessionCountSum gets a reference to the given int64 and assigns it to the MobileRumSessionCountSum field.
// Deprecated
func (o *UsageSummaryDate) SetMobileRumSessionCountSum(v int64) {
	o.MobileRumSessionCountSum = &v
}

// GetMobileRumUnitsSum returns the MobileRumUnitsSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumUnitsSum() int64 {
	if o == nil || o.MobileRumUnitsSum == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumUnitsSum
}

// GetMobileRumUnitsSumOk returns a tuple with the MobileRumUnitsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetMobileRumUnitsSumOk() (*int64, bool) {
	if o == nil || o.MobileRumUnitsSum == nil {
		return nil, false
	}
	return o.MobileRumUnitsSum, true
}

// HasMobileRumUnitsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMobileRumUnitsSum() bool {
	return o != nil && o.MobileRumUnitsSum != nil
}

// SetMobileRumUnitsSum gets a reference to the given int64 and assigns it to the MobileRumUnitsSum field.
// Deprecated
func (o *UsageSummaryDate) SetMobileRumUnitsSum(v int64) {
	o.MobileRumUnitsSum = &v
}

// GetNdmNetflowEventsSum returns the NdmNetflowEventsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetNdmNetflowEventsSum() int64 {
	if o == nil || o.NdmNetflowEventsSum == nil {
		var ret int64
		return ret
	}
	return *o.NdmNetflowEventsSum
}

// GetNdmNetflowEventsSumOk returns a tuple with the NdmNetflowEventsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetNdmNetflowEventsSumOk() (*int64, bool) {
	if o == nil || o.NdmNetflowEventsSum == nil {
		return nil, false
	}
	return o.NdmNetflowEventsSum, true
}

// HasNdmNetflowEventsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasNdmNetflowEventsSum() bool {
	return o != nil && o.NdmNetflowEventsSum != nil
}

// SetNdmNetflowEventsSum gets a reference to the given int64 and assigns it to the NdmNetflowEventsSum field.
func (o *UsageSummaryDate) SetNdmNetflowEventsSum(v int64) {
	o.NdmNetflowEventsSum = &v
}

// GetNetflowIndexedEventsCountSum returns the NetflowIndexedEventsCountSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSum() int64 {
	if o == nil || o.NetflowIndexedEventsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.NetflowIndexedEventsCountSum
}

// GetNetflowIndexedEventsCountSumOk returns a tuple with the NetflowIndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetNetflowIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil || o.NetflowIndexedEventsCountSum == nil {
		return nil, false
	}
	return o.NetflowIndexedEventsCountSum, true
}

// HasNetflowIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasNetflowIndexedEventsCountSum() bool {
	return o != nil && o.NetflowIndexedEventsCountSum != nil
}

// SetNetflowIndexedEventsCountSum gets a reference to the given int64 and assigns it to the NetflowIndexedEventsCountSum field.
// Deprecated
func (o *UsageSummaryDate) SetNetflowIndexedEventsCountSum(v int64) {
	o.NetflowIndexedEventsCountSum = &v
}

// GetNetworkDeviceWirelessTop99p returns the NetworkDeviceWirelessTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetNetworkDeviceWirelessTop99p() int64 {
	if o == nil || o.NetworkDeviceWirelessTop99p == nil {
		var ret int64
		return ret
	}
	return *o.NetworkDeviceWirelessTop99p
}

// GetNetworkDeviceWirelessTop99pOk returns a tuple with the NetworkDeviceWirelessTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetNetworkDeviceWirelessTop99pOk() (*int64, bool) {
	if o == nil || o.NetworkDeviceWirelessTop99p == nil {
		return nil, false
	}
	return o.NetworkDeviceWirelessTop99p, true
}

// HasNetworkDeviceWirelessTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasNetworkDeviceWirelessTop99p() bool {
	return o != nil && o.NetworkDeviceWirelessTop99p != nil
}

// SetNetworkDeviceWirelessTop99p gets a reference to the given int64 and assigns it to the NetworkDeviceWirelessTop99p field.
func (o *UsageSummaryDate) SetNetworkDeviceWirelessTop99p(v int64) {
	o.NetworkDeviceWirelessTop99p = &v
}

// GetNpmHostTop99p returns the NpmHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetNpmHostTop99p() int64 {
	if o == nil || o.NpmHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.NpmHostTop99p
}

// GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetNpmHostTop99pOk() (*int64, bool) {
	if o == nil || o.NpmHostTop99p == nil {
		return nil, false
	}
	return o.NpmHostTop99p, true
}

// HasNpmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasNpmHostTop99p() bool {
	return o != nil && o.NpmHostTop99p != nil
}

// SetNpmHostTop99p gets a reference to the given int64 and assigns it to the NpmHostTop99p field.
func (o *UsageSummaryDate) SetNpmHostTop99p(v int64) {
	o.NpmHostTop99p = &v
}

// GetObservabilityPipelinesBytesProcessedSum returns the ObservabilityPipelinesBytesProcessedSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetObservabilityPipelinesBytesProcessedSum() int64 {
	if o == nil || o.ObservabilityPipelinesBytesProcessedSum == nil {
		var ret int64
		return ret
	}
	return *o.ObservabilityPipelinesBytesProcessedSum
}

// GetObservabilityPipelinesBytesProcessedSumOk returns a tuple with the ObservabilityPipelinesBytesProcessedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetObservabilityPipelinesBytesProcessedSumOk() (*int64, bool) {
	if o == nil || o.ObservabilityPipelinesBytesProcessedSum == nil {
		return nil, false
	}
	return o.ObservabilityPipelinesBytesProcessedSum, true
}

// HasObservabilityPipelinesBytesProcessedSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasObservabilityPipelinesBytesProcessedSum() bool {
	return o != nil && o.ObservabilityPipelinesBytesProcessedSum != nil
}

// SetObservabilityPipelinesBytesProcessedSum gets a reference to the given int64 and assigns it to the ObservabilityPipelinesBytesProcessedSum field.
func (o *UsageSummaryDate) SetObservabilityPipelinesBytesProcessedSum(v int64) {
	o.ObservabilityPipelinesBytesProcessedSum = &v
}

// GetOciHostSum returns the OciHostSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetOciHostSum() int64 {
	if o == nil || o.OciHostSum == nil {
		var ret int64
		return ret
	}
	return *o.OciHostSum
}

// GetOciHostSumOk returns a tuple with the OciHostSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetOciHostSumOk() (*int64, bool) {
	if o == nil || o.OciHostSum == nil {
		return nil, false
	}
	return o.OciHostSum, true
}

// HasOciHostSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOciHostSum() bool {
	return o != nil && o.OciHostSum != nil
}

// SetOciHostSum gets a reference to the given int64 and assigns it to the OciHostSum field.
func (o *UsageSummaryDate) SetOciHostSum(v int64) {
	o.OciHostSum = &v
}

// GetOciHostTop99p returns the OciHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetOciHostTop99p() int64 {
	if o == nil || o.OciHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.OciHostTop99p
}

// GetOciHostTop99pOk returns a tuple with the OciHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetOciHostTop99pOk() (*int64, bool) {
	if o == nil || o.OciHostTop99p == nil {
		return nil, false
	}
	return o.OciHostTop99p, true
}

// HasOciHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOciHostTop99p() bool {
	return o != nil && o.OciHostTop99p != nil
}

// SetOciHostTop99p gets a reference to the given int64 and assigns it to the OciHostTop99p field.
func (o *UsageSummaryDate) SetOciHostTop99p(v int64) {
	o.OciHostTop99p = &v
}

// GetOnCallSeatHwm returns the OnCallSeatHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetOnCallSeatHwm() int64 {
	if o == nil || o.OnCallSeatHwm == nil {
		var ret int64
		return ret
	}
	return *o.OnCallSeatHwm
}

// GetOnCallSeatHwmOk returns a tuple with the OnCallSeatHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetOnCallSeatHwmOk() (*int64, bool) {
	if o == nil || o.OnCallSeatHwm == nil {
		return nil, false
	}
	return o.OnCallSeatHwm, true
}

// HasOnCallSeatHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOnCallSeatHwm() bool {
	return o != nil && o.OnCallSeatHwm != nil
}

// SetOnCallSeatHwm gets a reference to the given int64 and assigns it to the OnCallSeatHwm field.
func (o *UsageSummaryDate) SetOnCallSeatHwm(v int64) {
	o.OnCallSeatHwm = &v
}

// GetOnlineArchiveEventsCountSum returns the OnlineArchiveEventsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetOnlineArchiveEventsCountSum() int64 {
	if o == nil || o.OnlineArchiveEventsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.OnlineArchiveEventsCountSum
}

// GetOnlineArchiveEventsCountSumOk returns a tuple with the OnlineArchiveEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetOnlineArchiveEventsCountSumOk() (*int64, bool) {
	if o == nil || o.OnlineArchiveEventsCountSum == nil {
		return nil, false
	}
	return o.OnlineArchiveEventsCountSum, true
}

// HasOnlineArchiveEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOnlineArchiveEventsCountSum() bool {
	return o != nil && o.OnlineArchiveEventsCountSum != nil
}

// SetOnlineArchiveEventsCountSum gets a reference to the given int64 and assigns it to the OnlineArchiveEventsCountSum field.
func (o *UsageSummaryDate) SetOnlineArchiveEventsCountSum(v int64) {
	o.OnlineArchiveEventsCountSum = &v
}

// GetOpentelemetryApmHostTop99p returns the OpentelemetryApmHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetOpentelemetryApmHostTop99p() int64 {
	if o == nil || o.OpentelemetryApmHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryApmHostTop99p
}

// GetOpentelemetryApmHostTop99pOk returns a tuple with the OpentelemetryApmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetOpentelemetryApmHostTop99pOk() (*int64, bool) {
	if o == nil || o.OpentelemetryApmHostTop99p == nil {
		return nil, false
	}
	return o.OpentelemetryApmHostTop99p, true
}

// HasOpentelemetryApmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOpentelemetryApmHostTop99p() bool {
	return o != nil && o.OpentelemetryApmHostTop99p != nil
}

// SetOpentelemetryApmHostTop99p gets a reference to the given int64 and assigns it to the OpentelemetryApmHostTop99p field.
func (o *UsageSummaryDate) SetOpentelemetryApmHostTop99p(v int64) {
	o.OpentelemetryApmHostTop99p = &v
}

// GetOpentelemetryHostTop99p returns the OpentelemetryHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetOpentelemetryHostTop99p() int64 {
	if o == nil || o.OpentelemetryHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.OpentelemetryHostTop99p
}

// GetOpentelemetryHostTop99pOk returns a tuple with the OpentelemetryHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetOpentelemetryHostTop99pOk() (*int64, bool) {
	if o == nil || o.OpentelemetryHostTop99p == nil {
		return nil, false
	}
	return o.OpentelemetryHostTop99p, true
}

// HasOpentelemetryHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOpentelemetryHostTop99p() bool {
	return o != nil && o.OpentelemetryHostTop99p != nil
}

// SetOpentelemetryHostTop99p gets a reference to the given int64 and assigns it to the OpentelemetryHostTop99p field.
func (o *UsageSummaryDate) SetOpentelemetryHostTop99p(v int64) {
	o.OpentelemetryHostTop99p = &v
}

// GetOrgs returns the Orgs field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetOrgs() []UsageSummaryDateOrg {
	if o == nil || o.Orgs == nil {
		var ret []UsageSummaryDateOrg
		return ret
	}
	return o.Orgs
}

// GetOrgsOk returns a tuple with the Orgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetOrgsOk() (*[]UsageSummaryDateOrg, bool) {
	if o == nil || o.Orgs == nil {
		return nil, false
	}
	return &o.Orgs, true
}

// HasOrgs returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasOrgs() bool {
	return o != nil && o.Orgs != nil
}

// SetOrgs gets a reference to the given []UsageSummaryDateOrg and assigns it to the Orgs field.
func (o *UsageSummaryDate) SetOrgs(v []UsageSummaryDateOrg) {
	o.Orgs = v
}

// GetProductAnalyticsSum returns the ProductAnalyticsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetProductAnalyticsSum() int64 {
	if o == nil || o.ProductAnalyticsSum == nil {
		var ret int64
		return ret
	}
	return *o.ProductAnalyticsSum
}

// GetProductAnalyticsSumOk returns a tuple with the ProductAnalyticsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetProductAnalyticsSumOk() (*int64, bool) {
	if o == nil || o.ProductAnalyticsSum == nil {
		return nil, false
	}
	return o.ProductAnalyticsSum, true
}

// HasProductAnalyticsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasProductAnalyticsSum() bool {
	return o != nil && o.ProductAnalyticsSum != nil
}

// SetProductAnalyticsSum gets a reference to the given int64 and assigns it to the ProductAnalyticsSum field.
func (o *UsageSummaryDate) SetProductAnalyticsSum(v int64) {
	o.ProductAnalyticsSum = &v
}

// GetProfilingAasCountTop99p returns the ProfilingAasCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetProfilingAasCountTop99p() int64 {
	if o == nil || o.ProfilingAasCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.ProfilingAasCountTop99p
}

// GetProfilingAasCountTop99pOk returns a tuple with the ProfilingAasCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetProfilingAasCountTop99pOk() (*int64, bool) {
	if o == nil || o.ProfilingAasCountTop99p == nil {
		return nil, false
	}
	return o.ProfilingAasCountTop99p, true
}

// HasProfilingAasCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasProfilingAasCountTop99p() bool {
	return o != nil && o.ProfilingAasCountTop99p != nil
}

// SetProfilingAasCountTop99p gets a reference to the given int64 and assigns it to the ProfilingAasCountTop99p field.
func (o *UsageSummaryDate) SetProfilingAasCountTop99p(v int64) {
	o.ProfilingAasCountTop99p = &v
}

// GetProfilingHostTop99p returns the ProfilingHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetProfilingHostTop99p() int64 {
	if o == nil || o.ProfilingHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.ProfilingHostTop99p
}

// GetProfilingHostTop99pOk returns a tuple with the ProfilingHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetProfilingHostTop99pOk() (*int64, bool) {
	if o == nil || o.ProfilingHostTop99p == nil {
		return nil, false
	}
	return o.ProfilingHostTop99p, true
}

// HasProfilingHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasProfilingHostTop99p() bool {
	return o != nil && o.ProfilingHostTop99p != nil
}

// SetProfilingHostTop99p gets a reference to the given int64 and assigns it to the ProfilingHostTop99p field.
func (o *UsageSummaryDate) SetProfilingHostTop99p(v int64) {
	o.ProfilingHostTop99p = &v
}

// GetProxmoxHostSum returns the ProxmoxHostSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetProxmoxHostSum() int64 {
	if o == nil || o.ProxmoxHostSum == nil {
		var ret int64
		return ret
	}
	return *o.ProxmoxHostSum
}

// GetProxmoxHostSumOk returns a tuple with the ProxmoxHostSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetProxmoxHostSumOk() (*int64, bool) {
	if o == nil || o.ProxmoxHostSum == nil {
		return nil, false
	}
	return o.ProxmoxHostSum, true
}

// HasProxmoxHostSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasProxmoxHostSum() bool {
	return o != nil && o.ProxmoxHostSum != nil
}

// SetProxmoxHostSum gets a reference to the given int64 and assigns it to the ProxmoxHostSum field.
func (o *UsageSummaryDate) SetProxmoxHostSum(v int64) {
	o.ProxmoxHostSum = &v
}

// GetProxmoxHostTop99p returns the ProxmoxHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetProxmoxHostTop99p() int64 {
	if o == nil || o.ProxmoxHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.ProxmoxHostTop99p
}

// GetProxmoxHostTop99pOk returns a tuple with the ProxmoxHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetProxmoxHostTop99pOk() (*int64, bool) {
	if o == nil || o.ProxmoxHostTop99p == nil {
		return nil, false
	}
	return o.ProxmoxHostTop99p, true
}

// HasProxmoxHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasProxmoxHostTop99p() bool {
	return o != nil && o.ProxmoxHostTop99p != nil
}

// SetProxmoxHostTop99p gets a reference to the given int64 and assigns it to the ProxmoxHostTop99p field.
func (o *UsageSummaryDate) SetProxmoxHostTop99p(v int64) {
	o.ProxmoxHostTop99p = &v
}

// GetPublishedAppHwm returns the PublishedAppHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetPublishedAppHwm() int64 {
	if o == nil || o.PublishedAppHwm == nil {
		var ret int64
		return ret
	}
	return *o.PublishedAppHwm
}

// GetPublishedAppHwmOk returns a tuple with the PublishedAppHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetPublishedAppHwmOk() (*int64, bool) {
	if o == nil || o.PublishedAppHwm == nil {
		return nil, false
	}
	return o.PublishedAppHwm, true
}

// HasPublishedAppHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasPublishedAppHwm() bool {
	return o != nil && o.PublishedAppHwm != nil
}

// SetPublishedAppHwm gets a reference to the given int64 and assigns it to the PublishedAppHwm field.
func (o *UsageSummaryDate) SetPublishedAppHwm(v int64) {
	o.PublishedAppHwm = &v
}

// GetRumBrowserAndMobileSessionCount returns the RumBrowserAndMobileSessionCount field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumBrowserAndMobileSessionCount() int64 {
	if o == nil || o.RumBrowserAndMobileSessionCount == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserAndMobileSessionCount
}

// GetRumBrowserAndMobileSessionCountOk returns a tuple with the RumBrowserAndMobileSessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumBrowserAndMobileSessionCountOk() (*int64, bool) {
	if o == nil || o.RumBrowserAndMobileSessionCount == nil {
		return nil, false
	}
	return o.RumBrowserAndMobileSessionCount, true
}

// HasRumBrowserAndMobileSessionCount returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumBrowserAndMobileSessionCount() bool {
	return o != nil && o.RumBrowserAndMobileSessionCount != nil
}

// SetRumBrowserAndMobileSessionCount gets a reference to the given int64 and assigns it to the RumBrowserAndMobileSessionCount field.
func (o *UsageSummaryDate) SetRumBrowserAndMobileSessionCount(v int64) {
	o.RumBrowserAndMobileSessionCount = &v
}

// GetRumBrowserLegacySessionCountSum returns the RumBrowserLegacySessionCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumBrowserLegacySessionCountSum() int64 {
	if o == nil || o.RumBrowserLegacySessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserLegacySessionCountSum
}

// GetRumBrowserLegacySessionCountSumOk returns a tuple with the RumBrowserLegacySessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumBrowserLegacySessionCountSumOk() (*int64, bool) {
	if o == nil || o.RumBrowserLegacySessionCountSum == nil {
		return nil, false
	}
	return o.RumBrowserLegacySessionCountSum, true
}

// HasRumBrowserLegacySessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumBrowserLegacySessionCountSum() bool {
	return o != nil && o.RumBrowserLegacySessionCountSum != nil
}

// SetRumBrowserLegacySessionCountSum gets a reference to the given int64 and assigns it to the RumBrowserLegacySessionCountSum field.
func (o *UsageSummaryDate) SetRumBrowserLegacySessionCountSum(v int64) {
	o.RumBrowserLegacySessionCountSum = &v
}

// GetRumBrowserLiteSessionCountSum returns the RumBrowserLiteSessionCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumBrowserLiteSessionCountSum() int64 {
	if o == nil || o.RumBrowserLiteSessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserLiteSessionCountSum
}

// GetRumBrowserLiteSessionCountSumOk returns a tuple with the RumBrowserLiteSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumBrowserLiteSessionCountSumOk() (*int64, bool) {
	if o == nil || o.RumBrowserLiteSessionCountSum == nil {
		return nil, false
	}
	return o.RumBrowserLiteSessionCountSum, true
}

// HasRumBrowserLiteSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumBrowserLiteSessionCountSum() bool {
	return o != nil && o.RumBrowserLiteSessionCountSum != nil
}

// SetRumBrowserLiteSessionCountSum gets a reference to the given int64 and assigns it to the RumBrowserLiteSessionCountSum field.
func (o *UsageSummaryDate) SetRumBrowserLiteSessionCountSum(v int64) {
	o.RumBrowserLiteSessionCountSum = &v
}

// GetRumBrowserReplaySessionCountSum returns the RumBrowserReplaySessionCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumBrowserReplaySessionCountSum() int64 {
	if o == nil || o.RumBrowserReplaySessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.RumBrowserReplaySessionCountSum
}

// GetRumBrowserReplaySessionCountSumOk returns a tuple with the RumBrowserReplaySessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumBrowserReplaySessionCountSumOk() (*int64, bool) {
	if o == nil || o.RumBrowserReplaySessionCountSum == nil {
		return nil, false
	}
	return o.RumBrowserReplaySessionCountSum, true
}

// HasRumBrowserReplaySessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumBrowserReplaySessionCountSum() bool {
	return o != nil && o.RumBrowserReplaySessionCountSum != nil
}

// SetRumBrowserReplaySessionCountSum gets a reference to the given int64 and assigns it to the RumBrowserReplaySessionCountSum field.
func (o *UsageSummaryDate) SetRumBrowserReplaySessionCountSum(v int64) {
	o.RumBrowserReplaySessionCountSum = &v
}

// GetRumIndexedSessionsSum returns the RumIndexedSessionsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumIndexedSessionsSum() int64 {
	if o == nil || o.RumIndexedSessionsSum == nil {
		var ret int64
		return ret
	}
	return *o.RumIndexedSessionsSum
}

// GetRumIndexedSessionsSumOk returns a tuple with the RumIndexedSessionsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumIndexedSessionsSumOk() (*int64, bool) {
	if o == nil || o.RumIndexedSessionsSum == nil {
		return nil, false
	}
	return o.RumIndexedSessionsSum, true
}

// HasRumIndexedSessionsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumIndexedSessionsSum() bool {
	return o != nil && o.RumIndexedSessionsSum != nil
}

// SetRumIndexedSessionsSum gets a reference to the given int64 and assigns it to the RumIndexedSessionsSum field.
func (o *UsageSummaryDate) SetRumIndexedSessionsSum(v int64) {
	o.RumIndexedSessionsSum = &v
}

// GetRumIngestedSessionsSum returns the RumIngestedSessionsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumIngestedSessionsSum() int64 {
	if o == nil || o.RumIngestedSessionsSum == nil {
		var ret int64
		return ret
	}
	return *o.RumIngestedSessionsSum
}

// GetRumIngestedSessionsSumOk returns a tuple with the RumIngestedSessionsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumIngestedSessionsSumOk() (*int64, bool) {
	if o == nil || o.RumIngestedSessionsSum == nil {
		return nil, false
	}
	return o.RumIngestedSessionsSum, true
}

// HasRumIngestedSessionsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumIngestedSessionsSum() bool {
	return o != nil && o.RumIngestedSessionsSum != nil
}

// SetRumIngestedSessionsSum gets a reference to the given int64 and assigns it to the RumIngestedSessionsSum field.
func (o *UsageSummaryDate) SetRumIngestedSessionsSum(v int64) {
	o.RumIngestedSessionsSum = &v
}

// GetRumLiteSessionCountSum returns the RumLiteSessionCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumLiteSessionCountSum() int64 {
	if o == nil || o.RumLiteSessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.RumLiteSessionCountSum
}

// GetRumLiteSessionCountSumOk returns a tuple with the RumLiteSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumLiteSessionCountSumOk() (*int64, bool) {
	if o == nil || o.RumLiteSessionCountSum == nil {
		return nil, false
	}
	return o.RumLiteSessionCountSum, true
}

// HasRumLiteSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumLiteSessionCountSum() bool {
	return o != nil && o.RumLiteSessionCountSum != nil
}

// SetRumLiteSessionCountSum gets a reference to the given int64 and assigns it to the RumLiteSessionCountSum field.
func (o *UsageSummaryDate) SetRumLiteSessionCountSum(v int64) {
	o.RumLiteSessionCountSum = &v
}

// GetRumMobileLegacySessionCountAndroidSum returns the RumMobileLegacySessionCountAndroidSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountAndroidSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountAndroidSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountAndroidSum
}

// GetRumMobileLegacySessionCountAndroidSumOk returns a tuple with the RumMobileLegacySessionCountAndroidSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountAndroidSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountAndroidSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountAndroidSum, true
}

// HasRumMobileLegacySessionCountAndroidSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLegacySessionCountAndroidSum() bool {
	return o != nil && o.RumMobileLegacySessionCountAndroidSum != nil
}

// SetRumMobileLegacySessionCountAndroidSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountAndroidSum field.
func (o *UsageSummaryDate) SetRumMobileLegacySessionCountAndroidSum(v int64) {
	o.RumMobileLegacySessionCountAndroidSum = &v
}

// GetRumMobileLegacySessionCountFlutterSum returns the RumMobileLegacySessionCountFlutterSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountFlutterSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountFlutterSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountFlutterSum
}

// GetRumMobileLegacySessionCountFlutterSumOk returns a tuple with the RumMobileLegacySessionCountFlutterSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountFlutterSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountFlutterSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountFlutterSum, true
}

// HasRumMobileLegacySessionCountFlutterSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLegacySessionCountFlutterSum() bool {
	return o != nil && o.RumMobileLegacySessionCountFlutterSum != nil
}

// SetRumMobileLegacySessionCountFlutterSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountFlutterSum field.
func (o *UsageSummaryDate) SetRumMobileLegacySessionCountFlutterSum(v int64) {
	o.RumMobileLegacySessionCountFlutterSum = &v
}

// GetRumMobileLegacySessionCountIosSum returns the RumMobileLegacySessionCountIosSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountIosSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountIosSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountIosSum
}

// GetRumMobileLegacySessionCountIosSumOk returns a tuple with the RumMobileLegacySessionCountIosSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountIosSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountIosSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountIosSum, true
}

// HasRumMobileLegacySessionCountIosSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLegacySessionCountIosSum() bool {
	return o != nil && o.RumMobileLegacySessionCountIosSum != nil
}

// SetRumMobileLegacySessionCountIosSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountIosSum field.
func (o *UsageSummaryDate) SetRumMobileLegacySessionCountIosSum(v int64) {
	o.RumMobileLegacySessionCountIosSum = &v
}

// GetRumMobileLegacySessionCountReactnativeSum returns the RumMobileLegacySessionCountReactnativeSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountReactnativeSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountReactnativeSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountReactnativeSum
}

// GetRumMobileLegacySessionCountReactnativeSumOk returns a tuple with the RumMobileLegacySessionCountReactnativeSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountReactnativeSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountReactnativeSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountReactnativeSum, true
}

// HasRumMobileLegacySessionCountReactnativeSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLegacySessionCountReactnativeSum() bool {
	return o != nil && o.RumMobileLegacySessionCountReactnativeSum != nil
}

// SetRumMobileLegacySessionCountReactnativeSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountReactnativeSum field.
func (o *UsageSummaryDate) SetRumMobileLegacySessionCountReactnativeSum(v int64) {
	o.RumMobileLegacySessionCountReactnativeSum = &v
}

// GetRumMobileLegacySessionCountRokuSum returns the RumMobileLegacySessionCountRokuSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountRokuSum() int64 {
	if o == nil || o.RumMobileLegacySessionCountRokuSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLegacySessionCountRokuSum
}

// GetRumMobileLegacySessionCountRokuSumOk returns a tuple with the RumMobileLegacySessionCountRokuSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLegacySessionCountRokuSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLegacySessionCountRokuSum == nil {
		return nil, false
	}
	return o.RumMobileLegacySessionCountRokuSum, true
}

// HasRumMobileLegacySessionCountRokuSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLegacySessionCountRokuSum() bool {
	return o != nil && o.RumMobileLegacySessionCountRokuSum != nil
}

// SetRumMobileLegacySessionCountRokuSum gets a reference to the given int64 and assigns it to the RumMobileLegacySessionCountRokuSum field.
func (o *UsageSummaryDate) SetRumMobileLegacySessionCountRokuSum(v int64) {
	o.RumMobileLegacySessionCountRokuSum = &v
}

// GetRumMobileLiteSessionCountAndroidSum returns the RumMobileLiteSessionCountAndroidSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountAndroidSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountAndroidSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountAndroidSum
}

// GetRumMobileLiteSessionCountAndroidSumOk returns a tuple with the RumMobileLiteSessionCountAndroidSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountAndroidSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountAndroidSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountAndroidSum, true
}

// HasRumMobileLiteSessionCountAndroidSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLiteSessionCountAndroidSum() bool {
	return o != nil && o.RumMobileLiteSessionCountAndroidSum != nil
}

// SetRumMobileLiteSessionCountAndroidSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountAndroidSum field.
func (o *UsageSummaryDate) SetRumMobileLiteSessionCountAndroidSum(v int64) {
	o.RumMobileLiteSessionCountAndroidSum = &v
}

// GetRumMobileLiteSessionCountFlutterSum returns the RumMobileLiteSessionCountFlutterSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountFlutterSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountFlutterSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountFlutterSum
}

// GetRumMobileLiteSessionCountFlutterSumOk returns a tuple with the RumMobileLiteSessionCountFlutterSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountFlutterSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountFlutterSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountFlutterSum, true
}

// HasRumMobileLiteSessionCountFlutterSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLiteSessionCountFlutterSum() bool {
	return o != nil && o.RumMobileLiteSessionCountFlutterSum != nil
}

// SetRumMobileLiteSessionCountFlutterSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountFlutterSum field.
func (o *UsageSummaryDate) SetRumMobileLiteSessionCountFlutterSum(v int64) {
	o.RumMobileLiteSessionCountFlutterSum = &v
}

// GetRumMobileLiteSessionCountIosSum returns the RumMobileLiteSessionCountIosSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountIosSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountIosSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountIosSum
}

// GetRumMobileLiteSessionCountIosSumOk returns a tuple with the RumMobileLiteSessionCountIosSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountIosSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountIosSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountIosSum, true
}

// HasRumMobileLiteSessionCountIosSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLiteSessionCountIosSum() bool {
	return o != nil && o.RumMobileLiteSessionCountIosSum != nil
}

// SetRumMobileLiteSessionCountIosSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountIosSum field.
func (o *UsageSummaryDate) SetRumMobileLiteSessionCountIosSum(v int64) {
	o.RumMobileLiteSessionCountIosSum = &v
}

// GetRumMobileLiteSessionCountKotlinmultiplatformSum returns the RumMobileLiteSessionCountKotlinmultiplatformSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountKotlinmultiplatformSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountKotlinmultiplatformSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountKotlinmultiplatformSum
}

// GetRumMobileLiteSessionCountKotlinmultiplatformSumOk returns a tuple with the RumMobileLiteSessionCountKotlinmultiplatformSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountKotlinmultiplatformSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountKotlinmultiplatformSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountKotlinmultiplatformSum, true
}

// HasRumMobileLiteSessionCountKotlinmultiplatformSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLiteSessionCountKotlinmultiplatformSum() bool {
	return o != nil && o.RumMobileLiteSessionCountKotlinmultiplatformSum != nil
}

// SetRumMobileLiteSessionCountKotlinmultiplatformSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountKotlinmultiplatformSum field.
func (o *UsageSummaryDate) SetRumMobileLiteSessionCountKotlinmultiplatformSum(v int64) {
	o.RumMobileLiteSessionCountKotlinmultiplatformSum = &v
}

// GetRumMobileLiteSessionCountReactnativeSum returns the RumMobileLiteSessionCountReactnativeSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountReactnativeSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountReactnativeSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountReactnativeSum
}

// GetRumMobileLiteSessionCountReactnativeSumOk returns a tuple with the RumMobileLiteSessionCountReactnativeSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountReactnativeSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountReactnativeSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountReactnativeSum, true
}

// HasRumMobileLiteSessionCountReactnativeSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLiteSessionCountReactnativeSum() bool {
	return o != nil && o.RumMobileLiteSessionCountReactnativeSum != nil
}

// SetRumMobileLiteSessionCountReactnativeSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountReactnativeSum field.
func (o *UsageSummaryDate) SetRumMobileLiteSessionCountReactnativeSum(v int64) {
	o.RumMobileLiteSessionCountReactnativeSum = &v
}

// GetRumMobileLiteSessionCountRokuSum returns the RumMobileLiteSessionCountRokuSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountRokuSum() int64 {
	if o == nil || o.RumMobileLiteSessionCountRokuSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountRokuSum
}

// GetRumMobileLiteSessionCountRokuSumOk returns a tuple with the RumMobileLiteSessionCountRokuSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountRokuSumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountRokuSum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountRokuSum, true
}

// HasRumMobileLiteSessionCountRokuSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLiteSessionCountRokuSum() bool {
	return o != nil && o.RumMobileLiteSessionCountRokuSum != nil
}

// SetRumMobileLiteSessionCountRokuSum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountRokuSum field.
func (o *UsageSummaryDate) SetRumMobileLiteSessionCountRokuSum(v int64) {
	o.RumMobileLiteSessionCountRokuSum = &v
}

// GetRumMobileLiteSessionCountUnitySum returns the RumMobileLiteSessionCountUnitySum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountUnitySum() int64 {
	if o == nil || o.RumMobileLiteSessionCountUnitySum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileLiteSessionCountUnitySum
}

// GetRumMobileLiteSessionCountUnitySumOk returns a tuple with the RumMobileLiteSessionCountUnitySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileLiteSessionCountUnitySumOk() (*int64, bool) {
	if o == nil || o.RumMobileLiteSessionCountUnitySum == nil {
		return nil, false
	}
	return o.RumMobileLiteSessionCountUnitySum, true
}

// HasRumMobileLiteSessionCountUnitySum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileLiteSessionCountUnitySum() bool {
	return o != nil && o.RumMobileLiteSessionCountUnitySum != nil
}

// SetRumMobileLiteSessionCountUnitySum gets a reference to the given int64 and assigns it to the RumMobileLiteSessionCountUnitySum field.
func (o *UsageSummaryDate) SetRumMobileLiteSessionCountUnitySum(v int64) {
	o.RumMobileLiteSessionCountUnitySum = &v
}

// GetRumMobileReplaySessionCountAndroidSum returns the RumMobileReplaySessionCountAndroidSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileReplaySessionCountAndroidSum() int64 {
	if o == nil || o.RumMobileReplaySessionCountAndroidSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileReplaySessionCountAndroidSum
}

// GetRumMobileReplaySessionCountAndroidSumOk returns a tuple with the RumMobileReplaySessionCountAndroidSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileReplaySessionCountAndroidSumOk() (*int64, bool) {
	if o == nil || o.RumMobileReplaySessionCountAndroidSum == nil {
		return nil, false
	}
	return o.RumMobileReplaySessionCountAndroidSum, true
}

// HasRumMobileReplaySessionCountAndroidSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileReplaySessionCountAndroidSum() bool {
	return o != nil && o.RumMobileReplaySessionCountAndroidSum != nil
}

// SetRumMobileReplaySessionCountAndroidSum gets a reference to the given int64 and assigns it to the RumMobileReplaySessionCountAndroidSum field.
func (o *UsageSummaryDate) SetRumMobileReplaySessionCountAndroidSum(v int64) {
	o.RumMobileReplaySessionCountAndroidSum = &v
}

// GetRumMobileReplaySessionCountIosSum returns the RumMobileReplaySessionCountIosSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileReplaySessionCountIosSum() int64 {
	if o == nil || o.RumMobileReplaySessionCountIosSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileReplaySessionCountIosSum
}

// GetRumMobileReplaySessionCountIosSumOk returns a tuple with the RumMobileReplaySessionCountIosSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileReplaySessionCountIosSumOk() (*int64, bool) {
	if o == nil || o.RumMobileReplaySessionCountIosSum == nil {
		return nil, false
	}
	return o.RumMobileReplaySessionCountIosSum, true
}

// HasRumMobileReplaySessionCountIosSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileReplaySessionCountIosSum() bool {
	return o != nil && o.RumMobileReplaySessionCountIosSum != nil
}

// SetRumMobileReplaySessionCountIosSum gets a reference to the given int64 and assigns it to the RumMobileReplaySessionCountIosSum field.
func (o *UsageSummaryDate) SetRumMobileReplaySessionCountIosSum(v int64) {
	o.RumMobileReplaySessionCountIosSum = &v
}

// GetRumMobileReplaySessionCountKotlinmultiplatformSum returns the RumMobileReplaySessionCountKotlinmultiplatformSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileReplaySessionCountKotlinmultiplatformSum() int64 {
	if o == nil || o.RumMobileReplaySessionCountKotlinmultiplatformSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileReplaySessionCountKotlinmultiplatformSum
}

// GetRumMobileReplaySessionCountKotlinmultiplatformSumOk returns a tuple with the RumMobileReplaySessionCountKotlinmultiplatformSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileReplaySessionCountKotlinmultiplatformSumOk() (*int64, bool) {
	if o == nil || o.RumMobileReplaySessionCountKotlinmultiplatformSum == nil {
		return nil, false
	}
	return o.RumMobileReplaySessionCountKotlinmultiplatformSum, true
}

// HasRumMobileReplaySessionCountKotlinmultiplatformSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileReplaySessionCountKotlinmultiplatformSum() bool {
	return o != nil && o.RumMobileReplaySessionCountKotlinmultiplatformSum != nil
}

// SetRumMobileReplaySessionCountKotlinmultiplatformSum gets a reference to the given int64 and assigns it to the RumMobileReplaySessionCountKotlinmultiplatformSum field.
func (o *UsageSummaryDate) SetRumMobileReplaySessionCountKotlinmultiplatformSum(v int64) {
	o.RumMobileReplaySessionCountKotlinmultiplatformSum = &v
}

// GetRumMobileReplaySessionCountReactnativeSum returns the RumMobileReplaySessionCountReactnativeSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumMobileReplaySessionCountReactnativeSum() int64 {
	if o == nil || o.RumMobileReplaySessionCountReactnativeSum == nil {
		var ret int64
		return ret
	}
	return *o.RumMobileReplaySessionCountReactnativeSum
}

// GetRumMobileReplaySessionCountReactnativeSumOk returns a tuple with the RumMobileReplaySessionCountReactnativeSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumMobileReplaySessionCountReactnativeSumOk() (*int64, bool) {
	if o == nil || o.RumMobileReplaySessionCountReactnativeSum == nil {
		return nil, false
	}
	return o.RumMobileReplaySessionCountReactnativeSum, true
}

// HasRumMobileReplaySessionCountReactnativeSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumMobileReplaySessionCountReactnativeSum() bool {
	return o != nil && o.RumMobileReplaySessionCountReactnativeSum != nil
}

// SetRumMobileReplaySessionCountReactnativeSum gets a reference to the given int64 and assigns it to the RumMobileReplaySessionCountReactnativeSum field.
func (o *UsageSummaryDate) SetRumMobileReplaySessionCountReactnativeSum(v int64) {
	o.RumMobileReplaySessionCountReactnativeSum = &v
}

// GetRumReplaySessionCountSum returns the RumReplaySessionCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumReplaySessionCountSum() int64 {
	if o == nil || o.RumReplaySessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.RumReplaySessionCountSum
}

// GetRumReplaySessionCountSumOk returns a tuple with the RumReplaySessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumReplaySessionCountSumOk() (*int64, bool) {
	if o == nil || o.RumReplaySessionCountSum == nil {
		return nil, false
	}
	return o.RumReplaySessionCountSum, true
}

// HasRumReplaySessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumReplaySessionCountSum() bool {
	return o != nil && o.RumReplaySessionCountSum != nil
}

// SetRumReplaySessionCountSum gets a reference to the given int64 and assigns it to the RumReplaySessionCountSum field.
func (o *UsageSummaryDate) SetRumReplaySessionCountSum(v int64) {
	o.RumReplaySessionCountSum = &v
}

// GetRumSessionCountSum returns the RumSessionCountSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetRumSessionCountSum() int64 {
	if o == nil || o.RumSessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.RumSessionCountSum
}

// GetRumSessionCountSumOk returns a tuple with the RumSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetRumSessionCountSumOk() (*int64, bool) {
	if o == nil || o.RumSessionCountSum == nil {
		return nil, false
	}
	return o.RumSessionCountSum, true
}

// HasRumSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumSessionCountSum() bool {
	return o != nil && o.RumSessionCountSum != nil
}

// SetRumSessionCountSum gets a reference to the given int64 and assigns it to the RumSessionCountSum field.
// Deprecated
func (o *UsageSummaryDate) SetRumSessionCountSum(v int64) {
	o.RumSessionCountSum = &v
}

// GetRumSessionReplayAddOnSum returns the RumSessionReplayAddOnSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumSessionReplayAddOnSum() int64 {
	if o == nil || o.RumSessionReplayAddOnSum == nil {
		var ret int64
		return ret
	}
	return *o.RumSessionReplayAddOnSum
}

// GetRumSessionReplayAddOnSumOk returns a tuple with the RumSessionReplayAddOnSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumSessionReplayAddOnSumOk() (*int64, bool) {
	if o == nil || o.RumSessionReplayAddOnSum == nil {
		return nil, false
	}
	return o.RumSessionReplayAddOnSum, true
}

// HasRumSessionReplayAddOnSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumSessionReplayAddOnSum() bool {
	return o != nil && o.RumSessionReplayAddOnSum != nil
}

// SetRumSessionReplayAddOnSum gets a reference to the given int64 and assigns it to the RumSessionReplayAddOnSum field.
func (o *UsageSummaryDate) SetRumSessionReplayAddOnSum(v int64) {
	o.RumSessionReplayAddOnSum = &v
}

// GetRumTotalSessionCountSum returns the RumTotalSessionCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetRumTotalSessionCountSum() int64 {
	if o == nil || o.RumTotalSessionCountSum == nil {
		var ret int64
		return ret
	}
	return *o.RumTotalSessionCountSum
}

// GetRumTotalSessionCountSumOk returns a tuple with the RumTotalSessionCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetRumTotalSessionCountSumOk() (*int64, bool) {
	if o == nil || o.RumTotalSessionCountSum == nil {
		return nil, false
	}
	return o.RumTotalSessionCountSum, true
}

// HasRumTotalSessionCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumTotalSessionCountSum() bool {
	return o != nil && o.RumTotalSessionCountSum != nil
}

// SetRumTotalSessionCountSum gets a reference to the given int64 and assigns it to the RumTotalSessionCountSum field.
func (o *UsageSummaryDate) SetRumTotalSessionCountSum(v int64) {
	o.RumTotalSessionCountSum = &v
}

// GetRumUnitsSum returns the RumUnitsSum field value if set, zero value otherwise.
// Deprecated
func (o *UsageSummaryDate) GetRumUnitsSum() int64 {
	if o == nil || o.RumUnitsSum == nil {
		var ret int64
		return ret
	}
	return *o.RumUnitsSum
}

// GetRumUnitsSumOk returns a tuple with the RumUnitsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UsageSummaryDate) GetRumUnitsSumOk() (*int64, bool) {
	if o == nil || o.RumUnitsSum == nil {
		return nil, false
	}
	return o.RumUnitsSum, true
}

// HasRumUnitsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasRumUnitsSum() bool {
	return o != nil && o.RumUnitsSum != nil
}

// SetRumUnitsSum gets a reference to the given int64 and assigns it to the RumUnitsSum field.
// Deprecated
func (o *UsageSummaryDate) SetRumUnitsSum(v int64) {
	o.RumUnitsSum = &v
}

// GetScaFargateCountAvg returns the ScaFargateCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetScaFargateCountAvg() int64 {
	if o == nil || o.ScaFargateCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.ScaFargateCountAvg
}

// GetScaFargateCountAvgOk returns a tuple with the ScaFargateCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetScaFargateCountAvgOk() (*int64, bool) {
	if o == nil || o.ScaFargateCountAvg == nil {
		return nil, false
	}
	return o.ScaFargateCountAvg, true
}

// HasScaFargateCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasScaFargateCountAvg() bool {
	return o != nil && o.ScaFargateCountAvg != nil
}

// SetScaFargateCountAvg gets a reference to the given int64 and assigns it to the ScaFargateCountAvg field.
func (o *UsageSummaryDate) SetScaFargateCountAvg(v int64) {
	o.ScaFargateCountAvg = &v
}

// GetScaFargateCountHwm returns the ScaFargateCountHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetScaFargateCountHwm() int64 {
	if o == nil || o.ScaFargateCountHwm == nil {
		var ret int64
		return ret
	}
	return *o.ScaFargateCountHwm
}

// GetScaFargateCountHwmOk returns a tuple with the ScaFargateCountHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetScaFargateCountHwmOk() (*int64, bool) {
	if o == nil || o.ScaFargateCountHwm == nil {
		return nil, false
	}
	return o.ScaFargateCountHwm, true
}

// HasScaFargateCountHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasScaFargateCountHwm() bool {
	return o != nil && o.ScaFargateCountHwm != nil
}

// SetScaFargateCountHwm gets a reference to the given int64 and assigns it to the ScaFargateCountHwm field.
func (o *UsageSummaryDate) SetScaFargateCountHwm(v int64) {
	o.ScaFargateCountHwm = &v
}

// GetSdsApmScannedBytesSum returns the SdsApmScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSdsApmScannedBytesSum() int64 {
	if o == nil || o.SdsApmScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsApmScannedBytesSum
}

// GetSdsApmScannedBytesSumOk returns a tuple with the SdsApmScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSdsApmScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsApmScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsApmScannedBytesSum, true
}

// HasSdsApmScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsApmScannedBytesSum() bool {
	return o != nil && o.SdsApmScannedBytesSum != nil
}

// SetSdsApmScannedBytesSum gets a reference to the given int64 and assigns it to the SdsApmScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsApmScannedBytesSum(v int64) {
	o.SdsApmScannedBytesSum = &v
}

// GetSdsEventsScannedBytesSum returns the SdsEventsScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSdsEventsScannedBytesSum() int64 {
	if o == nil || o.SdsEventsScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsEventsScannedBytesSum
}

// GetSdsEventsScannedBytesSumOk returns a tuple with the SdsEventsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSdsEventsScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsEventsScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsEventsScannedBytesSum, true
}

// HasSdsEventsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsEventsScannedBytesSum() bool {
	return o != nil && o.SdsEventsScannedBytesSum != nil
}

// SetSdsEventsScannedBytesSum gets a reference to the given int64 and assigns it to the SdsEventsScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsEventsScannedBytesSum(v int64) {
	o.SdsEventsScannedBytesSum = &v
}

// GetSdsLogsScannedBytesSum returns the SdsLogsScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSdsLogsScannedBytesSum() int64 {
	if o == nil || o.SdsLogsScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsLogsScannedBytesSum
}

// GetSdsLogsScannedBytesSumOk returns a tuple with the SdsLogsScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSdsLogsScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsLogsScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsLogsScannedBytesSum, true
}

// HasSdsLogsScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsLogsScannedBytesSum() bool {
	return o != nil && o.SdsLogsScannedBytesSum != nil
}

// SetSdsLogsScannedBytesSum gets a reference to the given int64 and assigns it to the SdsLogsScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsLogsScannedBytesSum(v int64) {
	o.SdsLogsScannedBytesSum = &v
}

// GetSdsRumScannedBytesSum returns the SdsRumScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSdsRumScannedBytesSum() int64 {
	if o == nil || o.SdsRumScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsRumScannedBytesSum
}

// GetSdsRumScannedBytesSumOk returns a tuple with the SdsRumScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSdsRumScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsRumScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsRumScannedBytesSum, true
}

// HasSdsRumScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsRumScannedBytesSum() bool {
	return o != nil && o.SdsRumScannedBytesSum != nil
}

// SetSdsRumScannedBytesSum gets a reference to the given int64 and assigns it to the SdsRumScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsRumScannedBytesSum(v int64) {
	o.SdsRumScannedBytesSum = &v
}

// GetSdsTotalScannedBytesSum returns the SdsTotalScannedBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSdsTotalScannedBytesSum() int64 {
	if o == nil || o.SdsTotalScannedBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.SdsTotalScannedBytesSum
}

// GetSdsTotalScannedBytesSumOk returns a tuple with the SdsTotalScannedBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSdsTotalScannedBytesSumOk() (*int64, bool) {
	if o == nil || o.SdsTotalScannedBytesSum == nil {
		return nil, false
	}
	return o.SdsTotalScannedBytesSum, true
}

// HasSdsTotalScannedBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSdsTotalScannedBytesSum() bool {
	return o != nil && o.SdsTotalScannedBytesSum != nil
}

// SetSdsTotalScannedBytesSum gets a reference to the given int64 and assigns it to the SdsTotalScannedBytesSum field.
func (o *UsageSummaryDate) SetSdsTotalScannedBytesSum(v int64) {
	o.SdsTotalScannedBytesSum = &v
}

// GetServerlessAppsApmApmAzureAppserviceInstancesAvg returns the ServerlessAppsApmApmAzureAppserviceInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmApmAzureAppserviceInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmApmAzureAppserviceInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmAzureAppserviceInstancesAvg
}

// GetServerlessAppsApmApmAzureAppserviceInstancesAvgOk returns a tuple with the ServerlessAppsApmApmAzureAppserviceInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmApmAzureAppserviceInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmAzureAppserviceInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmAzureAppserviceInstancesAvg, true
}

// HasServerlessAppsApmApmAzureAppserviceInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmApmAzureAppserviceInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmApmAzureAppserviceInstancesAvg != nil
}

// SetServerlessAppsApmApmAzureAppserviceInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmAzureAppserviceInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmApmAzureAppserviceInstancesAvg(v int64) {
	o.ServerlessAppsApmApmAzureAppserviceInstancesAvg = &v
}

// GetServerlessAppsApmApmAzureAzurefunctionInstancesAvg returns the ServerlessAppsApmApmAzureAzurefunctionInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmApmAzureAzurefunctionInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg
}

// GetServerlessAppsApmApmAzureAzurefunctionInstancesAvgOk returns a tuple with the ServerlessAppsApmApmAzureAzurefunctionInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmApmAzureAzurefunctionInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg, true
}

// HasServerlessAppsApmApmAzureAzurefunctionInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmApmAzureAzurefunctionInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg != nil
}

// SetServerlessAppsApmApmAzureAzurefunctionInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmAzureAzurefunctionInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmApmAzureAzurefunctionInstancesAvg(v int64) {
	o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg = &v
}

// GetServerlessAppsApmApmAzureContainerappInstancesAvg returns the ServerlessAppsApmApmAzureContainerappInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmApmAzureContainerappInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmApmAzureContainerappInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmAzureContainerappInstancesAvg
}

// GetServerlessAppsApmApmAzureContainerappInstancesAvgOk returns a tuple with the ServerlessAppsApmApmAzureContainerappInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmApmAzureContainerappInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmAzureContainerappInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmAzureContainerappInstancesAvg, true
}

// HasServerlessAppsApmApmAzureContainerappInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmApmAzureContainerappInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmApmAzureContainerappInstancesAvg != nil
}

// SetServerlessAppsApmApmAzureContainerappInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmAzureContainerappInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmApmAzureContainerappInstancesAvg(v int64) {
	o.ServerlessAppsApmApmAzureContainerappInstancesAvg = &v
}

// GetServerlessAppsApmApmFargateEcsTasksAvg returns the ServerlessAppsApmApmFargateEcsTasksAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmApmFargateEcsTasksAvg() int64 {
	if o == nil || o.ServerlessAppsApmApmFargateEcsTasksAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmFargateEcsTasksAvg
}

// GetServerlessAppsApmApmFargateEcsTasksAvgOk returns a tuple with the ServerlessAppsApmApmFargateEcsTasksAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmApmFargateEcsTasksAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmFargateEcsTasksAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmFargateEcsTasksAvg, true
}

// HasServerlessAppsApmApmFargateEcsTasksAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmApmFargateEcsTasksAvg() bool {
	return o != nil && o.ServerlessAppsApmApmFargateEcsTasksAvg != nil
}

// SetServerlessAppsApmApmFargateEcsTasksAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmFargateEcsTasksAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmApmFargateEcsTasksAvg(v int64) {
	o.ServerlessAppsApmApmFargateEcsTasksAvg = &v
}

// GetServerlessAppsApmApmGcpCloudfunctionInstancesAvg returns the ServerlessAppsApmApmGcpCloudfunctionInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmApmGcpCloudfunctionInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg
}

// GetServerlessAppsApmApmGcpCloudfunctionInstancesAvgOk returns a tuple with the ServerlessAppsApmApmGcpCloudfunctionInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmApmGcpCloudfunctionInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg, true
}

// HasServerlessAppsApmApmGcpCloudfunctionInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmApmGcpCloudfunctionInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg != nil
}

// SetServerlessAppsApmApmGcpCloudfunctionInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmGcpCloudfunctionInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmApmGcpCloudfunctionInstancesAvg(v int64) {
	o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg = &v
}

// GetServerlessAppsApmApmGcpCloudrunInstancesAvg returns the ServerlessAppsApmApmGcpCloudrunInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmApmGcpCloudrunInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmApmGcpCloudrunInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmGcpCloudrunInstancesAvg
}

// GetServerlessAppsApmApmGcpCloudrunInstancesAvgOk returns a tuple with the ServerlessAppsApmApmGcpCloudrunInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmApmGcpCloudrunInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmGcpCloudrunInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmGcpCloudrunInstancesAvg, true
}

// HasServerlessAppsApmApmGcpCloudrunInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmApmGcpCloudrunInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmApmGcpCloudrunInstancesAvg != nil
}

// SetServerlessAppsApmApmGcpCloudrunInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmGcpCloudrunInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmApmGcpCloudrunInstancesAvg(v int64) {
	o.ServerlessAppsApmApmGcpCloudrunInstancesAvg = &v
}

// GetServerlessAppsApmAvg returns the ServerlessAppsApmAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmAvg() int64 {
	if o == nil || o.ServerlessAppsApmAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmAvg
}

// GetServerlessAppsApmAvgOk returns a tuple with the ServerlessAppsApmAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmAvg, true
}

// HasServerlessAppsApmAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmAvg() bool {
	return o != nil && o.ServerlessAppsApmAvg != nil
}

// SetServerlessAppsApmAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmAvg(v int64) {
	o.ServerlessAppsApmAvg = &v
}

// GetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg returns the ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg
}

// GetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgOk returns a tuple with the ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg, true
}

// HasServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg != nil
}

// SetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg(v int64) {
	o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg = &v
}

// GetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg returns the ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg
}

// GetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgOk returns a tuple with the ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg, true
}

// HasServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg != nil
}

// SetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg(v int64) {
	o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg = &v
}

// GetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg returns the ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg
}

// GetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgOk returns a tuple with the ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg, true
}

// HasServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg != nil
}

// SetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg(v int64) {
	o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg = &v
}

// GetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg returns the ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg
}

// GetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgOk returns a tuple with the ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg, true
}

// HasServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg != nil
}

// SetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg(v int64) {
	o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg = &v
}

// GetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg returns the ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg
}

// GetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgOk returns a tuple with the ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg, true
}

// HasServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg != nil
}

// SetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg(v int64) {
	o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg = &v
}

// GetServerlessAppsApmExclFargateAvg returns the ServerlessAppsApmExclFargateAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateAvg() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateAvg
}

// GetServerlessAppsApmExclFargateAvgOk returns a tuple with the ServerlessAppsApmExclFargateAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateAvg, true
}

// HasServerlessAppsApmExclFargateAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmExclFargateAvg() bool {
	return o != nil && o.ServerlessAppsApmExclFargateAvg != nil
}

// SetServerlessAppsApmExclFargateAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmExclFargateAvg(v int64) {
	o.ServerlessAppsApmExclFargateAvg = &v
}

// GetServerlessAppsAzureContainerAppInstancesAvg returns the ServerlessAppsAzureContainerAppInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsAzureContainerAppInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsAzureContainerAppInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsAzureContainerAppInstancesAvg
}

// GetServerlessAppsAzureContainerAppInstancesAvgOk returns a tuple with the ServerlessAppsAzureContainerAppInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsAzureContainerAppInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsAzureContainerAppInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsAzureContainerAppInstancesAvg, true
}

// HasServerlessAppsAzureContainerAppInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsAzureContainerAppInstancesAvg() bool {
	return o != nil && o.ServerlessAppsAzureContainerAppInstancesAvg != nil
}

// SetServerlessAppsAzureContainerAppInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsAzureContainerAppInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsAzureContainerAppInstancesAvg(v int64) {
	o.ServerlessAppsAzureContainerAppInstancesAvg = &v
}

// GetServerlessAppsAzureCountAvg returns the ServerlessAppsAzureCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsAzureCountAvg() int64 {
	if o == nil || o.ServerlessAppsAzureCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsAzureCountAvg
}

// GetServerlessAppsAzureCountAvgOk returns a tuple with the ServerlessAppsAzureCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsAzureCountAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsAzureCountAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsAzureCountAvg, true
}

// HasServerlessAppsAzureCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsAzureCountAvg() bool {
	return o != nil && o.ServerlessAppsAzureCountAvg != nil
}

// SetServerlessAppsAzureCountAvg gets a reference to the given int64 and assigns it to the ServerlessAppsAzureCountAvg field.
func (o *UsageSummaryDate) SetServerlessAppsAzureCountAvg(v int64) {
	o.ServerlessAppsAzureCountAvg = &v
}

// GetServerlessAppsAzureFunctionAppInstancesAvg returns the ServerlessAppsAzureFunctionAppInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsAzureFunctionAppInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsAzureFunctionAppInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsAzureFunctionAppInstancesAvg
}

// GetServerlessAppsAzureFunctionAppInstancesAvgOk returns a tuple with the ServerlessAppsAzureFunctionAppInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsAzureFunctionAppInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsAzureFunctionAppInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsAzureFunctionAppInstancesAvg, true
}

// HasServerlessAppsAzureFunctionAppInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsAzureFunctionAppInstancesAvg() bool {
	return o != nil && o.ServerlessAppsAzureFunctionAppInstancesAvg != nil
}

// SetServerlessAppsAzureFunctionAppInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsAzureFunctionAppInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsAzureFunctionAppInstancesAvg(v int64) {
	o.ServerlessAppsAzureFunctionAppInstancesAvg = &v
}

// GetServerlessAppsAzureWebAppInstancesAvg returns the ServerlessAppsAzureWebAppInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsAzureWebAppInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsAzureWebAppInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsAzureWebAppInstancesAvg
}

// GetServerlessAppsAzureWebAppInstancesAvgOk returns a tuple with the ServerlessAppsAzureWebAppInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsAzureWebAppInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsAzureWebAppInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsAzureWebAppInstancesAvg, true
}

// HasServerlessAppsAzureWebAppInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsAzureWebAppInstancesAvg() bool {
	return o != nil && o.ServerlessAppsAzureWebAppInstancesAvg != nil
}

// SetServerlessAppsAzureWebAppInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsAzureWebAppInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsAzureWebAppInstancesAvg(v int64) {
	o.ServerlessAppsAzureWebAppInstancesAvg = &v
}

// GetServerlessAppsEcsAvg returns the ServerlessAppsEcsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsEcsAvg() int64 {
	if o == nil || o.ServerlessAppsEcsAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsEcsAvg
}

// GetServerlessAppsEcsAvgOk returns a tuple with the ServerlessAppsEcsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsEcsAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsEcsAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsEcsAvg, true
}

// HasServerlessAppsEcsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsEcsAvg() bool {
	return o != nil && o.ServerlessAppsEcsAvg != nil
}

// SetServerlessAppsEcsAvg gets a reference to the given int64 and assigns it to the ServerlessAppsEcsAvg field.
func (o *UsageSummaryDate) SetServerlessAppsEcsAvg(v int64) {
	o.ServerlessAppsEcsAvg = &v
}

// GetServerlessAppsEksAvg returns the ServerlessAppsEksAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsEksAvg() int64 {
	if o == nil || o.ServerlessAppsEksAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsEksAvg
}

// GetServerlessAppsEksAvgOk returns a tuple with the ServerlessAppsEksAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsEksAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsEksAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsEksAvg, true
}

// HasServerlessAppsEksAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsEksAvg() bool {
	return o != nil && o.ServerlessAppsEksAvg != nil
}

// SetServerlessAppsEksAvg gets a reference to the given int64 and assigns it to the ServerlessAppsEksAvg field.
func (o *UsageSummaryDate) SetServerlessAppsEksAvg(v int64) {
	o.ServerlessAppsEksAvg = &v
}

// GetServerlessAppsExclFargateAvg returns the ServerlessAppsExclFargateAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateAvg() int64 {
	if o == nil || o.ServerlessAppsExclFargateAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateAvg
}

// GetServerlessAppsExclFargateAvgOk returns a tuple with the ServerlessAppsExclFargateAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateAvg, true
}

// HasServerlessAppsExclFargateAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsExclFargateAvg() bool {
	return o != nil && o.ServerlessAppsExclFargateAvg != nil
}

// SetServerlessAppsExclFargateAvg gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateAvg field.
func (o *UsageSummaryDate) SetServerlessAppsExclFargateAvg(v int64) {
	o.ServerlessAppsExclFargateAvg = &v
}

// GetServerlessAppsExclFargateAzureContainerAppInstancesAvg returns the ServerlessAppsExclFargateAzureContainerAppInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateAzureContainerAppInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg
}

// GetServerlessAppsExclFargateAzureContainerAppInstancesAvgOk returns a tuple with the ServerlessAppsExclFargateAzureContainerAppInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateAzureContainerAppInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg, true
}

// HasServerlessAppsExclFargateAzureContainerAppInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsExclFargateAzureContainerAppInstancesAvg() bool {
	return o != nil && o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg != nil
}

// SetServerlessAppsExclFargateAzureContainerAppInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateAzureContainerAppInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsExclFargateAzureContainerAppInstancesAvg(v int64) {
	o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg = &v
}

// GetServerlessAppsExclFargateAzureFunctionAppInstancesAvg returns the ServerlessAppsExclFargateAzureFunctionAppInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateAzureFunctionAppInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg
}

// GetServerlessAppsExclFargateAzureFunctionAppInstancesAvgOk returns a tuple with the ServerlessAppsExclFargateAzureFunctionAppInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateAzureFunctionAppInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg, true
}

// HasServerlessAppsExclFargateAzureFunctionAppInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsExclFargateAzureFunctionAppInstancesAvg() bool {
	return o != nil && o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg != nil
}

// SetServerlessAppsExclFargateAzureFunctionAppInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateAzureFunctionAppInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsExclFargateAzureFunctionAppInstancesAvg(v int64) {
	o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg = &v
}

// GetServerlessAppsExclFargateAzureWebAppInstancesAvg returns the ServerlessAppsExclFargateAzureWebAppInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateAzureWebAppInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsExclFargateAzureWebAppInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateAzureWebAppInstancesAvg
}

// GetServerlessAppsExclFargateAzureWebAppInstancesAvgOk returns a tuple with the ServerlessAppsExclFargateAzureWebAppInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateAzureWebAppInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateAzureWebAppInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateAzureWebAppInstancesAvg, true
}

// HasServerlessAppsExclFargateAzureWebAppInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsExclFargateAzureWebAppInstancesAvg() bool {
	return o != nil && o.ServerlessAppsExclFargateAzureWebAppInstancesAvg != nil
}

// SetServerlessAppsExclFargateAzureWebAppInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateAzureWebAppInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsExclFargateAzureWebAppInstancesAvg(v int64) {
	o.ServerlessAppsExclFargateAzureWebAppInstancesAvg = &v
}

// GetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg returns the ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg
}

// GetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgOk returns a tuple with the ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg, true
}

// HasServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg() bool {
	return o != nil && o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg != nil
}

// SetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg(v int64) {
	o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg = &v
}

// GetServerlessAppsExclFargateGoogleCloudRunInstancesAvg returns the ServerlessAppsExclFargateGoogleCloudRunInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateGoogleCloudRunInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg
}

// GetServerlessAppsExclFargateGoogleCloudRunInstancesAvgOk returns a tuple with the ServerlessAppsExclFargateGoogleCloudRunInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateGoogleCloudRunInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg, true
}

// HasServerlessAppsExclFargateGoogleCloudRunInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsExclFargateGoogleCloudRunInstancesAvg() bool {
	return o != nil && o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg != nil
}

// SetServerlessAppsExclFargateGoogleCloudRunInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateGoogleCloudRunInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsExclFargateGoogleCloudRunInstancesAvg(v int64) {
	o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg = &v
}

// GetServerlessAppsGoogleCloudFunctionsInstancesAvg returns the ServerlessAppsGoogleCloudFunctionsInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsGoogleCloudFunctionsInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsGoogleCloudFunctionsInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsGoogleCloudFunctionsInstancesAvg
}

// GetServerlessAppsGoogleCloudFunctionsInstancesAvgOk returns a tuple with the ServerlessAppsGoogleCloudFunctionsInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsGoogleCloudFunctionsInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsGoogleCloudFunctionsInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsGoogleCloudFunctionsInstancesAvg, true
}

// HasServerlessAppsGoogleCloudFunctionsInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsGoogleCloudFunctionsInstancesAvg() bool {
	return o != nil && o.ServerlessAppsGoogleCloudFunctionsInstancesAvg != nil
}

// SetServerlessAppsGoogleCloudFunctionsInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsGoogleCloudFunctionsInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsGoogleCloudFunctionsInstancesAvg(v int64) {
	o.ServerlessAppsGoogleCloudFunctionsInstancesAvg = &v
}

// GetServerlessAppsGoogleCloudRunInstancesAvg returns the ServerlessAppsGoogleCloudRunInstancesAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsGoogleCloudRunInstancesAvg() int64 {
	if o == nil || o.ServerlessAppsGoogleCloudRunInstancesAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsGoogleCloudRunInstancesAvg
}

// GetServerlessAppsGoogleCloudRunInstancesAvgOk returns a tuple with the ServerlessAppsGoogleCloudRunInstancesAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsGoogleCloudRunInstancesAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsGoogleCloudRunInstancesAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsGoogleCloudRunInstancesAvg, true
}

// HasServerlessAppsGoogleCloudRunInstancesAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsGoogleCloudRunInstancesAvg() bool {
	return o != nil && o.ServerlessAppsGoogleCloudRunInstancesAvg != nil
}

// SetServerlessAppsGoogleCloudRunInstancesAvg gets a reference to the given int64 and assigns it to the ServerlessAppsGoogleCloudRunInstancesAvg field.
func (o *UsageSummaryDate) SetServerlessAppsGoogleCloudRunInstancesAvg(v int64) {
	o.ServerlessAppsGoogleCloudRunInstancesAvg = &v
}

// GetServerlessAppsGoogleCountAvg returns the ServerlessAppsGoogleCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsGoogleCountAvg() int64 {
	if o == nil || o.ServerlessAppsGoogleCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsGoogleCountAvg
}

// GetServerlessAppsGoogleCountAvgOk returns a tuple with the ServerlessAppsGoogleCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsGoogleCountAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsGoogleCountAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsGoogleCountAvg, true
}

// HasServerlessAppsGoogleCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsGoogleCountAvg() bool {
	return o != nil && o.ServerlessAppsGoogleCountAvg != nil
}

// SetServerlessAppsGoogleCountAvg gets a reference to the given int64 and assigns it to the ServerlessAppsGoogleCountAvg field.
func (o *UsageSummaryDate) SetServerlessAppsGoogleCountAvg(v int64) {
	o.ServerlessAppsGoogleCountAvg = &v
}

// GetServerlessAppsTotalCountAvg returns the ServerlessAppsTotalCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsTotalCountAvg() int64 {
	if o == nil || o.ServerlessAppsTotalCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsTotalCountAvg
}

// GetServerlessAppsTotalCountAvgOk returns a tuple with the ServerlessAppsTotalCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsTotalCountAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsTotalCountAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsTotalCountAvg, true
}

// HasServerlessAppsTotalCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsTotalCountAvg() bool {
	return o != nil && o.ServerlessAppsTotalCountAvg != nil
}

// SetServerlessAppsTotalCountAvg gets a reference to the given int64 and assigns it to the ServerlessAppsTotalCountAvg field.
func (o *UsageSummaryDate) SetServerlessAppsTotalCountAvg(v int64) {
	o.ServerlessAppsTotalCountAvg = &v
}

// GetSiemAnalyzedLogsAddOnCountSum returns the SiemAnalyzedLogsAddOnCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSiemAnalyzedLogsAddOnCountSum() int64 {
	if o == nil || o.SiemAnalyzedLogsAddOnCountSum == nil {
		var ret int64
		return ret
	}
	return *o.SiemAnalyzedLogsAddOnCountSum
}

// GetSiemAnalyzedLogsAddOnCountSumOk returns a tuple with the SiemAnalyzedLogsAddOnCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSiemAnalyzedLogsAddOnCountSumOk() (*int64, bool) {
	if o == nil || o.SiemAnalyzedLogsAddOnCountSum == nil {
		return nil, false
	}
	return o.SiemAnalyzedLogsAddOnCountSum, true
}

// HasSiemAnalyzedLogsAddOnCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSiemAnalyzedLogsAddOnCountSum() bool {
	return o != nil && o.SiemAnalyzedLogsAddOnCountSum != nil
}

// SetSiemAnalyzedLogsAddOnCountSum gets a reference to the given int64 and assigns it to the SiemAnalyzedLogsAddOnCountSum field.
func (o *UsageSummaryDate) SetSiemAnalyzedLogsAddOnCountSum(v int64) {
	o.SiemAnalyzedLogsAddOnCountSum = &v
}

// GetSyntheticsBrowserCheckCallsCountSum returns the SyntheticsBrowserCheckCallsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSum() int64 {
	if o == nil || o.SyntheticsBrowserCheckCallsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsBrowserCheckCallsCountSum
}

// GetSyntheticsBrowserCheckCallsCountSumOk returns a tuple with the SyntheticsBrowserCheckCallsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSyntheticsBrowserCheckCallsCountSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsBrowserCheckCallsCountSum == nil {
		return nil, false
	}
	return o.SyntheticsBrowserCheckCallsCountSum, true
}

// HasSyntheticsBrowserCheckCallsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSyntheticsBrowserCheckCallsCountSum() bool {
	return o != nil && o.SyntheticsBrowserCheckCallsCountSum != nil
}

// SetSyntheticsBrowserCheckCallsCountSum gets a reference to the given int64 and assigns it to the SyntheticsBrowserCheckCallsCountSum field.
func (o *UsageSummaryDate) SetSyntheticsBrowserCheckCallsCountSum(v int64) {
	o.SyntheticsBrowserCheckCallsCountSum = &v
}

// GetSyntheticsCheckCallsCountSum returns the SyntheticsCheckCallsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSum() int64 {
	if o == nil || o.SyntheticsCheckCallsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsCheckCallsCountSum
}

// GetSyntheticsCheckCallsCountSumOk returns a tuple with the SyntheticsCheckCallsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSyntheticsCheckCallsCountSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsCheckCallsCountSum == nil {
		return nil, false
	}
	return o.SyntheticsCheckCallsCountSum, true
}

// HasSyntheticsCheckCallsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSyntheticsCheckCallsCountSum() bool {
	return o != nil && o.SyntheticsCheckCallsCountSum != nil
}

// SetSyntheticsCheckCallsCountSum gets a reference to the given int64 and assigns it to the SyntheticsCheckCallsCountSum field.
func (o *UsageSummaryDate) SetSyntheticsCheckCallsCountSum(v int64) {
	o.SyntheticsCheckCallsCountSum = &v
}

// GetSyntheticsMobileTestRunsSum returns the SyntheticsMobileTestRunsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSyntheticsMobileTestRunsSum() int64 {
	if o == nil || o.SyntheticsMobileTestRunsSum == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsMobileTestRunsSum
}

// GetSyntheticsMobileTestRunsSumOk returns a tuple with the SyntheticsMobileTestRunsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSyntheticsMobileTestRunsSumOk() (*int64, bool) {
	if o == nil || o.SyntheticsMobileTestRunsSum == nil {
		return nil, false
	}
	return o.SyntheticsMobileTestRunsSum, true
}

// HasSyntheticsMobileTestRunsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSyntheticsMobileTestRunsSum() bool {
	return o != nil && o.SyntheticsMobileTestRunsSum != nil
}

// SetSyntheticsMobileTestRunsSum gets a reference to the given int64 and assigns it to the SyntheticsMobileTestRunsSum field.
func (o *UsageSummaryDate) SetSyntheticsMobileTestRunsSum(v int64) {
	o.SyntheticsMobileTestRunsSum = &v
}

// GetSyntheticsParallelTestingMaxSlotsHwm returns the SyntheticsParallelTestingMaxSlotsHwm field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSyntheticsParallelTestingMaxSlotsHwm() int64 {
	if o == nil || o.SyntheticsParallelTestingMaxSlotsHwm == nil {
		var ret int64
		return ret
	}
	return *o.SyntheticsParallelTestingMaxSlotsHwm
}

// GetSyntheticsParallelTestingMaxSlotsHwmOk returns a tuple with the SyntheticsParallelTestingMaxSlotsHwm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSyntheticsParallelTestingMaxSlotsHwmOk() (*int64, bool) {
	if o == nil || o.SyntheticsParallelTestingMaxSlotsHwm == nil {
		return nil, false
	}
	return o.SyntheticsParallelTestingMaxSlotsHwm, true
}

// HasSyntheticsParallelTestingMaxSlotsHwm returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSyntheticsParallelTestingMaxSlotsHwm() bool {
	return o != nil && o.SyntheticsParallelTestingMaxSlotsHwm != nil
}

// SetSyntheticsParallelTestingMaxSlotsHwm gets a reference to the given int64 and assigns it to the SyntheticsParallelTestingMaxSlotsHwm field.
func (o *UsageSummaryDate) SetSyntheticsParallelTestingMaxSlotsHwm(v int64) {
	o.SyntheticsParallelTestingMaxSlotsHwm = &v
}

// GetTraceSearchIndexedEventsCountSum returns the TraceSearchIndexedEventsCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSum() int64 {
	if o == nil || o.TraceSearchIndexedEventsCountSum == nil {
		var ret int64
		return ret
	}
	return *o.TraceSearchIndexedEventsCountSum
}

// GetTraceSearchIndexedEventsCountSumOk returns a tuple with the TraceSearchIndexedEventsCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetTraceSearchIndexedEventsCountSumOk() (*int64, bool) {
	if o == nil || o.TraceSearchIndexedEventsCountSum == nil {
		return nil, false
	}
	return o.TraceSearchIndexedEventsCountSum, true
}

// HasTraceSearchIndexedEventsCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasTraceSearchIndexedEventsCountSum() bool {
	return o != nil && o.TraceSearchIndexedEventsCountSum != nil
}

// SetTraceSearchIndexedEventsCountSum gets a reference to the given int64 and assigns it to the TraceSearchIndexedEventsCountSum field.
func (o *UsageSummaryDate) SetTraceSearchIndexedEventsCountSum(v int64) {
	o.TraceSearchIndexedEventsCountSum = &v
}

// GetTwolIngestedEventsBytesSum returns the TwolIngestedEventsBytesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetTwolIngestedEventsBytesSum() int64 {
	if o == nil || o.TwolIngestedEventsBytesSum == nil {
		var ret int64
		return ret
	}
	return *o.TwolIngestedEventsBytesSum
}

// GetTwolIngestedEventsBytesSumOk returns a tuple with the TwolIngestedEventsBytesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetTwolIngestedEventsBytesSumOk() (*int64, bool) {
	if o == nil || o.TwolIngestedEventsBytesSum == nil {
		return nil, false
	}
	return o.TwolIngestedEventsBytesSum, true
}

// HasTwolIngestedEventsBytesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasTwolIngestedEventsBytesSum() bool {
	return o != nil && o.TwolIngestedEventsBytesSum != nil
}

// SetTwolIngestedEventsBytesSum gets a reference to the given int64 and assigns it to the TwolIngestedEventsBytesSum field.
func (o *UsageSummaryDate) SetTwolIngestedEventsBytesSum(v int64) {
	o.TwolIngestedEventsBytesSum = &v
}

// GetUniversalServiceMonitoringHostTop99p returns the UniversalServiceMonitoringHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetUniversalServiceMonitoringHostTop99p() int64 {
	if o == nil || o.UniversalServiceMonitoringHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.UniversalServiceMonitoringHostTop99p
}

// GetUniversalServiceMonitoringHostTop99pOk returns a tuple with the UniversalServiceMonitoringHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetUniversalServiceMonitoringHostTop99pOk() (*int64, bool) {
	if o == nil || o.UniversalServiceMonitoringHostTop99p == nil {
		return nil, false
	}
	return o.UniversalServiceMonitoringHostTop99p, true
}

// HasUniversalServiceMonitoringHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasUniversalServiceMonitoringHostTop99p() bool {
	return o != nil && o.UniversalServiceMonitoringHostTop99p != nil
}

// SetUniversalServiceMonitoringHostTop99p gets a reference to the given int64 and assigns it to the UniversalServiceMonitoringHostTop99p field.
func (o *UsageSummaryDate) SetUniversalServiceMonitoringHostTop99p(v int64) {
	o.UniversalServiceMonitoringHostTop99p = &v
}

// GetVsphereHostTop99p returns the VsphereHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetVsphereHostTop99p() int64 {
	if o == nil || o.VsphereHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.VsphereHostTop99p
}

// GetVsphereHostTop99pOk returns a tuple with the VsphereHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetVsphereHostTop99pOk() (*int64, bool) {
	if o == nil || o.VsphereHostTop99p == nil {
		return nil, false
	}
	return o.VsphereHostTop99p, true
}

// HasVsphereHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasVsphereHostTop99p() bool {
	return o != nil && o.VsphereHostTop99p != nil
}

// SetVsphereHostTop99p gets a reference to the given int64 and assigns it to the VsphereHostTop99p field.
func (o *UsageSummaryDate) SetVsphereHostTop99p(v int64) {
	o.VsphereHostTop99p = &v
}

// GetVulnManagementHostCountTop99p returns the VulnManagementHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetVulnManagementHostCountTop99p() int64 {
	if o == nil || o.VulnManagementHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.VulnManagementHostCountTop99p
}

// GetVulnManagementHostCountTop99pOk returns a tuple with the VulnManagementHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetVulnManagementHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.VulnManagementHostCountTop99p == nil {
		return nil, false
	}
	return o.VulnManagementHostCountTop99p, true
}

// HasVulnManagementHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasVulnManagementHostCountTop99p() bool {
	return o != nil && o.VulnManagementHostCountTop99p != nil
}

// SetVulnManagementHostCountTop99p gets a reference to the given int64 and assigns it to the VulnManagementHostCountTop99p field.
func (o *UsageSummaryDate) SetVulnManagementHostCountTop99p(v int64) {
	o.VulnManagementHostCountTop99p = &v
}

// GetWorkflowExecutionsUsageSum returns the WorkflowExecutionsUsageSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetWorkflowExecutionsUsageSum() int64 {
	if o == nil || o.WorkflowExecutionsUsageSum == nil {
		var ret int64
		return ret
	}
	return *o.WorkflowExecutionsUsageSum
}

// GetWorkflowExecutionsUsageSumOk returns a tuple with the WorkflowExecutionsUsageSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetWorkflowExecutionsUsageSumOk() (*int64, bool) {
	if o == nil || o.WorkflowExecutionsUsageSum == nil {
		return nil, false
	}
	return o.WorkflowExecutionsUsageSum, true
}

// HasWorkflowExecutionsUsageSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasWorkflowExecutionsUsageSum() bool {
	return o != nil && o.WorkflowExecutionsUsageSum != nil
}

// SetWorkflowExecutionsUsageSum gets a reference to the given int64 and assigns it to the WorkflowExecutionsUsageSum field.
func (o *UsageSummaryDate) SetWorkflowExecutionsUsageSum(v int64) {
	o.WorkflowExecutionsUsageSum = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSummaryDate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentHostTop99p != nil {
		toSerialize["agent_host_top99p"] = o.AgentHostTop99p
	}
	if o.ApmAzureAppServiceHostTop99p != nil {
		toSerialize["apm_azure_app_service_host_top99p"] = o.ApmAzureAppServiceHostTop99p
	}
	if o.ApmDevsecopsHostTop99p != nil {
		toSerialize["apm_devsecops_host_top99p"] = o.ApmDevsecopsHostTop99p
	}
	if o.ApmEnterpriseStandaloneHostsTop99p != nil {
		toSerialize["apm_enterprise_standalone_hosts_top99p"] = o.ApmEnterpriseStandaloneHostsTop99p
	}
	if o.ApmFargateCountAvg != nil {
		toSerialize["apm_fargate_count_avg"] = o.ApmFargateCountAvg
	}
	if o.ApmHostTop99p != nil {
		toSerialize["apm_host_top99p"] = o.ApmHostTop99p
	}
	if o.ApmProStandaloneHostsTop99p != nil {
		toSerialize["apm_pro_standalone_hosts_top99p"] = o.ApmProStandaloneHostsTop99p
	}
	if o.AppsecFargateCountAvg != nil {
		toSerialize["appsec_fargate_count_avg"] = o.AppsecFargateCountAvg
	}
	if o.AsmServerlessSum != nil {
		toSerialize["asm_serverless_sum"] = o.AsmServerlessSum
	}
	if o.AuditLogsLinesIndexedSum != nil {
		toSerialize["audit_logs_lines_indexed_sum"] = o.AuditLogsLinesIndexedSum
	}
	if o.AuditTrailEnabledHwm != nil {
		toSerialize["audit_trail_enabled_hwm"] = o.AuditTrailEnabledHwm
	}
	if o.AvgProfiledFargateTasks != nil {
		toSerialize["avg_profiled_fargate_tasks"] = o.AvgProfiledFargateTasks
	}
	if o.AwsHostTop99p != nil {
		toSerialize["aws_host_top99p"] = o.AwsHostTop99p
	}
	if o.AwsLambdaFuncCount != nil {
		toSerialize["aws_lambda_func_count"] = o.AwsLambdaFuncCount
	}
	if o.AwsLambdaInvocationsSum != nil {
		toSerialize["aws_lambda_invocations_sum"] = o.AwsLambdaInvocationsSum
	}
	if o.AzureAppServiceTop99p != nil {
		toSerialize["azure_app_service_top99p"] = o.AzureAppServiceTop99p
	}
	if o.BillableIngestedBytesSum != nil {
		toSerialize["billable_ingested_bytes_sum"] = o.BillableIngestedBytesSum
	}
	if o.BitsAiInvestigationsSum != nil {
		toSerialize["bits_ai_investigations_sum"] = o.BitsAiInvestigationsSum
	}
	if o.BrowserRumLiteSessionCountSum != nil {
		toSerialize["browser_rum_lite_session_count_sum"] = o.BrowserRumLiteSessionCountSum
	}
	if o.BrowserRumReplaySessionCountSum != nil {
		toSerialize["browser_rum_replay_session_count_sum"] = o.BrowserRumReplaySessionCountSum
	}
	if o.BrowserRumUnitsSum != nil {
		toSerialize["browser_rum_units_sum"] = o.BrowserRumUnitsSum
	}
	if o.CiPipelineIndexedSpansSum != nil {
		toSerialize["ci_pipeline_indexed_spans_sum"] = o.CiPipelineIndexedSpansSum
	}
	if o.CiTestIndexedSpansSum != nil {
		toSerialize["ci_test_indexed_spans_sum"] = o.CiTestIndexedSpansSum
	}
	if o.CiVisibilityItrCommittersHwm != nil {
		toSerialize["ci_visibility_itr_committers_hwm"] = o.CiVisibilityItrCommittersHwm
	}
	if o.CiVisibilityPipelineCommittersHwm != nil {
		toSerialize["ci_visibility_pipeline_committers_hwm"] = o.CiVisibilityPipelineCommittersHwm
	}
	if o.CiVisibilityTestCommittersHwm != nil {
		toSerialize["ci_visibility_test_committers_hwm"] = o.CiVisibilityTestCommittersHwm
	}
	if o.CloudCostManagementAwsHostCountAvg != nil {
		toSerialize["cloud_cost_management_aws_host_count_avg"] = o.CloudCostManagementAwsHostCountAvg
	}
	if o.CloudCostManagementAzureHostCountAvg != nil {
		toSerialize["cloud_cost_management_azure_host_count_avg"] = o.CloudCostManagementAzureHostCountAvg
	}
	if o.CloudCostManagementGcpHostCountAvg != nil {
		toSerialize["cloud_cost_management_gcp_host_count_avg"] = o.CloudCostManagementGcpHostCountAvg
	}
	if o.CloudCostManagementHostCountAvg != nil {
		toSerialize["cloud_cost_management_host_count_avg"] = o.CloudCostManagementHostCountAvg
	}
	if o.CloudCostManagementOciHostCountAvg != nil {
		toSerialize["cloud_cost_management_oci_host_count_avg"] = o.CloudCostManagementOciHostCountAvg
	}
	if o.CloudSiemEventsSum != nil {
		toSerialize["cloud_siem_events_sum"] = o.CloudSiemEventsSum
	}
	if o.CodeAnalysisSaCommittersHwm != nil {
		toSerialize["code_analysis_sa_committers_hwm"] = o.CodeAnalysisSaCommittersHwm
	}
	if o.CodeAnalysisScaCommittersHwm != nil {
		toSerialize["code_analysis_sca_committers_hwm"] = o.CodeAnalysisScaCommittersHwm
	}
	if o.CodeSecurityHostTop99p != nil {
		toSerialize["code_security_host_top99p"] = o.CodeSecurityHostTop99p
	}
	if o.ContainerAvg != nil {
		toSerialize["container_avg"] = o.ContainerAvg
	}
	if o.ContainerExclAgentAvg != nil {
		toSerialize["container_excl_agent_avg"] = o.ContainerExclAgentAvg
	}
	if o.ContainerHwm != nil {
		toSerialize["container_hwm"] = o.ContainerHwm
	}
	if o.CsmContainerEnterpriseComplianceCountSum != nil {
		toSerialize["csm_container_enterprise_compliance_count_sum"] = o.CsmContainerEnterpriseComplianceCountSum
	}
	if o.CsmContainerEnterpriseCwsCountSum != nil {
		toSerialize["csm_container_enterprise_cws_count_sum"] = o.CsmContainerEnterpriseCwsCountSum
	}
	if o.CsmContainerEnterpriseTotalCountSum != nil {
		toSerialize["csm_container_enterprise_total_count_sum"] = o.CsmContainerEnterpriseTotalCountSum
	}
	if o.CsmHostEnterpriseAasHostCountTop99p != nil {
		toSerialize["csm_host_enterprise_aas_host_count_top99p"] = o.CsmHostEnterpriseAasHostCountTop99p
	}
	if o.CsmHostEnterpriseAwsHostCountTop99p != nil {
		toSerialize["csm_host_enterprise_aws_host_count_top99p"] = o.CsmHostEnterpriseAwsHostCountTop99p
	}
	if o.CsmHostEnterpriseAzureHostCountTop99p != nil {
		toSerialize["csm_host_enterprise_azure_host_count_top99p"] = o.CsmHostEnterpriseAzureHostCountTop99p
	}
	if o.CsmHostEnterpriseComplianceHostCountTop99p != nil {
		toSerialize["csm_host_enterprise_compliance_host_count_top99p"] = o.CsmHostEnterpriseComplianceHostCountTop99p
	}
	if o.CsmHostEnterpriseCwsHostCountTop99p != nil {
		toSerialize["csm_host_enterprise_cws_host_count_top99p"] = o.CsmHostEnterpriseCwsHostCountTop99p
	}
	if o.CsmHostEnterpriseGcpHostCountTop99p != nil {
		toSerialize["csm_host_enterprise_gcp_host_count_top99p"] = o.CsmHostEnterpriseGcpHostCountTop99p
	}
	if o.CsmHostEnterpriseTotalHostCountTop99p != nil {
		toSerialize["csm_host_enterprise_total_host_count_top99p"] = o.CsmHostEnterpriseTotalHostCountTop99p
	}
	if o.CspmAasHostTop99p != nil {
		toSerialize["cspm_aas_host_top99p"] = o.CspmAasHostTop99p
	}
	if o.CspmAwsHostTop99p != nil {
		toSerialize["cspm_aws_host_top99p"] = o.CspmAwsHostTop99p
	}
	if o.CspmAzureHostTop99p != nil {
		toSerialize["cspm_azure_host_top99p"] = o.CspmAzureHostTop99p
	}
	if o.CspmContainerAvg != nil {
		toSerialize["cspm_container_avg"] = o.CspmContainerAvg
	}
	if o.CspmContainerHwm != nil {
		toSerialize["cspm_container_hwm"] = o.CspmContainerHwm
	}
	if o.CspmGcpHostTop99p != nil {
		toSerialize["cspm_gcp_host_top99p"] = o.CspmGcpHostTop99p
	}
	if o.CspmHostTop99p != nil {
		toSerialize["cspm_host_top99p"] = o.CspmHostTop99p
	}
	if o.CustomTsAvg != nil {
		toSerialize["custom_ts_avg"] = o.CustomTsAvg
	}
	if o.CwsContainerCountAvg != nil {
		toSerialize["cws_container_count_avg"] = o.CwsContainerCountAvg
	}
	if o.CwsFargateTaskAvg != nil {
		toSerialize["cws_fargate_task_avg"] = o.CwsFargateTaskAvg
	}
	if o.CwsHostTop99p != nil {
		toSerialize["cws_host_top99p"] = o.CwsHostTop99p
	}
	if o.DataJobsMonitoringHostHrSum != nil {
		toSerialize["data_jobs_monitoring_host_hr_sum"] = o.DataJobsMonitoringHostHrSum
	}
	if o.Date != nil {
		if o.Date.Nanosecond() == 0 {
			toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DbmHostTop99p != nil {
		toSerialize["dbm_host_top99p"] = o.DbmHostTop99p
	}
	if o.DbmQueriesCountAvg != nil {
		toSerialize["dbm_queries_count_avg"] = o.DbmQueriesCountAvg
	}
	if o.EphInfraHostAgentSum != nil {
		toSerialize["eph_infra_host_agent_sum"] = o.EphInfraHostAgentSum
	}
	if o.EphInfraHostAlibabaSum != nil {
		toSerialize["eph_infra_host_alibaba_sum"] = o.EphInfraHostAlibabaSum
	}
	if o.EphInfraHostAwsSum != nil {
		toSerialize["eph_infra_host_aws_sum"] = o.EphInfraHostAwsSum
	}
	if o.EphInfraHostAzureSum != nil {
		toSerialize["eph_infra_host_azure_sum"] = o.EphInfraHostAzureSum
	}
	if o.EphInfraHostEntSum != nil {
		toSerialize["eph_infra_host_ent_sum"] = o.EphInfraHostEntSum
	}
	if o.EphInfraHostGcpSum != nil {
		toSerialize["eph_infra_host_gcp_sum"] = o.EphInfraHostGcpSum
	}
	if o.EphInfraHostHerokuSum != nil {
		toSerialize["eph_infra_host_heroku_sum"] = o.EphInfraHostHerokuSum
	}
	if o.EphInfraHostOnlyAasSum != nil {
		toSerialize["eph_infra_host_only_aas_sum"] = o.EphInfraHostOnlyAasSum
	}
	if o.EphInfraHostOnlyVsphereSum != nil {
		toSerialize["eph_infra_host_only_vsphere_sum"] = o.EphInfraHostOnlyVsphereSum
	}
	if o.EphInfraHostOpentelemetryApmSum != nil {
		toSerialize["eph_infra_host_opentelemetry_apm_sum"] = o.EphInfraHostOpentelemetryApmSum
	}
	if o.EphInfraHostOpentelemetrySum != nil {
		toSerialize["eph_infra_host_opentelemetry_sum"] = o.EphInfraHostOpentelemetrySum
	}
	if o.EphInfraHostProSum != nil {
		toSerialize["eph_infra_host_pro_sum"] = o.EphInfraHostProSum
	}
	if o.EphInfraHostProplusSum != nil {
		toSerialize["eph_infra_host_proplus_sum"] = o.EphInfraHostProplusSum
	}
	if o.EphInfraHostProxmoxSum != nil {
		toSerialize["eph_infra_host_proxmox_sum"] = o.EphInfraHostProxmoxSum
	}
	if o.ErrorTrackingApmErrorEventsSum != nil {
		toSerialize["error_tracking_apm_error_events_sum"] = o.ErrorTrackingApmErrorEventsSum
	}
	if o.ErrorTrackingErrorEventsSum != nil {
		toSerialize["error_tracking_error_events_sum"] = o.ErrorTrackingErrorEventsSum
	}
	if o.ErrorTrackingEventsSum != nil {
		toSerialize["error_tracking_events_sum"] = o.ErrorTrackingEventsSum
	}
	if o.ErrorTrackingRumErrorEventsSum != nil {
		toSerialize["error_tracking_rum_error_events_sum"] = o.ErrorTrackingRumErrorEventsSum
	}
	if o.EventManagementCorrelationCorrelatedEventsSum != nil {
		toSerialize["event_management_correlation_correlated_events_sum"] = o.EventManagementCorrelationCorrelatedEventsSum
	}
	if o.EventManagementCorrelationCorrelatedRelatedEventsSum != nil {
		toSerialize["event_management_correlation_correlated_related_events_sum"] = o.EventManagementCorrelationCorrelatedRelatedEventsSum
	}
	if o.EventManagementCorrelationSum != nil {
		toSerialize["event_management_correlation_sum"] = o.EventManagementCorrelationSum
	}
	if o.FargateContainerProfilerProfilingFargateAvg != nil {
		toSerialize["fargate_container_profiler_profiling_fargate_avg"] = o.FargateContainerProfilerProfilingFargateAvg
	}
	if o.FargateContainerProfilerProfilingFargateEksAvg != nil {
		toSerialize["fargate_container_profiler_profiling_fargate_eks_avg"] = o.FargateContainerProfilerProfilingFargateEksAvg
	}
	if o.FargateTasksCountAvg != nil {
		toSerialize["fargate_tasks_count_avg"] = o.FargateTasksCountAvg
	}
	if o.FargateTasksCountHwm != nil {
		toSerialize["fargate_tasks_count_hwm"] = o.FargateTasksCountHwm
	}
	if o.FlexLogsComputeLargeAvg != nil {
		toSerialize["flex_logs_compute_large_avg"] = o.FlexLogsComputeLargeAvg
	}
	if o.FlexLogsComputeMediumAvg != nil {
		toSerialize["flex_logs_compute_medium_avg"] = o.FlexLogsComputeMediumAvg
	}
	if o.FlexLogsComputeSmallAvg != nil {
		toSerialize["flex_logs_compute_small_avg"] = o.FlexLogsComputeSmallAvg
	}
	if o.FlexLogsComputeXlargeAvg != nil {
		toSerialize["flex_logs_compute_xlarge_avg"] = o.FlexLogsComputeXlargeAvg
	}
	if o.FlexLogsComputeXsmallAvg != nil {
		toSerialize["flex_logs_compute_xsmall_avg"] = o.FlexLogsComputeXsmallAvg
	}
	if o.FlexLogsStarterAvg != nil {
		toSerialize["flex_logs_starter_avg"] = o.FlexLogsStarterAvg
	}
	if o.FlexLogsStarterStorageIndexAvg != nil {
		toSerialize["flex_logs_starter_storage_index_avg"] = o.FlexLogsStarterStorageIndexAvg
	}
	if o.FlexLogsStarterStorageRetentionAdjustmentAvg != nil {
		toSerialize["flex_logs_starter_storage_retention_adjustment_avg"] = o.FlexLogsStarterStorageRetentionAdjustmentAvg
	}
	if o.FlexStoredLogsAvg != nil {
		toSerialize["flex_stored_logs_avg"] = o.FlexStoredLogsAvg
	}
	if o.ForwardingEventsBytesSum != nil {
		toSerialize["forwarding_events_bytes_sum"] = o.ForwardingEventsBytesSum
	}
	if o.GcpHostTop99p != nil {
		toSerialize["gcp_host_top99p"] = o.GcpHostTop99p
	}
	if o.HerokuHostTop99p != nil {
		toSerialize["heroku_host_top99p"] = o.HerokuHostTop99p
	}
	if o.IncidentManagementMonthlyActiveUsersHwm != nil {
		toSerialize["incident_management_monthly_active_users_hwm"] = o.IncidentManagementMonthlyActiveUsersHwm
	}
	if o.IncidentManagementSeatsHwm != nil {
		toSerialize["incident_management_seats_hwm"] = o.IncidentManagementSeatsHwm
	}
	if o.IndexedEventsCountSum != nil {
		toSerialize["indexed_events_count_sum"] = o.IndexedEventsCountSum
	}
	if o.InfraHostTop99p != nil {
		toSerialize["infra_host_top99p"] = o.InfraHostTop99p
	}
	if o.IngestedEventsBytesSum != nil {
		toSerialize["ingested_events_bytes_sum"] = o.IngestedEventsBytesSum
	}
	if o.IotDeviceSum != nil {
		toSerialize["iot_device_sum"] = o.IotDeviceSum
	}
	if o.IotDeviceTop99p != nil {
		toSerialize["iot_device_top99p"] = o.IotDeviceTop99p
	}
	if o.LlmObservabilityMinSpendSum != nil {
		toSerialize["llm_observability_min_spend_sum"] = o.LlmObservabilityMinSpendSum
	}
	if o.LlmObservabilitySum != nil {
		toSerialize["llm_observability_sum"] = o.LlmObservabilitySum
	}
	if o.MobileRumLiteSessionCountSum != nil {
		toSerialize["mobile_rum_lite_session_count_sum"] = o.MobileRumLiteSessionCountSum
	}
	if o.MobileRumSessionCountAndroidSum != nil {
		toSerialize["mobile_rum_session_count_android_sum"] = o.MobileRumSessionCountAndroidSum
	}
	if o.MobileRumSessionCountFlutterSum != nil {
		toSerialize["mobile_rum_session_count_flutter_sum"] = o.MobileRumSessionCountFlutterSum
	}
	if o.MobileRumSessionCountIosSum != nil {
		toSerialize["mobile_rum_session_count_ios_sum"] = o.MobileRumSessionCountIosSum
	}
	if o.MobileRumSessionCountReactnativeSum != nil {
		toSerialize["mobile_rum_session_count_reactnative_sum"] = o.MobileRumSessionCountReactnativeSum
	}
	if o.MobileRumSessionCountRokuSum != nil {
		toSerialize["mobile_rum_session_count_roku_sum"] = o.MobileRumSessionCountRokuSum
	}
	if o.MobileRumSessionCountSum != nil {
		toSerialize["mobile_rum_session_count_sum"] = o.MobileRumSessionCountSum
	}
	if o.MobileRumUnitsSum != nil {
		toSerialize["mobile_rum_units_sum"] = o.MobileRumUnitsSum
	}
	if o.NdmNetflowEventsSum != nil {
		toSerialize["ndm_netflow_events_sum"] = o.NdmNetflowEventsSum
	}
	if o.NetflowIndexedEventsCountSum != nil {
		toSerialize["netflow_indexed_events_count_sum"] = o.NetflowIndexedEventsCountSum
	}
	if o.NetworkDeviceWirelessTop99p != nil {
		toSerialize["network_device_wireless_top99p"] = o.NetworkDeviceWirelessTop99p
	}
	if o.NpmHostTop99p != nil {
		toSerialize["npm_host_top99p"] = o.NpmHostTop99p
	}
	if o.ObservabilityPipelinesBytesProcessedSum != nil {
		toSerialize["observability_pipelines_bytes_processed_sum"] = o.ObservabilityPipelinesBytesProcessedSum
	}
	if o.OciHostSum != nil {
		toSerialize["oci_host_sum"] = o.OciHostSum
	}
	if o.OciHostTop99p != nil {
		toSerialize["oci_host_top99p"] = o.OciHostTop99p
	}
	if o.OnCallSeatHwm != nil {
		toSerialize["on_call_seat_hwm"] = o.OnCallSeatHwm
	}
	if o.OnlineArchiveEventsCountSum != nil {
		toSerialize["online_archive_events_count_sum"] = o.OnlineArchiveEventsCountSum
	}
	if o.OpentelemetryApmHostTop99p != nil {
		toSerialize["opentelemetry_apm_host_top99p"] = o.OpentelemetryApmHostTop99p
	}
	if o.OpentelemetryHostTop99p != nil {
		toSerialize["opentelemetry_host_top99p"] = o.OpentelemetryHostTop99p
	}
	if o.Orgs != nil {
		toSerialize["orgs"] = o.Orgs
	}
	if o.ProductAnalyticsSum != nil {
		toSerialize["product_analytics_sum"] = o.ProductAnalyticsSum
	}
	if o.ProfilingAasCountTop99p != nil {
		toSerialize["profiling_aas_count_top99p"] = o.ProfilingAasCountTop99p
	}
	if o.ProfilingHostTop99p != nil {
		toSerialize["profiling_host_top99p"] = o.ProfilingHostTop99p
	}
	if o.ProxmoxHostSum != nil {
		toSerialize["proxmox_host_sum"] = o.ProxmoxHostSum
	}
	if o.ProxmoxHostTop99p != nil {
		toSerialize["proxmox_host_top99p"] = o.ProxmoxHostTop99p
	}
	if o.PublishedAppHwm != nil {
		toSerialize["published_app_hwm"] = o.PublishedAppHwm
	}
	if o.RumBrowserAndMobileSessionCount != nil {
		toSerialize["rum_browser_and_mobile_session_count"] = o.RumBrowserAndMobileSessionCount
	}
	if o.RumBrowserLegacySessionCountSum != nil {
		toSerialize["rum_browser_legacy_session_count_sum"] = o.RumBrowserLegacySessionCountSum
	}
	if o.RumBrowserLiteSessionCountSum != nil {
		toSerialize["rum_browser_lite_session_count_sum"] = o.RumBrowserLiteSessionCountSum
	}
	if o.RumBrowserReplaySessionCountSum != nil {
		toSerialize["rum_browser_replay_session_count_sum"] = o.RumBrowserReplaySessionCountSum
	}
	if o.RumIndexedSessionsSum != nil {
		toSerialize["rum_indexed_sessions_sum"] = o.RumIndexedSessionsSum
	}
	if o.RumIngestedSessionsSum != nil {
		toSerialize["rum_ingested_sessions_sum"] = o.RumIngestedSessionsSum
	}
	if o.RumLiteSessionCountSum != nil {
		toSerialize["rum_lite_session_count_sum"] = o.RumLiteSessionCountSum
	}
	if o.RumMobileLegacySessionCountAndroidSum != nil {
		toSerialize["rum_mobile_legacy_session_count_android_sum"] = o.RumMobileLegacySessionCountAndroidSum
	}
	if o.RumMobileLegacySessionCountFlutterSum != nil {
		toSerialize["rum_mobile_legacy_session_count_flutter_sum"] = o.RumMobileLegacySessionCountFlutterSum
	}
	if o.RumMobileLegacySessionCountIosSum != nil {
		toSerialize["rum_mobile_legacy_session_count_ios_sum"] = o.RumMobileLegacySessionCountIosSum
	}
	if o.RumMobileLegacySessionCountReactnativeSum != nil {
		toSerialize["rum_mobile_legacy_session_count_reactnative_sum"] = o.RumMobileLegacySessionCountReactnativeSum
	}
	if o.RumMobileLegacySessionCountRokuSum != nil {
		toSerialize["rum_mobile_legacy_session_count_roku_sum"] = o.RumMobileLegacySessionCountRokuSum
	}
	if o.RumMobileLiteSessionCountAndroidSum != nil {
		toSerialize["rum_mobile_lite_session_count_android_sum"] = o.RumMobileLiteSessionCountAndroidSum
	}
	if o.RumMobileLiteSessionCountFlutterSum != nil {
		toSerialize["rum_mobile_lite_session_count_flutter_sum"] = o.RumMobileLiteSessionCountFlutterSum
	}
	if o.RumMobileLiteSessionCountIosSum != nil {
		toSerialize["rum_mobile_lite_session_count_ios_sum"] = o.RumMobileLiteSessionCountIosSum
	}
	if o.RumMobileLiteSessionCountKotlinmultiplatformSum != nil {
		toSerialize["rum_mobile_lite_session_count_kotlinmultiplatform_sum"] = o.RumMobileLiteSessionCountKotlinmultiplatformSum
	}
	if o.RumMobileLiteSessionCountReactnativeSum != nil {
		toSerialize["rum_mobile_lite_session_count_reactnative_sum"] = o.RumMobileLiteSessionCountReactnativeSum
	}
	if o.RumMobileLiteSessionCountRokuSum != nil {
		toSerialize["rum_mobile_lite_session_count_roku_sum"] = o.RumMobileLiteSessionCountRokuSum
	}
	if o.RumMobileLiteSessionCountUnitySum != nil {
		toSerialize["rum_mobile_lite_session_count_unity_sum"] = o.RumMobileLiteSessionCountUnitySum
	}
	if o.RumMobileReplaySessionCountAndroidSum != nil {
		toSerialize["rum_mobile_replay_session_count_android_sum"] = o.RumMobileReplaySessionCountAndroidSum
	}
	if o.RumMobileReplaySessionCountIosSum != nil {
		toSerialize["rum_mobile_replay_session_count_ios_sum"] = o.RumMobileReplaySessionCountIosSum
	}
	if o.RumMobileReplaySessionCountKotlinmultiplatformSum != nil {
		toSerialize["rum_mobile_replay_session_count_kotlinmultiplatform_sum"] = o.RumMobileReplaySessionCountKotlinmultiplatformSum
	}
	if o.RumMobileReplaySessionCountReactnativeSum != nil {
		toSerialize["rum_mobile_replay_session_count_reactnative_sum"] = o.RumMobileReplaySessionCountReactnativeSum
	}
	if o.RumReplaySessionCountSum != nil {
		toSerialize["rum_replay_session_count_sum"] = o.RumReplaySessionCountSum
	}
	if o.RumSessionCountSum != nil {
		toSerialize["rum_session_count_sum"] = o.RumSessionCountSum
	}
	if o.RumSessionReplayAddOnSum != nil {
		toSerialize["rum_session_replay_add_on_sum"] = o.RumSessionReplayAddOnSum
	}
	if o.RumTotalSessionCountSum != nil {
		toSerialize["rum_total_session_count_sum"] = o.RumTotalSessionCountSum
	}
	if o.RumUnitsSum != nil {
		toSerialize["rum_units_sum"] = o.RumUnitsSum
	}
	if o.ScaFargateCountAvg != nil {
		toSerialize["sca_fargate_count_avg"] = o.ScaFargateCountAvg
	}
	if o.ScaFargateCountHwm != nil {
		toSerialize["sca_fargate_count_hwm"] = o.ScaFargateCountHwm
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
	if o.ServerlessAppsApmApmAzureAppserviceInstancesAvg != nil {
		toSerialize["serverless_apps_apm_apm_azure_appservice_instances_avg"] = o.ServerlessAppsApmApmAzureAppserviceInstancesAvg
	}
	if o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg != nil {
		toSerialize["serverless_apps_apm_apm_azure_azurefunction_instances_avg"] = o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg
	}
	if o.ServerlessAppsApmApmAzureContainerappInstancesAvg != nil {
		toSerialize["serverless_apps_apm_apm_azure_containerapp_instances_avg"] = o.ServerlessAppsApmApmAzureContainerappInstancesAvg
	}
	if o.ServerlessAppsApmApmFargateEcsTasksAvg != nil {
		toSerialize["serverless_apps_apm_apm_fargate_ecs_tasks_avg"] = o.ServerlessAppsApmApmFargateEcsTasksAvg
	}
	if o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg != nil {
		toSerialize["serverless_apps_apm_apm_gcp_cloudfunction_instances_avg"] = o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg
	}
	if o.ServerlessAppsApmApmGcpCloudrunInstancesAvg != nil {
		toSerialize["serverless_apps_apm_apm_gcp_cloudrun_instances_avg"] = o.ServerlessAppsApmApmGcpCloudrunInstancesAvg
	}
	if o.ServerlessAppsApmAvg != nil {
		toSerialize["serverless_apps_apm_avg"] = o.ServerlessAppsApmAvg
	}
	if o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg"] = o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg
	}
	if o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg"] = o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg
	}
	if o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg"] = o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg
	}
	if o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg"] = o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg
	}
	if o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg"] = o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg
	}
	if o.ServerlessAppsApmExclFargateAvg != nil {
		toSerialize["serverless_apps_apm_excl_fargate_avg"] = o.ServerlessAppsApmExclFargateAvg
	}
	if o.ServerlessAppsAzureContainerAppInstancesAvg != nil {
		toSerialize["serverless_apps_azure_container_app_instances_avg"] = o.ServerlessAppsAzureContainerAppInstancesAvg
	}
	if o.ServerlessAppsAzureCountAvg != nil {
		toSerialize["serverless_apps_azure_count_avg"] = o.ServerlessAppsAzureCountAvg
	}
	if o.ServerlessAppsAzureFunctionAppInstancesAvg != nil {
		toSerialize["serverless_apps_azure_function_app_instances_avg"] = o.ServerlessAppsAzureFunctionAppInstancesAvg
	}
	if o.ServerlessAppsAzureWebAppInstancesAvg != nil {
		toSerialize["serverless_apps_azure_web_app_instances_avg"] = o.ServerlessAppsAzureWebAppInstancesAvg
	}
	if o.ServerlessAppsEcsAvg != nil {
		toSerialize["serverless_apps_ecs_avg"] = o.ServerlessAppsEcsAvg
	}
	if o.ServerlessAppsEksAvg != nil {
		toSerialize["serverless_apps_eks_avg"] = o.ServerlessAppsEksAvg
	}
	if o.ServerlessAppsExclFargateAvg != nil {
		toSerialize["serverless_apps_excl_fargate_avg"] = o.ServerlessAppsExclFargateAvg
	}
	if o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg != nil {
		toSerialize["serverless_apps_excl_fargate_azure_container_app_instances_avg"] = o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg
	}
	if o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg != nil {
		toSerialize["serverless_apps_excl_fargate_azure_function_app_instances_avg"] = o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg
	}
	if o.ServerlessAppsExclFargateAzureWebAppInstancesAvg != nil {
		toSerialize["serverless_apps_excl_fargate_azure_web_app_instances_avg"] = o.ServerlessAppsExclFargateAzureWebAppInstancesAvg
	}
	if o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg != nil {
		toSerialize["serverless_apps_excl_fargate_google_cloud_functions_instances_avg"] = o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg
	}
	if o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg != nil {
		toSerialize["serverless_apps_excl_fargate_google_cloud_run_instances_avg"] = o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg
	}
	if o.ServerlessAppsGoogleCloudFunctionsInstancesAvg != nil {
		toSerialize["serverless_apps_google_cloud_functions_instances_avg"] = o.ServerlessAppsGoogleCloudFunctionsInstancesAvg
	}
	if o.ServerlessAppsGoogleCloudRunInstancesAvg != nil {
		toSerialize["serverless_apps_google_cloud_run_instances_avg"] = o.ServerlessAppsGoogleCloudRunInstancesAvg
	}
	if o.ServerlessAppsGoogleCountAvg != nil {
		toSerialize["serverless_apps_google_count_avg"] = o.ServerlessAppsGoogleCountAvg
	}
	if o.ServerlessAppsTotalCountAvg != nil {
		toSerialize["serverless_apps_total_count_avg"] = o.ServerlessAppsTotalCountAvg
	}
	if o.SiemAnalyzedLogsAddOnCountSum != nil {
		toSerialize["siem_analyzed_logs_add_on_count_sum"] = o.SiemAnalyzedLogsAddOnCountSum
	}
	if o.SyntheticsBrowserCheckCallsCountSum != nil {
		toSerialize["synthetics_browser_check_calls_count_sum"] = o.SyntheticsBrowserCheckCallsCountSum
	}
	if o.SyntheticsCheckCallsCountSum != nil {
		toSerialize["synthetics_check_calls_count_sum"] = o.SyntheticsCheckCallsCountSum
	}
	if o.SyntheticsMobileTestRunsSum != nil {
		toSerialize["synthetics_mobile_test_runs_sum"] = o.SyntheticsMobileTestRunsSum
	}
	if o.SyntheticsParallelTestingMaxSlotsHwm != nil {
		toSerialize["synthetics_parallel_testing_max_slots_hwm"] = o.SyntheticsParallelTestingMaxSlotsHwm
	}
	if o.TraceSearchIndexedEventsCountSum != nil {
		toSerialize["trace_search_indexed_events_count_sum"] = o.TraceSearchIndexedEventsCountSum
	}
	if o.TwolIngestedEventsBytesSum != nil {
		toSerialize["twol_ingested_events_bytes_sum"] = o.TwolIngestedEventsBytesSum
	}
	if o.UniversalServiceMonitoringHostTop99p != nil {
		toSerialize["universal_service_monitoring_host_top99p"] = o.UniversalServiceMonitoringHostTop99p
	}
	if o.VsphereHostTop99p != nil {
		toSerialize["vsphere_host_top99p"] = o.VsphereHostTop99p
	}
	if o.VulnManagementHostCountTop99p != nil {
		toSerialize["vuln_management_host_count_top99p"] = o.VulnManagementHostCountTop99p
	}
	if o.WorkflowExecutionsUsageSum != nil {
		toSerialize["workflow_executions_usage_sum"] = o.WorkflowExecutionsUsageSum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageSummaryDate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentHostTop99p                                               *int64                `json:"agent_host_top99p,omitempty"`
		ApmAzureAppServiceHostTop99p                                  *int64                `json:"apm_azure_app_service_host_top99p,omitempty"`
		ApmDevsecopsHostTop99p                                        *int64                `json:"apm_devsecops_host_top99p,omitempty"`
		ApmEnterpriseStandaloneHostsTop99p                            *int64                `json:"apm_enterprise_standalone_hosts_top99p,omitempty"`
		ApmFargateCountAvg                                            *int64                `json:"apm_fargate_count_avg,omitempty"`
		ApmHostTop99p                                                 *int64                `json:"apm_host_top99p,omitempty"`
		ApmProStandaloneHostsTop99p                                   *int64                `json:"apm_pro_standalone_hosts_top99p,omitempty"`
		AppsecFargateCountAvg                                         *int64                `json:"appsec_fargate_count_avg,omitempty"`
		AsmServerlessSum                                              *int64                `json:"asm_serverless_sum,omitempty"`
		AuditLogsLinesIndexedSum                                      *int64                `json:"audit_logs_lines_indexed_sum,omitempty"`
		AuditTrailEnabledHwm                                          *int64                `json:"audit_trail_enabled_hwm,omitempty"`
		AvgProfiledFargateTasks                                       *int64                `json:"avg_profiled_fargate_tasks,omitempty"`
		AwsHostTop99p                                                 *int64                `json:"aws_host_top99p,omitempty"`
		AwsLambdaFuncCount                                            *int64                `json:"aws_lambda_func_count,omitempty"`
		AwsLambdaInvocationsSum                                       *int64                `json:"aws_lambda_invocations_sum,omitempty"`
		AzureAppServiceTop99p                                         *int64                `json:"azure_app_service_top99p,omitempty"`
		BillableIngestedBytesSum                                      *int64                `json:"billable_ingested_bytes_sum,omitempty"`
		BitsAiInvestigationsSum                                       *int64                `json:"bits_ai_investigations_sum,omitempty"`
		BrowserRumLiteSessionCountSum                                 *int64                `json:"browser_rum_lite_session_count_sum,omitempty"`
		BrowserRumReplaySessionCountSum                               *int64                `json:"browser_rum_replay_session_count_sum,omitempty"`
		BrowserRumUnitsSum                                            *int64                `json:"browser_rum_units_sum,omitempty"`
		CiPipelineIndexedSpansSum                                     *int64                `json:"ci_pipeline_indexed_spans_sum,omitempty"`
		CiTestIndexedSpansSum                                         *int64                `json:"ci_test_indexed_spans_sum,omitempty"`
		CiVisibilityItrCommittersHwm                                  *int64                `json:"ci_visibility_itr_committers_hwm,omitempty"`
		CiVisibilityPipelineCommittersHwm                             *int64                `json:"ci_visibility_pipeline_committers_hwm,omitempty"`
		CiVisibilityTestCommittersHwm                                 *int64                `json:"ci_visibility_test_committers_hwm,omitempty"`
		CloudCostManagementAwsHostCountAvg                            *int64                `json:"cloud_cost_management_aws_host_count_avg,omitempty"`
		CloudCostManagementAzureHostCountAvg                          *int64                `json:"cloud_cost_management_azure_host_count_avg,omitempty"`
		CloudCostManagementGcpHostCountAvg                            *int64                `json:"cloud_cost_management_gcp_host_count_avg,omitempty"`
		CloudCostManagementHostCountAvg                               *int64                `json:"cloud_cost_management_host_count_avg,omitempty"`
		CloudCostManagementOciHostCountAvg                            *int64                `json:"cloud_cost_management_oci_host_count_avg,omitempty"`
		CloudSiemEventsSum                                            *int64                `json:"cloud_siem_events_sum,omitempty"`
		CodeAnalysisSaCommittersHwm                                   *int64                `json:"code_analysis_sa_committers_hwm,omitempty"`
		CodeAnalysisScaCommittersHwm                                  *int64                `json:"code_analysis_sca_committers_hwm,omitempty"`
		CodeSecurityHostTop99p                                        *int64                `json:"code_security_host_top99p,omitempty"`
		ContainerAvg                                                  *int64                `json:"container_avg,omitempty"`
		ContainerExclAgentAvg                                         *int64                `json:"container_excl_agent_avg,omitempty"`
		ContainerHwm                                                  *int64                `json:"container_hwm,omitempty"`
		CsmContainerEnterpriseComplianceCountSum                      *int64                `json:"csm_container_enterprise_compliance_count_sum,omitempty"`
		CsmContainerEnterpriseCwsCountSum                             *int64                `json:"csm_container_enterprise_cws_count_sum,omitempty"`
		CsmContainerEnterpriseTotalCountSum                           *int64                `json:"csm_container_enterprise_total_count_sum,omitempty"`
		CsmHostEnterpriseAasHostCountTop99p                           *int64                `json:"csm_host_enterprise_aas_host_count_top99p,omitempty"`
		CsmHostEnterpriseAwsHostCountTop99p                           *int64                `json:"csm_host_enterprise_aws_host_count_top99p,omitempty"`
		CsmHostEnterpriseAzureHostCountTop99p                         *int64                `json:"csm_host_enterprise_azure_host_count_top99p,omitempty"`
		CsmHostEnterpriseComplianceHostCountTop99p                    *int64                `json:"csm_host_enterprise_compliance_host_count_top99p,omitempty"`
		CsmHostEnterpriseCwsHostCountTop99p                           *int64                `json:"csm_host_enterprise_cws_host_count_top99p,omitempty"`
		CsmHostEnterpriseGcpHostCountTop99p                           *int64                `json:"csm_host_enterprise_gcp_host_count_top99p,omitempty"`
		CsmHostEnterpriseTotalHostCountTop99p                         *int64                `json:"csm_host_enterprise_total_host_count_top99p,omitempty"`
		CspmAasHostTop99p                                             *int64                `json:"cspm_aas_host_top99p,omitempty"`
		CspmAwsHostTop99p                                             *int64                `json:"cspm_aws_host_top99p,omitempty"`
		CspmAzureHostTop99p                                           *int64                `json:"cspm_azure_host_top99p,omitempty"`
		CspmContainerAvg                                              *int64                `json:"cspm_container_avg,omitempty"`
		CspmContainerHwm                                              *int64                `json:"cspm_container_hwm,omitempty"`
		CspmGcpHostTop99p                                             *int64                `json:"cspm_gcp_host_top99p,omitempty"`
		CspmHostTop99p                                                *int64                `json:"cspm_host_top99p,omitempty"`
		CustomTsAvg                                                   *int64                `json:"custom_ts_avg,omitempty"`
		CwsContainerCountAvg                                          *int64                `json:"cws_container_count_avg,omitempty"`
		CwsFargateTaskAvg                                             *int64                `json:"cws_fargate_task_avg,omitempty"`
		CwsHostTop99p                                                 *int64                `json:"cws_host_top99p,omitempty"`
		DataJobsMonitoringHostHrSum                                   *int64                `json:"data_jobs_monitoring_host_hr_sum,omitempty"`
		Date                                                          *time.Time            `json:"date,omitempty"`
		DbmHostTop99p                                                 *int64                `json:"dbm_host_top99p,omitempty"`
		DbmQueriesCountAvg                                            *int64                `json:"dbm_queries_count_avg,omitempty"`
		EphInfraHostAgentSum                                          *int64                `json:"eph_infra_host_agent_sum,omitempty"`
		EphInfraHostAlibabaSum                                        *int64                `json:"eph_infra_host_alibaba_sum,omitempty"`
		EphInfraHostAwsSum                                            *int64                `json:"eph_infra_host_aws_sum,omitempty"`
		EphInfraHostAzureSum                                          *int64                `json:"eph_infra_host_azure_sum,omitempty"`
		EphInfraHostEntSum                                            *int64                `json:"eph_infra_host_ent_sum,omitempty"`
		EphInfraHostGcpSum                                            *int64                `json:"eph_infra_host_gcp_sum,omitempty"`
		EphInfraHostHerokuSum                                         *int64                `json:"eph_infra_host_heroku_sum,omitempty"`
		EphInfraHostOnlyAasSum                                        *int64                `json:"eph_infra_host_only_aas_sum,omitempty"`
		EphInfraHostOnlyVsphereSum                                    *int64                `json:"eph_infra_host_only_vsphere_sum,omitempty"`
		EphInfraHostOpentelemetryApmSum                               *int64                `json:"eph_infra_host_opentelemetry_apm_sum,omitempty"`
		EphInfraHostOpentelemetrySum                                  *int64                `json:"eph_infra_host_opentelemetry_sum,omitempty"`
		EphInfraHostProSum                                            *int64                `json:"eph_infra_host_pro_sum,omitempty"`
		EphInfraHostProplusSum                                        *int64                `json:"eph_infra_host_proplus_sum,omitempty"`
		EphInfraHostProxmoxSum                                        *int64                `json:"eph_infra_host_proxmox_sum,omitempty"`
		ErrorTrackingApmErrorEventsSum                                *int64                `json:"error_tracking_apm_error_events_sum,omitempty"`
		ErrorTrackingErrorEventsSum                                   *int64                `json:"error_tracking_error_events_sum,omitempty"`
		ErrorTrackingEventsSum                                        *int64                `json:"error_tracking_events_sum,omitempty"`
		ErrorTrackingRumErrorEventsSum                                *int64                `json:"error_tracking_rum_error_events_sum,omitempty"`
		EventManagementCorrelationCorrelatedEventsSum                 *int64                `json:"event_management_correlation_correlated_events_sum,omitempty"`
		EventManagementCorrelationCorrelatedRelatedEventsSum          *int64                `json:"event_management_correlation_correlated_related_events_sum,omitempty"`
		EventManagementCorrelationSum                                 *int64                `json:"event_management_correlation_sum,omitempty"`
		FargateContainerProfilerProfilingFargateAvg                   *int64                `json:"fargate_container_profiler_profiling_fargate_avg,omitempty"`
		FargateContainerProfilerProfilingFargateEksAvg                *int64                `json:"fargate_container_profiler_profiling_fargate_eks_avg,omitempty"`
		FargateTasksCountAvg                                          *int64                `json:"fargate_tasks_count_avg,omitempty"`
		FargateTasksCountHwm                                          *int64                `json:"fargate_tasks_count_hwm,omitempty"`
		FlexLogsComputeLargeAvg                                       *int64                `json:"flex_logs_compute_large_avg,omitempty"`
		FlexLogsComputeMediumAvg                                      *int64                `json:"flex_logs_compute_medium_avg,omitempty"`
		FlexLogsComputeSmallAvg                                       *int64                `json:"flex_logs_compute_small_avg,omitempty"`
		FlexLogsComputeXlargeAvg                                      *int64                `json:"flex_logs_compute_xlarge_avg,omitempty"`
		FlexLogsComputeXsmallAvg                                      *int64                `json:"flex_logs_compute_xsmall_avg,omitempty"`
		FlexLogsStarterAvg                                            *int64                `json:"flex_logs_starter_avg,omitempty"`
		FlexLogsStarterStorageIndexAvg                                *int64                `json:"flex_logs_starter_storage_index_avg,omitempty"`
		FlexLogsStarterStorageRetentionAdjustmentAvg                  *int64                `json:"flex_logs_starter_storage_retention_adjustment_avg,omitempty"`
		FlexStoredLogsAvg                                             *int64                `json:"flex_stored_logs_avg,omitempty"`
		ForwardingEventsBytesSum                                      *int64                `json:"forwarding_events_bytes_sum,omitempty"`
		GcpHostTop99p                                                 *int64                `json:"gcp_host_top99p,omitempty"`
		HerokuHostTop99p                                              *int64                `json:"heroku_host_top99p,omitempty"`
		IncidentManagementMonthlyActiveUsersHwm                       *int64                `json:"incident_management_monthly_active_users_hwm,omitempty"`
		IncidentManagementSeatsHwm                                    *int64                `json:"incident_management_seats_hwm,omitempty"`
		IndexedEventsCountSum                                         *int64                `json:"indexed_events_count_sum,omitempty"`
		InfraHostTop99p                                               *int64                `json:"infra_host_top99p,omitempty"`
		IngestedEventsBytesSum                                        *int64                `json:"ingested_events_bytes_sum,omitempty"`
		IotDeviceSum                                                  *int64                `json:"iot_device_sum,omitempty"`
		IotDeviceTop99p                                               *int64                `json:"iot_device_top99p,omitempty"`
		LlmObservabilityMinSpendSum                                   *int64                `json:"llm_observability_min_spend_sum,omitempty"`
		LlmObservabilitySum                                           *int64                `json:"llm_observability_sum,omitempty"`
		MobileRumLiteSessionCountSum                                  *int64                `json:"mobile_rum_lite_session_count_sum,omitempty"`
		MobileRumSessionCountAndroidSum                               *int64                `json:"mobile_rum_session_count_android_sum,omitempty"`
		MobileRumSessionCountFlutterSum                               *int64                `json:"mobile_rum_session_count_flutter_sum,omitempty"`
		MobileRumSessionCountIosSum                                   *int64                `json:"mobile_rum_session_count_ios_sum,omitempty"`
		MobileRumSessionCountReactnativeSum                           *int64                `json:"mobile_rum_session_count_reactnative_sum,omitempty"`
		MobileRumSessionCountRokuSum                                  *int64                `json:"mobile_rum_session_count_roku_sum,omitempty"`
		MobileRumSessionCountSum                                      *int64                `json:"mobile_rum_session_count_sum,omitempty"`
		MobileRumUnitsSum                                             *int64                `json:"mobile_rum_units_sum,omitempty"`
		NdmNetflowEventsSum                                           *int64                `json:"ndm_netflow_events_sum,omitempty"`
		NetflowIndexedEventsCountSum                                  *int64                `json:"netflow_indexed_events_count_sum,omitempty"`
		NetworkDeviceWirelessTop99p                                   *int64                `json:"network_device_wireless_top99p,omitempty"`
		NpmHostTop99p                                                 *int64                `json:"npm_host_top99p,omitempty"`
		ObservabilityPipelinesBytesProcessedSum                       *int64                `json:"observability_pipelines_bytes_processed_sum,omitempty"`
		OciHostSum                                                    *int64                `json:"oci_host_sum,omitempty"`
		OciHostTop99p                                                 *int64                `json:"oci_host_top99p,omitempty"`
		OnCallSeatHwm                                                 *int64                `json:"on_call_seat_hwm,omitempty"`
		OnlineArchiveEventsCountSum                                   *int64                `json:"online_archive_events_count_sum,omitempty"`
		OpentelemetryApmHostTop99p                                    *int64                `json:"opentelemetry_apm_host_top99p,omitempty"`
		OpentelemetryHostTop99p                                       *int64                `json:"opentelemetry_host_top99p,omitempty"`
		Orgs                                                          []UsageSummaryDateOrg `json:"orgs,omitempty"`
		ProductAnalyticsSum                                           *int64                `json:"product_analytics_sum,omitempty"`
		ProfilingAasCountTop99p                                       *int64                `json:"profiling_aas_count_top99p,omitempty"`
		ProfilingHostTop99p                                           *int64                `json:"profiling_host_top99p,omitempty"`
		ProxmoxHostSum                                                *int64                `json:"proxmox_host_sum,omitempty"`
		ProxmoxHostTop99p                                             *int64                `json:"proxmox_host_top99p,omitempty"`
		PublishedAppHwm                                               *int64                `json:"published_app_hwm,omitempty"`
		RumBrowserAndMobileSessionCount                               *int64                `json:"rum_browser_and_mobile_session_count,omitempty"`
		RumBrowserLegacySessionCountSum                               *int64                `json:"rum_browser_legacy_session_count_sum,omitempty"`
		RumBrowserLiteSessionCountSum                                 *int64                `json:"rum_browser_lite_session_count_sum,omitempty"`
		RumBrowserReplaySessionCountSum                               *int64                `json:"rum_browser_replay_session_count_sum,omitempty"`
		RumIndexedSessionsSum                                         *int64                `json:"rum_indexed_sessions_sum,omitempty"`
		RumIngestedSessionsSum                                        *int64                `json:"rum_ingested_sessions_sum,omitempty"`
		RumLiteSessionCountSum                                        *int64                `json:"rum_lite_session_count_sum,omitempty"`
		RumMobileLegacySessionCountAndroidSum                         *int64                `json:"rum_mobile_legacy_session_count_android_sum,omitempty"`
		RumMobileLegacySessionCountFlutterSum                         *int64                `json:"rum_mobile_legacy_session_count_flutter_sum,omitempty"`
		RumMobileLegacySessionCountIosSum                             *int64                `json:"rum_mobile_legacy_session_count_ios_sum,omitempty"`
		RumMobileLegacySessionCountReactnativeSum                     *int64                `json:"rum_mobile_legacy_session_count_reactnative_sum,omitempty"`
		RumMobileLegacySessionCountRokuSum                            *int64                `json:"rum_mobile_legacy_session_count_roku_sum,omitempty"`
		RumMobileLiteSessionCountAndroidSum                           *int64                `json:"rum_mobile_lite_session_count_android_sum,omitempty"`
		RumMobileLiteSessionCountFlutterSum                           *int64                `json:"rum_mobile_lite_session_count_flutter_sum,omitempty"`
		RumMobileLiteSessionCountIosSum                               *int64                `json:"rum_mobile_lite_session_count_ios_sum,omitempty"`
		RumMobileLiteSessionCountKotlinmultiplatformSum               *int64                `json:"rum_mobile_lite_session_count_kotlinmultiplatform_sum,omitempty"`
		RumMobileLiteSessionCountReactnativeSum                       *int64                `json:"rum_mobile_lite_session_count_reactnative_sum,omitempty"`
		RumMobileLiteSessionCountRokuSum                              *int64                `json:"rum_mobile_lite_session_count_roku_sum,omitempty"`
		RumMobileLiteSessionCountUnitySum                             *int64                `json:"rum_mobile_lite_session_count_unity_sum,omitempty"`
		RumMobileReplaySessionCountAndroidSum                         *int64                `json:"rum_mobile_replay_session_count_android_sum,omitempty"`
		RumMobileReplaySessionCountIosSum                             *int64                `json:"rum_mobile_replay_session_count_ios_sum,omitempty"`
		RumMobileReplaySessionCountKotlinmultiplatformSum             *int64                `json:"rum_mobile_replay_session_count_kotlinmultiplatform_sum,omitempty"`
		RumMobileReplaySessionCountReactnativeSum                     *int64                `json:"rum_mobile_replay_session_count_reactnative_sum,omitempty"`
		RumReplaySessionCountSum                                      *int64                `json:"rum_replay_session_count_sum,omitempty"`
		RumSessionCountSum                                            *int64                `json:"rum_session_count_sum,omitempty"`
		RumSessionReplayAddOnSum                                      *int64                `json:"rum_session_replay_add_on_sum,omitempty"`
		RumTotalSessionCountSum                                       *int64                `json:"rum_total_session_count_sum,omitempty"`
		RumUnitsSum                                                   *int64                `json:"rum_units_sum,omitempty"`
		ScaFargateCountAvg                                            *int64                `json:"sca_fargate_count_avg,omitempty"`
		ScaFargateCountHwm                                            *int64                `json:"sca_fargate_count_hwm,omitempty"`
		SdsApmScannedBytesSum                                         *int64                `json:"sds_apm_scanned_bytes_sum,omitempty"`
		SdsEventsScannedBytesSum                                      *int64                `json:"sds_events_scanned_bytes_sum,omitempty"`
		SdsLogsScannedBytesSum                                        *int64                `json:"sds_logs_scanned_bytes_sum,omitempty"`
		SdsRumScannedBytesSum                                         *int64                `json:"sds_rum_scanned_bytes_sum,omitempty"`
		SdsTotalScannedBytesSum                                       *int64                `json:"sds_total_scanned_bytes_sum,omitempty"`
		ServerlessAppsApmApmAzureAppserviceInstancesAvg               *int64                `json:"serverless_apps_apm_apm_azure_appservice_instances_avg,omitempty"`
		ServerlessAppsApmApmAzureAzurefunctionInstancesAvg            *int64                `json:"serverless_apps_apm_apm_azure_azurefunction_instances_avg,omitempty"`
		ServerlessAppsApmApmAzureContainerappInstancesAvg             *int64                `json:"serverless_apps_apm_apm_azure_containerapp_instances_avg,omitempty"`
		ServerlessAppsApmApmFargateEcsTasksAvg                        *int64                `json:"serverless_apps_apm_apm_fargate_ecs_tasks_avg,omitempty"`
		ServerlessAppsApmApmGcpCloudfunctionInstancesAvg              *int64                `json:"serverless_apps_apm_apm_gcp_cloudfunction_instances_avg,omitempty"`
		ServerlessAppsApmApmGcpCloudrunInstancesAvg                   *int64                `json:"serverless_apps_apm_apm_gcp_cloudrun_instances_avg,omitempty"`
		ServerlessAppsApmAvg                                          *int64                `json:"serverless_apps_apm_avg,omitempty"`
		ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg    *int64                `json:"serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg *int64                `json:"serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg  *int64                `json:"serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg   *int64                `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg        *int64                `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateAvg                               *int64                `json:"serverless_apps_apm_excl_fargate_avg,omitempty"`
		ServerlessAppsAzureContainerAppInstancesAvg                   *int64                `json:"serverless_apps_azure_container_app_instances_avg,omitempty"`
		ServerlessAppsAzureCountAvg                                   *int64                `json:"serverless_apps_azure_count_avg,omitempty"`
		ServerlessAppsAzureFunctionAppInstancesAvg                    *int64                `json:"serverless_apps_azure_function_app_instances_avg,omitempty"`
		ServerlessAppsAzureWebAppInstancesAvg                         *int64                `json:"serverless_apps_azure_web_app_instances_avg,omitempty"`
		ServerlessAppsEcsAvg                                          *int64                `json:"serverless_apps_ecs_avg,omitempty"`
		ServerlessAppsEksAvg                                          *int64                `json:"serverless_apps_eks_avg,omitempty"`
		ServerlessAppsExclFargateAvg                                  *int64                `json:"serverless_apps_excl_fargate_avg,omitempty"`
		ServerlessAppsExclFargateAzureContainerAppInstancesAvg        *int64                `json:"serverless_apps_excl_fargate_azure_container_app_instances_avg,omitempty"`
		ServerlessAppsExclFargateAzureFunctionAppInstancesAvg         *int64                `json:"serverless_apps_excl_fargate_azure_function_app_instances_avg,omitempty"`
		ServerlessAppsExclFargateAzureWebAppInstancesAvg              *int64                `json:"serverless_apps_excl_fargate_azure_web_app_instances_avg,omitempty"`
		ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg     *int64                `json:"serverless_apps_excl_fargate_google_cloud_functions_instances_avg,omitempty"`
		ServerlessAppsExclFargateGoogleCloudRunInstancesAvg           *int64                `json:"serverless_apps_excl_fargate_google_cloud_run_instances_avg,omitempty"`
		ServerlessAppsGoogleCloudFunctionsInstancesAvg                *int64                `json:"serverless_apps_google_cloud_functions_instances_avg,omitempty"`
		ServerlessAppsGoogleCloudRunInstancesAvg                      *int64                `json:"serverless_apps_google_cloud_run_instances_avg,omitempty"`
		ServerlessAppsGoogleCountAvg                                  *int64                `json:"serverless_apps_google_count_avg,omitempty"`
		ServerlessAppsTotalCountAvg                                   *int64                `json:"serverless_apps_total_count_avg,omitempty"`
		SiemAnalyzedLogsAddOnCountSum                                 *int64                `json:"siem_analyzed_logs_add_on_count_sum,omitempty"`
		SyntheticsBrowserCheckCallsCountSum                           *int64                `json:"synthetics_browser_check_calls_count_sum,omitempty"`
		SyntheticsCheckCallsCountSum                                  *int64                `json:"synthetics_check_calls_count_sum,omitempty"`
		SyntheticsMobileTestRunsSum                                   *int64                `json:"synthetics_mobile_test_runs_sum,omitempty"`
		SyntheticsParallelTestingMaxSlotsHwm                          *int64                `json:"synthetics_parallel_testing_max_slots_hwm,omitempty"`
		TraceSearchIndexedEventsCountSum                              *int64                `json:"trace_search_indexed_events_count_sum,omitempty"`
		TwolIngestedEventsBytesSum                                    *int64                `json:"twol_ingested_events_bytes_sum,omitempty"`
		UniversalServiceMonitoringHostTop99p                          *int64                `json:"universal_service_monitoring_host_top99p,omitempty"`
		VsphereHostTop99p                                             *int64                `json:"vsphere_host_top99p,omitempty"`
		VulnManagementHostCountTop99p                                 *int64                `json:"vuln_management_host_count_top99p,omitempty"`
		WorkflowExecutionsUsageSum                                    *int64                `json:"workflow_executions_usage_sum,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_host_top99p", "apm_azure_app_service_host_top99p", "apm_devsecops_host_top99p", "apm_enterprise_standalone_hosts_top99p", "apm_fargate_count_avg", "apm_host_top99p", "apm_pro_standalone_hosts_top99p", "appsec_fargate_count_avg", "asm_serverless_sum", "audit_logs_lines_indexed_sum", "audit_trail_enabled_hwm", "avg_profiled_fargate_tasks", "aws_host_top99p", "aws_lambda_func_count", "aws_lambda_invocations_sum", "azure_app_service_top99p", "billable_ingested_bytes_sum", "bits_ai_investigations_sum", "browser_rum_lite_session_count_sum", "browser_rum_replay_session_count_sum", "browser_rum_units_sum", "ci_pipeline_indexed_spans_sum", "ci_test_indexed_spans_sum", "ci_visibility_itr_committers_hwm", "ci_visibility_pipeline_committers_hwm", "ci_visibility_test_committers_hwm", "cloud_cost_management_aws_host_count_avg", "cloud_cost_management_azure_host_count_avg", "cloud_cost_management_gcp_host_count_avg", "cloud_cost_management_host_count_avg", "cloud_cost_management_oci_host_count_avg", "cloud_siem_events_sum", "code_analysis_sa_committers_hwm", "code_analysis_sca_committers_hwm", "code_security_host_top99p", "container_avg", "container_excl_agent_avg", "container_hwm", "csm_container_enterprise_compliance_count_sum", "csm_container_enterprise_cws_count_sum", "csm_container_enterprise_total_count_sum", "csm_host_enterprise_aas_host_count_top99p", "csm_host_enterprise_aws_host_count_top99p", "csm_host_enterprise_azure_host_count_top99p", "csm_host_enterprise_compliance_host_count_top99p", "csm_host_enterprise_cws_host_count_top99p", "csm_host_enterprise_gcp_host_count_top99p", "csm_host_enterprise_total_host_count_top99p", "cspm_aas_host_top99p", "cspm_aws_host_top99p", "cspm_azure_host_top99p", "cspm_container_avg", "cspm_container_hwm", "cspm_gcp_host_top99p", "cspm_host_top99p", "custom_ts_avg", "cws_container_count_avg", "cws_fargate_task_avg", "cws_host_top99p", "data_jobs_monitoring_host_hr_sum", "date", "dbm_host_top99p", "dbm_queries_count_avg", "eph_infra_host_agent_sum", "eph_infra_host_alibaba_sum", "eph_infra_host_aws_sum", "eph_infra_host_azure_sum", "eph_infra_host_ent_sum", "eph_infra_host_gcp_sum", "eph_infra_host_heroku_sum", "eph_infra_host_only_aas_sum", "eph_infra_host_only_vsphere_sum", "eph_infra_host_opentelemetry_apm_sum", "eph_infra_host_opentelemetry_sum", "eph_infra_host_pro_sum", "eph_infra_host_proplus_sum", "eph_infra_host_proxmox_sum", "error_tracking_apm_error_events_sum", "error_tracking_error_events_sum", "error_tracking_events_sum", "error_tracking_rum_error_events_sum", "event_management_correlation_correlated_events_sum", "event_management_correlation_correlated_related_events_sum", "event_management_correlation_sum", "fargate_container_profiler_profiling_fargate_avg", "fargate_container_profiler_profiling_fargate_eks_avg", "fargate_tasks_count_avg", "fargate_tasks_count_hwm", "flex_logs_compute_large_avg", "flex_logs_compute_medium_avg", "flex_logs_compute_small_avg", "flex_logs_compute_xlarge_avg", "flex_logs_compute_xsmall_avg", "flex_logs_starter_avg", "flex_logs_starter_storage_index_avg", "flex_logs_starter_storage_retention_adjustment_avg", "flex_stored_logs_avg", "forwarding_events_bytes_sum", "gcp_host_top99p", "heroku_host_top99p", "incident_management_monthly_active_users_hwm", "incident_management_seats_hwm", "indexed_events_count_sum", "infra_host_top99p", "ingested_events_bytes_sum", "iot_device_sum", "iot_device_top99p", "llm_observability_min_spend_sum", "llm_observability_sum", "mobile_rum_lite_session_count_sum", "mobile_rum_session_count_android_sum", "mobile_rum_session_count_flutter_sum", "mobile_rum_session_count_ios_sum", "mobile_rum_session_count_reactnative_sum", "mobile_rum_session_count_roku_sum", "mobile_rum_session_count_sum", "mobile_rum_units_sum", "ndm_netflow_events_sum", "netflow_indexed_events_count_sum", "network_device_wireless_top99p", "npm_host_top99p", "observability_pipelines_bytes_processed_sum", "oci_host_sum", "oci_host_top99p", "on_call_seat_hwm", "online_archive_events_count_sum", "opentelemetry_apm_host_top99p", "opentelemetry_host_top99p", "orgs", "product_analytics_sum", "profiling_aas_count_top99p", "profiling_host_top99p", "proxmox_host_sum", "proxmox_host_top99p", "published_app_hwm", "rum_browser_and_mobile_session_count", "rum_browser_legacy_session_count_sum", "rum_browser_lite_session_count_sum", "rum_browser_replay_session_count_sum", "rum_indexed_sessions_sum", "rum_ingested_sessions_sum", "rum_lite_session_count_sum", "rum_mobile_legacy_session_count_android_sum", "rum_mobile_legacy_session_count_flutter_sum", "rum_mobile_legacy_session_count_ios_sum", "rum_mobile_legacy_session_count_reactnative_sum", "rum_mobile_legacy_session_count_roku_sum", "rum_mobile_lite_session_count_android_sum", "rum_mobile_lite_session_count_flutter_sum", "rum_mobile_lite_session_count_ios_sum", "rum_mobile_lite_session_count_kotlinmultiplatform_sum", "rum_mobile_lite_session_count_reactnative_sum", "rum_mobile_lite_session_count_roku_sum", "rum_mobile_lite_session_count_unity_sum", "rum_mobile_replay_session_count_android_sum", "rum_mobile_replay_session_count_ios_sum", "rum_mobile_replay_session_count_kotlinmultiplatform_sum", "rum_mobile_replay_session_count_reactnative_sum", "rum_replay_session_count_sum", "rum_session_count_sum", "rum_session_replay_add_on_sum", "rum_total_session_count_sum", "rum_units_sum", "sca_fargate_count_avg", "sca_fargate_count_hwm", "sds_apm_scanned_bytes_sum", "sds_events_scanned_bytes_sum", "sds_logs_scanned_bytes_sum", "sds_rum_scanned_bytes_sum", "sds_total_scanned_bytes_sum", "serverless_apps_apm_apm_azure_appservice_instances_avg", "serverless_apps_apm_apm_azure_azurefunction_instances_avg", "serverless_apps_apm_apm_azure_containerapp_instances_avg", "serverless_apps_apm_apm_fargate_ecs_tasks_avg", "serverless_apps_apm_apm_gcp_cloudfunction_instances_avg", "serverless_apps_apm_apm_gcp_cloudrun_instances_avg", "serverless_apps_apm_avg", "serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg", "serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg", "serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg", "serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg", "serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg", "serverless_apps_apm_excl_fargate_avg", "serverless_apps_azure_container_app_instances_avg", "serverless_apps_azure_count_avg", "serverless_apps_azure_function_app_instances_avg", "serverless_apps_azure_web_app_instances_avg", "serverless_apps_ecs_avg", "serverless_apps_eks_avg", "serverless_apps_excl_fargate_avg", "serverless_apps_excl_fargate_azure_container_app_instances_avg", "serverless_apps_excl_fargate_azure_function_app_instances_avg", "serverless_apps_excl_fargate_azure_web_app_instances_avg", "serverless_apps_excl_fargate_google_cloud_functions_instances_avg", "serverless_apps_excl_fargate_google_cloud_run_instances_avg", "serverless_apps_google_cloud_functions_instances_avg", "serverless_apps_google_cloud_run_instances_avg", "serverless_apps_google_count_avg", "serverless_apps_total_count_avg", "siem_analyzed_logs_add_on_count_sum", "synthetics_browser_check_calls_count_sum", "synthetics_check_calls_count_sum", "synthetics_mobile_test_runs_sum", "synthetics_parallel_testing_max_slots_hwm", "trace_search_indexed_events_count_sum", "twol_ingested_events_bytes_sum", "universal_service_monitoring_host_top99p", "vsphere_host_top99p", "vuln_management_host_count_top99p", "workflow_executions_usage_sum"})
	} else {
		return err
	}
	o.AgentHostTop99p = all.AgentHostTop99p
	o.ApmAzureAppServiceHostTop99p = all.ApmAzureAppServiceHostTop99p
	o.ApmDevsecopsHostTop99p = all.ApmDevsecopsHostTop99p
	o.ApmEnterpriseStandaloneHostsTop99p = all.ApmEnterpriseStandaloneHostsTop99p
	o.ApmFargateCountAvg = all.ApmFargateCountAvg
	o.ApmHostTop99p = all.ApmHostTop99p
	o.ApmProStandaloneHostsTop99p = all.ApmProStandaloneHostsTop99p
	o.AppsecFargateCountAvg = all.AppsecFargateCountAvg
	o.AsmServerlessSum = all.AsmServerlessSum
	o.AuditLogsLinesIndexedSum = all.AuditLogsLinesIndexedSum
	o.AuditTrailEnabledHwm = all.AuditTrailEnabledHwm
	o.AvgProfiledFargateTasks = all.AvgProfiledFargateTasks
	o.AwsHostTop99p = all.AwsHostTop99p
	o.AwsLambdaFuncCount = all.AwsLambdaFuncCount
	o.AwsLambdaInvocationsSum = all.AwsLambdaInvocationsSum
	o.AzureAppServiceTop99p = all.AzureAppServiceTop99p
	o.BillableIngestedBytesSum = all.BillableIngestedBytesSum
	o.BitsAiInvestigationsSum = all.BitsAiInvestigationsSum
	o.BrowserRumLiteSessionCountSum = all.BrowserRumLiteSessionCountSum
	o.BrowserRumReplaySessionCountSum = all.BrowserRumReplaySessionCountSum
	o.BrowserRumUnitsSum = all.BrowserRumUnitsSum
	o.CiPipelineIndexedSpansSum = all.CiPipelineIndexedSpansSum
	o.CiTestIndexedSpansSum = all.CiTestIndexedSpansSum
	o.CiVisibilityItrCommittersHwm = all.CiVisibilityItrCommittersHwm
	o.CiVisibilityPipelineCommittersHwm = all.CiVisibilityPipelineCommittersHwm
	o.CiVisibilityTestCommittersHwm = all.CiVisibilityTestCommittersHwm
	o.CloudCostManagementAwsHostCountAvg = all.CloudCostManagementAwsHostCountAvg
	o.CloudCostManagementAzureHostCountAvg = all.CloudCostManagementAzureHostCountAvg
	o.CloudCostManagementGcpHostCountAvg = all.CloudCostManagementGcpHostCountAvg
	o.CloudCostManagementHostCountAvg = all.CloudCostManagementHostCountAvg
	o.CloudCostManagementOciHostCountAvg = all.CloudCostManagementOciHostCountAvg
	o.CloudSiemEventsSum = all.CloudSiemEventsSum
	o.CodeAnalysisSaCommittersHwm = all.CodeAnalysisSaCommittersHwm
	o.CodeAnalysisScaCommittersHwm = all.CodeAnalysisScaCommittersHwm
	o.CodeSecurityHostTop99p = all.CodeSecurityHostTop99p
	o.ContainerAvg = all.ContainerAvg
	o.ContainerExclAgentAvg = all.ContainerExclAgentAvg
	o.ContainerHwm = all.ContainerHwm
	o.CsmContainerEnterpriseComplianceCountSum = all.CsmContainerEnterpriseComplianceCountSum
	o.CsmContainerEnterpriseCwsCountSum = all.CsmContainerEnterpriseCwsCountSum
	o.CsmContainerEnterpriseTotalCountSum = all.CsmContainerEnterpriseTotalCountSum
	o.CsmHostEnterpriseAasHostCountTop99p = all.CsmHostEnterpriseAasHostCountTop99p
	o.CsmHostEnterpriseAwsHostCountTop99p = all.CsmHostEnterpriseAwsHostCountTop99p
	o.CsmHostEnterpriseAzureHostCountTop99p = all.CsmHostEnterpriseAzureHostCountTop99p
	o.CsmHostEnterpriseComplianceHostCountTop99p = all.CsmHostEnterpriseComplianceHostCountTop99p
	o.CsmHostEnterpriseCwsHostCountTop99p = all.CsmHostEnterpriseCwsHostCountTop99p
	o.CsmHostEnterpriseGcpHostCountTop99p = all.CsmHostEnterpriseGcpHostCountTop99p
	o.CsmHostEnterpriseTotalHostCountTop99p = all.CsmHostEnterpriseTotalHostCountTop99p
	o.CspmAasHostTop99p = all.CspmAasHostTop99p
	o.CspmAwsHostTop99p = all.CspmAwsHostTop99p
	o.CspmAzureHostTop99p = all.CspmAzureHostTop99p
	o.CspmContainerAvg = all.CspmContainerAvg
	o.CspmContainerHwm = all.CspmContainerHwm
	o.CspmGcpHostTop99p = all.CspmGcpHostTop99p
	o.CspmHostTop99p = all.CspmHostTop99p
	o.CustomTsAvg = all.CustomTsAvg
	o.CwsContainerCountAvg = all.CwsContainerCountAvg
	o.CwsFargateTaskAvg = all.CwsFargateTaskAvg
	o.CwsHostTop99p = all.CwsHostTop99p
	o.DataJobsMonitoringHostHrSum = all.DataJobsMonitoringHostHrSum
	o.Date = all.Date
	o.DbmHostTop99p = all.DbmHostTop99p
	o.DbmQueriesCountAvg = all.DbmQueriesCountAvg
	o.EphInfraHostAgentSum = all.EphInfraHostAgentSum
	o.EphInfraHostAlibabaSum = all.EphInfraHostAlibabaSum
	o.EphInfraHostAwsSum = all.EphInfraHostAwsSum
	o.EphInfraHostAzureSum = all.EphInfraHostAzureSum
	o.EphInfraHostEntSum = all.EphInfraHostEntSum
	o.EphInfraHostGcpSum = all.EphInfraHostGcpSum
	o.EphInfraHostHerokuSum = all.EphInfraHostHerokuSum
	o.EphInfraHostOnlyAasSum = all.EphInfraHostOnlyAasSum
	o.EphInfraHostOnlyVsphereSum = all.EphInfraHostOnlyVsphereSum
	o.EphInfraHostOpentelemetryApmSum = all.EphInfraHostOpentelemetryApmSum
	o.EphInfraHostOpentelemetrySum = all.EphInfraHostOpentelemetrySum
	o.EphInfraHostProSum = all.EphInfraHostProSum
	o.EphInfraHostProplusSum = all.EphInfraHostProplusSum
	o.EphInfraHostProxmoxSum = all.EphInfraHostProxmoxSum
	o.ErrorTrackingApmErrorEventsSum = all.ErrorTrackingApmErrorEventsSum
	o.ErrorTrackingErrorEventsSum = all.ErrorTrackingErrorEventsSum
	o.ErrorTrackingEventsSum = all.ErrorTrackingEventsSum
	o.ErrorTrackingRumErrorEventsSum = all.ErrorTrackingRumErrorEventsSum
	o.EventManagementCorrelationCorrelatedEventsSum = all.EventManagementCorrelationCorrelatedEventsSum
	o.EventManagementCorrelationCorrelatedRelatedEventsSum = all.EventManagementCorrelationCorrelatedRelatedEventsSum
	o.EventManagementCorrelationSum = all.EventManagementCorrelationSum
	o.FargateContainerProfilerProfilingFargateAvg = all.FargateContainerProfilerProfilingFargateAvg
	o.FargateContainerProfilerProfilingFargateEksAvg = all.FargateContainerProfilerProfilingFargateEksAvg
	o.FargateTasksCountAvg = all.FargateTasksCountAvg
	o.FargateTasksCountHwm = all.FargateTasksCountHwm
	o.FlexLogsComputeLargeAvg = all.FlexLogsComputeLargeAvg
	o.FlexLogsComputeMediumAvg = all.FlexLogsComputeMediumAvg
	o.FlexLogsComputeSmallAvg = all.FlexLogsComputeSmallAvg
	o.FlexLogsComputeXlargeAvg = all.FlexLogsComputeXlargeAvg
	o.FlexLogsComputeXsmallAvg = all.FlexLogsComputeXsmallAvg
	o.FlexLogsStarterAvg = all.FlexLogsStarterAvg
	o.FlexLogsStarterStorageIndexAvg = all.FlexLogsStarterStorageIndexAvg
	o.FlexLogsStarterStorageRetentionAdjustmentAvg = all.FlexLogsStarterStorageRetentionAdjustmentAvg
	o.FlexStoredLogsAvg = all.FlexStoredLogsAvg
	o.ForwardingEventsBytesSum = all.ForwardingEventsBytesSum
	o.GcpHostTop99p = all.GcpHostTop99p
	o.HerokuHostTop99p = all.HerokuHostTop99p
	o.IncidentManagementMonthlyActiveUsersHwm = all.IncidentManagementMonthlyActiveUsersHwm
	o.IncidentManagementSeatsHwm = all.IncidentManagementSeatsHwm
	o.IndexedEventsCountSum = all.IndexedEventsCountSum
	o.InfraHostTop99p = all.InfraHostTop99p
	o.IngestedEventsBytesSum = all.IngestedEventsBytesSum
	o.IotDeviceSum = all.IotDeviceSum
	o.IotDeviceTop99p = all.IotDeviceTop99p
	o.LlmObservabilityMinSpendSum = all.LlmObservabilityMinSpendSum
	o.LlmObservabilitySum = all.LlmObservabilitySum
	o.MobileRumLiteSessionCountSum = all.MobileRumLiteSessionCountSum
	o.MobileRumSessionCountAndroidSum = all.MobileRumSessionCountAndroidSum
	o.MobileRumSessionCountFlutterSum = all.MobileRumSessionCountFlutterSum
	o.MobileRumSessionCountIosSum = all.MobileRumSessionCountIosSum
	o.MobileRumSessionCountReactnativeSum = all.MobileRumSessionCountReactnativeSum
	o.MobileRumSessionCountRokuSum = all.MobileRumSessionCountRokuSum
	o.MobileRumSessionCountSum = all.MobileRumSessionCountSum
	o.MobileRumUnitsSum = all.MobileRumUnitsSum
	o.NdmNetflowEventsSum = all.NdmNetflowEventsSum
	o.NetflowIndexedEventsCountSum = all.NetflowIndexedEventsCountSum
	o.NetworkDeviceWirelessTop99p = all.NetworkDeviceWirelessTop99p
	o.NpmHostTop99p = all.NpmHostTop99p
	o.ObservabilityPipelinesBytesProcessedSum = all.ObservabilityPipelinesBytesProcessedSum
	o.OciHostSum = all.OciHostSum
	o.OciHostTop99p = all.OciHostTop99p
	o.OnCallSeatHwm = all.OnCallSeatHwm
	o.OnlineArchiveEventsCountSum = all.OnlineArchiveEventsCountSum
	o.OpentelemetryApmHostTop99p = all.OpentelemetryApmHostTop99p
	o.OpentelemetryHostTop99p = all.OpentelemetryHostTop99p
	o.Orgs = all.Orgs
	o.ProductAnalyticsSum = all.ProductAnalyticsSum
	o.ProfilingAasCountTop99p = all.ProfilingAasCountTop99p
	o.ProfilingHostTop99p = all.ProfilingHostTop99p
	o.ProxmoxHostSum = all.ProxmoxHostSum
	o.ProxmoxHostTop99p = all.ProxmoxHostTop99p
	o.PublishedAppHwm = all.PublishedAppHwm
	o.RumBrowserAndMobileSessionCount = all.RumBrowserAndMobileSessionCount
	o.RumBrowserLegacySessionCountSum = all.RumBrowserLegacySessionCountSum
	o.RumBrowserLiteSessionCountSum = all.RumBrowserLiteSessionCountSum
	o.RumBrowserReplaySessionCountSum = all.RumBrowserReplaySessionCountSum
	o.RumIndexedSessionsSum = all.RumIndexedSessionsSum
	o.RumIngestedSessionsSum = all.RumIngestedSessionsSum
	o.RumLiteSessionCountSum = all.RumLiteSessionCountSum
	o.RumMobileLegacySessionCountAndroidSum = all.RumMobileLegacySessionCountAndroidSum
	o.RumMobileLegacySessionCountFlutterSum = all.RumMobileLegacySessionCountFlutterSum
	o.RumMobileLegacySessionCountIosSum = all.RumMobileLegacySessionCountIosSum
	o.RumMobileLegacySessionCountReactnativeSum = all.RumMobileLegacySessionCountReactnativeSum
	o.RumMobileLegacySessionCountRokuSum = all.RumMobileLegacySessionCountRokuSum
	o.RumMobileLiteSessionCountAndroidSum = all.RumMobileLiteSessionCountAndroidSum
	o.RumMobileLiteSessionCountFlutterSum = all.RumMobileLiteSessionCountFlutterSum
	o.RumMobileLiteSessionCountIosSum = all.RumMobileLiteSessionCountIosSum
	o.RumMobileLiteSessionCountKotlinmultiplatformSum = all.RumMobileLiteSessionCountKotlinmultiplatformSum
	o.RumMobileLiteSessionCountReactnativeSum = all.RumMobileLiteSessionCountReactnativeSum
	o.RumMobileLiteSessionCountRokuSum = all.RumMobileLiteSessionCountRokuSum
	o.RumMobileLiteSessionCountUnitySum = all.RumMobileLiteSessionCountUnitySum
	o.RumMobileReplaySessionCountAndroidSum = all.RumMobileReplaySessionCountAndroidSum
	o.RumMobileReplaySessionCountIosSum = all.RumMobileReplaySessionCountIosSum
	o.RumMobileReplaySessionCountKotlinmultiplatformSum = all.RumMobileReplaySessionCountKotlinmultiplatformSum
	o.RumMobileReplaySessionCountReactnativeSum = all.RumMobileReplaySessionCountReactnativeSum
	o.RumReplaySessionCountSum = all.RumReplaySessionCountSum
	o.RumSessionCountSum = all.RumSessionCountSum
	o.RumSessionReplayAddOnSum = all.RumSessionReplayAddOnSum
	o.RumTotalSessionCountSum = all.RumTotalSessionCountSum
	o.RumUnitsSum = all.RumUnitsSum
	o.ScaFargateCountAvg = all.ScaFargateCountAvg
	o.ScaFargateCountHwm = all.ScaFargateCountHwm
	o.SdsApmScannedBytesSum = all.SdsApmScannedBytesSum
	o.SdsEventsScannedBytesSum = all.SdsEventsScannedBytesSum
	o.SdsLogsScannedBytesSum = all.SdsLogsScannedBytesSum
	o.SdsRumScannedBytesSum = all.SdsRumScannedBytesSum
	o.SdsTotalScannedBytesSum = all.SdsTotalScannedBytesSum
	o.ServerlessAppsApmApmAzureAppserviceInstancesAvg = all.ServerlessAppsApmApmAzureAppserviceInstancesAvg
	o.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg = all.ServerlessAppsApmApmAzureAzurefunctionInstancesAvg
	o.ServerlessAppsApmApmAzureContainerappInstancesAvg = all.ServerlessAppsApmApmAzureContainerappInstancesAvg
	o.ServerlessAppsApmApmFargateEcsTasksAvg = all.ServerlessAppsApmApmFargateEcsTasksAvg
	o.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg = all.ServerlessAppsApmApmGcpCloudfunctionInstancesAvg
	o.ServerlessAppsApmApmGcpCloudrunInstancesAvg = all.ServerlessAppsApmApmGcpCloudrunInstancesAvg
	o.ServerlessAppsApmAvg = all.ServerlessAppsApmAvg
	o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg = all.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg
	o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg = all.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg
	o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg = all.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg
	o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg = all.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg
	o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg = all.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg
	o.ServerlessAppsApmExclFargateAvg = all.ServerlessAppsApmExclFargateAvg
	o.ServerlessAppsAzureContainerAppInstancesAvg = all.ServerlessAppsAzureContainerAppInstancesAvg
	o.ServerlessAppsAzureCountAvg = all.ServerlessAppsAzureCountAvg
	o.ServerlessAppsAzureFunctionAppInstancesAvg = all.ServerlessAppsAzureFunctionAppInstancesAvg
	o.ServerlessAppsAzureWebAppInstancesAvg = all.ServerlessAppsAzureWebAppInstancesAvg
	o.ServerlessAppsEcsAvg = all.ServerlessAppsEcsAvg
	o.ServerlessAppsEksAvg = all.ServerlessAppsEksAvg
	o.ServerlessAppsExclFargateAvg = all.ServerlessAppsExclFargateAvg
	o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg = all.ServerlessAppsExclFargateAzureContainerAppInstancesAvg
	o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg = all.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg
	o.ServerlessAppsExclFargateAzureWebAppInstancesAvg = all.ServerlessAppsExclFargateAzureWebAppInstancesAvg
	o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg = all.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg
	o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg = all.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg
	o.ServerlessAppsGoogleCloudFunctionsInstancesAvg = all.ServerlessAppsGoogleCloudFunctionsInstancesAvg
	o.ServerlessAppsGoogleCloudRunInstancesAvg = all.ServerlessAppsGoogleCloudRunInstancesAvg
	o.ServerlessAppsGoogleCountAvg = all.ServerlessAppsGoogleCountAvg
	o.ServerlessAppsTotalCountAvg = all.ServerlessAppsTotalCountAvg
	o.SiemAnalyzedLogsAddOnCountSum = all.SiemAnalyzedLogsAddOnCountSum
	o.SyntheticsBrowserCheckCallsCountSum = all.SyntheticsBrowserCheckCallsCountSum
	o.SyntheticsCheckCallsCountSum = all.SyntheticsCheckCallsCountSum
	o.SyntheticsMobileTestRunsSum = all.SyntheticsMobileTestRunsSum
	o.SyntheticsParallelTestingMaxSlotsHwm = all.SyntheticsParallelTestingMaxSlotsHwm
	o.TraceSearchIndexedEventsCountSum = all.TraceSearchIndexedEventsCountSum
	o.TwolIngestedEventsBytesSum = all.TwolIngestedEventsBytesSum
	o.UniversalServiceMonitoringHostTop99p = all.UniversalServiceMonitoringHostTop99p
	o.VsphereHostTop99p = all.VsphereHostTop99p
	o.VulnManagementHostCountTop99p = all.VulnManagementHostCountTop99p
	o.WorkflowExecutionsUsageSum = all.WorkflowExecutionsUsageSum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
