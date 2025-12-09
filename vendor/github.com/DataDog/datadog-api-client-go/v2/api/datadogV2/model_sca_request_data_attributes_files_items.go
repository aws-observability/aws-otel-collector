// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesFilesItems
type ScaRequestDataAttributesFilesItems struct {
	//
	Name *string `json:"name,omitempty"`
	//
	Purl *string `json:"purl,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesFilesItems instantiates a new ScaRequestDataAttributesFilesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesFilesItems() *ScaRequestDataAttributesFilesItems {
	this := ScaRequestDataAttributesFilesItems{}
	return &this
}

// NewScaRequestDataAttributesFilesItemsWithDefaults instantiates a new ScaRequestDataAttributesFilesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesFilesItemsWithDefaults() *ScaRequestDataAttributesFilesItems {
	this := ScaRequestDataAttributesFilesItems{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesFilesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesFilesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesFilesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScaRequestDataAttributesFilesItems) SetName(v string) {
	o.Name = &v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesFilesItems) GetPurl() string {
	if o == nil || o.Purl == nil {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesFilesItems) GetPurlOk() (*string, bool) {
	if o == nil || o.Purl == nil {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesFilesItems) HasPurl() bool {
	return o != nil && o.Purl != nil
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *ScaRequestDataAttributesFilesItems) SetPurl(v string) {
	o.Purl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesFilesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Purl != nil {
		toSerialize["purl"] = o.Purl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesFilesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name,omitempty"`
		Purl *string `json:"purl,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "purl"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Purl = all.Purl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
