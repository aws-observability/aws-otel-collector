// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureStorageDestination The `azure_storage` destination forwards logs to an Azure Blob Storage container.
type AzureStorageDestination struct {
	// Optional prefix for blobs written to the container.
	BlobPrefix *string `json:"blob_prefix,omitempty"`
	// The name of the Azure Blob Storage container to store logs in.
	ContainerName string `json:"container_name"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The destination type. The value should always be `azure_storage`.
	Type AzureStorageDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureStorageDestination instantiates a new AzureStorageDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureStorageDestination(containerName string, id string, inputs []string, typeVar AzureStorageDestinationType) *AzureStorageDestination {
	this := AzureStorageDestination{}
	this.ContainerName = containerName
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewAzureStorageDestinationWithDefaults instantiates a new AzureStorageDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureStorageDestinationWithDefaults() *AzureStorageDestination {
	this := AzureStorageDestination{}
	var typeVar AzureStorageDestinationType = AZURESTORAGEDESTINATIONTYPE_AZURE_STORAGE
	this.Type = typeVar
	return &this
}

// GetBlobPrefix returns the BlobPrefix field value if set, zero value otherwise.
func (o *AzureStorageDestination) GetBlobPrefix() string {
	if o == nil || o.BlobPrefix == nil {
		var ret string
		return ret
	}
	return *o.BlobPrefix
}

// GetBlobPrefixOk returns a tuple with the BlobPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageDestination) GetBlobPrefixOk() (*string, bool) {
	if o == nil || o.BlobPrefix == nil {
		return nil, false
	}
	return o.BlobPrefix, true
}

// HasBlobPrefix returns a boolean if a field has been set.
func (o *AzureStorageDestination) HasBlobPrefix() bool {
	return o != nil && o.BlobPrefix != nil
}

// SetBlobPrefix gets a reference to the given string and assigns it to the BlobPrefix field.
func (o *AzureStorageDestination) SetBlobPrefix(v string) {
	o.BlobPrefix = &v
}

// GetContainerName returns the ContainerName field value.
func (o *AzureStorageDestination) GetContainerName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value
// and a boolean to check if the value has been set.
func (o *AzureStorageDestination) GetContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerName, true
}

// SetContainerName sets field value.
func (o *AzureStorageDestination) SetContainerName(v string) {
	o.ContainerName = v
}

// GetId returns the Id field value.
func (o *AzureStorageDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzureStorageDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AzureStorageDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *AzureStorageDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *AzureStorageDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *AzureStorageDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *AzureStorageDestination) GetType() AzureStorageDestinationType {
	if o == nil {
		var ret AzureStorageDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AzureStorageDestination) GetTypeOk() (*AzureStorageDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AzureStorageDestination) SetType(v AzureStorageDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureStorageDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BlobPrefix != nil {
		toSerialize["blob_prefix"] = o.BlobPrefix
	}
	toSerialize["container_name"] = o.ContainerName
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureStorageDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BlobPrefix    *string                      `json:"blob_prefix,omitempty"`
		ContainerName *string                      `json:"container_name"`
		Id            *string                      `json:"id"`
		Inputs        *[]string                    `json:"inputs"`
		Type          *AzureStorageDestinationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ContainerName == nil {
		return fmt.Errorf("required field container_name missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"blob_prefix", "container_name", "id", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BlobPrefix = all.BlobPrefix
	o.ContainerName = *all.ContainerName
	o.Id = *all.Id
	o.Inputs = *all.Inputs
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
