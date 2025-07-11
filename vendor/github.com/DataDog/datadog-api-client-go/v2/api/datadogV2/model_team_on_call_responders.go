// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamOnCallResponders Root object representing a team's on-call responder configuration.
type TeamOnCallResponders struct {
	// Defines the main on-call responder object for a team, including relationships and metadata.
	Data *TeamOnCallRespondersData `json:"data,omitempty"`
	// The `TeamOnCallResponders` `included`.
	Included []TeamOnCallRespondersIncluded `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamOnCallResponders instantiates a new TeamOnCallResponders object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamOnCallResponders() *TeamOnCallResponders {
	this := TeamOnCallResponders{}
	return &this
}

// NewTeamOnCallRespondersWithDefaults instantiates a new TeamOnCallResponders object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamOnCallRespondersWithDefaults() *TeamOnCallResponders {
	this := TeamOnCallResponders{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TeamOnCallResponders) GetData() TeamOnCallRespondersData {
	if o == nil || o.Data == nil {
		var ret TeamOnCallRespondersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamOnCallResponders) GetDataOk() (*TeamOnCallRespondersData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TeamOnCallResponders) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given TeamOnCallRespondersData and assigns it to the Data field.
func (o *TeamOnCallResponders) SetData(v TeamOnCallRespondersData) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *TeamOnCallResponders) GetIncluded() []TeamOnCallRespondersIncluded {
	if o == nil || o.Included == nil {
		var ret []TeamOnCallRespondersIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamOnCallResponders) GetIncludedOk() (*[]TeamOnCallRespondersIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *TeamOnCallResponders) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []TeamOnCallRespondersIncluded and assigns it to the Included field.
func (o *TeamOnCallResponders) SetIncluded(v []TeamOnCallRespondersIncluded) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamOnCallResponders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamOnCallResponders) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *TeamOnCallRespondersData      `json:"data,omitempty"`
		Included []TeamOnCallRespondersIncluded `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
