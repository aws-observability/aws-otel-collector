// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionChangeMetadataItem Object describing a single change within a version.
type SyntheticsTestVersionChangeMetadataItem struct {
	// The action that was performed (for example, `updated` or `created`).
	Action *string `json:"action,omitempty"`
	// Object containing metadata about a change action.
	ActionMetadata *SyntheticsTestVersionActionMetadata `json:"action_metadata,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestVersionChangeMetadataItem instantiates a new SyntheticsTestVersionChangeMetadataItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestVersionChangeMetadataItem() *SyntheticsTestVersionChangeMetadataItem {
	this := SyntheticsTestVersionChangeMetadataItem{}
	return &this
}

// NewSyntheticsTestVersionChangeMetadataItemWithDefaults instantiates a new SyntheticsTestVersionChangeMetadataItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestVersionChangeMetadataItemWithDefaults() *SyntheticsTestVersionChangeMetadataItem {
	this := SyntheticsTestVersionChangeMetadataItem{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SyntheticsTestVersionChangeMetadataItem) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionChangeMetadataItem) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SyntheticsTestVersionChangeMetadataItem) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SyntheticsTestVersionChangeMetadataItem) SetAction(v string) {
	o.Action = &v
}

// GetActionMetadata returns the ActionMetadata field value if set, zero value otherwise.
func (o *SyntheticsTestVersionChangeMetadataItem) GetActionMetadata() SyntheticsTestVersionActionMetadata {
	if o == nil || o.ActionMetadata == nil {
		var ret SyntheticsTestVersionActionMetadata
		return ret
	}
	return *o.ActionMetadata
}

// GetActionMetadataOk returns a tuple with the ActionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionChangeMetadataItem) GetActionMetadataOk() (*SyntheticsTestVersionActionMetadata, bool) {
	if o == nil || o.ActionMetadata == nil {
		return nil, false
	}
	return o.ActionMetadata, true
}

// HasActionMetadata returns a boolean if a field has been set.
func (o *SyntheticsTestVersionChangeMetadataItem) HasActionMetadata() bool {
	return o != nil && o.ActionMetadata != nil
}

// SetActionMetadata gets a reference to the given SyntheticsTestVersionActionMetadata and assigns it to the ActionMetadata field.
func (o *SyntheticsTestVersionChangeMetadataItem) SetActionMetadata(v SyntheticsTestVersionActionMetadata) {
	o.ActionMetadata = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestVersionChangeMetadataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.ActionMetadata != nil {
		toSerialize["action_metadata"] = o.ActionMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestVersionChangeMetadataItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action         *string                              `json:"action,omitempty"`
		ActionMetadata *SyntheticsTestVersionActionMetadata `json:"action_metadata,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "action_metadata"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Action = all.Action
	if all.ActionMetadata != nil && all.ActionMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ActionMetadata = all.ActionMetadata

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
