// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetAppResponse The definition of `GetAppResponse` object.
type GetAppResponse struct {
	// The definition of `GetAppResponseData` object.
	Data *GetAppResponseData `json:"data,omitempty"`
	// The `GetAppResponse` `included`.
	Included []DeploymentIncluded `json:"included,omitempty"`
	// The definition of `AppMeta` object.
	Meta *AppMeta `json:"meta,omitempty"`
	// The definition of `GetAppResponseRelationship` object.
	Relationship *GetAppResponseRelationship `json:"relationship,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetAppResponse instantiates a new GetAppResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetAppResponse() *GetAppResponse {
	this := GetAppResponse{}
	return &this
}

// NewGetAppResponseWithDefaults instantiates a new GetAppResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetAppResponseWithDefaults() *GetAppResponse {
	this := GetAppResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAppResponse) GetData() GetAppResponseData {
	if o == nil || o.Data == nil {
		var ret GetAppResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppResponse) GetDataOk() (*GetAppResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAppResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given GetAppResponseData and assigns it to the Data field.
func (o *GetAppResponse) SetData(v GetAppResponseData) {
	o.Data = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetAppResponse) GetIncluded() []DeploymentIncluded {
	if o == nil || o.Included == nil {
		var ret []DeploymentIncluded
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppResponse) GetIncludedOk() (*[]DeploymentIncluded, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetAppResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []DeploymentIncluded and assigns it to the Included field.
func (o *GetAppResponse) SetIncluded(v []DeploymentIncluded) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetAppResponse) GetMeta() AppMeta {
	if o == nil || o.Meta == nil {
		var ret AppMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppResponse) GetMetaOk() (*AppMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetAppResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given AppMeta and assigns it to the Meta field.
func (o *GetAppResponse) SetMeta(v AppMeta) {
	o.Meta = &v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *GetAppResponse) GetRelationship() GetAppResponseRelationship {
	if o == nil || o.Relationship == nil {
		var ret GetAppResponseRelationship
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppResponse) GetRelationshipOk() (*GetAppResponseRelationship, bool) {
	if o == nil || o.Relationship == nil {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *GetAppResponse) HasRelationship() bool {
	return o != nil && o.Relationship != nil
}

// SetRelationship gets a reference to the given GetAppResponseRelationship and assigns it to the Relationship field.
func (o *GetAppResponse) SetRelationship(v GetAppResponseRelationship) {
	o.Relationship = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetAppResponse) MarshalJSON() ([]byte, error) {
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
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Relationship != nil {
		toSerialize["relationship"] = o.Relationship
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetAppResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data         *GetAppResponseData         `json:"data,omitempty"`
		Included     []DeploymentIncluded        `json:"included,omitempty"`
		Meta         *AppMeta                    `json:"meta,omitempty"`
		Relationship *GetAppResponseRelationship `json:"relationship,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta", "relationship"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.Included = all.Included
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	if all.Relationship != nil && all.Relationship.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationship = all.Relationship

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
