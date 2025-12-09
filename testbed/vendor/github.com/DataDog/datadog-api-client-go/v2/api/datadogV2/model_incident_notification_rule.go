// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRule Response with a notification rule.
type IncidentNotificationRule struct {
	// Notification rule data from a response.
	Data IncidentNotificationRuleResponseData `json:"data"`
	// Related objects that are included in the response.
	Included []IncidentNotificationRuleIncludedItems `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationRule instantiates a new IncidentNotificationRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationRule(data IncidentNotificationRuleResponseData) *IncidentNotificationRule {
	this := IncidentNotificationRule{}
	this.Data = data
	return &this
}

// NewIncidentNotificationRuleWithDefaults instantiates a new IncidentNotificationRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationRuleWithDefaults() *IncidentNotificationRule {
	this := IncidentNotificationRule{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentNotificationRule) GetData() IncidentNotificationRuleResponseData {
	if o == nil {
		var ret IncidentNotificationRuleResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRule) GetDataOk() (*IncidentNotificationRuleResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentNotificationRule) SetData(v IncidentNotificationRuleResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentNotificationRule) GetIncluded() []IncidentNotificationRuleIncludedItems {
	if o == nil || o.Included == nil {
		var ret []IncidentNotificationRuleIncludedItems
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRule) GetIncludedOk() (*[]IncidentNotificationRuleIncludedItems, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentNotificationRule) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentNotificationRuleIncludedItems and assigns it to the Included field.
func (o *IncidentNotificationRule) SetIncluded(v []IncidentNotificationRuleIncludedItems) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *IncidentNotificationRuleResponseData   `json:"data"`
		Included []IncidentNotificationRuleIncludedItems `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
