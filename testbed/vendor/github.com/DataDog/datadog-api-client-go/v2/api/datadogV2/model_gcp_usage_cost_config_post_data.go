// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPUsageCostConfigPostData Google Cloud Usage Cost config post data.
type GCPUsageCostConfigPostData struct {
	// Attributes for Google Cloud Usage Cost config post request.
	Attributes *GCPUsageCostConfigPostRequestAttributes `json:"attributes,omitempty"`
	// Type of Google Cloud Usage Cost config post request.
	Type GCPUsageCostConfigPostRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPUsageCostConfigPostData instantiates a new GCPUsageCostConfigPostData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPUsageCostConfigPostData(typeVar GCPUsageCostConfigPostRequestType) *GCPUsageCostConfigPostData {
	this := GCPUsageCostConfigPostData{}
	this.Type = typeVar
	return &this
}

// NewGCPUsageCostConfigPostDataWithDefaults instantiates a new GCPUsageCostConfigPostData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPUsageCostConfigPostDataWithDefaults() *GCPUsageCostConfigPostData {
	this := GCPUsageCostConfigPostData{}
	var typeVar GCPUsageCostConfigPostRequestType = GCPUSAGECOSTCONFIGPOSTREQUESTTYPE_GCP_USAGE_COST_CONFIG_POST_REQUEST
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GCPUsageCostConfigPostData) GetAttributes() GCPUsageCostConfigPostRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret GCPUsageCostConfigPostRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostData) GetAttributesOk() (*GCPUsageCostConfigPostRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GCPUsageCostConfigPostData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given GCPUsageCostConfigPostRequestAttributes and assigns it to the Attributes field.
func (o *GCPUsageCostConfigPostData) SetAttributes(v GCPUsageCostConfigPostRequestAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *GCPUsageCostConfigPostData) GetType() GCPUsageCostConfigPostRequestType {
	if o == nil {
		var ret GCPUsageCostConfigPostRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostData) GetTypeOk() (*GCPUsageCostConfigPostRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GCPUsageCostConfigPostData) SetType(v GCPUsageCostConfigPostRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPUsageCostConfigPostData) MarshalJSON() ([]byte, error) {
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
func (o *GCPUsageCostConfigPostData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *GCPUsageCostConfigPostRequestAttributes `json:"attributes,omitempty"`
		Type       *GCPUsageCostConfigPostRequestType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
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
