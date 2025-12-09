// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamSyncAttributes Team sync attributes.
type TeamSyncAttributes struct {
	// The external source platform for team synchronization. Only "github" is supported.
	Source TeamSyncAttributesSource `json:"source"`
	// The type of synchronization operation. Only "link" is supported, which links existing teams by matching names.
	Type TeamSyncAttributesType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamSyncAttributes instantiates a new TeamSyncAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamSyncAttributes(source TeamSyncAttributesSource, typeVar TeamSyncAttributesType) *TeamSyncAttributes {
	this := TeamSyncAttributes{}
	this.Source = source
	this.Type = typeVar
	return &this
}

// NewTeamSyncAttributesWithDefaults instantiates a new TeamSyncAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamSyncAttributesWithDefaults() *TeamSyncAttributes {
	this := TeamSyncAttributes{}
	return &this
}

// GetSource returns the Source field value.
func (o *TeamSyncAttributes) GetSource() TeamSyncAttributesSource {
	if o == nil {
		var ret TeamSyncAttributesSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *TeamSyncAttributes) GetSourceOk() (*TeamSyncAttributesSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *TeamSyncAttributes) SetSource(v TeamSyncAttributesSource) {
	o.Source = v
}

// GetType returns the Type field value.
func (o *TeamSyncAttributes) GetType() TeamSyncAttributesType {
	if o == nil {
		var ret TeamSyncAttributesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TeamSyncAttributes) GetTypeOk() (*TeamSyncAttributesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TeamSyncAttributes) SetType(v TeamSyncAttributesType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamSyncAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["source"] = o.Source
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamSyncAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source *TeamSyncAttributesSource `json:"source"`
		Type   *TeamSyncAttributesType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"source", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
