// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateDegradationRequestDataAttributes The supported attributes for creating a degradation.
type CreateDegradationRequestDataAttributes struct {
	// The components affected by the degradation.
	ComponentsAffected []CreateDegradationRequestDataAttributesComponentsAffectedItems `json:"components_affected"`
	// The description of the degradation.
	Description *string `json:"description,omitempty"`
	// The status of the degradation.
	Status CreateDegradationRequestDataAttributesStatus `json:"status"`
	// The title of the degradation.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateDegradationRequestDataAttributes instantiates a new CreateDegradationRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateDegradationRequestDataAttributes(componentsAffected []CreateDegradationRequestDataAttributesComponentsAffectedItems, status CreateDegradationRequestDataAttributesStatus, title string) *CreateDegradationRequestDataAttributes {
	this := CreateDegradationRequestDataAttributes{}
	this.ComponentsAffected = componentsAffected
	this.Status = status
	this.Title = title
	return &this
}

// NewCreateDegradationRequestDataAttributesWithDefaults instantiates a new CreateDegradationRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateDegradationRequestDataAttributesWithDefaults() *CreateDegradationRequestDataAttributes {
	this := CreateDegradationRequestDataAttributes{}
	return &this
}

// GetComponentsAffected returns the ComponentsAffected field value.
func (o *CreateDegradationRequestDataAttributes) GetComponentsAffected() []CreateDegradationRequestDataAttributesComponentsAffectedItems {
	if o == nil {
		var ret []CreateDegradationRequestDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value
// and a boolean to check if the value has been set.
func (o *CreateDegradationRequestDataAttributes) GetComponentsAffectedOk() (*[]CreateDegradationRequestDataAttributesComponentsAffectedItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// SetComponentsAffected sets field value.
func (o *CreateDegradationRequestDataAttributes) SetComponentsAffected(v []CreateDegradationRequestDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDegradationRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDegradationRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDegradationRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDegradationRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value.
func (o *CreateDegradationRequestDataAttributes) GetStatus() CreateDegradationRequestDataAttributesStatus {
	if o == nil {
		var ret CreateDegradationRequestDataAttributesStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateDegradationRequestDataAttributes) GetStatusOk() (*CreateDegradationRequestDataAttributesStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CreateDegradationRequestDataAttributes) SetStatus(v CreateDegradationRequestDataAttributesStatus) {
	o.Status = v
}

// GetTitle returns the Title field value.
func (o *CreateDegradationRequestDataAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateDegradationRequestDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *CreateDegradationRequestDataAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateDegradationRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["components_affected"] = o.ComponentsAffected
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateDegradationRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentsAffected *[]CreateDegradationRequestDataAttributesComponentsAffectedItems `json:"components_affected"`
		Description        *string                                                          `json:"description,omitempty"`
		Status             *CreateDegradationRequestDataAttributesStatus                    `json:"status"`
		Title              *string                                                          `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ComponentsAffected == nil {
		return fmt.Errorf("required field components_affected missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components_affected", "description", "status", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComponentsAffected = *all.ComponentsAffected
	o.Description = all.Description
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
