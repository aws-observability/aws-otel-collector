// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyLinkResponse Team hierarchy link response
type TeamHierarchyLinkResponse struct {
	// Team hierarchy link
	Data *TeamHierarchyLink `json:"data,omitempty"`
	// Included teams
	Included []TeamHierarchyLinkTeam `json:"included,omitempty"`
	// When querying team hierarchy links, a set of links for navigation between different pages is included
	Links *TeamsHierarchyLinksResponseLinks `json:"links,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyLinkResponse instantiates a new TeamHierarchyLinkResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyLinkResponse() *TeamHierarchyLinkResponse {
	this := TeamHierarchyLinkResponse{}
	return &this
}

// NewTeamHierarchyLinkResponseWithDefaults instantiates a new TeamHierarchyLinkResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyLinkResponseWithDefaults() *TeamHierarchyLinkResponse {
	this := TeamHierarchyLinkResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TeamHierarchyLinkResponse) GetData() TeamHierarchyLink {
	if o == nil || o.Data == nil {
		var ret TeamHierarchyLink
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkResponse) GetDataOk() (*TeamHierarchyLink, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TeamHierarchyLinkResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given TeamHierarchyLink and assigns it to the Data field.
func (o *TeamHierarchyLinkResponse) SetData(v TeamHierarchyLink) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *TeamHierarchyLinkResponse) GetIncluded() []TeamHierarchyLinkTeam {
	if o == nil || o.Included == nil {
		var ret []TeamHierarchyLinkTeam
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkResponse) GetIncludedOk() (*[]TeamHierarchyLinkTeam, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *TeamHierarchyLinkResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []TeamHierarchyLinkTeam and assigns it to the Included field.
func (o *TeamHierarchyLinkResponse) SetIncluded(v []TeamHierarchyLinkTeam) {
	o.Included = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TeamHierarchyLinkResponse) GetLinks() TeamsHierarchyLinksResponseLinks {
	if o == nil || o.Links == nil {
		var ret TeamsHierarchyLinksResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamHierarchyLinkResponse) GetLinksOk() (*TeamsHierarchyLinksResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TeamHierarchyLinkResponse) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given TeamsHierarchyLinksResponseLinks and assigns it to the Links field.
func (o *TeamHierarchyLinkResponse) SetLinks(v TeamsHierarchyLinksResponseLinks) {
	o.Links = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyLinkResponse) MarshalJSON() ([]byte, error) {
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
func (o *TeamHierarchyLinkResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *TeamHierarchyLink                `json:"data,omitempty"`
		Included []TeamHierarchyLinkTeam           `json:"included,omitempty"`
		Links    *TeamsHierarchyLinksResponseLinks `json:"links,omitempty"`
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
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
