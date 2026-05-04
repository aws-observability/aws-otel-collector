// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseInsightsItems An insight of the case.
type CaseInsightsItems struct {
	// Reference of the insight.
	Ref *string `json:"ref,omitempty"`
	// Unique identifier of the resource. For example, the unique identifier of a security finding.
	ResourceId *string `json:"resource_id,omitempty"`
	// Type of the resource. For example, the type of a security finding is "SECURITY_FINDING".
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseInsightsItems instantiates a new CaseInsightsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseInsightsItems() *CaseInsightsItems {
	this := CaseInsightsItems{}
	return &this
}

// NewCaseInsightsItemsWithDefaults instantiates a new CaseInsightsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseInsightsItemsWithDefaults() *CaseInsightsItems {
	this := CaseInsightsItems{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *CaseInsightsItems) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseInsightsItems) GetRefOk() (*string, bool) {
	if o == nil || o.Ref == nil {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *CaseInsightsItems) HasRef() bool {
	return o != nil && o.Ref != nil
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *CaseInsightsItems) SetRef(v string) {
	o.Ref = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *CaseInsightsItems) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseInsightsItems) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *CaseInsightsItems) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *CaseInsightsItems) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CaseInsightsItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseInsightsItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CaseInsightsItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CaseInsightsItems) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseInsightsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ref != nil {
		toSerialize["ref"] = o.Ref
	}
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseInsightsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ref        *string `json:"ref,omitempty"`
		ResourceId *string `json:"resource_id,omitempty"`
		Type       *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ref", "resource_id", "type"})
	} else {
		return err
	}
	o.Ref = all.Ref
	o.ResourceId = all.ResourceId
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
