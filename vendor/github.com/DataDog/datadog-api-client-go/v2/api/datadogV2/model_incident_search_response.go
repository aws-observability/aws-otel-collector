// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentSearchResponse Response with incidents and facets.
type IncidentSearchResponse struct {
	// Data returned by an incident search.
	Data IncidentSearchResponseData `json:"data"`
	// Included related resources that the user requested.
	Included []IncidentResponseIncludedItem `json:"included,omitempty"`
	// The metadata object containing pagination metadata.
	Meta *IncidentSearchResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentSearchResponse instantiates a new IncidentSearchResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentSearchResponse(data IncidentSearchResponseData) *IncidentSearchResponse {
	this := IncidentSearchResponse{}
	this.Data = data
	return &this
}

// NewIncidentSearchResponseWithDefaults instantiates a new IncidentSearchResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentSearchResponseWithDefaults() *IncidentSearchResponse {
	this := IncidentSearchResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentSearchResponse) GetData() IncidentSearchResponseData {
	if o == nil {
		var ret IncidentSearchResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponse) GetDataOk() (*IncidentSearchResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentSearchResponse) SetData(v IncidentSearchResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentSearchResponse) GetIncluded() []IncidentResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []IncidentResponseIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponse) GetIncludedOk() (*[]IncidentResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentSearchResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentResponseIncludedItem and assigns it to the Included field.
func (o *IncidentSearchResponse) SetIncluded(v []IncidentResponseIncludedItem) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentSearchResponse) GetMeta() IncidentSearchResponseMeta {
	if o == nil || o.Meta == nil {
		var ret IncidentSearchResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentSearchResponse) GetMetaOk() (*IncidentSearchResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentSearchResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given IncidentSearchResponseMeta and assigns it to the Meta field.
func (o *IncidentSearchResponse) SetMeta(v IncidentSearchResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
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
func (o *IncidentSearchResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *IncidentSearchResponseData    `json:"data"`
		Included []IncidentResponseIncludedItem `json:"included,omitempty"`
		Meta     *IncidentSearchResponseMeta    `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	o.Included = all.Included
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
