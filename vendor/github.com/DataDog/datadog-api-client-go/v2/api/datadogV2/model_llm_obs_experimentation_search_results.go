// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationSearchResults The matching experimentation entities grouped by type.
type LLMObsExperimentationSearchResults struct {
	// Matching dataset records. Present when `dataset_records` is included in `filter.scope`.
	DatasetRecords []LLMObsDatasetRecordDataResponse `json:"dataset_records,omitempty"`
	// Matching datasets. Present when `datasets` is included in `filter.scope`.
	Datasets []LLMObsDatasetDataResponse `json:"datasets,omitempty"`
	// Matching experiment runs. Present when `experiment_runs` is included in `filter.scope`.
	ExperimentRuns []LLMObsExperimentRunDataResponse `json:"experiment_runs,omitempty"`
	// Matching experiments. Present when `experiments` is included in `filter.scope`.
	Experiments []LLMObsExperimentDataAttributesResponse `json:"experiments,omitempty"`
	// Matching projects. Present when `projects` is included in `filter.scope`.
	Projects []LLMObsProjectDataResponse `json:"projects,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationSearchResults instantiates a new LLMObsExperimentationSearchResults object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationSearchResults() *LLMObsExperimentationSearchResults {
	this := LLMObsExperimentationSearchResults{}
	return &this
}

// NewLLMObsExperimentationSearchResultsWithDefaults instantiates a new LLMObsExperimentationSearchResults object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationSearchResultsWithDefaults() *LLMObsExperimentationSearchResults {
	this := LLMObsExperimentationSearchResults{}
	return &this
}

// GetDatasetRecords returns the DatasetRecords field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentationSearchResults) GetDatasetRecords() []LLMObsDatasetRecordDataResponse {
	if o == nil {
		var ret []LLMObsDatasetRecordDataResponse
		return ret
	}
	return o.DatasetRecords
}

// GetDatasetRecordsOk returns a tuple with the DatasetRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentationSearchResults) GetDatasetRecordsOk() (*[]LLMObsDatasetRecordDataResponse, bool) {
	if o == nil || o.DatasetRecords == nil {
		return nil, false
	}
	return &o.DatasetRecords, true
}

// HasDatasetRecords returns a boolean if a field has been set.
func (o *LLMObsExperimentationSearchResults) HasDatasetRecords() bool {
	return o != nil && o.DatasetRecords != nil
}

// SetDatasetRecords gets a reference to the given []LLMObsDatasetRecordDataResponse and assigns it to the DatasetRecords field.
func (o *LLMObsExperimentationSearchResults) SetDatasetRecords(v []LLMObsDatasetRecordDataResponse) {
	o.DatasetRecords = v
}

// GetDatasets returns the Datasets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentationSearchResults) GetDatasets() []LLMObsDatasetDataResponse {
	if o == nil {
		var ret []LLMObsDatasetDataResponse
		return ret
	}
	return o.Datasets
}

// GetDatasetsOk returns a tuple with the Datasets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentationSearchResults) GetDatasetsOk() (*[]LLMObsDatasetDataResponse, bool) {
	if o == nil || o.Datasets == nil {
		return nil, false
	}
	return &o.Datasets, true
}

// HasDatasets returns a boolean if a field has been set.
func (o *LLMObsExperimentationSearchResults) HasDatasets() bool {
	return o != nil && o.Datasets != nil
}

// SetDatasets gets a reference to the given []LLMObsDatasetDataResponse and assigns it to the Datasets field.
func (o *LLMObsExperimentationSearchResults) SetDatasets(v []LLMObsDatasetDataResponse) {
	o.Datasets = v
}

// GetExperimentRuns returns the ExperimentRuns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentationSearchResults) GetExperimentRuns() []LLMObsExperimentRunDataResponse {
	if o == nil {
		var ret []LLMObsExperimentRunDataResponse
		return ret
	}
	return o.ExperimentRuns
}

// GetExperimentRunsOk returns a tuple with the ExperimentRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentationSearchResults) GetExperimentRunsOk() (*[]LLMObsExperimentRunDataResponse, bool) {
	if o == nil || o.ExperimentRuns == nil {
		return nil, false
	}
	return &o.ExperimentRuns, true
}

// HasExperimentRuns returns a boolean if a field has been set.
func (o *LLMObsExperimentationSearchResults) HasExperimentRuns() bool {
	return o != nil && o.ExperimentRuns != nil
}

// SetExperimentRuns gets a reference to the given []LLMObsExperimentRunDataResponse and assigns it to the ExperimentRuns field.
func (o *LLMObsExperimentationSearchResults) SetExperimentRuns(v []LLMObsExperimentRunDataResponse) {
	o.ExperimentRuns = v
}

// GetExperiments returns the Experiments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentationSearchResults) GetExperiments() []LLMObsExperimentDataAttributesResponse {
	if o == nil {
		var ret []LLMObsExperimentDataAttributesResponse
		return ret
	}
	return o.Experiments
}

// GetExperimentsOk returns a tuple with the Experiments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentationSearchResults) GetExperimentsOk() (*[]LLMObsExperimentDataAttributesResponse, bool) {
	if o == nil || o.Experiments == nil {
		return nil, false
	}
	return &o.Experiments, true
}

// HasExperiments returns a boolean if a field has been set.
func (o *LLMObsExperimentationSearchResults) HasExperiments() bool {
	return o != nil && o.Experiments != nil
}

// SetExperiments gets a reference to the given []LLMObsExperimentDataAttributesResponse and assigns it to the Experiments field.
func (o *LLMObsExperimentationSearchResults) SetExperiments(v []LLMObsExperimentDataAttributesResponse) {
	o.Experiments = v
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsExperimentationSearchResults) GetProjects() []LLMObsProjectDataResponse {
	if o == nil {
		var ret []LLMObsProjectDataResponse
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsExperimentationSearchResults) GetProjectsOk() (*[]LLMObsProjectDataResponse, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return &o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *LLMObsExperimentationSearchResults) HasProjects() bool {
	return o != nil && o.Projects != nil
}

// SetProjects gets a reference to the given []LLMObsProjectDataResponse and assigns it to the Projects field.
func (o *LLMObsExperimentationSearchResults) SetProjects(v []LLMObsProjectDataResponse) {
	o.Projects = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationSearchResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DatasetRecords != nil {
		toSerialize["dataset_records"] = o.DatasetRecords
	}
	if o.Datasets != nil {
		toSerialize["datasets"] = o.Datasets
	}
	if o.ExperimentRuns != nil {
		toSerialize["experiment_runs"] = o.ExperimentRuns
	}
	if o.Experiments != nil {
		toSerialize["experiments"] = o.Experiments
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationSearchResults) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetRecords []LLMObsDatasetRecordDataResponse        `json:"dataset_records,omitempty"`
		Datasets       []LLMObsDatasetDataResponse              `json:"datasets,omitempty"`
		ExperimentRuns []LLMObsExperimentRunDataResponse        `json:"experiment_runs,omitempty"`
		Experiments    []LLMObsExperimentDataAttributesResponse `json:"experiments,omitempty"`
		Projects       []LLMObsProjectDataResponse              `json:"projects,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dataset_records", "datasets", "experiment_runs", "experiments", "projects"})
	} else {
		return err
	}
	o.DatasetRecords = all.DatasetRecords
	o.Datasets = all.Datasets
	o.ExperimentRuns = all.ExperimentRuns
	o.Experiments = all.Experiments
	o.Projects = all.Projects

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
