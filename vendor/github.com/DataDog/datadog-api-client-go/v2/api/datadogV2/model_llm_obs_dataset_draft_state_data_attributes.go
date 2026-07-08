// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetDraftStateDataAttributes Attributes of an LLM Observability dataset draft state.
type LLMObsDatasetDraftStateDataAttributes struct {
	// Timestamp when the dataset draft session started.
	DraftingSince time.Time `json:"drafting_since"`
	// User information associated with a dataset draft state.
	User LLMObsDatasetDraftStateUser `json:"user"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetDraftStateDataAttributes instantiates a new LLMObsDatasetDraftStateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetDraftStateDataAttributes(draftingSince time.Time, user LLMObsDatasetDraftStateUser) *LLMObsDatasetDraftStateDataAttributes {
	this := LLMObsDatasetDraftStateDataAttributes{}
	this.DraftingSince = draftingSince
	this.User = user
	return &this
}

// NewLLMObsDatasetDraftStateDataAttributesWithDefaults instantiates a new LLMObsDatasetDraftStateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetDraftStateDataAttributesWithDefaults() *LLMObsDatasetDraftStateDataAttributes {
	this := LLMObsDatasetDraftStateDataAttributes{}
	return &this
}

// GetDraftingSince returns the DraftingSince field value.
func (o *LLMObsDatasetDraftStateDataAttributes) GetDraftingSince() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.DraftingSince
}

// GetDraftingSinceOk returns a tuple with the DraftingSince field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetDraftStateDataAttributes) GetDraftingSinceOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DraftingSince, true
}

// SetDraftingSince sets field value.
func (o *LLMObsDatasetDraftStateDataAttributes) SetDraftingSince(v time.Time) {
	o.DraftingSince = v
}

// GetUser returns the User field value.
func (o *LLMObsDatasetDraftStateDataAttributes) GetUser() LLMObsDatasetDraftStateUser {
	if o == nil {
		var ret LLMObsDatasetDraftStateUser
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetDraftStateDataAttributes) GetUserOk() (*LLMObsDatasetDraftStateUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *LLMObsDatasetDraftStateDataAttributes) SetUser(v LLMObsDatasetDraftStateUser) {
	o.User = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetDraftStateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DraftingSince.Nanosecond() == 0 {
		toSerialize["drafting_since"] = o.DraftingSince.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["drafting_since"] = o.DraftingSince.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["user"] = o.User

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetDraftStateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DraftingSince *time.Time                   `json:"drafting_since"`
		User          *LLMObsDatasetDraftStateUser `json:"user"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DraftingSince == nil {
		return fmt.Errorf("required field drafting_since missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"drafting_since", "user"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DraftingSince = *all.DraftingSince
	if all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.User = *all.User

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
