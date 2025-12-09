// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResolveVulnerableSymbolsResponseResultsVulnerableSymbols
type ResolveVulnerableSymbolsResponseResultsVulnerableSymbols struct {
	//
	AdvisoryId *string `json:"advisory_id,omitempty"`
	//
	Symbols []ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols `json:"symbols,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResolveVulnerableSymbolsResponseResultsVulnerableSymbols instantiates a new ResolveVulnerableSymbolsResponseResultsVulnerableSymbols object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResolveVulnerableSymbolsResponseResultsVulnerableSymbols() *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols {
	this := ResolveVulnerableSymbolsResponseResultsVulnerableSymbols{}
	return &this
}

// NewResolveVulnerableSymbolsResponseResultsVulnerableSymbolsWithDefaults instantiates a new ResolveVulnerableSymbolsResponseResultsVulnerableSymbols object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResolveVulnerableSymbolsResponseResultsVulnerableSymbolsWithDefaults() *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols {
	this := ResolveVulnerableSymbolsResponseResultsVulnerableSymbols{}
	return &this
}

// GetAdvisoryId returns the AdvisoryId field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) GetAdvisoryId() string {
	if o == nil || o.AdvisoryId == nil {
		var ret string
		return ret
	}
	return *o.AdvisoryId
}

// GetAdvisoryIdOk returns a tuple with the AdvisoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) GetAdvisoryIdOk() (*string, bool) {
	if o == nil || o.AdvisoryId == nil {
		return nil, false
	}
	return o.AdvisoryId, true
}

// HasAdvisoryId returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) HasAdvisoryId() bool {
	return o != nil && o.AdvisoryId != nil
}

// SetAdvisoryId gets a reference to the given string and assigns it to the AdvisoryId field.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) SetAdvisoryId(v string) {
	o.AdvisoryId = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) GetSymbols() []ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols {
	if o == nil || o.Symbols == nil {
		var ret []ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) GetSymbolsOk() (*[]ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols, bool) {
	if o == nil || o.Symbols == nil {
		return nil, false
	}
	return &o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) HasSymbols() bool {
	return o != nil && o.Symbols != nil
}

// SetSymbols gets a reference to the given []ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols and assigns it to the Symbols field.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) SetSymbols(v []ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) {
	o.Symbols = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdvisoryId != nil {
		toSerialize["advisory_id"] = o.AdvisoryId
	}
	if o.Symbols != nil {
		toSerialize["symbols"] = o.Symbols
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbols) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdvisoryId *string                                                           `json:"advisory_id,omitempty"`
		Symbols    []ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols `json:"symbols,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"advisory_id", "symbols"})
	} else {
		return err
	}
	o.AdvisoryId = all.AdvisoryId
	o.Symbols = all.Symbols

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
