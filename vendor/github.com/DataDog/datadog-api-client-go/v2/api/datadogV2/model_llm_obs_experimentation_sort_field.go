// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentationSortField A field and direction to sort results by.
type LLMObsExperimentationSortField struct {
	// Sort direction.
	Direction *LLMObsExperimentationSortFieldDirection `json:"direction,omitempty"`
	// The field name to sort on.
	Field string `json:"field"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentationSortField instantiates a new LLMObsExperimentationSortField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentationSortField(field string) *LLMObsExperimentationSortField {
	this := LLMObsExperimentationSortField{}
	this.Field = field
	return &this
}

// NewLLMObsExperimentationSortFieldWithDefaults instantiates a new LLMObsExperimentationSortField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentationSortFieldWithDefaults() *LLMObsExperimentationSortField {
	this := LLMObsExperimentationSortField{}
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *LLMObsExperimentationSortField) GetDirection() LLMObsExperimentationSortFieldDirection {
	if o == nil || o.Direction == nil {
		var ret LLMObsExperimentationSortFieldDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSortField) GetDirectionOk() (*LLMObsExperimentationSortFieldDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *LLMObsExperimentationSortField) HasDirection() bool {
	return o != nil && o.Direction != nil
}

// SetDirection gets a reference to the given LLMObsExperimentationSortFieldDirection and assigns it to the Direction field.
func (o *LLMObsExperimentationSortField) SetDirection(v LLMObsExperimentationSortFieldDirection) {
	o.Direction = &v
}

// GetField returns the Field field value.
func (o *LLMObsExperimentationSortField) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentationSortField) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *LLMObsExperimentationSortField) SetField(v string) {
	o.Field = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentationSortField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	toSerialize["field"] = o.Field

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentationSortField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Direction *LLMObsExperimentationSortFieldDirection `json:"direction,omitempty"`
		Field     *string                                  `json:"field"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"direction", "field"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Direction != nil && !all.Direction.IsValid() {
		hasInvalidField = true
	} else {
		o.Direction = all.Direction
	}
	o.Field = *all.Field

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
