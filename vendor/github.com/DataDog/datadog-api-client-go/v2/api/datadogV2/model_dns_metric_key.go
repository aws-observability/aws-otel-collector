// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DnsMetricKey The metric key for DNS metrics.
type DnsMetricKey string

// List of DnsMetricKey.
const (
	DNSMETRICKEY_DNS_TOTAL_REQUESTS             DnsMetricKey = "dns_total_requests"
	DNSMETRICKEY_DNS_FAILURES                   DnsMetricKey = "dns_failures"
	DNSMETRICKEY_DNS_SUCCESSFUL_RESPONSES       DnsMetricKey = "dns_successful_responses"
	DNSMETRICKEY_DNS_FAILED_RESPONSES           DnsMetricKey = "dns_failed_responses"
	DNSMETRICKEY_DNS_TIMEOUTS                   DnsMetricKey = "dns_timeouts"
	DNSMETRICKEY_DNS_RESPONSES_NXDOMAIN         DnsMetricKey = "dns_responses.nxdomain"
	DNSMETRICKEY_DNS_RESPONSES_SERVFAIL         DnsMetricKey = "dns_responses.servfail"
	DNSMETRICKEY_DNS_RESPONSES_OTHER            DnsMetricKey = "dns_responses.other"
	DNSMETRICKEY_DNS_SUCCESS_LATENCY_PERCENTILE DnsMetricKey = "dns_success_latency_percentile"
	DNSMETRICKEY_DNS_FAILURE_LATENCY_PERCENTILE DnsMetricKey = "dns_failure_latency_percentile"
)

var allowedDnsMetricKeyEnumValues = []DnsMetricKey{
	DNSMETRICKEY_DNS_TOTAL_REQUESTS,
	DNSMETRICKEY_DNS_FAILURES,
	DNSMETRICKEY_DNS_SUCCESSFUL_RESPONSES,
	DNSMETRICKEY_DNS_FAILED_RESPONSES,
	DNSMETRICKEY_DNS_TIMEOUTS,
	DNSMETRICKEY_DNS_RESPONSES_NXDOMAIN,
	DNSMETRICKEY_DNS_RESPONSES_SERVFAIL,
	DNSMETRICKEY_DNS_RESPONSES_OTHER,
	DNSMETRICKEY_DNS_SUCCESS_LATENCY_PERCENTILE,
	DNSMETRICKEY_DNS_FAILURE_LATENCY_PERCENTILE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DnsMetricKey) GetAllowedValues() []DnsMetricKey {
	return allowedDnsMetricKeyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DnsMetricKey) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DnsMetricKey(value)
	return nil
}

// NewDnsMetricKeyFromValue returns a pointer to a valid DnsMetricKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDnsMetricKeyFromValue(v string) (*DnsMetricKey, error) {
	ev := DnsMetricKey(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DnsMetricKey: valid values are %v", v, allowedDnsMetricKeyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DnsMetricKey) IsValid() bool {
	for _, existing := range allowedDnsMetricKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DnsMetricKey value.
func (v DnsMetricKey) Ptr() *DnsMetricKey {
	return &v
}
