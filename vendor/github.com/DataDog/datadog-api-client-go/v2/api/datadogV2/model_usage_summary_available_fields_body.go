// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSummaryAvailableFieldsBody Available-fields data.
type UsageSummaryAvailableFieldsBody struct {
	// The lists of field names returned by `GET /api/v1/usage/summary` at each
	// of its three response levels. Each list contains every key the data endpoint
	// emits—both typed fields declared in the OpenAPI spec and untyped keys
	// exposed through `additionalProperties`.
	Attributes *UsageSummaryAvailableFieldsAttributes `json:"attributes,omitempty"`
	// The identifier for the discovery scope. Always `"all"`.
	Id *string `json:"id,omitempty"`
	// Type of available-fields data.
	Type *UsageSummaryAvailableFieldsType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageSummaryAvailableFieldsBody instantiates a new UsageSummaryAvailableFieldsBody object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSummaryAvailableFieldsBody() *UsageSummaryAvailableFieldsBody {
	this := UsageSummaryAvailableFieldsBody{}
	var typeVar UsageSummaryAvailableFieldsType = USAGESUMMARYAVAILABLEFIELDSTYPE_USAGE_SUMMARY_AVAILABLE_FIELDS
	this.Type = &typeVar
	return &this
}

// NewUsageSummaryAvailableFieldsBodyWithDefaults instantiates a new UsageSummaryAvailableFieldsBody object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSummaryAvailableFieldsBodyWithDefaults() *UsageSummaryAvailableFieldsBody {
	this := UsageSummaryAvailableFieldsBody{}
	var typeVar UsageSummaryAvailableFieldsType = USAGESUMMARYAVAILABLEFIELDSTYPE_USAGE_SUMMARY_AVAILABLE_FIELDS
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UsageSummaryAvailableFieldsBody) GetAttributes() UsageSummaryAvailableFieldsAttributes {
	if o == nil || o.Attributes == nil {
		var ret UsageSummaryAvailableFieldsAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryAvailableFieldsBody) GetAttributesOk() (*UsageSummaryAvailableFieldsAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UsageSummaryAvailableFieldsBody) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given UsageSummaryAvailableFieldsAttributes and assigns it to the Attributes field.
func (o *UsageSummaryAvailableFieldsBody) SetAttributes(v UsageSummaryAvailableFieldsAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UsageSummaryAvailableFieldsBody) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryAvailableFieldsBody) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UsageSummaryAvailableFieldsBody) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UsageSummaryAvailableFieldsBody) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UsageSummaryAvailableFieldsBody) GetType() UsageSummaryAvailableFieldsType {
	if o == nil || o.Type == nil {
		var ret UsageSummaryAvailableFieldsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSummaryAvailableFieldsBody) GetTypeOk() (*UsageSummaryAvailableFieldsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UsageSummaryAvailableFieldsBody) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given UsageSummaryAvailableFieldsType and assigns it to the Type field.
func (o *UsageSummaryAvailableFieldsBody) SetType(v UsageSummaryAvailableFieldsType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSummaryAvailableFieldsBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageSummaryAvailableFieldsBody) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *UsageSummaryAvailableFieldsAttributes `json:"attributes,omitempty"`
		Id         *string                                `json:"id,omitempty"`
		Type       *UsageSummaryAvailableFieldsType       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
