// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsSpansResponseMeta Metadata accompanying the spans response.
type LLMObsSpansResponseMeta struct {
	// Time elapsed for the query in milliseconds.
	Elapsed int64 `json:"elapsed"`
	// Pagination cursor for the spans response.
	Page LLMObsSpansResponsePage `json:"page"`
	// Unique identifier for the request.
	RequestId string `json:"request_id"`
	// Status of the query execution.
	Status string `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsSpansResponseMeta instantiates a new LLMObsSpansResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsSpansResponseMeta(elapsed int64, page LLMObsSpansResponsePage, requestId string, status string) *LLMObsSpansResponseMeta {
	this := LLMObsSpansResponseMeta{}
	this.Elapsed = elapsed
	this.Page = page
	this.RequestId = requestId
	this.Status = status
	return &this
}

// NewLLMObsSpansResponseMetaWithDefaults instantiates a new LLMObsSpansResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsSpansResponseMetaWithDefaults() *LLMObsSpansResponseMeta {
	this := LLMObsSpansResponseMeta{}
	return &this
}

// GetElapsed returns the Elapsed field value.
func (o *LLMObsSpansResponseMeta) GetElapsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Elapsed
}

// GetElapsedOk returns a tuple with the Elapsed field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpansResponseMeta) GetElapsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Elapsed, true
}

// SetElapsed sets field value.
func (o *LLMObsSpansResponseMeta) SetElapsed(v int64) {
	o.Elapsed = v
}

// GetPage returns the Page field value.
func (o *LLMObsSpansResponseMeta) GetPage() LLMObsSpansResponsePage {
	if o == nil {
		var ret LLMObsSpansResponsePage
		return ret
	}
	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpansResponseMeta) GetPageOk() (*LLMObsSpansResponsePage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *LLMObsSpansResponseMeta) SetPage(v LLMObsSpansResponsePage) {
	o.Page = v
}

// GetRequestId returns the RequestId field value.
func (o *LLMObsSpansResponseMeta) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpansResponseMeta) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value.
func (o *LLMObsSpansResponseMeta) SetRequestId(v string) {
	o.RequestId = v
}

// GetStatus returns the Status field value.
func (o *LLMObsSpansResponseMeta) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LLMObsSpansResponseMeta) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *LLMObsSpansResponseMeta) SetStatus(v string) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsSpansResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["elapsed"] = o.Elapsed
	toSerialize["page"] = o.Page
	toSerialize["request_id"] = o.RequestId
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsSpansResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Elapsed   *int64                   `json:"elapsed"`
		Page      *LLMObsSpansResponsePage `json:"page"`
		RequestId *string                  `json:"request_id"`
		Status    *string                  `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Elapsed == nil {
		return fmt.Errorf("required field elapsed missing")
	}
	if all.Page == nil {
		return fmt.Errorf("required field page missing")
	}
	if all.RequestId == nil {
		return fmt.Errorf("required field request_id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"elapsed", "page", "request_id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Elapsed = *all.Elapsed
	if all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = *all.Page
	o.RequestId = *all.RequestId
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
