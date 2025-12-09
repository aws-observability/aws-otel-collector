// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateConnectionRequestDataAttributes
type UpdateConnectionRequestDataAttributes struct {
	//
	FieldsToAdd []CreateConnectionRequestDataAttributesFieldsItems `json:"fields_to_add,omitempty"`
	//
	FieldsToDelete []string `json:"fields_to_delete,omitempty"`
	//
	FieldsToUpdate []UpdateConnectionRequestDataAttributesFieldsToUpdateItems `json:"fields_to_update,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateConnectionRequestDataAttributes instantiates a new UpdateConnectionRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateConnectionRequestDataAttributes() *UpdateConnectionRequestDataAttributes {
	this := UpdateConnectionRequestDataAttributes{}
	return &this
}

// NewUpdateConnectionRequestDataAttributesWithDefaults instantiates a new UpdateConnectionRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateConnectionRequestDataAttributesWithDefaults() *UpdateConnectionRequestDataAttributes {
	this := UpdateConnectionRequestDataAttributes{}
	return &this
}

// GetFieldsToAdd returns the FieldsToAdd field value if set, zero value otherwise.
func (o *UpdateConnectionRequestDataAttributes) GetFieldsToAdd() []CreateConnectionRequestDataAttributesFieldsItems {
	if o == nil || o.FieldsToAdd == nil {
		var ret []CreateConnectionRequestDataAttributesFieldsItems
		return ret
	}
	return o.FieldsToAdd
}

// GetFieldsToAddOk returns a tuple with the FieldsToAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectionRequestDataAttributes) GetFieldsToAddOk() (*[]CreateConnectionRequestDataAttributesFieldsItems, bool) {
	if o == nil || o.FieldsToAdd == nil {
		return nil, false
	}
	return &o.FieldsToAdd, true
}

// HasFieldsToAdd returns a boolean if a field has been set.
func (o *UpdateConnectionRequestDataAttributes) HasFieldsToAdd() bool {
	return o != nil && o.FieldsToAdd != nil
}

// SetFieldsToAdd gets a reference to the given []CreateConnectionRequestDataAttributesFieldsItems and assigns it to the FieldsToAdd field.
func (o *UpdateConnectionRequestDataAttributes) SetFieldsToAdd(v []CreateConnectionRequestDataAttributesFieldsItems) {
	o.FieldsToAdd = v
}

// GetFieldsToDelete returns the FieldsToDelete field value if set, zero value otherwise.
func (o *UpdateConnectionRequestDataAttributes) GetFieldsToDelete() []string {
	if o == nil || o.FieldsToDelete == nil {
		var ret []string
		return ret
	}
	return o.FieldsToDelete
}

// GetFieldsToDeleteOk returns a tuple with the FieldsToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectionRequestDataAttributes) GetFieldsToDeleteOk() (*[]string, bool) {
	if o == nil || o.FieldsToDelete == nil {
		return nil, false
	}
	return &o.FieldsToDelete, true
}

// HasFieldsToDelete returns a boolean if a field has been set.
func (o *UpdateConnectionRequestDataAttributes) HasFieldsToDelete() bool {
	return o != nil && o.FieldsToDelete != nil
}

// SetFieldsToDelete gets a reference to the given []string and assigns it to the FieldsToDelete field.
func (o *UpdateConnectionRequestDataAttributes) SetFieldsToDelete(v []string) {
	o.FieldsToDelete = v
}

// GetFieldsToUpdate returns the FieldsToUpdate field value if set, zero value otherwise.
func (o *UpdateConnectionRequestDataAttributes) GetFieldsToUpdate() []UpdateConnectionRequestDataAttributesFieldsToUpdateItems {
	if o == nil || o.FieldsToUpdate == nil {
		var ret []UpdateConnectionRequestDataAttributesFieldsToUpdateItems
		return ret
	}
	return o.FieldsToUpdate
}

// GetFieldsToUpdateOk returns a tuple with the FieldsToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectionRequestDataAttributes) GetFieldsToUpdateOk() (*[]UpdateConnectionRequestDataAttributesFieldsToUpdateItems, bool) {
	if o == nil || o.FieldsToUpdate == nil {
		return nil, false
	}
	return &o.FieldsToUpdate, true
}

// HasFieldsToUpdate returns a boolean if a field has been set.
func (o *UpdateConnectionRequestDataAttributes) HasFieldsToUpdate() bool {
	return o != nil && o.FieldsToUpdate != nil
}

// SetFieldsToUpdate gets a reference to the given []UpdateConnectionRequestDataAttributesFieldsToUpdateItems and assigns it to the FieldsToUpdate field.
func (o *UpdateConnectionRequestDataAttributes) SetFieldsToUpdate(v []UpdateConnectionRequestDataAttributesFieldsToUpdateItems) {
	o.FieldsToUpdate = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateConnectionRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FieldsToAdd != nil {
		toSerialize["fields_to_add"] = o.FieldsToAdd
	}
	if o.FieldsToDelete != nil {
		toSerialize["fields_to_delete"] = o.FieldsToDelete
	}
	if o.FieldsToUpdate != nil {
		toSerialize["fields_to_update"] = o.FieldsToUpdate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateConnectionRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FieldsToAdd    []CreateConnectionRequestDataAttributesFieldsItems         `json:"fields_to_add,omitempty"`
		FieldsToDelete []string                                                   `json:"fields_to_delete,omitempty"`
		FieldsToUpdate []UpdateConnectionRequestDataAttributesFieldsToUpdateItems `json:"fields_to_update,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields_to_add", "fields_to_delete", "fields_to_update"})
	} else {
		return err
	}
	o.FieldsToAdd = all.FieldsToAdd
	o.FieldsToDelete = all.FieldsToDelete
	o.FieldsToUpdate = all.FieldsToUpdate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
