// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentGatesListResponse Response containing a paginated list of deployment gates.
type DeploymentGatesListResponse struct {
	// Array of deployment gates.
	Data []DeploymentGateResponseData `json:"data,omitempty"`
	// Metadata for a list of deployment gates response.
	Meta *DeploymentGatesListResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentGatesListResponse instantiates a new DeploymentGatesListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentGatesListResponse() *DeploymentGatesListResponse {
	this := DeploymentGatesListResponse{}
	return &this
}

// NewDeploymentGatesListResponseWithDefaults instantiates a new DeploymentGatesListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentGatesListResponseWithDefaults() *DeploymentGatesListResponse {
	this := DeploymentGatesListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeploymentGatesListResponse) GetData() []DeploymentGateResponseData {
	if o == nil || o.Data == nil {
		var ret []DeploymentGateResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesListResponse) GetDataOk() (*[]DeploymentGateResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeploymentGatesListResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []DeploymentGateResponseData and assigns it to the Data field.
func (o *DeploymentGatesListResponse) SetData(v []DeploymentGateResponseData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DeploymentGatesListResponse) GetMeta() DeploymentGatesListResponseMeta {
	if o == nil || o.Meta == nil {
		var ret DeploymentGatesListResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentGatesListResponse) GetMetaOk() (*DeploymentGatesListResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DeploymentGatesListResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given DeploymentGatesListResponseMeta and assigns it to the Meta field.
func (o *DeploymentGatesListResponse) SetMeta(v DeploymentGatesListResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentGatesListResponse) MarshalJSON() ([]byte, error) {
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
func (o *DeploymentGatesListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []DeploymentGateResponseData     `json:"data,omitempty"`
		Meta *DeploymentGatesListResponseMeta `json:"meta,omitempty"`
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
