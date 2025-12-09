// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSyslogNgSource The `syslog_ng` source listens for logs over TCP or UDP from a `syslog-ng` server using the syslog protocol.
type ObservabilityPipelineSyslogNgSource struct {
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline (e.g., as input to downstream components).
	Id string `json:"id"`
	// Protocol used by the syslog source to receive messages.
	Mode ObservabilityPipelineSyslogSourceMode `json:"mode"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The source type. The value should always be `syslog_ng`.
	Type ObservabilityPipelineSyslogNgSourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSyslogNgSource instantiates a new ObservabilityPipelineSyslogNgSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSyslogNgSource(id string, mode ObservabilityPipelineSyslogSourceMode, typeVar ObservabilityPipelineSyslogNgSourceType) *ObservabilityPipelineSyslogNgSource {
	this := ObservabilityPipelineSyslogNgSource{}
	this.Id = id
	this.Mode = mode
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSyslogNgSourceWithDefaults instantiates a new ObservabilityPipelineSyslogNgSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSyslogNgSourceWithDefaults() *ObservabilityPipelineSyslogNgSource {
	this := ObservabilityPipelineSyslogNgSource{}
	var typeVar ObservabilityPipelineSyslogNgSourceType = OBSERVABILITYPIPELINESYSLOGNGSOURCETYPE_SYSLOG_NG
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSyslogNgSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSyslogNgSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSyslogNgSource) SetId(v string) {
	o.Id = v
}

// GetMode returns the Mode field value.
func (o *ObservabilityPipelineSyslogNgSource) GetMode() ObservabilityPipelineSyslogSourceMode {
	if o == nil {
		var ret ObservabilityPipelineSyslogSourceMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSyslogNgSource) GetModeOk() (*ObservabilityPipelineSyslogSourceMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *ObservabilityPipelineSyslogNgSource) SetMode(v ObservabilityPipelineSyslogSourceMode) {
	o.Mode = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineSyslogNgSource) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSyslogNgSource) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineSyslogNgSource) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineSyslogNgSource) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSyslogNgSource) GetType() ObservabilityPipelineSyslogNgSourceType {
	if o == nil {
		var ret ObservabilityPipelineSyslogNgSourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSyslogNgSource) GetTypeOk() (*ObservabilityPipelineSyslogNgSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSyslogNgSource) SetType(v ObservabilityPipelineSyslogNgSourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSyslogNgSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["mode"] = o.Mode
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSyslogNgSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                                  `json:"id"`
		Mode *ObservabilityPipelineSyslogSourceMode   `json:"mode"`
		Tls  *ObservabilityPipelineTls                `json:"tls,omitempty"`
		Type *ObservabilityPipelineSyslogNgSourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "mode", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
