// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3DatastoreSpec The definition of Entity V3 Datastore Spec object.
type EntityV3DatastoreSpec struct {
	// A list of components the datastore is a part of
	ComponentOf []string `json:"componentOf,omitempty"`
	// The lifecycle state of the datastore.
	Lifecycle *string `json:"lifecycle,omitempty"`
	// The importance of the datastore.
	Tier *string `json:"tier,omitempty"`
	// The type of datastore.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3DatastoreSpec instantiates a new EntityV3DatastoreSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3DatastoreSpec() *EntityV3DatastoreSpec {
	this := EntityV3DatastoreSpec{}
	return &this
}

// NewEntityV3DatastoreSpecWithDefaults instantiates a new EntityV3DatastoreSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3DatastoreSpecWithDefaults() *EntityV3DatastoreSpec {
	this := EntityV3DatastoreSpec{}
	return &this
}

// GetComponentOf returns the ComponentOf field value if set, zero value otherwise.
func (o *EntityV3DatastoreSpec) GetComponentOf() []string {
	if o == nil || o.ComponentOf == nil {
		var ret []string
		return ret
	}
	return o.ComponentOf
}

// GetComponentOfOk returns a tuple with the ComponentOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3DatastoreSpec) GetComponentOfOk() (*[]string, bool) {
	if o == nil || o.ComponentOf == nil {
		return nil, false
	}
	return &o.ComponentOf, true
}

// HasComponentOf returns a boolean if a field has been set.
func (o *EntityV3DatastoreSpec) HasComponentOf() bool {
	return o != nil && o.ComponentOf != nil
}

// SetComponentOf gets a reference to the given []string and assigns it to the ComponentOf field.
func (o *EntityV3DatastoreSpec) SetComponentOf(v []string) {
	o.ComponentOf = v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EntityV3DatastoreSpec) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3DatastoreSpec) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EntityV3DatastoreSpec) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EntityV3DatastoreSpec) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *EntityV3DatastoreSpec) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3DatastoreSpec) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *EntityV3DatastoreSpec) HasTier() bool {
	return o != nil && o.Tier != nil
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *EntityV3DatastoreSpec) SetTier(v string) {
	o.Tier = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntityV3DatastoreSpec) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3DatastoreSpec) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntityV3DatastoreSpec) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntityV3DatastoreSpec) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3DatastoreSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComponentOf != nil {
		toSerialize["componentOf"] = o.ComponentOf
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
func (o *EntityV3DatastoreSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentOf []string `json:"componentOf,omitempty"`
		Lifecycle   *string  `json:"lifecycle,omitempty"`
		Tier        *string  `json:"tier,omitempty"`
		Type        *string  `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ComponentOf = all.ComponentOf
	o.Lifecycle = all.Lifecycle
	o.Tier = all.Tier
	o.Type = all.Type

	return nil
}
