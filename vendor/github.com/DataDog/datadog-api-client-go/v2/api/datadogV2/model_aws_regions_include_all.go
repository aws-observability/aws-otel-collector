// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSRegionsIncludeAll Include all regions. Defaults to `true`.
type AWSRegionsIncludeAll struct {
	// Include all regions.
	IncludeAll bool `json:"include_all"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSRegionsIncludeAll instantiates a new AWSRegionsIncludeAll object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSRegionsIncludeAll(includeAll bool) *AWSRegionsIncludeAll {
	this := AWSRegionsIncludeAll{}
	this.IncludeAll = includeAll
	return &this
}

// NewAWSRegionsIncludeAllWithDefaults instantiates a new AWSRegionsIncludeAll object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSRegionsIncludeAllWithDefaults() *AWSRegionsIncludeAll {
	this := AWSRegionsIncludeAll{}
	return &this
}

// GetIncludeAll returns the IncludeAll field value.
func (o *AWSRegionsIncludeAll) GetIncludeAll() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IncludeAll
}

// GetIncludeAllOk returns a tuple with the IncludeAll field value
// and a boolean to check if the value has been set.
func (o *AWSRegionsIncludeAll) GetIncludeAllOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeAll, true
}

// SetIncludeAll sets field value.
func (o *AWSRegionsIncludeAll) SetIncludeAll(v bool) {
	o.IncludeAll = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSRegionsIncludeAll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["include_all"] = o.IncludeAll

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSRegionsIncludeAll) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeAll *bool `json:"include_all"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncludeAll == nil {
		return fmt.Errorf("required field include_all missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_all"})
	} else {
		return err
	}
	o.IncludeAll = *all.IncludeAll

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
