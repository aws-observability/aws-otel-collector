// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpsertOAuthScopesRestrictionData Data object of an upsert OAuth2 scopes restriction request.
type UpsertOAuthScopesRestrictionData struct {
	// Attributes of an upsert OAuth2 scopes restriction request.
	Attributes *UpsertOAuthScopesRestrictionDataAttributes `json:"attributes,omitempty"`
	// JSON:API resource type for an upsert OAuth2 client scopes restriction request.
	Type UpsertOAuthScopesRestrictionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpsertOAuthScopesRestrictionData instantiates a new UpsertOAuthScopesRestrictionData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpsertOAuthScopesRestrictionData(typeVar UpsertOAuthScopesRestrictionType) *UpsertOAuthScopesRestrictionData {
	this := UpsertOAuthScopesRestrictionData{}
	this.Type = typeVar
	return &this
}

// NewUpsertOAuthScopesRestrictionDataWithDefaults instantiates a new UpsertOAuthScopesRestrictionData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpsertOAuthScopesRestrictionDataWithDefaults() *UpsertOAuthScopesRestrictionData {
	this := UpsertOAuthScopesRestrictionData{}
	var typeVar UpsertOAuthScopesRestrictionType = UPSERTOAUTHSCOPESRESTRICTIONTYPE_UPSERT_SCOPES_RESTRICTION
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpsertOAuthScopesRestrictionData) GetAttributes() UpsertOAuthScopesRestrictionDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret UpsertOAuthScopesRestrictionDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpsertOAuthScopesRestrictionData) GetAttributesOk() (*UpsertOAuthScopesRestrictionDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpsertOAuthScopesRestrictionData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UpsertOAuthScopesRestrictionDataAttributes and assigns it to the Attributes field.
func (o *UpsertOAuthScopesRestrictionData) SetAttributes(v UpsertOAuthScopesRestrictionDataAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *UpsertOAuthScopesRestrictionData) GetType() UpsertOAuthScopesRestrictionType {
	if o == nil {
		var ret UpsertOAuthScopesRestrictionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpsertOAuthScopesRestrictionData) GetTypeOk() (*UpsertOAuthScopesRestrictionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *UpsertOAuthScopesRestrictionData) SetType(v UpsertOAuthScopesRestrictionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpsertOAuthScopesRestrictionData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpsertOAuthScopesRestrictionData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UpsertOAuthScopesRestrictionDataAttributes `json:"attributes,omitempty"`
		Type       *UpsertOAuthScopesRestrictionType           `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
