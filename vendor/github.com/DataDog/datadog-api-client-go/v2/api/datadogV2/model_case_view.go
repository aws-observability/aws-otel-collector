// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseView A saved case view that provides a filtered, reusable list of cases matching a specific query. Views act as persistent dashboards for monitoring case subsets.
type CaseView struct {
	// Attributes of a case view, including the filter query and optional notification rule.
	Attributes CaseViewAttributes `json:"attributes"`
	// The view's identifier.
	Id string `json:"id"`
	// Related resources for the case view, including the creator, last modifier, and associated project.
	Relationships *CaseViewRelationships `json:"relationships,omitempty"`
	// JSON:API resource type for case views.
	Type CaseViewResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseView instantiates a new CaseView object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseView(attributes CaseViewAttributes, id string, typeVar CaseViewResourceType) *CaseView {
	this := CaseView{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCaseViewWithDefaults instantiates a new CaseView object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseViewWithDefaults() *CaseView {
	this := CaseView{}
	var typeVar CaseViewResourceType = CASEVIEWRESOURCETYPE_VIEW
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *CaseView) GetAttributes() CaseViewAttributes {
	if o == nil {
		var ret CaseViewAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CaseView) GetAttributesOk() (*CaseViewAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *CaseView) SetAttributes(v CaseViewAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *CaseView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CaseView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CaseView) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CaseView) GetRelationships() CaseViewRelationships {
	if o == nil || o.Relationships == nil {
		var ret CaseViewRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseView) GetRelationshipsOk() (*CaseViewRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CaseView) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given CaseViewRelationships and assigns it to the Relationships field.
func (o *CaseView) SetRelationships(v CaseViewRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *CaseView) GetType() CaseViewResourceType {
	if o == nil {
		var ret CaseViewResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaseView) GetTypeOk() (*CaseViewResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CaseView) SetType(v CaseViewResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseView) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *CaseViewAttributes    `json:"attributes"`
		Id            *string                `json:"id"`
		Relationships *CaseViewRelationships `json:"relationships,omitempty"`
		Type          *CaseViewResourceType  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
