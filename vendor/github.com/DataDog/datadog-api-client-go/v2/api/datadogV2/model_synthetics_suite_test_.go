// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsSuiteTest Object containing details about a Synthetic test included in a Synthetic suite.
type SyntheticsSuiteTest struct {
	// Alerting criticality for each the test.
	AlertingCriticality *SyntheticsSuiteTestAlertingCriticality `json:"alerting_criticality,omitempty"`
	// The public ID of the Synthetic test included in the suite.
	PublicId string `json:"public_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsSuiteTest instantiates a new SyntheticsSuiteTest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsSuiteTest(publicId string) *SyntheticsSuiteTest {
	this := SyntheticsSuiteTest{}
	this.PublicId = publicId
	return &this
}

// NewSyntheticsSuiteTestWithDefaults instantiates a new SyntheticsSuiteTest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsSuiteTestWithDefaults() *SyntheticsSuiteTest {
	this := SyntheticsSuiteTest{}
	return &this
}

// GetAlertingCriticality returns the AlertingCriticality field value if set, zero value otherwise.
func (o *SyntheticsSuiteTest) GetAlertingCriticality() SyntheticsSuiteTestAlertingCriticality {
	if o == nil || o.AlertingCriticality == nil {
		var ret SyntheticsSuiteTestAlertingCriticality
		return ret
	}
	return *o.AlertingCriticality
}

// GetAlertingCriticalityOk returns a tuple with the AlertingCriticality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsSuiteTest) GetAlertingCriticalityOk() (*SyntheticsSuiteTestAlertingCriticality, bool) {
	if o == nil || o.AlertingCriticality == nil {
		return nil, false
	}
	return o.AlertingCriticality, true
}

// HasAlertingCriticality returns a boolean if a field has been set.
func (o *SyntheticsSuiteTest) HasAlertingCriticality() bool {
	return o != nil && o.AlertingCriticality != nil
}

// SetAlertingCriticality gets a reference to the given SyntheticsSuiteTestAlertingCriticality and assigns it to the AlertingCriticality field.
func (o *SyntheticsSuiteTest) SetAlertingCriticality(v SyntheticsSuiteTestAlertingCriticality) {
	o.AlertingCriticality = &v
}

// GetPublicId returns the PublicId field value.
func (o *SyntheticsSuiteTest) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *SyntheticsSuiteTest) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value.
func (o *SyntheticsSuiteTest) SetPublicId(v string) {
	o.PublicId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsSuiteTest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AlertingCriticality != nil {
		toSerialize["alerting_criticality"] = o.AlertingCriticality
	}
	toSerialize["public_id"] = o.PublicId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsSuiteTest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AlertingCriticality *SyntheticsSuiteTestAlertingCriticality `json:"alerting_criticality,omitempty"`
		PublicId            *string                                 `json:"public_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PublicId == nil {
		return fmt.Errorf("required field public_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alerting_criticality", "public_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AlertingCriticality != nil && !all.AlertingCriticality.IsValid() {
		hasInvalidField = true
	} else {
		o.AlertingCriticality = all.AlertingCriticality
	}
	o.PublicId = *all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
