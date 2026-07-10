// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OCIConfigAttributes Attributes for an OCI config.
type OCIConfigAttributes struct {
	// The OCID of the OCI tenancy.
	AccountId string `json:"account_id"`
	// The timestamp when the OCI config was created.
	CreatedAt string `json:"created_at"`
	// The error messages for the OCI config.
	ErrorMessages datadog.NullableList[string] `json:"error_messages,omitempty"`
	// The status of the OCI config.
	Status string `json:"status"`
	// The timestamp when the OCI config status was last updated.
	StatusUpdatedAt string `json:"status_updated_at"`
	// The timestamp when the OCI config was last updated.
	UpdatedAt string `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOCIConfigAttributes instantiates a new OCIConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOCIConfigAttributes(accountId string, createdAt string, status string, statusUpdatedAt string, updatedAt string) *OCIConfigAttributes {
	this := OCIConfigAttributes{}
	this.AccountId = accountId
	this.CreatedAt = createdAt
	this.Status = status
	this.StatusUpdatedAt = statusUpdatedAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewOCIConfigAttributesWithDefaults instantiates a new OCIConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOCIConfigAttributesWithDefaults() *OCIConfigAttributes {
	this := OCIConfigAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *OCIConfigAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *OCIConfigAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *OCIConfigAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OCIConfigAttributes) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OCIConfigAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OCIConfigAttributes) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OCIConfigAttributes) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ErrorMessages.Get()
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OCIConfigAttributes) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessages.Get(), o.ErrorMessages.IsSet()
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *OCIConfigAttributes) HasErrorMessages() bool {
	return o != nil && o.ErrorMessages.IsSet()
}

// SetErrorMessages gets a reference to the given datadog.NullableList[string] and assigns it to the ErrorMessages field.
func (o *OCIConfigAttributes) SetErrorMessages(v []string) {
	o.ErrorMessages.Set(&v)
}

// SetErrorMessagesNil sets the value for ErrorMessages to be an explicit nil.
func (o *OCIConfigAttributes) SetErrorMessagesNil() {
	o.ErrorMessages.Set(nil)
}

// UnsetErrorMessages ensures that no value is present for ErrorMessages, not even an explicit nil.
func (o *OCIConfigAttributes) UnsetErrorMessages() {
	o.ErrorMessages.Unset()
}

// GetStatus returns the Status field value.
func (o *OCIConfigAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OCIConfigAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *OCIConfigAttributes) SetStatus(v string) {
	o.Status = v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value.
func (o *OCIConfigAttributes) GetStatusUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OCIConfigAttributes) GetStatusUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusUpdatedAt, true
}

// SetStatusUpdatedAt sets field value.
func (o *OCIConfigAttributes) SetStatusUpdatedAt(v string) {
	o.StatusUpdatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *OCIConfigAttributes) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OCIConfigAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *OCIConfigAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OCIConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["created_at"] = o.CreatedAt
	if o.ErrorMessages.IsSet() {
		toSerialize["error_messages"] = o.ErrorMessages.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["status_updated_at"] = o.StatusUpdatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OCIConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId       *string                      `json:"account_id"`
		CreatedAt       *string                      `json:"created_at"`
		ErrorMessages   datadog.NullableList[string] `json:"error_messages,omitempty"`
		Status          *string                      `json:"status"`
		StatusUpdatedAt *string                      `json:"status_updated_at"`
		UpdatedAt       *string                      `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.StatusUpdatedAt == nil {
		return fmt.Errorf("required field status_updated_at missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "created_at", "error_messages", "status", "status_updated_at", "updated_at"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.CreatedAt = *all.CreatedAt
	o.ErrorMessages = all.ErrorMessages
	o.Status = *all.Status
	o.StatusUpdatedAt = *all.StatusUpdatedAt
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
