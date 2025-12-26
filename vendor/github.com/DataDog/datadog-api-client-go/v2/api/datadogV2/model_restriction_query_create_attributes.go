// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionQueryCreateAttributes Attributes of the created restriction query.
type RestrictionQueryCreateAttributes struct {
	// The restriction query.
	RestrictionQuery string `json:"restriction_query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestrictionQueryCreateAttributes instantiates a new RestrictionQueryCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestrictionQueryCreateAttributes(restrictionQuery string) *RestrictionQueryCreateAttributes {
	this := RestrictionQueryCreateAttributes{}
	this.RestrictionQuery = restrictionQuery
	return &this
}

// NewRestrictionQueryCreateAttributesWithDefaults instantiates a new RestrictionQueryCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestrictionQueryCreateAttributesWithDefaults() *RestrictionQueryCreateAttributes {
	this := RestrictionQueryCreateAttributes{}
	return &this
}

// GetRestrictionQuery returns the RestrictionQuery field value.
func (o *RestrictionQueryCreateAttributes) GetRestrictionQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RestrictionQuery
}

// GetRestrictionQueryOk returns a tuple with the RestrictionQuery field value
// and a boolean to check if the value has been set.
func (o *RestrictionQueryCreateAttributes) GetRestrictionQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestrictionQuery, true
}

// SetRestrictionQuery sets field value.
func (o *RestrictionQueryCreateAttributes) SetRestrictionQuery(v string) {
	o.RestrictionQuery = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestrictionQueryCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["restriction_query"] = o.RestrictionQuery

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestrictionQueryCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RestrictionQuery *string `json:"restriction_query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RestrictionQuery == nil {
		return fmt.Errorf("required field restriction_query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"restriction_query"})
	} else {
		return err
	}
	o.RestrictionQuery = *all.RestrictionQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
