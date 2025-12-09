// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3DatadogPipelines CI Pipelines association.
type EntityV3DatadogPipelines struct {
	// A list of CI Fingerprints that associate CI Pipelines with the entity.
	Fingerprints []string `json:"fingerprints,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3DatadogPipelines instantiates a new EntityV3DatadogPipelines object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3DatadogPipelines() *EntityV3DatadogPipelines {
	this := EntityV3DatadogPipelines{}
	return &this
}

// NewEntityV3DatadogPipelinesWithDefaults instantiates a new EntityV3DatadogPipelines object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3DatadogPipelinesWithDefaults() *EntityV3DatadogPipelines {
	this := EntityV3DatadogPipelines{}
	return &this
}

// GetFingerprints returns the Fingerprints field value if set, zero value otherwise.
func (o *EntityV3DatadogPipelines) GetFingerprints() []string {
	if o == nil || o.Fingerprints == nil {
		var ret []string
		return ret
	}
	return o.Fingerprints
}

// GetFingerprintsOk returns a tuple with the Fingerprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3DatadogPipelines) GetFingerprintsOk() (*[]string, bool) {
	if o == nil || o.Fingerprints == nil {
		return nil, false
	}
	return &o.Fingerprints, true
}

// HasFingerprints returns a boolean if a field has been set.
func (o *EntityV3DatadogPipelines) HasFingerprints() bool {
	return o != nil && o.Fingerprints != nil
}

// SetFingerprints gets a reference to the given []string and assigns it to the Fingerprints field.
func (o *EntityV3DatadogPipelines) SetFingerprints(v []string) {
	o.Fingerprints = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3DatadogPipelines) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fingerprints != nil {
		toSerialize["fingerprints"] = o.Fingerprints
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3DatadogPipelines) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fingerprints []string `json:"fingerprints,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Fingerprints = all.Fingerprints

	return nil
}
