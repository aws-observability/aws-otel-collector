// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutcomesBatchResponseAttributes The JSON:API attributes for an outcome.
type OutcomesBatchResponseAttributes struct {
	// Creation time of the rule outcome.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Time of last rule outcome modification.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Any remarks regarding the scorecard rule's evaluation, and supports HTML hyperlinks.
	Remarks *string `json:"remarks,omitempty"`
	// The unique name for a service in the catalog.
	ServiceName *string `json:"service_name,omitempty"`
	// The state of the rule evaluation.
	State *State `json:"state,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOutcomesBatchResponseAttributes instantiates a new OutcomesBatchResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutcomesBatchResponseAttributes() *OutcomesBatchResponseAttributes {
	this := OutcomesBatchResponseAttributes{}
	return &this
}

// NewOutcomesBatchResponseAttributesWithDefaults instantiates a new OutcomesBatchResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutcomesBatchResponseAttributesWithDefaults() *OutcomesBatchResponseAttributes {
	this := OutcomesBatchResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OutcomesBatchResponseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesBatchResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OutcomesBatchResponseAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OutcomesBatchResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *OutcomesBatchResponseAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesBatchResponseAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *OutcomesBatchResponseAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *OutcomesBatchResponseAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *OutcomesBatchResponseAttributes) GetRemarks() string {
	if o == nil || o.Remarks == nil {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesBatchResponseAttributes) GetRemarksOk() (*string, bool) {
	if o == nil || o.Remarks == nil {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *OutcomesBatchResponseAttributes) HasRemarks() bool {
	return o != nil && o.Remarks != nil
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *OutcomesBatchResponseAttributes) SetRemarks(v string) {
	o.Remarks = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *OutcomesBatchResponseAttributes) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesBatchResponseAttributes) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *OutcomesBatchResponseAttributes) HasServiceName() bool {
	return o != nil && o.ServiceName != nil
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *OutcomesBatchResponseAttributes) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OutcomesBatchResponseAttributes) GetState() State {
	if o == nil || o.State == nil {
		var ret State
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesBatchResponseAttributes) GetStateOk() (*State, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OutcomesBatchResponseAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given State and assigns it to the State field.
func (o *OutcomesBatchResponseAttributes) SetState(v State) {
	o.State = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutcomesBatchResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Remarks != nil {
		toSerialize["remarks"] = o.Remarks
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OutcomesBatchResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time `json:"created_at,omitempty"`
		ModifiedAt  *time.Time `json:"modified_at,omitempty"`
		Remarks     *string    `json:"remarks,omitempty"`
		ServiceName *string    `json:"service_name,omitempty"`
		State       *State     `json:"state,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "modified_at", "remarks", "service_name", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.ModifiedAt = all.ModifiedAt
	o.Remarks = all.Remarks
	o.ServiceName = all.ServiceName
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
