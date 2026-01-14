// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceListDataAttributes
type ServiceListDataAttributes struct {
	//
	Metadata []ServiceListDataAttributesMetadataItems `json:"metadata,omitempty"`
	//
	Services []string `json:"services,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceListDataAttributes instantiates a new ServiceListDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceListDataAttributes() *ServiceListDataAttributes {
	this := ServiceListDataAttributes{}
	return &this
}

// NewServiceListDataAttributesWithDefaults instantiates a new ServiceListDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceListDataAttributesWithDefaults() *ServiceListDataAttributes {
	this := ServiceListDataAttributes{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServiceListDataAttributes) GetMetadata() []ServiceListDataAttributesMetadataItems {
	if o == nil || o.Metadata == nil {
		var ret []ServiceListDataAttributesMetadataItems
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceListDataAttributes) GetMetadataOk() (*[]ServiceListDataAttributesMetadataItems, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServiceListDataAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given []ServiceListDataAttributesMetadataItems and assigns it to the Metadata field.
func (o *ServiceListDataAttributes) SetMetadata(v []ServiceListDataAttributesMetadataItems) {
	o.Metadata = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ServiceListDataAttributes) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceListDataAttributes) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ServiceListDataAttributes) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *ServiceListDataAttributes) SetServices(v []string) {
	o.Services = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceListDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceListDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Metadata []ServiceListDataAttributesMetadataItems `json:"metadata,omitempty"`
		Services []string                                 `json:"services,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metadata", "services"})
	} else {
		return err
	}
	o.Metadata = all.Metadata
	o.Services = all.Services

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
