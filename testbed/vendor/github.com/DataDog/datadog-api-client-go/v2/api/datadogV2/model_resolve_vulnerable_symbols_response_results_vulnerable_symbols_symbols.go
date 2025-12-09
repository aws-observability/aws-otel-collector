// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols
type ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols struct {
	//
	Name *string `json:"name,omitempty"`
	//
	Type *string `json:"type,omitempty"`
	//
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols instantiates a new ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols() *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols {
	this := ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols{}
	return &this
}

// NewResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbolsWithDefaults instantiates a new ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbolsWithDefaults() *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols {
	this := ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResolveVulnerableSymbolsResponseResultsVulnerableSymbolsSymbols) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string `json:"name,omitempty"`
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "type", "value"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Type = all.Type
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
