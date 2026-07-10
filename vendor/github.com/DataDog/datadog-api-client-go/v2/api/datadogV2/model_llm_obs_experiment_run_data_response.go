// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentRunDataResponse Data object for an LLM Observability experiment run.
type LLMObsExperimentRunDataResponse struct {
	// Aggregated metric data for this run.
	AggregateData map[string]interface{} `json:"aggregate_data,omitempty"`
	// Timestamp when the run was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Identifier of the experiment this run belongs to.
	ExperimentId *string `json:"experiment_id,omitempty"`
	// Unique identifier of the experiment run.
	Id *string `json:"id,omitempty"`
	// Sequential number of this run within the experiment.
	RunNumber *int32 `json:"run_number,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentRunDataResponse instantiates a new LLMObsExperimentRunDataResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentRunDataResponse() *LLMObsExperimentRunDataResponse {
	this := LLMObsExperimentRunDataResponse{}
	return &this
}

// NewLLMObsExperimentRunDataResponseWithDefaults instantiates a new LLMObsExperimentRunDataResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentRunDataResponseWithDefaults() *LLMObsExperimentRunDataResponse {
	this := LLMObsExperimentRunDataResponse{}
	return &this
}

// GetAggregateData returns the AggregateData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentRunDataResponse) GetAggregateData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AggregateData
}

// GetAggregateDataOk returns a tuple with the AggregateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentRunDataResponse) GetAggregateDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.AggregateData == nil {
		return nil, false
	}
	return &o.AggregateData, true
}

// HasAggregateData returns a boolean if a field has been set.
func (o *LLMObsExperimentRunDataResponse) HasAggregateData() bool {
	return o != nil && o.AggregateData != nil
}

// SetAggregateData gets a reference to the given map[string]interface{} and assigns it to the AggregateData field.
func (o *LLMObsExperimentRunDataResponse) SetAggregateData(v map[string]interface{}) {
	o.AggregateData = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LLMObsExperimentRunDataResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentRunDataResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LLMObsExperimentRunDataResponse) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LLMObsExperimentRunDataResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExperimentId returns the ExperimentId field value if set, zero value otherwise.
func (o *LLMObsExperimentRunDataResponse) GetExperimentId() string {
	if o == nil || o.ExperimentId == nil {
		var ret string
		return ret
	}
	return *o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentRunDataResponse) GetExperimentIdOk() (*string, bool) {
	if o == nil || o.ExperimentId == nil {
		return nil, false
	}
	return o.ExperimentId, true
}

// HasExperimentId returns a boolean if a field has been set.
func (o *LLMObsExperimentRunDataResponse) HasExperimentId() bool {
	return o != nil && o.ExperimentId != nil
}

// SetExperimentId gets a reference to the given string and assigns it to the ExperimentId field.
func (o *LLMObsExperimentRunDataResponse) SetExperimentId(v string) {
	o.ExperimentId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LLMObsExperimentRunDataResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentRunDataResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LLMObsExperimentRunDataResponse) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LLMObsExperimentRunDataResponse) SetId(v string) {
	o.Id = &v
}

// GetRunNumber returns the RunNumber field value if set, zero value otherwise.
func (o *LLMObsExperimentRunDataResponse) GetRunNumber() int32 {
	if o == nil || o.RunNumber == nil {
		var ret int32
		return ret
	}
	return *o.RunNumber
}

// GetRunNumberOk returns a tuple with the RunNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentRunDataResponse) GetRunNumberOk() (*int32, bool) {
	if o == nil || o.RunNumber == nil {
		return nil, false
	}
	return o.RunNumber, true
}

// HasRunNumber returns a boolean if a field has been set.
func (o *LLMObsExperimentRunDataResponse) HasRunNumber() bool {
	return o != nil && o.RunNumber != nil
}

// SetRunNumber gets a reference to the given int32 and assigns it to the RunNumber field.
func (o *LLMObsExperimentRunDataResponse) SetRunNumber(v int32) {
	o.RunNumber = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentRunDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AggregateData != nil {
		toSerialize["aggregate_data"] = o.AggregateData
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ExperimentId != nil {
		toSerialize["experiment_id"] = o.ExperimentId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RunNumber != nil {
		toSerialize["run_number"] = o.RunNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentRunDataResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregateData map[string]interface{} `json:"aggregate_data,omitempty"`
		CreatedAt     *time.Time             `json:"created_at,omitempty"`
		ExperimentId  *string                `json:"experiment_id,omitempty"`
		Id            *string                `json:"id,omitempty"`
		RunNumber     *int32                 `json:"run_number,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregate_data", "created_at", "experiment_id", "id", "run_number"})
	} else {
		return err
	}
	o.AggregateData = all.AggregateData
	o.CreatedAt = all.CreatedAt
	o.ExperimentId = all.ExperimentId
	o.Id = all.Id
	o.RunNumber = all.RunNumber

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
