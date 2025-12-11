// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeDeleteResponseAttributes The EventBridge source delete response attributes.
type AWSEventBridgeDeleteResponseAttributes struct {
	// The event source status "empty".
	Status *AWSEventBridgeDeleteStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSEventBridgeDeleteResponseAttributes instantiates a new AWSEventBridgeDeleteResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSEventBridgeDeleteResponseAttributes() *AWSEventBridgeDeleteResponseAttributes {
	this := AWSEventBridgeDeleteResponseAttributes{}
	return &this
}

// NewAWSEventBridgeDeleteResponseAttributesWithDefaults instantiates a new AWSEventBridgeDeleteResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSEventBridgeDeleteResponseAttributesWithDefaults() *AWSEventBridgeDeleteResponseAttributes {
	this := AWSEventBridgeDeleteResponseAttributes{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AWSEventBridgeDeleteResponseAttributes) GetStatus() AWSEventBridgeDeleteStatus {
	if o == nil || o.Status == nil {
		var ret AWSEventBridgeDeleteStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteResponseAttributes) GetStatusOk() (*AWSEventBridgeDeleteStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AWSEventBridgeDeleteResponseAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AWSEventBridgeDeleteStatus and assigns it to the Status field.
func (o *AWSEventBridgeDeleteResponseAttributes) SetStatus(v AWSEventBridgeDeleteStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSEventBridgeDeleteResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSEventBridgeDeleteResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status *AWSEventBridgeDeleteStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
