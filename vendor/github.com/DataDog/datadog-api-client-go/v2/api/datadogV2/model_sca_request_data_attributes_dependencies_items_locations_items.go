// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesDependenciesItemsLocationsItems
type ScaRequestDataAttributesDependenciesItemsLocationsItems struct {
	//
	Block *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition `json:"block,omitempty"`
	//
	Name *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition `json:"name,omitempty"`
	//
	Namespace *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition `json:"namespace,omitempty"`
	//
	Version *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItems instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesDependenciesItemsLocationsItems() *ScaRequestDataAttributesDependenciesItemsLocationsItems {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItems{}
	return &this
}

// NewScaRequestDataAttributesDependenciesItemsLocationsItemsWithDefaults instantiates a new ScaRequestDataAttributesDependenciesItemsLocationsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesDependenciesItemsLocationsItemsWithDefaults() *ScaRequestDataAttributesDependenciesItemsLocationsItems {
	this := ScaRequestDataAttributesDependenciesItemsLocationsItems{}
	return &this
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetBlock() ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition {
	if o == nil || o.Block == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetBlockOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition, bool) {
	if o == nil || o.Block == nil {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) HasBlock() bool {
	return o != nil && o.Block != nil
}

// SetBlock gets a reference to the given ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition and assigns it to the Block field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetBlock(v ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) {
	o.Block = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetName() ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition {
	if o == nil || o.Name == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetNameOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition and assigns it to the Name field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetName(v ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetNamespace() ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition {
	if o == nil || o.Namespace == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetNamespaceOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition and assigns it to the Namespace field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetNamespace(v ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) {
	o.Namespace = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetVersion() ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition {
	if o == nil || o.Version == nil {
		var ret ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) GetVersionOk() (*ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition and assigns it to the Version field.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) SetVersion(v ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesDependenciesItemsLocationsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Block != nil {
		toSerialize["block"] = o.Block
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesDependenciesItemsLocationsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Block     *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition `json:"block,omitempty"`
		Name      *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition `json:"name,omitempty"`
		Namespace *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition `json:"namespace,omitempty"`
		Version   *ScaRequestDataAttributesDependenciesItemsLocationsItemsFilePosition `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"block", "name", "namespace", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Block != nil && all.Block.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Block = all.Block
	if all.Name != nil && all.Name.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Name = all.Name
	if all.Namespace != nil && all.Namespace.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Namespace = all.Namespace
	if all.Version != nil && all.Version.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
