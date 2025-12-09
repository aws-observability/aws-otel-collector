// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGoogleChronicleDestination The `google_chronicle` destination sends logs to Google Chronicle.
type ObservabilityPipelineGoogleChronicleDestination struct {
	// GCP credentials used to authenticate with Google Cloud Storage.
	Auth ObservabilityPipelineGcpAuth `json:"auth"`
	// The Google Chronicle customer ID.
	CustomerId string `json:"customer_id"`
	// The encoding format for the logs sent to Chronicle.
	Encoding *ObservabilityPipelineGoogleChronicleDestinationEncoding `json:"encoding,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The log type metadata associated with the Chronicle destination.
	LogType *string `json:"log_type,omitempty"`
	// The destination type. The value should always be `google_chronicle`.
	Type ObservabilityPipelineGoogleChronicleDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGoogleChronicleDestination instantiates a new ObservabilityPipelineGoogleChronicleDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGoogleChronicleDestination(auth ObservabilityPipelineGcpAuth, customerId string, id string, inputs []string, typeVar ObservabilityPipelineGoogleChronicleDestinationType) *ObservabilityPipelineGoogleChronicleDestination {
	this := ObservabilityPipelineGoogleChronicleDestination{}
	this.Auth = auth
	this.CustomerId = customerId
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineGoogleChronicleDestinationWithDefaults instantiates a new ObservabilityPipelineGoogleChronicleDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGoogleChronicleDestinationWithDefaults() *ObservabilityPipelineGoogleChronicleDestination {
	this := ObservabilityPipelineGoogleChronicleDestination{}
	var typeVar ObservabilityPipelineGoogleChronicleDestinationType = OBSERVABILITYPIPELINEGOOGLECHRONICLEDESTINATIONTYPE_GOOGLE_CHRONICLE
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetAuth() ObservabilityPipelineGcpAuth {
	if o == nil {
		var ret ObservabilityPipelineGcpAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetAuthOk() (*ObservabilityPipelineGcpAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) SetAuth(v ObservabilityPipelineGcpAuth) {
	o.Auth = v
}

// GetCustomerId returns the CustomerId field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetEncoding() ObservabilityPipelineGoogleChronicleDestinationEncoding {
	if o == nil || o.Encoding == nil {
		var ret ObservabilityPipelineGoogleChronicleDestinationEncoding
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetEncodingOk() (*ObservabilityPipelineGoogleChronicleDestinationEncoding, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *ObservabilityPipelineGoogleChronicleDestination) HasEncoding() bool {
	return o != nil && o.Encoding != nil
}

// SetEncoding gets a reference to the given ObservabilityPipelineGoogleChronicleDestinationEncoding and assigns it to the Encoding field.
func (o *ObservabilityPipelineGoogleChronicleDestination) SetEncoding(v ObservabilityPipelineGoogleChronicleDestinationEncoding) {
	o.Encoding = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetLogType returns the LogType field value if set, zero value otherwise.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetLogType() string {
	if o == nil || o.LogType == nil {
		var ret string
		return ret
	}
	return *o.LogType
}

// GetLogTypeOk returns a tuple with the LogType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetLogTypeOk() (*string, bool) {
	if o == nil || o.LogType == nil {
		return nil, false
	}
	return o.LogType, true
}

// HasLogType returns a boolean if a field has been set.
func (o *ObservabilityPipelineGoogleChronicleDestination) HasLogType() bool {
	return o != nil && o.LogType != nil
}

// SetLogType gets a reference to the given string and assigns it to the LogType field.
func (o *ObservabilityPipelineGoogleChronicleDestination) SetLogType(v string) {
	o.LogType = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetType() ObservabilityPipelineGoogleChronicleDestinationType {
	if o == nil {
		var ret ObservabilityPipelineGoogleChronicleDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGoogleChronicleDestination) GetTypeOk() (*ObservabilityPipelineGoogleChronicleDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineGoogleChronicleDestination) SetType(v ObservabilityPipelineGoogleChronicleDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGoogleChronicleDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth"] = o.Auth
	toSerialize["customer_id"] = o.CustomerId
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.LogType != nil {
		toSerialize["log_type"] = o.LogType
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineGoogleChronicleDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth       *ObservabilityPipelineGcpAuth                            `json:"auth"`
		CustomerId *string                                                  `json:"customer_id"`
		Encoding   *ObservabilityPipelineGoogleChronicleDestinationEncoding `json:"encoding,omitempty"`
		Id         *string                                                  `json:"id"`
		Inputs     *[]string                                                `json:"inputs"`
		LogType    *string                                                  `json:"log_type,omitempty"`
		Type       *ObservabilityPipelineGoogleChronicleDestinationType     `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Auth == nil {
		return fmt.Errorf("required field auth missing")
	}
	if all.CustomerId == nil {
		return fmt.Errorf("required field customer_id missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "customer_id", "encoding", "id", "inputs", "log_type", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = *all.Auth
	o.CustomerId = *all.CustomerId
	if all.Encoding != nil && !all.Encoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Encoding = all.Encoding
	}
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.LogType = all.LogType
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
