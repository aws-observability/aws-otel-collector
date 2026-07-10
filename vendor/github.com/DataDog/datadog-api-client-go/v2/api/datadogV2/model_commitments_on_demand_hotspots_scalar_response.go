// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsOnDemandHotspotsScalarResponse Response containing scalar on-demand hot-spots data for cloud commitment programs.
type CommitmentsOnDemandHotspotsScalarResponse struct {
	// Array of scalar columns in the response.
	Columns []CommitmentsScalarColumn `json:"columns"`
	// Metadata for the on-demand hot-spots scalar response.
	Meta *CommitmentsOnDemandHotspotsScalarMeta `json:"meta,omitempty"`
	// Array of scalar columns in the response.
	Total []CommitmentsScalarColumn `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsOnDemandHotspotsScalarResponse instantiates a new CommitmentsOnDemandHotspotsScalarResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsOnDemandHotspotsScalarResponse(columns []CommitmentsScalarColumn, total []CommitmentsScalarColumn) *CommitmentsOnDemandHotspotsScalarResponse {
	this := CommitmentsOnDemandHotspotsScalarResponse{}
	this.Columns = columns
	this.Total = total
	return &this
}

// NewCommitmentsOnDemandHotspotsScalarResponseWithDefaults instantiates a new CommitmentsOnDemandHotspotsScalarResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsOnDemandHotspotsScalarResponseWithDefaults() *CommitmentsOnDemandHotspotsScalarResponse {
	this := CommitmentsOnDemandHotspotsScalarResponse{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *CommitmentsOnDemandHotspotsScalarResponse) GetColumns() []CommitmentsScalarColumn {
	if o == nil {
		var ret []CommitmentsScalarColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CommitmentsOnDemandHotspotsScalarResponse) GetColumnsOk() (*[]CommitmentsScalarColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *CommitmentsOnDemandHotspotsScalarResponse) SetColumns(v []CommitmentsScalarColumn) {
	o.Columns = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CommitmentsOnDemandHotspotsScalarResponse) GetMeta() CommitmentsOnDemandHotspotsScalarMeta {
	if o == nil || o.Meta == nil {
		var ret CommitmentsOnDemandHotspotsScalarMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsOnDemandHotspotsScalarResponse) GetMetaOk() (*CommitmentsOnDemandHotspotsScalarMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CommitmentsOnDemandHotspotsScalarResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given CommitmentsOnDemandHotspotsScalarMeta and assigns it to the Meta field.
func (o *CommitmentsOnDemandHotspotsScalarResponse) SetMeta(v CommitmentsOnDemandHotspotsScalarMeta) {
	o.Meta = &v
}

// GetTotal returns the Total field value.
func (o *CommitmentsOnDemandHotspotsScalarResponse) GetTotal() []CommitmentsScalarColumn {
	if o == nil {
		var ret []CommitmentsScalarColumn
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CommitmentsOnDemandHotspotsScalarResponse) GetTotalOk() (*[]CommitmentsScalarColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *CommitmentsOnDemandHotspotsScalarResponse) SetTotal(v []CommitmentsScalarColumn) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsOnDemandHotspotsScalarResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsOnDemandHotspotsScalarResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns *[]CommitmentsScalarColumn             `json:"columns"`
		Meta    *CommitmentsOnDemandHotspotsScalarMeta `json:"meta,omitempty"`
		Total   *[]CommitmentsScalarColumn             `json:"total"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "meta", "total"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
