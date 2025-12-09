// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LeakedKeyAttributes The definition of LeakedKeyAttributes object.
type LeakedKeyAttributes struct {
	// The LeakedKeyAttributes date.
	Date time.Time `json:"date"`
	// The LeakedKeyAttributes leak_source.
	LeakSource *string `json:"leak_source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLeakedKeyAttributes instantiates a new LeakedKeyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLeakedKeyAttributes(date time.Time) *LeakedKeyAttributes {
	this := LeakedKeyAttributes{}
	this.Date = date
	return &this
}

// NewLeakedKeyAttributesWithDefaults instantiates a new LeakedKeyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLeakedKeyAttributesWithDefaults() *LeakedKeyAttributes {
	this := LeakedKeyAttributes{}
	return &this
}

// GetDate returns the Date field value.
func (o *LeakedKeyAttributes) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *LeakedKeyAttributes) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *LeakedKeyAttributes) SetDate(v time.Time) {
	o.Date = v
}

// GetLeakSource returns the LeakSource field value if set, zero value otherwise.
func (o *LeakedKeyAttributes) GetLeakSource() string {
	if o == nil || o.LeakSource == nil {
		var ret string
		return ret
	}
	return *o.LeakSource
}

// GetLeakSourceOk returns a tuple with the LeakSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeakedKeyAttributes) GetLeakSourceOk() (*string, bool) {
	if o == nil || o.LeakSource == nil {
		return nil, false
	}
	return o.LeakSource, true
}

// HasLeakSource returns a boolean if a field has been set.
func (o *LeakedKeyAttributes) HasLeakSource() bool {
	return o != nil && o.LeakSource != nil
}

// SetLeakSource gets a reference to the given string and assigns it to the LeakSource field.
func (o *LeakedKeyAttributes) SetLeakSource(v string) {
	o.LeakSource = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LeakedKeyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Date.Nanosecond() == 0 {
		toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["date"] = o.Date.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.LeakSource != nil {
		toSerialize["leak_source"] = o.LeakSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LeakedKeyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Date       *time.Time `json:"date"`
		LeakSource *string    `json:"leak_source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Date == nil {
		return fmt.Errorf("required field date missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"date", "leak_source"})
	} else {
		return err
	}
	o.Date = *all.Date
	o.LeakSource = all.LeakSource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
