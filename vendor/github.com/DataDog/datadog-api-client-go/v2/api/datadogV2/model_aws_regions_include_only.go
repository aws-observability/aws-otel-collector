// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSRegionsIncludeOnly Include only these regions.
type AWSRegionsIncludeOnly struct {
	// Include only these regions.
	IncludeOnly []string `json:"include_only"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSRegionsIncludeOnly instantiates a new AWSRegionsIncludeOnly object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSRegionsIncludeOnly(includeOnly []string) *AWSRegionsIncludeOnly {
	this := AWSRegionsIncludeOnly{}
	this.IncludeOnly = includeOnly
	return &this
}

// NewAWSRegionsIncludeOnlyWithDefaults instantiates a new AWSRegionsIncludeOnly object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSRegionsIncludeOnlyWithDefaults() *AWSRegionsIncludeOnly {
	this := AWSRegionsIncludeOnly{}
	return &this
}

// GetIncludeOnly returns the IncludeOnly field value.
func (o *AWSRegionsIncludeOnly) GetIncludeOnly() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IncludeOnly
}

// GetIncludeOnlyOk returns a tuple with the IncludeOnly field value
// and a boolean to check if the value has been set.
func (o *AWSRegionsIncludeOnly) GetIncludeOnlyOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeOnly, true
}

// SetIncludeOnly sets field value.
func (o *AWSRegionsIncludeOnly) SetIncludeOnly(v []string) {
	o.IncludeOnly = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSRegionsIncludeOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["include_only"] = o.IncludeOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSRegionsIncludeOnly) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeOnly *[]string `json:"include_only"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncludeOnly == nil {
		return fmt.Errorf("required field include_only missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_only"})
	} else {
		return err
	}
	o.IncludeOnly = *all.IncludeOnly

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
