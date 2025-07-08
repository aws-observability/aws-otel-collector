// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableGeoIp Uses a GeoIP database to enrich logs based on an IP field.
type ObservabilityPipelineEnrichmentTableGeoIp struct {
	// Path to the IP field in the log.
	KeyField string `json:"key_field"`
	// Locale used to resolve geographical names.
	Locale string `json:"locale"`
	// Path to the GeoIP database file.
	Path string `json:"path"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineEnrichmentTableGeoIp instantiates a new ObservabilityPipelineEnrichmentTableGeoIp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineEnrichmentTableGeoIp(keyField string, locale string, path string) *ObservabilityPipelineEnrichmentTableGeoIp {
	this := ObservabilityPipelineEnrichmentTableGeoIp{}
	this.KeyField = keyField
	this.Locale = locale
	this.Path = path
	return &this
}

// NewObservabilityPipelineEnrichmentTableGeoIpWithDefaults instantiates a new ObservabilityPipelineEnrichmentTableGeoIp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineEnrichmentTableGeoIpWithDefaults() *ObservabilityPipelineEnrichmentTableGeoIp {
	this := ObservabilityPipelineEnrichmentTableGeoIp{}
	return &this
}

// GetKeyField returns the KeyField field value.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) GetKeyField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.KeyField
}

// GetKeyFieldOk returns a tuple with the KeyField field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) GetKeyFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyField, true
}

// SetKeyField sets field value.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) SetKeyField(v string) {
	o.KeyField = v
}

// GetLocale returns the Locale field value.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) SetLocale(v string) {
	o.Locale = v
}

// GetPath returns the Path field value.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) SetPath(v string) {
	o.Path = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineEnrichmentTableGeoIp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["key_field"] = o.KeyField
	toSerialize["locale"] = o.Locale
	toSerialize["path"] = o.Path

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineEnrichmentTableGeoIp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		KeyField *string `json:"key_field"`
		Locale   *string `json:"locale"`
		Path     *string `json:"path"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.KeyField == nil {
		return fmt.Errorf("required field key_field missing")
	}
	if all.Locale == nil {
		return fmt.Errorf("required field locale missing")
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key_field", "locale", "path"})
	} else {
		return err
	}
	o.KeyField = *all.KeyField
	o.Locale = *all.Locale
	o.Path = *all.Path

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
