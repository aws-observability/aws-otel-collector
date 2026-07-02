// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSummaryAvailableFieldsAttributes The lists of field names returned by `GET /api/v1/usage/summary` at each
// of its three response levels. Each list contains every key the data endpoint
// emits—both typed fields declared in the OpenAPI spec and untyped keys
// exposed through `additionalProperties`.
type UsageSummaryAvailableFieldsAttributes struct {
	// Sorted list of every key returned inside each `UsageSummaryDate`
	// entry of `usage[]` (typed fields and `additionalProperties` keys
	// combined).
	DateFields []string `json:"date_fields,omitempty"`
	// Sorted list of every key returned inside each `UsageSummaryDateOrg`
	// entry of `usage[].orgs[]` (typed fields and `additionalProperties`
	// keys combined).
	DateOrgFields []string `json:"date_org_fields,omitempty"`
	// Sorted list of every key returned as a direct property of
	// `UsageSummaryResponse` (typed fields and `additionalProperties`
	// keys combined).
	ResponseFields []string `json:"response_fields,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageSummaryAvailableFieldsAttributes instantiates a new UsageSummaryAvailableFieldsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSummaryAvailableFieldsAttributes() *UsageSummaryAvailableFieldsAttributes {
	this := UsageSummaryAvailableFieldsAttributes{}
	return &this
}

// NewUsageSummaryAvailableFieldsAttributesWithDefaults instantiates a new UsageSummaryAvailableFieldsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSummaryAvailableFieldsAttributesWithDefaults() *UsageSummaryAvailableFieldsAttributes {
	this := UsageSummaryAvailableFieldsAttributes{}
	return &this
}

// GetDateFields returns the DateFields field value if set, zero value otherwise.
func (o *UsageSummaryAvailableFieldsAttributes) GetDateFields() []string {
	if o == nil || o.DateFields == nil {
		var ret []string
		return ret
	}
	return o.DateFields
}

// GetDateFieldsOk returns a tuple with the DateFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryAvailableFieldsAttributes) GetDateFieldsOk() (*[]string, bool) {
	if o == nil || o.DateFields == nil {
		return nil, false
	}
	return &o.DateFields, true
}

// HasDateFields returns a boolean if a field has been set.
func (o *UsageSummaryAvailableFieldsAttributes) HasDateFields() bool {
	return o != nil && o.DateFields != nil
}

// SetDateFields gets a reference to the given []string and assigns it to the DateFields field.
func (o *UsageSummaryAvailableFieldsAttributes) SetDateFields(v []string) {
	o.DateFields = v
}

// GetDateOrgFields returns the DateOrgFields field value if set, zero value otherwise.
func (o *UsageSummaryAvailableFieldsAttributes) GetDateOrgFields() []string {
	if o == nil || o.DateOrgFields == nil {
		var ret []string
		return ret
	}
	return o.DateOrgFields
}

// GetDateOrgFieldsOk returns a tuple with the DateOrgFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryAvailableFieldsAttributes) GetDateOrgFieldsOk() (*[]string, bool) {
	if o == nil || o.DateOrgFields == nil {
		return nil, false
	}
	return &o.DateOrgFields, true
}

// HasDateOrgFields returns a boolean if a field has been set.
func (o *UsageSummaryAvailableFieldsAttributes) HasDateOrgFields() bool {
	return o != nil && o.DateOrgFields != nil
}

// SetDateOrgFields gets a reference to the given []string and assigns it to the DateOrgFields field.
func (o *UsageSummaryAvailableFieldsAttributes) SetDateOrgFields(v []string) {
	o.DateOrgFields = v
}

// GetResponseFields returns the ResponseFields field value if set, zero value otherwise.
func (o *UsageSummaryAvailableFieldsAttributes) GetResponseFields() []string {
	if o == nil || o.ResponseFields == nil {
		var ret []string
		return ret
	}
	return o.ResponseFields
}

// GetResponseFieldsOk returns a tuple with the ResponseFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryAvailableFieldsAttributes) GetResponseFieldsOk() (*[]string, bool) {
	if o == nil || o.ResponseFields == nil {
		return nil, false
	}
	return &o.ResponseFields, true
}

// HasResponseFields returns a boolean if a field has been set.
func (o *UsageSummaryAvailableFieldsAttributes) HasResponseFields() bool {
	return o != nil && o.ResponseFields != nil
}

// SetResponseFields gets a reference to the given []string and assigns it to the ResponseFields field.
func (o *UsageSummaryAvailableFieldsAttributes) SetResponseFields(v []string) {
	o.ResponseFields = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSummaryAvailableFieldsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DateFields != nil {
		toSerialize["date_fields"] = o.DateFields
	}
	if o.DateOrgFields != nil {
		toSerialize["date_org_fields"] = o.DateOrgFields
	}
	if o.ResponseFields != nil {
		toSerialize["response_fields"] = o.ResponseFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageSummaryAvailableFieldsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DateFields     []string `json:"date_fields,omitempty"`
		DateOrgFields  []string `json:"date_org_fields,omitempty"`
		ResponseFields []string `json:"response_fields,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"date_fields", "date_org_fields", "response_fields"})
	} else {
		return err
	}
	o.DateFields = all.DateFields
	o.DateOrgFields = all.DateOrgFields
	o.ResponseFields = all.ResponseFields

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
