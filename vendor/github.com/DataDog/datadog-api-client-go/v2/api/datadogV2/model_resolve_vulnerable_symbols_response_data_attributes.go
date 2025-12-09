// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResolveVulnerableSymbolsResponseDataAttributes
type ResolveVulnerableSymbolsResponseDataAttributes struct {
	//
	Results []ResolveVulnerableSymbolsResponseResults `json:"results,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResolveVulnerableSymbolsResponseDataAttributes instantiates a new ResolveVulnerableSymbolsResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResolveVulnerableSymbolsResponseDataAttributes() *ResolveVulnerableSymbolsResponseDataAttributes {
	this := ResolveVulnerableSymbolsResponseDataAttributes{}
	return &this
}

// NewResolveVulnerableSymbolsResponseDataAttributesWithDefaults instantiates a new ResolveVulnerableSymbolsResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResolveVulnerableSymbolsResponseDataAttributesWithDefaults() *ResolveVulnerableSymbolsResponseDataAttributes {
	this := ResolveVulnerableSymbolsResponseDataAttributes{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseDataAttributes) GetResults() []ResolveVulnerableSymbolsResponseResults {
	if o == nil || o.Results == nil {
		var ret []ResolveVulnerableSymbolsResponseResults
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseDataAttributes) GetResultsOk() (*[]ResolveVulnerableSymbolsResponseResults, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseDataAttributes) HasResults() bool {
	return o != nil && o.Results != nil
}

// SetResults gets a reference to the given []ResolveVulnerableSymbolsResponseResults and assigns it to the Results field.
func (o *ResolveVulnerableSymbolsResponseDataAttributes) SetResults(v []ResolveVulnerableSymbolsResponseResults) {
	o.Results = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResolveVulnerableSymbolsResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResolveVulnerableSymbolsResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Results []ResolveVulnerableSymbolsResponseResults `json:"results,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"results"})
	} else {
		return err
	}
	o.Results = all.Results

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
