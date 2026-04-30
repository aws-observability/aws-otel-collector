// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserJourneySearchTarget Target for user journey search.
type UserJourneySearchTarget struct {
	// End value.
	End *string `json:"end,omitempty"`
	// Start value.
	Start *string `json:"start,omitempty"`
	// Target type.
	Type string `json:"type"`
	// Target value.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserJourneySearchTarget instantiates a new UserJourneySearchTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserJourneySearchTarget(typeVar string) *UserJourneySearchTarget {
	this := UserJourneySearchTarget{}
	this.Type = typeVar
	return &this
}

// NewUserJourneySearchTargetWithDefaults instantiates a new UserJourneySearchTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserJourneySearchTargetWithDefaults() *UserJourneySearchTarget {
	this := UserJourneySearchTarget{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *UserJourneySearchTarget) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearchTarget) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *UserJourneySearchTarget) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *UserJourneySearchTarget) SetEnd(v string) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *UserJourneySearchTarget) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearchTarget) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *UserJourneySearchTarget) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *UserJourneySearchTarget) SetStart(v string) {
	o.Start = &v
}

// GetType returns the Type field value.
func (o *UserJourneySearchTarget) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserJourneySearchTarget) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UserJourneySearchTarget) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UserJourneySearchTarget) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearchTarget) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UserJourneySearchTarget) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UserJourneySearchTarget) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserJourneySearchTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	toSerialize["type"] = o.Type
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserJourneySearchTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   *string `json:"end,omitempty"`
		Start *string `json:"start,omitempty"`
		Type  *string `json:"type"`
		Value *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "start", "type", "value"})
	} else {
		return err
	}
	o.End = all.End
	o.Start = all.Start
	o.Type = *all.Type
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
