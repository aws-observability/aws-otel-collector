// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems
type ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems struct {
	//
	Ref *string `json:"ref,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesVulnerabilitiesItemsAffectsItems instantiates a new ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesVulnerabilitiesItemsAffectsItems() *ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems {
	this := ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems{}
	return &this
}

// NewScaRequestDataAttributesVulnerabilitiesItemsAffectsItemsWithDefaults instantiates a new ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesVulnerabilitiesItemsAffectsItemsWithDefaults() *ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems {
	this := ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems) GetRefOk() (*string, bool) {
	if o == nil || o.Ref == nil {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems) HasRef() bool {
	return o != nil && o.Ref != nil
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems) SetRef(v string) {
	o.Ref = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ref != nil {
		toSerialize["ref"] = o.Ref
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesVulnerabilitiesItemsAffectsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ref *string `json:"ref,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ref"})
	} else {
		return err
	}
	o.Ref = all.Ref

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
