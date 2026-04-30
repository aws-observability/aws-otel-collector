// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamSyncSelectionStateItem Identifies a team or organization hierarchy to include in synchronization.
type TeamSyncSelectionStateItem struct {
	// The external identifier for a team or organization in the source platform.
	ExternalId TeamSyncSelectionStateExternalId `json:"external_id"`
	// The operation to perform on the selected hierarchy.
	// When set to `include`, synchronization covers the
	// referenced teams or organizations.
	Operation *TeamSyncSelectionStateOperation `json:"operation,omitempty"`
	// The scope of the selection. When set to `subtree`,
	// synchronization includes the referenced team or
	// organization and everything nested under it.
	Scope *TeamSyncSelectionStateScope `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamSyncSelectionStateItem instantiates a new TeamSyncSelectionStateItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamSyncSelectionStateItem(externalId TeamSyncSelectionStateExternalId) *TeamSyncSelectionStateItem {
	this := TeamSyncSelectionStateItem{}
	this.ExternalId = externalId
	return &this
}

// NewTeamSyncSelectionStateItemWithDefaults instantiates a new TeamSyncSelectionStateItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamSyncSelectionStateItemWithDefaults() *TeamSyncSelectionStateItem {
	this := TeamSyncSelectionStateItem{}
	return &this
}

// GetExternalId returns the ExternalId field value.
func (o *TeamSyncSelectionStateItem) GetExternalId() TeamSyncSelectionStateExternalId {
	if o == nil {
		var ret TeamSyncSelectionStateExternalId
		return ret
	}
	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *TeamSyncSelectionStateItem) GetExternalIdOk() (*TeamSyncSelectionStateExternalId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value.
func (o *TeamSyncSelectionStateItem) SetExternalId(v TeamSyncSelectionStateExternalId) {
	o.ExternalId = v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *TeamSyncSelectionStateItem) GetOperation() TeamSyncSelectionStateOperation {
	if o == nil || o.Operation == nil {
		var ret TeamSyncSelectionStateOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamSyncSelectionStateItem) GetOperationOk() (*TeamSyncSelectionStateOperation, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *TeamSyncSelectionStateItem) HasOperation() bool {
	return o != nil && o.Operation != nil
}

// SetOperation gets a reference to the given TeamSyncSelectionStateOperation and assigns it to the Operation field.
func (o *TeamSyncSelectionStateItem) SetOperation(v TeamSyncSelectionStateOperation) {
	o.Operation = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *TeamSyncSelectionStateItem) GetScope() TeamSyncSelectionStateScope {
	if o == nil || o.Scope == nil {
		var ret TeamSyncSelectionStateScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamSyncSelectionStateItem) GetScopeOk() (*TeamSyncSelectionStateScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *TeamSyncSelectionStateItem) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given TeamSyncSelectionStateScope and assigns it to the Scope field.
func (o *TeamSyncSelectionStateItem) SetScope(v TeamSyncSelectionStateScope) {
	o.Scope = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamSyncSelectionStateItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["external_id"] = o.ExternalId
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamSyncSelectionStateItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExternalId *TeamSyncSelectionStateExternalId `json:"external_id"`
		Operation  *TeamSyncSelectionStateOperation  `json:"operation,omitempty"`
		Scope      *TeamSyncSelectionStateScope      `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExternalId == nil {
		return fmt.Errorf("required field external_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"external_id", "operation", "scope"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ExternalId.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExternalId = *all.ExternalId
	if all.Operation != nil && !all.Operation.IsValid() {
		hasInvalidField = true
	} else {
		o.Operation = all.Operation
	}
	if all.Scope != nil && !all.Scope.IsValid() {
		hasInvalidField = true
	} else {
		o.Scope = all.Scope
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
