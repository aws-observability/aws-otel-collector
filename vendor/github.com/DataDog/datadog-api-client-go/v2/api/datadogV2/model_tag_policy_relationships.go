// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagPolicyRelationships Related resources for a tag policy. Only present when the corresponding `include` query parameter is supplied.
type TagPolicyRelationships struct {
	// A relationship to the compliance score resource for this policy.
	Score *TagPolicyScoreRelationship `json:"score,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagPolicyRelationships instantiates a new TagPolicyRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagPolicyRelationships() *TagPolicyRelationships {
	this := TagPolicyRelationships{}
	return &this
}

// NewTagPolicyRelationshipsWithDefaults instantiates a new TagPolicyRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagPolicyRelationshipsWithDefaults() *TagPolicyRelationships {
	this := TagPolicyRelationships{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *TagPolicyRelationships) GetScore() TagPolicyScoreRelationship {
	if o == nil || o.Score == nil {
		var ret TagPolicyScoreRelationship
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPolicyRelationships) GetScoreOk() (*TagPolicyScoreRelationship, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *TagPolicyRelationships) HasScore() bool {
	return o != nil && o.Score != nil
}

// SetScore gets a reference to the given TagPolicyScoreRelationship and assigns it to the Score field.
func (o *TagPolicyRelationships) SetScore(v TagPolicyScoreRelationship) {
	o.Score = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagPolicyRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagPolicyRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Score *TagPolicyScoreRelationship `json:"score,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"score"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Score != nil && all.Score.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Score = all.Score

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
