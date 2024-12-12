// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3SystemSpec The definition of Entity V3 System Spec object.
type EntityV3SystemSpec struct {
	// A list of components belongs to the system.
	Components []string `json:"components,omitempty"`
	// The lifecycle state of the component.
	Lifecycle *string `json:"lifecycle,omitempty"`
	// An entity reference to the owner of the component.
	Tier *string `json:"tier,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3SystemSpec instantiates a new EntityV3SystemSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3SystemSpec() *EntityV3SystemSpec {
	this := EntityV3SystemSpec{}
	return &this
}

// NewEntityV3SystemSpecWithDefaults instantiates a new EntityV3SystemSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3SystemSpecWithDefaults() *EntityV3SystemSpec {
	this := EntityV3SystemSpec{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *EntityV3SystemSpec) GetComponents() []string {
	if o == nil || o.Components == nil {
		var ret []string
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3SystemSpec) GetComponentsOk() (*[]string, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *EntityV3SystemSpec) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []string and assigns it to the Components field.
func (o *EntityV3SystemSpec) SetComponents(v []string) {
	o.Components = v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EntityV3SystemSpec) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3SystemSpec) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EntityV3SystemSpec) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EntityV3SystemSpec) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *EntityV3SystemSpec) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3SystemSpec) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *EntityV3SystemSpec) HasTier() bool {
	return o != nil && o.Tier != nil
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *EntityV3SystemSpec) SetTier(v string) {
	o.Tier = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3SystemSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if o.Tier != nil {
		toSerialize["tier"] = o.Tier
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3SystemSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components []string `json:"components,omitempty"`
		Lifecycle  *string  `json:"lifecycle,omitempty"`
		Tier       *string  `json:"tier,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Components = all.Components
	o.Lifecycle = all.Lifecycle
	o.Tier = all.Tier

	return nil
}
