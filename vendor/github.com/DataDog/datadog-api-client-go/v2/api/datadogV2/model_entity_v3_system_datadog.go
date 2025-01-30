// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3SystemDatadog Datadog product integrations for the service entity.
type EntityV3SystemDatadog struct {
	// Events associations.
	Events []EntityV3DatadogEventItem `json:"events,omitempty"`
	// Logs association.
	Logs []EntityV3DatadogLogItem `json:"logs,omitempty"`
	// Performance stats association.
	PerformanceData *EntityV3DatadogPerformance `json:"performanceData,omitempty"`
	// CI Pipelines association.
	Pipelines *EntityV3DatadogPipelines `json:"pipelines,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3SystemDatadog instantiates a new EntityV3SystemDatadog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3SystemDatadog() *EntityV3SystemDatadog {
	this := EntityV3SystemDatadog{}
	return &this
}

// NewEntityV3SystemDatadogWithDefaults instantiates a new EntityV3SystemDatadog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3SystemDatadogWithDefaults() *EntityV3SystemDatadog {
	this := EntityV3SystemDatadog{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *EntityV3SystemDatadog) GetEvents() []EntityV3DatadogEventItem {
	if o == nil || o.Events == nil {
		var ret []EntityV3DatadogEventItem
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3SystemDatadog) GetEventsOk() (*[]EntityV3DatadogEventItem, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *EntityV3SystemDatadog) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []EntityV3DatadogEventItem and assigns it to the Events field.
func (o *EntityV3SystemDatadog) SetEvents(v []EntityV3DatadogEventItem) {
	o.Events = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *EntityV3SystemDatadog) GetLogs() []EntityV3DatadogLogItem {
	if o == nil || o.Logs == nil {
		var ret []EntityV3DatadogLogItem
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3SystemDatadog) GetLogsOk() (*[]EntityV3DatadogLogItem, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return &o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *EntityV3SystemDatadog) HasLogs() bool {
	return o != nil && o.Logs != nil
}

// SetLogs gets a reference to the given []EntityV3DatadogLogItem and assigns it to the Logs field.
func (o *EntityV3SystemDatadog) SetLogs(v []EntityV3DatadogLogItem) {
	o.Logs = v
}

// GetPerformanceData returns the PerformanceData field value if set, zero value otherwise.
func (o *EntityV3SystemDatadog) GetPerformanceData() EntityV3DatadogPerformance {
	if o == nil || o.PerformanceData == nil {
		var ret EntityV3DatadogPerformance
		return ret
	}
	return *o.PerformanceData
}

// GetPerformanceDataOk returns a tuple with the PerformanceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3SystemDatadog) GetPerformanceDataOk() (*EntityV3DatadogPerformance, bool) {
	if o == nil || o.PerformanceData == nil {
		return nil, false
	}
	return o.PerformanceData, true
}

// HasPerformanceData returns a boolean if a field has been set.
func (o *EntityV3SystemDatadog) HasPerformanceData() bool {
	return o != nil && o.PerformanceData != nil
}

// SetPerformanceData gets a reference to the given EntityV3DatadogPerformance and assigns it to the PerformanceData field.
func (o *EntityV3SystemDatadog) SetPerformanceData(v EntityV3DatadogPerformance) {
	o.PerformanceData = &v
}

// GetPipelines returns the Pipelines field value if set, zero value otherwise.
func (o *EntityV3SystemDatadog) GetPipelines() EntityV3DatadogPipelines {
	if o == nil || o.Pipelines == nil {
		var ret EntityV3DatadogPipelines
		return ret
	}
	return *o.Pipelines
}

// GetPipelinesOk returns a tuple with the Pipelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3SystemDatadog) GetPipelinesOk() (*EntityV3DatadogPipelines, bool) {
	if o == nil || o.Pipelines == nil {
		return nil, false
	}
	return o.Pipelines, true
}

// HasPipelines returns a boolean if a field has been set.
func (o *EntityV3SystemDatadog) HasPipelines() bool {
	return o != nil && o.Pipelines != nil
}

// SetPipelines gets a reference to the given EntityV3DatadogPipelines and assigns it to the Pipelines field.
func (o *EntityV3SystemDatadog) SetPipelines(v EntityV3DatadogPipelines) {
	o.Pipelines = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3SystemDatadog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}
	if o.PerformanceData != nil {
		toSerialize["performanceData"] = o.PerformanceData
	}
	if o.Pipelines != nil {
		toSerialize["pipelines"] = o.Pipelines
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3SystemDatadog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Events          []EntityV3DatadogEventItem  `json:"events,omitempty"`
		Logs            []EntityV3DatadogLogItem    `json:"logs,omitempty"`
		PerformanceData *EntityV3DatadogPerformance `json:"performanceData,omitempty"`
		Pipelines       *EntityV3DatadogPipelines   `json:"pipelines,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	o.Events = all.Events
	o.Logs = all.Logs
	if all.PerformanceData != nil && all.PerformanceData.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PerformanceData = all.PerformanceData
	if all.Pipelines != nil && all.Pipelines.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pipelines = all.Pipelines

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
