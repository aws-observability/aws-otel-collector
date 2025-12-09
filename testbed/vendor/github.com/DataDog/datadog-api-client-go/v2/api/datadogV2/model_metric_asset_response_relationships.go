// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricAssetResponseRelationships Relationships to assets related to the metric.
type MetricAssetResponseRelationships struct {
	// An object containing the list of dashboards that can be referenced in the `included` data.
	Dashboards *MetricAssetDashboardRelationships `json:"dashboards,omitempty"`
	// A object containing the list of monitors that can be referenced in the `included` data.
	Monitors *MetricAssetMonitorRelationships `json:"monitors,omitempty"`
	// An object containing the list of notebooks that can be referenced in the `included` data.
	Notebooks *MetricAssetNotebookRelationships `json:"notebooks,omitempty"`
	// An object containing a list of SLOs that can be referenced in the `included` data.
	Slos *MetricAssetSLORelationships `json:"slos,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricAssetResponseRelationships instantiates a new MetricAssetResponseRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricAssetResponseRelationships() *MetricAssetResponseRelationships {
	this := MetricAssetResponseRelationships{}
	return &this
}

// NewMetricAssetResponseRelationshipsWithDefaults instantiates a new MetricAssetResponseRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricAssetResponseRelationshipsWithDefaults() *MetricAssetResponseRelationships {
	this := MetricAssetResponseRelationships{}
	return &this
}

// GetDashboards returns the Dashboards field value if set, zero value otherwise.
func (o *MetricAssetResponseRelationships) GetDashboards() MetricAssetDashboardRelationships {
	if o == nil || o.Dashboards == nil {
		var ret MetricAssetDashboardRelationships
		return ret
	}
	return *o.Dashboards
}

// GetDashboardsOk returns a tuple with the Dashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAssetResponseRelationships) GetDashboardsOk() (*MetricAssetDashboardRelationships, bool) {
	if o == nil || o.Dashboards == nil {
		return nil, false
	}
	return o.Dashboards, true
}

// HasDashboards returns a boolean if a field has been set.
func (o *MetricAssetResponseRelationships) HasDashboards() bool {
	return o != nil && o.Dashboards != nil
}

// SetDashboards gets a reference to the given MetricAssetDashboardRelationships and assigns it to the Dashboards field.
func (o *MetricAssetResponseRelationships) SetDashboards(v MetricAssetDashboardRelationships) {
	o.Dashboards = &v
}

// GetMonitors returns the Monitors field value if set, zero value otherwise.
func (o *MetricAssetResponseRelationships) GetMonitors() MetricAssetMonitorRelationships {
	if o == nil || o.Monitors == nil {
		var ret MetricAssetMonitorRelationships
		return ret
	}
	return *o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAssetResponseRelationships) GetMonitorsOk() (*MetricAssetMonitorRelationships, bool) {
	if o == nil || o.Monitors == nil {
		return nil, false
	}
	return o.Monitors, true
}

// HasMonitors returns a boolean if a field has been set.
func (o *MetricAssetResponseRelationships) HasMonitors() bool {
	return o != nil && o.Monitors != nil
}

// SetMonitors gets a reference to the given MetricAssetMonitorRelationships and assigns it to the Monitors field.
func (o *MetricAssetResponseRelationships) SetMonitors(v MetricAssetMonitorRelationships) {
	o.Monitors = &v
}

// GetNotebooks returns the Notebooks field value if set, zero value otherwise.
func (o *MetricAssetResponseRelationships) GetNotebooks() MetricAssetNotebookRelationships {
	if o == nil || o.Notebooks == nil {
		var ret MetricAssetNotebookRelationships
		return ret
	}
	return *o.Notebooks
}

// GetNotebooksOk returns a tuple with the Notebooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAssetResponseRelationships) GetNotebooksOk() (*MetricAssetNotebookRelationships, bool) {
	if o == nil || o.Notebooks == nil {
		return nil, false
	}
	return o.Notebooks, true
}

// HasNotebooks returns a boolean if a field has been set.
func (o *MetricAssetResponseRelationships) HasNotebooks() bool {
	return o != nil && o.Notebooks != nil
}

// SetNotebooks gets a reference to the given MetricAssetNotebookRelationships and assigns it to the Notebooks field.
func (o *MetricAssetResponseRelationships) SetNotebooks(v MetricAssetNotebookRelationships) {
	o.Notebooks = &v
}

// GetSlos returns the Slos field value if set, zero value otherwise.
func (o *MetricAssetResponseRelationships) GetSlos() MetricAssetSLORelationships {
	if o == nil || o.Slos == nil {
		var ret MetricAssetSLORelationships
		return ret
	}
	return *o.Slos
}

// GetSlosOk returns a tuple with the Slos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAssetResponseRelationships) GetSlosOk() (*MetricAssetSLORelationships, bool) {
	if o == nil || o.Slos == nil {
		return nil, false
	}
	return o.Slos, true
}

// HasSlos returns a boolean if a field has been set.
func (o *MetricAssetResponseRelationships) HasSlos() bool {
	return o != nil && o.Slos != nil
}

// SetSlos gets a reference to the given MetricAssetSLORelationships and assigns it to the Slos field.
func (o *MetricAssetResponseRelationships) SetSlos(v MetricAssetSLORelationships) {
	o.Slos = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricAssetResponseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Dashboards != nil {
		toSerialize["dashboards"] = o.Dashboards
	}
	if o.Monitors != nil {
		toSerialize["monitors"] = o.Monitors
	}
	if o.Notebooks != nil {
		toSerialize["notebooks"] = o.Notebooks
	}
	if o.Slos != nil {
		toSerialize["slos"] = o.Slos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricAssetResponseRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dashboards *MetricAssetDashboardRelationships `json:"dashboards,omitempty"`
		Monitors   *MetricAssetMonitorRelationships   `json:"monitors,omitempty"`
		Notebooks  *MetricAssetNotebookRelationships  `json:"notebooks,omitempty"`
		Slos       *MetricAssetSLORelationships       `json:"slos,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dashboards", "monitors", "notebooks", "slos"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Dashboards != nil && all.Dashboards.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dashboards = all.Dashboards
	if all.Monitors != nil && all.Monitors.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Monitors = all.Monitors
	if all.Notebooks != nil && all.Notebooks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Notebooks = all.Notebooks
	if all.Slos != nil && all.Slos.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Slos = all.Slos

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
