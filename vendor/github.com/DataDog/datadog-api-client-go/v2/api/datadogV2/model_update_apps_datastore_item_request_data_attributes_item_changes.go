// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateAppsDatastoreItemRequestDataAttributesItemChanges Changes to apply to a datastore item using set operations.
type UpdateAppsDatastoreItemRequestDataAttributesItemChanges struct {
	// Set operation that contains key-value pairs to set on the datastore item.
	OpsSet map[string]interface{} `json:"ops_set,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateAppsDatastoreItemRequestDataAttributesItemChanges instantiates a new UpdateAppsDatastoreItemRequestDataAttributesItemChanges object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateAppsDatastoreItemRequestDataAttributesItemChanges() *UpdateAppsDatastoreItemRequestDataAttributesItemChanges {
	this := UpdateAppsDatastoreItemRequestDataAttributesItemChanges{}
	return &this
}

// NewUpdateAppsDatastoreItemRequestDataAttributesItemChangesWithDefaults instantiates a new UpdateAppsDatastoreItemRequestDataAttributesItemChanges object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateAppsDatastoreItemRequestDataAttributesItemChangesWithDefaults() *UpdateAppsDatastoreItemRequestDataAttributesItemChanges {
	this := UpdateAppsDatastoreItemRequestDataAttributesItemChanges{}
	return &this
}

// GetOpsSet returns the OpsSet field value if set, zero value otherwise.
func (o *UpdateAppsDatastoreItemRequestDataAttributesItemChanges) GetOpsSet() map[string]interface{} {
	if o == nil || o.OpsSet == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.OpsSet
}

// GetOpsSetOk returns a tuple with the OpsSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAppsDatastoreItemRequestDataAttributesItemChanges) GetOpsSetOk() (*map[string]interface{}, bool) {
	if o == nil || o.OpsSet == nil {
		return nil, false
	}
	return &o.OpsSet, true
}

// HasOpsSet returns a boolean if a field has been set.
func (o *UpdateAppsDatastoreItemRequestDataAttributesItemChanges) HasOpsSet() bool {
	return o != nil && o.OpsSet != nil
}

// SetOpsSet gets a reference to the given map[string]interface{} and assigns it to the OpsSet field.
func (o *UpdateAppsDatastoreItemRequestDataAttributesItemChanges) SetOpsSet(v map[string]interface{}) {
	o.OpsSet = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateAppsDatastoreItemRequestDataAttributesItemChanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OpsSet != nil {
		toSerialize["ops_set"] = o.OpsSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateAppsDatastoreItemRequestDataAttributesItemChanges) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpsSet map[string]interface{} `json:"ops_set,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ops_set"})
	} else {
		return err
	}
	o.OpsSet = all.OpsSet

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
