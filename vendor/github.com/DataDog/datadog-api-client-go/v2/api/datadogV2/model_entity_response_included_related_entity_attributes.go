// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedRelatedEntityAttributes Related entity attributes.
type EntityResponseIncludedRelatedEntityAttributes struct {
	// Entity kind.
	Kind *string `json:"kind,omitempty"`
	// Entity name.
	Name *string `json:"name,omitempty"`
	// Entity namespace.
	Namespace *string `json:"namespace,omitempty"`
	// Entity relation type to the associated entity.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseIncludedRelatedEntityAttributes instantiates a new EntityResponseIncludedRelatedEntityAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseIncludedRelatedEntityAttributes() *EntityResponseIncludedRelatedEntityAttributes {
	this := EntityResponseIncludedRelatedEntityAttributes{}
	return &this
}

// NewEntityResponseIncludedRelatedEntityAttributesWithDefaults instantiates a new EntityResponseIncludedRelatedEntityAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseIncludedRelatedEntityAttributesWithDefaults() *EntityResponseIncludedRelatedEntityAttributes {
	this := EntityResponseIncludedRelatedEntityAttributes{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntityAttributes) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntityAttributes) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntityAttributes) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *EntityResponseIncludedRelatedEntityAttributes) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntityAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntityAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntityAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntityResponseIncludedRelatedEntityAttributes) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntityAttributes) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntityAttributes) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntityAttributes) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *EntityResponseIncludedRelatedEntityAttributes) SetNamespace(v string) {
	o.Namespace = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntityAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntityAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntityAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntityResponseIncludedRelatedEntityAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseIncludedRelatedEntityAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityResponseIncludedRelatedEntityAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kind      *string `json:"kind,omitempty"`
		Name      *string `json:"name,omitempty"`
		Namespace *string `json:"namespace,omitempty"`
		Type      *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"kind", "name", "namespace", "type"})
	} else {
		return err
	}
	o.Kind = all.Kind
	o.Name = all.Name
	o.Namespace = all.Namespace
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
