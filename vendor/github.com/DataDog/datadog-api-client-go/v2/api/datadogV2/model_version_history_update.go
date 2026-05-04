// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// VersionHistoryUpdate A change in a rule version.
type VersionHistoryUpdate struct {
	// The new value of the field.
	Change *string `json:"change,omitempty"`
	// The field that was changed.
	Field *string `json:"field,omitempty"`
	// The type of change.
	Type *VersionHistoryUpdateType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVersionHistoryUpdate instantiates a new VersionHistoryUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVersionHistoryUpdate() *VersionHistoryUpdate {
	this := VersionHistoryUpdate{}
	return &this
}

// NewVersionHistoryUpdateWithDefaults instantiates a new VersionHistoryUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVersionHistoryUpdateWithDefaults() *VersionHistoryUpdate {
	this := VersionHistoryUpdate{}
	return &this
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *VersionHistoryUpdate) GetChange() string {
	if o == nil || o.Change == nil {
		var ret string
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionHistoryUpdate) GetChangeOk() (*string, bool) {
	if o == nil || o.Change == nil {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *VersionHistoryUpdate) HasChange() bool {
	return o != nil && o.Change != nil
}

// SetChange gets a reference to the given string and assigns it to the Change field.
func (o *VersionHistoryUpdate) SetChange(v string) {
	o.Change = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *VersionHistoryUpdate) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionHistoryUpdate) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *VersionHistoryUpdate) HasField() bool {
	return o != nil && o.Field != nil
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *VersionHistoryUpdate) SetField(v string) {
	o.Field = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VersionHistoryUpdate) GetType() VersionHistoryUpdateType {
	if o == nil || o.Type == nil {
		var ret VersionHistoryUpdateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionHistoryUpdate) GetTypeOk() (*VersionHistoryUpdateType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VersionHistoryUpdate) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given VersionHistoryUpdateType and assigns it to the Type field.
func (o *VersionHistoryUpdate) SetType(v VersionHistoryUpdateType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o VersionHistoryUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Change != nil {
		toSerialize["change"] = o.Change
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
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
func (o *VersionHistoryUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Change *string                   `json:"change,omitempty"`
		Field  *string                   `json:"field,omitempty"`
		Type   *VersionHistoryUpdateType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change", "field", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Change = all.Change
	o.Field = all.Field
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
