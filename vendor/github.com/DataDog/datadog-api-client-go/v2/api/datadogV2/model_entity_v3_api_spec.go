// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3APISpec The definition of Entity V3 API Spec object.
type EntityV3APISpec struct {
	// Services which implemented the API.
	ImplementedBy []string `json:"implementedBy,omitempty"`
	// The API definition.
	Interface *EntityV3APISpecInterface `json:"interface,omitempty"`
	// The lifecycle state of the component.
	Lifecycle *string `json:"lifecycle,omitempty"`
	// The importance of the component.
	Tier *string `json:"tier,omitempty"`
	// The type of API.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3APISpec instantiates a new EntityV3APISpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3APISpec() *EntityV3APISpec {
	this := EntityV3APISpec{}
	return &this
}

// NewEntityV3APISpecWithDefaults instantiates a new EntityV3APISpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3APISpecWithDefaults() *EntityV3APISpec {
	this := EntityV3APISpec{}
	return &this
}

// GetImplementedBy returns the ImplementedBy field value if set, zero value otherwise.
func (o *EntityV3APISpec) GetImplementedBy() []string {
	if o == nil || o.ImplementedBy == nil {
		var ret []string
		return ret
	}
	return o.ImplementedBy
}

// GetImplementedByOk returns a tuple with the ImplementedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3APISpec) GetImplementedByOk() (*[]string, bool) {
	if o == nil || o.ImplementedBy == nil {
		return nil, false
	}
	return &o.ImplementedBy, true
}

// HasImplementedBy returns a boolean if a field has been set.
func (o *EntityV3APISpec) HasImplementedBy() bool {
	return o != nil && o.ImplementedBy != nil
}

// SetImplementedBy gets a reference to the given []string and assigns it to the ImplementedBy field.
func (o *EntityV3APISpec) SetImplementedBy(v []string) {
	o.ImplementedBy = v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *EntityV3APISpec) GetInterface() EntityV3APISpecInterface {
	if o == nil || o.Interface == nil {
		var ret EntityV3APISpecInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3APISpec) GetInterfaceOk() (*EntityV3APISpecInterface, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *EntityV3APISpec) HasInterface() bool {
	return o != nil && o.Interface != nil
}

// SetInterface gets a reference to the given EntityV3APISpecInterface and assigns it to the Interface field.
func (o *EntityV3APISpec) SetInterface(v EntityV3APISpecInterface) {
	o.Interface = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EntityV3APISpec) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3APISpec) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EntityV3APISpec) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EntityV3APISpec) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *EntityV3APISpec) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3APISpec) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *EntityV3APISpec) HasTier() bool {
	return o != nil && o.Tier != nil
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *EntityV3APISpec) SetTier(v string) {
	o.Tier = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntityV3APISpec) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3APISpec) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntityV3APISpec) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntityV3APISpec) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3APISpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ImplementedBy != nil {
		toSerialize["implementedBy"] = o.ImplementedBy
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
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
func (o *EntityV3APISpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ImplementedBy []string                  `json:"implementedBy,omitempty"`
		Interface     *EntityV3APISpecInterface `json:"interface,omitempty"`
		Lifecycle     *string                   `json:"lifecycle,omitempty"`
		Tier          *string                   `json:"tier,omitempty"`
		Type          *string                   `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ImplementedBy = all.ImplementedBy
	o.Interface = all.Interface
	o.Lifecycle = all.Lifecycle
	o.Tier = all.Tier
	o.Type = all.Type

	return nil
}
