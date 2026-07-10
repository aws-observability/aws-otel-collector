// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueRegression Regression information for an issue that was previously resolved and then reopened.
type IssueRegression struct {
	// Timestamp when the issue was reopened (regressed).
	RegressedAt time.Time `json:"regressed_at"`
	// Application version where the regression was observed.
	RegressedAtVersion *string `json:"regressed_at_version,omitempty"`
	// Timestamp when the issue was resolved before the regression.
	ResolvedAt time.Time `json:"resolved_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueRegression instantiates a new IssueRegression object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueRegression(regressedAt time.Time, resolvedAt time.Time) *IssueRegression {
	this := IssueRegression{}
	this.RegressedAt = regressedAt
	this.ResolvedAt = resolvedAt
	return &this
}

// NewIssueRegressionWithDefaults instantiates a new IssueRegression object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueRegressionWithDefaults() *IssueRegression {
	this := IssueRegression{}
	return &this
}

// GetRegressedAt returns the RegressedAt field value.
func (o *IssueRegression) GetRegressedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.RegressedAt
}

// GetRegressedAtOk returns a tuple with the RegressedAt field value
// and a boolean to check if the value has been set.
func (o *IssueRegression) GetRegressedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegressedAt, true
}

// SetRegressedAt sets field value.
func (o *IssueRegression) SetRegressedAt(v time.Time) {
	o.RegressedAt = v
}

// GetRegressedAtVersion returns the RegressedAtVersion field value if set, zero value otherwise.
func (o *IssueRegression) GetRegressedAtVersion() string {
	if o == nil || o.RegressedAtVersion == nil {
		var ret string
		return ret
	}
	return *o.RegressedAtVersion
}

// GetRegressedAtVersionOk returns a tuple with the RegressedAtVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueRegression) GetRegressedAtVersionOk() (*string, bool) {
	if o == nil || o.RegressedAtVersion == nil {
		return nil, false
	}
	return o.RegressedAtVersion, true
}

// HasRegressedAtVersion returns a boolean if a field has been set.
func (o *IssueRegression) HasRegressedAtVersion() bool {
	return o != nil && o.RegressedAtVersion != nil
}

// SetRegressedAtVersion gets a reference to the given string and assigns it to the RegressedAtVersion field.
func (o *IssueRegression) SetRegressedAtVersion(v string) {
	o.RegressedAtVersion = &v
}

// GetResolvedAt returns the ResolvedAt field value.
func (o *IssueRegression) GetResolvedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ResolvedAt
}

// GetResolvedAtOk returns a tuple with the ResolvedAt field value
// and a boolean to check if the value has been set.
func (o *IssueRegression) GetResolvedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResolvedAt, true
}

// SetResolvedAt sets field value.
func (o *IssueRegression) SetResolvedAt(v time.Time) {
	o.ResolvedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueRegression) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RegressedAt.Nanosecond() == 0 {
		toSerialize["regressed_at"] = o.RegressedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["regressed_at"] = o.RegressedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.RegressedAtVersion != nil {
		toSerialize["regressed_at_version"] = o.RegressedAtVersion
	}
	if o.ResolvedAt.Nanosecond() == 0 {
		toSerialize["resolved_at"] = o.ResolvedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["resolved_at"] = o.ResolvedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueRegression) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RegressedAt        *time.Time `json:"regressed_at"`
		RegressedAtVersion *string    `json:"regressed_at_version,omitempty"`
		ResolvedAt         *time.Time `json:"resolved_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RegressedAt == nil {
		return fmt.Errorf("required field regressed_at missing")
	}
	if all.ResolvedAt == nil {
		return fmt.Errorf("required field resolved_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"regressed_at", "regressed_at_version", "resolved_at"})
	} else {
		return err
	}
	o.RegressedAt = *all.RegressedAt
	o.RegressedAtVersion = all.RegressedAtVersion
	o.ResolvedAt = *all.ResolvedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
