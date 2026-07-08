// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsUtilizationScalarResponse Response containing scalar utilization metrics for cloud commitment programs.
type CommitmentsUtilizationScalarResponse struct {
	// Array of scalar columns in the response.
	Columns []CommitmentsScalarColumn `json:"columns"`
	// Array of per-product utilization breakdown entries.
	ProductBreakdown []CommitmentsUtilizationScalarProductBreakdownEntry `json:"product_breakdown,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsUtilizationScalarResponse instantiates a new CommitmentsUtilizationScalarResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsUtilizationScalarResponse(columns []CommitmentsScalarColumn) *CommitmentsUtilizationScalarResponse {
	this := CommitmentsUtilizationScalarResponse{}
	this.Columns = columns
	return &this
}

// NewCommitmentsUtilizationScalarResponseWithDefaults instantiates a new CommitmentsUtilizationScalarResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsUtilizationScalarResponseWithDefaults() *CommitmentsUtilizationScalarResponse {
	this := CommitmentsUtilizationScalarResponse{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *CommitmentsUtilizationScalarResponse) GetColumns() []CommitmentsScalarColumn {
	if o == nil {
		var ret []CommitmentsScalarColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CommitmentsUtilizationScalarResponse) GetColumnsOk() (*[]CommitmentsScalarColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *CommitmentsUtilizationScalarResponse) SetColumns(v []CommitmentsScalarColumn) {
	o.Columns = v
}

// GetProductBreakdown returns the ProductBreakdown field value if set, zero value otherwise.
func (o *CommitmentsUtilizationScalarResponse) GetProductBreakdown() []CommitmentsUtilizationScalarProductBreakdownEntry {
	if o == nil || o.ProductBreakdown == nil {
		var ret []CommitmentsUtilizationScalarProductBreakdownEntry
		return ret
	}
	return o.ProductBreakdown
}

// GetProductBreakdownOk returns a tuple with the ProductBreakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsUtilizationScalarResponse) GetProductBreakdownOk() (*[]CommitmentsUtilizationScalarProductBreakdownEntry, bool) {
	if o == nil || o.ProductBreakdown == nil {
		return nil, false
	}
	return &o.ProductBreakdown, true
}

// HasProductBreakdown returns a boolean if a field has been set.
func (o *CommitmentsUtilizationScalarResponse) HasProductBreakdown() bool {
	return o != nil && o.ProductBreakdown != nil
}

// SetProductBreakdown gets a reference to the given []CommitmentsUtilizationScalarProductBreakdownEntry and assigns it to the ProductBreakdown field.
func (o *CommitmentsUtilizationScalarResponse) SetProductBreakdown(v []CommitmentsUtilizationScalarProductBreakdownEntry) {
	o.ProductBreakdown = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsUtilizationScalarResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	if o.ProductBreakdown != nil {
		toSerialize["product_breakdown"] = o.ProductBreakdown
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsUtilizationScalarResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns          *[]CommitmentsScalarColumn                          `json:"columns"`
		ProductBreakdown []CommitmentsUtilizationScalarProductBreakdownEntry `json:"product_breakdown,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "product_breakdown"})
	} else {
		return err
	}
	o.Columns = *all.Columns
	o.ProductBreakdown = all.ProductBreakdown

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
