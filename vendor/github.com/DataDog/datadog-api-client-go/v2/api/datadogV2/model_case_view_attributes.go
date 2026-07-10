// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseViewAttributes Attributes of a case view, including the filter query and optional notification rule.
type CaseViewAttributes struct {
	// Timestamp when the view was created.
	CreatedAt time.Time `json:"created_at"`
	// Timestamp when the view was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// A human-readable name for the view, displayed in the Case Management UI.
	Name string `json:"name"`
	// The identifier of a notification rule linked to this view. When set, users subscribed to the view receive alerts for matching cases.
	NpRuleId *string `json:"np_rule_id,omitempty"`
	// The search query that determines which cases appear in this view. Uses the same syntax as the Case Management search bar (for example, `status:open priority:P1`).
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseViewAttributes instantiates a new CaseViewAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseViewAttributes(createdAt time.Time, name string, query string) *CaseViewAttributes {
	this := CaseViewAttributes{}
	this.CreatedAt = createdAt
	this.Name = name
	this.Query = query
	return &this
}

// NewCaseViewAttributesWithDefaults instantiates a new CaseViewAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseViewAttributesWithDefaults() *CaseViewAttributes {
	this := CaseViewAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CaseViewAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CaseViewAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CaseViewAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *CaseViewAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseViewAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *CaseViewAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *CaseViewAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value.
func (o *CaseViewAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CaseViewAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CaseViewAttributes) SetName(v string) {
	o.Name = v
}

// GetNpRuleId returns the NpRuleId field value if set, zero value otherwise.
func (o *CaseViewAttributes) GetNpRuleId() string {
	if o == nil || o.NpRuleId == nil {
		var ret string
		return ret
	}
	return *o.NpRuleId
}

// GetNpRuleIdOk returns a tuple with the NpRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseViewAttributes) GetNpRuleIdOk() (*string, bool) {
	if o == nil || o.NpRuleId == nil {
		return nil, false
	}
	return o.NpRuleId, true
}

// HasNpRuleId returns a boolean if a field has been set.
func (o *CaseViewAttributes) HasNpRuleId() bool {
	return o != nil && o.NpRuleId != nil
}

// SetNpRuleId gets a reference to the given string and assigns it to the NpRuleId field.
func (o *CaseViewAttributes) SetNpRuleId(v string) {
	o.NpRuleId = &v
}

// GetQuery returns the Query field value.
func (o *CaseViewAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CaseViewAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *CaseViewAttributes) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseViewAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["name"] = o.Name
	if o.NpRuleId != nil {
		toSerialize["np_rule_id"] = o.NpRuleId
	}
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseViewAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt  *time.Time `json:"created_at"`
		ModifiedAt *time.Time `json:"modified_at,omitempty"`
		Name       *string    `json:"name"`
		NpRuleId   *string    `json:"np_rule_id,omitempty"`
		Query      *string    `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "modified_at", "name", "np_rule_id", "query"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.ModifiedAt = all.ModifiedAt
	o.Name = *all.Name
	o.NpRuleId = all.NpRuleId
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
