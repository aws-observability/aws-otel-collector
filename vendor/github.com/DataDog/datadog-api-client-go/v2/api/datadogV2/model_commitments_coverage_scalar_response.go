// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsCoverageScalarResponse Response containing scalar coverage metrics for cloud commitment programs.
type CommitmentsCoverageScalarResponse struct {
	// Array of scalar columns in the response.
	Columns []CommitmentsScalarColumn `json:"columns"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsCoverageScalarResponse instantiates a new CommitmentsCoverageScalarResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsCoverageScalarResponse(columns []CommitmentsScalarColumn) *CommitmentsCoverageScalarResponse {
	this := CommitmentsCoverageScalarResponse{}
	this.Columns = columns
	return &this
}

// NewCommitmentsCoverageScalarResponseWithDefaults instantiates a new CommitmentsCoverageScalarResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsCoverageScalarResponseWithDefaults() *CommitmentsCoverageScalarResponse {
	this := CommitmentsCoverageScalarResponse{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *CommitmentsCoverageScalarResponse) GetColumns() []CommitmentsScalarColumn {
	if o == nil {
		var ret []CommitmentsScalarColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CommitmentsCoverageScalarResponse) GetColumnsOk() (*[]CommitmentsScalarColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *CommitmentsCoverageScalarResponse) SetColumns(v []CommitmentsScalarColumn) {
	o.Columns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsCoverageScalarResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsCoverageScalarResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns *[]CommitmentsScalarColumn `json:"columns"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns"})
	} else {
		return err
	}
	o.Columns = *all.Columns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
