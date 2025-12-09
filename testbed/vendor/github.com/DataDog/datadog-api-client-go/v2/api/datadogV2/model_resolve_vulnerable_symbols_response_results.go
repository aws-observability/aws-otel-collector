// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResolveVulnerableSymbolsResponseResults
type ResolveVulnerableSymbolsResponseResults struct {
	//
	Purl *string `json:"purl,omitempty"`
	//
	VulnerableSymbols []ResolveVulnerableSymbolsResponseResultsVulnerableSymbols `json:"vulnerable_symbols,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResolveVulnerableSymbolsResponseResults instantiates a new ResolveVulnerableSymbolsResponseResults object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResolveVulnerableSymbolsResponseResults() *ResolveVulnerableSymbolsResponseResults {
	this := ResolveVulnerableSymbolsResponseResults{}
	return &this
}

// NewResolveVulnerableSymbolsResponseResultsWithDefaults instantiates a new ResolveVulnerableSymbolsResponseResults object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResolveVulnerableSymbolsResponseResultsWithDefaults() *ResolveVulnerableSymbolsResponseResults {
	this := ResolveVulnerableSymbolsResponseResults{}
	return &this
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseResults) GetPurl() string {
	if o == nil || o.Purl == nil {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseResults) GetPurlOk() (*string, bool) {
	if o == nil || o.Purl == nil {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseResults) HasPurl() bool {
	return o != nil && o.Purl != nil
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *ResolveVulnerableSymbolsResponseResults) SetPurl(v string) {
	o.Purl = &v
}

// GetVulnerableSymbols returns the VulnerableSymbols field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseResults) GetVulnerableSymbols() []ResolveVulnerableSymbolsResponseResultsVulnerableSymbols {
	if o == nil || o.VulnerableSymbols == nil {
		var ret []ResolveVulnerableSymbolsResponseResultsVulnerableSymbols
		return ret
	}
	return o.VulnerableSymbols
}

// GetVulnerableSymbolsOk returns a tuple with the VulnerableSymbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseResults) GetVulnerableSymbolsOk() (*[]ResolveVulnerableSymbolsResponseResultsVulnerableSymbols, bool) {
	if o == nil || o.VulnerableSymbols == nil {
		return nil, false
	}
	return &o.VulnerableSymbols, true
}

// HasVulnerableSymbols returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseResults) HasVulnerableSymbols() bool {
	return o != nil && o.VulnerableSymbols != nil
}

// SetVulnerableSymbols gets a reference to the given []ResolveVulnerableSymbolsResponseResultsVulnerableSymbols and assigns it to the VulnerableSymbols field.
func (o *ResolveVulnerableSymbolsResponseResults) SetVulnerableSymbols(v []ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) {
	o.VulnerableSymbols = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResolveVulnerableSymbolsResponseResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Purl != nil {
		toSerialize["purl"] = o.Purl
	}
	if o.VulnerableSymbols != nil {
		toSerialize["vulnerable_symbols"] = o.VulnerableSymbols
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResolveVulnerableSymbolsResponseResults) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Purl              *string                                                    `json:"purl,omitempty"`
		VulnerableSymbols []ResolveVulnerableSymbolsResponseResultsVulnerableSymbols `json:"vulnerable_symbols,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"purl", "vulnerable_symbols"})
	} else {
		return err
	}
	o.Purl = all.Purl
	o.VulnerableSymbols = all.VulnerableSymbols

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
