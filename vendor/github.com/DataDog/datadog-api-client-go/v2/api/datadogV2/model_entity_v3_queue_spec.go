// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3QueueSpec The definition of Entity V3 Queue Spec object.
type EntityV3QueueSpec struct {
	// The lifecycle state of the queue.
	Lifecycle *string `json:"lifecycle,omitempty"`
	// The importance of the queue.
	Tier *string `json:"tier,omitempty"`
	// The type of queue.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3QueueSpec instantiates a new EntityV3QueueSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3QueueSpec() *EntityV3QueueSpec {
	this := EntityV3QueueSpec{}
	return &this
}

// NewEntityV3QueueSpecWithDefaults instantiates a new EntityV3QueueSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3QueueSpecWithDefaults() *EntityV3QueueSpec {
	this := EntityV3QueueSpec{}
	return &this
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EntityV3QueueSpec) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3QueueSpec) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EntityV3QueueSpec) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EntityV3QueueSpec) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *EntityV3QueueSpec) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3QueueSpec) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *EntityV3QueueSpec) HasTier() bool {
	return o != nil && o.Tier != nil
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *EntityV3QueueSpec) SetTier(v string) {
	o.Tier = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntityV3QueueSpec) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3QueueSpec) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntityV3QueueSpec) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntityV3QueueSpec) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3QueueSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *EntityV3QueueSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Lifecycle *string `json:"lifecycle,omitempty"`
		Tier      *string `json:"tier,omitempty"`
		Type      *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Lifecycle = all.Lifecycle
	o.Tier = all.Tier
	o.Type = all.Type

	return nil
}
