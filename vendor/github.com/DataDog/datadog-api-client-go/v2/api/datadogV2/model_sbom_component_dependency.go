// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMComponentDependency The dependencies of a component of the SBOM.
type SBOMComponentDependency struct {
	// The components that are dependencies of the ref component.
	DependsOn []string `json:"dependsOn,omitempty"`
	// The identifier for the related component.
	Ref *string `json:"ref,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSBOMComponentDependency instantiates a new SBOMComponentDependency object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSBOMComponentDependency() *SBOMComponentDependency {
	this := SBOMComponentDependency{}
	return &this
}

// NewSBOMComponentDependencyWithDefaults instantiates a new SBOMComponentDependency object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSBOMComponentDependencyWithDefaults() *SBOMComponentDependency {
	this := SBOMComponentDependency{}
	return &this
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise.
func (o *SBOMComponentDependency) GetDependsOn() []string {
	if o == nil || o.DependsOn == nil {
		var ret []string
		return ret
	}
	return o.DependsOn
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBOMComponentDependency) GetDependsOnOk() (*[]string, bool) {
	if o == nil || o.DependsOn == nil {
		return nil, false
	}
	return &o.DependsOn, true
}

// HasDependsOn returns a boolean if a field has been set.
func (o *SBOMComponentDependency) HasDependsOn() bool {
	return o != nil && o.DependsOn != nil
}

// SetDependsOn gets a reference to the given []string and assigns it to the DependsOn field.
func (o *SBOMComponentDependency) SetDependsOn(v []string) {
	o.DependsOn = v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *SBOMComponentDependency) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBOMComponentDependency) GetRefOk() (*string, bool) {
	if o == nil || o.Ref == nil {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *SBOMComponentDependency) HasRef() bool {
	return o != nil && o.Ref != nil
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *SBOMComponentDependency) SetRef(v string) {
	o.Ref = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SBOMComponentDependency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DependsOn != nil {
		toSerialize["dependsOn"] = o.DependsOn
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
func (o *SBOMComponentDependency) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DependsOn []string `json:"dependsOn,omitempty"`
		Ref       *string  `json:"ref,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dependsOn", "ref"})
	} else {
		return err
	}
	o.DependsOn = all.DependsOn
	o.Ref = all.Ref

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
