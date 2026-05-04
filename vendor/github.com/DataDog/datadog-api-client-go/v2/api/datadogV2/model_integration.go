// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Integration Integration resource object.
type Integration struct {
	// Attributes for an integration.
	Attributes IntegrationAttributes `json:"attributes"`
	// The unique identifier of the integration.
	Id string `json:"id"`
	// Links for the integration resource.
	Links *IntegrationLinks `json:"links,omitempty"`
	// Integration resource type.
	Type IntegrationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegration instantiates a new Integration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegration(attributes IntegrationAttributes, id string, typeVar IntegrationType) *Integration {
	this := Integration{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIntegrationWithDefaults instantiates a new Integration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationWithDefaults() *Integration {
	this := Integration{}
	var typeVar IntegrationType = INTEGRATIONTYPE_INTEGRATION
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *Integration) GetAttributes() IntegrationAttributes {
	if o == nil {
		var ret IntegrationAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *Integration) GetAttributesOk() (*IntegrationAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *Integration) SetAttributes(v IntegrationAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *Integration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Integration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *Integration) SetId(v string) {
	o.Id = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Integration) GetLinks() IntegrationLinks {
	if o == nil || o.Links == nil {
		var ret IntegrationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetLinksOk() (*IntegrationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Integration) HasLinks() bool {
	return o != nil && o.Links != nil
}

// SetLinks gets a reference to the given IntegrationLinks and assigns it to the Links field.
func (o *Integration) SetLinks(v IntegrationLinks) {
	o.Links = &v
}

// GetType returns the Type field value.
func (o *Integration) GetType() IntegrationType {
	if o == nil {
		var ret IntegrationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Integration) GetTypeOk() (*IntegrationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Integration) SetType(v IntegrationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Integration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Integration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *IntegrationAttributes `json:"attributes"`
		Id         *string                `json:"id"`
		Links      *IntegrationLinks      `json:"links,omitempty"`
		Type       *IntegrationType       `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "links", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Links != nil && all.Links.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Links = all.Links
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
