// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageNetworkHostsResponse Response containing the number of active NPM hosts for each hour for a given organization.
type UsageNetworkHostsResponse struct {
	// Get hourly usage for NPM hosts.
	Usage []UsageNetworkHostsHour `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageNetworkHostsResponse instantiates a new UsageNetworkHostsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageNetworkHostsResponse() *UsageNetworkHostsResponse {
	this := UsageNetworkHostsResponse{}
	return &this
}

// NewUsageNetworkHostsResponseWithDefaults instantiates a new UsageNetworkHostsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageNetworkHostsResponseWithDefaults() *UsageNetworkHostsResponse {
	this := UsageNetworkHostsResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageNetworkHostsResponse) GetUsage() []UsageNetworkHostsHour {
	if o == nil || o.Usage == nil {
		var ret []UsageNetworkHostsHour
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageNetworkHostsResponse) GetUsageOk() (*[]UsageNetworkHostsHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageNetworkHostsResponse) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given []UsageNetworkHostsHour and assigns it to the Usage field.
func (o *UsageNetworkHostsResponse) SetUsage(v []UsageNetworkHostsHour) {
	o.Usage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageNetworkHostsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageNetworkHostsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Usage []UsageNetworkHostsHour `json:"usage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"usage"})
	} else {
		return err
	}
	o.Usage = all.Usage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
