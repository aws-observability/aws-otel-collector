// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateData The object containing the query content.
type SpansAggregateData struct {
	// The object containing all the query parameters.
	Attributes *SpansAggregateRequestAttributes `json:"attributes,omitempty"`
	// The type of resource. The value should always be aggregate_request.
	Type *SpansAggregateRequestType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpansAggregateData instantiates a new SpansAggregateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpansAggregateData() *SpansAggregateData {
	this := SpansAggregateData{}
	var typeVar SpansAggregateRequestType = SPANSAGGREGATEREQUESTTYPE_AGGREGATE_REQUEST
	this.Type = &typeVar
	return &this
}

// NewSpansAggregateDataWithDefaults instantiates a new SpansAggregateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpansAggregateDataWithDefaults() *SpansAggregateData {
	this := SpansAggregateData{}
	var typeVar SpansAggregateRequestType = SPANSAGGREGATEREQUESTTYPE_AGGREGATE_REQUEST
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SpansAggregateData) GetAttributes() SpansAggregateRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret SpansAggregateRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateData) GetAttributesOk() (*SpansAggregateRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SpansAggregateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given SpansAggregateRequestAttributes and assigns it to the Attributes field.
func (o *SpansAggregateData) SetAttributes(v SpansAggregateRequestAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SpansAggregateData) GetType() SpansAggregateRequestType {
	if o == nil || o.Type == nil {
		var ret SpansAggregateRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpansAggregateData) GetTypeOk() (*SpansAggregateRequestType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SpansAggregateData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SpansAggregateRequestType and assigns it to the Type field.
func (o *SpansAggregateData) SetType(v SpansAggregateRequestType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpansAggregateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
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
func (o *SpansAggregateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *SpansAggregateRequestAttributes `json:"attributes,omitempty"`
		Type       *SpansAggregateRequestType       `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
