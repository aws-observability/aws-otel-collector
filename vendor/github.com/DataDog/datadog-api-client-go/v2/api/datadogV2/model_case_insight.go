// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseInsight A reference to an external Datadog resource that provides investigative context for a case, such as a security signal, monitor alert, error tracking issue, or incident.
type CaseInsight struct {
	// The URL path or deep link to the insight resource within Datadog (for example, `/monitors/12345?q=total`).
	Ref string `json:"ref"`
	// The unique identifier of the referenced Datadog resource (for example, a monitor ID, incident ID, or signal ID).
	ResourceId string `json:"resource_id"`
	// The type of Datadog resource linked to the case as contextual evidence. Each type corresponds to a different Datadog product signal (for example, a security finding, a monitor alert, or an incident).
	Type CaseInsightType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseInsight instantiates a new CaseInsight object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseInsight(ref string, resourceId string, typeVar CaseInsightType) *CaseInsight {
	this := CaseInsight{}
	this.Ref = ref
	this.ResourceId = resourceId
	this.Type = typeVar
	return &this
}

// NewCaseInsightWithDefaults instantiates a new CaseInsight object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseInsightWithDefaults() *CaseInsight {
	this := CaseInsight{}
	return &this
}

// GetRef returns the Ref field value.
func (o *CaseInsight) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *CaseInsight) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value.
func (o *CaseInsight) SetRef(v string) {
	o.Ref = v
}

// GetResourceId returns the ResourceId field value.
func (o *CaseInsight) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *CaseInsight) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *CaseInsight) SetResourceId(v string) {
	o.ResourceId = v
}

// GetType returns the Type field value.
func (o *CaseInsight) GetType() CaseInsightType {
	if o == nil {
		var ret CaseInsightType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaseInsight) GetTypeOk() (*CaseInsightType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CaseInsight) SetType(v CaseInsightType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseInsight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ref"] = o.Ref
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseInsight) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ref        *string          `json:"ref"`
		ResourceId *string          `json:"resource_id"`
		Type       *CaseInsightType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Ref == nil {
		return fmt.Errorf("required field ref missing")
	}
	if all.ResourceId == nil {
		return fmt.Errorf("required field resource_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ref", "resource_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Ref = *all.Ref
	o.ResourceId = *all.ResourceId
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
