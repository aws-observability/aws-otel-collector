// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NetworkHealthInsightAttributes Detailed attributes of a network health insight.
type NetworkHealthInsightAttributes struct {
	// AWS account identifier where the certificate is located. Only set for `tls-cert` insights.
	AccountId *string `json:"account_id,omitempty"`
	// ARN or identifier of the certificate. Only set for `tls-cert` insights.
	CertificateId *string `json:"certificate_id,omitempty"`
	// Percentage of the certificate's validity period that has elapsed, ranging from 0 to 100.
	// Only set for `tls-cert` insights.
	CertificateLifetimePercent *float64 `json:"certificate_lifetime_percent,omitempty"`
	// AWS region where the client is located. Only set for `tls-cert` insights.
	ClientRegion *string `json:"client_region,omitempty"`
	// Name of the service making the request (DNS query or TLS-secured connection).
	// Set to `N/A` when the client service cannot be determined.
	ClientService *string `json:"client_service,omitempty"`
	// Number of days remaining until the certificate expires. Negative values indicate the
	// certificate has already expired. Only set for `tls-cert` insights.
	DaysUntilExpiration *int64 `json:"days_until_expiration,omitempty"`
	// Domain name that was being resolved when the DNS failure occurred. Only set for `dns` insights.
	DnsQuery *string `json:"dns_query,omitempty"`
	// DNS server that received the failing query. Only set for `dns` insights.
	DnsServer *string `json:"dns_server,omitempty"`
	// Domain name covered by the certificate. Only set for `tls-cert` insights.
	DomainName *string `json:"domain_name,omitempty"`
	// Count of failed events observed during the query window. Only set for `dns`, `tcp`,
	// and `security-group` insights.
	FailureMagnitude *int64 `json:"failure_magnitude,omitempty"`
	// Percentage of requests that failed during the query window, ranging from 0 to 100.
	// Only set for `dns`, `tcp`, and `security-group` insights.
	FailureRate *float64 `json:"failure_rate,omitempty"`
	// Specific failure type within the insight category. For DNS insights: `timeout`, `nxdomain`,
	// `servfail`, or `general_failure`. For TLS certificate insights: `expired` or `expiring_soon`.
	// For security group insights: `denied`.
	FailureType *NetworkHealthInsightFailureType `json:"failure_type,omitempty"`
	// ARN of the load balancer using the certificate. Only set for `tls-cert` insights.
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	// AWS region where the server or load balancer is located. Only set for `tls-cert` insights.
	ServerRegion *string `json:"server_region,omitempty"`
	// Name of the target service the client was trying to reach.
	ServerService *string `json:"server_service,omitempty"`
	// Total number of requests observed during the query window. Provides context for
	// `failure_magnitude` and `failure_rate`. Only set for `dns`, `tcp`, and `security-group` insights.
	TotalRequests *int64 `json:"total_requests,omitempty"`
	// Network traffic volume metrics between the client and server services during the query window.
	TrafficVolume *NetworkHealthInsightTrafficVolume `json:"traffic_volume,omitempty"`
	// Category of network health insight. Indicates whether the insight relates to a DNS issue (`dns`),
	// a TCP issue (`tcp`), a TLS certificate issue (`tls-cert`), or a security group denial (`security-group`).
	Type *NetworkHealthInsightCategory `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNetworkHealthInsightAttributes instantiates a new NetworkHealthInsightAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNetworkHealthInsightAttributes() *NetworkHealthInsightAttributes {
	this := NetworkHealthInsightAttributes{}
	return &this
}

// NewNetworkHealthInsightAttributesWithDefaults instantiates a new NetworkHealthInsightAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNetworkHealthInsightAttributesWithDefaults() *NetworkHealthInsightAttributes {
	this := NetworkHealthInsightAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *NetworkHealthInsightAttributes) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetCertificateId() string {
	if o == nil || o.CertificateId == nil {
		var ret string
		return ret
	}
	return *o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetCertificateIdOk() (*string, bool) {
	if o == nil || o.CertificateId == nil {
		return nil, false
	}
	return o.CertificateId, true
}

// HasCertificateId returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasCertificateId() bool {
	return o != nil && o.CertificateId != nil
}

// SetCertificateId gets a reference to the given string and assigns it to the CertificateId field.
func (o *NetworkHealthInsightAttributes) SetCertificateId(v string) {
	o.CertificateId = &v
}

// GetCertificateLifetimePercent returns the CertificateLifetimePercent field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetCertificateLifetimePercent() float64 {
	if o == nil || o.CertificateLifetimePercent == nil {
		var ret float64
		return ret
	}
	return *o.CertificateLifetimePercent
}

// GetCertificateLifetimePercentOk returns a tuple with the CertificateLifetimePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetCertificateLifetimePercentOk() (*float64, bool) {
	if o == nil || o.CertificateLifetimePercent == nil {
		return nil, false
	}
	return o.CertificateLifetimePercent, true
}

// HasCertificateLifetimePercent returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasCertificateLifetimePercent() bool {
	return o != nil && o.CertificateLifetimePercent != nil
}

// SetCertificateLifetimePercent gets a reference to the given float64 and assigns it to the CertificateLifetimePercent field.
func (o *NetworkHealthInsightAttributes) SetCertificateLifetimePercent(v float64) {
	o.CertificateLifetimePercent = &v
}

// GetClientRegion returns the ClientRegion field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetClientRegion() string {
	if o == nil || o.ClientRegion == nil {
		var ret string
		return ret
	}
	return *o.ClientRegion
}

// GetClientRegionOk returns a tuple with the ClientRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetClientRegionOk() (*string, bool) {
	if o == nil || o.ClientRegion == nil {
		return nil, false
	}
	return o.ClientRegion, true
}

// HasClientRegion returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasClientRegion() bool {
	return o != nil && o.ClientRegion != nil
}

// SetClientRegion gets a reference to the given string and assigns it to the ClientRegion field.
func (o *NetworkHealthInsightAttributes) SetClientRegion(v string) {
	o.ClientRegion = &v
}

// GetClientService returns the ClientService field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetClientService() string {
	if o == nil || o.ClientService == nil {
		var ret string
		return ret
	}
	return *o.ClientService
}

// GetClientServiceOk returns a tuple with the ClientService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetClientServiceOk() (*string, bool) {
	if o == nil || o.ClientService == nil {
		return nil, false
	}
	return o.ClientService, true
}

// HasClientService returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasClientService() bool {
	return o != nil && o.ClientService != nil
}

// SetClientService gets a reference to the given string and assigns it to the ClientService field.
func (o *NetworkHealthInsightAttributes) SetClientService(v string) {
	o.ClientService = &v
}

// GetDaysUntilExpiration returns the DaysUntilExpiration field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetDaysUntilExpiration() int64 {
	if o == nil || o.DaysUntilExpiration == nil {
		var ret int64
		return ret
	}
	return *o.DaysUntilExpiration
}

// GetDaysUntilExpirationOk returns a tuple with the DaysUntilExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetDaysUntilExpirationOk() (*int64, bool) {
	if o == nil || o.DaysUntilExpiration == nil {
		return nil, false
	}
	return o.DaysUntilExpiration, true
}

// HasDaysUntilExpiration returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasDaysUntilExpiration() bool {
	return o != nil && o.DaysUntilExpiration != nil
}

// SetDaysUntilExpiration gets a reference to the given int64 and assigns it to the DaysUntilExpiration field.
func (o *NetworkHealthInsightAttributes) SetDaysUntilExpiration(v int64) {
	o.DaysUntilExpiration = &v
}

// GetDnsQuery returns the DnsQuery field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetDnsQuery() string {
	if o == nil || o.DnsQuery == nil {
		var ret string
		return ret
	}
	return *o.DnsQuery
}

// GetDnsQueryOk returns a tuple with the DnsQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetDnsQueryOk() (*string, bool) {
	if o == nil || o.DnsQuery == nil {
		return nil, false
	}
	return o.DnsQuery, true
}

// HasDnsQuery returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasDnsQuery() bool {
	return o != nil && o.DnsQuery != nil
}

// SetDnsQuery gets a reference to the given string and assigns it to the DnsQuery field.
func (o *NetworkHealthInsightAttributes) SetDnsQuery(v string) {
	o.DnsQuery = &v
}

// GetDnsServer returns the DnsServer field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetDnsServer() string {
	if o == nil || o.DnsServer == nil {
		var ret string
		return ret
	}
	return *o.DnsServer
}

// GetDnsServerOk returns a tuple with the DnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetDnsServerOk() (*string, bool) {
	if o == nil || o.DnsServer == nil {
		return nil, false
	}
	return o.DnsServer, true
}

// HasDnsServer returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasDnsServer() bool {
	return o != nil && o.DnsServer != nil
}

// SetDnsServer gets a reference to the given string and assigns it to the DnsServer field.
func (o *NetworkHealthInsightAttributes) SetDnsServer(v string) {
	o.DnsServer = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasDomainName() bool {
	return o != nil && o.DomainName != nil
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *NetworkHealthInsightAttributes) SetDomainName(v string) {
	o.DomainName = &v
}

// GetFailureMagnitude returns the FailureMagnitude field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetFailureMagnitude() int64 {
	if o == nil || o.FailureMagnitude == nil {
		var ret int64
		return ret
	}
	return *o.FailureMagnitude
}

// GetFailureMagnitudeOk returns a tuple with the FailureMagnitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetFailureMagnitudeOk() (*int64, bool) {
	if o == nil || o.FailureMagnitude == nil {
		return nil, false
	}
	return o.FailureMagnitude, true
}

// HasFailureMagnitude returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasFailureMagnitude() bool {
	return o != nil && o.FailureMagnitude != nil
}

// SetFailureMagnitude gets a reference to the given int64 and assigns it to the FailureMagnitude field.
func (o *NetworkHealthInsightAttributes) SetFailureMagnitude(v int64) {
	o.FailureMagnitude = &v
}

// GetFailureRate returns the FailureRate field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetFailureRate() float64 {
	if o == nil || o.FailureRate == nil {
		var ret float64
		return ret
	}
	return *o.FailureRate
}

// GetFailureRateOk returns a tuple with the FailureRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetFailureRateOk() (*float64, bool) {
	if o == nil || o.FailureRate == nil {
		return nil, false
	}
	return o.FailureRate, true
}

// HasFailureRate returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasFailureRate() bool {
	return o != nil && o.FailureRate != nil
}

// SetFailureRate gets a reference to the given float64 and assigns it to the FailureRate field.
func (o *NetworkHealthInsightAttributes) SetFailureRate(v float64) {
	o.FailureRate = &v
}

// GetFailureType returns the FailureType field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetFailureType() NetworkHealthInsightFailureType {
	if o == nil || o.FailureType == nil {
		var ret NetworkHealthInsightFailureType
		return ret
	}
	return *o.FailureType
}

// GetFailureTypeOk returns a tuple with the FailureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetFailureTypeOk() (*NetworkHealthInsightFailureType, bool) {
	if o == nil || o.FailureType == nil {
		return nil, false
	}
	return o.FailureType, true
}

// HasFailureType returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasFailureType() bool {
	return o != nil && o.FailureType != nil
}

// SetFailureType gets a reference to the given NetworkHealthInsightFailureType and assigns it to the FailureType field.
func (o *NetworkHealthInsightAttributes) SetFailureType(v NetworkHealthInsightFailureType) {
	o.FailureType = &v
}

// GetLoadbalancerId returns the LoadbalancerId field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetLoadbalancerId() string {
	if o == nil || o.LoadbalancerId == nil {
		var ret string
		return ret
	}
	return *o.LoadbalancerId
}

// GetLoadbalancerIdOk returns a tuple with the LoadbalancerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetLoadbalancerIdOk() (*string, bool) {
	if o == nil || o.LoadbalancerId == nil {
		return nil, false
	}
	return o.LoadbalancerId, true
}

// HasLoadbalancerId returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasLoadbalancerId() bool {
	return o != nil && o.LoadbalancerId != nil
}

// SetLoadbalancerId gets a reference to the given string and assigns it to the LoadbalancerId field.
func (o *NetworkHealthInsightAttributes) SetLoadbalancerId(v string) {
	o.LoadbalancerId = &v
}

// GetServerRegion returns the ServerRegion field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetServerRegion() string {
	if o == nil || o.ServerRegion == nil {
		var ret string
		return ret
	}
	return *o.ServerRegion
}

// GetServerRegionOk returns a tuple with the ServerRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetServerRegionOk() (*string, bool) {
	if o == nil || o.ServerRegion == nil {
		return nil, false
	}
	return o.ServerRegion, true
}

// HasServerRegion returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasServerRegion() bool {
	return o != nil && o.ServerRegion != nil
}

// SetServerRegion gets a reference to the given string and assigns it to the ServerRegion field.
func (o *NetworkHealthInsightAttributes) SetServerRegion(v string) {
	o.ServerRegion = &v
}

// GetServerService returns the ServerService field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetServerService() string {
	if o == nil || o.ServerService == nil {
		var ret string
		return ret
	}
	return *o.ServerService
}

// GetServerServiceOk returns a tuple with the ServerService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetServerServiceOk() (*string, bool) {
	if o == nil || o.ServerService == nil {
		return nil, false
	}
	return o.ServerService, true
}

// HasServerService returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasServerService() bool {
	return o != nil && o.ServerService != nil
}

// SetServerService gets a reference to the given string and assigns it to the ServerService field.
func (o *NetworkHealthInsightAttributes) SetServerService(v string) {
	o.ServerService = &v
}

// GetTotalRequests returns the TotalRequests field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetTotalRequests() int64 {
	if o == nil || o.TotalRequests == nil {
		var ret int64
		return ret
	}
	return *o.TotalRequests
}

// GetTotalRequestsOk returns a tuple with the TotalRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetTotalRequestsOk() (*int64, bool) {
	if o == nil || o.TotalRequests == nil {
		return nil, false
	}
	return o.TotalRequests, true
}

// HasTotalRequests returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasTotalRequests() bool {
	return o != nil && o.TotalRequests != nil
}

// SetTotalRequests gets a reference to the given int64 and assigns it to the TotalRequests field.
func (o *NetworkHealthInsightAttributes) SetTotalRequests(v int64) {
	o.TotalRequests = &v
}

// GetTrafficVolume returns the TrafficVolume field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetTrafficVolume() NetworkHealthInsightTrafficVolume {
	if o == nil || o.TrafficVolume == nil {
		var ret NetworkHealthInsightTrafficVolume
		return ret
	}
	return *o.TrafficVolume
}

// GetTrafficVolumeOk returns a tuple with the TrafficVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetTrafficVolumeOk() (*NetworkHealthInsightTrafficVolume, bool) {
	if o == nil || o.TrafficVolume == nil {
		return nil, false
	}
	return o.TrafficVolume, true
}

// HasTrafficVolume returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasTrafficVolume() bool {
	return o != nil && o.TrafficVolume != nil
}

// SetTrafficVolume gets a reference to the given NetworkHealthInsightTrafficVolume and assigns it to the TrafficVolume field.
func (o *NetworkHealthInsightAttributes) SetTrafficVolume(v NetworkHealthInsightTrafficVolume) {
	o.TrafficVolume = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkHealthInsightAttributes) GetType() NetworkHealthInsightCategory {
	if o == nil || o.Type == nil {
		var ret NetworkHealthInsightCategory
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHealthInsightAttributes) GetTypeOk() (*NetworkHealthInsightCategory, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkHealthInsightAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given NetworkHealthInsightCategory and assigns it to the Type field.
func (o *NetworkHealthInsightAttributes) SetType(v NetworkHealthInsightCategory) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NetworkHealthInsightAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.CertificateId != nil {
		toSerialize["certificate_id"] = o.CertificateId
	}
	if o.CertificateLifetimePercent != nil {
		toSerialize["certificate_lifetime_percent"] = o.CertificateLifetimePercent
	}
	if o.ClientRegion != nil {
		toSerialize["client_region"] = o.ClientRegion
	}
	if o.ClientService != nil {
		toSerialize["client_service"] = o.ClientService
	}
	if o.DaysUntilExpiration != nil {
		toSerialize["days_until_expiration"] = o.DaysUntilExpiration
	}
	if o.DnsQuery != nil {
		toSerialize["dns_query"] = o.DnsQuery
	}
	if o.DnsServer != nil {
		toSerialize["dns_server"] = o.DnsServer
	}
	if o.DomainName != nil {
		toSerialize["domain_name"] = o.DomainName
	}
	if o.FailureMagnitude != nil {
		toSerialize["failure_magnitude"] = o.FailureMagnitude
	}
	if o.FailureRate != nil {
		toSerialize["failure_rate"] = o.FailureRate
	}
	if o.FailureType != nil {
		toSerialize["failure_type"] = o.FailureType
	}
	if o.LoadbalancerId != nil {
		toSerialize["loadbalancer_id"] = o.LoadbalancerId
	}
	if o.ServerRegion != nil {
		toSerialize["server_region"] = o.ServerRegion
	}
	if o.ServerService != nil {
		toSerialize["server_service"] = o.ServerService
	}
	if o.TotalRequests != nil {
		toSerialize["total_requests"] = o.TotalRequests
	}
	if o.TrafficVolume != nil {
		toSerialize["traffic_volume"] = o.TrafficVolume
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NetworkHealthInsightAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId                  *string                            `json:"account_id,omitempty"`
		CertificateId              *string                            `json:"certificate_id,omitempty"`
		CertificateLifetimePercent *float64                           `json:"certificate_lifetime_percent,omitempty"`
		ClientRegion               *string                            `json:"client_region,omitempty"`
		ClientService              *string                            `json:"client_service,omitempty"`
		DaysUntilExpiration        *int64                             `json:"days_until_expiration,omitempty"`
		DnsQuery                   *string                            `json:"dns_query,omitempty"`
		DnsServer                  *string                            `json:"dns_server,omitempty"`
		DomainName                 *string                            `json:"domain_name,omitempty"`
		FailureMagnitude           *int64                             `json:"failure_magnitude,omitempty"`
		FailureRate                *float64                           `json:"failure_rate,omitempty"`
		FailureType                *NetworkHealthInsightFailureType   `json:"failure_type,omitempty"`
		LoadbalancerId             *string                            `json:"loadbalancer_id,omitempty"`
		ServerRegion               *string                            `json:"server_region,omitempty"`
		ServerService              *string                            `json:"server_service,omitempty"`
		TotalRequests              *int64                             `json:"total_requests,omitempty"`
		TrafficVolume              *NetworkHealthInsightTrafficVolume `json:"traffic_volume,omitempty"`
		Type                       *NetworkHealthInsightCategory      `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "certificate_id", "certificate_lifetime_percent", "client_region", "client_service", "days_until_expiration", "dns_query", "dns_server", "domain_name", "failure_magnitude", "failure_rate", "failure_type", "loadbalancer_id", "server_region", "server_service", "total_requests", "traffic_volume", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = all.AccountId
	o.CertificateId = all.CertificateId
	o.CertificateLifetimePercent = all.CertificateLifetimePercent
	o.ClientRegion = all.ClientRegion
	o.ClientService = all.ClientService
	o.DaysUntilExpiration = all.DaysUntilExpiration
	o.DnsQuery = all.DnsQuery
	o.DnsServer = all.DnsServer
	o.DomainName = all.DomainName
	o.FailureMagnitude = all.FailureMagnitude
	o.FailureRate = all.FailureRate
	if all.FailureType != nil && !all.FailureType.IsValid() {
		hasInvalidField = true
	} else {
		o.FailureType = all.FailureType
	}
	o.LoadbalancerId = all.LoadbalancerId
	o.ServerRegion = all.ServerRegion
	o.ServerService = all.ServerService
	o.TotalRequests = all.TotalRequests
	if all.TrafficVolume != nil && all.TrafficVolume.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TrafficVolume = all.TrafficVolume
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
