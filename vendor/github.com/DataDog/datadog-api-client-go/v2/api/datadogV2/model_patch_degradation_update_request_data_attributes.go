// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchDegradationUpdateRequestDataAttributes Attributes for editing a degradation update.
type PatchDegradationUpdateRequestDataAttributes struct {
	// The message body of the update.
	Description *string `json:"description,omitempty"`
	// The status of the degradation update.
	Status *PatchDegradationUpdateRequestDataAttributesStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchDegradationUpdateRequestDataAttributes instantiates a new PatchDegradationUpdateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchDegradationUpdateRequestDataAttributes() *PatchDegradationUpdateRequestDataAttributes {
	this := PatchDegradationUpdateRequestDataAttributes{}
	return &this
}

// NewPatchDegradationUpdateRequestDataAttributesWithDefaults instantiates a new PatchDegradationUpdateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchDegradationUpdateRequestDataAttributesWithDefaults() *PatchDegradationUpdateRequestDataAttributes {
	this := PatchDegradationUpdateRequestDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchDegradationUpdateRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationUpdateRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchDegradationUpdateRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchDegradationUpdateRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchDegradationUpdateRequestDataAttributes) GetStatus() PatchDegradationUpdateRequestDataAttributesStatus {
	if o == nil || o.Status == nil {
		var ret PatchDegradationUpdateRequestDataAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchDegradationUpdateRequestDataAttributes) GetStatusOk() (*PatchDegradationUpdateRequestDataAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchDegradationUpdateRequestDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given PatchDegradationUpdateRequestDataAttributesStatus and assigns it to the Status field.
func (o *PatchDegradationUpdateRequestDataAttributes) SetStatus(v PatchDegradationUpdateRequestDataAttributesStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchDegradationUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchDegradationUpdateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                            `json:"description,omitempty"`
		Status      *PatchDegradationUpdateRequestDataAttributesStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
