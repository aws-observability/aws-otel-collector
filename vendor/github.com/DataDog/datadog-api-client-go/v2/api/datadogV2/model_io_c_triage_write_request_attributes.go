// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IoCTriageWriteRequestAttributes Attributes for setting an indicator's triage state.
type IoCTriageWriteRequestAttributes struct {
	// The indicator value to triage (for example, an IP address or domain).
	Indicator string `json:"indicator"`
	// Current triage state of the indicator.
	TriageState IoCTriageState `json:"triage_state"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIoCTriageWriteRequestAttributes instantiates a new IoCTriageWriteRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIoCTriageWriteRequestAttributes(indicator string, triageState IoCTriageState) *IoCTriageWriteRequestAttributes {
	this := IoCTriageWriteRequestAttributes{}
	this.Indicator = indicator
	this.TriageState = triageState
	return &this
}

// NewIoCTriageWriteRequestAttributesWithDefaults instantiates a new IoCTriageWriteRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIoCTriageWriteRequestAttributesWithDefaults() *IoCTriageWriteRequestAttributes {
	this := IoCTriageWriteRequestAttributes{}
	return &this
}

// GetIndicator returns the Indicator field value.
func (o *IoCTriageWriteRequestAttributes) GetIndicator() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value
// and a boolean to check if the value has been set.
func (o *IoCTriageWriteRequestAttributes) GetIndicatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indicator, true
}

// SetIndicator sets field value.
func (o *IoCTriageWriteRequestAttributes) SetIndicator(v string) {
	o.Indicator = v
}

// GetTriageState returns the TriageState field value.
func (o *IoCTriageWriteRequestAttributes) GetTriageState() IoCTriageState {
	if o == nil {
		var ret IoCTriageState
		return ret
	}
	return o.TriageState
}

// GetTriageStateOk returns a tuple with the TriageState field value
// and a boolean to check if the value has been set.
func (o *IoCTriageWriteRequestAttributes) GetTriageStateOk() (*IoCTriageState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriageState, true
}

// SetTriageState sets field value.
func (o *IoCTriageWriteRequestAttributes) SetTriageState(v IoCTriageState) {
	o.TriageState = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IoCTriageWriteRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["indicator"] = o.Indicator
	toSerialize["triage_state"] = o.TriageState

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IoCTriageWriteRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Indicator   *string         `json:"indicator"`
		TriageState *IoCTriageState `json:"triage_state"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Indicator == nil {
		return fmt.Errorf("required field indicator missing")
	}
	if all.TriageState == nil {
		return fmt.Errorf("required field triage_state missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"indicator", "triage_state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Indicator = *all.Indicator
	if !all.TriageState.IsValid() {
		hasInvalidField = true
	} else {
		o.TriageState = *all.TriageState
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
