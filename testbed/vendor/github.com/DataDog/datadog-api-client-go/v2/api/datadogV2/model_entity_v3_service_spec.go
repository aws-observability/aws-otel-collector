// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3ServiceSpec The definition of Entity V3 Service Spec object.
type EntityV3ServiceSpec struct {
	// A list of components the service is a part of
	ComponentOf []string `json:"componentOf,omitempty"`
	// A list of components the service depends on.
	DependsOn []string `json:"dependsOn,omitempty"`
	// The service's programming language.
	Languages []string `json:"languages,omitempty"`
	// The lifecycle state of the component.
	Lifecycle *string `json:"lifecycle,omitempty"`
	// The importance of the component.
	Tier *string `json:"tier,omitempty"`
	// The type of service.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3ServiceSpec instantiates a new EntityV3ServiceSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3ServiceSpec() *EntityV3ServiceSpec {
	this := EntityV3ServiceSpec{}
	return &this
}

// NewEntityV3ServiceSpecWithDefaults instantiates a new EntityV3ServiceSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3ServiceSpecWithDefaults() *EntityV3ServiceSpec {
	this := EntityV3ServiceSpec{}
	return &this
}

// GetComponentOf returns the ComponentOf field value if set, zero value otherwise.
func (o *EntityV3ServiceSpec) GetComponentOf() []string {
	if o == nil || o.ComponentOf == nil {
		var ret []string
		return ret
	}
	return o.ComponentOf
}

// GetComponentOfOk returns a tuple with the ComponentOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3ServiceSpec) GetComponentOfOk() (*[]string, bool) {
	if o == nil || o.ComponentOf == nil {
		return nil, false
	}
	return &o.ComponentOf, true
}

// HasComponentOf returns a boolean if a field has been set.
func (o *EntityV3ServiceSpec) HasComponentOf() bool {
	return o != nil && o.ComponentOf != nil
}

// SetComponentOf gets a reference to the given []string and assigns it to the ComponentOf field.
func (o *EntityV3ServiceSpec) SetComponentOf(v []string) {
	o.ComponentOf = v
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise.
func (o *EntityV3ServiceSpec) GetDependsOn() []string {
	if o == nil || o.DependsOn == nil {
		var ret []string
		return ret
	}
	return o.DependsOn
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3ServiceSpec) GetDependsOnOk() (*[]string, bool) {
	if o == nil || o.DependsOn == nil {
		return nil, false
	}
	return &o.DependsOn, true
}

// HasDependsOn returns a boolean if a field has been set.
func (o *EntityV3ServiceSpec) HasDependsOn() bool {
	return o != nil && o.DependsOn != nil
}

// SetDependsOn gets a reference to the given []string and assigns it to the DependsOn field.
func (o *EntityV3ServiceSpec) SetDependsOn(v []string) {
	o.DependsOn = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *EntityV3ServiceSpec) GetLanguages() []string {
	if o == nil || o.Languages == nil {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3ServiceSpec) GetLanguagesOk() (*[]string, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return &o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *EntityV3ServiceSpec) HasLanguages() bool {
	return o != nil && o.Languages != nil
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *EntityV3ServiceSpec) SetLanguages(v []string) {
	o.Languages = v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EntityV3ServiceSpec) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3ServiceSpec) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EntityV3ServiceSpec) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EntityV3ServiceSpec) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *EntityV3ServiceSpec) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3ServiceSpec) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *EntityV3ServiceSpec) HasTier() bool {
	return o != nil && o.Tier != nil
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *EntityV3ServiceSpec) SetTier(v string) {
	o.Tier = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntityV3ServiceSpec) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3ServiceSpec) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntityV3ServiceSpec) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntityV3ServiceSpec) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3ServiceSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComponentOf != nil {
		toSerialize["componentOf"] = o.ComponentOf
	}
	if o.DependsOn != nil {
		toSerialize["dependsOn"] = o.DependsOn
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if o.Tier != nil {
		toSerialize["tier"] = o.Tier
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3ServiceSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentOf []string `json:"componentOf,omitempty"`
		DependsOn   []string `json:"dependsOn,omitempty"`
		Languages   []string `json:"languages,omitempty"`
		Lifecycle   *string  `json:"lifecycle,omitempty"`
		Tier        *string  `json:"tier,omitempty"`
		Type        *string  `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ComponentOf = all.ComponentOf
	o.DependsOn = all.DependsOn
	o.Languages = all.Languages
	o.Lifecycle = all.Lifecycle
	o.Tier = all.Tier
	o.Type = all.Type

	return nil
}
