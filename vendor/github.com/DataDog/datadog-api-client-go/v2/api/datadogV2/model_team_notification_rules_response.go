// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamNotificationRulesResponse Team notification rules response
type TeamNotificationRulesResponse struct {
	// Team notification rules response data
	Data []TeamNotificationRule `json:"data,omitempty"`
	// Metadata that is included in the response when querying the team notification rules
	Meta *TeamNotificationRulesResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamNotificationRulesResponse instantiates a new TeamNotificationRulesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamNotificationRulesResponse() *TeamNotificationRulesResponse {
	this := TeamNotificationRulesResponse{}
	return &this
}

// NewTeamNotificationRulesResponseWithDefaults instantiates a new TeamNotificationRulesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamNotificationRulesResponseWithDefaults() *TeamNotificationRulesResponse {
	this := TeamNotificationRulesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TeamNotificationRulesResponse) GetData() []TeamNotificationRule {
	if o == nil || o.Data == nil {
		var ret []TeamNotificationRule
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRulesResponse) GetDataOk() (*[]TeamNotificationRule, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []TeamNotificationRule and assigns it to the Data field.
func (o *TeamNotificationRulesResponse) SetData(v []TeamNotificationRule) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TeamNotificationRulesResponse) GetMeta() TeamNotificationRulesResponseMeta {
	if o == nil || o.Meta == nil {
		var ret TeamNotificationRulesResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRulesResponse) GetMetaOk() (*TeamNotificationRulesResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TeamNotificationRulesResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given TeamNotificationRulesResponseMeta and assigns it to the Meta field.
func (o *TeamNotificationRulesResponse) SetMeta(v TeamNotificationRulesResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamNotificationRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamNotificationRulesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []TeamNotificationRule             `json:"data,omitempty"`
		Meta *TeamNotificationRulesResponseMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
