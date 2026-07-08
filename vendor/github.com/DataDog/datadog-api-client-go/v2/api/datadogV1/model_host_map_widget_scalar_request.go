// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// HostMapWidgetScalarRequest Scalar formula request for the infrastructure host map widget. Each formula specifies
// which visual dimension it drives.
type HostMapWidgetScalarRequest struct {
	// List of formulas that operate on queries, each assigned to a visual dimension.
	Formulas []HostMapWidgetFormula `json:"formulas"`
	// List of queries that can be returned directly or used in formulas.
	Queries []FormulaAndFunctionQueryDefinition `json:"queries"`
	// Response format for the scalar formula request. Only `scalar` is supported.
	ResponseFormat HostMapWidgetScalarRequestResponseFormat `json:"response_format"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHostMapWidgetScalarRequest instantiates a new HostMapWidgetScalarRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHostMapWidgetScalarRequest(formulas []HostMapWidgetFormula, queries []FormulaAndFunctionQueryDefinition, responseFormat HostMapWidgetScalarRequestResponseFormat) *HostMapWidgetScalarRequest {
	this := HostMapWidgetScalarRequest{}
	this.Formulas = formulas
	this.Queries = queries
	this.ResponseFormat = responseFormat
	return &this
}

// NewHostMapWidgetScalarRequestWithDefaults instantiates a new HostMapWidgetScalarRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHostMapWidgetScalarRequestWithDefaults() *HostMapWidgetScalarRequest {
	this := HostMapWidgetScalarRequest{}
	return &this
}

// GetFormulas returns the Formulas field value.
func (o *HostMapWidgetScalarRequest) GetFormulas() []HostMapWidgetFormula {
	if o == nil {
		var ret []HostMapWidgetFormula
		return ret
	}
	return o.Formulas
}

// GetFormulasOk returns a tuple with the Formulas field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetScalarRequest) GetFormulasOk() (*[]HostMapWidgetFormula, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Formulas, true
}

// SetFormulas sets field value.
func (o *HostMapWidgetScalarRequest) SetFormulas(v []HostMapWidgetFormula) {
	o.Formulas = v
}

// GetQueries returns the Queries field value.
func (o *HostMapWidgetScalarRequest) GetQueries() []FormulaAndFunctionQueryDefinition {
	if o == nil {
		var ret []FormulaAndFunctionQueryDefinition
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetScalarRequest) GetQueriesOk() (*[]FormulaAndFunctionQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Queries, true
}

// SetQueries sets field value.
func (o *HostMapWidgetScalarRequest) SetQueries(v []FormulaAndFunctionQueryDefinition) {
	o.Queries = v
}

// GetResponseFormat returns the ResponseFormat field value.
func (o *HostMapWidgetScalarRequest) GetResponseFormat() HostMapWidgetScalarRequestResponseFormat {
	if o == nil {
		var ret HostMapWidgetScalarRequestResponseFormat
		return ret
	}
	return o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value
// and a boolean to check if the value has been set.
func (o *HostMapWidgetScalarRequest) GetResponseFormatOk() (*HostMapWidgetScalarRequestResponseFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseFormat, true
}

// SetResponseFormat sets field value.
func (o *HostMapWidgetScalarRequest) SetResponseFormat(v HostMapWidgetScalarRequestResponseFormat) {
	o.ResponseFormat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HostMapWidgetScalarRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["formulas"] = o.Formulas
	toSerialize["queries"] = o.Queries
	toSerialize["response_format"] = o.ResponseFormat

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HostMapWidgetScalarRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Formulas       *[]HostMapWidgetFormula                   `json:"formulas"`
		Queries        *[]FormulaAndFunctionQueryDefinition      `json:"queries"`
		ResponseFormat *HostMapWidgetScalarRequestResponseFormat `json:"response_format"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Formulas == nil {
		return fmt.Errorf("required field formulas missing")
	}
	if all.Queries == nil {
		return fmt.Errorf("required field queries missing")
	}
	if all.ResponseFormat == nil {
		return fmt.Errorf("required field response_format missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"formulas", "queries", "response_format"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Formulas = *all.Formulas
	o.Queries = *all.Queries
	if !all.ResponseFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.ResponseFormat = *all.ResponseFormat
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
