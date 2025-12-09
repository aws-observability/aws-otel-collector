// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedRelatedOncallAttributes Included related oncall attributes.
type EntityResponseIncludedRelatedOncallAttributes struct {
	// Oncall escalations.
	Escalations []EntityResponseIncludedRelatedOncallEscalationItem `json:"escalations,omitempty"`
	// Oncall provider.
	Provider *string `json:"provider,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseIncludedRelatedOncallAttributes instantiates a new EntityResponseIncludedRelatedOncallAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseIncludedRelatedOncallAttributes() *EntityResponseIncludedRelatedOncallAttributes {
	this := EntityResponseIncludedRelatedOncallAttributes{}
	return &this
}

// NewEntityResponseIncludedRelatedOncallAttributesWithDefaults instantiates a new EntityResponseIncludedRelatedOncallAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseIncludedRelatedOncallAttributesWithDefaults() *EntityResponseIncludedRelatedOncallAttributes {
	this := EntityResponseIncludedRelatedOncallAttributes{}
	return &this
}

// GetEscalations returns the Escalations field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedOncallAttributes) GetEscalations() []EntityResponseIncludedRelatedOncallEscalationItem {
	if o == nil || o.Escalations == nil {
		var ret []EntityResponseIncludedRelatedOncallEscalationItem
		return ret
	}
	return o.Escalations
}

// GetEscalationsOk returns a tuple with the Escalations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedOncallAttributes) GetEscalationsOk() (*[]EntityResponseIncludedRelatedOncallEscalationItem, bool) {
	if o == nil || o.Escalations == nil {
		return nil, false
	}
	return &o.Escalations, true
}

// HasEscalations returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedOncallAttributes) HasEscalations() bool {
	return o != nil && o.Escalations != nil
}

// SetEscalations gets a reference to the given []EntityResponseIncludedRelatedOncallEscalationItem and assigns it to the Escalations field.
func (o *EntityResponseIncludedRelatedOncallAttributes) SetEscalations(v []EntityResponseIncludedRelatedOncallEscalationItem) {
	o.Escalations = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedOncallAttributes) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedOncallAttributes) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedOncallAttributes) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *EntityResponseIncludedRelatedOncallAttributes) SetProvider(v string) {
	o.Provider = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseIncludedRelatedOncallAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Escalations != nil {
		toSerialize["escalations"] = o.Escalations
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityResponseIncludedRelatedOncallAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Escalations []EntityResponseIncludedRelatedOncallEscalationItem `json:"escalations,omitempty"`
		Provider    *string                                             `json:"provider,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"escalations", "provider"})
	} else {
		return err
	}
	o.Escalations = all.Escalations
	o.Provider = all.Provider

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
