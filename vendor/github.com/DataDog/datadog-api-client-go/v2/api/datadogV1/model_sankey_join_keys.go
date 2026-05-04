// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SankeyJoinKeys Join keys.
type SankeyJoinKeys struct {
	// Primary join key.
	Primary string `json:"primary"`
	// Secondary join keys.
	Secondary []string `json:"secondary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSankeyJoinKeys instantiates a new SankeyJoinKeys object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSankeyJoinKeys(primary string) *SankeyJoinKeys {
	this := SankeyJoinKeys{}
	this.Primary = primary
	return &this
}

// NewSankeyJoinKeysWithDefaults instantiates a new SankeyJoinKeys object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSankeyJoinKeysWithDefaults() *SankeyJoinKeys {
	this := SankeyJoinKeys{}
	return &this
}

// GetPrimary returns the Primary field value.
func (o *SankeyJoinKeys) GetPrimary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *SankeyJoinKeys) GetPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value.
func (o *SankeyJoinKeys) SetPrimary(v string) {
	o.Primary = v
}

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *SankeyJoinKeys) GetSecondary() []string {
	if o == nil || o.Secondary == nil {
		var ret []string
		return ret
	}
	return o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SankeyJoinKeys) GetSecondaryOk() (*[]string, bool) {
	if o == nil || o.Secondary == nil {
		return nil, false
	}
	return &o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *SankeyJoinKeys) HasSecondary() bool {
	return o != nil && o.Secondary != nil
}

// SetSecondary gets a reference to the given []string and assigns it to the Secondary field.
func (o *SankeyJoinKeys) SetSecondary(v []string) {
	o.Secondary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SankeyJoinKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["primary"] = o.Primary
	if o.Secondary != nil {
		toSerialize["secondary"] = o.Secondary
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SankeyJoinKeys) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Primary   *string  `json:"primary"`
		Secondary []string `json:"secondary,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Primary == nil {
		return fmt.Errorf("required field primary missing")
	}
	o.Primary = *all.Primary
	o.Secondary = all.Secondary

	return nil
}
