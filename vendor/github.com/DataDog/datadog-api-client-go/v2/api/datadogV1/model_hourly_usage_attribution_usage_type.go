// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HourlyUsageAttributionUsageType Supported products for hourly usage attribution requests.
// The following values have been **deprecated**: `estimated_indexed_spans_usage`, `estimated_ingested_spans_usage`.
type HourlyUsageAttributionUsageType string

// List of HourlyUsageAttributionUsageType.
const (
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_API_USAGE                                      HourlyUsageAttributionUsageType = "api_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_FARGATE_USAGE                              HourlyUsageAttributionUsageType = "apm_fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_HOST_USAGE                                 HourlyUsageAttributionUsageType = "apm_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_USM_USAGE                                  HourlyUsageAttributionUsageType = "apm_usm_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APPSEC_FARGATE_USAGE                           HourlyUsageAttributionUsageType = "appsec_fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APPSEC_USAGE                                   HourlyUsageAttributionUsageType = "appsec_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ASM_SERVERLESS_TRACED_INVOCATIONS_USAGE        HourlyUsageAttributionUsageType = "asm_serverless_traced_invocations_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ASM_SERVERLESS_TRACED_INVOCATIONS_PERCENTAGE   HourlyUsageAttributionUsageType = "asm_serverless_traced_invocations_percentage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_BROWSER_USAGE                                  HourlyUsageAttributionUsageType = "browser_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CI_PIPELINE_INDEXED_SPANS_USAGE                HourlyUsageAttributionUsageType = "ci_pipeline_indexed_spans_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CI_TEST_INDEXED_SPANS_USAGE                    HourlyUsageAttributionUsageType = "ci_test_indexed_spans_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CI_VISIBILITY_ITR_USAGE                        HourlyUsageAttributionUsageType = "ci_visibility_itr_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CLOUD_SIEM_USAGE                               HourlyUsageAttributionUsageType = "cloud_siem_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CODE_SECURITY_HOST_USAGE                       HourlyUsageAttributionUsageType = "code_security_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_EXCL_AGENT_USAGE                     HourlyUsageAttributionUsageType = "container_excl_agent_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_USAGE                                HourlyUsageAttributionUsageType = "container_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CSPM_CONTAINERS_USAGE                          HourlyUsageAttributionUsageType = "cspm_containers_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CSPM_HOSTS_USAGE                               HourlyUsageAttributionUsageType = "cspm_hosts_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_EVENT_USAGE                             HourlyUsageAttributionUsageType = "custom_event_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_INGESTED_TIMESERIES_USAGE               HourlyUsageAttributionUsageType = "custom_ingested_timeseries_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_TIMESERIES_USAGE                        HourlyUsageAttributionUsageType = "custom_timeseries_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_CONTAINERS_USAGE                           HourlyUsageAttributionUsageType = "cws_containers_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_FARGATE_TASK_USAGE                         HourlyUsageAttributionUsageType = "cws_fargate_task_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_HOSTS_USAGE                                HourlyUsageAttributionUsageType = "cws_hosts_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DATA_JOBS_MONITORING_USAGE                     HourlyUsageAttributionUsageType = "data_jobs_monitoring_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DATA_STREAM_MONITORING_USAGE                   HourlyUsageAttributionUsageType = "data_stream_monitoring_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DBM_HOSTS_USAGE                                HourlyUsageAttributionUsageType = "dbm_hosts_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DBM_QUERIES_USAGE                              HourlyUsageAttributionUsageType = "dbm_queries_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ERROR_TRACKING_USAGE                           HourlyUsageAttributionUsageType = "error_tracking_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ERROR_TRACKING_PERCENTAGE                      HourlyUsageAttributionUsageType = "error_tracking_percentage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INDEXED_SPANS_USAGE                  HourlyUsageAttributionUsageType = "estimated_indexed_spans_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INGESTED_SPANS_USAGE                 HourlyUsageAttributionUsageType = "estimated_ingested_spans_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FARGATE_USAGE                                  HourlyUsageAttributionUsageType = "fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FUNCTIONS_USAGE                                HourlyUsageAttributionUsageType = "functions_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INCIDENT_MANAGEMENT_MONTHLY_ACTIVE_USERS_USAGE HourlyUsageAttributionUsageType = "incident_management_monthly_active_users_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INDEXED_SPANS_USAGE                            HourlyUsageAttributionUsageType = "indexed_spans_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INFRA_HOST_USAGE                               HourlyUsageAttributionUsageType = "infra_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INGESTED_LOGS_BYTES_USAGE                      HourlyUsageAttributionUsageType = "ingested_logs_bytes_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INGESTED_SPANS_BYTES_USAGE                     HourlyUsageAttributionUsageType = "ingested_spans_bytes_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INVOCATIONS_USAGE                              HourlyUsageAttributionUsageType = "invocations_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LAMBDA_TRACED_INVOCATIONS_USAGE                HourlyUsageAttributionUsageType = "lambda_traced_invocations_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LLM_OBSERVABILITY_USAGE                        HourlyUsageAttributionUsageType = "llm_observability_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_15DAY_USAGE                       HourlyUsageAttributionUsageType = "logs_indexed_15day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_180DAY_USAGE                      HourlyUsageAttributionUsageType = "logs_indexed_180day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_1DAY_USAGE                        HourlyUsageAttributionUsageType = "logs_indexed_1day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_30DAY_USAGE                       HourlyUsageAttributionUsageType = "logs_indexed_30day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_360DAY_USAGE                      HourlyUsageAttributionUsageType = "logs_indexed_360day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_3DAY_USAGE                        HourlyUsageAttributionUsageType = "logs_indexed_3day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_45DAY_USAGE                       HourlyUsageAttributionUsageType = "logs_indexed_45day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_60DAY_USAGE                       HourlyUsageAttributionUsageType = "logs_indexed_60day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_7DAY_USAGE                        HourlyUsageAttributionUsageType = "logs_indexed_7day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_90DAY_USAGE                       HourlyUsageAttributionUsageType = "logs_indexed_90day_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_CUSTOM_RETENTION_USAGE            HourlyUsageAttributionUsageType = "logs_indexed_custom_retention_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_MOBILE_APP_TESTING_USAGE                       HourlyUsageAttributionUsageType = "mobile_app_testing_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NDM_NETFLOW_USAGE                              HourlyUsageAttributionUsageType = "ndm_netflow_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NETWORK_DEVICE_WIRELESS_USAGE                  HourlyUsageAttributionUsageType = "npm_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NPM_HOST_USAGE                                 HourlyUsageAttributionUsageType = "network_device_wireless_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_OBS_PIPELINE_BYTES_USAGE                       HourlyUsageAttributionUsageType = "obs_pipeline_bytes_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_OBS_PIPELINE_VCPU_USAGE                        HourlyUsageAttributionUsageType = "obs_pipelines_vcpu_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ONLINE_ARCHIVE_USAGE                           HourlyUsageAttributionUsageType = "online_archive_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PRODUCT_ANALYTICS_SESSION_USAGE                HourlyUsageAttributionUsageType = "product_analytics_session_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_CONTAINER_USAGE                       HourlyUsageAttributionUsageType = "profiled_container_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_FARGATE_USAGE                         HourlyUsageAttributionUsageType = "profiled_fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_HOST_USAGE                            HourlyUsageAttributionUsageType = "profiled_host_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PUBLISHED_APP_USAGE                            HourlyUsageAttributionUsageType = "published_app"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_BROWSER_MOBILE_SESSIONS_USAGE              HourlyUsageAttributionUsageType = "rum_browser_mobile_sessions_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_INGESTED_USAGE                             HourlyUsageAttributionUsageType = "rum_ingested_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_INVESTIGATE_USAGE                          HourlyUsageAttributionUsageType = "rum_investigate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_REPLAY_SESSIONS_USAGE                      HourlyUsageAttributionUsageType = "rum_replay_sessions_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_SESSION_REPLAY_ADD_ON_USAGE                HourlyUsageAttributionUsageType = "rum_session_replay_add_on_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SCA_FARGATE_USAGE                              HourlyUsageAttributionUsageType = "sca_fargate_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SDS_SCANNED_BYTES_USAGE                        HourlyUsageAttributionUsageType = "sds_scanned_bytes_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SERVERLESS_APPS_USAGE                          HourlyUsageAttributionUsageType = "serverless_apps_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SIEM_ANALYZED_LOGS_ADD_ON_USAGE                HourlyUsageAttributionUsageType = "siem_analyzed_logs_add_on_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SIEM_INGESTED_BYTES_USAGE                      HourlyUsageAttributionUsageType = "siem_ingested_bytes_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SNMP_USAGE                                     HourlyUsageAttributionUsageType = "snmp_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_UNIVERSAL_SERVICE_MONITORING_USAGE             HourlyUsageAttributionUsageType = "universal_service_monitoring_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_VULN_MANAGEMENT_HOSTS_USAGE                    HourlyUsageAttributionUsageType = "vuln_management_hosts_usage"
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_WORKFLOW_EXECUTIONS_USAGE                      HourlyUsageAttributionUsageType = "workflow_executions_usage"
)

