// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchDegradationRequestDataAttributes The supported attributes for updating a degradation.
type PatchDegradationRequestDataAttributes struct {
	// The components affected by the degradation.
	ComponentsAffected []PatchDegradationRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
	// The description of the degradation.
	Description *string `json:"description,omitempty"`
	// The status of the degradation.
	Status *PatchDegradationRequestDataAttributesStatus `json:"status,omitempty"`
	// The title of the degradation.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchDegradationRequestDataAttributes instantiates a new PatchDegradationRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchDegradationRequestDataAttributes() *PatchDegradationRequestDataAttributes {
	this := PatchDegradationRequestDataAttributes{}
	return &this
}

// NewPatchDegradationRequestDataAttributesWithDefaults instantiates a new PatchDegradationRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchDegradationRequestDataAttributesWithDefaults() *PatchDegradationRequestDataAttributes {
	this := PatchDegradationRequestDataAttributes{}
	return &this
}

// GetComponentsAffected returns the ComponentsAffected field value if set, zero value otherwise.
func (o *PatchDegradationRequestDataAttributes) GetComponentsAffected() []PatchDegradationRequestDataAttributesComponentsAffectedItems {
	if o == nil || o.ComponentsAffected == nil {
		var ret []PatchDegradationRequestDataAttributesComponentsAffectedItems
		return ret
	}
	return o.ComponentsAffected
}

// GetComponentsAffectedOk returns a tuple with the ComponentsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationRequestDataAttributes) GetComponentsAffectedOk() (*[]PatchDegradationRequestDataAttributesComponentsAffectedItems, bool) {
	if o == nil || o.ComponentsAffected == nil {
		return nil, false
	}
	return &o.ComponentsAffected, true
}

// HasComponentsAffected returns a boolean if a field has been set.
func (o *PatchDegradationRequestDataAttributes) HasComponentsAffected() bool {
	return o != nil && o.ComponentsAffected != nil
}

// SetComponentsAffected gets a reference to the given []PatchDegradationRequestDataAttributesComponentsAffectedItems and assigns it to the ComponentsAffected field.
func (o *PatchDegradationRequestDataAttributes) SetComponentsAffected(v []PatchDegradationRequestDataAttributesComponentsAffectedItems) {
	o.ComponentsAffected = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchDegradationRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchDegradationRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchDegradationRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchDegradationRequestDataAttributes) GetStatus() PatchDegradationRequestDataAttributesStatus {
	if o == nil || o.Status == nil {
		var ret PatchDegradationRequestDataAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationRequestDataAttributes) GetStatusOk() (*PatchDegradationRequestDataAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchDegradationRequestDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given PatchDegradationRequestDataAttributesStatus and assigns it to the Status field.
func (o *PatchDegradationRequestDataAttributes) SetStatus(v PatchDegradationRequestDataAttributesStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PatchDegradationRequestDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationRequestDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PatchDegradationRequestDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PatchDegradationRequestDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchDegradationRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComponentsAffected != nil {
		toSerialize["components_affected"] = o.ComponentsAffected
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchDegradationRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentsAffected []PatchDegradationRequestDataAttributesComponentsAffectedItems `json:"components_affected,omitempty"`
		Description        *string                                                        `json:"description,omitempty"`
		Status             *PatchDegradationRequestDataAttributesStatus                   `json:"status,omitempty"`
		Title              *string                                                        `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components_affected", "description", "status", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComponentsAffected = all.ComponentsAffected
	o.Description = all.Description
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
