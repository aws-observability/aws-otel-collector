// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnonymizeUsersResponseAttributes Attributes of an anonymize users response.
type AnonymizeUsersResponseAttributes struct {
	// List of errors encountered during anonymization, one entry per failed user.
	AnonymizeErrors []AnonymizeUserError `json:"anonymize_errors"`
	// List of user IDs (UUIDs) that were successfully anonymized.
	AnonymizedUserIds []string `json:"anonymized_user_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnonymizeUsersResponseAttributes instantiates a new AnonymizeUsersResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnonymizeUsersResponseAttributes(anonymizeErrors []AnonymizeUserError, anonymizedUserIds []string) *AnonymizeUsersResponseAttributes {
	this := AnonymizeUsersResponseAttributes{}
	this.AnonymizeErrors = anonymizeErrors
	this.AnonymizedUserIds = anonymizedUserIds
	return &this
}

// NewAnonymizeUsersResponseAttributesWithDefaults instantiates a new AnonymizeUsersResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnonymizeUsersResponseAttributesWithDefaults() *AnonymizeUsersResponseAttributes {
	this := AnonymizeUsersResponseAttributes{}
	return &this
}

// GetAnonymizeErrors returns the AnonymizeErrors field value.
func (o *AnonymizeUsersResponseAttributes) GetAnonymizeErrors() []AnonymizeUserError {
	if o == nil {
		var ret []AnonymizeUserError
		return ret
	}
	return o.AnonymizeErrors
}

// GetAnonymizeErrorsOk returns a tuple with the AnonymizeErrors field value
// and a boolean to check if the value has been set.
func (o *AnonymizeUsersResponseAttributes) GetAnonymizeErrorsOk() (*[]AnonymizeUserError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnonymizeErrors, true
}

// SetAnonymizeErrors sets field value.
func (o *AnonymizeUsersResponseAttributes) SetAnonymizeErrors(v []AnonymizeUserError) {
	o.AnonymizeErrors = v
}

// GetAnonymizedUserIds returns the AnonymizedUserIds field value.
func (o *AnonymizeUsersResponseAttributes) GetAnonymizedUserIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AnonymizedUserIds
}

// GetAnonymizedUserIdsOk returns a tuple with the AnonymizedUserIds field value
// and a boolean to check if the value has been set.
func (o *AnonymizeUsersResponseAttributes) GetAnonymizedUserIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnonymizedUserIds, true
}

// SetAnonymizedUserIds sets field value.
func (o *AnonymizeUsersResponseAttributes) SetAnonymizedUserIds(v []string) {
	o.AnonymizedUserIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnonymizeUsersResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["anonymize_errors"] = o.AnonymizeErrors
	toSerialize["anonymized_user_ids"] = o.AnonymizedUserIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnonymizeUsersResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnonymizeErrors   *[]AnonymizeUserError `json:"anonymize_errors"`
		AnonymizedUserIds *[]string             `json:"anonymized_user_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnonymizeErrors == nil {
		return fmt.Errorf("required field anonymize_errors missing")
	}
	if all.AnonymizedUserIds == nil {
		return fmt.Errorf("required field anonymized_user_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"anonymize_errors", "anonymized_user_ids"})
	} else {
		return err
	}
	o.AnonymizeErrors = *all.AnonymizeErrors
	o.AnonymizedUserIds = *all.AnonymizedUserIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
