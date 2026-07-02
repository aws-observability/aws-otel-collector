// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiCustomRuleItem An AI custom rule embedded within a ruleset response.
type AiCustomRuleItem struct {
	// The creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The identifier of the user who created the rule.
	CreatedBy string `json:"created_by"`
	// Response attributes of an AI custom rule revision.
	LastRevision AiCustomRuleRevisionResponseAttributes `json:"last_revision"`
	// The rule name.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiCustomRuleItem instantiates a new AiCustomRuleItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiCustomRuleItem(createdAt time.Time, createdBy string, lastRevision AiCustomRuleRevisionResponseAttributes, name string) *AiCustomRuleItem {
	this := AiCustomRuleItem{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.LastRevision = lastRevision
	this.Name = name
	return &this
}

// NewAiCustomRuleItemWithDefaults instantiates a new AiCustomRuleItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiCustomRuleItemWithDefaults() *AiCustomRuleItem {
	this := AiCustomRuleItem{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiCustomRuleItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiCustomRuleItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *AiCustomRuleItem) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleItem) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *AiCustomRuleItem) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetLastRevision returns the LastRevision field value.
func (o *AiCustomRuleItem) GetLastRevision() AiCustomRuleRevisionResponseAttributes {
	if o == nil {
		var ret AiCustomRuleRevisionResponseAttributes
		return ret
	}
	return o.LastRevision
}

// GetLastRevisionOk returns a tuple with the LastRevision field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleItem) GetLastRevisionOk() (*AiCustomRuleRevisionResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastRevision, true
}

// SetLastRevision sets field value.
func (o *AiCustomRuleItem) SetLastRevision(v AiCustomRuleRevisionResponseAttributes) {
	o.LastRevision = v
}

// GetName returns the Name field value.
func (o *AiCustomRuleItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiCustomRuleItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiCustomRuleItem) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiCustomRuleItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["last_revision"] = o.LastRevision
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiCustomRuleItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt    *time.Time                              `json:"created_at"`
		CreatedBy    *string                                 `json:"created_by"`
		LastRevision *AiCustomRuleRevisionResponseAttributes `json:"last_revision"`
		Name         *string                                 `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.LastRevision == nil {
		return fmt.Errorf("required field last_revision missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "last_revision", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	if all.LastRevision.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastRevision = *all.LastRevision
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
