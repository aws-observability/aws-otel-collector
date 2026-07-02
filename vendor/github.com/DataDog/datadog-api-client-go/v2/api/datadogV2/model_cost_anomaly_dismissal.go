// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostAnomalyDismissal Resolution metadata for an anomaly that has been dismissed.
type CostAnomalyDismissal struct {
	// Reason the anomaly was dismissed.
	Cause string `json:"cause"`
	// Unique identifier of the dismissal record.
	DismissalId string `json:"dismissal_id"`
	// Optional message explaining the dismissal.
	Message string `json:"message"`
	// Timestamp of the last dismissal update in Unix milliseconds.
	UpdatedAt int64 `json:"updated_at"`
	// Identifier of the user that last updated the dismissal.
	UpdatedBy string `json:"updated_by"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostAnomalyDismissal instantiates a new CostAnomalyDismissal object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostAnomalyDismissal(cause string, dismissalId string, message string, updatedAt int64, updatedBy string) *CostAnomalyDismissal {
	this := CostAnomalyDismissal{}
	this.Cause = cause
	this.DismissalId = dismissalId
	this.Message = message
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewCostAnomalyDismissalWithDefaults instantiates a new CostAnomalyDismissal object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostAnomalyDismissalWithDefaults() *CostAnomalyDismissal {
	this := CostAnomalyDismissal{}
	return &this
}

// GetCause returns the Cause field value.
func (o *CostAnomalyDismissal) GetCause() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *CostAnomalyDismissal) GetCauseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value.
func (o *CostAnomalyDismissal) SetCause(v string) {
	o.Cause = v
}

// GetDismissalId returns the DismissalId field value.
func (o *CostAnomalyDismissal) GetDismissalId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DismissalId
}

// GetDismissalIdOk returns a tuple with the DismissalId field value
// and a boolean to check if the value has been set.
func (o *CostAnomalyDismissal) GetDismissalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DismissalId, true
}

// SetDismissalId sets field value.
func (o *CostAnomalyDismissal) SetDismissalId(v string) {
	o.DismissalId = v
}

// GetMessage returns the Message field value.
func (o *CostAnomalyDismissal) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CostAnomalyDismissal) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *CostAnomalyDismissal) SetMessage(v string) {
	o.Message = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *CostAnomalyDismissal) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CostAnomalyDismissal) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *CostAnomalyDismissal) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value.
func (o *CostAnomalyDismissal) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *CostAnomalyDismissal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value.
func (o *CostAnomalyDismissal) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostAnomalyDismissal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cause"] = o.Cause
	toSerialize["dismissal_id"] = o.DismissalId
	toSerialize["message"] = o.Message
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["updated_by"] = o.UpdatedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostAnomalyDismissal) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cause       *string `json:"cause"`
		DismissalId *string `json:"dismissal_id"`
		Message     *string `json:"message"`
		UpdatedAt   *int64  `json:"updated_at"`
		UpdatedBy   *string `json:"updated_by"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cause == nil {
		return fmt.Errorf("required field cause missing")
	}
	if all.DismissalId == nil {
		return fmt.Errorf("required field dismissal_id missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.UpdatedBy == nil {
		return fmt.Errorf("required field updated_by missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cause", "dismissal_id", "message", "updated_at", "updated_by"})
	} else {
		return err
	}
	o.Cause = *all.Cause
	o.DismissalId = *all.DismissalId
	o.Message = *all.Message
	o.UpdatedAt = *all.UpdatedAt
	o.UpdatedBy = *all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
