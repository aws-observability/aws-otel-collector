// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsAPILimitReachedResponse Response returned by the Logs API when the max limit has been reached.
type LogsAPILimitReachedResponse struct {
	// Error returned by the Logs API
	Error *LogsAPIError `json:"error,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsAPILimitReachedResponse instantiates a new LogsAPILimitReachedResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsAPILimitReachedResponse() *LogsAPILimitReachedResponse {
	this := LogsAPILimitReachedResponse{}
	return &this
}

// NewLogsAPILimitReachedResponseWithDefaults instantiates a new LogsAPILimitReachedResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsAPILimitReachedResponseWithDefaults() *LogsAPILimitReachedResponse {
	this := LogsAPILimitReachedResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *LogsAPILimitReachedResponse) GetError() LogsAPIError {
	if o == nil || o.Error == nil {
		var ret LogsAPIError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsAPILimitReachedResponse) GetErrorOk() (*LogsAPIError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *LogsAPILimitReachedResponse) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given LogsAPIError and assigns it to the Error field.
func (o *LogsAPILimitReachedResponse) SetError(v LogsAPIError) {
	o.Error = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsAPILimitReachedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsAPILimitReachedResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error *LogsAPIError `json:"error,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Error != nil && all.Error.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Error = all.Error

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
