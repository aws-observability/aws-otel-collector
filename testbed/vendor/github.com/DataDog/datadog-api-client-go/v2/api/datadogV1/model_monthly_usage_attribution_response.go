// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyUsageAttributionResponse Response containing the monthly Usage Summary by tag(s).
type MonthlyUsageAttributionResponse struct {
	// The object containing document metadata.
	Metadata *MonthlyUsageAttributionMetadata `json:"metadata,omitempty"`
	// Get usage summary by tag(s).
	Usage []MonthlyUsageAttributionBody `json:"usage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonthlyUsageAttributionResponse instantiates a new MonthlyUsageAttributionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyUsageAttributionResponse() *MonthlyUsageAttributionResponse {
	this := MonthlyUsageAttributionResponse{}
	return &this
}

// NewMonthlyUsageAttributionResponseWithDefaults instantiates a new MonthlyUsageAttributionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyUsageAttributionResponseWithDefaults() *MonthlyUsageAttributionResponse {
	this := MonthlyUsageAttributionResponse{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionResponse) GetMetadata() MonthlyUsageAttributionMetadata {
	if o == nil || o.Metadata == nil {
		var ret MonthlyUsageAttributionMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionResponse) GetMetadataOk() (*MonthlyUsageAttributionMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionResponse) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given MonthlyUsageAttributionMetadata and assigns it to the Metadata field.
func (o *MonthlyUsageAttributionResponse) SetMetadata(v MonthlyUsageAttributionMetadata) {
	o.Metadata = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *MonthlyUsageAttributionResponse) GetUsage() []MonthlyUsageAttributionBody {
	if o == nil || o.Usage == nil {
		var ret []MonthlyUsageAttributionBody
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyUsageAttributionResponse) GetUsageOk() (*[]MonthlyUsageAttributionBody, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return &o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *MonthlyUsageAttributionResponse) HasUsage() bool {
	return o != nil && o.Usage != nil
}

// SetUsage gets a reference to the given []MonthlyUsageAttributionBody and assigns it to the Usage field.
func (o *MonthlyUsageAttributionResponse) SetUsage(v []MonthlyUsageAttributionBody) {
	o.Usage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyUsageAttributionResponse) MarshalJSON() ([]byte, error) {
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
func (o *MonthlyUsageAttributionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metadata *MonthlyUsageAttributionMetadata `json:"metadata,omitempty"`
		Usage    []MonthlyUsageAttributionBody    `json:"usage,omitempty"`
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
