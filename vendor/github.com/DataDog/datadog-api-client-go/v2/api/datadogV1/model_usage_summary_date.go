// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSummaryDate Response with hourly report of all data billed by Datadog for all organizations.
//
// For SDK users only: all fields at this response level are accessible through the
// `additionalProperties` map. Existing typed-field getters are unchanged. New billing
// dimensions will not have typed-field getters. Use
// [Get available fields for usage summary](https://docs.datadoghq.com/api/latest/usage-metering/#get-usage-summary-available-fields)
// to enumerate every available key.
type UsageSummaryDate struct {
	// Shows the 99th percentile of all agent hosts over all hours in the current date for all organizations.
	AgentHostTop99p *int64 `json:"agent_host_top99p,omitempty"`
	// Shows the sum of all AI credits used by Agent Builder over all hours in the current date for all organizations.
	// Values are returned in micro-credits. Divide by 1,000,000 to get AI credits.
	AiCreditsAgentBuilderAiCreditsSum *int64 `json:"ai_credits_agent_builder_ai_credits_sum,omitempty"`
	// Shows the sum of all AI credits used by Bits AI Assistant over all hours in the current date for all organizations.
	// Values are returned in micro-credits. Divide by 1,000,000 to get AI credits.
	AiCreditsBitsAssistantAiCreditsSum *int64 `json:"ai_credits_bits_assistant_ai_credits_sum,omitempty"`
	// Shows the sum of all AI credits used by Bits AI Dev over all hours in the current date for all organizations.
	// Values are returned in micro-credits. Divide by 1,000,000 to get AI credits.
	AiCreditsBitsDevAiCreditsSum *int64 `json:"ai_credits_bits_dev_ai_credits_sum,omitempty"`
	// Shows the sum of all AI credits used by Bits AI SRE over all hours in the current date for all organizations.
	// Values are returned in micro-credits. Divide by 1,000,000 to get AI credits.
	AiCreditsBitsSreAiCreditsSum *int64 `json:"ai_credits_bits_sre_ai_credits_sum,omitempty"`
	// Shows the sum of all AI credits over all hours in the current date for all organizations.
	// Values are returned in micro-credits. Divide by 1,000,000 to get AI credits.
	AiCreditsSum *int64 `json:"ai_credits_sum,omitempty"`
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
	// Shows the sum of all Audit Trail event forwarding events over all hours in the current date for all organizations.
	AuditTrailEventForwardingEventsSum *int64 `json:"audit_trail_event_forwarding_events_sum,omitempty"`
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
	// Shows the last value of Anthropic cloud spend monitored over all hours in the current date for all organizations.
	CcmAnthropicSpendLast *int64 `json:"ccm_anthropic_spend_last,omitempty"`
	// Shows the last value of AWS cloud spend monitored over all hours in the current date for all organizations.
	CcmAwsSpendLast *int64 `json:"ccm_aws_spend_last,omitempty"`
	// Shows the last value of Azure cloud spend monitored over all hours in the current date for all organizations.
	CcmAzureSpendLast *int64 `json:"ccm_azure_spend_last,omitempty"`
	// Shows the last value of Confluent cloud spend monitored over all hours in the current date for all organizations.
	CcmConfluentSpendLast *int64 `json:"ccm_confluent_spend_last,omitempty"`
	// Shows the last value of Databricks cloud spend monitored over all hours in the current date for all organizations.
	CcmDatabricksSpendLast *int64 `json:"ccm_databricks_spend_last,omitempty"`
	// Shows the last value of Elastic cloud spend monitored over all hours in the current date for all organizations.
	CcmElasticSpendLast *int64 `json:"ccm_elastic_spend_last,omitempty"`
	// Shows the last value of Fastly cloud spend monitored over all hours in the current date for all organizations.
	CcmFastlySpendLast *int64 `json:"ccm_fastly_spend_last,omitempty"`
	// Shows the last value of GCP cloud spend monitored over all hours in the current date for all organizations.
	CcmGcpSpendLast *int64 `json:"ccm_gcp_spend_last,omitempty"`
	// Shows the last value of GitHub cloud spend monitored over all hours in the current date for all organizations.
	CcmGithubSpendLast *int64 `json:"ccm_github_spend_last,omitempty"`
	// Shows the last value of MongoDB cloud spend monitored over all hours in the current date for all organizations.
	CcmMongodbSpendLast *int64 `json:"ccm_mongodb_spend_last,omitempty"`
	// Shows the last value of OCI cloud spend monitored over all hours in the current date for all organizations.
	CcmOciSpendLast *int64 `json:"ccm_oci_spend_last,omitempty"`
	// Shows the last value of OpenAI cloud spend monitored over all hours in the current date for all organizations.
	CcmOpenaiSpendLast *int64 `json:"ccm_openai_spend_last,omitempty"`
	// Shows the last value of Snowflake cloud spend monitored over all hours in the current date for all organizations.
	CcmSnowflakeSpendLast *int64 `json:"ccm_snowflake_spend_last,omitempty"`
	// Shows the last value of the amount of cloud spend monitored for Enterprise over all hours in the current date for all organizations.
	CcmSpendMonitoredEntLast *int64 `json:"ccm_spend_monitored_ent_last,omitempty"`
	// Shows the last value of the amount of cloud spend monitored for Pro over all hours in the current date for all organizations.
	CcmSpendMonitoredProLast *int64 `json:"ccm_spend_monitored_pro_last,omitempty"`
	// Shows the last value of Twilio cloud spend monitored over all hours in the current date for all organizations.
	CcmTwilioSpendLast *int64 `json:"ccm_twilio_spend_last,omitempty"`
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
	// Shows the sum of all Cloud SIEM Indexed Logs over all hours in the current date for the given org.
	CloudSiemIndexedLogsSum *int64 `json:"cloud_siem_indexed_logs_sum,omitempty"`
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
	// Shows the 99th percentile of all Cloud Security Management Enterprise OCI hosts over all hours in the current date for the given org.
	CsmHostEnterpriseOciHostCountTop99p *int64 `json:"csm_host_enterprise_oci_host_count_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Enterprise hosts over all hours in the current date for the given org.
	CsmHostEnterpriseTotalHostCountTop99p *int64 `json:"csm_host_enterprise_total_host_count_top99p,omitempty"`
	// Shows the sum of all Cloud Security Management Pro Agentless scanner hosts over all hours in the current date for the given org.
	CsmHostProHostsAgentlessScannersSum *int64 `json:"csm_host_pro_hosts_agentless_scanners_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro Agentless scanner hosts over all hours in the current date for the given org.
	CsmHostProHostsAgentlessScannersTop99p *int64 `json:"csm_host_pro_hosts_agentless_scanners_top99p,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro OCI hosts over all hours in the current date for the given org.
	CsmHostProOciHostCountTop99p *int64 `json:"csm_host_pro_oci_host_count_top99p,omitempty"`
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
	// Shows the sum of all Cloud Security Management Pro Agentless scanner hosts over all hours in the current date for all organizations.
	CspmHostsAgentlessScannersSum *int64 `json:"cspm_hosts_agentless_scanners_sum,omitempty"`
	// Shows the 99th percentile of all Cloud Security Management Pro Agentless scanner hosts over all hours in the current date for all organizations.
	CspmHostsAgentlessScannersTop99p *int64 `json:"cspm_hosts_agentless_scanners_top99p,omitempty"`
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
	// Shows the sum of all Data Streams Monitoring hosts over all hours in the current date for all organizations.
	DataStreamMonitoringHostCountSum *int64 `json:"data_stream_monitoring_host_count_sum,omitempty"`
	// Shows the 99th percentile of all Data Streams Monitoring hosts over all hours in the current date for all organizations.
	DataStreamMonitoringHostCountTop99p *int64 `json:"data_stream_monitoring_host_count_top99p,omitempty"`
	// The date for the usage.
	Date *time.Time `json:"date,omitempty"`
	// Shows the 99th percentile of all Database Monitoring hosts over all hours in the current date for all organizations.
	DbmHostTop99p *int64 `json:"dbm_host_top99p,omitempty"`
	// Shows the average of all normalized Database Monitoring queries over all hours in the current date for all organizations.
	DbmQueriesCountAvg *int64 `json:"dbm_queries_count_avg,omitempty"`
	// Shows the sum of all orchestrator job hours over all hours in the current date for all organizations.
	// Values are returned in seconds. Divide by 3,600 to convert to hours.
	DoJobsMonitoringOrchestratorsJobHoursSum *int64 `json:"do_jobs_monitoring_orchestrators_job_hours_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts with the Datadog Agent over all hours in the current date for the given org.
	EphInfraHostAgentSum *int64 `json:"eph_infra_host_agent_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on Alibaba over all hours in the current date for the given org.
	EphInfraHostAlibabaSum *int64 `json:"eph_infra_host_alibaba_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on AWS over all hours in the current date for the given org.
	EphInfraHostAwsSum *int64 `json:"eph_infra_host_aws_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts on Azure over all hours in the current date for the given org.
	EphInfraHostAzureSum *int64 `json:"eph_infra_host_azure_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts for Basic tier with the Datadog Agent over all hours in the current date for all organizations.
	EphInfraHostBasicInfraBasicAgentSum *int64 `json:"eph_infra_host_basic_infra_basic_agent_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts for Basic tier on vSphere over all hours in the current date for all organizations.
	EphInfraHostBasicInfraBasicVsphereSum *int64 `json:"eph_infra_host_basic_infra_basic_vsphere_sum,omitempty"`
	// Shows the sum of all ephemeral infrastructure hosts for Basic tier over all hours in the current date for all organizations.
	EphInfraHostBasicSum *int64 `json:"eph_infra_host_basic_sum,omitempty"`
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
	// Shows the sum of all Feature Flags Client-Side SDK config requests over all hours in the current date for all organizations.
	FeatureFlagsConfigRequestsSum *int64 `json:"feature_flags_config_requests_sum,omitempty"`
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
	// Shows the sum of all indexed custom metrics points over all hours in the current date for all organizations.
	IndexedPointsSum *int64 `json:"indexed_points_sum,omitempty"`
	// Shows the average of all Infrastructure vCPU cores over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuAvg *int64 `json:"infra_cpu_avg,omitempty"`
	// Shows the average of all default Infrastructure host vCPU cores reported by the Datadog Agent over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuAgentAvg *int64 `json:"infra_cpu_default_infra_host_vcpu_agent_avg,omitempty"`
	// Shows the average of all default basic Infrastructure host vCPU cores reported by the Datadog Agent over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuAgentBasicAvg *int64 `json:"infra_cpu_default_infra_host_vcpu_agent_basic_avg,omitempty"`
	// Shows the sum of all default basic Infrastructure host vCPU cores reported by the Datadog Agent over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuAgentBasicSum *int64 `json:"infra_cpu_default_infra_host_vcpu_agent_basic_sum,omitempty"`
	// Shows the sum of all default Infrastructure host vCPU cores reported by the Datadog Agent over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuAgentSum *int64 `json:"infra_cpu_default_infra_host_vcpu_agent_sum,omitempty"`
	// Shows the average of all default Infrastructure host vCPU cores on AWS over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuAwsAvg *int64 `json:"infra_cpu_default_infra_host_vcpu_aws_avg,omitempty"`
	// Shows the sum of all default Infrastructure host vCPU cores on AWS over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuAwsSum *int64 `json:"infra_cpu_default_infra_host_vcpu_aws_sum,omitempty"`
	// Shows the average of all default Infrastructure host vCPU cores on Azure over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuAzureAvg *int64 `json:"infra_cpu_default_infra_host_vcpu_azure_avg,omitempty"`
	// Shows the sum of all default Infrastructure host vCPU cores on Azure over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuAzureSum *int64 `json:"infra_cpu_default_infra_host_vcpu_azure_sum,omitempty"`
	// Shows the average of all default Infrastructure host vCPU cores on GCP over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuGcpAvg *int64 `json:"infra_cpu_default_infra_host_vcpu_gcp_avg,omitempty"`
	// Shows the sum of all default Infrastructure host vCPU cores on GCP over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuGcpSum *int64 `json:"infra_cpu_default_infra_host_vcpu_gcp_sum,omitempty"`
	// Shows the average of all default Infrastructure host vCPU cores on Nutanix over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuNutanixAvg *int64 `json:"infra_cpu_default_infra_host_vcpu_nutanix_avg,omitempty"`
	// Shows the average of all default basic Infrastructure host vCPU cores on Nutanix over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuNutanixBasicAvg *int64 `json:"infra_cpu_default_infra_host_vcpu_nutanix_basic_avg,omitempty"`
	// Shows the sum of all default basic Infrastructure host vCPU cores on Nutanix over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuNutanixBasicSum *int64 `json:"infra_cpu_default_infra_host_vcpu_nutanix_basic_sum,omitempty"`
	// Shows the sum of all default Infrastructure host vCPU cores on Nutanix over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuNutanixSum *int64 `json:"infra_cpu_default_infra_host_vcpu_nutanix_sum,omitempty"`
	// Shows the average of all default Infrastructure host vCPU cores reported by OpenTelemetry over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuOpentelemetryAvg *int64 `json:"infra_cpu_default_infra_host_vcpu_opentelemetry_avg,omitempty"`
	// Shows the sum of all default Infrastructure host vCPU cores reported by OpenTelemetry over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuDefaultInfraHostVcpuOpentelemetrySum *int64 `json:"infra_cpu_default_infra_host_vcpu_opentelemetry_sum,omitempty"`
	// Shows the average of all observed Infrastructure host vCPU cores reported by the Datadog Agent over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuAgentAvg *int64 `json:"infra_cpu_observed_infra_host_vcpu_agent_avg,omitempty"`
	// Shows the sum of all observed Infrastructure host vCPU cores reported by the Datadog Agent over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuAgentSum *int64 `json:"infra_cpu_observed_infra_host_vcpu_agent_sum,omitempty"`
	// Shows the average of all observed Infrastructure host vCPU cores on AWS over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuAwsAvg *int64 `json:"infra_cpu_observed_infra_host_vcpu_aws_avg,omitempty"`
	// Shows the sum of all observed Infrastructure host vCPU cores on AWS over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuAwsSum *int64 `json:"infra_cpu_observed_infra_host_vcpu_aws_sum,omitempty"`
	// Shows the average of all observed Infrastructure host vCPU cores on Azure over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuAzureAvg *int64 `json:"infra_cpu_observed_infra_host_vcpu_azure_avg,omitempty"`
	// Shows the sum of all observed Infrastructure host vCPU cores on Azure over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuAzureSum *int64 `json:"infra_cpu_observed_infra_host_vcpu_azure_sum,omitempty"`
	// Shows the average of all observed Infrastructure host vCPU cores on GCP over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuGcpAvg *int64 `json:"infra_cpu_observed_infra_host_vcpu_gcp_avg,omitempty"`
	// Shows the sum of all observed Infrastructure host vCPU cores on GCP over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuGcpSum *int64 `json:"infra_cpu_observed_infra_host_vcpu_gcp_sum,omitempty"`
	// Shows the average of all observed Infrastructure host vCPU cores on Nutanix over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuNutanixAvg *int64 `json:"infra_cpu_observed_infra_host_vcpu_nutanix_avg,omitempty"`
	// Shows the sum of all observed Infrastructure host vCPU cores on Nutanix over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuNutanixSum *int64 `json:"infra_cpu_observed_infra_host_vcpu_nutanix_sum,omitempty"`
	// Shows the average of all observed Infrastructure host vCPU cores reported by OpenTelemetry over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuOpentelemetryAvg *int64 `json:"infra_cpu_observed_infra_host_vcpu_opentelemetry_avg,omitempty"`
	// Shows the sum of all observed Infrastructure host vCPU cores reported by OpenTelemetry over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuObservedInfraHostVcpuOpentelemetrySum *int64 `json:"infra_cpu_observed_infra_host_vcpu_opentelemetry_sum,omitempty"`
	// Shows the sum of all Infrastructure vCPU cores over all hours in the current date for all organizations.
	// Values are returned in millicores. Divide by 1,000 to convert to cores.
	InfraCpuSum *int64 `json:"infra_cpu_sum,omitempty"`
	// Shows the 99th percentile of all Edge Devices Monitoring devices over all hours in the current date for all organizations.
	InfraEdgeMonitoringDevicesTop99p *int64 `json:"infra_edge_monitoring_devices_top99p,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts for Basic tier with the Datadog Agent over all hours in the current date for all organizations.
	InfraHostBasicInfraBasicAgentTop99p *int64 `json:"infra_host_basic_infra_basic_agent_top99p,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts for Basic tier on vSphere over all hours in the current date for all organizations.
	InfraHostBasicInfraBasicVsphereTop99p *int64 `json:"infra_host_basic_infra_basic_vsphere_top99p,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts for Basic tier over all hours in the current date for all organizations.
	InfraHostBasicTop99p *int64 `json:"infra_host_basic_top99p,omitempty"`
	// Shows the 99th percentile of all distinct infrastructure hosts over all hours in the current date for all organizations.
	InfraHostTop99p *int64 `json:"infra_host_top99p,omitempty"`
	// Shows the average number of storage management objects over all hours in the current date for all organizations.
	InfraStorageMgmtObjectsCountAvg *int64 `json:"infra_storage_mgmt_objects_count_avg,omitempty"`
	// Shows the sum of all ingested custom metrics points over all hours in the current date for all organizations.
	IngestPointsSum *int64 `json:"ingest_points_sum,omitempty"`
	// Shows the sum of all log bytes ingested over all hours in the current date for all organizations.
	IngestedEventsBytesSum *int64 `json:"ingested_events_bytes_sum,omitempty"`
	// Shows the sum of all Application Performance Monitoring IoT hosts over all hours in the current date for all organizations.
	IotApmHostSum *int64 `json:"iot_apm_host_sum,omitempty"`
	// Shows the 99th percentile of all Application Performance Monitoring IoT hosts over all hours in the current date for all organizations.
	IotApmHostTop99p *int64 `json:"iot_apm_host_top99p,omitempty"`
	// Shows the sum of all IoT devices over all hours in the current date for all organizations.
	IotDeviceSum *int64 `json:"iot_device_sum,omitempty"`
	// Shows the 99th percentile of all IoT devices over all hours in the current date all organizations.
	IotDeviceTop99p *int64 `json:"iot_device_top99p,omitempty"`
	// Shows the sum of all LLM Observability 15-day retention spans over all hours in the current date for all organizations.
	LlmObservability15dayRetentionSpansSum *int64 `json:"llm_observability_15day_retention_spans_sum,omitempty"`
	// Shows the sum of all LLM Observability 30-day retention spans over all hours in the current date for all organizations.
	LlmObservability30dayRetentionSpansSum *int64 `json:"llm_observability_30day_retention_spans_sum,omitempty"`
	// Shows the sum of all LLM Observability 60-day retention spans over all hours in the current date for all organizations.
	LlmObservability60dayRetentionSpansSum *int64 `json:"llm_observability_60day_retention_spans_sum,omitempty"`
	// Shows the sum of all LLM Observability 90-day retention spans over all hours in the current date for all organizations.
	LlmObservability90dayRetentionSpansSum *int64 `json:"llm_observability_90day_retention_spans_sum,omitempty"`
	// Sum of all LLM observability minimum spend over all hours in the current date for all organizations.
	LlmObservabilityMinSpendSum *int64 `json:"llm_observability_min_spend_sum,omitempty"`
	// Sum of all LLM observability sessions over all hours in the current date for all organizations.
	LlmObservabilitySum *int64 `json:"llm_observability_sum,omitempty"`
	// Shows the sum of all Logs Archive Search scanned data over all hours in the current date for all organizations.
	LogsArchiveSearchGbScannedSum *int64 `json:"logs_archive_search_gb_scanned_sum,omitempty"`
	// Shows the sum of all custom metric names over all hours in the current date for all organizations.
	MetricNamesSum *int64 `json:"metric_names_sum,omitempty"`
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
	// Shows the sum of all Network Path scheduled tests over all hours in the current date for all organizations.
	NetworkPathSum *int64 `json:"network_path_sum,omitempty"`
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
	// Shows the average number of Serverless Apps with Application Performance Monitoring for Google Kubernetes Engine Autopilot pods for the current date for all organizations.
	ServerlessAppsApmApmGcpGkeAutopilotPodsAvg *int64 `json:"serverless_apps_apm_apm_gcp_gke_autopilot_pods_avg,omitempty"`
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
	// Shows the average number of Serverless Apps with Application Performance Monitoring excluding Fargate for Google Kubernetes Engine Autopilot pods for the current date for all organizations.
	ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg *int64 `json:"serverless_apps_apm_excl_fargate_apm_gcp_gke_autopilot_pods_avg,omitempty"`
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
	// Shows the average number of DSM Fargate ECS tasks monitored under Serverless Apps DSM for the current date for all organizations.
	ServerlessAppsDsmFargateTasksAvg *int64 `json:"serverless_apps_dsm_fargate_tasks_avg,omitempty"`
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
	// Shows the average number of Serverless Apps excluding Fargate for Google Kubernetes Engine Autopilot pods for the current date for all organizations.
	ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg *int64 `json:"serverless_apps_excl_fargate_infra_gcp_gke_autopilot_pods_avg,omitempty"`
	// Shows the average number of Serverless Apps for Google Cloud Platform Cloud Functions instances for the current date for all organizations.
	ServerlessAppsGoogleCloudFunctionsInstancesAvg *int64 `json:"serverless_apps_google_cloud_functions_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps for Google Cloud Platform Cloud Run instances for the current date for all organizations.
	ServerlessAppsGoogleCloudRunInstancesAvg *int64 `json:"serverless_apps_google_cloud_run_instances_avg,omitempty"`
	// Shows the average number of Serverless Apps for Google Cloud for the given date and given org.
	ServerlessAppsGoogleCountAvg *int64 `json:"serverless_apps_google_count_avg,omitempty"`
	// Shows the average number of Serverless Apps for Google Kubernetes Engine Autopilot pods for the current date for all organizations.
	ServerlessAppsInfraGcpGkeAutopilotPodsAvg *int64 `json:"serverless_apps_infra_gcp_gke_autopilot_pods_avg,omitempty"`
	// Shows the average number of Serverless Apps for Azure and Google Cloud for the given date and given org.
	ServerlessAppsTotalCountAvg *int64 `json:"serverless_apps_total_count_avg,omitempty"`
	// Shows the sum of Cloud SIEM Indexed Logs (12-month retention) over all hours in the current date for the given org.
	Siem12moRetentionSum *int64 `json:"siem_12mo_retention_sum,omitempty"`
	// Shows the sum of Cloud SIEM Indexed Logs (6-month retention) over all hours in the current date for the given org.
	Siem6moRetentionSum *int64 `json:"siem_6mo_retention_sum,omitempty"`
	// Shows the sum of all log events analyzed by Cloud SIEM over all hours in the current date for the given org.
	SiemAnalyzedLogsAddOnCountSum *int64 `json:"siem_analyzed_logs_add_on_count_sum,omitempty"`
	// Shows the sum of all Network Device Monitoring devices over all hours in the current date for all organizations.
	SnmpDeviceCountSum *int64 `json:"snmp_device_count_sum,omitempty"`
	// Shows the 99th percentile of all Network Device Monitoring devices over all hours in the current date for all organizations.
	SnmpDeviceCountTop99p *int64 `json:"snmp_device_count_top99p,omitempty"`
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

// GetAiCreditsAgentBuilderAiCreditsSum returns the AiCreditsAgentBuilderAiCreditsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAiCreditsAgentBuilderAiCreditsSum() int64 {
	if o == nil || o.AiCreditsAgentBuilderAiCreditsSum == nil {
		var ret int64
		return ret
	}
	return *o.AiCreditsAgentBuilderAiCreditsSum
}

// GetAiCreditsAgentBuilderAiCreditsSumOk returns a tuple with the AiCreditsAgentBuilderAiCreditsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAiCreditsAgentBuilderAiCreditsSumOk() (*int64, bool) {
	if o == nil || o.AiCreditsAgentBuilderAiCreditsSum == nil {
		return nil, false
	}
	return o.AiCreditsAgentBuilderAiCreditsSum, true
}

// HasAiCreditsAgentBuilderAiCreditsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAiCreditsAgentBuilderAiCreditsSum() bool {
	return o != nil && o.AiCreditsAgentBuilderAiCreditsSum != nil
}

// SetAiCreditsAgentBuilderAiCreditsSum gets a reference to the given int64 and assigns it to the AiCreditsAgentBuilderAiCreditsSum field.
func (o *UsageSummaryDate) SetAiCreditsAgentBuilderAiCreditsSum(v int64) {
	o.AiCreditsAgentBuilderAiCreditsSum = &v
}

// GetAiCreditsBitsAssistantAiCreditsSum returns the AiCreditsBitsAssistantAiCreditsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAiCreditsBitsAssistantAiCreditsSum() int64 {
	if o == nil || o.AiCreditsBitsAssistantAiCreditsSum == nil {
		var ret int64
		return ret
	}
	return *o.AiCreditsBitsAssistantAiCreditsSum
}

// GetAiCreditsBitsAssistantAiCreditsSumOk returns a tuple with the AiCreditsBitsAssistantAiCreditsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAiCreditsBitsAssistantAiCreditsSumOk() (*int64, bool) {
	if o == nil || o.AiCreditsBitsAssistantAiCreditsSum == nil {
		return nil, false
	}
	return o.AiCreditsBitsAssistantAiCreditsSum, true
}

// HasAiCreditsBitsAssistantAiCreditsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAiCreditsBitsAssistantAiCreditsSum() bool {
	return o != nil && o.AiCreditsBitsAssistantAiCreditsSum != nil
}

// SetAiCreditsBitsAssistantAiCreditsSum gets a reference to the given int64 and assigns it to the AiCreditsBitsAssistantAiCreditsSum field.
func (o *UsageSummaryDate) SetAiCreditsBitsAssistantAiCreditsSum(v int64) {
	o.AiCreditsBitsAssistantAiCreditsSum = &v
}

// GetAiCreditsBitsDevAiCreditsSum returns the AiCreditsBitsDevAiCreditsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAiCreditsBitsDevAiCreditsSum() int64 {
	if o == nil || o.AiCreditsBitsDevAiCreditsSum == nil {
		var ret int64
		return ret
	}
	return *o.AiCreditsBitsDevAiCreditsSum
}

// GetAiCreditsBitsDevAiCreditsSumOk returns a tuple with the AiCreditsBitsDevAiCreditsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAiCreditsBitsDevAiCreditsSumOk() (*int64, bool) {
	if o == nil || o.AiCreditsBitsDevAiCreditsSum == nil {
		return nil, false
	}
	return o.AiCreditsBitsDevAiCreditsSum, true
}

// HasAiCreditsBitsDevAiCreditsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAiCreditsBitsDevAiCreditsSum() bool {
	return o != nil && o.AiCreditsBitsDevAiCreditsSum != nil
}

// SetAiCreditsBitsDevAiCreditsSum gets a reference to the given int64 and assigns it to the AiCreditsBitsDevAiCreditsSum field.
func (o *UsageSummaryDate) SetAiCreditsBitsDevAiCreditsSum(v int64) {
	o.AiCreditsBitsDevAiCreditsSum = &v
}

// GetAiCreditsBitsSreAiCreditsSum returns the AiCreditsBitsSreAiCreditsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAiCreditsBitsSreAiCreditsSum() int64 {
	if o == nil || o.AiCreditsBitsSreAiCreditsSum == nil {
		var ret int64
		return ret
	}
	return *o.AiCreditsBitsSreAiCreditsSum
}

// GetAiCreditsBitsSreAiCreditsSumOk returns a tuple with the AiCreditsBitsSreAiCreditsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAiCreditsBitsSreAiCreditsSumOk() (*int64, bool) {
	if o == nil || o.AiCreditsBitsSreAiCreditsSum == nil {
		return nil, false
	}
	return o.AiCreditsBitsSreAiCreditsSum, true
}

// HasAiCreditsBitsSreAiCreditsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAiCreditsBitsSreAiCreditsSum() bool {
	return o != nil && o.AiCreditsBitsSreAiCreditsSum != nil
}

// SetAiCreditsBitsSreAiCreditsSum gets a reference to the given int64 and assigns it to the AiCreditsBitsSreAiCreditsSum field.
func (o *UsageSummaryDate) SetAiCreditsBitsSreAiCreditsSum(v int64) {
	o.AiCreditsBitsSreAiCreditsSum = &v
}

// GetAiCreditsSum returns the AiCreditsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAiCreditsSum() int64 {
	if o == nil || o.AiCreditsSum == nil {
		var ret int64
		return ret
	}
	return *o.AiCreditsSum
}

// GetAiCreditsSumOk returns a tuple with the AiCreditsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAiCreditsSumOk() (*int64, bool) {
	if o == nil || o.AiCreditsSum == nil {
		return nil, false
	}
	return o.AiCreditsSum, true
}

// HasAiCreditsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAiCreditsSum() bool {
	return o != nil && o.AiCreditsSum != nil
}

// SetAiCreditsSum gets a reference to the given int64 and assigns it to the AiCreditsSum field.
func (o *UsageSummaryDate) SetAiCreditsSum(v int64) {
	o.AiCreditsSum = &v
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

// GetAuditTrailEventForwardingEventsSum returns the AuditTrailEventForwardingEventsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetAuditTrailEventForwardingEventsSum() int64 {
	if o == nil || o.AuditTrailEventForwardingEventsSum == nil {
		var ret int64
		return ret
	}
	return *o.AuditTrailEventForwardingEventsSum
}

// GetAuditTrailEventForwardingEventsSumOk returns a tuple with the AuditTrailEventForwardingEventsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetAuditTrailEventForwardingEventsSumOk() (*int64, bool) {
	if o == nil || o.AuditTrailEventForwardingEventsSum == nil {
		return nil, false
	}
	return o.AuditTrailEventForwardingEventsSum, true
}

// HasAuditTrailEventForwardingEventsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasAuditTrailEventForwardingEventsSum() bool {
	return o != nil && o.AuditTrailEventForwardingEventsSum != nil
}

// SetAuditTrailEventForwardingEventsSum gets a reference to the given int64 and assigns it to the AuditTrailEventForwardingEventsSum field.
func (o *UsageSummaryDate) SetAuditTrailEventForwardingEventsSum(v int64) {
	o.AuditTrailEventForwardingEventsSum = &v
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

// GetCcmAnthropicSpendLast returns the CcmAnthropicSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmAnthropicSpendLast() int64 {
	if o == nil || o.CcmAnthropicSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmAnthropicSpendLast
}

// GetCcmAnthropicSpendLastOk returns a tuple with the CcmAnthropicSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmAnthropicSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmAnthropicSpendLast == nil {
		return nil, false
	}
	return o.CcmAnthropicSpendLast, true
}

// HasCcmAnthropicSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmAnthropicSpendLast() bool {
	return o != nil && o.CcmAnthropicSpendLast != nil
}

// SetCcmAnthropicSpendLast gets a reference to the given int64 and assigns it to the CcmAnthropicSpendLast field.
func (o *UsageSummaryDate) SetCcmAnthropicSpendLast(v int64) {
	o.CcmAnthropicSpendLast = &v
}

// GetCcmAwsSpendLast returns the CcmAwsSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmAwsSpendLast() int64 {
	if o == nil || o.CcmAwsSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmAwsSpendLast
}

// GetCcmAwsSpendLastOk returns a tuple with the CcmAwsSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmAwsSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmAwsSpendLast == nil {
		return nil, false
	}
	return o.CcmAwsSpendLast, true
}

// HasCcmAwsSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmAwsSpendLast() bool {
	return o != nil && o.CcmAwsSpendLast != nil
}

// SetCcmAwsSpendLast gets a reference to the given int64 and assigns it to the CcmAwsSpendLast field.
func (o *UsageSummaryDate) SetCcmAwsSpendLast(v int64) {
	o.CcmAwsSpendLast = &v
}

// GetCcmAzureSpendLast returns the CcmAzureSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmAzureSpendLast() int64 {
	if o == nil || o.CcmAzureSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmAzureSpendLast
}

// GetCcmAzureSpendLastOk returns a tuple with the CcmAzureSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmAzureSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmAzureSpendLast == nil {
		return nil, false
	}
	return o.CcmAzureSpendLast, true
}

// HasCcmAzureSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmAzureSpendLast() bool {
	return o != nil && o.CcmAzureSpendLast != nil
}

// SetCcmAzureSpendLast gets a reference to the given int64 and assigns it to the CcmAzureSpendLast field.
func (o *UsageSummaryDate) SetCcmAzureSpendLast(v int64) {
	o.CcmAzureSpendLast = &v
}

// GetCcmConfluentSpendLast returns the CcmConfluentSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmConfluentSpendLast() int64 {
	if o == nil || o.CcmConfluentSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmConfluentSpendLast
}

// GetCcmConfluentSpendLastOk returns a tuple with the CcmConfluentSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmConfluentSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmConfluentSpendLast == nil {
		return nil, false
	}
	return o.CcmConfluentSpendLast, true
}

// HasCcmConfluentSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmConfluentSpendLast() bool {
	return o != nil && o.CcmConfluentSpendLast != nil
}

// SetCcmConfluentSpendLast gets a reference to the given int64 and assigns it to the CcmConfluentSpendLast field.
func (o *UsageSummaryDate) SetCcmConfluentSpendLast(v int64) {
	o.CcmConfluentSpendLast = &v
}

// GetCcmDatabricksSpendLast returns the CcmDatabricksSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmDatabricksSpendLast() int64 {
	if o == nil || o.CcmDatabricksSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmDatabricksSpendLast
}

// GetCcmDatabricksSpendLastOk returns a tuple with the CcmDatabricksSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmDatabricksSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmDatabricksSpendLast == nil {
		return nil, false
	}
	return o.CcmDatabricksSpendLast, true
}

// HasCcmDatabricksSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmDatabricksSpendLast() bool {
	return o != nil && o.CcmDatabricksSpendLast != nil
}

// SetCcmDatabricksSpendLast gets a reference to the given int64 and assigns it to the CcmDatabricksSpendLast field.
func (o *UsageSummaryDate) SetCcmDatabricksSpendLast(v int64) {
	o.CcmDatabricksSpendLast = &v
}

// GetCcmElasticSpendLast returns the CcmElasticSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmElasticSpendLast() int64 {
	if o == nil || o.CcmElasticSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmElasticSpendLast
}

// GetCcmElasticSpendLastOk returns a tuple with the CcmElasticSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmElasticSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmElasticSpendLast == nil {
		return nil, false
	}
	return o.CcmElasticSpendLast, true
}

// HasCcmElasticSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmElasticSpendLast() bool {
	return o != nil && o.CcmElasticSpendLast != nil
}

// SetCcmElasticSpendLast gets a reference to the given int64 and assigns it to the CcmElasticSpendLast field.
func (o *UsageSummaryDate) SetCcmElasticSpendLast(v int64) {
	o.CcmElasticSpendLast = &v
}

// GetCcmFastlySpendLast returns the CcmFastlySpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmFastlySpendLast() int64 {
	if o == nil || o.CcmFastlySpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmFastlySpendLast
}

// GetCcmFastlySpendLastOk returns a tuple with the CcmFastlySpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmFastlySpendLastOk() (*int64, bool) {
	if o == nil || o.CcmFastlySpendLast == nil {
		return nil, false
	}
	return o.CcmFastlySpendLast, true
}

// HasCcmFastlySpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmFastlySpendLast() bool {
	return o != nil && o.CcmFastlySpendLast != nil
}

// SetCcmFastlySpendLast gets a reference to the given int64 and assigns it to the CcmFastlySpendLast field.
func (o *UsageSummaryDate) SetCcmFastlySpendLast(v int64) {
	o.CcmFastlySpendLast = &v
}

// GetCcmGcpSpendLast returns the CcmGcpSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmGcpSpendLast() int64 {
	if o == nil || o.CcmGcpSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmGcpSpendLast
}

// GetCcmGcpSpendLastOk returns a tuple with the CcmGcpSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmGcpSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmGcpSpendLast == nil {
		return nil, false
	}
	return o.CcmGcpSpendLast, true
}

// HasCcmGcpSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmGcpSpendLast() bool {
	return o != nil && o.CcmGcpSpendLast != nil
}

// SetCcmGcpSpendLast gets a reference to the given int64 and assigns it to the CcmGcpSpendLast field.
func (o *UsageSummaryDate) SetCcmGcpSpendLast(v int64) {
	o.CcmGcpSpendLast = &v
}

// GetCcmGithubSpendLast returns the CcmGithubSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmGithubSpendLast() int64 {
	if o == nil || o.CcmGithubSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmGithubSpendLast
}

// GetCcmGithubSpendLastOk returns a tuple with the CcmGithubSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmGithubSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmGithubSpendLast == nil {
		return nil, false
	}
	return o.CcmGithubSpendLast, true
}

// HasCcmGithubSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmGithubSpendLast() bool {
	return o != nil && o.CcmGithubSpendLast != nil
}

// SetCcmGithubSpendLast gets a reference to the given int64 and assigns it to the CcmGithubSpendLast field.
func (o *UsageSummaryDate) SetCcmGithubSpendLast(v int64) {
	o.CcmGithubSpendLast = &v
}

// GetCcmMongodbSpendLast returns the CcmMongodbSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmMongodbSpendLast() int64 {
	if o == nil || o.CcmMongodbSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmMongodbSpendLast
}

// GetCcmMongodbSpendLastOk returns a tuple with the CcmMongodbSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmMongodbSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmMongodbSpendLast == nil {
		return nil, false
	}
	return o.CcmMongodbSpendLast, true
}

// HasCcmMongodbSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmMongodbSpendLast() bool {
	return o != nil && o.CcmMongodbSpendLast != nil
}

// SetCcmMongodbSpendLast gets a reference to the given int64 and assigns it to the CcmMongodbSpendLast field.
func (o *UsageSummaryDate) SetCcmMongodbSpendLast(v int64) {
	o.CcmMongodbSpendLast = &v
}

// GetCcmOciSpendLast returns the CcmOciSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmOciSpendLast() int64 {
	if o == nil || o.CcmOciSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmOciSpendLast
}

// GetCcmOciSpendLastOk returns a tuple with the CcmOciSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmOciSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmOciSpendLast == nil {
		return nil, false
	}
	return o.CcmOciSpendLast, true
}

// HasCcmOciSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmOciSpendLast() bool {
	return o != nil && o.CcmOciSpendLast != nil
}

// SetCcmOciSpendLast gets a reference to the given int64 and assigns it to the CcmOciSpendLast field.
func (o *UsageSummaryDate) SetCcmOciSpendLast(v int64) {
	o.CcmOciSpendLast = &v
}

// GetCcmOpenaiSpendLast returns the CcmOpenaiSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmOpenaiSpendLast() int64 {
	if o == nil || o.CcmOpenaiSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmOpenaiSpendLast
}

// GetCcmOpenaiSpendLastOk returns a tuple with the CcmOpenaiSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmOpenaiSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmOpenaiSpendLast == nil {
		return nil, false
	}
	return o.CcmOpenaiSpendLast, true
}

// HasCcmOpenaiSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmOpenaiSpendLast() bool {
	return o != nil && o.CcmOpenaiSpendLast != nil
}

// SetCcmOpenaiSpendLast gets a reference to the given int64 and assigns it to the CcmOpenaiSpendLast field.
func (o *UsageSummaryDate) SetCcmOpenaiSpendLast(v int64) {
	o.CcmOpenaiSpendLast = &v
}

// GetCcmSnowflakeSpendLast returns the CcmSnowflakeSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmSnowflakeSpendLast() int64 {
	if o == nil || o.CcmSnowflakeSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmSnowflakeSpendLast
}

// GetCcmSnowflakeSpendLastOk returns a tuple with the CcmSnowflakeSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmSnowflakeSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmSnowflakeSpendLast == nil {
		return nil, false
	}
	return o.CcmSnowflakeSpendLast, true
}

// HasCcmSnowflakeSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmSnowflakeSpendLast() bool {
	return o != nil && o.CcmSnowflakeSpendLast != nil
}

// SetCcmSnowflakeSpendLast gets a reference to the given int64 and assigns it to the CcmSnowflakeSpendLast field.
func (o *UsageSummaryDate) SetCcmSnowflakeSpendLast(v int64) {
	o.CcmSnowflakeSpendLast = &v
}

// GetCcmSpendMonitoredEntLast returns the CcmSpendMonitoredEntLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmSpendMonitoredEntLast() int64 {
	if o == nil || o.CcmSpendMonitoredEntLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmSpendMonitoredEntLast
}

// GetCcmSpendMonitoredEntLastOk returns a tuple with the CcmSpendMonitoredEntLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmSpendMonitoredEntLastOk() (*int64, bool) {
	if o == nil || o.CcmSpendMonitoredEntLast == nil {
		return nil, false
	}
	return o.CcmSpendMonitoredEntLast, true
}

// HasCcmSpendMonitoredEntLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmSpendMonitoredEntLast() bool {
	return o != nil && o.CcmSpendMonitoredEntLast != nil
}

// SetCcmSpendMonitoredEntLast gets a reference to the given int64 and assigns it to the CcmSpendMonitoredEntLast field.
func (o *UsageSummaryDate) SetCcmSpendMonitoredEntLast(v int64) {
	o.CcmSpendMonitoredEntLast = &v
}

// GetCcmSpendMonitoredProLast returns the CcmSpendMonitoredProLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmSpendMonitoredProLast() int64 {
	if o == nil || o.CcmSpendMonitoredProLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmSpendMonitoredProLast
}

// GetCcmSpendMonitoredProLastOk returns a tuple with the CcmSpendMonitoredProLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmSpendMonitoredProLastOk() (*int64, bool) {
	if o == nil || o.CcmSpendMonitoredProLast == nil {
		return nil, false
	}
	return o.CcmSpendMonitoredProLast, true
}

// HasCcmSpendMonitoredProLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmSpendMonitoredProLast() bool {
	return o != nil && o.CcmSpendMonitoredProLast != nil
}

// SetCcmSpendMonitoredProLast gets a reference to the given int64 and assigns it to the CcmSpendMonitoredProLast field.
func (o *UsageSummaryDate) SetCcmSpendMonitoredProLast(v int64) {
	o.CcmSpendMonitoredProLast = &v
}

// GetCcmTwilioSpendLast returns the CcmTwilioSpendLast field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCcmTwilioSpendLast() int64 {
	if o == nil || o.CcmTwilioSpendLast == nil {
		var ret int64
		return ret
	}
	return *o.CcmTwilioSpendLast
}

// GetCcmTwilioSpendLastOk returns a tuple with the CcmTwilioSpendLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCcmTwilioSpendLastOk() (*int64, bool) {
	if o == nil || o.CcmTwilioSpendLast == nil {
		return nil, false
	}
	return o.CcmTwilioSpendLast, true
}

// HasCcmTwilioSpendLast returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCcmTwilioSpendLast() bool {
	return o != nil && o.CcmTwilioSpendLast != nil
}

// SetCcmTwilioSpendLast gets a reference to the given int64 and assigns it to the CcmTwilioSpendLast field.
func (o *UsageSummaryDate) SetCcmTwilioSpendLast(v int64) {
	o.CcmTwilioSpendLast = &v
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

// GetCloudSiemIndexedLogsSum returns the CloudSiemIndexedLogsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCloudSiemIndexedLogsSum() int64 {
	if o == nil || o.CloudSiemIndexedLogsSum == nil {
		var ret int64
		return ret
	}
	return *o.CloudSiemIndexedLogsSum
}

// GetCloudSiemIndexedLogsSumOk returns a tuple with the CloudSiemIndexedLogsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCloudSiemIndexedLogsSumOk() (*int64, bool) {
	if o == nil || o.CloudSiemIndexedLogsSum == nil {
		return nil, false
	}
	return o.CloudSiemIndexedLogsSum, true
}

// HasCloudSiemIndexedLogsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCloudSiemIndexedLogsSum() bool {
	return o != nil && o.CloudSiemIndexedLogsSum != nil
}

// SetCloudSiemIndexedLogsSum gets a reference to the given int64 and assigns it to the CloudSiemIndexedLogsSum field.
func (o *UsageSummaryDate) SetCloudSiemIndexedLogsSum(v int64) {
	o.CloudSiemIndexedLogsSum = &v
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

// GetCsmHostEnterpriseOciHostCountTop99p returns the CsmHostEnterpriseOciHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostEnterpriseOciHostCountTop99p() int64 {
	if o == nil || o.CsmHostEnterpriseOciHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostEnterpriseOciHostCountTop99p
}

// GetCsmHostEnterpriseOciHostCountTop99pOk returns a tuple with the CsmHostEnterpriseOciHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostEnterpriseOciHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostEnterpriseOciHostCountTop99p == nil {
		return nil, false
	}
	return o.CsmHostEnterpriseOciHostCountTop99p, true
}

// HasCsmHostEnterpriseOciHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostEnterpriseOciHostCountTop99p() bool {
	return o != nil && o.CsmHostEnterpriseOciHostCountTop99p != nil
}

// SetCsmHostEnterpriseOciHostCountTop99p gets a reference to the given int64 and assigns it to the CsmHostEnterpriseOciHostCountTop99p field.
func (o *UsageSummaryDate) SetCsmHostEnterpriseOciHostCountTop99p(v int64) {
	o.CsmHostEnterpriseOciHostCountTop99p = &v
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

// GetCsmHostProHostsAgentlessScannersSum returns the CsmHostProHostsAgentlessScannersSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostProHostsAgentlessScannersSum() int64 {
	if o == nil || o.CsmHostProHostsAgentlessScannersSum == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostProHostsAgentlessScannersSum
}

// GetCsmHostProHostsAgentlessScannersSumOk returns a tuple with the CsmHostProHostsAgentlessScannersSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostProHostsAgentlessScannersSumOk() (*int64, bool) {
	if o == nil || o.CsmHostProHostsAgentlessScannersSum == nil {
		return nil, false
	}
	return o.CsmHostProHostsAgentlessScannersSum, true
}

// HasCsmHostProHostsAgentlessScannersSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostProHostsAgentlessScannersSum() bool {
	return o != nil && o.CsmHostProHostsAgentlessScannersSum != nil
}

// SetCsmHostProHostsAgentlessScannersSum gets a reference to the given int64 and assigns it to the CsmHostProHostsAgentlessScannersSum field.
func (o *UsageSummaryDate) SetCsmHostProHostsAgentlessScannersSum(v int64) {
	o.CsmHostProHostsAgentlessScannersSum = &v
}

// GetCsmHostProHostsAgentlessScannersTop99p returns the CsmHostProHostsAgentlessScannersTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostProHostsAgentlessScannersTop99p() int64 {
	if o == nil || o.CsmHostProHostsAgentlessScannersTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostProHostsAgentlessScannersTop99p
}

// GetCsmHostProHostsAgentlessScannersTop99pOk returns a tuple with the CsmHostProHostsAgentlessScannersTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostProHostsAgentlessScannersTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostProHostsAgentlessScannersTop99p == nil {
		return nil, false
	}
	return o.CsmHostProHostsAgentlessScannersTop99p, true
}

// HasCsmHostProHostsAgentlessScannersTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostProHostsAgentlessScannersTop99p() bool {
	return o != nil && o.CsmHostProHostsAgentlessScannersTop99p != nil
}

// SetCsmHostProHostsAgentlessScannersTop99p gets a reference to the given int64 and assigns it to the CsmHostProHostsAgentlessScannersTop99p field.
func (o *UsageSummaryDate) SetCsmHostProHostsAgentlessScannersTop99p(v int64) {
	o.CsmHostProHostsAgentlessScannersTop99p = &v
}

// GetCsmHostProOciHostCountTop99p returns the CsmHostProOciHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCsmHostProOciHostCountTop99p() int64 {
	if o == nil || o.CsmHostProOciHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CsmHostProOciHostCountTop99p
}

// GetCsmHostProOciHostCountTop99pOk returns a tuple with the CsmHostProOciHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCsmHostProOciHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.CsmHostProOciHostCountTop99p == nil {
		return nil, false
	}
	return o.CsmHostProOciHostCountTop99p, true
}

// HasCsmHostProOciHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCsmHostProOciHostCountTop99p() bool {
	return o != nil && o.CsmHostProOciHostCountTop99p != nil
}

// SetCsmHostProOciHostCountTop99p gets a reference to the given int64 and assigns it to the CsmHostProOciHostCountTop99p field.
func (o *UsageSummaryDate) SetCsmHostProOciHostCountTop99p(v int64) {
	o.CsmHostProOciHostCountTop99p = &v
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

// GetCspmHostsAgentlessScannersSum returns the CspmHostsAgentlessScannersSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCspmHostsAgentlessScannersSum() int64 {
	if o == nil || o.CspmHostsAgentlessScannersSum == nil {
		var ret int64
		return ret
	}
	return *o.CspmHostsAgentlessScannersSum
}

// GetCspmHostsAgentlessScannersSumOk returns a tuple with the CspmHostsAgentlessScannersSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCspmHostsAgentlessScannersSumOk() (*int64, bool) {
	if o == nil || o.CspmHostsAgentlessScannersSum == nil {
		return nil, false
	}
	return o.CspmHostsAgentlessScannersSum, true
}

// HasCspmHostsAgentlessScannersSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmHostsAgentlessScannersSum() bool {
	return o != nil && o.CspmHostsAgentlessScannersSum != nil
}

// SetCspmHostsAgentlessScannersSum gets a reference to the given int64 and assigns it to the CspmHostsAgentlessScannersSum field.
func (o *UsageSummaryDate) SetCspmHostsAgentlessScannersSum(v int64) {
	o.CspmHostsAgentlessScannersSum = &v
}

// GetCspmHostsAgentlessScannersTop99p returns the CspmHostsAgentlessScannersTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetCspmHostsAgentlessScannersTop99p() int64 {
	if o == nil || o.CspmHostsAgentlessScannersTop99p == nil {
		var ret int64
		return ret
	}
	return *o.CspmHostsAgentlessScannersTop99p
}

// GetCspmHostsAgentlessScannersTop99pOk returns a tuple with the CspmHostsAgentlessScannersTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetCspmHostsAgentlessScannersTop99pOk() (*int64, bool) {
	if o == nil || o.CspmHostsAgentlessScannersTop99p == nil {
		return nil, false
	}
	return o.CspmHostsAgentlessScannersTop99p, true
}

// HasCspmHostsAgentlessScannersTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasCspmHostsAgentlessScannersTop99p() bool {
	return o != nil && o.CspmHostsAgentlessScannersTop99p != nil
}

// SetCspmHostsAgentlessScannersTop99p gets a reference to the given int64 and assigns it to the CspmHostsAgentlessScannersTop99p field.
func (o *UsageSummaryDate) SetCspmHostsAgentlessScannersTop99p(v int64) {
	o.CspmHostsAgentlessScannersTop99p = &v
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

// GetDataStreamMonitoringHostCountSum returns the DataStreamMonitoringHostCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetDataStreamMonitoringHostCountSum() int64 {
	if o == nil || o.DataStreamMonitoringHostCountSum == nil {
		var ret int64
		return ret
	}
	return *o.DataStreamMonitoringHostCountSum
}

// GetDataStreamMonitoringHostCountSumOk returns a tuple with the DataStreamMonitoringHostCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetDataStreamMonitoringHostCountSumOk() (*int64, bool) {
	if o == nil || o.DataStreamMonitoringHostCountSum == nil {
		return nil, false
	}
	return o.DataStreamMonitoringHostCountSum, true
}

// HasDataStreamMonitoringHostCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDataStreamMonitoringHostCountSum() bool {
	return o != nil && o.DataStreamMonitoringHostCountSum != nil
}

// SetDataStreamMonitoringHostCountSum gets a reference to the given int64 and assigns it to the DataStreamMonitoringHostCountSum field.
func (o *UsageSummaryDate) SetDataStreamMonitoringHostCountSum(v int64) {
	o.DataStreamMonitoringHostCountSum = &v
}

// GetDataStreamMonitoringHostCountTop99p returns the DataStreamMonitoringHostCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetDataStreamMonitoringHostCountTop99p() int64 {
	if o == nil || o.DataStreamMonitoringHostCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.DataStreamMonitoringHostCountTop99p
}

// GetDataStreamMonitoringHostCountTop99pOk returns a tuple with the DataStreamMonitoringHostCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetDataStreamMonitoringHostCountTop99pOk() (*int64, bool) {
	if o == nil || o.DataStreamMonitoringHostCountTop99p == nil {
		return nil, false
	}
	return o.DataStreamMonitoringHostCountTop99p, true
}

// HasDataStreamMonitoringHostCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDataStreamMonitoringHostCountTop99p() bool {
	return o != nil && o.DataStreamMonitoringHostCountTop99p != nil
}

// SetDataStreamMonitoringHostCountTop99p gets a reference to the given int64 and assigns it to the DataStreamMonitoringHostCountTop99p field.
func (o *UsageSummaryDate) SetDataStreamMonitoringHostCountTop99p(v int64) {
	o.DataStreamMonitoringHostCountTop99p = &v
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

// GetDoJobsMonitoringOrchestratorsJobHoursSum returns the DoJobsMonitoringOrchestratorsJobHoursSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetDoJobsMonitoringOrchestratorsJobHoursSum() int64 {
	if o == nil || o.DoJobsMonitoringOrchestratorsJobHoursSum == nil {
		var ret int64
		return ret
	}
	return *o.DoJobsMonitoringOrchestratorsJobHoursSum
}

// GetDoJobsMonitoringOrchestratorsJobHoursSumOk returns a tuple with the DoJobsMonitoringOrchestratorsJobHoursSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetDoJobsMonitoringOrchestratorsJobHoursSumOk() (*int64, bool) {
	if o == nil || o.DoJobsMonitoringOrchestratorsJobHoursSum == nil {
		return nil, false
	}
	return o.DoJobsMonitoringOrchestratorsJobHoursSum, true
}

// HasDoJobsMonitoringOrchestratorsJobHoursSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasDoJobsMonitoringOrchestratorsJobHoursSum() bool {
	return o != nil && o.DoJobsMonitoringOrchestratorsJobHoursSum != nil
}

// SetDoJobsMonitoringOrchestratorsJobHoursSum gets a reference to the given int64 and assigns it to the DoJobsMonitoringOrchestratorsJobHoursSum field.
func (o *UsageSummaryDate) SetDoJobsMonitoringOrchestratorsJobHoursSum(v int64) {
	o.DoJobsMonitoringOrchestratorsJobHoursSum = &v
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

// GetEphInfraHostBasicInfraBasicAgentSum returns the EphInfraHostBasicInfraBasicAgentSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostBasicInfraBasicAgentSum() int64 {
	if o == nil || o.EphInfraHostBasicInfraBasicAgentSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostBasicInfraBasicAgentSum
}

// GetEphInfraHostBasicInfraBasicAgentSumOk returns a tuple with the EphInfraHostBasicInfraBasicAgentSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostBasicInfraBasicAgentSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostBasicInfraBasicAgentSum == nil {
		return nil, false
	}
	return o.EphInfraHostBasicInfraBasicAgentSum, true
}

// HasEphInfraHostBasicInfraBasicAgentSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostBasicInfraBasicAgentSum() bool {
	return o != nil && o.EphInfraHostBasicInfraBasicAgentSum != nil
}

// SetEphInfraHostBasicInfraBasicAgentSum gets a reference to the given int64 and assigns it to the EphInfraHostBasicInfraBasicAgentSum field.
func (o *UsageSummaryDate) SetEphInfraHostBasicInfraBasicAgentSum(v int64) {
	o.EphInfraHostBasicInfraBasicAgentSum = &v
}

// GetEphInfraHostBasicInfraBasicVsphereSum returns the EphInfraHostBasicInfraBasicVsphereSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostBasicInfraBasicVsphereSum() int64 {
	if o == nil || o.EphInfraHostBasicInfraBasicVsphereSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostBasicInfraBasicVsphereSum
}

// GetEphInfraHostBasicInfraBasicVsphereSumOk returns a tuple with the EphInfraHostBasicInfraBasicVsphereSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostBasicInfraBasicVsphereSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostBasicInfraBasicVsphereSum == nil {
		return nil, false
	}
	return o.EphInfraHostBasicInfraBasicVsphereSum, true
}

// HasEphInfraHostBasicInfraBasicVsphereSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostBasicInfraBasicVsphereSum() bool {
	return o != nil && o.EphInfraHostBasicInfraBasicVsphereSum != nil
}

// SetEphInfraHostBasicInfraBasicVsphereSum gets a reference to the given int64 and assigns it to the EphInfraHostBasicInfraBasicVsphereSum field.
func (o *UsageSummaryDate) SetEphInfraHostBasicInfraBasicVsphereSum(v int64) {
	o.EphInfraHostBasicInfraBasicVsphereSum = &v
}

// GetEphInfraHostBasicSum returns the EphInfraHostBasicSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetEphInfraHostBasicSum() int64 {
	if o == nil || o.EphInfraHostBasicSum == nil {
		var ret int64
		return ret
	}
	return *o.EphInfraHostBasicSum
}

// GetEphInfraHostBasicSumOk returns a tuple with the EphInfraHostBasicSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetEphInfraHostBasicSumOk() (*int64, bool) {
	if o == nil || o.EphInfraHostBasicSum == nil {
		return nil, false
	}
	return o.EphInfraHostBasicSum, true
}

// HasEphInfraHostBasicSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasEphInfraHostBasicSum() bool {
	return o != nil && o.EphInfraHostBasicSum != nil
}

// SetEphInfraHostBasicSum gets a reference to the given int64 and assigns it to the EphInfraHostBasicSum field.
func (o *UsageSummaryDate) SetEphInfraHostBasicSum(v int64) {
	o.EphInfraHostBasicSum = &v
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

// GetFeatureFlagsConfigRequestsSum returns the FeatureFlagsConfigRequestsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetFeatureFlagsConfigRequestsSum() int64 {
	if o == nil || o.FeatureFlagsConfigRequestsSum == nil {
		var ret int64
		return ret
	}
	return *o.FeatureFlagsConfigRequestsSum
}

// GetFeatureFlagsConfigRequestsSumOk returns a tuple with the FeatureFlagsConfigRequestsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetFeatureFlagsConfigRequestsSumOk() (*int64, bool) {
	if o == nil || o.FeatureFlagsConfigRequestsSum == nil {
		return nil, false
	}
	return o.FeatureFlagsConfigRequestsSum, true
}

// HasFeatureFlagsConfigRequestsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasFeatureFlagsConfigRequestsSum() bool {
	return o != nil && o.FeatureFlagsConfigRequestsSum != nil
}

// SetFeatureFlagsConfigRequestsSum gets a reference to the given int64 and assigns it to the FeatureFlagsConfigRequestsSum field.
func (o *UsageSummaryDate) SetFeatureFlagsConfigRequestsSum(v int64) {
	o.FeatureFlagsConfigRequestsSum = &v
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

// GetIndexedPointsSum returns the IndexedPointsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIndexedPointsSum() int64 {
	if o == nil || o.IndexedPointsSum == nil {
		var ret int64
		return ret
	}
	return *o.IndexedPointsSum
}

// GetIndexedPointsSumOk returns a tuple with the IndexedPointsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIndexedPointsSumOk() (*int64, bool) {
	if o == nil || o.IndexedPointsSum == nil {
		return nil, false
	}
	return o.IndexedPointsSum, true
}

// HasIndexedPointsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIndexedPointsSum() bool {
	return o != nil && o.IndexedPointsSum != nil
}

// SetIndexedPointsSum gets a reference to the given int64 and assigns it to the IndexedPointsSum field.
func (o *UsageSummaryDate) SetIndexedPointsSum(v int64) {
	o.IndexedPointsSum = &v
}

// GetInfraCpuAvg returns the InfraCpuAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuAvg() int64 {
	if o == nil || o.InfraCpuAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuAvg
}

// GetInfraCpuAvgOk returns a tuple with the InfraCpuAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuAvg == nil {
		return nil, false
	}
	return o.InfraCpuAvg, true
}

// HasInfraCpuAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuAvg() bool {
	return o != nil && o.InfraCpuAvg != nil
}

// SetInfraCpuAvg gets a reference to the given int64 and assigns it to the InfraCpuAvg field.
func (o *UsageSummaryDate) SetInfraCpuAvg(v int64) {
	o.InfraCpuAvg = &v
}

// GetInfraCpuDefaultInfraHostVcpuAgentAvg returns the InfraCpuDefaultInfraHostVcpuAgentAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAgentAvg() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAgentAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuAgentAvg
}

// GetInfraCpuDefaultInfraHostVcpuAgentAvgOk returns a tuple with the InfraCpuDefaultInfraHostVcpuAgentAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAgentAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAgentAvg == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuAgentAvg, true
}

// HasInfraCpuDefaultInfraHostVcpuAgentAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuAgentAvg() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuAgentAvg != nil
}

// SetInfraCpuDefaultInfraHostVcpuAgentAvg gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuAgentAvg field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuAgentAvg(v int64) {
	o.InfraCpuDefaultInfraHostVcpuAgentAvg = &v
}

// GetInfraCpuDefaultInfraHostVcpuAgentBasicAvg returns the InfraCpuDefaultInfraHostVcpuAgentBasicAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAgentBasicAvg() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAgentBasicAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuAgentBasicAvg
}

// GetInfraCpuDefaultInfraHostVcpuAgentBasicAvgOk returns a tuple with the InfraCpuDefaultInfraHostVcpuAgentBasicAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAgentBasicAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAgentBasicAvg == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuAgentBasicAvg, true
}

// HasInfraCpuDefaultInfraHostVcpuAgentBasicAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuAgentBasicAvg() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuAgentBasicAvg != nil
}

// SetInfraCpuDefaultInfraHostVcpuAgentBasicAvg gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuAgentBasicAvg field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuAgentBasicAvg(v int64) {
	o.InfraCpuDefaultInfraHostVcpuAgentBasicAvg = &v
}

// GetInfraCpuDefaultInfraHostVcpuAgentBasicSum returns the InfraCpuDefaultInfraHostVcpuAgentBasicSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAgentBasicSum() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAgentBasicSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuAgentBasicSum
}

// GetInfraCpuDefaultInfraHostVcpuAgentBasicSumOk returns a tuple with the InfraCpuDefaultInfraHostVcpuAgentBasicSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAgentBasicSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAgentBasicSum == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuAgentBasicSum, true
}

// HasInfraCpuDefaultInfraHostVcpuAgentBasicSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuAgentBasicSum() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuAgentBasicSum != nil
}

// SetInfraCpuDefaultInfraHostVcpuAgentBasicSum gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuAgentBasicSum field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuAgentBasicSum(v int64) {
	o.InfraCpuDefaultInfraHostVcpuAgentBasicSum = &v
}

// GetInfraCpuDefaultInfraHostVcpuAgentSum returns the InfraCpuDefaultInfraHostVcpuAgentSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAgentSum() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAgentSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuAgentSum
}

// GetInfraCpuDefaultInfraHostVcpuAgentSumOk returns a tuple with the InfraCpuDefaultInfraHostVcpuAgentSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAgentSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAgentSum == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuAgentSum, true
}

// HasInfraCpuDefaultInfraHostVcpuAgentSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuAgentSum() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuAgentSum != nil
}

// SetInfraCpuDefaultInfraHostVcpuAgentSum gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuAgentSum field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuAgentSum(v int64) {
	o.InfraCpuDefaultInfraHostVcpuAgentSum = &v
}

// GetInfraCpuDefaultInfraHostVcpuAwsAvg returns the InfraCpuDefaultInfraHostVcpuAwsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAwsAvg() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAwsAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuAwsAvg
}

// GetInfraCpuDefaultInfraHostVcpuAwsAvgOk returns a tuple with the InfraCpuDefaultInfraHostVcpuAwsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAwsAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAwsAvg == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuAwsAvg, true
}

// HasInfraCpuDefaultInfraHostVcpuAwsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuAwsAvg() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuAwsAvg != nil
}

// SetInfraCpuDefaultInfraHostVcpuAwsAvg gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuAwsAvg field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuAwsAvg(v int64) {
	o.InfraCpuDefaultInfraHostVcpuAwsAvg = &v
}

// GetInfraCpuDefaultInfraHostVcpuAwsSum returns the InfraCpuDefaultInfraHostVcpuAwsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAwsSum() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAwsSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuAwsSum
}

// GetInfraCpuDefaultInfraHostVcpuAwsSumOk returns a tuple with the InfraCpuDefaultInfraHostVcpuAwsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAwsSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAwsSum == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuAwsSum, true
}

// HasInfraCpuDefaultInfraHostVcpuAwsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuAwsSum() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuAwsSum != nil
}

// SetInfraCpuDefaultInfraHostVcpuAwsSum gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuAwsSum field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuAwsSum(v int64) {
	o.InfraCpuDefaultInfraHostVcpuAwsSum = &v
}

// GetInfraCpuDefaultInfraHostVcpuAzureAvg returns the InfraCpuDefaultInfraHostVcpuAzureAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAzureAvg() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAzureAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuAzureAvg
}

// GetInfraCpuDefaultInfraHostVcpuAzureAvgOk returns a tuple with the InfraCpuDefaultInfraHostVcpuAzureAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAzureAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAzureAvg == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuAzureAvg, true
}

// HasInfraCpuDefaultInfraHostVcpuAzureAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuAzureAvg() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuAzureAvg != nil
}

// SetInfraCpuDefaultInfraHostVcpuAzureAvg gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuAzureAvg field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuAzureAvg(v int64) {
	o.InfraCpuDefaultInfraHostVcpuAzureAvg = &v
}

// GetInfraCpuDefaultInfraHostVcpuAzureSum returns the InfraCpuDefaultInfraHostVcpuAzureSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAzureSum() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAzureSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuAzureSum
}

// GetInfraCpuDefaultInfraHostVcpuAzureSumOk returns a tuple with the InfraCpuDefaultInfraHostVcpuAzureSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuAzureSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuAzureSum == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuAzureSum, true
}

// HasInfraCpuDefaultInfraHostVcpuAzureSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuAzureSum() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuAzureSum != nil
}

// SetInfraCpuDefaultInfraHostVcpuAzureSum gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuAzureSum field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuAzureSum(v int64) {
	o.InfraCpuDefaultInfraHostVcpuAzureSum = &v
}

// GetInfraCpuDefaultInfraHostVcpuGcpAvg returns the InfraCpuDefaultInfraHostVcpuGcpAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuGcpAvg() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuGcpAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuGcpAvg
}

// GetInfraCpuDefaultInfraHostVcpuGcpAvgOk returns a tuple with the InfraCpuDefaultInfraHostVcpuGcpAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuGcpAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuGcpAvg == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuGcpAvg, true
}

// HasInfraCpuDefaultInfraHostVcpuGcpAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuGcpAvg() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuGcpAvg != nil
}

// SetInfraCpuDefaultInfraHostVcpuGcpAvg gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuGcpAvg field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuGcpAvg(v int64) {
	o.InfraCpuDefaultInfraHostVcpuGcpAvg = &v
}

// GetInfraCpuDefaultInfraHostVcpuGcpSum returns the InfraCpuDefaultInfraHostVcpuGcpSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuGcpSum() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuGcpSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuGcpSum
}

// GetInfraCpuDefaultInfraHostVcpuGcpSumOk returns a tuple with the InfraCpuDefaultInfraHostVcpuGcpSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuGcpSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuGcpSum == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuGcpSum, true
}

// HasInfraCpuDefaultInfraHostVcpuGcpSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuGcpSum() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuGcpSum != nil
}

// SetInfraCpuDefaultInfraHostVcpuGcpSum gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuGcpSum field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuGcpSum(v int64) {
	o.InfraCpuDefaultInfraHostVcpuGcpSum = &v
}

// GetInfraCpuDefaultInfraHostVcpuNutanixAvg returns the InfraCpuDefaultInfraHostVcpuNutanixAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuNutanixAvg() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuNutanixAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuNutanixAvg
}

// GetInfraCpuDefaultInfraHostVcpuNutanixAvgOk returns a tuple with the InfraCpuDefaultInfraHostVcpuNutanixAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuNutanixAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuNutanixAvg == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuNutanixAvg, true
}

// HasInfraCpuDefaultInfraHostVcpuNutanixAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuNutanixAvg() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuNutanixAvg != nil
}

// SetInfraCpuDefaultInfraHostVcpuNutanixAvg gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuNutanixAvg field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuNutanixAvg(v int64) {
	o.InfraCpuDefaultInfraHostVcpuNutanixAvg = &v
}

// GetInfraCpuDefaultInfraHostVcpuNutanixBasicAvg returns the InfraCpuDefaultInfraHostVcpuNutanixBasicAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuNutanixBasicAvg() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg
}

// GetInfraCpuDefaultInfraHostVcpuNutanixBasicAvgOk returns a tuple with the InfraCpuDefaultInfraHostVcpuNutanixBasicAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuNutanixBasicAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg, true
}

// HasInfraCpuDefaultInfraHostVcpuNutanixBasicAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuNutanixBasicAvg() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg != nil
}

// SetInfraCpuDefaultInfraHostVcpuNutanixBasicAvg gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuNutanixBasicAvg field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuNutanixBasicAvg(v int64) {
	o.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg = &v
}

// GetInfraCpuDefaultInfraHostVcpuNutanixBasicSum returns the InfraCpuDefaultInfraHostVcpuNutanixBasicSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuNutanixBasicSum() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuNutanixBasicSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuNutanixBasicSum
}

// GetInfraCpuDefaultInfraHostVcpuNutanixBasicSumOk returns a tuple with the InfraCpuDefaultInfraHostVcpuNutanixBasicSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuNutanixBasicSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuNutanixBasicSum == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuNutanixBasicSum, true
}

// HasInfraCpuDefaultInfraHostVcpuNutanixBasicSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuNutanixBasicSum() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuNutanixBasicSum != nil
}

// SetInfraCpuDefaultInfraHostVcpuNutanixBasicSum gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuNutanixBasicSum field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuNutanixBasicSum(v int64) {
	o.InfraCpuDefaultInfraHostVcpuNutanixBasicSum = &v
}

// GetInfraCpuDefaultInfraHostVcpuNutanixSum returns the InfraCpuDefaultInfraHostVcpuNutanixSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuNutanixSum() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuNutanixSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuNutanixSum
}

// GetInfraCpuDefaultInfraHostVcpuNutanixSumOk returns a tuple with the InfraCpuDefaultInfraHostVcpuNutanixSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuNutanixSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuNutanixSum == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuNutanixSum, true
}

// HasInfraCpuDefaultInfraHostVcpuNutanixSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuNutanixSum() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuNutanixSum != nil
}

// SetInfraCpuDefaultInfraHostVcpuNutanixSum gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuNutanixSum field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuNutanixSum(v int64) {
	o.InfraCpuDefaultInfraHostVcpuNutanixSum = &v
}

// GetInfraCpuDefaultInfraHostVcpuOpentelemetryAvg returns the InfraCpuDefaultInfraHostVcpuOpentelemetryAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuOpentelemetryAvg() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg
}

// GetInfraCpuDefaultInfraHostVcpuOpentelemetryAvgOk returns a tuple with the InfraCpuDefaultInfraHostVcpuOpentelemetryAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuOpentelemetryAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg, true
}

// HasInfraCpuDefaultInfraHostVcpuOpentelemetryAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuOpentelemetryAvg() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg != nil
}

// SetInfraCpuDefaultInfraHostVcpuOpentelemetryAvg gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuOpentelemetryAvg field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuOpentelemetryAvg(v int64) {
	o.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg = &v
}

// GetInfraCpuDefaultInfraHostVcpuOpentelemetrySum returns the InfraCpuDefaultInfraHostVcpuOpentelemetrySum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuOpentelemetrySum() int64 {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuOpentelemetrySum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuDefaultInfraHostVcpuOpentelemetrySum
}

// GetInfraCpuDefaultInfraHostVcpuOpentelemetrySumOk returns a tuple with the InfraCpuDefaultInfraHostVcpuOpentelemetrySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuDefaultInfraHostVcpuOpentelemetrySumOk() (*int64, bool) {
	if o == nil || o.InfraCpuDefaultInfraHostVcpuOpentelemetrySum == nil {
		return nil, false
	}
	return o.InfraCpuDefaultInfraHostVcpuOpentelemetrySum, true
}

// HasInfraCpuDefaultInfraHostVcpuOpentelemetrySum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuDefaultInfraHostVcpuOpentelemetrySum() bool {
	return o != nil && o.InfraCpuDefaultInfraHostVcpuOpentelemetrySum != nil
}

// SetInfraCpuDefaultInfraHostVcpuOpentelemetrySum gets a reference to the given int64 and assigns it to the InfraCpuDefaultInfraHostVcpuOpentelemetrySum field.
func (o *UsageSummaryDate) SetInfraCpuDefaultInfraHostVcpuOpentelemetrySum(v int64) {
	o.InfraCpuDefaultInfraHostVcpuOpentelemetrySum = &v
}

// GetInfraCpuObservedInfraHostVcpuAgentAvg returns the InfraCpuObservedInfraHostVcpuAgentAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAgentAvg() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAgentAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuAgentAvg
}

// GetInfraCpuObservedInfraHostVcpuAgentAvgOk returns a tuple with the InfraCpuObservedInfraHostVcpuAgentAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAgentAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAgentAvg == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuAgentAvg, true
}

// HasInfraCpuObservedInfraHostVcpuAgentAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuAgentAvg() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuAgentAvg != nil
}

// SetInfraCpuObservedInfraHostVcpuAgentAvg gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuAgentAvg field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuAgentAvg(v int64) {
	o.InfraCpuObservedInfraHostVcpuAgentAvg = &v
}

// GetInfraCpuObservedInfraHostVcpuAgentSum returns the InfraCpuObservedInfraHostVcpuAgentSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAgentSum() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAgentSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuAgentSum
}

// GetInfraCpuObservedInfraHostVcpuAgentSumOk returns a tuple with the InfraCpuObservedInfraHostVcpuAgentSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAgentSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAgentSum == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuAgentSum, true
}

// HasInfraCpuObservedInfraHostVcpuAgentSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuAgentSum() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuAgentSum != nil
}

// SetInfraCpuObservedInfraHostVcpuAgentSum gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuAgentSum field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuAgentSum(v int64) {
	o.InfraCpuObservedInfraHostVcpuAgentSum = &v
}

// GetInfraCpuObservedInfraHostVcpuAwsAvg returns the InfraCpuObservedInfraHostVcpuAwsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAwsAvg() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAwsAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuAwsAvg
}

// GetInfraCpuObservedInfraHostVcpuAwsAvgOk returns a tuple with the InfraCpuObservedInfraHostVcpuAwsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAwsAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAwsAvg == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuAwsAvg, true
}

// HasInfraCpuObservedInfraHostVcpuAwsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuAwsAvg() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuAwsAvg != nil
}

// SetInfraCpuObservedInfraHostVcpuAwsAvg gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuAwsAvg field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuAwsAvg(v int64) {
	o.InfraCpuObservedInfraHostVcpuAwsAvg = &v
}

// GetInfraCpuObservedInfraHostVcpuAwsSum returns the InfraCpuObservedInfraHostVcpuAwsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAwsSum() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAwsSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuAwsSum
}

// GetInfraCpuObservedInfraHostVcpuAwsSumOk returns a tuple with the InfraCpuObservedInfraHostVcpuAwsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAwsSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAwsSum == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuAwsSum, true
}

// HasInfraCpuObservedInfraHostVcpuAwsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuAwsSum() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuAwsSum != nil
}

// SetInfraCpuObservedInfraHostVcpuAwsSum gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuAwsSum field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuAwsSum(v int64) {
	o.InfraCpuObservedInfraHostVcpuAwsSum = &v
}

// GetInfraCpuObservedInfraHostVcpuAzureAvg returns the InfraCpuObservedInfraHostVcpuAzureAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAzureAvg() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAzureAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuAzureAvg
}

// GetInfraCpuObservedInfraHostVcpuAzureAvgOk returns a tuple with the InfraCpuObservedInfraHostVcpuAzureAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAzureAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAzureAvg == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuAzureAvg, true
}

// HasInfraCpuObservedInfraHostVcpuAzureAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuAzureAvg() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuAzureAvg != nil
}

// SetInfraCpuObservedInfraHostVcpuAzureAvg gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuAzureAvg field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuAzureAvg(v int64) {
	o.InfraCpuObservedInfraHostVcpuAzureAvg = &v
}

// GetInfraCpuObservedInfraHostVcpuAzureSum returns the InfraCpuObservedInfraHostVcpuAzureSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAzureSum() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAzureSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuAzureSum
}

// GetInfraCpuObservedInfraHostVcpuAzureSumOk returns a tuple with the InfraCpuObservedInfraHostVcpuAzureSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuAzureSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuAzureSum == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuAzureSum, true
}

// HasInfraCpuObservedInfraHostVcpuAzureSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuAzureSum() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuAzureSum != nil
}

// SetInfraCpuObservedInfraHostVcpuAzureSum gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuAzureSum field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuAzureSum(v int64) {
	o.InfraCpuObservedInfraHostVcpuAzureSum = &v
}

// GetInfraCpuObservedInfraHostVcpuGcpAvg returns the InfraCpuObservedInfraHostVcpuGcpAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuGcpAvg() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuGcpAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuGcpAvg
}

// GetInfraCpuObservedInfraHostVcpuGcpAvgOk returns a tuple with the InfraCpuObservedInfraHostVcpuGcpAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuGcpAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuGcpAvg == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuGcpAvg, true
}

// HasInfraCpuObservedInfraHostVcpuGcpAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuGcpAvg() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuGcpAvg != nil
}

// SetInfraCpuObservedInfraHostVcpuGcpAvg gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuGcpAvg field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuGcpAvg(v int64) {
	o.InfraCpuObservedInfraHostVcpuGcpAvg = &v
}

// GetInfraCpuObservedInfraHostVcpuGcpSum returns the InfraCpuObservedInfraHostVcpuGcpSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuGcpSum() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuGcpSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuGcpSum
}

// GetInfraCpuObservedInfraHostVcpuGcpSumOk returns a tuple with the InfraCpuObservedInfraHostVcpuGcpSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuGcpSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuGcpSum == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuGcpSum, true
}

// HasInfraCpuObservedInfraHostVcpuGcpSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuGcpSum() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuGcpSum != nil
}

// SetInfraCpuObservedInfraHostVcpuGcpSum gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuGcpSum field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuGcpSum(v int64) {
	o.InfraCpuObservedInfraHostVcpuGcpSum = &v
}

// GetInfraCpuObservedInfraHostVcpuNutanixAvg returns the InfraCpuObservedInfraHostVcpuNutanixAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuNutanixAvg() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuNutanixAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuNutanixAvg
}

// GetInfraCpuObservedInfraHostVcpuNutanixAvgOk returns a tuple with the InfraCpuObservedInfraHostVcpuNutanixAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuNutanixAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuNutanixAvg == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuNutanixAvg, true
}

// HasInfraCpuObservedInfraHostVcpuNutanixAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuNutanixAvg() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuNutanixAvg != nil
}

// SetInfraCpuObservedInfraHostVcpuNutanixAvg gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuNutanixAvg field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuNutanixAvg(v int64) {
	o.InfraCpuObservedInfraHostVcpuNutanixAvg = &v
}

// GetInfraCpuObservedInfraHostVcpuNutanixSum returns the InfraCpuObservedInfraHostVcpuNutanixSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuNutanixSum() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuNutanixSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuNutanixSum
}

// GetInfraCpuObservedInfraHostVcpuNutanixSumOk returns a tuple with the InfraCpuObservedInfraHostVcpuNutanixSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuNutanixSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuNutanixSum == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuNutanixSum, true
}

// HasInfraCpuObservedInfraHostVcpuNutanixSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuNutanixSum() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuNutanixSum != nil
}

// SetInfraCpuObservedInfraHostVcpuNutanixSum gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuNutanixSum field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuNutanixSum(v int64) {
	o.InfraCpuObservedInfraHostVcpuNutanixSum = &v
}

// GetInfraCpuObservedInfraHostVcpuOpentelemetryAvg returns the InfraCpuObservedInfraHostVcpuOpentelemetryAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuOpentelemetryAvg() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuOpentelemetryAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuOpentelemetryAvg
}

// GetInfraCpuObservedInfraHostVcpuOpentelemetryAvgOk returns a tuple with the InfraCpuObservedInfraHostVcpuOpentelemetryAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuOpentelemetryAvgOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuOpentelemetryAvg == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuOpentelemetryAvg, true
}

// HasInfraCpuObservedInfraHostVcpuOpentelemetryAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuOpentelemetryAvg() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuOpentelemetryAvg != nil
}

// SetInfraCpuObservedInfraHostVcpuOpentelemetryAvg gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuOpentelemetryAvg field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuOpentelemetryAvg(v int64) {
	o.InfraCpuObservedInfraHostVcpuOpentelemetryAvg = &v
}

// GetInfraCpuObservedInfraHostVcpuOpentelemetrySum returns the InfraCpuObservedInfraHostVcpuOpentelemetrySum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuOpentelemetrySum() int64 {
	if o == nil || o.InfraCpuObservedInfraHostVcpuOpentelemetrySum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuObservedInfraHostVcpuOpentelemetrySum
}

// GetInfraCpuObservedInfraHostVcpuOpentelemetrySumOk returns a tuple with the InfraCpuObservedInfraHostVcpuOpentelemetrySum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuObservedInfraHostVcpuOpentelemetrySumOk() (*int64, bool) {
	if o == nil || o.InfraCpuObservedInfraHostVcpuOpentelemetrySum == nil {
		return nil, false
	}
	return o.InfraCpuObservedInfraHostVcpuOpentelemetrySum, true
}

// HasInfraCpuObservedInfraHostVcpuOpentelemetrySum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuObservedInfraHostVcpuOpentelemetrySum() bool {
	return o != nil && o.InfraCpuObservedInfraHostVcpuOpentelemetrySum != nil
}

// SetInfraCpuObservedInfraHostVcpuOpentelemetrySum gets a reference to the given int64 and assigns it to the InfraCpuObservedInfraHostVcpuOpentelemetrySum field.
func (o *UsageSummaryDate) SetInfraCpuObservedInfraHostVcpuOpentelemetrySum(v int64) {
	o.InfraCpuObservedInfraHostVcpuOpentelemetrySum = &v
}

// GetInfraCpuSum returns the InfraCpuSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraCpuSum() int64 {
	if o == nil || o.InfraCpuSum == nil {
		var ret int64
		return ret
	}
	return *o.InfraCpuSum
}

// GetInfraCpuSumOk returns a tuple with the InfraCpuSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraCpuSumOk() (*int64, bool) {
	if o == nil || o.InfraCpuSum == nil {
		return nil, false
	}
	return o.InfraCpuSum, true
}

// HasInfraCpuSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraCpuSum() bool {
	return o != nil && o.InfraCpuSum != nil
}

// SetInfraCpuSum gets a reference to the given int64 and assigns it to the InfraCpuSum field.
func (o *UsageSummaryDate) SetInfraCpuSum(v int64) {
	o.InfraCpuSum = &v
}

// GetInfraEdgeMonitoringDevicesTop99p returns the InfraEdgeMonitoringDevicesTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraEdgeMonitoringDevicesTop99p() int64 {
	if o == nil || o.InfraEdgeMonitoringDevicesTop99p == nil {
		var ret int64
		return ret
	}
	return *o.InfraEdgeMonitoringDevicesTop99p
}

// GetInfraEdgeMonitoringDevicesTop99pOk returns a tuple with the InfraEdgeMonitoringDevicesTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraEdgeMonitoringDevicesTop99pOk() (*int64, bool) {
	if o == nil || o.InfraEdgeMonitoringDevicesTop99p == nil {
		return nil, false
	}
	return o.InfraEdgeMonitoringDevicesTop99p, true
}

// HasInfraEdgeMonitoringDevicesTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraEdgeMonitoringDevicesTop99p() bool {
	return o != nil && o.InfraEdgeMonitoringDevicesTop99p != nil
}

// SetInfraEdgeMonitoringDevicesTop99p gets a reference to the given int64 and assigns it to the InfraEdgeMonitoringDevicesTop99p field.
func (o *UsageSummaryDate) SetInfraEdgeMonitoringDevicesTop99p(v int64) {
	o.InfraEdgeMonitoringDevicesTop99p = &v
}

// GetInfraHostBasicInfraBasicAgentTop99p returns the InfraHostBasicInfraBasicAgentTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraHostBasicInfraBasicAgentTop99p() int64 {
	if o == nil || o.InfraHostBasicInfraBasicAgentTop99p == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostBasicInfraBasicAgentTop99p
}

// GetInfraHostBasicInfraBasicAgentTop99pOk returns a tuple with the InfraHostBasicInfraBasicAgentTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraHostBasicInfraBasicAgentTop99pOk() (*int64, bool) {
	if o == nil || o.InfraHostBasicInfraBasicAgentTop99p == nil {
		return nil, false
	}
	return o.InfraHostBasicInfraBasicAgentTop99p, true
}

// HasInfraHostBasicInfraBasicAgentTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraHostBasicInfraBasicAgentTop99p() bool {
	return o != nil && o.InfraHostBasicInfraBasicAgentTop99p != nil
}

// SetInfraHostBasicInfraBasicAgentTop99p gets a reference to the given int64 and assigns it to the InfraHostBasicInfraBasicAgentTop99p field.
func (o *UsageSummaryDate) SetInfraHostBasicInfraBasicAgentTop99p(v int64) {
	o.InfraHostBasicInfraBasicAgentTop99p = &v
}

// GetInfraHostBasicInfraBasicVsphereTop99p returns the InfraHostBasicInfraBasicVsphereTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraHostBasicInfraBasicVsphereTop99p() int64 {
	if o == nil || o.InfraHostBasicInfraBasicVsphereTop99p == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostBasicInfraBasicVsphereTop99p
}

// GetInfraHostBasicInfraBasicVsphereTop99pOk returns a tuple with the InfraHostBasicInfraBasicVsphereTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraHostBasicInfraBasicVsphereTop99pOk() (*int64, bool) {
	if o == nil || o.InfraHostBasicInfraBasicVsphereTop99p == nil {
		return nil, false
	}
	return o.InfraHostBasicInfraBasicVsphereTop99p, true
}

// HasInfraHostBasicInfraBasicVsphereTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraHostBasicInfraBasicVsphereTop99p() bool {
	return o != nil && o.InfraHostBasicInfraBasicVsphereTop99p != nil
}

// SetInfraHostBasicInfraBasicVsphereTop99p gets a reference to the given int64 and assigns it to the InfraHostBasicInfraBasicVsphereTop99p field.
func (o *UsageSummaryDate) SetInfraHostBasicInfraBasicVsphereTop99p(v int64) {
	o.InfraHostBasicInfraBasicVsphereTop99p = &v
}

// GetInfraHostBasicTop99p returns the InfraHostBasicTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraHostBasicTop99p() int64 {
	if o == nil || o.InfraHostBasicTop99p == nil {
		var ret int64
		return ret
	}
	return *o.InfraHostBasicTop99p
}

// GetInfraHostBasicTop99pOk returns a tuple with the InfraHostBasicTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraHostBasicTop99pOk() (*int64, bool) {
	if o == nil || o.InfraHostBasicTop99p == nil {
		return nil, false
	}
	return o.InfraHostBasicTop99p, true
}

// HasInfraHostBasicTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraHostBasicTop99p() bool {
	return o != nil && o.InfraHostBasicTop99p != nil
}

// SetInfraHostBasicTop99p gets a reference to the given int64 and assigns it to the InfraHostBasicTop99p field.
func (o *UsageSummaryDate) SetInfraHostBasicTop99p(v int64) {
	o.InfraHostBasicTop99p = &v
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

// GetInfraStorageMgmtObjectsCountAvg returns the InfraStorageMgmtObjectsCountAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetInfraStorageMgmtObjectsCountAvg() int64 {
	if o == nil || o.InfraStorageMgmtObjectsCountAvg == nil {
		var ret int64
		return ret
	}
	return *o.InfraStorageMgmtObjectsCountAvg
}

// GetInfraStorageMgmtObjectsCountAvgOk returns a tuple with the InfraStorageMgmtObjectsCountAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetInfraStorageMgmtObjectsCountAvgOk() (*int64, bool) {
	if o == nil || o.InfraStorageMgmtObjectsCountAvg == nil {
		return nil, false
	}
	return o.InfraStorageMgmtObjectsCountAvg, true
}

// HasInfraStorageMgmtObjectsCountAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasInfraStorageMgmtObjectsCountAvg() bool {
	return o != nil && o.InfraStorageMgmtObjectsCountAvg != nil
}

// SetInfraStorageMgmtObjectsCountAvg gets a reference to the given int64 and assigns it to the InfraStorageMgmtObjectsCountAvg field.
func (o *UsageSummaryDate) SetInfraStorageMgmtObjectsCountAvg(v int64) {
	o.InfraStorageMgmtObjectsCountAvg = &v
}

// GetIngestPointsSum returns the IngestPointsSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIngestPointsSum() int64 {
	if o == nil || o.IngestPointsSum == nil {
		var ret int64
		return ret
	}
	return *o.IngestPointsSum
}

// GetIngestPointsSumOk returns a tuple with the IngestPointsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIngestPointsSumOk() (*int64, bool) {
	if o == nil || o.IngestPointsSum == nil {
		return nil, false
	}
	return o.IngestPointsSum, true
}

// HasIngestPointsSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIngestPointsSum() bool {
	return o != nil && o.IngestPointsSum != nil
}

// SetIngestPointsSum gets a reference to the given int64 and assigns it to the IngestPointsSum field.
func (o *UsageSummaryDate) SetIngestPointsSum(v int64) {
	o.IngestPointsSum = &v
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

// GetIotApmHostSum returns the IotApmHostSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIotApmHostSum() int64 {
	if o == nil || o.IotApmHostSum == nil {
		var ret int64
		return ret
	}
	return *o.IotApmHostSum
}

// GetIotApmHostSumOk returns a tuple with the IotApmHostSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIotApmHostSumOk() (*int64, bool) {
	if o == nil || o.IotApmHostSum == nil {
		return nil, false
	}
	return o.IotApmHostSum, true
}

// HasIotApmHostSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIotApmHostSum() bool {
	return o != nil && o.IotApmHostSum != nil
}

// SetIotApmHostSum gets a reference to the given int64 and assigns it to the IotApmHostSum field.
func (o *UsageSummaryDate) SetIotApmHostSum(v int64) {
	o.IotApmHostSum = &v
}

// GetIotApmHostTop99p returns the IotApmHostTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetIotApmHostTop99p() int64 {
	if o == nil || o.IotApmHostTop99p == nil {
		var ret int64
		return ret
	}
	return *o.IotApmHostTop99p
}

// GetIotApmHostTop99pOk returns a tuple with the IotApmHostTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetIotApmHostTop99pOk() (*int64, bool) {
	if o == nil || o.IotApmHostTop99p == nil {
		return nil, false
	}
	return o.IotApmHostTop99p, true
}

// HasIotApmHostTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasIotApmHostTop99p() bool {
	return o != nil && o.IotApmHostTop99p != nil
}

// SetIotApmHostTop99p gets a reference to the given int64 and assigns it to the IotApmHostTop99p field.
func (o *UsageSummaryDate) SetIotApmHostTop99p(v int64) {
	o.IotApmHostTop99p = &v
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

// GetLlmObservability15dayRetentionSpansSum returns the LlmObservability15dayRetentionSpansSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetLlmObservability15dayRetentionSpansSum() int64 {
	if o == nil || o.LlmObservability15dayRetentionSpansSum == nil {
		var ret int64
		return ret
	}
	return *o.LlmObservability15dayRetentionSpansSum
}

// GetLlmObservability15dayRetentionSpansSumOk returns a tuple with the LlmObservability15dayRetentionSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetLlmObservability15dayRetentionSpansSumOk() (*int64, bool) {
	if o == nil || o.LlmObservability15dayRetentionSpansSum == nil {
		return nil, false
	}
	return o.LlmObservability15dayRetentionSpansSum, true
}

// HasLlmObservability15dayRetentionSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasLlmObservability15dayRetentionSpansSum() bool {
	return o != nil && o.LlmObservability15dayRetentionSpansSum != nil
}

// SetLlmObservability15dayRetentionSpansSum gets a reference to the given int64 and assigns it to the LlmObservability15dayRetentionSpansSum field.
func (o *UsageSummaryDate) SetLlmObservability15dayRetentionSpansSum(v int64) {
	o.LlmObservability15dayRetentionSpansSum = &v
}

// GetLlmObservability30dayRetentionSpansSum returns the LlmObservability30dayRetentionSpansSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetLlmObservability30dayRetentionSpansSum() int64 {
	if o == nil || o.LlmObservability30dayRetentionSpansSum == nil {
		var ret int64
		return ret
	}
	return *o.LlmObservability30dayRetentionSpansSum
}

// GetLlmObservability30dayRetentionSpansSumOk returns a tuple with the LlmObservability30dayRetentionSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetLlmObservability30dayRetentionSpansSumOk() (*int64, bool) {
	if o == nil || o.LlmObservability30dayRetentionSpansSum == nil {
		return nil, false
	}
	return o.LlmObservability30dayRetentionSpansSum, true
}

// HasLlmObservability30dayRetentionSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasLlmObservability30dayRetentionSpansSum() bool {
	return o != nil && o.LlmObservability30dayRetentionSpansSum != nil
}

// SetLlmObservability30dayRetentionSpansSum gets a reference to the given int64 and assigns it to the LlmObservability30dayRetentionSpansSum field.
func (o *UsageSummaryDate) SetLlmObservability30dayRetentionSpansSum(v int64) {
	o.LlmObservability30dayRetentionSpansSum = &v
}

// GetLlmObservability60dayRetentionSpansSum returns the LlmObservability60dayRetentionSpansSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetLlmObservability60dayRetentionSpansSum() int64 {
	if o == nil || o.LlmObservability60dayRetentionSpansSum == nil {
		var ret int64
		return ret
	}
	return *o.LlmObservability60dayRetentionSpansSum
}

// GetLlmObservability60dayRetentionSpansSumOk returns a tuple with the LlmObservability60dayRetentionSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetLlmObservability60dayRetentionSpansSumOk() (*int64, bool) {
	if o == nil || o.LlmObservability60dayRetentionSpansSum == nil {
		return nil, false
	}
	return o.LlmObservability60dayRetentionSpansSum, true
}

// HasLlmObservability60dayRetentionSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasLlmObservability60dayRetentionSpansSum() bool {
	return o != nil && o.LlmObservability60dayRetentionSpansSum != nil
}

// SetLlmObservability60dayRetentionSpansSum gets a reference to the given int64 and assigns it to the LlmObservability60dayRetentionSpansSum field.
func (o *UsageSummaryDate) SetLlmObservability60dayRetentionSpansSum(v int64) {
	o.LlmObservability60dayRetentionSpansSum = &v
}

// GetLlmObservability90dayRetentionSpansSum returns the LlmObservability90dayRetentionSpansSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetLlmObservability90dayRetentionSpansSum() int64 {
	if o == nil || o.LlmObservability90dayRetentionSpansSum == nil {
		var ret int64
		return ret
	}
	return *o.LlmObservability90dayRetentionSpansSum
}

// GetLlmObservability90dayRetentionSpansSumOk returns a tuple with the LlmObservability90dayRetentionSpansSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetLlmObservability90dayRetentionSpansSumOk() (*int64, bool) {
	if o == nil || o.LlmObservability90dayRetentionSpansSum == nil {
		return nil, false
	}
	return o.LlmObservability90dayRetentionSpansSum, true
}

// HasLlmObservability90dayRetentionSpansSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasLlmObservability90dayRetentionSpansSum() bool {
	return o != nil && o.LlmObservability90dayRetentionSpansSum != nil
}

// SetLlmObservability90dayRetentionSpansSum gets a reference to the given int64 and assigns it to the LlmObservability90dayRetentionSpansSum field.
func (o *UsageSummaryDate) SetLlmObservability90dayRetentionSpansSum(v int64) {
	o.LlmObservability90dayRetentionSpansSum = &v
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

// GetLogsArchiveSearchGbScannedSum returns the LogsArchiveSearchGbScannedSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetLogsArchiveSearchGbScannedSum() int64 {
	if o == nil || o.LogsArchiveSearchGbScannedSum == nil {
		var ret int64
		return ret
	}
	return *o.LogsArchiveSearchGbScannedSum
}

// GetLogsArchiveSearchGbScannedSumOk returns a tuple with the LogsArchiveSearchGbScannedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetLogsArchiveSearchGbScannedSumOk() (*int64, bool) {
	if o == nil || o.LogsArchiveSearchGbScannedSum == nil {
		return nil, false
	}
	return o.LogsArchiveSearchGbScannedSum, true
}

// HasLogsArchiveSearchGbScannedSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasLogsArchiveSearchGbScannedSum() bool {
	return o != nil && o.LogsArchiveSearchGbScannedSum != nil
}

// SetLogsArchiveSearchGbScannedSum gets a reference to the given int64 and assigns it to the LogsArchiveSearchGbScannedSum field.
func (o *UsageSummaryDate) SetLogsArchiveSearchGbScannedSum(v int64) {
	o.LogsArchiveSearchGbScannedSum = &v
}

// GetMetricNamesSum returns the MetricNamesSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetMetricNamesSum() int64 {
	if o == nil || o.MetricNamesSum == nil {
		var ret int64
		return ret
	}
	return *o.MetricNamesSum
}

// GetMetricNamesSumOk returns a tuple with the MetricNamesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetMetricNamesSumOk() (*int64, bool) {
	if o == nil || o.MetricNamesSum == nil {
		return nil, false
	}
	return o.MetricNamesSum, true
}

// HasMetricNamesSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasMetricNamesSum() bool {
	return o != nil && o.MetricNamesSum != nil
}

// SetMetricNamesSum gets a reference to the given int64 and assigns it to the MetricNamesSum field.
func (o *UsageSummaryDate) SetMetricNamesSum(v int64) {
	o.MetricNamesSum = &v
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

// GetNetworkPathSum returns the NetworkPathSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetNetworkPathSum() int64 {
	if o == nil || o.NetworkPathSum == nil {
		var ret int64
		return ret
	}
	return *o.NetworkPathSum
}

// GetNetworkPathSumOk returns a tuple with the NetworkPathSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetNetworkPathSumOk() (*int64, bool) {
	if o == nil || o.NetworkPathSum == nil {
		return nil, false
	}
	return o.NetworkPathSum, true
}

// HasNetworkPathSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasNetworkPathSum() bool {
	return o != nil && o.NetworkPathSum != nil
}

// SetNetworkPathSum gets a reference to the given int64 and assigns it to the NetworkPathSum field.
func (o *UsageSummaryDate) SetNetworkPathSum(v int64) {
	o.NetworkPathSum = &v
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

// GetServerlessAppsApmApmGcpGkeAutopilotPodsAvg returns the ServerlessAppsApmApmGcpGkeAutopilotPodsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmApmGcpGkeAutopilotPodsAvg() int64 {
	if o == nil || o.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg
}

// GetServerlessAppsApmApmGcpGkeAutopilotPodsAvgOk returns a tuple with the ServerlessAppsApmApmGcpGkeAutopilotPodsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmApmGcpGkeAutopilotPodsAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg, true
}

// HasServerlessAppsApmApmGcpGkeAutopilotPodsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmApmGcpGkeAutopilotPodsAvg() bool {
	return o != nil && o.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg != nil
}

// SetServerlessAppsApmApmGcpGkeAutopilotPodsAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmApmGcpGkeAutopilotPodsAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmApmGcpGkeAutopilotPodsAvg(v int64) {
	o.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg = &v
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

// GetServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg returns the ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg() int64 {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg
}

// GetServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvgOk returns a tuple with the ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg, true
}

// HasServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg() bool {
	return o != nil && o.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg != nil
}

// SetServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg gets a reference to the given int64 and assigns it to the ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg field.
func (o *UsageSummaryDate) SetServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg(v int64) {
	o.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg = &v
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

// GetServerlessAppsDsmFargateTasksAvg returns the ServerlessAppsDsmFargateTasksAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsDsmFargateTasksAvg() int64 {
	if o == nil || o.ServerlessAppsDsmFargateTasksAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsDsmFargateTasksAvg
}

// GetServerlessAppsDsmFargateTasksAvgOk returns a tuple with the ServerlessAppsDsmFargateTasksAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsDsmFargateTasksAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsDsmFargateTasksAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsDsmFargateTasksAvg, true
}

// HasServerlessAppsDsmFargateTasksAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsDsmFargateTasksAvg() bool {
	return o != nil && o.ServerlessAppsDsmFargateTasksAvg != nil
}

// SetServerlessAppsDsmFargateTasksAvg gets a reference to the given int64 and assigns it to the ServerlessAppsDsmFargateTasksAvg field.
func (o *UsageSummaryDate) SetServerlessAppsDsmFargateTasksAvg(v int64) {
	o.ServerlessAppsDsmFargateTasksAvg = &v
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

// GetServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg returns the ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg() int64 {
	if o == nil || o.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg
}

// GetServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvgOk returns a tuple with the ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg, true
}

// HasServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg() bool {
	return o != nil && o.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg != nil
}

// SetServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg gets a reference to the given int64 and assigns it to the ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg field.
func (o *UsageSummaryDate) SetServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg(v int64) {
	o.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg = &v
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

// GetServerlessAppsInfraGcpGkeAutopilotPodsAvg returns the ServerlessAppsInfraGcpGkeAutopilotPodsAvg field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetServerlessAppsInfraGcpGkeAutopilotPodsAvg() int64 {
	if o == nil || o.ServerlessAppsInfraGcpGkeAutopilotPodsAvg == nil {
		var ret int64
		return ret
	}
	return *o.ServerlessAppsInfraGcpGkeAutopilotPodsAvg
}

// GetServerlessAppsInfraGcpGkeAutopilotPodsAvgOk returns a tuple with the ServerlessAppsInfraGcpGkeAutopilotPodsAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetServerlessAppsInfraGcpGkeAutopilotPodsAvgOk() (*int64, bool) {
	if o == nil || o.ServerlessAppsInfraGcpGkeAutopilotPodsAvg == nil {
		return nil, false
	}
	return o.ServerlessAppsInfraGcpGkeAutopilotPodsAvg, true
}

// HasServerlessAppsInfraGcpGkeAutopilotPodsAvg returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasServerlessAppsInfraGcpGkeAutopilotPodsAvg() bool {
	return o != nil && o.ServerlessAppsInfraGcpGkeAutopilotPodsAvg != nil
}

// SetServerlessAppsInfraGcpGkeAutopilotPodsAvg gets a reference to the given int64 and assigns it to the ServerlessAppsInfraGcpGkeAutopilotPodsAvg field.
func (o *UsageSummaryDate) SetServerlessAppsInfraGcpGkeAutopilotPodsAvg(v int64) {
	o.ServerlessAppsInfraGcpGkeAutopilotPodsAvg = &v
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

// GetSiem12moRetentionSum returns the Siem12moRetentionSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSiem12moRetentionSum() int64 {
	if o == nil || o.Siem12moRetentionSum == nil {
		var ret int64
		return ret
	}
	return *o.Siem12moRetentionSum
}

// GetSiem12moRetentionSumOk returns a tuple with the Siem12moRetentionSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSiem12moRetentionSumOk() (*int64, bool) {
	if o == nil || o.Siem12moRetentionSum == nil {
		return nil, false
	}
	return o.Siem12moRetentionSum, true
}

// HasSiem12moRetentionSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSiem12moRetentionSum() bool {
	return o != nil && o.Siem12moRetentionSum != nil
}

// SetSiem12moRetentionSum gets a reference to the given int64 and assigns it to the Siem12moRetentionSum field.
func (o *UsageSummaryDate) SetSiem12moRetentionSum(v int64) {
	o.Siem12moRetentionSum = &v
}

// GetSiem6moRetentionSum returns the Siem6moRetentionSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSiem6moRetentionSum() int64 {
	if o == nil || o.Siem6moRetentionSum == nil {
		var ret int64
		return ret
	}
	return *o.Siem6moRetentionSum
}

// GetSiem6moRetentionSumOk returns a tuple with the Siem6moRetentionSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSiem6moRetentionSumOk() (*int64, bool) {
	if o == nil || o.Siem6moRetentionSum == nil {
		return nil, false
	}
	return o.Siem6moRetentionSum, true
}

// HasSiem6moRetentionSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSiem6moRetentionSum() bool {
	return o != nil && o.Siem6moRetentionSum != nil
}

// SetSiem6moRetentionSum gets a reference to the given int64 and assigns it to the Siem6moRetentionSum field.
func (o *UsageSummaryDate) SetSiem6moRetentionSum(v int64) {
	o.Siem6moRetentionSum = &v
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

// GetSnmpDeviceCountSum returns the SnmpDeviceCountSum field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSnmpDeviceCountSum() int64 {
	if o == nil || o.SnmpDeviceCountSum == nil {
		var ret int64
		return ret
	}
	return *o.SnmpDeviceCountSum
}

// GetSnmpDeviceCountSumOk returns a tuple with the SnmpDeviceCountSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSnmpDeviceCountSumOk() (*int64, bool) {
	if o == nil || o.SnmpDeviceCountSum == nil {
		return nil, false
	}
	return o.SnmpDeviceCountSum, true
}

// HasSnmpDeviceCountSum returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSnmpDeviceCountSum() bool {
	return o != nil && o.SnmpDeviceCountSum != nil
}

// SetSnmpDeviceCountSum gets a reference to the given int64 and assigns it to the SnmpDeviceCountSum field.
func (o *UsageSummaryDate) SetSnmpDeviceCountSum(v int64) {
	o.SnmpDeviceCountSum = &v
}

// GetSnmpDeviceCountTop99p returns the SnmpDeviceCountTop99p field value if set, zero value otherwise.
func (o *UsageSummaryDate) GetSnmpDeviceCountTop99p() int64 {
	if o == nil || o.SnmpDeviceCountTop99p == nil {
		var ret int64
		return ret
	}
	return *o.SnmpDeviceCountTop99p
}

// GetSnmpDeviceCountTop99pOk returns a tuple with the SnmpDeviceCountTop99p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryDate) GetSnmpDeviceCountTop99pOk() (*int64, bool) {
	if o == nil || o.SnmpDeviceCountTop99p == nil {
		return nil, false
	}
	return o.SnmpDeviceCountTop99p, true
}

// HasSnmpDeviceCountTop99p returns a boolean if a field has been set.
func (o *UsageSummaryDate) HasSnmpDeviceCountTop99p() bool {
	return o != nil && o.SnmpDeviceCountTop99p != nil
}

// SetSnmpDeviceCountTop99p gets a reference to the given int64 and assigns it to the SnmpDeviceCountTop99p field.
func (o *UsageSummaryDate) SetSnmpDeviceCountTop99p(v int64) {
	o.SnmpDeviceCountTop99p = &v
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
	if o.AiCreditsAgentBuilderAiCreditsSum != nil {
		toSerialize["ai_credits_agent_builder_ai_credits_sum"] = o.AiCreditsAgentBuilderAiCreditsSum
	}
	if o.AiCreditsBitsAssistantAiCreditsSum != nil {
		toSerialize["ai_credits_bits_assistant_ai_credits_sum"] = o.AiCreditsBitsAssistantAiCreditsSum
	}
	if o.AiCreditsBitsDevAiCreditsSum != nil {
		toSerialize["ai_credits_bits_dev_ai_credits_sum"] = o.AiCreditsBitsDevAiCreditsSum
	}
	if o.AiCreditsBitsSreAiCreditsSum != nil {
		toSerialize["ai_credits_bits_sre_ai_credits_sum"] = o.AiCreditsBitsSreAiCreditsSum
	}
	if o.AiCreditsSum != nil {
		toSerialize["ai_credits_sum"] = o.AiCreditsSum
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
	if o.AuditTrailEventForwardingEventsSum != nil {
		toSerialize["audit_trail_event_forwarding_events_sum"] = o.AuditTrailEventForwardingEventsSum
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
	if o.CcmAnthropicSpendLast != nil {
		toSerialize["ccm_anthropic_spend_last"] = o.CcmAnthropicSpendLast
	}
	if o.CcmAwsSpendLast != nil {
		toSerialize["ccm_aws_spend_last"] = o.CcmAwsSpendLast
	}
	if o.CcmAzureSpendLast != nil {
		toSerialize["ccm_azure_spend_last"] = o.CcmAzureSpendLast
	}
	if o.CcmConfluentSpendLast != nil {
		toSerialize["ccm_confluent_spend_last"] = o.CcmConfluentSpendLast
	}
	if o.CcmDatabricksSpendLast != nil {
		toSerialize["ccm_databricks_spend_last"] = o.CcmDatabricksSpendLast
	}
	if o.CcmElasticSpendLast != nil {
		toSerialize["ccm_elastic_spend_last"] = o.CcmElasticSpendLast
	}
	if o.CcmFastlySpendLast != nil {
		toSerialize["ccm_fastly_spend_last"] = o.CcmFastlySpendLast
	}
	if o.CcmGcpSpendLast != nil {
		toSerialize["ccm_gcp_spend_last"] = o.CcmGcpSpendLast
	}
	if o.CcmGithubSpendLast != nil {
		toSerialize["ccm_github_spend_last"] = o.CcmGithubSpendLast
	}
	if o.CcmMongodbSpendLast != nil {
		toSerialize["ccm_mongodb_spend_last"] = o.CcmMongodbSpendLast
	}
	if o.CcmOciSpendLast != nil {
		toSerialize["ccm_oci_spend_last"] = o.CcmOciSpendLast
	}
	if o.CcmOpenaiSpendLast != nil {
		toSerialize["ccm_openai_spend_last"] = o.CcmOpenaiSpendLast
	}
	if o.CcmSnowflakeSpendLast != nil {
		toSerialize["ccm_snowflake_spend_last"] = o.CcmSnowflakeSpendLast
	}
	if o.CcmSpendMonitoredEntLast != nil {
		toSerialize["ccm_spend_monitored_ent_last"] = o.CcmSpendMonitoredEntLast
	}
	if o.CcmSpendMonitoredProLast != nil {
		toSerialize["ccm_spend_monitored_pro_last"] = o.CcmSpendMonitoredProLast
	}
	if o.CcmTwilioSpendLast != nil {
		toSerialize["ccm_twilio_spend_last"] = o.CcmTwilioSpendLast
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
	if o.CloudSiemIndexedLogsSum != nil {
		toSerialize["cloud_siem_indexed_logs_sum"] = o.CloudSiemIndexedLogsSum
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
	if o.CsmHostEnterpriseOciHostCountTop99p != nil {
		toSerialize["csm_host_enterprise_oci_host_count_top99p"] = o.CsmHostEnterpriseOciHostCountTop99p
	}
	if o.CsmHostEnterpriseTotalHostCountTop99p != nil {
		toSerialize["csm_host_enterprise_total_host_count_top99p"] = o.CsmHostEnterpriseTotalHostCountTop99p
	}
	if o.CsmHostProHostsAgentlessScannersSum != nil {
		toSerialize["csm_host_pro_hosts_agentless_scanners_sum"] = o.CsmHostProHostsAgentlessScannersSum
	}
	if o.CsmHostProHostsAgentlessScannersTop99p != nil {
		toSerialize["csm_host_pro_hosts_agentless_scanners_top99p"] = o.CsmHostProHostsAgentlessScannersTop99p
	}
	if o.CsmHostProOciHostCountTop99p != nil {
		toSerialize["csm_host_pro_oci_host_count_top99p"] = o.CsmHostProOciHostCountTop99p
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
	if o.CspmHostsAgentlessScannersSum != nil {
		toSerialize["cspm_hosts_agentless_scanners_sum"] = o.CspmHostsAgentlessScannersSum
	}
	if o.CspmHostsAgentlessScannersTop99p != nil {
		toSerialize["cspm_hosts_agentless_scanners_top99p"] = o.CspmHostsAgentlessScannersTop99p
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
	if o.DataStreamMonitoringHostCountSum != nil {
		toSerialize["data_stream_monitoring_host_count_sum"] = o.DataStreamMonitoringHostCountSum
	}
	if o.DataStreamMonitoringHostCountTop99p != nil {
		toSerialize["data_stream_monitoring_host_count_top99p"] = o.DataStreamMonitoringHostCountTop99p
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
	if o.DoJobsMonitoringOrchestratorsJobHoursSum != nil {
		toSerialize["do_jobs_monitoring_orchestrators_job_hours_sum"] = o.DoJobsMonitoringOrchestratorsJobHoursSum
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
	if o.EphInfraHostBasicInfraBasicAgentSum != nil {
		toSerialize["eph_infra_host_basic_infra_basic_agent_sum"] = o.EphInfraHostBasicInfraBasicAgentSum
	}
	if o.EphInfraHostBasicInfraBasicVsphereSum != nil {
		toSerialize["eph_infra_host_basic_infra_basic_vsphere_sum"] = o.EphInfraHostBasicInfraBasicVsphereSum
	}
	if o.EphInfraHostBasicSum != nil {
		toSerialize["eph_infra_host_basic_sum"] = o.EphInfraHostBasicSum
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
	if o.FeatureFlagsConfigRequestsSum != nil {
		toSerialize["feature_flags_config_requests_sum"] = o.FeatureFlagsConfigRequestsSum
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
	if o.IndexedPointsSum != nil {
		toSerialize["indexed_points_sum"] = o.IndexedPointsSum
	}
	if o.InfraCpuAvg != nil {
		toSerialize["infra_cpu_avg"] = o.InfraCpuAvg
	}
	if o.InfraCpuDefaultInfraHostVcpuAgentAvg != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_agent_avg"] = o.InfraCpuDefaultInfraHostVcpuAgentAvg
	}
	if o.InfraCpuDefaultInfraHostVcpuAgentBasicAvg != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_agent_basic_avg"] = o.InfraCpuDefaultInfraHostVcpuAgentBasicAvg
	}
	if o.InfraCpuDefaultInfraHostVcpuAgentBasicSum != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_agent_basic_sum"] = o.InfraCpuDefaultInfraHostVcpuAgentBasicSum
	}
	if o.InfraCpuDefaultInfraHostVcpuAgentSum != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_agent_sum"] = o.InfraCpuDefaultInfraHostVcpuAgentSum
	}
	if o.InfraCpuDefaultInfraHostVcpuAwsAvg != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_aws_avg"] = o.InfraCpuDefaultInfraHostVcpuAwsAvg
	}
	if o.InfraCpuDefaultInfraHostVcpuAwsSum != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_aws_sum"] = o.InfraCpuDefaultInfraHostVcpuAwsSum
	}
	if o.InfraCpuDefaultInfraHostVcpuAzureAvg != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_azure_avg"] = o.InfraCpuDefaultInfraHostVcpuAzureAvg
	}
	if o.InfraCpuDefaultInfraHostVcpuAzureSum != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_azure_sum"] = o.InfraCpuDefaultInfraHostVcpuAzureSum
	}
	if o.InfraCpuDefaultInfraHostVcpuGcpAvg != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_gcp_avg"] = o.InfraCpuDefaultInfraHostVcpuGcpAvg
	}
	if o.InfraCpuDefaultInfraHostVcpuGcpSum != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_gcp_sum"] = o.InfraCpuDefaultInfraHostVcpuGcpSum
	}
	if o.InfraCpuDefaultInfraHostVcpuNutanixAvg != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_nutanix_avg"] = o.InfraCpuDefaultInfraHostVcpuNutanixAvg
	}
	if o.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_nutanix_basic_avg"] = o.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg
	}
	if o.InfraCpuDefaultInfraHostVcpuNutanixBasicSum != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_nutanix_basic_sum"] = o.InfraCpuDefaultInfraHostVcpuNutanixBasicSum
	}
	if o.InfraCpuDefaultInfraHostVcpuNutanixSum != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_nutanix_sum"] = o.InfraCpuDefaultInfraHostVcpuNutanixSum
	}
	if o.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_opentelemetry_avg"] = o.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg
	}
	if o.InfraCpuDefaultInfraHostVcpuOpentelemetrySum != nil {
		toSerialize["infra_cpu_default_infra_host_vcpu_opentelemetry_sum"] = o.InfraCpuDefaultInfraHostVcpuOpentelemetrySum
	}
	if o.InfraCpuObservedInfraHostVcpuAgentAvg != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_agent_avg"] = o.InfraCpuObservedInfraHostVcpuAgentAvg
	}
	if o.InfraCpuObservedInfraHostVcpuAgentSum != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_agent_sum"] = o.InfraCpuObservedInfraHostVcpuAgentSum
	}
	if o.InfraCpuObservedInfraHostVcpuAwsAvg != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_aws_avg"] = o.InfraCpuObservedInfraHostVcpuAwsAvg
	}
	if o.InfraCpuObservedInfraHostVcpuAwsSum != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_aws_sum"] = o.InfraCpuObservedInfraHostVcpuAwsSum
	}
	if o.InfraCpuObservedInfraHostVcpuAzureAvg != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_azure_avg"] = o.InfraCpuObservedInfraHostVcpuAzureAvg
	}
	if o.InfraCpuObservedInfraHostVcpuAzureSum != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_azure_sum"] = o.InfraCpuObservedInfraHostVcpuAzureSum
	}
	if o.InfraCpuObservedInfraHostVcpuGcpAvg != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_gcp_avg"] = o.InfraCpuObservedInfraHostVcpuGcpAvg
	}
	if o.InfraCpuObservedInfraHostVcpuGcpSum != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_gcp_sum"] = o.InfraCpuObservedInfraHostVcpuGcpSum
	}
	if o.InfraCpuObservedInfraHostVcpuNutanixAvg != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_nutanix_avg"] = o.InfraCpuObservedInfraHostVcpuNutanixAvg
	}
	if o.InfraCpuObservedInfraHostVcpuNutanixSum != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_nutanix_sum"] = o.InfraCpuObservedInfraHostVcpuNutanixSum
	}
	if o.InfraCpuObservedInfraHostVcpuOpentelemetryAvg != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_opentelemetry_avg"] = o.InfraCpuObservedInfraHostVcpuOpentelemetryAvg
	}
	if o.InfraCpuObservedInfraHostVcpuOpentelemetrySum != nil {
		toSerialize["infra_cpu_observed_infra_host_vcpu_opentelemetry_sum"] = o.InfraCpuObservedInfraHostVcpuOpentelemetrySum
	}
	if o.InfraCpuSum != nil {
		toSerialize["infra_cpu_sum"] = o.InfraCpuSum
	}
	if o.InfraEdgeMonitoringDevicesTop99p != nil {
		toSerialize["infra_edge_monitoring_devices_top99p"] = o.InfraEdgeMonitoringDevicesTop99p
	}
	if o.InfraHostBasicInfraBasicAgentTop99p != nil {
		toSerialize["infra_host_basic_infra_basic_agent_top99p"] = o.InfraHostBasicInfraBasicAgentTop99p
	}
	if o.InfraHostBasicInfraBasicVsphereTop99p != nil {
		toSerialize["infra_host_basic_infra_basic_vsphere_top99p"] = o.InfraHostBasicInfraBasicVsphereTop99p
	}
	if o.InfraHostBasicTop99p != nil {
		toSerialize["infra_host_basic_top99p"] = o.InfraHostBasicTop99p
	}
	if o.InfraHostTop99p != nil {
		toSerialize["infra_host_top99p"] = o.InfraHostTop99p
	}
	if o.InfraStorageMgmtObjectsCountAvg != nil {
		toSerialize["infra_storage_mgmt_objects_count_avg"] = o.InfraStorageMgmtObjectsCountAvg
	}
	if o.IngestPointsSum != nil {
		toSerialize["ingest_points_sum"] = o.IngestPointsSum
	}
	if o.IngestedEventsBytesSum != nil {
		toSerialize["ingested_events_bytes_sum"] = o.IngestedEventsBytesSum
	}
	if o.IotApmHostSum != nil {
		toSerialize["iot_apm_host_sum"] = o.IotApmHostSum
	}
	if o.IotApmHostTop99p != nil {
		toSerialize["iot_apm_host_top99p"] = o.IotApmHostTop99p
	}
	if o.IotDeviceSum != nil {
		toSerialize["iot_device_sum"] = o.IotDeviceSum
	}
	if o.IotDeviceTop99p != nil {
		toSerialize["iot_device_top99p"] = o.IotDeviceTop99p
	}
	if o.LlmObservability15dayRetentionSpansSum != nil {
		toSerialize["llm_observability_15day_retention_spans_sum"] = o.LlmObservability15dayRetentionSpansSum
	}
	if o.LlmObservability30dayRetentionSpansSum != nil {
		toSerialize["llm_observability_30day_retention_spans_sum"] = o.LlmObservability30dayRetentionSpansSum
	}
	if o.LlmObservability60dayRetentionSpansSum != nil {
		toSerialize["llm_observability_60day_retention_spans_sum"] = o.LlmObservability60dayRetentionSpansSum
	}
	if o.LlmObservability90dayRetentionSpansSum != nil {
		toSerialize["llm_observability_90day_retention_spans_sum"] = o.LlmObservability90dayRetentionSpansSum
	}
	if o.LlmObservabilityMinSpendSum != nil {
		toSerialize["llm_observability_min_spend_sum"] = o.LlmObservabilityMinSpendSum
	}
	if o.LlmObservabilitySum != nil {
		toSerialize["llm_observability_sum"] = o.LlmObservabilitySum
	}
	if o.LogsArchiveSearchGbScannedSum != nil {
		toSerialize["logs_archive_search_gb_scanned_sum"] = o.LogsArchiveSearchGbScannedSum
	}
	if o.MetricNamesSum != nil {
		toSerialize["metric_names_sum"] = o.MetricNamesSum
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
	if o.NetworkPathSum != nil {
		toSerialize["network_path_sum"] = o.NetworkPathSum
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
	if o.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg != nil {
		toSerialize["serverless_apps_apm_apm_gcp_gke_autopilot_pods_avg"] = o.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg
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
	if o.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg != nil {
		toSerialize["serverless_apps_apm_excl_fargate_apm_gcp_gke_autopilot_pods_avg"] = o.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg
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
	if o.ServerlessAppsDsmFargateTasksAvg != nil {
		toSerialize["serverless_apps_dsm_fargate_tasks_avg"] = o.ServerlessAppsDsmFargateTasksAvg
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
	if o.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg != nil {
		toSerialize["serverless_apps_excl_fargate_infra_gcp_gke_autopilot_pods_avg"] = o.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg
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
	if o.ServerlessAppsInfraGcpGkeAutopilotPodsAvg != nil {
		toSerialize["serverless_apps_infra_gcp_gke_autopilot_pods_avg"] = o.ServerlessAppsInfraGcpGkeAutopilotPodsAvg
	}
	if o.ServerlessAppsTotalCountAvg != nil {
		toSerialize["serverless_apps_total_count_avg"] = o.ServerlessAppsTotalCountAvg
	}
	if o.Siem12moRetentionSum != nil {
		toSerialize["siem_12mo_retention_sum"] = o.Siem12moRetentionSum
	}
	if o.Siem6moRetentionSum != nil {
		toSerialize["siem_6mo_retention_sum"] = o.Siem6moRetentionSum
	}
	if o.SiemAnalyzedLogsAddOnCountSum != nil {
		toSerialize["siem_analyzed_logs_add_on_count_sum"] = o.SiemAnalyzedLogsAddOnCountSum
	}
	if o.SnmpDeviceCountSum != nil {
		toSerialize["snmp_device_count_sum"] = o.SnmpDeviceCountSum
	}
	if o.SnmpDeviceCountTop99p != nil {
		toSerialize["snmp_device_count_top99p"] = o.SnmpDeviceCountTop99p
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

	typedKeys := map[string]bool{"agent_host_top99p": true, "ai_credits_agent_builder_ai_credits_sum": true, "ai_credits_bits_assistant_ai_credits_sum": true, "ai_credits_bits_dev_ai_credits_sum": true, "ai_credits_bits_sre_ai_credits_sum": true, "ai_credits_sum": true, "apm_azure_app_service_host_top99p": true, "apm_devsecops_host_top99p": true, "apm_enterprise_standalone_hosts_top99p": true, "apm_fargate_count_avg": true, "apm_host_top99p": true, "apm_pro_standalone_hosts_top99p": true, "appsec_fargate_count_avg": true, "asm_serverless_sum": true, "audit_logs_lines_indexed_sum": true, "audit_trail_enabled_hwm": true, "audit_trail_event_forwarding_events_sum": true, "avg_profiled_fargate_tasks": true, "aws_host_top99p": true, "aws_lambda_func_count": true, "aws_lambda_invocations_sum": true, "azure_app_service_top99p": true, "billable_ingested_bytes_sum": true, "bits_ai_investigations_sum": true, "browser_rum_lite_session_count_sum": true, "browser_rum_replay_session_count_sum": true, "browser_rum_units_sum": true, "ccm_anthropic_spend_last": true, "ccm_aws_spend_last": true, "ccm_azure_spend_last": true, "ccm_confluent_spend_last": true, "ccm_databricks_spend_last": true, "ccm_elastic_spend_last": true, "ccm_fastly_spend_last": true, "ccm_gcp_spend_last": true, "ccm_github_spend_last": true, "ccm_mongodb_spend_last": true, "ccm_oci_spend_last": true, "ccm_openai_spend_last": true, "ccm_snowflake_spend_last": true, "ccm_spend_monitored_ent_last": true, "ccm_spend_monitored_pro_last": true, "ccm_twilio_spend_last": true, "ci_pipeline_indexed_spans_sum": true, "ci_test_indexed_spans_sum": true, "ci_visibility_itr_committers_hwm": true, "ci_visibility_pipeline_committers_hwm": true, "ci_visibility_test_committers_hwm": true, "cloud_cost_management_aws_host_count_avg": true, "cloud_cost_management_azure_host_count_avg": true, "cloud_cost_management_gcp_host_count_avg": true, "cloud_cost_management_host_count_avg": true, "cloud_cost_management_oci_host_count_avg": true, "cloud_siem_events_sum": true, "cloud_siem_indexed_logs_sum": true, "code_analysis_sa_committers_hwm": true, "code_analysis_sca_committers_hwm": true, "code_security_host_top99p": true, "container_avg": true, "container_excl_agent_avg": true, "container_hwm": true, "csm_container_enterprise_compliance_count_sum": true, "csm_container_enterprise_cws_count_sum": true, "csm_container_enterprise_total_count_sum": true, "csm_host_enterprise_aas_host_count_top99p": true, "csm_host_enterprise_aws_host_count_top99p": true, "csm_host_enterprise_azure_host_count_top99p": true, "csm_host_enterprise_compliance_host_count_top99p": true, "csm_host_enterprise_cws_host_count_top99p": true, "csm_host_enterprise_gcp_host_count_top99p": true, "csm_host_enterprise_oci_host_count_top99p": true, "csm_host_enterprise_total_host_count_top99p": true, "csm_host_pro_hosts_agentless_scanners_sum": true, "csm_host_pro_hosts_agentless_scanners_top99p": true, "csm_host_pro_oci_host_count_top99p": true, "cspm_aas_host_top99p": true, "cspm_aws_host_top99p": true, "cspm_azure_host_top99p": true, "cspm_container_avg": true, "cspm_container_hwm": true, "cspm_gcp_host_top99p": true, "cspm_host_top99p": true, "cspm_hosts_agentless_scanners_sum": true, "cspm_hosts_agentless_scanners_top99p": true, "custom_ts_avg": true, "cws_container_count_avg": true, "cws_fargate_task_avg": true, "cws_host_top99p": true, "data_jobs_monitoring_host_hr_sum": true, "data_stream_monitoring_host_count_sum": true, "data_stream_monitoring_host_count_top99p": true, "date": true, "dbm_host_top99p": true, "dbm_queries_count_avg": true, "do_jobs_monitoring_orchestrators_job_hours_sum": true, "eph_infra_host_agent_sum": true, "eph_infra_host_alibaba_sum": true, "eph_infra_host_aws_sum": true, "eph_infra_host_azure_sum": true, "eph_infra_host_basic_infra_basic_agent_sum": true, "eph_infra_host_basic_infra_basic_vsphere_sum": true, "eph_infra_host_basic_sum": true, "eph_infra_host_ent_sum": true, "eph_infra_host_gcp_sum": true, "eph_infra_host_heroku_sum": true, "eph_infra_host_only_aas_sum": true, "eph_infra_host_only_vsphere_sum": true, "eph_infra_host_opentelemetry_apm_sum": true, "eph_infra_host_opentelemetry_sum": true, "eph_infra_host_pro_sum": true, "eph_infra_host_proplus_sum": true, "eph_infra_host_proxmox_sum": true, "error_tracking_apm_error_events_sum": true, "error_tracking_error_events_sum": true, "error_tracking_events_sum": true, "error_tracking_rum_error_events_sum": true, "event_management_correlation_correlated_events_sum": true, "event_management_correlation_correlated_related_events_sum": true, "event_management_correlation_sum": true, "fargate_container_profiler_profiling_fargate_avg": true, "fargate_container_profiler_profiling_fargate_eks_avg": true, "fargate_tasks_count_avg": true, "fargate_tasks_count_hwm": true, "feature_flags_config_requests_sum": true, "flex_logs_compute_large_avg": true, "flex_logs_compute_medium_avg": true, "flex_logs_compute_small_avg": true, "flex_logs_compute_xlarge_avg": true, "flex_logs_compute_xsmall_avg": true, "flex_logs_starter_avg": true, "flex_logs_starter_storage_index_avg": true, "flex_logs_starter_storage_retention_adjustment_avg": true, "flex_stored_logs_avg": true, "forwarding_events_bytes_sum": true, "gcp_host_top99p": true, "heroku_host_top99p": true, "incident_management_monthly_active_users_hwm": true, "incident_management_seats_hwm": true, "indexed_events_count_sum": true, "indexed_points_sum": true, "infra_cpu_avg": true, "infra_cpu_default_infra_host_vcpu_agent_avg": true, "infra_cpu_default_infra_host_vcpu_agent_basic_avg": true, "infra_cpu_default_infra_host_vcpu_agent_basic_sum": true, "infra_cpu_default_infra_host_vcpu_agent_sum": true, "infra_cpu_default_infra_host_vcpu_aws_avg": true, "infra_cpu_default_infra_host_vcpu_aws_sum": true, "infra_cpu_default_infra_host_vcpu_azure_avg": true, "infra_cpu_default_infra_host_vcpu_azure_sum": true, "infra_cpu_default_infra_host_vcpu_gcp_avg": true, "infra_cpu_default_infra_host_vcpu_gcp_sum": true, "infra_cpu_default_infra_host_vcpu_nutanix_avg": true, "infra_cpu_default_infra_host_vcpu_nutanix_basic_avg": true, "infra_cpu_default_infra_host_vcpu_nutanix_basic_sum": true, "infra_cpu_default_infra_host_vcpu_nutanix_sum": true, "infra_cpu_default_infra_host_vcpu_opentelemetry_avg": true, "infra_cpu_default_infra_host_vcpu_opentelemetry_sum": true, "infra_cpu_observed_infra_host_vcpu_agent_avg": true, "infra_cpu_observed_infra_host_vcpu_agent_sum": true, "infra_cpu_observed_infra_host_vcpu_aws_avg": true, "infra_cpu_observed_infra_host_vcpu_aws_sum": true, "infra_cpu_observed_infra_host_vcpu_azure_avg": true, "infra_cpu_observed_infra_host_vcpu_azure_sum": true, "infra_cpu_observed_infra_host_vcpu_gcp_avg": true, "infra_cpu_observed_infra_host_vcpu_gcp_sum": true, "infra_cpu_observed_infra_host_vcpu_nutanix_avg": true, "infra_cpu_observed_infra_host_vcpu_nutanix_sum": true, "infra_cpu_observed_infra_host_vcpu_opentelemetry_avg": true, "infra_cpu_observed_infra_host_vcpu_opentelemetry_sum": true, "infra_cpu_sum": true, "infra_edge_monitoring_devices_top99p": true, "infra_host_basic_infra_basic_agent_top99p": true, "infra_host_basic_infra_basic_vsphere_top99p": true, "infra_host_basic_top99p": true, "infra_host_top99p": true, "infra_storage_mgmt_objects_count_avg": true, "ingest_points_sum": true, "ingested_events_bytes_sum": true, "iot_apm_host_sum": true, "iot_apm_host_top99p": true, "iot_device_sum": true, "iot_device_top99p": true, "llm_observability_15day_retention_spans_sum": true, "llm_observability_30day_retention_spans_sum": true, "llm_observability_60day_retention_spans_sum": true, "llm_observability_90day_retention_spans_sum": true, "llm_observability_min_spend_sum": true, "llm_observability_sum": true, "logs_archive_search_gb_scanned_sum": true, "metric_names_sum": true, "mobile_rum_lite_session_count_sum": true, "mobile_rum_session_count_android_sum": true, "mobile_rum_session_count_flutter_sum": true, "mobile_rum_session_count_ios_sum": true, "mobile_rum_session_count_reactnative_sum": true, "mobile_rum_session_count_roku_sum": true, "mobile_rum_session_count_sum": true, "mobile_rum_units_sum": true, "ndm_netflow_events_sum": true, "netflow_indexed_events_count_sum": true, "network_device_wireless_top99p": true, "network_path_sum": true, "npm_host_top99p": true, "observability_pipelines_bytes_processed_sum": true, "oci_host_sum": true, "oci_host_top99p": true, "on_call_seat_hwm": true, "online_archive_events_count_sum": true, "opentelemetry_apm_host_top99p": true, "opentelemetry_host_top99p": true, "orgs": true, "product_analytics_sum": true, "profiling_aas_count_top99p": true, "profiling_host_top99p": true, "proxmox_host_sum": true, "proxmox_host_top99p": true, "published_app_hwm": true, "rum_browser_and_mobile_session_count": true, "rum_browser_legacy_session_count_sum": true, "rum_browser_lite_session_count_sum": true, "rum_browser_replay_session_count_sum": true, "rum_indexed_sessions_sum": true, "rum_ingested_sessions_sum": true, "rum_lite_session_count_sum": true, "rum_mobile_legacy_session_count_android_sum": true, "rum_mobile_legacy_session_count_flutter_sum": true, "rum_mobile_legacy_session_count_ios_sum": true, "rum_mobile_legacy_session_count_reactnative_sum": true, "rum_mobile_legacy_session_count_roku_sum": true, "rum_mobile_lite_session_count_android_sum": true, "rum_mobile_lite_session_count_flutter_sum": true, "rum_mobile_lite_session_count_ios_sum": true, "rum_mobile_lite_session_count_kotlinmultiplatform_sum": true, "rum_mobile_lite_session_count_reactnative_sum": true, "rum_mobile_lite_session_count_roku_sum": true, "rum_mobile_lite_session_count_unity_sum": true, "rum_mobile_replay_session_count_android_sum": true, "rum_mobile_replay_session_count_ios_sum": true, "rum_mobile_replay_session_count_kotlinmultiplatform_sum": true, "rum_mobile_replay_session_count_reactnative_sum": true, "rum_replay_session_count_sum": true, "rum_session_count_sum": true, "rum_session_replay_add_on_sum": true, "rum_total_session_count_sum": true, "rum_units_sum": true, "sca_fargate_count_avg": true, "sca_fargate_count_hwm": true, "sds_apm_scanned_bytes_sum": true, "sds_events_scanned_bytes_sum": true, "sds_logs_scanned_bytes_sum": true, "sds_rum_scanned_bytes_sum": true, "sds_total_scanned_bytes_sum": true, "serverless_apps_apm_apm_azure_appservice_instances_avg": true, "serverless_apps_apm_apm_azure_azurefunction_instances_avg": true, "serverless_apps_apm_apm_azure_containerapp_instances_avg": true, "serverless_apps_apm_apm_fargate_ecs_tasks_avg": true, "serverless_apps_apm_apm_gcp_cloudfunction_instances_avg": true, "serverless_apps_apm_apm_gcp_cloudrun_instances_avg": true, "serverless_apps_apm_apm_gcp_gke_autopilot_pods_avg": true, "serverless_apps_apm_avg": true, "serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg": true, "serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg": true, "serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg": true, "serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg": true, "serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg": true, "serverless_apps_apm_excl_fargate_apm_gcp_gke_autopilot_pods_avg": true, "serverless_apps_apm_excl_fargate_avg": true, "serverless_apps_azure_container_app_instances_avg": true, "serverless_apps_azure_count_avg": true, "serverless_apps_azure_function_app_instances_avg": true, "serverless_apps_azure_web_app_instances_avg": true, "serverless_apps_dsm_fargate_tasks_avg": true, "serverless_apps_ecs_avg": true, "serverless_apps_eks_avg": true, "serverless_apps_excl_fargate_avg": true, "serverless_apps_excl_fargate_azure_container_app_instances_avg": true, "serverless_apps_excl_fargate_azure_function_app_instances_avg": true, "serverless_apps_excl_fargate_azure_web_app_instances_avg": true, "serverless_apps_excl_fargate_google_cloud_functions_instances_avg": true, "serverless_apps_excl_fargate_google_cloud_run_instances_avg": true, "serverless_apps_excl_fargate_infra_gcp_gke_autopilot_pods_avg": true, "serverless_apps_google_cloud_functions_instances_avg": true, "serverless_apps_google_cloud_run_instances_avg": true, "serverless_apps_google_count_avg": true, "serverless_apps_infra_gcp_gke_autopilot_pods_avg": true, "serverless_apps_total_count_avg": true, "siem_12mo_retention_sum": true, "siem_6mo_retention_sum": true, "siem_analyzed_logs_add_on_count_sum": true, "snmp_device_count_sum": true, "snmp_device_count_top99p": true, "synthetics_browser_check_calls_count_sum": true, "synthetics_check_calls_count_sum": true, "synthetics_mobile_test_runs_sum": true, "synthetics_parallel_testing_max_slots_hwm": true, "trace_search_indexed_events_count_sum": true, "twol_ingested_events_bytes_sum": true, "universal_service_monitoring_host_top99p": true, "vsphere_host_top99p": true, "vuln_management_host_count_top99p": true, "workflow_executions_usage_sum": true}
	for key, value := range o.AdditionalProperties {
		if !typedKeys[key] {
			toSerialize[key] = value
		}
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageSummaryDate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentHostTop99p                                               *int64                `json:"agent_host_top99p,omitempty"`
		AiCreditsAgentBuilderAiCreditsSum                             *int64                `json:"ai_credits_agent_builder_ai_credits_sum,omitempty"`
		AiCreditsBitsAssistantAiCreditsSum                            *int64                `json:"ai_credits_bits_assistant_ai_credits_sum,omitempty"`
		AiCreditsBitsDevAiCreditsSum                                  *int64                `json:"ai_credits_bits_dev_ai_credits_sum,omitempty"`
		AiCreditsBitsSreAiCreditsSum                                  *int64                `json:"ai_credits_bits_sre_ai_credits_sum,omitempty"`
		AiCreditsSum                                                  *int64                `json:"ai_credits_sum,omitempty"`
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
		AuditTrailEventForwardingEventsSum                            *int64                `json:"audit_trail_event_forwarding_events_sum,omitempty"`
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
		CcmAnthropicSpendLast                                         *int64                `json:"ccm_anthropic_spend_last,omitempty"`
		CcmAwsSpendLast                                               *int64                `json:"ccm_aws_spend_last,omitempty"`
		CcmAzureSpendLast                                             *int64                `json:"ccm_azure_spend_last,omitempty"`
		CcmConfluentSpendLast                                         *int64                `json:"ccm_confluent_spend_last,omitempty"`
		CcmDatabricksSpendLast                                        *int64                `json:"ccm_databricks_spend_last,omitempty"`
		CcmElasticSpendLast                                           *int64                `json:"ccm_elastic_spend_last,omitempty"`
		CcmFastlySpendLast                                            *int64                `json:"ccm_fastly_spend_last,omitempty"`
		CcmGcpSpendLast                                               *int64                `json:"ccm_gcp_spend_last,omitempty"`
		CcmGithubSpendLast                                            *int64                `json:"ccm_github_spend_last,omitempty"`
		CcmMongodbSpendLast                                           *int64                `json:"ccm_mongodb_spend_last,omitempty"`
		CcmOciSpendLast                                               *int64                `json:"ccm_oci_spend_last,omitempty"`
		CcmOpenaiSpendLast                                            *int64                `json:"ccm_openai_spend_last,omitempty"`
		CcmSnowflakeSpendLast                                         *int64                `json:"ccm_snowflake_spend_last,omitempty"`
		CcmSpendMonitoredEntLast                                      *int64                `json:"ccm_spend_monitored_ent_last,omitempty"`
		CcmSpendMonitoredProLast                                      *int64                `json:"ccm_spend_monitored_pro_last,omitempty"`
		CcmTwilioSpendLast                                            *int64                `json:"ccm_twilio_spend_last,omitempty"`
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
		CloudSiemIndexedLogsSum                                       *int64                `json:"cloud_siem_indexed_logs_sum,omitempty"`
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
		CsmHostEnterpriseOciHostCountTop99p                           *int64                `json:"csm_host_enterprise_oci_host_count_top99p,omitempty"`
		CsmHostEnterpriseTotalHostCountTop99p                         *int64                `json:"csm_host_enterprise_total_host_count_top99p,omitempty"`
		CsmHostProHostsAgentlessScannersSum                           *int64                `json:"csm_host_pro_hosts_agentless_scanners_sum,omitempty"`
		CsmHostProHostsAgentlessScannersTop99p                        *int64                `json:"csm_host_pro_hosts_agentless_scanners_top99p,omitempty"`
		CsmHostProOciHostCountTop99p                                  *int64                `json:"csm_host_pro_oci_host_count_top99p,omitempty"`
		CspmAasHostTop99p                                             *int64                `json:"cspm_aas_host_top99p,omitempty"`
		CspmAwsHostTop99p                                             *int64                `json:"cspm_aws_host_top99p,omitempty"`
		CspmAzureHostTop99p                                           *int64                `json:"cspm_azure_host_top99p,omitempty"`
		CspmContainerAvg                                              *int64                `json:"cspm_container_avg,omitempty"`
		CspmContainerHwm                                              *int64                `json:"cspm_container_hwm,omitempty"`
		CspmGcpHostTop99p                                             *int64                `json:"cspm_gcp_host_top99p,omitempty"`
		CspmHostTop99p                                                *int64                `json:"cspm_host_top99p,omitempty"`
		CspmHostsAgentlessScannersSum                                 *int64                `json:"cspm_hosts_agentless_scanners_sum,omitempty"`
		CspmHostsAgentlessScannersTop99p                              *int64                `json:"cspm_hosts_agentless_scanners_top99p,omitempty"`
		CustomTsAvg                                                   *int64                `json:"custom_ts_avg,omitempty"`
		CwsContainerCountAvg                                          *int64                `json:"cws_container_count_avg,omitempty"`
		CwsFargateTaskAvg                                             *int64                `json:"cws_fargate_task_avg,omitempty"`
		CwsHostTop99p                                                 *int64                `json:"cws_host_top99p,omitempty"`
		DataJobsMonitoringHostHrSum                                   *int64                `json:"data_jobs_monitoring_host_hr_sum,omitempty"`
		DataStreamMonitoringHostCountSum                              *int64                `json:"data_stream_monitoring_host_count_sum,omitempty"`
		DataStreamMonitoringHostCountTop99p                           *int64                `json:"data_stream_monitoring_host_count_top99p,omitempty"`
		Date                                                          *time.Time            `json:"date,omitempty"`
		DbmHostTop99p                                                 *int64                `json:"dbm_host_top99p,omitempty"`
		DbmQueriesCountAvg                                            *int64                `json:"dbm_queries_count_avg,omitempty"`
		DoJobsMonitoringOrchestratorsJobHoursSum                      *int64                `json:"do_jobs_monitoring_orchestrators_job_hours_sum,omitempty"`
		EphInfraHostAgentSum                                          *int64                `json:"eph_infra_host_agent_sum,omitempty"`
		EphInfraHostAlibabaSum                                        *int64                `json:"eph_infra_host_alibaba_sum,omitempty"`
		EphInfraHostAwsSum                                            *int64                `json:"eph_infra_host_aws_sum,omitempty"`
		EphInfraHostAzureSum                                          *int64                `json:"eph_infra_host_azure_sum,omitempty"`
		EphInfraHostBasicInfraBasicAgentSum                           *int64                `json:"eph_infra_host_basic_infra_basic_agent_sum,omitempty"`
		EphInfraHostBasicInfraBasicVsphereSum                         *int64                `json:"eph_infra_host_basic_infra_basic_vsphere_sum,omitempty"`
		EphInfraHostBasicSum                                          *int64                `json:"eph_infra_host_basic_sum,omitempty"`
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
		FeatureFlagsConfigRequestsSum                                 *int64                `json:"feature_flags_config_requests_sum,omitempty"`
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
		IndexedPointsSum                                              *int64                `json:"indexed_points_sum,omitempty"`
		InfraCpuAvg                                                   *int64                `json:"infra_cpu_avg,omitempty"`
		InfraCpuDefaultInfraHostVcpuAgentAvg                          *int64                `json:"infra_cpu_default_infra_host_vcpu_agent_avg,omitempty"`
		InfraCpuDefaultInfraHostVcpuAgentBasicAvg                     *int64                `json:"infra_cpu_default_infra_host_vcpu_agent_basic_avg,omitempty"`
		InfraCpuDefaultInfraHostVcpuAgentBasicSum                     *int64                `json:"infra_cpu_default_infra_host_vcpu_agent_basic_sum,omitempty"`
		InfraCpuDefaultInfraHostVcpuAgentSum                          *int64                `json:"infra_cpu_default_infra_host_vcpu_agent_sum,omitempty"`
		InfraCpuDefaultInfraHostVcpuAwsAvg                            *int64                `json:"infra_cpu_default_infra_host_vcpu_aws_avg,omitempty"`
		InfraCpuDefaultInfraHostVcpuAwsSum                            *int64                `json:"infra_cpu_default_infra_host_vcpu_aws_sum,omitempty"`
		InfraCpuDefaultInfraHostVcpuAzureAvg                          *int64                `json:"infra_cpu_default_infra_host_vcpu_azure_avg,omitempty"`
		InfraCpuDefaultInfraHostVcpuAzureSum                          *int64                `json:"infra_cpu_default_infra_host_vcpu_azure_sum,omitempty"`
		InfraCpuDefaultInfraHostVcpuGcpAvg                            *int64                `json:"infra_cpu_default_infra_host_vcpu_gcp_avg,omitempty"`
		InfraCpuDefaultInfraHostVcpuGcpSum                            *int64                `json:"infra_cpu_default_infra_host_vcpu_gcp_sum,omitempty"`
		InfraCpuDefaultInfraHostVcpuNutanixAvg                        *int64                `json:"infra_cpu_default_infra_host_vcpu_nutanix_avg,omitempty"`
		InfraCpuDefaultInfraHostVcpuNutanixBasicAvg                   *int64                `json:"infra_cpu_default_infra_host_vcpu_nutanix_basic_avg,omitempty"`
		InfraCpuDefaultInfraHostVcpuNutanixBasicSum                   *int64                `json:"infra_cpu_default_infra_host_vcpu_nutanix_basic_sum,omitempty"`
		InfraCpuDefaultInfraHostVcpuNutanixSum                        *int64                `json:"infra_cpu_default_infra_host_vcpu_nutanix_sum,omitempty"`
		InfraCpuDefaultInfraHostVcpuOpentelemetryAvg                  *int64                `json:"infra_cpu_default_infra_host_vcpu_opentelemetry_avg,omitempty"`
		InfraCpuDefaultInfraHostVcpuOpentelemetrySum                  *int64                `json:"infra_cpu_default_infra_host_vcpu_opentelemetry_sum,omitempty"`
		InfraCpuObservedInfraHostVcpuAgentAvg                         *int64                `json:"infra_cpu_observed_infra_host_vcpu_agent_avg,omitempty"`
		InfraCpuObservedInfraHostVcpuAgentSum                         *int64                `json:"infra_cpu_observed_infra_host_vcpu_agent_sum,omitempty"`
		InfraCpuObservedInfraHostVcpuAwsAvg                           *int64                `json:"infra_cpu_observed_infra_host_vcpu_aws_avg,omitempty"`
		InfraCpuObservedInfraHostVcpuAwsSum                           *int64                `json:"infra_cpu_observed_infra_host_vcpu_aws_sum,omitempty"`
		InfraCpuObservedInfraHostVcpuAzureAvg                         *int64                `json:"infra_cpu_observed_infra_host_vcpu_azure_avg,omitempty"`
		InfraCpuObservedInfraHostVcpuAzureSum                         *int64                `json:"infra_cpu_observed_infra_host_vcpu_azure_sum,omitempty"`
		InfraCpuObservedInfraHostVcpuGcpAvg                           *int64                `json:"infra_cpu_observed_infra_host_vcpu_gcp_avg,omitempty"`
		InfraCpuObservedInfraHostVcpuGcpSum                           *int64                `json:"infra_cpu_observed_infra_host_vcpu_gcp_sum,omitempty"`
		InfraCpuObservedInfraHostVcpuNutanixAvg                       *int64                `json:"infra_cpu_observed_infra_host_vcpu_nutanix_avg,omitempty"`
		InfraCpuObservedInfraHostVcpuNutanixSum                       *int64                `json:"infra_cpu_observed_infra_host_vcpu_nutanix_sum,omitempty"`
		InfraCpuObservedInfraHostVcpuOpentelemetryAvg                 *int64                `json:"infra_cpu_observed_infra_host_vcpu_opentelemetry_avg,omitempty"`
		InfraCpuObservedInfraHostVcpuOpentelemetrySum                 *int64                `json:"infra_cpu_observed_infra_host_vcpu_opentelemetry_sum,omitempty"`
		InfraCpuSum                                                   *int64                `json:"infra_cpu_sum,omitempty"`
		InfraEdgeMonitoringDevicesTop99p                              *int64                `json:"infra_edge_monitoring_devices_top99p,omitempty"`
		InfraHostBasicInfraBasicAgentTop99p                           *int64                `json:"infra_host_basic_infra_basic_agent_top99p,omitempty"`
		InfraHostBasicInfraBasicVsphereTop99p                         *int64                `json:"infra_host_basic_infra_basic_vsphere_top99p,omitempty"`
		InfraHostBasicTop99p                                          *int64                `json:"infra_host_basic_top99p,omitempty"`
		InfraHostTop99p                                               *int64                `json:"infra_host_top99p,omitempty"`
		InfraStorageMgmtObjectsCountAvg                               *int64                `json:"infra_storage_mgmt_objects_count_avg,omitempty"`
		IngestPointsSum                                               *int64                `json:"ingest_points_sum,omitempty"`
		IngestedEventsBytesSum                                        *int64                `json:"ingested_events_bytes_sum,omitempty"`
		IotApmHostSum                                                 *int64                `json:"iot_apm_host_sum,omitempty"`
		IotApmHostTop99p                                              *int64                `json:"iot_apm_host_top99p,omitempty"`
		IotDeviceSum                                                  *int64                `json:"iot_device_sum,omitempty"`
		IotDeviceTop99p                                               *int64                `json:"iot_device_top99p,omitempty"`
		LlmObservability15dayRetentionSpansSum                        *int64                `json:"llm_observability_15day_retention_spans_sum,omitempty"`
		LlmObservability30dayRetentionSpansSum                        *int64                `json:"llm_observability_30day_retention_spans_sum,omitempty"`
		LlmObservability60dayRetentionSpansSum                        *int64                `json:"llm_observability_60day_retention_spans_sum,omitempty"`
		LlmObservability90dayRetentionSpansSum                        *int64                `json:"llm_observability_90day_retention_spans_sum,omitempty"`
		LlmObservabilityMinSpendSum                                   *int64                `json:"llm_observability_min_spend_sum,omitempty"`
		LlmObservabilitySum                                           *int64                `json:"llm_observability_sum,omitempty"`
		LogsArchiveSearchGbScannedSum                                 *int64                `json:"logs_archive_search_gb_scanned_sum,omitempty"`
		MetricNamesSum                                                *int64                `json:"metric_names_sum,omitempty"`
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
		NetworkPathSum                                                *int64                `json:"network_path_sum,omitempty"`
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
		ServerlessAppsApmApmGcpGkeAutopilotPodsAvg                    *int64                `json:"serverless_apps_apm_apm_gcp_gke_autopilot_pods_avg,omitempty"`
		ServerlessAppsApmAvg                                          *int64                `json:"serverless_apps_apm_avg,omitempty"`
		ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg    *int64                `json:"serverless_apps_apm_excl_fargate_apm_azure_appservice_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg *int64                `json:"serverless_apps_apm_excl_fargate_apm_azure_azurefunction_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg  *int64                `json:"serverless_apps_apm_excl_fargate_apm_azure_containerapp_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg   *int64                `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudfunction_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg        *int64                `json:"serverless_apps_apm_excl_fargate_apm_gcp_cloudrun_instances_avg,omitempty"`
		ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg         *int64                `json:"serverless_apps_apm_excl_fargate_apm_gcp_gke_autopilot_pods_avg,omitempty"`
		ServerlessAppsApmExclFargateAvg                               *int64                `json:"serverless_apps_apm_excl_fargate_avg,omitempty"`
		ServerlessAppsAzureContainerAppInstancesAvg                   *int64                `json:"serverless_apps_azure_container_app_instances_avg,omitempty"`
		ServerlessAppsAzureCountAvg                                   *int64                `json:"serverless_apps_azure_count_avg,omitempty"`
		ServerlessAppsAzureFunctionAppInstancesAvg                    *int64                `json:"serverless_apps_azure_function_app_instances_avg,omitempty"`
		ServerlessAppsAzureWebAppInstancesAvg                         *int64                `json:"serverless_apps_azure_web_app_instances_avg,omitempty"`
		ServerlessAppsDsmFargateTasksAvg                              *int64                `json:"serverless_apps_dsm_fargate_tasks_avg,omitempty"`
		ServerlessAppsEcsAvg                                          *int64                `json:"serverless_apps_ecs_avg,omitempty"`
		ServerlessAppsEksAvg                                          *int64                `json:"serverless_apps_eks_avg,omitempty"`
		ServerlessAppsExclFargateAvg                                  *int64                `json:"serverless_apps_excl_fargate_avg,omitempty"`
		ServerlessAppsExclFargateAzureContainerAppInstancesAvg        *int64                `json:"serverless_apps_excl_fargate_azure_container_app_instances_avg,omitempty"`
		ServerlessAppsExclFargateAzureFunctionAppInstancesAvg         *int64                `json:"serverless_apps_excl_fargate_azure_function_app_instances_avg,omitempty"`
		ServerlessAppsExclFargateAzureWebAppInstancesAvg              *int64                `json:"serverless_apps_excl_fargate_azure_web_app_instances_avg,omitempty"`
		ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg     *int64                `json:"serverless_apps_excl_fargate_google_cloud_functions_instances_avg,omitempty"`
		ServerlessAppsExclFargateGoogleCloudRunInstancesAvg           *int64                `json:"serverless_apps_excl_fargate_google_cloud_run_instances_avg,omitempty"`
		ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg          *int64                `json:"serverless_apps_excl_fargate_infra_gcp_gke_autopilot_pods_avg,omitempty"`
		ServerlessAppsGoogleCloudFunctionsInstancesAvg                *int64                `json:"serverless_apps_google_cloud_functions_instances_avg,omitempty"`
		ServerlessAppsGoogleCloudRunInstancesAvg                      *int64                `json:"serverless_apps_google_cloud_run_instances_avg,omitempty"`
		ServerlessAppsGoogleCountAvg                                  *int64                `json:"serverless_apps_google_count_avg,omitempty"`
		ServerlessAppsInfraGcpGkeAutopilotPodsAvg                     *int64                `json:"serverless_apps_infra_gcp_gke_autopilot_pods_avg,omitempty"`
		ServerlessAppsTotalCountAvg                                   *int64                `json:"serverless_apps_total_count_avg,omitempty"`
		Siem12moRetentionSum                                          *int64                `json:"siem_12mo_retention_sum,omitempty"`
		Siem6moRetentionSum                                           *int64                `json:"siem_6mo_retention_sum,omitempty"`
		SiemAnalyzedLogsAddOnCountSum                                 *int64                `json:"siem_analyzed_logs_add_on_count_sum,omitempty"`
		SnmpDeviceCountSum                                            *int64                `json:"snmp_device_count_sum,omitempty"`
		SnmpDeviceCountTop99p                                         *int64                `json:"snmp_device_count_top99p,omitempty"`
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
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
	} else {
		return err
	}
	o.AgentHostTop99p = all.AgentHostTop99p
	o.AiCreditsAgentBuilderAiCreditsSum = all.AiCreditsAgentBuilderAiCreditsSum
	o.AiCreditsBitsAssistantAiCreditsSum = all.AiCreditsBitsAssistantAiCreditsSum
	o.AiCreditsBitsDevAiCreditsSum = all.AiCreditsBitsDevAiCreditsSum
	o.AiCreditsBitsSreAiCreditsSum = all.AiCreditsBitsSreAiCreditsSum
	o.AiCreditsSum = all.AiCreditsSum
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
	o.AuditTrailEventForwardingEventsSum = all.AuditTrailEventForwardingEventsSum
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
	o.CcmAnthropicSpendLast = all.CcmAnthropicSpendLast
	o.CcmAwsSpendLast = all.CcmAwsSpendLast
	o.CcmAzureSpendLast = all.CcmAzureSpendLast
	o.CcmConfluentSpendLast = all.CcmConfluentSpendLast
	o.CcmDatabricksSpendLast = all.CcmDatabricksSpendLast
	o.CcmElasticSpendLast = all.CcmElasticSpendLast
	o.CcmFastlySpendLast = all.CcmFastlySpendLast
	o.CcmGcpSpendLast = all.CcmGcpSpendLast
	o.CcmGithubSpendLast = all.CcmGithubSpendLast
	o.CcmMongodbSpendLast = all.CcmMongodbSpendLast
	o.CcmOciSpendLast = all.CcmOciSpendLast
	o.CcmOpenaiSpendLast = all.CcmOpenaiSpendLast
	o.CcmSnowflakeSpendLast = all.CcmSnowflakeSpendLast
	o.CcmSpendMonitoredEntLast = all.CcmSpendMonitoredEntLast
	o.CcmSpendMonitoredProLast = all.CcmSpendMonitoredProLast
	o.CcmTwilioSpendLast = all.CcmTwilioSpendLast
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
	o.CloudSiemIndexedLogsSum = all.CloudSiemIndexedLogsSum
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
	o.CsmHostEnterpriseOciHostCountTop99p = all.CsmHostEnterpriseOciHostCountTop99p
	o.CsmHostEnterpriseTotalHostCountTop99p = all.CsmHostEnterpriseTotalHostCountTop99p
	o.CsmHostProHostsAgentlessScannersSum = all.CsmHostProHostsAgentlessScannersSum
	o.CsmHostProHostsAgentlessScannersTop99p = all.CsmHostProHostsAgentlessScannersTop99p
	o.CsmHostProOciHostCountTop99p = all.CsmHostProOciHostCountTop99p
	o.CspmAasHostTop99p = all.CspmAasHostTop99p
	o.CspmAwsHostTop99p = all.CspmAwsHostTop99p
	o.CspmAzureHostTop99p = all.CspmAzureHostTop99p
	o.CspmContainerAvg = all.CspmContainerAvg
	o.CspmContainerHwm = all.CspmContainerHwm
	o.CspmGcpHostTop99p = all.CspmGcpHostTop99p
	o.CspmHostTop99p = all.CspmHostTop99p
	o.CspmHostsAgentlessScannersSum = all.CspmHostsAgentlessScannersSum
	o.CspmHostsAgentlessScannersTop99p = all.CspmHostsAgentlessScannersTop99p
	o.CustomTsAvg = all.CustomTsAvg
	o.CwsContainerCountAvg = all.CwsContainerCountAvg
	o.CwsFargateTaskAvg = all.CwsFargateTaskAvg
	o.CwsHostTop99p = all.CwsHostTop99p
	o.DataJobsMonitoringHostHrSum = all.DataJobsMonitoringHostHrSum
	o.DataStreamMonitoringHostCountSum = all.DataStreamMonitoringHostCountSum
	o.DataStreamMonitoringHostCountTop99p = all.DataStreamMonitoringHostCountTop99p
	o.Date = all.Date
	o.DbmHostTop99p = all.DbmHostTop99p
	o.DbmQueriesCountAvg = all.DbmQueriesCountAvg
	o.DoJobsMonitoringOrchestratorsJobHoursSum = all.DoJobsMonitoringOrchestratorsJobHoursSum
	o.EphInfraHostAgentSum = all.EphInfraHostAgentSum
	o.EphInfraHostAlibabaSum = all.EphInfraHostAlibabaSum
	o.EphInfraHostAwsSum = all.EphInfraHostAwsSum
	o.EphInfraHostAzureSum = all.EphInfraHostAzureSum
	o.EphInfraHostBasicInfraBasicAgentSum = all.EphInfraHostBasicInfraBasicAgentSum
	o.EphInfraHostBasicInfraBasicVsphereSum = all.EphInfraHostBasicInfraBasicVsphereSum
	o.EphInfraHostBasicSum = all.EphInfraHostBasicSum
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
	o.FeatureFlagsConfigRequestsSum = all.FeatureFlagsConfigRequestsSum
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
	o.IndexedPointsSum = all.IndexedPointsSum
	o.InfraCpuAvg = all.InfraCpuAvg
	o.InfraCpuDefaultInfraHostVcpuAgentAvg = all.InfraCpuDefaultInfraHostVcpuAgentAvg
	o.InfraCpuDefaultInfraHostVcpuAgentBasicAvg = all.InfraCpuDefaultInfraHostVcpuAgentBasicAvg
	o.InfraCpuDefaultInfraHostVcpuAgentBasicSum = all.InfraCpuDefaultInfraHostVcpuAgentBasicSum
	o.InfraCpuDefaultInfraHostVcpuAgentSum = all.InfraCpuDefaultInfraHostVcpuAgentSum
	o.InfraCpuDefaultInfraHostVcpuAwsAvg = all.InfraCpuDefaultInfraHostVcpuAwsAvg
	o.InfraCpuDefaultInfraHostVcpuAwsSum = all.InfraCpuDefaultInfraHostVcpuAwsSum
	o.InfraCpuDefaultInfraHostVcpuAzureAvg = all.InfraCpuDefaultInfraHostVcpuAzureAvg
	o.InfraCpuDefaultInfraHostVcpuAzureSum = all.InfraCpuDefaultInfraHostVcpuAzureSum
	o.InfraCpuDefaultInfraHostVcpuGcpAvg = all.InfraCpuDefaultInfraHostVcpuGcpAvg
	o.InfraCpuDefaultInfraHostVcpuGcpSum = all.InfraCpuDefaultInfraHostVcpuGcpSum
	o.InfraCpuDefaultInfraHostVcpuNutanixAvg = all.InfraCpuDefaultInfraHostVcpuNutanixAvg
	o.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg = all.InfraCpuDefaultInfraHostVcpuNutanixBasicAvg
	o.InfraCpuDefaultInfraHostVcpuNutanixBasicSum = all.InfraCpuDefaultInfraHostVcpuNutanixBasicSum
	o.InfraCpuDefaultInfraHostVcpuNutanixSum = all.InfraCpuDefaultInfraHostVcpuNutanixSum
	o.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg = all.InfraCpuDefaultInfraHostVcpuOpentelemetryAvg
	o.InfraCpuDefaultInfraHostVcpuOpentelemetrySum = all.InfraCpuDefaultInfraHostVcpuOpentelemetrySum
	o.InfraCpuObservedInfraHostVcpuAgentAvg = all.InfraCpuObservedInfraHostVcpuAgentAvg
	o.InfraCpuObservedInfraHostVcpuAgentSum = all.InfraCpuObservedInfraHostVcpuAgentSum
	o.InfraCpuObservedInfraHostVcpuAwsAvg = all.InfraCpuObservedInfraHostVcpuAwsAvg
	o.InfraCpuObservedInfraHostVcpuAwsSum = all.InfraCpuObservedInfraHostVcpuAwsSum
	o.InfraCpuObservedInfraHostVcpuAzureAvg = all.InfraCpuObservedInfraHostVcpuAzureAvg
	o.InfraCpuObservedInfraHostVcpuAzureSum = all.InfraCpuObservedInfraHostVcpuAzureSum
	o.InfraCpuObservedInfraHostVcpuGcpAvg = all.InfraCpuObservedInfraHostVcpuGcpAvg
	o.InfraCpuObservedInfraHostVcpuGcpSum = all.InfraCpuObservedInfraHostVcpuGcpSum
	o.InfraCpuObservedInfraHostVcpuNutanixAvg = all.InfraCpuObservedInfraHostVcpuNutanixAvg
	o.InfraCpuObservedInfraHostVcpuNutanixSum = all.InfraCpuObservedInfraHostVcpuNutanixSum
	o.InfraCpuObservedInfraHostVcpuOpentelemetryAvg = all.InfraCpuObservedInfraHostVcpuOpentelemetryAvg
	o.InfraCpuObservedInfraHostVcpuOpentelemetrySum = all.InfraCpuObservedInfraHostVcpuOpentelemetrySum
	o.InfraCpuSum = all.InfraCpuSum
	o.InfraEdgeMonitoringDevicesTop99p = all.InfraEdgeMonitoringDevicesTop99p
	o.InfraHostBasicInfraBasicAgentTop99p = all.InfraHostBasicInfraBasicAgentTop99p
	o.InfraHostBasicInfraBasicVsphereTop99p = all.InfraHostBasicInfraBasicVsphereTop99p
	o.InfraHostBasicTop99p = all.InfraHostBasicTop99p
	o.InfraHostTop99p = all.InfraHostTop99p
	o.InfraStorageMgmtObjectsCountAvg = all.InfraStorageMgmtObjectsCountAvg
	o.IngestPointsSum = all.IngestPointsSum
	o.IngestedEventsBytesSum = all.IngestedEventsBytesSum
	o.IotApmHostSum = all.IotApmHostSum
	o.IotApmHostTop99p = all.IotApmHostTop99p
	o.IotDeviceSum = all.IotDeviceSum
	o.IotDeviceTop99p = all.IotDeviceTop99p
	o.LlmObservability15dayRetentionSpansSum = all.LlmObservability15dayRetentionSpansSum
	o.LlmObservability30dayRetentionSpansSum = all.LlmObservability30dayRetentionSpansSum
	o.LlmObservability60dayRetentionSpansSum = all.LlmObservability60dayRetentionSpansSum
	o.LlmObservability90dayRetentionSpansSum = all.LlmObservability90dayRetentionSpansSum
	o.LlmObservabilityMinSpendSum = all.LlmObservabilityMinSpendSum
	o.LlmObservabilitySum = all.LlmObservabilitySum
	o.LogsArchiveSearchGbScannedSum = all.LogsArchiveSearchGbScannedSum
	o.MetricNamesSum = all.MetricNamesSum
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
	o.NetworkPathSum = all.NetworkPathSum
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
	o.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg = all.ServerlessAppsApmApmGcpGkeAutopilotPodsAvg
	o.ServerlessAppsApmAvg = all.ServerlessAppsApmAvg
	o.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg = all.ServerlessAppsApmExclFargateApmAzureAppserviceInstancesAvg
	o.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg = all.ServerlessAppsApmExclFargateApmAzureAzurefunctionInstancesAvg
	o.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg = all.ServerlessAppsApmExclFargateApmAzureContainerappInstancesAvg
	o.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg = all.ServerlessAppsApmExclFargateApmGcpCloudfunctionInstancesAvg
	o.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg = all.ServerlessAppsApmExclFargateApmGcpCloudrunInstancesAvg
	o.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg = all.ServerlessAppsApmExclFargateApmGcpGkeAutopilotPodsAvg
	o.ServerlessAppsApmExclFargateAvg = all.ServerlessAppsApmExclFargateAvg
	o.ServerlessAppsAzureContainerAppInstancesAvg = all.ServerlessAppsAzureContainerAppInstancesAvg
	o.ServerlessAppsAzureCountAvg = all.ServerlessAppsAzureCountAvg
	o.ServerlessAppsAzureFunctionAppInstancesAvg = all.ServerlessAppsAzureFunctionAppInstancesAvg
	o.ServerlessAppsAzureWebAppInstancesAvg = all.ServerlessAppsAzureWebAppInstancesAvg
	o.ServerlessAppsDsmFargateTasksAvg = all.ServerlessAppsDsmFargateTasksAvg
	o.ServerlessAppsEcsAvg = all.ServerlessAppsEcsAvg
	o.ServerlessAppsEksAvg = all.ServerlessAppsEksAvg
	o.ServerlessAppsExclFargateAvg = all.ServerlessAppsExclFargateAvg
	o.ServerlessAppsExclFargateAzureContainerAppInstancesAvg = all.ServerlessAppsExclFargateAzureContainerAppInstancesAvg
	o.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg = all.ServerlessAppsExclFargateAzureFunctionAppInstancesAvg
	o.ServerlessAppsExclFargateAzureWebAppInstancesAvg = all.ServerlessAppsExclFargateAzureWebAppInstancesAvg
	o.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg = all.ServerlessAppsExclFargateGoogleCloudFunctionsInstancesAvg
	o.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg = all.ServerlessAppsExclFargateGoogleCloudRunInstancesAvg
	o.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg = all.ServerlessAppsExclFargateInfraGcpGkeAutopilotPodsAvg
	o.ServerlessAppsGoogleCloudFunctionsInstancesAvg = all.ServerlessAppsGoogleCloudFunctionsInstancesAvg
	o.ServerlessAppsGoogleCloudRunInstancesAvg = all.ServerlessAppsGoogleCloudRunInstancesAvg
	o.ServerlessAppsGoogleCountAvg = all.ServerlessAppsGoogleCountAvg
	o.ServerlessAppsInfraGcpGkeAutopilotPodsAvg = all.ServerlessAppsInfraGcpGkeAutopilotPodsAvg
	o.ServerlessAppsTotalCountAvg = all.ServerlessAppsTotalCountAvg
	o.Siem12moRetentionSum = all.Siem12moRetentionSum
	o.Siem6moRetentionSum = all.Siem6moRetentionSum
	o.SiemAnalyzedLogsAddOnCountSum = all.SiemAnalyzedLogsAddOnCountSum
	o.SnmpDeviceCountSum = all.SnmpDeviceCountSum
	o.SnmpDeviceCountTop99p = all.SnmpDeviceCountTop99p
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
