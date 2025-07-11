// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EPSS Vulnerability EPSS severity.
type EPSS struct {
	// Vulnerability EPSS severity score.
	Score float64 `json:"score"`
	// The vulnerability severity.
	Severity VulnerabilitySeverity `json:"severity"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEPSS instantiates a new EPSS object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEPSS(score float64, severity VulnerabilitySeverity) *EPSS {
	this := EPSS{}
	this.Score = score
	this.Severity = severity
	return &this
}

// NewEPSSWithDefaults instantiates a new EPSS object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEPSSWithDefaults() *EPSS {
	this := EPSS{}
	return &this
}

// GetScore returns the Score field value.
func (o *EPSS) GetScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *EPSS) GetScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value.
func (o *EPSS) SetScore(v float64) {
	o.Score = v
}

// GetSeverity returns the Severity field value.
func (o *EPSS) GetSeverity() VulnerabilitySeverity {
	if o == nil {
		var ret VulnerabilitySeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *EPSS) GetSeverityOk() (*VulnerabilitySeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *EPSS) SetSeverity(v VulnerabilitySeverity) {
	o.Severity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EPSS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["score"] = o.Score
	toSerialize["severity"] = o.Severity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EPSS) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Score    *float64               `json:"score"`
		Severity *VulnerabilitySeverity `json:"severity"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Score == nil {
		return fmt.Errorf("required field score missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"score", "severity"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Score = *all.Score
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
