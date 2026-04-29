// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionSearch Search configuration for retention queries.
type RetentionSearch struct {
	// Cohort criteria for retention queries.
	CohortCriteria RetentionCohortCriteria `json:"cohort_criteria"`
	// Filters for retention queries.
	Filters *RetentionFilters `json:"filters,omitempty"`
	// Entity to track for retention.
	RetentionEntity RetentionEntity `json:"retention_entity"`
	// Condition for counting user return.
	ReturnCondition RetentionReturnCondition `json:"return_condition"`
	// Return criteria for retention queries.
	ReturnCriteria *RetentionReturnCriteria `json:"return_criteria,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewRetentionSearch instantiates a new RetentionSearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionSearch(cohortCriteria RetentionCohortCriteria, retentionEntity RetentionEntity, returnCondition RetentionReturnCondition) *RetentionSearch {
	this := RetentionSearch{}
	this.CohortCriteria = cohortCriteria
	this.RetentionEntity = retentionEntity
	this.ReturnCondition = returnCondition
	return &this
}

// NewRetentionSearchWithDefaults instantiates a new RetentionSearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionSearchWithDefaults() *RetentionSearch {
	this := RetentionSearch{}
	return &this
}

// GetCohortCriteria returns the CohortCriteria field value.
func (o *RetentionSearch) GetCohortCriteria() RetentionCohortCriteria {
	if o == nil {
		var ret RetentionCohortCriteria
		return ret
	}
	return o.CohortCriteria
}

// GetCohortCriteriaOk returns a tuple with the CohortCriteria field value
// and a boolean to check if the value has been set.
func (o *RetentionSearch) GetCohortCriteriaOk() (*RetentionCohortCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CohortCriteria, true
}

// SetCohortCriteria sets field value.
func (o *RetentionSearch) SetCohortCriteria(v RetentionCohortCriteria) {
	o.CohortCriteria = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *RetentionSearch) GetFilters() RetentionFilters {
	if o == nil || o.Filters == nil {
		var ret RetentionFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionSearch) GetFiltersOk() (*RetentionFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *RetentionSearch) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given RetentionFilters and assigns it to the Filters field.
func (o *RetentionSearch) SetFilters(v RetentionFilters) {
	o.Filters = &v
}

// GetRetentionEntity returns the RetentionEntity field value.
func (o *RetentionSearch) GetRetentionEntity() RetentionEntity {
	if o == nil {
		var ret RetentionEntity
		return ret
	}
	return o.RetentionEntity
}

// GetRetentionEntityOk returns a tuple with the RetentionEntity field value
// and a boolean to check if the value has been set.
func (o *RetentionSearch) GetRetentionEntityOk() (*RetentionEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionEntity, true
}

// SetRetentionEntity sets field value.
func (o *RetentionSearch) SetRetentionEntity(v RetentionEntity) {
	o.RetentionEntity = v
}

// GetReturnCondition returns the ReturnCondition field value.
func (o *RetentionSearch) GetReturnCondition() RetentionReturnCondition {
	if o == nil {
		var ret RetentionReturnCondition
		return ret
	}
	return o.ReturnCondition
}

// GetReturnConditionOk returns a tuple with the ReturnCondition field value
// and a boolean to check if the value has been set.
func (o *RetentionSearch) GetReturnConditionOk() (*RetentionReturnCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnCondition, true
}

// SetReturnCondition sets field value.
func (o *RetentionSearch) SetReturnCondition(v RetentionReturnCondition) {
	o.ReturnCondition = v
}

// GetReturnCriteria returns the ReturnCriteria field value if set, zero value otherwise.
func (o *RetentionSearch) GetReturnCriteria() RetentionReturnCriteria {
	if o == nil || o.ReturnCriteria == nil {
		var ret RetentionReturnCriteria
		return ret
	}
	return *o.ReturnCriteria
}

// GetReturnCriteriaOk returns a tuple with the ReturnCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetentionSearch) GetReturnCriteriaOk() (*RetentionReturnCriteria, bool) {
	if o == nil || o.ReturnCriteria == nil {
		return nil, false
	}
	return o.ReturnCriteria, true
}

// HasReturnCriteria returns a boolean if a field has been set.
func (o *RetentionSearch) HasReturnCriteria() bool {
	return o != nil && o.ReturnCriteria != nil
}

// SetReturnCriteria gets a reference to the given RetentionReturnCriteria and assigns it to the ReturnCriteria field.
func (o *RetentionSearch) SetReturnCriteria(v RetentionReturnCriteria) {
	o.ReturnCriteria = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cohort_criteria"] = o.CohortCriteria
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["retention_entity"] = o.RetentionEntity
	toSerialize["return_condition"] = o.ReturnCondition
	if o.ReturnCriteria != nil {
		toSerialize["return_criteria"] = o.ReturnCriteria
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionSearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CohortCriteria  *RetentionCohortCriteria  `json:"cohort_criteria"`
		Filters         *RetentionFilters         `json:"filters,omitempty"`
		RetentionEntity *RetentionEntity          `json:"retention_entity"`
		ReturnCondition *RetentionReturnCondition `json:"return_condition"`
		ReturnCriteria  *RetentionReturnCriteria  `json:"return_criteria,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CohortCriteria == nil {
		return fmt.Errorf("required field cohort_criteria missing")
	}
	if all.RetentionEntity == nil {
		return fmt.Errorf("required field retention_entity missing")
	}
	if all.ReturnCondition == nil {
		return fmt.Errorf("required field return_condition missing")
	}

	hasInvalidField := false
	if all.CohortCriteria.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CohortCriteria = *all.CohortCriteria
	if all.Filters != nil && all.Filters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filters = all.Filters
	if !all.RetentionEntity.IsValid() {
		hasInvalidField = true
	} else {
		o.RetentionEntity = *all.RetentionEntity
	}
	if !all.ReturnCondition.IsValid() {
		hasInvalidField = true
	} else {
		o.ReturnCondition = *all.ReturnCondition
	}
	if all.ReturnCriteria != nil && all.ReturnCriteria.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReturnCriteria = all.ReturnCriteria

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
