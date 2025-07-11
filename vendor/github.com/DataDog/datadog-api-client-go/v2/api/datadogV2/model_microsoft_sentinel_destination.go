// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftSentinelDestination The `microsoft_sentinel` destination forwards logs to Microsoft Sentinel.
type MicrosoftSentinelDestination struct {
	// Azure AD client ID used for authentication.
	ClientId string `json:"client_id"`
	// The immutable ID of the Data Collection Rule (DCR).
	DcrImmutableId string `json:"dcr_immutable_id"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The name of the Log Analytics table where logs are sent.
	Table string `json:"table"`
	// Azure AD tenant ID.
	TenantId string `json:"tenant_id"`
	// The destination type. The value should always be `microsoft_sentinel`.
	Type MicrosoftSentinelDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMicrosoftSentinelDestination instantiates a new MicrosoftSentinelDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMicrosoftSentinelDestination(clientId string, dcrImmutableId string, id string, inputs []string, table string, tenantId string, typeVar MicrosoftSentinelDestinationType) *MicrosoftSentinelDestination {
	this := MicrosoftSentinelDestination{}
	this.ClientId = clientId
	this.DcrImmutableId = dcrImmutableId
	this.Id = id
	this.Inputs = inputs
	this.Table = table
	this.TenantId = tenantId
	this.Type = typeVar
	return &this
}

// NewMicrosoftSentinelDestinationWithDefaults instantiates a new MicrosoftSentinelDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMicrosoftSentinelDestinationWithDefaults() *MicrosoftSentinelDestination {
	this := MicrosoftSentinelDestination{}
	var typeVar MicrosoftSentinelDestinationType = MICROSOFTSENTINELDESTINATIONTYPE_MICROSOFT_SENTINEL
	this.Type = typeVar
	return &this
}

// GetClientId returns the ClientId field value.
func (o *MicrosoftSentinelDestination) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftSentinelDestination) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value.
func (o *MicrosoftSentinelDestination) SetClientId(v string) {
	o.ClientId = v
}

// GetDcrImmutableId returns the DcrImmutableId field value.
func (o *MicrosoftSentinelDestination) GetDcrImmutableId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DcrImmutableId
}

// GetDcrImmutableIdOk returns a tuple with the DcrImmutableId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftSentinelDestination) GetDcrImmutableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcrImmutableId, true
}

// SetDcrImmutableId sets field value.
func (o *MicrosoftSentinelDestination) SetDcrImmutableId(v string) {
	o.DcrImmutableId = v
}

// GetId returns the Id field value.
func (o *MicrosoftSentinelDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MicrosoftSentinelDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *MicrosoftSentinelDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *MicrosoftSentinelDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *MicrosoftSentinelDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *MicrosoftSentinelDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetTable returns the Table field value.
func (o *MicrosoftSentinelDestination) GetTable() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *MicrosoftSentinelDestination) GetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value.
func (o *MicrosoftSentinelDestination) SetTable(v string) {
	o.Table = v
}

// GetTenantId returns the TenantId field value.
func (o *MicrosoftSentinelDestination) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *MicrosoftSentinelDestination) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *MicrosoftSentinelDestination) SetTenantId(v string) {
	o.TenantId = v
}

// GetType returns the Type field value.
func (o *MicrosoftSentinelDestination) GetType() MicrosoftSentinelDestinationType {
	if o == nil {
		var ret MicrosoftSentinelDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MicrosoftSentinelDestination) GetTypeOk() (*MicrosoftSentinelDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *MicrosoftSentinelDestination) SetType(v MicrosoftSentinelDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MicrosoftSentinelDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_id"] = o.ClientId
	toSerialize["dcr_immutable_id"] = o.DcrImmutableId
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	toSerialize["table"] = o.Table
	toSerialize["tenant_id"] = o.TenantId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MicrosoftSentinelDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientId       *string                           `json:"client_id"`
		DcrImmutableId *string                           `json:"dcr_immutable_id"`
		Id             *string                           `json:"id"`
		Inputs         *[]string                         `json:"inputs"`
		Table          *string                           `json:"table"`
		TenantId       *string                           `json:"tenant_id"`
		Type           *MicrosoftSentinelDestinationType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientId == nil {
		return fmt.Errorf("required field client_id missing")
	}
	if all.DcrImmutableId == nil {
		return fmt.Errorf("required field dcr_immutable_id missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Table == nil {
		return fmt.Errorf("required field table missing")
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenant_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_id", "dcr_immutable_id", "id", "inputs", "table", "tenant_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ClientId = *all.ClientId
	o.DcrImmutableId = *all.DcrImmutableId
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.Table = *all.Table
	o.TenantId = *all.TenantId
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
