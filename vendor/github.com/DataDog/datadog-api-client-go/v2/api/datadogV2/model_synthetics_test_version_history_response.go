// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionHistoryResponse Response containing the paginated version history for a Synthetic test.
type SyntheticsTestVersionHistoryResponse struct {
	// List of version change records.
	Data []SyntheticsTestVersionChangeData `json:"data,omitempty"`
	// Pagination metadata for a version history response.
	Meta *SyntheticsTestVersionHistoryMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestVersionHistoryResponse instantiates a new SyntheticsTestVersionHistoryResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestVersionHistoryResponse() *SyntheticsTestVersionHistoryResponse {
	this := SyntheticsTestVersionHistoryResponse{}
	return &this
}

// NewSyntheticsTestVersionHistoryResponseWithDefaults instantiates a new SyntheticsTestVersionHistoryResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestVersionHistoryResponseWithDefaults() *SyntheticsTestVersionHistoryResponse {
	this := SyntheticsTestVersionHistoryResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SyntheticsTestVersionHistoryResponse) GetData() []SyntheticsTestVersionChangeData {
	if o == nil || o.Data == nil {
		var ret []SyntheticsTestVersionChangeData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionHistoryResponse) GetDataOk() (*[]SyntheticsTestVersionChangeData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SyntheticsTestVersionHistoryResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []SyntheticsTestVersionChangeData and assigns it to the Data field.
func (o *SyntheticsTestVersionHistoryResponse) SetData(v []SyntheticsTestVersionChangeData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SyntheticsTestVersionHistoryResponse) GetMeta() SyntheticsTestVersionHistoryMeta {
	if o == nil || o.Meta == nil {
		var ret SyntheticsTestVersionHistoryMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionHistoryResponse) GetMetaOk() (*SyntheticsTestVersionHistoryMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SyntheticsTestVersionHistoryResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given SyntheticsTestVersionHistoryMeta and assigns it to the Meta field.
func (o *SyntheticsTestVersionHistoryResponse) SetMeta(v SyntheticsTestVersionHistoryMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestVersionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestVersionHistoryResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []SyntheticsTestVersionChangeData `json:"data,omitempty"`
		Meta *SyntheticsTestVersionHistoryMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