var allowedHourlyUsageAttributionUsageTypeEnumValues = []HourlyUsageAttributionUsageType{
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_API_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APM_USM_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APPSEC_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_APPSEC_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ASM_SERVERLESS_TRACED_INVOCATIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ASM_SERVERLESS_TRACED_INVOCATIONS_PERCENTAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_BROWSER_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CI_PIPELINE_INDEXED_SPANS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CI_TEST_INDEXED_SPANS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CI_VISIBILITY_ITR_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CLOUD_SIEM_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CODE_SECURITY_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_EXCL_AGENT_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CONTAINER_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CSPM_CONTAINERS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CSPM_HOSTS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_EVENT_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_INGESTED_TIMESERIES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CUSTOM_TIMESERIES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_CONTAINERS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_FARGATE_TASK_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_CWS_HOSTS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DATA_JOBS_MONITORING_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DATA_STREAM_MONITORING_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DBM_HOSTS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_DBM_QUERIES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ERROR_TRACKING_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ERROR_TRACKING_PERCENTAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INDEXED_SPANS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ESTIMATED_INGESTED_SPANS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_FUNCTIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INCIDENT_MANAGEMENT_MONTHLY_ACTIVE_USERS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INDEXED_SPANS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INFRA_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INGESTED_LOGS_BYTES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INGESTED_SPANS_BYTES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_INVOCATIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LAMBDA_TRACED_INVOCATIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LLM_OBSERVABILITY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_15DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_180DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_1DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_30DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_360DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_3DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_45DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_60DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_7DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_90DAY_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_LOGS_INDEXED_CUSTOM_RETENTION_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_MOBILE_APP_TESTING_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NDM_NETFLOW_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NETWORK_DEVICE_WIRELESS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_NPM_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_OBS_PIPELINE_BYTES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_OBS_PIPELINE_VCPU_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_ONLINE_ARCHIVE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PRODUCT_ANALYTICS_SESSION_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_CONTAINER_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PROFILED_HOST_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_PUBLISHED_APP_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_BROWSER_MOBILE_SESSIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_INGESTED_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_INVESTIGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_REPLAY_SESSIONS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_RUM_SESSION_REPLAY_ADD_ON_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SCA_FARGATE_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SDS_SCANNED_BYTES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SERVERLESS_APPS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SIEM_ANALYZED_LOGS_ADD_ON_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SIEM_INGESTED_BYTES_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_SNMP_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_UNIVERSAL_SERVICE_MONITORING_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_VULN_MANAGEMENT_HOSTS_USAGE,
	HOURLYUSAGEATTRIBUTIONUSAGETYPE_WORKFLOW_EXECUTIONS_USAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *HourlyUsageAttributionUsageType) GetAllowedValues() []HourlyUsageAttributionUsageType {
	return allowedHourlyUsageAttributionUsageTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HourlyUsageAttributionUsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HourlyUsageAttributionUsageType(value)
	return nil
}

// NewHourlyUsageAttributionUsageTypeFromValue returns a pointer to a valid HourlyUsageAttributionUsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHourlyUsageAttributionUsageTypeFromValue(v string) (*HourlyUsageAttributionUsageType, error) {
	ev := HourlyUsageAttributionUsageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HourlyUsageAttributionUsageType: valid values are %v", v, allowedHourlyUsageAttributionUsageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HourlyUsageAttributionUsageType) IsValid() bool {
	for _, existing := range allowedHourlyUsageAttributionUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HourlyUsageAttributionUsageType value.
func (v HourlyUsageAttributionUsageType) Ptr() *HourlyUsageAttributionUsageType {
	return &v
}
