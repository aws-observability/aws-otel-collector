// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NotificationRulePreviewResponseAttributes Attributes of the notification preview response.
type NotificationRulePreviewResponseAttributes struct {
	// List of preview results for each rule type matched by the notification rule.
	PreviewResults []NotificationRulePreviewResult `json:"preview_results"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNotificationRulePreviewResponseAttributes instantiates a new NotificationRulePreviewResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNotificationRulePreviewResponseAttributes(previewResults []NotificationRulePreviewResult) *NotificationRulePreviewResponseAttributes {
	this := NotificationRulePreviewResponseAttributes{}
	this.PreviewResults = previewResults
	return &this
}

// NewNotificationRulePreviewResponseAttributesWithDefaults instantiates a new NotificationRulePreviewResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNotificationRulePreviewResponseAttributesWithDefaults() *NotificationRulePreviewResponseAttributes {
	this := NotificationRulePreviewResponseAttributes{}
	return &this
}

// GetPreviewResults returns the PreviewResults field value.
func (o *NotificationRulePreviewResponseAttributes) GetPreviewResults() []NotificationRulePreviewResult {
	if o == nil {
		var ret []NotificationRulePreviewResult
		return ret
	}
	return o.PreviewResults
}

// GetPreviewResultsOk returns a tuple with the PreviewResults field value
// and a boolean to check if the value has been set.
func (o *NotificationRulePreviewResponseAttributes) GetPreviewResultsOk() (*[]NotificationRulePreviewResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviewResults, true
}

// SetPreviewResults sets field value.
func (o *NotificationRulePreviewResponseAttributes) SetPreviewResults(v []NotificationRulePreviewResult) {
	o.PreviewResults = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NotificationRulePreviewResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["preview_results"] = o.PreviewResults

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NotificationRulePreviewResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PreviewResults *[]NotificationRulePreviewResult `json:"preview_results"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PreviewResults == nil {
		return fmt.Errorf("required field preview_results missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"preview_results"})
	} else {
		return err
	}
	o.PreviewResults = *all.PreviewResults

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
