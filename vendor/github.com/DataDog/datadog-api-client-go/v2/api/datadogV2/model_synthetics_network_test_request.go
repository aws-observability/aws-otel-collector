// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsNetworkTestRequest Object describing the request for a Network Path test.
type SyntheticsNetworkTestRequest struct {
	// An optional label displayed for the destination host in the Network Path visualization.
	DestinationService *string `json:"destination_service,omitempty"`
	// The number of packets sent to probe the destination to measure packet loss, latency and jitter.
	E2eQueries int64 `json:"e2e_queries"`
	// Host name to query.
	Host string `json:"host"`
	// The maximum time-to-live (max number of hops) used in outgoing probe packets.
	MaxTtl int64 `json:"max_ttl"`
	// For TCP or UDP tests, the port to use when performing the test.
	// If not set on a UDP test, a random port is assigned, which may affect the results.
	Port *int64 `json:"port,omitempty"`
	// An optional label displayed for the source host in the Network Path visualization.
	SourceService *string `json:"source_service,omitempty"`
	// For TCP tests, the TCP traceroute strategy.
	TcpMethod *SyntheticsNetworkTestRequestTCPMethod `json:"tcp_method,omitempty"`
	// Timeout in seconds.
	Timeout *int64 `json:"timeout,omitempty"`
	// The number of traceroute path tracings.
	TracerouteQueries int64 `json:"traceroute_queries"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsNetworkTestRequest instantiates a new SyntheticsNetworkTestRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsNetworkTestRequest(e2eQueries int64, host string, maxTtl int64, tracerouteQueries int64) *SyntheticsNetworkTestRequest {
	this := SyntheticsNetworkTestRequest{}
	this.E2eQueries = e2eQueries
	this.Host = host
	this.MaxTtl = maxTtl
	this.TracerouteQueries = tracerouteQueries
	return &this
}

// NewSyntheticsNetworkTestRequestWithDefaults instantiates a new SyntheticsNetworkTestRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsNetworkTestRequestWithDefaults() *SyntheticsNetworkTestRequest {
	this := SyntheticsNetworkTestRequest{}
	return &this
}

// GetDestinationService returns the DestinationService field value if set, zero value otherwise.
func (o *SyntheticsNetworkTestRequest) GetDestinationService() string {
	if o == nil || o.DestinationService == nil {
		var ret string
		return ret
	}
	return *o.DestinationService
}

// GetDestinationServiceOk returns a tuple with the DestinationService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestRequest) GetDestinationServiceOk() (*string, bool) {
	if o == nil || o.DestinationService == nil {
		return nil, false
	}
	return o.DestinationService, true
}

// HasDestinationService returns a boolean if a field has been set.
func (o *SyntheticsNetworkTestRequest) HasDestinationService() bool {
	return o != nil && o.DestinationService != nil
}

// SetDestinationService gets a reference to the given string and assigns it to the DestinationService field.
func (o *SyntheticsNetworkTestRequest) SetDestinationService(v string) {
	o.DestinationService = &v
}

// GetE2eQueries returns the E2eQueries field value.
func (o *SyntheticsNetworkTestRequest) GetE2eQueries() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.E2eQueries
}

// GetE2eQueriesOk returns a tuple with the E2eQueries field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestRequest) GetE2eQueriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.E2eQueries, true
}

// SetE2eQueries sets field value.
func (o *SyntheticsNetworkTestRequest) SetE2eQueries(v int64) {
	o.E2eQueries = v
}

// GetHost returns the Host field value.
func (o *SyntheticsNetworkTestRequest) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestRequest) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value.
func (o *SyntheticsNetworkTestRequest) SetHost(v string) {
	o.Host = v
}

// GetMaxTtl returns the MaxTtl field value.
func (o *SyntheticsNetworkTestRequest) GetMaxTtl() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxTtl
}

// GetMaxTtlOk returns a tuple with the MaxTtl field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestRequest) GetMaxTtlOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTtl, true
}

// SetMaxTtl sets field value.
func (o *SyntheticsNetworkTestRequest) SetMaxTtl(v int64) {
	o.MaxTtl = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SyntheticsNetworkTestRequest) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestRequest) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SyntheticsNetworkTestRequest) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SyntheticsNetworkTestRequest) SetPort(v int64) {
	o.Port = &v
}

// GetSourceService returns the SourceService field value if set, zero value otherwise.
func (o *SyntheticsNetworkTestRequest) GetSourceService() string {
	if o == nil || o.SourceService == nil {
		var ret string
		return ret
	}
	return *o.SourceService
}

// GetSourceServiceOk returns a tuple with the SourceService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestRequest) GetSourceServiceOk() (*string, bool) {
	if o == nil || o.SourceService == nil {
		return nil, false
	}
	return o.SourceService, true
}

// HasSourceService returns a boolean if a field has been set.
func (o *SyntheticsNetworkTestRequest) HasSourceService() bool {
	return o != nil && o.SourceService != nil
}

// SetSourceService gets a reference to the given string and assigns it to the SourceService field.
func (o *SyntheticsNetworkTestRequest) SetSourceService(v string) {
	o.SourceService = &v
}

// GetTcpMethod returns the TcpMethod field value if set, zero value otherwise.
func (o *SyntheticsNetworkTestRequest) GetTcpMethod() SyntheticsNetworkTestRequestTCPMethod {
	if o == nil || o.TcpMethod == nil {
		var ret SyntheticsNetworkTestRequestTCPMethod
		return ret
	}
	return *o.TcpMethod
}

// GetTcpMethodOk returns a tuple with the TcpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestRequest) GetTcpMethodOk() (*SyntheticsNetworkTestRequestTCPMethod, bool) {
	if o == nil || o.TcpMethod == nil {
		return nil, false
	}
	return o.TcpMethod, true
}

// HasTcpMethod returns a boolean if a field has been set.
func (o *SyntheticsNetworkTestRequest) HasTcpMethod() bool {
	return o != nil && o.TcpMethod != nil
}

// SetTcpMethod gets a reference to the given SyntheticsNetworkTestRequestTCPMethod and assigns it to the TcpMethod field.
func (o *SyntheticsNetworkTestRequest) SetTcpMethod(v SyntheticsNetworkTestRequestTCPMethod) {
	o.TcpMethod = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *SyntheticsNetworkTestRequest) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestRequest) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *SyntheticsNetworkTestRequest) HasTimeout() bool {
	return o != nil && o.Timeout != nil
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *SyntheticsNetworkTestRequest) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetTracerouteQueries returns the TracerouteQueries field value.
func (o *SyntheticsNetworkTestRequest) GetTracerouteQueries() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TracerouteQueries
}

// GetTracerouteQueriesOk returns a tuple with the TracerouteQueries field value
// and a boolean to check if the value has been set.
func (o *SyntheticsNetworkTestRequest) GetTracerouteQueriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TracerouteQueries, true
}

// SetTracerouteQueries sets field value.
func (o *SyntheticsNetworkTestRequest) SetTracerouteQueries(v int64) {
	o.TracerouteQueries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsNetworkTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DestinationService != nil {
		toSerialize["destination_service"] = o.DestinationService
	}
	toSerialize["e2e_queries"] = o.E2eQueries
	toSerialize["host"] = o.Host
	toSerialize["max_ttl"] = o.MaxTtl
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.SourceService != nil {
		toSerialize["source_service"] = o.SourceService
	}
	if o.TcpMethod != nil {
		toSerialize["tcp_method"] = o.TcpMethod
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	toSerialize["traceroute_queries"] = o.TracerouteQueries

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsNetworkTestRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DestinationService *string                                `json:"destination_service,omitempty"`
		E2eQueries         *int64                                 `json:"e2e_queries"`
		Host               *string                                `json:"host"`
		MaxTtl             *int64                                 `json:"max_ttl"`
		Port               *int64                                 `json:"port,omitempty"`
		SourceService      *string                                `json:"source_service,omitempty"`
		TcpMethod          *SyntheticsNetworkTestRequestTCPMethod `json:"tcp_method,omitempty"`
		Timeout            *int64                                 `json:"timeout,omitempty"`
		TracerouteQueries  *int64                                 `json:"traceroute_queries"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.E2eQueries == nil {
		return fmt.Errorf("required field e2e_queries missing")
	}
	if all.Host == nil {
		return fmt.Errorf("required field host missing")
	}
	if all.MaxTtl == nil {
		return fmt.Errorf("required field max_ttl missing")
	}
	if all.TracerouteQueries == nil {
		return fmt.Errorf("required field traceroute_queries missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destination_service", "e2e_queries", "host", "max_ttl", "port", "source_service", "tcp_method", "timeout", "traceroute_queries"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DestinationService = all.DestinationService
	o.E2eQueries = *all.E2eQueries
	o.Host = *all.Host
	o.MaxTtl = *all.MaxTtl
	o.Port = all.Port
	o.SourceService = all.SourceService
	if all.TcpMethod != nil && !all.TcpMethod.IsValid() {
		hasInvalidField = true
	} else {
		o.TcpMethod = all.TcpMethod
	}
	o.Timeout = all.Timeout
	o.TracerouteQueries = *all.TracerouteQueries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
