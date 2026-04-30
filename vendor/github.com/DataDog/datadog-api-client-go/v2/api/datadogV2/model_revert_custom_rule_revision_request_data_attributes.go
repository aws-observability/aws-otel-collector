// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RevertCustomRuleRevisionRequestDataAttributes Attributes specifying the current and target revision IDs for a revert operation.
type RevertCustomRuleRevisionRequestDataAttributes struct {
	// Current revision ID
	CurrentRevision *string `json:"currentRevision,omitempty"`
	// Target revision ID to revert to
	RevertToRevision *string `json:"revertToRevision,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRevertCustomRuleRevisionRequestDataAttributes instantiates a new RevertCustomRuleRevisionRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRevertCustomRuleRevisionRequestDataAttributes() *RevertCustomRuleRevisionRequestDataAttributes {
	this := RevertCustomRuleRevisionRequestDataAttributes{}
	return &this
}

// NewRevertCustomRuleRevisionRequestDataAttributesWithDefaults instantiates a new RevertCustomRuleRevisionRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRevertCustomRuleRevisionRequestDataAttributesWithDefaults() *RevertCustomRuleRevisionRequestDataAttributes {
	this := RevertCustomRuleRevisionRequestDataAttributes{}
	return &this
}

// GetCurrentRevision returns the CurrentRevision field value if set, zero value otherwise.
func (o *RevertCustomRuleRevisionRequestDataAttributes) GetCurrentRevision() string {
	if o == nil || o.CurrentRevision == nil {
		var ret string
		return ret
	}
	return *o.CurrentRevision
}

// GetCurrentRevisionOk returns a tuple with the CurrentRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevertCustomRuleRevisionRequestDataAttributes) GetCurrentRevisionOk() (*string, bool) {
	if o == nil || o.CurrentRevision == nil {
		return nil, false
	}
	return o.CurrentRevision, true
}

// HasCurrentRevision returns a boolean if a field has been set.
func (o *RevertCustomRuleRevisionRequestDataAttributes) HasCurrentRevision() bool {
	return o != nil && o.CurrentRevision != nil
}

// SetCurrentRevision gets a reference to the given string and assigns it to the CurrentRevision field.
func (o *RevertCustomRuleRevisionRequestDataAttributes) SetCurrentRevision(v string) {
	o.CurrentRevision = &v
}

// GetRevertToRevision returns the RevertToRevision field value if set, zero value otherwise.
func (o *RevertCustomRuleRevisionRequestDataAttributes) GetRevertToRevision() string {
	if o == nil || o.RevertToRevision == nil {
		var ret string
		return ret
	}
	return *o.RevertToRevision
}

// GetRevertToRevisionOk returns a tuple with the RevertToRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevertCustomRuleRevisionRequestDataAttributes) GetRevertToRevisionOk() (*string, bool) {
	if o == nil || o.RevertToRevision == nil {
		return nil, false
	}
	return o.RevertToRevision, true
}

// HasRevertToRevision returns a boolean if a field has been set.
func (o *RevertCustomRuleRevisionRequestDataAttributes) HasRevertToRevision() bool {
	return o != nil && o.RevertToRevision != nil
}

// SetRevertToRevision gets a reference to the given string and assigns it to the RevertToRevision field.
func (o *RevertCustomRuleRevisionRequestDataAttributes) SetRevertToRevision(v string) {
	o.RevertToRevision = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RevertCustomRuleRevisionRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CurrentRevision != nil {
		toSerialize["currentRevision"] = o.CurrentRevision
	}
	if o.RevertToRevision != nil {
		toSerialize["revertToRevision"] = o.RevertToRevision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RevertCustomRuleRevisionRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentRevision  *string `json:"currentRevision,omitempty"`
		RevertToRevision *string `json:"revertToRevision,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"currentRevision", "revertToRevision"})
	} else {
		return err
	}
	o.CurrentRevision = all.CurrentRevision
	o.RevertToRevision = all.RevertToRevision

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
