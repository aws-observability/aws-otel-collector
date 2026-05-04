// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFindingsSearchRequestPage Pagination attributes for the search request.
type SecurityFindingsSearchRequestPage struct {
	// Get the next page of results with a cursor provided in the previous query.
	Cursor *string `json:"cursor,omitempty"`
	// The maximum number of security findings in the response.
	Limit *int64 `json:"limit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityFindingsSearchRequestPage instantiates a new SecurityFindingsSearchRequestPage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityFindingsSearchRequestPage() *SecurityFindingsSearchRequestPage {
	this := SecurityFindingsSearchRequestPage{}
	var limit int64 = 10
	this.Limit = &limit
	return &this
}

// NewSecurityFindingsSearchRequestPageWithDefaults instantiates a new SecurityFindingsSearchRequestPage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityFindingsSearchRequestPageWithDefaults() *SecurityFindingsSearchRequestPage {
	this := SecurityFindingsSearchRequestPage{}
	var limit int64 = 10
	this.Limit = &limit
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *SecurityFindingsSearchRequestPage) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFindingsSearchRequestPage) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *SecurityFindingsSearchRequestPage) HasCursor() bool {
	return o != nil && o.Cursor != nil
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *SecurityFindingsSearchRequestPage) SetCursor(v string) {
	o.Cursor = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SecurityFindingsSearchRequestPage) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityFindingsSearchRequestPage) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SecurityFindingsSearchRequestPage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *SecurityFindingsSearchRequestPage) SetLimit(v int64) {
	o.Limit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityFindingsSearchRequestPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityFindingsSearchRequestPage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cursor *string `json:"cursor,omitempty"`
		Limit  *int64  `json:"limit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cursor", "limit"})
	} else {
		return err
	}
	o.Cursor = all.Cursor
	o.Limit = all.Limit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
