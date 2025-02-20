// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestinationElasticsearch The Elasticsearch destination.
type CustomDestinationResponseForwardDestinationElasticsearch struct {
	// Basic access authentication.
	Auth map[string]interface{} `json:"auth"`
	// The destination for which logs will be forwarded to.
	// Must have HTTPS scheme and forwarding back to Datadog is not allowed.
	Endpoint string `json:"endpoint"`
	// Name of the Elasticsearch index (must follow [Elasticsearch's criteria](https://www.elastic.co/guide/en/elasticsearch/reference/8.11/indices-create-index.html#indices-create-api-path-params)).
	IndexName string `json:"index_name"`
	// Date pattern with US locale and UTC timezone to be appended to the index name after adding `-`
	// (that is, `${index_name}-${indexPattern}`).
	// You can customize the index rotation naming pattern by choosing one of these options:
	// - Hourly: `yyyy-MM-dd-HH` (as an example, it would render: `2022-10-19-09`)
	// - Daily: `yyyy-MM-dd` (as an example, it would render: `2022-10-19`)
	// - Weekly: `yyyy-'W'ww` (as an example, it would render: `2022-W42`)
	// - Monthly: `yyyy-MM` (as an example, it would render: `2022-10`)
	//
	// If this field is missing or is blank, it means that the index name will always be the same
	// (that is, no rotation).
	IndexRotation *string `json:"index_rotation,omitempty"`
	// Type of the Elasticsearch destination.
	Type CustomDestinationResponseForwardDestinationElasticsearchType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationResponseForwardDestinationElasticsearch instantiates a new CustomDestinationResponseForwardDestinationElasticsearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationResponseForwardDestinationElasticsearch(auth map[string]interface{}, endpoint string, indexName string, typeVar CustomDestinationResponseForwardDestinationElasticsearchType) *CustomDestinationResponseForwardDestinationElasticsearch {
	this := CustomDestinationResponseForwardDestinationElasticsearch{}
	this.Auth = auth
	this.Endpoint = endpoint
	this.IndexName = indexName
	this.Type = typeVar
	return &this
}

// NewCustomDestinationResponseForwardDestinationElasticsearchWithDefaults instantiates a new CustomDestinationResponseForwardDestinationElasticsearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationResponseForwardDestinationElasticsearchWithDefaults() *CustomDestinationResponseForwardDestinationElasticsearch {
	this := CustomDestinationResponseForwardDestinationElasticsearch{}
	var typeVar CustomDestinationResponseForwardDestinationElasticsearchType = CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONELASTICSEARCHTYPE_ELASTICSEARCH
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetAuth() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetAuthOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) SetAuth(v map[string]interface{}) {
	o.Auth = v
}

// GetEndpoint returns the Endpoint field value.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetIndexName returns the IndexName field value.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) SetIndexName(v string) {
	o.IndexName = v
}

// GetIndexRotation returns the IndexRotation field value if set, zero value otherwise.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetIndexRotation() string {
	if o == nil || o.IndexRotation == nil {
		var ret string
		return ret
	}
	return *o.IndexRotation
}

// GetIndexRotationOk returns a tuple with the IndexRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetIndexRotationOk() (*string, bool) {
	if o == nil || o.IndexRotation == nil {
		return nil, false
	}
	return o.IndexRotation, true
}

// HasIndexRotation returns a boolean if a field has been set.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) HasIndexRotation() bool {
	return o != nil && o.IndexRotation != nil
}

// SetIndexRotation gets a reference to the given string and assigns it to the IndexRotation field.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) SetIndexRotation(v string) {
	o.IndexRotation = &v
}

// GetType returns the Type field value.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetType() CustomDestinationResponseForwardDestinationElasticsearchType {
	if o == nil {
		var ret CustomDestinationResponseForwardDestinationElasticsearchType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) GetTypeOk() (*CustomDestinationResponseForwardDestinationElasticsearchType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) SetType(v CustomDestinationResponseForwardDestinationElasticsearchType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationResponseForwardDestinationElasticsearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth"] = o.Auth
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["index_name"] = o.IndexName
	if o.IndexRotation != nil {
		toSerialize["index_rotation"] = o.IndexRotation
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationResponseForwardDestinationElasticsearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth          *map[string]interface{}                                       `json:"auth"`
		Endpoint      *string                                                       `json:"endpoint"`
		IndexName     *string                                                       `json:"index_name"`
		IndexRotation *string                                                       `json:"index_rotation,omitempty"`
		Type          *CustomDestinationResponseForwardDestinationElasticsearchType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Auth == nil {
		return fmt.Errorf("required field auth missing")
	}
	if all.Endpoint == nil {
		return fmt.Errorf("required field endpoint missing")
	}
	if all.IndexName == nil {
		return fmt.Errorf("required field index_name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "endpoint", "index_name", "index_rotation", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Auth = *all.Auth
	o.Endpoint = *all.Endpoint
	o.IndexName = *all.IndexName
	o.IndexRotation = all.IndexRotation
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
