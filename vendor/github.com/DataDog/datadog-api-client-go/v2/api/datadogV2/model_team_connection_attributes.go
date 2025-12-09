// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamConnectionAttributes Attributes of the team connection.
type TeamConnectionAttributes struct {
	// The entity that manages this team connection.
	ManagedBy *string `json:"managed_by,omitempty"`
	// The name of the external source.
	Source *string `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamConnectionAttributes instantiates a new TeamConnectionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamConnectionAttributes() *TeamConnectionAttributes {
	this := TeamConnectionAttributes{}
	return &this
}

// NewTeamConnectionAttributesWithDefaults instantiates a new TeamConnectionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamConnectionAttributesWithDefaults() *TeamConnectionAttributes {
	this := TeamConnectionAttributes{}
	return &this
}

// GetManagedBy returns the ManagedBy field value if set, zero value otherwise.
func (o *TeamConnectionAttributes) GetManagedBy() string {
	if o == nil || o.ManagedBy == nil {
		var ret string
		return ret
	}
	return *o.ManagedBy
}

// GetManagedByOk returns a tuple with the ManagedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConnectionAttributes) GetManagedByOk() (*string, bool) {
	if o == nil || o.ManagedBy == nil {
		return nil, false
	}
	return o.ManagedBy, true
}

// HasManagedBy returns a boolean if a field has been set.
func (o *TeamConnectionAttributes) HasManagedBy() bool {
	return o != nil && o.ManagedBy != nil
}

// SetManagedBy gets a reference to the given string and assigns it to the ManagedBy field.
func (o *TeamConnectionAttributes) SetManagedBy(v string) {
	o.ManagedBy = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TeamConnectionAttributes) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamConnectionAttributes) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TeamConnectionAttributes) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *TeamConnectionAttributes) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamConnectionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ManagedBy != nil {
		toSerialize["managed_by"] = o.ManagedBy
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamConnectionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ManagedBy *string `json:"managed_by,omitempty"`
		Source    *string `json:"source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"managed_by", "source"})
	} else {
		return err
	}
	o.ManagedBy = all.ManagedBy
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
