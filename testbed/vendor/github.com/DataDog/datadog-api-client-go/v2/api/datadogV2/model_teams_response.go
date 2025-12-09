// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamsResponse Response with multiple teams
type TeamsResponse struct {
	// Teams response data
	Data []Team `json:"data,omitempty"`
	// Resources related to the team
	Included []TeamIncluded `json:"included,omitempty"`
	// Teams response links.
	Links *TeamsResponseLinks `json:"links,omitempty"`
	// Teams response metadata.
	Meta *TeamsResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamsResponse instantiates a new TeamsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamsResponse() *TeamsResponse {
	this := TeamsResponse{}
	return &this
}

// NewTeamsResponseWithDefaults instantiates a new TeamsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamsResponseWithDefaults() *TeamsResponse {
	this := TeamsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TeamsResponse) GetData() []Team {
	if o == nil || o.Data == nil {
		var ret []Team
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponse) GetDataOk() (*[]Team, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TeamsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []Team and assigns it to the Data field.
func (o *TeamsResponse) SetData(v []Team) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *TeamsResponse) GetIncluded() []TeamIncluded {
	if o == nil || o.Included == nil {
		var ret []TeamIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponse) GetIncludedOk() (*[]TeamIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *TeamsResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []TeamIncluded and assigns it to the Included field.
func (o *TeamsResponse) SetIncluded(v []TeamIncluded) {
	o.Included = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TeamsResponse) GetLinks() TeamsResponseLinks {
	if o == nil || o.Links == nil {
		var ret TeamsResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponse) GetLinksOk() (*TeamsResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TeamsResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given TeamsResponseLinks and assigns it to the Links field.
func (o *TeamsResponse) SetLinks(v TeamsResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TeamsResponse) GetMeta() TeamsResponseMeta {
	if o == nil || o.Meta == nil {
		var ret TeamsResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsResponse) GetMetaOk() (*TeamsResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TeamsResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given TeamsResponseMeta and assigns it to the Meta field.
func (o *TeamsResponse) SetMeta(v TeamsResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamsResponse) MarshalJSON() ([]byte, error) {
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
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []Team              `json:"data,omitempty"`
		Included []TeamIncluded      `json:"included,omitempty"`
		Links    *TeamsResponseLinks `json:"links,omitempty"`
		Meta     *TeamsResponseMeta  `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "links", "meta"})
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
