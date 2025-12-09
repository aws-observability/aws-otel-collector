// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamLinkAttributes Team link attributes
type TeamLinkAttributes struct {
	// The link's label
	Label string `json:"label"`
	// The link's position, used to sort links for the team
	Position *int32 `json:"position,omitempty"`
	// ID of the team the link is associated with
	TeamId *string `json:"team_id,omitempty"`
	// The URL for the link
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamLinkAttributes instantiates a new TeamLinkAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamLinkAttributes(label string, url string) *TeamLinkAttributes {
	this := TeamLinkAttributes{}
	this.Label = label
	this.Url = url
	return &this
}

// NewTeamLinkAttributesWithDefaults instantiates a new TeamLinkAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamLinkAttributesWithDefaults() *TeamLinkAttributes {
	this := TeamLinkAttributes{}
	return &this
}

// GetLabel returns the Label field value.
func (o *TeamLinkAttributes) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *TeamLinkAttributes) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *TeamLinkAttributes) SetLabel(v string) {
	o.Label = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *TeamLinkAttributes) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamLinkAttributes) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *TeamLinkAttributes) HasPosition() bool {
	return o != nil && o.Position != nil
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *TeamLinkAttributes) SetPosition(v int32) {
	o.Position = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *TeamLinkAttributes) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamLinkAttributes) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *TeamLinkAttributes) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *TeamLinkAttributes) SetTeamId(v string) {
	o.TeamId = &v
}

// GetUrl returns the Url field value.
func (o *TeamLinkAttributes) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TeamLinkAttributes) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *TeamLinkAttributes) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamLinkAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["label"] = o.Label
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamLinkAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label    *string `json:"label"`
		Position *int32  `json:"position,omitempty"`
		TeamId   *string `json:"team_id,omitempty"`
		Url      *string `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"label", "position", "team_id", "url"})
	} else {
		return err
	}
	o.Label = *all.Label
	o.Position = all.Position
	o.TeamId = all.TeamId
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
