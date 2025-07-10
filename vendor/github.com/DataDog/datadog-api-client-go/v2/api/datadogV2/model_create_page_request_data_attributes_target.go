// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreatePageRequestDataAttributesTarget Information about the target to notify (such as a team or user).
type CreatePageRequestDataAttributesTarget struct {
	// Identifier for the target (for example, team handle or user ID).
	Identifier *string `json:"identifier,omitempty"`
	// The kind of target, `team_id` | `team_handle` | `user_id`.
	Type *OnCallPageTargetType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreatePageRequestDataAttributesTarget instantiates a new CreatePageRequestDataAttributesTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreatePageRequestDataAttributesTarget() *CreatePageRequestDataAttributesTarget {
	this := CreatePageRequestDataAttributesTarget{}
	return &this
}

// NewCreatePageRequestDataAttributesTargetWithDefaults instantiates a new CreatePageRequestDataAttributesTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreatePageRequestDataAttributesTargetWithDefaults() *CreatePageRequestDataAttributesTarget {
	this := CreatePageRequestDataAttributesTarget{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *CreatePageRequestDataAttributesTarget) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePageRequestDataAttributesTarget) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *CreatePageRequestDataAttributesTarget) HasIdentifier() bool {
	return o != nil && o.Identifier != nil
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *CreatePageRequestDataAttributesTarget) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreatePageRequestDataAttributesTarget) GetType() OnCallPageTargetType {
	if o == nil || o.Type == nil {
		var ret OnCallPageTargetType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePageRequestDataAttributesTarget) GetTypeOk() (*OnCallPageTargetType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreatePageRequestDataAttributesTarget) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given OnCallPageTargetType and assigns it to the Type field.
func (o *CreatePageRequestDataAttributesTarget) SetType(v OnCallPageTargetType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreatePageRequestDataAttributesTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
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
func (o *CreatePageRequestDataAttributesTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Identifier *string               `json:"identifier,omitempty"`
		Type       *OnCallPageTargetType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"identifier", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Identifier = all.Identifier
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
