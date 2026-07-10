// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ValidateAPIKeyResponse Response object for the API and application key validation status check.
type ValidateAPIKeyResponse struct {
	// Status of the validation. Always `ok` when both the API key and the application key are valid.
	Status ValidateAPIKeyStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewValidateAPIKeyResponse instantiates a new ValidateAPIKeyResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewValidateAPIKeyResponse(status ValidateAPIKeyStatus) *ValidateAPIKeyResponse {
	this := ValidateAPIKeyResponse{}
	this.Status = status
	return &this
}

// NewValidateAPIKeyResponseWithDefaults instantiates a new ValidateAPIKeyResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewValidateAPIKeyResponseWithDefaults() *ValidateAPIKeyResponse {
	this := ValidateAPIKeyResponse{}
	return &this
}

// GetStatus returns the Status field value.
func (o *ValidateAPIKeyResponse) GetStatus() ValidateAPIKeyStatus {
	if o == nil {
		var ret ValidateAPIKeyStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ValidateAPIKeyResponse) GetStatusOk() (*ValidateAPIKeyStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ValidateAPIKeyResponse) SetStatus(v ValidateAPIKeyStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ValidateAPIKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ValidateAPIKeyResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status *ValidateAPIKeyStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"status"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
