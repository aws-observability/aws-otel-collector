// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HourlyUsageAttributionResponse Response containing the hourly usage attribution by tag(s).
type HourlyUsageAttributionResponse struct {
	// The object containing document metadata.
	Metadata *HourlyUsageAttributionMetadata `json:"metadata,omitempty"`
	// Get the hourly usage attribution by tag(s).
	Usage []HourlyUsageAttributionBody `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHourlyUsageAttributionResponse instantiates a new HourlyUsageAttributionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHourlyUsageAttributionResponse() *HourlyUsageAttributionResponse {
	this := HourlyUsageAttributionResponse{}
	return &this
}

// NewHourlyUsageAttributionResponseWithDefaults instantiates a new HourlyUsageAttributionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHourlyUsageAttributionResponseWithDefaults() *HourlyUsageAttributionResponse {
	this := HourlyUsageAttributionResponse{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HourlyUsageAttributionResponse) GetMetadata() HourlyUsageAttributionMetadata {
	if o == nil || o.Metadata == nil {
		var ret HourlyUsageAttributionMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HourlyUsageAttributionResponse) GetMetadataOk() (*HourlyUsageAttributionMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HourlyUsageAttributionResponse) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given HourlyUsageAttributionMetadata and assigns it to the Metadata field.
func (o *HourlyUsageAttributionResponse) SetMetadata(v HourlyUsageAttributionMetadata) {
	o.Metadata = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *HourlyUsageAttributionResponse) GetUsage() []HourlyUsageAttributionBody {
	if o == nil || o.Usage == nil {
		var ret []HourlyUsageAttributionBody
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HourlyUsageAttributionResponse) GetUsageOk() (*[]HourlyUsageAttributionBody, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *HourlyUsageAttributionResponse) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given []HourlyUsageAttributionBody and assigns it to the Usage field.
func (o *HourlyUsageAttributionResponse) SetUsage(v []HourlyUsageAttributionBody) {
	o.Usage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HourlyUsageAttributionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
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
func (o *HourlyUsageAttributionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metadata *HourlyUsageAttributionMetadata `json:"metadata,omitempty"`
		Usage    []HourlyUsageAttributionBody    `json:"usage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metadata", "usage"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	o.Usage = all.Usage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
