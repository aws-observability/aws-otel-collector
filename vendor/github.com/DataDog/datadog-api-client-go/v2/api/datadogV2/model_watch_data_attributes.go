// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WatchDataAttributes Attributes for recording a session watch event, including the application, event reference, and timestamp.
type WatchDataAttributes struct {
	// Unique identifier of the RUM application containing the session.
	ApplicationId string `json:"application_id"`
	// Data source type indicating the origin of the session data (e.g., rum or product_analytics).
	DataSource *string `json:"data_source,omitempty"`
	// Unique identifier of the RUM event that was watched.
	EventId string `json:"event_id"`
	// Timestamp when the session was watched.
	Timestamp time.Time `json:"timestamp"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWatchDataAttributes instantiates a new WatchDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWatchDataAttributes(applicationId string, eventId string, timestamp time.Time) *WatchDataAttributes {
	this := WatchDataAttributes{}
	this.ApplicationId = applicationId
	this.EventId = eventId
	this.Timestamp = timestamp
	return &this
}

// NewWatchDataAttributesWithDefaults instantiates a new WatchDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWatchDataAttributesWithDefaults() *WatchDataAttributes {
	this := WatchDataAttributes{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *WatchDataAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *WatchDataAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *WatchDataAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *WatchDataAttributes) GetDataSource() string {
	if o == nil || o.DataSource == nil {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchDataAttributes) GetDataSourceOk() (*string, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *WatchDataAttributes) HasDataSource() bool {
	return o != nil && o.DataSource != nil
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *WatchDataAttributes) SetDataSource(v string) {
	o.DataSource = &v
}

// GetEventId returns the EventId field value.
func (o *WatchDataAttributes) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *WatchDataAttributes) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value.
func (o *WatchDataAttributes) SetEventId(v string) {
	o.EventId = v
}

// GetTimestamp returns the Timestamp field value.
func (o *WatchDataAttributes) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *WatchDataAttributes) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *WatchDataAttributes) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// MarshalJSON serializes the struct using spec logic.
func (o WatchDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["application_id"] = o.ApplicationId
	if o.DataSource != nil {
		toSerialize["data_source"] = o.DataSource
	}
	toSerialize["event_id"] = o.EventId
	if o.Timestamp.Nanosecond() == 0 {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WatchDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *string    `json:"application_id"`
		DataSource    *string    `json:"data_source,omitempty"`
		EventId       *string    `json:"event_id"`
		Timestamp     *time.Time `json:"timestamp"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.EventId == nil {
		return fmt.Errorf("required field event_id missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"application_id", "data_source", "event_id", "timestamp"})
	} else {
		return err
	}
	o.ApplicationId = *all.ApplicationId
	o.DataSource = all.DataSource
	o.EventId = *all.EventId
	o.Timestamp = *all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
