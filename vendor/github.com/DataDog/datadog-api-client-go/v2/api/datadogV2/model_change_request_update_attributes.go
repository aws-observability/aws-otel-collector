// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestUpdateAttributes Attributes for updating a change request.
type ChangeRequestUpdateAttributes struct {
	// The plan associated with the change request.
	ChangeRequestPlan *string `json:"change_request_plan,omitempty"`
	// The risk level of the change request.
	ChangeRequestRisk *ChangeRequestRiskLevel `json:"change_request_risk,omitempty"`
	// The type of the change request.
	ChangeRequestType *ChangeRequestChangeType `json:"change_request_type,omitempty"`
	// The planned end date of the change request.
	EndDate *time.Time `json:"end_date,omitempty"`
	// The identifier of the change request to update.
	Id *string `json:"id,omitempty"`
	// The planned start date of the change request.
	StartDate *time.Time `json:"start_date,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestUpdateAttributes instantiates a new ChangeRequestUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestUpdateAttributes() *ChangeRequestUpdateAttributes {
	this := ChangeRequestUpdateAttributes{}
	return &this
}

// NewChangeRequestUpdateAttributesWithDefaults instantiates a new ChangeRequestUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestUpdateAttributesWithDefaults() *ChangeRequestUpdateAttributes {
	this := ChangeRequestUpdateAttributes{}
	return &this
}

// GetChangeRequestPlan returns the ChangeRequestPlan field value if set, zero value otherwise.
func (o *ChangeRequestUpdateAttributes) GetChangeRequestPlan() string {
	if o == nil || o.ChangeRequestPlan == nil {
		var ret string
		return ret
	}
	return *o.ChangeRequestPlan
}

// GetChangeRequestPlanOk returns a tuple with the ChangeRequestPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateAttributes) GetChangeRequestPlanOk() (*string, bool) {
	if o == nil || o.ChangeRequestPlan == nil {
		return nil, false
	}
	return o.ChangeRequestPlan, true
}

// HasChangeRequestPlan returns a boolean if a field has been set.
func (o *ChangeRequestUpdateAttributes) HasChangeRequestPlan() bool {
	return o != nil && o.ChangeRequestPlan != nil
}

// SetChangeRequestPlan gets a reference to the given string and assigns it to the ChangeRequestPlan field.
func (o *ChangeRequestUpdateAttributes) SetChangeRequestPlan(v string) {
	o.ChangeRequestPlan = &v
}

// GetChangeRequestRisk returns the ChangeRequestRisk field value if set, zero value otherwise.
func (o *ChangeRequestUpdateAttributes) GetChangeRequestRisk() ChangeRequestRiskLevel {
	if o == nil || o.ChangeRequestRisk == nil {
		var ret ChangeRequestRiskLevel
		return ret
	}
	return *o.ChangeRequestRisk
}

// GetChangeRequestRiskOk returns a tuple with the ChangeRequestRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateAttributes) GetChangeRequestRiskOk() (*ChangeRequestRiskLevel, bool) {
	if o == nil || o.ChangeRequestRisk == nil {
		return nil, false
	}
	return o.ChangeRequestRisk, true
}

// HasChangeRequestRisk returns a boolean if a field has been set.
func (o *ChangeRequestUpdateAttributes) HasChangeRequestRisk() bool {
	return o != nil && o.ChangeRequestRisk != nil
}

// SetChangeRequestRisk gets a reference to the given ChangeRequestRiskLevel and assigns it to the ChangeRequestRisk field.
func (o *ChangeRequestUpdateAttributes) SetChangeRequestRisk(v ChangeRequestRiskLevel) {
	o.ChangeRequestRisk = &v
}

// GetChangeRequestType returns the ChangeRequestType field value if set, zero value otherwise.
func (o *ChangeRequestUpdateAttributes) GetChangeRequestType() ChangeRequestChangeType {
	if o == nil || o.ChangeRequestType == nil {
		var ret ChangeRequestChangeType
		return ret
	}
	return *o.ChangeRequestType
}

// GetChangeRequestTypeOk returns a tuple with the ChangeRequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateAttributes) GetChangeRequestTypeOk() (*ChangeRequestChangeType, bool) {
	if o == nil || o.ChangeRequestType == nil {
		return nil, false
	}
	return o.ChangeRequestType, true
}

// HasChangeRequestType returns a boolean if a field has been set.
func (o *ChangeRequestUpdateAttributes) HasChangeRequestType() bool {
	return o != nil && o.ChangeRequestType != nil
}

// SetChangeRequestType gets a reference to the given ChangeRequestChangeType and assigns it to the ChangeRequestType field.
func (o *ChangeRequestUpdateAttributes) SetChangeRequestType(v ChangeRequestChangeType) {
	o.ChangeRequestType = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ChangeRequestUpdateAttributes) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateAttributes) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ChangeRequestUpdateAttributes) HasEndDate() bool {
	return o != nil && o.EndDate != nil
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ChangeRequestUpdateAttributes) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChangeRequestUpdateAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChangeRequestUpdateAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChangeRequestUpdateAttributes) SetId(v string) {
	o.Id = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ChangeRequestUpdateAttributes) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateAttributes) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ChangeRequestUpdateAttributes) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ChangeRequestUpdateAttributes) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChangeRequestPlan != nil {
		toSerialize["change_request_plan"] = o.ChangeRequestPlan
	}
	if o.ChangeRequestRisk != nil {
		toSerialize["change_request_risk"] = o.ChangeRequestRisk
	}
	if o.ChangeRequestType != nil {
		toSerialize["change_request_type"] = o.ChangeRequestType
	}
	if o.EndDate != nil {
		if o.EndDate.Nanosecond() == 0 {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["end_date"] = o.EndDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.StartDate != nil {
		if o.StartDate.Nanosecond() == 0 {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["start_date"] = o.StartDate.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeRequestPlan *string                  `json:"change_request_plan,omitempty"`
		ChangeRequestRisk *ChangeRequestRiskLevel  `json:"change_request_risk,omitempty"`
		ChangeRequestType *ChangeRequestChangeType `json:"change_request_type,omitempty"`
		EndDate           *time.Time               `json:"end_date,omitempty"`
		Id                *string                  `json:"id,omitempty"`
		StartDate         *time.Time               `json:"start_date,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_request_plan", "change_request_risk", "change_request_type", "end_date", "id", "start_date"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ChangeRequestPlan = all.ChangeRequestPlan
	if all.ChangeRequestRisk != nil && !all.ChangeRequestRisk.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeRequestRisk = all.ChangeRequestRisk
	}
	if all.ChangeRequestType != nil && !all.ChangeRequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeRequestType = all.ChangeRequestType
	}
	o.EndDate = all.EndDate
	o.Id = all.Id
	o.StartDate = all.StartDate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
