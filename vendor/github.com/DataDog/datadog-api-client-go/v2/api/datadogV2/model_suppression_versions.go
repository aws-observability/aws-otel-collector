// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SuppressionVersions A suppression version with a list of updates.
type SuppressionVersions struct {
	// A list of changes.
	Changes []VersionHistoryUpdate `json:"changes,omitempty"`
	// The attributes of the suppression rule.
	Suppression *SecurityMonitoringSuppressionAttributes `json:"suppression,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSuppressionVersions instantiates a new SuppressionVersions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSuppressionVersions() *SuppressionVersions {
	this := SuppressionVersions{}
	return &this
}

// NewSuppressionVersionsWithDefaults instantiates a new SuppressionVersions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSuppressionVersionsWithDefaults() *SuppressionVersions {
	this := SuppressionVersions{}
	return &this
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *SuppressionVersions) GetChanges() []VersionHistoryUpdate {
	if o == nil || o.Changes == nil {
		var ret []VersionHistoryUpdate
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuppressionVersions) GetChangesOk() (*[]VersionHistoryUpdate, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return &o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *SuppressionVersions) HasChanges() bool {
	return o != nil && o.Changes != nil
}

// SetChanges gets a reference to the given []VersionHistoryUpdate and assigns it to the Changes field.
func (o *SuppressionVersions) SetChanges(v []VersionHistoryUpdate) {
	o.Changes = v
}

// GetSuppression returns the Suppression field value if set, zero value otherwise.
func (o *SuppressionVersions) GetSuppression() SecurityMonitoringSuppressionAttributes {
	if o == nil || o.Suppression == nil {
		var ret SecurityMonitoringSuppressionAttributes
		return ret
	}
	return *o.Suppression
}

// GetSuppressionOk returns a tuple with the Suppression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuppressionVersions) GetSuppressionOk() (*SecurityMonitoringSuppressionAttributes, bool) {
	if o == nil || o.Suppression == nil {
		return nil, false
	}
	return o.Suppression, true
}

// HasSuppression returns a boolean if a field has been set.
func (o *SuppressionVersions) HasSuppression() bool {
	return o != nil && o.Suppression != nil
}

// SetSuppression gets a reference to the given SecurityMonitoringSuppressionAttributes and assigns it to the Suppression field.
func (o *SuppressionVersions) SetSuppression(v SecurityMonitoringSuppressionAttributes) {
	o.Suppression = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SuppressionVersions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.Suppression != nil {
		toSerialize["suppression"] = o.Suppression
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SuppressionVersions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Changes     []VersionHistoryUpdate                   `json:"changes,omitempty"`
		Suppression *SecurityMonitoringSuppressionAttributes `json:"suppression,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"changes", "suppression"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Changes = all.Changes
	if all.Suppression != nil && all.Suppression.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Suppression = all.Suppression

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
