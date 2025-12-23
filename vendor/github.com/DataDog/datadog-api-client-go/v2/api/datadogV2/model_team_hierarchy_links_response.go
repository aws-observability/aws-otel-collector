// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyLinksResponse Team hierarchy links response
type TeamHierarchyLinksResponse struct {
	// Team hierarchy links response data
	Data []TeamHierarchyLink `json:"data,omitempty"`
	// Included teams
	Included []TeamHierarchyLinkTeam `json:"included,omitempty"`
	// When querying team hierarchy links, a set of links for navigation between different pages is included
	Links *TeamsHierarchyLinksResponseLinks `json:"links,omitempty"`
	// Metadata that is included in the response when querying the team hierarchy links
	Meta *TeamsHierarchyLinksResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyLinksResponse instantiates a new TeamHierarchyLinksResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyLinksResponse() *TeamHierarchyLinksResponse {
	this := TeamHierarchyLinksResponse{}
	return &this
}

// NewTeamHierarchyLinksResponseWithDefaults instantiates a new TeamHierarchyLinksResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyLinksResponseWithDefaults() *TeamHierarchyLinksResponse {
	this := TeamHierarchyLinksResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TeamHierarchyLinksResponse) GetData() []TeamHierarchyLink {
	if o == nil || o.Data == nil {
		var ret []TeamHierarchyLink
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinksResponse) GetDataOk() (*[]TeamHierarchyLink, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TeamHierarchyLinksResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []TeamHierarchyLink and assigns it to the Data field.
func (o *TeamHierarchyLinksResponse) SetData(v []TeamHierarchyLink) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *TeamHierarchyLinksResponse) GetIncluded() []TeamHierarchyLinkTeam {
	if o == nil || o.Included == nil {
		var ret []TeamHierarchyLinkTeam
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinksResponse) GetIncludedOk() (*[]TeamHierarchyLinkTeam, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *TeamHierarchyLinksResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []TeamHierarchyLinkTeam and assigns it to the Included field.
func (o *TeamHierarchyLinksResponse) SetIncluded(v []TeamHierarchyLinkTeam) {
	o.Included = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TeamHierarchyLinksResponse) GetLinks() TeamsHierarchyLinksResponseLinks {
	if o == nil || o.Links == nil {
		var ret TeamsHierarchyLinksResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinksResponse) GetLinksOk() (*TeamsHierarchyLinksResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TeamHierarchyLinksResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given TeamsHierarchyLinksResponseLinks and assigns it to the Links field.
func (o *TeamHierarchyLinksResponse) SetLinks(v TeamsHierarchyLinksResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TeamHierarchyLinksResponse) GetMeta() TeamsHierarchyLinksResponseMeta {
	if o == nil || o.Meta == nil {
		var ret TeamsHierarchyLinksResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinksResponse) GetMetaOk() (*TeamsHierarchyLinksResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TeamHierarchyLinksResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given TeamsHierarchyLinksResponseMeta and assigns it to the Meta field.
func (o *TeamHierarchyLinksResponse) SetMeta(v TeamsHierarchyLinksResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyLinksResponse) MarshalJSON() ([]byte, error) {
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
func (o *TeamHierarchyLinksResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []TeamHierarchyLink               `json:"data,omitempty"`
		Included []TeamHierarchyLinkTeam           `json:"included,omitempty"`
		Links    *TeamsHierarchyLinksResponseLinks `json:"links,omitempty"`
		Meta     *TeamsHierarchyLinksResponseMeta  `json:"meta,omitempty"`
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
