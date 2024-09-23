// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSpecifiedCustomReportsData Response containing date and type for specified custom reports.
type UsageSpecifiedCustomReportsData struct {
	// The response containing attributes for specified custom reports.
	Attributes *UsageSpecifiedCustomReportsAttributes `json:"attributes,omitempty"`
	// The date for specified custom reports.
	Id *string `json:"id,omitempty"`
	// The type of reports.
	Type *UsageReportsType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageSpecifiedCustomReportsData instantiates a new UsageSpecifiedCustomReportsData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSpecifiedCustomReportsData() *UsageSpecifiedCustomReportsData {
	this := UsageSpecifiedCustomReportsData{}
	var typeVar UsageReportsType = USAGEREPORTSTYPE_REPORTS
	this.Type = &typeVar
	return &this
}

// NewUsageSpecifiedCustomReportsDataWithDefaults instantiates a new UsageSpecifiedCustomReportsData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSpecifiedCustomReportsDataWithDefaults() *UsageSpecifiedCustomReportsData {
	this := UsageSpecifiedCustomReportsData{}
	var typeVar UsageReportsType = USAGEREPORTSTYPE_REPORTS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UsageSpecifiedCustomReportsData) GetAttributes() UsageSpecifiedCustomReportsAttributes {
	if o == nil || o.Attributes == nil {
		var ret UsageSpecifiedCustomReportsAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSpecifiedCustomReportsData) GetAttributesOk() (*UsageSpecifiedCustomReportsAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UsageSpecifiedCustomReportsData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UsageSpecifiedCustomReportsAttributes and assigns it to the Attributes field.
func (o *UsageSpecifiedCustomReportsData) SetAttributes(v UsageSpecifiedCustomReportsAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UsageSpecifiedCustomReportsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSpecifiedCustomReportsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UsageSpecifiedCustomReportsData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UsageSpecifiedCustomReportsData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UsageSpecifiedCustomReportsData) GetType() UsageReportsType {
	if o == nil || o.Type == nil {
		var ret UsageReportsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSpecifiedCustomReportsData) GetTypeOk() (*UsageReportsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UsageSpecifiedCustomReportsData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given UsageReportsType and assigns it to the Type field.
func (o *UsageSpecifiedCustomReportsData) SetType(v UsageReportsType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSpecifiedCustomReportsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *UsageSpecifiedCustomReportsData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UsageSpecifiedCustomReportsAttributes `json:"attributes,omitempty"`
		Id         *string                                `json:"id,omitempty"`
		Type       *UsageReportsType                      `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
