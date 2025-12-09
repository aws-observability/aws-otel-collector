// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonOpenSearchDestination The `amazon_opensearch` destination writes logs to Amazon OpenSearch.
type ObservabilityPipelineAmazonOpenSearchDestination struct {
	// Authentication settings for the Amazon OpenSearch destination.
	// The `strategy` field determines whether basic or AWS-based authentication is used.
	Auth ObservabilityPipelineAmazonOpenSearchDestinationAuth `json:"auth"`
	// The index to write logs to.
	BulkIndex *string `json:"bulk_index,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The destination type. The value should always be `amazon_opensearch`.
	Type ObservabilityPipelineAmazonOpenSearchDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonOpenSearchDestination instantiates a new ObservabilityPipelineAmazonOpenSearchDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonOpenSearchDestination(auth ObservabilityPipelineAmazonOpenSearchDestinationAuth, id string, inputs []string, typeVar ObservabilityPipelineAmazonOpenSearchDestinationType) *ObservabilityPipelineAmazonOpenSearchDestination {
	this := ObservabilityPipelineAmazonOpenSearchDestination{}
	this.Auth = auth
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAmazonOpenSearchDestinationWithDefaults instantiates a new ObservabilityPipelineAmazonOpenSearchDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonOpenSearchDestinationWithDefaults() *ObservabilityPipelineAmazonOpenSearchDestination {
	this := ObservabilityPipelineAmazonOpenSearchDestination{}
	var typeVar ObservabilityPipelineAmazonOpenSearchDestinationType = OBSERVABILITYPIPELINEAMAZONOPENSEARCHDESTINATIONTYPE_AMAZON_OPENSEARCH
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetAuth() ObservabilityPipelineAmazonOpenSearchDestinationAuth {
	if o == nil {
		var ret ObservabilityPipelineAmazonOpenSearchDestinationAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetAuthOk() (*ObservabilityPipelineAmazonOpenSearchDestinationAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) SetAuth(v ObservabilityPipelineAmazonOpenSearchDestinationAuth) {
	o.Auth = v
}

// GetBulkIndex returns the BulkIndex field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetBulkIndex() string {
	if o == nil || o.BulkIndex == nil {
		var ret string
		return ret
	}
	return *o.BulkIndex
}

// GetBulkIndexOk returns a tuple with the BulkIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetBulkIndexOk() (*string, bool) {
	if o == nil || o.BulkIndex == nil {
		return nil, false
	}
	return o.BulkIndex, true
}

// HasBulkIndex returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) HasBulkIndex() bool {
	return o != nil && o.BulkIndex != nil
}

// SetBulkIndex gets a reference to the given string and assigns it to the BulkIndex field.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) SetBulkIndex(v string) {
	o.BulkIndex = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetType() ObservabilityPipelineAmazonOpenSearchDestinationType {
	if o == nil {
		var ret ObservabilityPipelineAmazonOpenSearchDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) GetTypeOk() (*ObservabilityPipelineAmazonOpenSearchDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) SetType(v ObservabilityPipelineAmazonOpenSearchDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonOpenSearchDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth"] = o.Auth
	if o.BulkIndex != nil {
		toSerialize["bulk_index"] = o.BulkIndex
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAmazonOpenSearchDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth      *ObservabilityPipelineAmazonOpenSearchDestinationAuth `json:"auth"`
		BulkIndex *string                                               `json:"bulk_index,omitempty"`
		Id        *string                                               `json:"id"`
		Inputs    *[]string                                             `json:"inputs"`
		Type      *ObservabilityPipelineAmazonOpenSearchDestinationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Auth == nil {
		return fmt.Errorf("required field auth missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "bulk_index", "id", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = *all.Auth
	o.BulkIndex = all.BulkIndex
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
