// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateResponseMetadata The metadata associated with a request.
type SpansAggregateResponseMetadata struct {
	// The time elapsed in milliseconds.
	Elapsed *int64 `json:"elapsed,omitempty"`
	// The identifier of the request.
	RequestId *string `json:"request_id,omitempty"`
	// The status of the response.
	Status *SpansAggregateResponseStatus `json:"status,omitempty"`
	// A list of warnings (non fatal errors) encountered, partial results might be returned if
	// warnings are present in the response.
	Warnings []SpansWarning `json:"warnings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansAggregateResponseMetadata instantiates a new SpansAggregateResponseMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansAggregateResponseMetadata() *SpansAggregateResponseMetadata {
	this := SpansAggregateResponseMetadata{}
	return &this
}

// NewSpansAggregateResponseMetadataWithDefaults instantiates a new SpansAggregateResponseMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansAggregateResponseMetadataWithDefaults() *SpansAggregateResponseMetadata {
	this := SpansAggregateResponseMetadata{}
	return &this
}

// GetElapsed returns the Elapsed field value if set, zero value otherwise.
func (o *SpansAggregateResponseMetadata) GetElapsed() int64 {
	if o == nil || o.Elapsed == nil {
		var ret int64
		return ret
	}
	return *o.Elapsed
}

// GetElapsedOk returns a tuple with the Elapsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateResponseMetadata) GetElapsedOk() (*int64, bool) {
	if o == nil || o.Elapsed == nil {
		return nil, false
	}
	return o.Elapsed, true
}

// HasElapsed returns a boolean if a field has been set.
func (o *SpansAggregateResponseMetadata) HasElapsed() bool {
	return o != nil && o.Elapsed != nil
}

// SetElapsed gets a reference to the given int64 and assigns it to the Elapsed field.
func (o *SpansAggregateResponseMetadata) SetElapsed(v int64) {
	o.Elapsed = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *SpansAggregateResponseMetadata) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateResponseMetadata) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *SpansAggregateResponseMetadata) HasRequestId() bool {
	return o != nil && o.RequestId != nil
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *SpansAggregateResponseMetadata) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SpansAggregateResponseMetadata) GetStatus() SpansAggregateResponseStatus {
	if o == nil || o.Status == nil {
		var ret SpansAggregateResponseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateResponseMetadata) GetStatusOk() (*SpansAggregateResponseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SpansAggregateResponseMetadata) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SpansAggregateResponseStatus and assigns it to the Status field.
func (o *SpansAggregateResponseMetadata) SetStatus(v SpansAggregateResponseStatus) {
	o.Status = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SpansAggregateResponseMetadata) GetWarnings() []SpansWarning {
	if o == nil || o.Warnings == nil {
		var ret []SpansWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateResponseMetadata) GetWarningsOk() (*[]SpansWarning, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SpansAggregateResponseMetadata) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []SpansWarning and assigns it to the Warnings field.
func (o *SpansAggregateResponseMetadata) SetWarnings(v []SpansWarning) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansAggregateResponseMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Elapsed != nil {
		toSerialize["elapsed"] = o.Elapsed
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpansAggregateResponseMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Elapsed   *int64                        `json:"elapsed,omitempty"`
		RequestId *string                       `json:"request_id,omitempty"`
		Status    *SpansAggregateResponseStatus `json:"status,omitempty"`
		Warnings  []SpansWarning                `json:"warnings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"elapsed", "request_id", "status", "warnings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Elapsed = all.Elapsed
	o.RequestId = all.RequestId
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Warnings = all.Warnings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
