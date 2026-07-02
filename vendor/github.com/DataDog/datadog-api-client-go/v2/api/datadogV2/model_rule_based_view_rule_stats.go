// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleBasedViewRuleStats Counts of findings for the rule, grouped by their evaluation status.
type RuleBasedViewRuleStats struct {
	// Number of findings that failed evaluation.
	Fail int64 `json:"fail"`
	// Number of findings that have been muted.
	Muted int64 `json:"muted"`
	// Number of findings that passed evaluation.
	Pass int64 `json:"pass"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRuleBasedViewRuleStats instantiates a new RuleBasedViewRuleStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRuleBasedViewRuleStats(fail int64, muted int64, pass int64) *RuleBasedViewRuleStats {
	this := RuleBasedViewRuleStats{}
	this.Fail = fail
	this.Muted = muted
	this.Pass = pass
	return &this
}

// NewRuleBasedViewRuleStatsWithDefaults instantiates a new RuleBasedViewRuleStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRuleBasedViewRuleStatsWithDefaults() *RuleBasedViewRuleStats {
	this := RuleBasedViewRuleStats{}
	return &this
}

// GetFail returns the Fail field value.
func (o *RuleBasedViewRuleStats) GetFail() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Fail
}

// GetFailOk returns a tuple with the Fail field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRuleStats) GetFailOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fail, true
}

// SetFail sets field value.
func (o *RuleBasedViewRuleStats) SetFail(v int64) {
	o.Fail = v
}

// GetMuted returns the Muted field value.
func (o *RuleBasedViewRuleStats) GetMuted() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Muted
}

// GetMutedOk returns a tuple with the Muted field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRuleStats) GetMutedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Muted, true
}

// SetMuted sets field value.
func (o *RuleBasedViewRuleStats) SetMuted(v int64) {
	o.Muted = v
}

// GetPass returns the Pass field value.
func (o *RuleBasedViewRuleStats) GetPass() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pass
}

// GetPassOk returns a tuple with the Pass field value
// and a boolean to check if the value has been set.
func (o *RuleBasedViewRuleStats) GetPassOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pass, true
}

// SetPass sets field value.
func (o *RuleBasedViewRuleStats) SetPass(v int64) {
	o.Pass = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RuleBasedViewRuleStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fail"] = o.Fail
	toSerialize["muted"] = o.Muted
	toSerialize["pass"] = o.Pass

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RuleBasedViewRuleStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fail  *int64 `json:"fail"`
		Muted *int64 `json:"muted"`
		Pass  *int64 `json:"pass"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fail == nil {
		return fmt.Errorf("required field fail missing")
	}
	if all.Muted == nil {
		return fmt.Errorf("required field muted missing")
	}
	if all.Pass == nil {
		return fmt.Errorf("required field pass missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fail", "muted", "pass"})
	} else {
		return err
	}
	o.Fail = *all.Fail
	o.Muted = *all.Muted
	o.Pass = *all.Pass

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
