// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseBulkUpdateRequestAttributes Attributes for the bulk update, specifying which cases to update and the action to apply.
type CaseBulkUpdateRequestAttributes struct {
	// An array of case identifiers to apply the bulk action to.
	CaseIds []string `json:"case_ids"`
	// A key-value map of action-specific parameters. The required keys depend on the action type (for example, `priority` for the priority action, `assignee_id` for assign).
	Payload map[string]string `json:"payload,omitempty"`
	// The type of action to apply in a bulk update. Allowed values are `priority`, `status`, `assign`, `unassign`, `archive`, `unarchive`, `jira`, `servicenow`, `linear`, `update_project`.
	Type CaseBulkActionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseBulkUpdateRequestAttributes instantiates a new CaseBulkUpdateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseBulkUpdateRequestAttributes(caseIds []string, typeVar CaseBulkActionType) *CaseBulkUpdateRequestAttributes {
	this := CaseBulkUpdateRequestAttributes{}
	this.CaseIds = caseIds
	this.Type = typeVar
	return &this
}

// NewCaseBulkUpdateRequestAttributesWithDefaults instantiates a new CaseBulkUpdateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseBulkUpdateRequestAttributesWithDefaults() *CaseBulkUpdateRequestAttributes {
	this := CaseBulkUpdateRequestAttributes{}
	return &this
}

// GetCaseIds returns the CaseIds field value.
func (o *CaseBulkUpdateRequestAttributes) GetCaseIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CaseIds
}

// GetCaseIdsOk returns a tuple with the CaseIds field value
// and a boolean to check if the value has been set.
func (o *CaseBulkUpdateRequestAttributes) GetCaseIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseIds, true
}

// SetCaseIds sets field value.
func (o *CaseBulkUpdateRequestAttributes) SetCaseIds(v []string) {
	o.CaseIds = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CaseBulkUpdateRequestAttributes) GetPayload() map[string]string {
	if o == nil || o.Payload == nil {
		var ret map[string]string
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseBulkUpdateRequestAttributes) GetPayloadOk() (*map[string]string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CaseBulkUpdateRequestAttributes) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given map[string]string and assigns it to the Payload field.
func (o *CaseBulkUpdateRequestAttributes) SetPayload(v map[string]string) {
	o.Payload = v
}

// GetType returns the Type field value.
func (o *CaseBulkUpdateRequestAttributes) GetType() CaseBulkActionType {
	if o == nil {
		var ret CaseBulkActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaseBulkUpdateRequestAttributes) GetTypeOk() (*CaseBulkActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CaseBulkUpdateRequestAttributes) SetType(v CaseBulkActionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseBulkUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["case_ids"] = o.CaseIds
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseBulkUpdateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaseIds *[]string           `json:"case_ids"`
		Payload map[string]string   `json:"payload,omitempty"`
		Type    *CaseBulkActionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CaseIds == nil {
		return fmt.Errorf("required field case_ids missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"case_ids", "payload", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CaseIds = *all.CaseIds
	o.Payload = all.Payload
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
