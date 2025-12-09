// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsGrokParserRules Set of rules for the grok parser.
type LogsGrokParserRules struct {
	// List of match rules for the grok parser, separated by a new line.
	MatchRules string `json:"match_rules"`
	// List of support rules for the grok parser, separated by a new line.
	SupportRules *string `json:"support_rules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsGrokParserRules instantiates a new LogsGrokParserRules object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsGrokParserRules(matchRules string) *LogsGrokParserRules {
	this := LogsGrokParserRules{}
	this.MatchRules = matchRules
	var supportRules string = ""
	this.SupportRules = &supportRules
	return &this
}

// NewLogsGrokParserRulesWithDefaults instantiates a new LogsGrokParserRules object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsGrokParserRulesWithDefaults() *LogsGrokParserRules {
	this := LogsGrokParserRules{}
	var supportRules string = ""
	this.SupportRules = &supportRules
	return &this
}

// GetMatchRules returns the MatchRules field value.
func (o *LogsGrokParserRules) GetMatchRules() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MatchRules
}

// GetMatchRulesOk returns a tuple with the MatchRules field value
// and a boolean to check if the value has been set.
func (o *LogsGrokParserRules) GetMatchRulesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchRules, true
}

// SetMatchRules sets field value.
func (o *LogsGrokParserRules) SetMatchRules(v string) {
	o.MatchRules = v
}

// GetSupportRules returns the SupportRules field value if set, zero value otherwise.
func (o *LogsGrokParserRules) GetSupportRules() string {
	if o == nil || o.SupportRules == nil {
		var ret string
		return ret
	}
	return *o.SupportRules
}

// GetSupportRulesOk returns a tuple with the SupportRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsGrokParserRules) GetSupportRulesOk() (*string, bool) {
	if o == nil || o.SupportRules == nil {
		return nil, false
	}
	return o.SupportRules, true
}

// HasSupportRules returns a boolean if a field has been set.
func (o *LogsGrokParserRules) HasSupportRules() bool {
	return o != nil && o.SupportRules != nil
}

// SetSupportRules gets a reference to the given string and assigns it to the SupportRules field.
func (o *LogsGrokParserRules) SetSupportRules(v string) {
	o.SupportRules = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsGrokParserRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["match_rules"] = o.MatchRules
	if o.SupportRules != nil {
		toSerialize["support_rules"] = o.SupportRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsGrokParserRules) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MatchRules   *string `json:"match_rules"`
		SupportRules *string `json:"support_rules,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MatchRules == nil {
		return fmt.Errorf("required field match_rules missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"match_rules", "support_rules"})
	} else {
		return err
	}
	o.MatchRules = *all.MatchRules
	o.SupportRules = all.SupportRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
