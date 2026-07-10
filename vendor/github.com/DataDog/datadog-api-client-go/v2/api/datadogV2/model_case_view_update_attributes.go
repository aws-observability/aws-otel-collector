// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseViewUpdateAttributes Attributes that can be updated on a case view. All fields are optional; only provided fields are changed.
type CaseViewUpdateAttributes struct {
	// The name of the view.
	Name *string `json:"name,omitempty"`
	// The identifier of a notification rule linked to this view. When set, users subscribed to the view receive alerts for matching cases.
	NpRuleId *string `json:"np_rule_id,omitempty"`
	// The query used to filter cases in this view.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseViewUpdateAttributes instantiates a new CaseViewUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseViewUpdateAttributes() *CaseViewUpdateAttributes {
	this := CaseViewUpdateAttributes{}
	return &this
}

// NewCaseViewUpdateAttributesWithDefaults instantiates a new CaseViewUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseViewUpdateAttributesWithDefaults() *CaseViewUpdateAttributes {
	this := CaseViewUpdateAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CaseViewUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseViewUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CaseViewUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CaseViewUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetNpRuleId returns the NpRuleId field value if set, zero value otherwise.
func (o *CaseViewUpdateAttributes) GetNpRuleId() string {
	if o == nil || o.NpRuleId == nil {
		var ret string
		return ret
	}
	return *o.NpRuleId
}

// GetNpRuleIdOk returns a tuple with the NpRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseViewUpdateAttributes) GetNpRuleIdOk() (*string, bool) {
	if o == nil || o.NpRuleId == nil {
		return nil, false
	}
	return o.NpRuleId, true
}

// HasNpRuleId returns a boolean if a field has been set.
func (o *CaseViewUpdateAttributes) HasNpRuleId() bool {
	return o != nil && o.NpRuleId != nil
}

// SetNpRuleId gets a reference to the given string and assigns it to the NpRuleId field.
func (o *CaseViewUpdateAttributes) SetNpRuleId(v string) {
	o.NpRuleId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CaseViewUpdateAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseViewUpdateAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CaseViewUpdateAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *CaseViewUpdateAttributes) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseViewUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NpRuleId != nil {
		toSerialize["np_rule_id"] = o.NpRuleId
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseViewUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string `json:"name,omitempty"`
		NpRuleId *string `json:"np_rule_id,omitempty"`
		Query    *string `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "np_rule_id", "query"})
	} else {
		return err
	}
	o.Name = all.Name
	o.NpRuleId = all.NpRuleId
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
