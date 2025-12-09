// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleArray Response with notification rules.
type IncidentNotificationRuleArray struct {
	// The `NotificationRuleArray` `data`.
	Data []IncidentNotificationRuleResponseData `json:"data"`
	// Related objects that are included in the response.
	Included []IncidentNotificationRuleIncludedItems `json:"included,omitempty"`
	// Response metadata.
	Meta *IncidentNotificationRuleArrayMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationRuleArray instantiates a new IncidentNotificationRuleArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationRuleArray(data []IncidentNotificationRuleResponseData) *IncidentNotificationRuleArray {
	this := IncidentNotificationRuleArray{}
	this.Data = data
	return &this
}

// NewIncidentNotificationRuleArrayWithDefaults instantiates a new IncidentNotificationRuleArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationRuleArrayWithDefaults() *IncidentNotificationRuleArray {
	this := IncidentNotificationRuleArray{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentNotificationRuleArray) GetData() []IncidentNotificationRuleResponseData {
	if o == nil {
		var ret []IncidentNotificationRuleResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleArray) GetDataOk() (*[]IncidentNotificationRuleResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentNotificationRuleArray) SetData(v []IncidentNotificationRuleResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentNotificationRuleArray) GetIncluded() []IncidentNotificationRuleIncludedItems {
	if o == nil || o.Included == nil {
		var ret []IncidentNotificationRuleIncludedItems
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleArray) GetIncludedOk() (*[]IncidentNotificationRuleIncludedItems, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentNotificationRuleArray) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentNotificationRuleIncludedItems and assigns it to the Included field.
func (o *IncidentNotificationRuleArray) SetIncluded(v []IncidentNotificationRuleIncludedItems) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentNotificationRuleArray) GetMeta() IncidentNotificationRuleArrayMeta {
	if o == nil || o.Meta == nil {
		var ret IncidentNotificationRuleArrayMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleArray) GetMetaOk() (*IncidentNotificationRuleArrayMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentNotificationRuleArray) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given IncidentNotificationRuleArrayMeta and assigns it to the Meta field.
func (o *IncidentNotificationRuleArray) SetMeta(v IncidentNotificationRuleArrayMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationRuleArray) MarshalJSON() ([]byte, error) {
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
func (o *IncidentNotificationRuleArray) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]IncidentNotificationRuleResponseData `json:"data"`
		Included []IncidentNotificationRuleIncludedItems `json:"included,omitempty"`
		Meta     *IncidentNotificationRuleArrayMeta      `json:"meta,omitempty"`
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
