// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition
type ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition struct {
	//
	End *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition `json:"end,omitempty"`
	//
	FileName *string `json:"file_name,omitempty"`
	//
	Start *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition() *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition{}
	return &this
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItemsFilePositionWithDefaults instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesDependenciesItemsLocationsItemsFilePositionWithDefaults() *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) GetEnd() ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition {
	if o == nil || o.End == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) GetEndOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition and assigns it to the End field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) SetEnd(v ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) {
	o.End = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) SetFileName(v string) {
	o.FileName = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) GetStart() ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition {
	if o == nil || o.Start == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) GetStartOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition and assigns it to the Start field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) SetStart(v ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition) {
	o.Start = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.FileName != nil {
		toSerialize["file_name"] = o.FileName
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End      *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition `json:"end,omitempty"`
		FileName *string                                                          `json:"file_name,omitempty"`
		Start    *ScaRequestDataAttributesDependenciesItemsLocationsItemsPosition `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "file_name", "start"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.End != nil && all.End.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.End = all.End
	o.FileName = all.FileName
	if all.Start != nil && all.Start.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Start = all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
