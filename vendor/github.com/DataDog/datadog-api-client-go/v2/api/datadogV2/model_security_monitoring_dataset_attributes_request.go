// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetAttributesRequest The attributes of a dataset create or update request.
type SecurityMonitoringDatasetAttributesRequest struct {
	// The definition of the dataset. The shape depends on the value of `data_source`.
	// Use `reference_table` or `managed_resource` for a referential dataset, or one of the
	// event platform sources (for example `logs`, `audit`, `events`, `spans`, `rum`) for
	// an event platform dataset.
	Definition SecurityMonitoringDatasetDefinition `json:"definition"`
	// The description of the dataset. Maximum 255 characters.
	Description *string `json:"description,omitempty"`
	// The expected current version of the dataset for optimistic concurrency control on updates.
	// If the dataset's current version does not match, the request is rejected with a 409 Conflict.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetAttributesRequest instantiates a new SecurityMonitoringDatasetAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetAttributesRequest(definition SecurityMonitoringDatasetDefinition) *SecurityMonitoringDatasetAttributesRequest {
	this := SecurityMonitoringDatasetAttributesRequest{}
	this.Definition = definition
	return &this
}

// NewSecurityMonitoringDatasetAttributesRequestWithDefaults instantiates a new SecurityMonitoringDatasetAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetAttributesRequestWithDefaults() *SecurityMonitoringDatasetAttributesRequest {
	this := SecurityMonitoringDatasetAttributesRequest{}
	return &this
}

// GetDefinition returns the Definition field value.
func (o *SecurityMonitoringDatasetAttributesRequest) GetDefinition() SecurityMonitoringDatasetDefinition {
	if o == nil {
		var ret SecurityMonitoringDatasetDefinition
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesRequest) GetDefinitionOk() (*SecurityMonitoringDatasetDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value.
func (o *SecurityMonitoringDatasetAttributesRequest) SetDefinition(v SecurityMonitoringDatasetDefinition) {
	o.Definition = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityMonitoringDatasetAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SecurityMonitoringDatasetAttributesRequest) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetAttributesRequest) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SecurityMonitoringDatasetAttributesRequest) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SecurityMonitoringDatasetAttributesRequest) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["definition"] = o.Definition
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Definition  *SecurityMonitoringDatasetDefinition `json:"definition"`
		Description *string                              `json:"description,omitempty"`
		Version     *int64                               `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Definition == nil {
		return fmt.Errorf("required field definition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"definition", "description", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Definition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Definition = *all.Definition
	o.Description = all.Description
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
