// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateConnectionRequestDataAttributesFieldsToUpdateItems
type UpdateConnectionRequestDataAttributesFieldsToUpdateItems struct {
	//
	FieldId string `json:"field_id"`
	//
	UpdatedDescription *string `json:"updated_description,omitempty"`
	//
	UpdatedDisplayName *string `json:"updated_display_name,omitempty"`
	//
	UpdatedFieldId *string `json:"updated_field_id,omitempty"`
	//
	UpdatedGroups []string `json:"updated_groups,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateConnectionRequestDataAttributesFieldsToUpdateItems instantiates a new UpdateConnectionRequestDataAttributesFieldsToUpdateItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateConnectionRequestDataAttributesFieldsToUpdateItems(fieldId string) *UpdateConnectionRequestDataAttributesFieldsToUpdateItems {
	this := UpdateConnectionRequestDataAttributesFieldsToUpdateItems{}
	this.FieldId = fieldId
	return &this
}

// NewUpdateConnectionRequestDataAttributesFieldsToUpdateItemsWithDefaults instantiates a new UpdateConnectionRequestDataAttributesFieldsToUpdateItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateConnectionRequestDataAttributesFieldsToUpdateItemsWithDefaults() *UpdateConnectionRequestDataAttributesFieldsToUpdateItems {
	this := UpdateConnectionRequestDataAttributesFieldsToUpdateItems{}
	return &this
}

// GetFieldId returns the FieldId field value.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetFieldId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value
// and a boolean to check if the value has been set.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetFieldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldId, true
}

// SetFieldId sets field value.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) SetFieldId(v string) {
	o.FieldId = v
}

// GetUpdatedDescription returns the UpdatedDescription field value if set, zero value otherwise.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetUpdatedDescription() string {
	if o == nil || o.UpdatedDescription == nil {
		var ret string
		return ret
	}
	return *o.UpdatedDescription
}

// GetUpdatedDescriptionOk returns a tuple with the UpdatedDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetUpdatedDescriptionOk() (*string, bool) {
	if o == nil || o.UpdatedDescription == nil {
		return nil, false
	}
	return o.UpdatedDescription, true
}

// HasUpdatedDescription returns a boolean if a field has been set.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) HasUpdatedDescription() bool {
	return o != nil && o.UpdatedDescription != nil
}

// SetUpdatedDescription gets a reference to the given string and assigns it to the UpdatedDescription field.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) SetUpdatedDescription(v string) {
	o.UpdatedDescription = &v
}

// GetUpdatedDisplayName returns the UpdatedDisplayName field value if set, zero value otherwise.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetUpdatedDisplayName() string {
	if o == nil || o.UpdatedDisplayName == nil {
		var ret string
		return ret
	}
	return *o.UpdatedDisplayName
}

// GetUpdatedDisplayNameOk returns a tuple with the UpdatedDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetUpdatedDisplayNameOk() (*string, bool) {
	if o == nil || o.UpdatedDisplayName == nil {
		return nil, false
	}
	return o.UpdatedDisplayName, true
}

// HasUpdatedDisplayName returns a boolean if a field has been set.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) HasUpdatedDisplayName() bool {
	return o != nil && o.UpdatedDisplayName != nil
}

// SetUpdatedDisplayName gets a reference to the given string and assigns it to the UpdatedDisplayName field.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) SetUpdatedDisplayName(v string) {
	o.UpdatedDisplayName = &v
}

// GetUpdatedFieldId returns the UpdatedFieldId field value if set, zero value otherwise.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetUpdatedFieldId() string {
	if o == nil || o.UpdatedFieldId == nil {
		var ret string
		return ret
	}
	return *o.UpdatedFieldId
}

// GetUpdatedFieldIdOk returns a tuple with the UpdatedFieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetUpdatedFieldIdOk() (*string, bool) {
	if o == nil || o.UpdatedFieldId == nil {
		return nil, false
	}
	return o.UpdatedFieldId, true
}

// HasUpdatedFieldId returns a boolean if a field has been set.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) HasUpdatedFieldId() bool {
	return o != nil && o.UpdatedFieldId != nil
}

// SetUpdatedFieldId gets a reference to the given string and assigns it to the UpdatedFieldId field.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) SetUpdatedFieldId(v string) {
	o.UpdatedFieldId = &v
}

// GetUpdatedGroups returns the UpdatedGroups field value if set, zero value otherwise.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetUpdatedGroups() []string {
	if o == nil || o.UpdatedGroups == nil {
		var ret []string
		return ret
	}
	return o.UpdatedGroups
}

// GetUpdatedGroupsOk returns a tuple with the UpdatedGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) GetUpdatedGroupsOk() (*[]string, bool) {
	if o == nil || o.UpdatedGroups == nil {
		return nil, false
	}
	return &o.UpdatedGroups, true
}

// HasUpdatedGroups returns a boolean if a field has been set.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) HasUpdatedGroups() bool {
	return o != nil && o.UpdatedGroups != nil
}

// SetUpdatedGroups gets a reference to the given []string and assigns it to the UpdatedGroups field.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) SetUpdatedGroups(v []string) {
	o.UpdatedGroups = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateConnectionRequestDataAttributesFieldsToUpdateItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field_id"] = o.FieldId
	if o.UpdatedDescription != nil {
		toSerialize["updated_description"] = o.UpdatedDescription
	}
	if o.UpdatedDisplayName != nil {
		toSerialize["updated_display_name"] = o.UpdatedDisplayName
	}
	if o.UpdatedFieldId != nil {
		toSerialize["updated_field_id"] = o.UpdatedFieldId
	}
	if o.UpdatedGroups != nil {
		toSerialize["updated_groups"] = o.UpdatedGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateConnectionRequestDataAttributesFieldsToUpdateItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FieldId            *string  `json:"field_id"`
		UpdatedDescription *string  `json:"updated_description,omitempty"`
		UpdatedDisplayName *string  `json:"updated_display_name,omitempty"`
		UpdatedFieldId     *string  `json:"updated_field_id,omitempty"`
		UpdatedGroups      []string `json:"updated_groups,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FieldId == nil {
		return fmt.Errorf("required field field_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field_id", "updated_description", "updated_display_name", "updated_field_id", "updated_groups"})
	} else {
		return err
	}
	o.FieldId = *all.FieldId
	o.UpdatedDescription = all.UpdatedDescription
	o.UpdatedDisplayName = all.UpdatedDisplayName
	o.UpdatedFieldId = all.UpdatedFieldId
	o.UpdatedGroups = all.UpdatedGroups

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
