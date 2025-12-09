// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestsMobileApplication Mobile application for mobile synthetics test.
type SyntheticsMobileTestsMobileApplication struct {
	// Application ID of the mobile application.
	ApplicationId string `json:"applicationId"`
	// Reference ID of the mobile application.
	ReferenceId string `json:"referenceId"`
	// Reference type for the mobile application for a mobile synthetics test.
	ReferenceType SyntheticsMobileTestsMobileApplicationReferenceType `json:"referenceType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileTestsMobileApplication instantiates a new SyntheticsMobileTestsMobileApplication object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileTestsMobileApplication(applicationId string, referenceId string, referenceType SyntheticsMobileTestsMobileApplicationReferenceType) *SyntheticsMobileTestsMobileApplication {
	this := SyntheticsMobileTestsMobileApplication{}
	this.ApplicationId = applicationId
	this.ReferenceId = referenceId
	this.ReferenceType = referenceType
	return &this
}

// NewSyntheticsMobileTestsMobileApplicationWithDefaults instantiates a new SyntheticsMobileTestsMobileApplication object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileTestsMobileApplicationWithDefaults() *SyntheticsMobileTestsMobileApplication {
	this := SyntheticsMobileTestsMobileApplication{}
	return &this
}

// GetApplicationId returns the ApplicationId field value.
func (o *SyntheticsMobileTestsMobileApplication) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestsMobileApplication) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *SyntheticsMobileTestsMobileApplication) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetReferenceId returns the ReferenceId field value.
func (o *SyntheticsMobileTestsMobileApplication) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestsMobileApplication) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value.
func (o *SyntheticsMobileTestsMobileApplication) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetReferenceType returns the ReferenceType field value.
func (o *SyntheticsMobileTestsMobileApplication) GetReferenceType() SyntheticsMobileTestsMobileApplicationReferenceType {
	if o == nil {
		var ret SyntheticsMobileTestsMobileApplicationReferenceType
		return ret
	}
	return o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestsMobileApplication) GetReferenceTypeOk() (*SyntheticsMobileTestsMobileApplicationReferenceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceType, true
}

// SetReferenceType sets field value.
func (o *SyntheticsMobileTestsMobileApplication) SetReferenceType(v SyntheticsMobileTestsMobileApplicationReferenceType) {
	o.ReferenceType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileTestsMobileApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["applicationId"] = o.ApplicationId
	toSerialize["referenceId"] = o.ReferenceId
	toSerialize["referenceType"] = o.ReferenceType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileTestsMobileApplication) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApplicationId *string                                              `json:"applicationId"`
		ReferenceId   *string                                              `json:"referenceId"`
		ReferenceType *SyntheticsMobileTestsMobileApplicationReferenceType `json:"referenceType"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field applicationId missing")
	}
	if all.ReferenceId == nil {
		return fmt.Errorf("required field referenceId missing")
	}
	if all.ReferenceType == nil {
		return fmt.Errorf("required field referenceType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"applicationId", "referenceId", "referenceType"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApplicationId = *all.ApplicationId
	o.ReferenceId = *all.ReferenceId
	if !all.ReferenceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ReferenceType = *all.ReferenceType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
