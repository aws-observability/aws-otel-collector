// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConnectionAttributes Org connection attributes.
type OrgConnectionAttributes struct {
	// List of connection types.
	ConnectionTypes []OrgConnectionTypeEnum `json:"connection_types"`
	// Timestamp when the connection was created.
	CreatedAt time.Time `json:"created_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgConnectionAttributes instantiates a new OrgConnectionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgConnectionAttributes(connectionTypes []OrgConnectionTypeEnum, createdAt time.Time) *OrgConnectionAttributes {
	this := OrgConnectionAttributes{}
	this.ConnectionTypes = connectionTypes
	this.CreatedAt = createdAt
	return &this
}

// NewOrgConnectionAttributesWithDefaults instantiates a new OrgConnectionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgConnectionAttributesWithDefaults() *OrgConnectionAttributes {
	this := OrgConnectionAttributes{}
	return &this
}

// GetConnectionTypes returns the ConnectionTypes field value.
func (o *OrgConnectionAttributes) GetConnectionTypes() []OrgConnectionTypeEnum {
	if o == nil {
		var ret []OrgConnectionTypeEnum
		return ret
	}
	return o.ConnectionTypes
}

// GetConnectionTypesOk returns a tuple with the ConnectionTypes field value
// and a boolean to check if the value has been set.
func (o *OrgConnectionAttributes) GetConnectionTypesOk() (*[]OrgConnectionTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionTypes, true
}

// SetConnectionTypes sets field value.
func (o *OrgConnectionAttributes) SetConnectionTypes(v []OrgConnectionTypeEnum) {
	o.ConnectionTypes = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OrgConnectionAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrgConnectionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OrgConnectionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgConnectionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["connection_types"] = o.ConnectionTypes
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgConnectionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectionTypes *[]OrgConnectionTypeEnum `json:"connection_types"`
		CreatedAt       *time.Time               `json:"created_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConnectionTypes == nil {
		return fmt.Errorf("required field connection_types missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connection_types", "created_at"})
	} else {
		return err
	}
	o.ConnectionTypes = *all.ConnectionTypes
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
