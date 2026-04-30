// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseUpdateStatusAttributes Case update status attributes
type CaseUpdateStatusAttributes struct {
	// Deprecated way of representing the case status, which only supports OPEN, IN_PROGRESS, and CLOSED statuses. Use `status_name` instead.
	// Deprecated
	Status *CaseStatus `json:"status,omitempty"`
	// Status of the case. Must be one of the existing statuses for the case's type.
	StatusName *string `json:"status_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseUpdateStatusAttributes instantiates a new CaseUpdateStatusAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseUpdateStatusAttributes() *CaseUpdateStatusAttributes {
	this := CaseUpdateStatusAttributes{}
	return &this
}

// NewCaseUpdateStatusAttributesWithDefaults instantiates a new CaseUpdateStatusAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseUpdateStatusAttributesWithDefaults() *CaseUpdateStatusAttributes {
	this := CaseUpdateStatusAttributes{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
// Deprecated
func (o *CaseUpdateStatusAttributes) GetStatus() CaseStatus {
	if o == nil || o.Status == nil {
		var ret CaseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CaseUpdateStatusAttributes) GetStatusOk() (*CaseStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CaseUpdateStatusAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given CaseStatus and assigns it to the Status field.
// Deprecated
func (o *CaseUpdateStatusAttributes) SetStatus(v CaseStatus) {
	o.Status = &v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *CaseUpdateStatusAttributes) GetStatusName() string {
	if o == nil || o.StatusName == nil {
		var ret string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseUpdateStatusAttributes) GetStatusNameOk() (*string, bool) {
	if o == nil || o.StatusName == nil {
		return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *CaseUpdateStatusAttributes) HasStatusName() bool {
	return o != nil && o.StatusName != nil
}

// SetStatusName gets a reference to the given string and assigns it to the StatusName field.
func (o *CaseUpdateStatusAttributes) SetStatusName(v string) {
	o.StatusName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseUpdateStatusAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusName != nil {
		toSerialize["status_name"] = o.StatusName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseUpdateStatusAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status     *CaseStatus `json:"status,omitempty"`
		StatusName *string     `json:"status_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"status", "status_name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.StatusName = all.StatusName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
