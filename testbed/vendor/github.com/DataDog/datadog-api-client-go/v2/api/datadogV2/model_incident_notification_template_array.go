// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationTemplateArray Response with notification templates.
type IncidentNotificationTemplateArray struct {
	// The `NotificationTemplateArray` `data`.
	Data []IncidentNotificationTemplateResponseData `json:"data"`
	// Related objects that are included in the response.
	Included []IncidentNotificationTemplateIncludedItems `json:"included,omitempty"`
	// Response metadata.
	Meta *IncidentNotificationTemplateArrayMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationTemplateArray instantiates a new IncidentNotificationTemplateArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationTemplateArray(data []IncidentNotificationTemplateResponseData) *IncidentNotificationTemplateArray {
	this := IncidentNotificationTemplateArray{}
	this.Data = data
	return &this
}

// NewIncidentNotificationTemplateArrayWithDefaults instantiates a new IncidentNotificationTemplateArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationTemplateArrayWithDefaults() *IncidentNotificationTemplateArray {
	this := IncidentNotificationTemplateArray{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentNotificationTemplateArray) GetData() []IncidentNotificationTemplateResponseData {
	if o == nil {
		var ret []IncidentNotificationTemplateResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateArray) GetDataOk() (*[]IncidentNotificationTemplateResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentNotificationTemplateArray) SetData(v []IncidentNotificationTemplateResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentNotificationTemplateArray) GetIncluded() []IncidentNotificationTemplateIncludedItems {
	if o == nil || o.Included == nil {
		var ret []IncidentNotificationTemplateIncludedItems
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateArray) GetIncludedOk() (*[]IncidentNotificationTemplateIncludedItems, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentNotificationTemplateArray) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentNotificationTemplateIncludedItems and assigns it to the Included field.
func (o *IncidentNotificationTemplateArray) SetIncluded(v []IncidentNotificationTemplateIncludedItems) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentNotificationTemplateArray) GetMeta() IncidentNotificationTemplateArrayMeta {
	if o == nil || o.Meta == nil {
		var ret IncidentNotificationTemplateArrayMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateArray) GetMetaOk() (*IncidentNotificationTemplateArrayMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentNotificationTemplateArray) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given IncidentNotificationTemplateArrayMeta and assigns it to the Meta field.
func (o *IncidentNotificationTemplateArray) SetMeta(v IncidentNotificationTemplateArrayMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationTemplateArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationTemplateArray) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]IncidentNotificationTemplateResponseData `json:"data"`
		Included []IncidentNotificationTemplateIncludedItems `json:"included,omitempty"`
		Meta     *IncidentNotificationTemplateArrayMeta      `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	o.Included = all.Included
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
