// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestDecisionCreateAttributes Attributes for creating a change request decision.
type ChangeRequestDecisionCreateAttributes struct {
	// The status of a change request decision.
	ChangeRequestStatus *ChangeRequestDecisionStatusType `json:"change_request_status,omitempty"`
	// The reason for requesting the decision.
	RequestReason *string `json:"request_reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestDecisionCreateAttributes instantiates a new ChangeRequestDecisionCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestDecisionCreateAttributes() *ChangeRequestDecisionCreateAttributes {
	this := ChangeRequestDecisionCreateAttributes{}
	return &this
}

// NewChangeRequestDecisionCreateAttributesWithDefaults instantiates a new ChangeRequestDecisionCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestDecisionCreateAttributesWithDefaults() *ChangeRequestDecisionCreateAttributes {
	this := ChangeRequestDecisionCreateAttributes{}
	return &this
}

// GetChangeRequestStatus returns the ChangeRequestStatus field value if set, zero value otherwise.
func (o *ChangeRequestDecisionCreateAttributes) GetChangeRequestStatus() ChangeRequestDecisionStatusType {
	if o == nil || o.ChangeRequestStatus == nil {
		var ret ChangeRequestDecisionStatusType
		return ret
	}
	return *o.ChangeRequestStatus
}

// GetChangeRequestStatusOk returns a tuple with the ChangeRequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionCreateAttributes) GetChangeRequestStatusOk() (*ChangeRequestDecisionStatusType, bool) {
	if o == nil || o.ChangeRequestStatus == nil {
		return nil, false
	}
	return o.ChangeRequestStatus, true
}

// HasChangeRequestStatus returns a boolean if a field has been set.
func (o *ChangeRequestDecisionCreateAttributes) HasChangeRequestStatus() bool {
	return o != nil && o.ChangeRequestStatus != nil
}

// SetChangeRequestStatus gets a reference to the given ChangeRequestDecisionStatusType and assigns it to the ChangeRequestStatus field.
func (o *ChangeRequestDecisionCreateAttributes) SetChangeRequestStatus(v ChangeRequestDecisionStatusType) {
	o.ChangeRequestStatus = &v
}

// GetRequestReason returns the RequestReason field value if set, zero value otherwise.
func (o *ChangeRequestDecisionCreateAttributes) GetRequestReason() string {
	if o == nil || o.RequestReason == nil {
		var ret string
		return ret
	}
	return *o.RequestReason
}

// GetRequestReasonOk returns a tuple with the RequestReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionCreateAttributes) GetRequestReasonOk() (*string, bool) {
	if o == nil || o.RequestReason == nil {
		return nil, false
	}
	return o.RequestReason, true
}

// HasRequestReason returns a boolean if a field has been set.
func (o *ChangeRequestDecisionCreateAttributes) HasRequestReason() bool {
	return o != nil && o.RequestReason != nil
}

// SetRequestReason gets a reference to the given string and assigns it to the RequestReason field.
func (o *ChangeRequestDecisionCreateAttributes) SetRequestReason(v string) {
	o.RequestReason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestDecisionCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChangeRequestStatus != nil {
		toSerialize["change_request_status"] = o.ChangeRequestStatus
	}
	if o.RequestReason != nil {
		toSerialize["request_reason"] = o.RequestReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestDecisionCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeRequestStatus *ChangeRequestDecisionStatusType `json:"change_request_status,omitempty"`
		RequestReason       *string                          `json:"request_reason,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_request_status", "request_reason"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ChangeRequestStatus != nil && !all.ChangeRequestStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeRequestStatus = all.ChangeRequestStatus
	}
	o.RequestReason = all.RequestReason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
