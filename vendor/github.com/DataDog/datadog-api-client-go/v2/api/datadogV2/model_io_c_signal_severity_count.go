// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IoCSignalSeverityCount Count of security signals by severity level.
type IoCSignalSeverityCount struct {
	// Number of signals at this severity level.
	Count *int64 `json:"count,omitempty"`
	// Severity level (for example, critical, high, medium, low, info).
	Severity *string `json:"severity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIoCSignalSeverityCount instantiates a new IoCSignalSeverityCount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIoCSignalSeverityCount() *IoCSignalSeverityCount {
	this := IoCSignalSeverityCount{}
	return &this
}

// NewIoCSignalSeverityCountWithDefaults instantiates a new IoCSignalSeverityCount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIoCSignalSeverityCountWithDefaults() *IoCSignalSeverityCount {
	this := IoCSignalSeverityCount{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *IoCSignalSeverityCount) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCSignalSeverityCount) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *IoCSignalSeverityCount) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *IoCSignalSeverityCount) SetCount(v int64) {
	o.Count = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IoCSignalSeverityCount) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCSignalSeverityCount) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IoCSignalSeverityCount) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *IoCSignalSeverityCount) SetSeverity(v string) {
	o.Severity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IoCSignalSeverityCount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IoCSignalSeverityCount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count    *int64  `json:"count,omitempty"`
		Severity *string `json:"severity,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "severity"})
	} else {
		return err
	}
	o.Count = all.Count
	o.Severity = all.Severity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
