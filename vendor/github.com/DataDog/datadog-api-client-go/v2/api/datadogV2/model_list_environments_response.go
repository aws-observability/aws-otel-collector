// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListEnvironmentsResponse Response containing a list of environments.
type ListEnvironmentsResponse struct {
	// List of environments.
	Data []Environment `json:"data"`
	// Pagination metadata for environments.
	Meta *EnvironmentsPaginationMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListEnvironmentsResponse instantiates a new ListEnvironmentsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListEnvironmentsResponse(data []Environment) *ListEnvironmentsResponse {
	this := ListEnvironmentsResponse{}
	this.Data = data
	return &this
}

// NewListEnvironmentsResponseWithDefaults instantiates a new ListEnvironmentsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListEnvironmentsResponseWithDefaults() *ListEnvironmentsResponse {
	this := ListEnvironmentsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ListEnvironmentsResponse) GetData() []Environment {
	if o == nil {
		var ret []Environment
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListEnvironmentsResponse) GetDataOk() (*[]Environment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ListEnvironmentsResponse) SetData(v []Environment) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListEnvironmentsResponse) GetMeta() EnvironmentsPaginationMeta {
	if o == nil || o.Meta == nil {
		var ret EnvironmentsPaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEnvironmentsResponse) GetMetaOk() (*EnvironmentsPaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListEnvironmentsResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given EnvironmentsPaginationMeta and assigns it to the Meta field.
func (o *ListEnvironmentsResponse) SetMeta(v EnvironmentsPaginationMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListEnvironmentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListEnvironmentsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]Environment              `json:"data"`
		Meta *EnvironmentsPaginationMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
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
