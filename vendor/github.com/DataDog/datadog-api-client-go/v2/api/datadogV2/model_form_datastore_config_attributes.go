// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormDatastoreConfigAttributes The datastore configuration for a form.
type FormDatastoreConfigAttributes struct {
	// The ID of the datastore.
	DatastoreId uuid.UUID `json:"datastore_id"`
	// The name of the primary column in the datastore.
	PrimaryColumnName string `json:"primary_column_name"`
	// The strategy used to generate primary keys in the datastore.
	PrimaryKeyGenerationStrategy string `json:"primary_key_generation_strategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormDatastoreConfigAttributes instantiates a new FormDatastoreConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormDatastoreConfigAttributes(datastoreId uuid.UUID, primaryColumnName string, primaryKeyGenerationStrategy string) *FormDatastoreConfigAttributes {
	this := FormDatastoreConfigAttributes{}
	this.DatastoreId = datastoreId
	this.PrimaryColumnName = primaryColumnName
	this.PrimaryKeyGenerationStrategy = primaryKeyGenerationStrategy
	return &this
}

// NewFormDatastoreConfigAttributesWithDefaults instantiates a new FormDatastoreConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormDatastoreConfigAttributesWithDefaults() *FormDatastoreConfigAttributes {
	this := FormDatastoreConfigAttributes{}
	return &this
}

// GetDatastoreId returns the DatastoreId field value.
func (o *FormDatastoreConfigAttributes) GetDatastoreId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.DatastoreId
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value
// and a boolean to check if the value has been set.
func (o *FormDatastoreConfigAttributes) GetDatastoreIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatastoreId, true
}

// SetDatastoreId sets field value.
func (o *FormDatastoreConfigAttributes) SetDatastoreId(v uuid.UUID) {
	o.DatastoreId = v
}

// GetPrimaryColumnName returns the PrimaryColumnName field value.
func (o *FormDatastoreConfigAttributes) GetPrimaryColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrimaryColumnName
}

// GetPrimaryColumnNameOk returns a tuple with the PrimaryColumnName field value
// and a boolean to check if the value has been set.
func (o *FormDatastoreConfigAttributes) GetPrimaryColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryColumnName, true
}

// SetPrimaryColumnName sets field value.
func (o *FormDatastoreConfigAttributes) SetPrimaryColumnName(v string) {
	o.PrimaryColumnName = v
}

// GetPrimaryKeyGenerationStrategy returns the PrimaryKeyGenerationStrategy field value.
func (o *FormDatastoreConfigAttributes) GetPrimaryKeyGenerationStrategy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrimaryKeyGenerationStrategy
}

// GetPrimaryKeyGenerationStrategyOk returns a tuple with the PrimaryKeyGenerationStrategy field value
// and a boolean to check if the value has been set.
func (o *FormDatastoreConfigAttributes) GetPrimaryKeyGenerationStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryKeyGenerationStrategy, true
}

// SetPrimaryKeyGenerationStrategy sets field value.
func (o *FormDatastoreConfigAttributes) SetPrimaryKeyGenerationStrategy(v string) {
	o.PrimaryKeyGenerationStrategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormDatastoreConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["datastore_id"] = o.DatastoreId
	toSerialize["primary_column_name"] = o.PrimaryColumnName
	toSerialize["primary_key_generation_strategy"] = o.PrimaryKeyGenerationStrategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormDatastoreConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatastoreId                  *uuid.UUID `json:"datastore_id"`
		PrimaryColumnName            *string    `json:"primary_column_name"`
		PrimaryKeyGenerationStrategy *string    `json:"primary_key_generation_strategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatastoreId == nil {
		return fmt.Errorf("required field datastore_id missing")
	}
	if all.PrimaryColumnName == nil {
		return fmt.Errorf("required field primary_column_name missing")
	}
	if all.PrimaryKeyGenerationStrategy == nil {
		return fmt.Errorf("required field primary_key_generation_strategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"datastore_id", "primary_column_name", "primary_key_generation_strategy"})
	} else {
		return err
	}
	o.DatastoreId = *all.DatastoreId
	o.PrimaryColumnName = *all.PrimaryColumnName
	o.PrimaryKeyGenerationStrategy = *all.PrimaryKeyGenerationStrategy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
