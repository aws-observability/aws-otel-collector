// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultCdnCacheStatus Cache status reported by the CDN for the response.
type SyntheticsTestResultCdnCacheStatus struct {
	// Whether the response was served from the CDN cache.
	Cached *bool `json:"cached,omitempty"`
	// Raw cache status string reported by the CDN.
	Status *string `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultCdnCacheStatus instantiates a new SyntheticsTestResultCdnCacheStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultCdnCacheStatus() *SyntheticsTestResultCdnCacheStatus {
	this := SyntheticsTestResultCdnCacheStatus{}
	return &this
}

// NewSyntheticsTestResultCdnCacheStatusWithDefaults instantiates a new SyntheticsTestResultCdnCacheStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultCdnCacheStatusWithDefaults() *SyntheticsTestResultCdnCacheStatus {
	this := SyntheticsTestResultCdnCacheStatus{}
	return &this
}

// GetCached returns the Cached field value if set, zero value otherwise.
func (o *SyntheticsTestResultCdnCacheStatus) GetCached() bool {
	if o == nil || o.Cached == nil {
		var ret bool
		return ret
	}
	return *o.Cached
}

// GetCachedOk returns a tuple with the Cached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCdnCacheStatus) GetCachedOk() (*bool, bool) {
	if o == nil || o.Cached == nil {
		return nil, false
	}
	return o.Cached, true
}

// HasCached returns a boolean if a field has been set.
func (o *SyntheticsTestResultCdnCacheStatus) HasCached() bool {
	return o != nil && o.Cached != nil
}

// SetCached gets a reference to the given bool and assigns it to the Cached field.
func (o *SyntheticsTestResultCdnCacheStatus) SetCached(v bool) {
	o.Cached = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestResultCdnCacheStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCdnCacheStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestResultCdnCacheStatus) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyntheticsTestResultCdnCacheStatus) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultCdnCacheStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cached != nil {
		toSerialize["cached"] = o.Cached
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultCdnCacheStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cached *bool   `json:"cached,omitempty"`
		Status *string `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cached", "status"})
	} else {
		return err
	}
	o.Cached = all.Cached
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
