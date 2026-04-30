// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultCdnResource A CDN resource encountered while executing a browser step.
type SyntheticsTestResultCdnResource struct {
	// CDN provider details inferred from response headers.
	Cdn *SyntheticsTestResultCdnProviderInfo `json:"cdn,omitempty"`
	// Resolved IP address for the CDN resource.
	ResolvedIp *string `json:"resolved_ip,omitempty"`
	// Unix timestamp (ms) of when the resource was fetched.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// Timing breakdown for fetching the CDN resource.
	Timings map[string]interface{} `json:"timings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultCdnResource instantiates a new SyntheticsTestResultCdnResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultCdnResource() *SyntheticsTestResultCdnResource {
	this := SyntheticsTestResultCdnResource{}
	return &this
}

// NewSyntheticsTestResultCdnResourceWithDefaults instantiates a new SyntheticsTestResultCdnResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultCdnResourceWithDefaults() *SyntheticsTestResultCdnResource {
	this := SyntheticsTestResultCdnResource{}
	return &this
}

// GetCdn returns the Cdn field value if set, zero value otherwise.
func (o *SyntheticsTestResultCdnResource) GetCdn() SyntheticsTestResultCdnProviderInfo {
	if o == nil || o.Cdn == nil {
		var ret SyntheticsTestResultCdnProviderInfo
		return ret
	}
	return *o.Cdn
}

// GetCdnOk returns a tuple with the Cdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCdnResource) GetCdnOk() (*SyntheticsTestResultCdnProviderInfo, bool) {
	if o == nil || o.Cdn == nil {
		return nil, false
	}
	return o.Cdn, true
}

// HasCdn returns a boolean if a field has been set.
func (o *SyntheticsTestResultCdnResource) HasCdn() bool {
	return o != nil && o.Cdn != nil
}

// SetCdn gets a reference to the given SyntheticsTestResultCdnProviderInfo and assigns it to the Cdn field.
func (o *SyntheticsTestResultCdnResource) SetCdn(v SyntheticsTestResultCdnProviderInfo) {
	o.Cdn = &v
}

// GetResolvedIp returns the ResolvedIp field value if set, zero value otherwise.
func (o *SyntheticsTestResultCdnResource) GetResolvedIp() string {
	if o == nil || o.ResolvedIp == nil {
		var ret string
		return ret
	}
	return *o.ResolvedIp
}

// GetResolvedIpOk returns a tuple with the ResolvedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCdnResource) GetResolvedIpOk() (*string, bool) {
	if o == nil || o.ResolvedIp == nil {
		return nil, false
	}
	return o.ResolvedIp, true
}

// HasResolvedIp returns a boolean if a field has been set.
func (o *SyntheticsTestResultCdnResource) HasResolvedIp() bool {
	return o != nil && o.ResolvedIp != nil
}

// SetResolvedIp gets a reference to the given string and assigns it to the ResolvedIp field.
func (o *SyntheticsTestResultCdnResource) SetResolvedIp(v string) {
	o.ResolvedIp = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SyntheticsTestResultCdnResource) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCdnResource) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SyntheticsTestResultCdnResource) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *SyntheticsTestResultCdnResource) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetTimings returns the Timings field value if set, zero value otherwise.
func (o *SyntheticsTestResultCdnResource) GetTimings() map[string]interface{} {
	if o == nil || o.Timings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Timings
}

// GetTimingsOk returns a tuple with the Timings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCdnResource) GetTimingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Timings == nil {
		return nil, false
	}
	return &o.Timings, true
}

// HasTimings returns a boolean if a field has been set.
func (o *SyntheticsTestResultCdnResource) HasTimings() bool {
	return o != nil && o.Timings != nil
}

// SetTimings gets a reference to the given map[string]interface{} and assigns it to the Timings field.
func (o *SyntheticsTestResultCdnResource) SetTimings(v map[string]interface{}) {
	o.Timings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultCdnResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cdn != nil {
		toSerialize["cdn"] = o.Cdn
	}
	if o.ResolvedIp != nil {
		toSerialize["resolved_ip"] = o.ResolvedIp
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Timings != nil {
		toSerialize["timings"] = o.Timings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultCdnResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cdn        *SyntheticsTestResultCdnProviderInfo `json:"cdn,omitempty"`
		ResolvedIp *string                              `json:"resolved_ip,omitempty"`
		Timestamp  *int64                               `json:"timestamp,omitempty"`
		Timings    map[string]interface{}               `json:"timings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cdn", "resolved_ip", "timestamp", "timings"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Cdn != nil && all.Cdn.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cdn = all.Cdn
	o.ResolvedIp = all.ResolvedIp
	o.Timestamp = all.Timestamp
	o.Timings = all.Timings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
