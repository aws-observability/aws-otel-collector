// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFindingsMeta Metadata about the response.
type SecurityFindingsMeta struct {
	// The time elapsed in milliseconds.
	Elapsed *int64 `json:"elapsed,omitempty"`
	// Pagination information.
	Page *SecurityFindingsPage `json:"page,omitempty"`
	// The identifier of the request.
	RequestId *string `json:"request_id,omitempty"`
	// The status of the response.
	Status *SecurityFindingsStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityFindingsMeta instantiates a new SecurityFindingsMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityFindingsMeta() *SecurityFindingsMeta {
	this := SecurityFindingsMeta{}
	return &this
}

// NewSecurityFindingsMetaWithDefaults instantiates a new SecurityFindingsMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityFindingsMetaWithDefaults() *SecurityFindingsMeta {
	this := SecurityFindingsMeta{}
	return &this
}

// GetElapsed returns the Elapsed field value if set, zero value otherwise.
func (o *SecurityFindingsMeta) GetElapsed() int64 {
	if o == nil || o.Elapsed == nil {
		var ret int64
		return ret
	}
	return *o.Elapsed
}

// GetElapsedOk returns a tuple with the Elapsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFindingsMeta) GetElapsedOk() (*int64, bool) {
	if o == nil || o.Elapsed == nil {
		return nil, false
	}
	return o.Elapsed, true
}

// HasElapsed returns a boolean if a field has been set.
func (o *SecurityFindingsMeta) HasElapsed() bool {
	return o != nil && o.Elapsed != nil
}

// SetElapsed gets a reference to the given int64 and assigns it to the Elapsed field.
func (o *SecurityFindingsMeta) SetElapsed(v int64) {
	o.Elapsed = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *SecurityFindingsMeta) GetPage() SecurityFindingsPage {
	if o == nil || o.Page == nil {
		var ret SecurityFindingsPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFindingsMeta) GetPageOk() (*SecurityFindingsPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *SecurityFindingsMeta) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given SecurityFindingsPage and assigns it to the Page field.
func (o *SecurityFindingsMeta) SetPage(v SecurityFindingsPage) {
	o.Page = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *SecurityFindingsMeta) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFindingsMeta) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *SecurityFindingsMeta) HasRequestId() bool {
	return o != nil && o.RequestId != nil
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *SecurityFindingsMeta) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityFindingsMeta) GetStatus() SecurityFindingsStatus {
	if o == nil || o.Status == nil {
		var ret SecurityFindingsStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFindingsMeta) GetStatusOk() (*SecurityFindingsStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecurityFindingsMeta) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SecurityFindingsStatus and assigns it to the Status field.
func (o *SecurityFindingsMeta) SetStatus(v SecurityFindingsStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityFindingsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Elapsed != nil {
		toSerialize["elapsed"] = o.Elapsed
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityFindingsMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Elapsed   *int64                  `json:"elapsed,omitempty"`
		Page      *SecurityFindingsPage   `json:"page,omitempty"`
		RequestId *string                 `json:"request_id,omitempty"`
		Status    *SecurityFindingsStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"elapsed", "page", "request_id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Elapsed = all.Elapsed
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
	o.RequestId = all.RequestId
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
