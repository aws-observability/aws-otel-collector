// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationUpdateDataRelationshipsDegradationData The degradation linked to a degradation update.
type DegradationUpdateDataRelationshipsDegradationData struct {
	// The ID of the degradation.
	Id string `json:"id"`
	// Degradations resource type.
	Type PatchDegradationRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationUpdateDataRelationshipsDegradationData instantiates a new DegradationUpdateDataRelationshipsDegradationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationUpdateDataRelationshipsDegradationData(id string, typeVar PatchDegradationRequestDataType) *DegradationUpdateDataRelationshipsDegradationData {
	this := DegradationUpdateDataRelationshipsDegradationData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewDegradationUpdateDataRelationshipsDegradationDataWithDefaults instantiates a new DegradationUpdateDataRelationshipsDegradationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationUpdateDataRelationshipsDegradationDataWithDefaults() *DegradationUpdateDataRelationshipsDegradationData {
	this := DegradationUpdateDataRelationshipsDegradationData{}
	var typeVar PatchDegradationRequestDataType = PATCHDEGRADATIONREQUESTDATATYPE_DEGRADATIONS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *DegradationUpdateDataRelationshipsDegradationData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DegradationUpdateDataRelationshipsDegradationData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DegradationUpdateDataRelationshipsDegradationData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *DegradationUpdateDataRelationshipsDegradationData) GetType() PatchDegradationRequestDataType {
	if o == nil {
		var ret PatchDegradationRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DegradationUpdateDataRelationshipsDegradationData) GetTypeOk() (*PatchDegradationRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DegradationUpdateDataRelationshipsDegradationData) SetType(v PatchDegradationRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationUpdateDataRelationshipsDegradationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DegradationUpdateDataRelationshipsDegradationData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                          `json:"id"`
		Type *PatchDegradationRequestDataType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
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
