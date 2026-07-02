// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsSavingsScalarResponse Response containing scalar savings metrics for cloud commitment programs.
type CommitmentsSavingsScalarResponse struct {
	// Array of scalar columns in the response.
	Columns []CommitmentsScalarColumn `json:"columns"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsSavingsScalarResponse instantiates a new CommitmentsSavingsScalarResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsSavingsScalarResponse(columns []CommitmentsScalarColumn) *CommitmentsSavingsScalarResponse {
	this := CommitmentsSavingsScalarResponse{}
	this.Columns = columns
	return &this
}

// NewCommitmentsSavingsScalarResponseWithDefaults instantiates a new CommitmentsSavingsScalarResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsSavingsScalarResponseWithDefaults() *CommitmentsSavingsScalarResponse {
	this := CommitmentsSavingsScalarResponse{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *CommitmentsSavingsScalarResponse) GetColumns() []CommitmentsScalarColumn {
	if o == nil {
		var ret []CommitmentsScalarColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *CommitmentsSavingsScalarResponse) GetColumnsOk() (*[]CommitmentsScalarColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *CommitmentsSavingsScalarResponse) SetColumns(v []CommitmentsScalarColumn) {
	o.Columns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsSavingsScalarResponse) MarshalJSON() ([]byte, error) {
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
func (o *CommitmentsSavingsScalarResponse) UnmarshalJSON(bytes []byte) (err error) {
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
