// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseAttributes The attributes associated with the custom destination.
type CustomDestinationResponseAttributes struct {
	// Whether logs matching this custom destination should be forwarded or not.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether tags from the forwarded logs should be forwarded or not.
	ForwardTags *bool `json:"forward_tags,omitempty"`
	// List of [keys of tags](https://docs.datadoghq.com/getting_started/tagging/#define-tags) to be filtered.
	//
	// An empty list represents no restriction is in place and either all or no tags will be
	// forwarded depending on `forward_tags_restriction_list_type` parameter.
	ForwardTagsRestrictionList []string `json:"forward_tags_restriction_list,omitempty"`
	// How `forward_tags_restriction_list` parameter should be interpreted.
	// If `ALLOW_LIST`, then only tags whose keys on the forwarded logs match the ones on the restriction list
	// are forwarded.
	//
	// `BLOCK_LIST` works the opposite way. It does not forward the tags matching the ones on the list.
	ForwardTagsRestrictionListType *CustomDestinationAttributeTagsRestrictionListType `json:"forward_tags_restriction_list_type,omitempty"`
	// A custom destination's location to forward logs.
	ForwarderDestination *CustomDestinationResponseForwardDestination `json:"forwarder_destination,omitempty"`
	// The custom destination name.
	Name *string `json:"name,omitempty"`
	// The custom destination query filter. Logs matching this query are forwarded to the destination.
	Query *string `json:"query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationResponseAttributes instantiates a new CustomDestinationResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationResponseAttributes() *CustomDestinationResponseAttributes {
	this := CustomDestinationResponseAttributes{}
	var enabled bool = true
	this.Enabled = &enabled
	var forwardTags bool = true
	this.ForwardTags = &forwardTags
	var forwardTagsRestrictionListType CustomDestinationAttributeTagsRestrictionListType = CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_ALLOW_LIST
	this.ForwardTagsRestrictionListType = &forwardTagsRestrictionListType
	var query string = ""
	this.Query = &query
	return &this
}

// NewCustomDestinationResponseAttributesWithDefaults instantiates a new CustomDestinationResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationResponseAttributesWithDefaults() *CustomDestinationResponseAttributes {
	this := CustomDestinationResponseAttributes{}
	var enabled bool = true
	this.Enabled = &enabled
	var forwardTags bool = true
	this.ForwardTags = &forwardTags
	var forwardTagsRestrictionListType CustomDestinationAttributeTagsRestrictionListType = CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_ALLOW_LIST
	this.ForwardTagsRestrictionListType = &forwardTagsRestrictionListType
	var query string = ""
	this.Query = &query
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CustomDestinationResponseAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CustomDestinationResponseAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CustomDestinationResponseAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetForwardTags returns the ForwardTags field value if set, zero value otherwise.
func (o *CustomDestinationResponseAttributes) GetForwardTags() bool {
	if o == nil || o.ForwardTags == nil {
		var ret bool
		return ret
	}
	return *o.ForwardTags
}

// GetForwardTagsOk returns a tuple with the ForwardTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseAttributes) GetForwardTagsOk() (*bool, bool) {
	if o == nil || o.ForwardTags == nil {
		return nil, false
	}
	return o.ForwardTags, true
}

// HasForwardTags returns a boolean if a field has been set.
func (o *CustomDestinationResponseAttributes) HasForwardTags() bool {
	return o != nil && o.ForwardTags != nil
}

// SetForwardTags gets a reference to the given bool and assigns it to the ForwardTags field.
func (o *CustomDestinationResponseAttributes) SetForwardTags(v bool) {
	o.ForwardTags = &v
}

// GetForwardTagsRestrictionList returns the ForwardTagsRestrictionList field value if set, zero value otherwise.
func (o *CustomDestinationResponseAttributes) GetForwardTagsRestrictionList() []string {
	if o == nil || o.ForwardTagsRestrictionList == nil {
		var ret []string
		return ret
	}
	return o.ForwardTagsRestrictionList
}

// GetForwardTagsRestrictionListOk returns a tuple with the ForwardTagsRestrictionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseAttributes) GetForwardTagsRestrictionListOk() (*[]string, bool) {
	if o == nil || o.ForwardTagsRestrictionList == nil {
		return nil, false
	}
	return &o.ForwardTagsRestrictionList, true
}

// HasForwardTagsRestrictionList returns a boolean if a field has been set.
func (o *CustomDestinationResponseAttributes) HasForwardTagsRestrictionList() bool {
	return o != nil && o.ForwardTagsRestrictionList != nil
}

// SetForwardTagsRestrictionList gets a reference to the given []string and assigns it to the ForwardTagsRestrictionList field.
func (o *CustomDestinationResponseAttributes) SetForwardTagsRestrictionList(v []string) {
	o.ForwardTagsRestrictionList = v
}

// GetForwardTagsRestrictionListType returns the ForwardTagsRestrictionListType field value if set, zero value otherwise.
func (o *CustomDestinationResponseAttributes) GetForwardTagsRestrictionListType() CustomDestinationAttributeTagsRestrictionListType {
	if o == nil || o.ForwardTagsRestrictionListType == nil {
		var ret CustomDestinationAttributeTagsRestrictionListType
		return ret
	}
	return *o.ForwardTagsRestrictionListType
}

// GetForwardTagsRestrictionListTypeOk returns a tuple with the ForwardTagsRestrictionListType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseAttributes) GetForwardTagsRestrictionListTypeOk() (*CustomDestinationAttributeTagsRestrictionListType, bool) {
	if o == nil || o.ForwardTagsRestrictionListType == nil {
		return nil, false
	}
	return o.ForwardTagsRestrictionListType, true
}

// HasForwardTagsRestrictionListType returns a boolean if a field has been set.
func (o *CustomDestinationResponseAttributes) HasForwardTagsRestrictionListType() bool {
	return o != nil && o.ForwardTagsRestrictionListType != nil
}

// SetForwardTagsRestrictionListType gets a reference to the given CustomDestinationAttributeTagsRestrictionListType and assigns it to the ForwardTagsRestrictionListType field.
func (o *CustomDestinationResponseAttributes) SetForwardTagsRestrictionListType(v CustomDestinationAttributeTagsRestrictionListType) {
	o.ForwardTagsRestrictionListType = &v
}

// GetForwarderDestination returns the ForwarderDestination field value if set, zero value otherwise.
func (o *CustomDestinationResponseAttributes) GetForwarderDestination() CustomDestinationResponseForwardDestination {
	if o == nil || o.ForwarderDestination == nil {
		var ret CustomDestinationResponseForwardDestination
		return ret
	}
	return *o.ForwarderDestination
}

// GetForwarderDestinationOk returns a tuple with the ForwarderDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseAttributes) GetForwarderDestinationOk() (*CustomDestinationResponseForwardDestination, bool) {
	if o == nil || o.ForwarderDestination == nil {
		return nil, false
	}
	return o.ForwarderDestination, true
}

// HasForwarderDestination returns a boolean if a field has been set.
func (o *CustomDestinationResponseAttributes) HasForwarderDestination() bool {
	return o != nil && o.ForwarderDestination != nil
}

// SetForwarderDestination gets a reference to the given CustomDestinationResponseForwardDestination and assigns it to the ForwarderDestination field.
func (o *CustomDestinationResponseAttributes) SetForwarderDestination(v CustomDestinationResponseForwardDestination) {
	o.ForwarderDestination = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomDestinationResponseAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomDestinationResponseAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomDestinationResponseAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *CustomDestinationResponseAttributes) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseAttributes) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *CustomDestinationResponseAttributes) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *CustomDestinationResponseAttributes) SetQuery(v string) {
	o.Query = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ForwardTags != nil {
		toSerialize["forward_tags"] = o.ForwardTags
	}
	if o.ForwardTagsRestrictionList != nil {
		toSerialize["forward_tags_restriction_list"] = o.ForwardTagsRestrictionList
	}
	if o.ForwardTagsRestrictionListType != nil {
		toSerialize["forward_tags_restriction_list_type"] = o.ForwardTagsRestrictionListType
	}
	if o.ForwarderDestination != nil {
		toSerialize["forwarder_destination"] = o.ForwarderDestination
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled                        *bool                                              `json:"enabled,omitempty"`
		ForwardTags                    *bool                                              `json:"forward_tags,omitempty"`
		ForwardTagsRestrictionList     []string                                           `json:"forward_tags_restriction_list,omitempty"`
		ForwardTagsRestrictionListType *CustomDestinationAttributeTagsRestrictionListType `json:"forward_tags_restriction_list_type,omitempty"`
		ForwarderDestination           *CustomDestinationResponseForwardDestination       `json:"forwarder_destination,omitempty"`
		Name                           *string                                            `json:"name,omitempty"`
		Query                          *string                                            `json:"query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "forward_tags", "forward_tags_restriction_list", "forward_tags_restriction_list_type", "forwarder_destination", "name", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	o.ForwardTags = all.ForwardTags
	o.ForwardTagsRestrictionList = all.ForwardTagsRestrictionList
	if all.ForwardTagsRestrictionListType != nil && !all.ForwardTagsRestrictionListType.IsValid() {
		hasInvalidField = true
	} else {
		o.ForwardTagsRestrictionListType = all.ForwardTagsRestrictionListType
	}
	o.ForwarderDestination = all.ForwarderDestination
	o.Name = all.Name
	o.Query = all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
