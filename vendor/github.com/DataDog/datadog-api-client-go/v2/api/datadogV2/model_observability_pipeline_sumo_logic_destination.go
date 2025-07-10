// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSumoLogicDestination The `sumo_logic` destination forwards logs to Sumo Logic.
type ObservabilityPipelineSumoLogicDestination struct {
	// The output encoding format.
	Encoding *ObservabilityPipelineSumoLogicDestinationEncoding `json:"encoding,omitempty"`
	// A list of custom headers to include in the request to Sumo Logic.
	HeaderCustomFields []ObservabilityPipelineSumoLogicDestinationHeaderCustomFieldsItem `json:"header_custom_fields,omitempty"`
	// Optional override for the host name header.
	HeaderHostName *string `json:"header_host_name,omitempty"`
	// Optional override for the source category header.
	HeaderSourceCategory *string `json:"header_source_category,omitempty"`
	// Optional override for the source name header.
	HeaderSourceName *string `json:"header_source_name,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The destination type. The value should always be `sumo_logic`.
	Type ObservabilityPipelineSumoLogicDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSumoLogicDestination instantiates a new ObservabilityPipelineSumoLogicDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSumoLogicDestination(id string, inputs []string, typeVar ObservabilityPipelineSumoLogicDestinationType) *ObservabilityPipelineSumoLogicDestination {
	this := ObservabilityPipelineSumoLogicDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSumoLogicDestinationWithDefaults instantiates a new ObservabilityPipelineSumoLogicDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSumoLogicDestinationWithDefaults() *ObservabilityPipelineSumoLogicDestination {
	this := ObservabilityPipelineSumoLogicDestination{}
	var typeVar ObservabilityPipelineSumoLogicDestinationType = OBSERVABILITYPIPELINESUMOLOGICDESTINATIONTYPE_SUMO_LOGIC
	this.Type = typeVar
	return &this
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *ObservabilityPipelineSumoLogicDestination) GetEncoding() ObservabilityPipelineSumoLogicDestinationEncoding {
	if o == nil || o.Encoding == nil {
		var ret ObservabilityPipelineSumoLogicDestinationEncoding
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSumoLogicDestination) GetEncodingOk() (*ObservabilityPipelineSumoLogicDestinationEncoding, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *ObservabilityPipelineSumoLogicDestination) HasEncoding() bool {
	return o != nil && o.Encoding != nil
}

// SetEncoding gets a reference to the given ObservabilityPipelineSumoLogicDestinationEncoding and assigns it to the Encoding field.
func (o *ObservabilityPipelineSumoLogicDestination) SetEncoding(v ObservabilityPipelineSumoLogicDestinationEncoding) {
	o.Encoding = &v
}

// GetHeaderCustomFields returns the HeaderCustomFields field value if set, zero value otherwise.
func (o *ObservabilityPipelineSumoLogicDestination) GetHeaderCustomFields() []ObservabilityPipelineSumoLogicDestinationHeaderCustomFieldsItem {
	if o == nil || o.HeaderCustomFields == nil {
		var ret []ObservabilityPipelineSumoLogicDestinationHeaderCustomFieldsItem
		return ret
	}
	return o.HeaderCustomFields
}

// GetHeaderCustomFieldsOk returns a tuple with the HeaderCustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSumoLogicDestination) GetHeaderCustomFieldsOk() (*[]ObservabilityPipelineSumoLogicDestinationHeaderCustomFieldsItem, bool) {
	if o == nil || o.HeaderCustomFields == nil {
		return nil, false
	}
	return &o.HeaderCustomFields, true
}

// HasHeaderCustomFields returns a boolean if a field has been set.
func (o *ObservabilityPipelineSumoLogicDestination) HasHeaderCustomFields() bool {
	return o != nil && o.HeaderCustomFields != nil
}

// SetHeaderCustomFields gets a reference to the given []ObservabilityPipelineSumoLogicDestinationHeaderCustomFieldsItem and assigns it to the HeaderCustomFields field.
func (o *ObservabilityPipelineSumoLogicDestination) SetHeaderCustomFields(v []ObservabilityPipelineSumoLogicDestinationHeaderCustomFieldsItem) {
	o.HeaderCustomFields = v
}

// GetHeaderHostName returns the HeaderHostName field value if set, zero value otherwise.
func (o *ObservabilityPipelineSumoLogicDestination) GetHeaderHostName() string {
	if o == nil || o.HeaderHostName == nil {
		var ret string
		return ret
	}
	return *o.HeaderHostName
}

// GetHeaderHostNameOk returns a tuple with the HeaderHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSumoLogicDestination) GetHeaderHostNameOk() (*string, bool) {
	if o == nil || o.HeaderHostName == nil {
		return nil, false
	}
	return o.HeaderHostName, true
}

// HasHeaderHostName returns a boolean if a field has been set.
func (o *ObservabilityPipelineSumoLogicDestination) HasHeaderHostName() bool {
	return o != nil && o.HeaderHostName != nil
}

// SetHeaderHostName gets a reference to the given string and assigns it to the HeaderHostName field.
func (o *ObservabilityPipelineSumoLogicDestination) SetHeaderHostName(v string) {
	o.HeaderHostName = &v
}

// GetHeaderSourceCategory returns the HeaderSourceCategory field value if set, zero value otherwise.
func (o *ObservabilityPipelineSumoLogicDestination) GetHeaderSourceCategory() string {
	if o == nil || o.HeaderSourceCategory == nil {
		var ret string
		return ret
	}
	return *o.HeaderSourceCategory
}

// GetHeaderSourceCategoryOk returns a tuple with the HeaderSourceCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSumoLogicDestination) GetHeaderSourceCategoryOk() (*string, bool) {
	if o == nil || o.HeaderSourceCategory == nil {
		return nil, false
	}
	return o.HeaderSourceCategory, true
}

// HasHeaderSourceCategory returns a boolean if a field has been set.
func (o *ObservabilityPipelineSumoLogicDestination) HasHeaderSourceCategory() bool {
	return o != nil && o.HeaderSourceCategory != nil
}

// SetHeaderSourceCategory gets a reference to the given string and assigns it to the HeaderSourceCategory field.
func (o *ObservabilityPipelineSumoLogicDestination) SetHeaderSourceCategory(v string) {
	o.HeaderSourceCategory = &v
}

// GetHeaderSourceName returns the HeaderSourceName field value if set, zero value otherwise.
func (o *ObservabilityPipelineSumoLogicDestination) GetHeaderSourceName() string {
	if o == nil || o.HeaderSourceName == nil {
		var ret string
		return ret
	}
	return *o.HeaderSourceName
}

// GetHeaderSourceNameOk returns a tuple with the HeaderSourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSumoLogicDestination) GetHeaderSourceNameOk() (*string, bool) {
	if o == nil || o.HeaderSourceName == nil {
		return nil, false
	}
	return o.HeaderSourceName, true
}

// HasHeaderSourceName returns a boolean if a field has been set.
func (o *ObservabilityPipelineSumoLogicDestination) HasHeaderSourceName() bool {
	return o != nil && o.HeaderSourceName != nil
}

// SetHeaderSourceName gets a reference to the given string and assigns it to the HeaderSourceName field.
func (o *ObservabilityPipelineSumoLogicDestination) SetHeaderSourceName(v string) {
	o.HeaderSourceName = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSumoLogicDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSumoLogicDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSumoLogicDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineSumoLogicDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSumoLogicDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineSumoLogicDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSumoLogicDestination) GetType() ObservabilityPipelineSumoLogicDestinationType {
	if o == nil {
		var ret ObservabilityPipelineSumoLogicDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSumoLogicDestination) GetTypeOk() (*ObservabilityPipelineSumoLogicDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSumoLogicDestination) SetType(v ObservabilityPipelineSumoLogicDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSumoLogicDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.HeaderCustomFields != nil {
		toSerialize["header_custom_fields"] = o.HeaderCustomFields
	}
	if o.HeaderHostName != nil {
		toSerialize["header_host_name"] = o.HeaderHostName
	}
	if o.HeaderSourceCategory != nil {
		toSerialize["header_source_category"] = o.HeaderSourceCategory
	}
	if o.HeaderSourceName != nil {
		toSerialize["header_source_name"] = o.HeaderSourceName
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
func (o *ObservabilityPipelineSumoLogicDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Encoding             *ObservabilityPipelineSumoLogicDestinationEncoding                `json:"encoding,omitempty"`
		HeaderCustomFields   []ObservabilityPipelineSumoLogicDestinationHeaderCustomFieldsItem `json:"header_custom_fields,omitempty"`
		HeaderHostName       *string                                                           `json:"header_host_name,omitempty"`
		HeaderSourceCategory *string                                                           `json:"header_source_category,omitempty"`
		HeaderSourceName     *string                                                           `json:"header_source_name,omitempty"`
		Id                   *string                                                           `json:"id"`
		Inputs               *[]string                                                         `json:"inputs"`
		Type                 *ObservabilityPipelineSumoLogicDestinationType                    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
		datadog.DeleteKeys(additionalProperties, &[]string{"encoding", "header_custom_fields", "header_host_name", "header_source_category", "header_source_name", "id", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Encoding != nil && !all.Encoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Encoding = all.Encoding
	}
	o.HeaderCustomFields = all.HeaderCustomFields
	o.HeaderHostName = all.HeaderHostName
	o.HeaderSourceCategory = all.HeaderSourceCategory
	o.HeaderSourceName = all.HeaderSourceName
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
