// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CVSS Vulnerability severity.
type CVSS struct {
	// Vulnerability severity score.
	Score float64 `json:"score"`
	// The vulnerability severity.
	Severity VulnerabilitySeverity `json:"severity"`
	// Vulnerability CVSS vector.
	Vector string `json:"vector"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCVSS instantiates a new CVSS object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCVSS(score float64, severity VulnerabilitySeverity, vector string) *CVSS {
	this := CVSS{}
	this.Score = score
	this.Severity = severity
	this.Vector = vector
	return &this
}

// NewCVSSWithDefaults instantiates a new CVSS object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCVSSWithDefaults() *CVSS {
	this := CVSS{}
	return &this
}

// GetScore returns the Score field value.
func (o *CVSS) GetScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *CVSS) GetScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value.
func (o *CVSS) SetScore(v float64) {
	o.Score = v
}

// GetSeverity returns the Severity field value.
func (o *CVSS) GetSeverity() VulnerabilitySeverity {
	if o == nil {
		var ret VulnerabilitySeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *CVSS) GetSeverityOk() (*VulnerabilitySeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *CVSS) SetSeverity(v VulnerabilitySeverity) {
	o.Severity = v
}

// GetVector returns the Vector field value.
func (o *CVSS) GetVector() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Vector
}

// GetVectorOk returns a tuple with the Vector field value
// and a boolean to check if the value has been set.
func (o *CVSS) GetVectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vector, true
}

// SetVector sets field value.
func (o *CVSS) SetVector(v string) {
	o.Vector = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CVSS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["score"] = o.Score
	toSerialize["severity"] = o.Severity
	toSerialize["vector"] = o.Vector

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CVSS) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Score    *float64               `json:"score"`
		Severity *VulnerabilitySeverity `json:"severity"`
		Vector   *string                `json:"vector"`
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
	if all.Vector == nil {
		return fmt.Errorf("required field vector missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"score", "severity", "vector"})
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
	o.Vector = *all.Vector

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
