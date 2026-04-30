// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationIncidentFieldMappingsItems Mapping between an incident user-defined field and a case field.
type IntegrationIncidentFieldMappingsItems struct {
	// The case field to map the incident field value to.
	CaseField *string `json:"case_field,omitempty"`
	// The identifier of the incident user-defined field to map from.
	IncidentUserDefinedFieldId *string `json:"incident_user_defined_field_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationIncidentFieldMappingsItems instantiates a new IntegrationIncidentFieldMappingsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationIncidentFieldMappingsItems() *IntegrationIncidentFieldMappingsItems {
	this := IntegrationIncidentFieldMappingsItems{}
	return &this
}

// NewIntegrationIncidentFieldMappingsItemsWithDefaults instantiates a new IntegrationIncidentFieldMappingsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationIncidentFieldMappingsItemsWithDefaults() *IntegrationIncidentFieldMappingsItems {
	this := IntegrationIncidentFieldMappingsItems{}
	return &this
}

// GetCaseField returns the CaseField field value if set, zero value otherwise.
func (o *IntegrationIncidentFieldMappingsItems) GetCaseField() string {
	if o == nil || o.CaseField == nil {
		var ret string
		return ret
	}
	return *o.CaseField
}

// GetCaseFieldOk returns a tuple with the CaseField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationIncidentFieldMappingsItems) GetCaseFieldOk() (*string, bool) {
	if o == nil || o.CaseField == nil {
		return nil, false
	}
	return o.CaseField, true
}

// HasCaseField returns a boolean if a field has been set.
func (o *IntegrationIncidentFieldMappingsItems) HasCaseField() bool {
	return o != nil && o.CaseField != nil
}

// SetCaseField gets a reference to the given string and assigns it to the CaseField field.
func (o *IntegrationIncidentFieldMappingsItems) SetCaseField(v string) {
	o.CaseField = &v
}

// GetIncidentUserDefinedFieldId returns the IncidentUserDefinedFieldId field value if set, zero value otherwise.
func (o *IntegrationIncidentFieldMappingsItems) GetIncidentUserDefinedFieldId() string {
	if o == nil || o.IncidentUserDefinedFieldId == nil {
		var ret string
		return ret
	}
	return *o.IncidentUserDefinedFieldId
}

// GetIncidentUserDefinedFieldIdOk returns a tuple with the IncidentUserDefinedFieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationIncidentFieldMappingsItems) GetIncidentUserDefinedFieldIdOk() (*string, bool) {
	if o == nil || o.IncidentUserDefinedFieldId == nil {
		return nil, false
	}
	return o.IncidentUserDefinedFieldId, true
}

// HasIncidentUserDefinedFieldId returns a boolean if a field has been set.
func (o *IntegrationIncidentFieldMappingsItems) HasIncidentUserDefinedFieldId() bool {
	return o != nil && o.IncidentUserDefinedFieldId != nil
}

// SetIncidentUserDefinedFieldId gets a reference to the given string and assigns it to the IncidentUserDefinedFieldId field.
func (o *IntegrationIncidentFieldMappingsItems) SetIncidentUserDefinedFieldId(v string) {
	o.IncidentUserDefinedFieldId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationIncidentFieldMappingsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CaseField != nil {
		toSerialize["case_field"] = o.CaseField
	}
	if o.IncidentUserDefinedFieldId != nil {
		toSerialize["incident_user_defined_field_id"] = o.IncidentUserDefinedFieldId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationIncidentFieldMappingsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaseField                  *string `json:"case_field,omitempty"`
		IncidentUserDefinedFieldId *string `json:"incident_user_defined_field_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"case_field", "incident_user_defined_field_id"})
	} else {
		return err
	}
	o.CaseField = all.CaseField
	o.IncidentUserDefinedFieldId = all.IncidentUserDefinedFieldId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
