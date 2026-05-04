// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsSuiteSearchResponseDataAttributes Synthetics suite search response data attributes
type SyntheticsSuiteSearchResponseDataAttributes struct {
	// List of Synthetic suites matching the search query.
	Suites []SyntheticsSuite `json:"suites,omitempty"`
	// Total number of Synthetic suites matching the search query.
	Total *int32 `json:"total,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsSuiteSearchResponseDataAttributes instantiates a new SyntheticsSuiteSearchResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsSuiteSearchResponseDataAttributes() *SyntheticsSuiteSearchResponseDataAttributes {
	this := SyntheticsSuiteSearchResponseDataAttributes{}
	return &this
}

// NewSyntheticsSuiteSearchResponseDataAttributesWithDefaults instantiates a new SyntheticsSuiteSearchResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsSuiteSearchResponseDataAttributesWithDefaults() *SyntheticsSuiteSearchResponseDataAttributes {
	this := SyntheticsSuiteSearchResponseDataAttributes{}
	return &this
}

// GetSuites returns the Suites field value if set, zero value otherwise.
func (o *SyntheticsSuiteSearchResponseDataAttributes) GetSuites() []SyntheticsSuite {
	if o == nil || o.Suites == nil {
		var ret []SyntheticsSuite
		return ret
	}
	return o.Suites
}

// GetSuitesOk returns a tuple with the Suites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSuiteSearchResponseDataAttributes) GetSuitesOk() (*[]SyntheticsSuite, bool) {
	if o == nil || o.Suites == nil {
		return nil, false
	}
	return &o.Suites, true
}

// HasSuites returns a boolean if a field has been set.
func (o *SyntheticsSuiteSearchResponseDataAttributes) HasSuites() bool {
	return o != nil && o.Suites != nil
}

// SetSuites gets a reference to the given []SyntheticsSuite and assigns it to the Suites field.
func (o *SyntheticsSuiteSearchResponseDataAttributes) SetSuites(v []SyntheticsSuite) {
	o.Suites = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SyntheticsSuiteSearchResponseDataAttributes) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSuiteSearchResponseDataAttributes) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SyntheticsSuiteSearchResponseDataAttributes) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *SyntheticsSuiteSearchResponseDataAttributes) SetTotal(v int32) {
	o.Total = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsSuiteSearchResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Suites != nil {
		toSerialize["suites"] = o.Suites
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsSuiteSearchResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Suites []SyntheticsSuite `json:"suites,omitempty"`
		Total  *int32            `json:"total,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"suites", "total"})
	} else {
		return err
	}
	o.Suites = all.Suites
	o.Total = all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
