// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyScoreAttributes Attributes of a tag policy compliance score.
type TagPolicyScoreAttributes struct {
	// The compliance score for the policy over the requested time window, as a percentage
	// between 0 and 100. `null` indicates that no relevant telemetry was found.
	Score datadog.NullableFloat64 `json:"score"`
	// End of the time window the score was computed over, as a Unix timestamp in milliseconds.
	TsEnd int64 `json:"ts_end"`
	// Start of the time window the score was computed over, as a Unix timestamp in milliseconds.
	TsStart int64 `json:"ts_start"`
	// The version of the tag policy that the score was computed against.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPolicyScoreAttributes instantiates a new TagPolicyScoreAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPolicyScoreAttributes(score datadog.NullableFloat64, tsEnd int64, tsStart int64, version int64) *TagPolicyScoreAttributes {
	this := TagPolicyScoreAttributes{}
	this.Score = score
	this.TsEnd = tsEnd
	this.TsStart = tsStart
	this.Version = version
	return &this
}

// NewTagPolicyScoreAttributesWithDefaults instantiates a new TagPolicyScoreAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPolicyScoreAttributesWithDefaults() *TagPolicyScoreAttributes {
	this := TagPolicyScoreAttributes{}
	return &this
}

// GetScore returns the Score field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *TagPolicyScoreAttributes) GetScore() float64 {
	if o == nil || o.Score.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TagPolicyScoreAttributes) GetScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// SetScore sets field value.
func (o *TagPolicyScoreAttributes) SetScore(v float64) {
	o.Score.Set(&v)
}

// GetTsEnd returns the TsEnd field value.
func (o *TagPolicyScoreAttributes) GetTsEnd() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TsEnd
}

// GetTsEndOk returns a tuple with the TsEnd field value
// and a boolean to check if the value has been set.
func (o *TagPolicyScoreAttributes) GetTsEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TsEnd, true
}

// SetTsEnd sets field value.
func (o *TagPolicyScoreAttributes) SetTsEnd(v int64) {
	o.TsEnd = v
}

// GetTsStart returns the TsStart field value.
func (o *TagPolicyScoreAttributes) GetTsStart() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TsStart
}

// GetTsStartOk returns a tuple with the TsStart field value
// and a boolean to check if the value has been set.
func (o *TagPolicyScoreAttributes) GetTsStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TsStart, true
}

// SetTsStart sets field value.
func (o *TagPolicyScoreAttributes) SetTsStart(v int64) {
	o.TsStart = v
}

// GetVersion returns the Version field value.
func (o *TagPolicyScoreAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *TagPolicyScoreAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *TagPolicyScoreAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPolicyScoreAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["score"] = o.Score.Get()
	toSerialize["ts_end"] = o.TsEnd
	toSerialize["ts_start"] = o.TsStart
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagPolicyScoreAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Score   datadog.NullableFloat64 `json:"score"`
		TsEnd   *int64                  `json:"ts_end"`
		TsStart *int64                  `json:"ts_start"`
		Version *int64                  `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Score.IsSet() {
		return fmt.Errorf("required field score missing")
	}
	if all.TsEnd == nil {
		return fmt.Errorf("required field ts_end missing")
	}
	if all.TsStart == nil {
		return fmt.Errorf("required field ts_start missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"score", "ts_end", "ts_start", "version"})
	} else {
		return err
	}
	o.Score = all.Score
	o.TsEnd = *all.TsEnd
	o.TsStart = *all.TsStart
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
