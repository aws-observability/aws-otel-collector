// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OutcomesResponse Scorecard outcomes - the result of a rule for a service.
type OutcomesResponse struct {
	// List of rule outcomes.
	Data []OutcomesResponseDataItem `json:"data,omitempty"`
	// Array of rule details.
	Included []OutcomesResponseIncludedItem `json:"included,omitempty"`
	// Links attributes.
	Links *OutcomesResponseLinks `json:"links,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOutcomesResponse instantiates a new OutcomesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutcomesResponse() *OutcomesResponse {
	this := OutcomesResponse{}
	return &this
}

// NewOutcomesResponseWithDefaults instantiates a new OutcomesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutcomesResponseWithDefaults() *OutcomesResponse {
	this := OutcomesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OutcomesResponse) GetData() []OutcomesResponseDataItem {
	if o == nil || o.Data == nil {
		var ret []OutcomesResponseDataItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesResponse) GetDataOk() (*[]OutcomesResponseDataItem, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OutcomesResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []OutcomesResponseDataItem and assigns it to the Data field.
func (o *OutcomesResponse) SetData(v []OutcomesResponseDataItem) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *OutcomesResponse) GetIncluded() []OutcomesResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []OutcomesResponseIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesResponse) GetIncludedOk() (*[]OutcomesResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *OutcomesResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []OutcomesResponseIncludedItem and assigns it to the Included field.
func (o *OutcomesResponse) SetIncluded(v []OutcomesResponseIncludedItem) {
	o.Included = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OutcomesResponse) GetLinks() OutcomesResponseLinks {
	if o == nil || o.Links == nil {
		var ret OutcomesResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcomesResponse) GetLinksOk() (*OutcomesResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OutcomesResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given OutcomesResponseLinks and assigns it to the Links field.
func (o *OutcomesResponse) SetLinks(v OutcomesResponseLinks) {
	o.Links = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutcomesResponse) MarshalJSON() ([]byte, error) {
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
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OutcomesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []OutcomesResponseDataItem     `json:"data,omitempty"`
		Included []OutcomesResponseIncludedItem `json:"included,omitempty"`
		Links    *OutcomesResponseLinks         `json:"links,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "links"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	o.Included = all.Included
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
