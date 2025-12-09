// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DowntimeRelationships All relationships associated with downtime.
type DowntimeRelationships struct {
	// The user who created the downtime.
	CreatedBy *DowntimeRelationshipsCreatedBy `json:"created_by,omitempty"`
	// The monitor identified by the downtime.
	Monitor *DowntimeRelationshipsMonitor `json:"monitor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDowntimeRelationships instantiates a new DowntimeRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeRelationships() *DowntimeRelationships {
	this := DowntimeRelationships{}
	return &this
}

// NewDowntimeRelationshipsWithDefaults instantiates a new DowntimeRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeRelationshipsWithDefaults() *DowntimeRelationships {
	this := DowntimeRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DowntimeRelationships) GetCreatedBy() DowntimeRelationshipsCreatedBy {
	if o == nil || o.CreatedBy == nil {
		var ret DowntimeRelationshipsCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRelationships) GetCreatedByOk() (*DowntimeRelationshipsCreatedBy, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DowntimeRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given DowntimeRelationshipsCreatedBy and assigns it to the CreatedBy field.
func (o *DowntimeRelationships) SetCreatedBy(v DowntimeRelationshipsCreatedBy) {
	o.CreatedBy = &v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *DowntimeRelationships) GetMonitor() DowntimeRelationshipsMonitor {
	if o == nil || o.Monitor == nil {
		var ret DowntimeRelationshipsMonitor
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowntimeRelationships) GetMonitorOk() (*DowntimeRelationshipsMonitor, bool) {
	if o == nil || o.Monitor == nil {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *DowntimeRelationships) HasMonitor() bool {
	return o != nil && o.Monitor != nil
}

// SetMonitor gets a reference to the given DowntimeRelationshipsMonitor and assigns it to the Monitor field.
func (o *DowntimeRelationships) SetMonitor(v DowntimeRelationshipsMonitor) {
	o.Monitor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Monitor != nil {
		toSerialize["monitor"] = o.Monitor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy *DowntimeRelationshipsCreatedBy `json:"created_by,omitempty"`
		Monitor   *DowntimeRelationshipsMonitor   `json:"monitor,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "monitor"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	if all.Monitor != nil && all.Monitor.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Monitor = all.Monitor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
